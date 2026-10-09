package main

import (
	"maps"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

const (
	// mosaicTiles previews, spread over the eras, stand for "any era".
	mosaicTiles = 12
	// The "w" a page asks for a smaller preview with; they are the widths
	// the library prepares.
	smallPreviewParam = "540"
	tinyPreviewParam  = "270"
)

// homeData is built once: the library is fixed.
func (s *server) homeData() map[string]any {
	first, last := s.yearSpan()
	return map[string]any{
		"Title":     "Eras",
		"PageTitle": "Web design eras as style packs for coding agents",
		"Home":      true,
		"Nav":       "home",
		"Eras":      s.eraList,
		"Mosaic":    s.mosaic(),
		"First":     first,
		"Last":      last,
		"PageType": selectView{
			Name:        "archetype",
			Label:       "Page type",
			Options:     s.pageTypeOptions(),
			Multiple:    true,
			Placeholder: "Any page type",
		},
	}
}

// pageTypeOptions notes, for each page type, how many eras have a layout of it.
func (s *server) pageTypeOptions() []option {
	var options []option
	for _, a := range s.archetypes {
		n := 0
		for _, e := range s.lib.Eras {
			if slices.Contains(s.lib.Archetypes(e.Slug), a) {
				n++
			}
		}
		options = append(options, option{Value: a, Label: capitalise(a), Note: plural(n, "era")})
	}
	return options
}

func (s *server) mosaic() []string {
	var previews []string
	for i := 0; i < mosaicTiles && len(s.eraList) > 0; i++ {
		if p := s.eraList[i*len(s.eraList)/mosaicTiles].Preview; p != "" {
			previews = append(previews, p)
		}
	}
	return previews
}

// yearSpan is the years the home page's slider covers.
func (s *server) yearSpan() (first, last int) {
	for _, e := range s.eraList {
		if first == 0 || e.From < first {
			first = e.From
		}
		last = max(last, e.To)
	}
	return first, last
}

func (s *server) home(w http.ResponseWriter, r *http.Request) {
	// render adds to the map it is given.
	s.render(w, r, "home", maps.Clone(s.homeView))
}

func (s *server) era(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	at := slices.IndexFunc(s.eraList, func(e eraView) bool { return e.Slug == slug })
	if at < 0 {
		s.notFound(w, r)
		return
	}
	e := s.eraList[at]
	c := s.lib.Counts(slug)
	// A description reads "what it is: how it looks"; the page sets the
	// two apart.
	lead, traits, split := strings.Cut(e.Description, ": ")
	if split {
		traits = capitalise(traits)
	}
	s.render(w, r, "era", map[string]any{
		"Title":       e.Name,
		"PageTitle":   e.Name + " web design style, " + e.Years,
		"Nav":         "",
		"Era":         e,
		"Description": e.Description,
		"Image":       e.Preview,
		"Lead":        lead,
		"Traits":      traits,
		"Figures": []figure{
			{c.Structures, "layouts"},
			{c.Palettes, "palettes"},
			{c.Types, "type sets"},
			{c.Surfaces, "surface sets"},
		},
	})
}

type figure struct {
	N     int
	Label string
}

func (s *server) saved(w http.ResponseWriter, r *http.Request) {
	// The filter lists every era by the code a short ID carries; the page
	// shows only those the visitor has saved a style of.
	options := []option{{Label: "All saved eras", Selected: true}}
	for _, e := range s.eraList {
		options = append(options, option{Value: e.Code, Label: e.Name, Note: e.Years})
	}
	s.render(w, r, "saved", map[string]any{
		"Title":       "Saved",
		"Nav":         "saved",
		"Description": "The styles you saved in this browser.",
		"EraFilter":   selectView{Name: "saved-era", Label: "Era", Options: options},
	})
}

func (s *server) docs(w http.ResponseWriter, r *http.Request) {
	eras, layouts := s.totals()
	firstEra := s.lib.Eras[0]
	example := s.firstStyle()
	s.render(w, r, "docs", map[string]any{
		"Title":     "Docs",
		"PageTitle": "How chronoskin works",
		"Nav":       "docs",
		"Base":      s.base(r),
		"Totals": plural(eras, "era") + " and " + plural(layouts, "layout") +
			", and every layout of an era can wear any of that era's colour, type and surface sets",
		"Example":     s.lib.Short(example),
		"ExampleLong": example.String(),
		"EraCode":     firstEra.Code,
		"EraName":     eraName(firstEra.Name),
		"Description": "What a chronoskin style pack contains, how styles are put together, " +
			"and how to install one for Claude Code or any other coding agent.",
	})
}

// preview serves /preview/<era>/<n>.jpg, and with ?w=540 or ?w=270 a smaller one.
func (s *server) preview(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(strings.TrimSuffix(r.PathValue("file"), ".jpg"))
	jpeg := s.lib.Preview(r.PathValue("era"), n)
	switch r.URL.Query().Get("w") {
	case smallPreviewParam:
		jpeg = s.lib.PreviewSmall(r.PathValue("era"), n)
	case tinyPreviewParam:
		jpeg = s.lib.PreviewTiny(r.PathValue("era"), n)
	}
	if err != nil || len(jpeg) == 0 {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image/jpeg")
	if r.URL.Query().Get("v") == s.lib.Version {
		w.Header().Set("Cache-Control", cacheForever)
	} else {
		w.Header().Set("Cache-Control", cacheDay)
	}
	w.Write(jpeg)
}
