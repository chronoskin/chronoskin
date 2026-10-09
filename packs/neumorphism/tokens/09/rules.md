## Colour roles

- A dark soft UI. One surface: `--color-page`, `--color-canvas`, `--color-surface`, `--color-bar`, `--color-input`, `--color-notice` and `--color-button-secondary` are the same dark slate (39 55 76). `--color-surface-alt` (53 64 104) is a step lighter and bluer, for tiles and hovered rows; `--color-surface-strong` (31 34 37) is a step darker, for inline code.
- The relief pair: `--color-shadow` (31 34 37) is the dark side and `--color-border-muted` (53 64 104) the light side of every shadow. The light side is a lighter slate, not white.
- Text is near-white (`--color-text` 251 251 254, `--color-heading` white) with blue-grey for secondary text (`--color-text-muted` 178 184 201) and pale periwinkle for panel titles (`--color-heading-alt` 209 216 236).
- `--color-button` is pale periwinkle (209 216 236) with near-black text (31 34 37): the primary button is the lightest thing on the page. The code block uses the same periwinkle (`--color-inverse`) with dark text.
- `--color-accent` is bright blue (72 160 255) with near-black text; `--color-link` is the same blue. `--color-accent-alt` is amber (255 162 44), for "new" and the active link.
- The footer is near-black (`--color-bar-alt` 28 28 28) with periwinkle text.
- Status colours are light enough to read on slate: `--color-success` 34 197 94, `--color-warning` 255 162 44, `--color-danger` 255 154 162. `--color-fill-1` to `--color-fill-4` are violet, magenta, amber and blue chips.
- `--color-border` and `--color-input-border` (31 34 37) are dark hairlines for rules between rows.

## Typography roles

- One geometric family for everything: Poppins, then Montserrat, Futura and Century Gothic; `--font-mono` for code. The references used commercial geometric faces; this is the closest open and system stack.
- Small text under a very large title: `--text-base` and `--text-ui` 14px, `--text-small` 12px, `--text-large` 16px, `--text-h3` 16px, `--text-h2` 20px, `--text-h1` 36px, `--text-display` 60px.
- `--weight-body` 400, `--weight-ui` 500, `--weight-bold`, `--weight-heading` and `--weight-display` 700.
- `--line-body` 1.57, `--line-heading` 1.2, `--line-display` 1.1.
- Links in running text are underlined (`--link-decoration: underline`) and lose the underline on hover; quiet links and navigation are never underlined. No uppercase and no tracking from tokens.

## Surface roles

- Fully soft with tight corners: `--border-width` is 0; `--radius-control` 5px, `--radius-panel` and `--radius-page` 10px, `--radius-pill` 30px.
- Medium relief with deep wells: `--shadow-control` 4px offset and 8px blur, `--shadow-panel` and `--shadow-control-hover` 7px and 14px, `--shadow-dialog` 12px and 24px. `--shadow-control-pressed` is 7px and 14px inset, so inputs, the tab track and the empty state are visibly hollow. Resting shadows use both colours at 85%.
- Keys are convex: `--fill-button` and `--fill-button-secondary` are a 145 degree gradient, from the button colour mixed 40% with the light colour to the same mixed 30% with the shadow colour; the secondary key, which has the surface's own colour, runs from 55% light to 35% shadow so that its dome still reads. `--fill-button-hover` is the flat button colour. Panels, bars and wells stay flat.
- `--shadow-text` is a soft drop (0 3px 4px, the shadow colour at 40%) on the brand, the hero title and large figures.
- Transitions run 0.2s.

## Never

- `border-width <= 2px`: nothing is outlined; 1px rules between rows and the 2px ring are the only lines.
- `border-radius <= 10px`: controls 5px, panels, tiles and dialogs 10px.
- `box-shadow-blur <= 24px`: panels blur 14px, the dialog and the hero disc 24px.
- `gradient-fills <= 15%`: only the keys are convex; panels and wells are flat.
- `font-size <= 60px`: the hero title is the largest text.
- `font-size >= 12px`: meta text and badges are the smallest text.
- `font-weight <= 700`: headings are bold, never black.
- `font-families <= 2`: one sans and one monospace.
- `line-height <= 1.6`: body text is set at 1.57.
- `letter-spacing = 0`: no tracking on any text.
