package browser

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"testing"
)

func TestPixelColours(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 40, 40))
	for y := range 40 {
		for x := range 40 {
			c := color.RGBA{0, 102, 153, 255}
			if y < 10 {
				c = color.RGBA{255, 255, 255, 255}
			}
			img.Set(x, y, c)
		}
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		t.Fatal(err)
	}
	got, err := PixelColours(buf.Bytes(), 4)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0].Hex != "#006898" || got[0].Share != 75 || got[1].Hex != "#ffffff" {
		t.Errorf("got %+v", got)
	}
}
