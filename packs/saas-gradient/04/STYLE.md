# SaaS gradient blog and customer stories

## Summary

This is the company blog of a start-up product of 2016 to 2021, where customer stories and posts are told as cards: a plain bar of navigation, then the featured story as one large card of two-colour gradient with a small drawn scene of white cards, and below it a grid of story cards, each under a block of one flat colour with a line icon. Inner pages are quiet: a band in the card colour with a small gradient tile beside the title, an author strip between two hairlines, a pull quote on the gradient and text in a comfortable column beside small cards. It was the usual way for payment, analytics and support products to show who used them around 2018.

## Layout

- The page is designed for a 1440px viewport and is not fluid: content sits in `.ds-wrap`, centred, at most `--size-page` (1120px) wide with `--space-5` side padding. The navigation bar, the page header and the footer run the full width of the window.
- Navigation is a solid bar on `--fill-panel`, at least `--size-nav` (72px) tall, over a hairline: the `.ds-brand`, three links whose current one carries a `--border-width-strong` edge along the foot of the bar, then one quiet link and one small primary button at the right. Nothing of the navigation lies on the gradient.
- The home view stacks inside `.ds-home`: `.ds-hero`, the featured story as a gradient card in two columns (6 to 5); `.ds-stat`, four figures set straight on the page; a `.ds-block__head` with the title at the left and `.ds-tabs` as filter chips at the right; `.ds-grid`, six story cards three across; and a `.ds-block` of `.ds-cols` (2 to 1, `--level`) with the latest posts in a card beside the newsletter `.ds-cta`. Blocks are `--space-7` apart.
- Inner pages have no hero. They open with `.ds-page-header`: a band in the card colour over a hairline that holds the breadcrumb, then a `.ds-page-header__tile` (a square of gradient with an icon) beside the page title and one line of text, and at most one button at the right. `.ds-main` follows with `.ds-cols` (2 to 1) for an article, `--side` (1 to 3) for the archive with the `.ds-sidebar` at the left and the chips above the list, or `--even` (7 to 5) for a form.
- Views. The specimen is the blog of a shared inbox product, with four views. `home` is the customer stories page. `story` is one story: author strip, pull quote, the text and a table of results, beside the customer's facts, the features used, a newsletter card, more by the author and related links. `archive` lists all posts under chips, with topics and writers at the left, the pagination in the foot of the card and an empty topic. `subscribe` is the newsletter form with its notices and buttons beside the last issue and the open unsubscribe dialog.
- Below 960px the hero, the columns and the archive stack and cards stand two across; below 1040px `.ds-nav__links` is not shown and the menu `.ds-nav__menu` sits at the end of the bar after the brand, the quiet link and the small button, listing every view and marking the one showing; below 640px `.ds-nav__signin` is hidden too, the grid is one column, the gradient tile of the page header is dropped and a wide table scrolls inside `.ds-table__wrap`. The page never scrolls sideways.
- Spacing scale: `--space-1` 4px (label to field, chip padding), `--space-2` 8px (between chips and badges, kicker to title), `--space-3` 12px (avatar to name, button padding), `--space-4` 16px (list rows, author strip, drawn cards), `--space-5` 24px (card padding, grid gap, page gutter), `--space-6` 32px (column gap, page header padding), `--space-7` 48px (between blocks, hero padding), `--space-8` 80px (foot of a view).
- Other sizes: `--size-mark` 34px, `--size-icon` 44px (one and a half times that for the icon of a card and the page header tile), `--size-art` 168px (height of a card's colour block), `--size-avatar` 40px, `--size-control` 38px (pagination), `--size-dialog` 460px.

## Typography and colour roles

The conventions are the era's, as its first layout sets them out.

- The gradient is `--fill-inverse` (from `--color-fill-1` to `--color-fill-2`); text on it is `--color-inverse-text`. It fills the hero card, the newsletter card, the pull quote, the page header tile, the meter in the hero scene and one card block (`.ds-grid__art--mix`). Only `.ds-button--inverse` and `.ds-button--ghost` stand on it.
- `--color-fill-1` to `--color-fill-4` are the colour blocks of the story cards, the avatars and the left edges of the figures. A fill carries an icon or initials in `--color-inverse-text`, never running text; the circles on a block are `--color-inverse-text` at 16%.
- `--color-page` is the page, `--color-surface` (through `--fill-panel`) the cards, the navigation bar and the page header, `--color-surface-alt` the dialog foot and the empty state, `--color-surface-strong` the current sidebar row, the quiet badge, inline code and the bars of the hero scene.
- `--fill-bar` with `--color-bar-text` is the table head only. `--fill-bar-alt` with `--color-bar-alt-text` is the footer.
- Text is `--color-text` (`#4a5f61`), meta `--color-text-muted`, headings and figures `--color-heading` (`#03363d`); the kicker of a card, sidebar titles and h3 in prose are `--color-heading-alt`.
- Links are `--color-link`, medium weight, underlined only on hover. Navigation links are `--color-link-quiet` and turn `--color-heading` when current.
- `--color-button` (dark forest `#03363d`, white text) is the one action colour off the gradient: the primary button, the current chip, the current page number and the edge under the current navigation link. The secondary button is `--color-button-secondary-text` on `--fill-button-secondary` with a `--color-border` edge; the destructive button is `--color-danger` with `--color-button-text`.
- Badges, notices and the tinted status colours follow the era's rules: `--fill-accent` for the default badge, tints mixed into `--color-surface` for `--alt` and the status variants, `--color-notice` with `--color-notice-text`, the error notice in `--color-danger` on `--color-danger-surface`.
- Two families: `--font-heading` (Work Sans, Helvetica Neue) set light for titles and figures, `--font-body` and `--font-ui` (Source Sans Pro, Helvetica Neue); the references used commercial grotesques. `--font-mono` for code. `--text-base` 18px at `--line-body` 1.7; `--text-ui` 16px (controls, tables, card text); `--text-small` 14px (meta, badges, bylines, chips); `--text-large` 22px (leads); `--text-display` 54px (the featured title); `--text-h1` 40px (page titles, figures); `--text-h2` 28px (block titles, the pull quote); `--text-h3` 20px (card titles).
- Headings are light: `--weight-heading` and `--weight-display` 300. `--weight-bold` is 600 and `--weight-ui` 500. Nothing is uppercase.
- Surface: everything is nearly square. `--radius-control` and `--radius-panel` are 4px, `--radius-page` 6px (hero card, pull quote); only badges, chips and avatars are round (`--radius-pill`). Cards carry a hairline in `--color-border-muted` and `--shadow-panel`, a small close shadow; a hovered card and the cards of the hero scene take `--shadow-dialog`. `--border-width-strong` 3px is the edge of the current navigation link and of each figure.
- One-off values made with `calc()` and `color-mix()` because the vocabulary has no token: washes on the gradient (`--color-inverse-text` at 14 to 24%, the ghost border at 55%), tinted badges, the icon size of a card (`--size-icon` times 1.5), the footer rule (`--color-bar-alt-text` at 22%). Opacity 0.8 to 0.9 dims secondary text on the gradient.

## Components

- `.ds-page`: on `<body>`. Sets font, text colour and `--fill-page`.
- `.ds-wrap`: the centred content column.
- `.ds-brand`: the site's mark and name, first in the navigation and again in the footer. `.ds-brand__mark` is a tile dressed like the primary button holding the mark as inline SVG; `.ds-brand__name` is the name. Name and mark are placeholders for the installing project's own name and logo.
- `.ds-nav`: the solid bar. `.ds-nav__inner` centres it; `.ds-nav__links` holds `.ds-nav__link`, the current one edged (`is-current`); `.ds-nav__actions` holds a quiet link (`.ds-nav__signin`) and one small primary button. In the specimen the showing view marks its link with one `:has()` rule per view; a project sets `is-current` itself.
- `.ds-hero`: the featured story, one gradient card per site. `.ds-hero__copy` holds `.ds-hero__kicker`, `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__actions` (one `--inverse`, one `--ghost` button) and a `.ds-byline`. `.ds-hero__art` is a drawn scene: `.ds-hero__ring` circles (`--small`) and two or three `.ds-hero__card` (`--in`, `--out` to stagger them) holding an avatar, a `.ds-hero__grow` label with `.ds-hero__line` bars (`--short`), a badge, a `.ds-hero__figure` or a `.ds-hero__meter` with its `.ds-hero__meter-bar`.
- `.ds-avatar`: a circle of initials on a fill (`--1`, `--2`, `--4`; `--plain` on the gradient).
- `.ds-byline`: who wrote it: a `.ds-avatar`, `.ds-byline__name` and `.ds-byline__meta`. `--strip` is the author strip of an article, between two hairlines, with `.ds-byline__who` at the left and tags at the right.
- `.ds-stat`: the row of four figures; each `.ds-stat__item` has a left edge in one of the fills, a `.ds-stat__figure` and a `.ds-stat__label`.
- `.ds-home`, `.ds-main`: the body of the home view and of an inner view. `.ds-block` is a further block of the home view; `.ds-block__head` puts a `.ds-block__title` at the left and chips at the right.
- `.ds-cols`: main and side column, 2 to 1; `--even` 7 to 5, `--side` 1 to 3, `--level` for equal height. `.ds-stack` stacks cards 24px apart.
- `.ds-tabs`: filter chips: `.ds-tabs__tab` pills with a hairline, the current one (`is-current`) filled with `--fill-button`.
- `.ds-grid`: story cards, three across. Each `.ds-grid__cell` is a link: `.ds-grid__art` (a block of one fill with a line icon; `--2`, `--3`, `--4`, `--mix` for the gradient) over `.ds-grid__body` with a `.ds-kicker`, `.ds-grid__title`, `.ds-grid__text` and a `.ds-byline`. `is-hover` deepens the shadow.
- `.ds-page-header`: head of an inner page, used instead of `.ds-hero`. `.ds-page-header__inner` holds the breadcrumb, a `.ds-page-header__row` of the `.ds-page-header__tile`, `.ds-page-header__title` and `.ds-page-header__text`, and `.ds-page-header__actions` at the right.
- `.ds-cover`: the pull quote of an article on the gradient: `.ds-cover__text` and `.ds-cover__by`.
- `.ds-facts`: pairs of `.ds-facts__term` and `.ds-facts__value` in `.ds-facts__row`, for a customer's details.
- `.ds-cta`: the newsletter card on the gradient: `.ds-cta__title`, `.ds-cta__text` and the two gradient buttons.
- `.ds-prose`: running text: h1, h2, h3, p, lists, `code`, `strong`. `.ds-prose__lead` for the first paragraph.
- `.ds-link`: text link. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `--quiet` for meta; `--more` appends an arrow. `.ds-linkrow` lays links in a row.
- `.ds-button`: 16px medium, nearly square. Variants: `--secondary`, `--danger` (destructive actions only), `--inverse` and `--ghost` (on the gradient only), `--small`, `--wide`. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttonrow` spaces buttons.
- `.ds-form`: vertical form with labels above the fields. `.ds-form__row` puts two fields side by side; `.ds-form__field` wraps `.ds-form__label` and one of `.ds-form__input`, `.ds-form__select`, `.ds-form__textarea`. Error: `is-error` on the input plus `.ds-form__error`. `.ds-form__hint`, `.ds-form__check` with `.ds-form__checkbox`, `.ds-form__actions`.
- `.ds-table`: data table inside a card: a small bold head on `--fill-bar`, hairlines between rows; `is-hover` tints a row. `.ds-table__num` right-aligns numbers; `.ds-table__title` is a row's link. Wrap it in `.ds-table__wrap`; `.ds-table__foot` is the foot of a card for a count and the pagination.
- `.ds-list`: posts parted by hairlines. `.ds-list__item` holds a `.ds-list__title` (alone or in a `.ds-list__head` with a badge at the right), an optional `.ds-list__text` and `.ds-list__meta`.
- `.ds-panel`: the white card. `.ds-panel__head` holds the `.ds-panel__title` and an optional badge or link above a hairline; `.ds-panel__body` is padded; `.ds-panel__text` is a small paragraph; `--side` pads a card that holds a sidebar block.
- `.ds-badge`: pill tag on `--fill-accent`. `--alt` tinted, `--quiet` grey. Status variants: `--success`, `--warning`, `--danger`. `.ds-badgerow` for several.
- `.ds-sidebar`: a block of links: `.ds-sidebar__title` and a `.ds-sidebar__list` of `.ds-sidebar__item` with an optional `.ds-sidebar__count`; the current row (`is-current`) is grey.
- `.ds-notice`: tinted message; `--error` in the danger colours. `.ds-notice__title` for the bold lead-in. `.ds-noticerow` stacks notices.
- `.ds-pagination`: `.ds-pagination__link` squares; `is-current` is filled with `--fill-button`, `is-disabled` is grey.
- `.ds-breadcrumb`: a small trail of `.ds-breadcrumb__item` parted by slashes; the last is `.ds-breadcrumb__current`.
- `.ds-dialog`: a dimmed block (`--color-overlay`) centring `.ds-dialog__box`, a card in `--shadow-dialog`: `.ds-dialog__head` (`.ds-dialog__title`, `.ds-dialog__close`), `.ds-dialog__body`, and `.ds-dialog__actions` right-aligned on `--color-surface-alt`.
- `.ds-empty`: a centred message on `--color-surface-alt`: `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one small secondary button.
- `.ds-tile`: an icon tile on a fill (`--2`, `--3`, `--4`, `--round`), for small marks inside cards.
- `.ds-footer`: the closing band on `--fill-bar-alt`: `.ds-footer__inner` with the brand and a `.ds-footer__about` line at the left and `.ds-footer__links` of `.ds-footer__link` at the right; `.ds-footer__legal` under a rule.
- `.ds-menu`: a drop-down, a `<details>` whose `.ds-menu__button` is the summary and whose `.ds-menu__list` of `.ds-menu__link` opens as a small card under it. Use it for a group of links under one bar item and for the account.
- `.ds-nav__menu`: the `.ds-menu` that replaces the bar's links on a phone. It lists every view, marks the one showing, and sits last in the bar after the brand and the one action that stays.
- `.ds-switch`: an on or off setting drawn over a checkbox (`.ds-switch__input`, `.ds-switch__track`, then the label). Use it for a setting that takes effect at once, not for a choice sent with a form's other fields.
- `.ds-tooltip`: a short hint on hover or focus; `.ds-tooltip--mark` is the small round question mark beside a label, `.ds-tooltip__text` the hint, `.ds-tooltip--start` opens it from the left edge. One sentence at most.
- `.ds-progress`: a thin pill bar with `.ds-progress__bar` (`--20` to `--100` set the width). `.ds-meter` puts a name and a figure above it. Use it for steps done and for usage against a limit.

## Never

- `text-shadow = none`: text is flat, also on the gradient.
- `animation = none`: nothing moves by itself; only hover states ease.
- `border-width <= 3px`: borders are hairlines; 3px marks the current navigation link and the figures.
- `border-radius <= 6px`: controls and cards are 4px and the largest blocks 6px; only pills and circles are rounder.
- `box-shadow-blur <= 40px`: shadows are small and close; the deepest ends at 40px.
- `font-weight <= 600`: headings are light and labels semibold, never heavy.
- `font-size >= 14px`: meta, badges and bylines are the smallest text.
- `font-size <= 54px`: the featured title is the largest text.
- `font-families <= 3`: a heading face, a body sans and monospace for code.
- `line-height <= 1.7`: body text is set at 1.7.
- `uppercase-text <= 0%`: navigation, buttons and labels are in sentence case.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new kind of post is another `.ds-grid__cell` with its own fill and icon; a new part of an article is a card (`.ds-panel`) in the side column or a block in the text column; anything that lists posts is a `.ds-list` in a card with the pagination in a `.ds-table__foot`. Keep the gradient for the featured story, the pull quote, the newsletter card and the page header tile, with text in `--color-inverse-text` and only the `--inverse` and `--ghost` buttons on it; the four fills take an icon or initials only. Keep corners, shadows and borders on their tokens, spacing on the `--space-*` scale, and note any one-off `color-mix()` shade next to its rule.
