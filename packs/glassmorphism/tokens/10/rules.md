## Colour roles

- Backdrop: `--color-page` is a warm paper white (252 249 247) under a pastel mesh: `--color-fill-1` periwinkle (184 202 245), `--color-fill-2` mint (155 216 169), `--color-fill-3` butter (253 233 155) and `--color-fill-4` lilac (226 223 254) as wide soft ovals that overlap down the page. The same four tint the tiles.
- This palette is light. Glass is milky white: `--color-surface` at 50%, `--color-surface-alt` 68%, `--color-surface-strong` 85%, edged with white at 70% (`--color-border`). `--color-border-strong` is solid lavender (171 159 242); `--color-border-muted` is ink at 9% for hairlines between rows.
- Text is plum ink, not black: `--color-text`, `--color-heading-alt` and `--color-link-quiet` are 60 49 91, `--color-text-muted` the same at 72%. Only `--color-heading` is near-black (28 28 28).
- Bars are pale too: `--color-bar` (233 232 234 at 60%) and `--color-bar-alt` (253 252 254 at 60%) carry plum text. The dark block of the palette is `--color-inverse` (plum) with cream text.
- `--color-button` is lavender (171 159 242) with near-black text; the secondary button is white glass with plum text.
- `--color-accent`, `--color-link` and `--color-focus` are one electric blue (49 57 251) with cream text on it; `--color-accent-alt` is the mint, with plum text on it.
- Status: `--color-danger` 214 69 65 on a pale rose `--color-danger-surface`, `--color-success` 46 139 87, `--color-warning` 224 142 30. `--color-notice` is butter at 70% with near-black text.
- `--color-shadow` is plum at 14%, never grey; `--color-overlay` is plum at 40%.

## Typography roles

- Two families: headings and the hero title are a text serif (`--font-heading`: New York, Iowan Old Style, Source Serif 4, Georgia) at regular weight; body and controls are the system sans. The references used commercial serif and sans faces; these are the closest system and open stacks.
- `--weight-heading` and `--weight-display` are 400: the serif does the work, not the weight. `--weight-ui` 500, `--weight-bold` 600.
- Sizes: `--text-base` 16px, `--text-ui` 15px, `--text-small` 13px, `--text-large` 20px, `--text-h3` 18px, `--text-h2` 24px, `--text-h1` 32px, `--text-display` 60px.
- Line height: `--line-body` 1.4, `--line-heading` 1.2, `--line-display` 1.1.
- No uppercase and no tracking from tokens; links are marked by colour and underlined on hover only.

## Surface roles

- `--fill-page` is the paper colour with the four pastels as overlapping ovals 620px to 700px wide, alternating left and right; no part of the page is left plain for long.
- Corners are large and soft: `--radius-control` 14px, `--radius-panel` 24px, `--radius-page` and `--radius-pill` 32px.
- Glass is a thick, polished pebble: `--fill-panel` adds a diagonal sheen (white at 28% in the top left corner, gone by the middle) over `--color-surface`, and `--shadow-panel` and `--shadow-dialog` start with an inset 1px highlight along the top and left edges (white at 50%).
- Shadows are faint and tinted: `--shadow-panel` 0 2px 8px under the highlight, `--shadow-control` a 4px halo, `--shadow-control-hover` 0 5px 5px, `--shadow-dialog` 0 20px 40px, all in plum at 14%.
- `--shadow-text` is `none`: dark text on light glass takes no shadow.
- `--backdrop-blur` is 32px, the heaviest of the era: the pastel ovals behind a pane melt into one milky tint and no edge of them shows through. The strongest glass of a layout blurs three times that, 96px.
- Buttons carry a soft gloss: `--fill-button` and `--fill-button-secondary` fade from white at 28% or 22% at the top to nothing. Transitions run 0.2s.

## Never

- `border-width <= 2px`: borders are 1px; 2px is the strong width.
- `border-radius <= 32px`: controls are 14px, panels 24px, large containers and badges 32px.
- `box-shadow-blur <= 40px`: only the dialog shadow reaches 40px.
- `text-shadow = none`: no text shadow anywhere.
- `font-size <= 60px`: the hero title is the largest text.
- `font-size >= 13px`: meta text and badges are the smallest text.
- `font-weight <= 600`: headings are regular weight; 600 is for bold only.
- `font-families <= 3`: one serif for headings, one sans, one monospace.
- `line-height <= 1.5`: body text is set at 1.4.
- `letter-spacing = 0`: no tracking on any text.
- `underlined-links <= 25%`: underline is a hover state.
- `uppercase-text <= 10%`: uppercase is for small labels only.
- `gradient-fills <= 70%`: glass and buttons carry one soft sheen each; rows, bars and inputs stay flat.
