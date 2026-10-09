# Claymorphism collectibles landing

## Summary

Claymorphism is the soft 3D look of 2021 to 2025: thick, very round shapes that seem modelled from coloured clay, with a light inner glow above, a darker lip below and a soft coloured shadow under them. In 2021 and 2022 it dressed the landing pages of collectible and token projects, where floating coins, blobs and little characters rendered in 3D promised something friendly. This pack is such a landing page for a set of numbered clay figures: a centred hero between floating coins and critters, three big coins with a figure each, steps as soft bricks, a carousel of figures and a dark closing block, set in a wide heavy display face over a geometric sans.

## Layout

- Content column: `--size-page` (1080px), centred, with `--space-5` (26px) side padding. The page fill covers the rest of the window; the layout is fluid below that width.
- Navigation: `.ds-nav` has no bar. The brand stands on the bare page at the left, the links sit together in one small clay capsule in the middle (the last of them a `.ds-menu`), the primary button stands at the right.
- Home page: `.ds-page__sections` stacks centred sections `--space-8` (96px) apart: the hero (badge, title, lead, actions and the meter bar, with clay things floating at both sides), the `.ds-stat` coins, a `.ds-feature` row (a pressed tray of clay things beside text), the `.ds-grid` of four bricks under a centred `.ds-page__title`, a `.ds-page__split` of two even columns (the carousel of figures beside the accordion of questions) and the closing `.ds-page__band`.
- Inner pages have no hero. `--space-5` under the navigation come the `.ds-breadcrumb` and the `.ds-page-header` (title at `--text-h1`, one lead line, the action at the right) on the bare page. The drops page is one `.ds-page__stack`, 36px apart: a `.ds-page__tools` row with the `.ds-tabs` and a `.ds-switch`, the table across the column, the pagination, then a `.ds-page__split` (the list of recent draws beside a panel). The join page is a `.ds-page__split`: the form in one `.ds-panel`, notices and the dialog beside it. The guide uses `.ds-page__split--side`: the text panel at the left, a `--size-side` (300px) column with the `.ds-sidebar` and the empty state at the right. Keep the two columns about equally long.
- Views. The specimen is one site with four views, each `<section class="ds-view" id="...">` inside `.ds-page__inner`; navigation and footer stand outside them. `home` is the landing page, `drops` the table of drops, `join` the form, `guide` the help text. One `:has()` rule per view puts the link of the current view in the accent clay.
- Spacing scale: `--space-1` 4px (between capsule links), `--space-2` 8px (label to control, the thickness of a coin's edge), `--space-3` 12px (between list bars, the edge of a big coin), `--space-4` 18px (between buttons, row padding), `--space-5` 26px (card padding, grid gap, form gap), `--space-6` 36px (between panels, column gap of a split), `--space-7` 56px (band padding, gap in a feature row), `--space-8` 96px (between home sections, above the footer).
- `--size-coin` (232px) is the diameter of a stat coin, `--size-tray` (300px) the least height of a feature tray, `--size-clay` (112px) the base size of coins and critters.
- Below 960px the splits and the feature row stack, the capsule moves under the brand and the hero's clay things line up under its text instead of floating, and table cells close up to `--space-3` at their sides. Below 720px the bricks are one column and the stat coins lie straight. Below 560px the hero title takes `--text-h1`, the meter stacks, tabs fill the width and table columns marked `.ds-table__opt` are hidden, so nothing scrolls sideways.
- On a phone (560px and below) the capsule `.ds-nav__list` stays under the brand and does not wrap: its links become equal cells with centred text inside a pill-shaped capsule, and the current link stays a raised accent pill. Between 560px and 960px the capsule wraps and is centred. Below 960px the bar has two lines, brand and button above, the capsule below.

## Typography and colour roles

- `--font-heading` is a wide, heavy display face (Unbounded, then Syne, Verdana) for titles, figures and the brand; `--font-body` and `--font-ui` are a geometric sans (Plus Jakarta Sans, then Avenir). `--font-mono` is for code only.
- Sizes: `--text-base` 17px, `--text-ui` 15px for controls, navigation, tables and notices, `--text-small` 14px for meta text, labels and badges, `--text-large` 21px for lead text and large buttons, `--text-h3` 17px, `--text-h2` 25px, `--text-h1` 38px for section titles and the figures on the coins, `--text-display` 60px for the hero title only.
- Weights: 500 body, 700 headings, 800 controls, bold and display. Line height 1.6 for body, 1.2 for headings, 1.08 for display. Headings are tracked a little tight.
- The era's conventions for which text token sits on which fill are kept:
  - `--color-text`, `--color-text-muted`, `--color-heading`, `--color-heading-alt` and the link colours stand only on the page, on `--color-surface`, `--color-surface-alt` and `--color-surface-strong`. The brand in the navigation stands on the bare page and takes `--color-heading`.
  - `--color-accent-text` is the ink for the accent and for all four `--color-fill-*` clays (stat coins, bricks, avatars, the current link). Meta text on a fill is the same ink at opacity 0.8.
  - `--color-bar-text` on `--fill-bar` (the capsule), `--color-bar-alt-text` on `--fill-bar-alt` (the footer block, table head), `--color-inverse-text` on `--fill-inverse` (the closing band, tooltips, code blocks), `--color-button-text` on `--fill-button`, `--color-button-secondary-text` on `--fill-button-secondary` (secondary buttons, numbered balls, pagination keys, the current sidebar link), `--color-input-text` on `--fill-input` (controls, the tab track, the sidebar well, the empty state), `--color-notice-text` on `--color-notice`, `--color-danger` on `--color-danger-surface`.
  - Status colours never carry text: a status badge is the surface tinted with the status colour, a dot of the full colour and `--color-text`.
- The clay comes from the surface part: every shadow token holds a drop shadow in `--color-shadow` and two inset shadows mixed from white and black (here a soft glow above and a hard lip below), so the same tokens shape a white card, a coloured brick and the dark band. `--shadow-control-pressed` has only insets, reversed: it is a dent.
- Coins and stat coins get their thickness from one more hard shadow under them: `color-mix()` of their fill token and `--color-accent-text`, which is that fill's own dark. Bricks add a soft drop shadow in their own colour, mixed with `transparent`. Both are written in the component because the vocabulary has no token per fill.
- One-off values with no token: circles use `border-radius: 50%`, critters and their horns their own percentages; a list bar, a notice, an accordion bar and a figure tray take the smaller of `--radius-panel` and `--radius-control`; placeholders are `--color-input-text` at opacity 0.55.

## Components

- `.ds-page`: on `<body>`; paints `--fill-page` and sets the base type. `.ds-page__inner` is the width container, `.ds-page__sections` the stack of home sections, `.ds-page__stack` a vertical stack 36px apart, `.ds-page__split` two even columns (`--side` for a side column at the right), `.ds-page__narrow` a narrower centred column, `.ds-page__tools` a row with two ends, `.ds-page__title` a centred section heading with `.ds-page__title-note` (`--left` aligns it left), `.ds-page__band` the closing inverted block with `.ds-page__band-title`, `.ds-page__band-text` and `.ds-page__band-art`. `.ds-view` is one screen of the specimen; an application with one document per page does not need it.
- `.ds-brand`: the site's mark and name in the header: `.ds-brand__mark` is a clay tile `--size-brand` (46px) dressed like the primary button, holding the mark as inline SVG; `.ds-brand__name` is the name in the heading font. The name `chronoskin` and its mark are placeholders: replace them with the installing project's own and keep the tile.
- `.ds-nav`: the header row. `.ds-nav__list` is the capsule, `.ds-nav__link` one link (a raised accent pill when `is-current`, pressed in on `is-hover`), `.ds-nav__actions` the button at the right.
- `.ds-menu`: a `<details>` drop-down under a navigation item: `.ds-menu__toggle` on the summary, `.ds-menu__list` the clay card under it, `.ds-menu__link` with an optional `.ds-menu__note`. `--end` aligns it right, `--up` opens it upwards.
- `.ds-hero`: the page introduction, centred on the bare page: `.ds-hero__body` holds a badge, `.ds-hero__title` at the display size, `.ds-hero__lead`, `.ds-hero__actions` and `.ds-hero__meter` (a clay bar with overlapping avatars in `.ds-hero__faces` and a progress bar in `.ds-hero__meter-bar`). `.ds-hero__art` lies behind it and holds clay things, each placed with `.ds-hero__obj--a` to `--f`.
- `.ds-coin`: a clay coin with a pressed ring and a raised middle, tilted; `--2` to `--4` pick the fill, `--small` and `--large` the size. `.ds-critter`: a blob of clay on two round feet, without a face; it may hold two `.ds-critter__horn` and a pressed `.ds-critter__belly`; `--1`, `--3`, `--4` pick the fill, `--small` and `--large` the size. `.ds-clay` (ball, `--cube`, `--pill`, `--ring`, `--tiny` and other sizes) gives further objects. All are decorative: mark them `aria-hidden`.
- `.ds-feature`: a row of two halves: `.ds-feature__art`, a pressed tray of clay things, and a text half with a `.ds-page__title--left`, `.ds-feature__text` and one button.
- `.ds-page-header`: head of an inner page: `.ds-page-header__text` with `.ds-page-header__title` and `.ds-page-header__lead`, `.ds-page-header__actions` at the right with at most one primary button.
- `.ds-prose`: running text inside a panel: h1 to h3, paragraphs, lists, inline `code`; `.ds-prose__lead` for a lead paragraph, `.ds-prose__code` for an inverted code block.
- `.ds-link`: link in running text; states `is-visited`, `is-hover`, `is-active`, `is-focus`; `.ds-link--quiet` for titles and table links; `.ds-link__sample` lays out a row of links.
- `.ds-button`: a fat clay key 52px high. `is-hover` lifts it, `is-active` squashes it. `.ds-button--secondary` is the same key in the surface colour, `.ds-button--danger` the destructive action, `.ds-button--large` (64px) for the hero and the band, `.ds-button--wide` fills its card; `disabled` or `is-disabled` is flat. `.ds-button__row` spaces a group. One primary button per group.
- `.ds-form`: vertical form in a panel. `.ds-form__field` wraps a `.ds-form__label` and a control: `.ds-form__input`, `.ds-form__textarea`, `.ds-form__select` (in `.ds-form__select-wrap` with a `.ds-form__chevron`). Controls are dents. `.ds-form__check` is a label row with a `.ds-form__checkbox`. `is-error` rings the control in `--color-danger` and is followed by `.ds-form__error`; `.ds-form__hint` is help text, `.ds-form__actions` holds the buttons, `.ds-form__row` puts two fields side by side.
- `.ds-switch`: a fat clay toggle over a checkbox: `.ds-switch__input`, `.ds-switch__track` (a dent that fills with the accent when on) with a clay ball as thumb, `.ds-switch__label` and `.ds-switch__hint`; `--end` puts the track at the right.
- `.ds-table`: one clay sheet with a head in `--color-bar-alt` and rows divided by a soft rule. `.ds-table__num` right-aligns numbers, `.ds-table__meta` is muted, `.ds-table__strong` a row title, `.ds-table__opt` a column dropped on a phone, `is-hover` tints a row.
- `.ds-list`: items as separate clay bars: `.ds-list__item` holds an avatar, a `.ds-list__body` (`.ds-list__title`, `.ds-list__meta`) and a badge or a `.ds-list__aside` figure.
- `.ds-panel`: the basic clay card: `.ds-panel__title`, `.ds-panel__body` (a column with 16px gaps), `.ds-panel__note` for small muted text, `.ds-panel__line` for a row with two ends. `.ds-panel--accent` is the card in the accent clay.
- `.ds-stat`: the row; each `.ds-stat__item` is one big clay coin lying a little askew, in the first three fills, with a `.ds-stat__value` over a `.ds-stat__label`. Keep the figure to six characters.
- `.ds-grid`: four soft bricks, two by two. `.ds-grid__cell` with `--1` to `--4` takes a `--color-fill-*` clay and its own coloured drop shadow; inside, a numbered ball `.ds-grid__num` beside a `.ds-grid__title` and `.ds-grid__text`.
- `.ds-card`: a figure card, used as a carousel slide: `.ds-card__tray` is a pressed square holding a critter, `.ds-card__body` holds a badge, `.ds-card__title`, `.ds-card__text` and one button.
- `.ds-tabs`: a dented track of `.ds-tabs__tab` pills; `is-current` is a raised pill in the accent. For switching what a page shows, not for site navigation.
- `.ds-badge`: a pill in the accent. `.ds-badge--quiet` is neutral; `.ds-badge--success`, `.ds-badge--warning` and `.ds-badge--danger` are tinted pills with a status dot.
- `.ds-sidebar`: a pressed well of topics: `.ds-sidebar__heading`, `.ds-sidebar__list`, `.ds-sidebar__link`, a raised key when `is-current` or hovered.
- `.ds-notice`: a raised bar in `--color-notice` with `.ds-notice__icon` and `.ds-notice__title`; `.ds-notice--error` is the error. `.ds-notice__stack` stacks several.
- `.ds-pagination`: clay keys, `.ds-pagination__link`; the current page is `is-current`, in the accent.
- `.ds-breadcrumb`: small bold trail above the page header: `.ds-breadcrumb__item`, `.ds-breadcrumb__link`.
- `.ds-dialog`: a tinted sheet (`--color-overlay`) with a centred `.ds-dialog__box` under the deepest shadow: `.ds-dialog__title`, `.ds-dialog__body`, `.ds-dialog__actions`.
- `.ds-empty`: a dent holding a clay ball `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one primary button.
- `.ds-tooltip`: wraps a trigger; `.ds-tooltip__tip` is an inverted pill shown on hover or focus (`is-open` shows it statically).
- `.ds-progress`: a thick dented track with a puffy `.ds-progress__bar`; width by `--10`, `--25`, `--40`, `--60`, `--75`, `--90`; `.ds-progress__legend` is the line above it. `.ds-progress--ring` is a ring around a clay ball `.ds-progress__face`.
- `.ds-avatar`: initials on a clay ball `.ds-avatar__pic`, alone or with `.ds-avatar__text` (`.ds-avatar__name`, `.ds-avatar__meta`); `--1` to `--4` pick a fill.
- `.ds-accordion`: `<details>` sections as separate clay bars: `.ds-accordion__item`, `.ds-accordion__head` with a round `.ds-accordion__chevron`, `.ds-accordion__body`.
- `.ds-carousel`: `.ds-carousel__track` scrolls and snaps, one `.ds-carousel__slide` showing; `.ds-carousel__foot` holds two round `.ds-carousel__arrow` keys around `.ds-carousel__dots` (`.ds-carousel__dot`, `is-current`).
- `.ds-footer`: one clay block in `--fill-bar-alt` with everything centred: `.ds-footer__inner` holds the brand, `.ds-footer__links` with `.ds-footer__link` and `.ds-footer__text`.

## Never

- `box-shadow != none`: nothing raised is flat; a block without its inner glow, its lip and its drop shadow is not clay.
- `border-width <= 3px`: clay has no outlines; the only lines are 2px rules and the 3px focus or error ring.
- `box-shadow-blur <= 60px`: the deepest shadow, under the dialog and the band, blurs over 60px.
- `text-shadow = none`: the text is flat; only shapes are modelled.
- `font-size <= 60px`: the hero title is the largest text.
- `font-size >= 14px`: meta text and badges are the smallest text.
- `font-weight >= 500`: the type is sturdy; nothing is light.
- `font-families <= 3`: one display face, one sans and one monospace.
- `line-height <= 1.65`: body text is set at 1.6.
- `uppercase-text <= 5%`: nothing is set in capitals.
- `block-gap <= 96px`: 96px between home sections is the largest gap.
- `content-width <= 1080px`: the column is capped.

## Extending

Derive a new component from the nearest one in the specimen. Decide first whether it is raised or dented: a container starts from `.ds-panel` (fill, radius and `--shadow-panel` together), something pressed or filled in starts from `.ds-form__input` (`--fill-input`, `--color-input-text`, `--shadow-control-pressed`), something the user presses starts from `.ds-button--secondary`. For a coloured card take a `--color-fill-*` background, `--color-accent-text` for everything on it, and add the coloured drop shadow as `.ds-grid__cell` does. A new illustration is built like `.ds-coin` and `.ds-critter`: round shapes in a fill, parts that inherit that fill, raised with `--shadow-control` or pressed with `--shadow-control-pressed`, and no faces. Never draw a shadow in a raw colour and never flatten a shape. Use tokens only; for a size with no token use `calc()` over `--space-*` or `--size-*`. Keep each text token on the fill it belongs to.
