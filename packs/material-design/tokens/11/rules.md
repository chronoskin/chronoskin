## Colour roles

- White page: `--color-page`, `--color-canvas` and `--color-surface` are all `#ffffff`; a sheet is told from the page by its shadow alone. `--color-surface-alt` `#f5f5f5`, `--color-surface-strong` `#e0e0e0`.
- `--color-bar` coral `#ee6e73` with white text: the navigation bar and header band. `--color-bar-alt` is the lighter coral `#e57373`; `--color-inverse` blue-grey `#78909c` with white text is the footer and any inverted block.
- Teal `#26a69a` is the colour of action: `--color-button` with white text and `--color-focus`. Links are light blue `#039be5`, hover `#2196f3`, active `#1867c0`; visited links keep the link colour.
- `--color-accent` amber `#ffb731` with `#212121` text: floating button, indicators, count badges. `--color-accent-alt` is pink `#ff4081`.
- Text and headings are black at 87%, secondary text `#616161`; `--color-heading-alt` is coral, so section and panel titles pick up the bar colour.
- Block fills: coral `#ee6e73`, teal `#26a69a` and blue-grey `#78909c` carry white text, amber `#ffb731` carries `#212121`.
- Dividers are black at 14% (`--color-border`) and `#eeeeee` (`--color-border-muted`). `--color-danger` red `#f44336`, `--color-success` deep teal `#009688`, `--color-warning` amber; the notice is `#f5f5f5` with dark text.

## Typography roles

- One family, Roboto, then Segoe UI and Helvetica Neue; `--font-mono` for code only.
- Large light headings over small text: `--text-display` 56px on `--line-display` 1.1, `--text-h1` 34px, `--text-h2` 26px, all at weight 300; `--text-h3` 18px at `--weight-bold` 500. `--line-heading` 1.3.
- `--text-base` 15px on a 24px line (`--line-body` 1.6) at weight 400; `--text-large` 20px for the lead; `--text-small` and `--text-ui` are both 13px.
- `--ui-transform` uppercase at `--weight-ui` 500 for buttons, tabs and navigation; no tracking.
- Links have no underline at rest and gain one on hover (`--link-decoration-hover`); quiet links never do.

## Surface roles

- Soft paper: `--shadow-panel` is a wide, faint two-layer shadow (16px blur, 6px drop); `--shadow-control` a small soft pair (6px blur); `--shadow-control-hover` 14px blur; `--shadow-dialog` two layers up to 22px blur.
- Round corners that grow with the box: `--radius-control` 8px, `--radius-panel` 12px, `--radius-page` 16px; `--radius-pill` is a full pill.
- Flat fills; a hovered button lightens (15% white) and fields are filled with `--color-surface-alt` above their bottom line, with rounded top corners. `--border-width` 1px dividers, `--border-width-strong` 2px indicators; `--focus-ring` is a 4px halo of the focus colour at 30%; `--transition` 0.3s ease-out.

## Never

- `gradient-fills = 0%`: every fill is one flat colour.
- `text-shadow = none`: no text has a shadow.
- `border-radius <= 16px`: 8px controls, 12px cards, 16px for the largest containers; pills and discs are shapes.
- `box-shadow-blur <= 22px`: sheets rest on 6px; only the dialog reaches 22px.
- `font-weight >= 300`: light, never thin.
- `font-weight <= 500`: emphasis is medium, never bold.
- `font-size >= 13px`: captions are the smallest text.
- `font-size <= 56px`: the header band title is the largest text.
- `letter-spacing = 0px`: no tracking.
- `font-families <= 2`: one sans-serif plus the mono for code.
