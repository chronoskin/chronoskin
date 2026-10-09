# Claymorphism shop

## Summary

Claymorphism is the soft 3D look of 2021 to 2025: thick, very round shapes that seem modelled from coloured clay, with a light inner glow on one side, a darker inner shadow on the other and a large soft shadow under them. Around 2022 it also had a night form, on token, game and toy pages: a deep indigo page on which candy-coloured clays seem to glow. This pack is a small toy shop in that form: a clay bar across the top of the window, a hero card with a tray of clay toys, product cards with a coloured tray each, a product view, a basket and a checkout, set in a geometric display face over a rounded body face.

## Layout

- Content column: `--size-page` (1200px), centred, with `--space-5` (24px) side padding. The bar and the footer run across the whole window; the layout is fluid below that width.
- Navigation: `.ds-nav` is a clay bar attached to the top edge, rounded only at the bottom. Its first row has the brand at the left, the search well in the middle and the basket key and the account avatar at the right; its second row has the sections as pills (one of them a `.ds-menu`) and a small note at the right end.
- Home page: every view starts `--space-6` (32px) under the bar. The shop front is a `.ds-page__stack` at 32px gaps: the hero card (offer left, tray of toys right), the `.ds-stat` row of three pills, then `.ds-page__shop`: the `.ds-sidebar` of filters in a `--size-side` (248px) column and the shelf, a `.ds-grid` of product cards three across with the pagination under it.
- Inner pages have no hero. They start with the `.ds-breadcrumb` and the `.ds-page-header` (title at `--text-h1`, one lead line, actions at the right) on the bare page. A product page has `.ds-page__cols--even`: the picture carousel at the left, the buying panel and an accordion at the right; under it `.ds-page__cols` with the reviews list and a ratings panel. A basket page uses `.ds-page__cols`: notices, the table and the empty state in the wide column, the order panel and the dialog in a `--size-aside` (392px) column. A checkout page uses `.ds-page__cols`: the form in a `.ds-panel`, a prose panel in the aside. Keep the two columns about equally long.
- Views. The specimen is one shop with four views, each `<section class="ds-view" id="...">` inside `.ds-page__inner`; the bar and the footer stand outside them. `shop` is the home view, `toy` the product, `basket` the basket, `checkout` the form. The pill of the current view is pressed into the bar by one `:has()` rule per view.
- Spacing scale: `--space-1` 4px, `--space-2` 8px (between pills, label to control), `--space-3` 12px (card inset around a tray, between list bars), `--space-4` 16px (between buttons), `--space-5` 24px (panel padding, grid gap), `--space-6` 32px (between blocks and columns), `--space-7` 44px (footer padding), `--space-8` 64px (above the footer).
- `--size-pic` (208px) is the height of a product tray; `--size-rule` (2px) the weight of the rules between table rows.
- Below 1000px the columns, the shop front and the hero stack and the shelf goes two across. Below 560px the search drops to its own line in the bar, the shelf and the stats are one column, form rows stack, the tabs fill the width and table columns marked `.ds-table__opt` are hidden, so nothing scrolls sideways.
- On a phone (560px and below) `.ds-nav__top` wraps: the brand and the `.ds-nav__keys` on the first line, `.ds-nav__search` on a line of its own, then `.ds-nav__list` with smaller links. The Checkout link is not shown; the Basket link stands for both and is pressed in on the checkout view as well. The bar's bottom corners are rounded with `--radius-panel`.

## Typography and colour roles

- `--font-heading` and `--font-ui` are one geometric sans with round bowls (Outfit, then Lexend, Futura) for titles, prices, buttons and navigation; `--font-body` is a rounded sans (Quicksand, then Hiragino Maru Gothic, Varela Round). `--font-mono` is for code only. These are open stand-ins for the commercial rounded faces of the references.
- Sizes: `--text-base` 18px, `--text-ui` 17px for controls and navigation, `--text-small` 14px for meta text and badges, `--text-large` 22px for lead text and large buttons, `--text-h3` 21px, `--text-h2` 26px (prices on cards), `--text-h1` 36px, `--text-display` 58px for the hero title only.
- Weights: 400 body, 500 for buttons and navigation, 600 for headings and bold text, 700 for the hero title and the large price. Headings are tracked at 0.01em.
- Text tokens sit on fills as the era's first layout sets out: body, muted, heading and link colours only on the page and the three surface colours; `--color-accent-text` on the accent and on every `--color-fill-*` clay (badges on trays, avatars, faces), with nothing else set on a fill; `--color-bar-text` on `--fill-bar` (the top bar); `--color-bar-alt-text` on `--fill-bar-alt` (the footer) and on `--color-bar-alt` (table head); `--color-inverse-text` on `--fill-inverse` (code, tooltips); `--color-button-text` on `--fill-button`; `--color-button-secondary-text` on `--fill-button-secondary` (secondary buttons, round keys, pagination, switch thumbs); `--color-input-text` on `--fill-input` (controls, the search well, the hero tray, tab and progress tracks, the empty state); `--color-notice-text` on `--color-notice`; `--color-danger` on `--color-danger-surface`. Status colours never carry text.
- The palette is dark: the page is deep indigo, the surfaces are lighter indigo clays, and the fills, the accent and the button are candy colours light enough to carry the dark ink `--color-accent-text`.
- The clay comes from the surface part: panels have a dark drop shadow and two insets mixed from black and white; controls throw a glow in `--color-button` instead of a dark shadow, and the dialog and the hero card add a halo in `--color-accent`. `--shadow-text` is a faint glow on the text of primary buttons. Clay objects add a drop shadow in their own fill through `color-mix()` with `transparent`, written in the component.
- One-off values with no token: circles use `border-radius: 50%`, the cube 28%, checkboxes 34%; a tray takes the card's radius less its inset, by `calc()`; list bars, notices and accordion bars take the smaller of `--radius-panel` and `--radius-control`; placeholders are `--color-input-text` at opacity 0.55.

## Components

- `.ds-page`: on `<body>`; paints `--fill-page` and sets the base type. `.ds-page__inner` is the width container, `.ds-page__stack` a vertical stack 32px apart, `.ds-page__shop` the filters beside the shelf, `.ds-page__cols` a wide column with an aside (`--even` for two halves), `.ds-page__heading` a section heading on the bare page. `.ds-view` is one screen of the specimen; an application with one document per page does not need it.
- `.ds-brand`: the mark and name in the bar: `.ds-brand__mark` is a clay ball `--size-brand` (48px) dressed like the primary button, holding the mark as inline SVG; `.ds-brand__name` is the name in the heading font, in the colour of the bar. The name `chronoskin` and its mark are placeholders: replace them with the installing project's own and keep the ball.
- `.ds-nav`: the top bar. `.ds-nav__inner` is the width container, `.ds-nav__top` the first row with `.ds-nav__search` and `.ds-nav__keys` (round `.ds-nav__key` buttons, a count as a badge on the corner), `.ds-nav__list` the row of sections, `.ds-nav__link` one pill (pressed in when `is-current`, raised on `is-hover`), `.ds-nav__note` the note at the right end.
- `.ds-menu`: a `<details>` drop-down under a section: `.ds-menu__toggle` on the summary, `.ds-menu__list` the clay card, `.ds-menu__link` with a count in `.ds-menu__note`; `--end` aligns it right, `--up` opens it upwards.
- `.ds-hero`: the offer on one clay card: `.ds-hero__text` with a badge, `.ds-hero__title` at the display size, `.ds-hero__lead` and `.ds-hero__actions`; `.ds-hero__tray` is a dent holding `.ds-clay` toys placed with `.ds-hero__obj--a` to `--e`.
- `.ds-clay`: a clay object drawn in CSS: a ball; `--1` to `--4` pick the fill, `--tiny`, `--small`, `--large` the size, `--cube`, `--pill`, `--ring` the shape; a face holds two `.ds-clay__eye`, a `.ds-clay__smile` and two `.ds-clay__cheek`. Inside a tray the hole of a ring takes the tray's colour. Decorative only.
- `.ds-page-header`: head of an inner page: `.ds-page-header__text` with `.ds-page-header__title` and `.ds-page-header__lead`, `.ds-page-header__actions` at the right.
- `.ds-prose`: running text inside a panel: h1 to h3, paragraphs, lists, inline `code`; `.ds-prose__lead`, `.ds-prose__code`.
- `.ds-link`: link in running text, marked by colour alone; states `is-visited`, `is-hover`, `is-active`, `is-focus`; `.ds-link--quiet` for table links; `.ds-link__sample` lays out a row of links.
- `.ds-button`: a fat clay pill 50px high that glows in its own colour. `is-hover` lifts it and widens the glow, `is-active` squashes it into a dent. `.ds-button--secondary` in a surface clay, `.ds-button--danger` on `--color-danger-surface`, `.ds-button--large` (62px), `.ds-button--wide`; `disabled` or `is-disabled` is flat. `.ds-button__row` spaces a group.
- `.ds-form`: vertical form in a panel: `.ds-form__field`, `.ds-form__label`, `.ds-form__input`, `.ds-form__textarea`, `.ds-form__select` (in `.ds-form__select-wrap` with `.ds-form__chevron`), `.ds-form__check` with `.ds-form__checkbox`, `.ds-form__error` after an `is-error` control, `.ds-form__hint`, `.ds-form__actions`, `.ds-form__row`. Controls are dents, darker than the card.
- `.ds-switch`: a fat clay toggle over a checkbox: `.ds-switch__input`, `.ds-switch__track`, `.ds-switch__label`, `.ds-switch__hint`; `--end` puts the track at the right. Used for filters and options that apply at once.
- `.ds-table`: one clay sheet with a head in `--color-bar-alt`: `.ds-table__num`, `.ds-table__meta`, `.ds-table__strong`, `.ds-table__opt`, `is-hover` on a row.
- `.ds-list`: reviews as separate clay bars: `.ds-list__item` holds an avatar and a `.ds-list__body` (`.ds-list__title`, `.ds-list__meta`); `.ds-list__aside` is a figure at the right.
- `.ds-panel`: the basic clay card: `.ds-panel__title`, `.ds-panel__body`, `.ds-panel__note`, `.ds-panel__line`; `.ds-panel--accent` in the accent clay. `.ds-panel__price` is the large price, `.ds-panel__meter` a rating row (figure, bar, share), `.ds-panel__total` the ruled last line of an order.
- `.ds-grid`: the shelf, three cards across. `.ds-grid__cell` is a clay card (it tilts and lifts on `is-hover`); `--1` to `--4` colour its `.ds-grid__pic`, a dented tray in a `--color-fill-*` clay holding a `.ds-clay` toy and an optional badge; then `.ds-grid__title`, `.ds-grid__meta` and `.ds-grid__foot` with the `.ds-grid__price` and one button. `.ds-grid__pic--large` is the tall tray of the product view.
- `.ds-stat`: the row of three; each `.ds-stat__item` is a clay pill with a `.ds-stat__dot` ball (`--2`, `--3` pick a fill), a `.ds-stat__value` and a `.ds-stat__label`; `.ds-stat__meta` is a note beside the value.
- `.ds-tabs`: a dented track of `.ds-tabs__tab` pills; `is-current` is raised, in the accent. Here it picks a colour set of a toy.
- `.ds-badge`: a pill in the accent; `.ds-badge--quiet` neutral; `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger` tinted with a status dot.
- `.ds-sidebar`: the filters, a clay card: `.ds-sidebar__heading`, `.ds-sidebar__list`, `.ds-sidebar__link` (dented when `is-current`), `.ds-sidebar__options` for a column of switches.
- `.ds-notice`: a raised bar in `--color-notice` with `.ds-notice__icon` and `.ds-notice__title`; `.ds-notice--error`; `.ds-notice__stack`.
- `.ds-pagination`: clay keys, `.ds-pagination__link`, the current one `is-current` in the accent.
- `.ds-breadcrumb`: small bold trail above the page header: `.ds-breadcrumb__item`, `.ds-breadcrumb__link`.
- `.ds-dialog`: a tinted sheet with a centred `.ds-dialog__box` under a halo: `.ds-dialog__title`, `.ds-dialog__body`, `.ds-dialog__actions`.
- `.ds-empty`: a dent holding a clay ball `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one primary button.
- `.ds-tooltip`: wraps a trigger; `.ds-tooltip__tip` shows on hover or focus.
- `.ds-progress`: a thick dented track with a puffy `.ds-progress__bar` (`--10` to `--90`), `.ds-progress__legend` above it; `.ds-progress--ring` is a ring around a clay ball `.ds-progress__face`.
- `.ds-avatar`: initials on a clay ball `.ds-avatar__pic`, with optional `.ds-avatar__text` (`.ds-avatar__name`, `.ds-avatar__meta`); `--1` to `--4` pick a fill.
- `.ds-accordion`: `<details>` sections as separate clay bars: `.ds-accordion__item`, `.ds-accordion__head`, `.ds-accordion__chevron`, `.ds-accordion__body`.
- `.ds-carousel`: `.ds-carousel__track` with one `.ds-carousel__slide` showing (here a product card with a large tray), `.ds-carousel__foot` with two `.ds-carousel__arrow` keys around `.ds-carousel__dots` (`.ds-carousel__dot`).
- `.ds-footer`: a clay band across the bottom of the window, rounded only at the top, in `--fill-bar-alt`: `.ds-footer__inner`, `.ds-footer__about` with the brand and `.ds-footer__text`, `.ds-footer__heading`, `.ds-footer__links`, `.ds-footer__link`, `.ds-footer__toys` for three small clay balls.

## Never

- `box-shadow != none`: every card, key and toy is modelled; nothing raised is flat.
- `border-width <= 3px`: clay has no outlines; the only lines are 2px rules and the 3px focus or error ring.
- `border-radius <= 56px`: panels are 40px, the hero card and the dialog 56px; only pills and balls go higher.
- `box-shadow-blur <= 70px`: the halo of the dialog is the widest blur.
- `font-size <= 58px`: the hero title is the largest text.
- `font-size >= 14px`: meta text and badges are the smallest text.
- `font-weight <= 700`: the type is round, not heavy.
- `font-families <= 3`: one display face, one rounded sans and one monospace.
- `line-height <= 1.55`: body text is set at 1.5.
- `uppercase-text <= 5%`: nothing is set in capitals.
- `underlined-links <= 10%`: links are marked by colour alone.
- `block-gap <= 64px`: blocks sit 32px apart; 64px above the footer is the largest gap.
- `content-width <= 1200px`: the column is capped.

## Extending

Derive a new component from the nearest one in the specimen. A container starts from `.ds-panel`, something pressed or filled in from `.ds-form__input`, something the user presses from `.ds-button--secondary`. A picture is a dented tray in a `--color-fill-*` clay holding `.ds-clay` objects in other fills; text on a fill is `--color-accent-text` and nothing else. Never draw a shadow or a glow in a raw colour and never flatten a shape. Use tokens only; for a size with no token use `calc()` over `--space-*` or `--size-*`. Keep each text token on the fill it belongs to, as the list above says.
