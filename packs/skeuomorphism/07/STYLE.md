# Skeuomorphism app (recorder console, side rail)

## Summary

A small application drawn as a piece of audio equipment, from the years when screens imitated physical materials: one cabinet on a perforated page, a leatherette rail down its left side, and beside it a brushed faceplate with a display window, two knobs and a record key. It was a common dress for voice memo, radio, music and dictation tools on desktops, tablets and phones between about 2009 and 2013. The working parts are ordinary (a list, a table, a form) but each is built like hardware: counters sit in little windows, tabs are preset keys with a lamp, checkboxes are slide switches.

## Layout

- Fixed width: `--size-page` is 1040px, the width of the cabinet (`.ds-page__case`), centred on the page texture with `--space-6` above it. Nothing runs the full window width except the page texture.
- No bar across the top. The navigation is the rail, `.ds-nav`, `--size-rail` (210px) wide and as tall as the cabinet: the brand at its head, the navigation rows, one `.ds-sidebar` block (the reels) and, pushed to its foot, the tape gauge. To its right `.ds-page__main` (830px) holds the views and, at its foot, the status strip `.ds-footer`. The rail and the footer are written once and stay on every view.
- Order of the home view (`deck`): the `.ds-hero` faceplate across the main column, directly under it the `.ds-stats` counter strip, then one `.ds-section` split by `.ds-page__split` into a flexible column (title row and the `.ds-list` of latest takes) and a `--size-side` (300px) column (title row and the `.ds-grid` of four cassettes, two to a line). Both columns are one height: the list and the grid stretch.
- Inner pages keep the rail and drop the faceplate and counters. In their place: `.ds-page__bar`, a brushed strip with the breadcrumb at its left and a small `.ds-page__count` at its right; `.ds-page-header` (title and one line left, one or two keys right); where the screen has tabs, `.ds-page__tools`, a darker strip with the `.ds-tabs` preset keys at its left and a `.ds-page__note` at its right; then `.ds-section` parts, the second and later ones with `.ds-section--ruled`. A detail screen splits a section with `.ds-page__split`, running text on the left and stacked panels (`.ds-stack`) in the 300px column.
- Views. The specimen is a five-screen application, one `.ds-view` showing at a time. `deck` is the home view: faceplate, counters, latest takes, reels. `recordings` is the All tab: the table of takes with its state labels, pagination and the key to the labels. `recording` is one take: its trace in a display window, the notes as running text with links, and two panels of details and markers. `shared` is the Shared tab: an empty well, and the list of invitations sent. `settings` is the settings form with its notices and keys, and under it the alert that the destructive key opens. The rail leads to `deck`, `recordings`, `shared` and `settings`; `recording` is reached from any row and keeps Recordings lit.
- Spacing scale: `--space-1` 4px (inside badges, between rail rows), `--space-2` 8px (cell padding, icon gaps, between keys), `--space-3` 12px (row padding, panel padding, between cassettes), `--space-4` 16px (key side padding, display padding), `--space-5` 24px (section padding, column gap), `--space-6` 32px (above the cabinet, empty well), `--space-7` 48px (below the cabinet).
- Control sizes: `--size-pill` 20px (pagination keys, the switch), `--size-control` 28px (keys, fields, rail rows, tabs), `--size-control-large` 44px (the record key), `--size-knob` 52px, `--size-switch` 44px, `--size-meter` 8px (level meters and the tape gauge), `--size-wave` 96px. `--size-hair` (1px) is the offset of every highlight and groove.

## Typography and colour roles

- `--font-body` for running text, lists, tables and fields; `--font-heading` for the display title, page titles and section titles; `--font-ui` for keys, rail rows, tabs, table heads and form labels; `--font-mono` for everything a machine would print: counters, lengths, the status line of the display, numeric table cells and time scales.
- Sizes: `--text-base` for body, `--text-small` for meta, counts, badges and the status strip, `--text-ui` for controls, `--text-large` for the display's lead and the record key, `--text-display` for the title in the display window only, `--text-h1` for page titles, `--text-h2` for section titles and counter figures, `--text-h3` for the third heading level.
- `--fill-page` is the page; `--color-canvas` the cabinet's main column; `--color-surface` panels, rows and cassette labels; `--color-surface-alt` alternate rows and the foot of the alert; `--color-surface-strong` the tabs strip, the empty well and inline code.
- `--fill-inverse` with `--color-inverse-text` is the rail, and only the rail. `--fill-bar` with `--color-bar-text` is the faceplate, the table head and whatever is lit or current (the current rail row; pressed states use `--color-bar` darkened toward `--color-shadow`). `--fill-bar-alt` with `--color-bar-alt-text` is every quiet strip: the breadcrumb strip, the counter strip, panel titles (in `--color-heading-alt`), the tabs, the alert's title bar and the footer.
- A display window (the hero display, a counter figure, the trace of a take) is `--fill-input` with `--color-input-text`, a border darkened toward `--color-shadow` and `--shadow-control-pressed`: the same dress as a form field, so a palette's field colours are also its display colours.
- Cassettes use `--color-fill-1` to `--color-fill-4` with `--color-inverse-text`; their paper label is `--color-surface` with `--color-heading`. The dots beside the reels in the rail repeat the four fills.
- `--color-button` is the record key and every primary key; `--fill-button-secondary` the other keys, the knobs, the round play keys of a list row, the select and pagination. `--color-accent` is the badge and the lamp of the current tab; `--color-success`, `--color-warning` and `--color-danger` are state labels, the level meter and the destructive key. Text on the warning label is `--color-warning` darkened most of the way to `--color-shadow`, so that it reads in a light or a dark palette.
- One-off values built in `components.css` from tokens: the four screws of the faceplate (radial gradients of `--color-shadow` and `--color-bar-text`), the segmented level meter, the sheen of a cassette, the knob of the slide switch (a radial gradient of `--color-button-secondary`), `dotted` for the reels, `uppercase` and 1px tracking for the small titles of the rail and the knob labels, border colours darkened with `color-mix()` toward `--color-shadow`, and translucent mixes of a text token for its quieter form.

## Components

- `.ds-page`: on `<body>`; paints the page texture. `.ds-page__case` is the cabinet (`--radius-page`, `--shadow-dialog`), `.ds-page__main` its right column. `.ds-page__bar` is the brushed strip at the top of an inner screen with `.ds-page__count`; `.ds-page__tools` the strip that carries tabs, with `.ds-page__note`. `.ds-page__split` makes two columns of `.ds-page__col`; `.ds-page__col--narrow` is the 300px one. `.ds-section` is one padded part of the main column (`.ds-section--ruled` for a part that follows another), with `.ds-section__head` (a `.ds-section__title` and one link), and `.ds-section__foot` (pagination left, labels right). `.ds-stack` spaces stacked panels.
- `.ds-view`: one screen of the application, a `<section>` with an id; only the one named in the address shows, and `.ds-view--home` when none is named. Always a plain block.
- `.ds-nav`: the rail. `.ds-nav__items` holds `.ds-nav__item` rows, each a `.ds-nav__icon`, a word and at most one `.ds-nav__count`. The current row is a raised brushed key: `is-current`, or in the specimen one selector per view. `is-hover` lightens a row. `.ds-nav__foot` at the bottom holds the `.ds-nav__gauge` with its `.ds-nav__gauge-fill`.
- `.ds-brand`: the mark and name at the head of the rail. `.ds-brand__mark` is a 28px tile dressed as the primary key; `.ds-brand__name` the name in the heading font, always bold. Both are placeholders for the installing project's own name and logo.
- `.ds-hero`: the faceplate, once, at the top of the home view. `.ds-hero__display` is the display window with `.ds-hero__status` (a `.ds-hero__lamp` and one mono line), `.ds-hero__title`, `.ds-hero__lead` and two `.ds-hero__meter` rows (`.ds-hero__track`, `.ds-hero__level`, `.ds-hero__level--low`). `.ds-hero__controls` holds `.ds-hero__knobs` (each a `.ds-hero__dial` label with a `.ds-hero__knob`, turned by `--turned` or `--back`), `.ds-hero__actions` with the one `.ds-button--large`, and `.ds-hero__note`.
- `.ds-page-header`: the head of an inner screen in place of the faceplate: `.ds-page-header__copy` with `.ds-page-header__title` and one `.ds-page-header__text`; `.ds-page-header__actions` with at most two keys, the primary last.
- `.ds-prose`: running text: h1, h2, h3, paragraphs, lists, `code`, links.
- `.ds-link`: text link; states `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--quiet` for meta links. `.ds-links` is a wrapping row.
- `.ds-button`: a moulded key. `.ds-button--secondary`; `.ds-button--danger` for a destructive action, set apart with `.ds-buttons__end`; `.ds-button--large` for the record key only; `.ds-button--small`. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-button__icon` is a 16px glyph. `.ds-buttons` is a row.
- `.ds-form`: `.ds-form__row` with a `.ds-form__label` column and a `.ds-form__field`. `.ds-form__input` and `.ds-form__textarea` are sunk fields, `.ds-form__select` a secondary key. `is-error` with `.ds-form__error`; `.ds-form__hint`; `.ds-form__check` with `.ds-form__checkbox`, drawn as a slide switch that shows `--color-success` when on; `.ds-form__actions`.
- `.ds-table`: framed table. `.ds-table__head` cells are brushed; `.ds-table__row` with `is-first` and `is-alt`; `.ds-table__cell`, `.ds-table__cell--name`, and `is-numeric` for right-aligned mono figures.
- `.ds-list`: framed rows. `.ds-list__item` (`is-first`) holds a round `.ds-list__icon` key, `.ds-list__body` with `.ds-list__title` and `.ds-list__meta`, optionally a badge, and `.ds-list__time`.
- `.ds-panel`: framed plate with a `.ds-panel__title` strip and `.ds-panel__body`; `.ds-panel__text`, or `.ds-panel__rows` of `.ds-panel__row` (`is-first`) with `.ds-panel__key` and `.ds-panel__value`.
- `.ds-stat`: one counter: `.ds-stat__value`, a mono figure in a display window, beside a `.ds-stat__label`. Three to five sit in the `.ds-stats` strip directly under the faceplate; `is-first` on the first.
- `.ds-grid`: four `.ds-grid__cell` cassettes, two to a line (`--2`, `--3`, `--4` for the other fills). Each holds a `.ds-grid__label` (the paper label, a link), a `.ds-grid__window` with two `.ds-grid__reel` and one `.ds-grid__text` line.
- `.ds-tabs`: a joined row of preset keys. `.ds-tabs__tab` (`is-first`); the `is-current` key is held down and its lamp is lit in `--color-accent`.
- `.ds-badge`: a small lit plate for a word. `.ds-badge--count` is a number in a sunk window; `.ds-badge--success` (saved), `.ds-badge--warning` (copying, waiting) and `.ds-badge--danger` (clipped, bounced) are the state labels. `.ds-badges` is a row.
- `.ds-sidebar`: a block of the rail: `.ds-sidebar__title` over a `.ds-sidebar__list` of `.ds-sidebar__item`, each a `.ds-sidebar__dot` (`--2`, `--3`, `--4`), a `.ds-sidebar__link` and a `.ds-sidebar__meta` count.
- `.ds-notice`: a plate with a strong left edge; `.ds-notice__title` for the bold lead-in; `.ds-notice--error`. `.ds-notices` wraps several.
- `.ds-pagination`: small keys, `.ds-pagination__link`; `is-current` is held down, `is-disabled` is flat.
- `.ds-breadcrumb`: on the `.ds-page__bar` strip: `.ds-breadcrumb__link`, `.ds-breadcrumb__sep`, `.ds-breadcrumb__current`.
- `.ds-dialog`: an alert window: `.ds-dialog__title` bar with `.ds-dialog__close`, `.ds-dialog__body` with `.ds-dialog__heading` and `.ds-dialog__text`, `.ds-dialog__actions` with the confirming key last. `.ds-dialog__backdrop` is the dimmed area behind it.
- `.ds-empty`: a well sunk into the cabinet with `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one secondary key.
- `.ds-wave`: the trace of one take in a display window: `.ds-wave__trace` (an inline SVG painted with `currentColor`) over a `.ds-wave__scale` of times.
- `.ds-footer`: the status strip at the foot of the main column: `.ds-footer__note` left, `.ds-footer__links` of `.ds-footer__link` right.
- `.ds-switch`: glossy sliding ON/OFF switch over a checkbox (`.ds-switch__track`, inset, with a raised knob), set as rows of a `.ds-switches` list (`.ds-switches__row`, `.ds-switches__name`). Here: the Switches row of the settings form.
- `.ds-progress`: inset `<progress>` bar with a glossy fill and a `.ds-progress__label` line. Here: the Details panel of a recording.
- `.ds-tooltip`: `?` key with a dark glossy bubble and arrow (`.ds-tooltip__trigger`, `.ds-tooltip__bubble`) shown on hover or focus. Here: on the Switches label.
- `.ds-accordion`: `<details>` help list with glossy bar headings (`.ds-accordion__item`, `__head`, `__body`). Here: under the invitations of the shared view.

## Never

- `border-radius <= 10px`: only the cabinet is 10px; keys are 3px and plates 4px.
- `border-width <= 4px`: edges are hairlines; only the dotted rim of a reel is 4px.
- `box-shadow-blur <= 16px`: only the cabinet and the alert cast a wide shadow.
- `font-size <= 34px`: the title in the display window is the largest text.
- `font-size >= 11px`: meta text is the smallest.
- `font-weight >= 400`: no light weights.
- `font-families <= 3`: a body sans, a narrow sans for titles and keys, a monospace for figures.
- `line-height <= 1.55`: body text is set at 1.5.
- `underlined-links <= 10%`: links are coloured, not underlined, until hovered.
- `letter-spacing <= 1px`: titles and labels are tracked by one pixel at most.
- `transition = none`: states change at once.
- `animation = none`: nothing moves; meters and reels are drawn still.
- `content-width <= 1040px`: everything stays inside the cabinet.
- `row-gap <= 10px`: rows of a list or a table touch.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. Something raised (a key, a knob, a strip) takes a `--fill-*` gradient, a 1px border darkened toward `--color-shadow` and `--shadow-control`; something that shows a reading (a counter, a clock, a meter) is a display window in `--fill-input` and `--color-input-text` with `--shadow-control-pressed`, set in `--font-mono`. A new screen starts with `.ds-page__bar` and `.ds-page-header`; a new group of things in the rail is another `.ds-sidebar`. Do not add a bar across the top of the window, soft floating cards, borders thicker than a hairline on boxes, new font families or new colours.
