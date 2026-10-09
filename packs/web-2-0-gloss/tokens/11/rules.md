## Colour roles

- Olive, cream and rust on a white page. `--color-page`, `--color-canvas` and `--color-surface` are #ffffff; `--color-surface-alt` is the cream #f3f3cc (alternate rows, quoted lines, the empty state) and `--color-surface-strong` a light olive tint #dde7b0 (title strips, tags, inactive tabs).
- `--color-bar` #666633, a dark olive with cream `--color-bar-text` #ffffcc, fills the site bar, the table head and the dialog title. `--color-bar-alt` is the cream #f3f3cc with `--color-bar-alt-text` #333300. `--color-inverse` #333300 with #ffffcc text is the footer band.
- Text is a warm near-black: `--color-text` and `--color-heading` #333300, `--color-text-muted` #999999. `--color-heading-alt` #933100, a rust, for panel and sidebar titles and prose h2.
- `--color-link` #336699, a slate blue; hover turns rust (#933100), visited #666666, active #333300. `--color-link-quiet` #666633 for names, tags and minor tools.
- `--color-button` #a94400, a burnt orange with white text, is the primary button; `--color-button-secondary` is cream with #333300 text in a `--color-border-strong` #999966 outline. Fields are a very pale blue (`--color-input` #f8fdff) in the same olive outline.
- `--color-accent` #9eb847, a leaf green with #333300 `--color-accent-text`: count blocks, the pill badge, the current page. `--color-accent-alt` is the rust: the NEW marker only.
- Block fills: `--color-fill-1` #f3f3cc (hero), `--color-fill-2` #f8fdff, `--color-fill-3` #dde7b0, `--color-fill-4` #ffffcc.
- Borders: `--color-border` #aaaaaa, `--color-border-strong` #999966, `--color-border-muted` #d6d6ad (the cream, darkened) between rows.
- Status: `--color-danger` #aa0000 on `--color-danger-surface` #fde8dc, `--color-success` #5c7a1a, `--color-warning` #ffe840, as error text and badge fills only. `--color-notice` #ffffcc with #333300 text. `--color-focus` #336699; `--color-shadow` #333300.

## Typography roles

- A serif set: `--font-body` and `--font-heading` are Georgia, "Times New Roman", serif. `--font-ui` is "Lucida Grande", Arial, Helvetica, sans-serif for buttons, tabs, fields and table text. `--font-mono` ("Courier New") for inline code only.
- `--text-base` 13px at `--line-body` 1.3; `--text-small` 11px for meta lines and the footer; `--text-ui` 11px bold (`--weight-ui` 700).
- Headings are Georgia in regular weight (`--weight-heading` and `--weight-display` 400): `--text-display` 22px (the name in the header, hero title), `--text-h1` 20px, `--text-h2` 17px, `--text-h3` 14px, at `--line-heading` 1.2. `--text-large` 15px for the hero lead and the large button. Emphasis inside text is `--weight-bold` 700.
- No tracking and no uppercase: `--heading-tracking`, `--display-tracking` 0; `--heading-transform`, `--ui-transform` none.
- No link is underlined at rest: `--link-decoration` and `--link-decoration-quiet` are none, `--link-decoration-hover` underline.

## Surface roles

- Matte, not glossy: `--fill-bar-alt`, `--fill-inverse`, `--fill-accent`, `--fill-panel` and `--fill-button-secondary` are flat colours. Only `--fill-bar` (12% lighter at the top) and `--fill-button` (18% lighter at the top) carry a soft sheen with no hard break.
- Softly rounded boxes, nearly square controls: `--radius-panel` and `--radius-page` 6px, `--radius-control` and `--radius-pill` 3px.
- No shadows of any kind: `--shadow-panel`, `--shadow-control*`, `--shadow-dialog` and `--shadow-text` are none.
- `--border-width` and `--border-width-strong` are 1px solid. `--focus-ring` is a 1px dotted outline in `--color-focus`. `--transition` none.

## Never

- `palette-colours <= 32`: olive in three steps, cream, rust, burnt orange, one slate blue, greys and the status colours.
- `font-size <= 22px`: the name in the header and the hero title are the largest text.
- `font-size >= 11px`: meta text is 11px and nothing is smaller.
- `font-weight <= 700`: bold is the heaviest weight.
- `font-families <= 3`: Georgia, Lucida Grande and a monospace for inline code.
- `line-height <= 1.3`: body text is 13px on 17px.
- `letter-spacing = 0px`: no tracking anywhere.
- `uppercase-text = 0%`: nothing is uppercased by CSS.
- `underlined-links <= 5%`: links are not underlined until hovered.
- `border-radius <= 6px`: corners are 3 or 6px.
- `border-width <= 1px`: every border and rule is a 1px hairline.
- `box-shadow = none`: nothing casts a shadow.
- `text-shadow = none`: text is flat on its bar.
- `gradient-fills <= 10%`: only the bar, the table head and the primary button carry a sheen.
- `transition = none`: states change instantly.
- `animation = none`: nothing moves.
