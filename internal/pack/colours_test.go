package pack

import "testing"

func TestParseColour(t *testing.T) {
	for in, want := range map[string]string{
		"#2b55e0": "#2b55e0", "#FFF": "#ffffff", "#0008": "rgba(0, 0, 0, 0.533)", "#11223380": "rgba(17, 34, 51, 0.502)",
		"rgb(1, 2, 3)": "#010203", "rgba(1,2,3,0.5)": "rgba(1, 2, 3, 0.502)", "rgb(1 2 3 / 50%)": "rgba(1, 2, 3, 0.502)", "rgba(0, 0, 0, 1)": "#000000",
	} {
		c, ok := ParseColour(in)
		if !ok || c.CSS() != want {
			t.Errorf("ParseColour(%q) = %q, %v; want %q", in, c.CSS(), ok, want)
		}
		// What CSS writes must read back to the same colour.
		if back, ok := ParseColour(c.CSS()); !ok || back != c {
			t.Errorf("%q does not read back", c.CSS())
		}
	}
	for _, bad := range []string{"", "red", "#12", "#12345", "rgb(256, 0, 0)", "rgba(0,0,0,2)", "linear-gradient(#000, #fff)", "var(--x)"} {
		if _, ok := ParseColour(bad); ok {
			t.Errorf("ParseColour(%q) succeeded", bad)
		}
	}
}

func TestColoursRoundTrip(t *testing.T) {
	tokens := PaletteTokens()
	in := map[string]RGBA{tokens[0]: {1, 2, 3, 255}, tokens[len(tokens)-1]: {4, 5, 6, 7}, "--not-a-token": {9, 9, 9, 255}}
	text := EncodeColours(in)
	out, ok := DecodeColours(text)
	if !ok || len(out) != 2 || out[tokens[0]] != in[tokens[0]] || out[tokens[len(tokens)-1]] != in[tokens[len(tokens)-1]] {
		t.Fatalf("round trip of %q gave %v, %v", text, out, ok)
	}
	if EncodeColours(nil) != "" {
		t.Error("no overrides must encode to nothing")
	}
	// One style, one ID: text that decodes but is not the canonical form is refused.
	for _, bad := range []string{"a", "zzzz", text + "a", text[:len(text)-1]} {
		if _, ok := DecodeColours(bad); ok {
			t.Errorf("DecodeColours(%q) succeeded", bad)
		}
	}
	id := ID{Prefix: "v1", Era: "flat-design", Archetype: "landing", Structure: 2, Palette: 3, Type: 1, Surface: 4, Density: "roomy", Colours: text}
	if back, ok := ParseID(id.String()); !ok || back != id {
		t.Errorf("ParseID(%q) = %v, %v", id.String(), back, ok)
	}
	short, _ := id.Short("fl")
	if parsed, ok := ParseShortID(short); !ok || parsed.Colours != text || parsed.Density != "roomy" {
		t.Errorf("ParseShortID(%q) = %v, %v", short, parsed, ok)
	}
}
