# Fan wiki portal, banner and navigation boxes, 2003 to 2010

## Summary

This is the dressed-up skin of a fan or game wiki of the middle 2000s: a dark banner with the mark and a search box, a glossy bar of sections under it, navigation boxes down the left and a sheet with a thick coloured top edge at the right. The home page is a portal of browse tiles, dated news and a featured page; article pages carry a large profile box and a row of section tabs. Text is 11px Verdana, boxes are square and flat, bars are shaded with short gradients, and nothing casts a shadow or moves.

## Layout

- The page is fluid. `--size-page` is `100%`; `.ds-page` has `--space-5` (10px) above and `--space-6` (14px) at the sides. Designed at 1024px; it does not reflow for phones.
- Order from the top: the header `.ds-nav` in one frame (the banner `.ds-nav__banner`, at least `--size-banner` (64px) tall, in the inverse fill, then the bar `.ds-nav__menu` in the bar fill); then `.ds-page__body`, a grid of the `--size-side` (172px) box column `.ds-sidebar` and the sheet `.ds-page__sheet`, `--space-5` apart; then the footer bar. No text sits on the page colour.
- The sheet has a 1px frame and a top edge of twice `--border-width-strong`, both in `--color-border-strong`. It holds the views; inside a view `.ds-view__body` is one column of blocks `--space-6` apart.
- Views. The specimen is a four-screen example fan wiki; the bar has one item per view and the current one takes the accent fill. `home` is the portal: welcome box, six browse tiles, a news list beside the article of the week and "did you know", then the counters. `page` is a character page: trail, title, spoiler notice, section tabs, the floated profile box beside the text, the appearances table and see also. `category` is a category listing: letter index, subcategories and pages in three columns, paging, and an empty picture section. `community` is the community portal: two notices, the note form beside the to-do panel and the discard confirmation.
- Inner pages (every view but `home`) start with `.ds-breadcrumb` (a tinted bordered line), then `.ds-page-header` in place of the welcome box (title and one muted line at the left, buttons at the right, on a strong rule), then an optional `.ds-notice`, then `.ds-tabs` when the page has sections, then the working blocks. Floated material sits in `.ds-view__flow`.
- Two-column areas use `.ds-columns` (equal) or `.ds-columns--wide` (3 to 2), with `.ds-columns__stack` for a column of several boxes. The browse tiles and the category listing are three columns.
- The profile box `.ds-infobox` is `--size-infobox` (270px) wide and floats right. Form rows are a `--size-label` (130px) label cell and a field cell; inputs are `--size-field` (300px); the banner search field `--size-search` (180px). A dialog is `--size-dialog` (420px) wide, in a band in the flow.
- Spacing scale: `--space-1` 1px, `--space-2` 3px, `--space-3` 5px, `--space-4` 8px, `--space-5` 10px, `--space-6` 14px, `--space-7` 20px.
  - `--space-2`: vertical padding of strips, box titles, list rows, inputs and buttons.
  - `--space-3`: vertical padding of bar items, tabs and table cells.
  - `--space-4`: horizontal padding of cells and box rows, gap between tiles and between buttons.
  - `--space-5`: padding of panel bodies and notices, gap between the box column and the sheet.
  - `--space-6`: padding of the sheet and the banner, and the gap between blocks.
  - `--space-7`: list indent, gap between the listing columns.

## Typography and colour roles

- Text on fills follows the era's conventions (see the first layout of this era): no text on `--color-page`; body, heading and link colours on the canvas, the three surfaces, `--fill-panel` and the four fills; only `--color-bar-text` on `--fill-bar`, `--color-bar-alt-text` on `--fill-bar-alt`, `--color-inverse-text` on `--fill-inverse`, `--color-accent-text` on `--fill-accent`, `--color-notice-text` on `--color-notice`, and `--color-text` on `--color-danger-surface`.
- `--color-page` is charcoal with fine scan lines from `--fill-page`; the sheet `--color-canvas` is white. `--color-surface` fills the navigation boxes, the profile box and alternate news rows; `--color-surface-alt` (lavender) the breadcrumb, label cells, idle tabs, counters and alternate table rows; `--color-surface-strong` (pale blue) letter heads and plain panel titles.
- `--fill-bar` (near black, shaded) with amber `--color-bar-text` is the section bar, table heads, the profile title, the current tab, the dialog title, tile icons and the footer. `--fill-bar-alt` (deep blue) with white `--color-bar-alt-text` titles every navigation box, panel and form and the group strips of the profile box. `--fill-inverse` (dark maroon) with `--color-inverse-text` is the banner.
- `--color-heading` (maroon) is for page and section titles and figures, `--color-heading-alt` (blue) for third-level headings and news dates. Links are `--color-link` (blue), tool links `--color-link-quiet` (olive), hover `--color-link-hover` (burnt orange); a missing page is `.ds-link--new` in `--color-danger`.
- `--color-border-strong` (tan) frames the sheet, the header, the boxes and the buttons and rules the headings; `--color-border` (grey) is for ordinary box outlines and dotted row rules; `--color-border-muted` for panel feet.
- `--color-accent` (amber) with dark `--color-accent-text` marks the current bar item, the current page number and badges; `--color-accent-alt` is the left bar of an information notice.
- `--color-fill-1` to `--color-fill-4` are the lavender, cream, sage and sand of the browse tiles; the welcome box is `--color-fill-1`.
- Body and controls are `--font-body` / `--font-ui` (Verdana); headings, box titles, the wordmark and tile names are `--font-heading` (Trebuchet MS), bold. `--font-mono` is for code.
- Sizes: `--text-base` 11px; `--text-small` and `--text-ui` 10px for meta, bar items, tabs, buttons and badges; `--text-large` 13px for the lead, panel titles and tile names; `--text-h3` 12px, `--text-h2` 15px, `--text-h1` 20px, `--text-display` 22px for the wordmark and the welcome title. `--weight-ui` is 700.
- Bars, buttons and the accent are short vertical gradients from the surface set; panels, boxes and the sheet are flat. Row rules inside boxes are `dotted`.
- One-off choices outside the token set: the dividers between bar items are `--color-bar-text` at 25%; the notice's left bar is `calc(var(--border-width-strong) * 2)`; status badges are an 18 to 26% `color-mix()` of the status colour over `--color-surface`; `dotted` is used for row rules and the empty state instead of `--border-style`.

## Components

- `.ds-page`: on `<body>`. `.ds-page__frame` wraps everything, `.ds-page__body` is the grid of boxes and sheet, `.ds-page__sheet` the sheet. Views are `.ds-view` with a `.ds-view__body` column; `.ds-view__flow` contains floats. `.ds-columns` (`--wide`) and `.ds-columns__stack` make two-column areas. `.ds-sprite` hides the SVG symbol sheet, `.ds-icon` sizes an inline icon.
- `.ds-nav`: the framed header. `.ds-nav__banner` holds the brand and `.ds-nav__tools` (the `.ds-nav__search` form with its `.ds-nav__field`, and the `.ds-nav__personal` list of `.ds-nav__personal-link` items). `.ds-nav__menu` is the bar of `.ds-nav__link` items; the current one takes `is-current`.
- `.ds-brand`: the mark and name at the left of the banner. `.ds-brand__mark` is a square tile 1.9 times `--text-display`, dressed like the primary button with a `--border-width-strong` outline in `--color-border-strong`. `.ds-brand__text` stacks `.ds-brand__name` (the wordmark) and `.ds-brand__tag`, both in `--color-inverse-text`. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo: replace both, keep the tile.
- `.ds-sidebar`: the column of navigation boxes. Each `.ds-sidebar__block` has a `.ds-sidebar__title` strip in the secondary bar fill, a `.ds-sidebar__list` of dotted-ruled `.ds-sidebar__item` rows and an optional `.ds-sidebar__note`.
- `.ds-breadcrumb`: a tinted bordered line with bold `.ds-breadcrumb__link` items, `.ds-breadcrumb__sep` and `.ds-breadcrumb__current`.
- `.ds-hero`: the welcome box of the portal, in the first fill with an accent left edge: `.ds-hero__title`, `.ds-hero__lead`, buttons in `.ds-hero__action`, and a small bordered facts list `.ds-hero__aside`.
- `.ds-page-header`: the head of an inner page: `.ds-page-header__title`, `.ds-page-header__desc`, and `.ds-page-header__action` with one or two buttons, all on a strong rule.
- `.ds-section`: a section heading on a 1px strong-coloured rule, between blocks.
- `.ds-panel`: a titled box. `.ds-panel__title` is a secondary bar strip (`--plain` for a pale strip with heading text), `.ds-panel__body` the content with `.ds-panel__line` and `.ds-panel__items`, `.ds-panel__foot` a small right-aligned closing line. `.ds-thumb` with `.ds-thumb__image` is a framed picture floated left in a body.
- `.ds-grid`: three browse tiles in a row. `.ds-grid__cell` (`--2`, `--3`, `--4`) has a `.ds-grid__icon` tile, a `.ds-grid__name` and `.ds-grid__meta`, and a strong bottom edge.
- `.ds-list`: news entries: `.ds-list__item` with a bold `.ds-list__date` first line and a `.ds-list__meta` last line; even rows are tinted.
- `.ds-stat`: the row; four `.ds-stat__item` counters, each a large `.ds-stat__value` beside a small `.ds-stat__label`, with a strong left edge.
- `.ds-tabs`: sections of one page on a strong rule in the bar colour; `.ds-tabs__tab`, `is-current` takes the bar fill.
- `.ds-prose`: long text: h1, h2 (on a strong-coloured rule), h3, p, ul, strong, code, blockquote. `.ds-edit` is the bracketed edit link inside a heading.
- `.ds-link`: every link in text. States `is-visited`, `is-hover`, `is-active`, `is-focus`; `.ds-link--new` for a missing page, `.ds-link--quiet` for tool links, `.ds-link--strong` for bold.
- `.ds-infobox`: the profile box floated right: `.ds-infobox__title` (bar fill), `.ds-infobox__image`, `.ds-infobox__caption`, `.ds-infobox__group` strips and `.ds-infobox__rows` tables of label and value.
- `.ds-table`: a data table with a bar-fill head, ruled rows and tinted even rows; `.ds-table__num` right-aligns numbers.
- `.ds-index` is the letter line of a category (`.ds-index__letter`, `is-empty`); `.ds-catlist` the three-column listing with `.ds-catlist__head` letters and `.ds-catlist__items` lists.
- `.ds-badge`: a small square amber label. `.ds-badge--success`, `.ds-badge--warning` and `.ds-badge--danger` are tinted status labels with a 1px outline in the status colour.
- `.ds-notice`: a box with a thick left bar in `--color-accent-alt` and a `.ds-notice__title` line; `.ds-notice--error` uses `--color-danger` on `--color-danger-surface`. `.ds-notice__link` is an underlined link in the text colour.
- `.ds-pagination`: a centred bordered line of `.ds-pagination__link` boxes (`is-current` in the accent fill) and a `.ds-pagination__label`.
- `.ds-button`: primary action. `.ds-button--secondary` is the pale button, `.ds-button--danger` the destructive one (error text on the error surface). States `is-hover`, `is-active`, `is-focus`, `is-disabled`. Group with `.ds-buttons`.
- `.ds-form`: a box with a `.ds-form__title` strip and `.ds-form__row` lines of a tinted `.ds-form__label` cell and a `.ds-form__field` cell. Controls take `.ds-form__input` (`--wide`, `--auto`, `is-invalid`, `is-disabled`), with `.ds-form__error` and `.ds-form__hint`; `.ds-form__check` and `.ds-form__checkbox`; buttons in `.ds-form__actions`.
- `.ds-dialog`: a confirmation box in the page flow inside `.ds-dialog__backdrop`: `.ds-dialog__title` (bar fill), `.ds-dialog__body`, `.ds-dialog__text`, `.ds-dialog__actions`. It never floats over the page.
- `.ds-empty`: a dotted box with `.ds-empty__title`, `.ds-empty__text` and one button.
- `.ds-footer`: the bar under the page: `.ds-footer__text` and a `.ds-footer__row` of underlined `.ds-footer__link` items.
- `ds-accordion` (with `__summary`, `__body`): a collapsible box, a `<details>` section for contents, tips or rules.
- `ds-menu` (with `__summary`, `__list`, `__item`): a page-actions or language drop-down, a `<details>` opened by click.
- `ds-tooltip` (with `__tip`): a small hint box on a link, shown on hover or focus.
- `ds-progress` (inside `ds-meter`): a native progress bar for completeness or a drive.

## Never

- `border-radius <= 0px`: no reference has a rounded corner in CSS; every box is square.
- `box-shadow = none`: no box shadow was measured in any reference.
- `text-shadow = none`: no text shadow was measured in any reference.
- `transition = none`: hover states switch at once.
- `animation = none`: nothing moves.
- `font-size <= 22px`: the wordmark and the welcome title, 22px, are the largest text.
- `font-size >= 10px`: meta text and controls are 10px; nothing is smaller.
- `font-families <= 3`: Verdana, Trebuchet MS for headings, monospace for code.
- `border-width <= 6px`: borders are 1px, edges 3px, and only the top of the sheet and the left bar of a notice reach 6px.
- `letter-spacing <= 1px`: only the wordmark is tracked, by 1px.
- `uppercase-text <= 0%`: no text is transformed to uppercase.
- `row-gap <= 6px`: rows touch or are divided by a dotted rule.
- `block-gap <= 28px`: blocks are 14px apart.
- `content-width >= 85%`: the page fills the window.

## Extending

Derive a new component from the nearest one in the specimen: a new box in the left column from `.ds-sidebar__block`, a titled box in the sheet from `.ds-panel`, a profile row from `.ds-infobox__rows`, a message from `.ds-notice`. Give it a 1px outline, a title strip in `--fill-bar-alt` with `--color-bar-alt-text`, keep text on fills as the era's conventions say, and use tokens only: no raw colours or lengths, no radius or shadow, no font size outside the type tokens, and `--space-6` as the gap to the next block.
