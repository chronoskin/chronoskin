# Travel guide wiki, banner and guide column, 2003 to 2010

## Summary

This is the travel or city guide wiki of the later 2000s: a coloured banner with the name and a place search, a pale strip of sections with small icons, and one wide sheet. A destination guide keeps the same rhythm everywhere (understand, get in, see, do, eat, sleep, go next), with a column of contents and quick facts at the left and listings as small boxes with a round pin, an address line and a price. Text is a roomy 15px Trebuchet with teal links on white and sand, boxes are outlined with dashes, and nothing casts a shadow or moves.

## Layout

- The page is fluid. `--size-page` is `100%`; `.ds-page__frame` has `--space-4` (10px) above and `--space-6` (18px) at the sides. Designed at 1024px; it does not reflow for phones.
- Order from the top: the header `.ds-nav` (the banner `.ds-nav__top` in the bar fill with the brand at the left and the search and account links at the right; the strip `.ds-nav__menu` in the secondary bar fill); the sheet `.ds-page__sheet`, joined to the strip, on the canvas; the footer, a separate band in the inverse fill with three columns. No text sits on the page colour.
- Views. The specimen is a four-screen example travel guide; the strip has one item per view and the current one has a thick underline in `--color-accent-alt`. `home` is the welcome band with a drawn view, the figures as ticket stubs, three equal panels (`.ds-trio`) and the region cards. `destination` is one guide: trail, title, the row of page tabs, then the `--size-guide` (196px) column of contents and quick facts beside the text with its listings. `region` is a region as a table of places beside a `--size-aside` (260px) column. `add` is the listing form with two notices and the discard confirmation, beside a column of advice.
- Inner pages (every view but `home`) start with `.ds-breadcrumb` (the "is in" trail, parts joined by colons), then `.ds-page-header` in place of the welcome band (title and one muted line at the left, the action at the right, on a thick pale rule), then `.ds-tabs` when the page has several faces, then `.ds-layout`: the guide column at the left, or with `.ds-layout--aside` a help column at the right. `.ds-layout__col` stacks blocks `--space-5` apart.
- Form labels sit above their fields; `.ds-form__pair` puts two rows side by side. Inputs are `--size-field` (340px) unless wide, the banner search `--size-search` (190px), a text area `--size-editor` (90px) tall. A pin is `--size-pin` (28px) round, an icon `--size-icon` (16px). The drawn view is `--size-scene` (230px) wide. A dialog is `--size-dialog` (430px) wide.
- Spacing scale: `--space-1` 2px, `--space-2` 4px, `--space-3` 7px, `--space-4` 10px, `--space-5` 14px, `--space-6` 18px, `--space-7` 26px.
  - `--space-1`: vertical padding of buttons, tabs and fact rows.
  - `--space-2`: vertical padding of table cells, list rows and box titles.
  - `--space-3`: gap between an icon and its text, between buttons and between listings; vertical padding of strip items.
  - `--space-4`: horizontal padding of cells and boxes; gap between region cards.
  - `--space-5`: gap between blocks; padding of panels.
  - `--space-6`: padding of the sheet, the banner and the footer; gap between the columns; margin above a section heading.
  - `--space-7`: padding of the welcome band, list indent, gap between footer columns.

## Typography and colour roles

- Text on fills follows the era's conventions (see the first layout of this era): no text on `--color-page`; body, muted, heading and link colours, `--color-danger` and `--color-success` on the canvas, the three surfaces, `--fill-panel` and the four fills; only `--color-bar-text` on `--fill-bar`, `--color-bar-alt-text` on `--fill-bar-alt`, `--color-inverse-text` on `--fill-inverse`, `--color-accent-text` on `--fill-accent`, `--color-notice-text` on `--color-notice`, and `--color-text` on `--color-danger-surface`. Links on a bar take the bar's text colour and are told apart by underline.
- `--color-page` is pale sand, lighter at the top through `--fill-page`; the sheet `--color-canvas` is white. `--color-surface` fills listings, contents boxes, tables and the ticket stubs; `--color-surface-alt` idle tabs and alternate rows; `--color-surface-strong` box titles and the rule under a page title. Panels take `--fill-panel`.
- `--fill-bar` (teal) with white `--color-bar-text` is the banner, table heads, the current page tab and the dialog title. `--fill-bar-alt` (pale sun) with `--color-bar-alt-text` is the strip of sections. `--fill-inverse` (deep teal) with `--color-inverse-text` is the footer band.
- `--color-heading` (deep teal) is for titles and section headings; `--color-heading-alt` (terracotta) for third-level headings, panel titles, figures, prices and the drawn view. Links are `--color-link` (teal) without underline; trail links `--color-link-quiet` (olive); a missing guide is `.ds-link--new`.
- `--color-border` outlines boxes and `--color-border-strong` the sheet, the strip and section headings, all in `--border-style`, which this style sets to `dashed`.
- `--fill-accent` (orange) with `--color-accent-text` is the round pin, the badge and the current page number; `--color-accent` is also the left edge of a notice, and `--color-accent-alt` the underline of the current section.
- `--color-fill-1` to `--color-fill-4` are the sea, sun, olive and terracotta tints of the region cards; the welcome band is `--color-fill-1`.
- One family: `--font-body`, `--font-heading` and `--font-ui` are all Trebuchet MS; headings and controls are bold. `--font-mono` is for times.
- Sizes: `--text-base` 15px; `--text-ui` 13px for the strip, tabs, buttons, tables and contents; `--text-small` 12px for meta, facts and the footer; `--text-large` 17px for the lead, panel titles and card names; `--text-h3` 16px, `--text-h2` 21px, `--text-h1` 28px for page titles, figures and the wordmark, `--text-display` 34px for the welcome title. Line height `--line-body` 1.6.
- One-off choices outside the token set: the brand tile and the pins are round through `--radius-pill`; the current tab and the brand tile use a `solid` outline; ticket stubs are divided by `dashed` rules and fact rows by `dotted` ones whatever `--border-style` is; the notice's left edge is `--border-width-strong`; status badges are tints made with `color-mix()`; the drawn view uses opacity for its two paler layers.

## Components

- `.ds-page`: on `<body>`. `.ds-page__frame` is the padded page, `.ds-page__sheet` the sheet. Views are `.ds-view` with a `.ds-view__body` column; `.ds-layout` (`--aside`) and `.ds-layout__col` make columns, `.ds-trio` three equal ones. `.ds-sprite` hides the SVG symbol sheet and `.ds-icon` sizes an inline icon.
- `.ds-nav`: the header. `.ds-nav__top` holds the brand and `.ds-nav__side` (the `.ds-nav__search` form with its `.ds-nav__field`, and the `.ds-nav__personal` list of `.ds-nav__personal-link` items). `.ds-nav__menu` is the strip of `.ds-nav__link` items, each with an optional icon; the current one takes `is-current`.
- `.ds-brand`: mark and name at the left of the banner. `.ds-brand__mark` is a round tile 1.5 times `--text-h1`, dressed like the primary button; `.ds-brand__text` stacks `.ds-brand__name` and `.ds-brand__tag` in `--color-bar-text`. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo: replace both, keep the tile.
- `.ds-breadcrumb`: the "is in" trail: `.ds-breadcrumb__link`, `.ds-breadcrumb__current`.
- `.ds-hero`: the welcome band in the first fill: `.ds-hero__title`, `.ds-hero__lead`, buttons in `.ds-hero__action`, and `.ds-hero__scene`, a framed inline drawing in `--color-heading-alt`.
- `.ds-page-header`: the head of an inner page: `.ds-page-header__title`, `.ds-page-header__desc`, `.ds-page-header__action`.
- `.ds-tabs`: the faces of a guide as a row of lozenges; `.ds-tabs__tab`, `is-current` in the bar fill.
- `.ds-stat`: the row; four `.ds-stat__item` ticket stubs divided by dashed rules, each a `.ds-stat__value` over a `.ds-stat__label`.
- `.ds-section`: a plain heading between blocks.
- `.ds-panel`: a box in `--fill-panel` with a `.ds-panel__title`, a `.ds-panel__body` of `.ds-panel__line` paragraphs and a bold `.ds-panel__more` link line.
- `.ds-list`: short notes: `.ds-list__item` with a `.ds-list__meta` line, divided by rules.
- `.ds-grid`: region cards three to a row: `.ds-grid__cell` (`--2`, `--3`, `--4`) with a `.ds-pin`, a `.ds-grid__name` link and `.ds-grid__meta`. `.ds-pin` is the round accent tile that holds one icon, also used before each listing.
- `.ds-sidebar`: the guide column. Each `.ds-sidebar__block` has a `.ds-sidebar__title` strip and a `.ds-sidebar__list` of `.ds-sidebar__item` rows (icon and link), or a `.ds-facts` list of `.ds-facts__row` pairs (`.ds-facts__term`, `.ds-facts__value`).
- `.ds-prose`: guide text: h1, h2 (a section of the rhythm: icon, name and a `.ds-edit` link at the right, on a rule), h3 (a budget class or a way of travel), p, ul, strong, code.
- `.ds-listing`: the places of a section. Each `.ds-listing__item` is a pin, then the text (a bold `.ds-listing__name`, address and description, and a `.ds-listing__meta` line of hours), then a `.ds-listing__price`.
- `.ds-table`: a table with a bar-fill head and ruled rows; `.ds-table__num` right-aligns numbers, `.ds-table__row--alt` tints a row, `.ds-table__nowrap` keeps a cell on one line.
- `.ds-notice`: a traveller's warning with a thick left edge in `--color-accent`; `.ds-notice--error` has it in `--color-danger` on `--color-danger-surface`. `.ds-notice__title` is the bold first phrase, `.ds-notice__link` an underlined link in the text colour.
- `.ds-form`: `.ds-form__row` stacks a `.ds-form__label` over a `.ds-form__field` holding a `.ds-form__input` (`--wide`, `--auto`, `--text`; `is-invalid`, `is-disabled`) with `.ds-form__hint` or `.ds-form__error`; `.ds-form__pair` sets two rows side by side; `.ds-form__check` and `.ds-form__checkbox`; `.ds-form__actions` with `.ds-form__spacer` pushing the destructive button to the right.
- `.ds-link`: every link in text. States `is-visited`, `is-hover`, `is-active`, `is-focus`; `.ds-link--new` (in `--color-danger`) for a page nobody has written, `.ds-link--quiet` for tool and outside links, `.ds-link--strong` for bold.
- `.ds-button`: primary action. `.ds-button--secondary` is the plain button, `.ds-button--danger` the destructive one (error text on the error surface), `.ds-button--small` a smaller one. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. Group with `.ds-buttons`.
- `.ds-badge`: a small label in the accent fill for counts and "new". `.ds-badge--success`, `.ds-badge--warning` and `.ds-badge--danger` are status labels: a 14 to 22% tint of the status colour over `--color-surface`, body text and a 1px outline in the status colour.
- `.ds-pagination`: a line of small bordered `.ds-pagination__link` steps (`is-current` in the accent fill, `is-disabled` plain) with a muted `.ds-pagination__label`.
- `.ds-dialog`: a box in the page flow, centred in the `.ds-dialog__backdrop` band (`--color-overlay`, no text on it): `.ds-dialog__title` in the bar fill, `.ds-dialog__body`, `.ds-dialog__text`, `.ds-dialog__actions`. It never floats over the page.
- `.ds-empty`: a dashed box on `--color-surface` with `.ds-empty__title`, `.ds-empty__text` and one secondary button.
- `.ds-footer`: the inverted band: a wide first column with `.ds-footer__title` and `.ds-footer__text`, then columns of a `.ds-footer__title` over a `.ds-footer__row` of `.ds-footer__link` items.
- `.ds-clear` ends floats.
- `ds-accordion` (with `__summary`, `__body`): a collapsible box, a `<details>` section for contents, tips or rules.
- `ds-menu` (with `__summary`, `__list`, `__item`): a page-actions or language drop-down, a `<details>` opened by click.
- `ds-tooltip` (with `__tip`): a small hint box on a link, shown on hover or focus.
- `ds-progress` (inside `ds-meter`): a native progress bar for completeness or a drive.

## Never

- `border-radius <= 2px`: boxes have 2px corners at most; pins, tabs and badges are full pills.
- `box-shadow = none`: no box shadow was measured in any reference.
- `text-shadow = none`: no text shadow was measured in any reference.
- `transition = none`: hover states switch at once.
- `animation = none`: nothing moves.
- `font-size <= 34px`: the welcome title, 34px, is the largest text.
- `font-size >= 12px`: meta text is 12px; nothing is smaller.
- `font-families <= 2`: one sans-serif, plus monospace.
- `border-width <= 4px`: outlines are 1px; the page-title rule, notice edge and current-section underline are 4px.
- `letter-spacing <= 0.5px`: text is never spaced out; the display size is tightened by half a pixel.
- `uppercase-text <= 0%`: no text is transformed to uppercase.
- `line-height <= 1.7`: body text is set at 1.6.
- `content-width >= 75%`: the sheet stretches with the window.

## Extending

Derive a new component from the nearest one in the specimen: a box in the guide column from `.ds-sidebar__block`, a new kind of place from `.ds-listing__item` with its own icon in a `.ds-pin`, a boxed aside from `.ds-panel`, a warning from `.ds-notice`. Build it as a 1px `--color-border` outline in `--border-style` on `--color-surface`, keep text on fills as the conventions above say, and use tokens only: no raw colours or lengths, no shadow, gradient or transition of its own, no font size outside the type tokens, and `--space-5` as the gap to the next block.
