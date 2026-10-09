## Colour roles

- Backdrop: `--color-page` is a blue-black (7 8 10) with a fine grid of faint dots and one ember glow: `--color-fill-1` (150 28 36), a dark crimson, falling from the top centre and returning at the edges further down. `--color-fill-2` (graphite, 67 67 69), `--color-fill-3` (62 43 104) and `--color-fill-4` (80 70 228) are used only as tile tints.
- Glass is almost clear: `--color-surface` white at 5%, `--color-surface-alt` 9%, `--color-surface-strong` 16%. The edge does the work: `--color-border` white at 12%, `--color-border-strong` white at 40% for dialogs, notices and secondary buttons, `--color-border-muted` 6% for hairlines.
- Dark glass: `--color-bar` (black at 60%) for bars, table heads and code blocks, `--color-bar-alt` (17 18 20 at 70%) for tabs, footer and status badges, `--color-input` (17 18 20 at 80%) for form controls.
- Text is white; `--color-text-muted` is neutral grey (156 156 157) with no colour cast, and `--color-heading-alt` is light grey (230 230 230).
- `--color-button` is a light grey solid (230 230 230) with near-black text, the brightest thing on the page. The secondary button is clear glass with the bright edge.
- `--color-accent`, `--color-link` and `--color-focus` are sky blue (86 194 255) with near-black text on it. `--color-accent-alt` is indigo (80 70 228) with white text.
- Status: `--color-danger` is the coral red of the glow family (255 99 99), `--color-success` 89 212 153, `--color-warning` 255 197 49. `--color-notice` is the accent at 10% with light grey text.
- `--color-shadow` and `--color-overlay` are black at 60%.

## Typography roles

- Two families: Inter for body and headings, and a monospace (`--font-ui`: Geist Mono, JetBrains Mono, then the system monospace) for buttons, inputs, navigation, tabs and table cells as well as code.
- Sizes: `--text-base` 16px, `--text-ui` and `--text-small` 13px, `--text-large` 18px, `--text-h3` 18px, `--text-h2` 24px, `--text-h1` 36px, `--text-display` 56px.
- Body is airy and headings are light: `--line-body` 1.6, `--line-heading` 1.33, `--line-display` 1.14; `--weight-heading`, `--weight-display` and `--weight-ui` 500, `--weight-bold` 600.
- No uppercase and no tracking from tokens; links are marked by colour and underlined on hover only.

## Surface roles

- `--fill-page` layers a 24px grid of 1px dots over the ember glow (an oval 760px wide at the top centre, 620px ovals at alternating edges below) on the page colour.
- Corners are tight: `--radius-control` 3px, `--radius-panel` 6px, `--radius-page` 10px; only `--radius-pill` (30px) stays round.
- `--border-width` and `--border-width-strong` are both 2px: the outline is the surface. Glass is smoked, `--fill-panel` being `--color-surface` mixed with 35% black, so a panel is a darker pane inside a pale 2px frame.
- Nothing floats: `--shadow-panel` and `--shadow-control` are `none`. Only the dialog casts a shadow (`--shadow-dialog`, 0 24px 48px in black). A hovered button gains a 2px bright ring, and `--focus-ring` is a 2px dashed line.
- `--shadow-text` is `none`.
- `--backdrop-blur` is 3px, the lightest of the era: the dot grid and the glow stay readable through a pane, only softened. The strongest glass of a layout blurs three times that, 9px.
- Every fill is flat; transitions run 0.15s.

## Never

- `border-width <= 2px`: every border is 2px, never heavier.
- `border-radius <= 30px`: panels are 6px, large containers 10px, pills 30px.
- `box-shadow-blur <= 48px`: only the dialog casts a shadow, blurred 48px.
- `text-shadow = none`: no text shadow anywhere.
- `font-size <= 56px`: the hero title is the largest text.
- `font-size >= 13px`: controls and meta text are the smallest text.
- `font-weight <= 600`: headings are medium; 600 is for bold only.
- `font-families <= 2`: one sans and one monospace.
- `line-height <= 1.6`: body text is set at 1.6.
- `letter-spacing = 0`: no tracking on any text.
- `underlined-links <= 25%`: underline is a hover state.
- `uppercase-text <= 10%`: uppercase is for small labels only.
- `gradient-fills <= 10%`: glass is a flat translucent fill.
