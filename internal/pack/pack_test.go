package pack

import (
	"encoding/json"
	"strings"
	"testing"
)

// The example manifest, with the markup field.
const specExample = `{
  "id": "v1.bulletin-board-forums.forum.s17.p03.t11.u08",
  "library": "2026.10",
  "era": { "slug": "bulletin-board-forums", "years": [2002, 2007] },
  "archetype": "forum",
  "mode": "mix",
  "markup": "modern",
  "viewport": { "width": 1024, "fluid": false },
  "source": "https://example.test/s/v1.bulletin-board-forums.forum.s17.p03.t11.u08"
}`

func TestManifestDecodesSpecExample(t *testing.T) {
	var m Manifest
	dec := json.NewDecoder(strings.NewReader(specExample))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&m); err != nil {
		t.Fatal(err)
	}
	if m.Era.Years != [2]int{2002, 2007} || m.Viewport.Width != 1024 || m.Markup != "modern" {
		t.Errorf("unexpected manifest: %+v", m)
	}
}
