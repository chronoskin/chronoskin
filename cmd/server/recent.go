package main

import (
	"net/http"
	"slices"
	"strings"

	"github.com/chronoskin/chronoskin/internal/library"
	"github.com/chronoskin/chronoskin/internal/pack"
)

// A visitor who presses Generate again and again should keep seeing
// something new. The server stores nothing, so the last styles it sent a
// browser to travel in a cookie, and a new pick avoids them:
//
//   - the same style does not come back within recentStyles picks;
//   - the same layout does not come back within recentLayouts picks, or
//     within one fewer than there are layouts to choose from.
//
// Requests without the cookie (the API, MCP, a first visit) are unaffected,
// and so is the rule that one seed gives one style.
const (
	recentCookie    = "recent"
	recentCookieAge = 30 * 24 * 3600 // seconds the cookie is kept
	recentStyles    = 10
	recentLayouts   = 5
	recentTries     = 48 // further seeds tried before a repeat is accepted
)

// recent is short IDs without adjustments, newest first.
type recent []string

func readRecent(r *http.Request) recent {
	c, err := r.Cookie(recentCookie)
	if err != nil || c.Value == "" {
		return nil
	}
	list := strings.Split(c.Value, ".")
	if len(list) > recentStyles {
		list = list[:recentStyles]
	}
	return list
}

func (rc recent) write(w http.ResponseWriter, newest string) {
	older := slices.DeleteFunc(slices.Clone(rc), func(s string) bool { return s == newest })
	list := append([]string{newest}, older...)
	if len(list) > recentStyles {
		list = list[:recentStyles]
	}
	http.SetCookie(w, &http.Cookie{
		Name:     recentCookie,
		Value:    strings.Join(list, "."),
		Path:     "/",
		MaxAge:   recentCookieAge,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}

// layoutOf is the layout part of a short ID: v1-fl-2314 gives v1-fl-2.
func layoutOf(short string) string {
	if i := strings.LastIndexByte(short, '-'); i > 0 && i+1 < len(short) {
		return short[:i+2]
	}
	return short
}

// layoutChoices counts the layouts of the style's own era when regenerating,
// else those of every era the request allows.
func (s *server) layoutChoices(req library.Request, pickedEra string) int {
	choices := 0
	for _, e := range s.lib.Eras {
		if req.From != "" && e.Slug != pickedEra {
			continue
		}
		if req.From == "" && !library.Among(req.Era, e.Slug) {
			continue
		}
		choices += len(e.Structures)
	}
	return choices
}

// fresh generates a style the browser has not just seen, trying further
// seeds when a pick repeats one. accept, when set, must also hold.
func (s *server) fresh(req library.Request, seen recent, accept func(pack.ID) bool) (pack.ID, error) {
	id, err := s.lib.Generate(req)
	if err != nil {
		return id, err
	}
	// Holding all four parts is a request for this very style, adjusted.
	holdsAll := req.From != "" && len(req.Lock) >= 4
	// How many layouts the pick could have landed on decides how long a
	// layout stays away.
	window := min(recentLayouts, s.layoutChoices(req, id.Era)-1, len(seen))
	holdsLayout := req.From != "" && slices.Contains(req.Lock, "structure")
	ok := func(id pack.ID) bool {
		if accept != nil && !accept(id) {
			return false
		}
		if holdsAll {
			return true
		}
		short := s.lib.Short(id.Base())
		if slices.Contains(seen, short) {
			return false
		}
		if holdsLayout || window <= 0 {
			return true
		}
		layout := layoutOf(short)
		return !slices.ContainsFunc(seen[:window], func(old string) bool { return layoutOf(old) == layout })
	}
	best := id
	for try := 0; !ok(id) && try < recentTries; try++ {
		req.Seed++
		if id, err = s.lib.Generate(req); err != nil {
			return best, nil
		}
		// A pick that only repeats a layout beats one that fails outright.
		if accept == nil || accept(id) {
			best = id
		}
	}
	if ok(id) {
		return id, nil
	}
	return best, nil
}
