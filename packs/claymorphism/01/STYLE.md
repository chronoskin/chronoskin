# Claymorphism landing

## Summary

Claymorphism is the soft 3D look of 2021 to 2025: every block is a thick, very round shape that seems modelled from coloured clay, with a light inner glow along its upper edge, a darker inner shadow along its lower edge and one large soft drop shadow under it. Unlike neumorphism the shapes are coloured and float above a pastel page instead of being pressed out of it, and simple clay objects (a ball, a ring, a cube, a pill, a round face) stand in for illustrations. This pack is a playful product landing page in that look: a floating clay navigation bar, a hero with clay objects, feature cards in four pastel clays, a pricing trio and a chunky footer block, set in a rounded display face over a friendly sans.

## Layout

- Content column: `--size-page` (1160px), centred, with `--space-5` (24px) side padding. The page gradient fills the rest of the window; the layout is fluid below that width.
- Navigation: `.ds-nav` is one floating clay bar `--size-nav` (72px) high, as wide as the column and 24px below the top edge: the brand at the left, the links centred (one of them a `.ds-menu`), a secondary and a primary button at the right.
- Home page: `.ds-page__sections` stacks full-width sections `--space-8` (88px) apart: the hero (text left, clay objects right), the `.ds-stat` bar, the `.ds-grid` of four feature cards under a centred `.ds-page__title`, a `.ds-page__split` of two even columns (carousel of quotes beside a panel) and the closing `.ds-page__band`.
- Inner pages have no hero. `--space-6` (32px) under the navigation come the `.ds-breadcrumb` and the `.ds-page-header` (title at `--text-h1`, one lead line, actions at the right) on the bare page. A pricing page then stacks, 32px apart, the centred `.ds-tabs` with a `.ds-switch`, the `.ds-price` trio and a `.ds-page__split--aside` (table left, a `--size-aside` 380px column of the accordion right). A form page uses `.ds-page__split--aside`: the form in one `.ds-panel`, notices and the dialog in the aside. A text page uses `.ds-page__split--side`: a `--size-side` (264px) column with the `.ds-sidebar` and the empty state at the left, the prose panel, the list and the pagination at the right. Keep the two columns about equally long.
- Views. The specimen is one site with four views, each `<section class="ds-view" id="...">` inside `.ds-page__inner`; navigation and footer stand outside them. `home` is the landing page, `pricing` the plans, `signup` the form, `guide` the help text. The link of the current view is pressed into the bar by one `:has()` rule per view.
- Spacing scale: `--space-1` 4px (dot to text), `--space-2` 8px (label to control, between stacked links), `--space-3` 12px (icon to text, between list bars), `--space-4` 16px (between buttons, row padding), `--space-5` 24px (card padding, grid gap, form gap), `--space-6` 32px (between panels, dialog padding), `--space-7` 48px (band and footer padding, column gap of a split), `--space-8` 88px (between home sections, above the footer).
- Clay needs air: never put two raised shapes closer than `--space-3`, and give every card at least `--space-5` of padding so that the inner glow does not touch the text.
- `--size-rule` (2px) is the weight of the soft rules between table rows and stat figures.
- Below 960px the hero, the splits and the pricing trio stack, the grid goes to two columns and the navigation links wrap to a second line inside the bar. Below 560px the grid and stats are one column, the hero title takes `--text-h1`, form rows stack, the tabs fill the width and table columns marked `.ds-table__opt` are hidden, so nothing scrolls sideways.
- On a phone (560px and below) `.ds-nav__list` stays inside the bar as one row of equal cells (`flex: 1 1 auto`, no wrapping, text centred), under the brand and `.ds-nav__actions`; the current link stays pressed in. The More `.ds-menu` opens as a full row inset by `--space-3` under the bar instead of a card under its toggle. At every width a tooltip in a table opens at the left edge of its trigger, at most 60vw wide.

## Typography and colour roles

- `--font-heading` is a rounded display face (Fredoka, then Baloo 2, Arial Rounded) for titles, figures and the brand; `--font-body` and `--font-ui` are a soft sans (Nunito, then Avenir Next). `--font-mono` is for code only. The references used Fredoka One and commercial geometric rounded faces; these stacks are the closest open ones.
- Sizes: `--text-base` 17px, `--text-ui` 16px for controls, navigation, tables and notices, `--text-small` 14px for meta text, labels and badges, `--text-large` 20px for lead text and large buttons, `--text-h3` 20px, `--text-h2` 28px, `--text-h1` 40px, `--text-display` 64px for the hero title and the prices only.
- Weights: 500 body, 700 for headings, controls, labels and links. Line height 1.6 for body, 1.25 for headings, 1.05 for display.
- These are the era's conventions for which text token sits on which fill; every layout of the era keeps to them:
  - `--color-text`, `--color-text-muted`, `--color-heading`, `--color-heading-alt` and the link colours stand only on the page, on `--color-surface`, `--color-surface-alt` and `--color-surface-strong`.
  - `--color-accent-text` is the ink for the accent and for all four `--color-fill-*` clays. A palette keeps the accent and the fills in one lightness band so that this one ink reads on all five. Meta text on a fill is the same ink at opacity 0.8, never `--color-text-muted`.
  - `--color-bar-text` on `--fill-bar` (the navigation), `--color-bar-alt-text` on `--fill-bar-alt` (the footer block, table head), `--color-inverse-text` on `--fill-inverse` (the closing band, code blocks, tooltips), `--color-button-text` on `--fill-button`, `--color-button-secondary-text` on `--fill-button-secondary` (secondary buttons, pagination keys, round icon balls, switch thumbs), `--color-input-text` on `--fill-input` (controls, the tab track, progress tracks, the empty state), `--color-notice-text` on `--color-notice`, `--color-danger` on `--color-danger-surface` (error notice, destructive button).
  - Status colours never carry text: a status badge is the surface tinted 22% with the status colour, a dot of the full colour and `--color-text`.
- The clay itself comes from the surface part: every shadow token holds the drop shadow in `--color-shadow` (a tint of the palette, not grey) and two inset shadows mixed from white and black, so the same tokens shape a white card, a pastel card and a dark band. `--shadow-control-pressed` has only the insets, reversed: it is a dent.
- Coloured shapes (grid cards, clay objects) add a second drop shadow in their own colour: `color-mix()` of their fill token and `transparent`, written in the component because the vocabulary has no token per fill.
- `--fill-page` is a pastel gradient with soft blobs mixed from `--color-fill-1` to `--color-fill-3`.
- One-off values with no token: circles use `border-radius: 50%`, the cube 28% and icon tiles 34%; a list bar, a notice and an accordion bar take the smaller of `--radius-panel` and `--radius-control`; placeholders are `--color-input-text` at opacity 0.55.

## Components

- `.ds-page`: on `<body>`; paints `--fill-page` and sets the base type. `.ds-page__inner` is the width container, `.ds-page__sections` the stack of home sections, `.ds-page__stack` a vertical stack 32px apart, `.ds-page__split` two columns (`--aside` and `--side` for the inner pages), `.ds-page__title` a centred section heading with `.ds-page__title-note`, `.ds-page__center` a centred row, `.ds-page__band` the closing inverted block with `.ds-page__band-title`, `.ds-page__band-text` and `.ds-page__band-art`. `.ds-view` is one screen of the specimen; an application with one document per page does not need it.
- `.ds-brand`: the site's mark and name in the header: `.ds-brand__mark` is a clay ball `--size-brand` (44px) dressed like the primary button, holding the mark as inline SVG; `.ds-brand__name` is the name in the heading font. It inherits its colour from the bar it stands on. The name `chronoskin` and its mark are placeholders: replace them with the installing project's own and keep the ball.
- `.ds-nav`: the floating bar. `.ds-nav__inner` is the clay, `.ds-nav__list` the links, `.ds-nav__link` one link (pressed in when `is-current`, raised on `is-hover`), `.ds-nav__actions` the buttons at the right.
- `.ds-menu`: a `<details>` drop-down under a navigation item: `.ds-menu__toggle` on the summary, `.ds-menu__list` the clay card that opens under it, `.ds-menu__link` with an optional `.ds-menu__note`. `--end` aligns it right, `--up` opens it upwards.
- `.ds-hero`: the page introduction on the bare page: a badge, `.ds-hero__title` at the display size, `.ds-hero__lead`, `.ds-hero__actions`, `.ds-hero__proof` with overlapping avatars in `.ds-hero__faces`. `.ds-hero__art` holds the clay objects, each a `.ds-clay` placed with `.ds-hero__obj--a` to `--f`, and one floating `.ds-hero__card` with `.ds-hero__card-title`.
- `.ds-clay`: a clay object drawn in CSS. Plain it is a ball in `--color-fill-1`; `--1` to `--4` pick the fill, `--tiny`, `--small` and `--large` the size, `--cube`, `--pill` and `--ring` the shape. A face is a ball holding two `.ds-clay__eye`, a `.ds-clay__smile` and two `.ds-clay__cheek`. A ring's hole takes `--color-page`: place rings on the bare page or set the hole to the fill behind it. Decorative only: mark them `aria-hidden`.
- `.ds-page-header`: head of an inner page: `.ds-page-header__text` with `.ds-page-header__title` and `.ds-page-header__lead`, `.ds-page-header__actions` at the right with at most one primary button.
- `.ds-prose`: running text inside a panel: h1 to h3, paragraphs, lists, inline `code`; `.ds-prose__lead` for a lead paragraph, `.ds-prose__code` for an inverted code block.
- `.ds-link`: link in running text; states `is-visited`, `is-hover`, `is-active`, `is-focus`; `.ds-link--quiet` for titles and table links; `.ds-link__sample` lays out a row of links.
- `.ds-button`: a fat clay pill 48px high. `is-hover` lifts it, `is-active` squashes it (it scales to 97% by 93% and takes `--shadow-control-pressed`). `.ds-button--secondary` is the same key in the surface colour, `.ds-button--danger` the destructive action on `--color-danger-surface`, `.ds-button--large` (60px) for the hero, `.ds-button--wide` fills its card; `disabled` or `is-disabled` is flat. `.ds-button__row` spaces a group. One primary button per group.
- `.ds-form`: vertical form in a panel. `.ds-form__field` wraps a `.ds-form__label` and a control: `.ds-form__input`, `.ds-form__textarea`, `.ds-form__select` (in `.ds-form__select-wrap` with a `.ds-form__chevron`). Controls are dents: `--fill-input` with `--shadow-control-pressed`. `.ds-form__check` is a label row with a `.ds-form__checkbox` (a raised clay tile, accent with a tick when checked). `is-error` rings the control in `--color-danger` and is followed by `.ds-form__error`; `.ds-form__hint` is help text, `.ds-form__actions` holds the buttons, `.ds-form__row` puts two fields side by side.
- `.ds-switch`: a fat clay toggle over a checkbox: `.ds-switch__input`, `.ds-switch__track` (a dent that fills with the accent when on) with a clay ball as thumb, `.ds-switch__label` and `.ds-switch__hint`; `--end` puts the track at the right.
- `.ds-table`: one clay sheet with a head in `--color-bar-alt` and rows divided by a soft rule. `.ds-table__num` right-aligns numbers, `.ds-table__meta` is muted, `.ds-table__strong` a row title, `.ds-table__opt` a column dropped on a phone, `is-hover` tints a row.
- `.ds-list`: items as separate clay bars 12px apart: `.ds-list__item` holds a `.ds-list__body` (`.ds-list__title`, `.ds-list__meta`) and a badge, button or `.ds-list__aside` figure.
- `.ds-panel`: the basic clay card: `.ds-panel__title`, `.ds-panel__body` (a column with 16px gaps), `.ds-panel__note` for small muted text, `.ds-panel__line` for a row with two ends. `.ds-panel--accent` is the card in the accent clay, with all text in `--color-accent-text`.
- `.ds-price`: the pricing trio, three `.ds-panel` cards with a `.ds-price__amount` (display size, `.ds-price__per` after it) and a `.ds-price__list`; the middle one is `.ds-panel--accent` and taller.
- `.ds-grid`: four feature cards. `.ds-grid__cell` with `--1` to `--4` takes a `--color-fill-*` clay and its own coloured drop shadow; inside, a `.ds-grid__icon` ball, `.ds-grid__title`, `.ds-grid__text` and an underlined `.ds-grid__more`. Every second card sits 24px lower.
- `.ds-stat`: the row; one long clay bar holding three `.ds-stat__item` figures divided by rules, each a `.ds-stat__value` (40px) over a `.ds-stat__label`, with an optional `.ds-stat__meta`.
- `.ds-tabs`: a dented track of `.ds-tabs__tab` pills; `is-current` is a raised pill in the accent. For switching what a page shows, not for site navigation.
- `.ds-badge`: a pill in the accent. `.ds-badge--quiet` is neutral; `.ds-badge--success`, `.ds-badge--warning` and `.ds-badge--danger` are tinted pills with a status dot.
- `.ds-sidebar`: a clay card of topics: `.ds-sidebar__heading`, `.ds-sidebar__list`, `.ds-sidebar__link`, dented when `is-current`.
- `.ds-notice`: a raised bar in `--color-notice` with `.ds-notice__icon` and `.ds-notice__title`; `.ds-notice--error` is the error. `.ds-notice__stack` stacks several.
- `.ds-pagination`: round clay keys, `.ds-pagination__link`; the current page is `is-current`, in the accent.
- `.ds-breadcrumb`: small bold trail above the page header: `.ds-breadcrumb__item`, `.ds-breadcrumb__link`.
- `.ds-dialog`: a tinted sheet (`--color-overlay`) with a centred `.ds-dialog__box` under the deepest shadow: `.ds-dialog__title`, `.ds-dialog__body`, `.ds-dialog__actions`.
- `.ds-empty`: a dent holding a clay ball `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one primary button.
- `.ds-tooltip`: wraps a trigger; `.ds-tooltip__tip` is an inverted pill shown on hover or focus (`is-open` shows it statically).
- `.ds-progress`: a thick dented track with a puffy `.ds-progress__bar`; width by `--10`, `--25`, `--40`, `--60`, `--75`, `--90`; `.ds-progress__legend` is the line above it. `.ds-progress--ring` is a ring around a clay ball `.ds-progress__face`.
- `.ds-avatar`: initials on a clay ball `.ds-avatar__pic`, alone or with `.ds-avatar__text` (`.ds-avatar__name`, `.ds-avatar__meta`); `--1` to `--4` pick a fill.
- `.ds-accordion`: `<details>` sections as separate clay bars: `.ds-accordion__item`, `.ds-accordion__head` with a round `.ds-accordion__chevron`, `.ds-accordion__body`.
- `.ds-carousel`: `.ds-carousel__track` scrolls and snaps, one `.ds-carousel__slide` showing; `.ds-carousel__foot` holds two round `.ds-carousel__arrow` keys around `.ds-carousel__dots` (`.ds-carousel__dot`, `is-current`). `.ds-quote` with `.ds-quote__text` is the quote card used as a slide.
- `.ds-footer`: one chunky clay block in `--fill-bar-alt` that the page ends on: `.ds-footer__inner` (four columns), `.ds-footer__about` with the brand and `.ds-footer__text`, `.ds-footer__heading`, `.ds-footer__links`, `.ds-footer__link`.

## Never

- `box-shadow != none`: nothing raised is flat; a block without its inner glow and drop shadow is not clay.
- `border-width <= 3px`: clay has no outlines; the only lines are 2px rules and the 3px focus or error ring.
- `box-shadow-blur <= 80px`: the deepest shadow, under the dialog, blurs over 80px.
- `text-shadow = none`: the text is flat; only shapes are modelled.
- `font-size <= 64px`: the hero title and the prices are the largest text.
- `font-size >= 14px`: meta text and badges are the smallest text.
- `font-weight >= 500`: the type is round and sturdy; nothing is light.
- `font-families <= 3`: one display face, one sans and one monospace.
- `line-height <= 1.65`: body text is set at 1.6.
- `letter-spacing = 0`: no tracking on any text.
- `uppercase-text <= 5%`: nothing is set in capitals.
- `underlined-links <= 25%`: links are marked by colour and weight; only the small links on coloured cards are underlined.
- `block-gap <= 88px`: 88px between home sections is the largest gap.
- `content-width <= 1160px`: the column is capped.

## Extending

Derive a new component from the nearest one in the specimen. Decide first whether it is raised or dented: a container starts from `.ds-panel` (fill, radius and `--shadow-panel` together), something pressed or filled in starts from `.ds-form__input` (`--fill-input`, `--color-input-text`, `--shadow-control-pressed`), something the user presses starts from `.ds-button--secondary`. For a coloured card take a `--color-fill-*` background, `--color-accent-text` for everything on it, and add the coloured drop shadow as `.ds-grid__cell` does. Never draw a shadow in a raw colour and never flatten a shape: use the shadow tokens as they are. Use tokens only; for a size with no token use `calc()` over `--space-*` or `--size-*`. Keep each text token on the fill it belongs to.
