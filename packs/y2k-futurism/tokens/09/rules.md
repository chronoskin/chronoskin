## Colour roles

- White with pale blue bands and one strong blue. `--color-page`, `--color-canvas` and `--color-surface` are all #ffffff: the sheet is marked only by its outline. `--color-surface-alt` #f0f0f0 fills forms, the link column and alternate rows; `--color-surface-strong` #d0e0e8 is title strips inside the link column.
- `--color-bar` #9cc6e1 with black text is the navigation bar, panel and dialog titles, table heads and the current tab; `--color-bar-alt` #d8e0e8 with black text is the utility strip, section tabs, inactive tabs and the footer strip. `--color-inverse` #282828 with white text is the one dark panel and the closing band.
- `--color-text` and `--color-heading` #000000; `--color-text-muted` #606060; `--color-heading-alt` #0033cc.
- `--color-link` and `--color-link-quiet` #0033cc; visited #606060, hover #000000, active #282828. `--color-accent` #0033cc with white text is the current navigation capsule, badges and the current page number; `--color-accent-alt` #9cc6e1 is a lamp colour.
- `--color-fill-1` to `--color-fill-4` are pale blue-greys and near-whites (#d0e0e8, #f8f8ff, #d8e0e8, #f0f0f0). Lines are grey: `--color-border` #c0c0c0, `--color-border-strong` #a0a0a0, `--color-border-muted` #d8e0e8.
- `--color-button` #ffffff with black text, outlined; `--color-button-secondary` #f0f0f0. Inputs are white with a #a0a0a0 line.
- `--color-danger` #800000 is only used as text on `--color-danger-surface` #fffff0 and as a rim or lamp; `--color-success` #709000 and `--color-warning` #ffc800 are lamp and rim colours, never text. `--color-notice` #f8f8ff with black text.

## Typography roles

- Small humanist sans-serif for text, a light serif for headings. `--font-body` and `--font-ui` are Lucida Grande, then Geneva, Verdana; `--font-heading` is Garamond, then Times New Roman, at `--weight-heading` 400. The references drew their headlines in a commercial condensed Garamond; Garamond or Times at regular weight is the stand-in.
- `--text-base` 10px, `--text-small` 9px, `--text-ui` 11px, `--text-large` 12px. `--text-h3` 12px, `--text-h2` 15px, `--text-h1` and `--text-display` 18px.
- Interface text is regular weight (`--weight-ui` 400) and not uppercased; nothing is tracked.
- Every link is underlined, quiet and navigation links included.

## Surface roles

- Flat fills and outlined capsules. Every `--fill-*` token is the plain colour except `--fill-page`, a faint 2px pinstripe; there is no gloss and no bevel.
- `--radius-control` 12px makes buttons and inputs capsules; `--radius-pill` 10px for badges and lamps; `--radius-panel` and `--radius-page` 8px for panels, tabs, the sheet, hero and grid cells.
- `--border-width` and `--border-width-strong` are both 1px, solid: shapes are drawn by their outline.
- `--shadow-dialog` is the only resting shadow, 1px down with a 2px blur; `--shadow-control-pressed` is a soft inset. All other shadows and `--shadow-text` are `none`. `--focus-ring` is 1px solid; `--transition` none.

## Never

- `palette-colours <= 18`: white, greys, pale blues, one strong blue, plus status lamps.
- `font-size <= 18px`: headings stop at 18px.
- `font-size >= 9px`: nothing is set below 9px.
- `font-weight <= 700`: headings are regular weight; bold is for labels and list titles.
- `font-families <= 3`: a sans-serif, a serif for headings and a monospace.
- `letter-spacing <= 0px`: no text is tracked.
- `uppercase-text <= 0%`: nothing is uppercased.
- `underlined-links >= 85%`: every link is underlined.
- `border-radius <= 12px`: 8px on boxes; controls and badges are capsules.
- `border-width <= 1px`: every line is a hairline.
- `box-shadow-blur <= 2px`: only the dialog has a shadow, and it is tight.
- `text-shadow = none`: text is flat.
- `gradient-fills <= 2%`: only the page pinstripe is a gradient.
- `transition = none`: states change instantly.
- `animation = none`: nothing moves.
