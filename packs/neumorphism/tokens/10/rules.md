## Colour roles

- A warm soft UI. One surface: `--color-page`, `--color-canvas`, `--color-surface`, `--color-bar`, `--color-input` and `--color-button-secondary` are the same cream (242 235 227). `--color-surface-alt` and `--color-notice` (248 243 237) are a step lighter, for tiles, hovered rows and notices; `--color-surface-strong` (228 217 205) a step darker, for inline code, quiet badges and disabled controls.
- The relief pair: `--color-shadow` (211 198 184) is the dark side and `--color-border-muted` (warm white, 255 252 247) the light side of every shadow.
- Text is warm brown, never black: `--color-text` 84 70 64, `--color-text-muted` 124 106 96, `--color-heading` 62 48 44, `--color-heading-alt` 104 86 78.
- `--color-accent` and `--color-button` are one coral (226 88 68) with warm white text; `--color-link` is the darker coral 196 68 50, which is also the focus ring. `--color-accent-alt` is teal (46 140 140) and only marks "new".
- The footer and code block are the dark brown of the headings (`--color-bar-alt`, `--color-inverse`, 62 48 44) with cream text.
- Status: `--color-success` 70 140 90, `--color-warning` 200 124 16, `--color-danger` 184 36 48 on a blush `--color-danger-surface`. `--color-fill-1` to `--color-fill-4` are coral, teal, mustard and plum.
- `--color-border` (228 217 205) is for rules between rows; `--color-input-border` and `--color-border-strong` (212 198 183) for hairlines when a surface part draws them.

## Typography roles

- One rounded family for everything: Rubik, then Varela Round and Trebuchet MS; `--font-mono` for code.
- `--text-base` 15px at `--line-body` 1.6, `--text-ui` 14px, `--text-small` 13px, `--text-large` 18px, `--text-h3` 17px, `--text-h2` 22px, `--text-h1` 30px, `--text-display` 42px.
- Two weights: `--weight-body` 400; `--weight-bold`, `--weight-heading`, `--weight-display` and `--weight-ui` 700. `--line-heading` 1.25, `--line-display` 1.15.
- No uppercase and no tracking from tokens; links are marked by colour and underlined on hover only.

## Surface roles

- Fully soft and very round: `--border-width` is 0; `--radius-control` and `--radius-panel` 22px, so a 40px button is nearly a pill, `--radius-page` 32px, `--radius-pill` 40px. `--border-width-strong` (2px) is the focus and error ring.
- Medium relief: `--shadow-control` 5px offset and 10px blur, `--shadow-panel` and `--shadow-control-hover` 9px and 18px, `--shadow-control-pressed` 4px and 8px inset, `--shadow-dialog` 16px and 36px with its light side at 35%.
- Every fill is flat; `--fill-button-hover` mixes 12% of the light colour into the button colour. `--shadow-text` is `none`. Transitions run 0.2s.

## Never

- `border-width <= 2px`: nothing is outlined; 1px rules between rows and the 2px ring are the only lines.
- `border-radius <= 32px`: controls and panels 22px, tiles and dialogs 32px.
- `box-shadow-blur <= 36px`: panels blur 18px, the dialog and the hero disc 36px.
- `text-shadow = none`: text is flat.
- `gradient-fills <= 5%`: every fill is one flat colour.
- `font-size <= 42px`: the hero title is the largest text.
- `font-size >= 13px`: meta text and badges are the smallest text.
- `font-weight <= 700`: headings are bold, never black.
- `font-families <= 2`: one sans and one monospace.
- `line-height <= 1.65`: body text is set at 1.6.
- `letter-spacing = 0`: no tracking on any text.
- `underlined-links <= 10%`: underline is a hover state.
