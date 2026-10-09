package main

import (
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"

	"github.com/chronoskin/chronoskin/internal/library"
	"github.com/chronoskin/chronoskin/internal/pack"
)

// partView is one row of "This style": a part and the styles one step either way.
type partView struct {
	Key, Label string
	About      string // what the part is and what changing it changes
	Value      string // "02 / 7", or a word for spacing
	Prev, Next string
	CanStep    bool
	Keepable   bool // regenerate can hold it
	Kept       bool // the visitor is holding it
}

type heartView struct {
	ID, Title string
	Width     int // the design width, to draw the saved style at
}

// keepable are the parts Regenerate can hold, in the order a link lists them.
var keepable = []string{"structure", "palette", "type", "surface"}

// pageTitle names a style's page by its layout and its era, without saying
// the era twice when the layout's name already begins with it.
func pageTitle(era, layout string) string {
	if strings.HasPrefix(strings.ToLower(layout), strings.ToLower(era)) {
		return layout
	}
	return layout + ", " + era
}

// pageQuery carries what a visitor set on one style page to the next: the
// parts they keep, the page types they limited layouts to and the eras
// they regenerate among, as "?keep=structure,type&type=shop,blog&eras=...".
func pageQuery(parts, types, eras []string) string {
	var kept, query []string
	for _, p := range keepable {
		if slices.Contains(parts, p) {
			kept = append(kept, p)
		}
	}
	if len(kept) > 0 {
		query = append(query, "keep="+strings.Join(kept, ","))
	}
	types = slices.DeleteFunc(slices.Clone(types), func(t string) bool {
		return !slices.Contains(pack.Archetypes, t)
	})
	if len(types) > 0 {
		query = append(query, "type="+strings.Join(types, ","))
	}
	if len(eras) > 0 {
		query = append(query, "eras="+url.QueryEscape(strings.Join(eras, ",")))
	}
	if len(query) == 0 {
		return ""
	}
	return "?" + strings.Join(query, "&")
}

func (s *server) stylePage(w http.ResponseWriter, r *http.Request, id pack.ID, base string) {
	st, err := s.lib.Style(id)
	if err != nil {
		s.failWith(w, r, err)
		return
	}
	q := r.URL.Query()
	kept := strings.Split(q.Get("keep"), ",")
	// The page types limit which layouts the arrows step through; the eras
	// are those Regenerate picks among: every era when none is chosen.
	types := list(q["type"])
	within := strings.Join(types, ",")
	pool := slices.DeleteFunc(list(q["eras"]), func(slug string) bool { return !s.hasEra(slug) })
	query := pageQuery(kept, types, pool)

	parts, err := s.partViews(id, kept, within, query)
	if err != nil {
		s.failWith(w, r, err)
		return
	}
	palette, err := s.lib.Palette(id)
	if err != nil {
		s.failWith(w, r, err)
		return
	}
	name := eraName(st.Era.Name)
	layout := eraName(st.Name)
	short := s.lib.Short(id)
	s.render(w, r, "style", map[string]any{
		"Title":      name,
		"Nav":        "",
		"Style":      st,
		"EraName":    name,
		"LayoutName": layout,
		"ID":         short,
		"LongID":     id.String(),
		"Years":      years(st.Era.Years),
		"Base":       base,
		"NextSeed":   randomSeed(),
		"Parts":      parts,
		"EraSelect": selectView{
			Name:        "era",
			Label:       "Eras",
			Options:     s.eraOptions(pool),
			Multiple:    true,
			Placeholder: "All eras",
		},
		"Pool":          strings.Join(pool, ","),
		"Several":       len(pool) > 1,
		"PrevEra":       s.neighbourEra(id.Era, -1, pool, within),
		"NextEra":       s.neighbourEra(id.Era, 1, pool, within),
		"About":         st.Era.Description,
		"Colours":       colourGroups(palette),
		"CustomColours": id.Colours != "",
		"Keep":          strings.Join(kept, ","),
		"Within":        within,
		"Query":         query,
		"TypeSelect": selectView{
			Name:        "archetype",
			Label:       "Page types",
			Options:     s.typeOptions(types),
			Multiple:    true,
			Placeholder: "Any page type",
		},
		"Image":   s.previewURL(id.Era, id.Structure),
		"NoIndex": id != id.Base(),
		"Heart":   heartView{ID: short, Title: name, Width: st.Viewport.Width},
		"Description": layout + ", a layout in the " + name + " style of " + years(st.Era.Years) +
			": a style pack to install as .design/ for your coding agent.",
		// Many styles share one layout and differ in colours, type and
		// surface. Search engines are pointed at the layout's own style.
		"PageTitle":     pageTitle(name, layout),
		"CanonicalPath": s.link(s.pureID(id.Era, id.Archetype, id.Structure)),
	})
}

func (s *server) hasEra(slug string) bool {
	return slices.ContainsFunc(s.eraList, func(e eraView) bool { return e.Slug == slug })
}

// position is where a part stands among its kind: "02 / 7".
func position(n, count int) string {
	return twoDigits(n) + " / " + strconv.Itoa(count)
}

func (s *server) partViews(id pack.ID, kept []string, within, query string) ([]partView, error) {
	counts := s.lib.Counts(id.Era)
	colours := position(id.Palette, counts.Palettes)
	if id.Colours != "" {
		colours += " · edited"
	}
	spacing := "Normal"
	if id.Density != "" {
		spacing = capitalise(id.Density)
	}
	parts := []partView{
		{
			Key:      "structure",
			Label:    "Layout",
			About:    "The kind of page and how it is put together: where the navigation sits, how the content is divided, which components it shows. Changing it gives another page of the same era.",
			Value:    position(id.Structure, counts.Structures),
			Keepable: true,
		},
		{
			Key:      "palette",
			Label:    "Colours",
			About:    "Every colour of the style: the page, the text, the bars, the buttons and the status colours. Changing it repaints the same layout.",
			Value:    colours,
			Keepable: true,
		},
		{
			Key:      "type",
			Label:    "Type",
			About:    "The typefaces and how they are set: sizes, weights, line heights and letter spacing, for headings, text and controls.",
			Value:    position(id.Type, counts.Types),
			Keepable: true,
		},
		{
			Key:      "surface",
			Label:    "Surface",
			About:    "What things are made of: borders, corners, shadows, fills and textures. It is what makes a style flat, glossy, raised or glassy.",
			Value:    position(id.Surface, counts.Surfaces),
			Keepable: true,
		},
		{
			Key:   "density",
			Label: "Spacing",
			About: "How much room things get: compact, normal or roomy. It changes gaps and padding and nothing else.",
			Value: spacing,
		},
	}
	for i := range parts {
		p := &parts[i]
		prev, err := s.lib.Step(id, p.Key, -1, within)
		if err != nil {
			return nil, err
		}
		next, _ := s.lib.Step(id, p.Key, 1, within)
		p.Prev = s.link(prev) + query
		p.Next = s.link(next) + query
		p.CanStep = next != id
		p.Kept = p.Keepable && slices.Contains(kept, p.Key)
	}
	return parts, nil
}

// neighbourEra links to the era before (step -1) or after (step 1) this
// one, round the ends of the library's order: where the arrow keys go. The
// choice of eras and page types travels along.
func (s *server) neighbourEra(era string, step int, pool []string, within string) string {
	here := slices.IndexFunc(s.eraList, func(e eraView) bool { return e.Slug == era })
	n := len(s.eraList)
	to := s.eraList[((here+step)%n+n)%n].Slug
	q := url.Values{"mode": {"mix"}, "era": {to}, "pool": {strings.Join(pool, ",")}}
	if within != "" {
		q.Set("archetype", within)
	}
	return "/api/generate?" + q.Encode()
}

func (s *server) eraOptions(selected []string) []option {
	var options []option
	for _, e := range s.eraList {
		options = append(options, option{
			Value:    e.Slug,
			Label:    e.Name,
			Note:     e.Years,
			Selected: slices.Contains(selected, e.Slug),
		})
	}
	return options
}

func (s *server) typeOptions(selected []string) []option {
	var options []option
	for _, a := range s.archetypes {
		options = append(options, option{
			Value:    a,
			Label:    capitalise(a),
			Selected: slices.Contains(selected, a),
		})
	}
	return options
}

type colourView struct {
	Token, Label string
	Value, Own   string // as CSS: what the style uses, and the palette's own
	Hex          string // Value without alpha, for the colour picker
	Custom       bool
	Editable     bool
}

type colourGroup struct {
	Name    string
	Colours []colourView
}

// colourGroupNames says which group of the editor a palette token goes in,
// by the start of its name after "--color-"; the first match wins.
var colourGroupNames = []struct{ prefix, group string }{
	{"page", "Page and surfaces"}, {"canvas", "Page and surfaces"}, {"surface", "Page and surfaces"},
	{"text", "Text"}, {"heading", "Text"},
	{"link", "Links"},
	{"bar", "Bars"}, {"inverse", "Bars"},
	{"button", "Buttons"},
	{"input", "Fields"},
	{"accent", "Accents and fills"}, {"fill", "Accents and fills"}, {"notice", "Accents and fills"},
	{"success", "Status"}, {"warning", "Status"}, {"danger", "Status"},
	{"", "Lines and effects"},
}

// unreadableHex is what the colour picker shows for a value that is not a plain colour.
const unreadableHex = "#808080"

// colourGroups lists the editor's groups in the order each is first met.
func colourGroups(palette []library.Colour) []colourGroup {
	var groups []colourGroup
	for _, c := range palette {
		label := strings.TrimPrefix(c.Token, "--color-")
		v := colourView{
			Token:    c.Token,
			Label:    strings.ReplaceAll(label, "-", " "),
			Value:    c.Value,
			Own:      c.Own,
			Hex:      unreadableHex,
			Custom:   c.Custom,
			Editable: c.Editable,
		}
		if rgba, ok := pack.ParseColour(c.Value); ok {
			rgba.A = 255
			v.Hex = rgba.CSS()
		}
		name := colourGroupName(label)
		i := slices.IndexFunc(groups, func(have colourGroup) bool { return have.Name == name })
		if i < 0 {
			i = len(groups)
			groups = append(groups, colourGroup{Name: name})
		}
		groups[i].Colours = append(groups[i].Colours, v)
	}
	return groups
}

func colourGroupName(label string) string {
	for _, g := range colourGroupNames {
		if strings.HasPrefix(label, g.prefix) {
			return g.group
		}
	}
	return ""
}
