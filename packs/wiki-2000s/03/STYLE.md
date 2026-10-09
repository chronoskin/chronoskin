# How-to wiki, top navigation and numbered steps, 2003 to 2010

## Summary

This is the friendlier wiki of the later 2000s, the how-to manual or project handbook that ran on a different engine from the encyclopaedias: a coloured bar across the top with the name, the sections as tabs and a search box, and under it one wide column with a rail of link boxes at the right instead of a side column at the left. Pages are white sheets with softly rounded corners, section titles sit in rounded tinted strips, steps are numbered with round badges, and tips and warnings are boxes with a coloured left bar. Text is 14px Helvetica or Arial with plain blue links; there are no shadows and no motion.

## Layout

- The page is fluid. `--size-page` is `100%`; the header runs edge to edge and `.ds-page__frame` has `--space-6` (24px) at the sides. Designed at 1024px; it does not reflow for phones.
- Order from the top: the header `.ds-nav` (the bar `.ds-nav__bar` with brand, section tabs and search; the strip `.ds-nav__sub` with the tagline at the left and account links at the right); `.ds-page__frame` with the current view; the footer band, edge to edge. No text sits on the page colour: everything is on a sheet, a box or a band.
- Navigation is at the top, never in a left column. A view that needs side links uses `.ds-layout`: a fluid main column and a `--size-rail` (216px) rail `.ds-sidebar` at the right, `--space-5` (16px) apart.
- Views. The specimen is a five-screen example how-to manual; the bar has one tab per view and the current one takes the sheet colour. `home` is the home page: welcome sheet with the three figures, spotlight box, category cells, newest guides, and the rail. `guide` is one guide: trail, page tabs joined to the sheet, title, needs table, numbered steps, a tip and a warning, related guides, and the rail. `changes` is the recent changes list by day with paging. `search` is the results page: search line, the "no such title" box and the matching guides. `write` is the new-guide form with its error notice, buttons and the discard confirmation, beside a rail of help.
- Inner pages (every view but `home`) start with `.ds-breadcrumb` (a white bordered line), then a `.ds-sheet` that opens with `.ds-page-header` in place of the welcome sheet (title and one muted line at the left, buttons at the right, no rule). A page with several faces puts `.ds-tabs` directly above the sheet and marks the sheet `.ds-sheet--tabbed`. Blocks in a sheet are `--space-5` apart; blocks in a view `--space-4`.
- Category cells are three to a row. Form rows are a right-aligned `--size-label` (150px) label and the field; inputs are `--size-field` (340px); the bar search field `--size-search` (170px). The spotlight picture is `--size-thumb` (150px). A dialog is `--size-dialog` (430px) wide, in a tinted band in the flow.
- Spacing scale: `--space-1` 2px, `--space-2` 4px, `--space-3` 8px, `--space-4` 12px, `--space-5` 16px, `--space-6` 24px, `--space-7` 32px.
  - `--space-1`: gap between bar tabs, padding of rail rows and page numbers.
  - `--space-2`: vertical padding of tabs, strips, buttons, inputs and table cells.
  - `--space-3`: gap between buttons, cells and figures; padding of cells and notices.
  - `--space-4`: horizontal padding of cells, strips and buttons; gap between blocks of a view; padding of rail boxes.
  - `--space-5`: gap between main column and rail, between blocks in a sheet, top padding of a sheet.
  - `--space-6`: side padding of the page, the bar and a sheet; padding of the welcome sheet.
  - `--space-7`: reserved for a wider break between page regions.

## Typography and colour roles

- Text on fills follows the era's conventions (see the first layout of this era): no text on `--color-page`; body, heading and link colours on the canvas, the three surfaces, `--fill-panel` and the four fills; only `--color-bar-text` on `--fill-bar`, `--color-bar-alt-text` on `--fill-bar-alt`, `--color-inverse-text` on `--fill-inverse`, `--color-accent-text` on `--fill-accent`, `--color-notice-text` on `--color-notice`, and `--color-text` on `--color-danger-surface`.
- `--color-page` is a pale green tint, slightly darker at the top through `--fill-page`; sheets, rail boxes, lists and the breadcrumb are `--color-canvas` (white). `--color-surface-alt` (light green) fills section strips, figures, idle tabs, day rows and the search line; `--color-surface-strong` the table head; `--color-surface` the empty state and table foot.
- `--fill-bar` (leaf green) with white `--color-bar-text` is the top bar, the step numbers and the current page number; `--color-bar` also rules rail titles and the table head. `--fill-bar-alt` (pale blue) with dark blue `--color-bar-alt-text` is the strip under the bar. `--fill-inverse` (dark green) with `--color-inverse-text` is the footer band.
- `--color-heading` is near black; `--color-heading-alt` (dark green) is for section strips, panel titles, figures and third-level headings. Links are `--color-link` (blue), tool and author links `--color-link-quiet` (green); a missing page is `.ds-link--new` in `--color-danger`.
- `--color-border` (#dddddd) outlines sheets, boxes and cells; `--color-border-muted` rules rows inside them; `--color-border-strong` outlines secondary buttons, the current bar tab and the dialog.
- `--color-fill-1` to `--color-fill-4` are the pale blue, green, yellow and grey of the category cells; the spotlight box is `--color-fill-1` and search matches are marked with `--color-fill-3`. `--color-notice` (pale blue) is the tip box, with a `--color-success` left bar.
- `--color-accent` (orange) with white `--color-accent-text` is the count badge only.
- One family for everything: `--font-body`, `--font-heading` and `--font-ui` are the same Helvetica and Arial stack; headings are bold. `--font-mono` is for code and times.
- Sizes: `--text-base` 14px; `--text-small` 12px for meta, rails, the strip and the footer; `--text-ui` 13px for tabs, buttons and inputs; `--text-large` 16px for the lead, result titles and rail titles; `--text-h3` 14px, `--text-h2` 18px, `--text-h1` 26px for page titles, figures and the wordmark; `--text-display` 32px for the welcome title only. Line height `--line-body` 1.5.
- Corners are rounded by the surface set: `--radius-control` 4px on buttons, inputs and notices, `--radius-panel` 6px on boxes, `--radius-page` 8px on sheets, tabs and cells, and `--radius-pill` on badges, section strips and step numbers, which makes the numbers round.
- One-off choices outside the token set: the primary button's outline is `--color-button` mixed with 30% of `--color-shadow`; the notice's left bar is `calc(var(--border-width-strong) * 3)`; status badges are a 14 to 20% `color-mix()` of the status colour over `--color-surface`; an invalid field has a `--border-width-strong` outline.

## Components

- `.ds-page`: on `<body>`. `.ds-page__frame` holds the views. Each `.ds-view` has a `.ds-view__body` column; `.ds-layout` with `.ds-layout__main` adds the rail; `.ds-sheet` (`--tabbed` under a tab row) is the white sheet. `.ds-sprite` hides the SVG symbol sheet.
- `.ds-nav`: the header. `.ds-nav__bar` holds the brand, `.ds-nav__menu` (tabs `.ds-nav__link`, the current one `is-current` in the sheet colour) and `.ds-nav__search` with its `.ds-nav__field`. `.ds-nav__sub` is the strip with `.ds-nav__tagline` and the `.ds-nav__tools` list of `.ds-nav__sublink` items.
- `.ds-brand`: the mark and name at the left of the bar. `.ds-brand__mark` is a tile 1.35 times `--text-h1`, dressed like the primary button; `.ds-brand__name` is the bold name in `--color-bar-text`. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo: replace both, keep the tile.
- `.ds-breadcrumb`: a white bordered line under the header: `.ds-breadcrumb__link`, `.ds-breadcrumb__sep`, bold `.ds-breadcrumb__current`.
- `.ds-hero`: the welcome sheet of the home page: `.ds-hero__title` at the display size, `.ds-hero__lead`, two large buttons in `.ds-hero__action`, and the figures at the right.
- `.ds-stat`: the row; three `.ds-stat__item` tiles, each a large green `.ds-stat__value` over a small `.ds-stat__label`.
- `.ds-page-header`: the head of an inner page: `.ds-page-header__title`, `.ds-page-header__desc`, and `.ds-page-header__action` with one or two buttons.
- `.ds-section`: a section title in a rounded tinted strip.
- `.ds-panel`: the spotlight box in the first fill: a white rounded `.ds-panel__title` strip and a `.ds-panel__body` of a framed `.ds-panel__thumb` beside a `.ds-panel__name` link and `.ds-panel__line` paragraphs.
- `.ds-grid`: category cells three to a row: `.ds-grid__cell` (`--2`, `--3`, `--4`) with a bold `.ds-grid__name` and `.ds-grid__meta`.
- `.ds-list`: a bordered list of rows: `.ds-list__item` is time (`.ds-list__time`), a `.ds-list__title` with a `.ds-list__meta` line, and a right `.ds-list__aside`. `.ds-list__day` is a tinted date row; `.ds-list__plus` and `.ds-list__minus` colour byte changes.
- `.ds-sidebar`: the rail. Each `.ds-sidebar__block` is a white box with a `.ds-sidebar__title` on a bar-coloured rule and a bulleted `.ds-sidebar__list` of `.ds-sidebar__item` rows.
- `.ds-tabs`: the faces of one page (guide, discuss, edit, history), rounded at the top and joined to the sheet below; `.ds-tabs__tab`, `is-current` in the sheet colour.
- `.ds-prose`: long text: h1, h2, h3, p, ul, strong, code. `.ds-edit` is a small edit link inside a heading.
- `.ds-link`: every link in text. States `is-visited`, `is-hover`, `is-active`, `is-focus`; `.ds-link--new` for a missing page, `.ds-link--quiet` for tool and author links, `.ds-link--strong` for bold.
- `.ds-steps`: the numbered steps of a guide. Each `.ds-steps__item` has a `.ds-steps__num` badge in the bar fill, a bold `.ds-steps__title` and a `.ds-steps__text`. Write the number in the markup.
- `.ds-table`: a light table: tinted head on a bar-coloured rule, rows divided by faint rules; `.ds-table__num` right-aligns numbers, `.ds-table__foot` is a bold total row.
- `.ds-badge`: an orange pill for counts and marks. `.ds-badge--success`, `.ds-badge--warning` and `.ds-badge--danger` are tinted status pills with a 1px outline in the status colour.
- `.ds-notice`: a tip box with a thick left bar in `--color-success`; `.ds-notice--error` is the warning or error box with the bar in `--color-danger`. `.ds-notice__title` is the bold first word, `.ds-notice__link` an underlined link in the text colour.
- `.ds-pagination`: a line of rounded `.ds-pagination__link` boxes (`is-current` in the bar fill) with a right-aligned `.ds-pagination__label`.
- `.ds-button`: primary action. `.ds-button--secondary` is the grey button, `.ds-button--danger` the destructive one (error text on the error surface), `.ds-button--large` for the welcome sheet. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. Group with `.ds-buttons`.
- `.ds-form`: rows of a right-aligned `.ds-form__label` and a field. Controls take `.ds-form__input` (`--wide`, `--auto`, `is-invalid`, `is-disabled`), with `.ds-form__error` and `.ds-form__hint`; `.ds-form__check` and `.ds-form__checkbox`; buttons in `.ds-form__actions` under a faint rule. `.ds-searchbar` with `.ds-searchbar__field` is the one-line search form.
- `.ds-results`: search results: `.ds-results__title`, `.ds-results__text` with `.ds-results__match` marks, `.ds-results__meta`.
- `.ds-dialog`: a confirmation box in the page flow inside the tinted `.ds-dialog__backdrop`: `.ds-dialog__title`, `.ds-dialog__text`, `.ds-dialog__actions`. It never floats over the page.
- `.ds-empty`: a pale bordered box with `.ds-empty__title`, `.ds-empty__text` and one button.
- `.ds-footer`: the inverted band at the foot: `.ds-footer__text` and a `.ds-footer__row` of bold underlined `.ds-footer__link` items.
- `ds-accordion` (with `__summary`, `__body`): a collapsible box, a `<details>` section for contents, tips or rules.
- `ds-menu` (with `__summary`, `__list`, `__item`): a page-actions or language drop-down, a `<details>` opened by click.
- `ds-tooltip` (with `__tip`): a small hint box on a link, shown on hover or focus.
- `ds-progress` (inside `ds-meter`): a native progress bar for completeness or a drive.

## Never

- `border-radius <= 8px`: corners are softly rounded, 4 to 8px; only pills and round numbers go further.
- `box-shadow = none`: no box shadow was measured in any reference.
- `text-shadow = none`: no text shadow was measured in any reference.
- `transition = none`: hover states switch at once.
- `animation = none`: nothing moves.
- `font-size <= 32px`: the welcome title, 32px, is the largest text.
- `font-size >= 12px`: meta text is 12px; nothing is smaller.
- `font-families <= 2`: one sans-serif stack, plus monospace.
- `border-width <= 6px`: borders are 1px or 2px; only the left bar of a notice is 6px.
- `letter-spacing <= 1px`: at most 1px of tracking was measured; this pack uses none.
- `uppercase-text <= 0%`: no text is transformed to uppercase.
- `line-height <= 1.6`: body text is set at 1.5.
- `gradient-fills <= 10%`: only buttons carry a faint gradient.
- `block-gap <= 32px`: blocks are 12 to 16px apart.
- `content-width >= 85%`: the page fills the window.

## Extending

Derive a new component from the nearest one in the specimen: a new rail box from `.ds-sidebar__block`, a tinted feature box from `.ds-panel`, a row layout from `.ds-list`, a message from `.ds-notice`, a numbered procedure from `.ds-steps`. Put it on a `.ds-sheet` or give it a 1px `--color-border` outline on `--color-canvas` with `--radius-panel`, keep text on fills as the era's conventions say, and use tokens only: no raw colours or lengths, no shadow or transition, no font size outside the type tokens, and `--space-5` as the gap to the next block.
