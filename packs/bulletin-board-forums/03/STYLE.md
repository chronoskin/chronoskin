# Bulletin board forum, community front page, 2002 to 2007

## Summary

This is the forum as the community section of a larger site, common from about 2003 to 2006: a white masthead with the wordmark and a row of icons over labels, a section title with a few member links, and then a front page that leads with the latest discussions rather than the list of boards. The discussions run down a wide main column under date strips, each with an excerpt and a right-aligned byline, and a wide column of titled boxes on the right carries the forum tree and recent articles; inner pages go back to one full-width column, or keep the right column for a menu of their own. One petrol blue with orange bar titles on cool grey cells, flat fills, hairline borders, no underlines until hover.

## Layout

- The page is fluid. `--size-page` is `100%`; `.ds-page` has `--space-5` (10px) of padding, so the outlined sheet `.ds-page__frame` takes about 98% of the viewport. Designed at 1024px. Never centre a narrow column.
- Order from the top: a thin utility strip `.ds-nav__top` in the inverse fill with three right-aligned links; the masthead `.ds-nav__head` on the canvas, wordmark left and the icon navigation right, each item `--size-navitem` (78px) wide with a `--size-navicon` (26px) icon over its label, the current one an outlined pale cell; a 1px strong rule; then `.ds-page__body`; then the pale footer band and one closing line `.ds-page__legal`, both inside the sheet. No text sits on the page colour.
- The front page: `.ds-hero` across the full width (display-size section title and lead on the left, two rows of member links and the main action on the right, a ruled foot line with visit times and shortcuts), an optional `.ds-notice`, then `.ds-layout`: the main column first and a `--size-sidebar` (27%) sidebar on the right. The main column holds the `.ds-list` of latest discussions, the `.ds-stat` line and the `.ds-grid` of section boxes. A full-width `.ds-panel` closes the page.
- Spacing scale: `--space-1` 1px, `--space-2` 2px, `--space-3` 4px, `--space-4` 6px, `--space-5` 10px, `--space-6` 16px.
  - `--space-1`: the ruled gap between cells of a table, form or dialog; vertical padding of buttons and inputs.
  - `--space-2`: vertical padding of bars and strips, gap between stacked small lines.
  - `--space-3`: vertical padding of table cells, gap between inline items and between navigation items.
  - `--space-4`: padding of list items and strips, the gap between grid boxes and between buttons.
  - `--space-5`: page padding, the gap between neighbouring blocks and between the two columns, padding of panel bodies. This is the only block gap.
  - `--space-6`: horizontal padding of the masthead and the body, sub-item indent, padding of dialog and empty-state bodies.
- Bars and title strips are `--size-bar` (24px) tall. Topic table columns: status `--size-status` (28px), topic (fluid), replies, author and views `--size-count` (60px) each, last post `--size-last` (150px, centred); rows are at least `--size-row` (32px).
- Each screen is a `.ds-view` inside `.ds-page__body`; its `.ds-view__body` stacks the blocks `--space-5` apart. The utility strip, the masthead and the footer are written once, outside the views.
- Inner pages (a forum, a discussion, the posting form) keep the utility strip, the masthead and the footer and drop the hero and the sidebar: one full-width column. From the top: `.ds-breadcrumb` as a secondary strip, then `.ds-page-header` in place of the hero (an open title at the heading size with one line of description, the main action bottom-right, closed by a strong rule), then the working blocks. A forum has `.ds-tabs` as folder tabs on a rule, the `.ds-table`, a `.ds-pagination` text line under it and the `.ds-legend`. A discussion is a run of `.ds-post` boxes, each with a `--size-last` (150px) author cell on the left, closed by a `.ds-toolbar` with the topic buttons on the left and the page numbers on the right; a confirmation `.ds-dialog` sits in the flow between the posts and the toolbar. The posting page has any `.ds-notice--error` first, then `.ds-form` at full width. All are `--space-5` apart.
- Inner pages with a menu of their own (the member's desk, help) keep the breadcrumb and page header at full width and then use `.ds-layout` as the front page does: working blocks in the main column, the menu and one more box in the `--size-sidebar` column on the right. Fill both columns so they end near each other.
- Views. The specimen is a six-screen example community; the icon navigation has one item per view and marks the current one. `index` (home) is the community front page: section title with the main action, an announcement, the latest discussions beside the forum tree and recent articles, the figures line, the desks grid and who is online. `board` is one forum: tabs, the topic table, page numbers and the icon legend. `topic` is a discussion: three posts, the delete confirmation that a post's Delete link opens, and the topic buttons with page numbers. `post` is the new topic form with its error notice and a short box of posting notes. `desk` is the member's own page: watched topics and an empty private message inbox beside the account menu and the critique ledger. `rules` is the help page: the house rules as long text and related links beside the help contents and the moderator list.
- Form rows are a right-aligned `--size-label` (32%) label cell and a field cell; text inputs are `--size-field` (220px) unless `--wide`. A dialog is `--size-dialog` (48%) wide, centred in the flow.

## Typography and colour roles

- `--color-page` is the grey #e5e5e5 and `--color-canvas` the white sheet. Cells are cool greys: `--color-surface` #efefef for list items, sidebar and panel bodies and the main table cell, `--color-surface-alt` #dee3e7 for status, count and last-post cells, form fields, idle tabs and the footer band, `--color-surface-strong` #d1d7dc for group rows, plain panel titles, the form action row and the current page number.
- One saturated colour, the petrol blue #006699: it is `--color-bar`, `--color-inverse`, `--color-heading`, `--color-link`, `--color-link-quiet`, `--color-border-strong`, `--color-accent` and `--color-button`.
- `--color-bar-text` is the light orange #ffa34f, bold, on the petrol bar only (list, sidebar, panel, form and dialog titles, table heads); never put it on a pale surface. `--color-bar-alt` is the grey-blue #98aab1 with black `--color-bar-alt-text`: date strips, sidebar sub-strips and foot strips, the breadcrumb strip, and the colour of navigation and read-status icons. The utility strip is `--color-inverse` with `--color-inverse-text`.
- Text is `--color-text` #000000, secondary text `--color-text-muted` #444444, `--color-heading-alt` #333333 for second-level headings in prose. A hovered link turns `--color-link-hover` #333333; visited is the lighter #4691b6. Only `--color-text-muted` is ever set directly on the page colour.
- `--color-border` (#98aab1) outlines blocks and inputs; `--color-border-muted` (#d1d7dc) divides inside a box; `--color-border-strong` rules off the masthead, the page header and the footer and outlines the dialog.
- Markers: `--color-accent-alt` is the dark green #006600 for emphasised remarks, shared with `--color-success`. Orange #ffa34f returns as `--color-warning`. `--color-danger` #c60202 on `--color-danger-surface` #fddbcc. `--color-notice` is the pale #e8eef2. Text on the accent is `--color-accent-text`. Status badges are a `color-mix()` of 30% of the status colour (`--color-success`, `--color-warning`, `--color-danger`) over `--color-surface`, with `--color-text` and a 1px outline in the status colour, so they read under every palette. The destructive button is the error pair: `--color-danger` text on `--color-danger-surface` with a `--color-danger` outline.
- `--color-fill-1` to `--color-fill-4` are the four pale greys #efefef, #d1d7dc, #dee3e7, #e8eef2; they fill the section boxes of `.ds-grid`, with `--color-link-quiet` names and `--color-text-muted` meta.
- `--font-body` and `--font-heading` are Verdana with Helvetica. `--font-ui` is Tahoma: navigation labels, tabs and buttons, at `--text-ui` 10px, bold (`--weight-ui` 700) and uppercase (`--ui-transform`). Inputs use `--font-body` at `--text-small`.
- Two sizes do most of the work: `--text-small` 10px for excerpts, bylines, descriptions, table cells and strips, and `--text-base` / `--text-large` / `--text-h2` 13px for running text, discussion and topic titles and bar titles. `--text-h3` is 12px. `--text-h1` and `--text-display` are 18px for the page heading, the section title of the hero and the wordmark; nothing is larger.
- Line height is `normal`; headings are bold with no tracking and no transform.
- No link is underlined at rest: `--link-decoration` and `--link-decoration-quiet` are `none`, and `--link-decoration-hover` adds the underline. Links are told apart by colour, so link text must never share the body text colour.
- Completely flat: `--fill-bar`, `--fill-bar-alt`, `--fill-button`, `--fill-inverse` and `--fill-accent` are the plain colours. No gradient anywhere.
- `--fill-panel` is `--color-border-strong`: the one-pixel gaps between cells of a table, form or dialog are drawn in the strong border colour. Never set text directly on it. Lists, panels and sidebar boxes are not gapped; their rows touch or are divided by a `--color-border-muted` rule.
- Hairlines only: `--border-width` and `--border-width-strong` are both 1px solid. All `--radius-*` are 0 and every `--shadow-*` is `none`. `--focus-ring` is a 1px dotted line; `--transition` is `none`.
- One-off choices outside the token set: the idle navigation item has a `transparent` border so the current item's outline does not shift it; the current tab is moved down by `--space-1` to cover the rule.

## Components

- `ds-progress`: a quota or poll bar, a native `<progress>` with a label above (private-message quota).
- `ds-accordion`: a collapsible `<details>` entry for a forum jump or an FAQ answer.
- `ds-avatar`: the square picture of a member, placed under or beside the name.
- `.ds-page`: on `<body>`. Wrap everything in `.ds-page__frame`; put blocks in `.ds-page__body`. In the specimen each screen is a `.ds-view` (the first also `.ds-view--home`), shown one at a time by its id in the address, with a `.ds-view__body` column inside; a real page needs neither. `.ds-page__legal` is the closing line at the foot of the sheet. `.ds-layout` with `.ds-layout__main` first and `.ds-sidebar` second makes the two-column area. `.ds-icon` sizes a small inline SVG icon, `.ds-sprite` hides the symbol sheet, `.ds-block` puts a span on its own line.
- `.ds-nav`: the site header. `.ds-nav__top` is the utility strip of `.ds-nav__toplink` items. `.ds-nav__head` holds `.ds-nav__brand` (the `.ds-brand` with a `.ds-nav__tagline` line under it) and `.ds-nav__menu`, whose `.ds-nav__link` items are a `.ds-nav__icon` over a label; the current one takes `is-current` and becomes an outlined pale cell (the specimen marks it from the view in the address instead).
- `.ds-brand`: the site's mark and name at the left of the masthead, linking to the index; one per page. `.ds-brand__mark` is a small square tile dressed like the primary button (`--fill-button` over `--color-button`, a `--border-width` outline in `--color-border-strong`, `--radius-control`, `--shadow-control`) with the mark drawn in `--color-button-text`; it is 1.5 times `--text-display`, so it follows the type set. `.ds-brand__name` is the name in `--font-heading` at `--text-display` and `--weight-display`, with `--display-tracking` and `--heading-transform`, coloured `--color-heading` for the canvas masthead it sits on. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo: replace both, keep the tile.
- `.ds-breadcrumb`: a secondary strip with `.ds-breadcrumb__link`, `.ds-breadcrumb__sep` (`::`) and `.ds-breadcrumb__current`.
- `.ds-hero`: the section head of the front page. `.ds-hero__title` at the display size, `.ds-hero__lead` one small line; `.ds-hero__aside` holds `.ds-hero__links` (rows of `.ds-hero__linkitem`, a small icon and a link) and the main action `.ds-hero__action`; `.ds-hero__foot` is the ruled line of times and shortcuts with a right-aligned `.ds-hero__side`. No fill, no box.
- `.ds-page-header`: the head of an inner page. `.ds-page-header__title`, `.ds-page-header__desc`, and `.ds-page-header__action` with one primary button at the right; a strong rule closes it.
- `.ds-prose`: long text on the canvas in a hairline box: h1, h2, h3, p, ul, strong, code, hr.
- `.ds-link`: every link in text. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--quiet` and `.ds-link--strong`.
- `.ds-button`: primary action, flat, bold uppercase. `.ds-button--secondary` is the pale button, `.ds-button--danger` the destructive one, error text on the error surface with an error outline. `is-hover`, `is-active`, `is-focus`, `is-disabled` (or `disabled`). Group buttons in `.ds-buttons`.
- `.ds-form`: a ruled block. `.ds-form__title` (centred primary bar), `.ds-form__row` with a right-aligned `.ds-form__label` (plus `.ds-form__hint`) and a `.ds-form__field`; controls take `.ds-form__input` (`--wide`, `--auto`, `is-invalid`, `is-disabled`), errors `.ds-form__error`, checkboxes `.ds-form__check` and `.ds-form__checkbox`, buttons centred in `.ds-form__actions`.
- `.ds-table`: the topic index of an inner page and any data table. `th` is the primary bar (`.ds-table__head--left` for the text column); `.ds-table__group` is a pale group row; cells `.ds-table__status` (`is-read`), `.ds-table__num`, `.ds-table__last`; inside the main cell `.ds-table__title` and `.ds-table__meta`. It is not used on the front page.
- `.ds-list`: the latest discussions. `.ds-list__head` (centred primary bar), `.ds-list__date` strips between days, `.ds-list__item` (`is-alt`) with an icon and a `.ds-list__body` holding `.ds-list__title`, a quoted `.ds-list__excerpt` and a right-aligned `.ds-list__meta` byline.
- `.ds-post`: one message of a discussion, a hairline box on `--color-surface`. `.ds-post__head` is a secondary strip with the date left and the post number right; `.ds-post__author` is the centred `--color-surface-alt` cell with `.ds-post__name` (a quiet bold link at `--text-large`) and `.ds-post__meta` lines; `.ds-post__body` holds `.ds-post__text` paragraphs, an outlined `.ds-post__quote` with its `.ds-post__cite`, inline `.ds-post__code` and an italic ruled `.ds-post__sig`; `.ds-post__foot` is a `--color-surface-strong` strip of small links, profile left and post tools right.
- `.ds-toolbar`: the row under a discussion: `.ds-buttons` on the left, `.ds-pagination` on the right.
- `.ds-panel`: a full-width titled box. `.ds-panel__title` (`--bar` for the primary bar), `.ds-panel__sub` strip, `.ds-panel__row` for an icon beside a `.ds-panel__body` (`--alt`), `.ds-panel__line` paragraphs.
- `.ds-grid`: section boxes two across with a `--space-4` gap, each `.ds-grid__cell` (`--2`, `--3`, `--4` for the other fills) outlined, with `.ds-grid__name` and `.ds-grid__meta`; `.ds-grid__title` is an open heading on a faint rule.
- `.ds-stat`: `.ds-stat` is the row: one centred line of text on the canvas. Each figure is a `.ds-stat__item` made of a `.ds-stat__label` and a bold `.ds-stat__value`. Never boxes.
- `.ds-tabs`: folder tabs on a hairline; `.ds-tabs__tab` is a pale outlined tab, `is-current` is the canvas colour and joins the page below.
- `.ds-badge`: a small square label. `.ds-badge--quiet`, The status variants `.ds-badge--success` (solved, online, approved), `.ds-badge--warning` (pending, awaiting) and `.ds-badge--danger` (locked, banned, overdue) are a pale tint of the status colour with body text and a 1px outline in that colour. `.ds-new` is bold text for a remark. `.ds-legend` with `.ds-legend__item` (`is-read`) and `.ds-legend__label` explains the status icons.
- `.ds-sidebar`: the right column of `.ds-sidebar__block` boxes, each with a centred `.ds-sidebar__title` bar, an optional `.ds-sidebar__sub` strip, a `.ds-sidebar__list` of `.ds-sidebar__item` (`--head` for a section, `--sub` for an indented forum, `.ds-sidebar__byline` for an author) and an optional `.ds-sidebar__foot` with a `.ds-sidebar__footlink`.
- `.ds-notice`: a centred pale line in a hairline box with a `.ds-notice__title`. `.ds-notice--error` is the error box, title on its own line.
- `.ds-pagination`: a line of text under a table or at the right of a `.ds-toolbar`: `.ds-pagination__label` on the left, then bold `.ds-pagination__link` items separated by `.ds-pagination__sep` commas; `is-current` is a pale chip.
- `.ds-dialog`: a centred information box in the page flow: `.ds-dialog__title` bar, `.ds-dialog__body` with `.ds-dialog__text` and `.ds-dialog__actions`. `.ds-dialog__backdrop` is a flat band. It never floats over the page.
- `.ds-empty`: a pale box for a list with nothing in it: an optional centred `.ds-empty__title` in the primary bar, then `.ds-empty__body` with one centred bold `.ds-empty__text` and one button.
- `.ds-footer`: the pale band closing the sheet under a strong rule: a `.ds-footer__row` of `.ds-footer__link` items with `.ds-footer__sep` bars, then `.ds-footer__note`.

## Never

- `border-radius <= 0px`: every box is square.
- `box-shadow = none`: no shadow in any reference.
- `text-shadow = none`: no text shadow in any reference.
- `gradient-fills <= 0%`: bars and buttons are flat colour.
- `transition = none`: states switch at once.
- `animation = none`: nothing moves.
- `border-width <= 1px`: hairlines only.
- `font-size <= 18px`: the wordmark and page heading are the largest text.
- `font-size >= 10px`: 10px is the smallest size.
- `font-weight <= 700`: bold is the heaviest weight.
- `font-families <= 3`: Verdana, Tahoma for controls, a monospace for code.
- `letter-spacing <= 0px`: no tracking.
- `uppercase-text <= 15%`: only navigation, tabs and buttons are uppercase.
- `underlined-links <= 10%`: links are underlined on hover only.
- `row-gap <= 1px`: rows touch or are divided by a hairline.
- `block-gap <= 20px`: neighbouring blocks are 10px apart.
- `content-width >= 90%`: the two columns together fill the viewport.

## Extending

Derive a new component from the nearest one in the specimen: a new right-column box from `.ds-sidebar__block`, a new feed from `.ds-list`, a new full-width box from `.ds-panel`, a new data view from `.ds-table`. Give it a 1px `--color-border` outline, a title in `--fill-bar` (centred) or a strip in `--fill-bar-alt`, and a body in `--color-surface`; divide rows with `--color-border-muted` rather than gaps, except in tables and forms, which use `--space-1` gaps over `--fill-panel`. Keep text on the fills it already sits on in the specimen. Use tokens only: no raw colours or lengths, no new radius, shadow, gradient or transition, no font size outside the type tokens, and `--space-5` as the gap to the next block. On the front page, new boxes go into the right column or below the discussions; inner pages stay one column unless they have a menu of their own, which goes in the right column.
