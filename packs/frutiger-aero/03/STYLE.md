# Frutiger Aero account centre

## Summary

A fixed-width account and support centre in the fresh, daylight variety of the Frutiger Aero look: aqua and lime on white, with brushed silver for the top strip, the tabs and the status bar. Self-service centres of this kind were common from about 2005 to 2010 beside desktop software, phones and broadband accounts, and borrowed their layout from the windows of the operating system: large tabs, a task pane on the left, wizard steps, a status bar. Surfaces are crisp rather than glassy (small radii, short shadows, no blur), gloss is kept to buttons, tabs and pills, and the type pairs small Verdana with bold Trebuchet headings.

## Layout

- Page width is fixed: `--size-page` = 960px, centred in a 1024px window. It does not stretch and does not reflow for phones.
- The top strip spans the window (`--size-top` 46px) and holds a centred `--size-page` box; under it, 14px lower, a row of large tabs (`--size-tab` 34px) from which the sheet hangs.
- `.ds-page__wrap` (960px) holds each view's `.ds-page__frame`, the window-like sheet (92% `--color-canvas`, a hairline border, `--radius-page` on its upper right corner), and under it the status-bar footer, which closes the sheet with the two lower corners rounded.
- The frame is two columns of equal height: the task pane `.ds-page__pane` on the left (`--size-pane` 210px, a tinted fade, a hairline at its right) and `.ds-page__main` (748px, 20px padding). The pane is on every view and always on the left.
- On the overview the main column runs: hero band, `.ds-stats` row of three gauges, the two-by-two task grid under a `.ds-section__title`, then `.ds-page__split`, two equal `.ds-page__half` columns of panels 20px apart.
- Spacing scale: `--space-1` 3px (hairline gaps, pane row padding), `--space-2` 6px (gaps between buttons, list row padding, panel caption padding), `--space-3` 10px (control padding, table cell padding, gap between task tiles), `--space-4` 14px (panel and tile padding, gap between gauges), `--space-5` 20px (gutter, main padding, gap between stacked blocks), `--space-6` 28px (gap between sections; the largest gap).
- Fixed sizes: controls `--size-control` 24px tall, the hero button `--size-control-large` 34px, the brand tile `--size-mark` 26px, meters `--size-meter` 8px, form labels `--size-label` 150px, the dialog `--size-dialog` 420px.
- Inner pages drop the hero, gauges and task grid; the pane stays. In `.ds-page__main`, top to bottom: `.ds-breadcrumb`, then `.ds-page-header` across the column (title and one line on the left, the buttons that act on the page at the right, a 1px rule in `--color-accent` under it), then the working blocks at full column width: `.ds-tabs` directly above the table they switch, `.ds-pagination` under it, the wizard's `.ds-steps` above its notices and form, a dialog on its dimmed strip after the table. A page of reading uses `.ds-page__split`: the text in the left half, panels in the right. Further blocks are `.ds-section` blocks 28px apart.
- Views: the specimen is one account centre of four screens, each a `.ds-view` holding its own frame, with the header and the footer written once around them. `home` is the overview (welcome band, gauges, task tiles, recent activity, the plan). `devices` is the device list (button row, segmented tabs, table with status badges, pagination, the key to the states, the removing dialog). `setup` is step 2 of the set-up wizard (steps, notices, the form). `help` is one guide (long text, onward links, an empty state for open questions, a feedback panel). Links are plain `href="#name"`; the tab of the showing view is raised and joined to the sheet by a `.ds-page:has(#name:target)` rule, the Overview tab also when no view is named.

## Typography and colour roles

- Three families: `--font-body` Verdana, Geneva, "DejaVu Sans"; `--font-heading` "Trebuchet MS", "Lucida Sans", Tahoma; `--font-ui` Tahoma, Verdana, Geneva. All are system fonts. `--font-mono` ("Courier New") is for inline code.
- `--text-base` 11px at `--line-body` 1.4. `--text-small` 10px for meta lines, labels, hints and the status bar. `--text-ui` 11px bold (`--weight-ui` 700) for buttons, tabs and steps.
- Headings are bold (`--weight-heading`, `--weight-display` 700): `--text-display` 24px the welcome title and gauge figures; `--text-h1` 20px page titles; `--text-h2` 15px section and tile titles, the brand name and the section name in the strip; `--text-h3` 12px panel captions, pane headings, the large tabs. `--text-large` 13px the hero lead and large button. Nothing is uppercased or tracked.
- Links in text, lists and the pane are underlined (`--link-decoration` underline); titles (`.ds-link--title`), tabs and bar links are not.
- `--color-page` #dfe5e8 with `--fill-page`: brushed silver, fine vertical lines, lighter at the top.
- `--color-bar` #c4cace with `--fill-bar` (brushed metal: fine lines over a light-dark-light sheen) and dark `--color-bar-text` #25323a with a light `--shadow-text` under it: the top strip, the status bar, the dialog frame.
- `--color-bar-alt` #e4f5bd (pale lime) with `--fill-bar-alt`: panel captions, the table head, the steps strip; text `--color-bar-alt-text` #3a5200.
- `--color-inverse` #0e8f9e (aqua) with `--fill-inverse` (lighter at the top, a soft light rising from the lower left): the welcome band, white text.
- `--color-heading` #0a7a86 (aqua) titles and figures; `--color-heading-alt` #5a8a00 (lime, darkened for text) pane headings, prose h2, breadcrumb marks. `--color-text` #444444, `--color-text-muted` #686868. Surfaces: `--color-surface-alt` #f1f8f6 alternate rows and the pane, `--color-surface-strong` #dcefe9 meter tracks and the dialog's action strip.
- `--color-link` #00768c, `--color-link-quiet` #555555, visited #7a5fa8, hover #00a5ba, active #5a8a00.
- `--color-button` #12a0b5 with `--fill-button` (aqua, a hard break at 50%): the primary button, the brand tile, arrow buttons, the current step and page number. `--color-button-secondary` #d9dfe2 with `--fill-button-secondary` (silver): every other button, the large tabs at rest, the segmented tabs.
- `--color-accent` #9bd61a (lime) with dark `--color-accent-text`: pill badges, meter bars, the current segmented tab, the edge of the current large tab, the rule under the page header, pane arrows. `--color-accent-alt` #ff7a1a: the New marker.
- `--color-fill-1` to `--color-fill-4` (aqua tint, lime tint, silver, mint): the left end of each task tile's fade.
- Borders are 1px solid: `--color-border` #c5d3d6 for boxes, `--color-border-muted` #e3ebec between rows, `--color-border-strong` #8fa3a8 for the sheet and the dialog, `--color-input-border` #9fb3b8.
- `--color-notice` #f3fbd9; `--color-danger` #c8321f on `--color-danger-surface` #fdebe6; `--color-success` #5c9a08 and `--color-warning` #f2a100 fill status badges and the dots of the activity list.
- Surface: `--radius-control` 3px, `--radius-panel` 5px, `--radius-page` 6px, `--radius-pill` 9px. Shadows are short: `--shadow-panel` a 3px drop, `--shadow-control` a white inner top edge and a 1px drop, `--shadow-dialog` 12px. `--fill-panel` is opaque white ending in a faint tint; `--backdrop-blur` is 0px: nothing is translucent. `--transition` none.
- One-off shades are made with `color-mix()` over tokens: outlines of glossy fills (the fill's colour mixed with `--color-shadow`), the sheet (92% `--color-canvas`), bubbles (`--color-inverse-text` with transparent), gloss on status fills (mixed with `--color-button-text`).
- The era's conventions for which text sits on which fill, kept by every layout and every token set of the era:
  - `--color-text`, `--color-text-muted`, `--color-heading`, `--color-heading-alt` and the link colours sit on `--color-canvas`, `--color-surface`, `--color-surface-alt`, `--color-surface-strong`, `--fill-panel` and the four `--color-fill-*` tints. A palette keeps page, canvas, surfaces and fills on one side (all light or all dark) so that these stay readable on every one of them.
  - `--color-bar-text` only on `--fill-bar`; `--color-bar-alt-text` only on `--fill-bar-alt`; `--color-inverse-text` only on `--fill-inverse`. Links on a bar take the bar's text colour, never `--color-link`.
  - `--color-button-text` on `--fill-button` and on any fill made from `--color-danger`, `--color-success` or `--color-accent-alt`; it is always the light colour and also supplies the gloss on those fills (mixed in with `color-mix()`).
  - `--color-accent-text` on `--fill-accent` and on `--color-warning`; `--color-button-secondary-text` on `--fill-button-secondary`; `--color-notice-text` on `--color-notice`; `--color-danger` as text on `--color-danger-surface` and on the surfaces.
  - Nothing but decoration sits on `--fill-page`: no text is set directly on the page.
  - `--color-focus` is the glow colour: the focus ring, the halo of the search field and of a hovered button. `--color-inverse-text` draws the bubbles on the sky band. `--color-shadow` tints every shadow and darkens the outline of every glossy fill.

## Components

- `.ds-page`: on `<body>`. Sets the page fill, base font and colour. `.ds-page__wrap` is the 960px column; each view holds a `.ds-page__frame` with `.ds-page__pane` (the task pane) and `.ds-page__main`; `.ds-page__split` lays two `.ds-page__half` columns side by side. `.ds-section` with a `.ds-section__title` (15px bold) separates blocks by 28px. `.ds-view` is one screen of the example site (`.ds-view--home` the first).
- `.ds-nav`: the header. `.ds-nav__top` is the brushed strip across the window; its `.ds-nav__inner` holds the `.ds-brand`, the centre's name as `.ds-nav__section` behind a hairline, and right-aligned `.ds-nav__tools` (text, bold `.ds-nav__tool` links, the `.ds-nav__input` search field). `.ds-nav__tabs` under it is the row of large glossy tabs, each a `.ds-nav__link`: silver at rest, the current one (`is-current`) taller, white, joined to the sheet, with a lime edge on top. Four or five tabs.
- `.ds-brand`: the site's mark and name at the left end of the strip, linking to the overview; one per page. `.ds-brand__mark` is a round tile dressed like the primary button (`--fill-button`, a darker outline, `--shadow-control`) with the mark in `--color-button-text`; `.ds-brand__name` is the name at `--text-h2` in `--color-bar-text`. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo: replace both, keep the tile.
- `.ds-hero`: the welcome band at the top of the overview: a rounded block in `--fill-inverse` with up to three `.ds-hero__bubble` spans (`--1` to `--3`). `.ds-hero__main` holds `.ds-hero__title` (24px) and one `.ds-hero__lead` line; `.ds-hero__action` at the right holds one `.ds-button--large`. On the overview only.
- `.ds-page-header`: the head of an inner page instead of the hero: `.ds-page-header__main` with `.ds-page-header__title` (20px, the page's h1) and a grey `.ds-page-header__text` line; `.ds-page-header__actions` at the right with the page's buttons; a lime rule under the whole. One per inner page.
- `.ds-stat`: one gauge; three sit in a `.ds-stats` row under the hero. Each is a bordered tile with a small `.ds-stat__label`, a 24px `.ds-stat__number` (its unit in a grey `.ds-stat__unit`) and a `.ds-meter` whose `.ds-meter__bar` (`--full`, `--most`, `--some`) shows the share used.
- `.ds-grid`: the task tiles, two columns. Each `.ds-grid__cell` (`--1` to `--4` choose the tint it fades from) holds `.ds-grid__body` (`.ds-grid__title`, `.ds-grid__text`) and a `.ds-button--arrow` at the right. One tile per thing the visitor can do.
- `.ds-tabs`: a segmented control: `.ds-tabs__item` (first `is-first`, last `is-last`) each holding a `.ds-tabs__tab`; the current one `is-current` is lime and pressed in; `.ds-tabs__count` for a number. Directly above the table it filters.
- `.ds-list`: the activity log. Each `.ds-list__item` (first: `is-first`) is a round `.ds-list__dot` (green; `--warning`, `--danger`) beside `.ds-list__body`: a `.ds-list__title` line and a grey `.ds-list__meta`. Rows are divided by hairlines.
- `.ds-panel`: a small window (`is-last` on the last of a column): `.ds-panel__title` is a pale lime caption strip, `.ds-panel__body` holds `.ds-panel__text`, a `.ds-panel__facts` list of `.ds-panel__fact` rows (name in `.ds-panel__fact-name`, value at the right) and `.ds-panel__actions`.
- `.ds-button`: the primary action: aqua gloss, white bold 11px, 24px tall, 3px corners; may end with a `.ds-button__icon` arrow. Variants: `.ds-button--secondary` (silver, the usual button), `.ds-button--danger` (red, only for removing), `.ds-button--large` (34px, the hero), `.ds-button--arrow` (a round button holding only the arrow, at the right of a task tile). States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttons` is a row; `.ds-buttons--bar` a toolbar. One primary button per page header, form, panel or dialog. `.ds-steps` shows a wizard's progress above its form: joined `.ds-steps__step` segments (first `is-first`) with a round `.ds-steps__number`; finished steps `is-done`, the current one `is-current`.
- `.ds-sidebar`: one block of the task pane (`is-last` on the last): a lime-green `.ds-sidebar__title` over a hairline, then `.ds-sidebar__list` of `.ds-sidebar__item` links behind small arrows (current: `is-current`) or a `.ds-sidebar__text` line and a button.
- `.ds-prose`: long text: h1, h2 (in `--color-heading-alt` over a hairline), h3 (bold), p, ul, ol, code, strong and links inside it are styled. For articles and guides.
- `.ds-link`: any link. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--title` is bold, for titles; `.ds-link--quiet` for minor links. `.ds-links` lays several out in a row.
- `.ds-form`: label-left form: `.ds-form__row` with a right-aligned bold `.ds-form__label` and a `.ds-form__field` holding `.ds-form__input` (`--short` for half width), `.ds-form__select` or `.ds-form__textarea`; `.ds-form__hint` and `.ds-form__error` under a control; `is-error`, `is-focus` (a glow in `--color-focus`), `is-disabled` on the control; `.ds-form__check` with `.ds-form__checkbox`; `.ds-form__actions` for the buttons, under a hairline and aligned with the fields.
- `.ds-table`: data table in a rounded 1px border with the panel shadow: `.ds-table__head` cells form a strip in `--fill-bar-alt`; `.ds-table__row` with `is-alt` on alternate rows; `.ds-table__cell` with a hairline above; `is-numeric` right-aligns; `.ds-table__note` is a second grey line in a cell.
- `.ds-badge`: a small gel pill in `--fill-accent`. `.ds-badge--count` (flat, a number), `.ds-badge--new` (`--color-accent-alt`). Status variants: `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger`, one per row. `.ds-badges` lays several out in a row.
- `.ds-notice`: a rounded one-line message with a round `.ds-notice__sign`, `.ds-notice__text` and a bold `.ds-notice__title`. Variant `.ds-notice--error`. Notices sit above the form or list they report on.
- `.ds-pagination`: a row of small glass keys: `.ds-pagination__link`, `is-current` (the primary button's gel), `is-disabled`, `.ds-pagination__gap` for an ellipsis.
- `.ds-breadcrumb`: a small trail of `.ds-link` items with `.ds-breadcrumb__sep` guillemets in `--color-heading-alt` and a bold `.ds-breadcrumb__current`.
- `.ds-dialog`: a window: a glass frame in `--fill-bar` holding `.ds-dialog__title` (with a red `.ds-dialog__close`), a `.ds-dialog__body` in `--color-surface` with `.ds-dialog__text`, and a `.ds-dialog__actions` strip with the buttons at the right. It casts `--shadow-dialog` and sits on `.ds-dialog__backdrop`, a dimmed rounded strip in the page flow, directly under the table or form that opened it.
- `.ds-empty`: a pale well with a round glass `.ds-empty__orb` icon, `.ds-empty__title`, `.ds-empty__text` and one button.
- `.ds-footer`: the status bar that closes the sheet: a brushed strip with `.ds-footer__legal` at the left and `.ds-footer__links` (`.ds-footer__link`) at the right, the two lower corners rounded.
- `.ds-menu`: a glass drop-down under a navigation item (`.ds-menu__list` of `.ds-menu__link`), opened on hover or focus, never forced open.
- `.ds-progress`: the glossy accent progress bar, a native `<progress>` in a `.ds-progress__row` with label and value; for downloads, installs and storage.
- `.ds-accordion`: expanding help sections, one `<details class="ds-accordion__item">` per question with `__head` and `__body`.
- `.ds-avatar`: the framed account picture drawn as initials, beside a name in `.ds-person`; `--small` for tight places.

## Never

- `border-radius <= 9px`: corners are 3, 5 or 6px; only dots, round buttons and pills are rounder.
- `border-width <= 1px`: every border and rule is a hairline.
- `box-shadow-blur <= 12px`: shadows are short; the dialog's 12px is the deepest.
- `gradient-fills <= 26%`: gloss belongs to the strips, tabs, buttons, pills and tiles; text areas stay plain.
- `font-size <= 24px`: the welcome title and the gauge figures are the largest text.
- `font-size >= 10px`: nothing is smaller than the meta text.
- `font-weight <= 700`: bold is the heaviest weight.
- `font-families <= 4`: a body sans, a heading sans, a control sans, a monospace for inline code.
- `line-height <= 1.4`: body text is 11px on about 15px.
- `underlined-links >= 40%`: links in text, lists and the pane are underlined; only titles and tabs are not.
- `letter-spacing = 0px`: no tracking.
- `uppercase-text = 0%`: nothing is uppercased by CSS.
- `row-gap <= 20px`: rows are divided by a hairline, not by space.
- `block-gap <= 28px`: sections are at most 28px apart.
- `content-width <= 960px`: text stays inside the fixed column.
- `palette-colours <= 44`: aqua, lime, silver greys, orange, amber, red and four pale tints.
- `transition = none`: states change at once.
- `animation = none`: nothing moves.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new task is a tile in `.ds-grid` with an arrow button; a new fact box is a `.ds-panel` with a caption; a new group of shortcuts is a `.ds-sidebar` block in the pane; a new multi-step job reuses `.ds-steps` over a `.ds-form`. A new strip is `--fill-bar`, a caption `--fill-bar-alt`, a band `--fill-inverse`. Keep to the conventions above for which text token sits on which fill, give glossy fills an outline mixed from their colour and `--color-shadow`, and make any extra shade with `color-mix()` over existing tokens.
