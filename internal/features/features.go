// Package features turns a page measured by internal/browser into numbers
// that can be compared: a feature vector for "how far is this page from
// that one", and the checks behind a pack's Never rules.
package features

import (
	"encoding/json"
	"fmt"
	"math"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

// Metrics mirrors the `metrics` object of extract.js.
type Metrics struct {
	ViewportWidth       float64 `json:"viewportWidth"`
	BorderRadiusMax     float64 `json:"borderRadiusMax"`
	BorderRadiusMin     float64 `json:"borderRadiusMin"` // smallest radius among rounded boxes
	BorderRadiusMedian  float64 `json:"borderRadiusMedian"`
	RoundedBoxesPct     float64 `json:"roundedBoxesPct"`
	RoundShapesPct      float64 `json:"roundShapesPct"`
	BorderWidthMax      float64 `json:"borderWidthMax"`
	BorderWidthMin      float64 `json:"borderWidthMin"` // thinnest border among bordered boxes
	BorderWidthMedian   float64 `json:"borderWidthMedian"`
	BorderedBoxesPct    float64 `json:"borderedBoxesPct"`
	BoxShadowPct        float64 `json:"boxShadowPct"`
	BoxShadowBlurMax    float64 `json:"boxShadowBlurMax"`
	BoxShadowOffsetMax  float64 `json:"boxShadowOffsetMax"`
	TextShadowPct       float64 `json:"textShadowPct"`
	GradientFillsPct    float64 `json:"gradientFillsPct"`
	FontSizeMin         float64 `json:"fontSizeMin"`
	FontSizeMax         float64 `json:"fontSizeMax"`
	FontSizeBase        float64 `json:"fontSizeBase"`
	FontWeightMin       float64 `json:"fontWeightMin"`
	FontWeightMax       float64 `json:"fontWeightMax"`
	BoldTextPct         float64 `json:"boldTextPct"`
	FontFamilies        float64 `json:"fontFamilies"`
	FontFamilyMain      string  `json:"fontFamilyMain"`
	LineHeight          float64 `json:"lineHeight"` // 0 when the page leaves it at "normal"
	UnderlinedLinksPct  float64 `json:"underlinedLinksPct"`
	LetterSpacingMax    float64 `json:"letterSpacingMax"`
	UppercaseTextPct    float64 `json:"uppercaseTextPct"`
	RowGapMax           float64 `json:"rowGapMax"`
	RowGapMedian        float64 `json:"rowGapMedian"`
	RowGapLow           float64 `json:"rowGapLow"`  // 10th percentile
	RowGapHigh          float64 `json:"rowGapHigh"` // 90th percentile
	BlockGapMax         float64 `json:"blockGapMax"`
	BlockGapLow         float64 `json:"blockGapLow"`
	BlockGapHigh        float64 `json:"blockGapHigh"`
	BlockGapMedian      float64 `json:"blockGapMedian"`
	ContentWidth        float64 `json:"contentWidth"`
	ContentWidthPct     float64 `json:"contentWidthPct"`
	PaletteColours      float64 `json:"paletteColours"`
	Transitions         float64 `json:"transitions"`
	Animations          float64 `json:"animations"`
	TextDensity         float64 `json:"textDensity"`
	BackdropBlurPct     float64 `json:"backdropBlurPct"`
	TranslucentFillsPct float64 `json:"translucentFillsPct"`
}

// Hints mirrors `archetypeHints` of extract.js: counts that point at the page type.
type Hints struct {
	NumericRows      float64 `json:"numericRows"`
	ForumWords       float64 `json:"forumWords"`
	PriceCount       float64 `json:"priceCount"`
	CartWords        float64 `json:"cartWords"`
	BlogWords        float64 `json:"blogWords"`
	DateCount        float64 `json:"dateCount"`
	LinkCount        float64 `json:"linkCount"`
	LinkTextPct      float64 `json:"linkTextPct"`
	ParagraphTextPct float64 `json:"paragraphTextPct"`
	CodeBlocks       float64 `json:"codeBlocks"`
	NavListDepth     float64 `json:"navListDepth"`
	FormControls     float64 `json:"formControls"`
	PasswordInputs   float64 `json:"passwordInputs"`
	Buttons          float64 `json:"buttons"`
	SignupWords      float64 `json:"signupWords"`
}

// Archetype guesses the page type from DOM signals. The order of the
// rules matters: the more specific page types are tested first.
func (h Hints) Archetype() string {
	switch {
	case h.NumericRows >= 5 && h.ForumWords >= 5:
		return "forum"
	case h.PriceCount >= 6 && (h.CartWords >= 1 || h.PriceCount >= 12):
		return "shop"
	case h.CodeBlocks >= 4 || (h.NavListDepth >= 2 && h.ParagraphTextPct >= 30):
		return "docs"
	case h.BlogWords >= 4 && h.DateCount >= 2 && h.ParagraphTextPct >= 20:
		return "blog"
	case h.LinkTextPct >= 50 && h.LinkCount >= 80:
		return "portal"
	case h.FormControls >= 8 && h.SignupWords < 3:
		return "app"
	}
	return "landing"
}

type weighted struct {
	Value  string  `json:"value"`
	Weight float64 `json:"weight"`
}

// Page is the part of extract.js output the vector is built from.
type Page struct {
	Title             string     `json:"title"`
	URL               string     `json:"url"`
	TextSample        string     `json:"textSample"`
	TextChars         int        `json:"textChars"`
	BodyBackground    string     `json:"bodyBackground"`
	BackgroundsByArea []weighted `json:"backgroundsByArea"`
	TextColours       []weighted `json:"textColours"`
	Links             []weighted `json:"links"`
	Hints             Hints      `json:"archetypeHints"`
	Metrics           Metrics    `json:"metrics"`
}

func Parse(raw json.RawMessage) (*Page, error) {
	var p Page
	if err := json.Unmarshal(raw, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

// Dimension is one coordinate of the feature vector. Scale is the
// difference that counts as "completely different" for that coordinate.
type Dimension struct {
	Name  string
	Unit  string
	Scale float64
}

var Dimensions = []Dimension{
	{"page background lightness", "", 50},
	{"page background green-red", "", 60},
	{"page background blue-yellow", "", 60},
	{"main surface lightness", "", 50},
	{"main surface green-red", "", 60},
	{"main surface blue-yellow", "", 60},
	{"text lightness", "", 50},
	{"text green-red", "", 60},
	{"text blue-yellow", "", 60},
	{"link lightness", "", 50},
	{"link green-red", "", 60},
	{"link blue-yellow", "", 60},
	{"text contrast with surface", "", 60},
	{"colourfulness of fills", "", 50},
	{"distinct colours", "", 20},
	{"gradient fills", "%", 20},
	{"base font size", "px", 8},
	{"largest font size", "px", 48},
	{"heading scale", "x", 4},
	{"heaviest font weight", "", 500},
	{"bold text", "%", 40},
	{"font families", "", 3},
	{"serif body", "", 1},
	{"line height", "", 0.6},
	{"underlined links", "%", 100},
	{"letter spacing", "px", 3},
	{"uppercase text", "%", 30},
	{"text shadow", "%", 30},
	{"bordered boxes", "%", 30},
	{"border width, median", "px", 2},
	{"border width, largest", "px", 4},
	{"rounded boxes", "%", 30},
	{"corner radius, median", "px", 12},
	{"corner radius, largest", "px", 24},
	{"pill and circle shapes", "%", 10},
	{"boxes with shadow", "%", 20},
	{"shadow blur, largest", "px", 30},
	{"shadow offset, largest", "px", 12},
	{"row gap, median", "px", 16},
	{"block gap, median", "px", 40},
	{"content width", "% of viewport", 60},
	{"text density", "chars per 1000px²", 4},
	{"transitions", "", 1},
	{"backdrop blur", "%", 10},
	{"translucent fills", "%", 40},
}

// Vector holds one value per Dimension. NaN means "not measured" and is
// left out of comparisons.
type Vector []float64

var serifFamilies = []string{
	"times", "times new roman", "georgia", "serif", "palatino",
	"garamond", "book antiqua", "cambria", "charter",
}

// noColour is lab's answer for a colour it cannot read or that is see-through.
var noColour = [3]float64{math.NaN(), math.NaN(), math.NaN()}

func (p *Page) Vector() Vector {
	m := p.Metrics
	page := lab(p.BodyBackground)
	if math.IsNaN(page[0]) {
		// A transparent body shows the browser's white.
		page = [3]float64{100, 0, 0}
	}
	surface := p.largestOpaqueFill(page)
	text := firstColour(p.TextColours)
	link := firstColour(p.Links)
	lineHeight := m.LineHeight
	if lineHeight == 0 {
		lineHeight = math.NaN()
	}
	scale := math.NaN()
	if m.FontSizeBase > 0 {
		scale = m.FontSizeMax / m.FontSizeBase
	}
	return Vector{
		page[0], page[1], page[2],
		surface[0], surface[1], surface[2],
		text[0], text[1], text[2],
		link[0], link[1], link[2],
		math.Abs(text[0] - surface[0]),
		p.meanFillChroma(),
		m.PaletteColours,
		m.GradientFillsPct,
		m.FontSizeBase, m.FontSizeMax, scale,
		m.FontWeightMax, m.BoldTextPct, m.FontFamilies,
		flag(slices.Contains(serifFamilies, m.FontFamilyMain)),
		lineHeight,
		m.UnderlinedLinksPct, m.LetterSpacingMax, m.UppercaseTextPct, m.TextShadowPct,
		m.BorderedBoxesPct, m.BorderWidthMedian, m.BorderWidthMax,
		m.RoundedBoxesPct, m.BorderRadiusMedian, m.BorderRadiusMax, m.RoundShapesPct,
		m.BoxShadowPct, m.BoxShadowBlurMax, m.BoxShadowOffsetMax,
		m.RowGapMedian, m.BlockGapMedian,
		// Off-screen carousels can push text past the viewport.
		math.Min(100, m.ContentWidthPct), m.TextDensity,
		flag(m.Transitions > 0),
		m.BackdropBlurPct, m.TranslucentFillsPct,
	}
}

func (p *Page) largestOpaqueFill(fallback [3]float64) [3]float64 {
	for _, b := range p.BackgroundsByArea {
		if c := lab(b.Value); !math.IsNaN(c[0]) {
			return c
		}
	}
	return fallback
}

func (p *Page) meanFillChroma() float64 {
	var chroma, area float64
	for _, b := range p.BackgroundsByArea {
		if c := lab(b.Value); !math.IsNaN(c[0]) {
			chroma += math.Hypot(c[1], c[2]) * b.Weight
			area += b.Weight
		}
	}
	if area == 0 {
		return 0
	}
	return chroma / area
}

func firstColour(colours []weighted) [3]float64 {
	if len(colours) == 0 {
		return noColour
	}
	return lab(colours[0].Value)
}

func flag(b bool) float64 {
	if b {
		return 1
	}
	return 0
}

type Delta struct {
	Dimension Dimension
	Measured  float64
	Reference float64
	Share     float64 // 0 to 1: how much of the scale the difference covers
}

func (d Delta) String() string {
	unit := d.Dimension.unitSuffix()
	return fmt.Sprintf("%s measured %s%s, reference %s%s",
		d.Dimension.Name, num(d.Measured), unit, num(d.Reference), unit)
}

func (d Dimension) unitSuffix() string {
	if d.Unit == "" || d.Unit == "%" || d.Unit == "x" {
		return d.Unit
	}
	return " " + d.Unit
}

func num(f float64) string {
	return strconv.FormatFloat(math.Round(f*10)/10, 'f', -1, 64)
}

// Compare returns the distance between a measured vector and a reference,
// from 0 (the same) to 1 (different on every dimension), and the
// per-dimension differences, largest first.
func Compare(measured, reference Vector) (distance float64, deltas []Delta) {
	var sum float64
	var n int
	for i, dim := range Dimensions {
		a, b := measured[i], reference[i]
		if math.IsNaN(a) || math.IsNaN(b) {
			continue
		}
		share := math.Min(1, math.Abs(a-b)/dim.Scale)
		sum += share * share
		n++
		deltas = append(deltas, Delta{dim, a, b, share})
	}
	sort.SliceStable(deltas, func(i, j int) bool { return deltas[i].Share > deltas[j].Share })
	if n == 0 {
		return 0, nil
	}
	return math.Sqrt(sum / float64(n)), deltas
}

var rgbRe = regexp.MustCompile(`^rgba?\(\s*([\d.]+)[,\s]+([\d.]+)[,\s]+([\d.]+)(?:\s*[,/]\s*([\d.]+)(%?))?`)

// lab converts a computed CSS colour to CIE Lab (D65). Colours it cannot
// read come back as NaN.
func lab(css string) [3]float64 {
	g := rgbRe.FindStringSubmatch(strings.TrimSpace(css))
	if g == nil {
		return noColour
	}
	// A mostly transparent colour is not the colour that is seen.
	if g[4] != "" {
		alpha, _ := strconv.ParseFloat(g[4], 64)
		if g[5] == "%" {
			alpha /= 100
		}
		if alpha < 0.5 {
			return noColour
		}
	}
	var lin [3]float64
	for i := range 3 {
		c, _ := strconv.ParseFloat(g[i+1], 64)
		c /= 255
		if c <= 0.04045 {
			lin[i] = c / 12.92
		} else {
			lin[i] = math.Pow((c+0.055)/1.055, 2.4)
		}
	}
	x := (0.4124*lin[0] + 0.3576*lin[1] + 0.1805*lin[2]) / 0.95047
	y := 0.2126*lin[0] + 0.7152*lin[1] + 0.0722*lin[2]
	z := (0.0193*lin[0] + 0.1192*lin[1] + 0.9505*lin[2]) / 1.08883
	f := func(t float64) float64 {
		if t > 0.008856 {
			return math.Cbrt(t)
		}
		return 7.787*t + 16.0/116
	}
	return [3]float64{116*f(y) - 16, 500 * (f(x) - f(y)), 200 * (f(y) - f(z))}
}

// RangeDelta is how far a measurement lies outside the range its references span.
type RangeDelta struct {
	Dimension Dimension
	Measured  float64
	Low, High float64 // the range the references span
	Share     float64 // 0 inside the range, up to 1
}

func (d RangeDelta) String() string {
	unit := d.Dimension.unitSuffix()
	return fmt.Sprintf("%s measured %s%s, references range from %s%s to %s%s",
		d.Dimension.Name, num(d.Measured), unit, num(d.Low), unit, num(d.High), unit)
}

// CompareRange measures how far a vector lies outside the range a group of
// references spans, dimension by dimension. A blend of several references
// should sit inside what they show; it need not sit at their average,
// which for a group of light and dark pages is a grey that none of them is.
func CompareRange(measured Vector, references []Vector) (distance float64, deltas []RangeDelta) {
	var sum float64
	var n int
	for i, dim := range Dimensions {
		low, high := math.Inf(1), math.Inf(-1)
		for _, r := range references {
			if !math.IsNaN(r[i]) {
				low = math.Min(low, r[i])
				high = math.Max(high, r[i])
			}
		}
		if math.IsNaN(measured[i]) || math.IsInf(low, 1) {
			continue
		}
		var out float64
		if measured[i] < low {
			out = low - measured[i]
		} else if measured[i] > high {
			out = measured[i] - high
		}
		share := math.Min(1, out/dim.Scale)
		sum += share * share
		n++
		if share > 0 {
			deltas = append(deltas, RangeDelta{dim, measured[i], low, high, share})
		}
	}
	sort.SliceStable(deltas, func(i, j int) bool { return deltas[i].Share > deltas[j].Share })
	if n == 0 {
		return 0, nil
	}
	return math.Sqrt(sum / float64(n)), deltas
}
