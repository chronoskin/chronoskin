# Claymorphism home panel

## Summary

Claymorphism is the soft 3D look of 2021 to 2025: thick, very round shapes that seem modelled from coloured clay, with a light inner glow along the upper edge, a darker inner shadow along the lower edge and one large soft shadow under them. Around 2022 it was the look of smart-home, music and car control panels, usually at night: deep coloured clay in which fat toggles, dials and sliders sit like toys. This pack is such a panel for a house: a capsule of round keys at the left, a board of cards of mixed widths, device cards in four warm clays with a fat switch each, a temperature dial and a bar chart of pressed wells, set in a geometric display face over a plain sans.

## Layout

- The working column is `--size-page` (1200px) wide. To its left stands the rail, `--size-rail` (92px) wide and `--space-5` (22px) away; rail and column together are centred in the window. The layout is fluid below that width.
- Navigation: `.ds-nav` is a strip on the bare page across the top (brand, the house as a `.ds-menu`, at the right a bell key with a `.ds-tooltip` and the `.ds-avatar` of the person) and, hanging under it at the left, `.ds-nav__rail`: one tall clay capsule of five round keys, each with its label under it. The key of the current view is in the accent clay.
- Home page: `.ds-bento` is a board of twelve columns with `--space-5` gaps. A child spans all twelve unless it carries `.ds-bento__3`, `__4`, `__5` or `__7`; the cards of a row are equally tall. The rows are: the hero (7) beside the climate panel with the dial (5); the three `.ds-stat` tiles; a `.ds-page__title`; the `.ds-grid` of four device cards; the chart panel (5), the scenes carousel (4) and the people at home (3).
- Inner pages have no hero. `--space-6` (30px) under the strip come the `.ds-breadcrumb` and the `.ds-page-header` (title at `--text-h1`, one lead line, the action at the right) on the bare page. Below them a `.ds-page__stack`: a `.ds-page__tools` row with the `.ds-tabs` at the left and a `.ds-switch` at the right, then `.ds-page__split`, a working column with a `--size-aside` (360px) column at the right (devices and sliders beside the room's air and an empty state; the table beside the budget; the form beside notices and the dialog). The help page uses `.ds-page__split--side`: a `--size-side` (256px) sidebar at the left, the text panel and the accordion at the right. Keep the two columns about equally long.
- Views. The specimen is one site with five views, each `<section class="ds-view" id="...">` inside `.ds-page__inner`; the strip, the rail and the footer stand outside them. `home` is the board, `rooms` one room, `energy` the week's use, `routines` the form for a routine, `help` the help text. One `:has()` rule per view colours the key of the current view.
- Spacing scale: `--space-1` 4px (key to its label, between title and meta), `--space-2` 8px (label to control, between chart wells), `--space-3` 12px (icon to text, between list bars and between buttons), `--space-4` 16px (between rail keys, inside a row), `--space-5` 22px (card padding, every gap of the board), `--space-6` 30px (hero padding, under the strip, dialog padding), `--space-7` 44px (above the footer), `--space-8` 60px (not used on the board; the largest step for a new section break).
- `--size-rule` (2px) is the weight of the soft rules between table rows.
- Below 1180px the spans of 3 and 4 become half the board and those of 5 and 7 the whole, and the grid goes to two columns. Below 960px the rail lies down as a row of keys under the strip and the splits stack. Below 560px everything is one column, the hero stacks with its art under the text, the person's name is dropped from the strip, tabs fill the width and table columns marked `.ds-table__opt` are hidden, so nothing scrolls sideways.
- On a phone (560px and below) `.ds-nav__rail` is a floating dock: a capsule fixed `--space-3` above the bottom edge of the window, `--space-4` from each side, with a pill radius, holding the row of keys of `.ds-nav__list`. The `.ds-nav__home` menu takes a full row under the brand and its list opens as a full row; the person's name in `.ds-nav__actions` is not shown. The current key stays pressed in.

## Typography and colour roles

- `--font-heading` is a geometric display face (Sora, then Space Grotesk, Trebuchet MS) for titles, figures and the brand; `--font-body` and `--font-ui` are a plain sans (Manrope, then the system face). `--font-mono` is for code only.
- Sizes: `--text-base` 16px, `--text-ui` 15px for controls, tables and notices, `--text-small` 13px for meta text, labels, badges and the labels of the rail, `--text-large` 18px for lead text and large buttons, `--text-h3` 18px, `--text-h2` 24px, `--text-h1` 34px for the hero and page titles, `--text-display` 54px for the figure on the dial only.
- Weights: 400 body, 600 headings and controls, 700 bold. Line height 1.5 for body, 1.3 for headings, 1.15 for the dial.
- The era's conventions for which text token sits on which fill are kept:
  - `--color-text`, `--color-text-muted`, `--color-heading`, `--color-heading-alt` and the link colours stand only on the page, on `--color-surface`, `--color-surface-alt` and `--color-surface-strong`. The strip is on the bare page, so the brand there takes `--color-heading`.
  - `--color-accent-text` is the ink for the accent and for all four `--color-fill-*` clays (device cards, stat icons, avatars, the current key). Meta text on a fill is the same ink at opacity 0.8.
  - `--color-bar-text` on `--fill-bar` (the rail), `--color-bar-alt-text` on `--fill-bar-alt` (the footer strip, a scene slide, table head), `--color-inverse-text` on `--fill-inverse` (tooltips, code blocks), `--color-button-text` on `--fill-button`, `--color-button-secondary-text` on `--fill-button-secondary` (secondary buttons, round keys, icon balls, switch and slider thumbs, the house key), `--color-input-text` on `--fill-input` (controls, the tab track, the empty state), `--color-notice-text` on `--color-notice`, `--color-danger` on `--color-danger-surface`.
  - Status colours never carry text: a status badge is the surface tinted with the status colour, a dot of the full colour and `--color-text`.
- The clay comes from the surface part: every shadow token holds a drop shadow in `--color-shadow` and inset shadows mixed from white and black (here a thin bright rim and a soft glow above, a deep shade below), so the same tokens shape a dark card and a bright one. `--shadow-control-pressed` has only insets, reversed: it is a dent. Tracks, wells, the hero's tray and the empty state are dents in `--fill-input`.
- Coloured shapes (device cards, the lamp's shade) add a second drop shadow in their own colour: `color-mix()` of their fill token and `transparent`, written in the component because the vocabulary has no token per fill.
- A switch on a coloured card fills with `--fill-inverse` when on, not with the accent, because the accent may equal the card's fill.
- One-off values with no token: circles use `border-radius: 50%`, stat icon tiles 34%, the lamp's shade and the speaker their own percentages; a list bar, a notice and an accordion bar take the smaller of `--radius-panel` and `--radius-control`; placeholders are `--color-input-text` at opacity 0.55.

## Components

- `.ds-page`: on `<body>`; paints `--fill-page` and sets the base type. `.ds-page__inner` is the working column beside the rail, `.ds-page__stack` a vertical stack 22px apart, `.ds-page__split` a working column with an aside (`--side` for a sidebar at the left), `.ds-page__tools` a row with two ends, `.ds-page__title` a small heading on the bare page with a `.ds-page__more` link at its right. `.ds-view` is one screen of the specimen; an application with one document per page does not need it.
- `.ds-bento`: the home board; `.ds-bento__3`, `__4`, `__5`, `__7` set how many of the twelve columns a card takes.
- `.ds-brand`: the site's mark and name in the strip: `.ds-brand__mark` is a clay tile `--size-brand` (44px) dressed like the primary button, holding the mark as inline SVG; `.ds-brand__name` is the name in the heading font. The name `chronoskin` and its mark are placeholders: replace them with the installing project's own and keep the tile.
- `.ds-nav`: the strip and the rail. `.ds-nav__home` is the key that names the house (the toggle of a `.ds-menu`), `.ds-nav__actions` the right end, `.ds-nav__rail` the capsule, `.ds-nav__list` its keys, `.ds-nav__link` one item (a round `.ds-nav__key` over its label; `is-current` puts the key in the accent, `is-hover` lifts it).
- `.ds-key`: a round clay key with one icon: the bell, the dial's minus and plus. `is-hover` lifts it, `is-active` presses it.
- `.ds-menu`: a `<details>` drop-down: `.ds-menu__toggle` on the summary, `.ds-menu__list` the clay card under it, `.ds-menu__link` with an optional `.ds-menu__note`. `--end` aligns it right, `--up` opens it upwards.
- `.ds-hero`: the greeting on one clay card: `.ds-hero__text` with a badge, `.ds-hero__title`, `.ds-hero__lead` and `.ds-hero__actions`; `.ds-hero__art` is a pressed tray holding devices.
- `.ds-device`: a device drawn in clay, decorative. Plain it is a lamp built from `.ds-device__shade`, `.ds-device__stem` and `.ds-device__foot`; `.ds-device--speaker` is a speaker with two pressed cones. `.ds-clay` (ball, `--cube`, `--pill`, `--ring`, sizes and fills `--1` to `--4`) is there for further objects.
- `.ds-page-header`: head of an inner page: `.ds-page-header__text` with `.ds-page-header__title` and `.ds-page-header__lead`, `.ds-page-header__actions` at the right with at most one primary button.
- `.ds-prose`: running text inside a panel: h1 to h3, paragraphs, lists, inline `code`; `.ds-prose__lead` for a lead paragraph, `.ds-prose__code` for an inverted code block.
- `.ds-link`: link in running text; states `is-visited`, `is-hover`, `is-active`, `is-focus`; `.ds-link--quiet` for titles and table links; `.ds-link__sample` lays out a row of links.
- `.ds-button`: a clay key 46px high. `is-hover` lifts it, `is-active` squashes it. `.ds-button--secondary` is the same key in the secondary clay, `.ds-button--danger` the destructive action, `.ds-button--large` (58px) for the hero, `.ds-button--wide` fills its card; `disabled` or `is-disabled` is flat. `.ds-button__row` spaces a group. One primary button per group.
- `.ds-form`: vertical form in a panel. `.ds-form__field` wraps a `.ds-form__label` and a control: `.ds-form__input`, `.ds-form__textarea`, `.ds-form__select` (in `.ds-form__select-wrap` with a `.ds-form__chevron`). Controls are dents. `.ds-form__check` is a label row with a `.ds-form__checkbox`. `is-error` rings the control in `--color-danger` and is followed by `.ds-form__error`; `.ds-form__hint` is help text, `.ds-form__actions` holds the buttons, `.ds-form__row` puts two fields side by side.
- `.ds-switch`: the fat clay toggle over a checkbox, the panel's main control: `.ds-switch__input`, `.ds-switch__track` (a dent that fills when on) with a clay ball as thumb, `.ds-switch__label` and `.ds-switch__hint`; `--end` puts the track at the right. Without a visible label give the input an `aria-label`.
- `.ds-table`: one clay sheet with a head in `--color-bar-alt` and rows divided by a soft rule. `.ds-table__num` right-aligns numbers, `.ds-table__meta` is muted, `.ds-table__strong` a row title, `.ds-table__opt` a column dropped on a phone, `is-hover` tints a row.
- `.ds-list`: devices as separate clay bars: `.ds-list__item` holds a `.ds-list__icon` ball, a `.ds-list__body` (`.ds-list__title`, `.ds-list__meta`), a status badge and a switch; `.ds-list__aside` is a figure at the end.
- `.ds-panel`: the basic clay card: `.ds-panel__title`, `.ds-panel__body` (a column with 16px gaps), `.ds-panel__note` for small muted text, `.ds-panel__line` for a row with two ends. `.ds-panel--accent` is the card in the accent clay.
- `.ds-stat`: the row of three; each `.ds-stat__item` is a chunky tile with a `.ds-stat__icon` (a clay tile in a fill, `--2` and `--3` pick the others), a `.ds-stat__label` and a `.ds-stat__value` with an optional `.ds-stat__meta`.
- `.ds-grid`: four device cards. `.ds-grid__cell` with `--1` to `--4` takes a `--color-fill-*` clay and its own coloured drop shadow; inside, a `.ds-grid__icon` ball, `.ds-grid__title`, `.ds-grid__text` and a `.ds-grid__foot` with the state and the switch.
- `.ds-chart`: a bar chart from CSS: each `.ds-chart__col` is a pressed `.ds-chart__well` with a puffy `.ds-chart__bar` standing in it (height by `--30`, `--45`, `--60`, `--75`, `--90`) and its label; `is-current` puts a bar in the button colour.
- `.ds-dial`: a large `.ds-progress--ring` whose face shows a figure at the display size with a `.ds-dial__unit`, a `.ds-key` at either side.
- `.ds-tabs`: a dented track of `.ds-tabs__tab` pills; `is-current` is a raised pill in the accent. For switching what a page shows, not for site navigation.
- `.ds-badge`: a pill in the accent. `.ds-badge--quiet` is neutral; `.ds-badge--success`, `.ds-badge--warning` and `.ds-badge--danger` are tinted pills with a status dot.
- `.ds-sidebar`: a clay card of topics: `.ds-sidebar__heading`, `.ds-sidebar__list`, `.ds-sidebar__link`, dented when `is-current`.
- `.ds-notice`: a raised bar in `--color-notice` with `.ds-notice__icon` and `.ds-notice__title`; `.ds-notice--error` is the error. `.ds-notice__stack` stacks several.
- `.ds-pagination`: clay keys, `.ds-pagination__link`; the current page is `is-current`, in the accent.
- `.ds-breadcrumb`: small bold trail above the page header: `.ds-breadcrumb__item`, `.ds-breadcrumb__link`.
- `.ds-dialog`: a tinted sheet (`--color-overlay`) with a centred `.ds-dialog__box` under the deepest shadow: `.ds-dialog__title`, `.ds-dialog__body`, `.ds-dialog__actions`.
- `.ds-empty`: a dent holding a clay ball `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one primary button.
- `.ds-tooltip`: wraps a trigger; `.ds-tooltip__tip` is an inverted pill shown on hover or focus (`is-open` shows it statically). `--below` opens it under the trigger, for keys at the top of the window.
- `.ds-progress`: a thick dented track with a puffy `.ds-progress__bar`; width by `--10`, `--25`, `--40`, `--60`, `--75`, `--90`; `.ds-progress__legend` is the line above it. `.ds-progress--ring` is a ring around a clay ball `.ds-progress__face`. `.ds-progress--slider` puts a clay ball at the end of the bar: a slider shown at a fixed value.
- `.ds-avatar`: initials on a clay ball `.ds-avatar__pic`, alone or with `.ds-avatar__text` (`.ds-avatar__name`, `.ds-avatar__meta`); `--1` to `--4` pick a fill.
- `.ds-accordion`: `<details>` sections as separate clay bars: `.ds-accordion__item`, `.ds-accordion__head` with a round `.ds-accordion__chevron`, `.ds-accordion__body`.
- `.ds-carousel`: `.ds-carousel__track` scrolls and snaps, one `.ds-carousel__slide` showing; `.ds-carousel__foot` holds two round `.ds-carousel__arrow` keys around `.ds-carousel__dots` (`.ds-carousel__dot`, `is-current`). `.ds-scene` with `.ds-scene__title` and `.ds-scene__text` is the scene block used as a slide.
- `.ds-footer`: one slim clay strip in `--fill-bar-alt` at the end of the working column: `.ds-footer__inner` holds the brand, `.ds-footer__text` and `.ds-footer__links` with `.ds-footer__link`.

## Never

- `box-shadow != none`: nothing raised is flat; a block without its inner glow and drop shadow is not clay.
- `border-width <= 2px`: clay has no outlines; the only lines are 2px rules and the 2px focus or error ring.
- `box-shadow-blur <= 72px`: the deepest shadow, under the dialog, blurs over 72px.
- `text-shadow = none`: the text is flat; only shapes are modelled.
- `font-size <= 54px`: the figure on the dial is the largest text.
- `font-size >= 13px`: meta text, badges and the labels of the rail are the smallest text.
- `font-weight >= 400`: nothing is light.
- `font-families <= 3`: one display face, one sans and one monospace.
- `line-height <= 1.55`: body text is set at 1.5.
- `uppercase-text <= 5%`: nothing is set in capitals.
- `block-gap <= 44px`: cards sit 22px apart; 44px above the footer is the largest gap.
- `content-width <= 1200px`: the working column is capped.

## Extending

Derive a new component from the nearest one in the specimen. Decide first whether it is raised or dented: a container starts from `.ds-panel` (fill, radius and `--shadow-panel` together), something pressed or filled in starts from `.ds-form__input` (`--fill-input`, `--color-input-text`, `--shadow-control-pressed`), something the user presses starts from `.ds-button--secondary` or `.ds-key`. A new card on the board takes a `.ds-bento__*` span so that its row stays full. For a coloured card take a `--color-fill-*` background, `--color-accent-text` for everything on it, and add the coloured drop shadow as `.ds-grid__cell` does. A new device drawing is built like `.ds-device`: clay parts in fills and in the secondary clay, pressed parts in `--fill-input`, all with the shadow tokens. Never draw a shadow in a raw colour and never flatten a shape. Use tokens only; for a size with no token use `calc()` over `--space-*` or `--size-*`. Keep each text token on the fill it belongs to.
