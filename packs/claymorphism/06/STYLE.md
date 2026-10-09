# Claymorphism recipe cards

## Summary

Claymorphism is the soft 3D look of 2021 to 2025: thick, very round shapes that seem modelled from coloured clay, with a light inner glow on the side of the light, a darker inner shadow opposite and a large soft shadow under them. It suited recipe, fitness and habit pages of those years, where a deck of chunky cards and a big round timer made a daily task look like a game on the kitchen table. This pack is a small recipe site in that look: a header of separate clay keys, a hand of four tilted recipe cards with clay plates, a pressed tray of figures, a ring timer and a clay hill as footer, set in a soft serif over a humanist sans.

## Layout

- Content column: `--size-page` (1240px), centred, with `--space-5` (24px) side padding. The footer runs across the whole window; the layout is fluid below that width.
- Navigation: `.ds-nav` has no bar. On the bare page stand the brand, a pressed search well `--size-search` (300px) wide, then, pushed to the right, every section as a clay key of its own, and last the cook's avatar as the toggle of a `.ds-menu`. The key of the current view is pressed in.
- Home page: `.ds-page__sections` stacks sections `--space-7` (52px) apart: the hero (badge, title, lead and actions, centred) with the `.ds-grid` deck of four tilted cards directly under it, the `.ds-stat` tray, a `.ds-page__split` (the list of recipes beside the timer panel) and the closing `.ds-page__band`.
- Inner pages have no hero. `--space-6` (34px) under the navigation come the `.ds-breadcrumb` and the `.ds-page-header` (title at `--text-h1`, one lead line, actions at the right) on the bare page. A recipe page is a `.ds-page__stack`: a `.ds-page__tools` row with the `.ds-tabs` and a `.ds-switch`, then `.ds-page__split`, a reading column (the carousel of steps, the notes panel) with a `--size-aside` (400px) column at the right (ingredients, the table, the cook). The form page uses the same split: the form in one `.ds-panel`, notices and the dialog in the aside. The box page uses `.ds-page__split--side`: a `--size-side` (272px) column with the `.ds-sidebar` and the empty state at the left, the list, the pagination and the accordion at the right. Keep the two columns about equally long.
- Views. The specimen is one site with four views, each `<section class="ds-view" id="...">` inside `.ds-page__inner`; navigation and footer stand outside them. `home` is the day's hand of cards, `recipe` one recipe, `box` the saved recipes, `write` the form for a new one. One `:has()` rule per view presses the key of the current view.
- Spacing scale: `--space-1` 4px (between sidebar links), `--space-2` 8px (label to control, inside a card), `--space-3` 12px (between list bars, between navigation keys), `--space-4` 16px (between buttons and deck cards, row padding), `--space-5` 24px (card padding, form gap, between stacked panels), `--space-6` 34px (column gap of a split, under the navigation), `--space-7` 52px (between home sections, footer padding), `--space-8` 80px (above the footer).
- `--size-plate` (148px) is the diameter of a plate, `--size-timer` (216px) of the timer ring.
- Below 1080px the search well and the keys wrap under the brand, the splits stack and the deck is two by two with a slight tilt. Below 560px the deck is one column without tilt, the stats stack inside their tray, the hero title takes `--text-h1`, tabs fill the width and table columns marked `.ds-table__opt` are hidden, so nothing scrolls sideways.
- On a phone (560px and below) `.ds-nav__list` sits on its own line under the brand and the search well (`.ds-nav__search` is a full row from 1080px down), its keys stretch to share the width equally with centred labels and a smaller gap, and the current key stays pressed in. At every width a tooltip in a table opens at the left edge of its trigger, at most 60vw wide.

## Typography and colour roles

- `--font-heading` is a soft serif (Fraunces, then DM Serif Display, Iowan Old Style) for titles, figures and the brand; `--font-body` and `--font-ui` are a humanist sans (Karla, then Mulish, Gill Sans). `--font-mono` is for code only.
- Sizes: `--text-base` 18px, `--text-ui` 16px for controls, navigation, tables and notices, `--text-small` 15px for meta text, labels and badges, `--text-large` 21px for lead text, step text and large buttons, `--text-h3` 22px, `--text-h2` 29px, `--text-h1` 42px, `--text-display` 68px for the hero title only.
- Weights: 400 body, 600 controls, 700 headings and bold, 800 display. Line height 1.55 for body, 1.18 for headings, 1.02 for display.
- The era's conventions for which text token sits on which fill are kept:
  - `--color-text`, `--color-text-muted`, `--color-heading`, `--color-heading-alt` and the link colours stand only on the page, on `--color-surface`, `--color-surface-alt` and `--color-surface-strong`. The brand in the navigation stands on the bare page and takes `--color-heading`.
  - `--color-accent-text` is the ink for the accent and for all four `--color-fill-*` clays (deck cards, the minutes ball, step numbers, avatars, the current sidebar link). Meta text on a fill is the same ink at opacity 0.8.
  - `--color-bar-text` on `--fill-bar` (the navigation keys), `--color-bar-alt-text` on `--fill-bar-alt` (the footer, table head), `--color-inverse-text` on `--fill-inverse` (the closing band, tooltips, code blocks), `--color-button-text` on `--fill-button` (primary buttons, the timer's main key), `--color-button-secondary-text` on `--fill-button-secondary` (secondary buttons, round keys, pagination keys), `--color-input-text` on `--fill-input` (controls, the search well, the stat tray, the tab track, the empty state), `--color-notice-text` on `--color-notice`, `--color-danger` on `--color-danger-surface`.
  - Status colours never carry text: a status badge is the surface tinted with the status colour, a dot of the full colour and `--color-text`.
- The clay comes from the surface part: every shadow token holds drop shadows in `--color-shadow` (a tint of the palette) and two inset shadows mixed from white and black, so the same tokens shape a white card, a coloured card and the dark band. `--shadow-control-pressed` has only insets, reversed: it is a dent.
- Coloured cards add one more drop shadow in their own colour: `color-mix()` of their fill token and `transparent`, written in the component because the vocabulary has no token per fill.
- One-off values with no token: circles use `border-radius: 50%`, the bits of food their own percentages; a list bar, a notice and an accordion bar take the smaller of `--radius-panel` and `--radius-control`; placeholders are `--color-input-text` at opacity 0.55. The tilt of a deck card is between 2 and 5 degrees.

## Components

- `.ds-page`: on `<body>`; paints `--fill-page` and sets the base type. `.ds-page__inner` is the width container, `.ds-page__sections` the stack of home sections, `.ds-page__stack` a vertical stack 24px apart, `.ds-page__split` a reading column with an aside (`--side` for a side column at the left), `.ds-page__tools` a row with two ends, `.ds-page__title` a section heading with a `.ds-page__more` link at its right, `.ds-page__band` the closing inverted block with `.ds-page__band-title` and `.ds-page__band-text`. `.ds-view` is one screen of the specimen; an application with one document per page does not need it.
- `.ds-brand`: the site's mark and name in the header: `.ds-brand__mark` is a clay tile `--size-brand` (50px) dressed like the primary button, holding the mark as inline SVG; `.ds-brand__name` is the name in the heading font. The name `chronoskin` and its mark are placeholders: replace them with the installing project's own and keep the tile.
- `.ds-nav`: the header row. `.ds-nav__search` wraps an icon and a `.ds-form__input`, `.ds-nav__list` holds the keys, `.ds-nav__link` is one key (pressed in when `is-current`, lifted on `is-hover`), `.ds-nav__actions` holds the account menu.
- `.ds-menu`: a `<details>` drop-down: `.ds-menu__toggle` on the summary (here the avatar and a chevron), `.ds-menu__list` the clay card under it, `.ds-menu__link` with an optional `.ds-menu__note`. `--end` aligns it right, `--up` opens it upwards.
- `.ds-key`: a round clay key with one icon, under the timer; `.ds-key--primary` is larger and in the button colour. `is-hover` lifts it, `is-active` presses it.
- `.ds-hero`: the page introduction, centred on the bare page: a badge, `.ds-hero__title` at the display size, `.ds-hero__lead`, `.ds-hero__actions`. The deck follows it without a gap.
- `.ds-plate`: a clay dish with a pressed well and up to three `.ds-plate__bit` of food (`--1` to `--4` pick the fill); `.ds-plate--small` is the size for a list row. `.ds-clay` (ball, `--cube`, `--pill`, `--ring`, sizes and fills) gives further objects. Decorative: mark them `aria-hidden`.
- `.ds-page-header`: head of an inner page: `.ds-page-header__text` with `.ds-page-header__title` and `.ds-page-header__lead`, `.ds-page-header__actions` at the right with at most one primary button.
- `.ds-prose`: running text inside a panel: h1 to h3, paragraphs, lists, inline `code`; `.ds-prose__lead` for a lead paragraph, `.ds-prose__code` for an inverted code block.
- `.ds-link`: link in running text, underlined; states `is-visited`, `is-hover`, `is-active`, `is-focus`; `.ds-link--quiet` for titles and table links; `.ds-link__sample` lays out a row of links.
- `.ds-button`: a clay key 50px high. `is-hover` lifts it, `is-active` squashes it. `.ds-button--secondary` is the same key in the surface colour, `.ds-button--danger` the destructive action, `.ds-button--large` (62px) for the hero and the band, `.ds-button--wide` fills its card; `disabled` or `is-disabled` is flat. `.ds-button__row` spaces a group. One primary button per group.
- `.ds-form`: vertical form in a panel. `.ds-form__field` wraps a `.ds-form__label` and a control: `.ds-form__input`, `.ds-form__textarea`, `.ds-form__select` (in `.ds-form__select-wrap` with a `.ds-form__chevron`). Controls are dents. `.ds-form__check` is a label row with a `.ds-form__checkbox`; it also serves as a line of an ingredient list. `is-error` rings the control in `--color-danger` and is followed by `.ds-form__error`; `.ds-form__hint` is help text, `.ds-form__actions` holds the buttons, `.ds-form__row` puts two fields side by side.
- `.ds-switch`: a fat clay toggle over a checkbox: `.ds-switch__input`, `.ds-switch__track` (a dent that fills with the accent when on) with a clay ball as thumb, `.ds-switch__label` and `.ds-switch__hint`; `--end` puts the track at the right.
- `.ds-table`: one clay sheet with a head in `--color-bar-alt` and rows divided by a soft rule. `.ds-table__num` right-aligns numbers, `.ds-table__meta` is muted, `.ds-table__strong` a row title, `.ds-table__opt` a column dropped on a phone, `is-hover` tints a row.
- `.ds-list`: recipes as separate clay bars: `.ds-list__item` holds a `.ds-list__time` ball (minutes over a `.ds-list__unit`) or a small plate, a `.ds-list__body` (`.ds-list__title`, `.ds-list__meta`) and a badge; `.ds-list__aside` is a figure at the end.
- `.ds-panel`: the basic clay card: `.ds-panel__title`, `.ds-panel__body` (a column with 16px gaps), `.ds-panel__note` for small muted text, `.ds-panel__line` for a row with two ends. `.ds-panel--accent` is the card in the accent clay.
- `.ds-stat`: the row; one long pressed tray holding three `.ds-stat__item` figures divided by rules, each a `.ds-stat__value` (with an optional `.ds-stat__meta`) over a `.ds-stat__label`, all in `--color-input-text`.
- `.ds-grid`: the deck, four recipe cards fanned like a hand. `.ds-grid__cell` with `--1` to `--4` takes a `--color-fill-*` clay, its own coloured drop shadow and its tilt; inside, a `.ds-plate`, `.ds-grid__title`, `.ds-grid__text` and an underlined `.ds-grid__more`. `is-hover` straightens and lifts a card.
- `.ds-timer`: a large `.ds-progress--ring` whose face shows the time over a `.ds-timer__label`, and `.ds-timer__keys`, a row of `.ds-key`.
- `.ds-step`: one step of a recipe, used as a carousel slide: a `.ds-step__num` ball in the accent and a `.ds-step__body` with `.ds-step__title`, `.ds-step__text` and one secondary button.
- `.ds-tabs`: a dented track of `.ds-tabs__tab` pills; `is-current` is a raised pill in the accent. For switching what a page shows, not for site navigation.
- `.ds-badge`: a pill in the accent. `.ds-badge--quiet` is neutral; `.ds-badge--success`, `.ds-badge--warning` and `.ds-badge--danger` are tinted pills with a status dot.
- `.ds-sidebar`: a clay card of shelves: `.ds-sidebar__heading`, `.ds-sidebar__list`, `.ds-sidebar__link` with a quiet badge as count; `is-current` is a raised pill in the accent.
- `.ds-notice`: a raised bar in `--color-notice` with `.ds-notice__icon` and `.ds-notice__title`; `.ds-notice--error` is the error. `.ds-notice__stack` stacks several.
- `.ds-pagination`: clay keys, `.ds-pagination__link`; the current page is `is-current`, in the accent.
- `.ds-breadcrumb`: small bold trail above the page header: `.ds-breadcrumb__item`, `.ds-breadcrumb__link`.
- `.ds-dialog`: a tinted sheet (`--color-overlay`) with a centred `.ds-dialog__box` under the deepest shadow: `.ds-dialog__title`, `.ds-dialog__body`, `.ds-dialog__actions`.
- `.ds-empty`: a dent holding a clay ball `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one primary button.
- `.ds-tooltip`: wraps a trigger; `.ds-tooltip__tip` is an inverted pill shown on hover or focus (`is-open` shows it statically).
- `.ds-progress`: a thick dented track with a puffy `.ds-progress__bar`; width by `--10`, `--25`, `--40`, `--60`, `--75`, `--90`; `.ds-progress__legend` is the line above it. `.ds-progress--ring` is a ring around a clay ball `.ds-progress__face`; the timer is one.
- `.ds-avatar`: initials on a clay ball `.ds-avatar__pic`, alone or with `.ds-avatar__text` (`.ds-avatar__name`, `.ds-avatar__meta`); `--1` to `--4` pick a fill.
- `.ds-accordion`: `<details>` sections as separate clay bars: `.ds-accordion__item`, `.ds-accordion__head` with a round `.ds-accordion__chevron`, `.ds-accordion__body`.
- `.ds-carousel`: `.ds-carousel__track` scrolls and snaps, one `.ds-carousel__slide` showing; `.ds-carousel__foot` holds two round `.ds-carousel__arrow` keys around `.ds-carousel__dots` (`.ds-carousel__dot`, `is-current`).
- `.ds-footer`: a clay hill in `--fill-bar-alt` across the bottom of the window, round at the top only: `.ds-footer__inner` (three columns), `.ds-footer__about` with the brand and `.ds-footer__text`, `.ds-footer__heading`, `.ds-footer__links`, `.ds-footer__link`.

## Never

- `box-shadow != none`: nothing raised is flat; a block without its inner glow and drop shadow is not clay.
- `border-width <= 3px`: clay has no outlines; the only lines are 2px rules and the 3px focus or error ring.
- `box-shadow-blur <= 80px`: the deepest shadow, under the dialog and the band, blurs over 80px.
- `text-shadow = none`: the text is flat; only shapes are modelled.
- `font-size <= 68px`: the hero title is the largest text.
- `font-size >= 15px`: meta text and badges are the smallest text.
- `font-weight >= 400`: nothing is light.
- `font-families <= 3`: one serif, one sans and one monospace.
- `line-height <= 1.6`: body text is set at 1.55.
- `uppercase-text <= 5%`: nothing is set in capitals.
- `block-gap <= 80px`: 52px between home sections and 80px above the footer are the largest gaps.
- `content-width <= 1240px`: the column is capped.

## Extending

Derive a new component from the nearest one in the specimen. Decide first whether it is raised or dented: a container starts from `.ds-panel` (fill, radius and `--shadow-panel` together), something pressed or filled in starts from `.ds-form__input` (`--fill-input`, `--color-input-text`, `--shadow-control-pressed`), something the user presses starts from `.ds-button--secondary` or `.ds-key`. For a coloured card take a `--color-fill-*` background, `--color-accent-text` for everything on it, and add the coloured drop shadow as `.ds-grid__cell` does. A new picture of food or a tool is built like `.ds-plate`: a dish in the secondary clay with a pressed well, and round bits in the fills raised with `--shadow-control`. Never draw a shadow in a raw colour and never flatten a shape. Use tokens only; for a size with no token use `calc()` over `--space-*` or `--size-*`. Keep each text token on the fill it belongs to.
