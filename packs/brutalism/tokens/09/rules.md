## Colour roles

- Grey page, white boxes, black ink and one electric blue. `--color-page` and `--color-canvas` are #eeeeee; `--color-surface` #ffffff is table cells, panels and dialogs; `--color-surface-alt` #eeeeee; `--color-surface-strong` #e0e0e0.
- `--color-bar` #000000 with white `--color-bar-text` is the black line of navigation links, table heads and dialog titles. `--color-bar-alt` #ffffff with black text is the footer. `--color-inverse` #0000ff with white text is the current navigation item.
- `--color-text`, `--color-heading` and every border (`--color-border`, `--color-border-strong`, `--color-input-border`) are #000000; `--color-text-muted`, `--color-link-visited` and `--color-border-muted` are #606060.
- Blue #0000ff is `--color-link`, `--color-link-quiet`, `--color-link-active`, `--color-heading-alt`, `--color-accent` (white text), `--color-accent-alt` and `--color-focus`. `--color-link-hover` is #000000.
- The primary button is black with white text; the secondary is white with black text. Fills alternate #ffffff and #e0e0e0.
- Status: `--color-danger` #cc0000, `--color-success` #006600, `--color-warning` #7a5200. The reference has no status colours; these are dark enough to read as text on #eeeeee and on white and are not measured.

## Typography roles

- Monospace throughout: all four font tokens are "Courier New", Courier, monospace.
- One size for everything that is not a headline: `--text-base`, `--text-small`, `--text-ui`, `--text-large`, `--text-h2` and `--text-h3` are 13px, `--line-body` 1.3. `--text-h1` is 24px and `--text-display` 48px at `--line-display` 1.
- Headings are bold capitals (`--heading-transform` uppercase), which is how a 13px heading is told from 13px text. Navigation and buttons are bold (`--weight-ui` 700) and keep their case.
- Every link is underlined at rest (`--link-decoration` and `--link-decoration-quiet` underline) and loses the underline on hover (`--link-decoration-hover` none).

## Surface roles

- Heavy rules everywhere: `--border-width` 3px solid for table cells, boxes, row dividers and control edges; `--border-width-strong` 6px for the rules under the header, hero and page header and the outline of notices and dialogs.
- Boxes stand off the page on hard offset shadows in `--color-shadow`, with no blur: `--shadow-panel` 4px down and right, `--shadow-control` 2px (4px on hover, turned inward when pressed), `--shadow-dialog` 8px.
- All radii are 0, every fill is its plain palette colour; `--focus-ring` is a 3px solid outline; `--transition` none.

## Never

- `border-radius <= 0px`: every box and control is square.
- `box-shadow-blur <= 0px`: shadows are hard offset blocks; nothing is soft.
- `text-shadow = none`: text is flat.
- `gradient-fills <= 0%`: every fill is one flat colour.
- `letter-spacing <= 0px`: no text is tracked.
- `transition = none`: states change instantly.
- `animation = none`: nothing moves.
- `border-width <= 6px`: rules are 3px, emphasised rules 6px.
- `font-size <= 48px`: the hero headline is 48px.
- `font-size >= 13px`: everything else is 13px.
- `font-families <= 1`: one monospace for text, headings, controls and code.
- `line-height <= 1.4`: text is set at 1.3.
- `uppercase-text <= 15%`: only headings are capitals.
- `underlined-links >= 90%`: every link is underlined at rest; only a hovered link is bare.
