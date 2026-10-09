# Personal home page: the club notice board, 1999

## Summary

The home page of a small hobby club with its own notice board, as volunteers built them between 1996 and 2001: a banner box with the club's name and a line of bracketed links, then one table of three columns, small boxes at the left and right and the welcome, the boards and the newest messages in the middle. Text is dark brown Georgia on parchment and cream, headings are maroon and forest green in an old-style book face, links are blue and underlined, and every box sits in a 3px sunken border on a speckled paper tile. A visitor counter, "seen this week", a poll, a web ring and a join form are its furniture; the kind gave way to hosted forums after 2001.

## Layout

- Fixed, not fluid: `--size-page` is 776px, centred in an 800px window. `.ds-page__frame` has no fill and no border: the boxes stand on `--fill-page`, the tile.
- `.ds-nav` is written once at the top: a banner box on `--fill-panel` in a `--border-width-strong` border, the `.ds-brand` centred in it at `--text-display`, the italic `.ds-nav__tagline`, a rule and `.ds-nav__links`, text links in square brackets.
- The specimen is one example site, a garden birdwatchers' club, of five views.
- Home view, `.ds-page__main`, is three columns `--space-4` apart: the `.ds-sidebar` (`--size-left` 150px: club pages, seen this week, new members, the `.ds-counter`), `.ds-page__column` (the centred `.ds-hero`, the `.ds-stat` row of three plaques, a `.ds-heading` over the two by two `.ds-grid` of boards, a heading over the `.ds-list` of newest messages) and `.ds-page__aside` (`--size-right` 180px: three or four `.ds-panel` boxes and the `.ds-construction` strip). The three columns end near each other.
- Inner pages keep the banner and the footer and drop the side columns, the hero and the plaques. `.ds-page__inner` is one column the full 776px wide and opens with `.ds-page__head`: the `.ds-breadcrumb`, the `.ds-page-header` box (title and one italic line at the left, a button at the right) and, on a board, the `.ds-tabs` row. Below it: the table of topics, page numbers and the empty state; or `.ds-posts`, the messages of one topic, each with its writer in a `--size-author` (132px) cell at the left; or notices and the form. Only the guide uses `.ds-page__columns`: `.ds-prose` and a `--size-side` (190px) `.ds-sidebar` at the right.
- Views: `home`; `board` (tabs, the topic table with its badge key, page numbers, the empty list of your own messages); `topic` (three messages, page numbers, the message box for removing one); `guide` (running text, the sidebar, link states); `join` (notices, the form).
- Spacing: `--space-1` 2px between table cells; `--space-2` 5px strip and cell padding, between tabs; `--space-3` 8px box padding, the gap between plaques, grid cells and buttons, under a heading; `--space-4` 12px between boxes and between the three columns; `--space-5` 18px between the blocks of an inner page, hero padding; `--space-6` 30px list indent.
- Fixed sizes: `--size-label` 160px form labels, `--size-field` 240px inputs, `--size-area` 360px textarea, `--size-check` 16px, `--size-dialog` 380px, `--size-icon` 44px board drawings.

## Typography and colour roles

The conventions for which text sits on which fill are those of the era's first layout: text, heading, link and status colours are legible on page, canvas, surfaces and the four fills; bars, the inverse, the accent and notices bring their own text colour; status colours as fills are lettered in `--color-canvas`.

- `--font-body` Georgia 15px at line height 1.4, small print 12px. `--font-heading` Book Antiqua (then Palatino) bold: 16, 20 and 26px, the name 42px with 2px tracking. `--font-ui` Tahoma 12px bold sets the bracketed links, buttons, controls, table heads, tabs and badges. `--font-mono` Courier sets the counter, code and the textarea. Links are underlined and lose the line under the pointer.
- `--color-page` tan `#e8d8b0`, `--color-canvas` cream `#f8f0dc`, `--color-surface` `#fffaeb` (boxes, cells), `--color-surface-alt` `#f0e4c4` (alternate rows, the page-header box, the writer's cell, the foot of a plaque, the empty state), `--color-surface-strong` `#dcc898` (the ground between table cells).
- `--color-text` dark brown `#33200c`; `--color-heading` maroon `#7a1f1f` (the name, titles, drawings, h1 and h3); `--color-heading-alt` forest green `#006633` (section titles, the tagline, lead lines, writers' names, form labels). Links `#0000cc`, visited `#663399`, hover rust. `--color-accent-alt` ochre is the italic "Free!" word, list markers and the dots of the divider.
- `--fill-bar` green with pale yellow text is the title strips of the left column, table heads and the dialog title. `--fill-bar-alt` maroon with pale gold text is the title strips of the right column and the current tab. `--fill-inverse` dark brown with gold text is the figure of a plaque and the counter digits. `--fill-accent` rust with white text is the current bracketed link, the "NEW!" badge and the current page number. `--color-fill-1` to `-4` (sage, peach, duck-egg, straw) are the four boards, the quote in a message and the picture in a panel, and carry body text and links.
- The primary button is sand `--color-button` with brown text; the secondary is olive with white text.
- `--color-success`, `--color-warning` and `--color-danger` double as fills: a status badge and the destructive button are filled with one of them and lettered in `--color-canvas`. `--color-notice` with `--color-notice-text` is notices and the two stripes of the construction strip; `--color-danger-surface` with `--color-danger` is the error notice.
- Written literally in `components.css`: `outset` and `inset` bevels on buttons, the brand tile and form controls (drawn in the control's own colour), `dotted` rules inside lists and boxes, italic for taglines and lead lines, and the slow two-step `.ds-blink`, which only runs inside `prefers-reduced-motion: no-preference`.
- Surface: `--border-style` `inset`, 3px for boxes and 6px for the banner and the dialog, so every box looks let into the paper. Buttons stand on a hard 2px edge of `--color-shadow` (`--shadow-control`); nothing else has a shadow and nothing is blurred. `--fill-bar` is ruled with fine vertical lines and `--fill-bar-alt` with horizontal ones, like book cloth; `--fill-inverse` is hatched; buttons are a two-stop top-lit fill. `--fill-page` is a 64px tile of dark and light specks over `--color-page`. Radius 0, no text shadow, no transitions.

## Components

- `.ds-page`: on `<body>`; the tiled page and body type. `.ds-page__frame` the centred 776px column; `.ds-page__main` the three columns of the front page, `.ds-page__column` the middle one and `.ds-page__aside` the right one; `.ds-page__inner` the single column of an inner page; `.ds-page__head` stacks path, page header and tabs; `.ds-page__columns` text plus sidebar; `.ds-section` a heading with its block.
- `.ds-nav`: the banner box. `.ds-nav__tagline` the italic line; `.ds-nav__links` the line of `.ds-nav__item` / `.ds-nav__link` text links, each between square brackets; the link of the view that shows takes `--fill-accent` (one selector per view, or `is-current`).
- `.ds-brand`: mark and name, centred in the banner. `.ds-brand__mark` is a tile dressed like the button (`--fill-button`, `--color-button-text`, outset bevel, `--radius-control`, `--shadow-control`) holding the mark as inline SVG with square caps; `.ds-brand__name` is the name at `--text-display` in `--color-heading`. Name and mark are placeholders for the installing project's own.
- `.ds-hero`: the centred welcome box on `--fill-panel`: `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__action` (a `.ds-button--large` and a secondary button), `.ds-hero__hint`.
- `.ds-page-header`: a box on `--color-surface-alt`: `.ds-page-header__text` with `.ds-page-header__title` and the italic `.ds-page-header__lead`, and `.ds-page-header__action` at the right.
- `.ds-heading`: a section title in `--color-heading-alt` between two rules.
- `.ds-rule`: a plain carved rule; `.ds-rule--dots` is a row of dots in `--color-accent-alt`, the usual divider.
- `.ds-table`: cells that stand `--space-1` apart on `--color-surface-strong`: `.ds-table__head` (`--num`), `.ds-table__cell` (`--num` centred, `--meta` small, `--group` a closing row on `--color-fill-4`), `.ds-table__row--alt`, `.ds-table__desc`, `.ds-table__caption`.
- `.ds-stat`: the row of plaques; each `.ds-stat__item` is a `.ds-stat__value` on `--fill-inverse` over a `.ds-stat__label`.
- `.ds-list`: the newest messages in a box: each `.ds-list__item` is a `.ds-list__title` link over a `.ds-list__meta` line of who, where and when. `.ds-list__more` under it.
- `.ds-posts` and `.ds-post`: the messages of a topic. `.ds-post__author` (with `.ds-post__name`) at the left on `--color-surface-alt`, `.ds-post__body` with `.ds-post__meta`, `.ds-post__text` and `.ds-post__quote` at the right.
- `.ds-panel`: a box of the right column: `.ds-panel__title` strip on `--fill-bar-alt`, `.ds-panel__body`, `.ds-panel__text`, `.ds-panel__pic` (a drawing on `--color-fill-3`), `.ds-panel__form` (a one-question form). `.ds-ring` is the previous, random, next line of a web ring.
- `.ds-grid`: two boards to a row: `.ds-grid__cell` (`--2`, `--3`, `--4` pick the fill) holds a `.ds-grid__pic` drawing and a `.ds-grid__body` with the `.ds-grid__title` link and a `.ds-grid__text` count.
- `.ds-tabs`: a row of small boxes, `.ds-tabs__item` / `.ds-tabs__tab`, the current one (`is-current`) on `--fill-bar-alt`.
- `.ds-sidebar`: a column of `.ds-sidebar__block` boxes, each a `.ds-sidebar__title` strip on `--fill-bar` over a `.ds-sidebar__list` of `.ds-sidebar__item` lines. `.ds-counter` with `.ds-counter__digits` is the visitor counter inside one.
- `.ds-footer`: a dotted rule, `.ds-footer__links` with `.ds-footer__item`, the italic `.ds-footer__note` and `.ds-footer__legal`, all centred on the tile.
- `.ds-prose`: running text; styles `h1`, `h2`, `h3`, `p`, `ul`, `ol`, `strong`, `code`; `.ds-prose__lead` the italic opening line, `.ds-prose__code` a block to copy.
- `.ds-link`: every link; states `is-visited`, `is-hover`, `is-active`, `is-focus`; `.ds-link--quiet` small, `.ds-link--strong` bold, `.ds-link--plain` on a coloured fill. `.ds-links` wraps a line of them. `.ds-icon` holds a small drawing before a link.
- `.ds-bullets`: a link list with small square marks, `.ds-bullets__item` holding a `.ds-bullets__text`.
- `.ds-button`: a bevelled key; `.ds-button--secondary`, `.ds-button--danger`, `.ds-button--large`; states `is-hover`, `is-pressed`, `is-focus`, `is-disabled`. `.ds-buttons` lays out a row.
- `.ds-form`: two columns, right-aligned `.ds-form__label` and `.ds-form__field`; `.ds-form__input` (`--short`, `is-error`), `.ds-form__select`, `.ds-form__textarea`, `.ds-form__check`, `.ds-form__hint`, `.ds-form__error`, `.ds-form__actions`.
- `.ds-badge`: a solid "NEW!" label on `--fill-accent`; `.ds-badge--text` an italic coloured word, `.ds-badge--count` a muted count, `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger` status labels; `.ds-badges` the key line under a table.
- `.ds-notice`: a line on `--color-notice`; `.ds-notice__label`; `.ds-notice--error`. `.ds-notices` stacks them.
- `.ds-pagination`: a centred line: `.ds-pagination__label`, `.ds-pagination__link` with `is-current` (on `--fill-accent`) and `is-disabled`.
- `.ds-breadcrumb`: a small path, `.ds-breadcrumb__item` separated by a coloured mark, the last `is-current`.
- `.ds-dialog`: a message box in the page flow: `.ds-dialog__title` bar on `--fill-bar`, `.ds-dialog__body`, `.ds-dialog__actions`. No backdrop.
- `.ds-empty`: `.ds-empty__title`, `.ds-empty__text` and a button on `--color-surface-alt`.
- `.ds-construction`: the diagonal-striped strip with `.ds-construction__text` on it. `.ds-blink`: add to one small mark for the slow blink.
- `.ds-view`: one screen of the example site; `.ds-view--home` shows when the address names none.

## Never

- `border-radius <= 0px`: every box is square; only the dots of the divider are shapes.
- `box-shadow-blur <= 0px`: the one shadow is the hard edge under a button.
- `text-shadow = none`: text is flat.
- `border-width <= 6px`: the banner's sunken border is the thickest line.
- `font-size <= 42px`: the name is the largest text.
- `font-size >= 12px`: small print is 12px.
- `font-weight >= 400`: regular and bold only.
- `font-families <= 4`: Georgia, the book face of the headings, Tahoma and Courier.
- `uppercase-text <= 0%`: capitals are typed, never transformed.
- `content-width <= 776px`: the three columns fit an 800px window.
- `transition = none`: nothing eases on hover.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new small box goes in the left column as a `.ds-sidebar__block` or in the right as a `.ds-panel`; a new board is a `.ds-grid__cell`; a new block in the middle column is a `.ds-heading` over a sunken box; a new figure joins `.ds-stat`. Keep the banner on top and the three columns on the front page, keep inner pages to one wide column, keep text dark on paper and links blue and underlined, and never add radius, blur, transitions, a sheet behind the columns or a menu bar.
