# Personal home page: the fan page, 1998

## Summary

A fan page as people built them for themselves on free web space between 1996 and 2001: one centred column set straight on a tiled starfield, a huge coloured name at the top, a rainbow bar and a thick-bordered table of sections under it. Text is bright on black in the browser's own serif (white body text, yellow and lime headings, cyan links, pink visited links), every box has a 4px or 6px ridged border, and the page carries the furniture of its kind: a hit counter of boxed digits, "NEW!" marks, an "under construction" strip, awards, a web ring box and a guestbook. It is loud, centred and hand-made, and it was common until page builders and weblogs replaced it after 2001.

## Layout

- Fixed, not fluid: `--size-page` is 700px, centred in an 800px window by `.ds-page__frame`. The frame has no fill and no border: everything stands on `--fill-page`, the tile, as text stood on a background picture.
- One column, one axis. The header, welcome box, counters, section titles, pictures, boxes, awards and footer are centred; news, tables, forms and running text are left-aligned inside their boxes.
- The specimen is one example site, a fan page for a made-up television serial, of six views. `.ds-nav` and `.ds-footer` are written once outside the views.
- Home view from the top: `.ds-nav` (the `.ds-brand` at `--text-display`, an italic tagline, a `.ds-rule--rainbow`, then `.ds-nav__links`, a six-cell table of sections), then `.ds-page__main`: the `.ds-hero` welcome box, the `.ds-construction` strip, the `.ds-stat` counters, a `.ds-heading` over the `.ds-list` of dated news, a rainbow rule, a heading over the `.ds-grid` of four framed pictures, and `.ds-page__pair` with two `.ds-panel` boxes (web ring, poll).
- Inner pages keep the header and footer and drop the welcome box and counters. `.ds-page__inner` opens with `.ds-page__head`: the centred `.ds-breadcrumb`, the `.ds-page-header` strip (title and one italic line at the left, one button at the right) and, where the page has parts, the `.ds-tabs` row. Below it the page stays one column, with `.ds-rule--rainbow` between unrelated blocks. Only the page of running text uses `.ds-page__columns`: `.ds-prose` and a `--size-side` (190px) `.ds-sidebar` at the right.
- Views: `home`; `episodes` (tabs for the seasons, the episode table with its key of status badges, page numbers); `about` (running text, the sidebar, the line that shows link states); `gallery` (eight framed pictures and the empty state of a season with none); `links` (two boxes of bulleted links and the awards); `guestbook` (entries, notices, the form to sign, the delete message box).
- Spacing: `--space-1` 2px between table cells and digits; `--space-2` 4px cell and strip padding; `--space-3` 8px between a title and its block, between buttons and form rows, the height of the rainbow bar; `--space-4` 12px box padding and paragraph spacing; `--space-5` 16px between blocks and columns; `--space-6` 28px list indent, the gap between counters and the side padding of the welcome box.
- Fixed sizes: `--size-label` 170px form labels, `--size-field` 240px inputs, `--size-area` 340px textarea, `--size-check` 16px, `--size-dialog` 380px, `--size-icon` 40px, `--size-award` 124px, `--size-date` 84px date column of the news, `--size-thumb` 92px height of a framed picture.

## Typography and colour roles

These are the conventions of the whole era; every layout and token set of it follows them.

- The page and the sheet are of one lightness: `--color-page` and `--color-canvas` are both dark or both light, and `--fill-page` is a quiet tile built from palette tokens. `--color-text`, `--color-text-muted`, `--color-heading`, `--color-heading-alt`, every link colour, `--color-accent-alt`, `--color-danger`, `--color-success` and `--color-warning` are legible on the page, the canvas, all three surfaces and all four `--color-fill-*` blocks, at 2:1 or better.
- Fills that bring their own text: `--fill-bar` with `--color-bar-text` (the table of sections, table heads, the dialog title), `--fill-bar-alt` with `--color-bar-alt-text` (panel and sidebar title strips), `--fill-inverse` with `--color-inverse-text` (counter digits, awards, the current tab), `--fill-accent` with `--color-accent-text` (the current section, the "NEW!" badge, the current page number), `--color-notice` with `--color-notice-text` (notices and the two stripes of the construction strip), `--color-danger-surface` with `--color-danger`. A link on one of these takes the fill's text colour (`.ds-link--plain`).
- Status colours double as fills: a status badge and the destructive button are filled with `--color-success`, `--color-warning` or `--color-danger` and lettered in `--color-canvas`.
- `--color-surface-strong` is the ground that shows between the cells of a table; `--color-surface-alt` is alternate rows, inline code and the empty state.
- Headings are coloured: `--color-heading` (yellow here) for the name, page titles, h1 and h3; `--color-heading-alt` (lime) for section titles, h2, taglines, dates and form labels. `--color-accent-alt` (magenta) is the italic "Updated!" word, bullets and list markers.
- The rainbow bar is six hard steps: `--color-danger`, `--color-warning`, `--color-heading`, `--color-success`, `--color-link`, `--color-link-visited`.
- Type here: `--font-body` and `--font-heading` are Times New Roman at 16px, the unchanged browser default; headings 19, 24 and 32px bold, the name 48px. `--font-ui` Arial 13px sets buttons, controls, table heads, the section cells and badges; `--font-mono` Courier New sets counter digits, news dates and the textarea. Links are always underlined. Italic is written literally for taglines, lead lines and the footer note.
- Surface here: `--border-style` `ridge`, 4px for boxes and 6px (`--border-width-strong`) for the section table, the welcome box, tables, panels, pictures, awards and the dialog. No shadows, no radius, flat fills, no transitions. `--fill-page` is a starfield of 1px dots in the text, heading and link colours over `--color-page`.
- Written literally in `components.css`: `outset` and `inset` bevels on buttons, the brand tile and form controls (drawn in the control's own colour), `dotted` rules inside lists, the 45 degree turn of a bullet, and the slow two-step `.ds-blink`, which only runs inside `prefers-reduced-motion: no-preference`.

## Components

- `.ds-page`: on `<body>`; the tiled page and body type. `.ds-page__frame` the centred column; `.ds-page__main` and `.ds-page__inner` stack blocks `--space-5` apart; `.ds-page__head` stacks path, page header and tabs; `.ds-page__columns` text plus sidebar; `.ds-page__pair` two equal boxes; `.ds-section` a heading with its block.
- `.ds-nav`: the centred header. `.ds-nav__tagline` the italic line; `.ds-nav__links` the table of sections, six `.ds-nav__item` / `.ds-nav__link` cells in `--fill-bar` inside a `--border-width-strong` ridge; the cell of the view that shows takes `--fill-accent` (one selector per view, or `is-current`).
- `.ds-brand`: mark and name, centred. `.ds-brand__mark` is a tile dressed like the button (`--fill-button`, `--color-button-text`, outset bevel, `--radius-control`, `--shadow-control`) holding the mark as inline SVG with square caps; `.ds-brand__name` is the name at `--text-display` in `--color-heading`. Name and mark are placeholders for the installing project's own.
- `.ds-hero`: the welcome box on `--fill-panel` in a thick ridge with `--shadow-dialog`: `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__action` (a `.ds-button--large` and a secondary button), `.ds-hero__hint` (the "last updated" line).
- `.ds-page-header`: a ridged strip: `.ds-page-header__text` with `.ds-page-header__title` and the italic `.ds-page-header__lead`, and `.ds-page-header__action` at the right.
- `.ds-heading`: a centred section title in `--color-heading-alt`.
- `.ds-rule`: a carved rule; `.ds-rule--short` is 60% wide; `.ds-rule--rainbow` is the six-step bar, the usual divider.
- `.ds-construction`: the diagonal-striped strip with `.ds-construction__text` on it. `.ds-blink`: add to one small mark for the slow blink.
- `.ds-prose`: running text; styles `h1`, `h2`, `h3`, `p`, `ul`, `ol`, `strong`, `code`; `.ds-prose__lead`, `.ds-prose__code` for a block of markup.
- `.ds-link`: every link; states `is-visited`, `is-hover`, `is-active`, `is-focus`; `.ds-link--quiet` small, `.ds-link--strong` bold, `.ds-link--plain` on a coloured fill. `.ds-links` wraps a line of them. `.ds-icon` holds a small drawing before a link.
- `.ds-bullets`: a link list with coloured diamonds, `.ds-bullets__item` holding a `.ds-bullets__text`.
- `.ds-button`: a bevelled grey key; `.ds-button--secondary`, `.ds-button--danger`, `.ds-button--large`; states `is-hover`, `is-pressed`, `is-focus`, `is-disabled`. `.ds-buttons` lays out a row.
- `.ds-form`: two columns, right-aligned `.ds-form__label` and `.ds-form__field`; `.ds-form__input` (`--short`, `is-error`), `.ds-form__select`, `.ds-form__textarea`, `.ds-form__check`, `.ds-form__hint`, `.ds-form__error`, `.ds-form__actions`.
- `.ds-table`: a thick ridge round cells that stand `--space-1` apart on `--color-surface-strong`: `.ds-table__head` (`--num`), `.ds-table__cell` (`--num`, `--group`), `.ds-table__row--alt`, `.ds-table__desc`, `.ds-table__caption`.
- `.ds-stat`: the row of counters; each `.ds-stat__item` is a `.ds-stat__value` of boxed `.ds-stat__digit` spans in `--fill-inverse` over a `.ds-stat__label`.
- `.ds-list`: dated news in a ridged box: `.ds-list__item` with a `.ds-list__meta` date and a `.ds-list__body`; `.ds-list__more` under it. `.ds-list--entries` is the guestbook: `.ds-list__title`, meta, `.ds-list__text`.
- `.ds-panel`: a thick yellow-ridged box: `.ds-panel__title` strip, `.ds-panel__body`, `.ds-panel__text`, `.ds-panel__form`. `.ds-ring` with `.ds-ring__item` is the bracketed previous, random, next line of a web ring.
- `.ds-grid`: four framed pictures to a row: `.ds-grid__cell` (`--2`, `--3`, `--4` pick the fill) holds a `.ds-grid__pic` block with a line drawing, a `.ds-grid__title` link and a `.ds-grid__text` file size.
- `.ds-tabs`: a one-row table of `.ds-tabs__item` / `.ds-tabs__tab` cells, the current one (`is-current`) in `--fill-inverse`.
- `.ds-badge`: a solid "NEW!" label; `.ds-badge--text` an italic coloured word, `.ds-badge--count` a muted count, `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger` status labels; `.ds-badges` the key line.
- `.ds-awards` and `.ds-award`: prize pictures as bordered badges on `--fill-inverse`: `.ds-award__star`, `.ds-award__title`, `.ds-award__text`.
- `.ds-sidebar`: a column of `.ds-sidebar__block` boxes, each a `.ds-sidebar__title` strip over a `.ds-sidebar__list` of `.ds-sidebar__item` lines.
- `.ds-notice`: a centred line on `--color-notice`; `.ds-notice__label`; `.ds-notice--error`. `.ds-notices` stacks them.
- `.ds-pagination`: a centred line: `.ds-pagination__label`, `.ds-pagination__link` with `is-current` and `is-disabled`.
- `.ds-breadcrumb`: a small centred path, `.ds-breadcrumb__item` separated by a coloured mark, the last `is-current`.
- `.ds-dialog`: a message box in the page flow: `.ds-dialog__title` bar, `.ds-dialog__body`, `.ds-dialog__actions`. No backdrop.
- `.ds-empty`: `.ds-empty__title`, `.ds-empty__text` and a button on `--color-surface-alt`.
- `.ds-footer`: a rainbow bar, `.ds-footer__links` with `.ds-footer__item`, the `.ds-footer__mail` link, the italic `.ds-footer__note` ("best viewed at") and `.ds-footer__legal`.
- `.ds-view`: one screen of the example site; `.ds-view--home` shows when the address names none.

## Never

- `border-radius <= 0px`: nothing on these pages had a rounded corner.
- `box-shadow = none`: depth comes from ridged borders only.
- `text-shadow = none`: text is flat.
- `border-width <= 6px`: the thickest ridge is 6px.
- `font-size <= 48px`: the name is the largest text.
- `font-size >= 13px`: small print is 13px.
- `font-weight >= 400`: regular and bold only.
- `font-families <= 3`: Times, Arial and Courier.
- `uppercase-text <= 0%`: capitals are typed, never transformed.
- `content-width <= 700px`: one fixed 700px column.
- `transition = none`: nothing eases on hover.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new section is a centred `.ds-heading` over a ridged box, with a `.ds-rule--rainbow` before it; a new box is a `.ds-panel`; a new figure is a counter in `.ds-stat`. Keep everything on the centre axis, keep text bright on the tile, keep links underlined, and never add radius, soft shadows, hover transitions, a second navigation or a muted palette.
