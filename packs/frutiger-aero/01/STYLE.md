# Frutiger Aero product page

## Summary

A fixed-width product page in the glossy, optimistic look that desktop operating systems wore in the second half of the 2000s: a sky-blue page with a pale glow at the top, one segmented glass bar for navigation, and a single rounded sheet that opens with a sky band where a glass window floats over its own reflection above green hills. It was common from about 2004 to 2010 on the pages of operating systems, desktop software, phones and gadgets. Gloss is everywhere a finger would go (bars, buttons, orbs, pills, tabs), panels are white glass with a bright upper edge and a soft shadow, corners are rounded by 4 to 10px, and the text is a small humanist sans in deep blue and grass green.

## Layout

- Page width is fixed: `--size-page` = 940px, centred in a 1024px window. It does not stretch and does not reflow for phones.
- Everything sits in `.ds-page__wrap` (940px): the navigation bar (`--size-nav` 38px tall), then `.ds-page__sheet`, the one glass sheet of the page (88% `--color-canvas`, `--radius-page` corners, `--shadow-panel`), then the footer band. `--space-4` (12px) separates bar and sheet, `--space-5` (20px) sheet and footer.
- The home page fills the sheet edge to edge with the hero and the stats strip under it; all other content sits in `.ds-page__body` (20px padding at the sides and top, 32px at the bottom): the row of four feature cells, then two columns.
- Two columns in `.ds-page__columns`: main (600px) and a right column of `--size-sidebar` (280px), `--space-5` (20px) apart. The side column is always on the right and holds `.ds-panel` and `.ds-sidebar` blocks.
- Spacing scale: `--space-1` 2px (hairline gaps, date tile padding), `--space-2` 4px (between a title and its rule, inside controls), `--space-3` 8px (control padding, gaps between buttons, table cell padding), `--space-4` 12px (panel padding, gaps between grid cells, list row padding), `--space-5` 20px (gutter, sheet padding, gap between stacked blocks), `--space-6` 32px (hero padding, gap between sections; the largest gap).
- Fixed sizes: controls `--size-control` 26px tall, the hero buttons `--size-control-large` 38px, feature orbs `--size-orb` 56px, the hero window `--size-window` 350px with a reflection of `--size-reflection` 46px, date tiles `--size-date` 44px, form labels `--size-label` 150px, the dialog `--size-dialog` 440px.
- Inner pages drop the hero and the stats strip. In `.ds-page__body`, top to bottom: `.ds-breadcrumb`, then `.ds-page-header` across the full 900px (title and one line on the left, one or two buttons on the right, a hairline under it), then either a full-width block (the comparison table, the `.ds-steps` row, a boxed `.ds-stats--boxed` row) or the same two columns. `.ds-tabs` sit at the top of the main column directly above the text or list they switch; notices sit above the form they report on; a dialog follows the columns across the full width on its dimmed strip. Further blocks are `.ds-section` blocks 32px apart with a `.ds-section__title`.
- Views: the specimen is one product site of five screens, each a `.ds-view` inside the sheet, with the bar and the footer written once around them. `home` is the product page (hero, stats, feature cells, news list, requirements panel, popular answers). `tour` explains one feature (tabs, long text, onward links, the tour's contents and a call to action in the side column). `editions` compares the editions (table with status badges, the key to the labels, a panel and a link block). `download` is the order form (steps, notices, the form with its buttons, a summary panel, the cancelling dialog). `help` lists help articles (boxed stats, tabs, list, pagination, topics, an empty state for open questions). Links are plain `href="#name"`; the bar segment of the showing view is pressed in by a `.ds-page:has(#name:target)` rule, the Home segment also when no view is named.

## Typography and colour roles

- One family: `--font-body`, `--font-heading` and `--font-ui` are "Segoe UI", Tahoma, "Lucida Grande", sans-serif: system faces only, the humanist sans of the period's desktops, with Tahoma and Lucida Grande where Segoe UI is not installed. `--font-mono` (Consolas) is for inline code.
- `--text-base` 12px at `--line-body` 1.45. `--text-small` 11px for meta lines, hints, stat labels, footer. `--text-ui` 12px for buttons and fields; buttons are bold, navigation is regular (`--weight-ui` 400).
- Headings are regular weight (`--weight-heading`, `--weight-display` 400) and get their weight from size and colour: `--text-display` 28px hero title, `--text-h1` 22px page titles and stat figures, `--text-h2` 17px section, panel and cell titles and the brand name, `--text-h3` 13px (bold) list titles, sidebar titles, navigation, tabs. `--text-large` 15px hero lead and the large button. Nothing is uppercased or tracked.
- Links are not underlined at rest (`--link-decoration` none) and underline on hover; titles are bold (`.ds-link--title`).
- `--color-page` #3a8ccb with `--fill-page`: a sky that is darker at the top, with a white and a green glow behind the bar, fading to pale blue below.
- `--color-bar` #2d6fa3 with `--fill-bar` (two-step glass: a light upper half, a hard break at 50%, a darker lower half) and white `--color-bar-text` under `--shadow-text`: the navigation bar, the dialog frame, the blue orb.
- `--color-bar-alt` #cfe3f3 with `--fill-bar-alt` (pale glass with the same break): the stats strip, table head, sidebar titles, the hero window's caption, date tile heads.
- `--color-inverse` #256f95 with `--fill-inverse` (deep at the top, lighter towards the horizon, a soft glow at the upper right): the hero sky and the footer band, white text.
- `--color-heading` #0b4f86 page, section and cell titles and figures; `--color-heading-alt` #549c00 (grass green) panel titles, prose h2, breadcrumb marks. `--color-text` #333333, `--color-text-muted` #636363.
- `--color-link` #07519a, `--color-link-quiet` #006699, visited #5a4fa3, hover #2198c4, active #549c00.
- `--color-button` #36b701 with `--fill-button` (green gel: bright top, a break just under the middle, light again at the foot): the primary button, the brand tile, the arrow button, the current step and page number, the first orb. `--color-button-secondary` #d8e0e8 with `--fill-button-secondary` (silver-blue glass): every other button, tabs at rest, pagination keys.
- `--color-accent` #6eab26 with `--fill-accent`: pill badges, meter bars, list bullets, the edge of the current tab, the hills of the hero. `--color-accent-alt` #ff9c00: the New marker and the fourth orb.
- `--color-fill-1` to `--color-fill-4` (sky, mist, leaf, sun): the lower end of each feature cell's fade; never text or borders.
- Borders are 1px solid: `--color-border` #aaccee for boxes, `--color-border-muted` #d6e3e7 for rules inside a box, `--color-border-strong` #3274a4 for the sheet and the dialog, `--color-input-border` #7fb1d8 for fields.
- `--color-notice` #fff8c9 for notices; `--color-danger` #c3271b on `--color-danger-surface` #fde9e4 for errors; `--color-success` #3f9a0b and `--color-warning` #f0a000 fill the status badges and signs.
- Surface: `--radius-control` 4px, `--radius-panel` 7px, `--radius-pill` 10px, `--radius-page` 10px. `--shadow-panel` is a white inner top edge plus a soft 6px drop; `--shadow-control` the same on a button; a hovered button gains a halo in `--color-focus`; `--shadow-dialog` is the deep 24px shadow of a floating window. `--fill-panel` is translucent white glass and `--backdrop-blur` 6px blurs what lies behind it. `--transition` none.
- One-off shades are made with `color-mix()` over tokens: outlines of glossy fills (the fill's colour mixed with `--color-shadow`), dividers on the bar (`--color-bar-text` with transparent), the sheet (88% `--color-canvas`), bubbles (`--color-inverse-text` with transparent), the hills (`--color-accent` with transparent), gloss on danger, success and warning fills (mixed with `--color-button-text`).
- The era's conventions for which text sits on which fill, kept by every layout and every token set of the era:
  - `--color-text`, `--color-text-muted`, `--color-heading`, `--color-heading-alt` and the link colours sit on `--color-canvas`, `--color-surface`, `--color-surface-alt`, `--color-surface-strong`, `--fill-panel` and the four `--color-fill-*` tints. A palette keeps page, canvas, surfaces and fills on one side (all light or all dark) so that these stay readable on every one of them.
  - `--color-bar-text` only on `--fill-bar`; `--color-bar-alt-text` only on `--fill-bar-alt`; `--color-inverse-text` only on `--fill-inverse`. Links on a bar take the bar's text colour, never `--color-link`.
  - `--color-button-text` on `--fill-button` and on any fill made from `--color-danger`, `--color-success` or `--color-accent-alt`; it is always the light colour and also supplies the gloss on those fills (mixed in with `color-mix()`).
  - `--color-accent-text` on `--fill-accent` and on `--color-warning`; `--color-button-secondary-text` on `--fill-button-secondary`; `--color-notice-text` on `--color-notice`; `--color-danger` as text on `--color-danger-surface` and on the surfaces.
  - Nothing but decoration sits on `--fill-page`: no text is set directly on the page.
  - `--color-focus` is the glow colour: the focus ring, the halo of the search field and of a hovered button. `--color-inverse-text` draws the bubbles on the sky band. `--color-shadow` tints every shadow and darkens the outline of every glossy fill.

## Components

- `.ds-page`: on `<body>`. Sets the page fill, base font and colour. `.ds-page__wrap` is the 940px column; `.ds-page__sheet` the glass sheet; `.ds-page__body` its padded content area; `.ds-page__columns` holds `.ds-page__main` and `.ds-page__aside`. `.ds-section` with a `.ds-section__title` (17px over a hairline) separates blocks by 32px. `.ds-view` is one screen of the example site (`.ds-view--home` the first).
- `.ds-nav`: the header. `.ds-nav__bar` is one rounded glass bar: the `.ds-brand`, then `.ds-nav__links` of equal `.ds-nav__item` segments, each a centred `.ds-nav__link` divided from the next by a light hairline (current: `is-current`, pressed in and darker; hover: `is-hover`), then `.ds-nav__search` with a pill-shaped `.ds-nav__input`. Five segments at most.
- `.ds-brand`: the site's mark and name at the left end of the bar, linking home; one per page. `.ds-brand__mark` is a round tile dressed like the primary button (`--fill-button`, a darker outline, `--shadow-control`) with the mark in `--color-button-text`; `.ds-brand__name` is the name at `--text-h2` in `--color-bar-text`. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo: replace both, keep the tile.
- `.ds-hero`: the product page's opening band, on the home page only: `--fill-inverse` sky with two green hills along the bottom. `.ds-hero__main` holds `.ds-hero__title` (28px), `.ds-hero__lead` (15px), `.ds-hero__action` (a `.ds-button--large` and a secondary one) and a small `.ds-hero__note`. `.ds-hero__product` on the right holds `.ds-hero__window`, a glass panel with a `.ds-hero__window-bar` caption and `.ds-hero__rows` of `.ds-hero__row` (first: `is-first`; `.ds-hero__row-name`, a `.ds-meter`, `.ds-hero__row-state`), and under it `.ds-hero__reflection`, an empty mirrored copy that fades out. Up to four `.ds-hero__bubble` spans (`--1` to `--4`) float in the sky. `.ds-meter` with `.ds-meter__bar` (`--full`, `--most`, `--some`) is a small progress bar.
- `.ds-page-header`: the head of an inner page instead of the hero: `.ds-page-header__main` with `.ds-page-header__title` (22px, the page's h1) and one grey `.ds-page-header__text` line; `.ds-page-header__actions` at the right with one or two buttons; a hairline under the whole. One per inner page.
- `.ds-stat`: one figure; three or four sit in a `.ds-stats` row, a pale glass strip across the sheet directly under the hero, figures divided by hairlines (first: `is-first`). Each is a 22px `.ds-stat__number` with an 11px `.ds-stat__label` beside it. On an inner page the row takes `.ds-stats--boxed` (a rounded border, 20px below).
- `.ds-grid`: a row of four feature cells 12px apart. `.ds-grid__cell` (`--1` to `--4` choose the tint it fades into) holds a `.ds-orb`, a `.ds-grid__title`, a `.ds-grid__text` and a small `.ds-grid__more` link, all centred. `.ds-orb` is a 56px glossy ball with a line icon: green by default, `.ds-orb--bar` blue, `.ds-orb--accent` grass, `.ds-orb--alt` orange.
- `.ds-prose`: long text: h1, h2 (green over a hairline), h3 (bold 13px), p, ul, ol, code, strong and links inside it are styled. For tour and help pages.
- `.ds-link`: any link. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--title` is bold, for titles; `.ds-link--quiet` for minor links. `.ds-links` lays several out in a row.
- `.ds-button`: the primary action: green gel, white bold 12px, 26px tall, 4px corners; may end with a `.ds-button__icon` arrow. Variants: `.ds-button--secondary` (silver-blue glass, the usual button), `.ds-button--danger` (red gel, only for an action that removes or cancels), `.ds-button--large` (38px, the hero), `.ds-button--arrow` (a round button holding only the arrow icon, placed before the link it belongs to). States `is-hover` (a halo), `is-active`, `is-focus`, `is-disabled`. `.ds-buttons` is a row; `.ds-buttons--bar` a toolbar. One primary button per form, panel or dialog.
- `.ds-form`: label-left form: `.ds-form__row` with a right-aligned bold `.ds-form__label` and a `.ds-form__field` holding `.ds-form__input` (`--short` for half width), `.ds-form__select` or `.ds-form__textarea`; `.ds-form__hint` and `.ds-form__error` under a control; `is-error`, `is-focus` (a glow), `is-disabled` on the control; `.ds-form__check` with `.ds-form__checkbox`; `.ds-form__actions` for the buttons, under a hairline and aligned with the fields. `.ds-steps` above a form shows a wizard's progress: `.ds-steps__step` with a round `.ds-steps__number`, the current one `is-current`.
- `.ds-table`: data table in a rounded 1px border with the panel shadow: `.ds-table__head` cells form a pale glass strip; `.ds-table__row` with `is-alt` on alternate rows; `.ds-table__cell` with a hairline above; `is-numeric` right-aligns; `.ds-table__note` is a second grey line in a cell.
- `.ds-list`: news and article list. Each `.ds-list__item` (first: `is-first`) is an optional `.ds-list__date` tile (`.ds-list__month` on pale glass over `.ds-list__day`) beside `.ds-list__body`: a bold `.ds-list__title`, a `.ds-list__text` line and a grey `.ds-list__meta`. Rows are divided by hairlines, never boxed.
- `.ds-panel`: a white glass box in the side column (`is-last` on the last block of a column): green `.ds-panel__title` over a hairline, `.ds-panel__body` with `.ds-panel__text`, a `.ds-panel__facts` list of `.ds-panel__fact` rows (name in `.ds-panel__fact-name`, value at the right) and `.ds-panel__actions`.
- `.ds-tabs`: glass tabs with rounded tops on a hairline: `.ds-tabs__tab`, the current one `is-current` (white, joined to the sheet, a green edge on top). Two to five tabs over the text or list they switch.
- `.ds-badge`: a small gel pill, green by default. `.ds-badge--count` (flat pale, a number), `.ds-badge--new` (orange). Status variants: `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger`, one per table row. `.ds-badges` lays several out in a row.
- `.ds-sidebar`: a link block in the side column (`is-last` on the last): `.ds-sidebar__title` on a pale glass strip, `.ds-sidebar__list` of `.ds-sidebar__item` rows behind small round green bullets (current: `is-current`), with an optional right-aligned `.ds-sidebar__count`.
- `.ds-notice`: a rounded one-line message with a round `.ds-notice__sign`, `.ds-notice__text` and a bold `.ds-notice__title`. Variant `.ds-notice--error`.
- `.ds-pagination`: a row of small glass keys: `.ds-pagination__link`, `is-current` (green gel), `is-disabled`, `.ds-pagination__gap` for an ellipsis.
- `.ds-breadcrumb`: 11px trail of `.ds-link` items with green `.ds-breadcrumb__sep` guillemets and a bold `.ds-breadcrumb__current`.
- `.ds-dialog`: a window: a glass frame in `--fill-bar` holding `.ds-dialog__title` (with a red `.ds-dialog__close`), a white `.ds-dialog__body` with `.ds-dialog__text`, and a `.ds-dialog__actions` strip with the buttons at the right. It casts `--shadow-dialog` and sits on `.ds-dialog__backdrop`, a dimmed rounded strip in the page flow.
- `.ds-empty`: a pale well with a round glass `.ds-empty__orb` icon, `.ds-empty__title`, `.ds-empty__text` and one button.
- `.ds-footer`: a rounded sky band under the sheet: `.ds-footer__columns` of four link groups (`.ds-footer__heading`, `.ds-footer__list`, `.ds-footer__link`) and a `.ds-footer__legal` line over a light hairline.
- `.ds-menu`: a glass drop-down under a navigation item (`.ds-menu__list` of `.ds-menu__link`), opened on hover or focus, never forced open.
- `.ds-progress`: the glossy accent progress bar, a native `<progress>` in a `.ds-progress__row` with label and value; for downloads, installs and storage.
- `.ds-accordion`: expanding help sections, one `<details class="ds-accordion__item">` per question with `__head` and `__body`.
- `.ds-tooltip`: a round `?` hint that opens a pale balloon (`__trigger`, `__text`) on hover or focus, beside a label or title.

## Never

- `border-radius <= 10px`: corners are 4, 7 or 10px; only orbs, bullets and pills are rounder.
- `border-width <= 1px`: every border and rule is a hairline.
- `box-shadow-blur <= 24px`: shadows are soft but short; only the floating window reaches 24px.
- `gradient-fills <= 25%`: gloss belongs to bars, buttons, orbs, pills and cells; text areas stay plain.
- `font-size <= 28px`: the hero title is the largest text.
- `font-size >= 11px`: nothing is smaller than the meta text.
- `font-weight <= 700`: bold is the heaviest weight; headings are regular.
- `font-families <= 2`: one humanist sans, a monospace only for inline code.
- `line-height <= 1.5`: body text is 12px on about 17px.
- `underlined-links <= 10%`: links underline on hover only.
- `letter-spacing = 0px`: no tracking.
- `uppercase-text = 0%`: nothing is uppercased by CSS.
- `row-gap <= 24px`: list rows are divided by a hairline, not by space.
- `block-gap <= 32px`: sections are at most 32px apart.
- `content-width <= 940px`: text stays inside the fixed column.
- `palette-colours <= 44`: sky blues, one grass green, orange, amber, red, four pale tints and greys.
- `transition = none`: states change at once.
- `animation = none`: nothing moves.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new bar is `--fill-bar`, a pale strip `--fill-bar-alt`, a sky band `--fill-inverse`; a new box is a `.ds-panel` if it explains and a `.ds-sidebar` if it lists links; anything round and pressable is dressed like `.ds-orb` or `.ds-button--arrow`. Keep to the conventions above for which text token sits on which fill, give glossy fills an outline mixed from their colour and `--color-shadow`, take highlights from the surface tokens instead of drawing new ones, and make any extra shade with `color-mix()` over existing tokens.
