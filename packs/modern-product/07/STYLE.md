# Modern product support hub

## Summary

This is the support and help site of a big consumer brand between 2020 and 2026, here a maker of home audio: a full-bleed block of one hot pink with a very large, light serif question and one big search field, a row of rounded product tiles that overlaps its lower edge, and then sections that alternate blush white, pale pink and a deep wine. The bar across the top carries the brand in its middle and a wide drop-down of products, guides are read one step at a time beside a list of steps, and the products themselves are drawn in CSS. It was the common look of the help centres of device, audio and finance brands in the first half of the 2020s.

## Layout

- The page is fluid and designed at 1440px. Content sits in `.ds-wrap`, centred, at most `--size-page` (1240px) wide with `--space-5` side padding; sections are full bleed and pad themselves with `--space-9` (112px), `--tight` with `--space-8`.
- Navigation is `.ds-nav`, a bar `--size-nav` (68px) tall that stays at the top of the window, translucent, with a hairline under it. It has three columns: the sections at the left (one of them a `.ds-menu` that opens a wide panel of products), the brand in the middle, a contact link and the signed-in person's `.ds-avatar` at the right.
- The home view stacks: `.ds-hero`, the full-bleed accent block, flush under the bar and rounded by `--radius-page` only at its lower corners, everything centred, with the `.ds-search` field and a row of small ghost buttons; `.ds-picker`, which pulls the `.ds-grid` of four product tiles up over the hero's lower edge; a section with the `.ds-carousel` of quick fixes; a `.ds-tone--alt` section with the `.ds-accordion` of questions beside two panels (a list of guides and the service status with a `.ds-progress`); a full-bleed inverse section of `.ds-stat` figures; and a section of three `.ds-way` cards with people. The footer is a dark band with rounded upper corners.
- Inner pages have no hero. Inside `.ds-wrap` they open with `.ds-page-header` (breadcrumb, a large title, one line of text, at most two buttons at the right) on the page colour, then the working components. Tabs are a row of underlined words in `.ds-tabsrow`, with the badge key at its right. Content is `.ds-cols` (7 to 5), `.ds-cols--even`, or `.ds-cols--side` with the `.ds-sidebar` at the left in a `.ds-side` column that stays in place while the text scrolls.
- Views. The specimen is the support site of a maker of speakers. `home` is the hub. `guide` is one step of a guide: the steps and an empty saved-guides box at the left, the step with its progress bar, text, notice and pagination, a "did this work" strip and related links. `repairs` has the tabs and badge key, the table of registered products, a repair in progress with its bar and stages, the open cancel dialog and three switches for updates. `contact` is the message form with its notices and buttons, the people who answer, questions as an accordion and a call-back block.
- Below 1000px the bar becomes two rows (brand and account, then one strip that scrolls sideways from edge to edge and fades at the right: the sections, then the first link of each group of the product menu, whose panel is not shown), tiles go two by two, the contact cards and all columns stack. Below 640px the hero and the largest blocks take `--radius-panel`, the slide of the carousel stacks, tabs scroll sideways, the first name beside the `.ds-avatar` of `.ds-nav__me` is hidden and a wide table scrolls inside `.ds-table__wrap`. Below 640px a `.ds-tooltip__tip` takes the width of its paragraph, so it never leaves the window. The page never scrolls sideways.
- Spacing scale: `--space-1` 4px (hairline gaps, pagination), `--space-2` 8px (icon to text, progress bar height), `--space-3` 12px (between buttons, list details), `--space-4` 16px (form rows, list rows, tile gap), `--space-5` 24px (panel and tile padding, page gutter), `--space-6` 32px (hero to search, contact card padding), `--space-7` 48px (section head to content, slide padding), `--space-8` 72px (tight sections, top of the hero), `--space-9` 112px (sections, the overlap of the tiles).
- Other sizes: `--size-mark` 30px, `--size-icon` 52px, `--size-control` 42px, `--size-search` 720px, `--size-tile` 280px, `--size-device` 112px (height of a drawn product), `--size-menu` 660px, `--size-avatar` 40px, `--size-tip` 240px, `--size-dialog` 480px.
- Motion, in CSS only. An element with `.ds-reveal` fades in and rises by `--space-7` as it enters the window (`animation-timeline: view()`; `--2` and `--3` start later along a row). Buttons rise a little on hover and shrink to 97% when pressed, tiles rise by `--space-2`, the menu panel and tooltips fade and slide in, all through `--transition`. Animations are declared inside `@media (prefers-reduced-motion: no-preference)` and state only their start, so nothing is hidden where they do not run.

## Typography and colour roles

These are the conventions of the era; every layout and every token set keeps them.

- A block stands on one of four grounds, and each ground has its own text token. The page (`--fill-page`), panels (`--fill-panel`), `--color-surface-alt`, `--color-surface-strong` and the four fills carry `--color-text`, `--color-text-muted`, `--color-heading` and `--color-link`. An inverse block (`--fill-inverse`, class `.ds-tone--inverse`) carries only `--color-inverse-text`. An accent block (`--fill-accent`, class `.ds-tone--accent`, and the hero) carries only `--color-accent-text`. The navigation (`--fill-bar`) carries `--color-bar-text`; the footer (`--fill-bar-alt`) carries `--color-bar-alt-text`.
- A palette may be light or dark; nothing assumes that the page is light.
- `--color-fill-1` to `--color-fill-4` are the grounds of tiles, slides and avatars and carry `--color-text` and `--color-heading`.
- Buttons by ground. On the page, panels and fills: `.ds-button` and `.ds-button--secondary`. On an inverse block only `.ds-button--inverse` and `.ds-button--ghost`. On an accent block only `.ds-button--onaccent` and `.ds-button--ghost`. The search field is a form control (`--fill-input`, `--color-input-text`), so the primary button inside it is on its own ground.
- `--color-inverse` is the body of a drawn product (`.ds-device`), with details in `--color-inverse-text` and one light in `--color-accent`; a product is drawn on a fill or the page, never on an inverse block. `--color-inverse` with `--color-inverse-text` is also the tooltip.
- `--fill-bar` is translucent and the bar blurs what scrolls under it by `--backdrop-blur`.
- `--color-heading-alt` is the kicker above a title, the sidebar title and h3 in prose.
- `--fill-accent` with `--color-accent-text` is the default badge; `--color-accent` is the line under the current tab. `.ds-badge--alt` and the status badges are tints mixed into `--color-surface`.
- `--fill-button` is also the filled part of a progress bar, a switch that is on, the open accordion's sign and the current page number.
- `--color-notice` with `--color-notice-text` is the information notice; the error notice is `--color-danger` on `--color-danger-surface`.
- Borders are hairlines: `--border-width` in `--color-border-muted` around panels and tiles, `--color-border` between accordion rows and under tabs, `--color-input-border` around fields. `--border-width-strong` frames the empty state, underlines the current tab and draws the details of a product.
- Radii: `--radius-control` for buttons, fields, the search and the current row of a sidebar; `--radius-panel` for panels, the dialog and the empty state (half of it for notices); `--radius-page` for the hero's lower corners, tiles, slides, contact cards and the footer's upper corners; `--radius-pill` for badges, bars and switches.
- Shadows: `--shadow-panel` on panels and tiles, `--shadow-control` on buttons, `--shadow-control-hover` on a hovered button or tile, `--shadow-dialog` on the dialog, the menu panel and the search field.
- One-off values made with `calc()` and `color-mix()`: washes of the bar text at 8 to 13% for hovered and current navigation, tinted badges, the error notice's edge, the footer rule, the footer mark's outline (35% of the footer text). Opacity 0.25 to 0.8 dims dots and secondary text on coloured blocks.
- Here: the page is a blush white `#fff6f9`, panels are white, the accent is one hot pink `#ff4d94` carrying a near-black wine, and buttons, the inverse block and the footer are that deep wine `#24091a`. The four fills are pastel pink, peach, lilac and mint.
- Two families: a serif for headings, figures and the brand (`--font-heading`: Fraunces, Source Serif 4, Iowan Old Style; the references used commercial display serifs) and a plain sans for text and controls (`--font-body`, `--font-ui`: Public Sans, Helvetica Neue); `--font-mono` for code and serial numbers. `--text-base` 17px at `--line-body` 1.55; `--text-ui` 15px; `--text-small` 13px; `--text-large` 20px (leads, list and accordion titles); `--text-display` 96px (hero title and figures, scaled down with `clamp()` and `vw`); `--text-h1` 60px (section and page titles); `--text-h2` 38px; `--text-h3` 24px.
- Headlines are large, light and tight: `--weight-display` and `--weight-heading` 400 with `--display-tracking` -3px and `--heading-tracking` -0.6px. `--weight-ui` and `--weight-bold` are 600. Nothing is uppercase; links are not underlined until hovered.
- Surface: pills for controls and badges, `--radius-panel` 28px, `--radius-page` 48px, hairline borders, no resting shadows, a soft shadow under hovered controls and floating things. `--transition` is 0.28s; `--backdrop-blur` 10px.

## Components

- `.ds-page`: on `<body>`. Sets font, text colour and `--fill-page`.
- `.ds-tone--inverse`, `.ds-tone--accent`, `.ds-tone--alt`: put a block on the inverse fill, the accent fill or `--color-surface-alt` and set its text colour.
- `.ds-brand`: the site's mark and name, in the middle of the navigation and again in the footer. `.ds-brand__mark` is a small tile dressed like the primary button holding the mark as inline SVG; `.ds-brand__name` is the name. Name and mark are placeholders for the installing project's own name and logo.
- `.ds-wrap`: the centred content column. `.ds-section`: a full-width block (`--tight`, `--flush`) with a `.ds-section__head` of `.ds-section__kicker`, `.ds-section__title` and `.ds-section__lead`; `--left` aligns it left.
- `.ds-nav`: the bar. `.ds-nav__inner` lays out `.ds-nav__links` of `.ds-nav__link` (the current one on a wash, `is-current`), the brand, and `.ds-nav__actions` with a link and `.ds-nav__me` (avatar and first name).
- `.ds-menu`: a navigation item with a drop-down. `.ds-menu__toggle` is the item, with a small caret; `.ds-menu__panel` is a wide panel of three columns under it, shown on hover, on focus or with `is-open`; each column has a `.ds-menu__title` and a `.ds-menu__list` of `.ds-menu__link` (a name over a small line). `--end` aligns the panel to the item's right edge. Do not show it open over other content.
- `.ds-hero`: the accent block of the home view: `.ds-hero__kicker`, `.ds-hero__title`, `.ds-hero__lead`, the `.ds-search` and `.ds-hero__actions`, a row of `--ghost --small` buttons for popular searches. One per site.
- `.ds-search`: one large field (`.ds-search__input`) with a primary button inside its right end.
- `.ds-picker`: wraps the grid that overlaps the hero.
- `.ds-grid`: four rounded tiles on the four fills. Each `.ds-grid__cell` is a link with a `.ds-grid__art` holding a drawn product, a `.ds-grid__title` and a `.ds-grid__text`; it rises on hover.
- `.ds-device`: a product drawn with boxes: `--speaker`, `--bar`, `--phones`, `--app`; `--large` for a slide.
- `.ds-carousel`: slides of which one shows. `.ds-carousel__track` scrolls sideways with scroll snapping and holds `.ds-carousel__slide` items on the fills; `.ds-carousel__bar` under it has `.ds-carousel__dots` (`.ds-carousel__dot`, `is-current`) and `.ds-carousel__hint`, a small line that says to swipe. Here a slide holds a `.ds-fix`: a badge, `.ds-fix__title`, `.ds-fix__text`, buttons, and a `.ds-fix__art` with the product.
- `.ds-accordion`: questions that open and close. Each `.ds-accordion__item` is a `<details>` with a `.ds-accordion__head` summary (a plus that becomes a filled minus when open) and a `.ds-accordion__body`.
- `.ds-tooltip`: a dotted-underlined trigger holding a `.ds-tooltip__tip`, a small inverse box above it that shows on hover, on focus or with `is-open`.
- `.ds-switch`: an on and off control drawn over `<input type="checkbox">`; on when checked or `is-on`.
- `.ds-progress`: a bar; `.ds-progress__bar` with `--20`, `--40`, `--60`, `--70`, `--80`, `--100` sets how far; `--large` is thicker. `.ds-progress__label` above it has a bold name and a quiet remark. `.ds-steps` under it names four stages (`.ds-steps__step`, `is-done`, `is-current`).
- `.ds-avatar`: initials in a circle on a fill (`--2`, `--3`, `--large`). `.ds-person` puts it beside a `.ds-person__name` and `.ds-person__meta`.
- `.ds-status`: rows of a service and its badge (`.ds-status__row`).
- `.ds-stat`: the row of three figures; each `.ds-stat__item` is a `.ds-stat__figure` in display size over a `.ds-stat__label`.
- `.ds-ways`: three `.ds-panel.ds-way` cards, each a person, a `.ds-way__title`, a `.ds-way__text` and one button; `.ds-way--inverse` stands on the inverse fill and takes the `--inverse` button.
- `.ds-page-header`: head of an inner page, used instead of `.ds-hero`: `.ds-page-header__main` holds the breadcrumb, the `.ds-page-header__title` and one line of `.ds-page-header__text`; `.ds-page-header__actions` sits at the right.
- `.ds-cols`: main and side column, 7 to 5; `--side` 1 to 3 for a sidebar; `--even` two equal columns; `--fill` stretches both to one height. `.ds-stack` stacks blocks one `--space-5` apart; `.ds-side` keeps a column in place under the bar.
- `.ds-prose`: running text: h1, h2, h3, p, lists, `code`, `strong`. `.ds-prose__lead` for the first paragraph.
- `.ds-link`: text link. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `--quiet` for meta; `--more` appends a small arrow. `.ds-linkrow` lays links in a row.
- `.ds-button`: the pill button. Variants: `--secondary`, `--danger` (destructive actions only), `--inverse`, `--onaccent`, `--ghost`, `--small`, `--wide`. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttonrow` spaces buttons.
- `.ds-form`: vertical form with labels above the fields. `.ds-form__row` puts two fields side by side; `.ds-form__field` wraps `.ds-form__label` and one of `.ds-form__input`, `.ds-form__select`, `.ds-form__textarea`. Error: `is-error` on the input plus `.ds-form__error`. `.ds-form__hint`, `.ds-form__check` with `.ds-form__checkbox`, `.ds-form__actions`.
- `.ds-table`: data table without fills: a small quiet head, hairlines between rows; `is-hover` tints a row. `.ds-table__num` right-aligns numbers; `.ds-table__title` is a row's link; `.ds-table__code` sets a serial number in the monospace font. Wrap it in `.ds-table__wrap`; `.ds-table__foot` is a last strip for a count and a link or the pagination.
- `.ds-list`: rows parted by hairlines. `.ds-list__item` holds a `.ds-list__title`, an optional `.ds-list__text` and `.ds-list__meta`; `.ds-list__row` puts a `.ds-list__grow` block and a badge or switch on one line.
- `.ds-panel`: the rounded card. `.ds-panel__head` holds the `.ds-panel__title` and an optional badge or link; `.ds-panel__body` is padded; `.ds-panel__text` is a paragraph.
- `.ds-tabs`: a row of words on a hairline; the current `.ds-tabs__tab` (`is-current`) is underlined in the accent. `.ds-tabsrow` places it above what it changes.
- `.ds-badge`: small label on `--fill-accent`. `--alt` tinted, `--quiet` grey. Status variants: `--success`, `--warning`, `--danger`. `.ds-badgerow` for several.
- `.ds-sidebar`: a block of links: `.ds-sidebar__title` and a `.ds-sidebar__list` of `.ds-sidebar__item` with an optional `.ds-sidebar__count`; the current row (`is-current`) is filled.
- `.ds-notice`: tinted rounded message; `--error` in the danger colours. `.ds-notice__title` for the bold lead-in. `.ds-noticerow` stacks notices.
- `.ds-pagination`: `.ds-pagination__link` pills; `is-current` is filled with `--fill-button`, `is-disabled` is grey.
- `.ds-breadcrumb`: a small trail of `.ds-breadcrumb__item` parted by small arrows; the last is `.ds-breadcrumb__current`.
- `.ds-dialog`: a dimmed block (`--color-overlay`) centring `.ds-dialog__box`: `.ds-dialog__head` (`.ds-dialog__title`, a round `.ds-dialog__close`), `.ds-dialog__body` and `.ds-dialog__actions` with the main action first.
- `.ds-empty`: a centred message in a frame of `--border-width-strong`: `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one small secondary button.
- `.ds-helpful`: a strip on `--color-surface-alt` asking whether a step worked, with two small buttons. `.ds-callout`: a rounded block used with a tone class: `.ds-callout__title`, `.ds-callout__text` and one button.
- `.ds-reveal`: the scroll animation described under Layout; `--2`, `--3`, `--zoom`.
- `.ds-footer`: the closing band on `--fill-bar-alt` with rounded upper corners: `.ds-footer__cols` with the brand and a `.ds-footer__about` line, then four columns of `.ds-footer__title` and `.ds-footer__list` of `.ds-footer__link`; `.ds-footer__legal` under a rule.

## Never

- `text-shadow = none`: text is flat on every ground.
- `border-width <= 8px`: borders are hairlines and 3px accents; only the band of the drawn headphones is 8px.
- `border-radius <= 48px`: panels are 28px and the largest blocks 48px; only pills and circles are rounder.
- `box-shadow-blur <= 60px`: only hovered controls and floating things cast a shadow, and it ends at 60px.
- `font-weight <= 600`: headlines are light; 600 is the heaviest text.
- `font-size >= 13px`: meta, badges and the footer are the smallest text.
- `font-size <= 96px`: the hero title and the figures are the largest text.
- `font-families <= 3`: a serif for headings, a sans for the rest, monospace for codes.
- `line-height <= 1.55`: body text is set at 1.55.
- `uppercase-text <= 0%`: navigation, buttons and labels are in sentence case.
- `underlined-links <= 5%`: links are told apart by colour and weight; only a hovered link is underlined.
- `gradient-fills <= 0%`: every fill is flat.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. First decide which ground it stands on (page, panel or fill; inverse; accent) and take the text token and the buttons of that ground; do not put a primary button on an inverse or accent block, or `--color-inverse-text` on a fill. A new block of the hub is a `.ds-section`, with a tone class when it is not on the page colour, and a `.ds-section__head`; anything that holds data is a `.ds-panel`; a new product is another `.ds-device` variant drawn from `--color-inverse`, `--color-inverse-text` and one spot of `--color-accent`. Things that float (a menu, a tooltip, a popover) take `--shadow-dialog` and stay hidden until hover or focus. Keep corners, shadows and borders on their tokens, spacing on the `--space-*` scale, type sizes on the `--text-*` tokens (large titles through `clamp()` between two of them), and note any one-off `color-mix()` shade next to its rule.
