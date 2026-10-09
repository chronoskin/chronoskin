package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// The files in assets/ are copies of what the site draws from its own
// templates. This keeps the two the same: the mark's three shapes, and every
// icon's drawing.
func TestAssetsMatchTheSite(t *testing.T) {
	const shapes = `<path d="M12 3a9 9 0 1 1-9 9"/><path d="M12 17a5 5 0 1 1 5-5"/><circle cx="12" cy="12" r="1.4" fill="currentColor" stroke="none"/>`
	// Line breaks and indentation between tags are layout, not drawing.
	between := regexp.MustCompile(`>\s+<`)
	read := func(path string) string {
		t.Helper()
		b, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		return between.ReplaceAllString(string(b), "><")
	}
	logos, _ := filepath.Glob("../../assets/logo/*.svg")
	if len(logos) == 0 {
		t.Fatal("no logo files in assets/logo")
	}
	for _, path := range append(logos, "web/templates/partials/mark.html", "web/static/favicon.svg") {
		if !strings.Contains(read(path), shapes) {
			t.Errorf("%s: does not draw the mark with its three shapes", path)
		}
	}

	// An era's mark is what its first pack's specimen carries.
	mark := regexp.MustCompile(`(?s)<span class="ds-brand__mark"><svg[^>]*>(.*?)</svg></span>`)
	whole := regexp.MustCompile(`(?s)<svg[^>]*>(.*)</svg>`)
	specimens, _ := filepath.Glob("../../packs/*/01/specimen.html")
	if len(specimens) == 0 {
		t.Fatal("no packs found in packs/")
	}
	for _, path := range specimens {
		era := filepath.Base(filepath.Dir(filepath.Dir(path)))
		in := mark.FindStringSubmatch(read(path))
		file := whole.FindStringSubmatch(read("../../assets/logo/eras/" + era + ".svg"))
		if in == nil || file == nil || in[1] != file[1] {
			t.Errorf("assets/logo/eras/%s.svg is not the mark of packs/%s/01", era, era)
		}
	}

	partials := ""
	for _, name := range []string{"icons.html"} {
		partials += read("web/templates/partials/" + name)
	}
	icons, _ := filepath.Glob("../../assets/icons/*.svg")
	if len(icons) == 0 {
		t.Fatal("no icons in assets/icons")
	}
	inner := regexp.MustCompile(`(?s)<svg[^>]*>(.*)</svg>`)
	for _, path := range icons {
		m := inner.FindStringSubmatch(read(path))
		if m == nil || !strings.Contains(partials, m[1]) {
			t.Errorf("%s: the site's templates draw this icon differently, or not at all", path)
		}
	}
}
