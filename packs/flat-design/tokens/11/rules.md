## Colour roles

- Family: the pale blues of the 2015 customer-messaging page (`#f0f6fb`, `#e9edf0`, `#d5eafb`, its soft blue `#4d9de0` and green `#4fb67d`), with the greys and the sage green of the 2014 publishing-platform page, the pale band and the indigo of the 2015 board-tool page and the coral of the 2014 newsletter page.
- The page is tinted, not white: `--color-page` and `--color-canvas` are `#f0f6fb`. `--color-surface` white is for panels, table cells and the navigation (`--color-bar`); `--color-surface-alt` `#e4f0f6` is the alternate band and row, `--color-surface-strong` `#d5eafb` the title strip. `--color-bar-alt` `#e9edf0` with `#6b8499` text is the secondary strip.
- Text is soft: `--color-text` `#7d878a`, `--color-text-muted` `#999999`, headings `--color-heading` `#35393b`; small titles are blue `--color-heading-alt` `#5ba4e5`.
- Green acts, blue links: `--color-button` `#4fb67d` with white text; `--color-link` `#4d9de0`, hover `#2c9ab7`, active `#42548e`. The secondary button is a blue outline with no fill.
- `--color-fill-1` `#4d9de0` is the hero and the emphasis band with white text; `--color-fill-2` `#9fbb58` (sage), `--color-fill-3` `#6b8499` (steel) and `--color-fill-4` `#42548e` (indigo) are the other tiles.
- `--color-accent` coral `#e85c41` (white text) is the default badge, `--color-accent-alt` `#d5eafb` the second. `--color-danger` is the same coral; the error notice sits on white. `--color-success` `#4fb67d`, `--color-warning` `#f0a30a` (the amber of the 2014 tile-framework page). The notice is `--color-notice` `#d5eafb` with indigo text.
- Borders are hairlines: `--color-border` `#d9dee1`, `--color-border-muted` `#e9edf0`, `--color-border-strong` and `--color-input-border` `#d1dae8`. `--color-inverse` `#212b3a` with white text is the footer. `--color-shadow` is black at 7%, `--color-overlay` the inverse colour at 45%.

## Typography roles

- One thin sans, after the 2014 publishing-platform page and the 2014 list-app page: `--font-body`, `--font-heading` and `--font-ui` are Open Sans, then Helvetica Neue. `--font-mono` is Menlo.
- Body is light: `--weight-body` 300 at `--text-base` 16px and `--line-body` 1.5; `--text-large` 20px, `--text-ui` 14px, `--text-small` 12px.
- Headings are thinner than the body and large: `--weight-heading` and `--weight-display` 200 at `--text-display` 48px, `--text-h1` 36px, `--text-h2` 26px, `--text-h3` 18px. Hierarchy is size and colour, never weight.
- `--weight-ui` is 400: buttons, tabs and navigation are regular. `--weight-bold` 600 is for labels and table heads only.
- Sentence case, no tracking: both transforms `none`, both tracking tokens 0. All three link decorations are `none`.

## Surface roles

- `--border-width` 1px, `--border-width-strong` 2px for the current item.
- Soft corners everywhere: `--radius-control` 4px, `--radius-panel` 6px, `--radius-pill` 500px for badges, `--radius-page` 8px for tiles, pictures and grid boxes.
- This is the late, softened flat of 2015 and 2016: panels lift off the tinted page with `--shadow-panel` `0 1px 3px 0` in the faint `--color-shadow`, and a dialog floats on `--shadow-dialog` `0 12px 32px 0`. Controls and text stay shadowless.
- `--fill-button` is the flat `--color-button`, hovered mixed 15% with white; `--fill-button-secondary` is `transparent`, so the secondary button is an outline on whatever it stands on. `--focus-ring` is a soft halo: `3px solid` in `--color-focus` at 40%. `--transition` fades colour, background and border colour in 0.3s.

## Never

- `box-shadow-blur <= 32px`: a panel's shadow is a 3px blur and a dialog's 32px; controls have none.
- `text-shadow = none`: text is flat on every fill.
- `gradient-fills <= 0%`: every background is one solid colour.
- `border-radius <= 8px`: 4px on controls, 6px on panels, 8px on tiles; badges are pills.
- `border-width <= 3px`: rules are 1px, 2px marks the current item, and the focus halo is 3px.
- `font-weight <= 600`: headings are thin; 600 is only for labels.
- `font-size >= 12px`: meta text and badges are the smallest text.
- `font-size <= 48px`: the hero title is the largest text.
- `font-families <= 2`: one sans-serif family, plus monospace for code.
- `line-height <= 1.5`: body text is set at 1.5.
- `letter-spacing = 0`: text is not tracked.
- `uppercase-text <= 0%`: everything is in sentence case.
- `underlined-links <= 0%`: links are marked by colour alone.
- `animation = none`: no reference has a CSS animation.
