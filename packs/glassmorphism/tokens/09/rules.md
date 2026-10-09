## Colour roles

- Backdrop: `--color-page` is a warm near-black (18 16 16). `--color-fill-1` (64 0 249) is the one glow colour; `--color-fill-2` (110 72 255), `--color-fill-3` (120 64 184) and `--color-fill-4` (207 64 255) are its neighbours, used only as tile tints and in the hero orbs.
- Glass is faint: `--color-surface` white at 10%, `--color-surface-alt` 15%, `--color-surface-strong` 25%. Borders are white at 15% (`--color-border`), 30% (`--color-border-strong`) and 6% (`--color-border-muted`).
- Dark glass: `--color-bar` (28 28 28 at 75%) for the navigation bar, table head and code blocks, `--color-bar-alt` (40 39 38 at 60%) for tabs, footer and status badges, `--color-input` (the page colour at 60%) for form controls.
- Text is white; secondary text is not grey but pale lavender: `--color-text-muted` and `--color-heading-alt` are 216 205 249.
- `--color-button` is that lavender as a solid, with dark blue text (42 51 90). `--color-accent` is lime green (130 211 73) with dark green text (42 57 19); `--color-link` is the same lime. `--color-accent-alt` is violet (110 72 255) with white text.
- Status: `--color-danger` 242 84 91, `--color-success` 61 193 60, `--color-warning` 255 162 44. `--color-notice` is lime at 10% with near-white text.
- `--color-shadow` is black at 50%; `--color-focus` is lavender.

## Typography roles

- One family for everything, "Helvetica Neue" then Helvetica and Arial; `--font-mono` for code.
- The scale is small and dense: `--text-base` and `--text-ui` 14px, `--text-small` 12px, `--text-large` 18px, `--text-h3` 16px, `--text-h2` 22px, `--text-h1` 28px, `--text-display` 56px.
- Line height is tight: `--line-body` 1.375, `--line-heading` 1.3, `--line-display` 1.07.
- Weights: `--weight-body` 400, `--weight-ui` 500, `--weight-heading` and `--weight-display` 600, `--weight-bold` 700.
- No uppercase and no tracking from tokens; links are marked by colour and underlined on hover only.

## Surface roles

- `--fill-page` is the page colour with one glow colour, `--color-fill-1`, as large round glows (680px to 760px radius) alternating right and left down the page; between them the backdrop stays near-black.
- Corners are generous: `--radius-panel` 16px, `--radius-page` and `--radius-pill` 22px, `--radius-control` 8px.
- `--border-width` 1px; `--border-width-strong` 2px.
- Glass is tinted and lit from its edge, not lifted: `--fill-panel` lays a diagonal wash of `--color-accent` (22%) into `--color-fill-1` (26%) over `--color-surface`, and `--shadow-panel` and `--shadow-dialog` are a 1px ring of the accent at 45% with a glow of the same colour, 28px to 30px wide, on every side. Controls glow too: `--shadow-control` is a 12px halo of the accent, 20px when hovered. Nothing casts a dark drop shadow.
- `--shadow-text` is a wide soft 0 4px 10px on the brand, the hero title and large figures.
- Buttons and inputs are flat; the secondary button carries the same accent wash as the glass. Transitions run 0.3s.
- `--backdrop-blur` is 18px: a firm frost that turns the glow behind a pane into an even wash of colour. The strongest glass of a layout (bars, dialogs) blurs three times that, 54px.

## Never

- `border-width <= 2px`: borders are 1px; 2px is the strong width.
- `border-radius <= 22px`: panels are 16px, tiles and dialogs 22px.
- `box-shadow-blur <= 30px`: the widest glow, on the dialog, spreads 30px.
- `font-size <= 56px`: the hero title is the largest text.
- `font-size >= 12px`: meta text and badges are the smallest text.
- `font-weight <= 700`: headings are semibold; 700 is for bold only.
- `font-families <= 2`: one sans and one monospace.
- `line-height <= 1.4`: body text is set at 1.375.
- `letter-spacing = 0`: no tracking on any text.
- `underlined-links <= 25%`: underline is a hover state.
- `uppercase-text <= 10%`: uppercase is for small labels only.
- `gradient-fills <= 60%`: glass carries one diagonal tint; controls, bars and rows stay flat.
