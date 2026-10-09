## Colour roles

- Cream paper, brick red and near-black. `--color-page` #fbf8f3 is the cream page; `--color-canvas` and `--color-surface` #ffffff the sheet; `--color-surface-alt` #f8f8f0 the warm tone sections fade to; `--color-surface-strong` #f0f0e8 toolbar strips.
- `--color-bar` #ad2f1e (brick red, white text) is the navigation bar, table head, count badges and pressed "current" states; `--color-bar-alt` #d0d0d0 with #393939 text is title strips and the segmented control; `--color-inverse` #a0693b (wood brown, white text) only the closing band.
- Text is `--color-text` #333333, meta `--color-text-muted` #808080, headings `--color-heading` #1c1c1c, sidebar and panel titles `--color-heading-alt` #393939.
- Links are `--color-link` #0088cc, hover `--color-link-hover` #2b96f1, visited #4880a2, active #09517e.
- `--color-button` #454040 (warm near-black, white text) is the one primary action colour; the secondary key is the near-white `--color-button-secondary` #f6f6f6 with #393939 text.
- `--color-accent` #2b96f1 (azure) for badges, `--color-accent-alt` #a1cbef (pale azure, dark text) for "new" markers; `--color-success` #02b602, `--color-warning` #ba8f57 (light wood, dark text), `--color-danger` #f3201f on `--color-danger-surface` #fff8f8. Notices are `--color-notice` #fffff8 with #393939 text.
- `--color-fill-1` to `--color-fill-4` (#ad2f1e red, #454040 near-black, #a0693b wood, #2b96f1 azure) colour the grid blocks only; white text on all four.

## Typography roles

- Running text is a book serif: `--font-body` is Palatino Linotype, Palatino, Book Antiqua, Georgia. `--font-heading` is a geometric sans (Montserrat, Avenir Next, Century Gothic) standing in for the era's commercial one; `--font-ui` is Helvetica Neue; `--font-mono` Monaco.
- Larger than the other sets: `--text-base` 15px at `--line-body` 1.56; `--text-small` 12px; `--text-ui` 13px at `--weight-ui` 700; `--text-large` 20px for the lead and the large key.
- `--text-h3` 16px, `--text-h2` 20px, `--text-h1` 28px with `--weight-heading` 700; `--text-display` 40px at `--weight-display` 400 and `--display-tracking` -1px, with a bold phrase inside the regular headline as the emphasis.
- Links are not underlined until hovered (`--link-decoration: none`, `--link-decoration-hover: underline`). No uppercase from this part.

## Surface roles

- Felt and soft pills. `--fill-page` is a fine 4px speckle over `--color-page`, with no gradient. `--fill-bar` and `--fill-inverse` are felt: the same speckle and a pool of light from the top centre (white at 30%) over the flat bar colour.
- Controls are pills: `--radius-control` 13px rounds the 28px keys and fields fully, `--radius-pill` 20px the large key, navigation items and badges; `--radius-panel` and `--radius-page` are 6px.
- Keys are matte, lit from above: `--fill-button` goes from 30% lighter at the top through the plain colour at 55% to 14% darker, with a light inner top edge and a light line below (`--shadow-control`). No drop shadow under keys; pressed is a 1px inset shade.
- Nothing floats: `--shadow-panel` is only a 1px ring at 12% and `--shadow-dialog` a ring with a hard 3px step below. Blur comes only from the wells the structure sinks into the sheet.
- `--shadow-text` is a soft dark edge above light text (0 -1px 1px at 40%). `--fill-panel` fades evenly from `--color-surface` to `--color-surface-alt` over the whole height. `--fill-input` is flat. `--transition` is none.

## Never

- `border-radius <= 20px`: the 20px pill is the largest radius; panels and blocks are 6px.
- `border-width <= 1px`: every edge is a hairline.
- `box-shadow-blur <= 3px`: panels have rings, not shadows; only wells are shaded.
- `font-size <= 40px`: the hero headline is the largest text.
- `font-size >= 12px`: meta text is the smallest.
- `font-weight >= 400`: no light weights.
- `font-families <= 4`: a book serif, a geometric sans for headings, Helvetica for controls and a monospace.
- `line-height <= 1.6`: body text is set at 1.56.
- `underlined-links <= 10%`: links are coloured, not underlined, until hovered.
- `uppercase-text <= 5%`: only the small sidebar titles are capitals.
- `letter-spacing <= 1px`: nothing is tracked wider than the small capital titles.
- `transition = none`: states change at once.
- `animation = none`: nothing moves.
