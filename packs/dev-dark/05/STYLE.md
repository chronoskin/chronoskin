# Dark developer tools status dashboard

## Summary

This is the status and observability page of a developer service from 2019 to 2026: a dark page of hairline boxes that answers one question first, whether everything is up, and then shows the evidence. A banner states the verdict, tiles carry a figure over a sparkline, each service has a bar for every one of the last 60 days, latency is three lines on a chart and incidents are dated rows with a state label. In this set the page is true black, every letter is monospace and headings are uppercase, corners are square, nothing has a soft shadow and the one accent is a hot pink.

## Layout

- Designed for a 1440px viewport, not fluid beyond it. `.ds-main` is the only column: at most `--size-page` (1280px) wide, padded by `--space-6` above and `--space-5` at the sides, stacking its children `--space-5` (20px) apart.
- Navigation is `.ds-nav`, a bar of two rows on `--fill-bar` that scrolls with the page. The first row, `.ds-nav__inner` (`--size-top` 52px), holds the brand, the mono `.ds-nav__scope` (organisation and environment after slashes) and at the right a state badge and two small buttons. The second row, `.ds-nav__links`, is the sections as joined cells with hairlines between them; the current cell has a 9% wash of the bar's text colour and a `--border-width-strong` line in `--color-accent` along its top. There is no rail and no sidebar on the home view.
- The home view stacks: `.ds-hero`, a banner with a ringed dot, the verdict, two buttons and `.ds-hero__facts` at the right; `.ds-stat`, four tiles; `.ds-cols` with the services panel (`.ds-uptime`) beside a `--size-side` (380px) column holding the latency chart and the inverted `.ds-callout`; a `.ds-colhead` and `.ds-grid`, six region cells in one row; a `.ds-colhead` and `.ds-list`, the recent incidents.
- Inner pages have no hero. `.ds-main` opens with `.ds-page-header`: breadcrumb, a `--text-h1` title (a state badge may follow it on the same line) and one muted line at the left of a `--border-width-strong` rule in `--color-accent-alt`, at most two buttons at the right. Tabs are joined cells in a `.ds-toolbar` directly over the panel that holds the table; the pagination is in the `.ds-table__foot`. `.ds-cols` puts a 380px column at the right, `.ds-cols--even` halves.
- Views. The specimen is the status page of a hosted query service in four views. `overview` is the home view described above. `incidents` has the tabs, the table of incidents with pagination, then an empty state beside the key of state labels with a row of links. `incident` is one incident: its updates as `.ds-timeline` and the postmortem as prose in the main column; affected components, facts, the action buttons and the open dialog that resolves it at the right. `alerts` has two notices, the form for a new rule and, beside it, the existing rules with their switches and the payload a webhook receives.
- Spacing scale: `--space-1` 4px (meter height, inside key caps), `--space-2` 8px (button and tab vertical padding, between adjacent buttons and badges, sidebar rows), `--space-3` 12px (panel head, table cell and uptime row vertical padding), `--space-4` 16px (panel, tile and cell padding, between region cells), `--space-5` 20px (between blocks, side padding, between tiles), `--space-6` 28px (banner padding, top of the column), `--space-7` 40px (sparkline height), `--space-8` 64px (above the footer).
- Other sizes: `--size-label` 208px (the name column of a service and the date column of an incident), `--size-bar` 28px (height of the daily bars), `--size-pulse` 44px (the ringed dot), `--size-mark` 22px, `--size-dot` 8px, `--size-dialog` 460px.
- Below 1080px the side column goes under, the facts move under the verdict, tiles stand two across and regions three. Below 640px everything is one column, the section cells scroll sideways, a service's bars go under its name, an incident's date goes over it, and a wide table scrolls inside `.ds-table__wrap`; the page itself never scrolls sideways.
- On a phone (640px and below) `.ds-nav__inner` wraps and the second line of `.ds-nav__scope` (the environment) is not shown; `.ds-nav__links` is a second row of cells that scrolls sideways with tighter padding. The `.ds-menu` keeps only its avatar (`.ds-menu__label` is not shown) and still lists every view, the current one with the accent line.

## Typography and colour roles

The conventions for which text sits on which fill are those of the era's first layout and are kept here.

- `--color-text`, `--color-heading` and `--color-text-muted` sit on the page, on `--fill-panel`, `--color-surface-alt` (a hovered row, the current sidebar row, the dialog) and `--color-surface-strong` (the current tab, key caps, quiet badges). `--color-bar-text` sits on `--fill-bar` (both rows of the bar). `--color-bar-alt-text` sits on `--fill-bar-alt` (table heads, the sidebar title, code heads, the footer), with `--color-heading` for the brand there. `--color-input-text` sits on `--fill-input` (form controls and code). `--color-inverse-text` sits on `--fill-inverse`, used once per view, for `.ds-callout`.
- `--color-fill-1` to `--color-fill-4` never carry text. They colour the sparkline of a tile and the lines of the chart (78% or 82% of the fill mixed into `--color-heading`), rule the top of a region cell at full strength and tint it at 13%; `--color-fill-1` is the glow in the corner of the banner.
- `--color-accent` marks the current section and fills the default badge through `--fill-accent`. `--color-accent-alt` is a line and text colour: the mono label of the banner, the rule of the page header, the line under the current tab, keywords in code.
- `--color-success`, `--color-warning` and `--color-danger` carry the meaning of the page and are used at full strength only for things without text on them: the daily bars, dots, meters, the knob of a switch. With text they are a 12% tint with a 32% hairline and the colour as text. The destructive button is a tint too.
- Borders: `--color-border` on boxes, bars and between cells, `--color-border-strong` on the banner, secondary buttons, grid lines of the chart and the dialog, `--color-border-muted` between rows inside a box. All `--border-width`; `--border-width-strong` is the line of the current section and tab, the top of a region cell and the rule of the page header.
- Type. `--font-mono` for every figure, date, id and label; body, headings and controls use their own tokens. `--text-base` 14px at `--line-body` 1.7; `--text-ui` 13px for navigation, buttons, forms, tables, rows and rules; `--text-small` 12px for badges, mono labels, facts and legends; `--text-large` 15px for the banner's lead. `--text-display` 44px is the verdict only; `--text-h1` 26px page titles and tile figures; `--text-h2` 18px section heads and region figures; `--text-h3` 13px panel and cell titles. Headings take `--heading-transform` and `--heading-tracking`.
- Uppercase comes from `--heading-transform` on headings and from mono labels (`.ds-stat__label`, table heads, sidebar titles, region codes), tracked by 6% of the font size. Links in running text and tables follow `--link-decoration`; quiet links are `--color-link-quiet` and turn `--color-heading`.
- Surface. `--radius-control` for buttons, inputs and the tabs, `--radius-panel` for panels, tiles, the list, notices and the callout, `--radius-page` for the banner, region cells and the dialog box, `--radius-pill` for badges, meters and switches. `--shadow-panel` is on panels, tiles, cells and the list, `--shadow-control` on buttons and the brand tile, `--shadow-dialog` on the banner and the dialog.
- One-off values written with `calc()` and `color-mix()`: the banner glow (20% of a fill), region tints (13%), line colours (78% and 82%), status tints (12%, 14%, 22%, 32%, 40%, 55%), the wash of the current section (9%), the notice border (22% and 40%), the corners of a daily bar (30% of `--radius-control`), the gap between bars (half of `--space-1`), mono label tracking. Navigation cells are at `opacity` 0.6 until current, daily bars at 0.86, callout text at 0.78. The empty state's outline is dashed.

## Components

- `.ds-page`: on `<body>`. Sets the font, text colour and `--fill-page`. `.ds-nav`, the views and `.ds-footer` are its direct children.
- `.ds-main`, `.ds-stack`, `.ds-cols`, `.ds-colhead`: the column of a view, a vertical stack 20px apart (`.ds-panel--grow` in it takes spare height), the column pairs (`--even`) and a section title with a note at its right (`.ds-colhead__title`).
- `.ds-brand`: the site's mark and name. `.ds-brand__mark` is a 22px tile dressed like the primary button holding the mark as inline SVG; `.ds-brand__name` is the name. Both are placeholders for the installing project's own name and logo.
- `.ds-nav`: the bar. `.ds-nav__inner` holds the brand, `.ds-nav__scope` and `.ds-nav__actions`; `.ds-nav__links` is the second row of `.ds-nav__link` cells (current: `is-current`). In the specimen the showing view marks its link with one `:has()` rule per view; a project sets `is-current`.
- `.ds-hero`: home view only, the state banner: `.ds-hero__pulse`, then `.ds-eyebrow`, `.ds-hero__title`, `.ds-hero__lead` and `.ds-hero__actions`, then `.ds-hero__facts` of `.ds-hero__fact` rows. For a degraded or failing state change the wording; the dot keeps `--color-success` only when all is well.
- `.ds-stat`: the row of four tiles. Each `.ds-stat__item` (`--2`, `--3`, `--4` for the other fills) is a mono `.ds-stat__label`, a `.ds-stat__row` with `.ds-stat__figure` and `.ds-stat__delta` (`--down`), and `.ds-stat__spark`, an inline SVG line along the foot.
- `.ds-uptime`: one `.ds-uptime__row` per service: `.ds-uptime__name` with a dot, `.ds-uptime__days` of 60 `.ds-uptime__day` bars (`--warn`, `--down`, `--none`) and `.ds-uptime__pct`. `.ds-uptime__axis` closes the panel with the ends of the range and `.ds-uptime__key` entries.
- `.ds-chart`: `.ds-chart__plot` is an inline SVG whose groups take their colour from `.ds-chart__grid`, `.ds-chart__p50`, `.ds-chart__p95` and `.ds-chart__p99`; `.ds-chart__legend` of `.ds-chart__key` names the lines.
- `.ds-grid`: region cells. `.ds-grid__cell` (`--2`, `--3`, `--4`) holds `.ds-grid__code` with a dot, `.ds-grid__title`, `.ds-grid__figure` and a `.ds-meter` whose `.ds-meter__fill` (`--mid`, `--high`) shows the load.
- `.ds-list`: incidents as rows of one box. `.ds-list__item` is `.ds-list__meta` (mono date and duration), then `.ds-list__title` with `.ds-list__text`, then a state badge.
- `.ds-callout`: the inverted block: `.ds-callout__title`, `.ds-callout__text` and one `.ds-button--inverse`.
- `.ds-page-header`: head of an inner page: `.ds-page-header__copy` (breadcrumb, `.ds-page-header__title`, `.ds-page-header__text`) and `.ds-page-header__actions`.
- `.ds-prose`: running text: h1, h2, h3, p, lists, `code`, `strong`; `.ds-prose__lead` for the first paragraph.
- `.ds-code`: a code block on `--fill-input`: optional `.ds-code__head` with `.ds-code__name`, then `.ds-code__body` (a `pre`). Spans `.ds-code__k`, `__s`, `__c`, `__n` colour keywords, strings, comments and names.
- `.ds-eyebrow`, `.ds-mono`, `.ds-kbd`, `.ds-dot`, `.ds-avatar`: the mono label (`--accent`), mono data text, a key cap, a status dot (`--success`, `--warning`, `--danger`, `--accent`, `--ring`) and an initials disc.
- `.ds-link`: text link with `is-visited`, `is-hover`, `is-active`, `is-focus`; `--quiet`, `--more` (appends an arrow). `.ds-linkrow` lays links in a row.
- `.ds-button`: 13px medium text. `--secondary` (dark fill, strong hairline), `--danger` (tint, hairline and text in the danger colour), `--inverse` (on the callout only), `--small`, `--large`. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttonrow` spaces them 8px apart.
- `.ds-form`: labels above fields. `.ds-form__row`, `.ds-form__field`, `.ds-form__label`, `.ds-form__input` (`--mono`), `.ds-form__select`, `.ds-form__textarea`, `is-error` plus `.ds-form__error`, `.ds-form__hint`, `.ds-form__check` with `.ds-form__checkbox`, `.ds-form__actions`.
- `.ds-table`: inside a panel, edge to edge: a mono uppercase head on `--fill-bar-alt`, hairline rows, `is-hover` rows on `--color-surface-alt`, `.ds-table__num` for right-aligned mono numbers. Wrap in `.ds-table__wrap`; `.ds-table__foot` closes the panel and holds the pagination.
- `.ds-panel`: hairline box on `--fill-panel`: `.ds-panel__head` with `.ds-panel__title`, `.ds-panel__body`; `.ds-panel__rows` of `.ds-panel__row` (`.ds-panel__key`) for facts. Uptime rows, the chart, a timeline and rules sit in it edge to edge.
- `.ds-tabs`: joined cells of `.ds-tabs__tab`; the current one (`is-current`) is on `--color-surface-strong` over a line in `--color-accent-alt`. `.ds-toolbar` puts it in a row with a `.ds-toolbar__note`.
- `.ds-badge`: small label on `--fill-accent`. `--quiet`, `--outline`, `--new` (accent-alt tint), `--mono`. Status variants: `--success`, `--warning`, `--danger`. `.ds-badgerow` for several.
- `.ds-sidebar`: a box with a mono `.ds-sidebar__title` strip on `--fill-bar-alt` over `.ds-sidebar__list`; each `.ds-sidebar__item` is a dot, a link and a mono word at the right, the current one (`is-current`) on `--color-surface-alt`.
- `.ds-timeline`: the updates of an incident: `.ds-timeline__item` is a mono time, a state badge and `.ds-timeline__text`.
- `.ds-rules`: alert rules: `.ds-rules__item` is `.ds-rules__name` (with its condition as `.ds-mono`), a channel badge and `.ds-switch` (`is-on`), a static drawing of a toggle.
- `.ds-notice`: message strip on `--color-notice`; `--error` in the danger colours. `.ds-notice__title`; `.ds-noticerow` stacks notices.
- `.ds-pagination`: `.ds-pagination__link`; `is-current` on `--color-surface-strong`, `is-disabled` grey.
- `.ds-breadcrumb`: `.ds-breadcrumb__item` parted by slashes; the last is `.ds-breadcrumb__current`.
- `.ds-dialog`: a block in `--color-overlay` centring `.ds-dialog__box` with `--shadow-dialog`: `.ds-dialog__head` (`.ds-dialog__title`, `.ds-dialog__close`), `.ds-dialog__body`, `.ds-dialog__actions`.
- `.ds-empty`: a dashed box: `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one small button.
- `.ds-footer`: one line on `--fill-bar-alt`: `.ds-footer__inner` with the brand, a mono `.ds-footer__note` and `.ds-footer__links` of `.ds-footer__link` at the right.
- `.ds-menu`: a drop-down behind a button in the bar (`.ds-menu__button`, `.ds-menu__list`); it lists every view, so on a phone it is the full navigation.
- `.ds-menu__item`: one link of the drop-down; the current view carries the accent line and the stronger fill.
- `.ds-menu__item--view`: an item of the menu that names a view of the site; the current view's item is marked by one `:has()` rule per view, like the bar links, so the menu is a second way to every view on a phone.
- `.ds-menu__title`: a small mono label over a group of links in the drop-down.
- `.ds-tooltip`: a small question mark beside a title; `.ds-tooltip__tip` opens on hover or focus, never by default.
- `.ds-progress`: a thin bar for a build, a rollout or a quota, with a mono line above (`.ds-progress__head`, `.ds-progress__value`, `.ds-progress__bar`).

## Never

- `border-width <= 2px`: rules and outlines are 1px hairlines; 2px marks the current section, the current tab and the top of a region cell.
- `border-radius <= 0px`: every box is square; dots and the ringed dot are circles.
- `box-shadow-blur <= 0px`: shadows are hard offsets or lines, never soft.
- `text-shadow = none`: text is flat on every fill.
- `font-size >= 12px`: badges, mono labels and facts are the smallest text.
- `font-size <= 44px`: the verdict of the banner is the largest text.
- `font-weight <= 700`: headings are bold; nothing is heavier.
- `font-families <= 1`: one monospace family for everything.
- `line-height <= 1.7`: body text is set at 1.7.
- `uppercase-text <= 15%`: headings and mono labels are uppercase, running text never.
- `letter-spacing <= 1px`: labels are spaced by 6% of their size, headings by 0.4px.
- `transition = none`: states change at once.
- `animation = none`: nothing moves by itself, not even the status dot.
- `block-gap <= 64px`: blocks are 20px apart, the footer 64px below the last.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new block is a `.ds-panel` with a 1px border in `--color-border`, a direct child of `.ds-main` or of a `.ds-cols` column; a new measurement is a `.ds-stat__item` with a sparkline, a new service is a `.ds-uptime__row`, a new place or shard is a `.ds-grid__cell`. Draw charts as inline SVG painted with `currentColor` and give each line its colour through a class that mixes a `--color-fill-*` into `--color-heading`. Keep to the conventions above for which text sits on which fill: status colours are solid only where no text sits on them, tinted otherwise. Use one inverted block per view. Spacing comes from the `--space-*` scale, and any one-off shade is a `color-mix()` of tokens noted next to the rule.
