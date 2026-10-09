## Colour roles

- `--color-bar` #000000 with white `--color-bar-text` fills the site bar, the table head and the dialog title; `--color-bar-alt` #eeeeee is the topic strip and the footer link strip, with black `--color-bar-alt-text`. `--color-inverse` #444444 with white text is the small-print strip.
- One hue only: `--color-link` #0000ff, the pure blue, also `--color-button` (primary button, current page number, square count badges, white text). Visited #708cf2, hover #000000, active #444444.
- Greys carry the rest: `--color-text` #444444, `--color-heading` #000000, `--color-heading-alt` and `--color-link-quiet` #666666, `--color-text-muted` #888888, `--color-surface-alt` #f6f6f6, `--color-surface-strong` #eeeeee (title strips, tags).
- `--color-accent` #dddddd with black `--color-accent-text`: count blocks, the pill badge and the notice border are grey, not coloured. `--color-notice` is #eeeeee with #444444 text. `--color-accent-alt` #ff6600 (feed orange) is the NEW marker only.
- `--color-button-secondary` #efefef with black text in a `--color-border-strong` #767676 outline; fields have the same #767676 `--color-input-border`. Other borders are `--color-border` #cccccc and `--color-border-muted` #dddddd.
- Block fills are greys: `--color-fill-1` #eeeeee (hero), `--color-fill-2` #f6f6f6, `--color-fill-3` #dddddd, `--color-fill-4` #eeeeee.
- Status: `--color-danger` #cc0000 on `--color-danger-surface` #ffe8cd, `--color-success` #339900, `--color-warning` #ffa100, used as error text and badge fills only. `--color-focus` #708cf2; `--color-shadow` #000000.

## Typography roles

- One family everywhere: `--font-body`, `--font-heading` and `--font-ui` are "Helvetica Neue", Arial, sans-serif; `--font-mono` ("Courier New") for inline code only.
- Small and tight: `--text-base` 12px at `--line-body` 1.25, `--text-small` 11px for meta lines, tools, tags and the footer, `--text-ui` 12px bold (`--weight-ui` 700) for buttons, the topic strip, tabs, table text and labels.
- A short heading scale, all bold: `--text-display` 21px (brand name, hero title), `--text-h1` 18px, `--text-h2` 15px (list titles, panel and sidebar titles), `--text-h3` 13px, at `--line-heading` 1.2. `--text-large` 14px for the hero lead and the large button.
- No tracking and no uppercase: `--heading-tracking`, `--display-tracking` 0; `--heading-transform`, `--ui-transform` none.
- No link is underlined at rest: `--link-decoration` and `--link-decoration-quiet` are none, `--link-decoration-hover` underline. Links are told apart by colour and, in titles, by weight.

## Surface roles

- Nearly square: `--radius-panel` 0 (panels, tabs, notices, dialog, stats box), `--radius-control` and `--radius-pill` 2px (buttons, count blocks, badges), `--radius-page` 4px (hero, grid cells).
- A satin sheen instead of a gloss line: `--fill-bar` starts 45% lighter than the bar colour at the top edge and settles into the plain colour by the bottom, with no hard break; `--fill-button` runs from 12% lighter to 14% darker; `--fill-button-secondary` from near white to slightly darker than its colour; `--fill-inverse` has a darker upper edge.
- `--fill-bar-alt`, `--fill-accent` and `--fill-panel` are flat (`--fill-panel` is `--color-surface`): strips, count blocks and panels have no gradient, a panel is a white box in a hairline.
- No shadows: `--shadow-panel`, `--shadow-control*`, `--shadow-dialog` and `--shadow-text` are none.
- `--border-width` and `--border-width-strong` are 1px solid. `--focus-ring` is a 1px solid outline in `--color-focus`. `--transition` none.

## Never

- `palette-colours <= 28`: black, greys and one pure blue; orange, red, green and amber appear only as markers.
- `font-size <= 21px`: the brand name and hero title are the largest text.
- `font-size >= 11px`: meta text is 11px and nothing is smaller.
- `font-weight <= 700`: bold is the heaviest weight.
- `font-families <= 2`: Helvetica Neue for everything, a monospace only for inline code.
- `line-height <= 1.25`: body text is 12px on 15px.
- `letter-spacing = 0px`: no tracking anywhere.
- `uppercase-text = 0%`: nothing is uppercased by CSS.
- `underlined-links <= 5%`: links are not underlined until hovered.
- `border-radius <= 4px`: corners are 0, 2 or 4px.
- `border-width <= 1px`: every border and rule is a 1px hairline.
- `box-shadow = none`: nothing casts a shadow.
- `text-shadow = none`: text is flat on its bar.
- `gradient-fills <= 10%`: only bars, buttons and tabs carry a gradient; strips, blocks and panels are flat.
- `transition = none`: states change instantly.
- `animation = none`: nothing moves.
