# SaaS gradient developer portal

## Summary

This is the developer side of a start-up product of 2016 to 2021: a dark bar of navigation with a search field, a band of two-colour gradient whose lower edge rises to the left and carries the claim beside a request drawn as a code window, and a quickstart card that climbs onto the band. The reference is read in three columns, contents at the left, text and tables in the middle and the request and its reply in dark windows at the right, and the dashboard pages keep the same cards for keys, quotas and forms. Any product with an API had such a portal by 2018, friendly in the marketing colours and exact in the code.

## Layout

- The page is designed for a 1440px viewport and is not fluid: content sits in `.ds-wrap`, centred, at most `--size-page` (1180px) wide with `--space-5` side padding. The bar, the band and the footer run the full width of the window.
- Navigation is a solid bar on `--fill-bar-alt`, at least `--size-nav` (60px) tall, over a faint rule: the `.ds-brand`, a `.ds-nav__tag` naming the portal behind a hairline, four links (the current one on a translucent block), then at the right the `.ds-search` field, one `.ds-menu` and one small `.ds-button--bar`. Nothing of the navigation lies on the gradient.
- The home view stacks: `.ds-hero`, in two columns (5 to 6): kicker, title, one paragraph, two buttons and an install line at the left, a `.ds-code` window with a `.ds-reply` card over its corner at the right; its lower edge is cut by `clip-path` so that it rises 5vw to the left. Then `.ds-main--home`: `.ds-quick`, one card that climbs onto the band, a tinted lead with the progress bar and three `.ds-steps` across; `.ds-stat`, four figures set on the page; a `.ds-block__head` and the `.ds-grid` of six resources three across; `.ds-cols` (2 to 1, `--level`) with the recent changes beside the status rows; and `.ds-cta`, a gradient card. Blocks are `--space-7` apart.
- Inner pages have no hero. They open with `.ds-page-header`: the same gradient as a slim straight band holding the breadcrumb, the page title, one line of text and at most one button at the right. `.ds-main` follows. A reference page is `.ds-docs`, three columns: the `.ds-sidebar` lists (`--size-side`, 210px), the text column (the `.ds-endpoint` line, `.ds-prose`, the parameter table in a card, notices, the errors as an accordion, the pagination), and `.ds-docs__code` (`--size-code`, 400px), which stays in view while the text scrolls and holds the `.ds-tabs`, the code windows and the next links. Dashboard pages use `.ds-cols` (2 to 1); a list page uses `.ds-cols--side` with the sidebar at the left.
- Views. The specimen is the developer portal of a parcel shipping API, with four views. `home` is the overview. `reference` is the page of one request. `keys` is the list of API keys with the test mode switch, the form that makes a key with its buttons, the month's quotas, the open revoke dialog and an empty list of webhooks. `changelog` lists the changes of a year with the versions and the label key at the left and the pagination in the foot of the card.
- Below 1100px the code column moves under the text and the search field goes; below 960px the hero, the quickstart and the columns stack and cards stand two across; below 900px `.ds-nav__links` is not shown and the menu `.ds-nav__menu` takes its place at the right of the brand, while `.ds-nav__tag` and the search and menu in `.ds-nav__actions` are hidden; below 640px the menu goes again and the links return as one `.ds-nav__links` strip under the brand and button, scrolling sideways edge to edge and fading at the right (on the later views it starts from its far end, so that the current link is in sight), everything is one column and a wide table scrolls inside `.ds-table__wrap`. The page never scrolls sideways.
- Spacing scale: `--space-1` 4px (label to field, pagination gaps), `--space-2` 8px (icon to text, between badges), `--space-3` 12px (button and cell vertical padding), `--space-4` 16px (between form rows, code padding), `--space-5` 24px (card padding, grid gap, page gutter), `--space-6` 32px (column gap of the reference, page header padding), `--space-7` 48px (between blocks, hero top), `--space-8` 80px (hero foot, end of the page).
- Other sizes: `--size-mark` 28px, `--size-icon` 40px (icon tiles), `--size-control` 34px (pagination), `--size-avatar` 28px (avatars, step numbers), `--size-search` 200px, `--size-menu` 190px, `--size-tip` 220px (widest tooltip), `--size-dialog` 420px.

## Typography and colour roles

These are the conventions of the era; every layout and token set keeps them.

- The gradient is `--fill-inverse`, made in the surface part from `--color-fill-1` and `--color-fill-2`. Text on it is always `--color-inverse-text`. `--color-inverse` is the solid middle of the gradient: the text of the `--inverse` button on the band.
- On the gradient there are two buttons only: `.ds-button--inverse` (fill `--color-inverse-text`, text `--color-inverse`) and `.ds-button--ghost` (outline in `--color-inverse-text`). The primary button never stands on the gradient.
- `--fill-bar-alt` with `--color-bar-alt-text` is the navigation bar, the code windows and the footer; a palette may make it dark or pale, so everything on it is drawn from `--color-bar-alt-text` alone: rules and washes are that colour at 8 to 45%, the button on the bar is `.ds-button--bar`, and strings in code are `--color-accent-alt` mixed into it. The primary button never stands on it either.
- `--color-fill-1` to `--color-fill-4` are icon tiles, avatars and the rules beside the figures. A fill carries an icon or initials in `--color-inverse-text`, never running text.
- `--color-page` is the pale page, `--color-surface` the card, `--color-surface-alt` the quiet inset (lead of the quickstart, hovered row, dialog foot, empty state), `--color-surface-strong` the strongest neutral (quiet badge, inline code, `.ds-path`, current sidebar row, progress track).
- `--fill-bar` with `--color-bar-text` is the table head only.
- Text is `--color-text`, meta `--color-text-muted`, headings and figures `--color-heading`; sidebar titles and h3 in prose are `--color-heading-alt`. A tooltip is `--color-surface` text on `--color-heading`.
- Links are `--color-link`, medium weight, underlined only on hover; quiet links are `--color-link-quiet`.
- `--fill-button` with `--color-button-text` is the one action colour on cards: the primary button, the current page number, the switch when on, and under the current tab as `--color-button`. The secondary button is `--color-button-secondary-text` on `--fill-button-secondary` with a `--color-border` edge. The destructive button is `--color-danger` with `--color-button-text`.
- `--fill-accent` with `--color-accent-text` is the default badge (a version date). `--alt` and the status badges are tints: the colour mixed at 16 to 22% into `--color-surface`, with text mixed from the colour and `--color-heading`. `--color-success` also fills the number of a finished step.
- `--color-notice` with `--color-notice-text` is the information notice; the error notice is `--color-danger` on `--color-danger-surface` with a faint edge.
- One family for body, headings and controls (`--font-body`: Fira Sans, Helvetica Neue; the references used commercial grotesques) and `--font-mono` (Fira Mono) for code, keys, paths and the install line. `--text-base` 16px at `--line-body` 1.6; `--text-ui` 14px (controls, tables, card text); `--text-small` 12px (meta, badges, code, labels); `--text-large` 18px (leads); `--text-display` 44px (hero title); `--text-h1` 32px (page titles, figures); `--text-h2` 24px; `--text-h3` 16px (card titles).
- Headings are semibold: `--weight-heading`, `--weight-display` and `--weight-bold` 600, `--weight-ui` 500. Nothing is uppercase.
- Surface: corners are tight, `--radius-control` 4px, `--radius-panel` 6px, `--radius-pill` 3px for badges, `--radius-page` 10px for the call to action and the dialog backdrop. `--border-width` 1px in `--color-border-muted`; `--border-width-strong` 3px under the current tab and beside a figure. `--shadow-panel` is a close, short shadow; `--shadow-control` gives buttons and fields a darker lower lip; `--shadow-dialog` is the deepest and also lies under the hero's code window. `--fill-button` and `--fill-bar-alt` are faint gradients.
- One-off values made with `calc()` and `color-mix()` because the vocabulary has no token: the 5vw slant, washes on the gradient (`--color-inverse-text` at 14 to 16%, the ghost border at 55%), washes on the bar, tinted badges, the brand tile's edge, the error notice's edge. Opacity 0.6 to 0.9 dims secondary text on the band, the bar and in code. Code is set at a line height of 1.7.

## Components

- `.ds-page`: on `<body>`. Sets font, text colour and `--fill-page`.
- `.ds-wrap`: the centred content column.
- `.ds-brand`: the site's mark and name, first in the bar and again in the footer. `.ds-brand__mark` is a 28px tile dressed like the primary button (`--fill-button`, `--color-button-text`, `--radius-control`, `--shadow-control`) holding the mark as inline SVG; `.ds-brand__name` is the name in bold. Name and mark are placeholders for the installing project's own name and logo.
- `.ds-nav`: the bar. `.ds-nav__inner` centres it; `.ds-nav__tag` names the portal; `.ds-nav__links` holds `.ds-nav__link`, the current one on a translucent block (`is-current`); `.ds-nav__actions` holds the search, the menu and one `.ds-button--bar ds-button--small`. In the specimen the showing view marks its link with one `:has()` rule per view; a project sets `is-current` itself.
- `.ds-search`: the search field of the bar, an icon and a `.ds-search__input` in the bar's colours.
- `.ds-menu`: a drop-down, a `<details>` whose `.ds-menu__button` is the summary and whose `.ds-menu__list` of `.ds-menu__link` opens as a small card under it, aligned to the right.
- `.ds-hero`: the gradient band of the home view. `.ds-hero__inner` is two columns: `.ds-hero__copy` with `.ds-hero__kicker`, `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__actions` (one `--inverse`, one `--ghost` button) and `.ds-hero__note` in monospace; and `.ds-hero__art` with a `.ds-code--deep` and a `.ds-reply`. One per site.
- `.ds-code`: a code window: `.ds-code__bar` (optional `.ds-code__dot`s, `.ds-code__name`, a label or `.ds-code__copy`) over `.ds-code__body`, a `<pre>` that scrolls sideways. Inside, `.ds-code__b` is bold, `.ds-code__s` a string or value, `.ds-code__c` a comment. `--deep` takes the dialog shadow.
- `.ds-reply`: a small card of `.ds-reply__head` and `.ds-reply__row` pairs (`.ds-reply__value` at the right) laid over the corner of the hero's code window.
- `.ds-page-header`: head of an inner page, used instead of `.ds-hero`: `.ds-page-header__inner` holds the breadcrumb, `.ds-page-header__title`, `.ds-page-header__text` and `.ds-page-header__actions`.
- `.ds-main`: the area under the band; `--home` has no top padding. `.ds-block` is one block of it, `.ds-block__head` its `.ds-block__title` with a `.ds-block__text` at the right.
- `.ds-cols`: main and side column, 2 to 1; `--level` makes them equally tall; `--side` puts a narrow column first. `.ds-stack` stacks cards 24px apart.
- `.ds-quick`: the quickstart, a `.ds-panel` with `.ds-quick__lead` (title, `.ds-quick__note`, a progress bar) and `.ds-steps`, three `.ds-steps__item` each a `.ds-steps__num`, `.ds-steps__title`, `.ds-steps__text` and a link; `is-done` fills the number.
- `.ds-stat`: the row of figures; each `.ds-stat__item` is a `.ds-stat__figure` over a `.ds-stat__label` behind a rule in one of the fills.
- `.ds-grid`: three cards across. Each `.ds-grid__cell` is a `.ds-grid__head` (a `.ds-tile` and the `.ds-grid__title`), `.ds-grid__text` and a `.ds-path`.
- `.ds-tile`: an icon tile on `--color-fill-1` (`--2`, `--3`, `--4`; `--round` for a circle).
- `.ds-path`: a request path or a key prefix in monospace on a grey chip; `.ds-key` is the same without the chip, for tables and running text.
- `.ds-rows`: plain rows in a card, a `.ds-rows__name` with a badge or figure at the right, `.ds-rows__item` each.
- `.ds-cta`: the closing call to action, a gradient card with `.ds-cta__title`, `.ds-cta__text` and the two gradient buttons.
- `.ds-docs`: the three columns of a reference page; `.ds-docs__code` is the third. `.ds-endpoint` is the line that names the request: a badge for the method, `.ds-endpoint__path` and an optional label.
- `.ds-prose`: running text: h1, h2, h3, p, lists, `code`, `strong`. `.ds-prose__lead` for the first paragraph.
- `.ds-link`: text link. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `--quiet` for meta; `--more` appends an arrow. `.ds-linkrow` lays links in a row.
- `.ds-button`: 14px medium, tight corners. Variants: `--secondary`, `--danger` (destructive actions only), `--inverse` and `--ghost` (on the gradient only), `--bar` (on the bar only), `--small`, `--wide`. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttonrow` spaces buttons.
- `.ds-form`: vertical form with labels above the fields. `.ds-form__row` puts two fields side by side; `.ds-form__field` wraps `.ds-form__label` and one of `.ds-form__input`, `.ds-form__select`, `.ds-form__textarea`. Error: `is-error` on the input plus `.ds-form__error`. `.ds-form__hint`, `.ds-form__check` with `.ds-form__checkbox`, `.ds-form__actions`.
- `.ds-table`: data table inside a card: a small bold head on `--fill-bar`, hairlines between rows; `is-hover` tints a row. `.ds-table__num` right-aligns numbers; `.ds-table__title` is a row's link. Wrap it in `.ds-table__wrap`; `.ds-table__foot` is a last row for a count and the pagination.
- `.ds-list`: rows parted by hairlines. `.ds-list__item` holds a `.ds-list__title` (alone or in a `.ds-list__head` with a badge at the right), an optional `.ds-list__text` and `.ds-list__meta`.
- `.ds-panel`: the white card. `.ds-panel__head` holds the `.ds-panel__title` and an optional badge, link or `.ds-panel__tools` above a hairline; `.ds-panel__body` is padded; `.ds-panel__text` is a small paragraph.
- `.ds-tabs`: a row of `.ds-tabs__tab` over a hairline, the current one (`is-current`) underlined in `--color-button`. Here it chooses the language of the code windows.
- `.ds-badge`: small tag, 12px semibold, on `--fill-accent`. `--alt` tinted, `--quiet` grey. Status variants: `--success`, `--warning`, `--danger`. `.ds-badgerow` for several.
- `.ds-sidebar`: a block of links: `.ds-sidebar__title` and a `.ds-sidebar__list` of `.ds-sidebar__item` with an optional `.ds-sidebar__count`; the current row (`is-current`) is a grey block.
- `.ds-notice`: tinted message; `--error` in the danger colours. `.ds-notice__title` for the bold lead-in. `.ds-noticerow` stacks notices.
- `.ds-pagination`: `.ds-pagination__link` blocks; `is-current` is filled with `--fill-button`, `is-disabled` is grey.
- `.ds-breadcrumb`: a small trail of `.ds-breadcrumb__item` parted by slashes; the last is `.ds-breadcrumb__current`. It takes the colour of what it stands on.
- `.ds-dialog`: a dimmed block (`--color-overlay`) centring `.ds-dialog__box`, a card in `--shadow-dialog`: `.ds-dialog__head` (`.ds-dialog__title`, `.ds-dialog__close`), `.ds-dialog__body`, and `.ds-dialog__actions` right-aligned on `--color-surface-alt`.
- `.ds-empty`: a centred message on `--color-surface-alt`: `.ds-empty__icon` (a circle of `--color-fill-3`), `.ds-empty__title`, `.ds-empty__text` and one small secondary button.
- `.ds-accordion`: sections that open and close: each `.ds-accordion__item` is a `<details>` with a `.ds-accordion__title` summary (a plus that turns to a minus) and a `.ds-accordion__body`. Used for the errors of a request.
- `.ds-tooltip`: a hint shown on hover or focus of its trigger; `.ds-tooltip__text` is the hint, above the trigger. `--mark` draws the trigger as a small round question mark; `--end` aligns the hint to the trigger's right edge; `is-open` shows it.
- `.ds-switch`: an on/off control, a `<label>` with a hidden `.ds-switch__input` checkbox, the `.ds-switch__track` and the text; on, the track takes `--fill-button`.
- `.ds-progress`: a thin track with a `.ds-progress__bar` in the gradient (`--33`, `--62`, `--82`, `--90` set its width; `--over` fills it in `--color-danger`). `.ds-meter` puts a `.ds-meter__row` (`.ds-meter__name` and the figures) over one.
- `.ds-avatar`: a person as initials on a round fill (`--1`, `--2`, `--4`); `.ds-who` sets one beside a name.
- `.ds-footer`: one row on `--fill-bar-alt`: `.ds-footer__row` holds the brand, a `.ds-footer__list` of `.ds-footer__link` and `.ds-footer__legal` at the right.
- `.ds-nav__menu`: the `.ds-menu` that replaces the bar's links on a phone. It lists every view, marks the one showing, and sits last in the bar after the brand and the one action that stays.

## Never

- `text-shadow = none`: text is flat, also on the gradient and in code.
- `animation = none`: nothing moves by itself; only hover states ease.
- `border-width <= 3px`: borders are hairlines; 3px marks the current tab and stands beside a figure.
- `border-radius <= 10px`: cards are 6px and the largest blocks 10px; only circles are rounder.
- `box-shadow-blur <= 60px`: shadows are close and end at 60px.
- `font-weight <= 600`: headings and labels are semibold, never heavy.
- `font-size >= 12px`: meta, badges and code are the smallest text.
- `font-size <= 44px`: the hero title is the largest text.
- `font-families <= 2`: one sans-serif family, plus monospace for code.
- `line-height <= 1.7`: body text is set at 1.6 and code at 1.7.
- `uppercase-text <= 0%`: navigation, buttons and labels are in sentence case.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new block of the overview is a `.ds-block` with a `.ds-block__head`; a new reference page keeps the three columns of `.ds-docs` and puts every request and reply in a `.ds-code` window; anything that holds data is a `.ds-panel`. Text goes on cards, on the page and, in `--color-inverse-text`, on the gradient; on the bar and in code windows everything is drawn from `--color-bar-alt-text`. The four fills take an icon or initials only. Put only the `--inverse` and `--ghost` buttons on the gradient, only `--bar` on the bar, and at most one primary button in a card. Keep corners, shadows and borders on their tokens, spacing on the `--space-*` scale, and note any one-off `color-mix()` shade next to its rule.
