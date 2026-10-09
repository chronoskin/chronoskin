# SaaS gradient event page

## Summary

This is the event page of a start-up product of 2016 to 2021, made for a conference, a webinar or a launch: the gradient is one large curve that sweeps in from the top left, a white bar of navigation floats on it, and the claim, a countdown and the faces of those already coming stand beside a sign-up card that hangs over the curve's edge. Below come four tinted figures, the speakers as cards under blocks of flat colour, the day as a timed list, last year's praise as a slide on the gradient and the questions as an accordion. Products of the period ran such a page every spring, and the registration form was always in the first screen.

## Layout

- The page is designed for a 1440px viewport and is not fluid: content sits in `.ds-wrap`, centred, at most `--size-page` (1120px) wide with `--space-5` side padding. The curve, the white sections and the footer run the full width of the window.
- Navigation is a floating bar: `.ds-nav__inner` is a card on `--fill-panel` with `--radius-page` corners and the panel shadow, as wide as the content and `--space-4` below the top of the window. It holds the `.ds-brand`, three links (the current one on a grey pill), then at the right one `.ds-menu`, the date and one small primary button. The hero or the page header is pulled up under it by a negative margin, so the bar lies on the gradient.
- The home view stacks: `.ds-hero`, in two columns (7 to 5): kicker, title, one paragraph, the `.ds-countdown` and the `.ds-hero__crowd` at the left, the `.ds-signup` card at the right. The gradient is drawn behind them by `::before` and cut to an ellipse centred near the top left, so it is deepest under the copy and gives way under the card. Then `.ds-section--first` with `.ds-about` (6 to 5): copy and ticked points beside `.ds-stat`, four tinted figures two by two, the right pair stepped down; a white section with the `.ds-grid` of four speakers; a section of `.ds-cols--half`, the day's list in a card beside the `.ds-carousel` and the accordion; and a white section with `.ds-cta`, a centred gradient card. Sections have `--space-7` above and below.
- Inner pages have no hero. They open with `.ds-page-header`: the same curve, shallower, holding the breadcrumb, the page title, one line of text and at most one button at the right. `.ds-main` follows with `.ds-cols` (2 to 1) for a session, `--side` (1 to 3) with the `.ds-sidebar` at the left for the programme, or `--even` (7 to 5) for the form. On the programme a `.ds-toolbar` puts the `.ds-tabs`, here chips, at the left and a switch at the right, above the card with the table.
- Views. The specimen is the page of a two-day online conference about teaching courses, with four views. `home` is the event page. `agenda` is the programme: tracks, the label key and an empty personal schedule at the left; the day chips, the time zone switch and the table of sessions with the pagination in its foot at the right. `session` is one talk as running text with later sessions under it, beside the speaker, her workshop's places, a notice and related links. `register` is the form with its notices and buttons beside what the ticket holds and the open dialog for giving up a place.
- Below 960px the columns stack, the gradient ends above the foot of the sign-up card and speakers stand two across; below 640px the bar `.ds-nav__inner` wraps to two rows, the brand and `.ds-nav__actions` above and `.ds-nav__links` across the full width below, scrolling sideways, with the current link still on its grey pill, everything is one column and a wide table scrolls inside `.ds-table__wrap`. The page never scrolls sideways.
- Spacing scale: `--space-1` 4px (label to field, pagination gaps), `--space-2` 8px (icon to text, between badges and chips), `--space-3` 12px (button and cell vertical padding, between countdown tiles), `--space-4` 16px (between form rows, list row padding), `--space-5` 24px (card padding, grid gap, page gutter), `--space-6` 36px (under a section head, page top of inner pages), `--space-7` 60px (section padding, column gap of the hero), `--space-8` 100px (foot of the hero, end of the page).
- Other sizes: `--size-nav` 64px, `--size-mark` 30px, `--size-icon` 44px, `--size-control` 36px (pagination), `--size-avatar` 36px, `--size-face` 76px (a speaker), `--size-menu` 200px, `--size-tip` 220px, `--size-dialog` 440px.

## Typography and colour roles

These are the conventions of the era; every layout and token set keeps them.

- The gradient is `--fill-inverse`, made in the surface part from `--color-fill-1` and `--color-fill-2`. Text on it is always `--color-inverse-text`, which a palette may make light or dark; here the gradient is yellow to amber and the text ink. `--color-inverse` is the solid middle of the gradient: the text of the `--inverse` button.
- On the gradient there are two buttons only: `.ds-button--inverse` (fill `--color-inverse-text`, text `--color-inverse`) and `.ds-button--ghost` (outline in `--color-inverse-text`). The primary button never stands on the gradient itself; in the hero it stands in the white sign-up card and in the white bar.
- `--color-fill-1` to `--color-fill-4` are avatars, icon circles and, mixed at 24% into `--color-surface`, the tinted figures and the blocks above the speakers. A fill carries initials or an icon in `--color-inverse-text`, never running text; text on a tint is `--color-heading` and `--color-text`.
- `--color-page` is the pale page, `--color-surface` the card and the white sections, `--color-surface-alt` the quiet inset (hovered row and link, dialog foot, empty state), `--color-surface-strong` the strongest neutral (current link of the bar, quiet badge, inline code, current sidebar row, progress track).
- `--fill-bar` with `--color-bar-text` is the table head only. `--fill-bar-alt` with `--color-bar-alt-text` is the footer; nothing else stands on it.
- Text is `--color-text`, meta `--color-text-muted`, headings and figures `--color-heading`; kickers, sidebar titles, times of the day and h3 in prose are `--color-heading-alt`. A tooltip is `--color-surface` text on `--color-heading`.
- Links are `--color-link`, bold, underlined only on hover; quiet links and the links of the bar are `--color-link-quiet`.
- `--fill-button` with `--color-button-text` is the one action colour: the primary button, the current chip, the current page number, the current dot of the carousel, the switch when on. The secondary button is `--color-button-secondary-text` on `--fill-button-secondary` with a `--color-border` edge. The destructive button is `--color-danger` with `--color-button-text`.
- `--fill-accent` with `--color-accent-text` is the default badge. `--alt` and the status badges are tints: the colour mixed at 16 to 22% into `--color-surface`, with text mixed from the colour and `--color-heading`. `--color-success` is also the tick of a point.
- `--color-notice` with `--color-notice-text` is the information notice; the error notice is `--color-danger` on `--color-danger-surface` with a faint edge.
- Two families: `--font-heading` and `--font-ui` (Montserrat, Avenir Next) for headings, figures, buttons, labels and navigation, `--font-body` (Open Sans) for text; `--font-mono` for code. The references used commercial geometric faces. `--text-base` 16px at `--line-body` 1.62; `--text-ui` 14px (controls, tables, card text); `--text-small` 13px (meta, badges, labels); `--text-large` 20px (leads, the quote); `--text-display` 56px (hero title); `--text-h1` 38px (section and page titles, figures); `--text-h2` 26px; `--text-h3` 18px (card titles).
- Headings are heavy: `--weight-heading` and `--weight-display` 800, `--weight-bold` and `--weight-ui` 700, with tight tracking. Nothing is uppercase.
- Surface: buttons, inputs and badges are pills (`--radius-control`, `--radius-pill` 100px); cards are `--radius-panel` 18px; the bar and the call to action `--radius-page` 32px. `--border-width` is 2px in `--color-border-muted`, `--border-width-strong` 4px (the ring of a speaker's face). `--shadow-panel` is a solid lip of `--color-border` under the card and a wide soft shadow below; `--shadow-control` is a solid darker lip under buttons and fields that grows on hover; `--shadow-dialog` is the deepest and lies under the sign-up card.
- One-off values made with `calc()` and `color-mix()` because the vocabulary has no token: the ellipses that cut the gradient, washes on the gradient (`--color-inverse-text` at 14 to 18%, the ghost border at 55%), tinted badges, figures and blocks, the brand tile's edge, the error notice's edge, the footer rule. Opacity 0.85 to 0.9 dims secondary text on the gradient. The figures of the countdown are set at a line height of 1.1.

## Components

- `.ds-page`: on `<body>`. Sets font, text colour and `--fill-page`.
- `.ds-wrap`: the centred content column.
- `.ds-brand`: the site's mark and name, first in the bar and again in the footer. `.ds-brand__mark` is a 30px tile dressed like the primary button (`--fill-button`, `--color-button-text`, `--radius-control`, `--shadow-control`) holding the mark as inline SVG; `.ds-brand__name` is the name in bold. Name and mark are placeholders for the installing project's own name and logo.
- `.ds-nav`: the floating bar. `.ds-nav__inner` is the card; `.ds-nav__links` holds `.ds-nav__link`, the current one on a grey pill (`is-current`); `.ds-nav__actions` holds the menu, `.ds-nav__date` and one `.ds-button--small`. In the specimen the showing view marks its link with one `:has()` rule per view; a project sets `is-current` itself.
- `.ds-menu`: a drop-down, a `<details>` whose `.ds-menu__button` is the summary and whose `.ds-menu__list` of `.ds-menu__link` opens as a small card under it, aligned to the right.
- `.ds-hero`: the curve of the home view. `.ds-hero__inner` is two columns: `.ds-hero__copy` with `.ds-hero__kicker`, `.ds-hero__title`, `.ds-hero__lead`, the countdown and `.ds-hero__crowd` (a `.ds-avatars` stack and a line), and the `.ds-signup` card. One per site.
- `.ds-countdown`: the time left, a row of `.ds-countdown__item` tiles washed onto the gradient, each a `.ds-countdown__figure` over a `.ds-countdown__label`.
- `.ds-signup`: the card in the hero: `.ds-signup__title`, two `.ds-form__field`s, one `--wide` primary button, a `.ds-meter` and `.ds-signup__note`.
- `.ds-page-header`: head of an inner page, used instead of `.ds-hero`: `.ds-page-header__inner` holds the breadcrumb, `.ds-page-header__title`, `.ds-page-header__text` and `.ds-page-header__actions` (one `--inverse` button).
- `.ds-section`: a full-width block of the home view; `--alt` white, `--first` closer to the hero. `.ds-section__head` holds a `.ds-section__kicker`, `.ds-section__title` and `.ds-section__lead`, left-aligned. `.ds-main` is the area under a page header.
- `.ds-cols`: main and side column, 2 to 1; `--side` 1 to 3, `--even` 7 to 5, `--half` 1 to 1. `.ds-stack` stacks cards 24px apart.
- `.ds-about`: copy beside the figures. `.ds-points` is a list of ticked `.ds-points__item`.
- `.ds-stat`: the figures, four `.ds-stat__item` tiles two by two, each tinted with one of the fills and holding a `.ds-stat__figure` over a `.ds-stat__label`.
- `.ds-grid`: four cards across. Each `.ds-grid__cell` has a `.ds-grid__top` (a tinted block, `--2`, `--3`, `--4`, with a `.ds-avatar--face` and a badge) and a `.ds-grid__body`: `.ds-grid__title`, `.ds-grid__meta`, `.ds-grid__text` and a link.
- `.ds-slot`: on a `.ds-list__item`, puts a `.ds-slot__time` before the item.
- `.ds-carousel`: a row of slides of which one shows: `.ds-carousel__track` scrolls and snaps, each `.ds-carousel__slide` fills it; `.ds-carousel__bar` holds a caption and the `.ds-carousel__dots`, the current `.ds-carousel__dot` (`is-current`) longer and in the action colour.
- `.ds-quote`: a slide on the gradient: `.ds-quote__text` and `.ds-quote__by` with a washed `.ds-avatar` and `.ds-quote__name`.
- `.ds-speaker`: a `.ds-avatar--face` beside `.ds-speaker__name` and `.ds-speaker__role`, at the head of a card.
- `.ds-rows`: plain rows in a card, each `.ds-rows__item` a `.ds-rows__name` with a value or badge at the right.
- `.ds-cta`: the closing call to action, a centred gradient card with `.ds-cta__title`, `.ds-cta__text` and the two gradient buttons.
- `.ds-prose`: running text: h1, h2, h3, p, lists, `code`, `strong`. `.ds-prose__lead` for the first paragraph.
- `.ds-link`: text link. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `--quiet` for meta; `--more` appends an arrow. `.ds-linkrow` lays links in a row.
- `.ds-button`: pill button, 14px bold, with a solid lip. Variants: `--secondary`, `--danger` (destructive actions only), `--inverse` and `--ghost` (on the gradient only), `--small`, `--wide`. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttonrow` spaces buttons.
- `.ds-form`: vertical form with labels above the fields. `.ds-form__row` puts two fields side by side; `.ds-form__field` wraps `.ds-form__label` and one of `.ds-form__input`, `.ds-form__select`, `.ds-form__textarea`. Error: `is-error` on the input plus `.ds-form__error`. `.ds-form__hint`, `.ds-form__check` with `.ds-form__checkbox`, `.ds-form__actions`.
- `.ds-table`: data table inside a card: a small bold head on `--fill-bar`, rules between rows; `is-hover` tints a row. `.ds-table__num` right-aligns numbers; `.ds-table__title` is a row's link and `.ds-table__sub` a small line under it. Wrap it in `.ds-table__wrap`; `.ds-table__foot` is a last row for a count and the pagination.
- `.ds-list`: rows parted by rules. `.ds-list__item` holds a `.ds-list__title` (alone or in a `.ds-list__head` with a badge at the right), an optional `.ds-list__text` and `.ds-list__meta`.
- `.ds-panel`: the white card with the lip and the soft shadow. `.ds-panel__head` holds the `.ds-panel__title` and an optional badge, link or `.ds-panel__tools` above a rule; `.ds-panel__body` is padded; `.ds-panel__text` is a small paragraph.
- `.ds-tabs`: here a row of chips: `.ds-tabs__tab` pills with an edge, the current one (`is-current`) filled with `--fill-button`. `.ds-toolbar` sets them at the left with a tool at the right.
- `.ds-badge`: pill tag, 13px bold, on `--fill-accent`. `--alt` tinted, `--quiet` grey. Status variants: `--success`, `--warning`, `--danger`. `.ds-badgerow` for several.
- `.ds-sidebar`: a block of links: `.ds-sidebar__title` and a `.ds-sidebar__list` of `.ds-sidebar__item` with an optional `.ds-sidebar__count`; the current row (`is-current`) is a grey pill.
- `.ds-notice`: tinted rounded message; `--error` in the danger colours. `.ds-notice__title` for the bold lead-in. `.ds-noticerow` stacks notices.
- `.ds-pagination`: `.ds-pagination__link` pills; `is-current` is filled with `--fill-button`, `is-disabled` is grey.
- `.ds-breadcrumb`: a small trail of `.ds-breadcrumb__item` parted by slashes; the last is `.ds-breadcrumb__current`. It takes the colour of what it stands on.
- `.ds-dialog`: a dimmed block (`--color-overlay`) centring `.ds-dialog__box`, a card in `--shadow-dialog`: `.ds-dialog__head` (`.ds-dialog__title`, `.ds-dialog__close`), `.ds-dialog__body`, and `.ds-dialog__actions` right-aligned on `--color-surface-alt`.
- `.ds-empty`: a centred message on `--color-surface-alt`: `.ds-empty__icon` (a circle of `--color-fill-3`), `.ds-empty__title`, `.ds-empty__text` and one small secondary button.
- `.ds-accordion`: sections that open and close: each `.ds-accordion__item` is a `<details>` with a `.ds-accordion__title` summary (a plus that turns to a minus) and a `.ds-accordion__body`. Used for the questions people ask before they register.
- `.ds-tooltip`: a hint shown on hover or focus of its trigger; `.ds-tooltip__text` is the hint, above the trigger. `--mark` draws the trigger as a small round question mark; `--end` aligns the hint to the trigger's right edge; `is-open` shows it.
- `.ds-switch`: an on/off control, a `<label>` with a hidden `.ds-switch__input` checkbox, the `.ds-switch__track` and the text; on, the track takes `--fill-button`.
- `.ds-progress`: a thin track with a `.ds-progress__bar` in the gradient (`--33`, `--62`, `--82`, `--90` set its width; `--over` fills it in `--color-danger`). `.ds-meter` puts a `.ds-meter__row` (`.ds-meter__name` and the figures) over one.
- `.ds-avatar`: a person as initials on a round fill (`--1`, `--2`, `--4`). `--face` is the large one of a speaker, ringed in the card colour; `.ds-avatars` overlaps several in a row; `.ds-who` sets one beside a name.
- `.ds-tile`: an icon tile on one of the fills (`--2`, `--3`, `--4`; `--round` for a circle), for a new block that needs an icon.
- `.ds-footer`: the closing band on `--fill-bar-alt`, centred: the brand, a `.ds-footer__list` of `.ds-footer__link` and `.ds-footer__legal` over a rule.

## Never

- `text-shadow = none`: text is flat, also on the gradient.
- `animation = none`: nothing moves by itself, not even the countdown; only hover states ease.
- `border-width <= 4px`: borders are 2px; 4px rings a speaker's face.
- `border-radius <= 32px`: cards are 18px and the largest blocks 32px; only pills and circles are rounder.
- `box-shadow-blur <= 80px`: shadows are soft but end at 80px.
- `font-weight <= 800`: headings are heavy, nothing is black.
- `font-size >= 13px`: meta, badges and labels are the smallest text.
- `font-size <= 56px`: the hero title is the largest text.
- `font-families <= 3`: a display sans, a text sans and monospace for code.
- `line-height <= 1.7`: body text is set at 1.62.
- `uppercase-text <= 0%`: navigation, buttons and labels are in sentence case.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new block of the event page is a `.ds-section` with a `.ds-section__head`, alternating plain and `--alt`; a person is a `.ds-avatar`, large and ringed when it is their card; anything that holds data is a `.ds-panel`. Text goes on cards, on the page and, in `--color-inverse-text`, on the gradient; the four fills take initials or an icon only, and their tints take ordinary text. Put only the `--inverse` and `--ghost` buttons on the gradient, and at most one primary button in a card. Keep corners, shadows and borders on their tokens, spacing on the `--space-*` scale, and note any one-off `color-mix()` shade next to its rule.
