## Colour roles

- Neutral grey cells with navy bars and rust links, the stock look of the most widely installed commercial board of the period. `--color-page` and `--color-canvas` are both white: the board has no sheet and stands straight on the page.
- Cells: `--color-surface` #eeeeee for the main cell, `--color-surface-alt` #dadada for poster panels, status, count and last-post cells and form labels. `--color-surface-strong` is the pale blue #a2c0df of category rows and title strips, the one cool surface in the set.
- `--color-bar` is the saturated navy #0b198c with white `--color-bar-text`; `--color-bar-alt` the grey violet #666686 with white `--color-bar-alt-text`. The footer band is `--color-inverse` #666686 with white text.
- Text is `--color-text` #000000 and `--color-text-muted` #4a4a4a. `--color-heading` is the navy #0b198c, which also reads on the pale blue strips; `--color-heading-alt` is black.
- Links are rust: `--color-link` #b75a3e, `--color-link-quiet` the darker sienna #a0522d for titles and names, `--color-link-visited` #666686. A hovered link turns `--color-link-hover` #ff0000.
- `--color-border` is the mid grey #999999; `--color-border-strong` the navy #0b198c, used for outer and emphasised edges and button outlines; `--color-border-muted` #dadada. Inputs use `--color-input-border` #767676.
- Markers: `--color-accent` #b75a3e with white text for the badge and new-post icons; `--color-accent-alt` #22229c for a few emphasised words. `--color-danger` #a62a2a on `--color-danger-surface` #fddbcc, `--color-success` #008000, `--color-warning` #ffa34f, `--color-notice` the pale yellow #ffffdd.
- The primary button is the navy `--color-button` with white text; the secondary button is the system grey #efefef.
- `--color-fill-1` to `--color-fill-4` repeat the greys (#eeeeee, #dadada, #efefef, #eeeeee).

## Typography roles

- Verdana for text and headings (`--font-body`, `--font-heading`); `--font-ui` puts Tahoma first for navigation, tabs, buttons and inputs.
- `--text-base`, `--text-ui`, `--text-h2` and `--text-h3` are 11px, `--text-small` 10px, `--text-large` 13px for titles in lists.
- `--text-display` and `--text-h1` are 18px, and `--weight-display` is 400: the site name is large and not bold, as poster names and page titles were in this skin. `--weight-heading` stays 700, so the 18px page heading is bold. Nothing is larger than 18px.
- Line height is `normal`. No tracking, no uppercase. `--weight-ui` is 700.
- No link is underlined at rest: `--link-decoration` and `--link-decoration-quiet` are `none`, `--link-decoration-hover` is `underline`.

## Surface roles

- Every fill is flat: `--fill-bar`, `--fill-bar-alt`, `--fill-inverse`, `--fill-accent` and `--fill-button` are the plain colours. `--fill-button-hover` lightens the button by 18% white.
- `--fill-panel` is `--color-canvas`, so the 1px gaps between cells are white.
- Buttons are raised by `--shadow-control`, a hard 1px bevel: a light inner edge at top and left, a dark one at bottom and right. `--shadow-control-pressed` reverses it. There is no other shadow.
- Borders are `--border-width` 1px solid, `--border-width-strong` 2px. All `--radius-*` are 0. `--shadow-text` is `none`.
- `--focus-ring` is a 1px dotted line; `--transition` is `none`.

## Never

- `border-radius <= 0px`: every box is square.
- `box-shadow-blur <= 0px`: the only shadow is the hard bevel of a button.
- `text-shadow = none`: no text shadow.
- `gradient-fills <= 0%`: all fills are flat.
- `transition = none`: states switch at once.
- `animation = none`: nothing moves.
- `border-width <= 2px`: 1px, with 2px on emphasised edges.
- `font-size <= 18px`: the site name and page heading are the largest text.
- `font-size >= 10px`: 10px is the smallest size.
- `font-weight <= 700`: bold is the heaviest weight.
- `font-families <= 3`: Verdana, Tahoma for controls, a monospace for code.
- `letter-spacing <= 0px`: no tracking.
- `uppercase-text <= 0%`: nothing is transformed to uppercase.
