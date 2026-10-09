package library

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand/v2"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/chronoskin/chronoskin/internal/pack"
)

// Every field is optional.
type Request struct {
	Era       string // one slug or several joined by commas; "" or "any" for every era
	Archetype string // one page type or several joined by commas; "" for any
	Mode      string // "pure" (default) or "mix"
	Seed      uint64
	From      string // an existing Style ID to regenerate from, long or short
	// Density is "compact", "normal" or "roomy". Left empty, a regenerated
	// style keeps its own.
	Density string
	// Colours are palette tokens to set by hand, as CSS colours; nil leaves
	// them as they are.
	Colours   map[string]string
	NoColours bool     // drop the hand-set colours a regenerated style has
	Lock      []string // parts of From to keep: structure, palette, type, surface
}

// Among reports whether a request's era list, one slug or several joined by
// commas, includes the era. An empty list and "any" include every era.
func Among(list, era string) bool {
	return list == "" || list == "any" || slices.Contains(strings.Split(list, ","), era)
}

var ErrNotFound = errors.New("not found")

var partNames = []string{"structure", "palette", "type", "surface"}

const (
	// seedStream and the request's seed fix every random pick.
	seedStream = 0x9e3779b97f4a7c15
	// Regenerating tries again when chance gives back the style it
	// started from.
	rerollTries = 17
)

// The same request always gives the same ID: all randomness comes from the
// seed.
func (l *Library) Generate(req Request) (pack.ID, error) {
	for _, lock := range req.Lock {
		if !slices.Contains(partNames, lock) {
			return pack.ID{}, fmt.Errorf("unknown lock %q; use structure, palette, type or surface", lock)
		}
	}
	if req.Mode != "" && req.Mode != "pure" && req.Mode != "mix" {
		return pack.ID{}, fmt.Errorf("unknown mode %q; use pure or mix", req.Mode)
	}
	rng := rand.New(rand.NewPCG(req.Seed, seedStream))
	id, err := l.pickParts(req, rng)
	if err != nil {
		return id, err
	}
	return l.adjust(id, req)
}

func (l *Library) pickParts(req Request, rng *rand.Rand) (pack.ID, error) {
	if req.From == "" {
		return l.pickNew(req, rng)
	}
	// Regenerating must change something whenever anything can change.
	from, _ := l.Parse(req.From)
	var id pack.ID
	var err error
	for range rerollTries {
		id, err = l.reroll(req, rng)
		if err != nil || id != from {
			break
		}
	}
	return id, err
}

func (l *Library) reroll(req Request, rng *rand.Rand) (pack.ID, error) {
	id, ok := l.Parse(req.From)
	if !ok || id.Prefix != l.Prefix {
		return pack.ID{}, fmt.Errorf("%w: style %q", ErrNotFound, req.From)
	}
	if _, err := l.parts(id); err != nil {
		return pack.ID{}, err
	}
	e := l.eras[id.Era]
	locked := func(part string) bool { return slices.Contains(req.Lock, part) }
	if !locked("structure") {
		candidates := l.layouts(e, req.Archetype)
		if locked("palette") {
			candidates = accepting(e, candidates, id.Palette)
		}
		// Layouts are few, so a random pick soon repeats one. Taking the
		// next in order shows every layout before any comes round again.
		// An era's layouts may be different kinds of page; the page type
		// in the ID follows the layout.
		id.Structure = after(candidates, id.Structure)
		id.Archetype = e.structures[id.Structure].Archetype
	}
	if !locked("palette") {
		id.Palette = pick(rng, e.structures[id.Structure].Palettes)
		// Colours set by hand belong to the palette they were set on.
		id.Colours = ""
	}
	if !locked("type") {
		id.Type = pick(rng, every(e.Type))
	}
	if !locked("surface") {
		id.Surface = pick(rng, every(e.Surfaces))
	}
	return id, nil
}

func (l *Library) pickNew(req Request, rng *rand.Rand) (pack.ID, error) {
	var candidates []*era
	for _, ei := range l.Eras {
		e := l.eras[ei.Slug]
		if !Among(req.Era, e.Slug) {
			continue
		}
		if req.Archetype != "" && len(numbers(structureParts(e.Structures), req.Archetype)) == 0 {
			continue
		}
		candidates = append(candidates, e)
	}
	if len(candidates) == 0 {
		return pack.ID{}, fmt.Errorf("%w: no style for era %q and archetype %q", ErrNotFound, req.Era, req.Archetype)
	}
	e := candidates[rng.IntN(len(candidates))]
	id := pack.ID{Prefix: l.Prefix, Era: e.Slug}
	id.Structure = pick(rng, l.layouts(e, req.Archetype))
	id.Archetype = e.structures[id.Structure].Archetype
	if req.Mode == "mix" {
		id.Palette = pick(rng, e.structures[id.Structure].Palettes)
		id.Type = pick(rng, every(e.Type))
		id.Surface = pick(rng, every(e.Surfaces))
	} else {
		id.Palette = id.Structure
		id.Type = id.Structure
		id.Surface = id.Structure
	}
	return id, nil
}

func accepting(e *era, layouts []int, palette int) []int {
	var out []int
	for _, n := range layouts {
		if slices.Contains(e.structures[n].Palettes, palette) {
			out = append(out, n)
		}
	}
	return out
}

func (l *Library) adjust(id pack.ID, req Request) (pack.ID, error) {
	switch req.Density {
	case "":
	case "normal":
		id.Density = ""
	case "compact", "roomy":
		id.Density = req.Density
	default:
		return pack.ID{}, fmt.Errorf("unknown density %q; use compact, normal or roomy", req.Density)
	}
	if req.NoColours {
		id.Colours = ""
	}
	if req.Colours == nil {
		return id, nil
	}
	colours, err := l.encodeOverrides(id, req.Colours)
	if err != nil {
		return pack.ID{}, err
	}
	id.Colours = colours
	return id, nil
}

// A colour equal to the palette's own is not an override.
func (l *Library) encodeOverrides(id pack.ID, colours map[string]string) (string, error) {
	p, err := l.parts(id.Base())
	if err != nil {
		return "", err
	}
	own, err := pack.TokenValues(p.palette.Tokens)
	if err != nil {
		return "", err
	}
	custom := map[string]pack.RGBA{}
	for token, value := range colours {
		c, ok := pack.ParseColour(value)
		if !ok || !pack.IsPaletteToken(token) {
			return "", fmt.Errorf("%s: %q is not a palette token with a colour such as #2b55e0 or rgba(0, 0, 0, 0.5)", token, value)
		}
		if base, ok := pack.ParseColour(own[token]); !ok || base != c {
			custom[token] = c
		}
	}
	return pack.EncodeColours(custom), nil
}

// Archetypes is the page types an era has layouts of, sorted.
func (l *Library) Archetypes(eraSlug string) []string {
	if e := l.eras[eraSlug]; e != nil {
		return e.archetypes()
	}
	return nil
}

// Parse reads a Style ID in either form. The short form names its era by
// the library's code and leaves the page type to the layout.
func (l *Library) Parse(s string) (pack.ID, bool) {
	if id, ok := pack.ParseID(s); ok {
		return id, true
	}
	short, ok := pack.ParseShortID(s)
	if !ok {
		return pack.ID{}, false
	}
	for _, e := range l.Eras {
		if e.Code != short.Code {
			continue
		}
		st := l.eras[e.Slug].structures[short.Structure]
		if st == nil {
			return pack.ID{}, false
		}
		return pack.ID{
			Prefix:    short.Prefix,
			Era:       e.Slug,
			Archetype: st.Archetype,
			Structure: short.Structure,
			Palette:   short.Palette,
			Type:      short.Type,
			Surface:   short.Surface,
			Density:   short.Density,
			Colours:   short.Colours,
		}, true
	}
	return pack.ID{}, false
}

// Short is the short form of a style's ID, or the long form for an ID the
// short form cannot hold.
func (l *Library) Short(id pack.ID) string {
	if e := l.eras[id.Era]; e != nil {
		if s, ok := id.Short(e.Code); ok {
			return s
		}
	}
	return id.String()
}

// after returns the number that follows current in numbers, wrapping to the
// first; current itself when there is nothing else.
func after(numbers []int, current int) int {
	if len(numbers) == 0 {
		return current
	}
	sorted := slices.Clone(numbers)
	slices.Sort(sorted)
	for _, n := range sorted {
		if n > current {
			return n
		}
	}
	return sorted[0]
}

func before(numbers []int, current int) int {
	if len(numbers) == 0 {
		return current
	}
	sorted := slices.Clone(numbers)
	slices.Sort(sorted)
	for i := len(sorted) - 1; i >= 0; i-- {
		if sorted[i] < current {
			return sorted[i]
		}
	}
	return sorted[len(sorted)-1]
}

// layouts is the numbers of an era's layouts of one page type, or of all of
// them when the type is empty or the era has none of that type.
func (l *Library) layouts(e *era, archetype string) []int {
	if archetype != "" {
		if of := numbers(structureParts(e.Structures), archetype); len(of) > 0 {
			return of
		}
	}
	return every(structureParts(e.Structures))
}

// Step returns the style that differs from id in one part, which moves to
// the next number (direction 1) or the previous (-1) and wraps at the ends.
// A layout and a palette that may not be combined are skipped. "density"
// steps through compact, the layout's own spacing and roomy. A page type,
// when given, keeps a step of the layout among the layouts of that type.
func (l *Library) Step(id pack.ID, part string, direction int, archetype ...string) (pack.ID, error) {
	if _, err := l.parts(id); err != nil {
		return pack.ID{}, err
	}
	move := after
	if direction < 0 {
		move = before
	}
	e := l.eras[id.Era]
	switch part {
	case "structure":
		within := ""
		if len(archetype) > 0 {
			within = archetype[0]
		}
		candidates := accepting(e, l.layouts(e, within), id.Palette)
		id.Structure = move(candidates, id.Structure)
		id.Archetype = e.structures[id.Structure].Archetype
	case "palette":
		id.Palette = move(e.structures[id.Structure].Palettes, id.Palette)
		// Colours set by hand belong to the palette they were set on.
		id.Colours = ""
	case "type":
		id.Type = move(every(e.Type), id.Type)
	case "surface":
		id.Surface = move(every(e.Surfaces), id.Surface)
	case "density":
		order := []string{"compact", "", "roomy"}
		id.Density = order[move([]int{0, 1, 2}, slices.Index(order, id.Density))]
	default:
		return pack.ID{}, fmt.Errorf("unknown part %q; use structure, palette, type, surface or density", part)
	}
	return id, nil
}

func (e *era) archetypes() []string {
	var out []string
	for _, s := range e.Structures {
		if !slices.Contains(out, s.Archetype) {
			out = append(out, s.Archetype)
		}
	}
	slices.Sort(out)
	return out
}

func structureParts(s []StructureIndex) []PartIndex {
	out := make([]PartIndex, len(s))
	for i := range s {
		out[i] = s[i].PartIndex
	}
	return out
}

// numbers lists the parts made for a page type, or for any of several joined by commas.
func numbers(parts []PartIndex, archetype string) []int {
	var out []int
	wanted := strings.Split(archetype, ",")
	for _, p := range parts {
		if slices.Contains(wanted, p.Archetype) {
			out = append(out, p.N)
		}
	}
	return out
}

// Palettes, type sets and surface sets belong to the era, not to a page
// type: any layout of the era may wear them.
func every(parts []PartIndex) []int {
	out := make([]int, len(parts))
	for i, p := range parts {
		out[i] = p.N
	}
	return out
}

// pick panics on an empty list: validate guarantees every era has parts of
// each kind and every structure a palette.
func pick(rng *rand.Rand, from []int) int {
	return from[rng.IntN(len(from))]
}

type parts struct {
	era       *era
	structure *structure
	palette   *tokenPart
	typ       *tokenPart
	surface   *tokenPart
}

// parts fails for any ID that is not a style of this library; every other
// method relies on that.
func (l *Library) parts(id pack.ID) (parts, error) {
	missing := fmt.Errorf("%w: style %q", ErrNotFound, id)
	if id.Density != "" && pack.DensityFactor(id.Density) == 1 {
		return parts{}, missing
	}
	if _, ok := pack.DecodeColours(id.Colours); !ok {
		return parts{}, missing
	}
	e := l.eras[id.Era]
	if id.Prefix != l.Prefix || e == nil {
		return parts{}, missing
	}
	p := parts{
		era:       e,
		structure: e.structures[id.Structure],
		palette:   e.palettes[id.Palette],
		typ:       e.types[id.Type],
		surface:   e.surfaces[id.Surface],
	}
	if p.structure == nil || p.palette == nil || p.typ == nil || p.surface == nil ||
		p.structure.Archetype != id.Archetype {
		return parts{}, missing
	}
	// A palette the structure was not cleared for is not a style of this library.
	if !slices.Contains(p.structure.Palettes, id.Palette) {
		return parts{}, missing
	}
	return p, nil
}

type Colour struct {
	Token    string // such as --color-accent
	Value    string // what the style uses, as CSS
	Own      string // what the palette itself says
	Custom   bool   // set by hand
	Editable bool   // a plain colour, which an ID can carry
}

// Palette lists every colour of a style in vocabulary order.
func (l *Library) Palette(id pack.ID) ([]Colour, error) {
	p, err := l.parts(id)
	if err != nil {
		return nil, err
	}
	own, err := pack.TokenValues(p.palette.Tokens)
	if err != nil {
		return nil, err
	}
	custom, _ := pack.DecodeColours(id.Colours)
	var out []Colour
	for _, token := range pack.PaletteTokens() {
		c := Colour{Token: token, Value: own[token], Own: own[token]}
		_, c.Editable = pack.ParseColour(own[token])
		if set, ok := custom[token]; ok {
			c.Value = set.CSS()
			c.Custom = true
		}
		out = append(out, c)
	}
	return out, nil
}

type Style struct {
	ID          pack.ID
	Name        string // of the structure: the layout's own title
	Era         Era
	Archetype   string
	Mode        string
	Viewport    pack.Viewport
	Description string
}

func (l *Library) Style(id pack.ID) (Style, error) {
	p, err := l.parts(id)
	if err != nil {
		return Style{}, err
	}
	return Style{
		ID:          id,
		Name:        p.structure.Title,
		Era:         p.era.Era,
		Archetype:   id.Archetype,
		Mode:        id.Mode(),
		Viewport:    p.structure.Viewport,
		Description: p.structure.sections["Summary"],
	}, nil
}

// Pack assembles the five files of a style. source is the style's permalink.
func (l *Library) Pack(id pack.ID, source string) (map[string]string, error) {
	p, err := l.parts(id)
	if err != nil {
		return nil, err
	}
	custom, _ := pack.DecodeColours(id.Colours)
	manifest, err := json.MarshalIndent(pack.Manifest{
		ID:        id.String(),
		Library:   l.Version,
		Era:       pack.Era{Slug: p.era.Slug, Years: p.era.Years},
		Archetype: id.Archetype,
		Mode:      id.Mode(),
		Markup:    "modern",
		Viewport:  p.structure.Viewport,
		Source:    source,
		Short:     l.Short(id),
		Density:   id.Density,
		Colours:   coloursCSS(custom),
	}, "", "  ")
	if err != nil {
		return nil, err
	}

	density := pack.DensityFactor(id.Density)
	never := []string{p.structure.sections["Never"]}
	for _, part := range []*tokenPart{p.palette, p.typ, p.surface} {
		never = append(never, part.Never...)
	}
	sections := map[string]string{
		"Summary":                     p.structure.sections["Summary"],
		"Layout":                      p.structure.sections["Layout"],
		"Typography and colour roles": roles(id, p),
		"Components":                  p.structure.sections["Components"],
		"Never":                       scaleGapRules(strings.TrimSpace(strings.Join(never, "\n")), density),
		"Extending":                   p.structure.sections["Extending"],
	}
	tokens := []string{withColours(p.palette.Tokens, custom), p.typ.Tokens, p.surface.Tokens}
	return map[string]string{
		"style.json":     string(manifest) + "\n",
		"STYLE.md":       pack.JoinStyle(p.structure.Title, sections),
		"tokens.css":     pack.JoinTokens(tokens),
		"components.css": scaleSpacing(p.structure.components, density),
		"specimen.html":  p.structure.specimen,
	}, nil
}

// roles joins the role notes of the three token parts. In a mix they were
// written for different styles, so the tokens are named as the authority.
func roles(id pack.ID, p parts) string {
	var notes []string
	if !id.Pure() {
		notes = append(notes, "This style combines the parts of different styles of one era. "+
			"Where a note below names a value that differs from `tokens.css`, the tokens are right.")
	}
	for _, part := range []*tokenPart{p.typ, p.palette, p.surface} {
		if part.Roles != "" && !slices.Contains(notes, part.Roles) {
			notes = append(notes, part.Roles)
		}
	}
	if custom, _ := pack.DecodeColours(id.Colours); len(custom) > 0 {
		var names []string
		for _, token := range pack.PaletteTokens() {
			if c, ok := custom[token]; ok {
				names = append(names, "`"+token+"` ("+c.CSS()+")")
			}
		}
		notes = append(notes, "These colours were set by hand and replace the palette's own "+
			"wherever a note above names them: "+strings.Join(names, ", ")+".")
	}
	if id.Density != "" {
		factor := strconv.FormatFloat(pack.DensityFactor(id.Density), 'g', -1, 64)
		notes = append(notes, "Spacing is "+id.Density+": every `--space-*` token of the layout was multiplied by "+factor+
			". Where the Layout section names a spacing in pixels, the tokens are right.")
	}
	return strings.Join(notes, "\n\n")
}

func coloursCSS(custom map[string]pack.RGBA) map[string]string {
	if len(custom) == 0 {
		return nil
	}
	out := map[string]string{}
	for token, c := range custom {
		out[token] = c.CSS()
	}
	return out
}

var (
	spaceDeclRe = regexp.MustCompile(`(--space-[a-z0-9-]+\s*:\s*)(\d+(?:\.\d+)?)px`)
	gapRuleRe   = regexp.MustCompile("(`(?:row-gap|block-gap) (<=|>=) )(\\d+(?:\\.\\d+)?)px`")
)

func withColours(palette string, custom map[string]pack.RGBA) string {
	for token, c := range custom {
		decl := regexp.MustCompile(`(` + regexp.QuoteMeta(token) + `\s*:\s*)[^;]+;`)
		palette = decl.ReplaceAllString(palette, "${1}"+c.CSS()+";")
	}
	return palette
}

// scaleSpacing multiplies every --space-* token of a stylesheet by factor,
// to whole pixels and never down to nothing.
func scaleSpacing(css string, factor float64) string {
	if factor == 1 {
		return css
	}
	return spaceDeclRe.ReplaceAllStringFunc(css, func(decl string) string {
		g := spaceDeclRe.FindStringSubmatch(decl)
		v, _ := strconv.ParseFloat(g[2], 64)
		scaled := math.Round(v * factor)
		if v > 0 && scaled < 1 {
			scaled = 1
		}
		return g[1] + strconv.FormatFloat(scaled, 'f', -1, 64) + "px"
	})
}

// scaleGapRules moves the limits of the Never rules about gaps with the
// spacing, so that a roomy style is not told its own gaps are too wide.
func scaleGapRules(never string, factor float64) string {
	if factor == 1 {
		return never
	}
	return gapRuleRe.ReplaceAllStringFunc(never, func(rule string) string {
		g := gapRuleRe.FindStringSubmatch(rule)
		v, _ := strconv.ParseFloat(g[3], 64)
		// Only a limit the new spacing would break needs to move.
		if (g[2] == "<=") != (factor > 1) {
			return rule
		}
		scaled := math.Ceil(v * factor)
		if g[2] == ">=" {
			scaled = math.Floor(v * factor)
		}
		return g[1] + strconv.FormatFloat(scaled, 'f', -1, 64) + "px`"
	})
}
