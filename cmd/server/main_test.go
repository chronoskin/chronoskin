package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/chronoskin/chronoskin/internal/library"
	"github.com/chronoskin/chronoskin/internal/pack"
)

func testServer(t *testing.T) *server {
	t.Helper()
	dirs, _ := filepath.Glob("../../packs/*/[0-9][0-9]")
	if len(dirs) == 0 {
		t.Fatal("no packs found in packs/")
	}
	out := t.TempDir()
	if err := library.Build("test", "v1", out, dirs, library.Options{}); err != nil {
		t.Fatal(err)
	}
	lib, err := library.Load(out)
	if err != nil {
		t.Fatal(err)
	}
	return newServer(lib, "")
}

func get(t *testing.T, s *server, path string) *httptest.ResponseRecorder {
	t.Helper()
	rec := httptest.NewRecorder()
	s.handler().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
	return rec
}

func TestPages(t *testing.T) {
	s := testServer(t)
	for _, path := range []string{"/healthz", "/", "/docs", "/eras/" + s.lib.Eras[0].Slug} {
		if rec := get(t, s, path); rec.Code != http.StatusOK {
			t.Errorf("%s: status %d", path, rec.Code)
		}
	}
	for _, path := range []string{"/nope", "/s/garbage", "/s/v1.no-such-era.forum.s01.p01.t01.u01", "/api/generate?era=no-such-era"} {
		if rec := get(t, s, path); rec.Code != http.StatusNotFound {
			t.Errorf("%s: status %d, want 404", path, rec.Code)
		}
	}
	if rec := get(t, s, "/api/generate?seed=x"); rec.Code != http.StatusBadRequest {
		t.Errorf("bad seed: status %d", rec.Code)
	}
}

func TestGenerateAndPackRoutes(t *testing.T) {
	s := testServer(t)
	rec := get(t, s, "/api/generate?seed=5")
	loc := rec.Header().Get("Location")
	if rec.Code != http.StatusSeeOther || !strings.HasPrefix(loc, "/s/v1-") {
		t.Fatalf("generate: status %d, location %q", rec.Code, loc)
	}
	if again := get(t, s, "/api/generate?seed=5").Header().Get("Location"); again != loc {
		t.Errorf("same seed gave %s then %s", loc, again)
	}
	id := strings.TrimPrefix(loc, "/s/")

	page := get(t, s, loc)
	if page.Code != http.StatusOK || !strings.Contains(page.Body.String(), id+".tar.gz | tar xz") {
		t.Errorf("permalink page: status %d", page.Code)
	}

	var info struct {
		ID    string
		Files map[string]string
	}
	rec = get(t, s, loc+".json")
	if err := json.Unmarshal(rec.Body.Bytes(), &info); err != nil || info.ID != id || len(info.Files) != len(pack.Files) {
		t.Errorf("json: %v, id %q, %d files", err, info.ID, len(info.Files))
	}
	// The browser keeps the response but checks it: a rebuilt library may
	// give other files for the same ID.
	etag := rec.Header().Get("ETag")
	if cc := rec.Header().Get("Cache-Control"); cc != "no-cache, no-transform" || etag == "" {
		t.Errorf("json Cache-Control = %q, ETag = %q", cc, etag)
	}
	again := httptest.NewRequest("GET", loc+".json", nil)
	again.Header.Set("If-None-Match", etag)
	recAgain := httptest.NewRecorder()
	s.handler().ServeHTTP(recAgain, again)
	if recAgain.Code != http.StatusNotModified {
		t.Errorf("revalidation: status %d", recAgain.Code)
	}

	// The long form names the same style, and adjustments travel in the ID.
	long, _ := s.lib.Parse(id)
	if rec := get(t, s, "/s/"+long.String()+".json"); rec.Code != http.StatusOK {
		t.Errorf("long form: status %d", rec.Code)
	}
	adjusted := get(t, s, "/api/generate?from="+id+"&lock=structure,palette,type,surface&density=roomy&--color-accent=%232b55e0&keep=").Header().Get("Location")
	if !strings.HasPrefix(adjusted, loc+"-dr-c") {
		t.Errorf("adjusted style = %q", adjusted)
	}
	if rec := get(t, s, adjusted+".json"); !strings.Contains(rec.Body.String(), "--color-accent: #2b55e0") {
		t.Errorf("colour not applied: status %d", rec.Code)
	}
	if page := get(t, s, adjusted); !regexp.MustCompile(`value="#2b55e0"[^>]*name="--color-accent"`).MatchString(page.Body.String()) {
		t.Error("the editor does not show the colour set by hand")
	}
	if reset := get(t, s, "/api/generate?from="+strings.TrimPrefix(adjusted, "/s/")+"&lock=structure,palette,type,surface&colours=none&density=normal").Header().Get("Location"); reset != loc+"?keep=structure,palette,type,surface" {
		t.Errorf("reset gave %q", reset)
	}

	specimen := get(t, s, loc+"/specimen.html").Body.String()
	if strings.Contains(specimen, "<link") || !strings.Contains(specimen, "--color-page") || !strings.Contains(specimen, ".ds-page") {
		t.Error("specimen is not inlined")
	}

	// The tarball must unpack to .design/ holding a pack that passes lint.
	tarRec := get(t, s, loc+".tar.gz")
	if !bytes.Equal(tarRec.Body.Bytes(), get(t, s, loc+".tar.gz").Body.Bytes()) {
		t.Error("tarball differs between two requests")
	}
	gz, err := gzip.NewReader(tarRec.Body)
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	tr := tar.NewReader(gz)
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		path := filepath.Join(root, h.Name)
		if h.Typeflag == tar.TypeDir {
			os.MkdirAll(path, 0o755)
			continue
		}
		b, _ := io.ReadAll(tr)
		if err := os.WriteFile(path, b, 0o644); err != nil {
			t.Fatal(err)
		}
	}
	for _, p := range pack.Lint(filepath.Join(root, ".design")) {
		t.Errorf("unpacked pack: %s", p)
	}
}

func mcp(t *testing.T, s *server, body string) (int, map[string]any) {
	t.Helper()
	rec := httptest.NewRecorder()
	s.handler().ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/mcp", strings.NewReader(body)))
	var out map[string]any
	json.Unmarshal(rec.Body.Bytes(), &out)
	return rec.Code, out
}

func toolText(t *testing.T, resp map[string]any) (string, bool) {
	t.Helper()
	result, _ := resp["result"].(map[string]any)
	blocks, _ := result["content"].([]any)
	if len(blocks) == 0 {
		t.Fatalf("no content in %v", resp)
	}
	var all []string
	for _, b := range blocks {
		all = append(all, b.(map[string]any)["text"].(string))
	}
	isError, _ := result["isError"].(bool)
	return strings.Join(all, "\n"), isError
}

func TestMCP(t *testing.T) {
	s := testServer(t)

	_, resp := mcp(t, s, `{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2025-03-26","capabilities":{}}}`)
	if v := resp["result"].(map[string]any)["protocolVersion"]; v != "2025-03-26" {
		t.Errorf("protocolVersion = %v", v)
	}
	if code, _ := mcp(t, s, `{"jsonrpc":"2.0","method":"notifications/initialized"}`); code != http.StatusAccepted {
		t.Errorf("notification: status %d, want 202", code)
	}
	_, resp = mcp(t, s, `{"jsonrpc":"2.0","id":2,"method":"tools/list"}`)
	if n := len(resp["result"].(map[string]any)["tools"].([]any)); n != 4 {
		t.Errorf("%d tools, want 4", n)
	}
	_, resp = mcp(t, s, `{"jsonrpc":"2.0","id":3,"method":"nope"}`)
	if resp["error"] == nil {
		t.Error("unknown method: no error")
	}
	if rec := get(t, s, "/mcp"); rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("GET /mcp: status %d", rec.Code)
	}

	_, resp = mcp(t, s, `{"jsonrpc":"2.0","id":4,"method":"tools/call","params":{"name":"list_eras"}}`)
	if text, _ := toolText(t, resp); !strings.Contains(text, s.lib.Eras[0].Slug) {
		t.Error("list_eras does not name the first era")
	}

	_, resp = mcp(t, s, `{"jsonrpc":"2.0","id":5,"method":"tools/call","params":{"name":"generate_style","arguments":{"seed":11,"mode":"mix"}}}`)
	text, _ := toolText(t, resp)
	var style struct{ ID string }
	if err := json.Unmarshal([]byte(text), &style); err != nil || style.ID == "" {
		t.Fatalf("generate_style: %v in %s", err, text)
	}

	_, resp = mcp(t, s, `{"jsonrpc":"2.0","id":6,"method":"tools/call","params":{"name":"get_style","arguments":{"id":"`+style.ID+`"}}}`)
	text, _ = toolText(t, resp)
	for _, name := range pack.Files {
		if !strings.Contains(text, "FILE: .design/"+name+"\n") {
			t.Errorf("get_style lacks %s", name)
		}
	}

	call := func(css string) (string, bool) {
		args, _ := json.Marshal(map[string]string{"id": style.ID, "css": css})
		_, resp := mcp(t, s, `{"jsonrpc":"2.0","id":7,"method":"tools/call","params":{"name":"lint_css","arguments":`+string(args)+`}}`)
		return toolText(t, resp)
	}
	if text, _ := call(".a { color: var(--color-text); }"); !strings.HasPrefix(text, "ok") {
		t.Errorf("clean css: %s", text)
	}
	if text, _ := call(".a { color: #123456; border-radius: 37px; }"); !strings.HasPrefix(text, "2 violations") {
		t.Errorf("dirty css: %s", text)
	}

	_, resp = mcp(t, s, `{"jsonrpc":"2.0","id":8,"method":"tools/call","params":{"name":"get_style","arguments":{"id":"nope"}}}`)
	if _, isError := toolText(t, resp); !isError {
		t.Error("get_style with a bad id: isError not set")
	}
}

func TestMCPStdio(t *testing.T) {
	s := testServer(t)
	in := strings.NewReader(`{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}
{"jsonrpc":"2.0","method":"notifications/initialized"}
{"jsonrpc":"2.0","id":2,"method":"ping"}
`)
	var out bytes.Buffer
	if err := s.serveStdio(in, &out); err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(lines) != 2 || !strings.Contains(lines[0], `"serverInfo"`) || !strings.Contains(lines[1], `"id":2`) {
		t.Errorf("stdio output: %q", lines)
	}
}

func TestRateLimit(t *testing.T) {
	l := newLimiter(3, 0)
	for i := range 3 {
		if !l.allow("a") {
			t.Fatalf("request %d refused", i)
		}
	}
	if l.allow("a") {
		t.Error("fourth request allowed")
	}
	if !l.allow("b") {
		t.Error("another client refused")
	}
}

func TestEraName(t *testing.T) {
	for in, want := range map[string]string{
		"Web portal, 2000 to 2004":                            "Web portal",
		"Neubrutalism landing page (2021 to 2026)":            "Neubrutalism landing page",
		"Skeuomorphism landing (stitched cloth, glossy keys)": "Skeuomorphism landing",
		"Material paper docs, 2016":                           "Material paper docs",
		"Bulletin board forum":                                "Bulletin board forum",
	} {
		if got := eraName(in); got != want {
			t.Errorf("eraName(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestSecurityHeadersAndOrigin(t *testing.T) {
	s := testServer(t)
	home := get(t, s, "/")
	if home.Header().Get("X-Content-Type-Options") != "nosniff" || !strings.Contains(home.Header().Get("Content-Security-Policy"), "frame-ancestors 'none'") {
		t.Errorf("home headers: %v", home.Header())
	}
	id := strings.TrimPrefix(get(t, s, "/api/generate?seed=2").Header().Get("Location"), "/s/")
	specimen := get(t, s, "/s/"+id+"/specimen.html")
	csp := specimen.Header().Get("Content-Security-Policy")
	for _, want := range []string{"sandbox allow-scripts", "default-src 'none'", "script-src 'sha256-", "frame-ancestors 'self'"} {
		if !strings.Contains(csp, want) {
			t.Errorf("specimen CSP %q lacks %q", csp, want)
		}
	}
	// The second request is served from the cache and must be identical.
	if again := get(t, s, "/s/"+id+"/specimen.html"); again.Body.String() != specimen.Body.String() {
		t.Error("cached specimen differs")
	}

	post := func(origin string) int {
		req := httptest.NewRequest(http.MethodPost, "/mcp", strings.NewReader(`{"jsonrpc":"2.0","id":1,"method":"ping"}`))
		if origin != "" {
			req.Header.Set("Origin", origin)
		}
		rec := httptest.NewRecorder()
		s.handler().ServeHTTP(rec, req)
		return rec.Code
	}
	if code := post(""); code != http.StatusOK {
		t.Errorf("no Origin: %d", code)
	}
	if code := post("http://example.com"); code != http.StatusOK {
		t.Errorf("same origin (httptest host is example.com): %d", code)
	}
	if code := post("https://evil.example"); code != http.StatusForbidden {
		t.Errorf("foreign Origin: %d, want 403", code)
	}
}

func TestGenerateNotEra(t *testing.T) {
	s := testServer(t)
	not := s.lib.Eras[0].Slug
	for seed := range 40 {
		rec := httptest.NewRecorder()
		s.handler().ServeHTTP(rec, httptest.NewRequest("GET", "/api/generate?mode=mix&not="+not+"&seed="+strconv.Itoa(seed), nil))
		if loc := rec.Header().Get("Location"); rec.Code != http.StatusSeeOther || strings.Contains(loc, "-"+s.lib.Eras[0].Code+"-") {
			t.Fatalf("seed %d: status %d, location %q", seed, rec.Code, loc)
		}
	}
}

func TestGenerateDoesNotRepeat(t *testing.T) {
	s := testServer(t)
	s.limiter = newLimiter(1000, 1000) // this test presses Generate faster than a visitor could
	era := s.lib.Eras[0]
	for _, query := range []string{"mode=mix", "mode=mix&era=" + era.Slug} {
		var cookie *http.Cookie
		var styles, layouts []string
		for range 40 {
			req := httptest.NewRequest("GET", "/api/generate?"+query, nil)
			if cookie != nil {
				req.AddCookie(cookie)
			}
			rec := httptest.NewRecorder()
			s.handler().ServeHTTP(rec, req)
			id := strings.TrimPrefix(rec.Header().Get("Location"), "/s/")
			if rec.Code != http.StatusSeeOther || len(rec.Result().Cookies()) != 1 {
				t.Fatalf("%s: status %d, cookies %v", query, rec.Code, rec.Result().Cookies())
			}
			cookie = rec.Result().Cookies()[0]
			if n := len(styles); slices.Contains(styles[max(0, n-recentStyles):], id) {
				t.Fatalf("%s: %s came back within %d picks", query, id, recentStyles)
			}
			// Within one era there are fewer layouts than the window is long.
			window := min(recentLayouts, len(era.Structures)-1)
			if n := len(layouts); slices.Contains(layouts[max(0, n-window):], layoutOf(id)) {
				t.Fatalf("%s: layout of %s came back within %d picks", query, id, window)
			}
			styles, layouts = append(styles, id), append(layouts, layoutOf(id))
		}
	}
}

func TestClientIP(t *testing.T) {
	s := testServer(t)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.RemoteAddr = "10.0.0.1:1234"
	req.Header.Set("X-Forwarded-For", "203.0.113.9, 198.51.100.7")
	if got := s.clientIP(req); got != "10.0.0.1" {
		t.Errorf("without -trust-proxy: %q", got)
	}
	s.trustProxy = true
	if got := s.clientIP(req); got != "198.51.100.7" {
		t.Errorf("with -trust-proxy: %q, want the proxy-appended address", got)
	}
	// A proxy that adds its own header line: the client's line must not win.
	req.Header["X-Forwarded-For"] = []string{"1.2.3.4", "198.51.100.8"}
	if got := s.clientIP(req); got != "198.51.100.8" {
		t.Errorf("two header lines: %q", got)
	}
	req.Header["X-Forwarded-For"] = []string{"not-an-address"}
	if got := s.clientIP(req); got != "10.0.0.1" {
		t.Errorf("garbage header: %q, want the connection's address", got)
	}
	s.trustProxy = false
	req.RemoteAddr = "[2001:db8:1:2:aaaa:bbbb:cccc:dddd]:443"
	if got := s.clientIP(req); got != "2001:db8:1:2::" {
		t.Errorf("IPv6: %q, want the /64", got)
	}
}

// A person in a browser gets a page for an error; a program gets the plain
// message.
func TestErrorPages(t *testing.T) {
	s := testServer(t)
	browse := func(path string) *httptest.ResponseRecorder {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.Header.Set("Accept", "text/html,application/xhtml+xml")
		s.handler().ServeHTTP(rec, req)
		return rec
	}
	for _, path := range []string{"/nope", "/s/garbage", "/api/generate?era=no-such-era"} {
		rec := browse(path)
		if rec.Code != http.StatusNotFound || !strings.Contains(rec.Body.String(), "Nothing here") || !strings.Contains(rec.Body.String(), `content="noindex"`) {
			t.Errorf("%s in a browser: status %d, or not the error page", path, rec.Code)
		}
		if plain := get(t, s, path); plain.Code != http.StatusNotFound || strings.Contains(plain.Body.String(), "<html") {
			t.Errorf("%s for a program: status %d, or a page", path, plain.Code)
		}
	}
	// A file is never answered with a page, whatever the client accepts.
	if rec := browse("/s/garbage.json"); strings.Contains(rec.Body.String(), "<html") {
		t.Error("a .json address got an error page")
	}

	s.limiter = newLimiter(1, 0)
	browse("/")
	if rec := browse("/"); rec.Code != http.StatusTooManyRequests || !strings.Contains(rec.Body.String(), "Slow down") || rec.Header().Get("Retry-After") == "" {
		t.Errorf("rate limit in a browser: status %d", rec.Code)
	}
}
