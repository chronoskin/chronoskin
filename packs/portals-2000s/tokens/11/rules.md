## Colour roles

- Black, orange and sand. `--color-page` #f7eac7 (sand) surrounds a white `--color-canvas` and white `--color-surface`. `--color-surface-alt` #f3f3f3 fills sidebars, forms and alternate rows; `--color-surface-strong` #cccc99 (khaki) is sidebar title strips and inactive tabs.
- `--color-bar` #000000 with white text is the link bar, panel titles, table heads and the current tab; `--color-bar-alt` #ff6600 with white text is the utility strip, footer strip and alternate panel titles. `--color-inverse` #000066 with white text is the one dark block.
- `--color-text` #000000; `--color-text-muted` #666666; `--color-heading` #333333; `--color-heading-alt` #cc3300 (burnt orange) for section titles and h3.
- `--color-link` #0000cc, `--color-link-quiet` #000099, visited #333366, hover #ff6600, active #ff0000. `--color-accent` #ffcc00 with black text is the badge, the current navigation item and links on the dark block; `--color-accent-alt` #ff0000 marks promoted entries.
- Block fills are warm and pale: `--color-fill-1` #f7eac7, `--color-fill-2` #f3f3f3, `--color-fill-3` #ffffcc, `--color-fill-4` #eeeeee. Rules `--color-border` #cccccc, outlines `--color-border-strong` #808080.
- `--color-button` #cccccc, secondary #f3f3f3, black labels. `--color-danger` #cc0000 on `--color-danger-surface` #ffffcc. The palette has no green: `--color-success` is the olive #999966; `--color-warning` is #ffcc00. Both are badge fills only.

## Typography roles

- Small Verdana with Arial headings: `--font-body` and `--font-ui` are Verdana at `--text-base` 11px and `--line-body` 1.4; `--font-heading` is Arial bold.
- `--text-small` and `--text-ui` 10px (bars, tabs and buttons are 10px bold); `--text-large` 12px; `--text-h3` 12px, `--text-h2` 13px, `--text-h1` 16px; `--text-display` 20px is used for the wordmark and the lead headline only.
- No uppercase and no tracking. No link is underlined until it is hovered (`--link-decoration` and `--link-decoration-quiet` are none).

## Surface roles

- `--fill-panel` shades from `--color-surface-alt` at the top edge to `--color-surface` 40px down, the grey fade the period drew under its box headings. `--fill-button` shades from a near-white mix to `--color-button`. `--fill-bar` and `--fill-bar-alt` carry a soft highlight that fades out 60% of the way down; `--fill-page` darkens toward the top of the window; `--fill-input` has a 4px shade under its top edge, like a sunken field.
- `--border-width` 1px, `--border-width-strong` 2px, solid; controls keep the 2px bevel. `--radius-control` 2px softens buttons and inputs; the other radii are 0. `--shadow-panel` is a faint 3px soft shadow to the lower right of each box and `--shadow-dialog` a deeper 8px one; controls and text cast none.
- `--focus-ring` is 1px dotted; `--transition` none.

## Never

- `palette-colours <= 22`: black, orange, navy, sand, khaki, yellow and greys; no green, no violet.
- `font-size <= 20px`: the largest text is 20px bold.
- `font-size >= 10px`: nothing is set below 10px.
- `font-weight <= 700`: regular and bold only.
- `font-families <= 3`: Verdana, Arial for headings, a monospace for code.
- `letter-spacing <= 0px`: no text is tracked.
- `uppercase-text <= 0%`: nothing is uppercased by CSS.
- `underlined-links <= 5%`: links are underlined only on hover.
- `border-radius <= 2px`: boxes are square and controls barely rounded.
- `border-width <= 2px`: 1px rules, 2px bevels on controls.
- `box-shadow-blur <= 8px`: shadows are small and faint; nothing floats.
- `text-shadow = none`: text is flat.
- `gradient-fills <= 35%`: shading is for the page, bars, panels, inputs and the primary button; cells and rows are flat.
- `transition = none`: states change instantly.
- `animation = none`: nothing moves.
