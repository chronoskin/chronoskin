## Colour roles

- Family: bright tile colours on dark navy. Every value is measured in a reference: the navies and the pale text of the 2015 customer-messaging page (`#212b3a`, `#0b253b`, `#1e2024`, `#d1dae8`), the slate of the 2014 team-chat page (`#3d4b5b`), the tile colours of the 2014 tile-framework page (red `#e51400`, amber `#f0a30a`, green `#60a917`, steel `#647687`), the cyan and orange of the 2013 phone-platform page (`#00bcf2`, `#ff8c00`) and the green, blue and night blue of the 2014 language-course page (`#7eb530`, `#1caff6`, `#0b3e71`). Dark first screens with bright flat blocks are also what the 2014 ride-app and 2014 payments pages show.
- The page is dark: `--color-page`, `--color-canvas` and `--color-surface` are `#212b3a`; `--color-surface-alt` `#0b253b` is the alternate band and row (darker, not lighter), `--color-surface-strong` `#3d4b5b` the title strip. `--color-bar` `#3d4b5b` with white `--color-bar-text` is the navigation; `--color-bar-alt` `#0b3e71` with white text the secondary strip.
- Body text is `--color-text` `#d1dae8`, meta `--color-text-muted` `#999999`, headings white; small titles are blue `--color-heading-alt` `#1caff6`.
- Green acts, blue links: `--color-button` `#7eb530` with white text; `--color-link` `#1caff6`, hover `#00ccff`, active `#3ea5ce`. The secondary button is the slate `--color-button-secondary` `#3d4b5b` with white `--color-button-secondary-text` and a white edge: its text and edge must contrast with the dark page as well as with its own fill, because a surface set may make `--fill-button-secondary` transparent.
- `--color-fill-1` `#00bcf2` (cyan) is the hero and the emphasis band with `--color-inverse-text` white on it; `--color-fill-2` `#e51400`, `--color-fill-3` `#f0a30a` (icon only) and `--color-fill-4` `#60a917` are the other tiles.
- `--color-accent` `#e51400` (white text) is the default badge, `--color-accent-alt` amber `#f0a30a` the second. Status: `--color-success` `#7eb530`, `--color-warning` `#ff8c00`; `--color-danger` is the lighter red `#eb4d5c` of the team-chat page, because the tile red is too dark for small text on navy.
- Form controls stay light on the dark page: `--color-input` white with `#3c3c3c` text and a `#d1dae8` border. The notice is `--color-notice` `#4d9de0` with white text; the error notice is red on the pale `--color-danger-surface` `#f4fafc`.
- Borders are `--color-border` and `--color-border-muted` `#3d4b5b`, `--color-border-strong` `#647687`. `--color-inverse` `#1e2024` with white text is the footer: darker than the page, not a light block. `--color-shadow` is black at 25%, `--color-overlay` black at 50%.

## Typography roles

- One rounded sans at medium weight, after the 2014 language-course page (which used a commercial rounded face): `--font-body`, `--font-heading` and `--font-ui` are Nunito, Varela Round, then Helvetica Neue. `--font-mono` is Monaco.
- `--text-base` 17px at `--line-body` 1.65, `--text-small` 13px, `--text-ui` 14px, `--text-large` 22px.
- Headings are medium, not light and not black: `--weight-heading` and `--weight-body` 500 at `--text-h1` 30px, `--text-h2` 22px, `--text-h3` 18px. The hero title and the big figures are `--text-display` 36px at `--weight-display` 700.
- `--weight-ui` and `--weight-bold` are 700. Sentence case, no tracking: both transforms `none`, both tracking tokens 0. All three link decorations are `none`.

## Surface roles

- Thick lines: `--border-width` 2px for every rule and outline, `--border-width-strong` 3px for the current item.
- `--radius-control` and `--radius-panel` 6px, `--radius-pill` 4px (badges are rounded tags, not pills), `--radius-page` 0.
- Everything that can be picked up stands on a hard lower edge, after the 2015 board-tool page: `--shadow-control` and `--shadow-control-hover` are `0 2px 0 0` in `--color-shadow` under buttons and inputs, `--shadow-panel` the same under panels, `--shadow-dialog` `0 4px 0 0` under a dialog; pressed is `none`. No shadow is blurred and text has none.
- Buttons are solid: `--fill-button` is the flat `--color-button`, hovered mixed 20% with white; `--fill-button-secondary` is the solid `--color-button-secondary`. `--focus-ring` is `2px solid` in `--color-focus`. `--transition` is `none`.

## Never

- `box-shadow-blur <= 0px`: the only shadows are hard edges under controls, panels and dialogs.
- `text-shadow = none`: text is flat on every fill.
- `gradient-fills <= 0%`: every background is one solid colour.
- `border-radius <= 6px`: 6px on controls and panels, 4px on tags.
- `border-width <= 3px`: rules are 2px; 3px marks the current item.
- `font-weight >= 400`: this type set has no light weight.
- `font-weight <= 700`: the hero title and labels are the heaviest text.
- `font-size >= 13px`: meta text and tags are the smallest text.
- `font-size <= 36px`: the hero title is the largest text.
- `font-families <= 2`: one sans-serif family, plus monospace for code.
- `line-height <= 1.65`: body text is set at 1.65.
- `letter-spacing = 0`: text is not tracked.
- `uppercase-text <= 0%`: everything is in sentence case.
- `underlined-links <= 0%`: no reference underlines links.
- `transition = none`: states change without a fade.
- `animation = none`: no reference has a CSS animation.
