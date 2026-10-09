// Package browser drives headless Chromium over the DevTools Protocol. It
// talks to the browser through --remote-debugging-pipe, so it needs no
// WebSocket client: messages are JSON separated by a zero byte.
package browser

import (
	"bufio"
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

//go:embed extract.js
var extractJS string

const (
	callTimeout     = 60 * time.Second
	evalLoadTimeout = 20 * time.Second
	// A reply may be larger than the buffer: a screenshot arrives as one
	// message.
	readBufferBytes = 1 << 20
	messageBacklog  = 64
)

// Where a browser is looked for when $CHROME is not set: the macOS
// application bundles first, then these names on the PATH.
var (
	chromeAppPaths = []string{
		"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
		"/Applications/Chromium.app/Contents/MacOS/Chromium",
	}
	chromeCommands = []string{"chromium", "chromium-browser", "google-chrome"}
)

// chromeArgs come before the flags of $CHROME_FLAGS.
var chromeArgs = []string{
	"--headless=new",
	"--remote-debugging-pipe",
	"--hide-scrollbars",
	"--disable-gpu",
	"--no-first-run",
	"--mute-audio",
}

type message struct {
	ID        int             `json:"id,omitempty"`
	Method    string          `json:"method,omitempty"`
	Params    json.RawMessage `json:"params,omitempty"`
	SessionID string          `json:"sessionId,omitempty"`
	Result    json.RawMessage `json:"result,omitempty"`
	Error     *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

type obj = map[string]any

// A Browser is not safe for concurrent use.
type Browser struct {
	cmd     *exec.Cmd
	w       *os.File
	r       *os.File
	done    chan struct{}
	in      chan message
	events  []message
	nextID  int
	profile string
}

// ChromePath returns the browser binary: $CHROME, or the first known install.
func ChromePath() (string, error) {
	if p := os.Getenv("CHROME"); p != "" {
		return p, nil
	}
	for _, p := range chromeAppPaths {
		if _, err := os.Stat(p); err == nil {
			return p, nil
		}
	}
	for _, name := range chromeCommands {
		if p, err := exec.LookPath(name); err == nil {
			return p, nil
		}
	}
	return "", errors.New("no Chromium found; set CHROME to its path")
}

// Start launches a browser with a throwaway profile. Close it when done.
func Start() (*Browser, error) {
	path, err := ChromePath()
	if err != nil {
		return nil, err
	}
	profile, err := os.MkdirTemp("", "chronoskin-chrome-")
	if err != nil {
		return nil, err
	}
	// Chromium reads commands on fd 3 and writes replies on fd 4.
	cmdR, cmdW, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	outR, outW, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	args := append([]string{"--user-data-dir=" + profile}, chromeArgs...)
	// Extra flags for the environment, such as --no-sandbox in a container
	// that is itself the sandbox.
	args = append(args, strings.Fields(os.Getenv("CHROME_FLAGS"))...)
	cmd := exec.Command(path, args...)
	cmd.ExtraFiles = []*os.File{cmdR, outW}
	if err := cmd.Start(); err != nil {
		for _, f := range []*os.File{cmdR, cmdW, outR, outW} {
			f.Close()
		}
		os.RemoveAll(profile)
		return nil, err
	}
	cmdR.Close()
	outW.Close()

	b := &Browser{
		cmd:     cmd,
		w:       cmdW,
		r:       outR,
		done:    make(chan struct{}),
		in:      make(chan message, messageBacklog),
		profile: profile,
	}
	go b.readMessages()
	return b, nil
}

func (b *Browser) readMessages() {
	defer close(b.in)
	r := bufio.NewReaderSize(b.r, readBufferBytes)
	for {
		raw, err := r.ReadBytes(0)
		if err != nil {
			return
		}
		var m message
		if json.Unmarshal(raw[:len(raw)-1], &m) != nil {
			continue
		}
		select {
		case b.in <- m:
		case <-b.done:
			return
		}
	}
}

func (b *Browser) Close() {
	close(b.done)
	b.w.Close()
	b.cmd.Process.Kill()
	b.cmd.Wait()
	b.r.Close()
	os.RemoveAll(b.profile)
}

// call sends one command and waits for its reply, which is decoded into out
// unless out is nil. Events that arrive meanwhile are kept in b.events.
func (b *Browser) call(session, method string, params any, out any) error {
	b.nextID++
	id := b.nextID
	p, err := json.Marshal(params)
	if err != nil {
		return err
	}
	raw, _ := json.Marshal(message{ID: id, Method: method, Params: p, SessionID: session})
	if _, err := b.w.Write(append(raw, 0)); err != nil {
		return err
	}
	timeout := time.After(callTimeout)
	for {
		select {
		case m, ok := <-b.in:
			if !ok {
				return errors.New("browser exited")
			}
			if m.ID != id {
				b.events = append(b.events, m)
				continue
			}
			if m.Error != nil {
				return fmt.Errorf("%s: %s", method, m.Error.Message)
			}
			if out != nil {
				return json.Unmarshal(m.Result, out)
			}
			return nil
		case <-timeout:
			return fmt.Errorf("%s: no reply", method)
		}
	}
}

// waitEvent reports whether the event arrived within d.
func (b *Browser) waitEvent(session, method string, d time.Duration) bool {
	for _, m := range b.events {
		if m.Method == method && m.SessionID == session {
			return true
		}
	}
	timeout := time.After(d)
	for {
		select {
		case m, ok := <-b.in:
			if !ok {
				return false
			}
			b.events = append(b.events, m)
			if m.Method == method && m.SessionID == session {
				return true
			}
		case <-timeout:
			return false
		}
	}
}

// openTab loads url in a fresh tab and waits for its load event, at most
// loadTimeout: a page that never fires it (a hung sub-resource) is still
// usable. With network set, the tab's requests are recorded in b.events.
func (b *Browser) openTab(url string, width, height int, network bool, loadTimeout time.Duration) (session string, closeTab func(), err error) {
	var target struct {
		TargetID string `json:"targetId"`
	}
	if err := b.call("", "Target.createTarget", obj{"url": "about:blank"}, &target); err != nil {
		return "", nil, err
	}
	closeTab = func() {
		b.call("", "Target.closeTarget", obj{"targetId": target.TargetID}, nil)
	}
	session, err = b.navigate(target.TargetID, url, width, height, network)
	if err != nil {
		closeTab()
		return "", nil, err
	}
	b.waitEvent(session, "Page.loadEventFired", loadTimeout)
	return session, closeTab, nil
}

func (b *Browser) navigate(targetID, url string, width, height int, network bool) (session string, err error) {
	var attached struct {
		SessionID string `json:"sessionId"`
	}
	if err := b.call("", "Target.attachToTarget", obj{"targetId": targetID, "flatten": true}, &attached); err != nil {
		return "", err
	}
	s := attached.SessionID
	b.events = nil
	if err := b.call(s, "Page.enable", obj{}, nil); err != nil {
		return "", err
	}
	if network {
		if err := b.call(s, "Network.enable", obj{}, nil); err != nil {
			return "", err
		}
	}
	viewport := obj{"width": width, "height": height, "deviceScaleFactor": 1, "mobile": false}
	if err := b.call(s, "Emulation.setDeviceMetricsOverride", viewport, nil); err != nil {
		return "", err
	}
	var nav struct {
		ErrorText string `json:"errorText"`
	}
	if err := b.call(s, "Page.navigate", obj{"url": url}, &nav); err != nil {
		return "", err
	}
	if nav.ErrorText != "" {
		return "", errors.New(nav.ErrorText)
	}
	return s, nil
}

type Capture struct {
	PNG      []byte
	Features json.RawMessage // output of extract.js
	// Stylesheets the page asked for, and how many of them failed to load
	// or came back with an error status.
	Stylesheets       int
	StylesheetsFailed int
}

type CaptureOptions struct {
	Width, Height int
	MaxHeight     int           // screenshot height cap
	LoadTimeout   time.Duration // give up waiting for the load event after this
	Settle        time.Duration // extra wait after load
	Scale         float64       // pixels per CSS pixel in the screenshot; 0 means 1
}

func (b *Browser) Capture(url string, o CaptureOptions) (*Capture, error) {
	s, closeTab, err := b.openTab(url, o.Width, o.Height, true, o.LoadTimeout)
	if err != nil {
		return nil, err
	}
	defer closeTab()
	time.Sleep(o.Settle)

	var eval struct {
		Result struct {
			Value string `json:"value"`
		} `json:"result"`
		ExceptionDetails *struct {
			Text string `json:"text"`
		} `json:"exceptionDetails"`
	}
	if err := b.call(s, "Runtime.evaluate", obj{"expression": extractJS, "returnByValue": true}, &eval); err != nil {
		return nil, err
	}
	if eval.ExceptionDetails != nil {
		return nil, fmt.Errorf("extract.js: %s", eval.ExceptionDetails.Text)
	}
	var dims struct {
		DocumentHeight float64 `json:"documentHeight"`
	}
	if err := json.Unmarshal([]byte(eval.Result.Value), &dims); err != nil {
		return nil, fmt.Errorf("extract.js returned %q", eval.Result.Value)
	}
	height := max(float64(o.Height), min(dims.DocumentHeight, float64(o.MaxHeight)))

	var shot struct {
		Data string `json:"data"`
	}
	if err := b.call(s, "Page.captureScreenshot", obj{
		"format":                "png",
		"captureBeyondViewport": true,
		"clip":                  obj{"x": 0, "y": 0, "width": o.Width, "height": height, "scale": max(o.Scale, 1)},
	}, &shot); err != nil {
		return nil, err
	}
	png, err := base64.StdEncoding.DecodeString(shot.Data)
	if err != nil {
		return nil, err
	}
	c := &Capture{PNG: png, Features: json.RawMessage(eval.Result.Value)}
	c.Stylesheets, c.StylesheetsFailed = stylesheetStats(b.events, s)
	return c, nil
}

// stylesheetStats counts the stylesheet requests of a session and those
// that failed, from the Network events seen so far.
func stylesheetStats(events []message, session string) (total, failed int) {
	seen, bad := map[string]bool{}, map[string]bool{}
	for _, m := range events {
		if m.SessionID != session {
			continue
		}
		var p struct {
			RequestID string `json:"requestId"`
			Type      string `json:"type"`
			Canceled  bool   `json:"canceled"`
			Response  struct {
				Status int `json:"status"`
			} `json:"response"`
		}
		switch m.Method {
		case "Network.responseReceived":
			if json.Unmarshal(m.Params, &p) == nil && p.Type == "Stylesheet" {
				seen[p.RequestID] = true
				if p.Response.Status >= 400 {
					bad[p.RequestID] = true
				}
			}
		case "Network.loadingFailed":
			if json.Unmarshal(m.Params, &p) == nil && p.Type == "Stylesheet" {
				seen[p.RequestID] = true
				if !p.Canceled {
					bad[p.RequestID] = true
				}
			}
		}
	}
	return len(seen), len(bad)
}

// Eval loads url in a fresh tab of the given size, runs a script in the
// page and returns the string it evaluates to. The script may be a promise.
func (b *Browser) Eval(url string, width, height int, script string) (string, error) {
	s, closeTab, err := b.openTab(url, width, height, false, evalLoadTimeout)
	if err != nil {
		return "", err
	}
	defer closeTab()

	var eval struct {
		Result struct {
			Value json.RawMessage `json:"value"`
		} `json:"result"`
		ExceptionDetails *struct {
			Text      string `json:"text"`
			Exception struct {
				Description string `json:"description"`
			} `json:"exception"`
		} `json:"exceptionDetails"`
	}
	params := obj{"expression": script, "returnByValue": true, "awaitPromise": true}
	if err := b.call(s, "Runtime.evaluate", params, &eval); err != nil {
		return "", err
	}
	if eval.ExceptionDetails != nil {
		return "", fmt.Errorf("%s %s", eval.ExceptionDetails.Text, eval.ExceptionDetails.Exception.Description)
	}
	var text string
	if err := json.Unmarshal(eval.Result.Value, &text); err != nil {
		return "", fmt.Errorf("the script gave %s, not a string", eval.Result.Value)
	}
	return text, nil
}
