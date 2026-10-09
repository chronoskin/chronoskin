## Colour roles

- Ink is dark navy, not black: `--color-text`, `--color-heading`, `--color-border`, `--color-border-strong` and `--color-input-border` are #090b37. Panel and sidebar titles (`--color-heading-alt`), links in running text (`--color-link`, `--color-link-visited`) and the focus ring (`--color-focus`) are indigo #2c2680; quiet, hovered and active links go back to the navy ink.
- The page is a cool off-white (`--color-page`, `--color-canvas` #f8f8ff) and `--color-surface` is white for cards, panels, tables and inputs. The navigation bar and table head (`--color-bar`) are white with navy text.
- Fills are pale and never saturated, and all carry navy text: `--color-surface-alt` pink #ffdfe7 (empty state, current row, inline code, hovered tab), `--color-surface-strong` mint #d1ffee (panel and dialog title strips, hovered link), `--color-fill-1` to `--color-fill-4` pink #ffdfe7, mint #d1ffee, lavender #e0defa and peach #fadabe.
- `--color-accent` is peach #fadabe with navy text: the current navigation item, tab and page and the default badge. `--color-accent-alt` and `--color-button` are pale lime #d8f999: the marked words of the hero title and the primary button. The secondary button is white.
- The one loud colour is the shadow: `--color-shadow` is pink #ff8ac5, so every hard shadow is pink under a thin navy outline.
- The inverted block and the announcement strip are navy with white text (`--color-inverse`, `--color-bar-alt` #090b37).
- Status, all with navy text: `--color-success` pale green #d0f5c4, `--color-warning` peach-brown #e9aa74, `--color-danger` rose #d76b8e (error text, the destructive button, the danger badge), `--color-danger-surface` pale rose #f8d8d8 behind an error notice, `--color-notice` lavender #e0defa.
- `--color-text-muted` #5c5c66 for dates and hints; `--color-border-muted` #d0d8f0 for faint dividers; disabled is #e0e8e8 with #5c5c66 text; `--color-overlay` is the navy at 70%.

## Typography roles

- One humanist sans for everything: `--font-body`, `--font-heading` and `--font-ui` are Source Sans 3, then Hind, then Lucida Grande and Trebuchet MS. `--font-mono` is the system monospace, for inline code only.
- Text is large and light: `--text-base` 18px at `--line-body` 1.5, `--text-small` 14px, `--text-ui` 16px, `--text-large` 22px for leads and large buttons.
- `--text-display` 56px at `--line-display` 1.1 for the hero title; `--text-h1` 48px, `--text-h2` 36px, `--text-h3` 22px at `--line-heading` 1.2.
- Nothing is heavier than 600: `--weight-body` and `--weight-ui` 400, `--weight-bold`, `--weight-heading` and `--weight-display` 600. No uppercase headings, no tracking; buttons and navigation are in sentence case.
- Links in running text are underlined; navigation, titles and breadcrumbs are not.

## Surface roles

- Outlines are thin: `--border-width` 1px on every box and rule, `--border-width-strong` 2px under the navigation.
- Corners are soft and grow with the box: `--radius-control` 14px, `--radius-panel` 18px, `--radius-page` 24px, `--radius-pill` 9999px so badges are full pills.
- Shadows stay hard and unblurred, in pink, and fall straight down, so every box stands on a lip instead of a corner block: `--shadow-panel` `0 6px 0 0`, `--shadow-control` `0 4px 0 0`, `--shadow-control-hover` `0 2px 0 0`, `--shadow-control-pressed` zero offset, `--shadow-dialog` `0 10px 0 0`.
- All fills are flat (`--fill-page` is the plain page colour, no grid).
- `--focus-ring` is `2px solid` in `--color-focus`; `--transition` animates only `transform` and `box-shadow` over 0.12s.

## Never

- `box-shadow-blur <= 0px`: the outlines are thin but every shadow is still a hard offset block.
- `border-width >= 1px`: every box keeps its navy outline.
- `border-width <= 2px`: outlines are 1px, 2px under the navigation; nothing heavier.
- `border-radius <= 24px`: 14px on controls, 18px on boxes, 24px on the largest; badges are pills.
- `text-shadow = none`: no text shadow.
- `gradient-fills <= 0%`: every fill is flat, including the page.
- `font-size <= 56px`: the hero title is the largest text.
- `font-size >= 12px`: badge and footer labels are the smallest text.
- `font-weight >= 400`: no light weights.
- `font-weight <= 600`: emphasis stops at 600.
- `font-families <= 2`: one humanist sans and a monospace for code.
- `line-height <= 1.5`: body text is set at 1.5.
- `letter-spacing = 0px`: nothing is tracked.
- `uppercase-text <= 5%`: uppercase is limited to badges and footer labels.
- `palette-colours <= 18`: navy, indigo, white, the off-white page, pink, mint, lavender, peach, lime, the status tones, the pink shadow and two greys.
- `animation = none`: nothing moves on its own.
