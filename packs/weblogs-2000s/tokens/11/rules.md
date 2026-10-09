
## Colour roles

- The hosted-journal default of the period: a white sheet on a periwinkle page. `--color-page` #6666cc, `--color-canvas` and `--color-surface` #ffffff, `--color-surface-alt` #eeeeff (pale lavender) and `--color-surface-strong` #c0c0ff (lavender).
- `--color-bar` #c0c0ff with black text is the navigation bar or banner, table head and dialog title: a light bar, not a dark one. `--color-bar-alt` #eeeeff with navy #000050 text is the paler strip (sidebar titles, date labels); `--color-inverse` #6666cc with white text is the footer.
- `--color-text` and `--color-heading` are #000000, `--color-text-muted` #707070; `--color-heading-alt` #c00000, a plain red, is sub-headings, dates and code, as entry subjects were.
- `--color-link` #000050, a very dark navy, always underlined; visited #8b1a1a, hover and active #c00000, quiet #484848.
- Lines are `--color-border` #c0c0ff, `--color-border-muted` #eeeeff and `--color-border-strong` #6666cc: boxes are drawn in the lavenders, never in grey.
- Fills: `--color-fill-1` #eeeeff, `--color-fill-2` #f0f0f8, `--color-fill-3` #c0c0ff, `--color-fill-4` #f8f8f8. `--color-accent` #6666cc with white text is the count badge and the current day; `--color-accent-alt` #c00000 is the "new" word.
- Buttons are pale lavender: `--color-button` #eeeeff with #000050 text, the secondary #f8f8f8 with black. `--color-danger` is the measured red #c00000; `--color-success` #538620 and `--color-warning` #cc6600 are the era's usual ones, since the references show neither.

## Typography roles

- `--font-body` is Times at the browser's size: `--text-base` 16px with `--line-body` 1.375 (22px). It is the unstyled-serif look of hosted journals and early default themes; do not swap it for a sans.
- `--font-heading` and `--font-ui` are Arial: `--text-display` 24px bold, `--text-h1` 18px bold, `--text-h2` 16px bold, `--text-h3` 13px bold. `--text-small` and `--text-ui` are 13px. Nothing is uppercased or tracked.
- Every link is underlined, at rest and on hover, quiet links and navigation included (`--link-decoration`, `--link-decoration-hover` and `--link-decoration-quiet` all underline). `--font-mono` is Courier New.

## Surface roles

- Dashed and tiled: `--border-style` dashed at `--border-width` 1px for every outline; `--border-width-strong` 3px is the rule on top of a sheet or a hero.
- `--fill-bar`, `--fill-bar-alt` and `--fill-inverse` are scanlines: every third pixel row is 12% darker. `--fill-page` is an 8px check over `--color-page`. Controls and panels are flat.
- Panels and the dialog stand on a hard 2px and 4px offset in `--color-border` (`--shadow-panel`, `--shadow-dialog`), without blur. Tags and badges are eased by `--radius-pill` 6px; every other radius is 0.
- `--focus-ring` is a 1px dotted outline; `--transition` none.

## Never

- `border-radius <= 6px`: only tags and badges are eased; every block is square.
- `box-shadow-blur <= 0px`: the offset under a panel is hard, never soft.
- `text-shadow = none`: text is flat.
- `gradient-fills <= 20%`: only bars and the footer strip carry scanlines; tints are flat.
- `border-width <= 3px`: lines are 1px, bevels 2px, the top rule 3px.
- `font-size <= 32px`: the largest text measured is 32px; the set itself stops at 24px.
- `font-size >= 10px`: the smallest text measured is 10px; the set itself stops at 13px.
- `font-weight <= 700`: bold is the heaviest weight.
- `font-families <= 3`: Times, Arial and a monospace.
- `line-height >= 1.1`: text is set close, at 1.375.
- `line-height <= 1.5`: and never open.
- `underlined-links >= 60%`: links are underlined everywhere.
- `transition = none`: states change instantly.
- `animation = none`: nothing moves.
