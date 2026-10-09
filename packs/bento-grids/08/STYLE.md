# Line-up bento, 2024

## Summary

A festival site whose programme is a bento: three wide tracks on a short row unit, in which the poster cell, the headliners, the stages, the countdown and the ticket counter are rounded cells that grow downwards more than sideways. The ground is a deep wine red, cells are a shade lighter and lit from one corner, a few are filled flat with bottle green, plum or night blue, text is pale pink in a heavy narrow grotesque, and one lime carries links, progress and the closing call. The form was common for festivals, conferences, launches and year-in-review pages from about 2022 to 2026.

## Layout

- Fixed centred column: `--size-page` is 1160px, designed at a 1440px viewport, padded by `--space-5` (24px) at the sides. Nothing spans the window: no bar, no band, no rule.
- `.ds-nav` is a masthead that scrolls with the page. Its first line, `.ds-nav__top`, holds the `.ds-brand` at the left, the date and place in `.ds-nav__when` in the middle and one button at the right. Under it `.ds-nav__links` is one strip as wide as the column, filled with `--fill-bar-alt`, cut into five equal segments; the current view's segment is filled with `--color-surface-strong`, and the last segment opens a `.ds-menu`.
- `.ds-grid` has three equal tracks (about 375px) and rows of at least `--size-row` (148px). A cell takes one track and one row by default; `--w2` takes two tracks, `--full` all three, `--h2` two rows, `--h3` three. `.ds-grid--auto` lets rows follow their content.
- The home page is one grid. Rows one to three: the `.ds-hero` poster, two tracks by three rows, with a two-row act cell over a one-row countdown beside it. Rows four and five: three two-row cells (an act, the `.ds-carousel` of stages, the ticket counter). Row six: the `.ds-stats` row of four small tinted cells. Rows seven and eight: a two-row list, two one-row cells stacked in the middle, the two-row `--accent` cell. The footer is the last cell, inverted.
- Cells never touch: the gap is `--space-gap` (18px) in both directions, the padding inside a cell `--space-cell` (28px), in the hero and the footer `--space-7` (48px).
- Spacing scale: `--space-1` 4px (strip inset, pill padding), `--space-2` 8px (button gap, title to caption), `--space-3` 12px (control padding, table cell padding, list rows), `--space-4` 16px (stack gap inside a cell, form gap, cell head to body), `--space-5` 24px (page gutter, above the masthead, title to lead), `--space-6` 32px (page header padding, hero blocks), `--space-7` 48px (hero and footer padding, dialog stage), `--space-8` 72px, `--space-9` 104px (textarea height).
- Controls are `--size-control` (42px) tall, `--size-control-large` (54px) in the hero and on the accent cell, 32px with `--small`; an avatar and the brand mark are `--size-avatar` and `--size-icon` (44px), a large avatar 88px.
- Inner pages (the programme, one act, tickets) have no hero. A view starts with the `.ds-breadcrumb` on the bare page, then the `.ds-page-header` as a band of its own: a full-width cell with the title at `--text-h1`, one line of text, an optional large avatar before them and buttons or a badge at the right. Under it, optionally, a `.ds-toolbar` with the `.ds-tabs` pill at the left and badges at the right, then one `.ds-grid--auto` in the same three tracks: the working table or form takes `--w2` with a one-track cell beside it, everything else is one track.
- Views: `home` is the festival (poster, two headliners, countdown, stages, ticket counter, the four figures, more acts, getting there, reminders, the call for a weekend pass); `lineup` is one day's programme (tabs for the days, the status labels, the table with pagination, stages as a sidebar, announcements, an empty slot, the bookers); `artist` is one act (prose, set times in panels, how full the field is, neighbouring acts, links); `tickets` is the box office (three ticket cells, the order form with its buttons, extras as switches, notices, questions as an accordion, the dialog that empties the order). The strip marks Festival, Line-up, Headliner or Tickets.
- Below 960px the grid and the stats row drop to two tracks and the hero takes both. Below 640px everything is one column and rows follow their content, the date line of the masthead hides, display type steps down to `--text-h2`, a table hides its `.ds-table__extra` columns and scrolls inside `.ds-table__scroll`, and a tooltip opens as a full row under its line (from 960px down).
- On a phone (640px and below) the strip `.ds-nav__links` keeps its five segments on one row, with a smaller `--text-small` label and tighter padding, so the nav never wraps; the `.ds-nav__when` date line is not shown and the brand and the tickets button stay on the top row. The current segment stays filled with `--color-surface-strong`, and the list of the last segment's `.ds-menu--end` drop-down is at least 20 characters wide.

## Typography and colour roles

- Two sans families: `--font-heading` is Bricolage Grotesque (then Archivo and Arial Narrow), heavy and narrow, for titles, names, figures and prices; `--font-body` and `--font-ui` are Familjen Grotesk with Work Sans and system fallbacks. `--font-mono` is only for inline code. The references used commercial display grotesques; these are the open substitutes.
- Sizes: `--text-base` 17px at `--line-body` 1.45; `--text-small` 14.5px (captions, table, meta); `--text-ui` 15px (strip, buttons, form); `--text-large` 20px (hero lead); `--text-display` 72px at `--line-display` 0.96 with `--display-tracking` -2.4px (poster title, countdown); `--text-h1` 48px (page title, figures, prices, footer wordmark), `--text-h2` 34px (act names, the accent cell), `--text-h3` 22px (cell titles, brand name) at `--line-heading` 1.05 with `--heading-tracking` -0.6px. The smallest text is 12.4px, `calc(var(--text-small) * 0.857)`: badges, meta, hints.
- Weights: `--weight-body` 400, `--weight-ui` 600, `--weight-bold` 700, `--weight-heading` and `--weight-display` 800.
- `--color-page` `#780016` wine: the ground and the masthead. `--color-surface` `#8c1528`: cells, lit by `--fill-panel` (8% lighter at the top left corner). `--color-surface-alt` `#981f33`: panels, the dialog, menus, the current sidebar item and table row. `--color-surface-strong` `#a52c40`: the current segment and tab, neutral badges, avatars, progress tracks. `--color-bar-alt` `#650012`: the strip, the tab pill, the table head.
- `--color-text` `#f3d3f0` pale pink, `--color-text-muted` `#d9a7c6`, `--color-heading` `#fff0fb`. Text on a `--color-fill-*` cell uses the same tokens; the fills are dark enough for them.
- `--color-button` `#e9c0e9` pink with wine text is the primary button and the brand tile. `--color-accent` and `--color-link` `#d2e823` lime: links, icons, the dots of the masthead, progress fills, the accent badge and the closing cell, whose text is `--color-accent-text`. `--color-accent-alt` `#e8efd6` pale green marks "new", ticks and switched-on toggles. `--color-inverse` `#2b0008` is the footer cell, the button on the accent cell, tooltips and the knob of a switched-on toggle.
- `--color-fill-1` bottle green, `--color-fill-2` plum, `--color-fill-3` night blue, `--color-fill-4` deep wine: flat fills for act cells, stat cells and one ticket cell; the poster carries a quarter disc of fill 2 and a small disc of fill 3.
- Borders: `--color-border` (pink at 14%) at `--border-width` 1px on every cell, panel and badge; `--color-border-strong` (pink at 30%) on secondary buttons, the dialog, menus, avatars and the empty cell; `--color-border-muted` `#a02c40` for rows inside a cell.
- Links have no underline and gain one on hover.
- Radii: `--radius-page` 36px for cells, the page header and the footer, `--radius-panel` 22px for panels, tables, notices, menus and the dialog, `--radius-control` 16px for buttons, inputs and the strip's segments, `--radius-pill` for badges, tabs, switches, avatars, icon discs and the brand tile.
- Shadows: `--shadow-panel` is one soft drop tucked under the cell (24px blur, pulled in by 14px); `--shadow-control` is a 3px darker lip inside the bottom edge of a button, 4px on hover, turned to the top when pressed; `--shadow-dialog` a 28px drop. No text shadow.
- One-off mixes written in `components.css`: the button's outline, status badges and `--new` (the colour at 14% with a 30% outline), the danger button's outline and hover fill, notice outlines, the rule in the footer (20% of its text colour).
- Status: `--color-success` `#8ee29a`, `--color-warning` `#ffd166`, `--color-danger` `#ffb0a0` on `--color-danger-surface` `#4a000e`, as small text and tinted badges only; `--color-notice` night blue with `--color-notice-text`.
- `--transition` is a 0.2s fade of colour and shadow. Nothing moves except the accordion mark.

## Components

- `.ds-page` on `<body>`: wine ground, family and 17px text. `.ds-page__inner` is the 1160px column of a view.
- `.ds-nav`: the masthead. `.ds-nav__inner` holds `.ds-nav__top` (brand, `.ds-nav__when` with `.ds-nav__sep` dots, `.ds-nav__actions`) and the strip `.ds-nav__links` of `.ds-nav__link` (`is-current` by hand).
- `.ds-brand`: the site's mark and name at the left of the masthead. `.ds-brand__mark` is a 44px round tile dressed like the primary button (`--fill-button`, `--color-button-text`, outline, `--shadow-control`) holding the mark as inline SVG; `.ds-brand__name` is the name in the heading family at 22px. The name `chronoskin` and its mark are placeholders: the installing project puts its own name and logo here.
- `.ds-hero`: the poster cell, always with `.ds-grid__cell--w2` and `--h3`. Badges for date and place, `.ds-hero__title`, `.ds-hero__lead` at the top; `.ds-hero__actions` with two large buttons at the bottom.
- `.ds-grid`: the bento. `.ds-grid__cell` with spans `--w2`, `--full`, `--h2`, `--h3`; `--stack` (16px between children), `--between` (first child at the top, last at the bottom), `--flush` (no padding, clipped, for the dialog stage), fills `--fill-1` to `--fill-4`, `--accent`, `--bare`. Inside: `.ds-grid__head` (title left, badge or icon right), `.ds-grid__icon` with `.ds-grid__glyph`, `.ds-grid__title`, `.ds-grid__text`, `.ds-grid__visual` and `.ds-grid__actions` (pushed to the bottom). `.ds-grid--auto` is for cells that hold components.
- `.ds-act`: one artist as the whole content of a cell and a link to the act: `.ds-act__top` (a large `.ds-avatar` and a day badge), then `.ds-act__name` and `.ds-act__meta` at the bottom.
- `.ds-price`: a figure at 48px with `.ds-price__unit` saying what it buys; one per ticket cell.
- `.ds-page-header`: the head of an inner page as a full-width band. `.ds-page-header__copy` holds an optional large avatar and a block with `.ds-page-header__title` and `.ds-page-header__text`; `.ds-page-header__actions` sits at the right. Put a `.ds-breadcrumb` above it; never use it with a `.ds-hero`.
- `.ds-stat`: one figure: `.ds-stat__value` (48px, optional `.ds-stat__unit`) with a `.ds-stat__label` and an optional `.ds-stat__note` (`--success`, `--warning`, `--danger`). `.ds-stats` is the row, spanning the grid: four small cells, tinted with `.ds-stat--fill-1` to `--fill-4`. `.ds-stat--display` raises a single figure in a cell of its own to 72px.
- `.ds-prose`: long text in a `--w2` cell; styles h1 to h3, paragraphs, lists, `strong`, `code` and links.
- `.ds-link`: inline lime link with `is-visited`, `is-hover`, `is-active`, `is-focus`; `--quiet` for pink links. `.ds-links` lays several out in a row.
- `.ds-button`: pink primary with a darker lip. `--secondary` deep wine with a pale outline, `--danger` coral text on `--color-danger-surface` (never a solid block), `--inverse` for the accent cell, `--large`, `--small`; states `is-hover`, `is-active`, `is-focus`, `is-disabled`. Group with `.ds-buttons`.
- `.ds-badge`: 12px pill. `--accent` lime, `--new`, `--count`, optional `.ds-badge__dot`; status labels `--success` (on time), `--warning` (moved, going fast), `--danger` (cancelled, sold out). `.ds-chips` wraps several.
- `.ds-form`: two-column grid of `.ds-form__field` (`--full`) with `.ds-form__label`, `.ds-form__control` (`--area`; `is-invalid`, `is-focus`, `is-disabled`), `.ds-form__hint`, `.ds-form__error`, `.ds-form__check` with `.ds-form__checkbox`, and `.ds-form__actions`.
- `.ds-table`: a 22px-rounded box with a `--fill-bar-alt` head and faint row rules, inside `.ds-table__scroll`; `.ds-table__link` on the act's name, `.ds-table__num`, `.ds-table__muted`, `.ds-table__extra` (hidden on a phone), `tr.is-current` for the act now on.
- `.ds-list`: rule-separated `.ds-list__item` with `.ds-list__title`, `.ds-list__text`, `.ds-list__meta`, or holding a `.ds-person`; `.ds-list__row` puts a badge beside a title. `.ds-rows` with `.ds-rows__item` is the label-and-control list used for switches.
- `.ds-panel`: the small 22px card inside a cell, `.ds-panel__title`, `.ds-panel__body`, `.ds-panel__meta`; `.ds-panels` puts two side by side.
- `.ds-tabs`: a segmented pill of `.ds-tabs__tab`, `is-current` filled; one tab per day. `.ds-toolbar` puts it on one line with a row of badges.
- `.ds-sidebar`: a one-track cell with `.ds-sidebar__title` and a `.ds-sidebar__list` of `.ds-sidebar__link` (`is-current` filled): the stages of a day.
- `.ds-notice`: blue information box with `.ds-notice__title` and `.ds-notice__text`; `--error` for the dark red variant. `.ds-notices` stacks them.
- `.ds-pagination`: `.ds-pagination__link` items; `is-current` filled, `is-disabled` greyed.
- `.ds-breadcrumb`: 14.5px trail of `.ds-breadcrumb__item` parted by slashes, the last `is-current`; the first thing of an inner view.
- `.ds-dialog`: 480px box with `.ds-dialog__bar` (`.ds-dialog__title`, `.ds-dialog__close` with `.ds-dialog__glyph`), `.ds-dialog__body`, `.ds-dialog__actions`, shown on a `.ds-stage` filled with `--color-overlay` inside a `--flush` cell.
- `.ds-empty`: centred `.ds-empty__title`, `.ds-empty__text` and one secondary button inside a `--bare` cell.
- `.ds-checks`: a ticked list, `.ds-checks__item` with `.ds-checks__mark`, for what a ticket includes. `.ds-code` (`.ds-code__key`, `.ds-code__str`) and `.ds-chart` (`.ds-chart__bar`) are there for a cell that needs them.
- `.ds-menu`: a drop-down under its trigger, opened by hover or keyboard focus (`is-open` by hand): `.ds-menu__list` of `.ds-menu__item` with an optional `.ds-menu__hint`, parted by `.ds-menu__rule`; `.ds-menu__caret` in the trigger; `--end` aligns it to the right edge. Used under the strip's last segment.
- `.ds-accordion`: `<details>` as `.ds-accordion__item` with a `.ds-accordion__head` summary and `.ds-accordion__body`; the questions on the tickets page.
- `.ds-tooltip`: a focusable term with `.ds-tooltip__mark` and `.ds-tooltip__tip`, shown on hover or focus (`is-open` by hand) on `--fill-inverse`.
- `.ds-carousel`: `.ds-carousel__track` of `.ds-carousel__slide` with scroll snap, one showing; `.ds-carousel__foot` holds `.ds-carousel__dots` (`.ds-carousel__dot`, `is-current`) and `.ds-carousel__arrows` of `.ds-carousel__arrow` with a `.ds-carousel__glyph`. Used for the stages.
- `.ds-switch`: a label around `.ds-switch__input` (a checkbox) and `.ds-switch__track`; on is `--color-accent-alt` with a dark knob.
- `.ds-progress`: `.ds-progress__label` with `.ds-progress__value` over `.ds-progress__bar` with `.ds-progress__fill` (`--1` to `--5` set the width; `--success`, `--warning` the colour). Tickets sold, pitches taken.
- `.ds-avatar`: initials on a disc, `--accent`, `--large` (88px, in act cells and an act's page header); `.ds-avatars` overlaps several; `.ds-person` puts one beside `.ds-person__copy` with `.ds-person__name` and `.ds-person__meta`.
- `.ds-footer`: the last cell, inverted: `.ds-footer__inner` with `.ds-footer__about` (`.ds-footer__brand` as a 48px wordmark, `.ds-footer__text`), two columns of `.ds-footer__title` over a `.ds-footer__list` of `.ds-footer__link`, and `.ds-footer__strip`.

## Never

- `box-shadow-blur <= 28px`: one soft drop under a cell, 28px under the dialog; nothing glows.
- `border-width <= 2px`: hairlines are 1px.
- `border-radius >= 8px`: no box is square.
- `border-radius <= 36px`: the cell is the roundest box.
- `font-families <= 3`: a display grotesque, a text grotesque and the mono.
- `font-size >= 12px`: the smallest text is 12.4px.
- `font-size <= 72px`: the poster title is the largest text.
- `font-weight <= 800`: display type is extra bold, never black.
- `text-shadow = none`: no text shadows.
- `gradient-fills <= 30%`: the corner light of a cell is the only gradient.
- `underlined-links <= 5%`: links are told apart by colour; underline is the hover state.
- `uppercase-text <= 2%`: capitals only in avatar initials.
- `line-height <= 1.6`: body text sits at 1.45.
- `block-gap <= 48px`: cells sit 18px apart.
- `content-width <= 1160px`: content stays in the centred column.
- `animation = none`: nothing animates.

## Extending

Derive a new component from the nearest one in the specimen. A new item of the programme is a `.ds-grid__cell` that grows downwards (`--h2`) before it grows sideways, with one large name or figure and one caption; a new box inside a cell is a `.ds-panel`; a new label is a `.ds-badge`; a person is always a `.ds-avatar`, never a picture. Use tokens only: fills are `--fill-panel`, a `--color-surface*` or a flat `--color-fill-*`, outlines `--border-width` in `--color-border`, corners `--radius-page` for cells, `--radius-panel` inside them and `--radius-control` for controls, spacing from the `--space-*` scale. Keep one subject per cell, part cells with the gap and never with a rule, keep lime for what can be clicked or counted, and do not add a full-width band, a glow or a second display family.
