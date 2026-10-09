# Wiki encyclopaedia, side column and tabbed sheet, 2003 to 2010

## Summary

This is the collaborative reference site of the middle 2000s: a narrow column of small links under a square mark at the left, and beside it a white article sheet that stretches with the window and carries a row of tabs on its top edge. Text is small sans-serif at 13px with blue links, red links for pages nobody has written, serif titles over hairlines, and grey bordered boxes for contents, facts and plain ruled tables. Every box is square and flat: no shadow, no gradient, no rounded corner and no motion.

## Layout

- The page is fluid. `--size-page` is `100%`; `.ds-page__frame` is a grid of a fixed `--size-side` (158px) side column and a main column that takes the rest, `--space-5` (12px) apart. Designed at 1024px; it does not reflow for phones.
- Side column `.ds-page__side`, from the top: the header `.ds-nav` (the brand block, the navigation box, the search box), then the tool boxes `.ds-sidebar`. Boxes are `--space-5` apart.
- Main column `.ds-page__main`, from the top: the personal links line `.ds-personal`, right-aligned; the current view (its `.ds-tabs` row and, joined to it, the `.ds-sheet`); the footer box. The main column is a grid whose middle row takes the height left by the side column, and the sheet has a minimum height of that row less the tab row, so a short page still reaches the foot of the side column.
- No text sits on the page colour. The brand block and the personal line are painted with `--fill-bar-alt`; everything else is a box or the sheet.
- Views. The specimen is a four-screen example reference; the navigation box has one link per view and marks the current one in bold. `main` (home) is the main page: site notice, welcome box, two tinted columns (featured article and "did you know" at the left, news and "on this day" at the right), the figures line and the grid of other shelves. `article` is one article: trail, title, floated fact box, contents box, sections with bracketed edit links, a thumbnail, a notice, a data table, see also, references and the category strip. `history` is the revision table with its legend, paging line and compare buttons. `edit` is the edit form: two notices, the empty preview box, toolbar and tall text field, summary row, buttons, and the delete confirmation.
- Each view starts with its own tab row. The current tab is taller, bold, in the sheet colour and open at the bottom, so it reads as part of the sheet; idle tabs are `--color-surface` with link-coloured text. A tab to a page that does not exist takes `is-missing`.
- Inner pages (every view but `main`) open the sheet with `.ds-breadcrumb`, then `.ds-page-header` in place of the welcome box (serif title on a hairline, the action bottom-aligned at the right) and one muted line `.ds-page-header__desc`; then the working blocks. Blocks that are not running text sit in `.ds-sheet__stack`, `--space-5` apart.
- In an article the fact box `.ds-infobox` (`--size-infobox`, 250px) and thumbnails `.ds-thumb` (`--size-thumb`, 180px) float right and text flows around them; the contents box `.ds-toc` is at least `--size-toc` (280px) wide and sits in the text flow.
- Spacing scale: `--space-1` 2px, `--space-2` 4px, `--space-3` 6px, `--space-4` 8px, `--space-5` 12px, `--space-6` 16px, `--space-7` 24px.
  - `--space-1`: vertical padding of tabs, buttons, title strips; gap under a heading rule.
  - `--space-2`: gap between tabs, between inline items, vertical padding of table cells.
  - `--space-3`: padding inside tinted panels, gap between buttons.
  - `--space-4`: horizontal padding of cells, strips and boxes; gap between the two tinted columns and between grid cells.
  - `--space-5`: gap between blocks, between the columns of the page, padding of notices.
  - `--space-6`: padding of the sheet, margin above a section heading, gap between a float and the text.
  - `--space-7`: list indent.
- Text inputs are `--size-field` (320px); the edit box is `--size-editor` (260px) tall and full width. A dialog is `--size-dialog` (440px) wide, centred in a band in the flow.

## Typography and colour roles

- These are the era's conventions for which text sits on which fill; every layout and every token set of the era keeps to them.
  - `--color-page`: no text at all.
  - `--color-canvas`, `--color-surface`, `--color-surface-alt`, `--color-surface-strong`, `--fill-panel` and `--color-fill-1` to `--color-fill-4`: `--color-text`, `--color-text-muted`, `--color-heading`, `--color-heading-alt`, all link colours, `--color-danger` and `--color-success`.
  - `--fill-bar`: `--color-bar-text` only. `--fill-bar-alt`: `--color-bar-alt-text` only. `--fill-inverse`: `--color-inverse-text` only. Links on these fills take the same text colour and are told apart by underline.
  - `--fill-accent`: `--color-accent-text`. `--color-notice`: `--color-notice-text`. `--color-danger-surface`: `--color-text`, with `--color-danger` for the title.
  - `--fill-button`: `--color-button-text`; `--fill-button-secondary`: `--color-button-secondary-text`; `--fill-input`: `--color-input-text`; `--color-disabled`: `--color-disabled-text`.
  - `--color-overlay` carries no text; a dialog on it is a `--color-canvas` box.
- Here `--color-page` is pale grey (#f0f0f0) and the sheet `--color-canvas` white. `--color-surface` (#f9f9f9) fills the side boxes, contents box, fact box, table cells and idle tabs; `--color-surface-alt` alternate rows and picture frames; `--color-surface-strong` the small title strips of side boxes and the fact box.
- `--color-bar` is the grey of table heads and the dialog title, with black `--color-bar-text`. `--color-bar-alt` equals the page grey, with blue `--color-bar-alt-text`, so the brand and the personal links look as if they sat on the page. `--color-inverse` (teal) is the site notice line.
- Links are `--color-link` (blue) without underline until hovered; `--color-link-visited` purple; `--color-link-active` orange. A link to a missing page is `.ds-link--new` in `--color-danger` (red). Outside links, footer links and the breadcrumb are `--color-link-quiet`.
- `--color-border` (#aaaaaa) outlines the sheet, tabs and every box and rules the tables; `--color-border-muted` is for rules inside a box and picture frames; `--color-border-strong` only for button outlines, the mark and the dialog.
- `--color-fill-1` to `--color-fill-3` are the mint, blue and lilac tints of the main page panels; `--color-fill-4` is pale yellow. A panel's title strip is its tint mixed with 16% of `--color-success`, `--color-link` or `--color-link-visited`, and its border the tint mixed with `--color-border-strong`.
- `--color-accent` (amber) is the left bar of an information notice, the badge and the round letters of the grid; `--color-success` and `--color-danger` colour the plus and minus byte counts.
- `--font-body` and `--font-ui` are the plain sans-serif stack; `--font-heading` is the serif stack, weight 400, used for the page title, section headings, the brand name, the welcome title and the figures. `--font-mono` is for code and the edit box.
- Sizes: `--text-base` 13px running text; `--text-ui` 12px tabs, buttons, inputs, tables, contents box; `--text-small` 11px side boxes, personal line, fact box, captions, footer; `--text-large` 15px panel titles; `--text-h3` 15px bold sans; `--text-h2` 19px; `--text-h1` and `--text-display` 24px. Line height `--line-body` 1.5.
- `--ui-transform` is `lowercase`: tabs, side box titles, the personal line and figure labels are set in lowercase by the type set, never by typing them so.
- One-off choices outside the token set: the panel title tints above; the notice's left bar is `calc(var(--border-width-strong) * 4)`; status badges are a 16 to 22% `color-mix()` of the status colour over `--color-surface`; the empty state has a `dashed` outline.

## Components

- `.ds-page`: on `<body>`. `.ds-page__frame` is the two-column grid, `.ds-page__side` the side column, `.ds-page__main` the main column. `.ds-view` sections hold a tab row and a `.ds-sheet`; `.ds-sheet__stack` spaces blocks inside it; `.ds-clear` ends floats. `.ds-sprite` hides the SVG symbol sheet and `.ds-icon` sizes an inline icon.
- `.ds-nav`: the head of the side column. It holds the brand, then `.ds-nav__block` boxes with a `.ds-nav__title` strip. `.ds-nav__menu` lists `.ds-nav__item` entries; a `.ds-nav__link` leads to a view and takes `is-current` (bold, text colour). `.ds-nav__search` is the search box with `.ds-nav__go` for its two small buttons.
- `.ds-brand`: the mark and name stacked and centred at the top of the side column, on `--fill-bar-alt`. `.ds-brand__mark` is a square tile 2.5 times `--text-display`, dressed like the primary button (`--fill-button`, `--color-button-text`, a `--color-border-strong` outline, `--radius-control`, `--shadow-control`). `.ds-brand__name` is the serif name, `.ds-brand__tag` an italic line under it. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo: replace both, keep the tile.
- `.ds-personal`: the line of account links at the top right, each a `.ds-personal__link`.
- `.ds-hero`: the welcome box of the main page: centred `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__count`, one button in `.ds-hero__action`, and `.ds-hero__aside`, a three-column list of shelf links. Never a banner.
- `.ds-sitenotice`: one inverted line above the welcome box; its link is `.ds-sitenotice__link`.
- `.ds-page-header`: the head of an inner page. `.ds-page-header__title` over a hairline, `.ds-page-header__action` at the right, and `.ds-page-header__desc` as the next line.
- `.ds-breadcrumb`: the small trail above the title, with `.ds-breadcrumb__link`, `.ds-breadcrumb__sep` and `.ds-breadcrumb__current`.
- `.ds-tabs`: the views of one page, above the sheet. `.ds-tabs__tab` with `is-current` and `is-missing`; `.ds-tabs__gap` separates the watch tab.
- `.ds-prose`: the article body: h1, h2 (serif on a hairline), h3 (bold sans), p, ul, ol, strong, code, sup. `.ds-edit` is the bracketed edit link inside a heading; `.ds-hatnote` the italic pointer line.
- `.ds-link`: every link in text. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--new` for a missing page, `.ds-link--quiet` for outside and tool links, `.ds-link--strong` for bold.
- `.ds-toc`: the contents box, with `.ds-toc__title`, `.ds-toc__toggle`, nested `.ds-toc__list` and `.ds-toc__num`.
- `.ds-infobox`: the fact box floated right: `.ds-infobox__title`, `.ds-infobox__image`, `.ds-infobox__caption`, `.ds-infobox__group` and a `.ds-infobox__rows` table of label and value.
- `.ds-thumb`: a framed picture with `.ds-thumb__image` and `.ds-thumb__caption`; `.ds-thumb--left` is the small left-floated one.
- `.ds-refs` is the small numbered reference list and `.ds-catlinks` the bordered category strip at the foot of an article.
- `.ds-columns`: the two tinted columns of the main page.
- `.ds-panel`: a tinted bordered box holding one or more pairs of `.ds-panel__title` and `.ds-panel__body`. `.ds-panel--1`, `--2`, `--3` choose the tint. `.ds-panel__line` is a paragraph, `.ds-panel__items` a bullet list, `.ds-panel__more` the bold right-aligned closing line.
- `.ds-list`: dated entries: `.ds-list__item` with a bold `.ds-list__date` and text with a `.ds-list__meta` line.
- `.ds-grid`: four cells in a row, each `.ds-grid__cell` (`--2`, `--3`, `--4` for the other fills) with a `.ds-grid__icon` letter, a `.ds-grid__name` link and `.ds-grid__meta`.
- `.ds-stat`: the row; one bordered line of `.ds-stat__item` figures divided by faint rules, each a serif `.ds-stat__value` over a small `.ds-stat__label`.
- `.ds-table`: a plain ruled table: grey head, every cell bordered. `.ds-table__num` right-aligns numbers, `.ds-table__tools` keeps a cell on one line, `.ds-table__row--alt` is the alternate row, `.ds-table__plus` and `.ds-table__minus` colour byte changes, `.ds-table__summary` is the italic edit summary.
- `.ds-badge`: a small square amber label (the "m" of a minor edit, "new"). `.ds-badge--success`, `.ds-badge--warning` and `.ds-badge--danger` are tinted status labels with a 1px outline in the status colour.
- `.ds-notice`: a box with a thick left bar in `--color-accent`; `.ds-notice--error` has the bar in `--color-danger` on `--color-danger-surface`. `.ds-notice__title` is the bold first phrase, `.ds-notice__link` an underlined link in the text colour.
- `.ds-pagination`: a line of small bordered steps: `.ds-pagination__label`, `.ds-pagination__link` with `is-current` and `is-disabled`.
- `.ds-button`: primary action. `.ds-button--secondary` is the plain grey button, `.ds-button--danger` the destructive one (error text on the error surface), `.ds-button--small` for the side column. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. Group with `.ds-buttons`.
- `.ds-toolbar`: the strip of markup keys over the edit box: `.ds-toolbar__key` (`--bold`, `--italic`) and `.ds-toolbar__hint`.
- `.ds-form`: the edit form. `.ds-form__editor` wraps toolbar and text field; `.ds-form__row` is a line of `.ds-form__label` and `.ds-form__input` (`--wide`, `--auto`, `--text` for the tall field; `is-invalid`, `is-disabled`), with `.ds-form__error` and `.ds-form__hint`; `.ds-form__check` and `.ds-form__checkbox`; `.ds-form__legal`; `.ds-form__actions` with `.ds-form__spacer` pushing the destructive button to the right.
- `.ds-dialog`: a confirmation box in the page flow inside `.ds-dialog__backdrop`: `.ds-dialog__title`, `.ds-dialog__body`, `.ds-dialog__text`, `.ds-dialog__actions`. It never floats over the page.
- `.ds-empty`: a dashed box with `.ds-empty__title`, `.ds-empty__text` and one button.
- `.ds-sidebar`: the tool boxes of the side column: `.ds-sidebar__block`, `.ds-sidebar__title`, `.ds-sidebar__list`, `.ds-sidebar__item`.
- `.ds-section` is a serif heading on a hairline outside prose; `.ds-legend` a small muted explanatory line.
- `.ds-footer`: the small-print box under the sheet: `.ds-footer__badge`, `.ds-footer__body`, `.ds-footer__text`, `.ds-footer__row` and `.ds-footer__link`.
- `ds-accordion` (with `__summary`, `__body`): a collapsible box, a `<details>` section for contents, tips or rules.
- `ds-menu` (with `__summary`, `__list`, `__item`): a page-actions or language drop-down, a `<details>` opened by click.
- `ds-tooltip` (with `__tip`): a small hint box on a link, shown on hover or focus.
- `ds-progress` (inside `ds-meter`): a native progress bar for completeness or a drive.

## Never

- `border-radius <= 0px`: no reference has a rounded corner; every box is square.
- `box-shadow = none`: no box shadow was measured in any reference.
- `text-shadow = none`: no text shadow was measured in any reference.
- `transition = none`: hover states switch at once.
- `animation = none`: nothing moves.
- `gradient-fills <= 0%`: every fill is flat.
- `font-size <= 24px`: the page title, 24px, is the largest text.
- `font-size >= 11px`: side boxes and captions are 11px; nothing is smaller.
- `font-families <= 3`: one sans-serif, one serif for titles, monospace for code.
- `border-width <= 8px`: borders are 1px; only the left bar of a notice is thicker.
- `letter-spacing <= 0px`: no reference tracks its text.
- `line-height <= 1.6`: body text is set at 1.5.
- `underlined-links <= 15%`: links are not underlined until hovered.
- `row-gap <= 6px`: table rows touch and list rows are a few pixels apart.
- `block-gap <= 24px`: blocks are 12px apart; 24px is the widest gap.
- `content-width >= 75%`: the sheet stretches with the window.

## Extending

Derive a new component from the nearest one in the specimen: a new box in the side column from `.ds-sidebar__block`, a box inside an article from `.ds-toc` or `.ds-infobox`, a tinted main-page section from `.ds-panel`, a message from `.ds-notice`. Build it as a 1px `--color-border` outline on `--color-surface`, keep text on fills as the conventions above say, and use tokens only: no raw colours or lengths, no radius, shadow, gradient or transition, no font size outside the type tokens, and `--space-5` as the gap to the next block.
