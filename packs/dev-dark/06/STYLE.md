# Dark developer tools pricing and sign-up

## Summary

This is the buying side of a developer service from 2019 to 2026: a pricing page that leads into a short sign-up, an order review and the team settings behind it. A rounded bar floats at the top, the plans are joined columns of one box, a calculator of static sliders stands beside its estimate, and each later step sits under a numbered row that shows how far the visitor has come. In this set the page is a dark green-grey, type is a soft humanist sans, corners are large and shadows gentle, and the one accent is an amber that also glows behind the claim.

## Layout

- Designed for a 1440px viewport, not fluid beyond it. Content sits in `.ds-wrap`, at most `--size-page` (1080px) wide with `--space-5` (24px) of side padding. Section heads are centred and kept to `--size-narrow` (620px).
- Navigation is `.ds-nav`, which stays at the top of the window with `--space-3` of air above it; its `.ds-nav__inner` is a rounded bar (`--radius-page`, `--shadow-panel`) on `--fill-bar`, as wide as the content, `--size-top` (52px) tall: brand at the left, four links centred, whose current one has a hairline outline, a secondary and a primary small button at the right. There is no rail.
- The home view stacks: `.ds-hero` (a badge, the title, a lead, two large buttons and a mono note, all centred over the glow); `.ds-tiers`, four plans as joined columns; a `.ds-colhead` and `.ds-cols`, the calculator panel beside the inverted `.ds-estimate` (`--size-side` 344px); a `.ds-colhead` and `.ds-grid`, four cells of two widths; `.ds-stat`, four figures in one rounded strip; a `.ds-colhead` and `.ds-list`, questions with their answers. Sections are `.ds-section`, `--space-8` (88px) apart; `--tight` is 36px.
- Inner pages have no hero. Inside `.ds-wrap` they open with `.ds-page-header`, a rounded band holding the breadcrumb, a `--text-h1` title and one muted line at the left and at most two buttons at the right. A page of the sign-up follows it with `.ds-steps`; then `.ds-cols` puts the work at the left and a 344px column of summaries at the right. A settings page uses `.ds-cols--menu`, a `--size-menu` (220px) column with the `.ds-sidebar` at the left; tabs are a segmented control in a `.ds-toolbar` over the table, and the pagination is in the `.ds-table__foot`.
- Views. The specimen sells a build service in four views. `pricing` is the home view described above. `signup` is the account step: two notices and the form at the left; the chosen plan with its total, a terminal block and the terms of the trial at the right. `review` is the order: the table of line items and the billing rules as prose at the left; the inverted amount due, the payment panel with its buttons and a row of document links at the right. `team` is the settings page behind it: the section menu and the seats in use at the left; tabs, the table of members with pagination, the open dialog that removes a member, the key of role and state labels and an empty state at the right.
- Spacing scale: `--space-1` 4px (inside the segmented control, between navigation links), `--space-2` 8px (button vertical padding, between adjacent buttons and badges, menu rows), `--space-3` 12px (panel head and table cell vertical padding, air above the bar), `--space-4` 16px (panel padding, between grid cells, between form rows, slider rows), `--space-5` 24px (between stacked blocks, side padding, plan and cell padding), `--space-6` 36px (below a section head, above a page header), `--space-7` 56px (hero bottom, sections on a narrow window), `--space-8` 88px (between sections, above the footer).
- Other sizes: `--size-thumb` 18px and `--size-track` 6px (a slider), `--size-art` 96px (the drawing of a cell), `--size-glow` 420px, `--size-mark` 26px, `--size-dot` 8px, `--size-dialog` 420px.
- Below 960px plans stand two across, cells take the full row and every column pair stacks. Below 640px everything is one column, the links of the bar move to a second row that scrolls sideways, the steps stand two across without their rule, a question sits over its answer, and a wide table scrolls inside `.ds-table__wrap`; the page itself never scrolls sideways.
- On a phone (640px and below) the floating bar wraps inside `--radius-panel`: the brand and `.ds-nav__actions` stay on the first row, `.ds-nav__links` take a second row, centred when they fit and scrolling sideways when they do not, and the secondary button of the actions is not shown. The `.ds-menu` stays and lists every view, the current one with the accent line.

## Typography and colour roles

The conventions for which text sits on which fill are those of the era's first layout and are kept here.

- `--color-text`, `--color-heading` and `--color-text-muted` sit on the page, on `--fill-panel` (panels, plans, cells, the page header, the menu), `--color-surface-alt` (the stat strip, the segmented control, a hovered row, the dialog) and `--color-surface-strong` (slider tracks, the current menu row, key caps, quiet badges). `--color-bar-text` sits on `--fill-bar` (the floating bar). `--color-bar-alt-text` sits on `--fill-bar-alt` (table heads, code heads, the footer), with `--color-heading` for the brand there. `--color-button-text` sits on `--fill-button`, which also dresses the current segment of the tabs and the thumb of a slider. `--color-accent-text` sits on `--fill-accent` (the default badge, the disc of the current step). `--color-inverse-text` sits on `--fill-inverse`, used once per view, for `.ds-estimate`.
- `--color-fill-1` to `--color-fill-4` never carry text. Each tints one corner of a grid cell at 20% and colours its drawing (72% of the fill mixed into `--color-heading`); `--color-fill-1` is also the glow behind the hero, the tint of the recommended plan (17%) and of the page header (14%).
- `--color-accent` rules the top of the recommended plan and fills the used part of a slider through `--fill-accent`. `--color-accent-alt` is a text colour: the mono label of a section, the plus before each plan feature, keywords in code.
- `--color-success`, `--color-warning` and `--color-danger` are a tick, a dot, or a 12% tint with a 32% hairline and the colour as text. Nothing is filled solid with them; the destructive button is a tint too.
- Borders: `--color-border` on boxes and between plans, `--color-border-strong` on the bar, the box of plans, secondary buttons, step discs and their rule, the total and the dialog, `--color-border-muted` between rows inside a box. All `--border-width`; `--border-width-strong` is the rule of the recommended plan, the ring of a slider thumb and the focus offset.
- Type. `--font-body` for text and controls, `--font-heading` for titles, prices and figures, `--font-mono` for units, dates, counts and code. `--text-base` 15px at `--line-body` 1.5; `--text-ui` 14px for navigation, buttons, forms, tables, plan features and cell text; `--text-small` 13px for badges, mono labels and scales; `--text-large` 19px for the hero lead. `--text-display` 52px is the hero title and the estimate's figure; `--text-h1` 34px section and page titles, prices and stat figures; `--text-h2` 22px the total; `--text-h3` 16px panel, plan, cell and question titles. Headings are `--weight-heading` 600.
- Uppercase appears only on mono labels (`.ds-eyebrow`, `.ds-estimate__label`, table heads, the menu title), tracked by 6% of the font size. Links are `--color-link` without underline and gain one on hover; quiet links are `--color-link-quiet` and turn `--color-heading`.
- Surface. `--radius-control` 14px for buttons, inputs, menu rows and tab segments, `--radius-panel` 20px for panels, code, notices and the menu, `--radius-page` 28px for the bar, the box of plans, cells, the stat strip, the page header, the estimate, the dialog box and the top of the footer, `--radius-pill` for badges, navigation links and slider tracks. `--shadow-panel` is on the bar, panels, cells, the page header and the menu, `--shadow-control` on buttons, the brand tile, the slider thumb and the current segment, `--shadow-dialog` on the box of plans and the dialog.
- One-off values written with `calc()` and `color-mix()`: the glow (20% of a fill), plan, cell and header tints (17%, 20%, 14%), drawing colours (72%), status tints (12%, 32%, 45%), the outline of the current link (36% of the bar's text colour), rules inside the estimate (16% of its text colour), the notice border (22% and 40%), the corner of the segmented control (`--radius-control` plus `--space-1`), mono label tracking, the width of the bar (`--size-page` less two `--space-5`). Navigation links are at `opacity` 0.64 until current, secondary text of the estimate at 0.7. The empty state's outline is dashed.

## Components

- `.ds-page`: on `<body>`. Sets the font, text colour and `--fill-page`. `.ds-nav`, the views and `.ds-footer` are its direct children.
- `.ds-wrap`, `.ds-section`, `.ds-stack`, `.ds-cols`, `.ds-colhead`: the centred column, a section 88px below the last (`--tight`), a vertical stack 24px apart (`.ds-panel--grow` in it takes spare height), the column pairs (`--even`, `--menu`) and a centred section head (`.ds-eyebrow`, `.ds-colhead__title`, `.ds-colhead__text`).
- `.ds-brand`: the site's mark and name. `.ds-brand__mark` is a 26px tile dressed like the primary button holding the mark as inline SVG; `.ds-brand__name` is the name. Both are placeholders for the installing project's own name and logo.
- `.ds-nav`: the floating bar. `.ds-nav__inner` holds the brand, `.ds-nav__links` of `.ds-nav__link` (current: `is-current`) and `.ds-nav__actions`. In the specimen the showing view marks its link with one `:has()` rule per view; a project sets `is-current`.
- `.ds-hero`: home view only. A badge, `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__actions` and `.ds-hero__note`, centred. The glow is its `::before`.
- `.ds-tiers`: the plans as one box of joined columns. Each `.ds-tier` (`--featured` for the recommended one) holds `.ds-tier__name`, `.ds-tier__price` with `.ds-tier__per`, `.ds-tier__text`, one button and `.ds-tier__items` of `.ds-tier__item`.
- `.ds-range`: a static slider: `.ds-range__label`, `.ds-range__value`, `.ds-range__track` with `.ds-range__fill` (`--low`, `--high`; the thumb is its `::after`) and an optional `.ds-range__scale`. `.ds-calc__extras` rows the checkboxes under the sliders. A project that makes it live keeps the same parts.
- `.ds-estimate`: the inverted block: `.ds-estimate__label`, `.ds-estimate__figure`, `.ds-estimate__lines` of `.ds-estimate__line`, one `.ds-button--inverse` and `.ds-estimate__note`.
- `.ds-page-header`: head of an inner page: `.ds-page-header__inner` with `.ds-page-header__copy` (breadcrumb, `.ds-page-header__title`, `.ds-page-header__text`) and `.ds-page-header__actions`.
- `.ds-steps`: the steps of the sign-up: `.ds-steps__item` with `is-done` (a tick) and `is-current` (a filled disc), joined by a rule.
- `.ds-prose`: running text: h1, h2, h3, p, lists, `code`, `strong`; `.ds-prose__lead` for the first paragraph.
- `.ds-code`: a code block on `--fill-input`: optional `.ds-code__head` with `.ds-code__name`, then `.ds-code__body` (a `pre`). Spans `.ds-code__k`, `__s`, `__c`, `__n` colour keywords, strings, comments and names.
- `.ds-eyebrow`, `.ds-mono`, `.ds-kbd`, `.ds-dot`, `.ds-avatar`: the mono label (`--accent`), mono data text, a key cap, a status dot (`--success`, `--warning`, `--danger`, `--accent`, `--ring`) and an initials disc; `.ds-person` puts one before a name.
- `.ds-link`: text link with `is-visited`, `is-hover`, `is-active`, `is-focus`; `--quiet`, `--more` (appends an arrow). `.ds-linkrow` lays links in a row.
- `.ds-button`: 14px semibold text, 14px corners, a soft shadow. `--secondary` (dark fill, strong hairline), `--danger` (tint, hairline and text in the danger colour), `--inverse` (on the estimate only), `--small`, `--large`. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttonrow` spaces them 8px apart.
- `.ds-form`: labels above fields. `.ds-form__row`, `.ds-form__field`, `.ds-form__label`, `.ds-form__input` (`--mono`), `.ds-form__select`, `.ds-form__textarea`, `is-error` plus `.ds-form__error`, `.ds-form__hint`, `.ds-form__check` with `.ds-form__checkbox`, `.ds-form__actions`.
- `.ds-table`: inside a panel, edge to edge: a mono uppercase head on `--fill-bar-alt`, hairline rows, `is-hover` rows on `--color-surface-alt`, `.ds-table__num` for right-aligned mono numbers. Wrap in `.ds-table__wrap`; `.ds-table__foot` closes the panel and holds the pagination.
- `.ds-list`: questions between rules. `.ds-list__item` is the `.ds-list__title` with a mono `.ds-list__meta` under it at the left and the answer, `.ds-list__text`, at the right.
- `.ds-panel`: hairline box on `--fill-panel`: `.ds-panel__head` with `.ds-panel__title`, `.ds-panel__body`; `.ds-panel__rows` of `.ds-panel__row` (`.ds-panel__key`) for facts, closed by `.ds-total` with `.ds-total__figure` when they add up.
- `.ds-stat`: the row of figures, one rounded strip. Each `.ds-stat__item` is a `.ds-stat__figure` over a `.ds-stat__label`, parted by vertical hairlines.
- `.ds-grid`: what every plan includes. `.ds-grid__cell` (`--wide`; `--2`, `--3`, `--4` for the other fills) holds `.ds-grid__body` with a mono number, `.ds-grid__title` and `.ds-grid__text`, and `.ds-grid__art`, one inline SVG drawn in strokes.
- `.ds-tabs`: a segmented control of `.ds-tabs__tab`; the current one (`is-current`) is dressed like the primary button. `.ds-toolbar` puts it in a row with a `.ds-toolbar__note`.
- `.ds-badge`: small pill on `--fill-accent`. `--quiet`, `--outline`, `--new` (accent-alt tint), `--mono`. Status variants: `--success`, `--warning`, `--danger`. `.ds-badgerow` for several.
- `.ds-sidebar`: the menu of a settings page, a box with a mono `.ds-sidebar__title` over `.ds-sidebar__list`; the current `.ds-sidebar__item` (`is-current`) is on `--color-surface-strong`.
- `.ds-notice`: message strip on `--color-notice`; `--error` in the danger colours. `.ds-notice__title`; `.ds-noticerow` stacks notices.
- `.ds-pagination`: `.ds-pagination__link`; `is-current` on `--color-surface-strong`, `is-disabled` grey.
- `.ds-breadcrumb`: `.ds-breadcrumb__item` parted by slashes; the last is `.ds-breadcrumb__current`.
- `.ds-dialog`: a block in `--color-overlay` centring `.ds-dialog__box` with `--shadow-dialog`: `.ds-dialog__head` (`.ds-dialog__title`, `.ds-dialog__close`), `.ds-dialog__body`, `.ds-dialog__actions`.
- `.ds-empty`: a dashed box: `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one small button.
- `.ds-footer`: centred on `--fill-bar-alt`, rounded at the top: the brand, `.ds-footer__links` of `.ds-footer__link` and the mono `.ds-footer__line`.
- `.ds-menu`: a drop-down behind a button in the bar (`.ds-menu__button`, `.ds-menu__list`); it lists every view, so on a phone it is the full navigation.
- `.ds-menu__item`: one link of the drop-down; the current view carries the accent line and the stronger fill.
- `.ds-menu__item--view`: an item of the menu that names a view of the site; the current view's item is marked by one `:has()` rule per view, like the bar links, so the menu is a second way to every view on a phone.
- `.ds-menu__title`: a small mono label over a group of links in the drop-down.
- `.ds-tooltip`: a small question mark beside a title; `.ds-tooltip__tip` opens on hover or focus, never by default.
- `.ds-switch`: an on or off setting that takes effect at once, drawn over a checkbox (`.ds-switch__input`, `.ds-switch__track`, `.ds-switch__label`).
- `.ds-progress`: a thin bar for a build, a rollout or a quota, with a mono line above (`.ds-progress__head`, `.ds-progress__value`, `.ds-progress__bar`).

## Never

- `border-width <= 2px`: rules and outlines are 1px hairlines; 2px is the rule of the recommended plan and the ring of a slider thumb.
- `border-radius <= 28px`: 14px on controls, 20px on panels, 28px on the largest boxes; pills are shapes.
- `box-shadow-blur <= 60px`: only the box of plans and the dialog cast a wide shadow.
- `text-shadow = none`: text is flat on every fill.
- `font-size >= 13px`: badges, mono labels and scales are the smallest text.
- `font-size <= 52px`: the hero title and the estimate are the largest text.
- `font-weight <= 600`: headings and controls are semibold; nothing is bold or black.
- `font-families <= 3`: a sans for text, a sans for headings and one monospace.
- `line-height <= 1.5`: body text is set at 1.5.
- `uppercase-text <= 5%`: only mono labels are uppercase.
- `underlined-links <= 5%`: links are told apart by colour and underline on hover only.
- `letter-spacing <= 1px`: only mono labels are spaced out, by 6% of their size.
- `animation = none`: nothing moves by itself.
- `block-gap <= 88px`: sections are 88px apart at most, blocks inside them 24px.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new block is a `.ds-panel` with a 1px border in `--color-border`, or a `.ds-grid__cell` when it promotes something; a further plan is one more `.ds-tier` in the same box, a further setting of the calculator one more `.ds-range`, a further step of the flow a `.ds-steps__item` and a view that opens with `.ds-page-header`. Anything that adds up ends in a `.ds-total`. Keep to the conventions above for which text sits on which fill: never put text on a full-strength `--color-fill-*` or status colour, tint it instead. Use one glow and one inverted block per view. Spacing comes from the `--space-*` scale, and any one-off shade is a `color-mix()` of tokens noted next to the rule.
