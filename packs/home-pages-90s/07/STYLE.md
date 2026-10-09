# Personal home page: the band page, 1999

## Summary

The home page of a small band, as one of its members built it on free web space between 1996 and 2001: the name in wide engraved capitals on a dark harlequin tile, a menu bar with a list that drops from one of its cells, a poster box beside the next date, a strip of figures and three equal columns of news, a vote and songs to fetch. Text is small pale Helvetica on deep violet with mint and pink capitals for headings, periwinkle underlined links, flat 2px and 4px frames in cyan with a black line inside them and bars that glow in the middle. Gig tables, a one-at-a-time photo viewer, a web ring and a mailing list form are its furniture; the kind moved to hosted band profiles after 2001.

## Layout

- Fixed, not fluid: `--size-page` is 760px, centred in an 800px window. `.ds-page__frame` has no fill and no border: every box stands on `--fill-page`, the tile. Text is left-aligned; nothing is centred but button rows inside message boxes and picture captions.
- `.ds-nav` is written once at the top: `.ds-nav__top` (the `.ds-brand` at `--text-display` at the left, `.ds-nav__aside` with an italic tagline and one small line right-aligned at the foot of the name) and under it `.ds-nav__links`, one bar of five equal cells in a `--border-width-strong` frame. One cell is a `.ds-menu`: its list drops under it on hover or focus.
- The specimen is one example site, a band's page, of five views.
- Home view, `.ds-page__main`, is four rows `--space-5` apart: `.ds-page__poster` (the `.ds-hero` at three parts, a `.ds-panel` with the next date at two, equal in height); the `.ds-stat` strip of four figures; `.ds-page__trio`, three equal columns that end together, each a `.ds-heading` over one box (the `.ds-list` of dated news, a vote drawn with `.ds-progress`, a panel of songs with a one-field form and the web ring line); the `.ds-construction` strip; a heading over the `.ds-grid` of four prints.
- Inner pages keep the header and footer and drop the poster, the figures and the three columns. `.ds-page__inner` opens with `.ds-page__head`: the `.ds-breadcrumb`, the `.ds-page-header` (title and one italic line at the left, a button at the right, a thick line under) and, on the dates page, the `.ds-tabs` row. Below it: the table, page numbers and a `.ds-page__pair` of two boxes; or `.ds-page__columns`, `.ds-prose` and a `--size-side` (190px) `.ds-sidebar` at the right; or `.ds-page__poster` again with the `.ds-carousel` at three parts and a panel at two, then the grid and the empty state; or notices, the form and the message box in one column.
- Views: `home`; `gigs` (tabs, the table of dates with its badge key, page numbers, two boxes with a tooltip and a progress bar); `band` (running text, the sidebar, link states); `photos` (the carousel, a list of other nights, four prints, the empty state of a night with no photos); `join` (notices, the form, the message box for leaving the list).
- Spacing: `--space-1` 2px between figures, digits and tabs; `--space-2` 5px cell and strip padding; `--space-3` 8px between a title and its block, between buttons and form rows; `--space-4` 12px box padding, paragraph spacing, between prints; `--space-5` 18px between blocks and columns, poster padding; `--space-6` 26px list indent and the space under the page.
- Fixed sizes: `--size-menu` 170px drop-down list, `--size-label` 150px form labels, `--size-field` 240px inputs, `--size-area` 340px textarea, `--size-check` 15px, `--size-dialog` 380px, `--size-icon` 38px, `--size-thumb` 84px height of a print, `--size-day` 72px day block, `--size-slide` 210px height of the carousel picture, `--size-tip` 170px tooltip.

## Typography and colour roles

- The conventions of the era hold. `--color-page` and `--color-canvas` are of one lightness and every text, heading, link and status colour is legible on the page, the surfaces and the four `--color-fill-*` blocks. Fills that bring their own text: `--fill-bar` with `--color-bar-text` (the menu bar, table heads, the dialog title), `--fill-bar-alt` with `--color-bar-alt-text` (panel and sidebar title strips, the hovered menu cell), `--fill-inverse` with `--color-inverse-text` (the day block, the current tab, counter digits), `--fill-accent` with `--color-accent-text` (the current menu cell, "NEW!", the current page number, the hovered drop-down line, the blocks of a progress bar, the current carousel square), `--color-notice` with `--color-notice-text` (notices, tooltips, the construction strip), `--color-danger-surface` with `--color-danger`.
- Status colours double as fills: status badges and the destructive button are filled with `--color-success`, `--color-warning` or `--color-danger` and lettered in `--color-canvas`.
- `--color-surface-strong` is the ground between the figures of the strip, the track of a progress bar and the idle carousel squares; `--color-surface-alt` is alternate rows, idle tabs, inline code, the picture well of a print, the carousel's foot and the empty state.
- Headings: `--color-heading` (mint here) for the name, the poster title, page titles, h1, h3, figures and the place of the next date; `--color-heading-alt` (pink) for section titles, h2, taglines, dates, form labels, tooltip triggers and percentages. `--color-accent-alt` (yellow) marks bullets, list markers, the path separator and the italic "Updated!" word.
- Type here: `--font-body` and `--font-ui` are Helvetica at 13px and 11px, `--font-heading` is Copperplate (a commercial face of the period; the stack falls back to Helvetica) set in capitals with 1px tracking at 15, 19 and 27px and the name at 40px with 2px. `--font-mono` Lucida Console sets dates, figures, percentages and the textarea. Buttons, menu cells, tabs and table heads are bold 11px capitals. Links are underlined and lose the line under the pointer; menu, tab and page-number links carry none.
- Surface here: `--border-style` `solid`, 2px for boxes and 4px (`--border-width-strong`) for the menu bar, the poster, panels, tables, the carousel, the dialog and the lines under headings. `--shadow-panel` is a hard black line inside the frame, `--shadow-text` a hard 2px black drop under capitals on the tile and on bars. No radius, no soft shadow, no transitions. `--fill-page` is a harlequin of 44px diamonds; bars are brightest in the middle.
- Written literally in `components.css`: `outset` and `inset` bevels on buttons, the brand tile, form controls, the picture well and the progress track; `dotted` lines between rows; the dotted underline of a tooltip trigger; the percentage widths of `.ds-progress__bar--NN`; `calc()` over `--size-icon` for the carousel drawing.

## Components

- `.ds-page`: on `<body>`; the tiled page and body type. `.ds-page__frame` the 760px column; `.ds-page__main` and `.ds-page__inner` stack blocks `--space-5` apart; `.ds-page__head` stacks path, page header and tabs; `.ds-page__poster` three parts to two; `.ds-page__trio` three equal columns; `.ds-page__columns` text plus sidebar; `.ds-page__pair` two equal boxes; `.ds-section` a heading with its block.
- `.ds-nav`: the header. `.ds-nav__top` holds the brand and `.ds-nav__aside` (`.ds-nav__tagline`, `.ds-nav__note`); `.ds-nav__links` is the menu bar of `.ds-nav__item` cells with `.ds-nav__link` in `--fill-bar`; the cell of the view that shows takes `--fill-accent` (one selector per view, or `is-current`).
- `.ds-menu`: a drop-down under a cell of the bar. Put it on the `.ds-nav__item`, give the cell's link `.ds-menu__label` (it adds the small triangle) and follow it with `.ds-menu__list` of `.ds-menu__link` lines; the list shows on hover, on focus inside it, or with `is-open`.
- `.ds-brand`: mark and name. `.ds-brand__mark` is a tile dressed like the button (`--fill-button`, `--color-button-text`, outset bevel, `--radius-control`, `--shadow-control`) holding the mark with square caps; `.ds-brand__name` is the name at `--text-display` in `--color-heading`. Name and mark are placeholders for the installing project's own.
- `.ds-hero`: the poster box on `--fill-panel` in a thick frame with `--shadow-dialog`: `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__action` (a `.ds-button--large` and a secondary button), `.ds-hero__hint`.
- `.ds-gig`: the next date inside a panel: `.ds-gig__date` (a block on `--fill-inverse` with `.ds-gig__day` and `.ds-gig__month`) beside `.ds-gig__text` with `.ds-gig__place`.
- `.ds-page-header`: `.ds-page-header__text` with `.ds-page-header__title` and the italic `.ds-page-header__lead`, `.ds-page-header__action` at the right, a thick line under.
- `.ds-heading`: a section title in capitals with a thick line under it.
- `.ds-rule`: a plain rule; `.ds-rule--short` is 60% wide; `.ds-rule--rainbow` is a bar of six colours. `.ds-construction` with `.ds-construction__text` is the striped strip.
- `.ds-prose`: running text; styles `h1`, `h2`, `h3`, `p`, `ul`, `ol`, `strong`, `code`; `.ds-prose__lead`, `.ds-prose__code`.
- `.ds-link`: every link; states `is-visited`, `is-hover`, `is-active`, `is-focus`; `.ds-link--quiet` small, `.ds-link--strong` bold, `.ds-link--plain` on a coloured fill. `.ds-links` wraps a line of them. `.ds-icon` holds a small drawing.
- `.ds-bullets`: a link list with small coloured squares, `.ds-bullets__item` holding a `.ds-bullets__text`.
- `.ds-button`: a bevelled key; `.ds-button--secondary`, `.ds-button--danger`, `.ds-button--large`; states `is-hover`, `is-pressed`, `is-focus`, `is-disabled`. `.ds-buttons` lays out a row.
- `.ds-form`: two columns, right-aligned `.ds-form__label` and `.ds-form__field`; `.ds-form__input` (`--short`, `is-error`), `.ds-form__select`, `.ds-form__textarea`, `.ds-form__check`, `.ds-form__hint`, `.ds-form__error`, `.ds-form__actions`.
- `.ds-table`: one thick frame, rows parted by dotted lines: `.ds-table__head` (`--num`), `.ds-table__cell` (`--num`, `--date`, `--group`), `.ds-table__row--alt`, `.ds-table__desc`, `.ds-table__caption`.
- `.ds-stat`: the strip of four figures; each `.ds-stat__item` is a `.ds-stat__value` over a `.ds-stat__label`.
- `.ds-list`: dated news in a framed box: `.ds-list__item` with a `.ds-list__meta` date over a `.ds-list__body`; `.ds-list__more` under it.
- `.ds-panel`: a thick-framed box: `.ds-panel__title` strip, `.ds-panel__body`, `.ds-panel__text`, `.ds-panel__form` (one field and a button on a line). `.ds-ring` with `.ds-ring__item` is the bracketed line of a web ring.
- `.ds-grid`: four prints to a row: `.ds-grid__cell` (`--2`, `--3`, `--4` pick the fill) is a coloured card holding a sunken `.ds-grid__pic`, a `.ds-grid__title` link and a `.ds-grid__text` file size.
- `.ds-carousel`: one picture at a time, as a slide show page was: `.ds-carousel__slide`, `.ds-carousel__caption`, and `.ds-carousel__nav` with two buttons, the `.ds-carousel__dots` squares (`.ds-carousel__dot`, the showing one `is-current`) and `.ds-carousel__count`. Static: the buttons are links to the next page.
- `.ds-tabs`: a row of `.ds-tabs__item` / `.ds-tabs__tab` cells standing on a line, the current one (`is-current`) in `--fill-inverse`.
- `.ds-badge`: a solid "NEW!" label; `.ds-badge--text` an italic coloured word, `.ds-badge--count` a muted count, `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger` status labels; `.ds-badges` the key line.
- `.ds-tooltip`: a word with a dotted underline; its `.ds-tooltip__tip` child is a small box on `--color-notice` that shows above it on hover, on focus (give the trigger `tabindex="0"`) or with `is-open`.
- `.ds-progress`: `.ds-progress__label` (with `.ds-progress__value` at its right) over a sunken `.ds-progress__track` holding a `.ds-progress__bar` of blocks in `--fill-accent`; its length is `.ds-progress__bar--10` to `--100` in tenths. Use it for votes, funds and anything counted towards a total.
- `.ds-counter`: the visitor counter, `.ds-counter__digit` boxes on `--fill-inverse`.
- `.ds-sidebar`: a column of `.ds-sidebar__block` boxes, each a `.ds-sidebar__title` strip over a `.ds-sidebar__list` of `.ds-sidebar__item` lines.
- `.ds-notice`: a line on `--color-notice`; `.ds-notice__label`; `.ds-notice--error`. `.ds-notices` stacks them.
- `.ds-pagination`: a centred line: `.ds-pagination__label`, `.ds-pagination__link` with `is-current` and `is-disabled`.
- `.ds-breadcrumb`: a small path, `.ds-breadcrumb__item` separated by a coloured mark, the last `is-current`.
- `.ds-dialog`: a message box in the page flow: `.ds-dialog__title` bar, `.ds-dialog__body`, `.ds-dialog__actions`. No backdrop.
- `.ds-empty`: `.ds-empty__title`, `.ds-empty__text` and a button on `--color-surface-alt`.
- `.ds-footer`: a thick line, then `.ds-footer__row` (`.ds-footer__links` with `.ds-footer__item` at the left, `.ds-footer__count` with the counter at the right), the italic `.ds-footer__note` and `.ds-footer__legal`.
- `.ds-view`: one screen of the example site; `.ds-view--home` shows when the address names none.
- Left out as not of the period: `.ds-accordion`, `.ds-switch`, `.ds-avatar`.

## Never

- `border-radius <= 0px`: every box, key and field is square.
- `box-shadow-blur <= 0px`: the lines inside frames are hard, never soft.
- `border-width <= 4px`: the thick frame is 4px.
- `font-size <= 40px`: the name is the largest text.
- `font-size >= 11px`: small print is 11px.
- `font-weight >= 400`: regular and bold only.
- `font-families <= 3`: Helvetica, the engraved face of the headings and one monospace.
- `content-width <= 760px`: everything fits a 760px column.
- `transition = none`: nothing eases on hover.
- `animation = none`: nothing moves.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new section is a `.ds-heading` over one framed box; a new box is a `.ds-panel`; a new figure joins the `.ds-stat` strip; anything counted towards a total is a `.ds-progress`. Keep text left-aligned on the tile, keep headings in tracked capitals with the hard drop under them, keep links underlined, and never add radius, soft shadows, hover transitions, a second menu bar or a pale page.
