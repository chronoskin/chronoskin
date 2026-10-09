## Colour roles

- The page is black (`--color-page`, `--color-canvas`, `--color-bar` #000000) and text is white (`--color-text`, `--color-heading`, `--color-heading-alt`, the link colours #ffffff). `--color-surface` is near-black #242423 for cards, panels, tables, inputs and the notice; `--color-surface-alt` is warm dark grey #4a443c for the empty state, the current row, inline code and a hovered tab.
- Outlines are white (`--color-border`, `--color-border-strong`, `--color-input-border` #ffffff) and every hard shadow is pink (`--color-shadow` #ff90e8), as is the focus ring (`--color-focus`).
- Pink #ff90e8 with black text is the accent: the primary button (`--color-button`), the current navigation item, tab and page and the default badge (`--color-accent`), the announcement strip (`--color-bar-alt`), an active link. The secondary button is white with black text.
- Fills that carry white text are deep, never pastel: `--color-surface-strong` magenta #b23386 (panel and dialog title strips, hovered link), `--color-accent-alt` teal #23a094 (the marked words of the hero title), `--color-fill-1` to `--color-fill-4` teal #23a094, red #dc341e, dark grey #4a443c and magenta #b23386.
- The inverted block is white: `--color-inverse` #ffffff with `--color-inverse-text` #000000 for the footer.
- Status, all with white text: `--color-success` teal #23a094, `--color-warning` grey-brown #6f675a (the nearest dark tone; a bright amber cannot carry white text), `--color-danger` #e2442f for error text and the destructive button, `--color-danger-surface` #dc341e behind an error notice.
- `--color-text-muted` #b3ac9e for dates and hints; `--color-border-muted` #6f675a for faint dividers; disabled is #4a443c with #b3ac9e text; `--color-overlay` is #242423 at 90%.

## Typography roles

- Two families: `--font-body` is the system sans (system-ui, Segoe UI, Roboto) for running text; `--font-heading` and `--font-ui` are Space Grotesk (then Archivo, then the system sans) for headings, buttons, navigation and inputs. `--font-mono` (Geist Mono, then the system monospace) for inline code.
- Headings are uppercase: `--heading-transform: uppercase` on the hero title, band titles, the page title and prose headings, weight 700 (`--weight-heading`, `--weight-display`), `--line-heading` 1 and `--line-display` 1.1, no tracking. Buttons and navigation stay in sentence case (`--ui-transform: none`).
- Text is small and dense: `--text-base` 14px at `--line-body` 1.625, `--text-small` 12px (badges come out at about 10px), `--text-ui` 14px, `--text-large` 18px for leads and large buttons.
- `--text-display` 60px for the hero title; `--text-h1` 48px, `--text-h2` 30px, `--text-h3` 18px.
- Weights: 400 body, 500 (`--weight-ui`) navigation, links, inputs and leads, 700 (`--weight-bold`) buttons, labels and badges.
- Links in running text are underlined; navigation, titles and breadcrumbs are not.

## Surface roles

- `--border-width` is 2px on every box and rule; `--border-width-strong` 3px under the navigation.
- Every radius token is 0: all boxes, controls and badges are square.
- Shadows are hard and fall to the lower left, the mirror of the usual direction: `--shadow-panel` and `--shadow-control` are `-4px 4px 0 0` in `--color-shadow`, `--shadow-control-hover` (also pagination squares and icon boxes) `-2px 2px 0 0`, `--shadow-control-pressed` zero offset, `--shadow-dialog` `-6px 6px 0 0`.
- All fills are flat (`--fill-page` is the plain page colour, no grid).
- `--focus-ring` is `2px solid` in `--color-focus`; `--transition` animates only `transform` and `box-shadow`, 0.1s linear.

## Never

- `border-radius <= 0px`: every corner is square; no rounding and no pills.
- `border-width >= 2px`: outlines and rules are 2px, 3px under the navigation; no hairlines.
- `border-width <= 3px`: nothing heavier than 3px.
- `box-shadow-blur <= 0px`: every shadow is a hard offset block.
- `text-shadow = none`: no text shadow.
- `gradient-fills <= 0%`: every fill is flat, including the page.
- `font-size <= 60px`: the hero title is the largest text.
- `font-size >= 10px`: badge and footer labels, at about 10px, are the smallest text.
- `font-weight >= 400`: no light weights.
- `font-weight <= 700`: emphasis stops at 700.
- `font-families <= 3`: a system sans for text, one grotesque for headings and controls, a monospace for code.
- `line-height <= 1.63`: body text is set at 1.625.
- `letter-spacing = 0px`: neither uppercase headings nor labels are tracked.
- `uppercase-text <= 15%`: headings, badges and footer labels are uppercase; running text, buttons and navigation are not.
- `underlined-links <= 20%`: only links in running text are underlined.
- `palette-colours <= 14`: black, white, near-black, pink, magenta, teal, two reds and the warm greys.
- `animation = none`: nothing moves on its own.
