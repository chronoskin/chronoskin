## Colour roles

- Brushed silver with one blue. `--color-page` #686868 surrounds a pale grey sheet, `--color-canvas` #e8e8e8; `--color-surface` #ffffff is panel bodies and table rows, `--color-surface-alt` #f0f0f0 forms, the link column and alternate rows, `--color-surface-strong` #d0d0d0 title strips inside the link column.
- `--color-bar` #686868 with white text is the navigation bar, panel and dialog titles, table heads and the current tab; `--color-bar-alt` #d0d0d0 with black text is the utility strip, section tabs, inactive tabs and the footer strip. `--color-inverse` #003098, a deep blue, with white text is the one blue panel and the closing band.
- `--color-text` and `--color-heading` #000000; `--color-text-muted` #585858; `--color-heading-alt` #0060c8.
- `--color-link` #0060c8; `--color-link-quiet` #505050; visited #686868, hover #000000, active #1870c8. `--color-accent` #0060c8 with white text is the current navigation item, badges and the current page number; `--color-accent-alt` #1870c8 is a lamp colour.
- `--color-fill-1` to `--color-fill-4` are four greys (#f0f0f0, #e0e0e0, #d0d0d0, #f8f8f8): there are no coloured blocks. Lines are grey: `--color-border` #b8b8b8, `--color-border-strong` #686868, `--color-border-muted` #d8d8d8.
- `--color-button` #0868c8 with white text is the one saturated control; `--color-button-secondary` #d0d0d0 with black text. Inputs are white with a #686868 line.
- `--color-danger` #800000 is only used as text on `--color-danger-surface` #ffffff and as a rim or lamp; `--color-success` #709000 and `--color-warning` #ffc800 are lamp and rim colours, never text. `--color-notice` is white with black text.

## Typography roles

- Small bitmap-style text with tracked capitals. `--font-body` and `--font-ui` are Geneva, then Verdana; `--font-heading` is Arial bold, uppercased by `--heading-transform` and tracked by `--heading-tracking` 1px (`--display-tracking` 2px for the wordmark and hero title).
- `--text-base` 10px, `--text-small` and `--text-ui` 9px, `--text-large` 12px. `--text-h3` 10px, `--text-h2` 12px, `--text-h1` 14px, `--text-display` 18px.
- Interface text is 9px bold capitals (`--ui-transform` uppercase, `--weight-ui` 700): navigation, tabs, buttons, table heads.
- Links in running text are underlined; quiet and navigation links are not (`--link-decoration-quiet` none).

## Surface roles

- Pinstriped metal and square boxes. `--fill-page`, `--fill-bar` and `--fill-bar-alt` are 1px horizontal pinstripes of the colour and a slightly lighter or darker mix of it; every other fill is flat.
- Every radius is 0. `--border-width` and `--border-width-strong` are both 1px, solid.
- `--shadow-control` is a 1px inset highlight on the top and left of buttons; `--shadow-control-pressed` darkens it. `--shadow-dialog` is a hard 2px offset without blur. `--shadow-panel` and `--shadow-text` are `none`. `--focus-ring` is 1px dotted; `--transition` none.

## Never

- `palette-colours <= 18`: greys, white, black, three blues, plus status lamps.
- `font-size <= 18px`: the wordmark and hero title are the largest text.
- `font-size >= 9px`: nothing is set below 9px.
- `font-weight <= 700`: regular and bold only.
- `font-families <= 3`: Geneva or Verdana, Arial for headings and a monospace.
- `letter-spacing <= 2px`: headings are tracked by 1px, display text by 2px.
- `border-radius <= 0px`: every box is square.
- `border-width <= 1px`: every line is a hairline.
- `box-shadow-blur <= 0px`: the bevel and the dialog offset are hard.
- `text-shadow = none`: text is flat.
- `transition = none`: states change instantly.
- `animation = none`: nothing moves.
