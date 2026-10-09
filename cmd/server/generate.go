package main

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"

	"github.com/chronoskin/chronoskin/internal/library"
	"github.com/chronoskin/chronoskin/internal/pack"
)

func randomSeed() uint64 {
	var b [8]byte
	rand.Read(b[:])
	// Seeds travel in URLs and JSON, so keep them within 53 bits.
	return binary.LittleEndian.Uint64(b[:]) >> 11
}

// parseRequest reads the parameters shared by /api/generate and generate_style.
func parseRequest(get func(string) string, locks []string) (library.Request, error) {
	req := library.Request{
		Era:       get("era"),
		Archetype: get("archetype"),
		Mode:      get("mode"),
		From:      get("from"),
		Density:   get("density"),
	}
	for _, l := range locks {
		for _, part := range strings.Split(l, ",") {
			if part = strings.TrimSpace(part); part != "" {
				req.Lock = append(req.Lock, part)
			}
		}
	}
	seed := get("seed")
	if seed == "" {
		req.Seed = randomSeed()
		return req, nil
	}
	n, err := strconv.ParseUint(seed, 10, 64)
	if err != nil {
		return req, errors.New("seed must be a non-negative integer")
	}
	req.Seed = n
	return req, nil
}

// generate is /api/generate: it picks a style and sends a browser to its
// page, or describes it to a program as JSON.
func (s *server) generate(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	req, err := parseRequest(q.Get, q["lock"])
	if err != nil {
		s.fail(w, r, http.StatusBadRequest, err.Error())
		return
	}
	req.Colours = handSetColours(q)
	req.NoColours = q.Get("colours") == "none"
	// Eras and page types may be several: repeated parameters or commas.
	types := list(q["archetype"])
	pool := list(q["era"])
	req.Archetype = strings.Join(types, ",")
	req.Era = strings.Join(pool, ",")
	if from, ok := s.lib.Parse(req.From); ok {
		req.Era = ""
		// The style page's Regenerate says "among=eras": pick among the
		// eras chosen there, and among all of them when none is chosen. A
		// kept part belongs to its era, so then the style stays where it
		// is, as it does when the only era chosen is its own. Without
		// "among", regenerating stays in the style's era.
		ownEraOnly := len(pool) == 1 && pool[0] == from.Era
		if q.Get("among") == "eras" && len(req.Lock) == 0 && !ownEraOnly {
			req.From = ""
			req.Era = strings.Join(pool, ",")
			req.Mode = "mix"
		}
	}
	// A link that leaves the chosen eras ("Any era") still carries them on.
	if q.Has("pool") {
		pool = list(q["pool"])
	}
	// "not" asks for any era but the named one, so that the style page can
	// offer a way out of the era it shows.
	var accept func(pack.ID) bool
	anyEra := req.Era == "" || req.Era == "any"
	if not := q.Get("not"); not != "" && req.From == "" && anyEra && len(s.lib.Eras) > 1 {
		accept = func(id pack.ID) bool { return id.Era != not }
	}
	seen := readRecent(r)
	id, err := s.fresh(req, seen, accept)
	if errors.Is(err, library.ErrNotFound) && req.From == "" && req.Archetype != "" {
		// None of the chosen eras has a page of the chosen types.
		req.Archetype = ""
		id, err = s.fresh(req, seen, accept)
	}
	if err != nil {
		s.failWith(w, r, err)
		return
	}
	w.Header().Set("Cache-Control", "no-store")
	if q.Get("format") == "json" || strings.Contains(r.Header.Get("Accept"), "application/json") {
		writeJSON(w, s.describe(id, s.base(r)))
		return
	}
	// What was kept stays kept on the next page. A form that locks parts
	// only to hold the style still (the colour editor) says so with "keep".
	keep := req.Lock
	if q.Has("keep") {
		keep = strings.FieldsFunc(q.Get("keep"), func(r rune) bool { return r == ',' })
	}
	seen.write(w, s.lib.Short(id.Base()))
	http.Redirect(w, r, s.link(id)+pageQuery(keep, types, pool), http.StatusSeeOther)
}

// handSetColours reads one parameter per palette token. It is nil when
// there are none, which leaves a regenerated style's colours as they are.
func handSetColours(q url.Values) map[string]string {
	var colours map[string]string
	for name, values := range q {
		if !strings.HasPrefix(name, "--color-") || values[0] == "" {
			continue
		}
		if colours == nil {
			colours = map[string]string{}
		}
		colours[name] = values[0]
	}
	return colours
}

// list splits parameter values that may each hold several names joined by
// commas, dropping empty ones and "any".
func list(values []string) []string {
	var out []string
	for _, v := range values {
		for _, name := range strings.Split(v, ",") {
			name = strings.TrimSpace(name)
			if name != "" && name != "any" && !slices.Contains(out, name) {
				out = append(out, name)
			}
		}
	}
	return out
}
