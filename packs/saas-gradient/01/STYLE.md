# SaaS gradient landing page

## Summary

This is the start-up product page of 2016 to 2021: a wide band of two-colour gradient with a slanted lower edge carries the navigation, the claim and a product screen drawn as white cards, and everything below it is white rounded cards with soft, wide, tinted shadows on a very pale page. Type is one friendly geometric sans at a generous size, buttons are pills, links are never underlined, and small flat illustrations are built from tinted blocks, circles and card shapes. It was the common look of payment, analytics and messaging products around 2018.

## Layout

- The page is designed for a 1440px viewport and is not fluid: content sits in `.ds-wrap`, centred, at most `--size-page` (1100px) wide with `--space-5` side padding. Bands (the hero, `.ds-section--alt`, the footer) run the full width of the window.
- Navigation is one transparent row, `--size-nav` (76px) tall, lying on the gradient: the `.ds-brand`, three links, then "Sign in" and one small white button at the right. The hero or the page header is pulled up under it by a negative margin of the same height, so the gradient starts at the top of the window.
- The home view stacks: `.ds-hero` (copy at the left, the drawn product screen at the right, lower edge slanted by 8vw with `clip-path`); `.ds-stat`, one white card of four figures that overlaps the slant; the logo row as text; a white band with two `.ds-feature` rows whose sides alternate; a `.ds-grid` of four cards; a white band of three `.ds-quote` cards; and `.ds-cta`, a gradient card. Sections are `--space-8` (96px) apart, tight ones `--space-7`.
- Inner pages have no hero. They open with `.ds-page-header`: the same gradient band with a 4vw slant, holding the breadcrumb, the page title, one line of text and at most one white button at the right. `--deep` makes the band taller so that the first cards (`.ds-section--lift`) climb onto it. Below, content is either full width (the three plans) or `.ds-cols` (7 to 5), or `.ds-cols--side` (1 to 3) with a `.ds-sidebar` at the left. Tabs are a centred pill switch above the cards they change.
- Views. The specimen is the site of an invoicing product. `home` is the landing page. `pricing` has the switch between monthly and yearly, three plan cards, the comparison table with a notice under it, and the common questions as a list. `guide` is one page of the getting-started guide: its contents and related links at the left, the text, the label key, an empty questions box and the pagination. `signup` is the trial form with its notices and buttons beside what happens next and the open leave dialog.
- Below 960px the columns, feature rows, plans and quotes stack and cards stand two across; below 900px `.ds-nav__links` is not shown and the menu `.ds-nav__menu` takes its place at the end of the bar, after the brand, "Sign in" and the one `.ds-button--inverse`; the menu lists every view and marks the one showing. Below 640px `.ds-nav__signin` is hidden too, grids go to one column and a wide table scrolls inside `.ds-table__wrap`. The page never scrolls sideways.
- Spacing scale: `--space-1` 4px (label to field, pagination gaps), `--space-2` 8px (icon to text, between badges), `--space-3` 12px (button and cell vertical padding, between buttons), `--space-4` 16px (between form rows, list row padding, card head), `--space-5` 24px (card padding, grid gap, page gutter), `--space-6` 32px (column gap, hero actions), `--space-7` 56px (tight sections, hero top), `--space-8` 96px (between sections, between feature rows).
- Other sizes: `--size-mark` 32px, `--size-icon` 44px (icon tiles), `--size-control` 36px (pagination), `--size-avatar` 40px, `--size-art` 300px (least height of an illustration block), `--size-dialog` 440px.

## Typography and colour roles

These are the conventions of the era; every layout and token set keeps them.

- The gradient is `--fill-inverse`, made in the surface part from `--color-fill-1` and `--color-fill-2`; only its angle and stops change between surface sets. Text on it is always `--color-inverse-text`, here white. `--color-inverse` is the solid middle of the gradient: the text of a white button on the band.
- On the gradient there are two buttons only: `.ds-button--inverse` (fill `--color-inverse-text`, text `--color-inverse`) and `.ds-button--ghost` (outline in `--color-inverse-text`). The primary button never stands on the gradient, so `--color-button` may be any colour, close to the gradient or far from it.
- `--color-fill-1` to `--color-fill-4` are saturated mid tones: icon tiles, avatars, sparklines, chart lines and illustration shapes. A fill carries an icon or initials in `--color-inverse-text`, never running text. Illustration blocks are a fill mixed at about 20% into `--color-surface`.
- `--color-page` is the pale page, `--color-surface` the card, `--color-surface-alt` the quiet inset (dialog foot, hovered row, empty state), `--color-surface-strong` the strongest neutral (quiet badge, inline code, current sidebar row, progress track).
- `--fill-bar` with `--color-bar-text` is the table head only. `--fill-bar-alt` with `--color-bar-alt-text` is the footer, which a palette may make pale or dark; nothing else stands on it.
- Text is `--color-text` (`#525f7f`), meta `--color-text-muted`, headings and figures `--color-heading` (`#32325d`); kickers, sidebar titles and h3 in prose are `--color-heading-alt`.
- Links are `--color-link`, medium weight, never underlined (`--link-decoration` none); quiet links are `--color-link-quiet`.
- `--color-button` (coral `#ff5a6e`, white text) is the one action colour on cards: the primary button, the current pill of the switch, the current page number, the top edge of the chosen plan. The secondary button is `--color-button-secondary-text` on `--fill-button-secondary` with a `--color-border` edge. The destructive button is `--color-danger` with `--color-button-text`.
- `--fill-accent` with `--color-accent-text` is the default badge. `--alt` and the status badges are tints: the colour mixed at 16 to 22% into `--color-surface`, with text mixed from the colour and `--color-heading`, so they work on any palette. `--color-success` is also the tick of a feature point.
- `--color-notice` with `--color-notice-text` is the information notice; the error notice is `--color-danger` on `--color-danger-surface` with a faint edge.
- One family for body, headings and controls (`--font-body`: Nunito Sans, Avenir Next; the references used commercial geometric sans faces), `--font-mono` for code. `--text-base` 17px at `--line-body` 1.65; `--text-ui` 15px (controls, tables, card text); `--text-small` 13px (meta, badges, labels, breadcrumb); `--text-large` 21px (leads); `--text-display` 46px (hero title, plan price); `--text-h1` 36px (section and page titles, figures); `--text-h2` 26px; `--text-h3` 18px (card titles).
- Headings are medium, not bold: `--weight-heading` and `--weight-display` 500. `--weight-bold` and `--weight-ui` are 600. Nothing is uppercase.
- Surface: buttons, inputs and badges are pills (`--radius-control`, `--radius-pill` 100px); cards are `--radius-panel` 8px; the largest containers `--radius-page` 12px. `--border-width` 1px in `--color-border-muted` is almost invisible: cards are set off by `--shadow-panel`, two wide soft layers in the tinted `--color-shadow`. `--shadow-control` is a small lift that grows on hover while the button rises by 1px; `--shadow-dialog` is the deepest and is also used under the drawn product screen. `--transition` is 0.15s.
- One-off values made with `calc()` and `color-mix()` because the vocabulary has no token: the slants (8vw and 4vw), the washes on the gradient (`--color-inverse-text` at 14 to 18%, the ghost border at 55%), tinted badges and illustration blocks, the brand tile's edge (`--color-button-text` at 35%), the error notice's edge (`--color-danger` at 35%), the footer rule (`--color-bar-alt-text` at 22%). Opacity 0.75 to 0.9 dims secondary text on the gradient.

## Components

- `.ds-page`: on `<body>`. Sets font, text colour and `--fill-page`.
- `.ds-wrap`: the centred content column inside a band.
- `.ds-brand`: the site's mark and name, first in the navigation and again in the footer. `.ds-brand__mark` is a 32px tile dressed like the primary button (`--fill-button`, `--color-button-text`, `--radius-control`, `--shadow-control`) holding the mark as inline SVG; `.ds-brand__name` is the name in bold. Name and mark are placeholders for the installing project's own name and logo.
- `.ds-nav`: the transparent row on the gradient. `.ds-nav__inner` centres it; `.ds-nav__links` holds `.ds-nav__link`, the current one on a translucent pill (`is-current`); `.ds-nav__actions` holds the sign-in link (`.ds-nav__signin`) and one `.ds-button--inverse ds-button--small`. In the specimen the showing view marks its link with one `:has()` rule per view; a project sets `is-current` itself.
- `.ds-hero`: the slanted gradient band of the home view. `.ds-hero__inner` is two columns: `.ds-hero__copy` with an optional `.ds-hero__kicker`, the `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__actions` (one `--inverse`, one `--ghost` button) and `.ds-hero__note`; and `.ds-hero__art` with a `.ds-shot`. One per site.
- `.ds-shot`: a product screen drawn as a card: `.ds-shot__bar` (three `.ds-shot__dot` and a `.ds-shot__name`), `.ds-shot__row` lines of a `.ds-tile`, `.ds-shot__who`, `.ds-shot__sum` and a badge; `.ds-shot__float` is a small card laid over its corner with `.ds-shot__label`, `.ds-shot__figure` and a `.ds-spark`.
- `.ds-spark`: a sparkline, an inline SVG path stroked with `currentColor`; `--2`, `--3`, `--4` take the other fills.
- `.ds-tile`: an icon tile on `--color-fill-1` (`--2`, `--3`, `--4`; `--round` for a circle).
- `.ds-page-header`: head of an inner page, used instead of `.ds-hero`: the gradient band with a shallower slant. `.ds-page-header__inner` holds the breadcrumb, `.ds-page-header__title`, `.ds-page-header__text` and `.ds-page-header__actions`. `--deep` when the first cards overlap it.
- `.ds-section`: a full-width block of the page; `--alt` white, `--tight` with less padding, `--lift` pulled up onto a deep page header. `.ds-section__head` centres a `.ds-section__kicker`, `.ds-section__title` and `.ds-section__lead`.
- `.ds-cols`: main and side column, 7 to 5; `--side` 1 to 3. `.ds-stack` stacks cards 24px apart.
- `.ds-stat`: the row of figures, one card of four `.ds-stat__item`, each a `.ds-stat__figure` over a `.ds-stat__label`. On the home view it overlaps the slant.
- `.ds-logos`: customer names set as text, `.ds-logos__name`, under a `.ds-logos__lead` line.
- `.ds-feature`: a feature row, `.ds-feature__copy` beside a `.ds-art`; `--flip` puts the copy at the right. Copy is a kicker, `.ds-feature__title`, `.ds-feature__text`, `.ds-feature__points` of ticked `.ds-feature__point` and a `.ds-link--more`.
- `.ds-art`: a flat illustration block, a tinted fill (`--2`, `--3`) with a `.ds-art__blob` circle and two or three `.ds-art__card` (`--in`, `--out` to stagger them) holding a tile, a `.ds-art__grow` label with `.ds-art__line` bars, and a badge or figure.
- `.ds-grid`: four cards across. Each `.ds-grid__cell` is a `.ds-tile`, `.ds-grid__title`, `.ds-grid__text` and a link.
- `.ds-quotes`: three `.ds-quote` cards: `.ds-quote__text` and `.ds-quote__by` with a `.ds-quote__avatar` (`--4`, `--plain`) and `.ds-quote__name`. `--lead` puts one card on the gradient.
- `.ds-cta`: the closing call to action, a gradient card with `.ds-cta__title`, `.ds-cta__text` and the two gradient buttons.
- `.ds-prose`: running text: h1, h2, h3, p, lists, `code`, `strong`. `.ds-prose__lead` for the first paragraph.
- `.ds-link`: text link. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `--quiet` for meta; `--more` appends an arrow. `.ds-linkrow` lays links in a row.
- `.ds-button`: pill button, 15px semibold. Variants: `--secondary`, `--danger` (destructive actions only), `--inverse` and `--ghost` (on the gradient only), `--small`, `--wide`. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttonrow` spaces buttons.
- `.ds-form`: vertical form with labels above the fields. `.ds-form__row` puts two fields side by side; `.ds-form__field` wraps `.ds-form__label` and one of `.ds-form__input`, `.ds-form__select`, `.ds-form__textarea`. Error: `is-error` on the input plus `.ds-form__error`. `.ds-form__hint`, `.ds-form__check` with `.ds-form__checkbox`, `.ds-form__actions`.
- `.ds-table`: data table inside a card: a small bold head on `--fill-bar`, hairlines between rows, no vertical lines; `is-hover` tints a row. `.ds-table__num` right-aligns numbers; `.ds-table__title` is a row's link. Wrap it in `.ds-table__wrap`; `.ds-table__foot` is a last row for pagination and a count.
- `.ds-list`: rows parted by hairlines. `.ds-list__item` holds a `.ds-list__title` (alone or in a `.ds-list__head` with a badge at the right), an optional `.ds-list__text` and `.ds-list__meta`.
- `.ds-panel`: the white card with the soft wide shadow. `.ds-panel__head` holds the `.ds-panel__title` and an optional badge or link above a hairline; `.ds-panel__body` is padded; `.ds-panel__text` is a small paragraph.
- `.ds-plans`: three `.ds-panel.ds-plan` cards of equal height: `.ds-plan__price` with `.ds-plan__per`, `.ds-plan__text`, feature points and a `--wide` button. `.ds-plan--best` takes a top edge in `--color-button` and the primary button.
- `.ds-tabs`: here a pill switch inside `.ds-tabsrow`: `.ds-tabs__tab` pills on a white capsule, the current one (`is-current`) filled with `--fill-button`.
- `.ds-badge`: pill tag, 13px semibold, on `--fill-accent`. `--alt` tinted, `--quiet` grey. Status variants: `--success`, `--warning`, `--danger`. `.ds-badgerow` for several.
- `.ds-sidebar`: a block of links: `.ds-sidebar__title` and a `.ds-sidebar__list` of `.ds-sidebar__item` with an optional `.ds-sidebar__count`; the current row (`is-current`) is a grey pill.
- `.ds-notice`: tinted rounded message; `--error` in the danger colours. `.ds-notice__title` for the bold lead-in. `.ds-noticerow` stacks notices.
- `.ds-pagination`: `.ds-pagination__link` pills; `is-current` is filled with `--fill-button`, `is-disabled` is grey.
- `.ds-breadcrumb`: a small trail of `.ds-breadcrumb__item` parted by slashes; the last is `.ds-breadcrumb__current`. It takes the colour of what it stands on and sits above the page title.
- `.ds-dialog`: a dimmed block (`--color-overlay`) centring `.ds-dialog__box`, a card in `--shadow-dialog`: `.ds-dialog__head` (`.ds-dialog__title`, `.ds-dialog__close`), `.ds-dialog__body`, and `.ds-dialog__actions` right-aligned on `--color-surface-alt`.
- `.ds-empty`: a centred message on `--color-surface-alt`: `.ds-empty__icon` (a circle of `--color-fill-3`), `.ds-empty__title`, `.ds-empty__text` and one small secondary button.
- `.ds-footer`: the closing band on `--fill-bar-alt`: `.ds-footer__cols` with the brand and a `.ds-footer__about` line, then three columns of `.ds-footer__title` and `.ds-footer__list` of `.ds-footer__link`; `.ds-footer__legal` under a rule.
- `.ds-menu`: a drop-down, a `<details>` whose `.ds-menu__button` is the summary and whose `.ds-menu__list` of `.ds-menu__link` opens as a small card under it. Use it for a group of links under one bar item and for the account.
- `.ds-nav__menu`: the `.ds-menu` that replaces the bar's links on a phone. It lists every view, marks the one showing, and sits last in the bar after the brand and the one action that stays.
- `.ds-switch`: an on or off setting drawn over a checkbox (`.ds-switch__input`, `.ds-switch__track`, then the label). Use it for a setting that takes effect at once, not for a choice sent with a form's other fields.
- `.ds-tooltip`: a short hint on hover or focus; `.ds-tooltip--mark` is the small round question mark beside a label, `.ds-tooltip__text` the hint, `.ds-tooltip--start` opens it from the left edge. One sentence at most.
- `.ds-accordion`: questions that open one by one, each a `<details class="ds-accordion__item">` with `.ds-accordion__title` and `.ds-accordion__body`. Use it for common questions inside a panel.
- `.ds-progress`: a thin pill bar with `.ds-progress__bar` (`--20` to `--100` set the width). `.ds-meter` puts a name and a figure above it. Use it for steps done and for usage against a limit.
- `.ds-avatar`: a person as initials on a round fill (`--1`, `--2`, `--4` change the fill); `.ds-who` sets it beside a name. Use it wherever a named person appears.

## Never

- `text-shadow = none`: text is flat, also on the gradient.
- `animation = none`: nothing moves by itself; only hover states ease.
- `border-width <= 2px`: borders are hairlines; 2px marks the chosen plan.
- `border-radius <= 12px`: cards are 8px and the largest blocks 12px; only pills and circles are rounder.
- `box-shadow-blur <= 100px`: shadows are wide and soft but end at 100px.
- `font-weight <= 600`: headings are medium and labels semibold, never heavy.
- `font-size >= 13px`: meta, badges and labels are the smallest text.
- `font-size <= 46px`: the hero title and the plan price are the largest text.
- `font-families <= 2`: one sans-serif family, plus monospace for code.
- `line-height <= 1.65`: body text is set at 1.65.
- `uppercase-text <= 0%`: navigation, buttons and labels are in sentence case.
- `underlined-links <= 0%`: links are told apart by colour and weight.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new block of the landing page is a `.ds-section` with a `.ds-section__head`; a new feature is another `.ds-feature` row with its side flipped and an illustration built from `.ds-art` pieces; anything that holds data is a `.ds-panel`. Text goes on cards, on the page and, in `--color-inverse-text`, on the gradient; the four fills take an icon or initials only. Put only the `--inverse` and `--ghost` buttons on the gradient, and at most one primary button in a card. Keep corners, shadows and borders on their tokens, spacing on the `--space-*` scale, and note any one-off `color-mix()` shade next to its rule.
