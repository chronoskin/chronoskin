package pack

import (
	"strings"
	"testing"
)

func TestLintCSS(t *testing.T) {
	tokens := minimalPack()["tokens.css"]
	tokens = strings.Replace(tokens, "--radius-control: 0", "--radius-control: 4px", 1)
	tokens = strings.Replace(tokens, "--text-base: 0", "--text-base: 11px", 1)

	good := `.card { color: var(--color-text); border-radius: var(--radius-control); font-family: var(--font-body);
		box-shadow: none; padding: 12px; font-size: var(--text-base); }`
	got, err := LintCSS(good, tokens)
	if err != nil || len(got) != 0 {
		t.Errorf("clean CSS: %v, %v", got, err)
	}

	cases := map[string]string{
		`.a { color: #336699; }`:                                             "use var(--color-",
		`.a { color: #ff00ff; }`:                                             "not in the pack's palette",
		`.a { background: white; }`:                                          "named colour",
		`.a { color: rgb(0 0 0); }`:                                          "raw colour function",
		`.a { border-radius: 4px; }`:                                         "use var(--radius-control)",
		`.a { border-radius: 12px; }`:                                        "not one of the pack's --radius-",
		`.a { box-shadow: 0 2px 8px var(--color-shadow); }`:                  "use a --shadow-* token",
		`.a { color: var(--color-x); }`:                                      "not a pack token",
		`:root { --x: #ff0000 } .a { color: var(--x); }`:                     "not in the pack's palette",
		`:root { --r: 9px } .a { border-radius: var(--r); }`:                 "not one of the pack's --radius-",
		`.a { font-family: Comic Sans MS, var(--font-body); }`:               "use a --font-* token",
		`@keyframes k { from { color: #f00 } }`:                              "not in the pack's palette",
		`.q::before{content:"/*"} .a{color:#123456} .q::after{content:"*/"}`: "not in the pack's palette",
		`.a { box-shadow: 0 2px 8px black; }`:                                "use a --shadow-* token",
		`.a { font-family: Inter, sans-serif; }`:                             "use a --font-* token",
		`.a { font-size: 11px; }`:                                            "use var(--text-base)",
		`.a { font-size: 18px; }`:                                            "not one of the pack's --text-",
		`.a { transition: all 0.2s; }`:                                       "use var(--transition)",
	}
	for css, want := range cases {
		got, err := LintCSS(css, tokens)
		if err != nil {
			t.Errorf("%s: %v", css, err)
			continue
		}
		if !strings.Contains(strings.Join(got, "\n"), want) {
			t.Errorf("%s: want %q, got %v", css, want, got)
		}
	}
}

func TestLintCSSLimits(t *testing.T) {
	tokens := minimalPack()["tokens.css"]
	// Deep nesting must be refused quickly, not parsed quadratically.
	deep := strings.Repeat("@a{", 50000) + strings.Repeat("}", 50000)
	if _, err := LintCSS(deep, tokens); err == nil || !strings.Contains(err.Error(), "nested") {
		t.Errorf("deep nesting: %v", err)
	}
	if _, err := LintCSS(strings.Repeat("a", MaxLintCSSBytes+1), tokens); err == nil {
		t.Error("oversized CSS: no error")
	}
	many := ".a{color:" + strings.Repeat("#abcdef ", 5000) + "}"
	got, err := LintCSS(many, tokens)
	if err != nil || len(got) != maxLintCSSViolations+1 || !strings.Contains(got[len(got)-1], "more violations") {
		t.Errorf("violation cap: %d lines, %v", len(got), err)
	}
}
