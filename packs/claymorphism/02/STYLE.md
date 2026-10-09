# Claymorphism app

## Summary

Claymorphism is the soft 3D look of 2021 to 2025: thick, very round shapes that seem modelled from coloured clay, each with a light inner glow from the upper left, a darker inner shadow at the lower right and a soft drop shadow in a tint of the palette. It was the look of playful learning, habit and money apps, where chunky keys, fat progress bars and a round mascot made a daily task feel like a toy. This pack is such an app: a clay rail at the left, a greeting block with a clay character, tiles with one big figure each, lesson cards and thick rings and bars, set in a heavy geometric display face over a plain sans.

## Layout

- The window is two columns: the rail `--size-rail` (264px) at the left, sticking to the top, and the working column, capped at `--size-page` (1080px) with `--space-6` (28px) side padding and centred in what remains. The layout is fluid below that width.
- Navigation: `.ds-nav` is the rail, one tall clay card: the brand, four chunky links with icons, each `--size-control-large` (56px) high, and a small streak card at its foot. Over every view the working column starts with `.ds-page__top`: the search well at the left, the streak badge and the account `.ds-menu` at the right.
- Home page: a `.ds-page__stack` at 28px gaps: the hero (one inverted clay block, greeting left, clay character right), the `.ds-grid` of four figure tiles, then `.ds-page__cols`: the list of lesson cards under a `.ds-page__heading` in the wide column and a `--size-aside` (340px) column with the weekly panel and the `.ds-sidebar` of courses.
- Inner pages have no hero. They start with the `.ds-breadcrumb` and the `.ds-page-header` (title at `--text-h1`, one lead line, one action at the right) on the bare page. A list page stacks the `.ds-stat` row, the `.ds-tabs` and `.ds-page__cols`: table and pagination in the wide column, the empty state and a small panel in the aside. A practice page uses `.ds-page__cols`: the progress bar, the card carousel, the answer keys and the action row at the left, a prose panel and the dialog at the right. A settings page uses `.ds-page__cols`: the form in a `.ds-panel`, notices, a panel of switches and the accordion in the aside. Keep the two columns about equally long.
- Views. The specimen is one app with four views, each `<section class="ds-view" id="...">` inside `.ds-page__inner`; the rail, the top strip and the footer stand outside them. `today` is the home view, `lessons` the list, `lesson` the practice screen, `settings` the form. The rail link of the current view is pressed in by one `:has()` rule per view.
- Spacing scale: `--space-1` 4px, `--space-2` 8px (between rail links, label to control), `--space-3` 12px (icon to text, between list cards), `--space-4` 16px (between buttons, card padding of a list item), `--space-5` 20px (tile and panel padding, grid gap), `--space-6` 28px (between blocks and columns), `--space-7` 40px (hero padding), `--space-8` 56px (above the footer).
- `--size-rule` (2px) is the weight of the rules between table rows.
- Below 1100px the aside drops under the working column and the tiles go two across. Below 900px the rail becomes a block at the top of the page: the brand over a chunky tab bar of four equal keys with the icon above the label; `.ds-nav__card` is not shown there and the current key stays pressed in. Below 560px the hero stacks, stats are one column, form rows and answer keys stack, the tabs fill the width and table columns marked `.ds-table__opt` are hidden, so nothing scrolls sideways.

## Typography and colour roles

- `--font-heading` and `--font-ui` are one heavy geometric sans (Poppins, then Montserrat, Avenir Next) for titles, figures, buttons and navigation; `--font-body` is a plain sans (DM Sans, then Helvetica Neue). `--font-mono` is for code only. The references used commercial rounded grotesques; these stacks are the closest open ones.
- Sizes: `--text-base` 16px, `--text-ui` 14px for controls and navigation, `--text-small` 13px for meta text, labels and badges, `--text-large` 19px for lead text and large keys, `--text-h3` 19px, `--text-h2` 30px, `--text-h1` 44px for page titles and tile figures. `--text-display` (72px) is not used by this layout: the greeting is set at `--text-h1`.
- Weights: 400 body, 600 bold, 800 for headings, buttons and navigation, 900 for the greeting and tile figures. Buttons, navigation, tabs and table heads are uppercase through `--ui-transform`; headings are tracked at -0.02em.
- Text tokens sit on fills as the era's first layout sets out: body, muted, heading and link colours only on the page and the three surface colours; `--color-accent-text` on the accent and on every `--color-fill-*` clay (tiles, lesson numbers, flashcards, avatars), with meta text on a fill at opacity 0.8; `--color-bar-text` on `--fill-bar` (the rail); `--color-bar-alt-text` on `--fill-bar-alt` (the footer strip) and on `--color-bar-alt` (table head); `--color-inverse-text` on `--fill-inverse` (the hero block, code, tooltips); `--color-button-text` on `--fill-button`; `--color-button-secondary-text` on `--fill-button-secondary` (secondary keys, the account pill, icon tiles, the flashcard's sign ball, switch thumbs); `--color-input-text` on `--fill-input`; `--color-notice-text` on `--color-notice`; `--color-danger` on `--color-danger-surface`. Status colours never carry text.
- The clay comes from the surface part: each shadow token is a drop shadow in `--color-shadow` towards the lower right and two inset shadows mixed from black and white. Coloured tiles add a second drop shadow in their own fill through `color-mix()` with `transparent`, written in the component.
- `--shadow-text` puts a 2px light edge under the text of primary buttons.
- One-off values with no token: circles use `border-radius: 50%`, icon tiles 34%, the cube 28%; rail links, list cards, notices and accordion bars take the smaller of `--radius-panel` and `--radius-control`; placeholders are `--color-input-text` at opacity 0.55.

## Components

- `.ds-page`: on `<body>`; the two-column frame, painted with `--fill-page`. `.ds-page__inner` is the working column, `.ds-page__top` the strip over every view with `.ds-page__search` and `.ds-page__top-end`, `.ds-page__account` the account pill (`.ds-page__account-name` is hidden on a phone), `.ds-page__stack` a vertical stack, `.ds-page__cols` the wide column with an aside, `.ds-page__heading` a section heading with a link at its right end. `.ds-view` is one screen of the specimen; an application with one document per page does not need it.
- `.ds-brand`: the mark and name at the top of the rail: `.ds-brand__mark` is a clay key `--size-brand` (42px) dressed like the primary button, holding the mark as inline SVG; `.ds-brand__name` is the name in the heading font, in the colour of the bar. The name `chronoskin` and its mark are placeholders: replace them with the installing project's own and keep the key.
- `.ds-nav`: the rail. `.ds-nav__inner` is the clay card, `.ds-nav__list` the links, `.ds-nav__link` one chunky link with an icon (pressed in when `is-current`, raised on `is-hover`), `.ds-nav__card` with `.ds-nav__card-title` the streak card at its foot.
- `.ds-menu`: the account drop-down, a `<details>`: `.ds-menu__toggle` on the summary, `.ds-menu__list` the clay card, `.ds-menu__link` with an optional `.ds-menu__note`; `--end` aligns it right, `--up` opens it upwards.
- `.ds-hero`: the greeting on one inverted clay block: `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__actions`, and `.ds-hero__art` holding `.ds-clay` objects. Once per app, on the home view.
- `.ds-clay`: a clay object drawn in CSS: a ball; `--1` to `--4` pick the fill, `--tiny`, `--small`, `--large` the size, `--cube`, `--pill`, `--ring` the shape; a face holds two `.ds-clay__eye`, a `.ds-clay__smile` and two `.ds-clay__cheek`. Decorative only.
- `.ds-page-header`: head of an inner page: `.ds-page-header__text` with `.ds-page-header__title` and `.ds-page-header__lead`, `.ds-page-header__actions` at the right.
- `.ds-prose`: running text inside a panel: h1 to h3, paragraphs, lists, inline `code`; `.ds-prose__lead`, `.ds-prose__code`.
- `.ds-link`: link in running text, underlined; states `is-visited`, `is-hover`, `is-active`, `is-focus`; `.ds-link--quiet` for table links; `.ds-link__sample` lays out a row of links.
- `.ds-button`: a fat clay key 46px high with uppercase text. `is-hover` lifts it, `is-active` squashes it into a dent. `.ds-button--secondary` in the surface colour, `.ds-button--danger` on `--color-danger-surface`, `.ds-button--large` (56px), `.ds-button--wide`; `disabled` or `is-disabled` is flat. `.ds-button__row` spaces a group.
- `.ds-form`: vertical form in a panel: `.ds-form__field`, `.ds-form__label`, `.ds-form__input`, `.ds-form__textarea`, `.ds-form__select` (in `.ds-form__select-wrap` with `.ds-form__chevron`), `.ds-form__check` with `.ds-form__checkbox`, `.ds-form__error` after an `is-error` control, `.ds-form__hint`, `.ds-form__actions`, `.ds-form__row`. Controls are dents.
- `.ds-switch`: a fat clay toggle over a checkbox: `.ds-switch__input`, `.ds-switch__track`, `.ds-switch__label`, `.ds-switch__hint`; `--end` puts the track at the right. Use it for settings that apply at once.
- `.ds-table`: one clay sheet with a head in `--color-bar-alt`: `.ds-table__num`, `.ds-table__meta`, `.ds-table__strong`, `.ds-table__opt`, `is-hover` on a row.
- `.ds-list`: lesson cards 12px apart. `.ds-list__item` is a clay card holding a `.ds-list__step` ball with the lesson number (`--2` to `--4` pick a fill), a `.ds-list__body` (`.ds-list__title`, `.ds-list__meta`) and a badge or button; `is-locked` presses the card in. `.ds-list__aside` is a figure at the right.
- `.ds-panel`: the basic clay card: `.ds-panel__title`, `.ds-panel__body`, `.ds-panel__note`, `.ds-panel__line`; `.ds-panel--accent` in the accent clay. `.ds-panel__goal` sets a ring beside `.ds-panel__bars`; `.ds-panel__rank` and `.ds-panel__points` are a row of a board.
- `.ds-grid`: four figure tiles. `.ds-grid__cell` with `--1` to `--4` takes a `--color-fill-*` clay and its own coloured drop shadow; inside, a `.ds-grid__icon` tile, a `.ds-grid__value` (44px), a `.ds-grid__label` and a `.ds-grid__meta`.
- `.ds-stat`: the row of three; each `.ds-stat__item` is a clay card with a `.ds-stat__label` over a `.ds-stat__value` and an optional `.ds-stat__meta`. Used on inner pages; the grid is the home view's form of it.
- `.ds-tabs`: a dented track of `.ds-tabs__tab` pills; `is-current` is raised, in the accent.
- `.ds-badge`: a pill in the accent; `.ds-badge--quiet` neutral; `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger` tinted with a status dot.
- `.ds-sidebar`: a clay card of courses in the aside: `.ds-sidebar__heading`, `.ds-sidebar__list`, `.ds-sidebar__link` (a pill, dented when `is-current`).
- `.ds-notice`: a raised bar in `--color-notice` with `.ds-notice__icon` and `.ds-notice__title`; `.ds-notice--error`; `.ds-notice__stack`.
- `.ds-pagination`: clay keys, `.ds-pagination__link`, the current one `is-current` in the accent.
- `.ds-breadcrumb`: small bold trail above the page header: `.ds-breadcrumb__item`, `.ds-breadcrumb__link`.
- `.ds-dialog`: a tinted sheet with a centred `.ds-dialog__box`: `.ds-dialog__title`, `.ds-dialog__body`, `.ds-dialog__actions`.
- `.ds-empty`: a dent holding a clay ball `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one primary button.
- `.ds-tooltip`: wraps a trigger; `.ds-tooltip__tip` shows on hover or focus.
- `.ds-progress`: a thick dented track with a puffy `.ds-progress__bar` (`--10` to `--90`), `.ds-progress__legend` above it; `.ds-progress--ring` is a ring around a clay ball `.ds-progress__face`.
- `.ds-avatar`: initials on a clay ball `.ds-avatar__pic`, with optional `.ds-avatar__text` (`.ds-avatar__name`, `.ds-avatar__meta`); `--1` to `--4` pick a fill.
- `.ds-accordion`: `<details>` sections as separate clay bars: `.ds-accordion__item`, `.ds-accordion__head`, `.ds-accordion__chevron`, `.ds-accordion__body`.
- `.ds-carousel`: `.ds-carousel__track` with one `.ds-carousel__slide` showing, `.ds-carousel__foot` with two `.ds-carousel__arrow` keys around `.ds-carousel__dots` (`.ds-carousel__dot`). `.ds-card` is the flashcard used as a slide (`--2`, `--4` pick a fill): a `.ds-card__sign` ball, `.ds-card__title`, `.ds-card__text`; `.ds-card__answers` is the grid of answer keys under it.
- `.ds-footer`: one flat clay strip in `--fill-bar-alt` at the end of the working column: `.ds-footer__inner`, `.ds-footer__links`, `.ds-footer__link`.

## Never

- `box-shadow != none`: every key, card and tile is modelled; nothing raised is flat.
- `border-width <= 4px`: clay has no outlines; the only lines are 2px rules and the 4px focus or error ring.
- `border-radius <= 36px`: keys are 20px, cards 28px, the hero and flashcards 36px; only pills and balls go higher.
- `box-shadow-blur <= 44px`: the shadows are short and firm; the dialog's is the deepest.
- `font-size <= 44px`: page titles and tile figures are the largest text.
- `font-size >= 13px`: meta text and badges are the smallest text.
- `font-weight >= 400`: nothing is light.
- `font-families <= 3`: one display face, one sans and one monospace.
- `line-height <= 1.6`: body text is set at 1.55.
- `uppercase-text <= 30%`: capitals are for buttons, navigation, tabs and table heads only.
- `block-gap <= 56px`: blocks sit 28px apart; 56px above the footer is the largest gap.
- `content-width <= 1140px`: the working column is capped.

## Extending

Derive a new component from the nearest one in the specimen. A container starts from `.ds-panel`, something pressed or filled in from `.ds-form__input`, something the user presses from `.ds-button--secondary`. For a coloured tile take a `--color-fill-*` background, `--color-accent-text` for everything on it, and add the coloured drop shadow as `.ds-grid__cell` does. Never draw a shadow in a raw colour and never flatten a shape. Use tokens only; for a size with no token use `calc()` over `--space-*` or `--size-*`. Keep each text token on the fill it belongs to, as the list above says.
