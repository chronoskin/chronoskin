# Personal home page: the shareware author's page, 1998

## Summary

The page of one person who writes small programs and gives them away to try, as such pages looked between 1996 and 2001: the whole site dressed as a program window, a grey framed sheet on a dithered grey desktop with a blue title banner, a toolbar of raised keys and a status bar at the foot. Text is black Arial on silver and white, headings are navy and maroon, links are the default blue and underlined, boxes have 2px grooved borders and the bars fade from dark to light like a title bar. Rows of programs with file sizes, a "why register" table, dated news and a printed order form are its furniture; it faded when download sites and on-line shops took the author's place.

## Layout

- Fixed, not fluid: `--size-page` is 760px. `.ds-page__frame` is one sheet in `--color-canvas` with a `--border-width-strong` border in `--color-border-strong` and `--shadow-dialog`, centred in an 800px window with `--space-4` above it. Everything, header and footer included, is inside the sheet and left-aligned.
- `.ds-nav` is written once at the top of the sheet: `.ds-nav__top`, a banner on `--fill-bar` with the `.ds-brand` (mark and name at `--text-display`) at the left and `.ds-nav__aside` (tagline and one line) right-aligned; under it `.ds-nav__links`, a toolbar of bevelled keys on `--color-surface-strong` with the date of the last update at the right end.
- The specimen is one example site, a shareware author's page, of four views.
- Home view, `.ds-page__main`: the `.ds-hero` (words and buttons at the left, `.ds-shot`, a drawing of the program's window `--size-shot` 230px wide, at the right), the `.ds-stat` strip of four figures, a `.ds-heading` over the `.ds-list` of program rows, then `.ds-page__pair` at three parts to two: a heading over the two by two `.ds-grid` and the `.ds-construction` strip at the left, a `.ds-panel` of dated news at the right.
- Inner pages keep the banner, toolbar and status bar and drop the hero and the figures. `.ds-page__inner` opens with `.ds-page__head`: the `.ds-breadcrumb` and the `.ds-page-header` (title and one italic line at the left, a button at the right, a thick line under). Below it: a `.ds-tabs` row sitting on the table it sorts, then page numbers and a section with the empty state; or `.ds-page__columns`, a `--size-side` (190px) `.ds-sidebar` of fact boxes at the left and `.ds-prose` at the right; or notices over a `.ds-page__pair` of the form and a column with a panel and the message box.
- Views: `home`; `programs` (tabs, the table of programs with its badge key, page numbers, the empty list of beta versions); `clock` (one program: sidebar, running text, link states); `register` (notices, the order form, where to send it, the confirm message box).
- Spacing: `--space-1` 2px between toolbar keys and inside strips; `--space-2` 4px cell and strip padding; `--space-3` 8px between a title and its block, between buttons, form rows and grid cells; `--space-4` 12px box padding and paragraph spacing; `--space-5` 16px sheet padding and the gap between blocks and columns; `--space-6` 24px list indent and the space under the sheet.
- Fixed sizes: `--size-label` 150px form labels, `--size-field` 240px inputs, `--size-area` 360px textarea, `--size-check` 16px, `--size-dialog` 380px, `--size-icon` 40px program icon, `--size-meta` 132px version column of a program row.

## Typography and colour roles

The conventions for which text sits on which fill are those of the era's first layout: text, heading, link and status colours are legible on page, canvas, surfaces and the four fills; bars, the inverse, the accent and notices bring their own text colour; status colours as fills are lettered in `--color-canvas`. This layout adds one: the name stands on `--fill-bar`, so the brand and tagline use `--color-bar-text`.

- `--font-body` and `--font-heading` are Arial: body 13px at line height 1.3, small print 11px, headings 14, 18 and 24px bold, the name 36px. `--font-ui` (MS Sans Serif, then Tahoma) at 11px sets toolbar keys, buttons, controls, table heads, tabs and badges. `--font-mono` Courier New sets versions, file sizes, numeric cells, news dates and the read-out. Links are always underlined.
- `--color-page` silver `#c0c0c0`, `--color-canvas` `#d4d0c8` (the sheet and the sunken cells of the status bar), `--color-surface` white (boxes, table cells), `--color-surface-alt` `#e8e8e8` (alternate rows, the figures strip, the empty state), `--color-surface-strong` `#a8a8a8` (the ground of the toolbar and the status bar).
- `--color-text` black; `--color-heading` navy `#000080` (titles, figures, program icons, h1 and h3); `--color-heading-alt` maroon `#800000` (section titles, lead lines, dates, form labels). Links `#0000ff`, visited `#800080`, hover red.
- `--fill-bar` navy with white text is the banner, table heads, the current tab and the dialog title. `--fill-bar-alt` teal with white text is every title strip. `--fill-inverse` black with green text is the read-out in the hero. `--fill-accent` yellow with black text is the pressed toolbar key, the "NEW!" badge and the current page number. `--color-fill-1` to `-4` (pale blue, cream, pale green, pale lilac) are the icon tiles and the grid cells and carry body text.
- The primary button and the toolbar keys are `--color-button` silver with black text; the secondary button is teal with white text.
- `--color-success`, `--color-warning` and `--color-danger` double as fills: a status badge and the destructive button are filled with one of them and lettered in `--color-canvas`. `--color-notice` with `--color-notice-text` is notices and the two stripes of the construction strip; `--color-danger-surface` with `--color-danger` is the error notice.
- Written literally in `components.css`: `outset` and `inset` bevels on buttons, the brand tile and form controls (drawn in the control's own colour), `dotted` rules inside lists and boxes, italic for taglines and lead lines, and the slow two-step `.ds-blink`, which only runs inside `prefers-reduced-motion: no-preference`.
- Surface: `--border-style` `groove`, 2px for boxes, cells and rules, 3px for the sheet, the table, the page-header line and the program window. Buttons carry a hard 1px shadow, the sheet and the program window 2px; nothing is blurred. `--fill-bar` and `--fill-bar-alt` fade from the bar colour at the left to a lighter mix at the right; other fills are flat. `--fill-page` is a 4px dither of `--color-border-muted` over `--color-page`. Radius 0, no text shadow, no transitions.

## Components

- `.ds-page`: on `<body>`; the tiled page and body type. `.ds-page__frame` the sheet; `.ds-page__main` and `.ds-page__inner` stack blocks `--space-5` apart inside it; `.ds-page__head` stacks path and page header; `.ds-page__columns` sidebar plus text; `.ds-page__pair` two blocks at three parts to two; `.ds-section` a heading with its block.
- `.ds-nav`: banner and toolbar. `.ds-nav__top` the banner, `.ds-nav__aside` with `.ds-nav__tagline` and `.ds-nav__note`; `.ds-nav__links` the toolbar of `.ds-nav__item` / `.ds-nav__link` keys, `.ds-nav__item--end` a note pushed to the right. The key of the view that shows is pressed in on `--fill-accent` (one selector per view, or `is-current`).
- `.ds-brand`: mark and name on the banner. `.ds-brand__mark` is a tile dressed like the button (`--fill-button`, `--color-button-text`, outset bevel, `--radius-control`, `--shadow-control`) holding the mark as inline SVG with square caps; `.ds-brand__name` is the name at `--text-display`. Name and mark are placeholders for the installing project's own.
- `.ds-hero`: the featured program on `--fill-panel`: `.ds-hero__text` with `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__action` (a `.ds-button--large` and a secondary button) and `.ds-hero__hint`; beside it `.ds-shot`, a small window: `.ds-shot__title` strip, `.ds-shot__body`, `.ds-shot__readout` on `--fill-inverse`, `.ds-shot__caption`.
- `.ds-page-header`: `.ds-page-header__text` with `.ds-page-header__title` and the italic `.ds-page-header__lead`, `.ds-page-header__action` at the right, a thick line under.
- `.ds-heading`: a left-aligned section title in `--color-heading-alt` over a full-width line.
- `.ds-rule`: a plain carved rule.
- `.ds-table`: every cell ruled, the borders collapsed: `.ds-table__head` (`--num`), `.ds-table__cell` (`--num` in the monospace, `--group` a closing row on `--color-fill-2`), `.ds-table__row--alt`, `.ds-table__desc`, `.ds-table__caption`.
- `.ds-stat`: the strip of figures; each `.ds-stat__item` is a `.ds-stat__value` over a `.ds-stat__label`, with a rule between items.
- `.ds-list`: one row for each program: `.ds-list__item` with a `.ds-list__icon` tile (`--2`, `--3`, `--4` pick the fill), a `.ds-list__body` holding the `.ds-list__title` link and `.ds-list__text`, and `.ds-list__meta` (version and size) over the `.ds-list__get` link. `.ds-list__more` under it.
- `.ds-panel`: a small window: `.ds-panel__title` strip on `--fill-bar-alt`, `.ds-panel__body`, `.ds-panel__text`. `.ds-news` with `.ds-news__item` and `.ds-news__date` is the dated news inside it.
- `.ds-grid`: two cells to a row: `.ds-grid__cell` (`--2`, `--3`, `--4` pick the fill) with a `.ds-grid__title` and a `.ds-grid__text`.
- `.ds-tabs`: property-sheet tabs on a base line, `.ds-tabs__item` / `.ds-tabs__tab`, the current one (`is-current`) taller and on `--fill-bar`.
- `.ds-sidebar`: a column of `.ds-sidebar__block` boxes, each a `.ds-sidebar__title` strip over a `.ds-sidebar__list` of `.ds-sidebar__item` lines.
- `.ds-footer`: the foot of the window on `--color-surface-strong`: `.ds-footer__links` with `.ds-footer__item`, then `.ds-footer__status`, three sunken `.ds-footer__cell` boxes.
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

- `border-radius <= 0px`: every box, key and field is square.
- `box-shadow-blur <= 0px`: the only shadows are hard 1px and 2px edges.
- `text-shadow = none`: text is flat.
- `border-width <= 3px`: the sheet's groove is the thickest line.
- `font-size <= 36px`: the name is the largest text.
- `font-size >= 11px`: small print is 11px.
- `font-weight >= 400`: regular and bold only.
- `font-families <= 3`: Arial, the system face of the keys and Courier.
- `uppercase-text <= 0%`: capitals are typed, never transformed.
- `content-width <= 760px`: everything is inside the 760px sheet.
- `transition = none`: nothing eases on hover.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new program is a row of `.ds-list` and a row of `.ds-table`; a new page gets a key in the toolbar; a new block on the sheet is a `.ds-heading` over a grooved box, or a `.ds-panel` when it needs a title strip; a new figure joins `.ds-stat`. Keep everything inside the one sheet and left-aligned, keep the banner and the status bar, keep links underlined, and never add radius, soft shadows, hover transitions, centred columns of text or a menu down the side.
