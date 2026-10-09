// Command pack is the offline tool for style packs: it checks them
// against the format, measures pages against them and builds a library
// out of them. It is never deployed with the server.
//
//	pack lint <pack-dir>...            check style packs against the format
//	pack lint-css [-pack .design] <file.css>...
//	                                   check an agent's CSS against a pack's tokens
//	pack adhere -pack DIR [-css FILE]... [-width N] [-json FILE] <page>...
//	                                   measure built pages against a pack
//	pack views [-phone] [-closed] [-leftovers] <pack-dir>...
//	                                   click through each specimen's views in a browser
//	pack shot [-width N] [-height N] [-scale N] <page-or-url> <out.png>
//	                                   a picture of a page in a window of that size
//	pack capture [-width N] [-height N] [-max-height N] [-timeout S] <url> <out-prefix>
//	                                   write <out-prefix>.png and .features.json
//	pack build-library -version V -out DIR [-prefix P] [-previews] <packs-dir>
//	                                   split the packs of every era into parts and write a library
//	pack compose [-library DIR] [-base-url URL] -out DIR <style-id>
//	                                   write the pack of a Style ID from a library
//
// The browser is found through $CHROME and started with $CHROME_FLAGS; see
// internal/browser.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chronoskin/chronoskin/internal/browser"
	"github.com/chronoskin/chronoskin/internal/library"
	"github.com/chronoskin/chronoskin/internal/pack"
)

const usage = "usage: pack lint|lint-css|adhere|views|shot|capture|build-library|compose ..."

const (
	defaultPackDir    = ".design"
	defaultLibraryDir = "library"
	defaultPrefix     = "v1"
	defaultSite       = "https://chrono.skin"
	defaultWidth      = 1024 // window width of capture and shot, in pixels
	defaultHeight     = 768  // window height of capture
	defaultShotHeight = 2000 // window height of shot
	defaultMaxHeight  = 4000 // tallest screenshot capture takes
	defaultLoadWait   = 20   // seconds capture waits for a page to load
)

const (
	loadTimeout = 20 * time.Second
	// How long a page is left to settle after loading, before it is
	// measured or pictured.
	captureSettle = 2 * time.Second
	shotSettle    = time.Second
	adhereSettle  = 500 * time.Millisecond
	previewSettle = 300 * time.Millisecond

	pixelColourCount = 16
	// Enough for a card of the site on a screen of double density.
	previewWidth   = 1080
	previewQuality = 82

	// The window adhere measures in, and the cap on its screenshot.
	adhereHeight    = 900
	adhereMaxHeight = 4000
	// adhere names at most furthestCount dimensions per page, and only
	// those at least furthestShare of their scale away from the specimen.
	furthestCount = 5
	furthestShare = 0.15
	// cssViolationsShown are printed; the JSON report has all of them.
	cssViolationsShown = 20

	viewsHeight = 800
	// The narrow window a responsive specimen must also work in.
	phoneWidth  = 390
	tabletWidth = 768
)

func main() {
	log.SetFlags(0)
	if len(os.Args) < 2 {
		log.Print(usage)
		os.Exit(2)
	}
	args := os.Args[2:]
	switch os.Args[1] {
	case "lint":
		os.Exit(lint(args))
	case "lint-css":
		exit(lintCSS(args))
	case "views":
		exit(views(args))
	case "adhere":
		fatal(adhere(args))
	case "shot":
		fatal(shot(args))
	case "capture":
		fatal(capture(args))
	case "build-library":
		fatal(buildLibrary(args))
	case "compose":
		fatal(compose(args))
	default:
		log.Print(usage)
		os.Exit(2)
	}
}

// exit ends a command that reports findings: 0 for none, 1 for some, 2 when it could not run.
func exit(status int, err error) {
	if err != nil {
		log.Print(err)
	}
	os.Exit(status)
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func lint(dirs []string) int {
	if len(dirs) == 0 {
		log.Print("usage: pack lint <pack-dir>...")
		return 2
	}
	status := 0
	for _, dir := range dirs {
		problems := pack.Lint(dir)
		for _, p := range problems {
			fmt.Printf("%s/%s\n", dir, p)
			status = 1
		}
		if len(problems) == 0 {
			fmt.Printf("%s: ok\n", dir)
		}
	}
	return status
}

func capture(args []string) error {
	fs := flag.NewFlagSet("capture", flag.ExitOnError)
	width := fs.Int("width", defaultWidth, "viewport width")
	height := fs.Int("height", defaultHeight, "viewport height")
	maxHeight := fs.Int("max-height", defaultMaxHeight, "screenshot height cap")
	timeout := fs.Int("timeout", defaultLoadWait, "seconds to wait for the page to load")
	fs.Parse(args)
	if fs.NArg() != 2 {
		return fmt.Errorf("usage: pack capture [flags] <url> <out-prefix>")
	}
	url, prefix := fs.Arg(0), fs.Arg(1)

	b, err := browser.Start()
	if err != nil {
		return err
	}
	defer b.Close()
	c, err := b.Capture(url, browser.CaptureOptions{
		Width:       *width,
		Height:      *height,
		MaxHeight:   *maxHeight,
		LoadTimeout: time.Duration(*timeout) * time.Second,
		Settle:      captureSettle,
	})
	if err != nil {
		return err
	}
	colours, err := browser.PixelColours(c.PNG, pixelColourCount)
	if err != nil {
		return err
	}
	out, err := json.MarshalIndent(map[string]any{
		"url":          url,
		"viewport":     []int{*width, *height},
		"computed":     c.Features,
		"pixelColours": colours,
	}, "", " ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(prefix+".png", c.PNG, 0o644); err != nil {
		return err
	}
	return os.WriteFile(prefix+".features.json", out, 0o644)
}

func buildLibrary(args []string) error {
	fs := flag.NewFlagSet("build-library", flag.ExitOnError)
	version := fs.String("version", "", "library version, such as 2026.10")
	prefix := fs.String("prefix", defaultPrefix, "leading segment of the library's Style IDs")
	out := fs.String("out", "", "output directory; must not exist yet, a library is never edited")
	previews := fs.Bool("previews", false, "render a preview image of every structure, when a Chromium is installed")
	fs.Parse(args)
	if *version == "" || *out == "" || fs.NArg() != 1 {
		return fmt.Errorf("usage: pack build-library -version V -out DIR [-previews] <packs-dir>")
	}
	if _, err := os.Stat(*out); err == nil {
		return fmt.Errorf("%s already exists", *out)
	}
	packs, opts, err := library.Source(fs.Arg(0))
	if err != nil {
		return err
	}
	if *previews {
		// The pictures are a nicety: without a browser to draw them the
		// library is built all the same, and the site shows its cards
		// without pictures.
		b, err := browser.Start()
		if err != nil {
			fmt.Fprintf(os.Stderr, "no preview pictures: %v\n", err)
		} else {
			defer b.Close()
			opts.Preview = func(html string, width int) ([]byte, error) {
				return preview(b, html, width)
			}
		}
	}
	if err := library.Build(*version, *prefix, *out, packs, opts); err != nil {
		return err
	}
	fmt.Printf("%s: library %s from %d eras, %d packs and %d token sets\n",
		*out, *version, len(opts.Eras), len(packs), len(opts.Variants))
	return nil
}

// preview renders the top of a specimen at its design width as a JPEG.
func preview(b *browser.Browser, html string, width int) ([]byte, error) {
	dir, err := os.MkdirTemp("", "chronoskin-preview-")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "specimen.html")
	if err := os.WriteFile(path, []byte(html), 0o644); err != nil {
		return nil, err
	}
	height := width * 2 / 3
	shot, err := b.Capture("file://"+path, browser.CaptureOptions{
		Width:       width,
		Height:      height,
		MaxHeight:   height,
		LoadTimeout: loadTimeout,
		Settle:      previewSettle,
	})
	if err != nil {
		return nil, err
	}
	src, err := png.Decode(bytes.NewReader(shot.PNG))
	if err != nil {
		return nil, err
	}
	var out bytes.Buffer
	err = jpeg.Encode(&out, narrow(src, previewWidth), &jpeg.Options{Quality: previewQuality})
	return out.Bytes(), err
}

// narrow scales a picture down to at most maxWidth, each pixel the average
// of the block it covers. A narrower picture keeps its size.
func narrow(src image.Image, maxWidth int) *image.RGBA {
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	if w > maxWidth {
		w = maxWidth
		h = bounds.Dy() * maxWidth / bounds.Dx()
	}
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := range h {
		y0 := bounds.Min.Y + y*bounds.Dy()/h
		y1 := bounds.Min.Y + (y+1)*bounds.Dy()/h
		for x := range w {
			x0 := bounds.Min.X + x*bounds.Dx()/w
			x1 := bounds.Min.X + (x+1)*bounds.Dx()/w
			dst.SetRGBA(x, y, meanColour(src, image.Rect(x0, y0, max(x1, x0+1), max(y1, y0+1))))
		}
	}
	return dst
}

func meanColour(src image.Image, block image.Rectangle) color.RGBA {
	var r, g, b, n uint32
	for y := block.Min.Y; y < block.Max.Y; y++ {
		for x := block.Min.X; x < block.Max.X; x++ {
			pr, pg, pb, _ := src.At(x, y).RGBA()
			r += pr >> 8
			g += pg >> 8
			b += pb >> 8
			n++
		}
	}
	return color.RGBA{uint8(r / n), uint8(g / n), uint8(b / n), 255}
}

// compose writes the pack of a Style ID from a library directory: what the
// server would hand out, without the server.
func compose(args []string) error {
	fs := flag.NewFlagSet("compose", flag.ExitOnError)
	libDir := fs.String("library", defaultLibraryDir, "library directory")
	base := fs.String("base-url", defaultSite, "site the pack names as its source")
	out := fs.String("out", "", "directory to write the five files into")
	fs.Parse(args)
	if fs.NArg() != 1 || *out == "" {
		return fmt.Errorf("usage: pack compose [-library DIR] [-base-url URL] -out DIR <style-id>")
	}
	lib, err := library.Load(*libDir)
	if err != nil {
		return err
	}
	id, ok := pack.ParseID(fs.Arg(0))
	if !ok {
		return fmt.Errorf("%q is not a Style ID", fs.Arg(0))
	}
	files, err := lib.Pack(id, strings.TrimRight(*base, "/")+"/s/"+id.String())
	if err != nil {
		return err
	}
	if err := os.MkdirAll(*out, 0o755); err != nil {
		return err
	}
	for name, content := range files {
		if err := os.WriteFile(filepath.Join(*out, name), []byte(content), 0o644); err != nil {
			return err
		}
	}
	return nil
}
