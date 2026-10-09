## Colour roles

- Family: the 1999 bookshop front, a khaki ground with a white sheet on it, pale yellow boxes, dark green and maroon, with the purple of the 1998 university page for visited links. It is the only set of this era whose page is not the colour of its content: `--color-page` is khaki `#cccc99` and `--color-canvas` white `#ffffff`, so the centred page stands in the window as a white sheet with khaki at both sides.
- `--color-surface` is white; `--color-surface-alt` `#ffffcc` (pale yellow) is alternate rows, quiet strips and the empty state; `--color-surface-strong` is the khaki.
- `--color-bar` `#cccc99` (khaki, black text, blue links) is the search band, the table head, panel titles and the header and footer rules. `--color-bar-alt` `#ffffcc` (pale yellow, black text) is every other title strip and the notice border.
- `--color-inverse` `#003300` (dark green, white text) is the dialog title bar or a top bar. The same green is `--color-heading-alt` (the brand name, h3, positive figures), `--color-accent` (the solid badge, white text), `--color-border-strong`, `--color-success` and `--color-focus`.
- Headings are maroon: `--color-heading` `#990000`. Text is black; `--color-text-muted` `#686868`.
- Links are `--color-link` `#003399`, always underlined; visited `#680060`; active `#aa0000`; hover equals the link colour.
- `--color-accent-alt` `#aa0000` is the "New!" word, prices and negative figures; `--color-danger` is the same red, as text on white and as a fill under white text.
- `--color-fill-1` to `--color-fill-4` are `#ffffcc`, `#cccc99`, `#cc9900` (the gold of the shop's section labels), `#ffffff`. They carry black text and blue links.
- `--color-warning` `#cc9900` is a fill only and carries black text. Notices sit on pale yellow; the error notice is red on white.
- Borders: `--color-border` `#989898`, `--color-border-muted` `#c0c0c0` for row rules and bevels, `--color-input-border` `#c0c0c0`. Buttons are the grey of the operating system, `#efefef` and `#c0c0c0`.

## Typography roles

- `--font-heading` is Arial Black, the face the bookshop sets its section labels in, falling back to Arial; it is used at `--weight-heading` and `--weight-display` 900, its only weight. `--font-body` and `--font-ui` are Arial, Helvetica. `--font-mono` is Lucida Console, falling back to Monaco and Courier New.
- `--text-base` and `--text-ui` 13px; `--text-small` 11px; `--text-h3` 13px, `--text-h2` and `--text-large` 16px, `--text-h1` 18px, `--text-display` 26px for the brand name only. Headings are heavy, not large.
- `--weight-bold` is 700 and so is `--weight-ui`: button and tab labels are bold. Line height is set like a printed catalogue, `--line-body` 1.3 and `--line-heading` 1.15; no transform, no tracking. All three link decorations are `underline`.

## Surface roles

- Rules are the double rules of a printed catalogue: `--border-style` is `double` and `--border-width` and `--border-width-strong` are both 3px, so every outline and rule is two hairlines a pixel apart.
- `--fill-page` is ruled like laid paper, a faint darker and a faint lighter line every 4px over `--color-page`, which shows in the window beside the page column. The sheet and its boxes are outlined and cast a hard edge in `--color-border` with no blur: `--shadow-panel` is a 1px outline with a 3px offset edge, `--shadow-dialog` 3px, `--shadow-control` and `--shadow-control-hover` 1px. Radius 0; every other fill is the flat palette colour and `--fill-panel` is `--color-surface`.
- `--focus-ring` is `3px double` in `--color-focus`; `--transition` is `none`.

## Never

- `border-radius <= 0px`: no reference has a rounded corner.
- `box-shadow-blur <= 0px`: shadows are hard offset edges; nothing is blurred.
- `text-shadow = none`: no reference has a text shadow.
- `gradient-fills <= 3%`: only the ruled page ground carries a pattern; every other fill is flat.
- `border-width <= 3px`: every rule and bevel is 3px or less.
- `transition = none`: nothing in the references changes on hover.
- `animation = none`: no reference has a CSS animation.
- `font-size >= 10px`: small print is 10px, nothing smaller.
- `font-size <= 26px`: the brand name is the largest text; headings stop at 18px.
- `font-families <= 3`: Arial, Arial Black and Lucida Console.
- `font-weight >= 400`: no light weights were measured.
- `underlined-links >= 95%`: every reference underlines 100% of its link text.
- `letter-spacing = 0`: no reference tracks its text.
- `uppercase-text <= 0%`: no reference uses text-transform; capitals are typed.
