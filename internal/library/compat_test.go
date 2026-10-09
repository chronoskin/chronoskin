package library

import (
	"testing"
)

func palette(text, surface string) string {
	return ":root { --color-text: " + text + "; --color-surface: " + surface + "; --color-canvas: " + surface +
		"; --color-bar-text: #ffffff; --color-bar: #003366; --color-button-text: #ffffff; --color-button: #0066cc; --color-link: #0044aa; }"
}

func TestCompatible(t *testing.T) {
	palettes := map[int]string{
		1: palette("#000000", "#ffffff"),
		2: palette("#111111", "#f5f5f5"),
		3: palette("#222222", "#eeeeee"),
		4: palette("#1a1a1a", "#fafafa"),
		5: palette("#bbbbbb", "#cccccc"), // nearly no contrast between text and surface
		6: ":root { --color-text: var(--nope); }",
		7: palette("#0a0a0a", "#f0f0f0"),
		8: palette("#151515", "#ffffff"),
	}
	got := compatible(palettes)
	for _, n := range []int{1, 2, 3, 4, 7, 8} {
		if !got[n] {
			t.Errorf("palette %d should mix", n)
		}
	}
	if got[5] {
		t.Error("the low-contrast palette should stay with its own structure")
	}
	if got[6] {
		t.Error("an unreadable palette should stay with its own structure")
	}
	// The floor is fixed: a palette is judged alone, not against the others.
	if few := compatible(map[int]string{1: palettes[1], 5: palettes[5]}); !few[1] || few[5] {
		t.Errorf("two palettes: %v", few)
	}
	// White text on faintly white glass over a dark page is legible, and
	// must not be measured as white on white.
	glass := ":root { --color-page: #101828; --color-text: #ffffff; --color-surface: rgba(255, 255, 255, 0.2); --color-canvas: rgba(255, 255, 255, 0.1);" +
		" --color-bar-text: #ffffff; --color-bar: rgba(0, 0, 0, 0.4); --color-button-text: #101828; --color-button: #ffffff; --color-link: #9ad1ff; }"
	if got := compatible(map[int]string{1: glass}); !got[1] {
		t.Error("a glass palette on a dark page should mix")
	}
}
