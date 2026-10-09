## Colour roles

- Dark theme: `--color-page` `#212121`, sheets `--color-surface` `#333333`, `--color-surface-alt` `#484848`, `--color-surface-strong` `#555555`. Text is `#eeeeee`, headings white, secondary text and section titles `#bdbdbd`. Dividers are white at 17% (`--color-border`) and 8% (`--color-border-muted`).
- `--color-bar` teal `#009688` with white text: the navigation bar and header band. `--color-bar-alt` blue-grey `#78909c` with white text is the secondary toolbar. `--color-inverse` `#eeeeee` with `#212121` text: the inverted block is the light one on this dark page.
- Cyan `#00bcd4` is the colour of action: `--color-button` with dark `#212121` text, links and the focus line. Hover is light blue `#5cbbf6` for both; active links are white. `--color-button-secondary` is `#eeeeee` with dark text.
- `--color-accent` pink `#ff80ab` with `#212121` text: floating button, indicators, count badges. `--color-accent-alt` is amber `#ffb731`.
- Block fills: teal `#009688` and cyan `#00bcd4` carry white text, pink `#ff80ab` and `#eeeeee` carry `#212121`.
- Status colours are light with dark text: `--color-danger` soft red `#e57373` on a `#333333` surface, `--color-success` light green `#8bc34a`, `--color-warning` amber `#ffb731`. The notice is `#555555` with white text.

## Typography roles

- No webfont: `--font-body`, `--font-heading` and `--font-ui` are Helvetica Neue, Helvetica, Arial; `--font-mono` is Menlo, Monaco, Consolas.
- `--text-base` 15px on a 25px line (`--line-body` 1.67); `--text-small` 13px; `--text-ui` 14px.
- Modest regular headings: `--text-display` 35px, `--text-h1` 24px, `--text-h2` 20px, `--text-h3` 15px at `--weight-bold` 700; `--text-large` 20px for the lead. `--line-heading` 1.4, `--line-display` 1.2.
- Body, headings and display are 400; `--weight-ui` 500; `--weight-bold` 700, because the system faces have no medium weight.
- `--ui-transform` uppercase for buttons, tabs and navigation; no underlines, no tracking.

## Surface roles

- Tonal sheets: `--shadow-panel` is `none` and sheets have no outline; `--fill-page` is `--color-page` darkened with 14% black, so the page is always a step darker than the sheets on it, under a light palette and under a dark one, and a sheet is told from the page by its tone alone.
- Buttons keep a tight lift: `--shadow-control` is a 2px-blur pair, `--shadow-control-hover` a 10px-blur pair; the dialog alone drops a 10px shadow.
- `--radius-control` and `--radius-panel` are 2px; `--radius-page` is 0, so the largest containers and grid cards are square; `--radius-pill` is a full pill for badges, chips and pagination.
- Flat fills; fields are filled with `--color-input` above their bottom line. `--border-width` 1px dividers, `--border-width-strong` 2px indicators, `--focus-ring` a 1px line; `--transition` is slow and soft, 0.4s.

## Never

- `gradient-fills = 0%`: every fill is one flat colour.
- `text-shadow = none`: no text has a shadow.
- `border-radius <= 2px`: every box and control has a 2px corner; pills and discs are the only round shapes.
- `box-shadow-blur <= 10px`: sheets have no shadow at all, buttons and the dialog reach at most 10px.
- `font-weight >= 400`: no light text.
- `font-weight <= 700`: bold is the heaviest weight.
- `font-size >= 13px`: captions are the smallest text.
- `font-size <= 35px`: the header band title is the largest text.
- `underlined-links = 0%`: links are marked by colour and weight.
- `letter-spacing = 0px`: no tracking.
- `font-families <= 2`: one system sans-serif plus the mono for code.
