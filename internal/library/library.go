// Package library holds a released style library: the eras and, for each,
// the structure, palette, type and surface parts that styles are composed
// from. A library is a directory of plain files; the server loads it into
// memory once and never writes to it.
//
//	<dir>/index.json
//	<dir>/eras/<slug>/era.json
//	<dir>/eras/<slug>/structures/<nn>/{specimen.html,components.css,rules.md}
//	<dir>/eras/<slug>/palettes/<nn>.json
//	<dir>/eras/<slug>/type/<nn>.json
//	<dir>/eras/<slug>/surfaces/<nn>.json
package library

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"path/filepath"
	"regexp"
	"slices"

	"github.com/chronoskin/chronoskin/internal/pack"
)

// Index is index.json: which files to read and which parts may be combined.
type Index struct {
	Version string     `json:"version"` // such as "2026.10"
	Prefix  string     `json:"prefix"`  // leading segment of this library's Style IDs
	Eras    []EraIndex `json:"eras"`
}

const (
	// Style IDs give each part two digits.
	maxParts = 99
	// Pixels: a half and a quarter of the widest preview the builder stores,
	// for a card and for a tile within a card.
	smallWidth   = 540
	tinyWidth    = 270
	smallQuality = 80
)

type Era struct {
	Slug        string `json:"slug"`
	Code        string `json:"code"` // two characters that stand for the era in a short Style ID
	Name        string `json:"name"`
	Years       [2]int `json:"years"`
	Description string `json:"description"`
}

type EraIndex struct {
	Era
	Structures []StructureIndex `json:"structures"`
	Palettes   []PartIndex      `json:"palettes"`
	Type       []PartIndex      `json:"type"`
	Surfaces   []PartIndex      `json:"surfaces"`
}

type PartIndex struct {
	N         int    `json:"n"`
	Archetype string `json:"archetype"`
}

type StructureIndex struct {
	PartIndex
	Title    string        `json:"title"`
	Viewport pack.Viewport `json:"viewport"`
	// Palettes lists the palettes this structure may be combined with: its
	// own, and those whose roles can be told apart. Computed by Build.
	Palettes []int `json:"palettes"`
}

// tokenPart is a palettes/, type/ or surfaces/ file.
type tokenPart struct {
	Tokens string   `json:"tokens"` // the part's :root block
	Never  []string `json:"never"`  // Never rules owned by this part
	// Roles is this part's share of "Typography and colour roles".
	Roles string `json:"roles,omitempty"`
}

type structure struct {
	StructureIndex
	specimen   string
	components string
	sections   map[string]string // Summary, Layout, Components, Extending, Never
	preview    []byte            // JPEG of the specimen with its own tokens; may be empty
	small      []byte            // the same at half the width, for a card
	tiny       []byte            // and at a quarter, for a tile within a card
}

type era struct {
	EraIndex
	structures map[int]*structure
	palettes   map[int]*tokenPart
	types      map[int]*tokenPart
	surfaces   map[int]*tokenPart
}

// A Library is read-only and safe for concurrent use.
type Library struct {
	Index
	eras map[string]*era
}

func Load(dir string) (*Library, error) {
	l := &Library{eras: map[string]*era{}}
	if err := readJSON(filepath.Join(dir, "index.json"), &l.Index); err != nil {
		return nil, err
	}
	if l.Version == "" || l.Prefix == "" {
		return nil, fmt.Errorf("%s: index.json needs version and prefix", dir)
	}
	for _, ei := range l.Eras {
		e, err := loadEra(filepath.Join(dir, "eras", ei.Slug), ei)
		if err != nil {
			return nil, err
		}
		l.eras[ei.Slug] = e
	}
	l.fillMissingCodes()
	if err := l.validate(); err != nil {
		return nil, fmt.Errorf("%s: %w", dir, err)
	}
	return l, nil
}

func loadEra(dir string, ei EraIndex) (*era, error) {
	e := &era{EraIndex: ei, structures: map[int]*structure{}}
	for _, si := range ei.Structures {
		st, err := loadStructure(filepath.Join(dir, "structures", nn(si.N)), si)
		if err != nil {
			return nil, err
		}
		e.structures[si.N] = st
	}
	var err error
	if e.palettes, err = loadTokenParts(filepath.Join(dir, "palettes"), ei.Palettes); err != nil {
		return nil, err
	}
	if e.types, err = loadTokenParts(filepath.Join(dir, "type"), ei.Type); err != nil {
		return nil, err
	}
	if e.surfaces, err = loadTokenParts(filepath.Join(dir, "surfaces"), ei.Surfaces); err != nil {
		return nil, err
	}
	return e, nil
}

func loadStructure(dir string, si StructureIndex) (*structure, error) {
	specimen, err := os.ReadFile(filepath.Join(dir, "specimen.html"))
	if err != nil {
		return nil, err
	}
	components, err := os.ReadFile(filepath.Join(dir, "components.css"))
	if err != nil {
		return nil, err
	}
	rules, err := os.ReadFile(filepath.Join(dir, "rules.md"))
	if err != nil {
		return nil, err
	}
	_, sections := splitRules(string(rules))
	// A library may be built without previews.
	preview, _ := os.ReadFile(filepath.Join(dir, "preview.jpg"))
	return &structure{
		StructureIndex: si,
		specimen:       string(specimen),
		components:     string(components),
		sections:       sections,
		preview:        preview,
		small:          shrink(preview, smallWidth),
		tiny:           shrink(preview, tinyWidth),
	}, nil
}

func loadTokenParts(dir string, index []PartIndex) (map[int]*tokenPart, error) {
	parts := map[int]*tokenPart{}
	for _, pi := range index {
		var part tokenPart
		if err := readJSON(filepath.Join(dir, nn(pi.N)+".json"), &part); err != nil {
			return nil, err
		}
		parts[pi.N] = &part
	}
	return parts, nil
}

// fillMissingCodes serves a library built before eras had codes.
func (l *Library) fillMissingCodes() {
	used := map[string]bool{}
	for i := range l.Eras {
		e := &l.Eras[i]
		if e.Code == "" {
			e.Code = freeCode(e.Slug, used)
			l.eras[e.Slug].Code = e.Code
		}
		used[e.Code] = true
	}
}

var eraCodeRe = regexp.MustCompile(`^[a-z0-9]{2}$`)

func nn(n int) string { return fmt.Sprintf("%02d", n) }

func readJSON(path string, v any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, v); err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}
	return nil
}

// Preview returns the JPEG preview of a structure rendered with its own
// tokens, or nil when the library was built without previews.
func (l *Library) Preview(eraSlug string, n int) []byte {
	if e := l.eras[eraSlug]; e != nil && e.structures[n] != nil {
		return e.structures[n].preview
	}
	return nil
}

// PreviewSmall is Preview at most smallWidth wide, or nil.
func (l *Library) PreviewSmall(eraSlug string, n int) []byte {
	if e := l.eras[eraSlug]; e != nil && e.structures[n] != nil {
		return e.structures[n].small
	}
	return nil
}

// PreviewTiny is Preview at most tinyWidth wide, or nil.
func (l *Library) PreviewTiny(eraSlug string, n int) []byte {
	if e := l.eras[eraSlug]; e != nil && e.structures[n] != nil {
		return e.structures[n].tiny
	}
	return nil
}

// shrink returns a JPEG narrowed to the given width, or as it is when it is
// already that narrow. Done once, at load, so that a card which shows a
// preview small does not download the large one.
func shrink(jpg []byte, width int) []byte {
	if len(jpg) == 0 {
		return nil
	}
	src, err := jpeg.Decode(bytes.NewReader(jpg))
	if err != nil {
		return nil
	}
	b := src.Bounds()
	if b.Dx() <= width {
		return jpg
	}
	w := width
	h := b.Dy() * width / b.Dx()
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := range h {
		y0 := b.Min.Y + y*b.Dy()/h
		y1 := b.Min.Y + (y+1)*b.Dy()/h
		for x := range w {
			x0 := b.Min.X + x*b.Dx()/w
			x1 := b.Min.X + (x+1)*b.Dx()/w
			dst.SetRGBA(x, y, meanColour(src, image.Rect(x0, y0, max(x1, x0+1), max(y1, y0+1))))
		}
	}
	var out bytes.Buffer
	if err := jpeg.Encode(&out, dst, &jpeg.Options{Quality: smallQuality}); err != nil {
		return nil
	}
	return out.Bytes()
}

func meanColour(src image.Image, block image.Rectangle) color.RGBA {
	var sum [3]uint32
	var n uint32
	for y := block.Min.Y; y < block.Max.Y; y++ {
		for x := block.Min.X; x < block.Max.X; x++ {
			r, g, b, _ := src.At(x, y).RGBA()
			sum[0] += r >> 8
			sum[1] += g >> 8
			sum[2] += b >> 8
			n++
		}
	}
	return color.RGBA{uint8(sum[0] / n), uint8(sum[1] / n), uint8(sum[2] / n), 255}
}

type Counts struct{ Structures, Palettes, Types, Surfaces int }

// Counts is all zero for an era the library lacks.
func (l *Library) Counts(eraSlug string) Counts {
	e := l.eras[eraSlug]
	if e == nil {
		return Counts{}
	}
	return Counts{
		Structures: len(e.Structures),
		Palettes:   len(e.Palettes),
		Types:      len(e.Type),
		Surfaces:   len(e.Surfaces),
	}
}

// validate checks what composition relies on, so that a damaged index is
// refused at start-up instead of failing on a request.
func (l *Library) validate() error {
	if len(l.Eras) == 0 {
		return fmt.Errorf("the library has no eras")
	}
	codes := map[string]string{}
	for _, e := range l.Eras {
		if !eraCodeRe.MatchString(e.Code) {
			return fmt.Errorf("era %s: code %q must be two characters, a to z or 0 to 9", e.Slug, e.Code)
		}
		if other, taken := codes[e.Code]; taken {
			return fmt.Errorf("eras %s and %s share the code %q", other, e.Slug, e.Code)
		}
		codes[e.Code] = e.Slug
		if len(e.Structures) == 0 {
			return fmt.Errorf("era %s has no structures", e.Slug)
		}
		for _, s := range e.Structures {
			if s.N < 1 || s.N > maxParts {
				return fmt.Errorf("era %s: structure number %d is out of range", e.Slug, s.N)
			}
			if len(s.Palettes) == 0 || !slices.Contains(s.Palettes, s.N) {
				return fmt.Errorf("era %s: structure %d does not list its own palette", e.Slug, s.N)
			}
			for _, p := range s.Palettes {
				if l.eras[e.Slug].palettes[p] == nil {
					return fmt.Errorf("era %s: structure %d lists palette %d, which does not exist", e.Slug, s.N, p)
				}
			}
			for name, parts := range map[string][]PartIndex{"type": e.Type, "surface": e.Surfaces} {
				if !slices.Contains(every(parts), s.N) {
					return fmt.Errorf("era %s: structure %d has no %s part of its own", e.Slug, s.N, name)
				}
			}
		}
	}
	return nil
}
