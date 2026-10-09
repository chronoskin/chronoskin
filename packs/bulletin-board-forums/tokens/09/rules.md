## Colour roles

- A steel blue-grey family on near-white. `--color-page` is the pale #e9f0f6, `--color-canvas` white. Cells are almost neutral: `--color-surface` #f6f6f6 for the main cell, `--color-surface-alt` #ecedf3 for status, count and last-post cells and form labels, `--color-surface-strong` #e0e1e8 for category rows and title strips.
- `--color-bar` is the mid steel #476c8e with white `--color-bar-text`; `--color-bar-alt` the light steel #88a6c0 with black `--color-bar-alt-text` (white does not read on it). The footer band is `--color-inverse` #375576 with `--color-inverse-text` #e9f0f6.
- Text is `--color-text` #000000 and `--color-text-muted` the grey #696969. `--color-heading` is the dark steel #375576, `--color-heading-alt` #476c8e.
- `--color-link` is #476c8e, `--color-link-quiet` the darker #375576, `--color-link-visited` #587898; a hovered link turns the strong blue `--color-link-hover` #0022bb.
- `--color-border` and `--color-border-strong` are both the dark steel #375576, so every block has a dark outline; `--color-border-muted` (#b8c6d3) is for dividers inside a box. Inputs use `--color-input-border` #696969.
- Markers: `--color-accent` #375576 with white text for the badge and new-post icons; `--color-accent-alt` is the violet #920ac4, for a few emphasised words only. `--color-danger` #a00000 on `--color-danger-surface` #fddbcc, `--color-success` #006600, `--color-warning` #ffa34f, `--color-notice` #e9f0f6.
- `--color-fill-1` to `--color-fill-4` repeat the pale surfaces (#f6f6f6, #e0e1e8, #ecedf3, #e9f0f6).

## Typography roles

- Text is small and of one size: `--text-base`, `--text-small` and `--text-ui` are all 10px Verdana (`--font-body`, `--font-ui`). Hierarchy comes from weight and from one jump to `--text-large` / `--text-h2` 13px bold for board titles and bar titles.
- Headings use `--font-heading`, Arial with Geneva and Helvetica, bold: `--text-h1` and `--text-display` 16px, `--text-h2` 13px, `--text-h3` 11px. Nothing is larger than 16px.
- Line height is `normal`. No tracking, no uppercase. `--weight-ui` is 700.
- Every link is underlined, at rest and hovered: `--link-decoration`, `--link-decoration-quiet` and `--link-decoration-hover` are all `underline`, including bold board and topic titles.

## Surface roles

- `--fill-bar` is a two-band gloss with a hard break at half height: the upper band runs from the bar colour with 14% white down to the plain bar colour, the lower band from 18% to 30% black. `--fill-button` is the same build on the button colour; `--fill-button-hover` flattens both bands.
- `--fill-bar-alt`, `--fill-inverse` and `--fill-accent` are flat.
- `--fill-panel` is `--color-border-muted`: the one-pixel gaps between cells are a faint rule, so pale cells stay separated without a dark grid.
- Borders are `--border-width` 1px solid, `--border-width-strong` 2px under the navigation bar and tabs and around the dialog. All `--radius-*` are 0, every `--shadow-*` is `none`.
- `--focus-ring` is a 1px dotted line; `--transition` is `none`.

## Never

- `border-radius <= 0px`: every box is square.
- `box-shadow = none`: no shadow in any reference.
- `text-shadow = none`: no text shadow in any reference.
- `gradient-fills <= 15%`: the gloss is on primary bars and primary buttons only.
- `transition = none`: states switch at once.
- `animation = none`: nothing moves.
- `border-width <= 2px`: 1px, with 2px on emphasised edges.
- `font-size <= 16px`: the wordmark and page heading are the largest text.
- `font-size >= 10px`: 10px is the smallest and by far the most common size.
- `font-weight <= 700`: bold is the heaviest weight.
- `font-families <= 3`: Verdana, Arial for headings, a monospace for code.
- `letter-spacing <= 0px`: no tracking.
- `uppercase-text <= 0%`: nothing is transformed to uppercase.
- `underlined-links >= 50%`: links are underlined, titles included.
