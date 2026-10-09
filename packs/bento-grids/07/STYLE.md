# Documentation hub bento, 2025

## Summary

A documentation home page laid out as a bento: the ways into the docs (a quickstart, the reference, the libraries, the guides, the status, what is new) are rounded cells of unequal width in a grid of five tracks, with the largest cell asking the reader where to start. Cells are cream on sage paper and every one is drawn with a 2px ink outline and no shadow, a few of them filled flat with lime, butter or clay, and the header is the two-row bar of a documentation site with a search field in its middle. The form was common for developer documentation, help centres and product handbooks from about 2022 to 2026.

## Layout

- Fixed centred column: `--size-page` is 1240px, designed at a 1440px viewport, padded by `--space-5` (20px) at the sides. The header's and the footer's rules run the full width; everything else stays in the column.
- `.ds-nav` is sticky and has two rows. The first, `--size-nav` (56px) tall, holds the `.ds-brand`, a version badge, the search field `.ds-nav__search` (at most `--size-search`, 400px, centred) and at the right one small button and the account `.ds-avatar` with its `.ds-menu`. Under a faint rule the second row, `--size-subnav` (44px), holds the sections as plain text links; the current one is underlined with `--border-width-strong` in `--color-link`, and one of them opens a `.ds-menu`. A status line sits at the right end of the second row.
- `.ds-grid` has five equal tracks (about 237px). A cell takes one track by default, two with `--w2`, three with `--w3`, all five with `--full`, and two rows with `--tall`. A cell is at least `--size-row` (212px) tall; inside `.ds-grid--auto`, in the rail and with `--flush` it follows its content.
- The home page is one grid. Rows one and two: the `.ds-hero` cell, three tracks by two rows, with a two-track code cell over two one-track cells beside it. Row three: a two-track list, a one-track status cell, a two-track `.ds-carousel`. Row four: the `.ds-stats` strip in a `--full --flush` cell. Row five: a two-track cell, a one-track cell and the two-track `--accent` cell. No two neighbouring rows split the same way.
- Cells never touch: the gap is `--space-gap` (14px) in both directions, the padding inside a cell `--space-cell` (24px), in the hero `--space-7` (40px).
- Spacing scale: `--space-1` 4px (pill padding, title to caption), `--space-2` 8px (button gap, label to control, chips), `--space-3` 12px (control padding, table cell padding, cell head to body), `--space-4` 16px (stack gap inside a cell, form gap, list row padding), `--space-5` 20px (page gutter, tab gap, under the page header), `--space-6` 28px (above the first cell, footer padding), `--space-7` 40px (hero padding, above a section label), `--space-8` 56px (under the last cell, dialog stage), `--space-9` 88px (textarea height).
- Controls are `--size-control` (36px) tall, `--size-control-large` (46px) in the hero, 28px with `--small`; icon tiles are `--size-icon` (36px) with a `--size-glyph` (18px) line icon; an avatar is `--size-avatar` (32px).
- Inner pages (a list of guides, one guide, a reference page, a settings page) have no hero. A view starts `--space-6` under the bar with the `.ds-breadcrumb`, then the `.ds-page-header` on the bare page: title at `--text-h1` with an optional badge beside it, one line of text, buttons at the right, and a rule under it. Below, `.ds-layout` splits the column into `.ds-layout__main` and a `.ds-layout__rail` of `--size-rail` (252px) at the right, sticky under the bar. The main column starts with an optional `.ds-toolbar` holding the `.ds-tabs` and continues with a `.ds-grid--auto` of the same five tracks, in which cells pair as three and two or take the full width. The rail is a stack of small cells: a `.ds-sidebar`, a progress cell, a legend. A guide has no page header: its `.ds-prose` starts with the h1 in the first cell.
- Views: `home` is the hub (hero, quickstart sample, reference, libraries, starter guides, status, what is new, the figures, the forum, change notes, the call for a key); `guides` is the list (tabs, the list with pagination, two reading paths in panels, the empty state for migration guides, topics and authors in the rail); `guide` is one guide (prose, code, notices, links, questions as an accordion, the page's contents and its author in the rail); `reference` is one resource (the endpoint table, parameters as an accordion, a sample answer, usage and the label legend in the rail); `keys` is the account page (the form for a key with its buttons, the dialog that revokes one, the table of keys, safety switches and a notice in the rail). The second row marks Overview, Guides, API reference or Keys; a guide marks Guides.
- Below 1080px the grid drops to two tracks, `--w3` and the hero take both, and the rail moves under the main column as two columns. Below 640px everything is one column, the search field takes the room between brand and avatar, the version badge, the header button, the key hint and the status line hide, the links of the second row wrap, the stats strip becomes two by two, a table hides its `.ds-table__extra` columns and scrolls inside `.ds-table__scroll`, a drop-down in the second row opens as a full row inset by `--space-4`, and a tooltip opens as a full row under its line.

## Typography and colour roles

- Two sans families: `--font-body` and `--font-ui` are IBM Plex Sans with system fallbacks, `--font-heading` is Space Grotesk; `--font-mono` (IBM Plex Mono) is for code, endpoints and the key hint. The references used commercial grotesques for headings; Space Grotesk is the open substitute.
- Sizes: `--text-base` 15px at `--line-body` 1.55; `--text-small` and `--text-ui` 14px (captions, navigation, buttons, table, form); `--text-large` 17px (hero lead, cell titles); `--text-display` 50px at `--line-display` 1.04 with `--display-tracking` -1.6px (hero title); `--text-h1` 36px (page title, figures), `--text-h2` 26px, `--text-h3` 19px at `--line-heading` 1.15 with `--heading-tracking` -0.4px. The smallest text is 12px, written `calc(var(--text-small) * 0.857)` because the vocabulary has no token for it: badges, meta lines, hints. Code is `calc(var(--text-small) * 0.886)`, an endpoint `calc(var(--text-small) * 0.928)`.
- Weights: `--weight-body` 400, `--weight-ui` 500, `--weight-bold` 600, `--weight-heading` and `--weight-display` 700. Nothing is uppercase except the initials in an avatar.
- `--color-page` `#d2d8c4` sage paper, also the bar. `--color-surface` `#fdfdf8` cream: cells. `--color-surface-alt` `#f1f3e6`: panels, the dialog, menus, the current sidebar item and the current table row. `--color-surface-strong` `#e0e4d2`: neutral badges, progress tracks, avatars, inline code, the verb of an endpoint.
- `--color-text` `#33362b`, `--color-text-muted` `#5c5e55` for captions and meta, `--color-heading` `#1b1d15`. Text on a `--color-fill-*` cell uses the same three tokens; the fills are pale enough for them.
- `--color-button` and `--color-link` `#254f1a` deep green: primary button, links, icons, the underline of the current section and tab, the accordion mark. `--color-accent` `#a0b878` sage green with `--color-accent-text` ink: the closing cell, the accent badge and avatar, progress fills. `--color-accent-alt` `#b17816` ochre marks "new", code keywords and switched-on toggles only.
- `--color-fill-1` lime, `--color-fill-2` butter, `--color-fill-3` clay, `--color-fill-4` pale sage: flat cell fills, on at most half the cells of a grid. The hero carries a soft fade of `--color-fill-1` from its top right corner.
- Borders: `--color-border` `#2b2d24` ink at `--border-width` 2px on every cell, panel, control, badge and bar; `--color-border-strong` `#111309` on secondary buttons, the dialog, menus and the empty cell; `--color-border-muted` `#c9cdbb` for rows inside a cell, the stats rules and the rule between the header's rows. `--border-width-strong` 3px is the underline of the current section and tab.
- Links in running text are underlined (`--link-decoration`); quiet links in navigation, sidebar and footer are not.
- Radii: `--radius-page` 22px for cells, `--radius-panel` 14px for panels, tables, notices, code, menus and the dialog, `--radius-control` 10px for buttons, inputs and icon tiles, `--radius-pill` for badges, switches, avatars and progress bars.
- Shadows are hard and small: `--shadow-control` is a 2px ink lip under buttons and inputs, 3px on hover, none when pressed; `--shadow-dialog` a 6px lip under the dialog, menus and tooltips; cells have none. No blur anywhere, no text shadow.
- One-off mixes written in `components.css`: the button's outline (25% button text in the button colour), status badges and `--new` (the colour at 14% with a 30% outline), the danger button's outline (45%) and hover fill, notice outlines.
- Status: `--color-success` `#2f7d32`, `--color-warning` `#a86a00`, `--color-danger` `#b3261e` on `--color-danger-surface`; `--color-notice` with `--color-notice-text` is the information notice. `--color-inverse` ink carries the button on the accent cell, the tooltip and the knob of a switched-on toggle.
- `--transition` is a 0.1s colour change. Nothing moves except the accordion mark, which turns.

## Components

- `.ds-page` on `<body>`: sage page, family and 15px text. `.ds-page__inner` is the 1240px column of a view.
- `.ds-nav`: the two-row header. `.ds-nav__inner` (brand, badge, `.ds-nav__search` with `.ds-nav__glass`, `.ds-nav__field` and `.ds-nav__key`, `.ds-nav__actions` with `.ds-nav__account`), `.ds-nav__rule`, then `.ds-nav__sub` with `.ds-nav__links` of `.ds-nav__entry` holding a `.ds-nav__link` (`is-current` by hand) and `.ds-nav__meta` with a `.ds-nav__dot`. `.ds-nav__extra` marks what hides on a phone.
- `.ds-brand`: the site's mark and name at the left of the first row. `.ds-brand__mark` is a 26px tile dressed like the primary button (`--fill-button`, `--color-button-text`, outline, `--radius-control`, `--shadow-control`) holding the mark as inline SVG; `.ds-brand__name` is the name in the heading family at 17px. The name `chronoskin` and its mark are placeholders: the installing project puts its own name and logo here.
- `.ds-hero`: the largest cell of the home grid, always together with `.ds-grid__cell--w3` and `--tall`. A badge, `.ds-hero__title`, `.ds-hero__lead` at the top; `.ds-hero__actions` with two large buttons and `.ds-hero__note` with a row of badge links at the bottom.
- `.ds-grid`: the bento. `.ds-grid__cell` with spans `--w2`, `--w3`, `--full`, `--tall`; `--stack` (16px between children), `--row` (children side by side), `--flush` (no padding, clipped, for the stats strip and the dialog stage), fills `--fill-1` to `--fill-4`, `--accent` (the closing call), `--bare` (outline only). Inside: `.ds-grid__head` (title left, badge right), `.ds-grid__icon` with `.ds-grid__glyph`, `.ds-grid__title`, `.ds-grid__text`, `.ds-grid__copy`, `.ds-grid__visual` and `.ds-grid__actions` (both pushed to the bottom), `.ds-grid__more` (the cell's one link). `.ds-grid--auto` is for cells that hold components.
- `.ds-layout`: main column and rail of an inner page, `.ds-layout__main` and `.ds-layout__rail`. `.ds-section__label` with `.ds-section__title` is a small heading on the bare page above a further run of cells.
- `.ds-page-header`: the head of an inner page, on the bare page with a rule under it. `.ds-page-header__copy` holds `.ds-page-header__title` (an h1, optionally with a badge) and `.ds-page-header__text`; `.ds-page-header__actions` sits at the right. Put a `.ds-breadcrumb` above it; never use it with a `.ds-hero`.
- `.ds-stat`: one figure: `.ds-stat__label` over `.ds-stat__value` (36px, with an optional `.ds-stat__unit`) over a 12px `.ds-stat__note` (`--success`, `--warning`, `--danger`). `.ds-stats` is the row: four figures in one `--full --flush` cell, parted by upright rules. `.ds-stat--display` raises a single figure to the display size.
- `.ds-prose`: long text in a full-width cell; styles h1 to h3, paragraphs, lists, `strong`, `code` and links.
- `.ds-link`: underlined inline link with `is-visited`, `is-hover`, `is-active`, `is-focus`; `--quiet` for grey links without underline. `.ds-links` lays several out in a row.
- `.ds-button`: green primary with an ink lip. `--secondary` cream with an ink outline, `--danger` red text on `--color-danger-surface` (never a solid red block), `--inverse` ink for the accent cell, `--large`, `--small`; states `is-hover`, `is-active`, `is-focus`, `is-disabled`. Group with `.ds-buttons`.
- `.ds-badge`: 12px pill with an outline. `--accent`, `--new`, `--count` (a number in a sidebar row), optional `.ds-badge__dot`; status labels `--success` (stable, live), `--warning` (beta, unused), `--danger` (leaving, failed). `.ds-chips` wraps several.
- `.ds-endpoint`: a method in `.ds-endpoint__verb` and a path, in mono; used in tables and cell heads.
- `.ds-form`: two-column grid of `.ds-form__field` (`--full`) with `.ds-form__label`, `.ds-form__control` (`--area`; `is-invalid`, `is-focus`, `is-disabled`), `.ds-form__hint`, `.ds-form__error`, `.ds-form__check` with `.ds-form__checkbox`, and `.ds-form__actions`.
- `.ds-table`: a 14px-rounded box with a `--fill-bar-alt` head and faint row rules, inside `.ds-table__scroll`; `.ds-table__link` on a row's first cell, `.ds-table__num`, `.ds-table__muted`, `.ds-table__extra` (hidden on a phone), `tr.is-current`.
- `.ds-list`: rule-separated `.ds-list__item` with `.ds-list__title`, `.ds-list__text`, `.ds-list__meta`; `.ds-list__row` puts a badge beside the title. `.ds-rows` with `.ds-rows__item` is the plain label-and-control list used for switches.
- `.ds-panel`: the small 14px card inside a cell, `.ds-panel__title`, `.ds-panel__body`, `.ds-panel__meta`; `.ds-panels` puts two side by side.
- `.ds-tabs`: text labels on a rule, `.ds-tabs__tab`, the current one (`is-current`) underlined in `--color-link`. `.ds-toolbar` holds it above the main column's grid.
- `.ds-sidebar`: a rail cell with `.ds-sidebar__title` and a `.ds-sidebar__list` of `.ds-sidebar__link` (`is-current` filled): topics of a list, the contents of a page.
- `.ds-notice`: information box with `.ds-notice__title` and `.ds-notice__text`; `--error` for the red variant. `.ds-notices` stacks them.
- `.ds-pagination`: `.ds-pagination__link` items; `is-current` outlined, `is-disabled` greyed.
- `.ds-breadcrumb`: 14px trail of `.ds-breadcrumb__item` parted by slashes, the last `is-current`; the first thing of an inner view.
- `.ds-dialog`: 460px box with `.ds-dialog__bar` (`.ds-dialog__title`, `.ds-dialog__close` with `.ds-dialog__glyph`), `.ds-dialog__body`, `.ds-dialog__actions`, shown on a `.ds-stage` filled with `--color-overlay` inside a `--flush` cell.
- `.ds-empty`: centred `.ds-empty__title`, `.ds-empty__text` and one secondary button inside a `--bare` cell.
- `.ds-code`: a mono sample in a cell, with `.ds-code__key` and `.ds-code__str`. `.ds-checks` with `.ds-checks__item` and `.ds-checks__mark` is a ticked list; `.ds-chart` with `.ds-chart__bar` a small bar chart.
- `.ds-menu`: a drop-down under its trigger, opened by hover or keyboard focus (`is-open` by hand). `.ds-menu__list` of `.ds-menu__item` with an optional `.ds-menu__hint`, parted by `.ds-menu__rule`; `.ds-menu__caret` in the trigger; `--end` aligns it to the right edge. Used under Libraries and under the account avatar.
- `.ds-accordion`: `<details>` as `.ds-accordion__item` with a `.ds-accordion__head` summary and `.ds-accordion__body`; questions under a guide, parameters of an endpoint.
- `.ds-tooltip`: a focusable term with `.ds-tooltip__mark` and `.ds-tooltip__tip`, shown on hover or focus (`is-open` by hand) on `--fill-inverse`. For one sentence of explanation, never for an action.
- `.ds-carousel`: `.ds-carousel__track` of `.ds-carousel__slide` with scroll snap, one showing; `.ds-carousel__foot` holds `.ds-carousel__dots` (`.ds-carousel__dot`, `is-current`) and `.ds-carousel__arrows` of `.ds-carousel__arrow` with a `.ds-carousel__glyph`. Used for what is new.
- `.ds-switch`: a label around `.ds-switch__input` (a checkbox) and `.ds-switch__track`; on is `--color-accent-alt` with an ink knob.
- `.ds-progress`: `.ds-progress__label` with `.ds-progress__value` over `.ds-progress__bar` with `.ds-progress__fill` (`--1` to `--5` set the width; `--success`, `--warning` the colour). Uptime, quota, place in a reading path.
- `.ds-avatar`: initials on a disc, `--accent`, `--large`; `.ds-avatars` overlaps several; `.ds-person` puts one beside `.ds-person__copy` with `.ds-person__name` and `.ds-person__meta`.
- `.ds-footer`: one line under a rule: `.ds-footer__inner` with `.ds-footer__brand`, a `.ds-footer__list` of `.ds-footer__link` and `.ds-footer__text`.

## Never

- `box-shadow-blur <= 0px`: shadows are hard ink lips; nothing is blurred.
- `border-width <= 3px`: outlines are 2px; 3px is only the underline of the current section and tab.
- `border-radius >= 5px`: no box is square; the smallest radius is half a control's.
- `border-radius <= 22px`: the cell is the roundest box.
- `font-families <= 3`: a body sans, a heading sans and the mono.
- `font-size >= 12px`: the smallest text is 12px.
- `font-size <= 50px`: the hero title is the largest text.
- `font-weight <= 700`: no black weights.
- `text-shadow = none`: no text shadows.
- `gradient-fills <= 5%`: fills are flat; the only gradient is the fade in the corner of the hero.
- `uppercase-text <= 2%`: capitals only in avatar initials.
- `line-height <= 1.7`: body text sits at 1.55.
- `block-gap <= 40px`: cells sit 14px apart and nothing stands further than 40px from its neighbour.
- `content-width <= 1240px`: content stays in the centred column.
- `animation = none`: nothing animates.

## Extending

Derive a new component from the nearest one in the specimen. A new way into the docs is a `.ds-grid__cell` whose span keeps its row uneven, with a title, one caption and either a `.ds-grid__more` link or one small visual; a new box inside a cell is a `.ds-panel`; a new label is a `.ds-badge`; anything that belongs beside the working column goes in the rail as a small cell. Use tokens only: fills are `--fill-panel`, a `--color-surface*` or a flat `--color-fill-*`, outlines are `--border-width` in `--color-border`, corners `--radius-page` for cells, `--radius-panel` inside them and `--radius-control` for controls, spacing from the `--space-*` scale. Keep one subject per cell, part cells with the gap and never with a rule, and do not add a blurred shadow, a gradient fill or a second accent.
