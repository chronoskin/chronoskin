# Software project wiki, button tiles and download box, 2003 to 2010

## Summary

This is the wiki of a free software project in the middle 2000s, the kind that served as home page, manual and workshop at once: one white sheet from edge to edge with the name at the top, a row of small bevelled button tiles for navigation and a strip that traces the pages just visited, and no column of links at either side. The start page puts the introduction beside a box of current versions, then hubs for each kind of reader, dated news and small boxes of links; manual pages float a contents box at the right and end, like every page, in a second row of tiles for the page itself. Text is 12px Lucida with bold Palatino titles, plum and lilac boxes with an amber edge, square corners, hard one-pixel bevels and offsets instead of soft shadows, and no motion.

## Layout

- The page is fluid. `--size-page` is `100%`; `.ds-page` has `--space-6` (18px) around the single sheet `.ds-page__sheet`, which is `--color-canvas` with a 1px `--color-border-strong` outline and `--space-6` of side padding. Designed at 1024px; it does not reflow for phones.
- Everything is inside the sheet, from the top: the header `.ds-nav` (the masthead `.ds-nav__top` with the brand at the left and the account line and search at the right; the menu `.ds-nav__menu`, a row of tiles closed by a `--border-width-strong` rule); the trace strip `.ds-trace` in the secondary bar fill; the current view; the page tools `.ds-pagetools`, a second row of tiles over a strong rule; the footer box. No text sits on the page colour. There are no bars as chrome and no side column: `--fill-bar` is the head of tables and dialogs, `--fill-inverse` the title of the download box.
- Views. The specimen is a four-screen example project wiki; the menu has one tile per view and the current one takes the primary button fill. `home` is the start page: the introduction beside the download box (`.ds-hero`, a fluid column and `--size-download`, 276px), four figure cells, four hub cells, then news beside a `--size-side` (232px) column of link boxes. `manual` is one manual page: trail, title, page tabs, text flowing round a floated `--size-toc` (236px) column with the contents box and a requirements panel, a command block, a notice, and the chapter paging. `releases` is the release table with its legend, paging and an empty nightly-build section. `report` is the ticket form with two notices, its buttons and the discard confirmation.
- Inner pages (every view but `home`) start with `.ds-breadcrumb`, then `.ds-page-header` in place of the introduction (title and one muted line at the left, the action at the right, on a 2px rule), then `.ds-tabs` when the page has several faces, then the working blocks. All blocks of a view are in `.ds-view__body`, `--space-5` (12px) apart.
- Form rows are a right-aligned `--size-label` (136px) label cell and a field cell; inputs are `--size-field` (320px), the search field `--size-search` (168px), a text area `--size-editor` (112px) tall. A dialog is `--size-dialog` (440px) wide, centred in a band in the flow.
- Spacing scale: `--space-1` 2px, `--space-2` 4px, `--space-3` 6px, `--space-4` 9px, `--space-5` 12px, `--space-6` 18px, `--space-7` 26px.
  - `--space-1`: vertical padding of tiles, buttons, tabs, strips and inputs.
  - `--space-2`: vertical padding of table cells, gap between tabs and page numbers.
  - `--space-3`: gap between tiles and between buttons, vertical padding of news rows and notices.
  - `--space-4`: horizontal padding of cells, boxes and strips; gap between figure and hub cells.
  - `--space-5`: gap between blocks; horizontal padding of tiles and notices; padding of the masthead.
  - `--space-6`: padding of the sheet and the page, gap between columns, margin above a section heading.
  - `--space-7`: list indent, padding under the sheet.

## Typography and colour roles

- The era's conventions for which text sits on which fill hold here as in every layout of the era.
  - `--color-page`: no text at all.
  - `--color-canvas`, `--color-surface`, `--color-surface-alt`, `--color-surface-strong`, `--fill-panel` and `--color-fill-1` to `--color-fill-4`: `--color-text`, `--color-text-muted`, `--color-heading`, `--color-heading-alt`, all link colours, `--color-danger` and `--color-success`.
  - `--fill-bar`: `--color-bar-text` only. `--fill-bar-alt`: `--color-bar-alt-text` only. `--fill-inverse`: `--color-inverse-text` only. Links on these fills take the same text colour and are underlined.
  - `--fill-accent`: `--color-accent-text`. `--color-notice`: `--color-notice-text`. `--color-danger-surface`: `--color-text`, with `--color-danger` for the title.
  - `--fill-button`: `--color-button-text`; `--fill-button-secondary`: `--color-button-secondary-text`; `--fill-input`: `--color-input-text`; `--color-disabled`: `--color-disabled-text`.
  - `--color-overlay` carries no text; a dialog on it is a `--color-canvas` box.
- Here `--color-page` is lilac grey (#d9d2e4) under a plum and amber stripe, and the sheet white. `--color-surface` fills the boxes, table cells and the footer; `--color-surface-alt` the figure cells, alternate rows, idle tabs and inline code; `--color-surface-strong` the title strips of link boxes and panels.
- `--color-bar` is deep plum with white text and an amber foot: table heads and the dialog title. `--color-bar-alt` (pale lilac) is the trace strip. `--color-inverse` (aubergine) is the title of the download box.
- Navigation tiles and page tools are dressed as secondary buttons (`--fill-button-secondary`, `--color-button-secondary-text`, a `--color-border-strong` outline, `--shadow-control`); the current tile and the mark are dressed as the primary button.
- Links are `--color-link` (blue) without underline until hovered, when they turn `--color-link-hover` (burnt orange); `--color-link-visited` is purple. A link to a page not yet written is `.ds-link--new` in `--color-danger`. Account, footer, outside links and the breadcrumb are `--color-link-quiet`.
- `--color-border` outlines boxes and rules tables; `--color-border-strong` outlines the sheet, tiles, buttons, the download box and the rules under the menu and over the page tools; `--color-border-muted` divides rows inside a box.
- `--color-fill-1` to `--color-fill-4` (lilac, blue, cream, rose) are the four hub cells. `--color-accent` (amber) is the left edge of a figure cell, the top edge of the current tab, the frame of a notice, the hub icons, the badge and the current page number.
- `--font-body` and `--font-ui` are the Lucida stack; `--font-heading` is Palatino, bold, for the name, page titles, section headings, box titles and figures; `--font-mono` is for commands, version numbers and news dates.
- Sizes: `--text-base` 12px text; `--text-ui` 11px tiles, buttons, inputs, tabs, tables; `--text-small` 11px meta, contents box, footer; `--text-large` 14px lead and the large button; `--text-h3` 14px; `--text-h2` 17px; `--text-h1` 23px; `--text-display` 26px name and welcome title. Line height `--line-body` 1.55.
- `--ui-transform` is `lowercase`: tiles, buttons, tabs, box titles and figure labels are lowercased by the type set, never by typing them so.
- One-off choices outside the token set: status badges are a 16 to 22% `color-mix()` of the status colour over `--color-surface`; the left edge of a figure cell is `calc(var(--border-width-strong) * 2)`; news rows are divided by a `dotted` rule; the command block and the empty state have a `dashed` outline.

## Components

- `.ds-page`: on `<body>`. `.ds-page__sheet` is the one sheet. `.ds-view` sections hold a `.ds-view__body`, the column of blocks; `.ds-flow` contains floats and `.ds-clear` ends them.
- `.ds-nav`: the header. `.ds-nav__top` is the masthead; `.ds-nav__tools` stacks `.ds-nav__account` (the account line) over `.ds-nav__search` (a `.ds-nav__field` input and two small buttons). `.ds-nav__menu` is the row of tiles: `.ds-nav__link` leads to a view and takes `is-current`; `.ds-nav__tool` is the lighter tile for site tools, right of `.ds-nav__spacer`.
- `.ds-brand`: the mark and name in a row at the left of the masthead. `.ds-brand__mark` is a square tile 1.7 times `--text-display`, dressed like the primary button (`--fill-button`, `--color-button-text`, a `--color-border-strong` outline, `--radius-control`, `--shadow-control`). `.ds-brand__text` holds `.ds-brand__name` in the heading colour and the muted `.ds-brand__tag`. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo: replace both, keep the tile.
- `.ds-trace`: the strip of pages just visited, with `.ds-trace__label`, `.ds-trace__link` and `.ds-trace__note` at the right.
- `.ds-hero`: the introduction of the start page: `.ds-hero__text` (a `--fill-panel` box with `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__line`, the buttons in `.ds-hero__action` and the bold link line `.ds-hero__jump` with `.ds-hero__sep`) beside the download box. Never a banner.
- `.ds-download`: the box of current versions: `.ds-download__title` in the inverse fill, a `.ds-table--plain`, then `.ds-download__body` with `.ds-download__note` and one button.
- `.ds-page-header`: the head of an inner page: `.ds-page-header__title` and `.ds-page-header__desc` at the left, `.ds-page-header__action` at the right, on a 2px rule.
- `.ds-breadcrumb`: the small trail above the title, with `.ds-breadcrumb__link`, `.ds-breadcrumb__sep` and `.ds-breadcrumb__current`.
- `.ds-tabs`: the faces of one page, on a rule. `.ds-tabs__tab` with `is-current` (sheet colour, amber top edge, open at the bottom) and `is-missing`.
- `.ds-prose`: manual text: h1, h2 (on a hairline), h3 (in `--color-heading-alt`), p, ul, ol, strong, code. `.ds-code` is a block of commands; `.ds-edit` the bracketed edit link in a heading.
- `.ds-link`: every link in text. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--new` for a missing page, `.ds-link--quiet` for outside and tool links, `.ds-link--strong` for bold.
- `.ds-aside`: the floated right column of a manual page; it holds `.ds-toc` (the contents box: `.ds-toc__title`, nested `.ds-toc__list`) and panels.
- `.ds-panel`: a box with a `.ds-panel__title` strip and a `.ds-panel__body`; `.ds-panel__line` is a paragraph and `.ds-panel__rows` a small table of label and value.
- `.ds-stat`: the row; four cells `.ds-stat__item`, each a Palatino `.ds-stat__value` and a small `.ds-stat__label` on one line, with an amber left edge.
- `.ds-grid`: four hub cells in a row, each `.ds-grid__cell` (`--2`, `--3`, `--4` for the other fills) with `.ds-grid__head` (a `.ds-grid__icon` tile, `.ds-grid__name`, `.ds-grid__meta`) and a `.ds-grid__links` list.
- `.ds-columns`: news beside the link boxes; `.ds-columns__main` stacks the news blocks. `.ds-section` is a heading on a hairline outside prose; `.ds-legend` a small muted line.
- `.ds-list`: dated news: `.ds-list__item` with a monospace `.ds-list__date`, a bold `.ds-list__title` link and a `.ds-list__meta` line.
- `.ds-sidebar`: the column of link boxes: `.ds-sidebar__block`, `.ds-sidebar__title`, `.ds-sidebar__list`, `.ds-sidebar__item`.
- `.ds-table`: a ruled table with a bar-coloured head. `.ds-table__num` right-aligns numbers, `.ds-table__nowrap` keeps a cell on one line, `.ds-table__row--alt` is the alternate row, `.ds-table__version` a monospace version number. `.ds-table--plain` drops the outer frame inside another box.
- `.ds-badge`: a small amber label. `.ds-badge--success`, `.ds-badge--warning` and `.ds-badge--danger` are tinted status labels with a 1px outline in the status colour.
- `.ds-notice`: a message framed in `--color-accent` with a thicker top edge; `.ds-notice--error` is framed in `--color-danger` on `--color-danger-surface`. `.ds-notice__title` is the bold first phrase, `.ds-notice__link` an underlined link in the text colour; `.ds-notices` stacks several.
- `.ds-pagination`: small bordered steps: `.ds-pagination__label`, `.ds-pagination__link` with `is-current` and `is-disabled`.
- `.ds-button`: primary action. `.ds-button--secondary` is the pale button, `.ds-button--danger` the destructive one (error text on the error surface), `.ds-button--large` the download button, `.ds-button--small` for the search box. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. Group with `.ds-buttons`.
- `.ds-form`: a boxed form. `.ds-form__row` is a `.ds-form__label` cell and a `.ds-form__field` cell holding `.ds-form__input` (`--wide`, `--auto`, `--text` for the text area; `is-invalid`, `is-disabled`), `.ds-form__hint`, `.ds-form__error`, and `.ds-form__check` with `.ds-form__checkbox`. `.ds-form__actions` is the button row, with `.ds-form__spacer` pushing the destructive button to the right.
- `.ds-dialog`: a confirmation box in the page flow inside `.ds-dialog__backdrop`: `.ds-dialog__title`, `.ds-dialog__body`, `.ds-dialog__text`, `.ds-dialog__actions`. It never floats over the page.
- `.ds-empty`: a dashed box with `.ds-empty__title`, `.ds-empty__text` and one button.
- `.ds-pagetools`: the row of `.ds-nav__tool` tiles at the foot of every page, with `.ds-pagetools__note` at the right.
- `.ds-footer`: the small-print box: `.ds-footer__text`, `.ds-footer__row` of `.ds-footer__link`, and `.ds-footer__badges`, a stack of two-part `.ds-footer__badge` labels (`.ds-footer__key` on the accent fill, `.ds-footer__value`).
- `ds-accordion` (with `__summary`, `__body`): a collapsible box, a `<details>` section for contents, tips or rules.
- `ds-menu` (with `__summary`, `__list`, `__item`): a page-actions or language drop-down, a `<details>` opened by click.
- `ds-tooltip` (with `__tip`): a small hint box on a link, shown on hover or focus.
- `ds-progress` (inside `ds-meter`): a native progress bar for completeness or a drive.

## Never

- `border-radius <= 2px`: every box is square; only badges have 2px corners.
- `box-shadow-blur <= 0px`: bevels and offsets are hard, one to four pixels, never blurred.
- `text-shadow = none`: no text shadow was measured in any reference.
- `transition = none`: hover states switch at once.
- `animation = none`: nothing moves.
- `font-size <= 26px`: the name and the welcome title, 26px, are the largest text.
- `font-size >= 11px`: tiles, tables and meta are 11px; nothing is smaller.
- `font-families <= 3`: one sans-serif, one serif for titles, monospace for commands.
- `border-width <= 4px`: outlines are 1px, rules 2px; only the left edge of a figure cell is 4px.
- `letter-spacing <= 0px`: no reference tracks its text.
- `line-height <= 1.6`: body text is set at 1.55.
- `block-gap <= 26px`: blocks are 12px apart.
- `content-width >= 80%`: the sheet fills the window.

## Extending

Derive a new component from the nearest one in the specimen: a navigation or tool link from `.ds-nav__tool`, a box of links from `.ds-sidebar__block`, a box beside manual text from `.ds-toc` or `.ds-panel` inside `.ds-aside`, a message from `.ds-notice`. Build it as a 1px `--color-border` outline on `--color-surface` or one of the four fills, keep text on fills as the conventions above say, and use tokens only: no raw colours or lengths, no radius, blur or transition of its own, no font size outside the type tokens, and `--space-5` as the gap to the next block.
