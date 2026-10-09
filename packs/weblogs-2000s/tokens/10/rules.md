
## Colour roles

- The early WordPress look: a white sheet on a light grey page. `--color-page` #e3e3e3, `--color-canvas` and `--color-surface` #ffffff, `--color-surface-alt` #f8f8f8 and `--color-surface-strong` #eeeeee. Almost everything is grey; colour is kept for links and one green.
- `--color-bar` #707070, a mid grey with white text, is the navigation bar or banner, table head and dialog title; `--color-bar-alt` #90b068, a soft leaf green with white text, is the one coloured strip (sidebar titles, date labels); `--color-inverse` #303030 with #e0e0e0 text is the footer.
- `--color-text` #383838, `--color-heading` #2d2d2d, `--color-text-muted` #888888; `--color-heading-alt` #555555 is sub-headings, dates and code.
- `--color-link` #192a72, a dark navy, underlined; visited #666666, hover and active #000000; `--color-link-quiet` #555555 is the grey link of menus and meta lines.
- Lines are `--color-border` #bbbbbb, `--color-border-muted` #e0e0e0 and `--color-border-strong` #000000: the black rule under an entry or on top of a hero.
- Fills are greys (#f8f8f8, #eeeeee, #e0e0e0, #f0f0f0). `--color-accent` #90b068 with white text is the count badge, the current day and bullets; `--color-accent-alt` #686868 is the "new" word.
- Buttons are the operating system's grey: `--color-button` #eeeeee with #2d2d2d text, the secondary #f8f8f8. Status colours are the era's usual ones, since these references show none: `--color-success` #538620, `--color-warning` #cc6600, `--color-danger` #cc0000.

## Typography roles

- `--font-body` is "Lucida Grande" (then Verdana, Arial) at `--text-base` 11px with `--line-body` 1.5; `--text-small` 10px is meta lines, sidebars and tables.
- `--font-heading` is Georgia, bold and set tight (`--weight-heading` and `--weight-display` 700, `--heading-tracking` and `--display-tracking` -1px): `--text-display` 25px, `--text-h1` 24px, `--text-h2` 15px, `--text-h3` 11px. Big, heavy serif titles over very small sans text are the character of this set.
- `--font-ui` is Verdana at `--text-ui` 10px, bold and in capitals (`--weight-ui` 700, `--ui-transform` uppercase): navigation, tabs and buttons read as small labels.
- Links in running text are underlined (`--link-decoration` underline) and lose the underline on hover; quiet links, titles and navigation are not underlined. `--font-mono` is Courier New.

## Surface roles

- Etched frames: `--border-width` 2px with `--border-style` groove, so every outline is a cut line with a light and a dark edge, and control bevels are 4px. `--border-width-strong` 5px is the heavy rule on top of a sheet or a hero.
- Bars, the footer strip and the primary button are bevelled, not shaded: `--fill-bar`, `--fill-bar-alt`, `--fill-inverse` and `--fill-button` are the flat colour with a 1px highlight along the top and a 1px shade along the bottom. Page, panels, inputs and the secondary button are flat.
- All radii are 0, all shadows `none`; `--focus-ring` is a 1px dotted outline; `--transition` none.

## Never

- `border-radius <= 0px`: every block is square.
- `box-shadow = none`: nothing casts a shadow.
- `text-shadow = none`: text is flat.
- `gradient-fills <= 20%`: bars and buttons carry a 1px bevel edge drawn as a gradient; nothing is shaded across its height.
- `border-width <= 5px`: lines are 1px, bevels 2px, the heavy rule 5px.
- `font-size <= 25px`: the largest title measured is 25px.
- `font-size >= 9px`: the smallest text measured is 9px; the set itself stops at 10px.
- `font-weight <= 700`: bold is the heaviest weight.
- `font-families <= 4`: Lucida Grande, Georgia, Verdana for controls and a monospace.
- `line-height >= 1.3`: body text is set at 1.5.
- `line-height <= 1.6`: and never looser.
- `transition = none`: states change instantly.
- `animation = none`: nothing moves.
