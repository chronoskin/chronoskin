# Tonal component kit docs, 2025

## Summary

A documentation and component gallery site in the tonal language of 2021 to 2026: a top app bar with a pill search field, a navigation drawer whose current page sits on a tonal pill, and the content on one rounded sheet. The palette is a warm light scheme from an orange key colour, with brown-tinted neutrals, a sand secondary and an olive tertiary. Headings are a serif (Roboto Serif) over a plain sans body at 17px, corners run from 16 to 32px, and cards carry one faint shadow at most.

## Layout

- Designed at 1440px; `--size-page` is 1440px and the page does not stretch beyond it. `.ds-nav` is the top app bar (`--size-bar` 64px) across both columns; below it the `.ds-sidebar` drawer (`--size-drawer` 296px, sticky) on the left and the views on the right. Each view is one `.ds-main` sheet (`--color-canvas`, `--radius-page`, padding `--space-6` by `--space-7`).
- The bar holds `.ds-brand` (as wide as the drawer), the `.ds-search` pill (up to `--size-search` 520px) and the links at the right, each a pill.
- Home view: a `.ds-banner` on `--fill-bar-alt`, `.ds-hero` (two cards `--space-2` apart: the title card at 7 parts on `--color-surface-alt`, the artwork card at 5 on `--fill-accent`), `.ds-stats` `--space-2` under it, then `.ds-section` blocks: a `.ds-toolbar` (section title and filter chips) over the `.ds-grid` of six component cards in three columns, and `.ds-columns` with the changes list at 3 parts and a tonal card at 2. `.ds-footer` is an inverted rounded block under the sheet.
- Spacing: `--space-1` 4px, `--space-2` 8px (between hero cards and stat tiles, inside a component card), `--space-3` 12px, `--space-4` 16px (between component cards, sheet to window edge), `--space-5` 24px (card padding, bar padding, columns gap), `--space-6` 36px (sheet top padding), `--space-7` 56px (sheet side padding, hero card padding, above a section), `--space-8` 80px.
- Heights: buttons `--size-control` 40px; fields `--size-field` 56px; drawer items, the search bar and table rows `--size-row-dense` 56px; list rows and the tall tabs `--size-row` 64px; previews `--size-preview` 152px. `--size-measure` 720px caps running text, `--size-toc` 220px is the contents column, `--size-dialog` 360px.
- Inner pages (every view but the home view) have no hero and no banner. The bar and drawer stay; the sheet opens with `.ds-breadcrumb` (where the view is below another), then `.ds-page-header`, then `.ds-tabs` across the sheet, then the content: `.ds-article` (text with `.ds-guides`, notices and pagination, beside the `.ds-toc` contents column); a `.ds-toolbar` over the table; or `.ds-columns` with the form at 3 parts and a `.ds-stack` at 2.
- Views: `components` (home) is the gallery. `guidance` is one component's guidance article with do and do not cards. `tokens` is the token table with status badges. `feedback` is the request form beside the discard dialog and the open requests. `search` shows an empty result.
- Below 1100px the article and columns become one column, the contents column is dropped and the grid has two columns. Below 800px the bar (`.ds-nav`) wraps to a brand row, the full-width search pill and a full-width `.ds-nav__links` row of spread `.ds-nav__link` items that scrolls sideways, the current one keeping `is-current`, and a menu in the bar hangs its list from the right edge; the drawer becomes one row of outlined chips that scrolls sideways; hero, stats, grid and guides stack; the tabs scroll sideways; `.ds-table__extra` cells are hidden.

## Typography and colour roles

- Two families: `--font-heading` is Roboto Serif (then Noto Serif, Georgia) for every title, the brand name and the stat figures; `--font-body` and `--font-ui` are Noto Sans. `--font-mono` is Noto Sans Mono.
- Sizes: `--text-base` 17px at `--line-body` 1.55; `--text-small` 13px; `--text-ui` 14px; `--text-large` 21px (leads); `--text-display` 52px at `--line-display` 1.15 with `--display-tracking` -0.5px; `--text-h1` 34px, `--text-h2` 26px, `--text-h3` 21px at `--line-heading` 1.3.
- Weights: display 400, headings 500, `--weight-ui` 600, `--weight-bold` 700. `--ui-transform` is `none`.
- Links are not underlined at rest, underlined on hover.
- The era's rule for text on fills, which every layout and every token set of this era keeps: on `--fill-page`, `--color-canvas`, `--color-surface`, `--color-surface-alt`, `--color-surface-strong` and on all four `--color-fill-*` blocks, text is `--color-text`, titles `--color-heading` and captions `--color-text-muted`; the fills are container tones, so the ordinary text colours stay readable on them in a light and in a dark scheme. On `--fill-bar` text is `--color-bar-text`; on `--fill-bar-alt` it is `--color-bar-alt-text`; on `--fill-inverse` it is `--color-inverse-text`; on `--fill-accent` it is `--color-accent-text`; on `--fill-button` it is `--color-button-text`; on `--fill-button-secondary` it is `--color-button-secondary-text`; on `--color-notice` it is `--color-notice-text`; on `--color-danger-surface` the body is `--color-text` and the title and icon are `--color-danger`. A solid block of `--color-danger`, `--color-success`, `--color-warning` or `--color-accent-alt` carries `--color-canvas` as its text colour, because the status colours are dark in a light scheme and pale in a dark one, always opposite to the canvas.
- Tonal roles: `--color-page` is the surface container behind everything; the bar, rail or drawer and the table head are on `--fill-bar`, which the surface part sets to a neighbouring container tone, so each surface set shows there; `--color-canvas` is the plain surface of the rounded content sheet; `--color-surface` (container low) fills cards and list cards, `--color-surface-alt` (container high) stat tiles and dialogs, `--color-surface-strong` (container highest) tonal cards, inline code and hovered rows. `--color-accent` is the primary container: hero card, floating action button, default badge, the disc of an empty state. `--color-button-secondary` is the secondary container: tonal buttons and every pill indicator behind a current navigation item, chip or tab. `--color-accent-alt` is the tertiary colour, used solid for "new" markers. `--color-fill-1` to `--color-fill-4` are the primary, secondary and tertiary containers and the highest surface container.
- `--color-link` is the primary colour: links, outlined and text button labels, the current tab, list markers. Visited links take the tertiary `--color-link-visited`. `--color-border` (outline variant) draws card outlines, chips and dividers, `--color-border-strong` (outline) the outlined button and the off switch, `--color-border-muted` the dividers inside lists and tables. `--color-input-border` outlines fields; `--color-focus` is the focus line.
- Status badges are tonal: `--color-success`, `--color-warning` or `--color-danger` mixed at 22% into `--color-canvas`, with `--color-text` and a dot in the full status colour.
- Shape: every button, badge, search bar, navigation indicator and pagination link uses `--radius-pill`; fields and chips use `--radius-control`; cards, notices, tables and the floating action button use `--radius-panel`; the content sheet, hero, grid tiles, dialogs and empty states use `--radius-page`.
- Elevation is shown by tone, not shadow. `--shadow-panel` and `--shadow-control` sit on cards and buttons and are `none` or one faint layer; `--shadow-control-hover` is on hovered buttons and on the search bar, `--shadow-control-pressed` on a pressed button and the selected chip, `--shadow-dialog` on dialogs and the floating action button. `--shadow-text` is on the hero title and button labels (`none`). `--backdrop-blur` blurs what lies behind the dialog scrim (`0px`).
- One-off values built with `color-mix()` because the vocabulary has no token: an 8% or 10% wash of the label colour for hovered tonal, outlined, text and icon buttons, chips and navigation items; the 22% status tints; the danger colour mixed with 12% of `--color-text` for a hovered danger button.

## Components

- `.ds-page` on `<body>`: the shell grid on `--fill-page`.
- `.ds-nav`: the top app bar on `--fill-bar`, holding the brand, the search bar and `.ds-nav__links` of pill-shaped `.ds-nav__link` items; the current one is filled with `--fill-button-secondary`.
- `.ds-sidebar`: the navigation drawer, written once outside the views. `.ds-sidebar__title` heads a group (a rule separates groups), `.ds-sidebar__list` holds 56px pill-shaped `.ds-sidebar__link` rows with an icon and an optional `.ds-sidebar__count`. A link marked `--page` is the page of a view and is put on the secondary container while that view is open; `is-current` does the same by hand.
- `.ds-main`: the content sheet of a view. `.ds-banner`: a pill on `--fill-bar-alt` with a badge, `.ds-banner__text` and a text button.
- `.ds-hero`: two cards. `.ds-hero__copy` holds `.ds-hero__title`, `.ds-hero__lead` and a `.ds-buttons` group on `--color-surface-alt`; `.ds-hero__art` holds `.ds-shapes` on `--fill-accent`.
- `.ds-stat` here is a low row, figure and label on one baseline, on `--fill-panel`.
- `.ds-toolbar`: a title or chips on the left and chips or badges on the right.
- `.ds-grid`: three columns of component cards. Each `.ds-grid__cell` is a filled card with a `.ds-grid__preview` (`--2`, `--3`, `--4` pick the fill) above `.ds-grid__title` and `.ds-grid__text`. Previews are drawn with `.ds-mini` stand-ins: a filled pill, `--tonal`, `--outlined`, `--chip`, `--fab`, and `--card` with `.ds-mini__line` (`--short`) bars.
- `.ds-tabs`: here one row of tall equal `.ds-tabs__tab` pills on `--fill-panel`; `is-current` is on the secondary container.
- `.ds-article`: text column and contents column. `.ds-toc` with `.ds-toc__title`: the contents list. `.ds-guides` holds two `.ds-guide` cards: `.ds-guide__preview` ruled underneath in `--color-success`, `.ds-guide__body` with a `.ds-guide__label`; `--dont` turns the rule and label to `--color-danger`.
- `.ds-swatch` (`--2`, `--3`, `--4`, `--button`, `--canvas`): a colour disc in a table cell.
- `.ds-footer`: a rounded block on `--fill-inverse` with `.ds-footer__legal` and `.ds-footer__links` of `.ds-footer__link`.
- `.ds-brand`: the site's mark and name, one link. `.ds-brand__mark` is a round tile dressed like the filled button (`--fill-button`, `--color-button-text`, `--radius-pill`, `--shadow-control`, a transparent `--border-width` border) holding the mark as inline SVG; `.ds-brand__name` is the name in the heading font. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo.
- `.ds-page-header`: the head of an inner view, unboxed on the sheet. `.ds-page-header__body` holds `.ds-page-header__title` (`--text-h1`) and `.ds-page-header__text` (one muted line); `.ds-page-header__actions` sits at the right with one filled button at most.
- `.ds-prose`: long-form text; styles h1 to h3, paragraphs, lists, `strong`, inline `code` and links. `.ds-prose__lead` is the larger opening paragraph.
- `.ds-link`: inline link with `is-visited`, `is-hover`, `is-active`, `is-focus`; `--quiet` is text-coloured, `--more` carries a chevron. `.ds-links` lays several out in a row.
- `.ds-button`: the filled pill button. `--secondary` is the tonal button, `--outlined` has a `--color-border-strong` outline and a link-coloured label, `--text` has no container, `--danger` is filled with `--color-danger` (with `--text` it is a danger-coloured text button for dialogs), `--large` is the tall form for a hero or an order form, `--icon` is a round icon button, `--fab` is the floating action button, a rounded square in `--fill-accent`, and `--on-inverse` is the button on an inverted block. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. Group with `.ds-buttons`. Never more than one filled button in a group.
- `.ds-icon`: stroked inline SVG in the current text colour; `--small` inside buttons, chips and tabs.
- `.ds-chip`: an outlined 32px chip; `is-current` fills it with the secondary container. Group with `.ds-chips`.
- `.ds-form`: stacked fields. `.ds-form__field` wraps a small `.ds-form__label` above a `.ds-form__control` (56px, outlined, `--fill-input`; `--area` for a textarea; wrapped in `.ds-form__select` for a select); `.ds-form__row` puts two fields side by side. `is-focus` shows the focus ring, `is-invalid` on the field plus `.ds-form__error` shows an error, `.ds-form__hint` is help text. `.ds-form__check` with `.ds-form__checkbox` is a checkbox row; `.ds-form__check--between` with `.ds-form__check-text` and a `.ds-switch` is a setting row. `.ds-form__actions` holds the buttons.
- `.ds-switch`: a checkbox drawn as a pill track; the thumb is small and outline-coloured when off, larger and in `--color-button-text` on a `--fill-button` track when on.
- `.ds-table`: inside `.ds-table-wrap`, an outlined rounded container. The head row is on `--fill-bar`; rows are divided by `--color-border-muted`; `.ds-table__num` right-aligns numbers, `.ds-table__name` sets a code name, `.ds-table__extra` marks cells a phone can do without; `tr.is-current` is tinted.
- `.ds-list`: rows of `.ds-list__item`, each with an optional `.ds-list__icon` disc (`--1`, `--3`, `--4` pick the fill), `.ds-list__body` with `.ds-list__title` and `.ds-list__text`, and `.ds-list__meta` or a button on the right. `--cards` turns every row into its own filled card, `--space-2` apart, and `is-current` puts the selected one on the secondary container.
- `.ds-panel`: a filled card with `.ds-panel__title`, `.ds-panel__body` and optional `.ds-panel__head` and `.ds-panel__actions`. `--outlined` is the outlined card on the canvas, `--tonal` the stronger tint, `--accent` the primary container. A card is filled or outlined, never both.
- `.ds-stat`: one summary figure, a tile with `.ds-stat__value` and `.ds-stat__label`. Three or more sit in a `.ds-stats` row.
- `.ds-badge`: a small pill in `--fill-accent`; `--new` is the solid tertiary colour, `--count` a solid danger count, and `--success`, `--warning`, `--danger` are the tonal status labels with a dot. Group with `.ds-badges`.
- `.ds-notice`: a tonal card with a leading icon, `.ds-notice__body` and `.ds-notice__title`; `--error` is on `--color-danger-surface` with a danger-coloured title. Group with `.ds-notices`.
- `.ds-pagination`: a row of round `.ds-pagination__link`; `is-current` is filled like the primary button, `is-disabled` is grey; previous and next are chevron icons.
- `.ds-breadcrumb`: a muted trail of `.ds-breadcrumb__item` with `.ds-breadcrumb__link`, separated by drawn chevrons; the last carries `is-current`.
- `.ds-dialog`: a `--color-surface-alt` sheet with `--radius-page` corners, an optional `.ds-dialog__icon`, `.ds-dialog__title`, `.ds-dialog__body` and right-aligned text buttons in `.ds-dialog__actions`. It sits on `.ds-overlay`, the scrim.
- `.ds-empty`: a centred block with `.ds-empty__mark` (a rounded square in `--fill-accent` holding an icon), `.ds-empty__title`, `.ds-empty__text` and one button.
- `.ds-search`: the search bar, one pill on `--fill-input` with an icon and `.ds-search__text`. `.ds-avatar`: a disc with initials. `.ds-snackbar`: a short message on `--fill-inverse` with `.ds-snackbar__action`. `.ds-shapes` with `.ds-shape--1` to `--4`: the circle, rounded square and pill drawn as artwork.
- `.ds-section` with `.ds-section__head`, `.ds-section__title` and `.ds-section__label` titles a block; `.ds-columns` (`--even`) puts two blocks side by side and `.ds-stack` piles cards in one of them.
- `.ds-menu`: a `<details>` drop-down. `.ds-menu__button` is its tonal pill (`--dense` at chip height; `.ds-menu__label` is the text); `.ds-menu__list` is a `--color-surface-alt` sheet of `.ds-menu__item` rows, the chosen one `is-current`. `.ds-menu--end` hangs the sheet from the right edge, `--end-wide` only above a phone. Use it for a choice of one among a few: sort order, version, speed.
- `.ds-accordion`: a group of `<details>`; each `.ds-accordion__item` is a tonal row with `.ds-accordion__summary` (a title, a `.ds-accordion__hint` line, a chevron) and `.ds-accordion__body`. Use it for questions and for details that most readers skip.
- `.ds-progress`: a thick track in two parts with a gap, `.ds-progress__bar` in `--fill-button` and `.ds-progress__rest` with a stop dot at its end; `--25`, `--40`, `--60`, `--75` set the share. `.ds-progress__label` is the line above it and `.ds-progress-group` stacks several. Use it for a known share of a whole, never for waiting.
- `.ds-avatar`: a disc with a person's initials; add it to a list's leading disc when the row is about a person.

## Never

- `gradient-fills = 0%`: every fill is one flat tone.
- `text-shadow = none`: no text has a shadow.
- `uppercase-text = 0%`: labels are sentence case.
- `border-width <= 4px`: outlines are 1px; the rule under a guidance preview is 4px.
- `box-shadow-blur <= 10px`: shadows are one or two small warm layers.
- `font-weight <= 700`: bold is for strong text only.
- `font-weight >= 400`: no thin text.
- `font-size >= 13px`: the smallest labels are 13px.
- `font-size <= 52px`: the hero title is the largest text.
- `font-families <= 3`: a serif for headings, a sans for text and the mono.
- `line-height <= 1.7`: body text is 17px on a 26px line.
- `underlined-links <= 10%`: links are marked by colour.
- `border-radius <= 32px`: containers stop at 32px; anything rounder is a full pill.
- `row-gap <= 8px`: list rows touch and are divided by hairlines.
- `block-gap <= 56px`: blocks on the sheet are at most 56px apart.
- `content-width <= 80%`: the drawer takes the left fifth of the window.

## Extending

Derive a new component from the nearest one in the specimen. Anything that holds content is a `.ds-panel`: filled with a container tone, or outlined on the canvas, with `--radius-panel`. Anything pressed is a `.ds-button` variant and stays a full pill. A selected state is the secondary container (`--fill-button-secondary` with `--color-button-secondary-text`) on a pill or on the item itself. Raise something by moving it to the next surface tone, not by adding a shadow. Use tokens only, take heights and gaps from the `--size-*` and `--space-*` scale, and follow the rule for text on fills in "Typography and colour roles". Never add a gradient, an uppercase label, a square corner on a container or a second filled button beside the first.
