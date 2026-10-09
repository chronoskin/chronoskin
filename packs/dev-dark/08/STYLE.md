# Dark developer tools customer stories

## Summary

This is the customer stories section of a developer product in the "technical grid" strand of 2022 to 2026, shown light as that look often is: off-white paper with a faint dot grid, dark hairlines ruled across the whole window and one hot colour used sparingly. A masthead with the brand in its middle stands over a mono strip of facts, the lead story sets a very large bold headline beside a flat tinted plate that carries its one figure, and the other stories are cells that share their borders, each with its own number. In this set headings are a heavy tight grotesque, labels are small monospace capitals, corners are square with pill tags, and the accent is a vermilion.

## Layout

- Designed for a 1440px viewport, not fluid beyond it. Content sits in `.ds-wrap`, at most `--size-page` (1320px) wide with `--space-6` (36px) of side padding and no frame lines. `.ds-band` is a row: a hairline across the whole window above it and `--space-7` (56px) of padding; `--flush` has none (its columns pad themselves), `--last` leaves `--space-8` above the footer.
- Navigation is `.ds-nav`, which stays at the top of the window. Its first row, `.ds-nav__inner` (`--size-top` 68px) on `--fill-bar`, has four links at the left (the first drops a `.ds-menu`), the brand centred at `--text-h2`, and a secondary and a primary small button at the right; the current link carries a `--border-width-strong` line in `--color-accent-alt` along the bottom of the row. Under it `.ds-nav__strip`, a mono strip on `--fill-bar-alt` between two hairlines, spreads four facts across the width. There is no rail.
- The home view stacks: `.ds-hero`, two columns 7 to 5 (label, the `--text-display` headline, lead, two large buttons and the byline at the left; `.ds-plate` with the story's figure and two meters at the right of a hairline); a row with `.ds-stat`, four figures each under a strong rule; a row with a `.ds-colhead` and `.ds-grid`, six story cells in one bordered box; a flush row with `.ds-split`, the quote `.ds-carousel` beside the numbered `.ds-list`; and `.ds-cta`, the one inverted block, a band across the whole window.
- Inner pages have no hero. They open with `.ds-page-header`: the breadcrumb, then a `--text-h1` title and one line at `--text-large` at the left with at most two buttons at the right; a story adds `.ds-page-header__meta`, a row of byline, date and tag under a strong rule. Below it one row holds the work: `.ds-article` (facts `--size-facts` 264px, text, results `--size-side` 304px), `.ds-cols--menu` (a 264px column of sidebars at the left) or `.ds-cols` (a 304px column at the right). Tabs sit on the rule of a `.ds-toolbar` directly over the table they filter; the pagination is in the `.ds-table__foot`.
- Views. The specimen is the customer stories of chronoskin in four views. `stories` is the home view described above. `story` is one case study: the panel of facts and a row of links at the left, the text as prose with a pull quote in the middle, four results with meters and a tooltip at the right. `index` lists every story: the sidebars and the key of state labels at the left; tabs with a switch, the table with pagination and an empty state at the right. `submit` has two notices, the form with a switch and the questions as an accordion at the left; the draft with its meter and buttons, the open dialog that discards it and the editor at the right.
- Spacing scale: `--space-1` 4px (small button vertical padding, sidebar rows, between carousel marks), `--space-2` 8px (button and input vertical padding, between adjacent buttons and badges, strip padding), `--space-3` 12px (table cell and list row vertical padding, above a figure's label), `--space-4` 16px (panel and story cell padding, between form rows), `--space-5` 24px (between stacked blocks, between figures, plate padding), `--space-6` 36px (side padding, under a column head, between the columns of a split), `--space-7` 56px (row padding, between main and side column), `--space-8` 88px (above the footer).
- Other sizes: `--size-plate` 148px (the tinted block of a story cell), `--size-index` 56px (the number column of the list), `--size-tick` 10px (crop marks), `--size-mark` 28px, `--size-dot` 8px, `--size-dialog` 460px, `--size-narrow` 660px.
- Below 1080px the brand moves to the left of the masthead, every column pair stacks, the results of a story go under its text and cells and figures stand two across. Below 900px the strip keeps two facts (`.ds-nav__fact--wide` is not shown). Below 640px everything is one column, the links take a second row of the masthead, and a wide table scrolls inside `.ds-table__wrap`; the page itself never scrolls sideways.
- On a phone (640px and below) the masthead is two columns: the brand and `.ds-nav__actions` on the first row and `.ds-nav__links` on a second row spread across it; the strip keeps the two facts it has had since 900px. The drop-down of a link that has a `.ds-menu` opens as a list `70vw` wide whose items may wrap onto several lines.

## Typography and colour roles

The conventions for which text sits on which fill are those of the era's first pack; this layout keeps to them and assumes neither a light nor a dark palette.

- `--color-text`, `--color-heading` and `--color-text-muted` sit on `--color-page`, `--fill-panel`, `--color-surface-alt` and `--color-surface-strong`. `--color-bar-text` sits on `--fill-bar` (the masthead). `--color-bar-alt-text` sits on `--fill-bar-alt` (the strip, a table head, the title strip of a code block, the footer). `--color-inverse-text` sits on `--fill-inverse`, used for the call to action and the tooltip. `--color-button-text` sits on `--fill-button`, `--color-accent-text` on `--fill-accent`, `--color-input-text` on `--fill-input`, `--color-notice-text` on `--color-notice`.
- `--color-fill-1` to `--color-fill-4` never carry text at full strength. They are flat tints of 16% to 18% behind text in the ordinary text colours (the plate and the block of a story cell), and 22% in an avatar. Nothing glows.
- `--color-accent-alt` is a text and line colour: numbered labels, the line of the current navigation link and tab, the rule of a pull quote, the numbers of the list, the filled part of a meter, the knob of a switch that is on, the unit after the plate's figure.
- `--color-success`, `--color-warning` and `--color-danger` appear as a dot, as text, or as a 12% tint with a 32% hairline; the destructive button is built the same way.
- Borders: `--color-border` on rows, boxes and between cells; `--color-border-strong` on secondary buttons, the dialog, avatars, the dashed empty state and, at `--border-width-strong`, the rule over a figure, a result and a story's meta row; `--color-border-muted` between rows inside a box. Crop marks are `--color-text-muted`.
- Type. A heavy grotesque for headings, a plainer one for text and controls, one monospace for labels and data. `--text-base` 16px at `--line-body` 1.55; `--text-ui` 14px for controls, tables and card text; `--text-small` 12px for mono labels, badges and dates; `--text-large` 21px for leads. `--text-display` 76px is the lead headline, and the plate's figure is one and a half times that; `--text-h1` 46px page titles, figures and quotes; `--text-h2` 28px the pull quote and the brand in the masthead; `--text-h3` 18px cell and panel titles.
- Uppercase is for mono labels only (`.ds-eyebrow`, the strip, table heads, form labels, the breadcrumb, footer titles), tracked by 8% of their size; `--ui-transform` applies to buttons, tabs, badges and navigation.
- Links are `--color-link` with `--link-decoration`; quiet links in the breadcrumb, sidebar, table and footer are `--color-link-quiet` and turn `--color-heading`.
- Surface. `--radius-control` on buttons, inputs, the brand tile and carousel arrows; `--radius-panel` on panels, notices, code blocks, the block of a story cell, the menu and the empty state; `--radius-page` on the plate, the box of story cells and the dialog; `--radius-pill` on badges, avatars, dots, meters and switches. `--shadow-panel` is on panels and the box of story cells, `--shadow-control` on buttons and the brand tile, `--shadow-dialog` on the plate, the dialog and the menu. `--backdrop-blur` blurs what scrolls under the masthead.
- One-off values written with `calc()` and `color-mix()` because the vocabulary has no token: the tints named above with a 40% hairline around the plate, navigation links at `opacity` 0.68 until current, the text of the inverted block at 0.72 and footer titles at 0.7, the hovered destructive button (22%), the hovered inverse button (82%), the notice border (24% and 40%), key cap and inline code corners (half of `--radius-control`), mono label tracking, the headline capped at 10vw and the plate's figure at 22vw so that they shrink on a narrow window.

## Components

- `.ds-page`: on `<body>`. Sets the font, text colour and `--fill-page`. `.ds-nav`, the views and `.ds-footer` are its direct children.
- `.ds-wrap`, `.ds-band`, `.ds-stack`, `.ds-cols`, `.ds-split`, `.ds-colhead`: the centred column, a ruled row (`--flush`, `--last`), a vertical stack 24px apart, the column pairs (`--menu`, `--even`), two columns of a flush row with a hairline between (`.ds-split__cell`), and the head of a block (`.ds-colhead__title` under a `.ds-eyebrow`, a link at the right).
- `.ds-brand`: the site's mark and name. `.ds-brand__mark` is a 28px tile dressed like the primary button holding the mark as inline SVG; `.ds-brand__name` is the name. Both are placeholders for the installing project's own name and logo.
- `.ds-nav`: the masthead. `.ds-nav__inner` holds `.ds-nav__links` of `.ds-nav__link` (current: `is-current`), the brand and `.ds-nav__actions`; `.ds-nav__strip` holds `.ds-nav__facts` of `.ds-nav__fact` (`--wide` is dropped on a narrow window). In the specimen the showing view marks its link with one `:has()` rule per view; a project sets `is-current`.
- `.ds-menu`: on the list item of a navigation link; its `.ds-menu__list` of `.ds-menu__item` drops under the link while it is hovered or focused (`is-open` shows it statically). No script.
- `.ds-hero`: home view only. `.ds-hero__inner` is two columns: `.ds-hero__copy` with a `.ds-eyebrow`, `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__actions` and the byline `.ds-hero__by`; `.ds-hero__side` with the plate.
- `.ds-plate`: a flat tinted block with crop marks in two corners: a label and `.ds-plate__text`, the large `.ds-plate__figure` (its `small` is the unit) and `.ds-plate__bars` for meters.
- `.ds-page-header`: head of an inner page: the breadcrumb, `.ds-page-header__inner` with `.ds-page-header__copy` (`.ds-page-header__title`, `.ds-page-header__text`) and `.ds-page-header__actions`, then an optional `.ds-page-header__meta`.
- `.ds-prose`: running text: h1, h2, h3, p, lists, `code`, `strong`; `.ds-prose__lead` for the first paragraph. `.ds-quote` is a pull quote inside it.
- `.ds-article`: the three columns of a story. `.ds-result` is one result in its right column: `.ds-result__label`, `.ds-result__figure` and a meter.
- `.ds-code`: a code block on `--fill-input`: optional `.ds-code__head` with `.ds-code__name`, then `.ds-code__body` (a `pre`). Spans `.ds-code__k`, `__s`, `__c`, `__n` colour keywords, strings, comments and names.
- `.ds-eyebrow`, `.ds-mono`, `.ds-kbd`, `.ds-dot`: the mono label (`--accent`), mono data text, a key cap and a status dot (`--success`, `--warning`, `--danger`, `--accent`).
- `.ds-avatar`: initials on a tile shaped by `--radius-pill`; `--2` and `--3` tint it, `--large` doubles it, `.ds-avatars` overlaps several. `.ds-person` puts one beside `.ds-person__name` and `.ds-person__role`.
- `.ds-link`: text link with `is-visited`, `is-hover`, `is-active`, `is-focus`; `--quiet`, `--more` (appends an arrow). `.ds-linkrow` lays links in a row.
- `.ds-button`: 14px medium text. `--secondary` (strong hairline), `--danger` (tint, hairline and text in the danger colour), `--inverse` (on the inverted block only), `--small`, `--large`. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttonrow` spaces them 8px apart.
- `.ds-form`: mono labels above fields. `.ds-form__row`, `.ds-form__field`, `.ds-form__label`, `.ds-form__input` (`--mono`), `.ds-form__select`, `.ds-form__textarea`, `is-error` plus `.ds-form__error`, `.ds-form__hint`, `.ds-form__check` with `.ds-form__checkbox`, `.ds-form__actions`.
- `.ds-switch`: a label holding `.ds-switch__input` (a real checkbox, invisible over the whole control), `.ds-switch__track` and an optional `.ds-switch__label`. `is-on` and `is-focus` show the states statically.
- `.ds-progress`: a meter: `.ds-progress__head` (name and `.ds-progress__value`) over `.ds-progress__bar`, a native `progress` element painted flat; `--quiet` fills it in the muted colour.
- `.ds-tooltip`: a small ringed trigger whose `.ds-tooltip__tip` shows under it on hover or focus (`is-open` statically), on `--fill-inverse`.
- `.ds-accordion`: ruled `details` elements: `.ds-accordion__item`, `.ds-accordion__summary` with a plus that becomes a minus, `.ds-accordion__body`.
- `.ds-carousel`: `.ds-carousel__track` is a row of `.ds-carousel__slide` items that snaps to one slide (`.ds-carousel__quote` over a `.ds-person`); `.ds-carousel__nav` holds `.ds-carousel__dots` of `.ds-carousel__dot` (`is-current`), a count and two `.ds-carousel__arrow` marks (`is-disabled`). It moves by scrolling, without script.
- `.ds-table`: inside a panel, edge to edge: a mono uppercase head on `--fill-bar-alt`, hairline rows, `is-hover` rows on `--color-surface-alt`, `.ds-table__num` for right-aligned mono numbers, `.ds-table__state` for a dot with a word; `--ruled`, `.ds-table__yes` and `.ds-table__no` for a matrix. Wrap in `.ds-table__wrap`; `.ds-table__foot` closes the panel.
- `.ds-list`: numbered rows under a rule: `.ds-list__item` is `.ds-list__no` beside a `.ds-list__title` link and a mono `.ds-list__meta`.
- `.ds-panel`: hairline box on `--fill-panel`: `.ds-panel__head` with `.ds-panel__title`, `.ds-panel__body`; `.ds-panel__rows` of `.ds-panel__row` (`.ds-panel__key`) for facts.
- `.ds-stat`: the row of figures. Each `.ds-stat__item` is a mono `.ds-stat__label` under a strong rule, a `.ds-stat__figure` and a muted `.ds-stat__text`.
- `.ds-grid`: story cells in one bordered box. `.ds-grid__cell` (a link) holds `.ds-grid__plate` (`--2`, `--3`, `--4` for the other fills) with `.ds-grid__figure` and `.ds-grid__unit`, then `.ds-grid__meta`, `.ds-grid__title` and `.ds-grid__text`. Keep the number of cells a multiple of the columns.
- `.ds-cta`: the inverted band: `.ds-cta__inner` with `.ds-cta__title`, `.ds-cta__text` and one `.ds-button--inverse`.
- `.ds-tabs`: words on a rule, `.ds-tabs__tab` with an optional `.ds-tabs__count`; the current one (`is-current`) has a strong line under it. `.ds-toolbar` is the ruled row that holds it and a control at the right.
- `.ds-badge`: small mono pill on `--fill-accent`. `--quiet`, `--outline`. Status variants: `--success`, `--warning`, `--danger`. `.ds-badgerow` for several.
- `.ds-sidebar`: a ruled mono `.ds-sidebar__title` over `.ds-sidebar__list`; `.ds-sidebar__item`, current with `is-current` and a strong line at its left.
- `.ds-notice`: message strip on `--color-notice` with a strong line at its left; `--error` in the danger colours. `.ds-notice__title`; `.ds-noticerow` stacks notices.
- `.ds-pagination`: joined mono cells, `.ds-pagination__link`; `is-current` on `--color-surface-strong`, `is-disabled` grey.
- `.ds-breadcrumb`: mono capitals, `.ds-breadcrumb__item` parted by slashes; the last is `.ds-breadcrumb__current`.
- `.ds-dialog`: a block in `--color-overlay` centring `.ds-dialog__box` with `--shadow-dialog`: `.ds-dialog__head` (`.ds-dialog__title`, `.ds-dialog__close`), `.ds-dialog__body`, `.ds-dialog__actions`.
- `.ds-empty`: a dashed box: `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one small button.
- `.ds-footer`: on `--fill-bar-alt`: `.ds-footer__inner` with `.ds-footer__brand` (`.ds-footer__about`) and three columns of `.ds-footer__title` and `.ds-footer__links` of `.ds-footer__link`; `.ds-footer__line` is the mono last line with `.ds-footer__status`.

## Never

- `border-width <= 3px`: rules are 1px hairlines; 3px is the strong rule over a figure and under the current link.
- `border-radius <= 3px`: boxes are square and controls eased by 3px; pills and circles are shapes.
- `box-shadow-blur <= 48px`: only the plate and the dialog cast a soft shadow.
- `text-shadow = none`: text is flat on every fill.
- `font-size >= 12px`: mono labels, badges and dates are the smallest text.
- `font-size <= 114px`: the figure on the plate is the largest text, the headline is 76px.
- `font-weight <= 700`: headings are bold; nothing is black.
- `font-families <= 3`: a heavy grotesque, a text grotesque and one monospace.
- `line-height <= 1.6`: body text is set at 1.55.
- `uppercase-text <= 10%`: only mono labels are uppercase.
- `letter-spacing <= 4px`: only the headline and the plate's figure are pulled together, by 3.2px.
- `gradient-fills <= 3%`: the dot grid of the page is the only gradient; no box is filled with one.
- `animation = none`: nothing moves by itself.
- `block-gap <= 88px`: rows are padded by 56px, the footer stands 88px below the last.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new block of the page is a `.ds-band` with a `.ds-wrap`, opened by a `.ds-colhead` with a numbered mono label; a result or a claim gets its number first, large, under a strong rule or on a tinted plate, and the words after it. Several like things become cells of one bordered box in the manner of `.ds-grid`, not separate cards with space between. Labels are mono capitals tracked by 8%, data is mono. Keep to the conventions above for which text sits on which fill: never put text on a full-strength `--color-fill-*` or status colour, tint it instead. Use the accent for one thing per block and `--fill-inverse` for one band per view. Spacing comes from the `--space-*` scale, and any one-off shade is a `color-mix()` of tokens noted next to the rule.
