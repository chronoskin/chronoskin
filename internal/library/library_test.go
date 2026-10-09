package library

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/chronoskin/chronoskin/internal/pack"
)

func fixtureLibrary(t *testing.T) *Library {
	t.Helper()
	dirs, _ := filepath.Glob("../../packs/*/[0-9][0-9]")
	if len(dirs) == 0 {
		t.Fatal("no packs found in packs/")
	}
	out := t.TempDir()
	variants, _ := filepath.Glob("../../packs/*/tokens/[0-9][0-9]")
	if err := Build("test", "v1", out, dirs, Options{Variants: variants}); err != nil {
		t.Fatal(err)
	}
	l, err := Load(out)
	if err != nil {
		t.Fatal(err)
	}
	return l
}

// Splitting a pack into parts and composing it again must give a valid pack.
func TestEveryStyleComposesToAValidPack(t *testing.T) {
	l := fixtureLibrary(t)
	for _, e := range l.Eras {
		for _, s := range e.Structures {
			id := pack.ID{Prefix: l.Prefix, Era: e.Slug, Archetype: s.Archetype,
				Structure: s.N, Palette: s.N, Type: s.N, Surface: s.N}
			files, err := l.Pack(id, "http://localhost:8080/s/"+id.String())
			if err != nil {
				t.Fatal(err)
			}
			dir := t.TempDir()
			for name, content := range files {
				os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644)
			}
			for _, p := range pack.Lint(dir) {
				t.Errorf("%s: %s", id, p)
			}
		}
	}
}

func TestGenerate(t *testing.T) {
	l := fixtureLibrary(t)
	a, err := l.Generate(Request{Seed: 7})
	if err != nil {
		t.Fatal(err)
	}
	b, _ := l.Generate(Request{Seed: 7})
	if a != b {
		t.Errorf("same seed gave %s and %s", a, b)
	}
	seen := map[string]bool{}
	for seed := range uint64(40) {
		id, _ := l.Generate(Request{Seed: seed})
		seen[id.Era] = true
	}
	if len(seen) < 2 {
		t.Errorf("40 seeds reached only %d era", len(seen))
	}

	era := l.Eras[0].Slug
	id, err := l.Generate(Request{Era: era, Mode: "mix", Seed: 3})
	if err != nil || id.Era != era {
		t.Fatalf("era %s: got %s, %v", era, id, err)
	}
	again, err := l.Generate(Request{From: id.String(), Lock: []string{"palette"}, Seed: 9})
	if err != nil || again.Palette != id.Palette || again.Era != id.Era {
		t.Errorf("regenerate with locked palette: %s -> %s, %v", id, again, err)
	}

	if _, err := l.Generate(Request{Era: "no-such-era"}); err == nil {
		t.Error("unknown era: no error")
	}
	if _, err := l.Generate(Request{Lock: []string{"colour"}}); err == nil {
		t.Error("unknown lock: no error")
	}
	if _, err := l.Pack(pack.ID{Prefix: "v1", Era: era, Archetype: "forum", Structure: 99}, ""); err == nil {
		t.Error("unknown style: no error")
	}
}

// What mix mode can produce must still be a valid pack. The full product of
// layouts, palettes, type sets and surface sets runs into the hundreds of
// thousands, so each layout is linted with each of its palettes, and the
// type and surface sets are walked along with them: every part is in some
// combination with every layout.
func TestEveryMixIsAValidPack(t *testing.T) {
	if testing.Short() {
		t.Skip("lints many combinations")
	}
	l := fixtureLibrary(t)
	combos := 0
	for _, e := range l.Eras {
		types, surfaces := every(e.Type), every(e.Surfaces)
		for _, s := range e.Structures {
			for i := range max(len(s.Palettes), len(types), len(surfaces)) {
				p, ty, u := s.Palettes[i%len(s.Palettes)], types[(i+s.N)%len(types)], surfaces[(i+2*s.N)%len(surfaces)]
				id := pack.ID{Prefix: l.Prefix, Era: e.Slug, Archetype: s.Archetype, Structure: s.N, Palette: p, Type: ty, Surface: u}
				files, err := l.Pack(id, "http://localhost:8080/s/"+id.String())
				if err != nil {
					t.Fatal(err)
				}
				dir := t.TempDir()
				for name, content := range files {
					os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644)
				}
				for _, problem := range pack.Lint(dir) {
					t.Errorf("%s: %s", id, problem)
				}
				combos++
			}
		}
	}
	t.Logf("%d combinations", combos)
}
