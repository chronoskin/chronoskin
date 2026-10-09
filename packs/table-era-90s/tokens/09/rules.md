## Colour roles

- Family: the 1999 guide page (olive `#cccc66`, khaki `#cccc99`, pale yellow, light grey, on white), with the red, black, link blue and navy of the 1998 community page and the dark green of the 1998 personal start page.
- `--color-page`, `--color-canvas`, `--color-surface` are white `#ffffff`; `--color-surface-alt` `#eeeeee` is the tab strip, alternate rows and the empty state; `--color-surface-strong` `#cccc99`.
- `--color-bar` `#cccc66` (olive, black text) is the search band, the table head and the header and footer rules. `--color-bar-alt` `#cccc99` (khaki, black text) is every title strip and the notice border.
- `--color-inverse` `#000000` with white text is the dialog title bar only.
- Links are `--color-link` `#0000cc`, visited `#6060a0`, active `#cc0000`; hover equals the link colour.
- `--color-heading-alt` `#003060` (navy) is the wordmark, tagline, h3 and positive figures. `--color-accent` and `--color-accent-alt` `#cc0000` are the mark, the solid badge, the "New!" word and negative figures.
- `--color-fill-1` to `--color-fill-4` are `#ffffcc`, `#eeeeee`, `#cccc99`, `#ffcc00`; `--color-fill-3` is also the khaki underline of section headings.
- `--color-warning` `#ffcc00` is a fill only and carries black text; `--color-success` `#006600` and `--color-danger` `#cc0000` carry white. Notices sit on pale yellow `#ffffcc`; the error notice is red on white. The secondary button is khaki `#cccc99`.
- Borders: `--color-border` `#a0a0a0`, `--color-border-strong` `#666633` (dark olive, for boxed outlines), `--color-border-muted` `#cccccc` and `--color-input-border` `#c0c0c0` for bevels. `--color-focus` is the navy `#003060`.

## Typography roles

- One wide screen sans-serif throughout: `--font-body`, `--font-heading` and `--font-ui` are Verdana, Helvetica, Arial, the stack of the 1999 guide page. `--font-mono` is Courier New.
- Text is small: `--text-base`, `--text-small` and `--text-h3` are all 10px, the size that carries most of the text on the 1999 guide page and the 1998 community page.
- `--text-ui` is 13px: the search band, form labels, controls and buttons stay at the size browsers gave form controls, a step above the 10px text around them, as on both pages.
- Titles are 13px bold (`--text-h2`, `--text-large`); `--text-h1` 16px; `--text-display` 18px bold (`--weight-display` 700) for the wordmark only.
- Line height `normal`; no transform, no tracking. All three link decorations are `underline`.

## Surface roles

- `--fill-panel` is `--color-surface-alt`: tables, figure rows, panels and dialogs stand on the light grey ground the 1999 guide page puts under its modules, so white cells show a grey gap between them.
- Everything is thin and dotted: `--border-style` is `dotted`, the rule the first stylesheets made possible: `--border-width` 1px for cells and outlines, `--border-width-strong` 2px for the header and footer rules, heading underlines and control bevels.
- Colour is dithered, as it was on a 256-colour screen: `--fill-bar`, `--fill-bar-alt`, `--fill-inverse` and `--fill-accent` are a 4px checker of the palette colour and a slightly darker (for the inverse, lighter) mix of it. `--fill-page` tiles a fine grey speckle over `--color-page`, which shows in the window beside the page column. The sheet and its panels are boxed: `--shadow-panel` is a 1px outline in `--color-border` with no offset and no blur; every other shadow is `none`. Radius 0; button and field fills flat.
- `--focus-ring` is `1px dashed` in `--color-focus`; `--transition` is `none`.

## Never

- `border-radius <= 0px`: no reference has a rounded corner.
- `box-shadow-blur <= 0px`: the only shadow is the 1px outline of the sheet and its panels; nothing is blurred.
- `text-shadow = none`: no reference has a text shadow.
- `gradient-fills <= 12%`: only the dithered bars, title strips and labels carry a pattern; every other fill is flat.
- `border-width <= 2px`: outlines are 1px and rules 2px; nothing is wider than a browser default field bevel.
- `transition = none`: nothing in the references changes on hover.
- `animation = none`: no reference has a CSS animation.
- `font-size >= 10px`: body and small print are 10px, nothing smaller.
- `font-size <= 18px`: the wordmark is the largest text; headings stop at 16px.
- `font-families <= 2`: Verdana and Courier.
- `font-weight >= 400`: only regular, bold and black weights were measured.
- `underlined-links >= 95%`: every reference underlines 100% of its link text.
- `letter-spacing = 0`: no reference tracks its text.
- `uppercase-text <= 0%`: no reference uses text-transform; capitals are typed.
