package pack

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func minimalPack() map[string]string {
	var tokens, css, html, md strings.Builder
	for _, p := range Parts {
		tokens.WriteString("/* @part " + p.Name + " */\n:root {\n")
		for _, t := range p.Tokens {
			v := "0"
			if p.Name == "palette" {
				v = "#336699"
			}
			tokens.WriteString("  " + t + ": " + v + ";\n")
		}
		tokens.WriteString("}\n")
	}

	css.WriteString(":root { --size-page: 760px; --space-1: 4px; }\n")
	html.WriteString("<!doctype html>\n<html><head><title>Specimen</title>\n" +
		"<link rel=\"stylesheet\" href=\"tokens.css\">\n<link rel=\"stylesheet\" href=\"components.css\">\n</head><body class=\"ds-page\">\n")
	md.WriteString("# Minimal\n\n## Summary\n\nText.\n\n## Layout\n\nInner pages start with the page header.\n\n## Typography and colour roles\n\nText.\n\n## Components\n\n")
	for _, c := range Components {
		css.WriteString("." + c.Class + " { color: var(--color-text); padding: var(--space-1); }\n")
		html.WriteString("<div class=\"" + c.Class + "\" data-component=\"" + c.ID + "\"></div>\n")
		md.WriteString("- `." + c.Class + "`: text.\n")
	}
	for _, v := range Variants {
		css.WriteString("." + v + " { color: var(--color-text); }\n")
		html.WriteString("<span class=\"" + v + "\"></span>\n")
	}
	css.WriteString(".ds-brand__name { color: var(--color-text); }\n")
	css.WriteString(strings.Join(viewRules, "\n") + "\n.ds-view--home { color: var(--color-text); }\n")
	for i, view := range []string{"home", "list", "detail"} {
		class := "ds-view"
		if i == 0 {
			class += " ds-view--home"
		}
		html.WriteString("<section class=\"" + class + "\" id=\"" + view + "\"><a class=\"ds-link\" href=\"#" + view + "\">" + view + "</a></section>\n")
	}
	html.WriteString("<span class=\"ds-brand__name\">chronoskin</span>\n")
	html.WriteString("</body></html>\n")
	md.WriteString("\n## Never\n\n" +
		"- `border-radius <= 0px`: corners are square.\n" +
		"- `box-shadow = none`: nothing casts a shadow.\n" +
		"- `row-gap <= 12px`: rows sit close together.\n" +
		"\n## Extending\n\nText.\n")

	const id = "v1.test-era.forum.s01.p01.t01.u01"
	return map[string]string{
		"style.json": `{"id":"` + id + `","library":"fixture","era":{"slug":"test-era","years":[2002,2007]},` +
			`"archetype":"forum","mode":"pure","markup":"modern","viewport":{"width":1024,"fluid":false},` +
			`"source":"http://localhost:8080/s/` + id + `"}`,
		"tokens.css":     tokens.String(),
		"components.css": css.String(),
		"specimen.html":  html.String(),
		"STYLE.md":       md.String(),
	}
}

func lintFiles(t *testing.T, files map[string]string) []Problem {
	t.Helper()
	dir := t.TempDir()
	for name, content := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return Lint(dir)
}

func TestLintMinimalPack(t *testing.T) {
	for _, p := range lintFiles(t, minimalPack()) {
		t.Error(p)
	}
}

func TestLintViolations(t *testing.T) {
	cases := []struct {
		name string
		file string
		edit func(string) string
		want string // substring of an expected problem
	}{
		{"raw colour in components", "components.css",
			func(s string) string { return s + ".ds-x { color: #ff0000; }\n" }, "raw colour"},
		{"named colour in components", "components.css",
			func(s string) string { return s + ".ds-x { background: white; }\n" }, "named colour"},
		{"raw length in components", "components.css",
			func(s string) string { return s + ".ds-x { border-radius: 8px; }\n" }, "raw length 8px"},
		{"unscoped selector", "components.css",
			func(s string) string { return s + "a { color: var(--color-link); }\n" }, "must be scoped"},
		{"foreign class", "components.css",
			func(s string) string { return s + ".ds-x .card { margin: 0; }\n" }, "must start with ds- or is-"},
		{"unknown token", "components.css",
			func(s string) string { return s + ".ds-x { color: var(--brand); }\n" }, "not a pack token"},
		{"font shorthand", "components.css",
			func(s string) string { return s + ".ds-x { font: inherit; }\n" }, "shorthand"},
		{"import", "components.css",
			func(s string) string { return "@import 'x.css';\n" + s }, "@import"},
		{"missing component class", "components.css",
			func(s string) string { return strings.Replace(s, ".ds-footer {", ".ds-foot {", 1) }, "missing component class .ds-footer"},
		{"missing variant", "components.css",
			func(s string) string { return strings.Replace(s, ".ds-button--danger {", ".ds-button--delete {", 1) }, "missing variant class .ds-button--danger"},
		{"unused variant", "specimen.html",
			func(s string) string { return strings.Replace(s, `class="ds-badge--warning"`, "", 1) }, "nothing uses the variant class ds-badge--warning"},
		{"no inner page guidance", "STYLE.md",
			func(s string) string { return strings.Replace(s, "Inner pages", "Pages", 1) }, "inner pages"},
		{"missing token", "tokens.css",
			func(s string) string { return strings.Replace(s, "  --color-link: #336699;\n", "", 1) }, "--color-link is not defined"},
		{"extra token", "tokens.css",
			func(s string) string {
				return strings.Replace(s, "  --font-body:", "  --font-fancy: 0;\n  --font-body:", 1)
			}, "not a type token"},
		{"raw colour in surface", "tokens.css",
			func(s string) string { return strings.Replace(s, "--fill-bar: 0", "--fill-bar: #cccccc", 1) }, "raw colour in the surface part"},
		{"remote url", "tokens.css",
			func(s string) string {
				return strings.Replace(s, "--fill-page: 0", "--fill-page: url(//example.test/bg.gif)", 1)
			}, "data: URI"},
		{"missing checklist item", "specimen.html",
			func(s string) string { return strings.Replace(s, `data-component="dialog"`, "", 1) }, `data-component="dialog"`},
		{"inline style", "specimen.html",
			func(s string) string { return strings.Replace(s, "</body>", "<p\nstyle=\"margin:0\"></p></body>", 1) }, `"style="`},
		{"event handler", "specimen.html",
			func(s string) string { return strings.Replace(s, "</body>", `<p onclick="x()"></p></body>`, 1) }, "event handler"},
		{"protocol-relative link", "specimen.html",
			func(s string) string {
				return strings.Replace(s, "</body>", `<a class="ds-link" href="//evil.example/">x</a></body>`, 1)
			}, "must be a fragment or a relative path"},
		{"meta refresh", "specimen.html",
			func(s string) string {
				return strings.Replace(s, "<title>", `<meta http-equiv="refresh" content="0;url=x"><title>`, 1)
			}, "http-equiv"},
		{"body without ds-page", "specimen.html",
			func(s string) string { return strings.Replace(s, `<body class="ds-page">`, "<body>", 1) }, "must carry the ds-page class"},
		{"negative raw length", "components.css",
			func(s string) string { return s + ".ds-x { margin-top: -12px; }\n" }, "raw length 12px"},
		{"rule hidden between strings", "components.css",
			func(s string) string {
				return s + `.ds-q::before{content:"/*"} .ds-x{color:#123456} .ds-q::after{content:"*/"}` + "\n"
			}, "raw colour"},
		{"raw colour in keyframes", "components.css",
			func(s string) string { return s + "@keyframes ds-k { from { color: #ff0000; } }\n" }, "raw colour"},
		{"scoped only inside :not", "components.css",
			func(s string) string { return s + "*:not(.ds-nothing) { margin: 0; }\n" }, "must be scoped"},
		{"less common colour name", "components.css",
			func(s string) string { return s + ".ds-x { color: rebeccapurple; }\n" }, "named colour"},
		{"at-rule around tokens", "tokens.css",
			func(s string) string {
				return strings.Replace(s, "/* @part type */\n:root {", "/* @part type */\n@media print { :root {", 1) + "}"
			}, "at-rules are not allowed"},
		{"presence rule with a number", "STYLE.md",
			func(s string) string { return strings.Replace(s, "`box-shadow = none`", "`box-shadow <= 5`", 1) }, "must be"},
		{"undefined class", "specimen.html",
			func(s string) string { return strings.Replace(s, "</body>", `<p class="ds-ghost"></p></body>`, 1) }, "not defined in components.css"},
		{"svg colour", "specimen.html",
			func(s string) string {
				return strings.Replace(s, "</body>", `<svg><rect fill="#ff0000"/></svg></body>`, 1)
			}, "inline SVG may only paint"},
		{"script", "specimen.html",
			func(s string) string { return strings.Replace(s, "</body>", "<script>1</script></body>", 1) }, `"<script"`},
		{"bad never rule", "STYLE.md",
			func(s string) string { return strings.Replace(s, "- `box-shadow = none`:", "- No shadows,", 1) }, "Never rule is not in the form"},
		{"another site name", "specimen.html",
			func(s string) string { return strings.Replace(s, ">chronoskin<", ">Northgate<", 1) }, "the brand must be named"},
		{"unreachable view", "specimen.html",
			func(s string) string { return strings.Replace(s, `href="#detail"`, `href="#"`, 1) }, "is not reachable"},
		{"no view rule", "components.css",
			func(s string) string { return strings.Replace(s, ".ds-view { display: none; }", "", 1) }, "missing the view rule"},
		{"unknown never metric", "STYLE.md",
			func(s string) string { return strings.Replace(s, "`row-gap <=", "`cosiness <=", 1) }, "unknown metric"},
		{"missing section", "STYLE.md",
			func(s string) string { return strings.Replace(s, "## Extending", "## Other", 1) }, "sections must be exactly"},
		{"id and era disagree", "style.json",
			func(s string) string { return strings.Replace(s, `"slug":"test-era"`, `"slug":"other"`, 1) }, "era.slug"},
		{"pure with mixed parts", "style.json",
			func(s string) string { return strings.ReplaceAll(s, "p01", "p02") }, "mode is pure"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			files := minimalPack()
			edited := c.edit(files[c.file])
			if edited == files[c.file] {
				t.Fatal("edit changed nothing")
			}
			files[c.file] = edited
			problems := lintFiles(t, files)
			for _, p := range problems {
				if strings.Contains(p.String(), c.want) {
					return
				}
			}
			t.Errorf("no problem containing %q; got %v", c.want, problems)
		})
	}
}

// The vocabulary lives in vocab.go; the format document must list all of it.
func TestFormatDocListsVocabulary(t *testing.T) {
	b, err := os.ReadFile("../../docs/pack-format.md")
	if err != nil {
		t.Fatal(err)
	}
	doc := string(b)
	var want []string
	for _, p := range Parts {
		want = append(want, p.Tokens...)
	}
	for _, c := range Components {
		want = append(want, "`"+c.ID+"`", "`."+c.Class+"`")
	}
	for _, v := range Variants {
		want = append(want, "`."+v+"`")
	}
	for _, m := range NeverMetrics {
		want = append(want, "`"+m+"`")
	}
	want = append(want, StyleSections...)
	for _, w := range want {
		if !strings.Contains(doc, w) {
			t.Errorf("docs/pack-format.md does not mention %s", w)
		}
	}
}

func TestFixturePacks(t *testing.T) {
	dirs, _ := filepath.Glob("../../packs/*/[0-9][0-9]")
	for _, dir := range dirs {
		for _, p := range Lint(dir) {
			t.Errorf("%s: %s", strings.TrimPrefix(dir, "../../"), p)
		}
	}
}
