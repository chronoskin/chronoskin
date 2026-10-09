## Colour roles

- A royal blue gaming skin: `--color-page` is the saturated blue #17479e and the board is a light sheet on it, `--color-canvas` #ececec.
- Cells are pale steel blue: `--color-surface` #c6d2de for the main cell, `--color-surface-alt` #bacbdc for poster panels, status, count and last-post cells and form labels, `--color-surface-strong` #98b5e2 for category rows and title strips.
- `--color-bar` is the page blue #17479e with white `--color-bar-text`. `--color-bar-alt` is the bright sky blue #40a2e5 and takes black `--color-bar-alt-text`, because white does not read on it. The footer band is `--color-inverse` #184078 with white text.
- Text is `--color-text` #000000; `--color-text-muted` is the dark blue #183870, so it only reads on the light surfaces. `--color-heading` is #000099, `--color-heading-alt` the dark blue #183870.
- `--color-link` is the deep blue #000099; `--color-link-quiet` is black, for titles and names; `--color-link-visited` #184078; a hovered link turns `--color-link-hover` #2060a0.
- `--color-border` is the light blue #98cafd, which also fills the 1px gaps between cells; `--color-border-strong` the page blue #17479e; `--color-border-muted` #98b5e2. Inputs use `--color-input-border` #767676.
- Markers: `--color-accent` #2060a0 with white text for the badge and new-post icons; `--color-accent-alt` is the green #008000. `--color-danger` #a62a2a on `--color-danger-surface` #fddbcc, `--color-success` #008000, `--color-warning` #ffa34f. `--color-notice` is white, so a notice stands out from the grey sheet.
- The primary button is the page blue with white text and a pale blue hover text; the secondary button is the sheet grey #ececec.
- `--color-fill-1` to `--color-fill-4` are #c6d2de, #bacbdc, #ececec and #98b5e2.

## Typography roles

- Arial for everything: `--font-body`, `--font-heading` and `--font-ui` are the same Arial stack with Verdana second.
- `--text-base`, `--text-large`, `--text-h2` and `--text-h3` are 13px; `--text-small` and `--text-ui` 11px; `--text-h1` 16px; `--text-display` 18px. Hierarchy comes from weight, not size. Nothing is larger than 18px.
- Line height is `normal`. No tracking, no uppercase. `--weight-ui` is 700.
- Every link is underlined, at rest and hovered: `--link-decoration`, `--link-decoration-quiet` and `--link-decoration-hover` are all `underline`.

## Surface roles

- `--fill-bar` is a vertical gradient that starts 30% darker than the bar colour, passes the bar colour at 40% and ends 30% lighter at the foot: the bars glow from below.
- `--fill-button` runs from a 30% lighter top to the button colour; `--fill-button-hover` from the button colour to 25% darker.
- `--fill-panel` is `--color-border`: the 1px gaps between cells are light blue lines, so every cell looks outlined.
- `--fill-bar-alt`, `--fill-inverse` and `--fill-accent` are flat.
- Borders are `--border-width` 1px solid, `--border-width-strong` 2px. All `--radius-*` are 0, every `--shadow-*` is `none`.
- `--focus-ring` is a 1px dotted line; `--transition` is `none`.

## Never

- `border-radius <= 0px`: every box is square.
- `box-shadow = none`: no shadow.
- `text-shadow = none`: no text shadow.
- `gradient-fills <= 20%`: gradients are on primary bars and the primary button only.
- `transition = none`: states switch at once.
- `animation = none`: nothing moves.
- `border-width <= 2px`: 1px, with 2px on emphasised edges.
- `font-size <= 18px`: the site name is the largest text.
- `font-size >= 11px`: 11px is the smallest size.
- `font-weight <= 700`: bold is the heaviest weight.
- `font-families <= 2`: Arial and a monospace for code.
- `letter-spacing <= 0px`: no tracking.
- `uppercase-text <= 0%`: nothing is transformed to uppercase.
- `underlined-links >= 50%`: links are underlined, titles included.
