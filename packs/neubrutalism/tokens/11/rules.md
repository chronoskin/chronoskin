## Colour roles

- Two colours and no more: black and one acid lime. The page, the canvas and every surface are white (`--color-page`, `--color-canvas`, `--color-surface` #ffffff); text, headings, every link colour, borders, focus and shadow are black #000000.
- Lime #ccff00 with black text is everything that is coloured: the primary button (`--color-button`), the current navigation item, tab and page and the default badge (`--color-accent`), the marked words of the hero title (`--color-accent-alt`), panel and dialog title strips and the highlight behind a hovered link (`--color-surface-strong`), and `--color-fill-1` and `--color-fill-4`.
- The only tint is pale lime #ecfcca: `--color-surface-alt` (empty state, current row, inline code, hovered tab), `--color-fill-3` and `--color-notice`. `--color-fill-2` is white, so blocks alternate lime, white, pale lime, lime. Do not add a third hue for decoration.
- Black blocks carry the contrast: `--color-inverse` is black with white text (the footer), `--color-bar-alt` is black with lime text (the announcement or tag strip). The navigation bar and the table head (`--color-bar`) stay white with black text. The secondary button is white.
- Status is the one place other colours appear, always as a flat fill under black text: `--color-success` lime #ccff00, `--color-warning` orange #ff9500, `--color-danger` and `--color-danger-surface` hot pink #ff006e (error text, the destructive button, the danger badge, the error notice).
- `--color-text-muted` #5c5c66 for dates and hints; `--color-border-muted` #d8d8d8 for faint dividers; disabled is #e8e8e8 with #5c5c66 text; `--color-overlay` is black at 80%.

## Typography roles

- Two voices: a monospace for everything that is read or pressed (`--font-body`, `--font-ui`, `--font-mono`: Fragment Mono, then DM Mono, JetBrains Mono and the system monospace) and a high-contrast serif for headings (`--font-heading`: Playfair Display, then Bodoni Moda, Didot and Georgia).
- Headings are uppercase serif capitals: `--heading-transform: uppercase`, weight 700 (`--weight-heading`, `--weight-display`), `--line-heading` 1.1 and `--line-display` 1, no tracking.
- Buttons, navigation and tabs are uppercase monospace: `--ui-transform: uppercase`, weight 500 (`--weight-ui`), 700 (`--weight-bold`) on buttons, labels and badges.
- Text is small and airy: `--text-base` 14px at `--line-body` 1.625, `--text-small` 12px (badges come out at about 10px), `--text-ui` 14px, `--text-large` 18px for leads and large buttons.
- `--text-display` 60px for the hero title; `--text-h1` 48px, `--text-h2` 36px, `--text-h3` 20px.
- Links in running text are underlined; navigation, titles and breadcrumbs are not.

## Surface roles

- Outlines are the heaviest of the era: `--border-width` and `--border-width-strong` are both 4px, on every box and every rule.
- Every radius token is 0: all boxes, controls and badges are square.
- Shadows are large, hard and black: `--shadow-panel` and `--shadow-dialog` `8px 8px 0 0`, `--shadow-control` `6px 6px 0 0`, `--shadow-control-hover` `4px 4px 0 0`, `--shadow-control-pressed` `2px 2px 0 0`.
- Three fills are hatched with thin diagonal lines, 2px in every 10px at 45 degrees: the page (`--fill-page`, `--color-border` at 9%), the secondary bar (`--fill-bar-alt`, its text colour at 22%) and the inverted block (`--fill-inverse`, its text colour at 14%). Every other fill is flat, and no box that holds running text is hatched.
- `--focus-ring` is `4px solid` in `--color-focus`; `--transition` animates only `transform` and `box-shadow` over 0.15s.

## Never

- `border-radius <= 0px`: every corner is square; no rounding and no pills.
- `border-width >= 4px`: every outline and rule is 4px; no hairlines.
- `border-width <= 4px`: nothing heavier than 4px.
- `box-shadow-blur <= 0px`: every shadow is a hard offset block.
- `text-shadow = none`: no text shadow.
- `gradient-fills <= 5%`: only the page, the secondary bar and the inverted block carry the diagonal hatch; cards, buttons and panels are flat.
- `font-size <= 60px`: the hero title is the largest text.
- `font-size >= 10px`: badge and footer labels, at about 10px, are the smallest text.
- `font-weight >= 400`: no light weights.
- `font-weight <= 700`: emphasis stops at 700.
- `font-families <= 2`: one monospace and one serif.
- `line-height <= 1.63`: body text is set at 1.625.
- `letter-spacing = 0px`: neither uppercase headings nor labels are tracked.
- `palette-colours <= 10`: black, white, lime, pale lime, the two status colours and the greys.
- `animation = none`: nothing moves on its own.
