# SaaS gradient dashboard

## Summary

This is the signed-in side of a start-up product of 2016 to 2021: a white rail of navigation at the left, and beside it a band of two-colour gradient that carries the search field, the greeting and the page titles, with white rounded cards climbing onto it from the pale work area below. Figures stand in soft-shadowed cards with a round icon tile and a sparkline, charts are thin coloured lines, and an onboarding checklist with a gradient progress bar sits beside them. Buttons are pills and type is a rounded, friendly sans set bold for headings.

## Layout

- The page is designed for a 1440px viewport. `.ds-page` is a grid of two columns: the rail (`.ds-nav`, `--size-rail` 252px) for the whole height, and beside it the top bar, the showing view and the footer line. The work area is centred in its column: `.ds-content` is at most `--size-page` (1140px) wide with `--space-6` side padding and stacks its children `--space-5` apart.
- Navigation is the rail on `--fill-panel`: the `.ds-brand`, then groups of links under small group names, each link an icon in one of the four fills and a label with an optional count, the current one on a grey pill; a gradient card at the foot holds the trial notice. `.ds-topbar` is `--size-top` (72px) tall and transparent, lying on the gradient: the search pill at the left, two links and the account's initials at the right.
- The gradient band is the `.ds-hero` on the home view and the `.ds-page-header` on every other. It is pulled up under the top bar by a negative margin and is tall enough for `.ds-content--lift` to rise `--space-8` onto it, so the first row of cards always overlaps the band. Its lower edge is straight.
- The home view stacks: the band with the greeting and two buttons; `.ds-stat`, four figure cards; `.ds-cols` (2 to 1, `--level`), the chart card beside the setup checklist; a `.ds-section-title`; and `.ds-grid`, four project cards with progress bars.
- Inner pages have no hero. The `.ds-page-header` band holds the breadcrumb, the page title, one line of text and at most one white button at the right. Below, a list page is one card with `.ds-tabs` on its head, the table and a `.ds-table__foot` with the pagination, then two level cards; other pages use `.ds-cols` (2 to 1), `--even`, or `--side` (1 to 3) with the `.ds-sidebar` in a card at the left.
- Views. The specimen is a time tracker for a small studio with five views. `overview` is the home view. `projects` is the table of projects under tabs with pagination, the status key and an empty archive. `team` lists people beside the waiting invitations and the open remove dialog. `settings` is the workspace form with its notices and buttons beside the settings sections. `news` is one release note as running text beside earlier notes and related links.
- Below 1100px cards stand two across and columns stack; below 900px the rail becomes a white block across the top: its groups, icons and card are not shown, the links `.ds-nav__links` are replaced by the menu `.ds-nav__menu` at the right of the brand, which lists every view and marks the one showing, and the top bar keeps only search and initials; below 640px the menu goes and the two `.ds-nav__links` lists become a dock fixed to the foot of the window (the first list 60% of the width, the second 40%, items in equal cells with the `.ds-nav__icon` above a small `.ds-nav__label`, the current one `is-current`); below 520px everything is one column. A wide table scrolls inside `.ds-table__wrap`; the page never scrolls sideways.
- Spacing scale: `--space-1` 4px (label to field, between rail links), `--space-2` 8px (rail link padding, between badges), `--space-3` 12px (table cell and button vertical padding, checklist rows), `--space-4` 16px (figure card padding, list rows, rail padding), `--space-5` 24px (between cards, card padding), `--space-6` 32px (work area side padding), `--space-7` 48px (around a dialog), `--space-8` 72px (how far cards rise onto the band).
- Other sizes: `--size-mark` 30px, `--size-icon` 40px (icon tiles), `--size-avatar` 36px, `--size-control` 34px (pagination), `--size-chart` 220px, `--size-search` 300px, `--size-dialog` 420px.

## Typography and colour roles

The conventions are the era's, as its first layout sets them out.

- The gradient is `--fill-inverse` (from `--color-fill-1` to `--color-fill-2`); text on it is `--color-inverse-text`. It fills the band, the trial card in the rail and the progress bars. Only `.ds-button--inverse` and `.ds-button--ghost` stand on it; the search field on the band is a translucent pill in `--color-inverse-text`.
- `--color-fill-1` to `--color-fill-4` colour the rail icons, the round icon tiles, avatars, sparklines and the two chart lines. A fill carries an icon or initials in `--color-inverse-text`, never running text.
- `--color-page` is the work area, `--color-surface` (through `--fill-panel`) the cards and the rail, `--color-surface-alt` a hovered table row, the dialog foot and the empty state, `--color-surface-strong` the current rail link, the current sidebar row, the quiet badge and the progress track.
- `--fill-bar` with `--color-bar-text` is the table head. `--fill-bar-alt` with `--color-bar-alt-text` is the footer line.
- Text is `--color-text` (`#576f7f`), meta `--color-text-muted`, headings and figures `--color-heading` (`#0e304b`); sidebar titles and h3 in prose are `--color-heading-alt`. The change under a figure is `--color-success` (`.ds-stat__up`) or `--color-danger` (`.ds-stat__down`).
- `--color-button` (navy `#033156`, white text) is the action colour on cards, the current page number and the edge under the current tab. Links are `--color-link`, bold, underlined only on hover.
- Badges, notices and the destructive button follow the era's rules: `--fill-accent` for the default badge, tints mixed into `--color-surface` for `--alt` and the status variants, `--color-notice` with `--color-notice-text`, `--color-danger` with `--color-button-text`.
- Two families: `--font-heading` is a rounded sans (Varela Round, Arial Rounded MT Bold) for titles and figures, `--font-body` and `--font-ui` a plain humanist sans (Nunito, Open Sans, Helvetica Neue); `--font-mono` for code. `--text-base` 16px at `--line-body` 1.6; `--text-ui` 15px; `--text-small` 13px (labels, meta, badges); `--text-large` 20px (band text); `--text-h1` 38px (band titles, figures); `--text-h2` 27px; `--text-h3` 19px (card titles). `--text-display` 50px is not used in this layout: the greeting is `--text-h1`.
- Headings and controls are bold (`--weight-heading`, `--weight-ui` 700). Nothing is uppercase.
- Surface: pills for buttons, inputs, rail links and badges (`--radius-control`, `--radius-pill` 100px); cards `--radius-panel` 16px; `--radius-page` 28px on the dialog's backdrop. `--shadow-panel` is one very wide, soft, tinted shadow (60px blur); `--shadow-dialog` is deeper and also marks a hovered project card. `--border-width-strong` 3px is the edge under the current tab, the ring of an unticked checklist mark and the ring of the account's initials.
- One-off values made with `calc()` and `color-mix()`: the translucent search field (`--color-inverse-text` at 16%, its edge at 30%), the ghost button border (55%), the account ring's wash (20%), tinted badges, the hairlines of the chart (a repeating gradient of `--color-border-muted` every 25%), the area under the first chart line (its own colour at 10% opacity). Progress widths are percentages.

## Components

- `.ds-page`: on `<body>`. Sets font, text colour and `--fill-page` and makes the two-column grid; `.ds-nav`, `.ds-topbar`, the views and `.ds-footer` are its direct children.
- `.ds-content`: the centred column of a view; `--lift` raises it onto the band.
- `.ds-brand`: the site's mark and name, first in the rail. `.ds-brand__mark` is a 30px tile dressed like the primary button (`--fill-button`, `--color-button-text`, `--radius-control`, `--shadow-control`) holding the mark as inline SVG; `.ds-brand__name` is the name in bold. Name and mark are placeholders for the installing project's own name and logo.
- `.ds-nav`: the white rail. `.ds-nav__group` names a group; `.ds-nav__links` holds `.ds-nav__link` rows of a `.ds-nav__icon` (`--2`, `--3`, `--4` for the other fills), a `.ds-nav__label` and an optional `.ds-badge`; the current row is `is-current`. `.ds-nav__foot` is a gradient card with `.ds-nav__foot-title`, `.ds-nav__foot-text` and a small inverse button. In the specimen the showing view marks its link with one `:has()` rule per view; a project sets `is-current` itself.
- `.ds-topbar`: the transparent bar on the band: `.ds-topbar__search`, up to two `.ds-topbar__link` and a `.ds-avatar--plain`.
- `.ds-avatar`: a circle of initials on a fill (`--1`, `--2`, `--4`; `--plain` is a ring for use on the gradient).
- `.ds-hero`: the band of the home view. `.ds-hero__inner` holds the `.ds-hero__title` and `.ds-hero__lead` at the left and `.ds-hero__actions` (one `--inverse`, one `--ghost` button) at the right. One per site.
- `.ds-page-header`: the band of an inner page, used instead of `.ds-hero`. `.ds-page-header__inner` holds the breadcrumb, `.ds-page-header__title` and `.ds-page-header__text` at the left and `.ds-page-header__actions` at the right.
- `.ds-cols`: main and side column, 2 to 1; `--even` for halves, `--side` 1 to 3, `--level` to end the columns on one line. `.ds-stack` stacks cards; `.ds-panel--grow` fills a column.
- `.ds-stat`: the row of four figure cards. Each `.ds-stat__item` has a `.ds-stat__head` (the `.ds-stat__label` over the `.ds-stat__figure`, and a round `.ds-tile`), a `.ds-stat__delta` line with `.ds-stat__up` or `.ds-stat__down`, and a `.ds-spark`.
- `.ds-spark`: a sparkline, an inline SVG path stroked with `currentColor`; `--2`, `--3`, `--4` take the other fills.
- `.ds-tile`: an icon tile on `--color-fill-1` (`--2`, `--3`, `--4`; `--round`).
- `.ds-chart`: a line chart in a card: `.ds-chart__svg` over hairlines, its paths stroked with `currentColor`, a second line marked `.ds-chart__alt`; `.ds-chart__axis` labels the steps and `.ds-chart__legend` of `.ds-chart__key` (`--alt`) names the two series. No more than two lines.
- `.ds-progress`: a track with a gradient `.ds-progress__bar` (`--25`, `--40`, `--60`, `--75`, `--90`; `--over` is full and in the danger colour).
- `.ds-steps`: the onboarding checklist: `.ds-steps__item` rows of a round `.ds-steps__mark`, a `.ds-steps__text` and an optional link; `is-done` fills the mark with `--color-success` and strikes the text.
- `.ds-grid`: four project cards. Each `.ds-grid__cell` is a link holding a `.ds-tile`, `.ds-grid__title`, `.ds-grid__text` and a `.ds-progress`. `.ds-section-title` is the heading above such a row.
- `.ds-prose`: running text inside a card: h1, h2, h3, p, lists, `code`, `strong`. `.ds-prose__lead` for the first paragraph.
- `.ds-link`: text link. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `--quiet` for meta and row actions; `--more` appends an arrow. `.ds-linkrow` lays links in a row.
- `.ds-button`: pill button, 15px bold. Variants: `--secondary`, `--danger` (destructive actions only), `--inverse` and `--ghost` (on the gradient only), `--small`, `--wide`. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttonrow` spaces buttons.
- `.ds-form`: vertical form with labels above the fields, inside a card. `.ds-form__row` puts two fields side by side; `.ds-form__field` wraps `.ds-form__label` and one of `.ds-form__input`, `.ds-form__select`, `.ds-form__textarea`. Error: `is-error` on the input plus `.ds-form__error`. `.ds-form__hint`, `.ds-form__check` with `.ds-form__checkbox`, `.ds-form__actions`.
- `.ds-table`: data table inside a card, edge to edge: a small bold head on `--fill-bar`, hairlines between rows; `is-hover` tints a row. `.ds-table__num` right-aligns numbers; `.ds-table__title` is the row's link. Wrap it in `.ds-table__wrap`; `.ds-table__foot` is the card's last row, pagination at the left and a count at the right.
- `.ds-list`: rows parted by hairlines. `.ds-list__item` holds a `.ds-list__title` (alone or in a `.ds-list__head` with a badge or link at the right) and `.ds-list__meta`. `.ds-list--people` puts a `.ds-avatar` before a `.ds-list__grow` block and a badge after it.
- `.ds-panel`: the white card with the wide soft shadow. `.ds-panel__head` holds the `.ds-panel__title` and an optional badge, link or legend above a hairline; `.ds-panel__body` is padded; `.ds-panel__text` is a small paragraph; `--side` is a slim card around a `.ds-sidebar`.
- `.ds-tabs`: underlined tabs on the head of a card: `.ds-tabs__tab`, the current one (`is-current`) in the heading colour over a 3px edge in `--color-button`.
- `.ds-badge`: pill tag, 13px bold, on `--fill-accent`. `--alt` tinted, `--quiet` grey. Status variants: `--success`, `--warning`, `--danger`. `.ds-badgerow` for several.
- `.ds-sidebar`: a block of links: `.ds-sidebar__title` and a `.ds-sidebar__list` of `.ds-sidebar__item` with an optional `.ds-sidebar__count`; the current row (`is-current`) is a grey pill.
- `.ds-notice`: tinted rounded message; `--error` in the danger colours. `.ds-notice__title` for the bold lead-in. `.ds-noticerow` stacks notices.
- `.ds-pagination`: `.ds-pagination__link` pills; `is-current` is filled with `--fill-button`, `is-disabled` is grey.
- `.ds-breadcrumb`: a small trail of `.ds-breadcrumb__item` parted by slashes; the last is `.ds-breadcrumb__current`. It stands on the band, above the page title.
- `.ds-dialog`: a dimmed block (`--color-overlay`) centring `.ds-dialog__box`, a card in `--shadow-dialog`: `.ds-dialog__head` (`.ds-dialog__title`, `.ds-dialog__close`), `.ds-dialog__body`, and `.ds-dialog__actions` right-aligned on `--color-surface-alt`. Place it on the page, not on the band.
- `.ds-empty`: a centred message on `--color-surface-alt`: `.ds-empty__icon` (a circle of `--color-fill-3`), `.ds-empty__title`, `.ds-empty__text` and one small secondary button. It stands where a card would be.
- `.ds-footer`: one line on `--fill-bar-alt` under the work area: `.ds-footer__name`, a short line, and `.ds-footer__links` of `.ds-footer__link` at the right.
- `.ds-menu`: a drop-down, a `<details>` whose `.ds-menu__button` is the summary and whose `.ds-menu__list` of `.ds-menu__link` opens as a small card under it. Use it for a group of links under one bar item and for the account.
- `.ds-nav__menu`: the `.ds-menu` that replaces the bar's links on a phone. It lists every view, marks the one showing, and sits last in the bar after the brand and the one action that stays.
- `.ds-topbar__account`: the account `.ds-menu` at the end of the top bar, the name and the avatar as its button.
- `.ds-switch`: an on or off setting drawn over a checkbox (`.ds-switch__input`, `.ds-switch__track`, then the label). Use it for a setting that takes effect at once, not for a choice sent with a form's other fields.
- `.ds-tooltip`: a short hint on hover or focus; `.ds-tooltip--mark` is the small round question mark beside a label, `.ds-tooltip__text` the hint, `.ds-tooltip--start` opens it from the left edge. One sentence at most.

## Never

- `text-shadow = none`: text is flat, also on the gradient.
- `animation = none`: nothing moves by itself; only hover states ease.
- `border-width <= 3px`: borders are hairlines; 3px marks the current tab and the rings.
- `border-radius <= 28px`: cards are 16px and the largest block 28px; only pills and circles are rounder.
- `box-shadow-blur <= 80px`: shadows are very wide and soft but end at 80px.
- `font-weight <= 700`: headings and figures are bold, never black.
- `font-size >= 13px`: labels, meta and badges are the smallest text.
- `font-size <= 38px`: band titles and figures are the largest text.
- `font-families <= 3`: a rounded heading face, a body sans and monospace for code.
- `line-height <= 1.6`: body text is set at 1.6.
- `uppercase-text <= 0%`: navigation, buttons and labels are in sentence case.
- `block-gap <= 48px`: cards are 24px apart; nothing is further from its neighbour than 48px.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. Anything new in the work area is a `.ds-panel` with a `.ds-panel__head`, placed in `.ds-content` or in a column of `.ds-cols`, 24px from its neighbours; a new figure is another `.ds-stat__item` with its tile and sparkline, a new destination another `.ds-nav__link` with a 20px line icon in one of the fills. A new page opens with a `.ds-page-header` band and lifts its first card onto it. Text goes on cards and, in `--color-inverse-text`, on the gradient; fills take an icon or initials only. Put only `--inverse` and `--ghost` buttons on the gradient, and keep at most one primary button in a card. Spacing comes from the `--space-*` scale, and any one-off shade is a `color-mix()` of tokens noted next to the rule.
