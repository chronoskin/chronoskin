# Frutiger Aero media gallery

## Summary

A fixed-width media gallery in the glossy black and pearl variety of the Frutiger Aero look: the whole site is one player window on a pale page, with a dark glass title bar, a black stage where album covers stand on their own reflections, and a control strip with round transport keys to close it. Galleries, players and media centres of this kind were common from about 2005 to 2010, when photos, music and films first shared one library on the home computer. Dark glass with an arched highlight is kept for bars and keys, the working area is opaque white and pale grey, controls are capsules, one green lights whatever is current, and the type is a neutral grotesque with large light headings.

## Layout

- Page width is fixed: `--size-page` = 960px, centred in a 1024px window. It does not stretch and does not reflow for phones.
- Everything sits in `.ds-page__wrap` (960px) inside `.ds-page__player`, one window with `--radius-page` corners: the title bar (`.ds-nav`, `--size-nav` 44px), then the showing view, then the control strip (`.ds-footer`). Under the window `.ds-page__floor` mirrors the strip and fades out: the window stands on a reflecting floor.
- The home page fills the window edge to edge with the stage (`.ds-hero`) and the readout strip (`.ds-stats`); the rest sits in `.ds-page__body` (18px padding, 30px at the bottom): the row of four album cells, then two columns.
- Columns in `.ds-page__columns`, `--space-5` (18px) apart: `.ds-page__main` beside a right column `.ds-page__aside` of `--size-side` (300px) for panels, or beside a left source list `.ds-page__source` of `--size-source` (190px) on the library page. Never three columns.
- Spacing scale: `--space-1` 2px (hairline gaps, the lamp under the current link), `--space-2` 5px (inside controls, table cell padding), `--space-3` 8px (gaps between buttons, row padding, panel title padding), `--space-4` 12px (panel padding, gaps between cells and thumbnails), `--space-5` 18px (gutter, body padding, gap between stacked blocks), `--space-6` 30px (stage padding, gap between sections; the largest gap).
- Fixed sizes: controls `--size-control` 24px tall, large keys `--size-control-large` 34px, shelf covers `--size-tile` 124px (the current one 1.22 times that), list pictures `--size-thumb` 38px, the now-playing cover `--size-cover` 64px, form labels `--size-label` 170px, the dialog `--size-dialog` 430px.
- Inner pages drop the stage and the readout strip. In `.ds-page__body`, top to bottom: `.ds-breadcrumb`, then `.ds-page-header` across the full width, then the columns. `.ds-tabs` sit at the top of the main column directly above the table or form they switch; notices sit between the tabs and the form; a dialog follows the columns across the full width on its dimmed strip. Further blocks are `.ds-section` blocks 30px apart with a `.ds-section__title`.
- Views: the specimen is one gallery of four screens, each a `.ds-view` inside the window, with the title bar and the control strip written once around them. `home` is the gallery front (stage with the shelf, readout, albums, recently added, now playing). `library` lists everything (source list, tabs, table with status badges, pagination and the key to the badges, an empty playlist). `album` shows one album (thumbnails, the text about it, onward links, details and comments at the side). `settings` is the settings form (tabs, notices, the form with its buttons, disk space, screens, the dialog that empties the library). Links are plain `href="#name"`; the title bar link of the showing view is pressed in over a green lamp by a `.ds-page:has(#name:target)` rule, the Home link also when no view is named.

## Typography and colour roles

- One family: `--font-body`, `--font-heading` and `--font-ui` are "Helvetica Neue", Helvetica, Arial, "Liberation Sans", sans-serif: system and open faces only. `--font-mono` (Menlo) is for inline code.
- `--text-base` 12px at `--line-body` 1.5. `--text-small` 11px for meta lines, hints, readout labels, the control strip. `--text-ui` 11px bold for buttons, tabs and table heads (`--weight-ui` 700).
- Headings are light (`--weight-heading`, `--weight-display` 300) and get their weight from size: `--text-display` 34px for the stage title with `--display-tracking` -1px, `--text-h1` 24px page titles and readout figures, `--text-h2` 16px section titles and the brand name, `--text-h3` 12px (bold) panel and source list titles, navigation. `--text-large` 14px (bold) for album names, the stage lead and large keys. Nothing is uppercased.
- Links are not underlined at rest (`--link-decoration` none) and underline on hover; titles are bold (`.ds-link--title`).
- `--color-page` #e8e8e8 with `--fill-page`: pearl, lit white from the top and a little darker towards the floor.
- `--color-bar` #303030 with `--fill-bar` (dark glass under an arched highlight) and white `--color-bar-text` under `--shadow-text`: the title bar, the control strip, the current tab, the dialog's title bar.
- `--color-bar-alt` #d0d0d0 with `--fill-bar-alt` (silver, light to mid grey): the readout strip, table heads, the current source.
- `--color-inverse` #080808 with `--fill-inverse` (black under a soft spotlight): the stage only, with `--color-inverse-text` #f0f0f0.
- `--color-heading` #202020 page, section and panel titles; `--color-heading-alt` #4a8a05 (green) source list titles, prose h2, breadcrumb marks. `--color-text` #565656, `--color-text-muted` #6e6e6e.
- `--color-link` #2971a7, `--color-link-quiet` #666666, visited #6c5a9c, hover #0060d0, active #59990e.
- `--color-button` #59a80f with `--fill-button` (green gel under the same arch): the primary key, the brand tile, the play key, the current page number. `--color-button-secondary` #d0d0d0 with `--fill-button-secondary` (silver): every other key, tabs at rest.
- `--color-accent` #6bbf12 with `--fill-accent`: badges, meter bars, the lamp under the current link, the frame of the current thumbnail. `--color-accent-alt` #2971a7: the New label.
- `--color-fill-1` to `--color-fill-4` (pearl, mist, leaf, lilac grey): the upper end of each album cell's fade; never text or borders.
- Borders are 1px solid: `--color-border` #cccccc for boxes, `--color-border-muted` #e5e5e5 for rules inside a box, `--color-border-strong` #767676 for the window and the dialog, `--color-input-border` #a8a8a8 for fields.
- `--color-notice` #eef6e2 for notices; `--color-danger` #b3261e on `--color-danger-surface` #fbeceb for errors; `--color-success` #4a9a0a and `--color-warning` #e89a00 fill the status badges and signs.
- Surface: `--radius-control` 12px (controls are capsules), `--radius-panel` 6px, `--radius-pill` 3px (labels and pictures are nearly square), `--radius-page` 8px. `--shadow-panel` is a white upper edge and a short shadow thrown on the floor under a box; `--shadow-control` a bright rim; a hovered key gains a ring in `--color-focus`; `--shadow-dialog` is the deep 30px shadow of a floating window and of the cover in front. Nothing is translucent: `--backdrop-blur` 0px. `--transition` none.
- One-off shades are made with `color-mix()` over tokens: outlines of glossy fills, dividers on the bars (`--color-bar-text` with transparent), the sheen and the motifs of `.ds-art` (mixed from `--color-bar`, `--color-accent`, `--color-accent-alt`, `--color-inverse`, `--color-button` and `--color-button-text`), the floor line of the stage (`--color-inverse-text` with transparent).
- The era's conventions for which text sits on which fill, kept by every layout and every token set of the era:
  - `--color-text`, `--color-text-muted`, `--color-heading`, `--color-heading-alt` and the link colours sit on `--color-canvas`, `--color-surface`, `--color-surface-alt`, `--color-surface-strong`, `--fill-panel` and the four `--color-fill-*` tints. A palette keeps canvas, surfaces and fills on one side (all light or all dark) so that these stay readable on every one of them.
  - `--color-bar-text` only on `--fill-bar`; `--color-bar-alt-text` only on `--fill-bar-alt`; `--color-inverse-text` only on `--fill-inverse`. Links on a bar or band take that bar's text colour, never `--color-link`.
  - `--color-button-text` on `--fill-button` and on any fill made from `--color-danger`, `--color-success` or `--color-accent-alt`; it is always the light colour and also supplies the gloss on those fills (mixed in with `color-mix()`).
  - `--color-accent-text` on `--fill-accent` and on `--color-warning`; `--color-button-secondary-text` on `--fill-button-secondary`; `--color-notice-text` on `--color-notice`; `--color-danger` and `--color-success` as text on the surfaces and on their own pale mixes.
  - Nothing but decoration sits on `--fill-page`: no text is set directly on the page.
  - `--color-focus` is the glow colour: the focus ring, the halo of a focused field and of a hovered button. `--color-shadow` tints every shadow and darkens the outline of every glossy fill (the fill's colour mixed with `--color-shadow`).
  - A `--fill-*` token may end in a plain colour layer, so it is always the last layer of a `background`.

## Components

- `.ds-page`: on `<body>`. Sets the page fill, base font and colour. `.ds-page__wrap` is the 960px column; `.ds-page__player` the window; `.ds-page__floor` its reflection; `.ds-page__body` the padded content area of a view; `.ds-page__columns` holds `.ds-page__main` with `.ds-page__aside` or `.ds-page__source`. `.ds-section` (last: `is-last`) with a `.ds-section__title` (16px, an optional small `.ds-section__more` link at its right end) separates blocks by 30px. `.ds-view` is one screen of the example site (`.ds-view--home` the first).
- `.ds-nav`: the window's title bar: the `.ds-brand`, then `.ds-nav__links` of `.ds-nav__item` segments, each a `.ds-nav__link` divided from the next by a light hairline (current: `is-current`, pressed in over a 2px lamp in `--color-accent`; hover: `is-hover`), then `.ds-nav__search` at the right end with a capsule `.ds-nav__input`. Five links at most.
- `.ds-brand`: the site's mark and name at the left end of the title bar, linking home; one per page. `.ds-brand__mark` is a tile dressed like the primary button (`--fill-button`, a darker outline, `--radius-control`, `--shadow-control`) with the mark in `--color-button-text`; `.ds-brand__name` is the name at `--text-h2` in `--color-bar-text`. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo: replace both, keep the tile.
- `.ds-hero`: the stage, on the home page only: `--fill-inverse` with a bright floor line. `.ds-hero__main` holds a small `.ds-hero__kicker` behind a glowing `.ds-hero__lamp`, `.ds-hero__title` (34px light), `.ds-hero__lead` (14px) and `.ds-hero__action` (a `.ds-button--large` and a secondary one). `.ds-shelf` on the right is a list of three or four `.ds-shelf__tile` links, each a mirrored `.ds-art` cover; the one that is showing is `is-current` (larger, a white frame, `--shadow-dialog`).
- `.ds-art`: a square picture drawn from the palette, used wherever a cover or a photo stands: `.ds-art--1` hills, `--2` sun on the sea, `--3` a record, `--4` a film still, `--5` a peak, `--6` blossoms. `.ds-art--wide` is 4 to 3; `.ds-art--mirror` adds the fading mirror image under it (leave 46% of its height free below). An installing project replaces the motif classes with its own pictures and keeps the frame, the sheen and the mirror.
- `.ds-page-header`: the head of an inner page instead of the hero: `.ds-page-header__main` with `.ds-page-header__title` (`--text-h1`, the page's h1) and one grey `.ds-page-header__text` line; `.ds-page-header__actions` at the right with one or two buttons or a status badge; a hairline under the whole. One per inner page, directly under the breadcrumb.
- `.ds-stat`: one figure; four or five sit in a `.ds-stats` row, the silver readout strip across the window directly under the stage, divided by hairlines (first: `is-first`). Each is an 11px `.ds-stat__label` over a 24px light `.ds-stat__number`. On an inner page the row takes `.ds-stats--boxed`.
- `.ds-grid`: a row of four album cells 12px apart. `.ds-grid__cell` (`--1` to `--4` choose the tint it fades from) holds a `.ds-grid__cover` link with a mirrored `.ds-art`, a bold `.ds-grid__title` and a grey `.ds-grid__text` line, all centred.
- `.ds-thumbs`: the pictures of one album, four to a row: `.ds-thumbs__item` (the one that is showing: `is-current`) with a `.ds-art--wide` and a `.ds-thumbs__name`.
- `.ds-prose`: long text: h1, h2 (in `--color-heading-alt`), h3 (bold), p, ul, ol, code, strong and links inside it are styled.
- `.ds-link`: any link on the sheet. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--title` is bold, for titles; `.ds-link--quiet` for minor links and tools. `.ds-links` lays several out in a row.
- `.ds-button`: the primary action: a glossy key in `--fill-button`, bold `--text-ui`, `--size-control` tall, `--radius-control` corners; may start with a `.ds-button__icon`. Variants: `.ds-button--secondary` (the usual button, `--fill-button-secondary`), `.ds-button--danger` (red gel, only for an action that removes or cancels), `.ds-button--large` (`--size-control-large`, the hero and the main action of a panel), `.ds-button--round` (a circle holding only an icon). States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttons` is a row. One primary button per form, panel or dialog.
- `.ds-form`: settings rows: `.ds-form__row` (first: `is-first`) with a bold `.ds-form__label` at the left (170px) and a `.ds-form__field` holding `.ds-form__input` (`--short` for half width), `.ds-form__select` or `.ds-form__textarea`; `.ds-form__hint` and `.ds-form__error` under a control; `is-error`, `is-focus` (a glow), `is-disabled` on the control; `.ds-form__check` with `.ds-form__checkbox`; rows are divided by hairlines; `.ds-form__actions` holds the buttons under a rule.
- `.ds-table`: the library table in a rounded 1px border: `.ds-table__head` cells form a silver strip divided by hairlines (first: `is-first`); `.ds-table__row` with `is-alt` on alternate rows; `.ds-table__cell` with a hairline above; `is-numeric` right-aligns; `.ds-table__cell--thumb` holds a small `.ds-art`; `.ds-table__note` is a second grey line in a cell. `.ds-table__foot` under it holds the pagination at the left and the key to the badges at the right.
- `.ds-list`: recently added items. Each `.ds-list__item` (first: `is-first`) is a `.ds-list__thumb` picture, a `.ds-list__body` with a bold `.ds-list__title` and a grey `.ds-list__meta` line, a `.ds-list__length` and a round play key. Rows are divided by hairlines, never boxed.
- `.ds-panel`: a box in the side column (`is-last` on the last block of a column): a bold `.ds-panel__title` on a grey strip (`--color-surface-strong`), `.ds-panel__body` with `.ds-panel__text`, a `.ds-panel__facts` list of `.ds-panel__fact` rows (name in `.ds-panel__fact-name`, value at the right) and `.ds-panel__actions`. `.ds-now` inside a panel is what is playing: `.ds-now__cover`, `.ds-now__title`, `.ds-now__meta`, a meter and `.ds-now__time`.
- `.ds-tabs`: one segmented capsule: `.ds-tabs__item` holding a `.ds-tabs__tab`, the current one `is-current` (dark glass, pressed in). Two to five segments over the table or form they switch.
- `.ds-badge`: a small glossy label in `--fill-accent` with `--radius-pill` corners. `.ds-badge--count` (flat, a number), `.ds-badge--new` (in `--color-accent-alt`). Status variants: `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger`, one per row or tile. `.ds-badges` lays several out in a row.
- `.ds-sidebar`: a source list in the left column (`is-last` on the last): a green `.ds-sidebar__title`, a `.ds-sidebar__list` of `.ds-sidebar__item` rows with a right-aligned `.ds-sidebar__count`; the current source (`is-current`) lies on a silver capsule.
- `.ds-notice`: a rounded one-line message with a round `.ds-notice__sign`, `.ds-notice__text` and a bold `.ds-notice__title`. Variant `.ds-notice--error`. Notices sit directly above the form or list they report on.
- `.ds-pagination`: a row of small glossy keys: `.ds-pagination__link`, `is-current` (in `--fill-button`), `is-disabled`, `.ds-pagination__gap` for an ellipsis. It shares the `.ds-table__foot` row under a table or grid with the key to the badges.
- `.ds-breadcrumb`: `--text-small` trail of `.ds-link` items with `.ds-breadcrumb__sep` marks in `--color-heading-alt` and a bold `.ds-breadcrumb__current`.
- `.ds-dialog`: a small window: `.ds-dialog__title` on dark glass (with a round `.ds-dialog__close`), a white `.ds-dialog__body` with `.ds-dialog__text`, and a pale `.ds-dialog__actions` strip with the buttons at the right. It casts `--shadow-dialog` and sits on `.ds-dialog__backdrop`, a dimmed rounded strip in the page flow.
- `.ds-empty`: a pale well with a round glass `.ds-empty__orb` icon, `.ds-empty__title`, `.ds-empty__text` and one secondary button.
- `.ds-meter`: a small progress bar; `.ds-meter__bar` with `--full`, `--most`, `--half` or `--some`; `.ds-meter--spaced` when a list follows.
- `.ds-footer`: the control strip that closes the window, in `--fill-bar`: `.ds-footer__controls` (round previous, play and next keys, the play key large), a `.ds-footer__list` of `.ds-footer__link` items and the `.ds-footer__legal` line at the right.
- `.ds-menu`: a glass drop-down under a navigation item (`.ds-menu__list` of `.ds-menu__link`), opened on hover or focus, never forced open.
- `.ds-progress`: the glossy accent progress bar, a native `<progress>` in a `.ds-progress__row` with label and value; for downloads, installs and storage.
- `.ds-carousel`: a picture viewer with previous and next orbs and dots (`.ds-carousel__stage`, `__track`, `__slide`, `__dots`); static, dots are not links.

## Never

- `border-radius <= 12px`: corners are 3, 6, 8 or 12px; only round keys and capsules are rounder.
- `border-width <= 1px`: every border and rule is a hairline.
- `box-shadow-blur <= 30px`: shadows are short; only the floating window and the front cover reach 30px.
- `gradient-fills <= 30%`: gloss belongs to bars, keys, labels, covers and cells; text areas stay plain.
- `font-size <= 34px`: the stage title is the largest text.
- `font-size >= 11px`: nothing is smaller than the meta text.
- `font-weight <= 700`: bold is the heaviest weight; headings are light.
- `font-families <= 2`: one grotesque, a monospace only for inline code.
- `line-height <= 1.5`: body text is 12px on 18px.
- `underlined-links <= 10%`: links underline on hover only.
- `uppercase-text = 0%`: nothing is uppercased by CSS.
- `row-gap <= 24px`: rows are divided by a hairline, not by space.
- `block-gap <= 30px`: sections are at most 30px apart.
- `content-width <= 960px`: text stays inside the window.
- `palette-colours <= 48`: greys, black, one green, one blue for links, amber and red for status.
- `transition = none`: states change at once.
- `animation = none`: nothing moves.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new bar is `--fill-bar`, a pale strip `--fill-bar-alt`, a dark band `--fill-inverse`; anything round and pressable is dressed like `.ds-button--round`. Keep to the conventions above for which text token sits on which fill, give glossy fills an outline mixed from their colour and `--color-shadow`, take highlights from the surface tokens instead of drawing new ones, put a `--fill-*` token last in a layered background, and make any extra shade with `color-mix()` over existing tokens. A new picture slot takes a `.ds-art` frame; a new list of sources is a `.ds-sidebar`; a new readout is a `.ds-stat`.
