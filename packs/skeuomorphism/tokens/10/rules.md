## Colour roles

- Manila card, cream paper and dark saddle leather. `--color-page` #d9c59c is the manila desk; `--color-canvas` and `--color-surface` #fbf5e6 the cream sheet; `--color-surface-alt` #f3ead3 the tone sections fade to; `--color-surface-strong` #e9dcbd toolbar and title strips.
- `--color-bar` #4b2e1c (dark brown leather, cream text #f7ecd2) is the navigation bar, table head, count badges and pressed "current" states; `--color-bar-alt` #d6c193 (card, dark brown text #47301c) is title strips and the segmented control; `--color-inverse` #6f2a22 (oxblood leather, cream text) the closing band or hero block.
- Text is the brown ink `--color-text` #4a3b2b, meta `--color-text-muted` #7d6a52, headings `--color-heading` #2e2015, sidebar and panel titles `--color-heading-alt` #5a3a1f. Nothing is pure black.
- Links are a teal ink, `--color-link` #1c6672, darker on hover (#0f4852), slate when visited (#5c5a7a), red when active.
- `--color-button` #3d6b47 (bookcloth green, white text) is the one primary action colour; the secondary key is `--color-button-secondary` #eadfc4 with #47301c text.
- `--color-accent` #b5382c (ribbon red) for badges, `--color-accent-alt` #e3b23c (brass, dark text) for "new" markers; `--color-success` #4a7d3a, `--color-warning` #e3b23c (dark text), `--color-danger` #a3271c on `--color-danger-surface` #f8e4da. Notices are legal-pad yellow, `--color-notice` #fbf0a8 with #4a3b12 text.
- `--color-fill-1` to `--color-fill-4` (#2b6f73 teal, #3d6b47 green, #2f4a6b navy, #a8701c ochre) are book-cover colours for the grid blocks only; cream text on all four.
- `--color-shadow` is a dark brown, #1e0f04, not black, so every shadow is warm.

## Typography roles

- One plain sans for everything: `--font-body`, `--font-heading` and `--font-ui` are Helvetica, then Arial and Verdana; `--font-mono` Courier New.
- Text is large: `--text-base` 15px at `--line-body` 1.4; `--text-small` 12px; `--text-large` 18px for the lead and the large key.
- Headings are heavy and tightly set: `--weight-heading` and `--weight-display` 700, `--line-heading` 1.2. `--text-h3` 16px, `--text-h2` 21px, `--text-h1` 28px, `--text-display` 36px at `--line-display` 1.1 with `--display-tracking` -0.5px.
- Keys, navigation and table heads are small bold capitals: `--text-ui` 11px, `--weight-ui` 700, `--ui-transform: uppercase`.
- Links are coloured and not underlined until hovered (`--link-decoration: none`, `--link-decoration-hover: underline`).

## Surface roles

- Leather with stitching. `--fill-bar` is grained leather, lit from the top, with a row of cream stitches 3px inside its top and bottom edges; `--fill-inverse` is the same grain with one row of stitches along the top and a pool of light. The stitches are drawn in the bar's own text colour at 50 to 55%.
- `--fill-page` is card: an 8px fleck pattern and a soft pool of light at the top over `--color-page`.
- Paper lies on the desk: `--shadow-panel` is 0 1px 3px at 32% of the brown shadow; `--shadow-dialog` a ring and 0 2px 5px.
- Keys are stitched leather tabs, matte, not glossy: `--fill-button` is a soft top-lit gradient under a faint grain; `--shadow-control` a light inner top edge and a 1px shadow. Pressed keys take a 2px inset shade.
- `--fill-input` is ruled paper: a faint teal line every 20px on the field colour. `--fill-panel` stays the sheet colour and warms only in its last third.
- Every edge is a seam: `--border-style` is `dashed`, so the borders of keys, fields, panels, tables and rules are drawn as 1px stitches.
- Corners: `--radius-control` 4px, `--radius-panel` 4px, `--radius-pill` 12px, `--radius-page` 10px. `--shadow-text` is a dark line one pixel above light text. `--transition` is none.

## Never

- `border-radius <= 20px`: the pill is 12px; nothing is rounder than the first structure's large key.
- `border-width <= 1px`: every edge is a hairline of stitches, dashed and never thicker.
- `box-shadow-blur <= 5px`: paper casts a short shadow.
- `font-size <= 36px`: the hero headline is the largest text.
- `font-size >= 11px`: the small capitals of keys are the smallest text.
- `font-weight >= 400`: no light weights.
- `font-families <= 2`: Helvetica and a monospace.
- `line-height <= 1.45`: body text is set at 1.4.
- `underlined-links <= 10%`: links are coloured, not underlined, until hovered.
- `letter-spacing <= 1px`: nothing is tracked wider than small capital titles.
- `transition = none`: states change at once.
- `animation = none`: nothing moves.
