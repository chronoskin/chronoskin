## Colour roles

- Dark theme: `--color-page` `#212121`, sheets `--color-surface` `#3c3c3c`, `--color-surface-alt` `#424242`, `--color-surface-strong` and `--color-border` `#505050`. Text is white, secondary text and section titles white at 70%.
- `--color-bar` is white with `#212121` text: the navigation bar and header band are the light sheet above the dark page. `--color-inverse` and `--color-bar-alt` `#eeeeee` with `#212121` text: footer and toolbar.
- Amber `#ffc400` is primary and accent: `--color-button`, `--color-link`, `--color-accent`, all with `#212121` text. Hover is yellow `--color-link-hover` `#ffeb3b`; `--color-focus` is the same yellow.
- Block fills: yellow `#ffeb3b`, amber `#ffc400`, green `#00e676`, white; text on all four is `#212121`.
- Status colours are bright with dark text: `--color-danger` pink `#ff4081`, `--color-success` `#00e676`, `--color-warning` `#ffeb3b`, `--color-accent-alt` lime `#b2ff59`. `--color-button-secondary` is `#eeeeee`.

## Typography roles

- One family, Roboto; body, headings and display are light: `--weight-body`, `--weight-heading`, `--weight-display` 300.
- Large headings: `--text-display` 55px on `--line-display` 1.05, `--text-h1` 40px, `--text-h2` 26px, `--text-h3` 20px at `--weight-bold` 500; `--line-heading` 1.2.
- `--text-base` 16px on a 24px line; `--text-small` and `--text-ui` are both 14px, so nothing is smaller than 14px.
- `--ui-transform` uppercase at `--weight-ui` 500 for buttons, tabs and navigation; no underlines, no tracking.

## Surface roles

- Everything is square: `--radius-control`, `--radius-panel`, `--radius-page` and `--radius-pill` are 0. Badges, chips and pagination cells are rectangles; only discs are round.
- Hard edges instead of soft shadows: `--shadow-panel` is a faint 1px ring with a solid 2px edge under the sheet and no blur; `--shadow-control` is the 2px edge alone, darker; `--shadow-control-hover` adds an 8px-blur shadow; `--shadow-dialog` has a 4px edge and an 8px blur.
- Heavy indicators: `--border-width-strong` is 4px, so the line under the current tab and navigation item and the focus and error line of a field are thick bars; `--border-width` 1px dividers.
- Flat fills; fields are a bottom line on a transparent fill. `--focus-ring` is a 2px dashed line; `--transition` is a snap, 0.1s linear.

## Never

- `gradient-fills = 0%`: every fill is one flat colour.
- `text-shadow = none`: no text has a shadow.
- `border-radius = 0px`: every box, button and tag is square; discs are the only round shape.
- `box-shadow-blur <= 8px`: shadows stay close to the sheet.
- `font-weight >= 300`: light, never thin.
- `font-weight <= 500`: emphasis is medium, never bold.
- `font-size >= 14px`: no caption size below 14px.
- `font-size <= 55px`: the header band title is the largest text.
- `underlined-links = 0%`: links are marked by colour and weight.
- `letter-spacing = 0px`: no tracking.
- `font-families <= 2`: one sans-serif plus the mono for code.
