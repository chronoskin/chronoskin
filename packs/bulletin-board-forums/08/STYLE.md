# Board calendar, 2002 to 2007

## Summary

This is the calendar that came with board software in the first half of the 2000s and that busy communities used for chat nights, releases, contests and meetups: a month drawn as a ruled grid of day cells across the whole page, with the events as short links in their days. A dark masthead carries the name and a jump box, section tabs join the sheet under it, and below the month three equal columns list what is coming up, the calendars and the week's birthdays; an event has its own page with sign-up buttons, details and replies. The skin is the night one of gaming boards: navy cells on a darker sheet, glossy burnt orange bars, amber uppercase titles in Trebuchet over 12px Arial, square boxes and no motion.

## Layout

- The page is fluid. `--size-page` is `100%`; `.ds-page` has `--space-5` (10px) of padding, so the sheet `.ds-page__frame` takes about 98% of the viewport. Designed at 1024px. Never centre a narrow column.
- Order from the top, all inside the sheet: `.ds-nav__head`, the inverted masthead band with the brand and a tagline on the left and the month jump form on the right; `.ds-nav__bar`, a secondary bar on whose lower edge the section tabs stand, with the date line at its right; `.ds-page__body` with the views; the footer band. No text sits on the page colour.
- A view is a `.ds-view` section whose `.ds-view__body` stacks its blocks `--space-5` apart. The month view is full-width blocks, then `.ds-columns`: as many equal columns as fit at `--size-col` (230px), three at 1024px.
- The month grid `.ds-cal` is seven equal columns; a day cell is at least `--size-day` (82px) tall, the weekday head `--size-bar` (22px). The date tile of an event row is `--size-tile` (44px) wide. Table count columns are `--size-count` (60px), its date column `--size-date` (120px), rows at least `--size-row` (28px).
- Spacing scale: `--space-1` 1px, `--space-2` 2px, `--space-3` 4px, `--space-4` 7px, `--space-5` 10px, `--space-6` 14px.
  - `--space-1`: the gap between cells and rows inside any bordered block, and between day cells.
  - `--space-2`: vertical padding of bars, strips and buttons; gap between events in a day and between tabs.
  - `--space-3`: vertical padding of cells and tabs, padding inside a day cell, gap between inline items.
  - `--space-4`: horizontal padding of cells and strips, padding of panel bodies, gap between buttons.
  - `--space-5`: page and body padding and the gap between neighbouring blocks and columns. This is the only block gap.
  - `--space-6`: side padding of the masthead, horizontal padding of tabs, list indents, padding of dialog and empty-state bodies.
- Inner pages keep the masthead, the tab bar and the footer and drop the month heading. From the top: `.ds-breadcrumb`, then `.ds-page-header` in place of the hero (title and one line on the left, the action or the sign-up buttons on the right, no box), then the working blocks at full width: the calendar cells `.ds-grid`, `.ds-tabs` and `.ds-table` on the event list; `.ds-form` on the add page. An event page uses `.ds-layout`: text and replies in `.ds-layout__main` (at least `--size-main`, 440px), details and attendance in `.ds-layout__rail` (`--size-rail`, 230px) on the right. The member's calendar uses `.ds-columns` again.
- Views. The specimen is a five-screen example calendar; the tab bar has one tab per view and marks the current one. `month` (home) is the month: heading with the previous and next links and the add button, notice, the month grid with its legend, then coming up, the calendars and birthdays in three columns, and the figures strip. `events` is every event as a list: the four calendars as cells, filter tabs, the event table in two groups, count line and page numbers. `event` is one event: sign-up buttons, the withdraw confirmation that the Withdraw button opens, the description as long text and replies beside the details and attendance. `add` is the event form as it came back with an error. `mine` is the member's calendar: sign-ups, an empty reminders box and the notification notes.
- Form rows are a `--size-label` (180px) label cell and a field cell that takes the rest; text inputs are `--size-field` (260px) unless `--wide` or `--auto`; the year box of the jump form is `--size-year` (56px). A dialog is `--size-dialog` (52%) wide, centred in the flow on a band of the overlay colour.

## Typography and colour roles

- Three families. `--font-body` (Arial) for running text; `--font-heading` (Trebuchet MS) for the wordmark, page titles and every title strip, set in uppercase by `--heading-transform`; `--font-ui` (Tahoma) for tabs, buttons, inputs, table and weekday heads and badges. `--font-mono` is for inline code only.
- Sizes: `--text-small` and `--text-ui` 11px for meta, day cells, strips, buttons and tabs; `--text-h2` and `--text-h3` 11px for the uppercase title strips; `--text-base` 12px for running text and day numbers; `--text-large` 14px for names in tables and cells, figures and the day of a date tile; `--text-h1` 17px for the month and page titles; `--text-display` 20px for the wordmark only.
- `--line-body` is 1.35, `--line-heading` 1.25. Headings are bold; nothing is tracked. Only headings are uppercase; navigation and buttons are not (`--ui-transform` none).
- Links have no underline until hovered. Links in text are `--color-link`; event names, titles, breadcrumb links and member names are `--color-link-quiet`. Links in a notice and in the footer take the colour of their block and are always underlined. Hover turns a link `--color-link-hover`.
- Surfaces: `--color-page` outside the sheet, `--color-canvas` for the sheet. `--color-surface` is the main cell and a weekday; `--color-surface-alt` a weekend day, count and date cells, form labels and figures; `--color-surface-strong` group rows, date tiles, panel and column titles and the form action row. `--color-fill-3` is a day of another month and the alternate row; `--color-fill-4` is today, outlined in `--color-accent`. `--fill-panel` shows in the 1px gaps between cells.
- Bars: `--fill-bar` with `--color-bar-text` for the weekday row, table heads, list heads, form and dialog titles and the current tab of `.ds-tabs`. `--fill-bar-alt` with `--color-bar-alt-text` for the tab bar, idle tabs, panel sub-strips, comment heads and quiet badges. `--fill-inverse` with `--color-inverse-text` for the masthead and the footer; the wordmark stands on it, so it is `--color-inverse-text`.
- The current section tab is a `--color-canvas` cell with `--color-heading` text and a `--color-border-strong` outline on three sides, so it joins the sheet.
- Event kinds are marked by a 2px left edge or a small square `.ds-swatch`: `--color-bar-alt` for meetups, `--color-success` for releases, `--color-warning` for contests, `--color-danger` for upkeep.
- Borders: blocks have a 1px `--color-border` outline; `--color-border-muted` divides rows inside a column box and outlines the breadcrumb; `--color-border-strong` outlines buttons and, at `--border-width-strong` (2px), rules the tab row and the dialog.
- Markers: `--color-accent` for the badge and today's outline; `--color-accent-alt` for "new" remarks; `--color-heading-alt` for column titles. Status badges are a `color-mix()` of 30% of the status colour over `--color-surface` with `--color-text` and an outline in the status colour. The destructive button is `--color-danger` on `--color-danger-surface`.
- `--color-fill-1` to `--color-fill-4` tint the four calendar cells; in this palette they are four steps of navy.
- One-off values outside the token set: the status badge tints above; the notice edge is an equal `color-mix()` of `--color-accent` and `--color-notice`; the left edge of an event in a day cell is always `solid`, whatever `--border-style` is, so a double or ridged style does not blur a 2px mark; always-underlined links write `underline` directly.
- The dark reference set Arial before Verdana, and the reference stacks name Trebuchet MS and Tahoma; all are system fonts and no substitution was needed.

## Components

- `ds-progress`: a quota or poll bar, a native `<progress>` with a label above (private-message quota).
- `ds-accordion`: a collapsible `<details>` entry for a forum jump or an FAQ answer.
- `ds-avatar`: the square picture of a member, placed under or beside the name.
- `.ds-page`: on `<body>`. Wrap everything in `.ds-page__frame`; `.ds-page__body` holds the views. Each view is a `.ds-view` section (`.ds-view--home` for the month) and its blocks go in `.ds-view__body`. `.ds-columns` is the row of equal columns; `.ds-layout` with `.ds-layout__main` and `.ds-layout__rail` the two-column area. `.ds-page__legal` is an optional line of small print inside the sheet. `.ds-icon` sizes an inline SVG icon, `.ds-sprite` hides a symbol sheet, `.ds-block` puts a span on its own line, `.ds-scroll` lets a wide table scroll inside its own box.
- `.ds-nav`: the site header. `.ds-nav__head` is the masthead with `.ds-nav__brand` (the `.ds-brand` and a `.ds-nav__tagline`) and the `.ds-nav__jump` form, whose year box also takes `.ds-nav__year`. `.ds-nav__bar` holds `.ds-nav__menu`, the row of `.ds-nav__link` tabs, and the `.ds-nav__today` line; the current tab takes `is-current`, and in the specimen one `:has()` rule per view applies the same dress.
- `.ds-brand`: the site's mark and name at the left of the masthead; one per page. `.ds-brand__mark` is a square tile dressed like the primary button (`--fill-button`, `--color-button-text`, a `--color-border-strong` outline, `--radius-control`, `--shadow-control`), 1.4 times `--text-display`. `.ds-brand__name` is the name in `--font-heading` at `--text-display`. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo: replace both, keep the tile.
- `.ds-breadcrumb`: one bordered line at the top of an inner page, with `.ds-breadcrumb__link`, `.ds-breadcrumb__sep` and `.ds-breadcrumb__current`.
- `.ds-hero`: the month heading, one bordered cell: `.ds-hero__text` with the month as `.ds-hero__title` and a `.ds-hero__lead` count line, and `.ds-hero__action` with the previous and next links (`.ds-hero__step`) and the add button. Never a banner.
- `.ds-cal`: the month grid. Seven `.ds-cal__head` weekday cells, then one `.ds-cal__day` per day (`--weekend`, `--other` for a day of another month, `is-today`), each with a `.ds-cal__num` and its events as `.ds-cal__event` links (`--release`, `--contest`, `--upkeep`). An event name that does not fit is cut with an ellipsis; never let a day cell grow sideways.
- `.ds-swatch`: the small square that names an event kind (`--release`, `--contest`, `--upkeep`). `.ds-legend` with `.ds-legend__item` explains the kinds under the month.
- `.ds-page-header`: the head of an inner page. `.ds-page-header__text` holds `.ds-page-header__title` and one `.ds-page-header__desc` line; `.ds-page-header__action` or a `.ds-buttons` group sits at the right.
- `.ds-prose`: long text such as an event description: h1, h2, h3, p, ul, strong, code, hr.
- `.ds-link`: every link in text. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--quiet` for names and tools, `.ds-link--strong` for bold.
- `.ds-button`: the primary action, glossy orange with white bold text. `.ds-button--secondary` is the plain navy button. `.ds-button--danger` is for a destructive action such as withdrawing or deleting an event; use it once per confirmation, beside a secondary button. `is-hover`, `is-active`, `is-focus`, `is-disabled` (or `disabled`). Group buttons in `.ds-buttons`.
- `.ds-form`: a ruled block of rows. `.ds-form__title`, `.ds-form__row` with `.ds-form__label` (plus `.ds-form__hint`) and `.ds-form__field`; controls take `.ds-form__input` (`--wide`, `--auto`, `is-invalid`, `is-disabled`), errors `.ds-form__error`, checkboxes `.ds-form__check` and `.ds-form__checkbox`, buttons in `.ds-form__actions`.
- `.ds-table`: the event list and any data table. `th` is the primary bar (`.ds-table__head--left` for text columns); `.ds-table__group` is a group row; `.ds-table__row--alt` the alternate row; cells `.ds-table__num` and `.ds-table__last` (the date); inside the main cell `.ds-table__title`, `.ds-table__desc`, `.ds-table__meta`.
- `.ds-list`: event rows outside a table. `.ds-list__head`, then `.ds-list__item` (`is-alt` for the alternate row) made of a `.ds-list__date` tile with the day in `.ds-list__day` and a `.ds-list__text` cell with `.ds-list__title` and `.ds-list__meta`.
- `.ds-panel`: any titled box. `.ds-panel__title` (add `--bar` for a main block), `.ds-panel__sub`, `.ds-panel__body` (`--alt`), `.ds-panel__line`. `.ds-spec` is the label and value list of an event inside a panel: `.ds-spec__label`, `.ds-spec__value`.
- `.ds-stat`: summary figures as one ruled strip of equal cells. `.ds-stat__title` is optional; each `.ds-stat__item` has a `.ds-stat__value` over a `.ds-stat__label`. Use three to six.
- `.ds-grid`: the calendars as cells with 1px gaps in one bordered block. `.ds-grid__title`, `.ds-grid__cell` (`--2`, `--3`, `--4` for the other fills) with a swatch, `.ds-grid__name`, `.ds-grid__meta` and `.ds-grid__count`.
- `.ds-tabs`: a row of `.ds-tabs__tab` items on a 2px rule; `is-current` takes the primary bar fill. Use for the filters of one list.
- `.ds-badge`: a small square label in the accent fill. `.ds-badge--quiet` is the secondary bar. `.ds-badge--success` (open, going), `.ds-badge--warning` (few places, maybe) and `.ds-badge--danger` (closed, cancelled) are tints with an outline. `.ds-new` is bold text for a "new" remark.
- `.ds-sidebar`: a column of link boxes. Each `.ds-sidebar__block` has a `.ds-sidebar__title`, a `.ds-sidebar__list` of ruled `.ds-sidebar__item` rows with a `.ds-sidebar__count` figure, and an optional `.ds-sidebar__foot` with a `.ds-sidebar__footlink`.
- `.ds-notice`: a tinted line for calendar news; `.ds-notice--error` is the error box with a `.ds-notice__title`. `.ds-notice__link` is a link inside either.
- `.ds-pagination`: `.ds-pagination__label`, then small bordered `.ds-pagination__link` boxes; `is-current` for the current page. `.ds-toolbar` puts a count line on the left and the page numbers on the right.
- `.ds-comment`: replies under an event, one ruled block: a `.ds-comment__head` strip per reply, then `.ds-comment__body` (`--alt`) with `.ds-comment__text` and an optional `.ds-comment__sig` signature line.
- `.ds-dialog`: a confirmation box in the page flow with a 2px outline: `.ds-dialog__title`, `.ds-dialog__body`, `.ds-dialog__text`, `.ds-dialog__actions`, on a `.ds-dialog__backdrop` band. It never floats over the page.
- `.ds-empty`: a ruled block with one centred message and one button for a list with nothing in it: `.ds-empty__title`, `.ds-empty__body`, `.ds-empty__text`.
- `.ds-footer`: the dark band closing the sheet: a centred `.ds-footer__row` of `.ds-footer__link` items with `.ds-footer__sep` between them, and one `.ds-footer__line` under it.

## Never

- `border-radius <= 0px`: every box is square in all references.
- `box-shadow-blur <= 0px`: the only shadows are hard one-pixel highlights inside buttons and inputs.
- `transition = none`: states switch at once.
- `animation = none`: nothing moves.
- `font-size <= 20px`: only the wordmark is 20px.
- `font-size >= 11px`: 11px is the smallest size of this type set.
- `font-weight <= 700`: bold is the heaviest weight.
- `font-families <= 4`: Arial, Trebuchet MS, Tahoma and a monospace for code.
- `border-width <= 2px`: 1px everywhere, 2px on emphasised edges and event marks.
- `letter-spacing <= 0px`: no tracking.
- `uppercase-text <= 15%`: only headings and title strips are uppercase.
- `underlined-links <= 20%`: links have no underline until hovered.
- `gradient-fills <= 25%`: the gloss is on bars, tabs, buttons and inputs only.
- `row-gap <= 1px`: rows touch or are separated by a 1px line.
- `block-gap <= 20px`: neighbouring blocks are 10px apart.
- `content-width >= 90%`: the month fills the viewport.

## Extending

Derive a new component from the nearest one in the specimen: a new titled box from `.ds-panel`, a new row layout from `.ds-list` or `.ds-table`, a week or day view from `.ds-cal`, a new figure list from `.ds-spec`. Build it as a 1px `--color-border` outline on `--fill-panel`, with cells in the three surface colours `--space-1` apart and a title in `--fill-bar` or `--color-surface-strong`. Keep text tokens on the fills named above (bar text on bars, heading text on the strong surface, quiet links on cells). Use tokens only: no raw colours or lengths, no new radius, shadow or transition, no font size outside the type tokens, and `--space-5` as the gap to the next block.
