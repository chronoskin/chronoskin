# Game fantasy official site (black iron, ember, ash)

## Summary

The long-running official site of an action role-playing game as such sites looked from about 2012 to 2018: a dark iron sheet under a wide bar with the crest hung in its middle, a painted key-art banner, league announcements as framed tiles and a column of shop boxes beside the news. It is wider and airier than the portals of the decade before, still a fixed sheet, and it earns its keep with supporter packs. Here the iron, the ember light and the key art are gradients, shadows and clipped shapes fed by tokens.

## Layout

- Fixed centred sheet: `--size-page` is 1180px (`.ds-page__wrap`), made for a 1280px screen. The page iron shows left and right of it and below the footer.
- Order of the home page: a slim strip of account links and platforms, the bar (`--size-nav`, 54px) with three links left and three right of the diamond crest, which overhangs the bar and carries the name on a plate below it; then the sheet in `--color-canvas`: the key-art banner (`.ds-hero`, `--size-banner` 520px, text centred on its dark foot), three framed tiles across, two columns (news posts, flexible, beside a `--size-aside` 320px column of boxes: deal of the day, league progress, links for new players), the closing band and the footer plate with six columns of links. Strip, bar and footer are written once and stay on every view.
- Inner pages keep the strip, the bar and the footer; the banner is left out. They start with the breadcrumb at the top left of the sheet, then `.ds-page-header` (title and one line on the left, one or two buttons on the right, a 2px iron rule under it), then the working parts: full-width blocks (the `.ds-stats` strip, `.ds-tabs` standing on the framed table, the `.ds-grid` of packs), a `.ds-split` of two columns (`--wide` is 3 to 2), or the two columns of the home page again.
- Views: `home` (banner, tiles with an item tooltip, news, shop and league boxes, band), `league` (figures, the announcement text, challenge progress, the trailer, tabs and the ladder table with class crests, pagination), `shop` (four packs as grid cells, a carousel of pack art, the accordion of pack contents, an item window), `order` (notices, the order form with its buttons, the dialog that cancels the order, the basket), `support` (the accordion of questions, the empty ticket list, server status, links).
- Spacing scale: `--space-1` 2px (label to bar), `--space-2` 4px (small gaps, cell insets), `--space-3` 8px (between buttons, row padding), `--space-4` 12px (panel padding, list gaps), `--space-5` 18px (between blocks and columns, sheet padding), `--space-6` 28px (banner and band padding, large button sides), `--space-7` 44px (below the footer, the star grid of the key art).
- Control heights: `--size-control-small` 22px, `--size-control` 30px, `--size-control-large` 46px. `--size-date` 54px is the date plate of a post, `--size-play` 60px the round play plate, `--size-brand` 74px the crest, `--size-corner` 14px a frame bracket.

## Typography and colour roles

- Two families and a monospace. `--font-heading` and `--font-ui` are a small-capital old-style serif (Alegreya SC or Cormorant SC where installed, else Hoefler Text, Baskerville, Palatino), not transformed, tracked 0.5px, for headings, navigation, buttons, labels and prices. `--font-body` is Tahoma for everything read. The references set titles in a commercial small-capitals face; these are the closest free stacks.
- Sizes: `--text-base` 13px at `--line-body` 1.55; `--text-small` 11px for meta, badges and the strip; `--text-ui` 15px; `--text-large` 18px for navigation, the banner lead and large buttons; `--text-h3` 16px, `--text-h2` 22px, `--text-h1` 30px; `--text-display` 44px with 1.5px tracking for the banner title only.
- Palette: black iron page and sheet, three grey-black surfaces, a warm iron bar with pale ember lettering (`--color-bar`), charred brown for the strip, the key art and the footer (`--color-bar-alt`), ash-grey text under bone-white headings, ember orange for second headings, links and the primary button, and a spark yellow (`--color-accent-alt`) for frame brackets and glows. The inverse block is ash.
- Surface: 2px iron frames with a dark inner line (`--shadow-panel`), corners of 2px to 4px and 8px badges, a chequer-plate page with a faint ember glow at its foot (`--fill-page`), brushed bars, buttons lit from above, and a hover that glows in the spark colour (`--shadow-control-hover`). Hover eases over 0.15s.
- The era's rule for text on fills is kept: on `--color-canvas` and the three surfaces (and `--fill-panel`) text is `--color-text`, `--color-text-muted`, `--color-heading`, `--color-heading-alt` and the link colours; on `--fill-bar` only `--color-bar-text`; on `--fill-bar-alt` only `--color-bar-alt-text`; on `--fill-inverse` only `--color-inverse-text`; on `--fill-accent` only `--color-accent-text`; on `--fill-button` `--color-button-text`, on `--fill-button-secondary` `--color-button-secondary-text`; in a form well `--color-input-text`; on `--color-notice` `--color-notice-text` and on `--color-danger-surface` `--color-danger`. No text is written on `--color-page`.
- `--color-fill-1` to `--color-fill-4`, the three status colours and `--color-accent-alt` are inks and lights: item names by rarity (`.ds-item`), class crests, status badges on a 16% tint of themselves, grid cells as the surface tinted 24% by their fill, bars of a progress track. They are never a fill under text.
- `--color-border-strong` is the iron of every frame and of the sheet's edges; `--color-border` the ordinary rule; `--color-border-muted` the rule between rows. `--fill-accent` marks the current navigation item, tab and page number, the ribbon and the default badge.
- The key art (`.ds-scene`) is filled with `--fill-bar-alt` and written in `--color-bar-alt-text`. Its ground and silhouettes are `--color-bar-alt` mixed toward `--color-shadow`; the light behind the gate is `--color-accent-alt` inside `--color-accent`; the far glow is `--color-fill-3`; the other two skies use `--color-fill-1` (a moon) and `--color-fill-2` (a rift). So every palette repaints the picture.
- One-off values built in `components.css` from tokens: the mixes named above, percentages inside `clip-path` polygons and gradients of the scene, `50%` radius for lamps and the play plate, `aspect-ratio` for pictures, a 45 degree turn for the crest and for diamonds. The month on a date plate is always uppercase.

## Components

- `.ds-page`: on `<body>`; paints the iron. `.ds-page__wrap` is the 1180px column, `.ds-page__body` the sheet, `.ds-page__inner` the stack of a view, `.ds-page__columns` the news column and the 320px column, `.ds-page__main` and `.ds-page__aside` their stacks.
- `.ds-nav`: the strip and the bar. `.ds-nav__top` holds two `.ds-nav__group`s of `.ds-nav__action` links and `.ds-nav__chip` platform labels; `.ds-nav__bar` holds `.ds-nav__links`, the brand, and `.ds-nav__links--end`. A `.ds-nav__link` is plain bar text; the current one is a plate in `--fill-accent` (`is-current`, or in the specimen one selector per view). Items with children are a `.ds-menu`.
- `.ds-brand`: the crest in the middle of the bar. `.ds-brand__mark` is a 74px tile turned to a diamond and dressed as the primary button (`--fill-button`, `--color-button-text`, iron border, `--radius-control`, `--shadow-control`) inside a second iron ring; the mark inside stays upright. `.ds-brand__name` is a small bar plate under it. Both are placeholders for the installing project's own name and logo.
- `.ds-hero`: the key-art banner: a `.ds-scene` whose picture fills the upper part, with `.ds-hero__copy` centred on the dark ground: a `.ds-ribbon`, `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__actions` (one large button, one secondary) and `.ds-hero__note`. Once, at the top of the home page.
- `.ds-scene`: the drawn picture, used for the banner, tile art, pack art, the carousel and the video still. Eight empty spans in this order: `.ds-scene__sky` (sky, the light and stars), `.ds-scene__far` (crags), `.ds-scene__ruins` (columns), `.ds-scene__gate` (the broken arch, lit on its inner edges), `.ds-scene__haze`, `.ds-scene__near` (the ground and a cloaked figure), `.ds-scene__embers`, `.ds-scene__shade` (vignette). `.ds-scene--2` is a moonlit sky, `.ds-scene--3` a rift.
- `.ds-page-header`: the head of an inner page: `.ds-page-header__title`, `.ds-page-header__text`, `.ds-page-header__actions` on the right, an iron rule below.
- `.ds-frame`: the iron-bound box: `--fill-panel` in a 2px border of `--color-border-strong` with `--shadow-panel` and a bracket in `--color-accent-alt` on each corner. Add it to a panel, a tile, a side box, the carousel. `.ds-frame--plain` leaves the brackets off (tables, the strip of figures).
- `.ds-ribbon`: a short label on `--fill-accent` with notched ends.
- `.ds-tiles`: three `.ds-tile` across, each a frame with `.ds-tile__art` (a scene), and a `.ds-tile__body` holding a badge, `.ds-tile__title`, `.ds-tile__text` and a secondary button.
- `.ds-video`: a framed `.ds-video__still` (a scene) with the round `.ds-video__play` plate and a `.ds-video__caption`.
- `.ds-prose`: running text: h1, h2 (ruled underneath), h3 (in `--color-heading-alt`), paragraphs, lists, `code`, links. `.ds-prose__lead` is the opening line.
- `.ds-link`: text link; states `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--quiet` for meta links. `.ds-links` is a wrapping row.
- `.ds-button`: a plate lit from above, 30px high. `.ds-button--secondary` is the iron plate; `.ds-button--danger` the destructive one (iron, with danger ink and edge, because the primary plate is itself ember red); `.ds-button--large` the 46px call in the heading face; `.ds-button--small` 22px. States `is-hover` (glows through `--shadow-control-hover`), `is-active`, `is-focus`, `is-disabled`. `.ds-buttons` is a row.
- `.ds-form`: rows (`.ds-form__row`) of a `.ds-form__label` (150px) and a `.ds-form__field`. `.ds-form__input`, `.ds-form__select`, `.ds-form__textarea` are sunk wells; `.ds-form__input--short`; `is-error` with `.ds-form__error`; `.ds-form__hint`; `.ds-form__check` with `.ds-form__checkbox`; `.ds-form__actions`, where `.ds-form__actions-end` pushes the destructive button right. `.ds-form--inline` is one well and one button in a line.
- `.ds-table`: a data table inside a `.ds-frame--plain`. Head cells take `--fill-bar`; rows are ruled and every second one tinted; `is-numeric`, `is-centre`; `.ds-table__who` a crest beside a name; `.ds-table__sub` a small second line; a row with `is-own` is the reader's own, tinted with `--color-accent-alt`. `.ds-table__wrap` lets a wide table scroll.
- `.ds-list`: ruled rows: `.ds-list__item` with `.ds-list__body` (`.ds-list__title`, `.ds-list__text`, `.ds-list__meta`) and perhaps a badge. `.ds-list--posts` makes each row a small slab with a `.ds-list__date` plate (`.ds-list__month` over the day); `.ds-list--tight` is the dense form.
- `.ds-panel`: a box with a `.ds-panel__title` (a bar plate), a `.ds-panel__body` (`--flush` without padding), `.ds-panel__text`, `.ds-panel__foot`. Use it with `.ds-frame`.
- `.ds-stat`: one figure: `.ds-stat__value` over `.ds-stat__label`. Three to five sit in a `.ds-stats` strip, divided by rules.
- `.ds-grid`: four `.ds-grid__cell` across (`--two`, `--three`), tinted and edged by `--color-fill-1`; `.ds-grid__cell--2`, `--3`, `--4` use the others. Inside: `.ds-grid__art` (a scene across the top) or `.ds-grid__icon`, `.ds-grid__title`, `.ds-grid__text`, a `.ds-price`, a button or `.ds-grid__more`.
- `.ds-price`: a price in the heading face; `.ds-price__was` the struck old price.
- `.ds-tabs`: plates standing on the frame that follows; `.ds-tabs__tab`, `is-current` in `--fill-accent`; `.ds-tabs__count`.
- `.ds-badge`: a small label in `--fill-accent`. `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger` (ink on a tint of itself), `.ds-badge--quiet`, `.ds-badge--new` (outlined in `--color-accent-alt`). `.ds-badges` is a row; `.ds-dot` with `--success`, `--warning`, `--danger` a status lamp.
- `.ds-item`: an item name in its rarity colour: plain, `--uncommon`, `--rare`, `--epic`, `--legendary`, `--set`.
- `.ds-tooltip`: the item window. The trigger wraps a `.ds-tooltip__tip` that shows on hover or focus (`is-open` by class, `.ds-tooltip--end` hangs it from the right). Inside: `.ds-tooltip__name`, `.ds-tooltip__line` rows, `--bonus`, `--muted`, `--lore`. `.ds-tooltip--card` writes the window into the page.
- `.ds-progress`: a sunk track with a `.ds-progress__bar`; width by `--10`, `--25`, `--40`, `--60`, `--75`, `--90`, `--100`; colour by `--health`, `--xp`, `--rep`, `--danger`. `.ds-progress__label` is the line above.
- `.ds-avatar`: a class crest (inline SVG) on a `--color-surface-strong` tile with an iron border; `.ds-avatar--small`, `.ds-avatar--large`; `--1` to `--4` colour the crest.
- `.ds-accordion`: `<details>` sections: `.ds-accordion__item`, `.ds-accordion__summary` with a small plate reading + or -, `.ds-accordion__meta`, `.ds-accordion__body`.
- `.ds-menu`: the drop-down under a navigation item: `.ds-menu__list` of `.ds-menu__item`, shown on hover or focus (`is-open` by class); `.ds-menu__caret`.
- `.ds-carousel`: one framed picture: `.ds-carousel__stage` holding a `.ds-carousel__picture` and two `.ds-carousel__arrow` plates (`--next`), and a `.ds-carousel__foot` with the caption and `.ds-carousel__dots` of diamond `.ds-carousel__dot`s, `is-current` lit.
- `.ds-switch`: defined for settings rows (`.ds-switch__input`, `.ds-switch__track`, `.ds-switch__label`, `.ds-switch__hint`); the specimen has no settings screen and does not show it.
- `.ds-sidebar`: a framed box of the right column: `.ds-sidebar__title` (a bar plate, may hold a second span at the right), `.ds-sidebar__body` or a `.ds-sidebar__list` of `.ds-sidebar__item` rows with a `.ds-sidebar__link` and a `.ds-sidebar__meta`.
- `.ds-notice`: a ruled message with `.ds-notice__title` and `.ds-notice__text`; `.ds-notice--error`. `.ds-notices` stacks several.
- `.ds-pagination`: small iron plates, `.ds-pagination__link`; `is-current` in `--fill-accent`, `is-disabled` flat.
- `.ds-breadcrumb`: a small trail of `.ds-breadcrumb__link`, `.ds-breadcrumb__sep` (a diamond) and `.ds-breadcrumb__current`.
- `.ds-dialog`: a small window: `.ds-dialog__title` bar plate with `.ds-dialog__close`, `.ds-dialog__body`, `.ds-dialog__text`, `.ds-dialog__actions`. `.ds-dialog__backdrop` is the dimmed area behind it.
- `.ds-empty`: an empty slot: `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one secondary button.
- `.ds-band`: the closing call in `--fill-inverse`: `.ds-band__title`, `.ds-band__text`, one large button. At most one per page.
- `.ds-footer`: the closing plate in `--fill-bar-alt` with a diamond on its top edge: `.ds-footer__columns` of `.ds-footer__title` and `.ds-footer__links` of `.ds-footer__link`, then `.ds-footer__note`.
- `.ds-view`: one screen of the example site; only the one named in the address shows, the home view when none is.
- `.ds-heading`: a section title followed by a fading rule (`--centre` for both sides); `.ds-rule` a rule with a diamond; `.ds-stack` (`--tight`), `.ds-split` (`--wide`), `.ds-row` (`--between`), `.ds-muted`, `.ds-icon`: helpers.

## Never

- `border-radius <= 8px`: plates and frames are nearly square; only badges reach 8px, and lamps and the play plate are round.
- `border-width <= 2px`: frames are 2px iron, every rule 1px.
- `font-size <= 44px`: the banner title is the largest text.
- `font-size >= 11px`: meta text is 11px and nothing is smaller.
- `font-families <= 3`: a small-capital serif, Tahoma and a monospace.
- `line-height <= 1.55`: body text is set at 1.55.
- `letter-spacing <= 1.5px`: headings are tracked half a pixel, the banner title 1.5px.
- `uppercase-text <= 5%`: capitals come from the small-capital face, not from a transform.
- `box-shadow-blur <= 28px`: shadows are lips and short glows.
- `animation = none`: nothing on the site moves by itself.
- `content-width <= 1180px`: the page is a fixed sheet.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A box that stands on the sheet is a `.ds-frame` (or `.ds-frame--plain` inside another); a raised thing takes a `--fill-*` gradient, an edge mixed toward `--color-shadow` and `--shadow-control`; a sunk thing takes `--fill-input` and `--shadow-control-pressed`. A new picture is a `.ds-scene` with its eight spans and, if it needs another light, a new sky variant built from one fill colour. Keep the rule for text on fills, and keep rarity and status colours as ink. Do not add images, heavy carved frames, a second accent, or text on the page iron.
