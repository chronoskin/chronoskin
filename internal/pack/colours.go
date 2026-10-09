package pack

import (
	"encoding/base32"
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// A style may carry its own colours: any palette token set by hand. They
// travel in the Style ID, so the same ID still gives the same files and
// nothing has to be stored. Each override is the token's position in the
// palette vocabulary and its colour; the list is written in base 32.

// RGBA is a colour with straight alpha, 255 being opaque.
type RGBA struct{ R, G, B, A uint8 }

// CSS writes the colour the way palettes do: six hex digits, or rgba()
// when it is translucent.
func (c RGBA) CSS() string {
	if c.A == 255 {
		return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
	}
	alpha := strconv.FormatFloat(float64(c.A)/255, 'f', 3, 64)
	alpha = strings.TrimRight(strings.TrimRight(alpha, "0"), ".")
	return fmt.Sprintf("rgba(%d, %d, %d, %s)", c.R, c.G, c.B, alpha)
}

var (
	hexColourRe = regexp.MustCompile(`^#([0-9a-fA-F]{3,8})$`)
	rgbaRe      = regexp.MustCompile(`^rgba?\(\s*(\d{1,3})\s*[, ]\s*(\d{1,3})\s*[, ]\s*(\d{1,3})\s*(?:[,/]\s*([0-9.]+%?)\s*)?\)$`)
)

// ParseColour reads a colour written as #rgb, #rgba, #rrggbb, #rrggbbaa,
// rgb() or rgba().
func ParseColour(css string) (RGBA, bool) {
	css = strings.TrimSpace(css)
	if m := hexColourRe.FindStringSubmatch(css); m != nil {
		return parseHexColour(m[1])
	}
	if m := rgbaRe.FindStringSubmatch(css); m != nil {
		return parseRGBColour(m[1:4], m[4])
	}
	return RGBA{}, false
}

func parseHexColour(h string) (RGBA, bool) {
	if len(h) == 3 || len(h) == 4 {
		var wide []byte
		for i := range len(h) {
			wide = append(wide, h[i], h[i])
		}
		h = string(wide)
	}
	if len(h) == 6 {
		h += "ff"
	}
	if len(h) != 8 {
		return RGBA{}, false
	}
	v, err := strconv.ParseUint(h, 16, 32)
	if err != nil {
		return RGBA{}, false
	}
	return RGBA{uint8(v >> 24), uint8(v >> 16), uint8(v >> 8), uint8(v)}, true
}

// alpha is "" for an opaque colour, else a fraction or a percentage.
func parseRGBColour(channels []string, alpha string) (RGBA, bool) {
	c := RGBA{A: 255}
	for i, dst := range []*uint8{&c.R, &c.G, &c.B} {
		v, _ := strconv.Atoi(channels[i])
		if v > 255 {
			return RGBA{}, false
		}
		*dst = uint8(v)
	}
	if alpha == "" {
		return c, true
	}
	number, percent := strings.CutSuffix(alpha, "%")
	a, err := strconv.ParseFloat(number, 64)
	if percent {
		a /= 100
	}
	if err != nil || a < 0 || a > 1 {
		return RGBA{}, false
	}
	c.A = uint8(math.Round(a * 255))
	return c, true
}

// PaletteTokens is the palette vocabulary, in the order overrides count by.
func PaletteTokens() []string {
	for _, p := range Parts {
		if p.Name == "palette" {
			return p.Tokens
		}
	}
	return nil
}

func IsPaletteToken(name string) bool {
	return slices.Contains(PaletteTokens(), name)
}

var colourEncoding = base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567").WithPadding(base32.NoPadding)

// EncodeColours writes colour overrides, by token name, as the text an ID
// carries. Tokens outside the palette are ignored. It is "" for none.
func EncodeColours(overrides map[string]RGBA) string {
	var raw []byte
	for i, token := range PaletteTokens() {
		c, ok := overrides[token]
		if !ok {
			continue
		}
		// The top bit of the position says that an alpha byte follows.
		if c.A == 255 {
			raw = append(raw, byte(i), c.R, c.G, c.B)
		} else {
			raw = append(raw, byte(i)|0x80, c.R, c.G, c.B, c.A)
		}
	}
	return colourEncoding.EncodeToString(raw)
}

// DecodeColours reads what EncodeColours wrote. It is false for text that
// is not the one way to write those overrides, so that a style has one ID.
func DecodeColours(text string) (map[string]RGBA, bool) {
	if text == "" {
		return nil, true
	}
	raw, err := colourEncoding.DecodeString(text)
	if err != nil {
		return nil, false
	}
	tokens := PaletteTokens()
	out := map[string]RGBA{}
	last := -1
	for len(raw) > 0 {
		i := int(raw[0] & 0x7f)
		size := 4
		if raw[0]&0x80 != 0 {
			size = 5
		}
		if len(raw) < size || i >= len(tokens) || i <= last {
			return nil, false
		}
		c := RGBA{raw[1], raw[2], raw[3], 255}
		if size == 5 {
			c.A = raw[4]
			if c.A == 255 {
				return nil, false
			}
		}
		out[tokens[i]] = c
		last = i
		raw = raw[size:]
	}
	if EncodeColours(out) != text {
		return nil, false
	}
	return out, true
}
