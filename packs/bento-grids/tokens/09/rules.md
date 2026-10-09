## Colour roles

- A blue-black slate ramp: `--color-page`, `--color-canvas`, `--color-bar` and `--color-input` `#030712`; cells `--color-surface` `#111827`; panels, the dialog, the secondary button and current navigation items `--color-surface-alt` `#1f2937`; current tab, neutral badges, chart bars and inline code `--color-surface-strong` `#374151`. `--color-bar-alt` `#0f172a` is the table head and tab row.
- Text is white with sky-tinted captions: `--color-heading` `#ffffff`, `--color-text` and `--color-heading-alt` `#f1f5f9`; `--color-text-muted`, `--color-link-quiet` and `--color-bar-alt-text` are pale sky `#bae6fd` at 70%, so every caption carries a little of the accent.
- The primary button is light, not coloured: `--color-button` `#f3f4f6` with `--color-button-text` `#030712`. `--color-inverse` is the same pair, used for the button on the accent cell and the toggle knob.
- Sky blue is the only hue: `--color-link` and `--color-focus` `#38bdf8` (links, icons, eyebrow, hero accent phrase), `--color-accent` `#0284c7` with white text (accent badge, highlighted chart bar, call-to-action cell, hero glow), `--color-accent-alt` `#0ea5e9` ("new", checks, switched-on toggles), `--color-link-visited` `#bae6fd`.
- Feature cells: `--color-fill-1` `#0c4a6e` (deep sky), `--color-fill-2` `#1a3045` (steel), `--color-fill-3` `#374151` and `--color-fill-4` `#1f2937` (slate).
- Borders are white at low alpha: `--color-border` 10%, `--color-border-strong` 20%; `--color-border-muted` `#1f2937` and `--color-input-border` `#374151` are solid.
- Status: `--color-success` `#3ecf8e`, `--color-warning` `#f7be2b`, `--color-danger` `#f24822` on `--color-danger-surface` `#0f172a`; the notice is `--color-notice` `#1a3045` with `--color-notice-text` `#bae6fd`. `--color-shadow` is the page colour `#030712`, `--color-overlay` `#111827` at 80%.

## Typography roles

- One grotesque, Inter first with the platform display sans behind it, for `--font-body`, `--font-heading` and `--font-ui`; `--font-mono` is Geist Mono or JetBrains Mono, then the system mono.
- Medium throughout: `--weight-heading`, `--weight-display` and `--weight-ui` are all 500, `--weight-body` 400, `--weight-bold` 600. No title is bold.
- Large but calm headings set tight: `--text-display` 64px at `--line-display` 1.06, `--text-h1` 40px, `--text-h2` 32px, `--text-h3` 20px at `--line-heading` 1.1.
- `--text-base` 16px at an open `--line-body` 1.6; `--text-large` 17px for the lead and cell titles; `--text-small` 14px for captions; `--text-ui` 13px makes navigation, buttons, tabs and form controls a size smaller than the captions.
- Links carry no underline (`--link-decoration` none) and gain one on hover. No tracking, no transforms.

## Surface roles

- Big soft corners: `--radius-page` 32px on cells, `--radius-panel` 24px on panels, tables, notices and the dialog, `--radius-control` 13px on buttons, inputs and icon tiles, `--radius-pill` for badges, tabs and switches.
- Cells float on a long soft drop: `--shadow-panel` is a 2px contact shadow at 8% plus `0 18px 30px -10px` of `--color-shadow` at 22%, so a cell reads as a cushion lifted off the page. Controls take 4px at 10% (`--shadow-control`), a 12px drop at 16% on hover, and a 4px inset shade when pressed. `--shadow-dialog` stacks three drops up to a 31px blur at 14%.
- Fills are flat: `--fill-panel`, `--fill-button` and the rest are the plain colour tokens, with no gradient. Only `--fill-bar` is translucent, `--color-bar` at 72%, for a layout that blurs what scrolls under its bar.
- `--border-width` is 1px and takes whatever hairline the palette gives; `--border-width-strong` is 2px. `--focus-ring` is a wide soft halo: 3px of `--color-focus` at 50%. `--fill-button-hover` moves the button 15% toward its own text colour.
- `--transition` is a 0.2s fade of background, colour and shadow.

## Never

- `palette-colours <= 36`: slate, white and sky blue; the three status colours appear only as small marks.
- `font-weight <= 600`: titles are medium; 600 is inline emphasis only.
- `font-size <= 64px`: the hero heading is the largest text, 64px.
- `font-size >= 12px`: the smallest text is 12px.
- `font-families <= 2`: one sans for everything plus the mono for code.
- `underlined-links <= 1%`: links are told apart by colour; underline appears only on hover.
- `line-height <= 1.7`: body text sits at 1.6.
- `border-radius <= 32px`: the largest box radius is the 32px cell.
- `border-radius >= 4px`: no box is square.
- `border-width <= 2px`: every hairline is 1px; 2px exists only as the focus ring.
- `box-shadow-blur <= 31px`: the dialog reaches a 31px blur and a cell 30px; nothing is blurred further.
- `text-shadow = none`: no text shadows.
- `animation = none`: nothing animates.
