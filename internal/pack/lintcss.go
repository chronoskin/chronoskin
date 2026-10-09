package pack

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

// Limits of LintCSS, which takes input from anyone through the MCP tool.
const (
	MaxLintCSSBytes = 256 << 10
	// maxLintCSSViolations are listed; the rest are only counted.
	maxLintCSSViolations = 200
)

var (
	// tokenOnlyRe matches what may remain of a value once its var()
	// references are removed, if the value is built from tokens alone:
	// calc(), numbers without a unit, operators and commas.
	tokenOnlyRe = regexp.MustCompile(`^(calc|[\s,()*/+\-0-9.%])*$`)
	varCallRe   = regexp.MustCompile(`var\(\s*--[A-Za-z0-9_-]+\s*\)`)
)

// cssKeywords are values that need no token on any property.
var cssKeywords = []string{"inherit", "initial", "unset", "none", "0", "normal"}

var packNamespaces = []string{
	"--color-", "--font-", "--text-", "--line-", "--weight-",
	"--radius-", "--shadow-", "--fill-", "--border-",
}

// LintCSS checks CSS written by an agent against a pack's tokens.css and
// returns its violations: colours, fonts, sizes, radii and shadows that do
// not come from the pack. An empty result means the CSS adheres.
func LintCSS(css, tokensCSS string) ([]string, error) {
	if len(css) > MaxLintCSSBytes {
		return nil, fmt.Errorf("the CSS is larger than %d KB; check it in parts", MaxLintCSSBytes>>10)
	}
	tokens, err := TokenValues(tokensCSS)
	if err != nil {
		return nil, fmt.Errorf("tokens.css: %w", err)
	}
	rules, err := parseCSS(css)
	if err != nil {
		return nil, err
	}
	l := &cssLinter{tokens: tokens}
	for _, r := range rules {
		if strings.HasPrefix(r.selector, "@") {
			continue
		}
		for _, d := range r.decls {
			// Colours are checked everywhere, custom properties included:
			// a raw colour behind a variable is still a raw colour.
			l.colours(r, d)
			if !strings.HasPrefix(d.prop, "--") {
				l.property(r, d)
			}
		}
	}
	if l.total > len(l.violations) {
		l.violations = append(l.violations, fmt.Sprintf("and %d more violations not listed", l.total-len(l.violations)))
	}
	return l.violations, nil
}

type cssLinter struct {
	tokens     map[string]string // the pack's custom properties and their values
	violations []string
	total      int // violations found, listed or not
}

func (l *cssLinter) add(r cssRule, d cssDecl, format string, args ...any) {
	l.total++
	if len(l.violations) >= maxLintCSSViolations {
		return
	}
	where := fmt.Sprintf("%s { %s: %s }: ", clip(r.selector), d.prop, clip(d.value))
	l.violations = append(l.violations, where+fmt.Sprintf(format, args...))
}

// tokenFor is the token of a family that holds a raw value, or "".
func (l *cssLinter) tokenFor(prefix, value string) string {
	var names []string
	for name, v := range l.tokens {
		if strings.HasPrefix(name, prefix) && strings.EqualFold(v, value) {
			names = append(names, name)
		}
	}
	if len(names) == 0 {
		return ""
	}
	return slices.Min(names)
}

// fromTokens reports whether the whole value comes from the pack: every
// var() a pack token, and nothing else beside them.
func (l *cssLinter) fromTokens(d cssDecl) bool {
	refs := varRe.FindAllStringSubmatch(d.value, -1)
	if len(refs) == 0 {
		return false
	}
	for _, m := range refs {
		if _, ok := l.tokens[m[1]]; !ok {
			return false
		}
	}
	rest := varCallRe.ReplaceAllString(strings.ToLower(d.value), "")
	return tokenOnlyRe.MatchString(rest)
}

func (l *cssLinter) colours(r cssRule, d cssDecl) {
	for _, hex := range hexRe.FindAllString(d.value, -1) {
		if name := l.tokenFor("--color-", hex); name != "" {
			l.add(r, d, "use var(%s) instead of %s", name, hex)
		} else {
			l.add(r, d, "colour %s is not in the pack's palette", hex)
		}
	}
	if colourFnRe.MatchString(d.value) {
		l.add(r, d, "raw colour function; use a --color-* token")
	}
	if !strings.HasPrefix(d.prop, "--") && !isPaintProperty(d.prop) {
		return
	}
	// Custom property names inside var() are not colour keywords.
	plain := varRe.ReplaceAllString(d.value, "var(")
	for _, w := range wordRe.FindAllString(plain, -1) {
		if slices.Contains(namedColours, strings.ToLower(w)) {
			l.add(r, d, "named colour %q; use a --color-* token", w)
		}
	}
}

// property checks the properties whose value has to be a token.
func (l *cssLinter) property(r cssRule, d cssDecl) {
	value := strings.ToLower(d.value)
	free := slices.Contains(cssKeywords, value)
	switch d.prop {
	case "font":
		l.add(r, d, "use font-family, font-size and font-weight with tokens, not the shorthand")
	case "font-family":
		if !free && !l.fromTokens(d) {
			l.add(r, d, "use a --font-* token, and only that")
		}
	case "font-size":
		if !free && !strings.HasSuffix(value, "%") && !l.fromTokens(d) {
			l.suggest(r, d, "--text-", "size is not one of the pack's --text-* sizes")
		}
	case "border-radius":
		if !free && value != "50%" && !l.fromTokens(d) {
			l.suggest(r, d, "--radius-", "radius is not one of the pack's --radius-* values")
		}
	case "box-shadow", "text-shadow":
		if !free && !l.fromTokens(d) {
			l.add(r, d, "use a --shadow-* token or none")
		}
	case "transition":
		if !free && !l.fromTokens(d) {
			l.add(r, d, "use var(--transition)")
		}
	default:
		for _, m := range varRe.FindAllStringSubmatch(d.value, -1) {
			if _, ok := l.tokens[m[1]]; !ok && isPackNamespace(m[1]) {
				l.add(r, d, "%s is not a pack token", m[1])
			}
		}
	}
}

func (l *cssLinter) suggest(r cssRule, d cssDecl, prefix, otherwise string) {
	if name := l.tokenFor(prefix, d.value); name != "" {
		l.add(r, d, "use var(%s)", name)
		return
	}
	l.add(r, d, "%s", otherwise)
}

// isPackNamespace tells a mistyped pack token from the author's own variable.
func isPackNamespace(name string) bool {
	for _, prefix := range packNamespaces {
		if strings.HasPrefix(name, prefix) {
			return true
		}
	}
	return false
}

func isPaintProperty(prop string) bool {
	return prop == "color" || prop == "fill" || prop == "stroke" || prop == "caret-color" ||
		strings.HasPrefix(prop, "background") || strings.HasPrefix(prop, "border") ||
		strings.HasPrefix(prop, "outline") || strings.HasSuffix(prop, "-shadow")
}
