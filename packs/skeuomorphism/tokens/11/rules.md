## Colour roles

- Green baize, mahogany and brass; the set is dark. `--color-page` #1f5c3d is the felt table; `--color-canvas` #184a31 a sheet of darker felt; `--color-surface` #1b5137 rows and cells; `--color-surface-alt` #17472f alternate rows and wells; `--color-surface-strong` #123a26 title strips.
- `--color-bar` #5b2f18 (mahogany, ivory text #f8e9c6) is the navigation bar, table head, count badges and pressed "current" states; `--color-bar-alt` #8b6a3b (oak, ivory text #fff4d8) is title strips and the segmented control; `--color-inverse` #24160f (ebony, ivory text) the closing band or hero block.
- Text is `--color-text` #d8e5d6, meta `--color-text-muted` #9cbba5, headings ivory `--color-heading` #fdf6e0, sidebar and panel titles `--color-heading-alt` #f3e4b8.
- Links are gilt: `--color-link` #f2cf6b, paler on hover (#ffe7a0), tan when visited (#d8bb8c), white when active.
- `--color-button` #a87722 (brass, ivory text) is the one primary action colour; the secondary key is a lighter felt, `--color-button-secondary` #2c7850 with #f6f8ee text. Fields are wells darker than the felt (`--color-input` #0f3524, pale text).
- `--color-accent` #c0302b (the red chip) for badges, `--color-accent-alt` #2d5aa3 (the blue chip, ivory text) for "new" markers; `--color-success` #3f9d2f, `--color-warning` #b3741a (ivory text), `--color-danger` #f0695b on `--color-danger-surface` #4a2420. Notices are ivory cards lying on the felt: `--color-notice` #f3e6c0 with dark text #3a2a12.
- `--color-fill-1` to `--color-fill-4` (#b02a26 red, #2c58a0 blue, #5c3a7e violet, #7b4a22 walnut) are chip and wood colours for the grid blocks only; ivory text on all four.

## Typography roles

- One plain sans for everything: `--font-body`, `--font-heading` and `--font-ui` are Helvetica, then Arial, Lucida Grande and Verdana; `--font-mono` Monaco.
- Text is small and dense, pale on the dark ground: `--text-base` 11px at `--line-body` 1.42; `--text-small` 10px; `--text-ui` 11px at `--weight-ui` 700; `--text-large` 14px for the lead and the large key.
- Headings are bold: `--weight-heading` and `--weight-display` 700. `--text-h3` 14px, `--text-h2` 18px, `--text-h1` 22px, `--text-display` 30px at `--line-display` 1.1.
- Links are not underlined until hovered (`--link-decoration: none`, `--link-decoration-hover: underline`). No uppercase and no tracking from this part.

## Surface roles

- Felt and wood. `--fill-page` is felt: a 6px fleck of dark and light fibres and a pool of light from the top centre. `--fill-inverse` is the same fleck and light over ebony. `--fill-bar` and `--fill-bar-alt` are planks: uneven horizontal grain lines over a top-lit gradient of the bar colour.
- Keys are polished brass: `--fill-button` runs from 45% lighter at the top through the plain colour to 22% darker, with a bright inner top edge and a soft 0 3px 7px shadow (`--shadow-control`), so they sit padded above the felt. Pressed keys take a 5px inset shade.
- Sheets are padded cushions lying on the table: `--shadow-panel` is a dark 1px ring and a soft 0 6px 14px at 50%; `--shadow-dialog` a ring and 0 10px 20px. `--fill-panel` is flat felt.
- Corners are large and soft: `--radius-control` 10px, `--radius-panel` 16px, `--radius-page` 20px, `--radius-pill` 16px. `--fill-accent` is a chip with a highlight across its top.
- `--shadow-text` is a soft dark shadow one pixel below light text (0 1px 1px at 60%). `--fill-input` shades the top 6px of a well. `--transition` is none.

## Never

- `border-radius <= 20px`: the largest containers are 20px; nothing is rounder than the first structure's large key.
- `border-width <= 1px`: every edge is a hairline.
- `box-shadow-blur <= 20px`: the dialog's 20px shadow is the widest.
- `font-size <= 30px`: the hero headline is the largest text.
- `font-size >= 10px`: meta text is the smallest.
- `font-weight >= 400`: no light weights.
- `font-families <= 2`: Helvetica and a monospace.
- `line-height <= 1.45`: body text is set at 1.42.
- `underlined-links <= 10%`: links are gilt, not underlined, until hovered.
- `letter-spacing <= 1px`: nothing is tracked wider than small capital titles.
- `transition = none`: states change at once.
- `animation = none`: nothing moves.
