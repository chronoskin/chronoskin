# Bulletin board forum, 2002 to 2007

## Summary

This is the look of hosted bulletin-board software in the first half of the 2000s: a white sheet on a grey page, filled edge to edge with dense tables of boards, topics and counts in 10px and 11px Verdana. Colour comes from one cold blue family, a navy primary bar with a short highlight at its top edge, a slate secondary bar and three pale blue-grey cell surfaces, with orange and dark red used only as small markers. Every box is square, flat and separated from its neighbour by a one-pixel white gap or a one-pixel rule; there are no shadows, no rounded corners and no motion.

## Layout

- The page is fluid. `--size-page` is `100%`; `.ds-page` has `--space-5` (10px) of padding, so the sheet `.ds-page__frame` takes about 98% of the viewport at every width. Designed at 1024px. Never centre a narrow column.
- Order from the top: masthead (wordmark left, three stacked tool links right), primary navigation bar, sub-navigation strip, then `.ds-page__body`, then the footer band and one legal line, both inside the sheet. No text sits on the page colour.
- `.ds-page__body` holds the views; inside a view, `.ds-view__body` is a single column of blocks `--space-5` apart. An optional two-column area, `.ds-layout`, puts a `--size-sidebar` (200px) sidebar on the left of the board table.
- Board table columns: status `--size-status` (28px), board (fluid), each count `--size-count` (56px), last post `--size-last` (150px). Body rows are at least `--size-row` (34px); bars and title strips are `--size-bar` (22px).
- Spacing scale: `--space-1` 1px, `--space-2` 2px, `--space-3` 4px, `--space-4` 6px, `--space-5` 10px, `--space-6` 16px.
  - `--space-1`: the gap between cells and rows inside any bordered block (table, panel, list, grid, form), and the vertical padding of buttons and inputs.
  - `--space-2`: vertical padding of bars and title strips, gap between stacked small lines.
  - `--space-3`: vertical padding of cells, gap between inline items (breadcrumb, footer links).
  - `--space-4`: horizontal padding of cells, strips and panel bodies; gap between buttons.
  - `--space-5`: page padding, body padding, and the gap between neighbouring blocks. This is the only block gap.
  - `--space-6`: indent of lists and sub-items, padding of dialog and empty-state bodies, horizontal padding of tabs.
- Inner pages (a board, a topic, a settings page) keep the masthead, both navigation bars and the footer, and drop the hero and the board sidebar: `.ds-page__body` is one full-width column. From the top: `.ds-breadcrumb`, then `.ds-page-header` in place of the hero (title and one line of description on the left, the main action bottom-aligned on the right, no fill and no rule), then `.ds-tabs` when the page has several views, then the working block (`.ds-table`, `.ds-list` or `.ds-form`) at full width, then `.ds-pagination` right-aligned under it. All of these are `--space-5` apart, like every other block. A statistics strip (`.ds-stat`) belongs at the foot of an index page, above the who-is-online panel.
- A thread is a stack of `.ds-post` blocks, each a `--size-last` (150px) author cell beside a fluid message cell, with a slate date strip above and a pale tool strip below. Under the last post, `.ds-toolbar` puts the topic buttons on the left and the page numbers on the right.
- Some inner pages have a menu of their own (control panel, help): they keep the breadcrumb and page header at full width and then use `.ds-layout`: the menu in the `--size-sidebar` column, the working blocks beside it.
- Views. The specimen is a six-screen example board; the primary bar has one item per view and marks the current one. `index` (home) is the board index: welcome block, announcement, sidebar beside the board table and the active topics list, then statistics and who is online. `board` is one board: sub-board grid, tabs, topic table, legend and page numbers. `topic` is a thread: three posts, the delete confirmation that a post's Delete link opens, and the topic buttons. `post` is the reply form with its error notice and a review of the latest posts. `panel` is the member's control panel: account menu, activity figures, subscribed topics and an empty inbox. `rules` is the help page: contents menu, the rules as long text, and related links.
- Form rows are a `--size-label` (24%) label cell and a field cell; text inputs are `--size-field` (280px) unless `--wide`. A dialog is `--size-dialog` (56%) wide, centred in the flow.

## Typography and colour roles

- One family for everything: `--font-body`, `--font-heading` and `--font-ui` are the same Verdana stack. `--font-mono` (Courier New) is for inline code only.
- Sizes: `--text-small` 10px for descriptions, meta, counts captions, strips and sub-navigation, and it carries most of the text on a page; `--text-base` and `--text-ui` 11px for running text, numbers, buttons, inputs and navigation; `--text-h2` 12px for bar and strip titles; `--text-h1` and `--text-large` 13px for the page heading, the hero title and board and topic titles; `--text-display` 16px for the wordmark only. Nothing is larger than 16px.
- Line height is `normal` everywhere. Headings are bold, never uppercase, with no tracking.
- Weight: `--weight-ui` is 700 for navigation, tabs and the primary button; the secondary button and inputs use `--weight-body`.
- Links in running text are `--color-link` and underlined. Titles of boards and topics, breadcrumb links, moderator and member names use `--color-link-quiet` (navy), also underlined (`--link-decoration-quiet`). Only links that sit on a bar (navigation, sub-navigation, tabs) have no underline until hovered. Hover turns any link `--color-link-hover` (orange).
- Surfaces: `--color-page` outside the sheet; `--color-canvas` is the sheet and, through `--fill-panel`, the white that shows in the 1px gaps between cells. `--color-surface` is the main cell, `--color-surface-alt` the status, count and last-post cells and form labels, `--color-surface-strong` the category rows, panel and sidebar title strips, form action row and current page number.
- Bars: `--fill-bar` (navy `--color-bar` with a lighter top edge) for the primary navigation, table heads and the title of a main board, list, form or dialog, always with `--color-bar-text`. `--fill-bar-alt` (flat slate) for the sub-navigation strip, date strips, panel sub-strips, the closing strip of a table, idle tabs and quiet badges. `--fill-inverse` for the footer band.
- Borders: blocks have a 1px `--color-border` outline; `--color-border-muted` is for rules inside a box and the breadcrumb outline; `--color-border-strong` at `--border-width-strong` (2px) is only for the rule under the navigation bar and tabs and for the dialog outline. Inputs use `--color-input-border`.
- Markers: `--color-accent` for new-post icons and the badge; `--color-accent-alt` for emphasised text that is not an error; `--color-danger` and `--color-danger-surface` for errors; `--color-notice` for announcements. `--color-success` (dark green) and `--color-warning` (amber) are the two member-group name colours of the era; here they tint status badges only. Status badges are a `color-mix()` of 30% of the status colour (`--color-success`, `--color-warning`, `--color-danger`) over `--color-surface`, with `--color-text` and a 1px outline in the status colour, so they read under every palette. The destructive button is the error pair: `--color-danger` text on `--color-danger-surface` with a `--color-danger` outline.
- `--color-fill-1` to `--color-fill-4` repeat the pale surfaces; the style has no decorative block colours.
- One-off mixes outside the token set: the status badge tints described above; the dashed edge of `.ds-notice` is `color-mix()` of `--color-accent` and `--color-notice` in equal parts; the secondary button uses the `outset` border style and `inset` when pressed, and the notice uses `dashed`, instead of `--border-style`.
- The references used Verdana with Geneva, Arial and Helvetica fallbacks, all system fonts; no substitution was needed.

## Components

- `ds-progress`: a quota or poll bar, a native `<progress>` with a label above (private-message quota).
- `ds-accordion`: a collapsible `<details>` entry for a forum jump or an FAQ answer.
- `ds-avatar`: the square picture of a member, placed under or beside the name.
- `.ds-page`: on `<body>`. Wrap everything in `.ds-page__frame`; `.ds-page__body` holds the views. Each view is a `.ds-view` section (`.ds-view--home` for the index) and its blocks go in `.ds-view__body`. `.ds-page__legal` is the small print at the foot of the sheet, under the footer. `.ds-layout` with `.ds-layout__main` makes the sidebar-plus-main area. `.ds-icon` sizes an inline SVG icon, `.ds-sprite` hides the SVG symbol sheet, `.ds-block` puts a span on its own line.
- `.ds-nav`: the site header. `.ds-nav__head` holds `.ds-nav__brand` (the `.ds-brand` with a `.ds-nav__tagline` line under it) and `.ds-nav__tools`. `.ds-nav__menu` is the primary bar of `.ds-nav__link` items; the current one takes `is-current` and becomes a pale cell; in the specimen the same dress is applied by one `:has()` rule per view. `.ds-nav__sub` is the slate strip of `.ds-nav__sublink` items for account links.
- `.ds-brand`: the site's mark and name at the left of the masthead, linking to the index; one per page. `.ds-brand__mark` is a small square tile dressed like the primary button (`--fill-button` over `--color-button`, a `--border-width` outline in `--color-border-strong`, `--radius-control`, `--shadow-control`) with the mark drawn in `--color-button-text`; it is 1.5 times `--text-display`, so it follows the type set. `.ds-brand__name` is the name in `--font-heading` at `--text-display` and `--weight-display`, with `--display-tracking` and `--heading-transform`, coloured `--color-heading` for the white masthead it sits on. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo: replace both, keep the tile.
- `.ds-breadcrumb`: one pale bordered line under the header, with `.ds-breadcrumb__link`, `.ds-breadcrumb__sep` and `.ds-breadcrumb__current`.
- `.ds-hero`: the welcome block at the top of an index page. A 13px bold `.ds-hero__title`, a 10px `.ds-hero__lead`, the main action in `.ds-hero__action`, and board totals right-aligned in `.ds-hero__aside`. It has no fill and ends in a 1px rule. Never make it a banner.
- `.ds-page-header`: the head of an inner page, used instead of `.ds-hero` on every page below the index. `.ds-page-header__title` is 13px bold navy, `.ds-page-header__desc` one 10px line (description, moderators), `.ds-page-header__action` holds one primary button at the right. No box, no fill.
- `.ds-prose`: long text such as rules or a post body: h1, h2, h3, p, ul, strong, code, hr.
- `.ds-link`: every link in text. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--quiet` for navy tool, meta and name links; `.ds-link--strong` for bold.
- `.ds-button`: primary action, blue with white bold text and a navy 1px border. `.ds-button--secondary` is the plain grey raised system button. `.ds-button--danger` is for a destructive action (delete, ban, prune): error-red text on the pale error surface with a red outline; use it once per confirmation, next to a secondary button. `is-hover`, `is-active`, `is-focus`, `is-disabled` (or `disabled`). Group buttons in `.ds-buttons`.
- `.ds-form`: a bordered block of rows. `.ds-form__title` (primary bar), `.ds-form__row` with `.ds-form__label` (plus `.ds-form__hint`) and `.ds-form__field`; controls take `.ds-form__input` (`--wide`, `--auto`, `is-invalid`, `is-disabled`), errors `.ds-form__error`, checkboxes `.ds-form__check` and `.ds-form__checkbox`, buttons in `.ds-form__actions`.
- `.ds-table`: the board or topic index and any data table. `th` is the primary bar (`.ds-table__head--left` for the first text column); `.ds-table__group` is a category row; cells `.ds-table__status` (`is-read`), `.ds-table__num`, `.ds-table__last`; inside the main cell `.ds-table__title`, `.ds-table__desc`, `.ds-table__meta`. End a board with a `.ds-table__foot` row.
- `.ds-infobar`: the 10px line of tools and time zone under a table; `.ds-infobar__side` is right-aligned.
- `.ds-list`: topic or post rows outside a table. `.ds-list__head` (primary bar), `.ds-list__date` (slate strip between groups), `.ds-list__item` (`is-alt` for the alternate row) with `.ds-list__title`, `.ds-list__meta`, `.ds-list__aside`.
- `.ds-post`: one message of a thread. `.ds-post__head` is the slate strip with the date and post number; `.ds-post__author` the pale cell with `.ds-post__name` and `.ds-post__meta` lines; `.ds-post__body` holds `.ds-post__text` paragraphs, a `.ds-post__quote` with its `.ds-post__cite`, inline `.ds-post__code` and the `.ds-post__sig` signature under a faint rule; `.ds-post__foot` is the pale strong strip of quiet tool links. `.ds-toolbar` is the row under a thread: `.ds-buttons` left, `.ds-pagination` right.
- `.ds-panel`: any titled box, such as statistics or who is online. `.ds-panel__title` is a pale strong strip; add `.ds-panel__title--bar` for a main block. `.ds-panel__sub` is a slate sub-strip, `.ds-panel__body` (`--alt`) the cells, `.ds-panel__line` a paragraph.
- `.ds-grid`: four equal cells with 1px gaps in one bordered block, for sub-boards or statistics. `.ds-grid__title`, `.ds-grid__cell` (`--2`, `--3`, `--4` for the other fills), `.ds-grid__name`, `.ds-grid__meta`. Cells never have gaps wider than 1px or their own borders.
- `.ds-stat`: summary figures as one bordered strip of equal count cells with 1px gaps, never separate cards. `.ds-stat__title` is an optional pale strong strip across the top; each `.ds-stat__item` is a centred `--color-surface-alt` cell with a 13px bold navy `.ds-stat__value` over a 10px muted `.ds-stat__label`. Use three to six items.
- `.ds-tabs`: a row of square slate `.ds-tabs__tab` items on a 2px navy rule; `is-current` takes the primary bar fill. Use for views of one list.
- `.ds-badge`: a small square orange label, 10px bold. `.ds-badge--quiet` is slate. The status variants `.ds-badge--success` (solved, online, approved), `.ds-badge--warning` (pending, awaiting) and `.ds-badge--danger` (locked, banned, overdue) are a pale tint of the status colour with body text and a 1px outline in that colour. Put a status badge after a title or in a table cell. `.ds-new` is bold dark red text for a "new" remark. `.ds-legend` with `.ds-legend__item` (`is-read`) and `.ds-legend__label` explains the status icons.
- `.ds-sidebar`: a column of `.ds-sidebar__block` boxes, each with `.ds-sidebar__title`, a `.ds-sidebar__list` of `.ds-sidebar__item` (`--head` for a group link, `--sub` for an indented one) and an optional slate `.ds-sidebar__foot` with a `.ds-sidebar__footlink`.
- `.ds-notice`: a pale yellow centred line with a dashed edge for announcements. `.ds-notice--error` is the salmon error box with a solid red edge and a `.ds-notice__title`.
- `.ds-pagination`: right-aligned under a list. `.ds-pagination__label`, then small bordered `.ds-pagination__link` boxes; `is-current` for the current page.
- `.ds-dialog`: a confirmation box in the page flow, 2px navy outline, `.ds-dialog__title` bar, `.ds-dialog__body` with `.ds-dialog__text` and `.ds-dialog__actions`. `.ds-dialog__backdrop` is a flat band in the page colour. It never floats over the page.
- `.ds-empty`: a bordered block with one centred message (`.ds-empty__body`, `.ds-empty__text`) and one button, for a list that has nothing in it yet. `.ds-empty__title` is an optional pale strong strip naming the list.
- `.ds-footer`: the navy band closing the sheet, a `.ds-footer__row` of `.ds-footer__link` items with `.ds-footer__sep` between them.

## Never

- `border-radius <= 0px`: no reference has a rounded corner in CSS; every box is square.
- `box-shadow = none`: no box shadow was measured in any reference.
- `text-shadow = none`: no text shadow was measured in any reference.
- `transition = none`: hover states switch at once.
- `animation = none`: nothing moves.
- `font-size <= 16px`: the largest text in any reference is 16px, and only the wordmark uses it.
- `font-size >= 10px`: 10px is the smallest and the most common size.
- `font-families <= 2`: one sans-serif stack, plus monospace for code.
- `border-width <= 2px`: borders are 1px, with 2px only on emphasised edges.
- `letter-spacing <= 1px`: at most one reference tracks its headings, by 1px; this pack uses none.
- `uppercase-text <= 0%`: no reference transforms text to uppercase.
- `underlined-links >= 50%`: between 54% and 98% of link text is underlined in the references.
- `gradient-fills <= 10%`: gradients appear only on bars and on the primary and destructive buttons.
- `row-gap <= 1px`: rows touch or are separated by a 1px line.
- `block-gap <= 20px`: neighbouring blocks are 10px apart; 20px is the widest measured.
- `content-width >= 90%`: the board fills the viewport; no reference used less than 95%.

## Extending

Derive a new component from the nearest one in the specimen: a new titled box from `.ds-panel`, a new row layout from `.ds-table` or `.ds-list`, a new strip from `.ds-nav__sub`. Build it as a 1px `--color-border` outline on `--fill-panel`, with cells in the three surface colours separated by `--space-1`, and a title in `--fill-bar` or `--color-surface-strong`. Use tokens only: no raw colours or lengths, no new radius, shadow or transition, no font size outside the type tokens, and `--space-5` as the gap to the next block.
