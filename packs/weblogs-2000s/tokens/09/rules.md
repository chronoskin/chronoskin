## Colour roles

- White throughout: `--color-page`, `--color-canvas` and `--color-surface` are #ffffff, `--color-surface-alt` #f8f8f8 and `--color-surface-strong` #efefef; the sheet shows only by its 1px #c0c0c0 sides.
- `--color-bar` #cc0000 with white text is the navigation bar, table head and dialog title; `--color-bar-alt` #000000 with white text is the sidebar strips and the band on top of the sheet; `--color-inverse` #282828 with #e0e0e0 text is the footer.
- `--color-text` and `--color-heading` are #000000, `--color-text-muted` #666666; `--color-heading-alt` #cc0000 is dates, sub-headings, code and the second half of the wordmark.
- `--color-link` #0000ee, the browser's blue; visited #173797, hover #cc0000, active #ff0000, quiet #808080.
- Lines are `--color-border` #c0c0c0, `--color-border-muted` #e0e0e0 and `--color-border-strong` #000000 (hero rule, notice outline).
- The one loud tint is #dfff0c: `--color-notice` and `--color-fill-3`. The other fills are greys (#efefef, #f8f8f8, #e0e0e0). `--color-accent` #ff0000 with white text is the count badge and bullets.
- Status: `--color-success` #006d00, `--color-warning` #cc6600, `--color-danger` #ff0000. Buttons are grey #e0e0e0 and #f0f0f0 with black text.

## Typography roles

- `--font-body` is Verdana, "Bitstream Vera Sans" at `--text-base` 11px with `--line-body` 1.8, the loosest setting of the era; `--text-small` 10px is meta, sidebar and tables.
- `--font-heading` and `--font-ui` are "Gill Sans" (Gill Sans Std, Gill Sans MT, then Verdana). Headings are regular weight (`--weight-heading`, `--weight-display` 400): `--text-display` 22px tracked 1px (`--display-tracking`), `--text-h1` 17px, `--text-h2` 12px, `--text-h3` 11px.
- Navigation, tabs and buttons are `--text-ui` 12px Gill Sans in capitals (`--ui-transform` uppercase). Headings are not uppercased.
- Links in running text are underlined (`--link-decoration` underline); quiet links, titles and navigation are not (`--link-decoration-quiet` none) and gain the underline on hover. `--font-mono` is Monaco.

## Surface roles

- Heavy and hard-edged: bars, bands, page and controls are their plain palette colour; no gradient, no pattern.
- `--border-width` 1px solid for outlines; `--border-width-strong` 10px is the heavy band on top of the sheet and of the hero.
- Boxes stand on hard offset shadows without blur, the print-like "drop shadow" drawn with a darker right and bottom edge: `--shadow-panel` 3px, `--shadow-dialog` 5px, buttons 1px (`--shadow-control`), 2px when hovered and pressed in by an inset 1px.
- All radii are 0; `--focus-ring` is a 2px solid outline; `--transition` none.

## Never

- `border-radius <= 0px`: every block is square.
- `box-shadow-blur <= 0px`: shadows are hard offsets of 1px to 5px, never soft.
- `text-shadow = none`: text is flat.
- `gradient-fills <= 0%`: every bar and band is one flat colour.
- `border-width <= 10px`: lines are 1px, bevels 2px, the top band 10px.
- `font-size <= 22px`: the wordmark is 22px; titles are 17px regular.
- `font-size >= 10px`: meta text stops at 10px.
- `font-families <= 3`: Verdana, Gill Sans and a monospace.
- `line-height >= 1.7`: body text is set very open, at 1.8.
- `line-height <= 1.8`: the loosest measured setting.
- `letter-spacing <= 2px`: the wordmark is tracked 1px, sidebar titles 2px.
- `uppercase-text <= 12%`: capitals are for navigation, tabs, buttons and sidebar titles only.
- `transition = none`: states change instantly.
- `animation = none`: nothing moves.
