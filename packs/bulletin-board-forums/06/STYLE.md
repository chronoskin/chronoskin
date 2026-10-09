# Bulletin board forum, clan site front page, 2002 to 2007

## Summary

This is the clan or fan site that small communities built on top of their board software between about 2003 and 2007: a dark banner with the name, a menu box at the head of a tinted rail that runs down the left side of every page, and beside it news posts with comment counts, a roster table, a poll and a shoutbox. Inner pages keep the rail and show one working column: the roster, the match results, a match report with its comments, the application form. Pale grey-teal title strips with dark text sit on near-white cells in 11px Tahoma, links are underlined, boxes are square with a light inner line, and nothing moves.

## Layout

- The sheet `.ds-page__frame` is `--size-page` (88%) wide and centred, with a 1px `--color-border-strong` outline; the page colour shows as a band on both sides and `.ds-page` has `--space-6` (16px) above and below. Designed at 1024px. Never let the sheet touch the window edge.
- Order from the top, all inside the sheet: the banner `.ds-banner` in the inverse fill, at least `--size-banner` (78px) tall, brand and tagline on the left and three short status lines on the right; the strip `.ds-strip` in the secondary bar with the date on the left and the member links on the right; then `.ds-page__body`, a grid of the rail `.ds-page__rail` (`--size-rail`, 178px, tinted `--color-surface-alt` for its full height, a 1px rule on its right) and `.ds-page__main`; then the footer band and the small print `.ds-page__legal`. Nothing with text sits on the page colour.
- The rail is written once and stays on every page. From the top, `--space-4` apart: the menu box `.ds-nav` (title strip, one row per section, the current one in the accent fill), then `.ds-sidebar` with the shoutbox and a box of links. There is no horizontal navigation bar.
- The front page (home view) in `.ds-page__main`, blocks `--space-5` apart: the welcome box `.ds-hero` (text and buttons on the left, the next match in a `--size-next`, 230px, tinted cell on the right), an optional `.ds-notice`, the `.ds-stat` row of four boxed figures, the news `.ds-list` (one box per post, `--space-4` apart), then `.ds-split`: the roster `.ds-table` in the wide column and the poll `.ds-panel` in the `--size-side` (32%) column.
- Inner pages keep the banner, strip, rail and footer and drop the hero. `.ds-page__main` is one column, from the top: `.ds-breadcrumb` (a plain small line starting "You are here:"), `.ds-page-header` in place of the hero (an open title with a small accent square before it, one muted line of description, the main action bottom-right, closed by a 1px rule), an error `.ds-notice` when there is one, then the working blocks `--space-5` apart. The roster page has the squad `.ds-grid`, `.ds-tabs`, the full `.ds-table` and a `.ds-toolbar` with links on the left and `.ds-pagination` on the right. The match page has the next-match `.ds-panel` with its buttons, the confirmation `.ds-dialog` in the flow under it, tabs, the results table and an `.ds-empty` block. A report and the application form use `.ds-split`: text or form in the wide column, one or two `.ds-panel` boxes in the narrow one.
- Views. The specimen is a five-screen example; the menu box has one row per view and marks the current one. `news` (home) is the front page: welcome box with the next match, notice, the four figures, three news posts, the first team roster and the poll. `roster` is the member list: squad cards, squad tabs, the roster table and page links. `matches` is the match page: next match and line-up buttons, the withdraw confirmation, result tabs and table, and an empty scrim list. `report` is one news post in full with its comments beside the match facts and related links. `join` is the trial application as it came back with an error, beside two boxes of notes.
- Table columns: text columns share the width; a numeric column is `--size-num` (58px, right-aligned), a date column `--size-date` (104px). Rows are at least `--size-row` (26px); title strips and head rows `--size-bar` (22px). Alternate rows take `is-alt`.
- Spacing scale: `--space-1` 1px, `--space-2` 2px, `--space-3` 4px, `--space-4` 8px, `--space-5` 12px, `--space-6` 16px.
  - `--space-1`: the gap between cells inside a bordered block and between menu rows.
  - `--space-2`: vertical padding of strips, cells, buttons and inputs; gap between stacked small lines.
  - `--space-3`: vertical padding of menu rows and form rows; gap between inline items and between comments.
  - `--space-4`: horizontal padding of cells and strips, padding of panel bodies and the rail, the gap between rail boxes, news posts, grid cards and buttons.
  - `--space-5`: padding of the main column and the hero, and the gap between neighbouring blocks and columns. This is the only block gap in the main column.
  - `--space-6`: page padding above and below the sheet, banner padding, list indent, padding of the empty state.
- Form rows are open: a `--size-label` (28%) label and the field on the box surface, divided by faint rules; text inputs are `--size-field` (240px) unless `--wide`; the action row is right-aligned. A dialog is `--size-dialog` (62%) of the main column, centred in the flow. Meters are `--size-meter` (8px) high.

## Typography and colour roles

- `--font-body` and `--font-ui` are Tahoma; `--font-heading` is Verdana, for the wordmark, page and box titles, news titles and figures. `--font-mono` is for inline code only.
- Sizes: `--text-small` 10px for meta, strips, the rail lists and captions; `--text-base` and `--text-ui` 11px for running text, table cells, buttons, inputs and the menu; `--text-h2` 11px bold for box titles; `--text-large` 12px for news titles and card names; `--text-h1` 14px for the page title, the hero title and the figures; `--text-display` 18px for the wordmark only. Nothing is larger than 18px.
- `--line-body` is 1.45, `--line-heading` 1.3. Headings are bold, never uppercase, no tracking. `--weight-ui` is 400: menu rows and tabs are not bold until current; buttons are always bold.
- Links are underlined at rest and lose the underline when hovered (`--link-decoration`, `--link-decoration-hover`), turning `--color-link-hover`. Text links and titles in tables and cards use `--color-link`; menu rows, bylines and comment authors use `--color-link-quiet`. News titles are `--color-heading`.
- Surfaces: `--color-page` beside the sheet, `--color-canvas` the sheet and, through `--fill-panel`, the 1px gaps. `--color-surface` is the body of every box, table cells and menu rows; `--color-surface-alt` the rail, the next-match cell, the figures, panel and form foot rows and meter tracks; `--color-surface-strong`, always with `--color-heading` text, the head strip of a news post or comment and the table head row.
- Bars: `--fill-bar` with `--color-bar-text` for box titles (menu, rail boxes, panels, form, dialog, table caption). `--fill-bar-alt` with `--color-bar-alt-text` for the strip under the banner, the tab row and quiet badges. `--fill-inverse` with `--color-inverse-text` for the banner and the footer; the wordmark is `--color-inverse-text`.
- Borders: boxes have a 1px `--color-border` outline; `--color-border-muted` divides rows inside a box; `--color-border-strong` is the sheet outline, the rules around the strip, the top edge of a squad card (`--border-width-strong`), the dialog outline and button outlines. Inputs use `--color-input-border`.
- Markers: `--fill-accent` with `--color-accent-text` for the current menu row, the badge, the page title square and meter bars; `--color-accent-alt` for a bold remark (`.ds-new`). Status badges are a `color-mix()` of 26% of the status colour over `--color-surface`, with `--color-text` and a 1px solid outline in the status colour: won, drawn, lost; active, on trial, unavailable. The destructive button is `--color-danger` text on `--color-danger-surface` with a `--color-danger` outline.
- `--color-fill-1` to `--color-fill-4` colour the squad cards; `--color-fill-3` is also the alternate table row.
- Every box takes `--shadow-panel`, which in this surface set is a light 1px inner line; the hero and the dialog take `--shadow-dialog`, a hard 3px ring.
- One-off values outside the token set: the status badge tints above; the notice edge is a `color-mix()` of 40% `--color-accent` in `--color-notice`; footer separators are at 0.6 opacity; poll bars carry their share as a percentage width.
- Tahoma and Verdana are system fonts; no substitution was needed.

## Components

- `ds-progress`: a quota or poll bar, a native `<progress>` with a label above (private-message quota).
- `ds-accordion`: a collapsible `<details>` entry for a forum jump or an FAQ answer.
- `ds-avatar`: the square picture of a member, placed under or beside the name.
- `.ds-page`: on `<body>`. Wrap everything in `.ds-page__frame`. `.ds-page__body` is the grid of `.ds-page__rail` and `.ds-page__main`; the views go in `.ds-page__main`, each a `.ds-view` section (`.ds-view--home` for the front page) whose blocks go in `.ds-view__body`. `.ds-page__legal` is the small print. `.ds-split` with `.ds-split__col` makes a wide and a narrow column inside a view. `.ds-block` puts a span on its own line.
- `.ds-banner`: the site header band. `.ds-banner__brand` holds the `.ds-brand` and a `.ds-banner__tagline`; `.ds-banner__aside` is a right-aligned column of short status lines with a bold `.ds-banner__link`. `.ds-strip` is the line under it with `.ds-strip__link` items.
- `.ds-brand`: the site's mark and name in the banner, linking home; one per page. `.ds-brand__mark` is a square tile dressed like the primary button (`--fill-button`, `--color-button-text`, a `--color-border-strong` outline, `--radius-control`, `--shadow-control`), twice `--text-display`, the mark drawn with square caps and a heavier stroke. `.ds-brand__name` is the name in `--font-heading` at `--text-display`, coloured `--color-inverse-text` for the banner it sits on. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo: replace both, keep the tile.
- `.ds-nav`: the site navigation, a menu box at the head of the rail. `.ds-nav__title` is its title strip, `.ds-nav__menu` the list, `.ds-nav__link` one row; the current row takes `is-current` (accent fill, bold).
- `.ds-sidebar`: the other boxes of the rail. Each `.ds-sidebar__block` has a `.ds-sidebar__title` and a `.ds-sidebar__list` of `.ds-sidebar__item` links. A shout is `.ds-sidebar__item--shout` with `.ds-sidebar__who` and `.ds-sidebar__when`; `.ds-sidebar__form` is the one-line field and button under it.
- `.ds-breadcrumb`: a plain small line: `.ds-breadcrumb__label`, `.ds-breadcrumb__link`, `.ds-breadcrumb__sep`, a bold `.ds-breadcrumb__current`.
- `.ds-hero`: the welcome box of the front page. `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__action` with the main action and one secondary button; `.ds-hero__aside` is the tinted cell for the next event, with `.ds-hero__label` and `.ds-hero__versus`. Never a banner image, never taller than its text.
- `.ds-page-header`: the head of an inner page, used instead of `.ds-hero`. `.ds-page-header__title` (the accent square is drawn by the class), `.ds-page-header__desc`, and `.ds-page-header__action` with one button on the right.
- `.ds-stat`: the row of summary figures: four boxed cells in one bordered block. Each `.ds-stat__item` has a `.ds-stat__value` in the heading font over a muted `.ds-stat__label`. `.ds-stat` is the row.
- `.ds-list`: news posts, one bordered box per `.ds-list__item`: `.ds-list__head` (the `.ds-list__title`, the byline, the date on the right), `.ds-list__body` (the excerpt) and `.ds-list__meta` (views, comment count, read more).
- `.ds-table`: roster, results and any data table. A `caption` is its title bar; `th` is a pale strong row with heading text; rows alternate with `is-alt`; `.ds-table__num` and `.ds-table__date` cells; `.ds-table__title` is the bold link of a row, `.ds-table__meta` a muted remark after it.
- `.ds-panel`: any titled box: `.ds-panel__title`, `.ds-panel__body` with `.ds-panel__line` paragraphs, and a tinted `.ds-panel__foot` row for a count and a button. A poll answer is `.ds-poll` (label, share, then a `.ds-meter` with its `.ds-meter__bar`, `--2` and `--3` for the other shares).
- `.ds-grid`: four squad or section cards, `--space-4` apart, each a `.ds-grid__cell` (`--2`, `--3`, `--4` for the other fills) with a strong top edge, a `.ds-grid__name` and `.ds-grid__meta` lines.
- `.ds-tabs`: a row of links on the secondary bar; `.ds-tabs__tab`, and `is-current` is a canvas cell with bold heading text. Use for the views of one table.
- `.ds-badge`: a small square label in the accent fill (captain, count). `.ds-badge--quiet` is the secondary bar. The status variants `.ds-badge--success` (won, active), `.ds-badge--warning` (drawn, on trial, open) and `.ds-badge--danger` (lost, unavailable) are a pale tint with a 1px outline. `.ds-new` is a bold remark in the second accent.
- `.ds-comment`: a reply under a post: `.ds-comment__head` (author, date) and `.ds-comment__body`. Stack them in `.ds-comments`.
- `.ds-prose`: long text such as a match report: h1, h2, h3, p, ul, strong, code.
- `.ds-link`: every link in text. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--quiet` for bylines and meta; `.ds-link--strong` for bold.
- `.ds-button`: the primary action, always bold. `.ds-button--secondary` is the pale button; `.ds-button--danger` the destructive action, once per confirmation; `is-hover`, `is-active`, `is-focus`, `is-disabled` (or `disabled`). Group buttons in `.ds-buttons`.
- `.ds-form`: a bordered box: `.ds-form__title`, a `.ds-form__body` of open `.ds-form__row` lines (a `.ds-form__label` with an optional `.ds-form__hint`, then the field), and `.ds-form__actions`. Controls take `.ds-form__input` (`--wide`, `--auto`, `--grow` inside a flex row, `is-invalid`, `is-disabled`), errors `.ds-form__error`, checkboxes `.ds-form__check` and `.ds-form__checkbox`.
- `.ds-notice`: a pale line for information; `.ds-notice--error` is the error box with a `.ds-notice__title`.
- `.ds-pagination`: a line of page links: `.ds-pagination__label`, `.ds-pagination__link`; `is-current` is a small bordered cell. `.ds-toolbar` puts links or buttons on the left and the page links on the right.
- `.ds-dialog`: a confirmation box in the page flow, strong outline, `.ds-dialog__title`, `.ds-dialog__body` with `.ds-dialog__text` and centred `.ds-dialog__actions`. `.ds-dialog__backdrop` is a flat band in the overlay colour. It never floats over the page.
- `.ds-empty`: a bordered box with a centred `.ds-empty__title`, one muted `.ds-empty__text` and one button.
- `.ds-footer`: the inverse band closing the sheet: a centred `.ds-footer__row` of `.ds-footer__link` items with `.ds-footer__sep` between them, and a `.ds-footer__line` under it.

## Never

- `border-radius <= 0px`: every box is square in all references.
- `box-shadow-blur <= 0px`: the only shadows are hard 1px lines and the 3px ring of the dialog; nothing is soft.
- `text-shadow = none`: no text shadow was measured in the period references.
- `transition = none`: states switch at once.
- `animation = none`: nothing moves.
- `font-size <= 18px`: only the wordmark is 18px.
- `font-size >= 10px`: 10px is the smallest size.
- `font-weight <= 700`: bold is the heaviest weight.
- `font-families <= 3`: Tahoma, Verdana for headings, a monospace for code.
- `border-width <= 2px`: 1px everywhere, 2px on the emphasised edges.
- `letter-spacing <= 0px`: no tracking.
- `uppercase-text <= 0%`: nothing is transformed to uppercase.
- `underlined-links >= 50%`: links are underlined; 88% to 94% in the references this follows.
- `gradient-fills <= 20%`: gradients are on title strips, the secondary bar, badges and buttons only.
- `row-gap <= 8px`: table rows are 1px apart; news posts are separate boxes 8px apart.
- `block-gap <= 20px`: neighbouring blocks are 12px apart.
- `content-width >= 80%`: the sheet takes 88% of the viewport; never a narrow column.

## Extending

Derive a new component from the nearest one in the specimen: a new rail box from `.ds-sidebar__block`, a new feed from `.ds-list`, a new side box from `.ds-panel`, a new data view from `.ds-table`. Build it as a 1px `--color-border` outline on `--fill-panel` with `--space-1` of padding and gap, a title in `--fill-bar` with `--color-bar-text` or a head strip in `--color-surface-strong` with `--color-heading` text, and a body in `--color-surface`. Keep text on the fills it already sits on in the specimen and never place text on the page colour. Use tokens only: no raw colours or lengths, no new radius, shadow or transition, no font size outside the type tokens, and `--space-5` as the gap to the next block. New site sections get a row in the menu box; new always-visible boxes go into the rail under it; everything else goes into the main column.
