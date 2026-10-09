## Colour roles

- Grey, red and violet on white. `--color-page`, `--color-canvas` and `--color-surface` are all #ffffff: there is no visible sheet. `--color-surface-alt` #eeeeee fills the sidebar, forms and alternate rows; `--color-surface-strong` #c0c0c0 is sidebar title strips and inactive tabs.
- `--color-bar` #666666 with white text is the link bar, panel titles, table heads and the current tab; `--color-bar-alt` #cc0000 with white text is the utility strip, footer strip and alternate panel titles. `--color-inverse` #000000 with white text is the one dark block.
- `--color-text` and `--color-heading` #000000; `--color-text-muted` #666666; `--color-heading-alt` #330066 (deep violet) for section titles and h3.
- `--color-link` #330066, `--color-link-quiet` #000000, visited #808080, hover #cc0000, active #000000. `--color-accent` #fefe99 with black text is the badge, the current navigation item and links on the dark block; `--color-accent-alt` #cc0000 marks promoted entries.
- The style has no block colours: `--color-fill-1` to `--color-fill-4` repeat the greys #eeeeee, #f8f8f8, #e0e0e0, #eeeeee. Rules `--color-border` #c0c0c0, outlines `--color-border-strong` #666666.
- `--color-button` #e0e0e0, secondary #eeeeee, black labels. `--color-danger` #cc0000 is the same red as the secondary bar, on `--color-danger-surface` #f8f8f8; `--color-success` #009933 and `--color-warning` #e8c068 are badge fills only.

## Typography roles

- Serif body, sans-serif furniture: `--font-body` is Times New Roman at `--text-base` 16px (the browser default of the period, left unstyled); `--font-heading` and `--font-ui` are Arial.
- `--text-ui` 13px Arial for bars, tabs, buttons and inputs; `--text-small` 10px Arial for meta lines, breadcrumb and footer. `--text-large` 16px; `--text-h3` 16px, `--text-h2` 18px, `--text-h1` and `--text-display` 24px: the browser's own heading steps of the period, set in Arial bold.
- `--weight-ui` 400: bar and secondary button labels are regular weight. No uppercase, no tracking; line heights are `normal`.
- Every link is underlined, including sidebar and navigation links (`--link-decoration` and `--link-decoration-quiet` both underline).

## Surface roles

- Glossy bars: `--fill-bar` and `--fill-bar-alt` are vertical gradients from a lighter mix of the bar colour through the colour itself to a darker mix, like the shaded bar graphics of the period. `--fill-button` and `--fill-button-secondary` shade from near white to the button colour, `--fill-accent` from a pale mix to the accent, and `--fill-inverse` from a lighter mix at the top to the colour halfway down. Pages, panels and inputs are flat.
- `--radius-pill` 8px turns badges and status labels into capsules; `--radius-control` 3px takes the corners off buttons and inputs; `--radius-panel` and `--radius-page` stay 0.
- `--border-width` 1px, `--border-width-strong` 2px, solid; controls keep the 2px bevel. No box shadows. `--shadow-text` is a 1px dark drop under text on bars and buttons, the engraved lettering of shaded bar graphics.
- `--focus-ring` is 1px dotted in `--color-focus`; `--transition` none.

## Never

- `palette-colours <= 12`: greys, one red, one violet; no tints.
- `font-size <= 24px`: the largest text is a 24px bold heading.
- `font-size >= 10px`: nothing is set below 10px.
- `font-weight <= 700`: regular and bold only.
- `font-families <= 3`: Times for text, Arial for headings and controls, a monospace for code.
- `letter-spacing <= 0px`: no text is tracked.
- `uppercase-text <= 0%`: nothing is uppercased by CSS.
- `underlined-links >= 90%`: every link is underlined.
- `border-radius <= 3px`: boxes are square and controls barely rounded; capsule badges are shapes.
- `border-width <= 2px`: 1px rules, 2px bevels on controls.
- `box-shadow = none`: nothing casts a shadow.
- `gradient-fills <= 25%`: gradients are for bars, buttons, badges and the one dark block; pages and panels are flat.
- `transition = none`: states change instantly.
- `animation = none`: nothing moves.
