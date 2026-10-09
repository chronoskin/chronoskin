# SaaS gradient help centre

## Summary

This is the help centre of a start-up product of 2016 to 2021: a band of two-colour gradient whose lower edge is one long curve carries a centred question and a large white search field, and the collections sit below as white cards with round icon tiles that climb onto the band. Inner pages keep a shorter band for the breadcrumb and title and put articles, lists and forms in soft-shadowed cards on a warm pale page. The type pairs a geometric display face with a plain humanist sans, and corners are tight.

## Layout

- The page is designed for a 1440px viewport and is not fluid: content sits in `.ds-wrap`, centred, at most `--size-page` (980px) wide with `--space-5` side padding. The band and the footer run the full width of the window.
- Navigation is one transparent row, `--size-nav` (68px) tall, on the gradient: the `.ds-brand`, a `.ds-nav__tag` naming the help centre behind a hairline, and three links at the right, the current one underlined. The hero or page header is pulled up under it by a negative margin.
- The home view stacks: `.ds-hero`, centred, with the title, one line, the `.ds-search` field and a line of often-asked links, its lower edge curved with `clip-path: ellipse()`; then `.ds-main--home`, which rises onto the band: `.ds-grid`, six collection cards three across; `.ds-cols` (5 to 3, `--level`) with the most-read list beside the service status; `.ds-stat`, three figures on a tinted strip; and `.ds-cta`, a gradient card. Blocks are `--space-6` apart.
- Inner pages have no hero. They open with `.ds-page-header`: the same curved band, shorter, holding the breadcrumb, the title, one line of text and at most one white button at the right. `.ds-main` rises `--space-8` onto it. A collection is `.ds-cols--side` (2 to 5): the collections `.ds-sidebar` in a card at the left, and at the right one card with `.ds-tabs` on its head, the article list and a `.ds-table__foot` with the pagination. An article and the contact form use `.ds-cols` (5 to 3) with small cards at the right.
- Views. The specimen is the help centre of an online booking product. `home` is the search page. `collection` lists the articles of one collection under tabs, with the collections at the left and an empty video box. `article` is one article with a table of plan limits and the feedback row, beside a notice, related links and what comes next. `contact` is the message form with its notices and buttons beside two suggested articles and the open discard dialog.
- Below 860px the columns stack and collections stand two across; below 640px `.ds-nav__links` is not shown and the menu `.ds-nav__menu` takes its place at the end of the bar after the brand, listing every view and marking the one showing, the curve flattens, and grids and figures go to one column. A wide table scrolls inside `.ds-table__wrap`; the page never scrolls sideways.
- Spacing scale: `--space-1` 4px (label to field, pagination gaps), `--space-2` 8px (search field padding, between badges, service rows), `--space-3` 12px (button and cell vertical padding), `--space-4` 16px (list rows, between form rows, tab padding), `--space-5` 24px (card padding, grid gap, page gutter), `--space-6` 36px (between blocks, footer padding), `--space-7` 60px (hero top, foot of a view), `--space-8` 88px (how far the body rises onto the band).
- Other sizes: `--size-mark` 30px, `--size-icon` 48px (icon tiles), `--size-control` 36px (pagination), `--size-search` 620px, `--size-dialog` 400px.

## Typography and colour roles

The conventions are the era's, as its first layout sets them out.

- The gradient is `--fill-inverse` (from `--color-fill-1` to `--color-fill-2`); text on it is `--color-inverse-text`. It fills the band and the `.ds-cta` card. Only `.ds-button--inverse` and `.ds-button--ghost` stand directly on it. The search field is a card on the band (`--fill-input`), so its button is the primary one.
- `--color-fill-1` to `--color-fill-4` are the round icon tiles of the collections and the empty state's icon. A fill carries an icon in `--color-inverse-text`, never running text.
- `--color-page` is the warm pale page, `--color-surface` (through `--fill-panel`) the cards, `--color-surface-alt` the dialog foot and the empty state, `--color-surface-strong` the strip of figures, the current sidebar row, the quiet badge and inline code.
- `--fill-bar` with `--color-bar-text` is the table head. `--fill-bar-alt` with `--color-bar-alt-text` is the footer.
- Text is `--color-text` (`#46585c`), meta `--color-text-muted`, headings and figures `--color-heading` (`#17262b`); sidebar titles and h3 in prose are `--color-heading-alt`.
- `--color-button` (violet `#5b4bd6`, white text) is the action colour on cards, the current page number and the edge under the current tab. Links are `--color-link`, bold, not underlined; only the often-asked links on the band are underlined, because colour cannot mark them there.
- Badges, notices and the destructive button follow the era's rules: `--fill-accent` for the default badge, tints mixed into `--color-surface` for `--alt` and the status variants, `--color-notice` with `--color-notice-text`, `--color-danger` with `--color-button-text`.
- Two families: `--font-heading` is a geometric display sans (Poppins, Montserrat, Futura) for titles and figures, `--font-body` and `--font-ui` a humanist sans (Lato, Lucida Grande); `--font-mono` for code. `--text-base` 15px at `--line-body` 1.55; `--text-ui` 13px (controls, card text, navigation); `--text-small` 12px (meta, badges, labels); `--text-large` 18px (leads); `--text-display` 42px (the hero question); `--text-h1` 32px (page titles, figures, article title); `--text-h2` 23px; `--text-h3` 16px (card and list titles).
- The hero and page titles are bold (`--weight-display` 700); other headings are medium (`--weight-heading` 500). Buttons, tabs and navigation links are uppercase and bold (`--ui-transform`, `--weight-ui` 700); nothing else is uppercase.
- Surface: tight corners (`--radius-control` and `--radius-panel` 5px, `--radius-pill` 3px for badges, `--radius-page` 8px for collection cards, the strip and the call to action). `--shadow-panel` is a close 3px shadow under a wide 28px one; `--shadow-dialog` is deeper and also lifts the search field and a hovered collection card. The primary button has a faint top-to-bottom gradient of its own colour (`--fill-button`).
- One-off values made with `calc()` and `color-mix()`: the curve (`ellipse(90% 100% at 50% 0%)`), the hairline before the navigation tag (`--color-inverse-text` at 40%), the ghost button border (55%), tinted badges, the error notice's edge (`--color-danger` at 35%). Opacity 0.75 to 0.9 dims secondary text on the gradient and in the footer.

## Components

- `.ds-page`: on `<body>`. Sets font, text colour and `--fill-page`.
- `.ds-wrap`: the centred content column. `.ds-main` is the body of a view, stacked `--space-6` apart and raised onto the band (`--home` a little further).
- `.ds-brand`: the site's mark and name, first in the navigation and again in the footer. `.ds-brand__mark` is a 30px tile dressed like the primary button (`--fill-button`, `--color-button-text`, `--radius-control`, `--shadow-control`) holding the mark as inline SVG; `.ds-brand__name` is the name in bold. Name and mark are placeholders for the installing project's own name and logo.
- `.ds-nav`: the transparent row on the gradient. `.ds-nav__inner` centres it; `.ds-nav__tag` names the section; `.ds-nav__links` holds `.ds-nav__link`, the current one underlined (`is-current`). In the specimen the showing view marks its link with one `:has()` rule per view; a project sets `is-current` itself.
- `.ds-hero`: the curved gradient band of the home view, centred: `.ds-hero__title`, `.ds-hero__lead`, a `.ds-search` and `.ds-hero__topics` of `.ds-hero__topic` links. One per site.
- `.ds-search`: the large search field, a white bar holding a `.ds-search__icon`, the `.ds-search__input` and one primary button.
- `.ds-page-header`: the band of an inner page, used instead of `.ds-hero`. `.ds-page-header__inner` holds the breadcrumb, `.ds-page-header__title` and `.ds-page-header__text` at the left and `.ds-page-header__actions` at the right.
- `.ds-cols`: main and side column, 5 to 3; `--side` 2 to 5 with the sidebar first; `--level` ends the columns on one line. `.ds-stack` stacks cards; `.ds-panel--grow` fills a column.
- `.ds-grid`: six collection cards, three across. Each `.ds-grid__cell` is a centred link: a round `.ds-tile`, `.ds-grid__title`, `.ds-grid__text` and `.ds-grid__count`.
- `.ds-tile`: an icon tile on `--color-fill-1` (`--2`, `--3`, `--4`; `--round`).
- `.ds-stat`: the strip of three figures on `--color-surface-strong`; each `.ds-stat__item` is a `.ds-stat__figure` over a `.ds-stat__label`.
- `.ds-rows`: plain rows in a card, `.ds-rows__row`, a name at the left and a badge at the right; used for the service status.
- `.ds-cta`: the closing call to action, a gradient card with `.ds-cta__title`, `.ds-cta__text` and one inverse button.
- `.ds-prose`: running text of an article: h1, h2, h3, p, lists, `code`, `strong`. `.ds-prose__lead` for the first paragraph. `.ds-article__meta` is the line above it (a badge, the date, the author); `.ds-article__block` frames a table inside the text; `.ds-feedback` is the closing row with the question and two small buttons.
- `.ds-link`: text link. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `--quiet` for meta; `--more` appends an arrow. `.ds-linkrow` lays links in a row.
- `.ds-button`: button with 5px corners, 13px bold uppercase. Variants: `--secondary`, `--danger` (destructive actions only), `--inverse` and `--ghost` (on the gradient only), `--small`, `--wide`. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttonrow` spaces buttons.
- `.ds-form`: vertical form with labels above the fields, inside a card. `.ds-form__row` puts two fields side by side; `.ds-form__field` wraps `.ds-form__label` and one of `.ds-form__input`, `.ds-form__select`, `.ds-form__textarea`. Error: `is-error` on the input plus `.ds-form__error`. `.ds-form__hint`, `.ds-form__check` with `.ds-form__checkbox`, `.ds-form__actions`.
- `.ds-table`: data table: a small bold head on `--fill-bar`, hairlines between rows; `is-hover` tints a row. `.ds-table__num` right-aligns numbers; `.ds-table__title` is a row's link. Wrap it in `.ds-table__wrap`; `.ds-table__foot` is the last row of a card, pagination at the left and a count at the right.
- `.ds-list`: rows parted by hairlines. `.ds-list__item` holds a `.ds-list__title` (alone or in a `.ds-list__head` with a badge at the right), an optional `.ds-list__text` and `.ds-list__meta`.
- `.ds-panel`: the white card with the soft shadow. `.ds-panel__head` holds the `.ds-panel__title` and an optional link above a hairline; `.ds-panel__body` is padded; `.ds-panel__text` is a small paragraph; `--side` is a slim card around a `.ds-sidebar`.
- `.ds-tabs`: underlined tabs on the head of a card: `.ds-tabs__tab`, the current one (`is-current`) in the heading colour over a 2px edge in `--color-button`.
- `.ds-badge`: small tag, 12px bold, on `--fill-accent`. `--alt` tinted, `--quiet` grey. Status variants: `--success`, `--warning`, `--danger`. `.ds-badgerow` for several.
- `.ds-sidebar`: a block of links: `.ds-sidebar__title` and a `.ds-sidebar__list` of `.ds-sidebar__item` with a `.ds-sidebar__count`; the current row (`is-current`) is on `--color-surface-strong`.
- `.ds-notice`: tinted message; `--error` in the danger colours. `.ds-notice__title` for the bold lead-in. `.ds-noticerow` stacks notices.
- `.ds-pagination`: `.ds-pagination__link` squares; `is-current` is filled with `--fill-button`, `is-disabled` is grey.
- `.ds-breadcrumb`: a small trail of `.ds-breadcrumb__item` parted by slashes; the last is `.ds-breadcrumb__current`. It stands on the band, above the page title.
- `.ds-dialog`: a dimmed block (`--color-overlay`) centring `.ds-dialog__box`, a card in `--shadow-dialog`: `.ds-dialog__head` (`.ds-dialog__title`, `.ds-dialog__close`), `.ds-dialog__body`, and `.ds-dialog__actions` right-aligned on `--color-surface-alt`.
- `.ds-empty`: a centred message on `--color-surface-alt`: `.ds-empty__icon` (a circle of `--color-fill-3`), `.ds-empty__title`, `.ds-empty__text` and one small secondary button.
- `.ds-footer`: the closing band on `--fill-bar-alt`, centred: the brand, `.ds-footer__links` of `.ds-footer__link`, and a `.ds-footer__legal` line.
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
- `border-width <= 2px`: borders are hairlines; 2px marks the current tab and navigation link.
- `border-radius <= 8px`: controls and cards are 5px and the largest blocks 8px; only circles are rounder.
- `box-shadow-blur <= 64px`: shadows are soft and wide but end at 64px.
- `font-weight <= 700`: titles are bold, never black.
- `font-size >= 12px`: meta, badges and labels are the smallest text.
- `font-size <= 42px`: the hero question is the largest text.
- `font-families <= 3`: a display face, a body sans and monospace for code.
- `line-height <= 1.55`: body text is set at 1.55.
- `uppercase-text <= 10%`: only buttons, tabs and navigation links are uppercase.
- `underlined-links <= 10%`: only the often-asked links on the band are underlined.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new kind of page opens with a `.ds-page-header` band and puts its content in `.ds-main`, in cards; a new collection is another `.ds-grid__cell` with a round tile on one of the fills; anything that lists or explains is a `.ds-panel` in a column of `.ds-cols`. Text goes on cards, on the page and, in `--color-inverse-text`, on the gradient; fills take an icon only. Put only `--inverse` and `--ghost` buttons directly on the gradient, and keep at most one primary button in a card. Keep corners tight on their tokens, spacing on the `--space-*` scale, and note any one-off `color-mix()` shade next to its rule.
