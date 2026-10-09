## Colour roles

- A white page with two loud bands. `--color-page`, `--color-canvas` and `--color-surface` are #ffffff; `--color-surface-alt`, `--color-surface-strong`, `--color-notice` and `--color-overlay` are the one grey, #efefef.
- `--color-bar` #6487dc (cornflower blue) with white `--color-bar-text` is the line of navigation, table heads and dialog titles. `--color-bar-alt` #ff6d06 (orange) with black text is the footer: the second band. `--color-inverse` #000000 with white text is the current navigation item.
- Text is #000000; `--color-text-muted` is #555555. Headings are not black: `--color-heading` and `--color-heading-alt` are #0066cc, the colour of the links, and a heading is told from a link by its size and the rule under it.
- `--color-link`, `--color-link-quiet` and `--color-link-visited` are #0066cc; `--color-link-hover` is #000000 and `--color-link-active` #ff6d06.
- Rules are blue: `--color-border` #6487dc for cells, boxes and row dividers, `--color-border-strong` #0066cc for the rule under the header, hero and page header and around notices and dialogs. `--color-border-muted` is #dddddd.
- `--color-accent` #ff6d06 with black text is the count badge; `--color-accent-alt` is the same orange, as a mark only.
- Fills alternate #ffffff and #efefef; there are no tinted blocks. Buttons are #efefef with black text and white with blue text; inputs are white with a #767676 edge.
- Status: `--color-danger` #c00000, `--color-success` #1e6e00, `--color-warning` #a04400. None was measured; they are dark enough to read on white and on #efefef.

## Typography roles

- One small sans: `--font-body`, `--font-heading` and `--font-ui` are Helvetica, Verdana, Arial, "Liberation Sans", sans-serif; `--font-mono` is Monaco, "Liberation Mono", monospace.
- `--text-base` 13px with `--line-body` 1.3; `--text-small` 11px; `--text-ui` 12px, so navigation and controls are a size under the text; `--text-large` 14px. Headings are bold: `--text-display` 27px, `--text-h1` 20px, `--text-h2` 16px, `--text-h3` 13px.
- No link is underlined at rest (`--link-decoration` and `--link-decoration-quiet` are none): links are known by their blue. They underline on hover. Nothing is uppercased or tracked.

## Surface roles

- Dashed rules: `--border-width` 1px for cells, boxes, row dividers and control edges, `--border-width-strong` 3px for the rule under the header, hero and page header and the outline of notices and dialogs, all in `--border-style` dashed, like a line to cut along.
- The page is a desk: `--fill-page` is `--color-surface-strong`, so a sheet narrower than the window shows grey beside it; `--fill-panel` is `--color-surface-alt`. Every other fill is its plain palette colour, except that a hovered button darkens a little.
- All radii are 0. Nothing casts a shadow at rest; a hovered button thickens its edge by 1px outside and a pressed one by 1px inside, without blur. `--focus-ring` is a 1px dashed outline; `--transition` none.

## Never

- `border-radius <= 0px`: every box and control is square.
- `box-shadow-blur <= 0px`: the only shadow is the hard extra edge of a hovered or pressed button.
- `text-shadow = none`: text is flat.
- `gradient-fills <= 0%`: every fill is one flat colour.
- `letter-spacing <= 0px`: no text is tracked.
- `transition = none`: states change instantly.
- `animation = none`: nothing moves.
- `border-width <= 3px`: rules are 1px, emphasised rules 3px, both dashed.
- `font-size <= 27px`: the hero headline is 27px.
- `font-size >= 11px`: small text is 11px.
- `font-families <= 2`: one sans and a monospace.
- `line-height <= 1.4`: text is set at 1.3.
- `uppercase-text <= 0%`: nothing is uppercased.
- `underlined-links <= 10%`: links are blue, not underlined; only a hovered link is.
