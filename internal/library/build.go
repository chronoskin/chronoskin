package library

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/chronoskin/chronoskin/internal/pack"
)

// structureSections are the STYLE.md sections kept with the structure part.
// "Typography and colour roles" goes with the type part, and Never rules
// are divided between the parts by metric.
var structureSections = []string{"Summary", "Layout", "Components", "Never", "Extending"}

// In the order of pack.Parts.
var tokenDirs = []string{"palettes", "type", "surfaces"}

// variantRoleSections names the rules.md section of each part's role notes.
var variantRoleSections = map[string]string{
	"palette": "Colour roles",
	"type":    "Typography roles",
	"surface": "Surface roles",
}

type Options struct {
	// Variants are directories named <era-slug>/tokens/<nn> holding a
	// tokens.css and a rules.md: a palette, type and surface set without a
	// structure, which widens what mix mode can combine.
	Variants []string
	// Eras gives eras a name and a description of their own. Without an
	// entry an era is named after its first pack, which describes one
	// layout rather than the period.
	Eras map[string]Era
	// Preview renders a self-contained specimen page to a JPEG. When set,
	// every structure gets a preview.jpg.
	Preview func(html string, width int) ([]byte, error)
}

// Build splits whole style packs into parts and writes them as a library.
// Packs of the same era are numbered in sorted order; token sets without a
// structure follow.
func Build(version, prefix, out string, packDirs []string, opts Options) error {
	b := &builder{
		out:           out,
		opts:          opts,
		eras:          map[string]*EraIndex{},
		paletteBlocks: map[string]map[int]string{},
	}
	sort.Strings(packDirs)
	for _, dir := range packDirs {
		if err := b.addPack(dir); err != nil {
			return err
		}
	}
	sort.Strings(opts.Variants)
	for _, dir := range opts.Variants {
		if err := b.addVariant(dir); err != nil {
			return err
		}
	}

	index := Index{Version: version, Prefix: prefix}
	usedCodes := map[string]bool{}
	for _, e := range b.erasOldestFirst() {
		b.allowPalettes(e)
		b.nameEra(e, usedCodes)
		if err := writeJSON(filepath.Join(out, "eras", e.Slug, "era.json"), e.Era); err != nil {
			return err
		}
		index.Eras = append(index.Eras, *e)
	}
	return writeJSON(filepath.Join(out, "index.json"), index)
}

type builder struct {
	out           string
	opts          Options
	eras          map[string]*EraIndex
	paletteBlocks map[string]map[int]string // era slug -> palette number -> :root block
}

func (b *builder) addPack(dir string) error {
	if problems := pack.Lint(dir); len(problems) > 0 {
		return fmt.Errorf("%s: %s (and %d more)", dir, problems[0], len(problems)-1)
	}
	files := map[string]string{}
	for _, name := range pack.Files {
		raw, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			return err
		}
		files[name] = string(raw)
	}
	var m pack.Manifest
	if err := json.Unmarshal([]byte(files["style.json"]), &m); err != nil {
		return err
	}
	title, sections, err := pack.SplitStyle(files["STYLE.md"])
	if err != nil {
		return fmt.Errorf("%s: %w", dir, err)
	}
	blocks, err := pack.SplitTokens(files["tokens.css"])
	if err != nil {
		return fmt.Errorf("%s: %w", dir, err)
	}

	slug := m.Era.Slug
	e := b.eras[slug]
	if e == nil {
		e = &EraIndex{Era: Era{
			Slug:        slug,
			Name:        title,
			Years:       m.Era.Years,
			Description: sections["Summary"],
		}}
		b.eras[slug] = e
		b.paletteBlocks[slug] = map[int]string{}
	}
	n := len(e.Structures) + 1
	if n > maxParts {
		return fmt.Errorf("%s: era %s already has %d styles; a Style ID holds two digits per part", dir, slug, maxParts)
	}

	rules, rest := pack.NeverRules(sections["Never"])
	never := neverByPart(rules)
	own := map[string]string{}
	for _, s := range structureSections {
		own[s] = sections[s]
	}
	own["Never"] = strings.TrimSpace(rest + "\n\n" + strings.Join(never["structure"], "\n"))

	sdir := filepath.Join(b.out, "eras", slug, "structures", nn(n))
	if err := os.MkdirAll(sdir, 0o755); err != nil {
		return err
	}
	for name, content := range map[string]string{
		"specimen.html":  files["specimen.html"],
		"components.css": files["components.css"],
		"rules.md":       joinRules(title, own),
	} {
		if err := os.WriteFile(filepath.Join(sdir, name), []byte(content), 0o644); err != nil {
			return err
		}
	}
	if b.opts.Preview != nil {
		jpeg, err := b.opts.Preview(pack.InlineSpecimen(files), m.Viewport.Width)
		if err != nil {
			return fmt.Errorf("%s: preview: %w", dir, err)
		}
		if err := os.WriteFile(filepath.Join(sdir, "preview.jpg"), jpeg, 0o644); err != nil {
			return err
		}
	}
	roles := map[string]string{"type": sections["Typography and colour roles"]}
	if err := b.writeTokenParts(slug, n, blocks, never, roles); err != nil {
		return err
	}

	part := PartIndex{N: n, Archetype: m.Archetype}
	e.Structures = append(e.Structures, StructureIndex{PartIndex: part, Title: title, Viewport: m.Viewport})
	e.addTokenParts(part)
	return nil
}

// A variant's era must have a pack.
func (b *builder) addVariant(dir string) error {
	slug := filepath.Base(filepath.Dir(filepath.Dir(dir)))
	e := b.eras[slug]
	if e == nil {
		return fmt.Errorf("%s: no pack for era %q", dir, slug)
	}
	if problems := pack.LintTokens(dir); len(problems) > 0 {
		return fmt.Errorf("%s: %s (and %d more)", dir, problems[0], len(problems)-1)
	}
	tokens, err := os.ReadFile(filepath.Join(dir, "tokens.css"))
	if err != nil {
		return err
	}
	blocks, err := pack.SplitTokens(string(tokens))
	if err != nil {
		return fmt.Errorf("%s: %w", dir, err)
	}
	rulesMD, err := os.ReadFile(filepath.Join(dir, "rules.md"))
	if err != nil {
		return err
	}
	sections := variantSections(string(rulesMD))
	rules, _ := pack.NeverRules(sections["Never"])
	n := len(e.Palettes) + 1
	if n > maxParts {
		return fmt.Errorf("%s: era %s already has %d token sets; a Style ID holds two digits per part", dir, slug, maxParts)
	}
	roles := map[string]string{}
	for part, section := range variantRoleSections {
		roles[part] = sections[section]
	}
	if err := b.writeTokenParts(slug, n, blocks, neverByPart(rules), roles); err != nil {
		return err
	}
	// A variant serves the page type of the era's first pack.
	e.addTokenParts(PartIndex{N: n, Archetype: e.Structures[0].Archetype})
	return nil
}

func neverByPart(rules []pack.NeverRule) map[string][]string {
	never := map[string][]string{}
	for _, r := range rules {
		owner := pack.NeverPart(r.Metric)
		never[owner] = append(never[owner], r.Line)
	}
	return never
}

// never and roles are keyed by part name.
func (b *builder) writeTokenParts(slug string, n int, blocks []string, never map[string][]string, roles map[string]string) error {
	for i, dir := range tokenDirs {
		name := pack.Parts[i].Name
		part := tokenPart{Tokens: blocks[i], Never: never[name], Roles: roles[name]}
		if part.Never == nil {
			// Written as [] rather than null.
			part.Never = []string{}
		}
		path := filepath.Join(b.out, "eras", slug, dir, nn(n)+".json")
		if err := writeJSON(path, part); err != nil {
			return err
		}
	}
	b.paletteBlocks[slug][n] = blocks[0]
	return nil
}

func (e *EraIndex) addTokenParts(part PartIndex) {
	e.Palettes = append(e.Palettes, part)
	e.Type = append(e.Type, part)
	e.Surfaces = append(e.Surfaces, part)
}

func (b *builder) erasOldestFirst() []*EraIndex {
	eras := make([]*EraIndex, 0, len(b.eras))
	for _, e := range b.eras {
		eras = append(eras, e)
	}
	sort.Slice(eras, func(i, j int) bool {
		x, y := eras[i], eras[j]
		if x.Years[0] != y.Years[0] {
			return x.Years[0] < y.Years[0]
		}
		if x.Years[1] != y.Years[1] {
			return x.Years[1] < y.Years[1]
		}
		return x.Slug < y.Slug
	})
	return eras
}

// allowPalettes lets each structure take its own palette and any other of
// the era whose roles can be told apart, whatever page type it was made for.
func (b *builder) allowPalettes(e *EraIndex) {
	blocks := map[int]string{}
	for _, p := range e.Palettes {
		blocks[p.N] = b.paletteBlocks[e.Slug][p.N]
	}
	fits := compatible(blocks)
	for i := range e.Structures {
		st := &e.Structures[i]
		for _, p := range e.Palettes {
			if p.N == st.N || fits[p.N] {
				st.Palettes = append(st.Palettes, p.N)
			}
		}
	}
}

// nameEra applies the era's era.json, and gives it a free code when it has none.
func (b *builder) nameEra(e *EraIndex, usedCodes map[string]bool) {
	if named, ok := b.opts.Eras[e.Slug]; ok {
		if named.Name != "" {
			e.Name = named.Name
		}
		if named.Description != "" {
			e.Description = named.Description
		}
		e.Code = named.Code
	}
	if e.Code == "" {
		e.Code = freeCode(e.Slug, usedCodes)
	}
	usedCodes[e.Code] = true
}

func writeJSON(path string, v any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(b, '\n'), 0o644)
}

func joinRules(title string, sections map[string]string) string {
	var b strings.Builder
	b.WriteString("# " + title + "\n")
	for _, s := range structureSections {
		b.WriteString("\n## " + s + "\n\n" + sections[s] + "\n")
	}
	return b.String()
}

// splitSections cuts markdown at the level-2 headings isSection accepts. It
// returns the text under each, and the lines before the first.
func splitSections(md string, isSection func(heading string) bool) (sections map[string]string, preamble []string) {
	sections = map[string]string{}
	current := ""
	var body []string
	flush := func() {
		if current == "" {
			preamble = body
		} else {
			sections[current] = strings.TrimSpace(strings.Join(body, "\n"))
		}
		body = nil
	}
	for _, line := range strings.Split(md, "\n") {
		if h, ok := strings.CutPrefix(line, "## "); ok && isSection(strings.TrimSpace(h)) {
			flush()
			current = strings.TrimSpace(h)
			continue
		}
		body = append(body, line)
	}
	flush()
	return sections, preamble
}

// In a structure's rules.md, a level-2 heading that is not a structure
// section is text of the section it stands in.
func splitRules(md string) (title string, sections map[string]string) {
	sections, preamble := splitSections(md, func(heading string) bool {
		return slices.Contains(structureSections, heading)
	})
	for _, line := range preamble {
		if t, ok := strings.CutPrefix(line, "# "); ok && strings.TrimSpace(t) != "" {
			return strings.TrimSpace(t), sections
		}
	}
	return "", sections
}

func variantSections(md string) map[string]string {
	sections, _ := splitSections(md, func(string) bool { return true })
	return sections
}

// freeCode makes an era code from a slug: its first letter with each later
// character in turn, the first pair not taken yet.
func freeCode(slug string, used map[string]bool) string {
	var chars []byte
	for i := 0; i < len(slug); i++ {
		if c := slug[i]; c != '-' {
			chars = append(chars, c)
		}
	}
	for _, c := range chars[1:] {
		if code := string([]byte{chars[0], c}); !used[code] {
			return code
		}
	}
	for c := byte('0'); c <= '9'; c++ {
		if code := string([]byte{chars[0], c}); !used[code] {
			return code
		}
	}
	return string(chars[:2])
}
