# Material tabbed handbook, 2016

## Summary

A documentation page in the paper-and-ink manner where the whole head of the page is one tall band of the primary colour: a toolbar row, a centred title with one line of lead text and a raised button, and a centred row of tabs along the band's bottom edge. Everything below is a single centred column of white sheets on a pale page, with no drawer and no second column; small captions sit on the page between the sheets. Component libraries, handbooks and product help pages built on the 2014 guidelines looked like this until about 2018.

## Layout

- The page is fluid: `--size-page` is `100%` and the band and footer span the window. Designed at 1440px. The content is one centred column capped at `--size-content` (880px); running text is capped at `--size-measure` (520px).
- Order on the home page: `.ds-nav` (`--size-bar` 64px: menu icon and brand at the left, a few links at the right), then a `.ds-band` in the same fill holding `.ds-hero` (centred title, lead and one button) and `.ds-tabs--band` (centred, the current tab marked by a 2px line in `--color-bar-text` on the band's bottom edge). Then `.ds-main`, the column, then the one-row `.ds-footer`.
- `.ds-main` is a vertical stack with `--space-4` (24px) between blocks. Every block is a sheet: `.ds-panel`, `.ds-grid` tiles, `.ds-stats` cards, `.ds-notice`. A group of sheets is introduced by `.ds-subhead`, a small caption set directly on the page and indented to the sheet's text edge. Sections are never separated by rules on the page, and nothing sits beside the column.
- The long article is one `.ds-panel--article` sheet with `--space-6` (48px) padding. Inside it `.ds-article` puts the `.ds-prose` text at the left and the `.ds-sidebar` contents block (`--size-sidebar` 200px) at the right. This contents block is the only side element the layout has.
- Spacing is on an 8px grid: `--space-1` 4px (helper text, pagination gap), `--space-2` 8px (between flat buttons, card action padding), `--space-3` 16px (button padding, gap between grid tiles and stat cards, tile text padding), `--space-4` 24px (gap between sheets, sheet side padding, bar side padding), `--space-5` 32px (top of the column, page header side padding, under the hero lead), `--space-6` 48px (hero padding, article sheet padding, overlay padding), `--space-7` 64px (spare step), `--space-8` 96px (bottom of the column).
- Heights: toolbar and card head `--size-bar` 64px; tabs and table rows `--size-row` 48px; avatars and the menu button `--size-row-dense` 40px; buttons, fields and nav links `--size-control` 36px; floating button `--size-fab` 56px; two-line list rows 72px. Other sizes: `--size-tile` 112px (colour block of a grid tile), `--size-overlap` 72px, `--size-stripe` 4px, `--size-icon` 24px, `--size-check` 18px, `--size-badge` 22px, `--size-dialog` 440px, `--size-area` 96px.
- Inner pages (a resource, settings, a record and its table) have no `.ds-hero` and no tabs in the band. Under the `.ds-nav` row the band takes `.ds-band--inner`: it is empty and continues the bar's fill for `--size-overlap` plus `--space-5`. `.ds-main--overlap` pulls the column up by `--size-overlap`, so the first sheet lies over the band. That first `.ds-panel` is the head of the page: `.ds-breadcrumb` at its top, `.ds-page-header` (title, one grey line, one raised button at the right) and `.ds-tabs` along its bottom edge, separated by a hairline. Below it, `--space-4` apart, come `.ds-stats` (separate small cards), then a `.ds-panel` with a `.ds-panel__head` and the `.ds-table`, or, on an edit screen, a `.ds-panel--form` with the form and its buttons in `.ds-panel__actions--end`. A chapter page puts its breadcrumb at the top of the `.ds-panel--article` sheet. The column stays single; there is no drawer on inner pages either.
- Views: the specimen is one site with five views. `handbook` is the home view: hero and band tabs, the grid of chapter tiles, a card and the list of recent edits with its pagination. `chapter` is one handbook chapter: the article sheet with its contents block, two notices and a sheet of related links. `api` is a resource page: head sheet with breadcrumb, page header and tabs, the stat cards, and the table with its status key. `webhook` is the edit screen for one table row: head sheet with the page actions, the form sheet and the confirmation dialog. `starred` is the reader's starred chapters, shown empty, above the list of most starred chapters.
- Below 900px the contents block drops under the text, grid tiles become one column, the page header stacks and the band tabs scroll. Below 600px the toolbar wraps: `.ds-nav__links` becomes a full-width second row under the menu icon and brand, each item an equal cell with its `.ds-nav__link` centred (the current one keeps `.is-current`), the tabs scroll sideways, the stat cards sit two by two, and table cells marked `.ds-table__extra` are hidden.

## Typography and colour roles

- `--color-bar` deep pink `#c2185b` with white `--color-bar-text`: navigation bar and header band. `--color-bar-alt` indigo `#3f51b5` is the secondary toolbar; `--color-inverse` `#424242` with `#fafafa` text is the footer.
- Indigo `#3f51b5` is the colour of action: `--color-button`, `--color-link`, `--color-heading-alt` (section and drawer titles). Hover darkens to `--color-link-hover` `#284888`; visited links are purple `#9c27b0`.
- `--color-accent` cyan `#00bcd4` with dark `--color-accent-text` `#212121`: the floating button and count badges. Tab indicators do not use it: on the band the line is `--color-bar-text`, on a sheet `--color-link`, so it stays readable under every palette. `--color-accent-alt` and `--color-focus` are bright pink `#ff4081`.
- Page `#fafafa`, sheets white, `--color-surface-alt` `#f5f5f5`, `--color-surface-strong` and `--color-border` `#dddddd`. Body text is `#424242`, headings `#212121`, secondary text `#757575`.
- Block fills: pink `#e91e63`, orange `#ff9800` (dark text), indigo `#3f51b5`, slate `#37474f`. Status: `--color-danger` `#e91e63`, `--color-success` teal `#009688`, `--color-warning` `#ff9800`; the notice is pale orange `#ffe0b2` with `#212121` text.
- `--font-heading` and `--font-ui` are the monospace stack (Roboto Mono, Menlo): headings, brand name, buttons, navigation, tabs, drawer. `--font-body` stays Roboto for running text.
- Headings are uppercase (`--heading-transform`) at regular weight 400 and modest sizes: `--text-display` 40px, `--text-h1` 30px, `--text-h2` 20px, `--text-h3` 16px, `--line-heading` 1.4.
- `--text-base` 16px on a 24px line; `--text-large` 18px; `--text-ui` 13px uppercase at `--weight-ui` 400; `--text-small` 12px. `--weight-bold` 500 is the heaviest weight.
- Links in running text are underlined in every state (`--link-decoration`, `--link-decoration-hover`); quiet links, titles and navigation are not (`--link-decoration-quiet` `none`).
- Sheets are flat and outlined: `--shadow-panel` is a 1px ring in `--color-border` with no blur; nothing rests above the page. `--fill-page` is the sheet colour (`--color-surface`), so page and sheets are one tone and only the outline separates them.
- `--radius-control`, `--radius-panel` and `--radius-page` are 4px; `--radius-pill` is 16px, so badges and chips are pills and pagination cells are rounded squares.
- Buttons rest flat with a 1px inner ring (`--shadow-control`), lift to a soft two-layer shadow on hover (`--shadow-control-hover`, 4px blur, also the floating button), and only the dialog floats (`--shadow-dialog`, three layers up to 14px blur). A hovered button lightens: `--fill-button-hover` is the button colour mixed with 10% white.
- Text fields are filled: `--fill-input` is the text colour at 6% over the sheet, with the top corners at `--radius-control`, a side inset of twice that radius and the bottom line. `--focus-ring` is a 3px halo of the focus colour at 40%. `--transition` is 0.28s on the standard curve.
- Every fill is one flat colour; `--border-width` 1px for dividers, `--border-width-strong` 2px for indicators and focused fields.
- Text on fills follows fixed pairs: `--color-bar-text` on the band, on `--color-fill-1` and on `--color-fill-3`; `--color-accent-text` on the accent, on `--color-fill-2` and on the warning badge; `--color-inverse-text` on the footer and on `--color-fill-4`; `--color-button-text` on the primary and danger buttons and on the new, success and danger badges. Links, muted text and headings are used only on the page and on white sheets.
- One-off values built with `color-mix()` or `calc()` because the vocabulary has no token: 70% and 87% of the bar text colour for band tabs, toolbar links and the hero lead, a 16% wash of it behind the current toolbar link, 66% of the quiet link colour for contents links, 70% of the inverse text in the footer, a 20% grey wash behind a hovered flat button, and line heights 1.18, 1.2, 1.43 and 1.6.

## Components

- `.ds-page` on `<body>`: page fill, body family and size.
- `.ds-band`: the coloured head of the page under the nav; wraps the hero and the band tabs and carries their shadow. `--inner` is the short form for inner pages.
- `.ds-nav`: the toolbar row inside the band. `.ds-nav__menu` is a round icon button, the summary of a `.ds-menu` that lists the views; `.ds-brand` follows it; `.ds-nav__links` > `.ds-nav__item` > `.ds-nav__link` sit at the right, the current one (`is-current`) on a faint wash of the bar text colour.
- `.ds-brand`: the site's mark and name at the left of the toolbar, one link. `.ds-brand__mark` is a `--size-brand` (28px) tile dressed like the primary raised button (`--fill-button`, `--color-button-text`, `--radius-control`, `--shadow-control`) with a `--border-width` ring of `--color-bar-text` at 70%, which keeps it apart from the band when a palette's button and bar colours are close, holding the mark as inline SVG in the current colour; `.ds-brand__name` is the name at `--text-large` in the heading font with `--weight-heading`, `--heading-tracking` and `--heading-transform`, in `--color-bar-text` because the band is filled with `--fill-bar`. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo.
- `.ds-hero`: centred `.ds-hero__title`, `.ds-hero__lead` and one raised `.ds-button--secondary`. No floating button and no image.
- `.ds-tabs`: a row of `.ds-tabs__tab`; `is-current` has a 2px line, `--color-link` on a sheet and `--color-bar-text` on the band. The plain form closes a sheet (hairline above); `--band` is the centred row at the bottom of the band, in the bar text colour.
- `.ds-main`: the single column; `--overlap` lifts it over an inner band. `.ds-subhead` is the caption above a group of sheets. `.ds-article` splits the article sheet into text and contents. `.ds-center` centres one block.
- `.ds-page-header`: head of an inner page, inside the first sheet. `.ds-page-header__body` holds `.ds-page-header__title` and `.ds-page-header__text`; `.ds-page-header__actions` at the right holds one raised `.ds-button`.
- `.ds-breadcrumb`: small grey trail of `.ds-breadcrumb__item` with `.ds-breadcrumb__link`, separated by slashes, at the top of the head sheet; the last carries `is-current`.
- `.ds-prose`: long-form text; styles h1 to h3, paragraphs, lists, `strong`, inline `code`, `pre` and links. `.ds-prose__lead` is the opening paragraph.
- `.ds-sidebar`: the contents block beside the article: a 4px stripe in the link colour at the left, `.ds-sidebar__title`, `.ds-sidebar__list` of `.ds-sidebar__link`; `is-current` is in the link colour. It has no fill and no shadow.
- `.ds-link`: inline link with `is-visited`, `is-hover`, `is-active`, `is-focus`; `--quiet` is text-coloured. `.ds-links` lays several out in a row.
- `.ds-button`: raised primary button. `--secondary` is the white raised button (use it on the band), `--flat` has no fill (use it inside sheets and dialogs), `--danger` is the destructive form (with `--flat`: danger-coloured text), `--fab` the accent disc with one `.ds-icon`. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. Group with `.ds-buttons`.
- `.ds-icon`: 24px stroked inline SVG in the current text colour.
- `.ds-form`: one column of fields. `.ds-form__field` wraps `.ds-form__label` and a `.ds-form__control` filled with `--fill-input` over a bottom line (`--area` for a textarea, inside `.ds-form__select` for a select); `is-focus`, `is-invalid` with `.ds-form__error`, `.ds-form__hint`, `.ds-form__check` with `.ds-form__checkbox`. The buttons go in the sheet's `.ds-panel__actions--end`, not in the form.
- `.ds-table`: 48px rows divided by hairlines, small grey head; it has no fill of its own and always sits in a `.ds-panel` under a `.ds-panel__head`. `.ds-table__num` right-aligns numbers, `.ds-table__name` sets a code name, `.ds-table__extra` marks cells a phone can do without, `tr.is-current` is tinted.
- `.ds-list`: two-line rows inside a `.ds-panel`. `.ds-list__item` holds a `.ds-list__avatar` disc (`--alt`, `--accent`), `.ds-list__body` with `.ds-list__title` and `.ds-list__text`, and `.ds-list__meta` at the right.
- `.ds-panel`: the sheet. `.ds-panel__title` and `.ds-panel__body` for a simple card; `.ds-panel__head` is a 64px title row with an optional flat button and a hairline under it; `.ds-panel__actions` is the ruled button row (`--end` right-aligns it). `--article` is the padded article sheet, `--form` and `--pad` give the body normal text colour.
- `.ds-stat`: one summary figure as its own small centred card: `.ds-stat__value` over `.ds-stat__label`. Three or more sit in a `.ds-stats` row, `--space-3` apart.
- `.ds-grid`: two columns of horizontal tiles. Each `.ds-grid__cell` (`--2`, `--3`, `--4` pick the fill) has a square `.ds-grid__media` colour block with a number or initial at the left and `.ds-grid__body` with `.ds-grid__title`, `.ds-grid__text` and `.ds-grid__actions`.
- `.ds-badge`: accent count pill; `--new`, `--chip`, and the status fills `--success`, `--warning`, `--danger`. Group with `.ds-badges`.
- `.ds-notice`: tinted block with a 4px left stripe and `.ds-notice__title`; `--error` is a white sheet in the danger colour. `.ds-notices` stacks them at column width.
- `.ds-pagination`: `.ds-pagination__link` cells, `is-current` filled, `is-disabled` grey; it sits right-aligned in the `.ds-panel__actions--end` of the sheet it pages.
- `.ds-dialog`: sheet with the deepest shadow, `.ds-dialog__title`, `.ds-dialog__body`, right-aligned flat buttons in `.ds-dialog__actions`; shown on `.ds-overlay`.
- `.ds-empty`: centred `.ds-empty__mark` disc, `.ds-empty__title`, `.ds-empty__text` and one flat button, inside a `.ds-panel`.
- `.ds-footer`: one row in the inverse colour: `.ds-footer__brand`, a `.ds-footer__list` of `.ds-footer__link` and `.ds-footer__legal` at the right. No columns.
- `.ds-menu`: a `<details>` whose `.ds-menu__button` (the `<summary>`: an icon, or the account's `.ds-avatar`) opens `.ds-menu__list`, a white sheet with the dialog shadow holding 48px `.ds-menu__link` rows. Use it for the account menu and the overflow (three dots) at the end of the bar; `--start` hangs the sheet from the left, for the menu button that carries the navigation on a phone.
- `.ds-tooltip`: on an icon-only button, with a `.ds-tooltip__text` child naming it; the small dark label shows on hover and focus (`is-open` statically). It sits below; `--left`, `--right` and `--above` move it. Use it on the floating button and bar icons, never on a button that already has a text label.
- `.ds-switch`: an on-off setting that takes effect at once: a `<span>` or `<label>` holding `.ds-switch__input` (a checkbox) and `.ds-switch__track`, which draws the track and the round thumb; `is-on` shows the on state statically. Keep the checkbox for choices that are saved with a form.
- `.ds-progress`: a `--size-progress` line with `.ds-progress__bar` (`--low`, `--high`) for work whose length is known; `--ring` is the small ring for work whose length is not. Put it inside the card or notice it reports on.
- `.ds-accordion`: expansion panels inside a card: `<details class="ds-accordion__item">` with a `.ds-accordion__head` summary (a chevron is drawn at its right) and a `.ds-accordion__body` paragraph. Use it for short questions and answers, closed by default.

## Never

- `gradient-fills = 0%`: every fill is one flat colour.
- `text-shadow = none`: no text has a shadow.
- `border-radius <= 16px`: 4px on sheets and controls, 16px on tags; nothing rounder except discs.
- `box-shadow-blur <= 14px`: outlines and one tight lift; only the dialog reaches 14px.
- `border-width <= 4px`: 1px dividers, 2px indicators, the 4px notice stripe.
- `font-weight <= 500`: emphasis is medium, never bold.
- `font-size >= 12px`: captions are the smallest text.
- `font-size <= 40px`: the mono display line stays small.
- `font-families <= 2`: one sans-serif for text, one monospace for headings, labels and code.
- `letter-spacing = 0px`: the mono face is not tracked further.
- `row-gap <= 4px`: table and list rows touch and are divided by hairlines; bullet items sit 4px apart at most.
- `block-gap <= 48px`: sheets in the column are 24px apart; nothing is spaced wider than a caption plus a gap.
- `content-width <= 66%`: one centred 880px column in a 1440px window; no drawer and no second column.

## Extending

Derive a new component from the nearest one in the specimen. New content is a new sheet in the column: a `.ds-panel` with a `.ds-panel__head` when it holds data, with a `.ds-subhead` above it when it starts a group. Never add a drawer, a second column beside the sheets, or a rule across the page; navigation between sections belongs in the band's tabs, navigation inside a page in the head sheet's tabs or the contents block. Use tokens only, pick heights and gaps from the `--size-*` and `--space-*` scale, and colour a block with one flat `--color-fill-*` and the text colour paired with it above.
