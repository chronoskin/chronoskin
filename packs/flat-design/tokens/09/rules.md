## Colour roles

- Family: the 2014 site-builder page (black, white, warm off-white, two greys) with the browns, beige and terracotta read from its photograph; `--color-success` `#2ebd59` and `--color-warning` `#f59b23` are borrowed from the 2015 music page, because the family has neither.
- `--color-page`, `--color-canvas`, `--color-surface` and `--color-bar` are white `#ffffff`; `--color-surface-alt` `#f6f5f3` (warm off-white) is the alternate band and row, `--color-surface-strong` `#e8e8e8` the panel title strip.
- Body text is grey `--color-text` `#666666`, meta `--color-text-muted` `#777777`, headings `--color-heading` `#222222`; small titles are dark brown `--color-heading-alt` `#482818`.
- There is no colour for action: `--color-button` `#222222` with white text, and the secondary button `#222222` on white. Links are black `--color-link` `#0c0c0c` and stand out from the grey text by darkness alone; hover fades to `--color-link-hover` `#777777`, active is `#482818`.
- `--color-fill-1` `#482818` (dark brown) is the hero and the emphasis band; `--color-fill-2` `#0c0c0c`, `--color-fill-3` `#d4ccbf` (beige, icon only) and `--color-fill-4` `#724f35` (wood) are the other tiles.
- `--color-accent` `#724f35` (white text) is the default badge, `--color-accent-alt` `#d4ccbf` the second. `--color-danger` is the terracotta `#a64538`; the error notice sits on `--color-danger-surface` `#f6f5f3`. The notice is `--color-notice` `#e8e8e8` with `#222222` text.
- Borders: `--color-border` `#e8e8e8`, `--color-border-muted` `#f0f0f0`; `--color-border-strong` `#222222` is the black rule under the table head; `--color-input-border` `#777777`. `--color-inverse` `#0c0c0c` with white text is the footer. `--color-shadow` is black at 15%.

## Typography roles

- One heavy geometric sans, after the 2015 music page: `--font-body`, `--font-heading` and `--font-ui` are Montserrat, Avenir Next, Century Gothic. `--font-mono` is Menlo.
- Large body: `--text-base` 16px at `--line-body` 1.5, `--text-large` 18px, `--text-small` 12px, `--text-ui` 14px.
- Headings are bold and tight: `--weight-heading` 600 with `--heading-tracking` -0.5px at `--text-h1` 36px, `--text-h2` 24px, `--text-h3` 18px. The hero title and the big figures are `--text-display` 44px at `--weight-display` 800, `--line-display` 1 and `--display-tracking` -1px.
- `--ui-transform` is `uppercase` at `--weight-ui` 600: navigation, tabs and buttons. `--weight-bold` is 600.
- All three link decorations are `none`.

## Surface roles

- Everything is square and thin: `--radius-control`, `--radius-panel` and `--radius-page` are 0, `--radius-pill` 2px, so badges are small tags with barely softened corners.
- Depth is the short hard diagonal of the period's long-shadow drawings, never a blur: `--shadow-control` is `2px 2px 0 0` in `--color-shadow` under buttons and inputs, `3px 3px` when hovered and gone when pressed; `--shadow-panel` is `4px 4px 0 0` under panels and `--shadow-dialog` `8px 8px 0 0` under a dialog. Text has no shadow.
- Buttons are solid: `--fill-button` is the flat `--color-button`, hovered mixed 20% with white; `--fill-button-secondary` is the solid `--color-button-secondary` with a `--border-width` 1px outline.
- `--border-width` 1px, `--border-width-strong` 3px for the current tab and navigation item and under the table head. `--focus-ring` is `1px solid` in `--color-focus`. `--transition` fades colour, background and border colour in 0.2s.

## Never

- `box-shadow-blur <= 0px`: shadows are hard diagonal offsets under controls, panels and dialogs; nothing is blurred.
- `text-shadow = none`: text is flat on every fill.
- `gradient-fills <= 0%`: every background is one solid colour.
- `border-radius <= 2px`: boxes and controls are square; only tags are eased, by 2px.
- `border-width <= 3px`: rules are 1px; 3px marks the current item and the table head.
- `font-weight >= 400`: this type set has no light weight.
- `font-weight <= 800`: the hero title and figures are the heaviest text.
- `font-size >= 12px`: meta text and badges are the smallest text.
- `font-size <= 44px`: the hero title is the largest text.
- `font-families <= 2`: one sans-serif family, plus monospace for code.
- `line-height <= 1.5`: body text is set at 1.5.
- `letter-spacing <= 1px`: headings are tightened by at most 1px; nothing is spaced out further.
- `underlined-links <= 0%`: no reference underlines links.
- `animation = none`: no reference has a CSS animation.
