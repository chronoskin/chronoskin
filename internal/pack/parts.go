package pack

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type ID struct {
	Prefix    string // library the style belongs to, such as "v1"
	Era       string
	Archetype string
	Structure int
	Palette   int
	Type      int
	Surface   int
	Density   string // "" for the layout's own spacing, or "compact" or "roomy"
	Colours   string // "" for the palette as it is, or overrides as EncodeColours writes them
}

// Densities are the spacing adjustments an ID may carry: the letter that
// stands for each and the factor applied to the layout's spacing tokens.
var Densities = []struct {
	Name   string
	Letter string
	Factor float64
}{
	{"compact", "c", 0.8},
	{"roomy", "r", 1.25},
}

func densityLetter(name string) string {
	for _, d := range Densities {
		if d.Name == name {
			return d.Letter
		}
	}
	return ""
}

func densityName(letter string) string {
	for _, d := range Densities {
		if d.Letter == letter {
			return d.Name
		}
	}
	return ""
}

// DensityFactor is what a density multiplies spacing by; 1 for none.
func DensityFactor(name string) float64 {
	for _, d := range Densities {
		if d.Name == name {
			return d.Factor
		}
	}
	return 1
}

var idRe = regexp.MustCompile(`^(v\d+)\.([a-z0-9]+(?:-[a-z0-9]+)*)\.([a-z]+)\.s(\d\d)\.p(\d\d)\.t(\d\d)\.u(\d\d)(?:\.d([cr]))?(?:\.c([a-z2-7]+))?$`)

// ParseID reads the long form of a Style ID, the one that names its era:
// v1.flat-design.landing.s02.p03.t01.u04, optionally followed by a density
// (.dc or .dr) and colour overrides (.c followed by letters and digits).
func ParseID(s string) (ID, bool) {
	g := idRe.FindStringSubmatch(s)
	if g == nil {
		return ID{}, false
	}
	if _, ok := DecodeColours(g[9]); !ok {
		return ID{}, false
	}
	return ID{
		Prefix:    g[1],
		Era:       g[2],
		Archetype: g[3],
		Structure: partNumber(g[4], 10),
		Palette:   partNumber(g[5], 10),
		Type:      partNumber(g[6], 10),
		Surface:   partNumber(g[7], 10),
		Density:   densityName(g[8]),
		Colours:   g[9],
	}, true
}

func partNumber(digits string, base int) int {
	n, _ := strconv.ParseInt(digits, base, 0)
	return int(n)
}

func (id ID) String() string {
	return fmt.Sprintf("%s.%s.%s.s%02d.p%02d.t%02d.u%02d",
		id.Prefix, id.Era, id.Archetype, id.Structure, id.Palette, id.Type, id.Surface) + id.adjustments(".")
}

// adjustments is the tail of an ID that carries density and colours.
func (id ID) adjustments(sep string) string {
	var out string
	if l := densityLetter(id.Density); l != "" {
		out += sep + "d" + l
	}
	if id.Colours != "" {
		out += sep + "c" + id.Colours
	}
	return out
}

// Base is the ID without adjustments: the four parts alone.
func (id ID) Base() ID {
	id.Density = ""
	id.Colours = ""
	return id
}

// ShortID is the short form of a Style ID: v1-fl-2314, the library prefix,
// a two-character era code and one base-36 digit for each of the four
// parts, optionally followed by -dc or -dr and -c with colour overrides.
// The era code belongs to a library, which turns a ShortID into an ID and
// back.
type ShortID struct {
	Prefix, Code                      string
	Structure, Palette, Type, Surface int
	Density, Colours                  string
}

var shortRe = regexp.MustCompile(`^(v\d+)-([a-z0-9]{2})-([1-9a-z])([1-9a-z])([1-9a-z])([1-9a-z])(?:-d([cr]))?(?:-c([a-z2-7]+))?$`)

func ParseShortID(s string) (ShortID, bool) {
	g := shortRe.FindStringSubmatch(s)
	if g == nil {
		return ShortID{}, false
	}
	if _, ok := DecodeColours(g[8]); !ok {
		return ShortID{}, false
	}
	return ShortID{
		Prefix:    g[1],
		Code:      g[2],
		Structure: partNumber(g[3], 36),
		Palette:   partNumber(g[4], 36),
		Type:      partNumber(g[5], 36),
		Surface:   partNumber(g[6], 36),
		Density:   densityName(g[7]),
		Colours:   g[8],
	}, true
}

// Short writes id in its short form, given its era's code. It is false
// when a part number does not fit one base-36 digit.
func (id ID) Short(code string) (string, bool) {
	var digits string
	for _, n := range []int{id.Structure, id.Palette, id.Type, id.Surface} {
		if n < 1 || n > 35 {
			return "", false
		}
		digits += strconv.FormatInt(int64(n), 36)
	}
	return id.Prefix + "-" + code + "-" + digits + id.adjustments("-"), true
}

// Pure reports whether all four parts come from the same style.
func (id ID) Pure() bool {
	return id.Structure == id.Palette && id.Palette == id.Type && id.Type == id.Surface
}

// Mode is "pure" or "mix", as style.json records it.
func (id ID) Mode() string {
	if id.Pure() {
		return "pure"
	}
	return "mix"
}

// SplitTokens cuts tokens.css into its palette, type and surface blocks, in
// that order, without the marker comments.
func SplitTokens(src string) ([]string, error) {
	marks := partRe.FindAllStringSubmatchIndex(src, -1)
	if len(marks) != len(Parts) {
		return nil, fmt.Errorf("want %d @part markers, found %d", len(Parts), len(marks))
	}
	blocks := make([]string, len(Parts))
	for i, part := range Parts {
		if name := src[marks[i][2]:marks[i][3]]; name != part.Name {
			return nil, fmt.Errorf("part %d is %q, want %q", i+1, name, part.Name)
		}
		end := len(src)
		if i+1 < len(marks) {
			end = marks[i+1][0]
		}
		blocks[i] = strings.TrimSpace(src[marks[i][1]:end])
	}
	return blocks, nil
}

func JoinTokens(blocks []string) string {
	var b strings.Builder
	for i, part := range Parts {
		if i > 0 {
			b.WriteString("\n")
		}
		b.WriteString("/* @part " + part.Name + " */\n" + blocks[i] + "\n")
	}
	return b.String()
}

func TokenValues(src string) (map[string]string, error) {
	rules, err := parseCSS(src)
	if err != nil {
		return nil, err
	}
	values := map[string]string{}
	for _, r := range rules {
		for _, d := range r.decls {
			if strings.HasPrefix(d.prop, "--") {
				values[d.prop] = d.value
			}
		}
	}
	return values, nil
}

// SplitStyle cuts STYLE.md into its title and the body of each section.
func SplitStyle(md string) (title string, sections map[string]string, err error) {
	lines := strings.Split(md, "\n")
	t, ok := strings.CutPrefix(lines[0], "# ")
	if !ok {
		return "", nil, fmt.Errorf("STYLE.md does not start with a title")
	}
	sections = map[string]string{}
	cur := ""
	var body []string
	flush := func() {
		if cur != "" {
			sections[cur] = strings.TrimSpace(strings.Join(body, "\n"))
		}
		body = nil
	}
	for _, line := range lines[1:] {
		if h, ok := strings.CutPrefix(line, "## "); ok {
			flush()
			cur = strings.TrimSpace(h)
			continue
		}
		body = append(body, line)
	}
	flush()
	for _, s := range StyleSections {
		if _, ok := sections[s]; !ok {
			return "", nil, fmt.Errorf("STYLE.md has no %q section", s)
		}
	}
	return strings.TrimSpace(t), sections, nil
}

func JoinStyle(title string, sections map[string]string) string {
	var b strings.Builder
	b.WriteString("# " + title + "\n")
	for _, s := range StyleSections {
		b.WriteString("\n## " + s + "\n\n" + sections[s] + "\n")
	}
	return b.String()
}

type NeverRule struct {
	Metric string
	Op     string // <=, >=, = or !=
	Value  string // "none", or a number with an optional px or % unit
	Line   string // the whole bullet
}

// Number returns the rule's value and its unit ("", "px" or "%"). ok is
// false for "none".
func (r NeverRule) Number() (n float64, unit string, ok bool) {
	v := r.Value
	for _, u := range []string{"px", "%"} {
		if rest, ok := strings.CutSuffix(v, u); ok {
			v = rest
			unit = u
		}
	}
	n, err := strconv.ParseFloat(v, 64)
	return n, unit, err == nil
}

// NeverRules returns the valid rules of a Never section body, and the rest
// of its text (introductions, notes) with those bullets removed.
func NeverRules(body string) (rules []NeverRule, rest string) {
	var other []string
	for _, line := range strings.Split(body, "\n") {
		if m := neverRe.FindStringSubmatch(line); m != nil {
			rules = append(rules, NeverRule{m[1], m[2], m[3], line})
		} else {
			other = append(other, line)
		}
	}
	return rules, strings.TrimSpace(strings.Join(other, "\n"))
}

// NeverPart names the part of a style that a Never metric belongs to.
func NeverPart(metric string) string {
	switch metric {
	case "border-radius", "border-width", "box-shadow", "box-shadow-blur", "text-shadow",
		"gradient-fills", "transition", "animation":
		return "surface"
	case "font-size", "font-weight", "font-families", "line-height", "letter-spacing",
		"uppercase-text", "underlined-links":
		return "type"
	case "palette-colours":
		return "palette"
	}
	return "structure"
}
