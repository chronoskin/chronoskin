## Colour roles

- Broadsheet blues on white with one red. `--color-page`, `--color-canvas` and `--color-surface` are #ffffff. `--color-surface-alt` #cfdfef (pale blue) fills sidebars, forms and alternate rows; `--color-surface-strong` #95b7e0 is sidebar title strips and inactive tabs.
- `--color-bar` #006699 with white text is the link bar, panel titles, table heads and the current tab; `--color-bar-alt` #ff0000 with white text is the dateline strip: utility and footer strips and alternate panel titles. `--color-inverse` #003366 with white text is the one dark block.
- `--color-text` and `--color-heading` #000000; `--color-text-muted` #333333; `--color-heading-alt` #004371 for section titles and h3.
- `--color-link` and `--color-link-quiet` #003366 (dark navy, close to text), visited #336699, hover #cc0000, active #ff0000. `--color-accent` #f6dd42 with black text is the badge, the current navigation item and links on the dark block; `--color-accent-alt` #cc0000 marks promoted entries.
- Block fills are all cool: `--color-fill-1` #eaeff4, `--color-fill-2` #d0dbe8, `--color-fill-3` #cfdfef, `--color-fill-4` #eeeeee. Rules `--color-border` #6699cc, outlines `--color-border-strong` #003366, faint dividers #dddddd.
- `--color-button` #dddddd, secondary #eeeeee. `--color-danger` #cc0000 on `--color-danger-surface` #eeeeee. The palette has no green: `--color-success` is the mid blue #336699; `--color-warning` is the yellow #f6dd42. Both are badge fills only.

## Typography roles

- Sans-serif text under serif headlines: `--font-body` and `--font-ui` are Arial at `--text-base` 12px with `--line-body` 1.33 (12px on 16px); `--font-heading` is Times New Roman bold.
- `--text-display` and `--text-h1` 22px, `--text-h2` 16px, `--text-h3` 14px, `--text-large` 14px; `--text-small` 10px for meta lines and footers; `--text-ui` 12px bold for bars, tabs and buttons (`--weight-ui` 700). Heading lines are tight, 1.1.
- No uppercase and no tracking. Links in running text are underlined; quiet links, titles and navigation are not until hovered.

## Surface roles

- Ruled like a broadsheet. `--border-style` is dotted: every 1px rule and outline is a row of dots. `--border-width-strong` 3px: heavy rules are thick dotted lines and controls have a 3px bevel. Every radius is 0 and nothing casts a shadow at rest; a pressed button gains a 1px inset line.
- No gradients. `--fill-bar`, `--fill-bar-alt` and `--fill-inverse` are the plain colour underscored along the bottom edge by a light hairline over a dark 2px line, drawn as a pattern; `--fill-page` lays a faint newsprint dither over the page colour; `--fill-panel` is `--color-surface-alt`, so boxes are tinted. A hovered button takes `--color-surface-strong`. `--focus-ring` is 2px solid; `--transition` none.

## Never

- `palette-colours <= 20`: blues, one red, one yellow and greys; no green, no warm tint.
- `font-size <= 22px`: the largest text is a 22px serif headline.
- `font-size >= 10px`: nothing is set below 10px.
- `font-weight <= 700`: regular and bold only.
- `font-families <= 3`: a sans-serif for text, Times for headings, a monospace for code.
- `letter-spacing <= 0px`: no text is tracked.
- `uppercase-text <= 0%`: nothing is uppercased by CSS.
- `border-radius <= 0px`: every box is square.
- `border-width <= 3px`: 1px dotted rules, 3px heavy rules and bevels.
- `box-shadow = none`: nothing casts a shadow.
- `text-shadow = none`: text is flat.
- `gradient-fills <= 0%`: fills are flat colour.
- `transition = none`: states change instantly.
- `animation = none`: nothing moves.
