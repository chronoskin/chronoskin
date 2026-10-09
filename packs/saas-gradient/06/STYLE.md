# SaaS gradient integrations directory

## Summary

This is the integrations directory of a start-up product of 2016 to 2021: a straight band of two-colour gradient carries the navigation, the claim, a small cloud of app marks drawn as initials on white tiles and three figures along its foot, and below it the catalogue stands in two columns, categories at the left and app tiles under a search field and filter chips at the right. An app's own page keeps a shorter band with its mark beside the title and shows the install dialog beside the description; a gradient card with a few lines of code invites developers. Most products with an API had such a marketplace by 2018.

## Layout

- The page is designed for a 1440px viewport and is not fluid: content sits in `.ds-wrap`, centred, at most `--size-page` (1160px) wide with `--space-5` side padding. The band and the footer run the full width of the window.
- Navigation is one transparent row, `--size-nav` (70px) tall, lying on the gradient: the `.ds-brand`, a `.ds-nav__tag` naming the directory behind a hairline, then at the right three links (the current one with a `--border-width-strong` edge under it), the sign-in link and one small button. The hero or the page header is pulled up under it by a negative margin of the same height. The band's lower edge is straight.
- The home view stacks: `.ds-hero`, in two columns (7 to 5): title, one paragraph and two buttons at the left, six `.ds-hero__app` tiles at the right, and `.ds-stat` along the foot of the band; then `.ds-main` with `.ds-cols--side` (1 to 4): the categories `.ds-sidebar` and a small card at the left; at the right a `.ds-toolbar` with the `.ds-tabs`, a second with the `.ds-search` field and the `.ds-chips`, the `.ds-grid` of six app tiles three across, a `.ds-gridfoot` with the count and the pagination, and the developers' `.ds-cta`. Things in a column are `--space-5` apart.
- Inner pages have no hero. They open with `.ds-page-header`: the same band, shorter, holding the breadcrumb, a `.ds-page-header__row` (an app's large mark, if the page is about one app, beside the title and one line of text) and at most two buttons at the right. `.ds-main` follows: `.ds-cols` (2 to 1) for an app, with the tabs on the head of the first card; `--side` (1 to 4) with a `.ds-sidebar` at the left for the workspace's own lists; `--even` (7 to 5) for a form.
- Views. The specimen is the app directory of a contacts and deals product, with four views. `home` is the directory. `app` is one app: its description under tabs and the table of what it syncs, beside the open install dialog, its details and labels, and related links. `installed` lists the workspace's apps in a table under two notices, with an empty list of private apps. `build` is the form that registers an app, with its buttons, beside the developer guides and a code sample.
- Below 960px the hero, the columns and the code card stack, the categories wrap into rows and tiles stand two across; below 1040px `.ds-nav__links` is not shown and the menu `.ds-nav__menu` sits at the end of the bar after the brand, the sign-in link and the one `.ds-button--inverse`; the menu lists every view and marks the one showing. `.ds-nav__tag` is hidden below 860px; below 640px `.ds-nav__signin` and the `.ds-nav__long` part of the button are hidden too, tiles are one column, the search field takes a full row and the figures stack their labels. A wide table scrolls inside `.ds-table__wrap`; the page never scrolls sideways.
- Spacing scale: `--space-1` 4px (label to field, chip padding), `--space-2` 8px (between chips and badges), `--space-3` 12px (mark to name, button padding, toolbar gaps), `--space-4` 16px (hero tiles, list rows, between form rows), `--space-5` 24px (card padding, grid gap, page gutter, figures), `--space-6` 32px (column gap, top of the body, the code card), `--space-7` 52px (hero padding, footer), `--space-8` 84px (between the hero's columns, foot of a view).
- Other sizes: `--size-mark` 32px, `--size-logo` 52px and `--size-logo-large` 76px (app marks), `--size-icon` 44px, `--size-control` 36px (pagination), `--size-search` 280px, `--size-dialog` 420px.

## Typography and colour roles

The conventions are the era's, as its first layout sets them out.

- The gradient is `--fill-inverse` (from `--color-fill-1` to `--color-fill-2`); text on it is `--color-inverse-text`, here charcoal on pastel. It fills the band, the developers' card and one app mark (`.ds-logo--mix`). Only `.ds-button--inverse` (fill `--color-inverse-text`, text `--color-inverse`) and `.ds-button--ghost` stand on it. The tag, the rules between the figures and the edge of the current link are `--color-inverse-text`, mixed down to 28 to 40% where they are hairlines.
- `--color-fill-1` to `--color-fill-4` are the app marks. A fill carries initials or an icon in `--color-inverse-text`, never running text.
- `--color-page` is the page, `--color-surface` (through `--fill-panel`) the tiles, cards, chips and the code block, `--color-surface-alt` the dialog foot, a hovered table row and the empty state, `--color-surface-strong` the current category, the quiet badge, inline code and the bars of the hero tiles.
- `--fill-bar` with `--color-bar-text` is the table head. `--fill-bar-alt` with `--color-bar-alt-text` is the footer.
- Text is `--color-text` (`#5a5568`), meta and counts `--color-text-muted`, headings `--color-heading` (`#2c2a36`); sidebar titles and h3 in prose are `--color-heading-alt`.
- `--color-button` (charcoal `#2c2a36`, white text) is the one action colour off the gradient: the primary button, the edge under the current tab and the current page number. The secondary button is `--color-button-secondary-text` on `--fill-button-secondary` with a `--color-border` edge; the destructive button is `--color-danger` with `--color-button-text`.
- A chosen filter chip is tinted with `--color-accent` (14% in `--color-surface`, its text mixed with `--color-heading`) and edged in it. `--color-success` is the tick of a permission and the status figure in a code sample.
- Links are `--color-link`, bold and underlined (`--link-decoration`); quiet links, titles and navigation are not underlined.
- Badges and notices follow the era's rules: `--fill-accent` for the default badge, tints mixed into `--color-surface` for `--alt` and the status variants, `--color-notice` with `--color-notice-text`, the error notice in `--color-danger` on `--color-danger-surface`.
- A serif display face for titles, figures and app marks (`--font-heading`: DM Serif Display, Playfair Display, Georgia) with a humanist sans for text and controls (`--font-body`, `--font-ui`: Karla, Gill Sans); `--font-mono` for code. `--text-base` 17px at `--line-body` 1.58; `--text-ui` 16px (controls, tables, card text); `--text-small` 14px (meta, badges, chips, code); `--text-large` 20px (leads); `--text-display` 58px (the claim); `--text-h1` 42px (page titles, figures); `--text-h2` 30px; `--text-h3` 21px (app names, card titles).
- Headings are regular weight (`--weight-heading`, `--weight-display` 400); controls and labels are bold (`--weight-ui`, `--weight-bold` 700). Nothing is uppercase.
- Surface: soft and round. `--radius-control` 12px, `--radius-panel` 20px (cards, tiles, app marks), `--radius-pill` 100px (badges, chips), `--radius-page` 36px (the developers' card, the dialog's backdrop). Cards have a hairline in `--color-border-muted` and `--shadow-panel`, one low tinted shadow; a hovered tile, the hero tiles and the code block take `--shadow-dialog`. `--border-width-strong` 2px is the edge under the current navigation link and tab and the rim of a large app mark.
- One-off values made with `calc()` and `color-mix()` because the vocabulary has no token: hairlines on the gradient (`--color-inverse-text` at 28 to 40%, the ghost border at 55%), the chosen chip, tinted badges, the inset of the search field's text, the footer rule (`--color-bar-alt-text` at 22%). Opacity 0.85 to 0.9 dims secondary text on the gradient.

## Components

- `.ds-page`: on `<body>`. Sets font, text colour and `--fill-page`.
- `.ds-wrap`: the centred content column inside a band.
- `.ds-brand`: the site's mark and name, first in the navigation and again in the footer. `.ds-brand__mark` is a tile dressed like the primary button holding the mark as inline SVG; `.ds-brand__name` is the name. Name and mark are placeholders for the installing project's own name and logo.
- `.ds-nav`: the transparent row on the band. `.ds-nav__inner` centres it; `.ds-nav__tag` names the section; `.ds-nav__links` holds `.ds-nav__link`, the current one edged (`is-current`); `.ds-nav__actions` holds the sign-in link (`.ds-nav__signin`) and one `.ds-button--inverse ds-button--small`. In the specimen the showing view marks its link with one `:has()` rule per view; a project sets `is-current` itself.
- `.ds-logo`: an app's mark, two letters on `--color-fill-1` (`--2`, `--3`, `--4`, `--mix` for the gradient; `--round` for a circle; `--large` on an app's page; `--card` adds a rim and a shadow on the band).
- `.ds-hero`: the band of the home view. `.ds-hero__inner` is two columns: the `.ds-hero__title`, `.ds-hero__lead` and `.ds-hero__actions` (one `--inverse`, one `--ghost` button), and `.ds-hero__art`, six `.ds-hero__app` tiles each a `.ds-logo` over `.ds-hero__line` bars (`--short`). One per site.
- `.ds-stat`: the row of three figures along the foot of the band, parted by hairlines; each `.ds-stat__item` sets a `.ds-stat__figure` beside its `.ds-stat__label`.
- `.ds-page-header`: head of an inner page, used instead of `.ds-hero`: the same band, shorter. `.ds-page-header__inner` holds the breadcrumb, a `.ds-page-header__row` with an optional large mark, `.ds-page-header__title` and `.ds-page-header__text`, and `.ds-page-header__actions`.
- `.ds-main`: the body of a view. `.ds-cols`: main and side column, 2 to 1; `--side` 1 to 4, `--even` 7 to 5. `.ds-stack` stacks things 24px apart.
- `.ds-toolbar`: a row above the grid; `.ds-toolbar__count` is a small count in it.
- `.ds-tabs`: a row of `.ds-tabs__tab` over a hairline; the current one (`is-current`) has an edge in `--color-button`. On an app's page it sits on the head of the first card, inside `.ds-panel__tabs`.
- `.ds-search`: the search field: a `.ds-search__icon` laid over a `.ds-form__input`.
- `.ds-chips`: filter chips, `.ds-chip` pills; a chosen one (`is-current`) is tinted and ticked.
- `.ds-grid`: app tiles, three across. Each `.ds-grid__cell` is a link: `.ds-grid__head` (a `.ds-logo`, `.ds-grid__title`, `.ds-grid__kind`), `.ds-grid__text` and `.ds-grid__foot` with a badge and the installs. `is-hover` deepens the shadow. `.ds-gridfoot` under the grid holds the count and the pagination.
- `.ds-cta`: the developers' card on the gradient: `.ds-cta__title`, `.ds-cta__text`, the two gradient buttons and a `.ds-code` block.
- `.ds-code`: a few lines of code on a card: `.ds-code__line`, with `.ds-code__key` for the method and `.ds-code__ok` for a good status.
- `.ds-facts`: pairs of `.ds-facts__term` and `.ds-facts__value` in `.ds-facts__row`, for an app's details.
- `.ds-points`: ticked `.ds-points__item` lines, for what an app may do.
- `.ds-prose`: running text: h1, h2, h3, p, lists, `code`, `strong`. `.ds-prose__lead` for the first paragraph.
- `.ds-link`: text link. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `--quiet` for meta; `--more` appends an arrow. `.ds-linkrow` lays links in a row.
- `.ds-button`: 16px bold, softly rounded. Variants: `--secondary`, `--danger` (destructive actions only), `--inverse` and `--ghost` (on the gradient only), `--small`, `--wide`. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttonrow` spaces buttons.
- `.ds-form`: vertical form with labels above the fields. `.ds-form__row` puts two fields side by side; `.ds-form__field` wraps `.ds-form__label` and one of `.ds-form__input`, `.ds-form__select`, `.ds-form__textarea`. Error: `is-error` on the input plus `.ds-form__error`. `.ds-form__hint`, `.ds-form__check` with `.ds-form__checkbox`, `.ds-form__actions`.
- `.ds-table`: data table inside a card: a small bold head on `--fill-bar`, hairlines between rows; `is-hover` tints a row. `.ds-table__num` right-aligns numbers; `.ds-table__title` is a row's link. Wrap it in `.ds-table__wrap`; `.ds-table__foot` is the foot of the card.
- `.ds-list`: rows parted by hairlines. `.ds-list__item` holds a `.ds-list__title` (alone or in a `.ds-list__head` with a badge at the right), an optional `.ds-list__text` and `.ds-list__meta`.
- `.ds-panel`: the white card. `.ds-panel__head` holds the `.ds-panel__title` and an optional badge or link above a hairline; `.ds-panel__body` is padded; `.ds-panel__text` is a small paragraph; `--side` pads a card that holds a sidebar block.
- `.ds-badge`: pill tag on `--fill-accent`. `--alt` tinted, `--quiet` grey. Status variants: `--success`, `--warning`, `--danger`. `.ds-badgerow` for several.
- `.ds-sidebar`: the categories: `.ds-sidebar__title` and a `.ds-sidebar__list` of `.ds-sidebar__item` with an optional `.ds-sidebar__count`; the current row (`is-current`) is filled.
- `.ds-notice`: tinted message; `--error` in the danger colours. `.ds-notice__title` for the bold lead-in. `.ds-noticerow` stacks notices.
- `.ds-pagination`: `.ds-pagination__link`; `is-current` is filled with `--fill-button`, `is-disabled` is grey.
- `.ds-breadcrumb`: a small trail of `.ds-breadcrumb__item` parted by slashes; the last is `.ds-breadcrumb__current`. It takes the colour of the band.
- `.ds-dialog`: a dimmed block (`--color-overlay`) centring `.ds-dialog__box`, a card in `--shadow-dialog`: `.ds-dialog__head` (`.ds-dialog__title`, `.ds-dialog__close`), `.ds-dialog__body`, and `.ds-dialog__actions` right-aligned on `--color-surface-alt`. The install dialog lists permissions as `.ds-points`.
- `.ds-empty`: a centred message on `--color-surface-alt`: `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one small secondary button.
- `.ds-tile`: an icon tile on a fill (`--2`, `--3`, `--4`, `--round`), for marks that are icons, not initials.
- `.ds-footer`: the closing band on `--fill-bar-alt`: `.ds-footer__cols` with the brand and a `.ds-footer__about` line, then columns of `.ds-footer__title` and `.ds-footer__list` (two across, `--one` for a single column) of `.ds-footer__link`; `.ds-footer__legal` under a rule.
- `.ds-menu`: a drop-down, a `<details>` whose `.ds-menu__button` is the summary and whose `.ds-menu__list` of `.ds-menu__link` opens as a small card under it. Use it for a group of links under one bar item and for the account.
- `.ds-nav__menu`: the `.ds-menu` that replaces the bar's links on a phone. It lists every view, marks the one showing, and sits last in the bar after the brand and the one action that stays.
- `.ds-switch`: an on or off setting drawn over a checkbox (`.ds-switch__input`, `.ds-switch__track`, then the label). Use it for a setting that takes effect at once, not for a choice sent with a form's other fields.
- `.ds-tooltip`: a short hint on hover or focus; `.ds-tooltip--mark` is the small round question mark beside a label, `.ds-tooltip__text` the hint, `.ds-tooltip--start` opens it from the left edge. One sentence at most.
- `.ds-progress`: a thin pill bar with `.ds-progress__bar` (`--20` to `--100` set the width). `.ds-meter` puts a name and a figure above it. Use it for steps done and for usage against a limit.
- `.ds-avatar`: a person as initials on a round fill (`--1`, `--2`, `--4` change the fill); `.ds-who` sets it beside a name. Use it wherever a named person appears.

## Never

- `text-shadow = none`: text is flat, also on the gradient.
- `animation = none`: nothing moves by itself; only hover states ease.
- `border-width <= 2px`: borders are hairlines; 2px marks the current link and tab.
- `border-radius <= 36px`: cards are 20px and the largest blocks 36px; only pills and circles are rounder.
- `box-shadow-blur <= 90px`: shadows are low and soft but end at 90px.
- `font-weight <= 700`: labels are bold, never black; headings are regular.
- `font-size >= 14px`: meta, badges, chips and code are the smallest text.
- `font-size <= 58px`: the claim is the largest text.
- `font-families <= 3`: a serif display face, a body sans and monospace for code.
- `line-height <= 1.58`: body text is set at 1.58.
- `uppercase-text <= 0%`: navigation, buttons and labels are in sentence case.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new app is another `.ds-grid__cell` with a `.ds-logo` of two letters on one of the fills; a new filter is a `.ds-chip`; a new list of the workspace is a view with the `.ds-sidebar` at the left and a table in a `.ds-panel` at the right. Anything that asks for consent is a `.ds-dialog` with its permissions as `.ds-points`. Text goes on cards, on the page and, in `--color-inverse-text`, on the gradient; the four fills take initials or an icon only, and only the `--inverse` and `--ghost` buttons stand on the gradient. Keep corners, shadows and borders on their tokens, spacing on the `--space-*` scale, and note any one-off `color-mix()` shade next to its rule.
