package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/chronoskin/chronoskin/internal/browser"
)

func shot(args []string) error {
	fs := flag.NewFlagSet("shot", flag.ExitOnError)
	width := fs.Int("width", defaultWidth, "window width")
	height := fs.Int("height", defaultShotHeight, "window height")
	fs.Parse(args)
	if fs.NArg() != 2 {
		return fmt.Errorf("usage: pack shot [-width N] [-height N] <page-or-url> <out.png>")
	}
	page, out := fs.Arg(0), fs.Arg(1)
	if !strings.Contains(page, "://") {
		abs, err := filepath.Abs(page)
		if err != nil {
			return err
		}
		page = "file://" + abs
	}
	b, err := browser.Start()
	if err != nil {
		return err
	}
	defer b.Close()
	c, err := b.Capture(page, browser.CaptureOptions{
		Width:       *width,
		Height:      *height,
		MaxHeight:   *height,
		LoadTimeout: loadTimeout,
		Settle:      shotSettle,
	})
	if err != nil {
		return err
	}
	return os.WriteFile(out, c.PNG, 0o644)
}
