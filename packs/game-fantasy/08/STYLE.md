# Game fantasy card table (tavern wood, green felt, brass)

## Summary

The site of a collectible card game as such sites have looked since about 2014: a wooden beam of navigation with the mark hanging from it, a felt table in a wooden rail with cards fanned on it, and below it boards nailed to a plank wall that hold the new cards, tavern news, the deck of the week and the sets in play. It is the warm, chunky end of the era: walnut and oak, green felt, brass rims and nails, sapphire gems for costs and buttons, a bold slab serif over a plain sans, everything with a thick lower lip as if cut from wood. The page is responsive, and every card is drawn in CSS from the palette: a framed small scene with a cost gem, a name plate and a text box.

## Layout

- One centred column on the plank wall. `.ds-wrap` is at most `--size-page` (1240px) with `--space-5` gutters and stacks its blocks `--space-5` apart; the wall (`--fill-page`) shows around and between them, and no text is written on it. The navigation beam and the footer run from edge to edge. Designed at 1440px; under 1100px two-column blocks become one column and grids two across; under 700px the navigation links move to a second row, the gallery shows two cards across and grids become one column. Nothing scrolls sideways at 390px; the hand of cards and tables scroll inside themselves.
- Order of the home page: the beam (`--size-nav`, 68px) with the brand's gem hanging below its edge; the hero, a `.ds-felt` table inside a wooden rail, copy on the left and four fanned cards (one face down) on the right; the `.ds-shelf` of figures; a `.ds-board` with the carousel of new cards; `.ds-columns` (a flexible board with the news list beside `--size-aside`, 360px: the deck of the week panel with its mana curve, and the collection box); a board with the four sets as grid cells; the closing band on card stock; the footer, a strip of felt under a wooden rail.
- Inner pages keep the beam and the footer; the hero is left out. Each is one `.ds-board` that starts with the breadcrumb and the `.ds-page-header` (title and one line on the left, buttons on the right, a brass rule under it), then the working parts: `.ds-library` (the filter rail, `--size-rail` 270px, beside tabs, the gallery and pagination), `.ds-showcase` (one large card on its mat, 2 to 3 beside the text), `.ds-columns`, or a `.ds-split` (`--wide` is 3 to 2).
- Views: `home`; `cards` (classes and switches in the rail, tabs by kind, eight cards, pagination); `card` (the large card, a keyword window, notes on the card with keyword tooltips, related links and badges); `decks` (notices, the deck list as a table, the accordion of questions, the mana curve, an empty deck slot); `build` (the deck form with a switch and its buttons, the dialog that deletes a deck, the draft so far).
- Spacing scale: `--space-1` 4px (label to bar, inside a card), `--space-2` 8px (small gaps, card padding), `--space-3` 12px (between buttons, tight stacks), `--space-4` 16px (between cards and cells), `--space-5` 24px (between blocks and columns, panel padding), `--space-6` 40px (board padding), `--space-7` 60px (felt padding), `--space-8` 88px (the grid of dust in the lamp light).
- Control heights: `--size-control-small` 32px, `--size-control` 44px, `--size-control-large` 56px. `--size-card` 190px is a card in a hand or gallery, `--size-card-large` 330px the card of its own page, `--size-gem` 38px a cost gem, `--size-curve` 120px the mana curve.
- On a tablet or phone (1000px and below) the beam wraps: the brand and `.ds-nav__end` on the first row, `.ds-nav__links` below across the full width as a row of equal `.ds-nav__link` cells with centred text, the `.ds-menu` among them opening a list as wide as the row; `.ds-nav__purse` is already hidden below 1100px, and at 700px and below `.ds-tabs` scrolls sideways. Where the page is one column (1100px and below) a `.ds-tooltip__tip` opens under its line between the gutters; a table in `.ds-table__wrap` keeps its cells on one line and scrolls inside the wrapper.

## Typography and colour roles

- Three families. `--font-heading` and `--font-ui` are a bold slab serif (Zilla Slab or Roboto Slab where installed, else Rockwell), in mixed case, for titles, navigation, buttons, tabs, badges, card names and numbers. `--font-body` is a plain sans (Open Sans, Noto Sans, Lucida Grande) for everything read, including card text. `--font-mono` only appears in inline code.
- Sizes: `--text-base` 16px at `--line-body` 1.6; `--text-small` 13px (card text, meta); `--text-ui` 17px; `--text-large` 20px; `--text-h3` 19px, `--text-h2` 26px, `--text-h1` 38px for board titles; `--text-display` 56px for page titles. The hero title is 1.25 times the display size, capped by the window width.
- Palette: a dark walnut wall, darker brown boards and three brown surfaces, an oak bar with cream lettering (`--color-bar`), green felt for the table, mats and footer (`--color-bar-alt`), cream text under paler headings, brass for every rim, rule and second heading (`--color-border-strong`, `--color-heading-alt`), a sapphire accent and primary button, bright brass for sparks and nails (`--color-accent-alt`), and cream card stock as the inverse block.
- Surface: 2px borders and 3px brass rims, corners of 8px to 18px, a lit top edge and a dark lower lip on every panel and button (`--shadow-panel`, `--shadow-control`), a brass glow on hover, planks on the wall (`--fill-page`), grain on the bars, a fine weave on the felt, gems with a highlight at the top, quick 0.12s easing. No glass: `--backdrop-blur` is 0.
- The era's rule for text on fills is kept: on `--color-canvas` and the three surfaces (and `--fill-panel`) text is `--color-text`, `--color-text-muted`, `--color-heading`, `--color-heading-alt` and the link colours; on `--fill-bar` only `--color-bar-text` (the beam, the shelf and its figures, table heads, a card's name plate, the dialog title); on `--fill-bar-alt` only `--color-bar-alt-text` (the felt, a mat, the footer); on `--fill-inverse` only `--color-inverse-text` (a card's text box, the date leaf of a news post, the band); on `--fill-accent` only `--color-accent-text` (cost gems, the current navigation item and tab, the tag); on `--fill-button` `--color-button-text`, on `--fill-button-secondary` `--color-button-secondary-text` (also a card's attack and health); in a form well `--color-input-text`; on `--color-notice` `--color-notice-text` and on `--color-danger-surface` `--color-danger`. No text is written on `--color-page`.
- `--color-fill-1` to `--color-fill-4`, the status colours and `--color-accent-alt` are inks and lights: card names by rarity (`.ds-item`), the sky of a card's art, class crests, set cells as the surface tinted 24% by their fill, status badges on a 16% tint of themselves, progress bars, the edge of an attack or health figure. Never a fill under text.
- The pictures. `.ds-card__art`: the sky is one fill colour over `--color-bar-alt`, the low sun and the stars are `--color-accent-alt`, the hills and the creature's crest are `--color-bar-alt` mixed toward `--color-shadow`, and a vignette of the shadow colour closes the corners. `.ds-felt`: `--fill-bar-alt` under lamp light and dust in `--color-accent-alt`, a chalk ring in `--color-bar-alt-text`, the rail in `--fill-bar`. Every palette repaints them.
- One-off values built in `components.css` from tokens: the mixes named above, percentages in `clip-path` polygons and gradients, `50%` radius for gems, avatars and coins, turns of 9 and 16 degrees for fanned cards, a card's 5 to 7 proportion, `min()` of a size token and a share of the window width for the largest titles.

## Components

- `.ds-page`: on `<body>`; paints the plank wall. `.ds-wrap` is the column of a view.
- `.ds-nav`: the wooden beam: the brand, `.ds-nav__links` of `.ds-nav__link` (the current one a gem plate in `--fill-accent`; `is-current`, or in the specimen one selector per view), and `.ds-nav__end` with a `.ds-nav__purse` (a `.ds-coin` and a count) and one small primary button. One item is a `.ds-menu`.
- `.ds-brand`: `.ds-brand__mark`, a 60px tile dressed as the primary button (`--fill-button`, `--color-button-text`, the brass rim, `--radius-control`, `--shadow-control`) that hangs below the beam, and `.ds-brand__name` in the heading face. Both are placeholders for the installing project's own name and logo.
- `.ds-hero`: the wooden rail around a `.ds-felt` table. The felt holds `.ds-hero__copy` (a `.ds-tag`, `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__actions`, `.ds-hero__note`) and a `.ds-fan` of four cards, the first face down.
- `.ds-card`: one card, 5 to 7: `.ds-card__cost` (the gem), `.ds-card__art` (holding one inline SVG crest), `.ds-card__name` (the plate), `.ds-card__text` (the text box) and `.ds-card__stats` (a `.ds-card__stat` for attack, the kind, a `.ds-card__stat--life` for health; spells leave the figures out). `.ds-card--2`, `--3`, `--4` change the colour of the art; `.ds-card--large` is the card of its own page; `.ds-card--back` is a card face down. As a link in a hand or gallery it lifts on hover (`is-hover`).
- `.ds-tag`: a gem label on `--fill-accent`.
- `.ds-shelf`: the oak shelf under the hero that holds the `.ds-stats` row.
- `.ds-board`: a slab in `--color-canvas` with a brass nail in each corner; `.ds-board__head` holds `.ds-board__title`, `.ds-board__sub` and `.ds-board__more` over a brass rule.
- `.ds-carousel`: one hand of cards showing: `.ds-carousel__stage` holds a `.ds-carousel__slide` (a row of cards that scrolls inside itself with scroll snapping on a narrow window) and two round `.ds-carousel__arrow`s (`--next`); `.ds-carousel__foot` a caption and `.ds-carousel__dots` of `.ds-carousel__dot`.
- `.ds-keyword`: a keyword in running text, the trigger of a tooltip.
- `.ds-curve`: the mana curve: `.ds-curve__col` per cost with a count, a `.ds-curve__bar` (height by `--25`, `--40`, `--60`, `--75`, `--100`) and a `.ds-curve__cost`.
- `.ds-library`: the filter rail beside the main column. `.ds-gallery`: the grid of cards, each a `.ds-gallery__item` with the card, its name as a `.ds-item` and a `.ds-gallery__meta`.
- `.ds-showcase`: the large card beside the text. `.ds-mat`: the felt mat a large card lies on, with a `.ds-mat__note`.
- `.ds-who`: an avatar beside `.ds-who__name` and `.ds-who__meta`.
- `.ds-page-header`: `.ds-page-header__title`, `.ds-page-header__text`, `.ds-page-header__actions`, a brass rule below.
- `.ds-prose`: running text: h1, h2 (ruled underneath), h3 (in `--color-heading-alt`), paragraphs, lists, `code`, links. `.ds-prose__lead` is the opening line.
- `.ds-link`: text link; states `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--quiet`. `.ds-links` is a wrapping row.
- `.ds-button`: a gem plate, 44px high, with a lit top and a dark lip. `.ds-button--secondary` (a wooden plate); `.ds-button--danger` (wooden plate, danger ink and edge); `.ds-button--large` (56px); `.ds-button--small` (32px). States `is-hover` (brass glow from `--shadow-control-hover`), `is-active`, `is-focus`, `is-disabled`. `.ds-buttons` is a row.
- `.ds-form`: rows (`.ds-form__row`) of a `.ds-form__label` (170px; above the field on a phone) and a `.ds-form__field`. `.ds-form__input`, `.ds-form__select`, `.ds-form__textarea`; `.ds-form__input--short`; `is-error` with `.ds-form__error`; `.ds-form__hint`; `.ds-form__check` with `.ds-form__checkbox`; `.ds-form__actions`, `.ds-form__actions-end`. `.ds-form--inline` is one field and one button in a line.
- `.ds-switch`: an on and off row over a checkbox: `.ds-switch__label`, `.ds-switch__hint`, the hidden `.ds-switch__input` and the `.ds-switch__track`, which takes `--fill-accent` when on (`is-on` by class). Used for library filters and deck options.
- `.ds-table`: a data table in a flush panel body with `.ds-table__wrap`. The head row takes `--fill-bar`; `is-numeric`, `is-centre`; `.ds-table__sub` a second line; `.ds-table__who`; `is-own` marks the reader's row.
- `.ds-list`: ruled rows: `.ds-list__item` with a `.ds-list__date` (a leaf of card stock, `.ds-list__month` over the day), a `.ds-list__body` (`.ds-list__title`, `.ds-list__text`, `.ds-list__meta`) and perhaps a badge. `.ds-list--tight` is the dense form.
- `.ds-panel`: a box with a lit top and a lip: `.ds-panel__title` (may hold a badge at the right), `.ds-panel__body` (`--flush`), `.ds-panel__text`, `.ds-panel__foot`.
- `.ds-stat`: one figure, `.ds-stat__value` over `.ds-stat__label`, with a rule at its left; four sit in a `.ds-stats` row. On the shelf both are written in `--color-bar-text`.
- `.ds-grid`: four `.ds-grid__cell` across (`--two`, `--three`), tinted and edged by `--color-fill-1`; `.ds-grid__cell--2`, `--3`, `--4`. Inside: `.ds-grid__icon` (a crest on a round brass-rimmed disc), `.ds-grid__title`, `.ds-grid__text`, a badge; `.ds-grid__more` is available.
- `.ds-tabs`: wooden plates standing on a brass rule; `.ds-tabs__tab`, `is-current` in `--fill-accent`; `.ds-tabs__count` a number.
- `.ds-badge`: a label in `--fill-accent`; `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger`, `.ds-badge--quiet`, `.ds-badge--new`. `.ds-badges` is a row; `.ds-dot` with `--success`, `--warning`, `--danger` a lamp.
- `.ds-item`: a card name in its rarity colour: plain for common, `--uncommon`, `--rare`, `--epic`, `--legendary`, `--set`.
- `.ds-tooltip`: the keyword window: a trigger wrapping a `.ds-tooltip__tip` shown on hover or focus (`is-open`, `.ds-tooltip--end`), with `.ds-tooltip__name` and `.ds-tooltip__line` rows (`--bonus`, `--muted`, `--lore`). `.ds-tooltip--card` writes the window into the page.
- `.ds-progress`: a sunk track with a `.ds-progress__bar`; width by `--10` to `--100`, colour by `--health`, `--xp`, `--rep`, `--danger`; `.ds-progress__label` above. Used for the collection, win rates and a draft.
- `.ds-avatar`: a round class crest (inline SVG) on a `--color-surface-strong` disc with a brass rim; `.ds-avatar--small`, `.ds-avatar--large`; `--1` to `--4` colour it.
- `.ds-accordion`: `<details>` sections: `.ds-accordion__item`, `.ds-accordion__summary`, `.ds-accordion__meta`, `.ds-accordion__body`.
- `.ds-menu`: the drop-down under a navigation item: `.ds-menu__list` of `.ds-menu__item`, shown on hover or focus (`is-open`); `.ds-menu__caret`.
- `.ds-columns`: a flexible column beside a 360px one. `.ds-sidebar`: a box of a narrow column: `.ds-sidebar__title`, a `.ds-sidebar__list` of `.ds-sidebar__item` rows with a `.ds-sidebar__link` and `.ds-sidebar__meta`.
- `.ds-notice`: a ruled message with `.ds-notice__title` and `.ds-notice__text`; `.ds-notice--error`. `.ds-notices` stacks several.
- `.ds-pagination`: wooden plates, `.ds-pagination__link`; `is-current` in `--fill-accent`, `is-disabled` flat.
- `.ds-breadcrumb`: a small trail: `.ds-breadcrumb__link`, `.ds-breadcrumb__sep` (an angle), `.ds-breadcrumb__current`.
- `.ds-dialog`: `.ds-dialog__title` on `--fill-bar` with `.ds-dialog__close`, `.ds-dialog__body`, `.ds-dialog__text`, `.ds-dialog__actions`; `.ds-dialog__backdrop` behind it.
- `.ds-empty`: an empty slot: `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text`, one secondary button.
- `.ds-band`: the closing call on `--fill-inverse` with a brass rim: `.ds-band__title` and `.ds-band__text` on the left, one large button on the right. At most one per page.
- `.ds-footer`: the closing strip of felt in `--fill-bar-alt` under a wooden rail: `.ds-footer__inner` holds the brand, `.ds-footer__links` of `.ds-footer__link` and a `.ds-footer__note`.
- `.ds-view`: one screen of the example site; only the one named in the address shows, the home view when none is.
- `.ds-heading` (`--centre`), `.ds-rule`, `.ds-stack` (`--tight`), `.ds-split` (`--wide`), `.ds-row` (`--between`), `.ds-muted`, `.ds-icon`: helpers.

## Never

- `border-radius <= 18px`: corners are eased like sanded wood, never soft bubbles; gems, crests and coins are round.
- `border-width <= 12px`: rims are 2px and 3px; only the rail above the footer is thicker.
- `font-size >= 13px`: card text and meta are the smallest text.
- `font-size <= 70px`: the hero title is the largest text.
- `font-families <= 3`: a slab serif, a plain sans and a monospace for code.
- `line-height <= 1.6`: body text is set at 1.6.
- `letter-spacing <= 1px`: nothing is tracked out; the slab is set tight.
- `uppercase-text <= 2%`: titles and buttons are in mixed case.
- `underlined-links <= 5%`: links are told apart by colour; only a hovered link is underlined.
- `box-shadow-blur <= 60px`: shadows are lips and short falls; the longest is the dialog's.
- `animation = none`: nothing moves by itself; a card only lifts under the pointer.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new part of a page is a `.ds-board` with its head; nothing is written on the wall. A box inside a board is a `.ds-panel` or a `.ds-sidebar`; anything a player would pick up is a `.ds-card` with its five parts, and a new card picture is one fill colour and one filled crest, never a face. Raised things (plates, gems, cards) take a `--fill-*` gradient, the brass rim and `--shadow-control` or `--shadow-panel`; sunk things (wells, tracks) take `--fill-input` and `--shadow-control-pressed`; felt (`--fill-bar-alt`) is only for where cards lie. Keep the rule for text on fills and keep rarity and status colours as ink. Do not add images, thin hairline cards, capitals, a second accent or movement.
