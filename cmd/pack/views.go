package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/chronoskin/chronoskin/internal/browser"
	"github.com/chronoskin/chronoskin/internal/pack"
)

//go:embed views.js
var viewsJS string

// listLeftovers makes views print every rule it counts as left over, one to
// a line, so that a script can remove them.
var listLeftovers bool

// views opens each pack's specimen in a browser, at its design width and,
// when it is responsive, at tablet and phone width, and clicks through its
// views.
// -phone checks every pack at phone width, whatever it declares; -closed
// skips the pass with every menu and tooltip held open; -leftovers lists
// each rule that styles nothing and that STYLE.md does not describe.
func views(args []string) (int, error) {
	everyPhone, skipOpen := false, false
	listLeftovers = false
	dirs := args
	for len(dirs) > 0 && strings.HasPrefix(dirs[0], "-") {
		switch dirs[0] {
		case "-phone":
			everyPhone = true
		case "-closed":
			skipOpen = true
		case "-leftovers":
			listLeftovers = true
		}
		dirs = dirs[1:]
	}
	if len(dirs) == 0 {
		return 2, fmt.Errorf("usage: pack views [-phone] [-closed] [-leftovers] <pack-dir>...")
	}
	b, err := browser.Start()
	if err != nil {
		return 2, err
	}
	defer b.Close()
	status := 0
	for _, dir := range dirs {
		problems, err := checkViews(b, dir, everyPhone, skipOpen)
		if err != nil {
			return 2, err
		}
		if problems == 0 {
			fmt.Printf("%s: ok\n", dir)
		} else {
			status = 1
		}
	}
	return status, nil
}

func checkViews(b *browser.Browser, dir string, everyPhone, skipOpen bool) (problems int, err error) {
	raw, err := os.ReadFile(filepath.Join(dir, "style.json"))
	if err != nil {
		return 0, err
	}
	var m pack.Manifest
	if err := json.Unmarshal(raw, &m); err != nil {
		return 0, fmt.Errorf("%s: %w", dir, err)
	}
	specimen, err := filepath.Abs(filepath.Join(dir, "specimen.html"))
	if err != nil {
		return 0, err
	}
	pages := []string{specimen}
	if !skipOpen {
		open, err := heldOpen(dir)
		if err != nil {
			return 0, err
		}
		defer os.Remove(open)
		pages = append(pages, open)
	}
	widths := []int{m.Viewport.Width}
	if m.Viewport.Responsive || everyPhone {
		widths = append(widths, tabletWidth, phoneWidth)
	}
	for _, width := range widths {
		for _, page := range pages {
			url := "file://" + page
			if listLeftovers && page != specimen {
				url += "?list"
			}
			n, err := printViewProblems(b, dir, url, width)
			if err != nil {
				return problems, err
			}
			problems += n
		}
	}
	return problems, nil
}

func printViewProblems(b *browser.Browser, dir, url string, width int) (int, error) {
	res, err := b.Eval(url, width, viewsHeight, viewsJS)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", dir, err)
	}
	var problems []string
	if err := json.Unmarshal([]byte(res), &problems); err != nil {
		return 0, fmt.Errorf("%s: %q", dir, res)
	}
	for _, p := range problems {
		fmt.Printf("%s: at %dpx %s\n", dir, width, p)
	}
	return len(problems), nil
}

var (
	ruleHeadRe        = regexp.MustCompile(`[^{}]+\{`)
	openStateRe       = regexp.MustCompile(`:(hover|focus-within|focus-visible|focus)`)
	menuDetailRe      = regexp.MustCompile(`<details([^>]*ds-menu[^>]*)>`)
	documentedClassRe = regexp.MustCompile(`\.ds-[a-z0-9_-]+`)
)

// heldOpen writes a copy of a pack's specimen in which every menu and
// tooltip shows as if it were pointed at, and returns its path. The states
// that open them are taken out of the selectors that name a menu or a
// tooltip, and a menu made of "<details>" is opened.
func heldOpen(dir string) (string, error) {
	files := map[string]string{}
	for _, name := range pack.Files {
		b, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			return "", err
		}
		files[name] = string(b)
	}
	files["components.css"] = ruleHeadRe.ReplaceAllStringFunc(files["components.css"], func(head string) string {
		if !strings.Contains(head, "ds-menu") && !strings.Contains(head, "ds-tooltip") {
			return head
		}
		return openStateRe.ReplaceAllString(head, "")
	})
	page := menuDetailRe.ReplaceAllString(pack.InlineSpecimen(files), "<details$1 open>")
	documented := strings.Join(documentedClassRe.FindAllString(files["STYLE.md"], -1), " ")
	page = strings.Replace(page, "<html", `<html data-open data-documented="`+documented+`"`, 1)
	f, err := os.CreateTemp("", "chronoskin-open-*.html")
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = f.WriteString(page)
	return f.Name(), err
}
