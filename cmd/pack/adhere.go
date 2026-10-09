package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/chronoskin/chronoskin/internal/browser"
	"github.com/chronoskin/chronoskin/internal/features"
	"github.com/chronoskin/chronoskin/internal/pack"
)

type pageReport struct {
	Page     string   `json:"page"`
	Distance float64  `json:"distance"` // to the pack's specimen, 0 to 1
	Furthest []string `json:"furthest"` // the dimensions that differ most
	Never    []string `json:"neverFailed"`
	Checked  int      `json:"neverChecked"`
}

type adhereReport struct {
	Pack string `json:"pack"`
	// Rules the pack's own specimen does not satisfy, or that cannot be
	// measured, are not held against the pages.
	Unusable []string     `json:"neverUnusable"`
	Pages    []pageReport `json:"pages"`
	CSS      []string     `json:"cssViolations"`
	CSSFiles int          `json:"cssFiles"`
}

// fileList is a flag that may be given several times.
type fileList []string

func (f *fileList) String() string { return strings.Join(*f, ",") }

func (f *fileList) Set(s string) error {
	*f = append(*f, s)
	return nil
}

// adhere measures pages an agent built from a pack and reports how well
// they hold to it: distance to the specimen, Never rules, token use.
func adhere(args []string) error {
	fs := flag.NewFlagSet("adhere", flag.ExitOnError)
	packDir := fs.String("pack", "", "pack directory (.design)")
	width := fs.Int("width", 0, "viewport width (default: the pack's)")
	jsonOut := fs.String("json", "", "also write the report to this file")
	var cssFiles fileList
	fs.Var(&cssFiles, "css", "stylesheet the agent wrote; repeatable")
	fs.Parse(args)
	if *packDir == "" || fs.NArg() == 0 {
		return fmt.Errorf("usage: pack adhere -pack DIR [-css FILE]... <page>...")
	}

	manifest, err := readManifest(*packDir)
	if err != nil {
		return err
	}
	if *width == 0 {
		*width = manifest.Viewport.Width
	}
	styleMD, err := os.ReadFile(filepath.Join(*packDir, "STYLE.md"))
	if err != nil {
		return err
	}
	_, sections, err := pack.SplitStyle(string(styleMD))
	if err != nil {
		return err
	}
	rules, _ := pack.NeverRules(sections["Never"])
	tokens, err := os.ReadFile(filepath.Join(*packDir, "tokens.css"))
	if err != nil {
		return err
	}

	b, err := browser.Start()
	if err != nil {
		return err
	}
	defer b.Close()

	specimen, err := measure(b, filepath.Join(*packDir, "specimen.html"), *width)
	if err != nil {
		return err
	}
	report := adhereReport{
		Pack:     manifest.ID,
		CSSFiles: len(cssFiles),
		CSS:      []string{},
		Unusable: []string{},
	}
	var usable []pack.NeverRule
	for _, r := range features.CheckNever(rules, specimen.Metrics) {
		rule := fmt.Sprintf("%s %s %s", r.Rule.Metric, r.Rule.Op, r.Rule.Value)
		switch {
		case r.Skipped != "":
			report.Unusable = append(report.Unusable, rule+": "+r.Skipped)
		case !r.Pass:
			report.Unusable = append(report.Unusable, rule+": the specimen itself measures "+r.Measured)
		default:
			usable = append(usable, r.Rule)
		}
	}
	for _, path := range fs.Args() {
		page, err := measure(b, path, *width)
		if err != nil {
			return err
		}
		report.Pages = append(report.Pages, comparePage(path, page, specimen, usable))
	}
	for _, path := range cssFiles {
		css, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		violations, err := pack.LintCSS(string(css), string(tokens))
		if err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}
		for _, v := range violations {
			report.CSS = append(report.CSS, filepath.Base(path)+": "+v)
		}
	}

	printReport(report)
	if *jsonOut == "" {
		return nil
	}
	out, _ := json.MarshalIndent(report, "", "  ")
	return os.WriteFile(*jsonOut, append(out, '\n'), 0o644)
}

func readManifest(packDir string) (pack.Manifest, error) {
	var m pack.Manifest
	raw, err := os.ReadFile(filepath.Join(packDir, "style.json"))
	if err != nil {
		return m, err
	}
	err = json.Unmarshal(raw, &m)
	return m, err
}

func measure(b *browser.Browser, path string, width int) (*features.Page, error) {
	c, err := b.Capture(fileURL(path), browser.CaptureOptions{
		Width:       width,
		Height:      adhereHeight,
		MaxHeight:   adhereMaxHeight,
		LoadTimeout: loadTimeout,
		Settle:      adhereSettle,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return features.Parse(c.Features)
}

func comparePage(path string, page, specimen *features.Page, usable []pack.NeverRule) pageReport {
	distance, deltas := features.Compare(page.Vector(), specimen.Vector())
	pr := pageReport{
		Page:     path,
		Distance: float64(int(distance*1000)) / 1000, // three decimal places
		Never:    []string{},
	}
	for _, d := range deltas[:min(furthestCount, len(deltas))] {
		if d.Share >= furthestShare {
			pr.Furthest = append(pr.Furthest, d.String())
		}
	}
	for _, r := range features.CheckNever(usable, page.Metrics) {
		if r.Skipped != "" {
			continue
		}
		pr.Checked++
		if !r.Pass {
			failed := fmt.Sprintf("%s %s %s: measured %s", r.Rule.Metric, r.Rule.Op, r.Rule.Value, r.Measured)
			pr.Never = append(pr.Never, failed)
		}
	}
	return pr
}

func printReport(report adhereReport) {
	fmt.Printf("pack %s\n", report.Pack)
	for _, u := range report.Unusable {
		fmt.Printf("  unusable rule: %s\n", u)
	}
	for _, p := range report.Pages {
		fmt.Printf("\n%s\n  distance to specimen: %.3f\n  never rules: %d of %d pass\n",
			p.Page, p.Distance, p.Checked-len(p.Never), p.Checked)
		for _, n := range p.Never {
			fmt.Printf("    FAIL %s\n", n)
		}
		for _, f := range p.Furthest {
			fmt.Printf("    far: %s\n", f)
		}
	}
	if report.CSSFiles == 0 {
		return
	}
	fmt.Printf("\ncss: %d violations in %d files\n", len(report.CSS), report.CSSFiles)
	for _, v := range report.CSS[:min(cssViolationsShown, len(report.CSS))] {
		fmt.Printf("    %s\n", v)
	}
}

func fileURL(path string) string {
	if strings.Contains(path, "://") {
		return path
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return path
	}
	return "file://" + abs
}
