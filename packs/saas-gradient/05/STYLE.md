# SaaS gradient sign-up and onboarding flow

## Summary

This is the sign-up and set-up flow of a start-up product of 2016 to 2021: the window is split, with a tall panel of two-colour gradient at the left that carries the brand, the numbered steps and one customer's words, and at the right a narrow working column where each step asks one thing in soft-shadowed cards. Plans are chosen from cards with an icon tile, forms are short, and every step ends in a row of numbered pages and one primary button. Almost every product of the period greeted a new account this way.

## Layout

- The page is designed for a 1440px viewport. `.ds-page` is a grid of two columns: the gradient panel (`.ds-nav`, `--size-rail` 460px) for the whole height, and beside it the top line (`.ds-topbar`), the showing view and the footer line. The working column is centred in its half: `.ds-content` is at most `--size-page` (760px) wide plus `--space-7` side padding and stacks its children `--space-5` apart.
- Navigation is the panel: the `.ds-brand`, a `.ds-nav__title`, then the five steps as an ordered list of links, each a numbered ring, a label and a hint, the current one on a translucent block with its ring filled; a `.ds-quote` closes the panel. Its content stays in view while the working column scrolls.
- The home view is the first step: `.ds-hero` set straight on the page (a badge naming the step, the title, one paragraph, a field with the primary button beside it and a note), `.ds-stat` (three figures under a hairline) and one `.ds-aside` card for people who were invited.
- Inner pages have no hero. Each opens with `.ds-page-header` on the page: the breadcrumb, the step's title, one line of text and at most one secondary button at the right. The working components follow in one column: `.ds-tabs` and the `.ds-grid` of plans, or one `.ds-panel` with the form; `.ds-cols` (3 to 2, `--level` for equal height) holds two things side by side. A step ends with `.ds-stepfoot`: the pagination at the left and the primary button at the right.
- Views. The specimen is the set-up of a rota planner for cafes and shops, with five views. `start` is the home view. `plan` has the switch between monthly and yearly, four plan cards, the comparison table and the step foot. `workspace` is the form with its notices and buttons. `team` is the invite line and the list of people, then the open remove dialog beside an empty list of stand-ins. `welcome` closes the flow: a gradient card, the first week as running text, and reading, links and status labels at the right.
- Below 1180px the panel narrows to `--size-rail-narrow`; below 900px it becomes a band across the top holding the brand and the steps as one row of pills that scrolls sideways, without title, hints or quote; below 640px the step pills of `.ds-nav__links` shrink to their `.ds-nav__num` rings spread across the width and only the open step (its `.ds-nav__link`) keeps its `.ds-nav__label`, and plans, columns and the invite line are one column and the figures become rows. A wide table scrolls inside `.ds-table__wrap`; the page never scrolls sideways.
- Spacing scale: `--space-1` 4px (label to field, between steps), `--space-2` 8px (between badges), `--space-3` 12px (step padding, people rows, button padding), `--space-4` 16px (ring to label, between form rows), `--space-5` 24px (between cards, card padding, grid gap), `--space-6` 36px (panel padding, between parts of the panel, the closing card), `--space-7` 52px (side padding of the panel and the working column), `--space-8` 72px (foot of a view).
- Other sizes: `--size-mark` 32px, `--size-icon` 44px (icon tiles), `--size-control` 36px (step rings, pagination), `--size-avatar` 40px, `--size-dialog` 440px.

## Typography and colour roles

The conventions are the era's, as its first layout sets them out.

- The gradient is `--fill-inverse` (from `--color-fill-1` to `--color-fill-2`); text on it is `--color-inverse-text`. It fills the side panel and the closing card of the last step. Only `.ds-button--inverse` and `.ds-button--ghost` stand on it. `--color-inverse` is the number in the ring of the current step.
- `--color-fill-1` to `--color-fill-4` colour the icon tiles of the plans and the avatars. A fill carries an icon or initials in `--color-inverse-text`, never running text.
- `--color-page` (through `--fill-page`) is the working column, `--color-surface` (through `--fill-panel`) the cards, `--color-surface-alt` the dialog foot, a hovered table row and the empty state, `--color-surface-strong` the current sidebar row, the quiet badge and inline code.
- `--fill-bar` with `--color-bar-text` is the table head. `--fill-bar-alt` with `--color-bar-alt-text` is the footer line.
- Text is `--color-text` (`#b4c0e4`), meta and the top line `--color-text-muted`, headings, figures and prices `--color-heading` (white); sidebar titles and h3 in prose are `--color-heading-alt` (yellow).
- `--color-button` (yellow `#ffd23f`, midnight text) is the one action colour off the gradient: the primary button, the edge of the chosen plan, the edge under the current tab and the current page number. The secondary button is `--color-button-secondary-text` on `--fill-button-secondary` with a `--color-border` edge; the destructive button is `--color-danger` with `--color-button-text`.
- Links are `--color-link`, medium weight, underlined only on hover. `--color-success` is the tick of a plan's points and the dot of the top line.
- Badges and notices follow the era's rules: `--fill-accent` for the default badge, tints mixed into `--color-surface` for `--alt` and the status variants, `--color-notice` with `--color-notice-text`, the error notice in `--color-danger` on `--color-danger-surface`.
- One sans for text and controls (`--font-body`, `--font-ui`: IBM Plex Sans, then the system face), a second for headings (`--font-heading`: Rubik); the references used commercial grotesques. `--font-mono` for code. `--text-base` 16px at `--line-body` 1.5; `--text-ui` 14px (controls, tables, card text); `--text-small` 13px (hints, meta, badges); `--text-large` 19px (leads, the quote); `--text-display` 52px (the first question); `--text-h1` 34px (step titles, figures, prices, the panel title); `--text-h2` 24px; `--text-h3` 17px (card titles).
- Headings are heavy and tight: `--weight-heading` 600, `--weight-display` 700 with `--display-tracking` -1.5px. `--weight-ui` is 500. Nothing is uppercase.
- Surface: `--radius-control` 8px, `--radius-panel` 12px (cards, the block of the current step), `--radius-pill` 6px (badges), `--radius-page` 20px (the quote, the closing card, the dialog's backdrop). `--shadow-panel` is a faint light edge at the top of a card over one deep, dark shadow; `--shadow-dialog` adds a hairline ring and also marks the chosen plan. `--fill-page` puts a glow of `--color-surface-alt` at the top of the page and `--fill-inverse` a highlight in the corner of the gradient. `--border-width-strong` 2px is the ring of a step, the edge of a plan card and the edge under the current tab.
- One-off values made with `calc()` and `color-mix()` because the vocabulary has no token: the washes on the gradient (`--color-inverse-text` at 10 to 24%, ring borders at 50%, the ghost border at 55%), tinted badges, the width of the working column (`--size-page` plus its padding). Opacity 0.8 to 0.9 dims secondary text on the gradient.

## Components

- `.ds-page`: on `<body>`. The two-column grid; sets font, text colour and `--fill-page`.
- `.ds-brand`: the site's mark and name, at the head of the panel and again in the footer line. `.ds-brand__mark` is a tile dressed like the primary button holding the mark as inline SVG; `.ds-brand__name` is the name. Name and mark are placeholders for the installing project's own name and logo.
- `.ds-nav`: the gradient side panel. `.ds-nav__inner` holds the brand, `.ds-nav__title`, `.ds-nav__links` (an ordered list of `.ds-nav__link`, each a `.ds-nav__num` ring, a `.ds-nav__label` and a `.ds-nav__hint`) and the quote. In the specimen the showing view marks its step with one `:has()` rule per view; a project sets `is-current` itself.
- `.ds-quote`: one customer's words on a translucent block: `.ds-quote__text` and `.ds-quote__by` with a `.ds-avatar--plain` and `.ds-quote__name`.
- `.ds-avatar`: a circle of initials on a fill (`--1`, `--2`, `--4`; `--plain` on the gradient).
- `.ds-topbar`: the line above the working column: `.ds-topbar__note` with a `.ds-topbar__dot` at the left, the sign-in link at the right.
- `.ds-content`: the working column of a view. `.ds-stack` stacks things 24px apart; `.ds-cols` puts two side by side (3 to 2), `--level` at equal height.
- `.ds-hero`: the first question, on the home view only: `.ds-hero__kicker` (a badge), `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__form` (one field and the primary button) and `.ds-hero__note`.
- `.ds-stat`: the row of three figures under a hairline; each `.ds-stat__item` is a `.ds-stat__figure` over a `.ds-stat__label`.
- `.ds-aside`: a card that asks one side question: `.ds-aside__body` (a round `.ds-tile`, `.ds-aside__title`, `.ds-aside__text`) and one small secondary button.
- `.ds-page-header`: head of a step, used instead of `.ds-hero`: the breadcrumb, then `.ds-page-header__inner` with `.ds-page-header__title`, `.ds-page-header__text` and `.ds-page-header__actions`.
- `.ds-tabs`: a row of `.ds-tabs__tab` over a hairline; the current one (`is-current`) has an edge in `--color-button`.
- `.ds-grid`: plan cards, two across. Each `.ds-grid__cell` is a link: `.ds-grid__head` (a `.ds-tile`, `.ds-grid__title`, an optional badge), `.ds-grid__price` with `.ds-grid__per`, `.ds-grid__text` and `.ds-points`. `is-current` is the chosen plan.
- `.ds-points`: ticked `.ds-points__item` lines.
- `.ds-tile`: an icon tile on `--color-fill-1` (`--2`, `--3`, `--4`; `--round` for a circle).
- `.ds-stepfoot`: the foot of a step: the pagination and one primary button over a hairline.
- `.ds-invite`: a field, a select and a button in one line.
- `.ds-cta`: the gradient card that closes the flow: `.ds-cta__title`, `.ds-cta__text` and the two gradient buttons.
- `.ds-prose`: running text: h1, h2, h3, p, lists, `code`, `strong`. `.ds-prose__lead` for the first paragraph.
- `.ds-link`: text link. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `--quiet` for meta; `--more` appends an arrow. `.ds-linkrow` lays links in a row.
- `.ds-button`: 14px medium. Variants: `--secondary`, `--danger` (destructive actions only), `--inverse` and `--ghost` (on the gradient only), `--small`, `--wide`. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttonrow` spaces buttons.
- `.ds-form`: vertical form with labels above the fields. `.ds-form__row` puts two fields side by side; `.ds-form__field` wraps `.ds-form__label` and one of `.ds-form__input`, `.ds-form__select`, `.ds-form__textarea`. Error: `is-error` on the input plus `.ds-form__error`. `.ds-form__hint`, `.ds-form__check` with `.ds-form__checkbox`, `.ds-form__actions`.
- `.ds-table`: data table inside a card: a small bold head on `--fill-bar`, hairlines between rows; `is-hover` tints a row. `.ds-table__num` right-aligns numbers; `.ds-table__title` is a row's link. Wrap it in `.ds-table__wrap`; `.ds-table__foot` is a foot row.
- `.ds-list`: rows parted by hairlines: `.ds-list__item` with a `.ds-list__title` (alone or in a `.ds-list__head`), `.ds-list__text` and `.ds-list__meta`. `--people` lays a row out as avatar, `.ds-list__grow` (name and meta) and a status badge.
- `.ds-panel`: the card. `.ds-panel__head` holds the `.ds-panel__title` and an optional badge above a hairline; `.ds-panel__body` is padded; `.ds-panel__text` is a small paragraph; `--side` pads a card that holds a sidebar block.
- `.ds-badge`: small tag on `--fill-accent`. `--alt` tinted, `--quiet` grey. Status variants: `--success`, `--warning`, `--danger`. `.ds-badgerow` for several.
- `.ds-sidebar`: a block of links: `.ds-sidebar__title` and a `.ds-sidebar__list` of `.ds-sidebar__item` with an optional `.ds-sidebar__count`; the current row (`is-current`) is filled.
- `.ds-notice`: tinted message; `--error` in the danger colours. `.ds-notice__title` for the bold lead-in. `.ds-noticerow` stacks notices.
- `.ds-pagination`: the steps as numbers: `.ds-pagination__link`; `is-current` is filled with `--fill-button`, `is-disabled` is grey.
- `.ds-breadcrumb`: a small trail of `.ds-breadcrumb__item` parted by slashes; the last is `.ds-breadcrumb__current`.
- `.ds-dialog`: a dimmed block (`--color-overlay`) centring `.ds-dialog__box`, a card in `--shadow-dialog`: `.ds-dialog__head` (`.ds-dialog__title`, `.ds-dialog__close`), `.ds-dialog__body`, and `.ds-dialog__actions` right-aligned on `--color-surface-alt`.
- `.ds-empty`: a centred message on `--color-surface-alt`: `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one small secondary button.
- `.ds-footer`: one line on `--fill-bar-alt` at the foot of the working column: the brand, `.ds-footer__links` of `.ds-footer__link` and `.ds-footer__legal`.
- `.ds-menu`: a drop-down, a `<details>` whose `.ds-menu__button` is the summary and whose `.ds-menu__list` of `.ds-menu__link` opens as a small card under it. Use it for a group of links under one bar item and for the account.
- `.ds-switch`: an on or off setting drawn over a checkbox (`.ds-switch__input`, `.ds-switch__track`, then the label). Use it for a setting that takes effect at once, not for a choice sent with a form's other fields.
- `.ds-tooltip`: a short hint on hover or focus; `.ds-tooltip--mark` is the small round question mark beside a label, `.ds-tooltip__text` the hint, `.ds-tooltip--start` opens it from the left edge. One sentence at most.
- `.ds-progress`: a thin pill bar with `.ds-progress__bar` (`--20` to `--100` set the width). `.ds-meter` puts a name and a figure above it. Use it for steps done and for usage against a limit.
- `.ds-nav__progress`: the `.ds-progress` under the steps in the bar; it fills by one fifth with each view.

## Never

- `text-shadow = none`: text is flat, also on the gradient.
- `animation = none`: nothing moves by itself; only hover states ease.
- `border-width <= 2px`: borders are hairlines; 2px marks the step rings, the plan cards and the current tab.
- `border-radius <= 20px`: cards are 12px and the largest blocks 20px; only circles are rounder.
- `box-shadow-blur <= 90px`: shadows are deep and soft but end at 90px.
- `font-weight <= 700`: titles are bold, never black.
- `font-size >= 13px`: hints, meta and badges are the smallest text.
- `font-size <= 52px`: the first question is the largest text.
- `font-families <= 3`: a heading face, a body sans and monospace for code.
- `line-height <= 1.5`: body text is set at 1.5.
- `uppercase-text <= 0%`: steps, buttons and labels are in sentence case.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new step is another `.ds-nav__link` in the panel and a view that opens with `.ds-page-header`, asks one thing in a `.ds-panel` or a `.ds-grid` of choices, and ends in a `.ds-stepfoot`; a side question is a `.ds-aside`. Keep the working column to `--size-page` and one primary button per step. Text goes on cards, on the page and, in `--color-inverse-text`, on the gradient; the four fills take an icon or initials only, and only the `--inverse` and `--ghost` buttons stand on the gradient. Keep corners, shadows and borders on their tokens, spacing on the `--space-*` scale, and note any one-off `color-mix()` shade next to its rule.
