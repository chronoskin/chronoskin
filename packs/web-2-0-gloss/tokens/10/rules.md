## Colour roles

- A dark set. `--color-page` and `--color-canvas` are charcoal #1a1a1a; boxes, table cells and panels are `--color-surface` #0e0e0e, darker than the page, in a `--color-border` #444444 hairline. `--color-surface-alt` #1a1a1a is the alternate row and `--color-surface-strong` #2f2f2f the title strips, tags and inactive tabs; rules inside a box are `--color-border-muted` #2f2f2f.
- `--color-text` #cccccc, `--color-text-muted` #888888, `--color-heading` #ffffff. `--color-heading-alt` #c8edbc, a pale green, for panel and sidebar titles.
- Links are pale blue: `--color-link` #c5e6ff, hover #ffffff, visited #bbbbbb, active the pale green. `--color-link-quiet` #aaaaaa for names, tags and minor tools.
- `--color-bar` #2a5fb0, a saturated blue with white `--color-bar-text`, fills the site bar, the table head and the dialog title. `--color-bar-alt` #0e0e0e with #cccccc text is the strip under it. `--color-inverse` #f3f3f3 with #555555 text: on this set the inverted block is a light band on the dark page.
- `--color-button` #3c9632, a glossy green with white text, is the primary button; `--color-button-secondary` #f3f3f3 with black text is a light button on the dark page. Fields stay white (`--color-input` #ffffff, black text, #777777 border).
- `--color-accent` #a4d62c, a lime with black `--color-accent-text`: count blocks, the pill badge, the current page. `--color-accent-alt` #e2501c: the NEW marker only.
- Block fills are near-blacks and greys: `--color-fill-1` #0e0e0e (hero), `--color-fill-2` #2f2f2f, `--color-fill-3` #444444, `--color-fill-4` #000000.
- Messages are light boxes on the dark page: `--color-notice` #ffffcc with #333300 text, `--color-danger` #f03840 on `--color-danger-surface` #fee5f2. `--color-success` #3c9632 and `--color-warning` #ffe28a are badge fills. `--color-focus` #c5e6ff; `--color-shadow` #000000; `--color-overlay` is a white veil, since a black one would not show.

## Typography roles

- Two families: `--font-body` is "Lucida Grande", Verdana, sans-serif; `--font-heading` and `--font-ui` are "Trebuchet MS", Tahoma, sans-serif. `--font-mono` (Monaco) for inline code only.
- `--text-base` 12px at `--line-body` 1.4; `--text-small` 10px for meta lines, tags and the footer; `--text-ui` 12px bold (`--weight-ui` 700) for buttons, tabs, table text and labels.
- Headings are bold Trebuchet: `--text-display` 26px (the name in the header, hero title), `--text-h1` 22px, `--text-h2` 16px (list, panel and sidebar titles), `--text-h3` 14px, at `--line-heading` 1.2. `--text-large` 14px for the hero lead and the large button.
- No tracking and no uppercase: `--heading-tracking`, `--display-tracking` 0; `--heading-transform`, `--ui-transform` none.
- No link is underlined at rest: `--link-decoration` and `--link-decoration-quiet` are none, `--link-decoration-hover` underline. Links are told apart by their pale colour and, in titles, by weight.

## Surface roles

- Rounded: `--radius-control` 4px, `--radius-panel` and `--radius-page` 6px, `--radius-pill` 10px.
- Hard gloss on the coloured parts: `--fill-bar`, `--fill-button` and `--fill-accent` are lighter in the upper half, break at 50% and darken toward the lower edge. `--fill-button-secondary` runs from near white to a light grey. `--fill-bar-alt` and `--fill-inverse` have a faint lighter upper edge; `--fill-panel` runs from the page colour down into the darker surface.
- `--shadow-text` is a 1px dark edge under the letters on bars. No box shadows: `--shadow-panel`, `--shadow-control*` and `--shadow-dialog` are none; on a dark page depth comes from gloss and hairlines.
- `--border-width` 1px, `--border-width-strong` 2px, solid. `--focus-ring` is a 1px dotted outline in `--color-focus`. `--transition` none.

## Never

- `palette-colours <= 32`: charcoal, near-blacks and greys, one blue, one green, one lime, two pale link tints and the light message colours.
- `font-size <= 26px`: the name in the header and the hero title are the largest text.
- `font-size >= 10px`: meta text is 10px and nothing is smaller.
- `font-weight <= 700`: bold is the heaviest weight.
- `font-families <= 3`: Lucida Grande, Trebuchet MS and a monospace for inline code.
- `line-height <= 1.4`: body text is 12px on 16.8px.
- `letter-spacing = 0px`: no tracking anywhere.
- `uppercase-text = 0%`: nothing is uppercased by CSS.
- `underlined-links <= 5%`: links are not underlined until hovered.
- `border-radius <= 10px`: corners are 4, 6 or 10px.
- `border-width <= 2px`: hairlines are 1px; only an emphasised frame is 2px.
- `box-shadow = none`: nothing casts a shadow.
- `transition = none`: states change instantly.
- `animation = none`: nothing moves.
