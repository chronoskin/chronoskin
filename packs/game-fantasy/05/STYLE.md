# Game fantasy early access (forest green, bronze, moonlight)

## Summary

The early-access or launch page of a new fantasy game as the big studios have built them since about 2019: full-bleed cinematic sections stacked down the page, a thin clear top bar with one bright button, a huge title in display capitals over drawn key art, and feature sections with the picture on alternating sides. The heavy carved frames of the older game sites are gone; the ornament survives as thin gilded rules, diamond dividers, studded buttons and wide-tracked capitals. The page is responsive and the key art is drawn in CSS from the palette.

## Layout

- Full-bleed and fluid. Content sits in `.ds-wrap`, at most `--size-page` (1200px) wide with `--space-5` gutters; pictures, bands and dividers run from edge to edge. Designed at 1440px; under 1000px every two-column block becomes one column and the navigation links fold into the menu of the bar; under 640px grids become one column, figures two by two, and form labels sit above their fields. Nothing scrolls sideways at 390px.
- Order of the home page: the top bar (`--size-nav`, 72px) lying over the hero; the hero (`--size-hero`, 920px): the picture in its upper half and, centred on its dark ground, a kicker, the title, a lead, the platform buttons, the main action and a note; a strip of four figures overlapping the hero's foot; a feature with the trailer; a divider; the carousel of classes on `--color-surface`; a divider; two features with the picture on alternating sides; a band on the page ground with a title plate, the four platform cells and three edition cards; the newsletter band; the footer.
- Inner pages keep the top bar and the footer; the hero is left out. Each starts with `.ds-masthead`, a low picture (`--size-masthead`, 340px) that carries the breadcrumb and the `.ds-page-header` (title, one line, actions on the right) and ends in a gilded rule. Below it one `.ds-section--tight` holds the working parts: `.ds-tabs` across the top, then a `.ds-split` (`--wide` is 3 to 2) or `.ds-columns` (flexible column beside `--size-aside`, 340px), then full-width tables.
- Views: `home`; `classes` (tabs by class, the description, attribute bars, an item window, the skill table); `editions` (the comparison table, the accordion of questions, platforms with status badges); `access` (notices, the account form with switches and buttons, the dialog that deletes an account, the chosen edition); `news` (the list of updates with pagination and links, the roadmap box with a progress bar, an empty streams slot).
- Spacing scale: `--space-1` 4px (label to bar, badge), `--space-2` 8px (small gaps), `--space-3` 12px (between buttons, rows), `--space-4` 16px (inside controls and small boxes), `--space-5` 24px (gutters, panel padding, between cards), `--space-6` 40px (between blocks of a tight section), `--space-7` 64px (between the two sides of a feature, band padding, divider height), `--space-8` 104px (section padding).
- Control heights: `--size-control-small` 30px, `--size-control` 44px, `--size-control-large` 60px. `--size-art` 440px is the least height of a feature picture, `--size-play` 84px the play button, `--size-brand` 40px the mark.
- On a phone (1000px and below) `.ds-nav__links` is not shown and `.ds-nav__more`, a `.ds-menu` in the bar, opens a full-width list of every view as `.ds-menu__item` rows, the sub-pages indented as `.ds-menu__item--sub` and the open view marked by an accent bar at its left; `.ds-nav__action` is hidden, and at 640px and below the summary shows only its bars (`.ds-nav__label` is hidden). `.ds-tabs` becomes one row that scrolls sideways. A `.ds-tooltip__tip` opens under its line between the gutters; a table in `.ds-table__wrap` keeps its cells on one line and scrolls inside the wrapper.

## Typography and colour roles

- Two families and a monospace. `--font-heading` is a high-contrast display serif (Cormorant Garamond where installed, else Cochin, Baskerville, Palatino), always in capitals (`--heading-transform`) and tracked 2px, 5px for kickers and the hero. `--font-body` and `--font-ui` are a humanist sans (Alegreya Sans or Source Sans 3, else Segoe UI, Helvetica Neue); UI text is small, bold, in tracked capitals.
- Sizes: `--text-base` 16px at `--line-body` 1.6; `--text-small` 13px; `--text-ui` 14px; `--text-large` 19px for leads; `--text-h3` 18px, `--text-h2` 25px, `--text-h1` 36px; `--text-display` 52px for section and feature titles and figures. The hero title is 1.75 times the display size and the class name 1.3 times, both capped by the window width.
- Palette: forest black-green page and sheet, three green surfaces, a moss bar with pale gold lettering, a deep night green for pictures and footer (`--color-bar-alt`), sage text under pale gold headings, bronze for second headings, rules (`--color-border-strong`), the accent and the primary button (dark lettering on bronze), and a wisp green (`--color-accent-alt`) for moonlight, diamonds, glows and focus. The inverse block is lichen-pale linen.
- Surface: hairline borders, square corners (only badges are rounded, 2px), long soft shadows under panels, a bronze hairline with a faint green glow around the dialog and the featured card (`--shadow-dialog`), a page of fine vertical grain (`--fill-page`), hover that eases over 0.25s and lights the control in the wisp colour.
- The era's rule for text on fills is kept: on `--color-canvas` and the three surfaces (and `--fill-panel`) text is `--color-text`, `--color-text-muted`, `--color-heading`, `--color-heading-alt` and the link colours; on `--fill-bar` only `--color-bar-text` (table heads, the dialog title, `.ds-plate`); on `--fill-bar-alt` only `--color-bar-alt-text` (every `.ds-scene`, the top bar that lies on it, the footer); on `--fill-inverse` only `--color-inverse-text`; on `--fill-accent` only `--color-accent-text`; on `--fill-button` `--color-button-text`, on `--fill-button-secondary` `--color-button-secondary-text`; in a form well `--color-input-text`; on `--color-notice` `--color-notice-text` and on `--color-danger-surface` `--color-danger`. No text is written on `--color-page`: the page shows in the dividers and behind the platform and edition cards, whose title stands on a plate.
- `--color-fill-1` to `--color-fill-4`, the status colours and `--color-accent-alt` are inks and lights: item names by rarity, crests, status badges on a 16% tint of themselves, grid cells as the surface tinted 24% by their fill, progress bars. Never a fill under text.
- The key art (`.ds-scene`): ground, trunks and silhouettes are `--color-bar-alt` mixed toward `--color-shadow`; the moon is `--color-bar-alt-text` with a halo and shafts of `--color-accent-alt`; the horizon takes `--color-accent`; `.ds-scene--2` lights it with `--color-fill-2` (a rift), `.ds-scene--3` with `--color-fill-3` (a fire). Every palette repaints it.
- One-off values built in `components.css` from tokens: the mixes named above, percentages in `clip-path` polygons and gradients, `50%` radius for the play button and lamps, 45 degree turns for diamonds and studs, `min()` of a size token and a share of the window width for the largest titles.

## Components

- `.ds-page`: on `<body>`; paints the page ground. `.ds-wrap` is the centred content column.
- `.ds-nav`: the clear top bar over the picture: the brand, `.ds-nav__links` of `.ds-nav__link` (the current one underlined in the rule colour with a diamond; `is-current`, or in the specimen one selector per view), and `.ds-nav__end` with a `.ds-nav__action` link and one small primary button. One item is a `.ds-menu`.
- `.ds-nav__more`: the phone navigation, a `.ds-menu` hidden above 1000px. Its summary (`.ds-menu__summary`: `.ds-nav__bars` and a `.ds-nav__label`) opens a sheet of every view, the sub-pages as `.ds-menu__item--sub`.
- `.ds-brand`: `.ds-brand__mark`, a 40px tile dressed as the primary button (`--fill-button`, `--color-button-text`, its border, `--radius-control`, `--shadow-control`), and `.ds-brand__name` in the heading face. Both are placeholders for the installing project's own name and logo.
- `.ds-hero`: the full picture of the home view, a `.ds-scene` with `.ds-hero__copy` centred on its ground: `.ds-kicker`, `.ds-hero__title`, `.ds-hero__lead`, `.ds-platforms` (a row of secondary buttons), `.ds-hero__actions` (one large button), `.ds-hero__note`.
- `.ds-scene`: the drawn picture, used for the hero, mastheads, feature art and the class portrait. Ten empty spans in this order: `.ds-scene__sky`, `.ds-scene__light` (shafts), `.ds-scene__far` and `.ds-scene__mid` (two tree lines), `.ds-scene__arch` (the ruined arch, lit on its inner edges), `.ds-scene__mist`, `.ds-scene__near` (ground and a traveller), `.ds-scene__trunks` (trunks and canopy at the edges), `.ds-scene__flies`, `.ds-scene__shade`. Variants `.ds-scene--2`, `.ds-scene--3`.
- `.ds-masthead`: the low scene that starts an inner page; holds a `.ds-wrap` with the breadcrumb and the page header.
- `.ds-page-header`: `.ds-page-header__title`, `.ds-page-header__text`, `.ds-page-header__actions`; written in the scene's text colour.
- `.ds-section`: a full-bleed band in `--color-canvas`; `--alt` in `--color-surface`, `--ground` clear so the page shows, `--tight` with less padding. `.ds-section__head` centres a `.ds-section__kicker`, `.ds-section__title`, a `.ds-rule` and perhaps a `.ds-section__lead`. `.ds-plate` is the title of a ground section, on `--fill-bar` with a diamond at each end. `.ds-divider` is a strip of page ground between two bands with a gilded rule and a diamond.
- `.ds-feature`: a picture and text side by side; `.ds-feature--flip` puts the picture right. `.ds-feature__art` (a scene in a double gilded line; may hold a `.ds-play` button and a `.ds-feature__caption`), `.ds-feature__copy` with a kicker, `.ds-feature__title`, `.ds-feature__text`, `.ds-feature__points` (a list with diamond bullets).
- `.ds-editions`: three `.ds-edition` cards (a `.ds-panel` each) with a `.ds-price`, points and a button; `.ds-edition--featured` is lit and may carry a `.ds-edition__flag` badge on its top edge.
- `.ds-prose`: running text: h1, h2 (ruled underneath), h3 (in `--color-heading-alt`), paragraphs, lists, `code`, links. `.ds-prose__lead` is the opening line.
- `.ds-link`: text link; states `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--quiet`. `.ds-links` is a wrapping row.
- `.ds-button`: a flat plate with a diamond stud on each end, 44px high. `.ds-button--secondary`; `.ds-button--danger` (secondary plate, danger ink and edge); `.ds-button--large` (60px, heading face); `.ds-button--small` (30px, no studs). States `is-hover` (lit through `--shadow-control-hover`), `is-active`, `is-focus`, `is-disabled`. `.ds-buttons` is a row.
- `.ds-form`: rows (`.ds-form__row`) of a `.ds-form__label` (170px; above the field on a phone) and a `.ds-form__field`. `.ds-form__input`, `.ds-form__select`, `.ds-form__textarea`; `.ds-form__input--short`; `is-error` with `.ds-form__error`; `.ds-form__hint`; `.ds-form__check` with `.ds-form__checkbox`; `.ds-form__actions`, `.ds-form__actions-end`. `.ds-form--inline` is one field and one button in a line (the newsletter band).
- `.ds-switch`: an on and off row over a checkbox: `.ds-switch__label`, `.ds-switch__hint`, the hidden `.ds-switch__input` and the `.ds-switch__track`, which takes `--fill-accent` when on (`is-on` by class).
- `.ds-table`: a data table in a flush panel body with `.ds-table__wrap` so that it scrolls inside itself on a phone. The head row takes `--fill-bar`; `is-numeric`, `is-centre`; `.ds-table__sub`, `.ds-table__who`; `is-own` marks a row.
- `.ds-list`: ruled rows: `.ds-list__item` with an optional `.ds-list__date`, a `.ds-list__body` (`.ds-list__title`, `.ds-list__text`, `.ds-list__meta`) and perhaps a badge. `.ds-list--tight` is the dense form.
- `.ds-panel`: a box with a gilded top edge: `.ds-panel__title`, `.ds-panel__body` (`--flush`), `.ds-panel__text`, `.ds-panel__foot`.
- `.ds-stat`: one figure, `.ds-stat__value` over `.ds-stat__label`; four sit in the `.ds-stats` strip that overlaps the hero's foot.
- `.ds-grid`: four `.ds-grid__cell` across (`--two`, `--three`; two, then one on narrow windows), tinted and edged by `--color-fill-1`; `.ds-grid__cell--2`, `--3`, `--4`. Inside: `.ds-grid__icon`, `.ds-grid__title`, `.ds-grid__text`, a badge or `.ds-grid__more`.
- `.ds-tabs`: a row of text tabs on a rule; `.ds-tabs__tab`, `is-current` underlined in the rule colour; `.ds-tabs__count`.
- `.ds-badge`: a small label in `--fill-accent`; `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger`, `.ds-badge--quiet`, `.ds-badge--new`. `.ds-badges` is a row; `.ds-dot` with `--success`, `--warning`, `--danger` a lamp.
- `.ds-item`: an item name in its rarity colour: plain, `--uncommon`, `--rare`, `--epic`, `--legendary`, `--set`.
- `.ds-tooltip`: the item window: a trigger wrapping a `.ds-tooltip__tip` shown on hover or focus (`is-open`, `.ds-tooltip--end`), with `.ds-tooltip__name` and `.ds-tooltip__line` rows (`--bonus`, `--muted`, `--lore`). `.ds-tooltip--card` writes the window into the page.
- `.ds-progress`: a track with a `.ds-progress__bar`; width by `--10` to `--100`, colour by `--health`, `--xp`, `--rep`, `--danger`; `.ds-progress__label` above.
- `.ds-avatar`: a class crest (inline SVG) on a `--color-surface-strong` tile; `.ds-avatar--small`, `.ds-avatar--large`; `--1` to `--4` colour it.
- `.ds-carousel`: the classes: `.ds-carousel__stage` holds one `.ds-carousel__slide` (`.ds-carousel__copy` with `.ds-carousel__name`, `.ds-carousel__text`, `.ds-carousel__bars`; and `.ds-carousel__portrait`, a scene with a `.ds-carousel__sigil`) between two `.ds-carousel__arrow`s (`--next`); `.ds-carousel__foot` holds `.ds-carousel__picks` (a `.ds-carousel__pick` crest per class, `is-current` lit) and `.ds-carousel__dots` of `.ds-carousel__dot`.
- `.ds-accordion`: `<details>` sections: `.ds-accordion__item`, `.ds-accordion__summary`, `.ds-accordion__meta`, `.ds-accordion__body`.
- `.ds-menu`: the drop-down under a navigation item: `.ds-menu__list` of `.ds-menu__item`, shown on hover or focus (`is-open`); `.ds-menu__caret`.
- `.ds-columns`: a flexible column beside a 340px one. `.ds-sidebar`: a box of the narrow column: `.ds-sidebar__title`, a `.ds-sidebar__list` of `.ds-sidebar__item` rows with a `.ds-sidebar__link` and `.ds-sidebar__meta`.
- `.ds-notice`: a ruled message with `.ds-notice__title` and `.ds-notice__text`; `.ds-notice--error`. `.ds-notices` stacks several.
- `.ds-pagination`: `.ds-pagination__link`; `is-current` in `--fill-accent`, `is-disabled` flat.
- `.ds-breadcrumb`: `.ds-breadcrumb__link`, `.ds-breadcrumb__sep` (a diamond), `.ds-breadcrumb__current`; it inherits the colour of the scene it lies on.
- `.ds-dialog`: `.ds-dialog__title` on `--fill-bar` with `.ds-dialog__close`, `.ds-dialog__body`, `.ds-dialog__text`, `.ds-dialog__actions`; `.ds-dialog__backdrop` behind it.
- `.ds-empty`: an empty slot: `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text`, one secondary button.
- `.ds-band`: the newsletter call in `--fill-inverse` between two gilded rules: `.ds-band__title`, `.ds-band__text` and an inline form. At most one per page.
- `.ds-footer`: the closing band in `--fill-bar-alt`: centred `.ds-footer__links` of `.ds-footer__link`, a `.ds-rule`, `.ds-footer__note`.
- `.ds-view`: one screen of the example site; only the one named in the address shows, the home view when none is.
- `.ds-heading` (`--centre`), `.ds-rule`, `.ds-stack` (`--tight`), `.ds-split` (`--wide`), `.ds-row` (`--between`), `.ds-muted`, `.ds-icon`, `.ds-price`: helpers.

## Never

- `border-radius <= 2px`: everything is square; only badges are eased by 2px, and lamps and the play button are round.
- `border-width <= 2px`: borders are hairlines.
- `font-size >= 13px`: meta, badges and labels are the smallest text.
- `font-size <= 91px`: the hero title is the largest text.
- `font-families <= 3`: a display serif, a humanist sans and a monospace.
- `line-height <= 1.6`: body text is set at 1.6.
- `letter-spacing <= 5px`: kickers and the hero title are tracked 5px, headings 2px.
- `box-shadow-blur <= 80px`: the longest shadow is the dialog's.
- `animation = none`: nothing moves by itself; only hovers ease.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new part of the home page is a `.ds-section` with a centred head, separated from its neighbours by a `.ds-divider`; put pictures in a `.ds-feature` and alternate their side. A new picture is a `.ds-scene` with its ten spans and, for another light, a variant built from one fill colour. Boxes are panels with a gilded top edge, not frames; ornament is a hairline, a diamond or a stud. Keep the rule for text on fills, keep rarity and status colours as ink, and write nothing on the page ground without a plate. Do not add images, rounded cards, a second accent or movement.
