## Colour roles

- Black on white with the browser's own blue: `--color-page`, `--color-canvas` and `--color-surface` are #ffffff; `--color-text`, `--color-heading`, `--color-heading-alt` and every border are #000000; `--color-text-muted` is #444444.
- `--color-link`, `--color-link-quiet`, `--color-link-visited` and `--color-focus` are #0000ee, the default link colour: on a text-only front page nearly everything is a blue underlined link and visited links are not told apart. `--color-link-hover` and `--color-link-active` are #cc0000.
- `--color-bar` #000000 with white `--color-bar-text` is the black line of navigation, table heads and dialog titles. `--color-bar-alt` (footer) is white with black text. `--color-inverse` #0000ee with white text is the current navigation item.
- The only tints are the faint blues that a blue link leaves around itself: `--color-surface-alt` #f8f8ff, `--color-surface-strong`, `--color-notice` and `--color-disabled` #f0f0ff, `--color-fill-4` and `--color-overlay` #d8d8ff, `--color-border-muted` #b8b8f8. `--color-fill-1` to `--color-fill-3` repeat white and the two faintest.
- `--color-accent` #cc0000 with white text is the count badge, the one red on the page; `--color-accent-alt` is the same red, as a mark.
- Buttons are white with a black edge: the primary has black text, the secondary blue text. Inputs are white with a black edge.
- Status: `--color-danger` #cc0000 is measured; `--color-success` #006400 and `--color-warning` #8a4b00 are derived, dark enough to read on white.

## Typography roles

- The system's own sans: `--font-body`, `--font-heading` and `--font-ui` are -apple-system, system-ui, "Segoe UI", Roboto, Helvetica, Arial, sans-serif; `--font-mono` is Consolas, Monaco, monospace.
- `--text-base` and `--text-ui` 16px with `--line-body` 1.4; `--text-small` 14px; `--text-large` 18px. Headings are bold: `--text-display` 36px, `--text-h1` 24px, `--text-h2` 21px, `--text-h3` 16px. Navigation and buttons are regular weight (`--weight-ui` 400).
- Every link is underlined at rest and on hover; nothing is uppercased or tracked.

## Surface roles

- One weight for everything: `--border-width` and `--border-width-strong` are both 2px solid, so cells, boxes, controls and the rules under the header are the same black line.
- Boxes are framed twice: `--shadow-panel` draws a second 2px line in `--color-border` 3px outside a panel, and `--shadow-dialog` a 4px band of `--color-border-strong` around a dialog, both without blur. Controls are flat; a hovered button gets a 2px line along its foot and a pressed one a 2px inner frame.
- The page is inverted: `--fill-page` is `--color-inverse`, so whatever of the window lies outside the sheet is a solid block of the inverse colour. Every other fill is its plain palette colour.
- All radii are 0; `--focus-ring` is a 2px solid outline; `--transition` none.

## Never

- `border-radius <= 0px`: every box and control is square.
- `box-shadow-blur <= 0px`: shadows only draw second frames; nothing is soft.
- `text-shadow = none`: text is flat.
- `gradient-fills <= 0%`: every fill is one flat colour.
- `letter-spacing <= 0px`: no text is tracked.
- `transition = none`: states change instantly.
- `animation = none`: nothing moves.
- `border-width <= 2px`: every rule is 2px.
- `font-size <= 36px`: the hero headline is 36px.
- `font-size >= 14px`: small text is 14px.
- `font-families <= 2`: the system sans and a monospace.
- `line-height >= 1.3`: text is set at 1.4.
- `line-height <= 1.5`: and no looser.
- `uppercase-text <= 0%`: nothing is uppercased.
- `underlined-links >= 90%`: every link is underlined at rest.
