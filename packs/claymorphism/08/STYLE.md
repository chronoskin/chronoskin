# Claymorphism clubhouse forum

## Summary

Claymorphism is the soft 3D look of 2021 to 2025: thick, very round shapes that seem modelled from coloured clay, each with a light inner glow above, a darker inner shade below and a soft drop shadow under it. This pack is a small community forum in that look, in bubblegum colours: pink and sky blue clay with lilac and aqua beside them, a cherry red button, and every raised shape standing on a firm ledge of its own shadow like a sweet on a tray. It was common for the community, club and fan pages of playful consumer products, where posts are speech bubbles and members are clay balls.

## Layout

- Content column: `--size-page` (1180px), centred, with `--space-5` (24px) side padding; the layout is fluid below that width.
- Navigation: `.ds-nav` has no bar. On the bare page stand the brand, then a row of separate clay keys `--size-nav` (56px) high in `--fill-bar`, each an icon and a word, and at the right end the primary button and the account (an avatar that opens a `.ds-menu`).
- Every page is the same two columns from the top, `.ds-page__split`: a working column and an aside of `--size-aside` (340px) at its right. Blocks in a column are `.ds-page__stack`ed `--space-6` (34px) apart.
- Home page: in the working column the hero (one wide clay slab: the title runs across it, lead and buttons below at the left, three clay plants in the lower right corner), the `.ds-grid` of four boards two to a row under a `.ds-page__head`, and the list of current topics. In the aside three panels: the club's figures with a progress ring, the carousel, the podium.
- Inner pages have no hero. Under the navigation come the `.ds-breadcrumb` and, on a board or a form, the `.ds-page-header` on the bare page. A board has the tabs, the table of topics and the pagination in the working column, the `.ds-sidebar` of boards and the accordion of rules in the aside. A topic is a run of `.ds-post` (the member in a `--size-author` 148px column, the words in a speech bubble) and the reply panel, with the topic's own panels in the aside; the first post's h1 stands in for the page header. A form page has the form panel in the working column and notices, empty state and dialog in the aside.
- Views. The specimen is one site with four views, each `<section class="ds-view" id="...">` inside `.ds-page__inner`: `home` (the boards), `board` (one board's topics), `topic` (one topic) and `new` (the form for a new topic). The key of the current view is pressed in by one `:has()` rule per view.
- Spacing scale: `--space-1` 4px (dot to text), `--space-2` 8px (label to control, icon to word), `--space-3` 12px (between list bars, keys and figures), `--space-4` 18px (between buttons, row padding, member to bubble), `--space-5` 24px (card padding, grid gap, form gap), `--space-6` 34px (between panels and columns, hero padding), `--space-7` 52px (dialog sheet padding), `--space-8` 80px (above the footer).
- `--size-rule` (2px) is the weight of the soft rules between table rows.
- Below 960px the aside moves under the working column and the keys of the navigation take a line of their own, equally wide. Below 560px the brand shows only its mark and the account only its avatar, the keys lose their icons, the boards are one column, the hero stacks with its plants centred, a post puts its member in a row above the bubble, form rows stack, the tabs fill the width and table columns marked `.ds-table__opt` are hidden, so nothing scrolls sideways.

## Typography and colour roles

- `--font-heading` and `--font-ui` are a fat rounded face (Baloo 2, then Grandstander, Arial Rounded) for titles, figures, keys and names; `--font-body` is a soft geometric sans (Be Vietnam Pro, then Red Hat Text). `--font-mono` is for code only.
- Sizes: `--text-base` 16px, `--text-ui` 16px for controls, navigation, tables and notices, `--text-small` 14px for meta text, labels and badges, `--text-large` 18px for lead text and large buttons, `--text-h3` 19px, `--text-h2` 26px, `--text-h1` 38px, `--text-display` 52px for the hero title only.
- Weights: 500 body, 600 for controls, 700 for links and labels, 800 for headings. Line height 1.5 for body, 1.15 for headings, 1 for display.
- The era's conventions for which text token sits on which fill are kept:
  - `--color-text`, `--color-text-muted`, `--color-heading`, `--color-heading-alt` and the link colours stand only on the page, on `--color-surface`, `--color-surface-alt` and `--color-surface-strong`.
  - `--color-accent-text` is the ink for the accent and for all four `--color-fill-*` clays (boards, figures, plinths, avatars, the picked answer, the accent panel). Meta text on a fill is the same ink at opacity 0.8.
  - `--color-bar-text` on `--fill-bar` (the navigation keys), `--color-bar-alt-text` on `--fill-bar-alt` (the footer pill, the table head), `--color-inverse-text` on `--fill-inverse` (code blocks, tooltips), `--color-button-text` on `--fill-button`, `--color-button-secondary-text` on `--fill-button-secondary` (secondary buttons, the account, reaction keys, icon balls, pagination keys, switch thumbs), `--color-input-text` on `--fill-input` (controls, the tab track, the reply count of a topic, a used reaction, the shelf, the empty state), `--color-notice-text` on `--color-notice`, `--color-danger` on `--color-danger-surface`.
  - Status colours never carry text: a status badge is the surface tinted with the status colour, a dot of the full colour and `--color-text`. The leaves of the clay plants are `--color-success` mixed into `--color-surface`; they carry no text either.
- The clay comes from the surface part. Each raised shadow token holds a firm ledge (a drop shadow without blur, in half of `--color-shadow`), a soft drop shadow in `--color-shadow`, an inner shade of black along the lower edge and an inner glow of white along the upper edge. `--shadow-control-pressed` has only the insets, reversed: it is a dent.
- One-off values with no token: circles use `border-radius: 50%`; a speech bubble and a plinth take `--space-2` for their one or two tight corners; a list bar, a notice, an accordion bar and a figure take the smaller of `--radius-panel` and `--radius-control`; placeholders are `--color-input-text` at opacity 0.55; the parts of a plant are placed, sized and rounded in percent of its box.

## Components

- `.ds-page`: on `<body>`; paints `--fill-page` and sets the base type. `.ds-page__inner` is the width container, `.ds-page__split` the two columns, `.ds-page__stack` a vertical stack 34px apart, `.ds-page__head` a section heading row with `.ds-page__title` and a link or `.ds-page__title-note` at its right end. `.ds-page__icon` sizes an icon in a button. `.ds-view` is one screen of the specimen; an application with one document per page does not need it.
- `.ds-brand`: the site's mark and name in the header: `.ds-brand__mark` is a clay ball `--size-brand` (48px) dressed like the primary button, holding the mark as inline SVG; `.ds-brand__name` is the name in the heading font. The name `chronoskin` and its mark are placeholders: replace them with the installing project's own and keep the ball.
- `.ds-nav`: the header row on the bare page. `.ds-nav__inner` lays it out, `.ds-nav__list` holds the keys, `.ds-nav__link` is one key (pressed in when `is-current`, lifted on `is-hover`), `.ds-nav__actions` the right end, `.ds-nav__account` the account pill with an avatar and `.ds-nav__account-name`.
- `.ds-menu`: a `<details>` drop-down, here under the account: `.ds-menu__toggle` on the summary, `.ds-menu__list` the clay card that opens under it, `.ds-menu__link` with an optional `.ds-menu__note`. `--end` aligns it right, `--up` opens it upwards.
- `.ds-hero`: the welcome slab of the home page: `.ds-hero__head` (a badge and `.ds-hero__title` at the display size) across the top, then `.ds-hero__lead` with `.ds-hero__actions`, and `.ds-hero__art` holding the plants.
- `.ds-pot`: a plant drawn in CSS from the same clay as the cards: `.ds-pot__body` and `.ds-pot__rim` in a block colour (`--1` to `--4`), and either three `.ds-pot__leaf` (`--left`, `--right`, `--tall`) or a `.ds-pot__ball` for a cactus; `--small` and `--large` size it. Decorative only: mark it `aria-hidden`.
- `.ds-page-header`: head of an inner page: `.ds-page-header__text` with `.ds-page-header__title` and `.ds-page-header__lead`, `.ds-page-header__actions` at the right with at most one primary button.
- `.ds-prose`: running text inside a post or a panel: h1 to h3, paragraphs, lists, inline `code`; `.ds-prose__lead` for a lead paragraph, `.ds-prose__code` for an inverted code block.
- `.ds-link`: link in running text; states `is-visited`, `is-hover`, `is-active`, `is-focus`; `.ds-link--quiet` for titles and table links; `.ds-link__sample` lays out a row of links.
- `.ds-button`: a fat clay key 50px high on its ledge. `is-hover` lifts it, `is-active` squashes it and takes `--shadow-control-pressed`. `.ds-button--secondary` is the same key in the surface colour, `.ds-button--danger` the destructive action on `--color-danger-surface`, `.ds-button--large` (62px) for the hero, `.ds-button--wide` fills its card; `disabled` or `is-disabled` is flat. `.ds-button__row` spaces a group. One primary button per group.
- `.ds-form`: vertical form in a panel. `.ds-form__field` wraps a `.ds-form__label` and a control: `.ds-form__input`, `.ds-form__textarea`, `.ds-form__select` (in `.ds-form__select-wrap` with a `.ds-form__chevron`). Controls are dents: `--fill-input` with `--shadow-control-pressed`. `.ds-form__check` is a label row with a `.ds-form__checkbox` (a raised clay tile, accent with a tick when checked). `is-error` rings the control in `--color-danger` and is followed by `.ds-form__error`; `.ds-form__hint` is help text, `.ds-form__actions` holds the buttons, `.ds-form__row` puts two fields side by side.
- `.ds-switch`: a fat clay toggle over a checkbox: `.ds-switch__input`, `.ds-switch__track` (a dent that fills with the accent when on) with a clay ball as thumb, `.ds-switch__label` and `.ds-switch__hint`; `--end` puts the track at the right.
- `.ds-table`: the topics of a board on one clay sheet with a head in `--color-bar-alt` and rows divided by a soft rule. `.ds-table__num` right-aligns numbers, `.ds-table__meta` is muted, `.ds-table__strong` a row title, `.ds-table__opt` a column dropped on a phone, `is-hover` tints a row.
- `.ds-list`: topics as separate clay bars 12px apart: `.ds-list__item` holds an avatar, a `.ds-list__body` (`.ds-list__title`, `.ds-list__meta`), a badge and `.ds-list__aside`, the number of replies in a dent.
- `.ds-panel`: the basic clay card: `.ds-panel__title`, `.ds-panel__body` (a column with 18px gaps), `.ds-panel__note` for small muted text, `.ds-panel__line` for a row with two ends. `.ds-panel--accent` is the card in the accent clay, with all text in `--color-accent-text`.
- `.ds-grid`: the boards, two to a row. `.ds-grid__cell` with `--1` to `--4` takes a `--color-fill-*` clay; inside, a `.ds-grid__icon` ball, the board's name as `.ds-grid__title` (a link), `.ds-grid__text` and a `.ds-grid__foot` with a `.ds-grid__count` and the `.ds-faces` of who posted last. A card lifts and tilts on `is-hover`.
- `.ds-faces`: a few small avatars that overlap.
- `.ds-stat`: the row of three figures; each `.ds-stat__item` is a tile with a `.ds-stat__value` over a `.ds-stat__label`. `--1` to `--3` put a tile in a block colour.
- `.ds-podium`: the three most helpful members: each `.ds-podium__place` (`--1`, `--3`; plain is second) is an avatar, a `.ds-podium__name` and a `.ds-podium__block`, a plinth in a block colour with the place on it.
- `.ds-tabs`: a dented track of `.ds-tabs__tab` pills; `is-current` is a raised pill in the accent. For filtering a board, not for site navigation.
- `.ds-badge`: a pill in the accent, also used for a member's rank. `.ds-badge--quiet` is neutral; `.ds-badge--success`, `.ds-badge--warning` and `.ds-badge--danger` are tinted pills with a status dot.
- `.ds-post`: one message of a topic: `.ds-post__author` (a large avatar, `.ds-post__name`, a rank badge, `.ds-post__meta`) and `.ds-post__bubble`, a clay card with one tight corner towards the member, holding `.ds-post__head`, the text and `.ds-post__foot`. `.ds-post--picked` puts the bubble in the accent clay: the answer the asker chose.
- `.ds-react`: a small clay key with an icon and a count under a post; `is-on` presses it in.
- `.ds-sidebar`: a clay card listing the boards: `.ds-sidebar__heading`, `.ds-sidebar__list`, `.ds-sidebar__link` (dented when `is-current`) with a `.ds-sidebar__dot` in the board's colour (`--2` to `--4`) and a `.ds-sidebar__count`.
- `.ds-notice`: a raised bar in `--color-notice` with `.ds-notice__icon` and `.ds-notice__title`; `.ds-notice--error` is the error. `.ds-notice__stack` stacks several.
- `.ds-pagination`: clay keys, `.ds-pagination__link`; the current page is `is-current`, in the accent.
- `.ds-breadcrumb`: small bold trail at the top of an inner page: `.ds-breadcrumb__item`, `.ds-breadcrumb__link`.
- `.ds-dialog`: a tinted sheet (`--color-overlay`) with a centred `.ds-dialog__box` under the deepest shadow: `.ds-dialog__title`, `.ds-dialog__body`, `.ds-dialog__actions`.
- `.ds-empty`: a dent holding a clay ball `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one button.
- `.ds-tooltip`: wraps a trigger, here a rank badge; `.ds-tooltip__tip` is an inverted pill shown on hover or focus (`is-open` shows it statically).
- `.ds-progress`: a thick dented track with a puffy `.ds-progress__bar`; width by `--10`, `--25`, `--40`, `--60`, `--75`, `--90`; `.ds-progress__legend` is the line above it. `.ds-progress--ring` is a ring around a clay ball `.ds-progress__face`.
- `.ds-avatar`: initials on a clay ball `.ds-avatar__pic`, alone or with `.ds-avatar__text` (`.ds-avatar__name`, `.ds-avatar__meta`); `--1` to `--4` pick a fill.
- `.ds-accordion`: `<details>` sections as separate clay bars: `.ds-accordion__item`, `.ds-accordion__head` with a round `.ds-accordion__chevron`, `.ds-accordion__body`.
- `.ds-carousel`: `.ds-carousel__track` scrolls and snaps, one `.ds-carousel__slide` showing; `.ds-carousel__foot` holds two round `.ds-carousel__arrow` keys around `.ds-carousel__dots` (`.ds-carousel__dot`, `is-current`). `.ds-shelf` is the slide: `.ds-shelf__plants` (a dented niche with a row of `.ds-pot`), `.ds-shelf__board` (the shelf, a bar of accent clay) and an avatar with a name as caption.
- `.ds-footer`: one clay pill in `--fill-bar-alt` at the end of the page: `.ds-footer__inner` holds the brand, `.ds-footer__links` with `.ds-footer__link`, and `.ds-footer__text`.

## Never

- `box-shadow != none`: nothing raised is flat; a block without its ledge, inner glow and drop shadow is not clay.
- `border-width <= 4px`: clay has no outlines; the only lines are 2px rules and the 4px focus or error ring.
- `box-shadow-blur <= 60px`: the deepest shadow, under the dialog, blurs over 60px.
- `text-shadow = none`: the text is flat; only shapes are modelled.
- `font-size <= 52px`: the hero title is the largest text.
- `font-size >= 14px`: meta text and badges are the smallest text.
- `font-weight >= 500`: the type is round and sturdy; nothing is light.
- `font-families <= 3`: one fat rounded face, one sans and one monospace.
- `line-height <= 1.55`: body text is set at 1.5.
- `letter-spacing = 0`: no tracking on any text.
- `uppercase-text <= 5%`: nothing is set in capitals.
- `underlined-links <= 15%`: links are marked by colour and weight; only a hovered link is underlined.
- `block-gap <= 80px`: blocks sit 34px apart; 80px above the footer is the largest gap.
- `content-width <= 1180px`: the column is capped.

## Extending

Derive a new component from the nearest one in the specimen. Decide first whether it is raised or dented: a container starts from `.ds-panel` (fill, radius and `--shadow-panel` together), something pressed or filled in starts from `.ds-form__input` (`--fill-input`, `--color-input-text`, `--shadow-control-pressed`), something the user presses starts from `.ds-button--secondary` or `.ds-react`. For a coloured card take a `--color-fill-*` background and `--color-accent-text` for everything on it, as `.ds-grid__cell` does. A new illustration is built like `.ds-pot`: rounded boxes in block colours under the shadow tokens, never an image and never a face. Never draw a shadow in a raw colour and never flatten a shape: use the shadow tokens as they are. Use tokens only; for a size with no token use `calc()` over `--space-*` or `--size-*`. Keep each text token on the fill it belongs to.
