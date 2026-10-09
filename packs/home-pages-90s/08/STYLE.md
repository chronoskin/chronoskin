# Personal home page: the links directory and web ring hub, 1999

## Summary

A hobbyist's directory of other people's home pages and the web rings that join them, as volunteers kept them between 1996 and 2001: the name in big italic serif on a sky tile with small clouds, a row of folder tabs, and under them one white sheet holding a search band, a line of figures and the directory as two columns of categories with their sub-links. Text is small black Verdana on white with navy Times headings, default blue underlined links, pale yellow and pale blue blocks, thin dashed frames with a hard blue shadow and bars split into a lighter and a darker half. Lists of sites with two lines each, a table of rings, a link check, questions that open and an add-your-page form are its content; search engines and hosted rings replaced the kind after 2001.

## Layout

- Fixed, not fluid: `--size-page` is 744px, centred in an 800px window. `.ds-page__frame` holds the header on the tile and then `.ds-page__sheet`, one sheet in `--color-canvas` with a `--border-width-strong` border in `--color-border-strong`, `--shadow-dialog` and `--space-5` of padding. Views and footer are inside the sheet. Text is left-aligned; only the search band, the empty state and the footer are centred.
- `.ds-nav` is written once: `.ds-nav__top` on the tile (the `.ds-brand` at `--text-display` in italic at the left, `.ds-nav__aside` with a bold tagline and a line of small links right-aligned), then `.ds-nav__links`, folder tabs standing on the sheet's top edge. The tab of the view that shows is the sheet's colour, taller, and covers the edge. One tab is a `.ds-menu`: the categories drop under it.
- The specimen is one example site, a hobby links directory with web rings, of five views.
- Home view, `.ds-page__main`, is five rows `--space-5` apart: the `.ds-hero` search band; the `.ds-stat` line of four figures; a `.ds-heading` bar over the `.ds-grid`, the directory, two columns by three rows of categories; `.ds-page__pair` (a heading over the `.ds-list` of new sites at three parts, a `.ds-page__stack` of two `.ds-panel` boxes at two); the `.ds-construction` strip.
- Inner pages keep the header, tabs and footer and drop the search band, the figures and the directory. `.ds-page__inner` opens with `.ds-page__head`: the `.ds-breadcrumb`, the `.ds-page-header` (a tinted box, title and one line at the left, a button at the right) and, in a category, the `.ds-tabs` line of choices. Below it: `.ds-page__columns`, the list of sites with page numbers and a `--size-side` (180px) `.ds-sidebar` at the right; or the table over a `.ds-page__pair` of the empty state and a panel; or notices, the form and the message box in one column; or a `.ds-page__pair` of `.ds-prose` at three parts and the `.ds-accordion` at two.
- Views: `home`; `category` (tabs, five sites, page numbers, the sidebar); `rings` (the table of rings with its badge key, the empty list of your own rings, a progress bar); `add` (notices, the form, the message box for removing a listing); `help` (running text, link states, the questions).
- Spacing: `--space-1` 2px strip padding and the lines between categories; `--space-2` 4px cell padding, between tabs; `--space-3` 7px between a title and its block, between buttons, form rows and list items; `--space-4` 10px box padding, paragraph spacing, between stacked boxes; `--space-5` 15px sheet padding and the gap between blocks and columns; `--space-6` 22px list indent and the space under the sheet.
- Fixed sizes: `--size-menu` 180px drop-down list, `--size-search` 260px search field, `--size-label` 140px form labels, `--size-field` 250px inputs, `--size-area` 360px textarea, `--size-check` 14px, `--size-dialog` 360px, `--size-icon` 34px category tile, `--size-tip` 180px tooltip.

## Typography and colour roles

- The conventions of the era hold. `--color-page` and `--color-canvas` are of one lightness and every text, heading, link and status colour is legible on the page, the sheet, the surfaces and the four `--color-fill-*` blocks; the name and the tagline stand straight on the tile. Fills that bring their own text: `--fill-bar` with `--color-bar-text` (idle tabs, section bars, table heads, the dialog title), `--fill-bar-alt` with `--color-bar-alt-text` (the line of figures, panel title strips, a hovered tab), `--fill-inverse` with `--color-inverse-text` (the current choice of `.ds-tabs`, counter digits), `--fill-accent` with `--color-accent-text` ("NEW!", the current page number, the hovered drop-down line, the blocks of a progress bar), `--color-notice` with `--color-notice-text` (notices, tooltips, the construction strip), `--color-danger-surface` with `--color-danger`.
- The current folder tab is `--color-canvas` lettered in `--color-heading`: it is a piece of the sheet, not a bar.
- Status colours double as fills: status badges and the destructive button are filled with `--color-success`, `--color-warning` or `--color-danger` and lettered in `--color-canvas`.
- `--color-surface-alt` (pale yellow here) is the search band, the page header box, sidebar boxes, alternate rows, question titles, the ring line, inline code and the empty state. `--color-surface-strong` is the lines between categories and the track of a progress bar. The `--color-fill-*` blocks are the small tiles beside each category.
- Headings: `--color-heading` (navy) for the name, the search title, page titles, h1, h3 and the drawings on category tiles; `--color-heading-alt` (rust) for h2, the tagline, page header lines, sidebar titles, form labels, tooltip triggers and percentages. `--color-accent-alt` marks bullets, list markers and the italic "Cool!" word.
- Type here: `--font-body` and `--font-ui` are Verdana at 11px and 10px with a 1.5 line; `--font-heading` is Times New Roman bold at 15, 18 and 26px, and the name at 38px in italic. `--font-mono` sets the line under each site and the textarea. Every link is underlined, navigation included.
- Surface here: `--border-style` `dashed`, 1px for boxes and 2px (`--border-width-strong`) for the sheet, the tabs and the dialog. `--shadow-panel` and `--shadow-dialog` are hard offsets in a deeper sky blue. No radius, no soft shadow, no transitions. `--fill-page` is small white clouds on `--color-page`; bars are a flat upper half over a darker lower half.
- Written literally in `components.css`: `outset` and `inset` bevels on buttons, the brand tile, form controls and the progress track; `dotted` lines between list items; the italic of the name; the dotted underline of a tooltip trigger; the percentage widths of `.ds-progress__bar--NN`; `calc()` over `--border-width-strong` to seat the tabs on the sheet.

## Components

- `.ds-page`: on `<body>`; the tiled page and body type. `.ds-page__frame` the 744px column; `.ds-page__sheet` the sheet; `.ds-page__main` and `.ds-page__inner` stack blocks `--space-5` apart; `.ds-page__head` stacks path, page header and tabs; `.ds-page__pair` three parts to two; `.ds-page__columns` content plus sidebar; `.ds-page__stack` a column of boxes; `.ds-section` a heading with its block.
- `.ds-nav`: the header. `.ds-nav__top` holds the brand and `.ds-nav__aside` (`.ds-nav__tagline`, `.ds-nav__tools`); `.ds-nav__links` is the row of folder tabs, `.ds-nav__item` holding a `.ds-nav__link`; the tab of the view that shows joins the sheet (one selector per view, or `is-current`).
- `.ds-menu`: a drop-down under a tab. Put it on the `.ds-nav__item`, give the tab `.ds-menu__label` (it adds the small triangle) and follow it with `.ds-menu__list` of `.ds-menu__link` lines; the list shows on hover, on focus inside it, or with `is-open`.
- `.ds-brand`: mark and name. `.ds-brand__mark` is a tile dressed like the button (`--fill-button`, `--color-button-text`, outset bevel, `--radius-control`, `--shadow-control`); `.ds-brand__name` is the name at `--text-display` in `--color-heading`, italic. Name and mark are placeholders for the installing project's own.
- `.ds-hero`: the search band on `--color-surface-alt`: `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__action` (a field, a select, a `.ds-button--large` and a secondary button on one line), `.ds-hero__hint`.
- `.ds-page-header`: a tinted box: `.ds-page-header__text` with `.ds-page-header__title` and `.ds-page-header__lead`, `.ds-page-header__action` at the right.
- `.ds-heading`: a section title as a bar on `--fill-bar`.
- `.ds-rule`: a plain rule; `.ds-rule--short` is 60% wide; `.ds-rule--rainbow` is a bar of six colours. `.ds-construction` with `.ds-construction__text` is the striped strip.
- `.ds-prose`: running text; styles `h1`, `h2`, `h3`, `p`, `ul`, `ol`, `strong`, `code`; `.ds-prose__lead`, `.ds-prose__code`.
- `.ds-link`: every link; states `is-visited`, `is-hover`, `is-active`, `is-focus`; `.ds-link--quiet` small, `.ds-link--strong` bold, `.ds-link--plain` on a coloured fill. `.ds-links` wraps a line of them. `.ds-icon` holds a small drawing.
- `.ds-bullets`: a link list with small coloured squares, `.ds-bullets__item` holding a `.ds-bullets__text`.
- `.ds-button`: a bevelled key; `.ds-button--secondary`, `.ds-button--danger`, `.ds-button--large`; states `is-hover`, `is-pressed`, `is-focus`, `is-disabled`. `.ds-buttons` lays out a row.
- `.ds-form`: two columns, right-aligned `.ds-form__label` and `.ds-form__field`; `.ds-form__input` (`--short`, `is-error`), `.ds-form__select`, `.ds-form__textarea`, `.ds-form__check`, `.ds-form__hint`, `.ds-form__error`, `.ds-form__actions`.
- `.ds-table`: every cell ruled: `.ds-table__head` (`--num`), `.ds-table__cell` (`--num`, `--group`), `.ds-table__row--alt`, `.ds-table__desc`, `.ds-table__caption`.
- `.ds-stat`: the line of figures on `--fill-bar-alt`; each `.ds-stat__item` is a `.ds-stat__value` followed by its `.ds-stat__label`.
- `.ds-list`: sites, parted by dotted lines: `.ds-list__item` with a `.ds-list__title` link, a `.ds-list__text` of two lines and a `.ds-list__meta` line; `.ds-list__more` under it.
- `.ds-panel`: a framed box: `.ds-panel__title` strip, `.ds-panel__body`, `.ds-panel__text`. `.ds-ring` with `.ds-ring__item` is the bracketed previous, random, next line of a web ring.
- `.ds-grid`: the directory: `.ds-grid__cell` (`--2`, `--3`, `--4` pick the fill of its tile) holds a `.ds-grid__pic` tile and a `.ds-grid__body` with the `.ds-grid__title` category link and a `.ds-grid__text` line of sub-links.
- `.ds-accordion`: questions that open: each `.ds-accordion__item` is a `<details>` with a `.ds-accordion__title` summary and a `.ds-accordion__body`. Use it for help and rules only.
- `.ds-tabs`: a line of choices parted by bars: `.ds-tabs__label`, `.ds-tabs__item` holding a `.ds-tabs__tab`, the current one (`is-current`) on `--fill-inverse`.
- `.ds-badge`: a solid "NEW!" label; `.ds-badge--text` an italic coloured word, `.ds-badge--count` a muted count, `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger` status labels; `.ds-badges` the key line.
- `.ds-tooltip`: a word with a dotted underline; its `.ds-tooltip__tip` child is a small box on `--color-notice` that shows above it on hover, on focus (give the trigger `tabindex="0"`) or with `is-open`.
- `.ds-progress`: `.ds-progress__label` (with `.ds-progress__value` at its right) over a sunken `.ds-progress__track` holding a `.ds-progress__bar` of blocks in `--fill-accent`; its length is `.ds-progress__bar--10` to `--100` in tenths.
- `.ds-counter`: the visitor counter, `.ds-counter__digit` boxes on `--fill-inverse`.
- `.ds-sidebar`: a column of tinted `.ds-sidebar__block` boxes, each a ruled `.ds-sidebar__title` over a `.ds-sidebar__list` of `.ds-sidebar__item` lines.
- `.ds-notice`: a line on `--color-notice`; `.ds-notice__label`; `.ds-notice--error`. `.ds-notices` stacks them.
- `.ds-pagination`: a centred line: `.ds-pagination__label`, `.ds-pagination__link` with `is-current` and `is-disabled`.
- `.ds-breadcrumb`: a small path, `.ds-breadcrumb__item` separated by a coloured mark, the last `is-current`.
- `.ds-dialog`: a message box in the page flow: `.ds-dialog__title` bar, `.ds-dialog__body`, `.ds-dialog__actions`. No backdrop.
- `.ds-empty`: `.ds-empty__title`, `.ds-empty__text` and a button on `--color-surface-alt`.
- `.ds-footer`: inside the sheet under a rule, centred: `.ds-footer__links` with `.ds-footer__item`, `.ds-footer__count` with the counter, `.ds-footer__note` and `.ds-footer__legal`.
- `.ds-view`: one screen of the example site; `.ds-view--home` shows when the address names none.
- Left out as not of the period: `.ds-carousel`, `.ds-switch`, `.ds-avatar`.

## Never

- `border-radius <= 0px`: every box, tab and field is square.
- `box-shadow-blur <= 0px`: shadows are hard offsets.
- `text-shadow = none`: text is flat.
- `border-width <= 2px`: the sheet's edge is the thickest line.
- `font-size <= 38px`: the name is the largest text.
- `font-size >= 10px`: small print is 10px.
- `font-weight >= 400`: regular and bold only.
- `font-families <= 3`: Verdana, Times and one monospace.
- `uppercase-text <= 0%`: capitals are typed, never transformed.
- `content-width <= 744px`: everything is inside the 744px sheet.
- `transition = none`: nothing eases on hover.
- `animation = none`: nothing moves.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new section is a `.ds-heading` bar over its block; a new box is a `.ds-panel`; a new category is a `.ds-grid__cell`; a new kind of entry is a `.ds-list__item` with a title, two lines and a meta line. Keep everything inside the sheet, keep links blue and underlined, keep frames thin, and never add radius, soft shadows, hover transitions, a side menu or a dark band across the page.
