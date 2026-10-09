package features

import (
	"fmt"
	"math"

	"github.com/chronoskin/chronoskin/internal/pack"
)

type NeverResult struct {
	Rule     pack.NeverRule
	Measured string
	Pass     bool
	Skipped  string // why the rule could not be checked, if it could not
}

func CheckNever(rules []pack.NeverRule, m Metrics) []NeverResult {
	out := make([]NeverResult, len(rules))
	for i, r := range rules {
		out[i] = check(r, m)
	}
	return out
}

func check(r pack.NeverRule, m Metrics) NeverResult {
	res := NeverResult{Rule: r}
	want, unit, numeric := r.Number()

	if amount, ok := presence(r.Metric, m); ok {
		res.Measured, res.Pass = checkPresence(r.Op, amount > 0, !numeric || want == 0)
		return res
	}
	if !numeric {
		res.Skipped = "value must be a number for this metric"
		return res
	}
	sp, ok := measuredSpan(r.Metric, unit, m)
	if !ok {
		res.Skipped = "no measurement for this metric"
		return res
	}
	if r.Metric == "line-height" && m.LineHeight == 0 {
		res.Skipped = "the page leaves line-height at normal"
		return res
	}

	slack := tolerance(r.Metric, want)
	switch r.Op {
	case "<=":
		res.Measured = num(sp.hi) + unit
		res.Pass = sp.hi <= want+slack
	case ">=":
		res.Measured = num(sp.lo) + unit
		res.Pass = sp.lo >= want-slack
	case "=":
		res.Measured = num(sp.hi) + unit
		if sp.lo != sp.hi {
			res.Measured = fmt.Sprintf("%s to %s%s", num(sp.lo), num(sp.hi), unit)
		}
		res.Pass = math.Max(math.Abs(sp.lo-want), math.Abs(sp.hi-want)) <= slack
	case "!=":
		res.Measured = num(sp.hi) + unit
		res.Pass = math.Abs(sp.hi-want) > slack
	}
	return res
}

// presence is the amount of a metric that a rule can only forbid or require.
func presence(metric string, m Metrics) (amount float64, ok bool) {
	switch metric {
	case "box-shadow":
		return m.BoxShadowPct, true
	case "text-shadow":
		return m.TextShadowPct, true
	case "transition":
		return m.Transitions, true
	case "animation":
		return m.Animations, true
	}
	return 0, false
}

// "=" and "<=" name the state that must never occur.
func checkPresence(op string, present, wantNone bool) (measured string, pass bool) {
	measured = "none"
	if present {
		measured = "present"
	}
	if op == "=" || op == "<=" {
		return measured, present != wantNone
	}
	return measured, present == wantNone
}

// span is the smallest and largest measurement of a metric: an upper bound
// is checked against the largest, a lower bound against the smallest.
type span struct{ lo, hi float64 }

func single(v float64) span { return span{v, v} }

func measuredSpan(metric, unit string, m Metrics) (span, bool) {
	switch metric {
	case "border-radius":
		return span{m.BorderRadiusMin, m.BorderRadiusMax}, true
	case "border-width":
		return span{m.BorderWidthMin, m.BorderWidthMax}, true
	case "box-shadow-blur":
		return span{0, m.BoxShadowBlurMax}, true
	case "gradient-fills":
		return single(m.GradientFillsPct), true
	case "font-size":
		return span{m.FontSizeMin, m.FontSizeMax}, true
	case "font-weight":
		return span{m.FontWeightMin, m.FontWeightMax}, true
	case "font-families":
		return single(m.FontFamilies), true
	case "line-height":
		return single(m.LineHeight), true
	case "underlined-links":
		return single(m.UnderlinedLinksPct), true
	case "letter-spacing":
		return span{0, m.LetterSpacingMax}, true
	case "uppercase-text":
		return single(m.UppercaseTextPct), true
	// Gaps use the 10th and 90th percentiles: one stray margin on a page
	// should not decide the rule.
	case "row-gap":
		return span{m.RowGapLow, m.RowGapHigh}, true
	case "block-gap":
		return span{m.BlockGapLow, m.BlockGapHigh}, true
	case "content-width":
		if unit == "%" {
			return single(m.ContentWidthPct), true
		}
		return single(m.ContentWidth), true
	case "palette-colours":
		return single(m.PaletteColours), true
	}
	return span{}, false
}

// tolerance is the measuring error a rule allows: half a unit of rounding
// for pixels and percentages, or 2% for large values such as widths; ratios
// and counts are held much tighter.
func tolerance(metric string, want float64) float64 {
	switch metric {
	case "line-height":
		return 0.05
	case "letter-spacing":
		return 0.1
	case "font-families", "palette-colours", "font-weight":
		return 0
	}
	return math.Max(0.5, math.Abs(want)*0.02)
}
