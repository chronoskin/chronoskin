package pack

import "regexp"

// LintTokens checks a tokens.css on its own, as the linter checks the one
// in a pack. It is what a token variant (a palette, type and surface set
// without a structure) is held to.
func LintTokens(dir string) []Problem {
	l := &linter{dir: dir}
	if _, ok := l.read("tokens.css"); !ok {
		return []Problem{{"tokens.css", "missing"}}
	}
	l.tokens()
	return l.problems
}

var stylesheetRe = regexp.MustCompile(`(?i)<link\b[^>]*\bhref="(tokens\.css|components\.css)"[^>]*>`)

// InlineSpecimen returns a pack's specimen.html with its two stylesheets
// embedded, so that it can be served or rendered as one file.
func InlineSpecimen(files map[string]string) string {
	return stylesheetRe.ReplaceAllStringFunc(files["specimen.html"], func(tag string) string {
		name := stylesheetRe.FindStringSubmatch(tag)[1]
		return "<style>\n" + files[name] + "</style>"
	})
}
