package pack

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Problem struct {
	File string
	Msg  string
}

func (p Problem) String() string { return p.File + ": " + p.Msg }

// The messages spell the view and rule counts out, and docs/pack-format.md
// states them.
const (
	minViews      = 3
	maxViews      = 6
	minNeverRules = 3 // of a pure style; a mix may inherit fewer
	firstYear     = 1990
	lastYear      = 2100
)

var (
	hexRe       = regexp.MustCompile(`#[0-9a-fA-F]{3,8}\b`)
	colourFnRe  = regexp.MustCompile(`(?i)\b(rgba?|hsla?|hwb|lab|lch|oklab|oklch|color)\(`)
	lengthRe    = regexp.MustCompile(`(?i)(?:^|[^\w.#-])-?(\d*\.?\d+)(px|rem|em|pt)\b`)
	styleAttrRe = regexp.MustCompile(`(?i)[\s"'/]style\s*=`)
	eventAttrRe = regexp.MustCompile(`(?i)[\s"'/]on[a-z]+\s*=`)
	// A URL attribute may hold a fragment or a relative path, nothing with a
	// scheme or a host.
	urlAttrRe    = regexp.MustCompile(`(?i)\b(href|src|action|formaction|poster|data|xlink:href|srcset|ping|background)\s*=\s*["']?\s*(//|[a-z][a-z0-9+.-]*:)`)
	functionalRe = regexp.MustCompile(`:(not|has|is|where)\([^()]*\)`)
	varRe        = regexp.MustCompile(`var\(\s*(--[A-Za-z0-9_-]+)`)
	wordRe       = regexp.MustCompile(`[A-Za-z]+`)
	classSelRe   = regexp.MustCompile(`\.(-?[A-Za-z_][\w-]*)`)
	urlRe        = regexp.MustCompile(`(?i)url\(\s*["']?([^"')]{0,5})`)
	partRe       = regexp.MustCompile(`/\*\s*@part\s+([a-z]+)\s*\*/`)
	viewRe       = regexp.MustCompile(`<section class="ds-view( ds-view--home)?" id="([a-z][a-z0-9-]*)"`)
	neverRe      = regexp.MustCompile("^- `([a-z-]+) (<=|>=|=|!=) (none|\\d+(?:\\.\\d+)?(?:px|%)?)`: \\S.*$")
	classAttrRe  = regexp.MustCompile(`\bclass="([^"]*)"`)
	linkRe       = regexp.MustCompile(`(?i)<link\b[^>]*>`)
	hrefRe       = regexp.MustCompile(`\bhref="([^"]*)"`)
	paintAttrRe  = regexp.MustCompile(`\b(fill|stroke|stop-color|flood-color|color)="([^"]*)"`)
	bodyClassRe  = regexp.MustCompile(`(?i)<body\b[^>]*\bclass="[^"]*\bds-page\b`)
)

var bannedMarkup = []string{
	"<style", "<script", "<img", "<iframe", "<video", "<audio", "<object", "<embed",
	"<font", "<center", "<marquee", "<frame", "<base", "http-equiv", "<image", "<foreignobject",
}

// paintKeywords: an inline SVG takes its colour from the surrounding text.
var paintKeywords = []string{"none", "currentcolor", "transparent", "inherit"}

var namedColours = strings.Fields(`aliceblue antiquewhite aqua aquamarine azure beige bisque black blanchedalmond blue
	blueviolet brown burlywood cadetblue chartreuse chocolate coral cornflowerblue cornsilk crimson cyan darkblue darkcyan
	darkgoldenrod darkgray darkgreen darkgrey darkkhaki darkmagenta darkolivegreen darkorange darkorchid darkred darksalmon
	darkseagreen darkslateblue darkslategray darkslategrey darkturquoise darkviolet deeppink deepskyblue dimgray dimgrey
	dodgerblue firebrick floralwhite forestgreen fuchsia gainsboro ghostwhite gold goldenrod gray green greenyellow grey
	honeydew hotpink indianred indigo ivory khaki lavender lavenderblush lawngreen lemonchiffon lightblue lightcoral lightcyan
	lightgoldenrodyellow lightgray lightgreen lightgrey lightpink lightsalmon lightseagreen lightskyblue lightslategray
	lightslategrey lightsteelblue lightyellow lime limegreen linen magenta maroon mediumaquamarine mediumblue mediumorchid
	mediumpurple mediumseagreen mediumslateblue mediumspringgreen mediumturquoise mediumvioletred midnightblue mintcream
	mistyrose moccasin navajowhite navy oldlace olive olivedrab orange orangered orchid palegoldenrod palegreen paleturquoise
	palevioletred papayawhip peachpuff peru pink plum powderblue purple rebeccapurple red rosybrown royalblue saddlebrown
	salmon sandybrown seagreen seashell sienna silver skyblue slateblue slategray slategrey snow springgreen steelblue tan
	teal thistle tomato turquoise violet wheat white whitesmoke yellow yellowgreen
	canvas canvastext linktext visitedtext activetext buttonface buttontext buttonborder field fieldtext highlight
	highlighttext selecteditem selecteditemtext mark marktext graytext accentcolor accentcolortext`)

// Lint returns every violation found; none means the pack is valid.
func Lint(dir string) []Problem {
	l := &linter{dir: dir}
	l.files()
	l.manifest()
	l.tokens()
	l.components()
	l.specimen()
	l.styleMD()
	return l.problems
}

type linter struct {
	dir      string
	problems []Problem
	mode     string          // from style.json
	layout   []string        // custom properties defined by components.css
	classes  map[string]bool // classes defined by components.css
}

func (l *linter) add(file, format string, args ...any) {
	l.problems = append(l.problems, Problem{file, fmt.Sprintf(format, args...)})
}

// read returns a pack file. A missing file is reported once, by files().
func (l *linter) read(name string) (string, bool) {
	b, err := os.ReadFile(filepath.Join(l.dir, name))
	if err != nil {
		return "", false
	}
	return string(b), true
}

func (l *linter) files() {
	entries, err := os.ReadDir(l.dir)
	if err != nil {
		l.add(".", "%v", err)
		return
	}
	var have []string
	for _, e := range entries {
		if strings.HasPrefix(e.Name(), ".") {
			continue
		}
		have = append(have, e.Name())
		if !slices.Contains(Files, e.Name()) {
			l.add(e.Name(), "not a pack file")
		}
	}
	for _, f := range Files {
		if !slices.Contains(have, f) {
			l.add(f, "missing")
		}
	}
}

func (l *linter) manifest() {
	const file = "style.json"
	src, ok := l.read(file)
	if !ok {
		return
	}
	var m Manifest
	dec := json.NewDecoder(strings.NewReader(src))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&m); err != nil {
		l.add(file, "%v", err)
		return
	}
	l.mode = m.Mode
	if id, ok := ParseID(m.ID); ok {
		l.manifestID(file, m, id)
	} else {
		l.add(file, "id %q does not match v<N>.<era-slug>.<archetype>.sNN.pNN.tNN.uNN", m.ID)
	}
	if !slices.Contains(Archetypes, m.Archetype) {
		l.add(file, "unknown archetype %q", m.Archetype)
	}
	if m.Mode != "pure" && m.Mode != "mix" {
		l.add(file, "mode must be pure or mix")
	}
	if m.Markup != "modern" {
		l.add(file, `markup must be "modern"`)
	}
	if m.Library == "" {
		l.add(file, "library is empty")
	}
	if y := m.Era.Years; y[0] < firstYear || y[1] < y[0] || y[1] > lastYear {
		l.add(file, "era.years %v is not a valid range", y)
	}
	if m.Viewport.Width <= 0 {
		l.add(file, "viewport.width must be positive")
	}
}

func (l *linter) manifestID(file string, m Manifest, id ID) {
	if id.Era != m.Era.Slug {
		l.add(file, "era.slug %q differs from the id", m.Era.Slug)
	}
	if id.Archetype != m.Archetype {
		l.add(file, "archetype %q differs from the id", m.Archetype)
	}
	if m.Mode == "pure" && !id.Pure() {
		l.add(file, "mode is pure but the id combines parts of different styles")
	}
	endsWithLong := strings.HasSuffix(m.Source, "/s/"+m.ID)
	endsWithShort := m.Short != "" && strings.HasSuffix(m.Source, "/s/"+m.Short)
	if !endsWithLong && !endsWithShort {
		l.add(file, "source must end with /s/ and the id, long or short")
	}
	if m.Short != "" && !isShortOf(m.Short, id) {
		l.add(file, "short %q does not name the same style as id", m.Short)
	}
	if m.Density != id.Density {
		l.add(file, "density %q differs from the id", m.Density)
	}
	custom := map[string]RGBA{}
	for token, value := range m.Colours {
		c, ok := ParseColour(value)
		if !ok || !IsPaletteToken(token) {
			l.add(file, "colours: %s: %q is not a palette token with a colour", token, value)
		}
		custom[token] = c
	}
	if EncodeColours(custom) != id.Colours {
		l.add(file, "colours differ from the id")
	}
}

// isShortOf leaves the era out: its code belongs to a library.
func isShortOf(text string, id ID) bool {
	short, ok := ParseShortID(text)
	return ok &&
		short.Prefix == id.Prefix &&
		short.Structure == id.Structure &&
		short.Palette == id.Palette &&
		short.Type == id.Type &&
		short.Surface == id.Surface &&
		short.Density == id.Density &&
		short.Colours == id.Colours
}

func (l *linter) tokens() {
	const file = "tokens.css"
	src, ok := l.read(file)
	if !ok {
		return
	}
	l.network(file, src)
	marks := partRe.FindAllStringSubmatchIndex(src, -1)
	if len(marks) != len(Parts) {
		l.add(file, "want %d /* @part <name> */ markers, found %d", len(Parts), len(marks))
		return
	}
	for i, part := range Parts {
		if name := src[marks[i][2]:marks[i][3]]; name != part.Name {
			l.add(file, "part %d is %q, want %q", i+1, name, part.Name)
			return
		}
		end := len(src)
		if i+1 < len(marks) {
			end = marks[i+1][0]
		}
		l.tokenPart(file, part, stripCSSComments(src[marks[i][1]:end]))
	}
}

// tokenPart checks one part: a single :root block that defines exactly the
// part's tokens, with raw colours only in the palette.
func (l *linter) tokenPart(file string, part Part, segment string) {
	if strings.Contains(segment, "@") {
		l.add(file, "%s part: at-rules are not allowed in tokens.css", part.Name)
		return
	}
	rules, err := parseCSS(segment)
	if err != nil {
		l.add(file, "%s part: %v", part.Name, err)
		return
	}
	if len(rules) != 1 || rules[0].selector != ":root" {
		l.add(file, "%s part must be exactly one :root block", part.Name)
		return
	}
	seen := map[string]bool{}
	for _, d := range rules[0].decls {
		if !slices.Contains(part.Tokens, d.prop) {
			l.add(file, "%s is not a %s token", d.prop, part.Name)
			continue
		}
		if seen[d.prop] {
			l.add(file, "%s is defined twice", d.prop)
		}
		seen[d.prop] = true
		if part.Name == "palette" {
			continue
		}
		l.vars(file, d, nil)
		if hexRe.MatchString(d.value) || colourFnRe.MatchString(d.value) {
			l.add(file, "%s: raw colour in the %s part; use a --color-* token", d.prop, part.Name)
		}
		if part.Name == "surface" {
			l.named(file, d, "white", "black")
		}
	}
	for _, t := range part.Tokens {
		if !seen[t] {
			l.add(file, "%s is not defined", t)
		}
	}
}

// viewRules are the rules every components.css carries for the specimen's
// views, compared with white space collapsed.
var viewRules = []string{
	".ds-view { display: none; }",
	".ds-view:target, .ds-page:not(:has(.ds-view:target)) .ds-view--home { display: block; }",
	".ds-view:target { scroll-margin-top: 100vh; }",
}

func (l *linter) components() {
	const file = "components.css"
	src, ok := l.read(file)
	if !ok {
		return
	}
	l.network(file, src)
	rules, err := parseCSS(src)
	if err != nil {
		l.add(file, "%v", err)
		return
	}
	l.classes = map[string]bool{}
	l.layoutTokens(file, rules)
	for _, r := range rules {
		switch {
		case r.selector == ":root":
			for _, d := range r.decls {
				l.colours(file, d)
			}
		case strings.HasPrefix(r.selector, "@"):
			l.add(file, "%s is not allowed", strings.Fields(r.selector)[0])
		default:
			l.componentRule(file, r)
		}
	}
	flat := strings.Join(strings.Fields(src), " ")
	for _, rule := range viewRules {
		if !strings.Contains(flat, rule) {
			l.add(file, "missing the view rule: %s", rule)
		}
	}
	for _, c := range Components {
		if !l.classes[c.Class] {
			l.add(file, "missing component class .%s", c.Class)
		}
	}
	for _, v := range Variants {
		if !l.classes[v] {
			l.add(file, "missing variant class .%s", v)
		}
	}
}

func (l *linter) layoutTokens(file string, rules []cssRule) {
	for _, r := range rules {
		if r.selector != ":root" {
			continue
		}
		for _, d := range r.decls {
			if !strings.HasPrefix(d.prop, "--size-") && !strings.HasPrefix(d.prop, "--space-") {
				l.add(file, ":root may only define --size-* and --space-* layout tokens, not %s", d.prop)
			}
			l.layout = append(l.layout, d.prop)
		}
	}
	if !slices.Contains(l.layout, "--size-page") {
		l.add(file, ":root must define --size-page")
	}
}

func (l *linter) componentRule(file string, r cssRule) {
	if !r.keyframe {
		for rest := r.selector; rest != ""; {
			var one string
			one, rest = cutTop(rest, ",")
			// A .ds- class inside :not() or :has() does not scope the rule.
			if !strings.Contains(functionalRe.ReplaceAllString(one, ""), ".ds-") {
				l.add(file, "selector %q must be scoped by a .ds- class", strings.TrimSpace(one))
			}
		}
	}
	for _, m := range classSelRe.FindAllStringSubmatch(r.selector, -1) {
		if !strings.HasPrefix(m[1], "ds-") && !strings.HasPrefix(m[1], "is-") {
			l.add(file, "class .%s must start with ds- or is-", m[1])
		}
		l.classes[m[1]] = true
	}
	for _, d := range r.decls {
		l.componentDecl(file, r, d)
	}
}

func (l *linter) componentDecl(file string, r cssRule, d cssDecl) {
	where := fmt.Sprintf("%s { %s }", clip(r.selector), d.prop)
	switch {
	case strings.HasPrefix(d.prop, "--"):
		l.add(file, "%s: custom properties belong in :root", where)
		return
	case d.prop == "font":
		l.add(file, "%s: use font-family, font-size and font-weight with tokens, not the shorthand", where)
	case d.prop == "font-family" && !strings.Contains(d.value, "var(") && d.value != "inherit":
		l.add(file, "%s: font-family must be a --font-* token", where)
	}
	if strings.Contains(strings.ToLower(d.value), "url(") {
		l.add(file, "%s: url() is not allowed here; patterns belong in a --fill-* token", where)
	}
	l.colours(file, d)
	if isPaintProperty(d.prop) {
		l.named(file, d)
	}
	l.vars(file, d, l.layout)
	for _, m := range lengthRe.FindAllStringSubmatch(d.value, -1) {
		if n, _ := strconv.ParseFloat(m[1], 64); n != 0 {
			l.add(file, "%s: raw length %s%s; use a token", where, m[1], m[2])
		}
	}
}

func (l *linter) specimen() {
	const file = "specimen.html"
	src, ok := l.read(file)
	if !ok {
		return
	}
	l.network(file, src)
	lower := strings.ToLower(src)
	if !strings.HasPrefix(strings.TrimSpace(lower), "<!doctype html>") {
		l.add(file, "must start with <!doctype html>")
	}
	for _, bad := range bannedMarkup {
		if strings.Contains(lower, bad) {
			l.add(file, "%q is not allowed", bad)
		}
	}
	if styleAttrRe.MatchString(src) {
		l.add(file, `"style=" is not allowed`)
	}
	if eventAttrRe.MatchString(src) {
		l.add(file, "event handler attributes are not allowed")
	}
	if m := urlAttrRe.FindStringSubmatch(src); m != nil {
		l.add(file, "%s must be a fragment or a relative path, not %q", strings.ToLower(m[1]), m[2])
	}
	if !bodyClassRe.MatchString(src) {
		l.add(file, "<body> must carry the ds-page class")
	}
	var hrefs []string
	for _, tag := range linkRe.FindAllString(src, -1) {
		if m := hrefRe.FindStringSubmatch(tag); m != nil {
			hrefs = append(hrefs, m[1])
		}
	}
	if !slices.Equal(hrefs, []string{"tokens.css", "components.css"}) {
		l.add(file, `must have exactly two <link> elements, href="tokens.css" then href="components.css"`)
	}
	for _, c := range Components {
		if !strings.Contains(src, `data-component="`+c.ID+`"`) {
			l.add(file, `checklist: no element with data-component="%s"`, c.ID)
		}
	}
	l.views(file, src)
	// Every specimen is the same site, so that a library reads as one mark
	// in many skins.
	if !strings.Contains(src, `class="ds-brand__name">chronoskin<`) {
		l.add(file, `the brand must be named: <span class="ds-brand__name">chronoskin</span>`)
	}
	for _, m := range paintAttrRe.FindAllStringSubmatch(src, -1) {
		if !slices.Contains(paintKeywords, strings.ToLower(m[2])) {
			l.add(file, `%s="%s": inline SVG may only paint with currentColor or none`, m[1], m[2])
		}
	}
	l.specimenClasses(file, src)
}

// views checks that the first view is the home view and each one is linked to.
func (l *linter) views(file, src string) {
	views := viewRe.FindAllStringSubmatch(src, -1)
	if len(views) < minViews || len(views) > maxViews {
		l.add(file, `must have three to six views, each <section class="ds-view" id="name">; found %d`, len(views))
	}
	for i, v := range views {
		home, name := v[1] != "", v[2]
		if home != (i == 0) {
			l.add(file, `the first view, and only the first, carries class="ds-view ds-view--home" (view %q)`, name)
		}
		if !strings.Contains(src, `href="#`+name+`"`) {
			l.add(file, `view %q is not reachable: no link with href="#%s"`, name, name)
		}
	}
}

func (l *linter) specimenClasses(file, src string) {
	used := map[string]bool{}
	var inOrder []string
	for _, m := range classAttrRe.FindAllStringSubmatch(src, -1) {
		for _, c := range strings.Fields(m[1]) {
			if !used[c] {
				used[c] = true
				inOrder = append(inOrder, c)
			}
		}
	}
	for _, v := range Variants {
		if !used[v] {
			l.add(file, "checklist: nothing uses the variant class %s", v)
		}
	}
	// Without a components.css that parses there is nothing to compare with.
	if l.classes == nil {
		return
	}
	for _, c := range inOrder {
		if !l.classes[c] {
			l.add(file, "class %q is not defined in components.css", c)
		}
	}
}

func (l *linter) styleMD() {
	const file = "STYLE.md"
	src, ok := l.read(file)
	if !ok {
		return
	}
	lines := strings.Split(src, "\n")
	if !strings.HasPrefix(lines[0], "# ") {
		l.add(file, "first line must be a level-1 heading with the style name")
	}
	var headings []string
	section := map[string][]string{}
	for _, line := range lines {
		if h, ok := strings.CutPrefix(line, "## "); ok {
			headings = append(headings, strings.TrimSpace(h))
			continue
		}
		if len(headings) > 0 {
			cur := headings[len(headings)-1]
			section[cur] = append(section[cur], line)
		}
	}
	if !slices.Equal(headings, StyleSections) {
		l.add(file, "sections must be exactly, in order: %s", strings.Join(StyleSections, ", "))
		return
	}
	if !strings.Contains(strings.ToLower(strings.Join(section["Layout"], "\n")), "inner page") {
		l.add(file, "Layout section must say how inner pages are laid out (use the words \"inner pages\")")
	}
	components := strings.Join(section["Components"], "\n")
	for _, c := range Components {
		if !strings.Contains(components, "`."+c.Class+"`") {
			l.add(file, "Components section does not describe `.%s`", c.Class)
		}
	}
	rules := l.neverRules(file, section["Never"])
	// A mix inherits its rules from four parts and may end up with fewer.
	if rules < minNeverRules && l.mode != "mix" {
		l.add(file, "Never section needs at least %d valid rules, has %d", minNeverRules, rules)
	}
}

// neverRules reports malformed bullets and returns how many are well formed.
func (l *linter) neverRules(file string, lines []string) int {
	rules := 0
	for _, line := range lines {
		if !strings.HasPrefix(line, "- ") {
			continue
		}
		m := neverRe.FindStringSubmatch(line)
		if m == nil {
			l.add(file, "Never rule is not in the form \"- `metric op value`: explanation\": %s", clip(line))
			continue
		}
		metric, op, value := m[1], m[2], m[3]
		if !slices.Contains(NeverMetrics, metric) {
			l.add(file, "Never rule uses unknown metric %q", metric)
		}
		if slices.Contains(PresenceMetrics, metric) && !((op == "=" || op == "!=") && value == "none") {
			l.add(file, "Never rule on %s must be \"= none\" or \"!= none\"", metric)
		}
		rules++
	}
	return rules
}

// network reports anything that could make the file fetch a resource.
func (l *linter) network(file, src string) {
	s := strings.ReplaceAll(src, "http://www.w3.org/", "")
	for _, bad := range []string{"http:", "https:", "@import", "@font-face"} {
		if strings.Contains(strings.ToLower(s), bad) {
			l.add(file, "%q is not allowed: packs make no network requests and ship no fonts", bad)
		}
	}
	for _, m := range urlRe.FindAllStringSubmatch(src, -1) {
		if !strings.HasPrefix(strings.ToLower(m[1]), "data:") {
			l.add(file, "url() may only hold a data: URI")
		}
	}
}

// colours reports a colour written out instead of taken from a token.
func (l *linter) colours(file string, d cssDecl) {
	if hexRe.MatchString(d.value) || colourFnRe.MatchString(d.value) {
		l.add(file, "%s: raw colour %q; use a token", d.prop, clip(d.value))
	}
}

// named reports colour keywords, except those in allow.
func (l *linter) named(file string, d cssDecl, allow ...string) {
	// Custom property names inside var() are not colour keywords.
	value := varRe.ReplaceAllString(d.value, "var(")
	for _, w := range wordRe.FindAllString(value, -1) {
		w = strings.ToLower(w)
		if slices.Contains(namedColours, w) && !slices.Contains(allow, w) {
			l.add(file, "%s: named colour %q; use a token", d.prop, w)
		}
	}
}

// vars reports var() references to anything but pack tokens and extra.
func (l *linter) vars(file string, d cssDecl, extra []string) {
	for _, m := range varRe.FindAllStringSubmatch(d.value, -1) {
		if !slices.Contains(extra, m[1]) && !isPackToken(m[1]) {
			l.add(file, "%s: var(%s) is not a pack token", d.prop, m[1])
		}
	}
}

func isPackToken(name string) bool {
	for _, p := range Parts {
		if slices.Contains(p.Tokens, name) {
			return true
		}
	}
	return false
}
