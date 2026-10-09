# Dark developer tools landing page

## Summary

This is the product page of a developer tool as it looked from 2019 to 2026: a near-black page, hairline borders in translucent white, small precise sans-serif type and one soft glow of colour behind the hero. The claim is centred over the glow, the product is shown as a window with a status list and a build log, and below it come customer names set as text, a ruled strip of figures and a grid of bordered cards that each hold a small line drawing. Monospace is used for every detail that a developer would read as data: commands, dates, ids, labels above sections and the heads of tables.

## Layout

- The page is designed for a 1440px viewport and is not fluid: content sits in one centred column, `.ds-wrap`, at most `--size-page` (1200px) wide with `--space-5` (24px) of side padding. Running text is kept to `--size-narrow` (720px).
- Navigation is `.ds-nav`, a bar `--size-top` (56px) tall that stays at the top of the window on the translucent `--fill-bar`, blurred by `--backdrop-blur`, with a hairline below. Brand at the left, four links after it, a secondary and a primary small button at the right. There is no sidebar on the home view.
- The home view stacks, top to bottom: `.ds-hero` (announcement pill, title, lead, two large buttons, a mono command line, then `.ds-frame`, the product window); `.ds-logos`; `.ds-stat`, four figures in one ruled strip; a band with `.ds-section__head` and `.ds-grid` (three columns, the first card two columns wide); a band with `.ds-cols`, a code panel beside the dated `.ds-list`; and `.ds-cta`, the one inverted block. Bands are `.ds-band`, padded by `--space-8` (88px) and parted by a hairline in `--color-border-muted`.
- Inner pages have no hero. They open with `.ds-page-header`: breadcrumb, a `--text-h1` title and one muted line at the left, at most two buttons at the right, over a small glow. Below it the work sits in `.ds-wrap`: a `.ds-stack` (24px apart) or `.ds-cols`. `.ds-cols` is two equal columns; `--wide` is 2 to 1 and `--side` is a main column beside a `--size-side` (300px) column. Tabs are a segmented control in a `.ds-toolbar` directly under the page header.
- Views. The specimen is a deployment service with four views. `home` is the landing page described above. `pricing` has the tabs for the billing period, three `.ds-plan` cards, the comparison table beside the panel of the current plan with its buttons, and the open dialog that cancels the plan. `changelog` has two `.ds-entry` articles and the pagination in the main column; the releases sidebar, the label key, an empty state and a row of links at the right. `start` is the sign-up form under its notices, beside the steps, a terminal block and the limits of the free plan.
- Spacing scale: `--space-1` 4px (inside pills and key caps, between pagination links), `--space-2` 8px (button vertical padding, between adjacent buttons and badges), `--space-3` 12px (panel head and table cell vertical padding), `--space-4` 16px (panel padding, between grid cards, between form rows), `--space-5` 24px (between stacked blocks, side padding, plan padding), `--space-6` 32px (between columns, around a dialog and an empty state), `--space-7` 48px (page header top, hero to frame, inverted block padding), `--space-8` 88px (band and hero padding, above the footer).
- Other sizes: `--size-art` 168px (the drawing area of a card), `--size-glow` 560px (height of the hero glow), `--size-date` 120px (the date column of lists and entries), `--size-mark` 24px, `--size-dot` 8px, `--size-dialog` 440px.
- Below 980px cards stand two across and every column pair stacks. Below 640px everything is one column, the links of the bar move to a second row that scrolls sideways, and a wide table scrolls inside `.ds-table__wrap`; the page itself never scrolls sideways.
- On a phone (640px and below) `.ds-nav__inner` wraps: the brand and `.ds-nav__actions` stay on the first row, `.ds-nav__links` take a second row of their own that scrolls sideways, each link sharing the width equally with centred text; the secondary button of the actions is not shown. The `.ds-menu` stays in the actions and lists every view as `.ds-menu__item--view` rows, the current one with the accent line. A tooltip in a panel title opens under its line with a capped width.

## Typography and colour roles

These are the conventions of the era. Every layout and every token set of dev-dark keeps to them.

- Which text sits on which fill. `--color-text`, `--color-heading` and `--color-text-muted` sit on `--color-page`, `--color-canvas`, `--fill-panel`, `--color-surface-alt` and `--color-surface-strong`. `--color-bar-text` sits on `--fill-bar` (the top bar). `--color-bar-alt-text` sits on `--fill-bar-alt` (a side rail, the title strip of a code block, a table head, the footer); a heading on that fill may use `--color-heading`. `--color-inverse-text` sits on `--fill-inverse`, which is used for exactly one block per view. `--color-button-text` sits on `--fill-button`, `--color-button-secondary-text` on `--fill-button-secondary`, `--color-accent-text` on `--fill-accent`, `--color-input-text` on `--fill-input` (form controls and code blocks), `--color-notice-text` on `--color-notice`.
- `--color-fill-1` to `--color-fill-4` never carry text. They are glows and tints: mixed with `transparent` at 12% to 34% behind text in the ordinary text colours, and at full strength or mixed 72% into `--color-heading` only for drawings, bars, cells and icons. `--color-fill-1` is the palette's glow colour: the hero, the page header, the featured plan. A palette may make any fill light or dark.
- `--color-accent-alt` is a text and line colour, not a fill: mono labels, keywords in code, the marker of the current item, the "new" badge. It must read on the page and on a panel.
- `--color-success`, `--color-warning` and `--color-danger` are used as a dot, as text, or as a 12% tint with a 32% hairline and the colour itself as text. Nothing is ever filled solid with them, so they need no text colour of their own. The destructive button is built the same way.
- Borders: `--color-border` on every box and bar, `--color-border-strong` on secondary buttons, the dialog, the product window and the dashed empty state, `--color-border-muted` between rows inside a box and between bands. All `--border-width`; `--border-width-strong` is the focus offset and the ring around a status dot.
- Type. One sans family for body, headings and controls, with `--font-mono` for data. `--text-base` 15px at `--line-body` 1.6; `--text-ui` 13px for navigation, buttons, forms, tables and card text; `--text-small` 12px for badges, mono labels, dates and table heads; `--text-large` 17px for the hero lead. `--text-display` 56px is the hero title only; `--text-h1` 36px page titles, section statements, figures and prices; `--text-h2` 24px column heads; `--text-h3` 15px panel and card titles in `--weight-bold`. Headings are `--weight-heading` 500 with slight negative tracking.
- Uppercase appears only on mono labels (`.ds-eyebrow`, table heads, sidebar and footer titles), tracked by 6% of the font size; `--ui-transform` applies to buttons, tabs and navigation.
- Links are `--color-link` without underline and gain one on hover; quiet links in breadcrumbs, the sidebar and the footer are `--color-link-quiet` and turn `--color-heading`.
- Surface. `--radius-control` 6px for buttons, inputs and the brand tile, `--radius-panel` 8px for panels, code blocks, notices and the empty state, `--radius-page` 12px for the product window, grid cards, plans, the dialog and the inverted block, `--radius-pill` for badges and the announcement pill. `--shadow-panel` is on panels and cards, `--shadow-control` on buttons and the brand tile, `--shadow-dialog` on the dialog and the product window. `--shadow-text` reaches navigation links and buttons.
- One-off values written with `calc()` and `color-mix()` because the vocabulary has no token: the glows (a fill at 18% to 34% over transparent), the drawing colour (72% of a fill in `--color-heading`), the wash behind the current navigation link (`--color-bar-text` at 9%), status tints (12% and 32%), the hovered destructive button (22%), the hovered inverse button (82% of `--color-inverse-text` in `--color-inverse`), the notice border (22% and 40%), key cap and inline code corners (half of `--radius-control`), the segmented control's corner (`--radius-control` plus `--space-1`), mono label tracking. Navigation links are at `opacity` 0.66 until current, customer names at 0.82, the text of the inverted block at 0.7. The empty state's outline is dashed.

## Components

- `.ds-page`: on `<body>`. Sets the font, text colour and `--fill-page`. `.ds-nav`, the views and `.ds-footer` are its direct children.
- `.ds-wrap`, `.ds-band`, `.ds-stack`, `.ds-cols`: the centred column, a padded section with a rule above (`--tight`, `--flush`), a vertical stack, and the column pairs (`--wide`, `--side`). `.ds-panel--grow` in a stack takes up spare height.
- `.ds-brand`: the site's mark and name. `.ds-brand__mark` is a 24px tile dressed like the primary button holding the mark as inline SVG; `.ds-brand__name` is the name. Both are placeholders for the installing project's own name and logo.
- `.ds-nav`: the top bar. `.ds-nav__inner` holds the brand, `.ds-nav__links` of `.ds-nav__link` (current: `is-current`) and `.ds-nav__actions`. In the specimen the showing view marks its link with one `:has()` rule per view; a project sets `is-current`.
- `.ds-hero`: home view only. `.ds-hero__pill`, `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__actions`, `.ds-hero__cmd`, then `.ds-frame`. The glow is its `::before`.
- `.ds-frame`: the product drawn as a window: `.ds-frame__bar` (lights, path, a status badge), `.ds-frame__body` with `.ds-frame__rows` of `.ds-frame__row` and the mono `.ds-frame__log`.
- `.ds-logos`: `.ds-logos__row` of `.ds-logos__name` (`--mono`, `--light`, `--italic`) and a mono caption. Text only, never an image.
- `.ds-page-header`: head of an inner page: `.ds-page-header__inner` with `.ds-page-header__copy` (breadcrumb, `.ds-page-header__title`, `.ds-page-header__text`) and `.ds-page-header__actions`.
- `.ds-prose`: running text: h1, h2, h3, p, lists, `code`, `strong`; `.ds-prose__lead` for the first paragraph.
- `.ds-code`: a code block on `--fill-input`: optional `.ds-code__head` with `.ds-code__name`, then `.ds-code__body` (a `pre`). Spans `.ds-code__k`, `__s`, `__c`, `__n` colour keywords, strings, comments and names.
- `.ds-eyebrow`, `.ds-mono`, `.ds-kbd`, `.ds-dot`, `.ds-avatar`: the mono label (`--accent`), mono data text, a key cap, a status dot (`--success`, `--warning`, `--danger`, `--accent`, `--ring`) and an initials disc.
- `.ds-link`: text link with `is-visited`, `is-hover`, `is-active`, `is-focus`; `--quiet`, `--more` (appends an arrow). `.ds-linkrow` lays links in a row.
- `.ds-button`: 13px medium text, 6px corners. `--secondary` (translucent fill, strong hairline), `--danger` (tint, hairline and text in the danger colour), `--inverse` (on the inverted block only), `--small`, `--large`. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttonrow` spaces them 8px apart.
- `.ds-form`: labels above fields. `.ds-form__row`, `.ds-form__field`, `.ds-form__label`, `.ds-form__input` (`--mono`), `.ds-form__select`, `.ds-form__textarea`, `is-error` plus `.ds-form__error`, `.ds-form__hint`, `.ds-form__check` with `.ds-form__checkbox`, `.ds-form__actions`.
- `.ds-table`: inside a panel, edge to edge: a mono uppercase head on `--fill-bar-alt`, hairline rows, `is-hover` rows on `--color-surface-alt`, `.ds-table__num` for right-aligned mono numbers, `.ds-table__state` for a dot with a word. Wrap in `.ds-table__wrap`; `.ds-table__foot` closes the panel.
- `.ds-list`: dated rows: `.ds-list__item` is `.ds-list__meta` (mono date) beside a `.ds-list__title` link and `.ds-list__text`.
- `.ds-panel`: hairline box on `--fill-panel`: `.ds-panel__head` with `.ds-panel__title`, `.ds-panel__body`; `.ds-panel__rows` of `.ds-panel__row` (`.ds-panel__key`) for facts. A `.ds-code` directly inside loses its own border.
- `.ds-stat`: the row of figures. Each `.ds-stat__item` is a `.ds-stat__figure` over a mono `.ds-stat__label`, parted by vertical hairlines.
- `.ds-section__head`: `.ds-eyebrow` and `.ds-section__title`, whose `strong` part is bright and the rest muted.
- `.ds-grid`: feature cards. `.ds-grid__cell` (`--wide`) holds `.ds-grid__art` (`--2`, `--3`, `--4` for the other fills) with one inline SVG drawn in strokes, and `.ds-grid__body` with `.ds-grid__title` and `.ds-grid__text`.
- `.ds-cta`: the inverted block: `.ds-cta__title`, `.ds-cta__text` and one `.ds-button--inverse`.
- `.ds-tabs`: a segmented control of `.ds-tabs__tab`; the current one (`is-current`) is on `--color-surface-strong`. `.ds-toolbar` puts it in a row with a `.ds-toolbar__note`.
- `.ds-plans`: three `.ds-plan` cards (`--featured` carries the glow): `.ds-plan__name`, `.ds-plan__price` with `.ds-plan__per`, `.ds-plan__text`, a button, `.ds-plan__items` of `.ds-plan__item`.
- `.ds-entry`: a changelog entry: `.ds-entry__meta` (date and badge) beside a `.ds-prose`.
- `.ds-badge`: small pill on `--fill-accent`. `--quiet`, `--outline`, `--new` (accent-alt tint), `--mono`. Status variants: `--success`, `--warning`, `--danger`. `.ds-badgerow` for several.
- `.ds-sidebar`: a mono `.ds-sidebar__title` over `.ds-sidebar__list` along a rule; `.ds-sidebar__item`, current with `is-current`.
- `.ds-steps`: numbered `.ds-steps__item` rows with a `.ds-steps__title`.
- `.ds-notice`: message strip on `--color-notice`; `--error` in the danger colours. `.ds-notice__title`; `.ds-noticerow` stacks notices.
- `.ds-pagination`: `.ds-pagination__link`; `is-current` on `--color-surface-strong`, `is-disabled` grey.
- `.ds-breadcrumb`: `.ds-breadcrumb__item` parted by slashes; the last is `.ds-breadcrumb__current`.
- `.ds-dialog`: a block in `--color-overlay` centring `.ds-dialog__box` with `--shadow-dialog`: `.ds-dialog__head` (`.ds-dialog__title`, `.ds-dialog__close`), `.ds-dialog__body`, `.ds-dialog__actions`.
- `.ds-empty`: a dashed box: `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one small button.
- `.ds-footer`: on `--fill-bar-alt`: `.ds-footer__inner` with `.ds-footer__brand` (`.ds-footer__about`) and three columns of `.ds-footer__title` and `.ds-footer__links` of `.ds-footer__link`; `.ds-footer__line` is the mono last line with `.ds-footer__status`.
- `.ds-menu`: a drop-down behind a button in the bar (`.ds-menu__button`, `.ds-menu__list`); it lists every view, so on a phone it is the full navigation.
- `.ds-menu__item`: one link of the drop-down; the current view carries the accent line and the stronger fill.
- `.ds-menu__item--view`: an item of the menu that names a view of the site; the current view's item is marked by one `:has()` rule per view, like the bar links, so the menu is a second way to every view on a phone.
- `.ds-menu__title`: a small mono label over a group of links in the drop-down.
- `.ds-carousel`: customer quotes in a row that scrolls and snaps (`.ds-carousel__track`, `.ds-carousel__slide`); the dots only count and are not links.
- `.ds-avatar`: initials in a circle for a person, in the account menu, quotes and lists.
- `.ds-tooltip`: a small question mark beside a title; `.ds-tooltip__tip` opens on hover or focus, never by default.
- `.ds-progress`: a thin bar for a build, a rollout or a quota, with a mono line above (`.ds-progress__head`, `.ds-progress__value`, `.ds-progress__bar`).
- `.ds-switch`: an on or off setting that takes effect at once, drawn over a checkbox (`.ds-switch__input`, `.ds-switch__track`, `.ds-switch__label`).
- `.ds-accordion`: questions as `details` items (`.ds-accordion__item`, `.ds-accordion__summary`, `.ds-accordion__body`); the first is open.

## Never

- `border-width <= 2px`: every rule and outline is a 1px hairline; 2px is the most any palette's emphasis takes.
- `border-radius <= 12px`: 6px on controls, 8px on panels, 12px on the largest boxes; pills are shapes.
- `box-shadow-blur <= 32px`: only the dialog and the product window cast a wide shadow.
- `text-shadow = none`: text is flat on every fill.
- `font-size >= 12px`: badges, mono labels and dates are the smallest text.
- `font-size <= 56px`: the hero title is the largest text.
- `font-weight <= 600`: headings are medium, labels semibold; nothing is bold or black.
- `font-families <= 2`: one sans-serif family and one monospace.
- `line-height <= 1.6`: body text is set at 1.6.
- `uppercase-text <= 5%`: only mono labels are uppercase.
- `underlined-links <= 5%`: links are told apart by colour and underline on hover only.
- `letter-spacing <= 3px`: only the mono customer names are spaced out, by 2px.
- `gradient-fills <= 5%`: gradients are the glows behind the hero, the page header and the card drawings, never a filled box.
- `animation = none`: nothing moves by itself.
- `block-gap <= 88px`: bands are 88px apart at most, blocks inside them 24px.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new block on the page is a `.ds-panel` or a `.ds-grid__cell` with a 1px border in `--color-border`, placed in `.ds-wrap` inside a `.ds-band`; a new feature gets a card with one stroke drawing in the next fill colour; data a developer would copy is set in `--font-mono`. Keep to the conventions above for which text sits on which fill: never put text on a full-strength `--color-fill-*` or status colour, tint it instead. Use one glow per view, always from `--color-fill-1`, and one inverted block. Spacing comes from the `--space-*` scale, and any one-off shade is a `color-mix()` of tokens noted next to the rule.
