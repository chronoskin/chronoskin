package library

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/chronoskin/chronoskin/internal/pack"
)

func TestShortIDRoundTrip(t *testing.T) {
	lib := fixtureLibrary(t)
	for _, e := range lib.Eras {
		st := e.Structures[len(e.Structures)-1]
		id := pack.ID{Prefix: lib.Prefix, Era: e.Slug, Archetype: st.Archetype, Structure: st.N,
			Palette: st.Palettes[0], Type: st.N, Surface: st.N, Density: "roomy", Colours: testColours}
		short := lib.Short(id)
		if strings.Contains(short, ".") || len(short) > 40 {
			t.Errorf("%s: short form %q", id, short)
		}
		for _, form := range []string{short, id.String()} {
			if got, ok := lib.Parse(form); !ok || got != id {
				t.Errorf("Parse(%q) = %v, %v; want %v", form, got, ok, id)
			}
		}
	}
	for _, bad := range []string{"v1-zz-1111", "v1-" + lib.Eras[0].Code + "-z111", "v1-te-111", "v1.nowhere.blog.s01.p01.t01.u01.dx"} {
		if id, ok := lib.Parse(bad); ok {
			if _, err := lib.Style(id); err == nil {
				t.Errorf("Parse(%q) gave a style", bad)
			}
		}
	}
}

func TestAdjustedPackIsValid(t *testing.T) {
	lib := fixtureLibrary(t)
	for _, e := range lib.Eras {
		st := e.Structures[0]
		base := pack.ID{Prefix: lib.Prefix, Era: e.Slug, Archetype: st.Archetype, Structure: st.N, Palette: st.N, Type: st.N, Surface: st.N}
		plain, err := lib.Pack(base, "")
		if err != nil {
			t.Fatal(err)
		}
		for _, density := range []string{"compact", "roomy"} {
			id, err := lib.Generate(Request{From: base.String(), Lock: partNames, Density: density,
				Colours: map[string]string{"--color-accent": "#ff5500", "--color-overlay": "rgba(10, 20, 30, 0.5)"}})
			if err != nil {
				t.Fatal(err)
			}
			files, err := lib.Pack(id, "http://localhost:8080/s/"+lib.Short(id))
			if err != nil {
				t.Fatalf("%s: %v", id, err)
			}
			if problems := lintFiles(t, files); len(problems) > 0 {
				t.Errorf("%s: %v", id, problems[0])
			}
			if !strings.Contains(files["tokens.css"], "--color-accent: #ff5500;") || !strings.Contains(files["tokens.css"], "--color-overlay: rgba(10, 20, 30, 0.502);") {
				t.Errorf("%s: colours not applied", id)
			}
			// Setting a colour to what the palette already has is no change.
			colours, _ := lib.Palette(base)
			same := map[string]string{}
			for _, c := range colours {
				if c.Editable {
					same[c.Token] = c.Own
				}
			}
			if back, err := lib.Generate(Request{From: id.String(), Lock: partNames, Colours: same, Density: "normal"}); err != nil || back != base {
				t.Errorf("%s: resetting gave %v, %v; want %v", id, back, err, base)
			}
			if files["components.css"] == plain["components.css"] {
				t.Errorf("%s: spacing did not change", id)
			}
			if !strings.Contains(files["style.json"], `"density": "`+density+`"`) {
				t.Errorf("%s: style.json lacks the density", id)
			}
		}
	}
}

func TestNextKeepsAdjustments(t *testing.T) {
	lib := fixtureLibrary(t)
	e := lib.Eras[0]
	st := e.Structures[0]
	id := pack.ID{Prefix: lib.Prefix, Era: e.Slug, Archetype: st.Archetype, Structure: st.N, Palette: st.N, Type: st.N, Surface: st.N, Density: "compact", Colours: testColours}
	for _, part := range []string{"structure", "palette", "type", "surface"} {
		for _, step := range []int{1, -1} {
			next, err := lib.Step(id, part, step)
			if err != nil {
				t.Fatal(err)
			}
			// Stepping the palette drops colours set by hand; nothing else does.
			want := id.Colours
			if part == "palette" {
				want = ""
			}
			if next.Density != "compact" || next.Colours != want {
				t.Errorf("%s %+d: adjustments are %q, %q", part, step, next.Density, next.Colours)
			}
			back, _ := lib.Step(next, part, -step)
			if part != "palette" && back != id {
				t.Errorf("%s: %+d then %+d gives %v, want %v", part, step, -step, back, id)
			}
		}
	}
}

func lintFiles(t *testing.T, files map[string]string) []pack.Problem {
	t.Helper()
	dir := t.TempDir()
	for name, content := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return pack.Lint(dir)
}

// testColours is one opaque and one translucent override, encoded.
var testColours = pack.EncodeColours(map[string]pack.RGBA{
	"--color-accent":  {R: 0x2b, G: 0x55, B: 0xe0, A: 255},
	"--color-overlay": {R: 0, G: 0, B: 0, A: 128},
})

// Layouts of one era may be different kinds of page. Stepping and
// regenerating must still reach every one of them, and the page type in the
// ID must follow the layout.
func TestEveryLayoutIsReachable(t *testing.T) {
	lib := fixtureLibrary(t)
	for _, e := range lib.Eras {
		first := e.Structures[0]
		start := pack.ID{Prefix: lib.Prefix, Era: e.Slug, Archetype: first.Archetype, Structure: first.N, Palette: first.N, Type: first.N, Surface: first.N}
		for name, next := range map[string]func(pack.ID) (pack.ID, error){
			"step": func(id pack.ID) (pack.ID, error) { return lib.Step(id, "structure", 1) },
			"regenerate": func(id pack.ID) (pack.ID, error) {
				return lib.Generate(Request{From: id.String(), Lock: []string{"palette"}})
			},
		} {
			seen := map[int]bool{}
			for id := start; !seen[id.Structure]; {
				seen[id.Structure] = true
				var err error
				if id, err = next(id); err != nil {
					t.Fatalf("%s %s: %v", e.Slug, name, err)
				}
				if _, err := lib.Style(id); err != nil {
					t.Fatalf("%s %s gave %v: %v", e.Slug, name, id, err)
				}
			}
			// Every layout that accepts the starting palette comes round.
			want := 0
			for _, st := range e.Structures {
				if slices.Contains(st.Palettes, start.Palette) {
					want++
				}
			}
			if len(seen) != want {
				t.Errorf("%s: %s reaches %d of %d layouts", e.Slug, name, len(seen), want)
			}
		}
	}
}
