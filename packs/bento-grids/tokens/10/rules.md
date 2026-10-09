## Colour roles

- A neutral near-black ramp with no hue in it: `--color-page`, `--color-canvas`, `--color-bar` and `--color-input` `#0d0d0d`; tiles `--color-surface` `#121212`; panels, the dialog, the secondary button, current navigation items and `--color-bar-alt` `#1c1c1c`; current tab, neutral badges, chart tracks and inline code `--color-surface-strong` `#262a2d`.
- Text is off-white and grey: `--color-heading` and `--color-bar-text` `#fafafa`, `--color-heading-alt` `#f0f0f0`, `--color-text` `#ebeced`; `--color-text-muted`, `--color-link-quiet` and `--color-bar-alt-text` `#a1a4a5`; `--color-disabled-text` `#6c6c6c`.
- One vivid accent, orange `#ff9800`, with near-black text on it: `--color-accent`, `--color-button`, `--color-link` and `--color-focus`. It carries the primary button, links, icons, the highlighted chart mark, the accent badge and the call-to-action tile, and nothing else on the page has a hue of its own.
- `--color-accent-alt` and `--color-warning` are the yellow `#ffd60a` of the same family: second chart series, "new" markers, switched-on toggles, the active link.
- Feature tiles stay neutral: `--color-fill-1` `#1c1c1c`, `--color-fill-2` `#0b0e14` (the one cold black), `--color-fill-3` `#141517`, `--color-fill-4` `#191b1e`. Tiles differ by a step of grey, never by colour.
- `--color-inverse` `#fafafa` with `--color-inverse-text` `#121212` is the single white tile or button.
- Borders are solid dark greys: `--color-border` and `--color-input-border` `#262a2d`, `--color-border-muted` `#1c1c1c`; `--color-border-strong` is a cold white at 19%.
- Status: `--color-success` `#4caf50`, `--color-warning` `#ffd60a`, `--color-danger` `#ff6465` on `--color-danger-surface` `#191b1e`; the notice is `--color-notice` `#0b0e14` with `--color-notice-text` `#ebeced`. `--color-shadow` is black, `--color-overlay` black at 70%.

## Typography roles

- A serif for headings against a plain sans for everything else: `--font-heading` is Hedvig Letters Serif, then Instrument Serif, Iowan Old Style and Georgia; `--font-body` and `--font-ui` are Hedvig Letters Sans, then Inter and the system sans. One reference set its display type in a commercial serif; these open faces stand in for it. `--font-mono` is Commit Mono or Geist Mono, then the system mono.
- Regular throughout: `--weight-heading`, `--weight-display`, `--weight-ui` and `--weight-body` are 400, `--weight-bold` 500. Size and the serif carry the hierarchy, not weight.
- `--text-display` 56px at `--line-display` 1.1 with `--display-tracking` -1px; `--text-h1` 40px, `--text-h2` 28px, `--text-h3` 20px at `--line-heading` 1.2.
- Small, open body text: `--text-base` 15px at `--line-body` 1.6; `--text-small` and `--text-ui` 14px; `--text-large` 18px.
- Links carry no underline (`--link-decoration` none) and gain one on hover. No transforms.

## Surface roles

- Tight corners: `--radius-page` 10px on tiles, `--radius-panel` 6px, `--radius-control` 4px, `--radius-pill` for badges and switches.
- Every tile sits in a bezel: `--shadow-panel` is a solid 3px ring of `--color-surface-strong` at 55% outside the 1px hairline, with no blur. The gaps of a grid read narrower for it, as if the tiles were set into a frame.
- Controls take the same frame, thinner: `--shadow-control` is a 2px ring, 3px at 80% on hover, none when pressed. The dialog is lifted by a solid 8px ring of `--color-shadow`.
- Every fill is flat: `--fill-panel`, `--fill-bar`, `--fill-button` and the rest are the plain colour tokens. `--fill-button-hover` moves the button 15% toward white.
- `--border-width` 1px, `--border-width-strong` 2px for the focus ring. `--transition` is a 0.15s colour fade.

## Never

- `palette-colours <= 36`: greys, one orange and its yellow; the status colours appear only as small marks.
- `font-weight <= 500`: nothing is bold; 500 is inline emphasis only.
- `font-size <= 56px`: the hero heading is the largest text.
- `font-size >= 12px`: the smallest text is 12px.
- `font-families <= 3`: the serif, the sans and the mono.
- `underlined-links <= 1%`: links are told apart by the accent colour; underline appears only on hover.
- `line-height <= 1.7`: body text sits at 1.6.
- `border-radius <= 10px`: the largest box radius is the 10px tile.
- `border-width <= 2px`: every hairline is 1px; 2px exists only as the focus ring.
- `box-shadow-blur <= 0px`: no shadow is blurred; the bezel and the dialog ring are solid.
- `text-shadow = none`: no text shadows.
- `animation = none`: nothing animates.
