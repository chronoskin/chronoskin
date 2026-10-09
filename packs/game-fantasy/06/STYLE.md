# Game fantasy live hub (gunmetal, white, electric violet)

## Summary

The hub of a live science-fantasy game as such sites have looked since about 2019: a clear glass top bar over a full-bleed picture with the season's title set large and left, a strip of numbered anchors, and below it decks of news, the season pass, the roadmap and the week's figures. It is the cold, sleek end of the era: condensed capitals, monospace labels, hairlines, cut corners and one electric colour, with the old ornament reduced to diamonds and brackets. The page is responsive and the key art is drawn in CSS from the palette.

## Layout

- Full-bleed and fluid, everything aligned left. The hero and mastheads run from edge to edge; their content sits in `.ds-wrap` (at most `--size-page`, 1320px, with `--space-6` gutters). Below them `.ds-decks` is a column of `.ds-deck` slabs in `--color-canvas` with two cut corners, laid `--space-5` apart on the page ground, whose grid shows between and around them. Designed at 1440px; under 1100px two-column blocks become one column and grids two across; under 700px the navigation links move to a second row and grids become one column. Nothing scrolls sideways at 390px; the pass track and tables scroll inside themselves.
- Order of the home page: the top bar (`--size-nav`, 64px) lying over the hero; the hero (`--size-hero`, 780px) with the tag, title, lead, two large buttons and a note at the bottom left and the reader's season card (`--size-card`, 400px) at the bottom right; the `.ds-subnav` strip; then decks: the news mosaic, the season pass (progress and a track of tier cards), the roadmap grid, the week's figures; the closing band; the footer in four columns.
- Inner pages keep the top bar and the footer; the hero is left out. Each starts with `.ds-masthead`, a low picture (`--size-masthead`, 300px) that carries the breadcrumb and the `.ds-page-header` and ends in a 2px accent line. Below it one deck holds the working parts: `.ds-tabs` or notices across the top, then `.ds-columns` (flexible column beside `--size-aside`, 360px) or a `.ds-split` (`--wide` is 3 to 2).
- Views: `home`; `season` (tabs by track, the tier table with an item tooltip, pagination, an item window, display switches); `news` (the patch notes, the list of posts, topics, links); `community` (crews recruiting with avatars, the accordion of rules, the empty crew slot, creators live); `account` (notices, the settings form with switches and buttons, the dialog that unlinks a platform, linked platforms).
- Spacing scale: `--space-1` 4px (label to bar, tag), `--space-2` 8px (small gaps), `--space-3` 12px (between buttons and tier cards), `--space-4` 16px (between tiles, inside small boxes), `--space-5` 24px (between decks and columns, panel padding), `--space-6` 36px (deck padding, gutters), `--space-7` 56px (hero and band padding, the star grid), `--space-8` 88px (the accent line on a deck).
- Control heights: `--size-control-small` 30px, `--size-control` 42px, `--size-control-large` 56px. `--size-tile` 230px is a row of the mosaic, `--size-tier` 168px a tier card, `--size-cut` 20px a cut corner or bracket.
- On a tablet or phone (1100px and below) the bar wraps: the brand and `.ds-nav__end` on the first row, `.ds-nav__links` below as one row that scrolls sideways, with the items of the menu flattened into it (its caret is hidden), the row fading at the right edge and the link of the open view moved to the front on the later views. At 700px and below `.ds-subnav` and `.ds-tabs` scroll sideways too, and `.ds-subnav__note` is hidden; a `.ds-tooltip__tip` opens under its line between the gutters; a table in `.ds-table__wrap` keeps its cells on one line and scrolls inside the wrapper.

## Typography and colour roles

- Three families. `--font-heading` and `--font-ui` are a bold condensed sans (Barlow Condensed where installed, else Oswald, Roboto Condensed, Arial Narrow), always in capitals, for titles, navigation, buttons, table heads and badges. `--font-body` is a plain sans (Inter, Roboto, Segoe UI). `--font-mono` labels things the way a game's interface does: breadcrumbs, dates, indexes, tier numbers, figure labels.
- Sizes: `--text-base` 15px at `--line-body` 1.55; `--text-small` 12px; `--text-ui` 16px; `--text-large` 18px; `--text-h3` 20px, `--text-h2` 27px, `--text-h1` 40px for deck titles; `--text-display` 60px for mastheads, figures and the band. The hero title is 1.6 times the display size, capped by the window width.
- Palette: gunmetal black page and sheet, three graphite surfaces, a lighter steel bar with white lettering, near-black for pictures and footer (`--color-bar-alt`), cool grey text under white headings, and one electric violet for the accent, links, the primary button and every light in the key art; `--color-accent-alt` is its pale spark. The inverse block is white.
- Surface: hairline borders, corners of 2px to 6px and fully round badges, panels that let a little of what is behind them through and blur it (`--fill-panel`, `--backdrop-blur` 14px, used by the top bar and the hero card), a page ruled as a 48px grid (`--fill-page`), a violet ring and glow on hover and around the dialog, hover easing over 0.15s.
- The era's rule for text on fills is kept: on `--color-canvas` and the three surfaces (and `--fill-panel`) text is `--color-text`, `--color-text-muted`, `--color-heading`, `--color-heading-alt` and the link colours; on `--fill-bar` only `--color-bar-text` (the subnav, table heads, the dialog title); on `--fill-bar-alt` only `--color-bar-alt-text` (every `.ds-scene` and what lies on it: the top bar, hero copy, mastheads, tile titles; and the footer); on `--fill-inverse` only `--color-inverse-text`; on `--fill-accent` only `--color-accent-text`; on `--fill-button` `--color-button-text`, on `--fill-button-secondary` `--color-button-secondary-text`; in a form well `--color-input-text`; on `--color-notice` `--color-notice-text` and on `--color-danger-surface` `--color-danger`. No text is written on `--color-page`.
- `--color-fill-1` to `--color-fill-4`, the status colours and `--color-accent-alt` are inks and lights: item names by rarity, crests, status badges on a 16% tint of themselves, roadmap cells as the surface tinted 24% by their fill, progress bars, the edge of a claimed or current tier. Never a fill under text.
- The key art (`.ds-scene`): silhouettes are `--color-bar-alt` mixed toward `--color-shadow`; the horizon, the beam, the seams of the spires and the floor lines are `--color-accent`; the ringed world is rimmed in `--color-bar-alt-text`; `.ds-scene--2`, `--3` and `--4` light the sky and the beam with `--color-fill-1`, `--color-fill-3` and `--color-fill-4`. Every palette repaints it.
- One-off values built in `components.css` from tokens: the mixes named above, percentages in `clip-path` polygons and gradients, `50%` radius for avatars and lamps, 45 degree turns for diamonds, `min()` of a size token and a share of the window width for the largest titles.

## Components

- `.ds-page`: on `<body>`; paints the ruled page ground. `.ds-wrap` is the content column of the hero, mastheads, subnav and footer.
- `.ds-nav`: the clear top bar: the brand, then `.ds-nav__links` of `.ds-nav__link` straight after it (the current one fully opaque with a line of `--color-accent-alt` under it; `is-current`, or in the specimen one selector per view), and `.ds-nav__end` with a `.ds-nav__status` lamp and one small primary button. One item is a `.ds-menu`.
- `.ds-brand`: `.ds-brand__mark`, a 36px tile dressed as the primary button (`--fill-button`, `--color-button-text`, its border, `--radius-control`, `--shadow-control`), and `.ds-brand__name` in the heading face. Both are placeholders for the installing project's own name and logo.
- `.ds-hero`: the season's picture, a `.ds-scene` whose `.ds-wrap` holds `.ds-hero__copy` (a `.ds-tag`, `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__actions`, `.ds-hero__note`) and, at the right, a panel with `.ds-hero__card` that blurs the picture behind it.
- `.ds-scene`: the drawn picture, used for the hero, mastheads and news tiles. Nine empty spans in this order: `.ds-scene__sky`, `.ds-scene__world` (the ringed world), `.ds-scene__far` (a skyline), `.ds-scene__beam`, `.ds-scene__spires`, `.ds-scene__haze`, `.ds-scene__floor` (lines to a vanishing point), `.ds-scene__near` (rocks and a lone figure), `.ds-scene__shade` (darkens the left for the text). Variants `.ds-scene--2`, `--3`, `--4`.
- `.ds-tag`: a slanted label on `--fill-accent`.
- `.ds-subnav`: the strip under the hero on `--fill-bar`: numbered `.ds-subnav__link`s and a `.ds-subnav__note` at the right.
- `.ds-masthead`: the low scene that starts an inner page; holds the breadcrumb and the page header.
- `.ds-page-header`: `.ds-page-header__title`, `.ds-page-header__text`, `.ds-page-header__actions`; written in the scene's text colour.
- `.ds-decks`: the column of slabs; `.ds-deck` one slab (`--alt` in `--color-surface`) with a short accent line on its top edge and a `.ds-deck__head`: `.ds-deck__index` (a number), `.ds-deck__title`, `.ds-deck__more` (a link at the right).
- `.ds-mosaic`: the news tiles: a `.ds-tile` is a scene with a badge, `.ds-tile__title` and `.ds-tile__meta` at its foot; `.ds-tile--lead` is two rows tall, `.ds-tile--wide` two columns wide; `is-hover` lights its edge.
- `.ds-pass`: the pass progress beside its button. `.ds-carousel`: the track of tiers: `.ds-carousel__stage` scrolls sideways inside itself with scroll snapping and holds `.ds-carousel__slide` cards (`.ds-carousel__tier`, a square avatar, an item name; `is-claimed`, `is-current`, `is-locked`); `.ds-carousel__foot` holds a caption, `.ds-carousel__dots` of `.ds-carousel__dot` and `.ds-carousel__arrows` with two `.ds-carousel__arrow`s (`--next`).
- `.ds-who`: an avatar beside `.ds-who__name` and `.ds-who__meta`.
- `.ds-prose`: running text: h1, h2 (ruled underneath), h3 (in `--color-heading-alt`), paragraphs, lists, `code`, links. `.ds-prose__lead` is the opening line.
- `.ds-link`: text link; states `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--quiet`. `.ds-links` is a wrapping row.
- `.ds-button`: a flat plate, 42px high, with a small lit tick in its top right corner. `.ds-button--secondary`; `.ds-button--danger` (secondary plate, danger ink and edge); `.ds-button--large` (56px); `.ds-button--small` (30px). States `is-hover` (ring and glow from `--shadow-control-hover`), `is-active`, `is-focus`, `is-disabled`. `.ds-buttons` is a row.
- `.ds-form`: rows (`.ds-form__row`) of a `.ds-form__label` (180px; above the field on a phone) and a `.ds-form__field`. `.ds-form__input`, `.ds-form__select`, `.ds-form__textarea`; `.ds-form__input--short`; `is-error` with `.ds-form__error`; `.ds-form__hint`; `.ds-form__check` with `.ds-form__checkbox`; `.ds-form__actions`, `.ds-form__actions-end`. `.ds-form--inline` is one field and one button in a line.
- `.ds-switch`: an on and off row over a checkbox: `.ds-switch__label`, `.ds-switch__hint`, the hidden `.ds-switch__input` and the `.ds-switch__track`, which takes `--fill-accent` when on (`is-on` by class). Used for settings and display options.
- `.ds-table`: a data table in a flush panel body with `.ds-table__wrap`. The head row takes `--fill-bar`; `is-numeric`, `is-centre`; `.ds-table__sub` a second line; `.ds-table__who`; `is-own` marks the reader's row.
- `.ds-list`: ruled rows: `.ds-list__item` with an optional `.ds-list__date` (monospace), a `.ds-list__body` (`.ds-list__title`, `.ds-list__text`, `.ds-list__meta`) and perhaps a badge. `.ds-list--tight` is the dense form.
- `.ds-panel`: a box with a bracket in `--color-accent-alt` on its top left corner: `.ds-panel__title` (may hold a badge at the right), `.ds-panel__body` (`--flush`), `.ds-panel__text`, `.ds-panel__foot`.
- `.ds-stat`: one figure, `.ds-stat__value` over a monospace `.ds-stat__label`, with a rule at its left; four sit in a `.ds-stats` row.
- `.ds-grid`: four `.ds-grid__cell` across (`--two`, `--three`), tinted and edged by `--color-fill-1` with a heavier top edge; `.ds-grid__cell--2`, `--3`, `--4`. Inside: `.ds-grid__when` (a monospace date), `.ds-grid__title`, `.ds-grid__text`, a badge; `.ds-grid__icon` and `.ds-grid__more` are available.
- `.ds-tabs`: text tabs on a rule; `.ds-tabs__tab`, `is-current` underlined in the accent; `.ds-tabs__count` a round counter.
- `.ds-badge`: a round-ended label in `--fill-accent`; `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger`, `.ds-badge--quiet`, `.ds-badge--new`. `.ds-badges` is a row; `.ds-dot` with `--success`, `--warning`, `--danger` a lamp.
- `.ds-item`: an item name in its rarity colour: plain, `--uncommon`, `--rare`, `--epic`, `--legendary`, `--set`.
- `.ds-tooltip`: the item window: a trigger wrapping a `.ds-tooltip__tip` shown on hover or focus (`is-open`, `.ds-tooltip--end`), with `.ds-tooltip__name` and `.ds-tooltip__line` rows (`--bonus`, `--muted`, `--lore`). `.ds-tooltip--card` writes the window into the page.
- `.ds-progress`: a thin track with a `.ds-progress__bar`; width by `--10` to `--100`, colour by `--health`, `--xp`, `--rep`, `--danger`; `.ds-progress__label` above.
- `.ds-avatar`: a round emblem (inline SVG) on a `--color-surface-strong` disc; `.ds-avatar--small`, `.ds-avatar--large`; `--1` to `--4` colour it. In a tier card it is square.
- `.ds-accordion`: `<details>` sections: `.ds-accordion__item`, `.ds-accordion__summary`, `.ds-accordion__meta`, `.ds-accordion__body`.
- `.ds-menu`: the drop-down under a navigation item: `.ds-menu__list` of `.ds-menu__item`, shown on hover or focus (`is-open`); `.ds-menu__caret`.
- `.ds-columns`: a flexible column beside a 360px one. `.ds-sidebar`: a box of the narrow column: `.ds-sidebar__title`, a `.ds-sidebar__list` of `.ds-sidebar__item` rows with a `.ds-sidebar__link` and `.ds-sidebar__meta`.
- `.ds-notice`: a ruled message with `.ds-notice__title` and `.ds-notice__text`; `.ds-notice--error`. `.ds-notices` stacks several.
- `.ds-pagination`: `.ds-pagination__link`; `is-current` in `--fill-accent`, `is-disabled` flat.
- `.ds-breadcrumb`: a monospace trail: `.ds-breadcrumb__link`, `.ds-breadcrumb__sep` (a slash), `.ds-breadcrumb__current`; it inherits the colour of the scene it lies on.
- `.ds-dialog`: `.ds-dialog__title` on `--fill-bar` with `.ds-dialog__close`, `.ds-dialog__body`, `.ds-dialog__text`, `.ds-dialog__actions`; `.ds-dialog__backdrop` behind it.
- `.ds-empty`: an empty slot: `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text`, one secondary button.
- `.ds-band`: the closing call in `--fill-inverse`, a slab with cut corners: `.ds-band__title`, `.ds-band__text` and an inline form. At most one per page.
- `.ds-footer`: the closing band in `--fill-bar-alt`: `.ds-footer__about` (a home link and `.ds-footer__note`) and three columns, each a monospace `.ds-footer__title` over `.ds-footer__links` of `.ds-footer__link`.
- `.ds-view`: one screen of the example site; only the one named in the address shows, the home view when none is.
- `.ds-heading` (`--centre`), `.ds-rule`, `.ds-stack` (`--tight`), `.ds-split` (`--wide`), `.ds-row` (`--between`), `.ds-muted`, `.ds-icon`: helpers.

## Never

- `border-radius <= 6px`: corners are barely eased; only badges, switches, avatars and lamps are round.
- `border-width <= 2px`: borders are hairlines, accent lines 2px.
- `font-size >= 12px`: monospace labels are the smallest text.
- `font-size <= 96px`: the hero title is the largest text.
- `font-families <= 3`: a condensed sans, a plain sans and a monospace.
- `line-height <= 1.55`: body text is set at 1.55.
- `letter-spacing <= 1.2px`: capitals are tracked by a fiftieth of their size, 1.2px at the display size.
- `underlined-links <= 5%`: links are told apart by the violet; only a hovered link is underlined.
- `box-shadow-blur <= 90px`: the longest shadow is the dialog's.
- `animation = none`: nothing moves by itself; only hovers ease.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new part of the home page is a `.ds-deck` with a numbered head; never a free-floating card on the page ground, and never text on it. A box inside a deck is a `.ds-panel` with its bracket; a picture is a `.ds-scene` with its nine spans and, for another light, a variant built from one fill colour. Label machine-like facts (dates, counts, indexes) in the monospace and everything a person would say in the sans. Keep the rule for text on fills and keep rarity and status colours as ink. Do not add images, carved frames, serif type, a second accent or movement.
