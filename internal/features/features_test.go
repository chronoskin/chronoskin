package features

import (
	"math"
	"strings"
	"testing"

	"github.com/chronoskin/chronoskin/internal/pack"
)

func TestLab(t *testing.T) {
	white, black := lab("rgb(255, 255, 255)"), lab("rgba(0, 0, 0, 0.5)")
	if math.Abs(white[0]-100) > 0.1 || math.Abs(black[0]) > 0.1 {
		t.Errorf("white L = %v, black L = %v", white[0], black[0])
	}
	if red := lab("rgb(255, 0, 0)"); red[1] < 70 {
		t.Errorf("red a* = %v", red[1])
	}
	if !math.IsNaN(lab("color(srgb 1 0 0)")[0]) {
		t.Error("unreadable colour should be NaN")
	}
	if !math.IsNaN(lab("rgba(0, 0, 0, 0)")[0]) || !math.IsNaN(lab("rgb(255 255 255 / 20%)")[0]) {
		t.Error("a transparent colour should be NaN, not black or white")
	}
	transparent := &Page{BodyBackground: "rgba(0, 0, 0, 0)", BackgroundsByArea: []weighted{{"rgba(255, 255, 255, 0)", 9}, {"rgb(0, 0, 0)", 5}}}
	if v := transparent.Vector(); v[0] != 100 || v[3] > 1 {
		t.Errorf("transparent body: page L %v (want 100), surface L %v (want the opaque black)", v[0], v[3])
	}
}

func TestCompare(t *testing.T) {
	a := &Page{BodyBackground: "rgb(229, 229, 229)", Metrics: Metrics{FontSizeBase: 11, FontSizeMax: 12, BorderWidthMax: 1}}
	b := &Page{BodyBackground: "rgb(255, 255, 255)", Metrics: Metrics{FontSizeBase: 16, FontSizeMax: 48, BorderRadiusMax: 12, BoxShadowBlurMax: 24, BoxShadowPct: 10}}
	if d, _ := Compare(a.Vector(), a.Vector()); d != 0 {
		t.Errorf("distance to itself = %v", d)
	}
	d, deltas := Compare(b.Vector(), a.Vector())
	if d <= 0.1 || d > 1 {
		t.Errorf("distance = %v", d)
	}
	if deltas[0].Share < deltas[len(deltas)-1].Share {
		t.Error("deltas are not sorted largest first")
	}
}

func TestCompareRange(t *testing.T) {
	light := (&Page{BodyBackground: "rgb(255, 255, 255)", Metrics: Metrics{FontSizeBase: 12, BorderRadiusMax: 0}}).Vector()
	dark := (&Page{BodyBackground: "rgb(0, 0, 0)", Metrics: Metrics{FontSizeBase: 16, BorderRadiusMax: 8}}).Vector()
	// A dark page with 14px text lies between the two on every axis.
	inside := (&Page{BodyBackground: "rgb(10, 10, 10)", Metrics: Metrics{FontSizeBase: 14, BorderRadiusMax: 4}}).Vector()
	if d, deltas := CompareRange(inside, []Vector{light, dark}); d != 0 || len(deltas) != 0 {
		t.Errorf("inside the range: distance %v, %v", d, deltas)
	}
	outside := (&Page{BodyBackground: "rgb(10, 10, 10)", Metrics: Metrics{FontSizeBase: 24, BorderRadiusMax: 30}}).Vector()
	d, deltas := CompareRange(outside, []Vector{light, dark})
	if d <= 0 || len(deltas) != 2 || !strings.Contains(deltas[0].String(), "references range from") {
		t.Errorf("outside the range: distance %v, %v", d, deltas)
	}
}

func TestCheckNever(t *testing.T) {
	rules, _ := pack.NeverRules("- `border-radius <= 0px`: a\n- `box-shadow = none`: b\n- `font-size <= 16px`: c\n" +
		"- `font-size >= 10px`: d\n- `line-height <= 1.3`: e\n- `content-width >= 90%`: f\n- `transition = none`: g")
	m := Metrics{BorderRadiusMax: 8, BoxShadowPct: 0, FontSizeMin: 10, FontSizeMax: 16, ContentWidthPct: 96, Transitions: 3}
	want := []string{"fail", "pass", "pass", "pass", "skip", "pass", "fail"}
	for i, r := range CheckNever(rules, m) {
		got := "fail"
		if r.Skipped != "" {
			got = "skip"
		} else if r.Pass {
			got = "pass"
		}
		if got != want[i] {
			t.Errorf("%s: %s (measured %s), want %s", r.Rule.Metric, got, r.Measured, want[i])
		}
	}
}
