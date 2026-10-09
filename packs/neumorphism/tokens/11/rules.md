## Colour roles

- A dark soft UI in neutral graphite. One surface: `--color-page`, `--color-canvas`, `--color-surface`, `--color-bar`, `--color-input`, `--color-notice`, `--color-danger-surface` and `--color-button-secondary` are the same graphite (42 44 49). `--color-surface-alt` (52 55 61) is a step lighter, for tiles and hovered rows; `--color-surface-strong` (30 32 36) a step darker, for inline code and quiet badges.
- The relief pair: `--color-shadow` (26 27 31) is the dark side and `--color-border-muted` (62 65 72) the light side of every shadow. The light side is a lighter graphite, not white.
- Text is pale grey (`--color-text` 232 233 237, `--color-heading` 246 246 248) with mid grey for secondary text (`--color-text-muted` 160 164 174) and `--color-heading-alt` 200 203 211 for panel titles.
- `--color-accent` and `--color-button` are one orange (255 138 48) with near-black brown text (32 22 12): the primary button is the brightest thing on the page. `--color-link` is a lighter orange (255 160 88). `--color-accent-alt` is cyan (90 200 220) and only marks "new".
- The footer is darker graphite (`--color-bar-alt` 30 32 36); the code block is inverted, pale grey with graphite text (`--color-inverse`).
- Status colours are light enough to read on graphite: `--color-success` 120 210 130, `--color-warning` 255 200 80, `--color-danger` 255 120 120. `--color-fill-1` to `--color-fill-4` are orange, cyan, violet and green.
- `--color-border` and `--color-input-border` (30 32 36) are dark hairlines for rules between rows.

## Typography roles

- One wide humanist family for everything: Oxygen, then Ubuntu and Verdana; `--font-mono` for code.
- Small text, because the face is wide: `--text-base` and `--text-ui` 13px at `--line-body` 1.6, `--text-small` 11px, `--text-large` 15px, `--text-h3` 15px, `--text-h2` 19px, `--text-h1` 28px, `--text-display` 40px.
- Two weights: `--weight-body` 400; `--weight-bold`, `--weight-heading`, `--weight-display` and `--weight-ui` 700. `--line-heading` 1.25, `--line-display` 1.15.
- No uppercase and no tracking from tokens; links are marked by colour and underlined on hover only.

## Surface roles

- Cut, not moulded: corners are nearly square. `--border-width` is 0; `--radius-control` 3px, `--radius-panel` 4px, `--radius-page` 6px, and `--radius-pill` is 4px too, so badges are small tags and a pill track becomes a slot. `--border-width-strong` (2px) is the focus and error ring.
- Tight, crisp relief with a lit edge: every raised shadow ends in `inset 1px 1px 0` of the light colour, a one pixel highlight along the top and left edges. `--shadow-control` is 3px offset and 5px blur on the dark side, 2px and 4px on the light side at 80%; `--shadow-panel` and `--shadow-control-hover` 5px and 10px; `--shadow-dialog` 10px and 20px with its light side at 50%. `--shadow-control-pressed` is a hard 3px and 5px inset on the dark side with a one pixel light edge at the bottom right.
- Every fill is flat. `--shadow-text` is `none`. Transitions are short and linear, 0.1s.

## Never

- `border-width <= 2px`: nothing is outlined; 1px rules between rows and the 2px ring are the only lines.
- `border-radius <= 6px`: controls 3px, panels and tags 4px, tiles and dialogs 6px; only discs are round.
- `box-shadow-blur <= 20px`: panels blur 10px, the dialog and the hero disc 20px.
- `text-shadow = none`: text is flat.
- `gradient-fills <= 5%`: every fill is one flat colour.
- `font-size <= 40px`: the hero title is the largest text.
- `font-size >= 11px`: meta text and badges are the smallest text.
- `font-weight <= 700`: headings are bold, never black.
- `font-families <= 2`: one sans and one monospace.
- `line-height <= 1.65`: body text is set at 1.6.
- `letter-spacing = 0`: no tracking on any text.
- `underlined-links <= 10%`: underline is a hover state.
