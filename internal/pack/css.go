package pack

import (
	"fmt"
	"strings"
)

// cssRule is a style rule or a block-less at-statement such as @import.
// Rules nested in @media and similar are flattened.
type cssRule struct {
	selector string
	decls    []cssDecl
	keyframe bool // a step of @keyframes: "from", "to" or a percentage
}

type cssDecl struct {
	prop, value string
}

// stripCSSComments removes comments. A comment opener inside a quoted
// string is text, not a comment: otherwise a pair of strings could hide the
// rules between them from the linter while a browser still applies them.
func stripCSSComments(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	var quote byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case quote != 0:
			b.WriteByte(c)
			if c == '\\' && i+1 < len(s) {
				i++
				b.WriteByte(s[i])
			} else if c == quote || c == '\n' {
				quote = 0
			}
		case c == '"' || c == '\'':
			quote = c
			b.WriteByte(c)
		case c == '/' && i+1 < len(s) && s[i+1] == '*':
			end := strings.Index(s[i+2:], "*/")
			if end < 0 {
				return b.String()
			}
			i += end + 3
			b.WriteByte(' ')
		default:
			b.WriteByte(c)
		}
	}
	return b.String()
}

// scanTo returns the index of the first top-level occurrence of any byte in
// stops, skipping strings and parentheses, or -1.
func scanTo(s string, stops string) int {
	depth := 0
	var quote byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case quote != 0:
			if c == '\\' {
				i++
			} else if c == quote {
				quote = 0
			}
		case c == '"' || c == '\'':
			quote = c
		case c == '(':
			depth++
		case c == ')':
			if depth > 0 {
				depth--
			}
		case depth == 0 && strings.IndexByte(stops, c) >= 0:
			return i
		}
	}
	return -1
}

// matchBrace returns the index of the "}" closing a block whose "{" has
// already been consumed, or -1.
func matchBrace(s string) int {
	depth := 1
	off := 0
	for {
		i := scanTo(s[off:], "{}")
		if i < 0 {
			return -1
		}
		if s[off+i] == '{' {
			depth++
		} else if depth--; depth == 0 {
			return off + i
		}
		off += i + 1
	}
}

// cutTop splits s at its first top-level sep; after is "" when there is none.
func cutTop(s, sep string) (before, after string) {
	i := scanTo(s, sep)
	if i < 0 {
		return s, ""
	}
	return s[:i], s[i+1:]
}

const (
	// Real stylesheets nest two or three levels; the limit keeps hostile
	// input from costing more than one pass per level.
	maxNesting     = 16
	maxQuoteLength = 60
)

func parseCSS(src string) ([]cssRule, error) {
	return parseBlocks(stripCSSComments(src), 0)
}

func parseBlocks(s string, depth int) ([]cssRule, error) {
	if depth > maxNesting {
		return nil, fmt.Errorf("at-rules nested more than %d deep", maxNesting)
	}
	var rules []cssRule
	for {
		s = strings.TrimSpace(s)
		if s == "" {
			return rules, nil
		}
		i := scanTo(s, "{;")
		if i < 0 {
			return nil, fmt.Errorf("unexpected text %q", clip(s))
		}
		prelude := strings.Join(strings.Fields(s[:i]), " ")
		if s[i] == ';' {
			rules = append(rules, cssRule{selector: prelude})
			s = s[i+1:]
			continue
		}
		end := matchBrace(s[i+1:])
		if end < 0 {
			return nil, fmt.Errorf("unclosed block after %q", clip(prelude))
		}
		body := s[i+1 : i+1+end]
		s = s[i+1+end+1:]

		if holdsRules(prelude) {
			inner, err := parseBlocks(body, depth+1)
			if err != nil {
				return nil, err
			}
			keyframes := strings.HasPrefix(prelude, "@keyframes") || strings.HasPrefix(prelude, "@-webkit-keyframes")
			for _, r := range inner {
				r.keyframe = r.keyframe || keyframes
				rules = append(rules, r)
			}
			continue
		}
		if scanTo(body, "{") >= 0 {
			return nil, fmt.Errorf("nested rules are not supported (in %q)", clip(prelude))
		}
		decls, err := parseDecls(body, prelude)
		if err != nil {
			return nil, err
		}
		rules = append(rules, cssRule{selector: prelude, decls: decls})
	}
}

// holdsRules reports whether an at-rule's block holds rules (@media,
// @supports, @keyframes) rather than declarations.
func holdsRules(prelude string) bool {
	return strings.HasPrefix(prelude, "@") &&
		!strings.HasPrefix(prelude, "@font-face") &&
		!strings.HasPrefix(prelude, "@page")
}

func parseDecls(body, selector string) ([]cssDecl, error) {
	var decls []cssDecl
	for body != "" {
		var d string
		d, body = cutTop(body, ";")
		d = strings.TrimSpace(d)
		if d == "" {
			continue
		}
		prop, value, ok := strings.Cut(d, ":")
		if !ok {
			return nil, fmt.Errorf("bad declaration %q in %q", clip(d), clip(selector))
		}
		decls = append(decls, cssDecl{strings.TrimSpace(prop), strings.TrimSpace(value)})
	}
	return decls, nil
}

func clip(s string) string {
	if len(s) > maxQuoteLength {
		return s[:maxQuoteLength] + "..."
	}
	return s
}
