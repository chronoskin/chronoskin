# Modern product order tracking

## Summary

This is the order and delivery tracking of a consumer brand between 2020 and 2026, here a maker of electric bikes: a dark forest-green app in which the state of the order is the headline, set very large in a geometric sans, over a bar of five stages, a map tile and a summary card that stays in place while the events scroll. One acid lime is the only loud colour: the primary button, the progress bar, the block that holds the drawn bike and a full-width band of figures. The screens are reached from a dock that floats at the lower edge of the window, as in the brand's phone app; it was the common look of the order, delivery and account pages of vehicle, device and delivery brands in the first half of the 2020s.

## Layout

- The page is fluid and designed at 1440px. Content sits in `.ds-wrap`, centred, at most `--size-page` (1280px) wide with `--space-5` side padding; sections pad themselves with `--space-9` (96px), `--tight` with `--space-8`.
- Navigation is `.ds-nav` in two places. A plain top row, `--size-nav` (64px) tall and not fixed, has the brand at the left and, at the right, a `.ds-menu` that opens a wide panel of the customer's orders and the signed-in person's `.ds-avatar`. The screens themselves are `.ds-nav__links`, a dock: a small translucent, blurred bar of four icon links that is `position: fixed`, centred, `--space-4` above the lower edge of the window; the current screen is filled like a primary button. The footer keeps `--space-9` free under itself for the dock. Do not give `.ds-nav` a `backdrop-filter` or a transform, or the dock stops being fixed to the window.
- The home view stacks, inside `.ds-wrap`: `.ds-hero`, on the page colour, with the order number and a badge, the state as a display headline whose second half is greyed, one line of text and two buttons, and at the right, on the baseline, the `.ds-rider` card; `.ds-track`, a rounded block with a large `.ds-progress` and five named stages; and `.ds-board`, two columns: at the left, `--size-side` (400px) wide and `position: sticky`, the order summary panel with the bike on an accent block, at the right the `.ds-map` tile and a panel with the timeline. Then, full bleed: `.ds-band`, an inset rounded block of the accent with the `.ds-stat` figures; a section with the `.ds-grid` of four tiles; and a `.ds-tone--alt` section with the `.ds-accordion` beside a `.ds-carousel` of tips.
- Inner pages have no hero. Inside `.ds-wrap` they open with `.ds-page-header` (breadcrumb, a large title, one line of text, one button at the right), then the working components. Tabs are a pill switch at the left of `.ds-tabsrow` with the badge key at its right. Content is full width (the table of orders), `.ds-cols` (7 to 5), `.ds-cols--even`, or `.ds-cols--side` with the `.ds-sidebar` at the left.
- Views. The specimen is the order area of a bike maker. `home` tracks one order. `orders` has the switch and badge key, the table of orders with its pagination, an empty returns box and the service plan with a progress bar. `delivery` is the change-delivery form with its notices and buttons, the open cancel dialog and three switches for messages. `help` has the topics and a person at the left, the text on the handover with a row of links, and more questions as an accordion.
- Below 1000px the hero, the board and all columns stack, the summary is no longer sticky, tiles go two by two and the menu panel is not shown. Below 640px the dock's links put the icon over the word and the dock sits `--space-2` above the lower edge, the first name beside the `.ds-avatar` of `.ds-nav__me` is hidden, the stages lose their dates, the band and the largest blocks take `--radius-panel` and a wide table scrolls inside `.ds-table__wrap`. Below 640px a `.ds-tooltip__tip` takes the width of its paragraph, so it never leaves the window. The page never scrolls sideways.
- Spacing scale: `--space-1` 4px (hairline gaps, dock padding), `--space-2` 8px (icon to text, progress bar height), `--space-3` 12px (between buttons, tile gap), `--space-4` 16px (form rows, list rows, dock to window edge), `--space-5` 20px (panel and tile padding, page gutter, column gap), `--space-6` 28px (track padding, hero actions), `--space-7` 40px (hero to track, section head to content, timeline indent), `--space-8` 64px (top of the hero, band padding, tight sections), `--space-9` 96px (sections, room for the dock).
- Other sizes: `--size-mark` 28px, `--size-icon` 44px, `--size-control` 40px, `--size-side` 400px, `--size-map` 340px, `--size-art` 240px (least height of the bike's block), `--size-tile` 220px, `--size-menu` 620px, `--size-avatar` 36px, `--size-tip` 220px, `--size-dialog` 440px.
- Motion, in CSS only. An element with `.ds-reveal` fades in and rises as it enters the window (`animation-timeline: view()`; `--2` and `--3` start later along a row). Buttons rise a little on hover and shrink to 97% when pressed, tiles rise by `--space-1`, the menu panel and tooltips fade and slide in, all through `--transition`. Animations are declared inside `@media (prefers-reduced-motion: no-preference)` and state only their start, so nothing is hidden where they do not run.

## Typography and colour roles

These are the conventions of the era; every layout and every token set keeps them.

- A block stands on one of four grounds, and each ground has its own text token. The page (`--fill-page`), panels (`--fill-panel`), `--color-surface-alt`, `--color-surface-strong` and the four fills carry `--color-text`, `--color-text-muted`, `--color-heading` and `--color-link`. An inverse block (`--fill-inverse`, class `.ds-tone--inverse`) carries only `--color-inverse-text`. An accent block (`--fill-accent`, class `.ds-tone--accent`, the band and the bike's block) carries only `--color-accent-text`. The navigation and the dock (`--fill-bar`) carry `--color-bar-text`; the footer (`--fill-bar-alt`) carries `--color-bar-alt-text`.
- A palette may be light or dark; nothing assumes that the page is dark.
- `--color-fill-1` to `--color-fill-4` are the grounds of tiles, slides, the map and avatars and carry `--color-text` and `--color-heading`.
- Buttons by ground. On the page, panels and fills: `.ds-button` and `.ds-button--secondary`. On an inverse block only `.ds-button--inverse` and `.ds-button--ghost`. On an accent block only `.ds-button--onaccent` and `.ds-button--ghost`.
- The product is drawn as inline SVG in `currentColor` on an accent block, so it takes `--color-accent-text`. The map's roads and route are SVG in `--color-heading` on `--color-fill-1`; its pins and note are small panels.
- `--fill-bar` is translucent and the dock blurs what scrolls under it by `--backdrop-blur`.
- `--fill-button` with `--color-button-text` is also the current screen in the dock, the filled part of a progress bar, a switch that is on, the open accordion's sign, the current page number and the current event of the timeline.
- `--color-heading-alt` is the kicker above a title, the sidebar title and h3 in prose.
- `--fill-accent` with `--color-accent-text` is the default badge. `.ds-badge--alt` and the status badges are tints mixed into `--color-surface`.
- `--color-notice` with `--color-notice-text` is the information notice; the error notice is `--color-danger` on `--color-danger-surface`.
- Borders are hairlines: `--border-width` in `--color-border-muted` around panels, `--color-border` between accordion rows, above a total and along the timeline, `--color-input-border` around fields. `--border-width-strong` frames the empty state and the dots of the timeline.
- Radii: `--radius-control` for buttons, fields, tabs, the dock's links and the current row of a sidebar; `--radius-panel` for panels, tiles, the dialog and the empty state; `--radius-page` for the track, the map, the band and slides; `--radius-pill` for badges, bars, switches and map pins.
- Shadows: `--shadow-panel` on panels, `--shadow-control` on buttons, `--shadow-control-hover` on a hovered button or tile, `--shadow-dialog` on the dialog, the dock, the menu panel and what floats on the map.
- One-off values made with `calc()` and `color-mix()`: washes of the bar text at 10 to 16% for the dock's edge and hovered links, tinted badges, the error notice's edge, the footer rule, the ring around the current timeline dot (30% of `--color-button`). Opacity 0.16 dims the map's roads, 0.25 to 0.8 dots and secondary text on coloured blocks.
- Here: the page is a deep forest green `#0b1a10`, panels one step lighter, text a pale sage; the accent and the primary button are one acid lime `#b7f34a` carrying the page's green. The four fills are deep green, moss, olive and teal. The inverse block is off-white.
- Two families: a geometric sans for headings, figures, buttons and navigation (`--font-heading`, `--font-ui`: Space Grotesk, Futura; the references used commercial grotesques) and a neutral sans for text (`--font-body`: Geist, Roboto, Helvetica Neue); `--font-mono` for order numbers and code. `--text-base` 16px at `--line-body` 1.5; `--text-ui` 15px; `--text-small` 13px; `--text-large` 20px; `--text-display` 100px (the state of the order and the figures, scaled down with `clamp()` and `vw`); `--text-h1` 68px (section and page titles); `--text-h2` 36px; `--text-h3` 22px.
- Headlines are medium weight and tight: `--weight-display` and `--weight-heading` 500 with `--display-tracking` -3px and `--heading-tracking` -1px. `--weight-ui` is 500, `--weight-bold` 700. Nothing is uppercase; links are not underlined until hovered.
- Surface: small 8px corners on controls against 20px panels and 32px blocks, pill badges, hairline borders, no resting shadows; a hovered control gets a ring of its own colour, floating things a hairline ring and a deep shadow. `--transition` is 0.22s; `--backdrop-blur` 14px.

## Components

- `.ds-page`: on `<body>`. Sets font, text colour and `--fill-page`.
- `.ds-tone--inverse`, `.ds-tone--accent`, `.ds-tone--alt`: put a block on the inverse fill, the accent fill or `--color-surface-alt` and set its text colour.
- `.ds-brand`: the site's mark and name, first in the top row and again in the footer. `.ds-brand__mark` is a small tile dressed like the primary button holding the mark as inline SVG; `.ds-brand__name` is the name. Name and mark are placeholders for the installing project's own name and logo.
- `.ds-wrap`: the centred content column. `.ds-section`: a full-width block (`--tight`, `--flush`) with a `.ds-section__head` of `.ds-section__kicker`, `.ds-section__title` and `.ds-section__lead`; `--left` aligns it left.
- `.ds-nav`: the top row and the dock. `.ds-nav__inner` lays out the brand and `.ds-nav__actions` (the menu and `.ds-nav__me`, avatar and first name); `.ds-nav__links` is the dock of `.ds-nav__link` items, each a small SVG icon and a word, the current one filled (`is-current`).
- `.ds-menu`: a top-row item with a drop-down. `.ds-menu__toggle` is the item, with a small caret; `.ds-menu__panel` is a wide panel under it, shown on hover, on focus or with `is-open`; `--end` aligns it to the item's right edge. Here it holds two `.ds-menu__order` cards (a badge, the order's name, a line and a progress bar) and a `.ds-menu__foot` of links; a plain menu uses `.ds-menu__title`, `.ds-menu__list` and `.ds-menu__link`. Do not show it open over other content.
- `.ds-hero`: the head of the home view: `.ds-hero__kicker` (order number and badge), `.ds-hero__title` (an `<em>` greys its second half), `.ds-hero__lead`, `.ds-hero__actions`, and the `.ds-rider` card at the right. One per site.
- `.ds-rider`: a small panel with a `.ds-person`, a `.ds-rider__text` and optionally one small button.
- `.ds-track`: the stages of the order: a `.ds-progress__label`, a `.ds-progress--large` and `.ds-track__steps` of five `.ds-track__step` (`is-done`, `is-current`), each a word over a small date.
- `.ds-progress`: a bar; `.ds-progress__bar` with `--20`, `--40`, `--60`, `--70`, `--80`, `--100` sets how far; `--large` is thicker. `.ds-progress__label` above it has a bold name and a quiet remark.
- `.ds-board`: the two columns of the home view; `.ds-board__side` is the sticky one.
- `.ds-panel`: the rounded card. `.ds-panel__head` holds the `.ds-panel__title` and an optional badge or link; `.ds-panel__body` is padded; `.ds-panel__text` is a paragraph. `.ds-summary` is the order's panel: `.ds-summary__art` (the product in SVG on the accent), a head, `.ds-lines` and a `.ds-table__foot` with the address.
- `.ds-lines`: priced rows: `.ds-lines__row` with a name over a small line and a `.ds-lines__sum`; `--total` is the last, larger row above a rule.
- `.ds-map`: a tile on `--color-fill-1` with two SVG layers, `.ds-map__roads` and `.ds-map__route`, `.ds-map__pin` labels (`--rider`, `--home`) and a `.ds-map__note` with a large first line.
- `.ds-list`: rows parted by hairlines. `.ds-list__item` holds a `.ds-list__title`, an optional `.ds-list__text` and `.ds-list__meta`; `.ds-list__row` puts a `.ds-list__grow` block and a badge or switch on one line. `.ds-list--timeline` turns it into a line of events with dots (`is-done`, `is-current`), newest first.
- `.ds-band`: an inset rounded block for the figures, used with `.ds-tone--accent`; `.ds-band__head` is its small heading.
- `.ds-stat`: the row of three figures; each `.ds-stat__item` is a `.ds-stat__figure` in display size over a `.ds-stat__label`.
- `.ds-grid`: four tiles on the four fills. Each `.ds-grid__cell` is a link with a `.ds-grid__num` in monospace at the top and a `.ds-grid__title` and `.ds-grid__text` at the bottom; `is-done` strikes the title through.
- `.ds-accordion`: questions that open and close. Each `.ds-accordion__item` is a `<details>` with a `.ds-accordion__head` summary (a plus that becomes a filled minus when open) and a `.ds-accordion__body`.
- `.ds-carousel`: slides of which one shows. `.ds-carousel__track` scrolls sideways with scroll snapping and holds `.ds-carousel__slide` items on the fills; `.ds-carousel__bar` under it has `.ds-carousel__dots` (`.ds-carousel__dot`, `is-current`) and `.ds-carousel__hint`, a small line that says to swipe. Here a slide holds a `.ds-tip`: a badge, a `.ds-tip__title`, a `.ds-tip__text` and at most one button.
- `.ds-tooltip`: a dotted-underlined trigger holding a `.ds-tooltip__tip`, a small inverse box above it that shows on hover, on focus or with `is-open`.
- `.ds-switch`: an on and off control drawn over `<input type="checkbox">`; on when checked or `is-on`.
- `.ds-avatar`: initials in a circle on a fill (`--2`, `--3`, `--large`). `.ds-person` puts it beside a `.ds-person__name` and `.ds-person__meta`.
- `.ds-page-header`: head of an inner page, used instead of `.ds-hero`: `.ds-page-header__main` holds the breadcrumb, the `.ds-page-header__title` and one line of `.ds-page-header__text`; `.ds-page-header__actions` sits at the right.
- `.ds-cols`: main and side column, 7 to 5; `--side` 1 to 3 for a sidebar; `--even` two equal columns; `--fill` stretches both to one height. `.ds-stack` stacks blocks one `--space-5` apart.
- `.ds-prose`: running text: h1, h2, h3, p, lists, `code`, `strong`. `.ds-prose__lead` for the first paragraph.
- `.ds-link`: text link. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `--quiet` for meta; `--more` appends a small arrow. `.ds-linkrow` lays links in a row.
- `.ds-button`: the button, in `--font-ui`. Variants: `--secondary`, `--danger` (destructive actions only), `--inverse`, `--onaccent`, `--ghost`, `--small`, `--wide`. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttonrow` spaces buttons.
- `.ds-form`: vertical form with labels above the fields. `.ds-form__row` puts two fields side by side; `.ds-form__field` wraps `.ds-form__label` and one of `.ds-form__input`, `.ds-form__select`, `.ds-form__textarea`. Error: `is-error` on the input plus `.ds-form__error`. `.ds-form__hint`, `.ds-form__check` with `.ds-form__checkbox`, `.ds-form__actions`.
- `.ds-table`: data table without fills: a small quiet head, hairlines between rows; `is-hover` tints a row. `.ds-table__num` right-aligns numbers; `.ds-table__title` is a row's link; `.ds-table__code` sets a code in the monospace font. Wrap it in `.ds-table__wrap`; `.ds-table__foot` is a last strip for a count and the pagination.
- `.ds-tabs`: a pill switch: `.ds-tabs__tab` on a block of `--color-surface-strong`, the current one (`is-current`) raised on `--fill-panel`. `.ds-tabsrow` places it above what it changes.
- `.ds-badge`: small label on `--fill-accent`. `--alt` tinted, `--quiet` grey. Status variants: `--success`, `--warning`, `--danger`. `.ds-badgerow` for several.
- `.ds-sidebar`: a block of links: `.ds-sidebar__title` and a `.ds-sidebar__list` of `.ds-sidebar__item` with an optional `.ds-sidebar__count`; the current row (`is-current`) is filled.
- `.ds-notice`: tinted rounded message; `--error` in the danger colours. `.ds-notice__title` for the bold lead-in. `.ds-noticerow` stacks notices.
- `.ds-pagination`: `.ds-pagination__link` blocks; `is-current` is filled with `--fill-button`, `is-disabled` is grey.
- `.ds-breadcrumb`: a small trail of `.ds-breadcrumb__item` parted by small arrows; the last is `.ds-breadcrumb__current`.
- `.ds-dialog`: a dimmed block (`--color-overlay`) centring `.ds-dialog__box`: `.ds-dialog__head` (`.ds-dialog__title`, a round `.ds-dialog__close`), `.ds-dialog__body` and `.ds-dialog__actions` with the main action first.
- `.ds-empty`: a centred message in a frame of `--border-width-strong`: `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one small secondary button.
- `.ds-plan__row`: a last line in a panel with a remark at the left and one small button at the right.
- `.ds-reveal`: the scroll animation described under Layout; `--2`, `--3`, `--zoom`.
- `.ds-footer`: the closing band on `--fill-bar-alt`: `.ds-footer__cols` with the brand and a `.ds-footer__about` line, then four columns of `.ds-footer__title` and `.ds-footer__list` of `.ds-footer__link`; `.ds-footer__legal` under a rule.

## Never

- `text-shadow = none`: text is flat on every ground.
- `border-width <= 2px`: borders are hairlines; 2px frames the empty state and the dots of the timeline.
- `border-radius <= 32px`: controls are 8px, panels 20px and the largest blocks 32px; only pills and circles are rounder.
- `box-shadow-blur <= 64px`: only floating things cast a shadow, and it ends at 64px.
- `font-weight <= 700`: headlines are medium; 700 is used only for bold text.
- `font-size >= 13px`: meta, badges and the footer are the smallest text.
- `font-size <= 100px`: the state of the order and the figures are the largest text.
- `font-families <= 3`: a geometric sans, a text sans and monospace for codes.
- `line-height <= 1.5`: body text is set at 1.5.
- `uppercase-text <= 0%`: navigation, buttons and labels are in sentence case.
- `underlined-links <= 5%`: links are told apart by colour; only a hovered link is underlined.
- `gradient-fills <= 0%`: every fill is flat.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. First decide which ground it stands on (page, panel or fill; inverse; accent) and take the text token and the buttons of that ground; do not put a primary button on an inverse or accent block, or `--color-inverse-text` on a fill. A new screen is one more link in the dock (keep it to five) and opens with a `.ds-page-header`; anything that holds data is a `.ds-panel`; anything that shows how far along something is uses `.ds-progress`, with named stages under it when there are few. Keep the accent rare: one accent block per screen at most, besides buttons and bars. Things that float (a menu, a tooltip, a pin) take `--shadow-dialog` and stay hidden until hover or focus unless they are part of a picture. Keep corners, shadows and borders on their tokens, spacing on the `--space-*` scale, type sizes on the `--text-*` tokens (large titles through `clamp()` between two of them), and note any one-off `color-mix()` shade next to its rule.
