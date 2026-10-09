package library

import (
	"math"

	"github.com/chronoskin/chronoskin/internal/pack"
)

// contrastPairs are the role pairs a structure always puts on top of each
// other, whatever palette fills them.
var contrastPairs = [][2]string{
	{"--color-text", "--color-surface"},
	{"--color-text", "--color-canvas"},
	{"--color-bar-text", "--color-bar"},
	{"--color-button-text", "--color-button"},
	{"--color-link", "--color-surface"},
}

// contrasts returns the log contrast ratio of each role pair of a palette
// block, or false when a colour cannot be read. A translucent colour is
// measured as it shows, laid over what lies under it: glass palettes are
// white on white otherwise.
func contrasts(paletteBlock string) ([]float64, bool) {
	tokens, err := pack.TokenValues(paletteBlock)
	if err != nil {
		return nil, false
	}
	page, ok := pack.ParseColour(tokens["--color-page"])
	if !ok {
		// A palette that does not say what its page is gets a mid grey,
		// which favours neither light nor dark glass.
		page = pack.RGBA{R: 128, G: 128, B: 128, A: 255}
	}
	page.A = 255
	out := make([]float64, len(contrastPairs))
	for i, pair := range contrastPairs {
		fg, okA := pack.ParseColour(tokens[pair[0]])
		bg, okB := pack.ParseColour(tokens[pair[1]])
		if !okA || !okB {
			return nil, false
		}
		bg = over(bg, page)
		fg = over(fg, bg)
		a, b := relativeLuminance(fg), relativeLuminance(bg)
		out[i] = math.Log((math.Max(a, b) + 0.05) / (math.Min(a, b) + 0.05))
	}
	return out, true
}

func over(top, under pack.RGBA) pack.RGBA {
	a := float64(top.A) / 255
	mix := func(t, u uint8) uint8 {
		return uint8(math.Round(float64(t)*a + float64(u)*(1-a)))
	}
	return pack.RGBA{R: mix(top.R, under.R), G: mix(top.G, under.G), B: mix(top.B, under.B), A: 255}
}

// relativeLuminance is the WCAG relative luminance of an opaque colour.
func relativeLuminance(c pack.RGBA) float64 {
	var lin [3]float64
	for i, v := range [3]uint8{c.R, c.G, c.B} {
		f := float64(v) / 255
		if f <= 0.04045 {
			lin[i] = f / 12.92
		} else {
			lin[i] = math.Pow((f+0.055)/1.055, 2.4)
		}
	}
	return 0.2126*lin[0] + 0.7152*lin[1] + 0.0722*lin[2]
}

// legibleContrast is the log of a 2:1 contrast ratio. Below it a pair of
// roles is hard to tell apart at all, whatever the size of the text.
var legibleContrast = math.Log(2)

// compatible decides which palettes may be combined with structures other
// than their own: those in which every pair of roles a structure puts on
// top of each other can be told apart. The floor is fixed, not relative to
// the era's other palettes, so that adding a palette never pushes a sound
// one out.
func compatible(palettes map[int]string) map[int]bool {
	ok := map[int]bool{}
	for n, block := range palettes {
		c, readable := contrasts(block)
		// Colours that cannot be read stay with their own structure.
		ok[n] = readable
		for _, contrast := range c {
			if contrast < legibleContrast {
				ok[n] = false
			}
		}
	}
	return ok
}
