# Frutiger Aero developer guide

## Summary

A fixed-width developer guide in the violet glass variety of the Frutiger Aero look: an amethyst page with an orchid and a lilac glow, a contents rail of frosted glass on the left, and one glass sheet beside it that opens with a dark band holding a small window of sample code. Guides of this kind were common from about 2006 to 2010, when desktops grew sidebars and every platform invited people to write gadgets and widgets for them. Bars, keys and labels are smooth gels lit along the top with no hard break, panels are translucent with a bright upper edge and a soft violet shadow, corners are rounded by 8 to 20px, and the type pairs a small humanist body with light, wide headings.

## Layout

- Page width is fixed: `--size-page` = 960px, centred in a 1024px window. It does not stretch and does not reflow for phones.
- `.ds-page__wrap` is a grid of two columns of equal height, `--space-5` (18px) apart: the contents rail (`.ds-nav`, `--size-rail` 196px) on the left and `.ds-page__content` (746px) with the glass sheet `.ds-page__sheet` (90% `--color-canvas`, `--radius-page` corners). The footer band runs under both, 18px lower.
- The navigation is the rail, never a bar across the top. Top to bottom: the name on a bar of `--size-nav` (48px), the search field, the contents tree, a list of the kit's samples, and at the foot a pale strip with the kit version and its download button. The rail is as tall as the sheet.
- The home page fills the sheet edge to edge with the hero; the rest sits in `.ds-page__body` (18px padding, 28px at the bottom): the row of four figures, the six topics in three columns under a `.ds-section__title`, then two columns.
- Two columns in `.ds-page__columns`: main (462px) and a right column `.ds-page__aside` of `--size-side` (230px), 18px apart. With the rail that makes three columns on the screen; never more.
- Spacing scale: `--space-1` 2px (hairline gaps, sub-link padding), `--space-2` 4px (inside controls, between a label and its field), `--space-3` 8px (gaps between buttons, row and cell padding), `--space-4` 12px (panel and cell padding, gaps between cells and figures), `--space-5` 18px (gutter, sheet padding, gap between stacked blocks), `--space-6` 28px (hero padding, gap between sections; the largest gap).
- Fixed sizes: controls `--size-control` 26px tall, the hero buttons `--size-control-large` 36px, the brand tile `--size-mark` 26px, topic tiles `--size-tile` 34px, icons `--size-icon` 16px, the code window `--size-code` 290px, the dialog `--size-dialog` 430px.
- Inner pages drop the hero and the figures; the rail stays. In `.ds-page__body`, top to bottom: `.ds-breadcrumb`, then `.ds-page-header` across the sheet (title and one line on the left, one or two buttons on the right, a 2px rule under it), then either the two columns or a full-width block. `.ds-tabs` sit directly above the text or table they switch; under a table `.ds-table__foot` holds the key to the labels on the left and the pagination on the right; notices sit above the form; a dialog follows the columns across the sheet on its dimmed strip. Further blocks are `.ds-section` blocks 28px apart.
- Views: the specimen is one guide of four screens, each a `.ds-view` inside the sheet, with the rail and the footer written once around them. `home` is the overview (hero with the code window, figures, six topics, recent updates, the kit panel, most read). `guide` is one guide (tabs, long text with a code block, the contents of the page, the series, onward links). `reference` lists the objects (tabs, table with status labels, the key and pagination, a note on reading an entry, an empty bookmark list). `submit` is the gallery form (notices, the form with its buttons, the checks, the rules, the discarding dialog). Links are plain `href="#name"`; the rail link of the showing view becomes a gel bar by a `.ds-page:has(#name:target)` rule, the Overview link also when no view is named.

## Typography and colour roles

- Two families and a monospace: `--font-body` and `--font-ui` are Calibri (Carlito, Geneva), `--font-heading` is Corbel (Skia, Lucida Sans): the humanist faces that came with the period's desktops, with free or system stand-ins. `--font-mono` (Consolas) sets code, which a guide has a lot of.
- `--text-base` 13px at `--line-body` 1.5. `--text-small` 11px for meta lines, hints, labels, code and the footer. `--text-ui` 12px for buttons, fields, tabs and sub-links; controls are bold (`--weight-ui` 700).
- Headings are regular weight and wide: `--text-display` 32px hero title, `--text-h1` 25px page titles and figures, `--text-h2` 19px section titles, the brand name and the kit version, `--text-h3` 14px (bold) list, cell, panel and sidebar titles and rail links. `--text-large` 16px hero lead and the large button. Nothing is uppercased or tracked.
- Links are not underlined at rest and underline on hover; titles are bold (`.ds-link--title`), object names are bold monospace (`.ds-link--code`).
- `--color-page` #5a3fb0 with `--fill-page`: violet, darker at the top, with an orchid glow at the upper left and a lilac one at the right, fading to pale lavender below.
- `--color-bar` #5a3d9e with `--fill-bar` (a smooth gel: bright top, darkest at two thirds, lighter at the foot) and white `--color-bar-text` under `--shadow-text`: the rail's name bar, the current rail link, table heads, dialog titles, the second topic tile.
- `--color-bar-alt` #ddd3f6 with `--fill-bar-alt`: the figures, the tab strip, the code window's caption, the foot of the rail.
- `--color-inverse` #2b1a66 with `--fill-inverse` (deep indigo with two soft glows): the hero and the footer band.
- `--color-heading` #40288a titles and figures; `--color-heading-alt` #1585a6 (teal) panel and sidebar titles, prose h2, breadcrumb marks. `--color-text` #2f2b3a, `--color-text-muted` #6a6578.
- `--color-link` #5a2fc2, `--color-link-quiet` #5f5a80, visited #8a4a8f, hover #1a9fc8, active #c2338f.
- `--color-button` #1392c4 with `--fill-button` (aqua gel): the primary button, the brand tile, the first topic tile, the current page number. `--color-button-secondary` #e2dcf0 with `--fill-button-secondary`: every other button, pagination keys.
- `--color-accent` #b455d8 with `--fill-accent`: labels, list dots, the mark of the current sidebar item. `--color-accent-alt` #ff8a1f: the New marker and the fourth tile.
- `--color-fill-1` to `--color-fill-4` (lavender, ice, rose, cream): the right end of each topic cell's fade; never text or borders.
- Borders are 1px solid, 2px (`--border-width-strong`) under the page header, beside sidebar lists and around the dialog. `--color-border` #c9bdea, `--color-border-muted` #e4dcf5 inside a box, `--color-border-strong` #7a5fc0 for the sheet, the rail, the dialog and the current tab.
- `--color-notice` #e8f6fb for notices; `--color-danger` #c7283a on `--color-danger-surface` #fde8ea for errors; `--color-success` #3a9c3a and `--color-warning` #eda400 fill the status labels.
- Surface: `--radius-control` 8px, `--radius-panel` 12px, `--radius-pill` 14px, `--radius-page` 20px. `--shadow-panel` is a white upper edge, a faint lower one and a soft 18px drop; a hovered button gains a halo in `--color-focus`; `--shadow-dialog` is the 36px shadow of a floating window. `--fill-panel` is translucent and `--backdrop-blur` 12px blurs what lies behind it. `--transition` none.
- One-off shades are made with `color-mix()` over tokens: outlines of gels (the fill's colour mixed with `--color-shadow`), the sheet and the rail (90% and 84% `--color-canvas`), the code window (88% `--color-canvas` under `--fill-panel`), the notice border (`--color-notice-text` into `--color-notice`), gloss on danger, success and warning fills (mixed with `--color-button-text`). Half radii are `calc()` over `--radius-control`.
- The era's conventions for which text sits on which fill are kept: body, heading and link colours on canvas, surfaces, `--fill-panel` and the four tints; `--color-bar-text` only on `--fill-bar`, `--color-bar-alt-text` only on `--fill-bar-alt`, `--color-inverse-text` only on `--fill-inverse`; `--color-button-text` on `--fill-button` and on fills made from `--color-danger`, `--color-success` or `--color-accent-alt`; `--color-accent-text` on `--fill-accent` and on `--color-warning`; `--color-notice-text` on `--color-notice`. No text is set directly on the page.

## Components

- `.ds-page`: on `<body>`. `.ds-page__wrap` is the 960px grid; `.ds-page__content` its right column; `.ds-page__sheet` the glass sheet; `.ds-page__body` its padded area; `.ds-page__columns` holds `.ds-page__main` and `.ds-page__aside`. `.ds-section` separates blocks by 28px, `.ds-section__title` (19px) heads one. `.ds-view` is one screen (`.ds-view--home` the first).
- `.ds-nav`: the contents rail. `.ds-nav__head` is the bar with the `.ds-brand`; `.ds-nav__search` holds a capsule `.ds-nav__input`; `.ds-nav__links` is the tree: each `.ds-nav__item` has a `.ds-nav__link` with a `.ds-nav__icon` (current: `is-current`, a gel bar; hover: `is-hover`) and may hold a `.ds-nav__sub` list of `.ds-nav__sublink` pages on a hairline. A `.ds-nav__label` heads a further plain list of pages (`.ds-nav__sub--plain`). `.ds-nav__foot` at the bottom holds a `.ds-nav__note` with the `.ds-nav__version` and one `.ds-button--block`. Four to six top links.
- `.ds-brand`: the site's mark and name on the rail's bar, linking home; one per page. `.ds-brand__mark` is a rounded tile dressed like the primary button; `.ds-brand__name` is the name at `--text-h2` in `--color-bar-text`. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo: replace both, keep the tile.
- `.ds-hero`: the overview's opening band on `--fill-inverse`, home page only. `.ds-hero__main` holds `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__action` (a `.ds-button--large` and a secondary one) and a small `.ds-hero__note`. `.ds-hero__window` on the right is a glass window: `.ds-hero__window-bar` caption and a `.ds-hero__code` block of at most twelve short lines (`.ds-hero__code-key` for object names, `.ds-hero__code-note` for comments).
- `.ds-page-header`: the head of an inner page instead of the hero: `.ds-page-header__main` with `.ds-page-header__title` (25px) and one grey `.ds-page-header__text` line; `.ds-page-header__actions` at the right; a 2px rule under the whole. One per inner page.
- `.ds-stat`: one figure, a pale glass tile with a 25px `.ds-stat__number` over an 11px `.ds-stat__label`; four sit in a `.ds-stats` row at the top of the overview.
- `.ds-grid`: topics in three columns, 12px apart. `.ds-grid__cell` (`--1` to `--4` choose the tint it fades into) holds a `.ds-grid__icon` tile (aqua by default, `--bar` violet, `--accent` orchid, `--alt` orange) beside `.ds-grid__body` with a `.ds-grid__title` and one grey `.ds-grid__text` line.
- `.ds-prose`: long text: h1, h2 (teal), h3 (bold 14px), p, ul, ol, code, strong and links inside it are styled. `.ds-code` is a block of sample code inside or beside it.
- `.ds-link`: any link. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--title` bold, `.ds-link--quiet` for minor links, `.ds-link--code` for object names. `.ds-links` lays several out in a row.
- `.ds-button`: the primary action: aqua gel, white bold 12px, 26px tall, 8px corners; may end with a `.ds-button__icon`. Variants: `.ds-button--secondary` (lavender glass, the usual button), `.ds-button--danger` (red gel, only for an action that removes), `.ds-button--large` (36px, the hero), `.ds-button--block` (full width, the rail). States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttons` is a row. One primary button per form, panel or dialog.
- `.ds-form`: labels above their fields. `.ds-form__row` holds a bold `.ds-form__label` and a `.ds-form__input`, `.ds-form__select` or `.ds-form__textarea`; `.ds-form__pair` sets two short rows side by side; `.ds-form__hint` and `.ds-form__error` under a control; `is-error`, `is-focus`, `is-disabled` on the control; `.ds-form__check` with `.ds-form__checkbox`; `.ds-form__actions` for the buttons over a hairline.
- `.ds-table`: data table in a rounded border with the panel shadow: `.ds-table__head` cells are a strip of `--fill-bar`; `.ds-table__row` with `is-alt` on alternate rows; `.ds-table__cell` with a hairline above; `is-numeric` right-aligns; `.ds-table__note` is a second grey line. `.ds-table__foot` under it holds the key and the pagination.
- `.ds-list`: list of updates. Each `.ds-list__item` is a `.ds-list__mark` dot beside `.ds-list__body`: a bold `.ds-list__title`, a `.ds-list__text` line and a grey `.ds-list__meta`. Rows are divided by hairlines.
- `.ds-panel`: a glass box in the side column (`is-last` on the last block): teal `.ds-panel__title` over a hairline, `.ds-panel__body` with `.ds-panel__text`, a `.ds-panel__facts` list of `.ds-panel__fact` rows (name in `.ds-panel__fact-name`, value at the right) and `.ds-panel__actions`.
- `.ds-tabs`: a pale toolbar strip of `.ds-tabs__tab` links; the current one `is-current` is a white raised plate. Two to five tabs.
- `.ds-badge`: a small gel label, orchid by default. `.ds-badge--count` (flat pale), `.ds-badge--new` (orange). Status variants: `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger`. `.ds-badges` lays several out in a row.
- `.ds-sidebar`: an unboxed link block in the side column (`is-last` on the last): `.ds-sidebar__title` and a `.ds-sidebar__list` hanging from a 2px rule; each `.ds-sidebar__item` (current: `is-current`, marked in `--color-accent`) may end with a `.ds-sidebar__count`. For the contents of a page and short link lists.
- `.ds-notice`: a rounded message with a square `.ds-notice__sign`, `.ds-notice__text` and a bold `.ds-notice__title`. Variant `.ds-notice--error`.
- `.ds-pagination`: one segmented key: `.ds-pagination__item` holding `.ds-pagination__link`, `is-current` (aqua gel), `is-disabled`.
- `.ds-breadcrumb`: 11px trail of `.ds-link` items with teal `.ds-breadcrumb__sep` marks and a bold `.ds-breadcrumb__current`.
- `.ds-dialog`: a small window with a 2px rim: `.ds-dialog__title` on `--fill-bar` (with a round red `.ds-dialog__close`), `.ds-dialog__body` with `.ds-dialog__text`, and a pale `.ds-dialog__actions` strip with the buttons at the right. It casts `--shadow-dialog` and sits on `.ds-dialog__backdrop`, a dimmed rounded strip in the page flow.
- `.ds-empty`: a pale well with a glass `.ds-empty__icon` tile, `.ds-empty__title`, `.ds-empty__text` and one button.
- `.ds-footer`: a dark band under the rail and the sheet: a `.ds-footer__list` of `.ds-footer__link` items on the left, the `.ds-footer__legal` line on the right.
- `.ds-progress`: the glossy accent progress bar, a native `<progress>` in a `.ds-progress__row` with label and value; for downloads, installs and storage.
- `.ds-accordion`: expanding help sections, one `<details class="ds-accordion__item">` per question with `__head` and `__body`.
- `.ds-tooltip`: a round `?` hint that opens a pale balloon (`__trigger`, `__text`) on hover or focus, beside a label or title.

## Never

- `border-radius <= 20px`: corners are 4, 6, 8, 12 or 20px; only dots, round signs and capsules are rounder.
- `border-width <= 2px`: hairlines, and 2px under the page header, beside sidebar lists and around the dialog.
- `box-shadow-blur <= 36px`: shadows are soft; only the floating windows reach 36px.
- `gradient-fills <= 30%`: gloss belongs to bars, keys, labels, tiles and cells; text areas stay plain.
- `font-size <= 32px`: the hero title is the largest text.
- `font-size >= 11px`: nothing is smaller than the meta text.
- `font-weight <= 700`: bold is the heaviest weight; headings are regular.
- `font-families <= 3`: a body sans, a heading sans, a monospace for code.
- `line-height <= 1.5`: body text is 13px on about 19px.
- `underlined-links <= 10%`: links underline on hover only.
- `letter-spacing = 0px`: no tracking.
- `uppercase-text = 0%`: nothing is uppercased by CSS.
- `row-gap <= 24px`: rows are divided by a hairline, not by space.
- `block-gap <= 28px`: sections are at most 28px apart.
- `content-width <= 960px`: text stays inside the rail and the sheet.
- `palette-colours <= 60`: violets and lavenders, aqua, orchid, teal, orange, green, amber and red for status.
- `transition = none`: states change at once.
- `animation = none`: nothing moves.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new bar is `--fill-bar`, a pale strip `--fill-bar-alt`, a dark band `--fill-inverse`; a new box is a `.ds-panel` if it explains and a `.ds-sidebar` if it lists links; a new entry in the rail is a `.ds-nav__item`, never a second bar. Keep to the conventions above for which text token sits on which fill, give gels an outline mixed from their colour and `--color-shadow`, take highlights from the surface tokens instead of drawing new ones, and make any extra shade with `color-mix()` over existing tokens.
