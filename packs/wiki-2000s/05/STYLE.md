# Team wiki, header bar and page tree, 2003 to 2010

## Summary

This is the team or project wiki that offices ran from about 2005: a glossy blue header bar with the name, the main sections as tabs and a search box, a pale strip of space links under it, and below that a pane with the tree of pages beside the pane with the page itself. A page has a trail, a title with its author line, a row of operation tabs, tinted note and tip boxes, ruled tables, a strip of labels, its attachments and a thread of comments with people drawn as initials. Text is 12px Tahoma, corners are barely rounded, bars and buttons carry a short gradient, and only dialogs cast a shadow.

## Layout

- The page is fluid. `--size-page` is `100%`; the header `.ds-nav` runs edge to edge. Designed at 1024px; it does not reflow for phones.
- Order from the top: the header (the bar `.ds-nav__bar` in the bar fill, the strip `.ds-nav__sub` in the secondary bar fill); `.ds-page__body`, a grid of the `--size-tree` (214px) tree pane `.ds-sidebar` and the content pane `.ds-page__main`, with `--space-4` (8px) around and between them; the footer band in the inverse fill. Both panes are `--color-canvas` with a 1px outline. No text sits on the page colour, which shows only as the gaps.
- Views. The specimen is a four-screen example team wiki; the bar has one tab per view and the current one takes the canvas colour. `dashboard` (home) is the welcome panel, four counters, the spaces grid, then the recently updated list beside two panels. `page` is one wiki page: trail, title, operation tabs, the text beside a tinted facts panel and a note, the steps table, labels, attachments and comments. `pages` lists the pages of a space: tabs, table, paging and an empty section. `edit` is the editor: two notices, the form beside notation help, and the page properties dialog.
- Inner pages (every view but `dashboard`) start with `.ds-breadcrumb`, then `.ds-page-header` in place of the welcome panel (title and a muted author line at the left, buttons at the right, no rule), then `.ds-tabs` on a thick bar-coloured rule, then the working blocks in `.ds-view__body`, `--space-5` (10px) apart.
- Two-column areas inside the content pane use `.ds-layout` (fluid column and a `--size-help` (220px) column) or `.ds-layout--even` (3 to 2), with `.ds-layout__col` for a column of several blocks.
- Form rows are a right-aligned `--size-label` (120px) label cell and the field cell; inputs are `--size-field` (320px), the header search `--size-search` (150px), the editor `--size-editor` (150px) tall. An avatar is `--size-avatar` (26px) square. A dialog is `--size-dialog` (460px) wide.
- Spacing scale: `--space-1` 2px, `--space-2` 3px, `--space-3` 5px, `--space-4` 8px, `--space-5` 10px, `--space-6` 15px, `--space-7` 22px.
  - `--space-1`: gap between tabs, padding of tree rows.
  - `--space-2`: vertical padding of tabs, table cells, strips and buttons.
  - `--space-3`: vertical padding of bar tabs and list rows, gap between buttons.
  - `--space-4`: horizontal padding of cells and boxes; the gap around and between the panes; gap between cards.
  - `--space-5`: gap between blocks, padding of panel bodies.
  - `--space-6`: side padding of the content pane, gap between its columns, indent of a tree level.
  - `--space-7`: indent of a reply in the comments.

## Typography and colour roles

- Text on fills follows the era's conventions (see the first layout of this era): no text on `--color-page`; body, muted, heading and link colours, `--color-danger` and `--color-success` on the canvas, the three surfaces, `--fill-panel` and the four fills; only `--color-bar-text` on `--fill-bar`, `--color-bar-alt-text` on `--fill-bar-alt`, `--color-inverse-text` on `--fill-inverse`, `--color-accent-text` on `--fill-accent`, `--color-notice-text` on `--color-notice`, and `--color-text` on `--color-danger-surface`. Links on a bar take the bar's text colour and are told apart by underline.
- `--color-page` is a cool grey-blue; panes are `--color-canvas` (white). `--color-surface` fills comment boxes, the labels strip and alternate rows; `--color-surface-alt` idle tabs, label cells of forms and comment heads; `--color-surface-strong` tree-pane titles, avatars and the editor toolbar.
- `--fill-bar` (blue, glossy) with white `--color-bar-text` is the header bar, table heads, the current operation tab and the dialog title; `--color-bar` also rules the tabs and tops the counters. `--fill-bar-alt` (pale blue) with `--color-bar-alt-text` is the strip under the bar and every panel title. `--fill-inverse` is the footer band.
- `--color-heading` (navy) is for titles and figures; `--color-heading-alt` (blue) for third-level headings and tree-pane titles. Links are `--color-link`; breadcrumb and label links `--color-link-quiet`; a missing page is `.ds-link--new`.
- `--color-border` outlines panes, boxes and table cells; `--color-border-strong` the welcome panel, avatars and buttons; `--color-border-muted` rules rows.
- `--color-fill-1` to `--color-fill-4` are the tinted panels of this kind of wiki: information blue, note yellow, tip green and warning pink. They fill the space cards and the avatars; the welcome panel is `--color-fill-1`, a note panel `--color-fill-2`, a tip panel `--color-fill-3`. The current tree row is marked with `--color-fill-2`.
- `--color-accent` (green) with `--color-accent-text` is the count badge and the current page number.
- One family: `--font-body`, `--font-heading` and `--font-ui` are all Tahoma; headings are bold. `--font-mono` is for code, space keys and the editor.
- Sizes: `--text-base` 12px; `--text-ui` and `--text-small` 11px for tabs, buttons, tables, the tree and meta; `--text-large` 14px for the welcome lead and card names; `--text-h3` 13px, `--text-h2` 16px, `--text-h1` and `--text-display` 21px for page titles, figures and the wordmark.
- One-off choices outside the token set: `dotted` rules in the tree and in panel rows; the notice outline is `--color-warning` mixed into `--color-border`; counters and tabs use a `--border-width-strong` edge in `--color-bar`; status badges are tints made with `color-mix()`.

## Components

- `.ds-page`: on `<body>`. `.ds-page__body` is the grid of the two panes, `.ds-page__main` the content pane. Views are `.ds-view` with a `.ds-view__body` column; `.ds-layout` (`--even`) and `.ds-layout__col` make columns. `.ds-sprite` hides the SVG symbol sheet and `.ds-icon` sizes an inline icon.
- `.ds-nav`: the header. `.ds-nav__bar` holds the brand, the `.ds-nav__menu` of `.ds-nav__link` tabs (`is-current` in the canvas colour) and the `.ds-nav__search` form with its `.ds-nav__field`. `.ds-nav__sub` is the strip with two `.ds-nav__tools` lists of `.ds-nav__sublink` items.
- `.ds-brand`: mark and name at the left of the bar. `.ds-brand__mark` is a tile 1.25 times `--text-display`, dressed like the primary button; `.ds-brand__name` is the bold name in `--color-bar-text`. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo: replace both, keep the tile.
- `.ds-sidebar`: the tree pane. Each `.ds-sidebar__block` has a `.ds-sidebar__title` strip and either a `.ds-tree` or a `.ds-sidebar__list` of `.ds-sidebar__item` rows. `.ds-tree` is the nested page tree: `.ds-tree__item`, a `.ds-tree__row` (`is-current`) with a `.ds-tree__toggle` box (`--leaf` for a page without children) and a link.
- `.ds-avatar`: a person as two initials on a small tile; `--2`, `--3`, `--4` take the other fills, `--small` is for running text.
- `.ds-breadcrumb`: the location line: `.ds-breadcrumb__link`, `.ds-breadcrumb__current`.
- `.ds-hero`: the welcome panel of the dashboard in the first fill: `.ds-hero__title`, `.ds-hero__lead`, and buttons in `.ds-hero__action` at the right.
- `.ds-page-header`: the head of an inner page: `.ds-page-header__title`, `.ds-page-header__desc` (the author line), `.ds-page-header__action`.
- `.ds-tabs`: the operations of a page; `.ds-tabs__tab`, `is-current` in the bar fill, with an optional `.ds-tabs__count`.
- `.ds-stat`: the row; four `.ds-stat__item` counters with a bar-coloured top edge, each a `.ds-stat__value` over a `.ds-stat__label`.
- `.ds-section`: a heading on a rule, with an optional small `.ds-section__more` link at the right.
- `.ds-grid`: four space cards in a row: `.ds-grid__cell` (`--2`, `--3`, `--4`) with a monospace `.ds-grid__key`, a `.ds-grid__name` link and `.ds-grid__meta`.
- `.ds-list`: recent changes: `.ds-list__item` is an avatar, a `.ds-list__title` link with a `.ds-list__meta` line, and a `.ds-list__time`; even rows are tinted.
- `.ds-panel`: a macro box: `.ds-panel__title` in the secondary bar fill and a `.ds-panel__body` of `.ds-panel__line` paragraphs or a `.ds-panel__items` list of `.ds-panel__item` rows (label left, value right). `.ds-panel--note` and `.ds-panel--tip` tint the box.
- `.ds-prose`: page text: h1, h2 (on a hairline), h3, p, ul, ol, strong, code. `.ds-edit` is a small edit link in a heading.
- `.ds-labels`: the strip of labels under a page, each a `.ds-labels__tag`.
- `.ds-table`: a fully ruled table with a bar-fill head; `.ds-table__num` right-aligns numbers, `.ds-table__row--alt` tints a row, `.ds-table__nowrap` keeps a cell on one line, `.ds-table__file` puts an icon before a file or page name.
- `.ds-comments`: the thread under a page: `.ds-comments__item` (`--reply` is indented) is an avatar beside a `.ds-comments__box` with a `.ds-comments__head` line and a `.ds-comments__text`.
- `.ds-notice`: a macro box with a `.ds-notice__mark` at the left and, beside it, text that opens with a bold `.ds-notice__title`; `.ds-notice--error` is outlined in `--color-danger` on `--color-danger-surface`. `.ds-notice__link` is an underlined link in the text colour.
- `.ds-form`: a bordered table of rows: `.ds-form__row` is a tinted right-aligned `.ds-form__label` and a `.ds-form__field` holding a `.ds-form__input` (`--wide`, `--auto`, `--text`; `is-invalid`, `is-disabled`) with `.ds-form__hint` or `.ds-form__error`, `.ds-form__check` and `.ds-form__checkbox`; `.ds-form__actions` is the foot, with `.ds-form__spacer` pushing the destructive button to the right. `.ds-toolbar` with `.ds-toolbar__key` is the strip of markup keys over the editor.
- `.ds-link`: every link in text. States `is-visited`, `is-hover`, `is-active`, `is-focus`; `.ds-link--new` (in `--color-danger`) for a page nobody has written, `.ds-link--quiet` for tool and outside links, `.ds-link--strong` for bold.
- `.ds-button`: primary action. `.ds-button--secondary` is the plain button, `.ds-button--danger` the destructive one (error text on the error surface), `.ds-button--small` a smaller one. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. Group with `.ds-buttons`.
- `.ds-badge`: a small label in the accent fill for counts and "new". `.ds-badge--success`, `.ds-badge--warning` and `.ds-badge--danger` are status labels: a 14 to 22% tint of the status colour over `--color-surface`, body text and a 1px outline in the status colour.
- `.ds-pagination`: a line of small bordered `.ds-pagination__link` steps (`is-current` in the accent fill, `is-disabled` plain) with a muted `.ds-pagination__label`.
- `.ds-dialog`: a box in the page flow, centred in the `.ds-dialog__backdrop` band (`--color-overlay`, no text on it): `.ds-dialog__title` in the bar fill, `.ds-dialog__body`, `.ds-dialog__text`, `.ds-dialog__actions`. It never floats over the page.
- `.ds-empty`: a dashed box on `--color-surface` with `.ds-empty__title`, `.ds-empty__text` and one secondary button. The dialog here also shows `.ds-dialog__props`, a list of `.ds-dialog__term` and `.ds-dialog__value` pairs.
- `.ds-footer`: the band under the panes: `.ds-footer__text` and a `.ds-footer__row` of `.ds-footer__link` items.
- `.ds-clear` ends floats.
- `ds-accordion` (with `__summary`, `__body`): a collapsible box, a `<details>` section for contents, tips or rules.
- `ds-menu` (with `__summary`, `__list`, `__item`): a page-actions or language drop-down, a `<details>` opened by click.
- `ds-tooltip` (with `__tip`): a small hint box on a link, shown on hover or focus.
- `ds-progress` (inside `ds-meter`): a native progress bar for completeness or a drive.

## Never

- `border-radius <= 9px`: corners are 3px; only badges and label tags are rounder.
- `box-shadow-blur <= 10px`: only the dialog casts a shadow, 10px soft.
- `text-shadow = none`: no text shadow was measured in any reference.
- `transition = none`: hover states switch at once.
- `animation = none`: nothing moves.
- `font-size <= 21px`: page titles, 21px, are the largest text.
- `font-size >= 11px`: tables and meta are 11px; nothing is smaller.
- `font-families <= 2`: one sans-serif, plus monospace.
- `border-width <= 3px`: outlines are 1px; the tab rule and counter edge are 3px.
- `letter-spacing <= 0px`: no reference tracks its text.
- `uppercase-text <= 0%`: no text is transformed to uppercase.
- `line-height <= 1.5`: body text is set at 1.4.
- `content-width >= 75%`: the panes fill the window.

## Extending

Derive a new component from the nearest one in the specimen: a box in the tree pane from `.ds-sidebar__block`, a macro from `.ds-panel` or `.ds-notice`, a row of people from `.ds-avatar`, a threaded item from `.ds-comments__item`. Build it as a 1px `--color-border` outline on `--color-surface` or one of the four fills, keep text on fills as the conventions above say, and use tokens only: no raw colours or lengths, no shadow or transition of its own, no font size outside the type tokens, and `--space-5` as the gap to the next block.
