package main

import (
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/chronoskin/chronoskin/internal/pack"
)

var nameTailRe = regexp.MustCompile(`\s*(\(.*\)|,?\s*\(?\d{4}(\s*(to|-)\s*\d{4})?\)?)\s*$`)

// eraName is an era's display name: the library's title without trailing
// years or asides, which the page shows separately.
func eraName(name string) string {
	for {
		short := nameTailRe.ReplaceAllString(name, "")
		if short == name || short == "" {
			return name
		}
		name = short
	}
}

func years(y [2]int) string {
	return strconv.Itoa(y[0]) + " to " + strconv.Itoa(y[1])
}

// capitalise expects ASCII, and not the empty string.
func capitalise(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

// plural is "1 layout" or "3 layouts".
func plural(n int, one string) string {
	if n == 1 {
		return "1 " + one
	}
	return strconv.Itoa(n) + " " + one + "s"
}

func twoDigits(n int) string {
	if n < 10 {
		return "0" + strconv.Itoa(n)
	}
	return strconv.Itoa(n)
}

type layoutView struct {
	Name, Link string
	Preview    string
	Archetype  string // the kind of page the layout is
}

type eraView struct {
	Slug, Name, Years, Description string
	Code                           string       // the two characters of a short style ID
	From, To                       int          // the era's years
	Preview                        string       // empty when the library has no previews
	Link                           string       // the era's first style
	LayoutViews                    []layoutView // one per structure
}

func (s *server) previewURL(era string, structure int) string {
	if len(s.lib.Preview(era, structure)) == 0 {
		return ""
	}
	// The library version in the address lets a browser keep the image for good.
	return "/preview/" + era + "/" + strconv.Itoa(structure) + ".jpg?v=" + s.lib.Version
}

// pureID is the style made of one structure and its own tokens.
func (s *server) pureID(era, archetype string, n int) pack.ID {
	return pack.ID{
		Prefix:    s.lib.Prefix,
		Era:       era,
		Archetype: archetype,
		Structure: n,
		Palette:   n,
		Type:      n,
		Surface:   n,
	}
}

// firstStyle is the library's first style, the one the documentation shows.
func (s *server) firstStyle() pack.ID {
	era := s.lib.Eras[0]
	first := era.Structures[0]
	return s.pureID(era.Slug, first.Archetype, first.N)
}

func (s *server) link(id pack.ID) string { return "/s/" + s.lib.Short(id) }

// eraViews also returns the page types the library has layouts of, sorted.
func (s *server) eraViews() (views []eraView, archetypes []string) {
	for _, e := range s.lib.Eras {
		first := e.Structures[0]
		v := eraView{
			Slug:        e.Slug,
			Name:        eraName(e.Name),
			Years:       years(e.Years),
			Description: e.Description,
			Code:        e.Code,
			From:        e.Years[0],
			To:          e.Years[1],
			Link:        s.link(s.pureID(e.Slug, first.Archetype, first.N)),
			Preview:     s.previewURL(e.Slug, first.N),
		}
		for _, st := range e.Structures {
			v.LayoutViews = append(v.LayoutViews, layoutView{
				Name:      eraName(st.Title),
				Archetype: st.Archetype,
				Link:      s.link(s.pureID(e.Slug, st.Archetype, st.N)),
				Preview:   s.previewURL(e.Slug, st.N),
			})
		}
		for _, a := range s.lib.Archetypes(e.Slug) {
			if !slices.Contains(archetypes, a) {
				archetypes = append(archetypes, a)
			}
		}
		views = append(views, v)
	}
	slices.Sort(archetypes)
	return views, archetypes
}

type option struct {
	Value, Label, Note string
	Selected           bool
}

type selectView struct {
	Name, Label string
	Options     []option
	Submit      bool   // submit the form as soon as an option is chosen
	Multiple    bool   // checkboxes: any number of options, or none
	Placeholder string // what a multiple select shows with nothing chosen
}

func (s *server) totals() (eras, layouts int) {
	for _, e := range s.lib.Eras {
		layouts += len(e.Structures)
	}
	return len(s.lib.Eras), layouts
}
