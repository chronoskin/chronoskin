## Colour roles

- Monochrome on white: `--color-page`, `--color-canvas`, `--color-bar`, `--color-input` and `--color-surface` are all `#ffffff`, so a tile has the same fill as the page and is drawn only by its hairline. `--color-surface-alt` `#f5f5f5` (panels, current navigation item), `--color-surface-strong` `#e5e5e5` (chart tracks, neutral badges, current tab, inline code), `--color-bar-alt` `#fafafa` (table head, tab strip).
- Ink: `--color-heading` and `--color-link` `#171717`; `--color-text`, `--color-heading-alt`, `--color-bar-text` and `--color-button-secondary-text` one step softer, `#262626`; `--color-text-muted`, `--color-link-quiet` and `--color-bar-alt-text` `#737373`; `--color-disabled-text` `#a3a3a3`.
- No hue leads. `--color-button` and `--color-focus` are the ink `#171717` with white text; `--color-accent` is mid grey `#737373` with white text (accent badge, highlighted chart mark, the call-to-action tile, and the faint haze where a layout draws a glow), and `--color-accent-alt` is the ink again: second chart series, checks, "new" markers and switched-on toggles.
- `--color-inverse` is `#ffffff` with `--color-inverse-text` `#171717`: on this palette the inverted block is the white one, used for the button on the grey call-to-action tile and for toggle knobs.
- Feature tiles are the faintest greys or nothing: `--color-fill-1` `#f5f5f5`, `--color-fill-2` and `--color-fill-4` `#fafafa`, `--color-fill-3` `#ffffff`. None is as dark as `--color-surface-strong`, so a chart track or a switch stays visible on them.
- Hairlines do the work: `--color-border` is black at 20% on every tile, `--color-border-strong` is solid `#171717` (secondary buttons, the dialog, the empty tile), `--color-border-muted` `#e5e5e5` for rows inside a tile, `--color-input-border` `#d4d4d4`.
- Colour appears only as status: `--color-success` `#37996b`, `--color-warning` `#b17816`, `--color-danger` `#f24822` on `--color-danger-surface` `#fafafa`. The notice is `--color-notice` `#fafafa` with `--color-notice-text` `#262626`. `--color-overlay` is `#171717` at 40%.

## Typography roles

- Mono for everything that labels, sans for everything that reads: `--font-heading`, `--font-ui` and `--font-mono` are Geist Mono, then JetBrains Mono, IBM Plex Mono and the system mono; `--font-body` is Inter, then the system sans.
- Headings and controls are capitals: `--heading-transform` and `--ui-transform` are `uppercase`. Tile titles, which do not take the transform, stay in lower case mono.
- Medium, small and flat: `--weight-heading`, `--weight-display` and `--weight-ui` 500, `--weight-body` 400, `--weight-bold` 600. `--text-display` 44px at `--line-display` 1.1 with `--display-tracking` -1px; `--text-h1` 32px, `--text-h2` 24px, `--text-h3` 18px at `--line-heading` 1.25.
- `--text-base` and `--text-small` 14px at `--line-body` 1.5; `--text-large` 16px; `--text-ui` 12px for navigation, buttons, tabs and form controls.
- Links in running text are ink with an underline (`--link-decoration` underline); quiet links are grey without one.

## Surface roles

- Dashed like a wireframe, with one small radius for everything: `--border-style` is `dashed`, so a tile is a dashed 1px outline on the page. `--radius-page` is 14px on tiles; `--radius-panel`, `--radius-control` and `--radius-pill` are all 8px, so badges, tabs and switches are rounded rectangles, never pills.
- No shadow on tiles or resting controls: `--shadow-panel` and `--shadow-control` are `none`. A hovered button gets a hard 2px offset of `--color-border-strong`; the dialog a solid 1px ring of the same colour and one short 6px drop at 10%.
- Every fill is flat: `--fill-panel` is `--color-surface`, `--fill-button` is `--color-button`. `--fill-button-hover` moves the button 15% toward its text colour.
- `--border-width` 1px is the look; `--border-width-strong` is 2px. `--focus-ring` is a 2px dashed line of `--color-focus`. `--transition` is a 0.1s linear change of colour and border colour, nothing else.

## Never

- `palette-colours <= 24`: black, white and greys; the three status colours appear only as small marks.
- `gradient-fills <= 5%`: no token is a gradient; a layout's own glow is the only one that may remain.
- `font-weight <= 600`: titles are medium; 600 is inline emphasis only.
- `font-size <= 44px`: the hero heading is the largest text.
- `font-size >= 12px`: the smallest text is 12px.
- `font-families <= 2`: the mono and the sans.
- `line-height <= 1.6`: body text sits at 1.5.
- `border-radius <= 14px`: the largest box radius is the 14px tile; nothing is a pill.
- `border-width <= 2px`: every hairline is 1px; 2px exists only as the focus ring.
- `box-shadow-blur <= 6px`: only the dialog carries a blurred shadow, 6px.
- `text-shadow = none`: no text shadows.
- `animation = none`: nothing animates.
