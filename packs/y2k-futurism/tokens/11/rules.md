## Colour roles

- Black with acid green, one aqua and one blue. `--color-page` #000000 surrounds a near-black sheet: `--color-canvas` and `--color-surface` #181818, `--color-surface-alt` #282828 for forms, the link column and alternate rows. `--color-surface-strong` #000000 is darker than the sheet, for title strips inside the link column.
- `--color-bar` #98ff00 with black text is the navigation bar, panel and dialog titles, table heads and the current tab; `--color-bar-alt` #484848 with white text is the utility strip, section tabs, inactive tabs and the footer strip. `--color-inverse` #ffffff with black text is the one white panel and the closing band.
- `--color-text` #e0e0e8; `--color-text-muted` #c0c0c0; `--color-heading` #ffffff; `--color-heading-alt` #98ff00.
- `--color-link` #98ff00; `--color-link-quiet` #e0e0e8; visited #98d0d0, hover #ffffff, active #a8ff00. `--color-accent` #98d0d0, aqua, with black text is the current navigation item, badges and the current page number; `--color-accent-alt` #a8ff00 is a lamp colour.
- `--color-fill-1` to `--color-fill-3` are near-blacks and a grey (#282828, #202020, #484848) and `--color-fill-4` is a steel blue #385090. Lines: `--color-border` #686868, `--color-border-strong` #98ff00, `--color-border-muted` #484848.
- `--color-button` #98ff00 with black text; `--color-button-secondary` #686868 with white text. Inputs are black with green text and a #989898 line, like a terminal field.
- `--color-danger` #ff0000 is only used as text on `--color-danger-surface` #ffffff and as a rim or lamp; `--color-success` #98ff00 and `--color-warning` #ffc800 are lamp and rim colours, never text. `--color-notice` is the steel blue with white text.

## Typography roles

- Verdana for text, a monospace for headings and for the interface, like a terminal read-out. `--font-heading` is Courier New bold, uppercased by `--heading-transform`; `--font-ui` is Courier New too, so navigation, tabs, buttons, inputs and table heads are monospaced; `--font-body` is Verdana.
- `--text-base` 10px, `--text-ui` 11px (the monospace runs small, so it is set one step up), `--text-small` 9px, `--text-large` 12px. `--text-h3` 11px, `--text-h2` 13px, `--text-h1` 16px, `--text-display` 18px.
- Interface text is regular weight and not uppercased; nothing is tracked.
- Links are not underlined (`--link-decoration` and `--link-decoration-quiet` are `none`); they are told from text by colour and underline on hover only.

## Surface roles

- Flat fills, capsule controls and a glow. Every `--fill-*` token is the plain colour except `--fill-page`, a grid of faint dots 8px apart.
- `--radius-control` and `--radius-pill` 8px make buttons, inputs and badges capsules; `--radius-panel` 0 keeps panels and tabs square; `--radius-page` 14px rounds the sheet, the hero and grid cells.
- `--border-width` 1px, `--border-width-strong` 2px, solid.
- No resting shadow on panels or controls. `--shadow-control-hover` is a 4px glow in the button colour and `--shadow-dialog` a 4px glow in the strong border colour: the only blurred shadows. `--shadow-text` is `none`. `--focus-ring` is 1px solid; `--transition` none.

## Never

- `palette-colours <= 18`: black, greys, white, acid green, aqua, steel blue, plus status lamps.
- `font-size <= 18px`: the wordmark and hero title are the largest text.
- `font-size >= 9px`: nothing is set below 9px.
- `font-weight <= 700`: regular and bold only.
- `font-families <= 2`: Verdana and one monospace, which also sets the headings and the interface.
- `letter-spacing <= 0px`: no text is tracked.
- `underlined-links <= 10%`: links are not underlined until hovered.
- `border-radius <= 14px`: 8px capsules, 14px on the largest containers, panels square.
- `border-width <= 2px`: hairlines, 2px for the dialog frame and tab rule.
- `box-shadow-blur <= 4px`: the hover and dialog glows are the only blur.
- `text-shadow = none`: text is flat.
- `gradient-fills <= 2%`: only the page's dot grid is a gradient.
- `transition = none`: states change instantly.
- `animation = none`: nothing moves.
