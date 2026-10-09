package browser

import (
	"bytes"
	"fmt"
	"image/png"
	"sort"
)

type PixelColour struct {
	Hex   string  `json:"hex"`
	Share float64 `json:"sharePct"`
}

// PixelColours returns the n most common colours of a PNG. Unlike computed
// styles it sees what the page drew with images: gradients, bevels, buttons.
// Channels are rounded to steps of 8 so that anti-aliasing noise groups up.
func PixelColours(data []byte, n int) ([]PixelColour, error) {
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	counts := map[[3]uint8]int{}
	total := 0
	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y += 2 {
		for x := b.Min.X; x < b.Max.X; x += 2 {
			r, g, bl, _ := img.At(x, y).RGBA()
			counts[[3]uint8{round8(r), round8(g), round8(bl)}]++
			total++
		}
	}
	keys := make([][3]uint8, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return counts[keys[i]] > counts[keys[j]] })
	var out []PixelColour
	for _, k := range keys[:min(n, len(keys))] {
		share := float64(counts[k]*1000/total) / 10
		out = append(out, PixelColour{fmt.Sprintf("#%02x%02x%02x", k[0], k[1], k[2]), share})
	}
	return out, nil
}

func round8(c uint32) uint8 {
	v := (c>>8 + 4) / 8 * 8
	return uint8(min(v, 255))
}
