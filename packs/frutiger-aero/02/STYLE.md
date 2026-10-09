# Frutiger Aero start page

## Summary

A fixed-width personal start page in the dark-glass variety of the Frutiger Aero look: a night-blue page lit by cyan and violet aurora glows, a black glass bar, a search band with one glowing field, and three columns of glass gadgets for news, mail, weather and markets. Portals and start pages of this kind were common from about 2005 to 2010, when every provider let visitors arrange small windows of content. The gadgets are translucent panels with a lit top edge and a deep soft shadow, buttons are capsule-shaped gels with a halo, and the type pairs a small Lucida body with heavier Gill Sans headings.

## Layout

- Page width is fixed: `--size-page` = 980px, centred in a 1024px window. It does not stretch and does not reflow for phones.
- The header and the footer span the window; their content sits in a centred `--size-page` box (`.ds-nav__inner`, `.ds-footer__legal-text`). The header is two strips: a small-print strip of `--size-top` (24px) and the glass bar of `--size-nav` (36px).
- Everything else sits in `.ds-page__wrap`, one translucent sheet (90% `--color-canvas`) that hangs from the bar: no top border, `--radius-page` on the two lower corners, 14px padding. Inside it the content is 950px wide.
- The home page is the search band (`.ds-hero`) and under it `.ds-page__gadgets`, three columns in the proportion 1.3 to 1 to 1 with `--space-5` (14px) between columns and between the gadgets of a column. Each `.ds-page__column` is a stack of `.ds-panel` gadgets; the last block of a column may be a `.ds-sidebar`.
- Spacing scale (dense): `--space-1` 2px (hairline gaps), `--space-2` 4px (list row padding, gadget head padding), `--space-3` 6px (gaps between buttons, between grid tiles and forecast tiles), `--space-4` 10px (gadget body padding, control padding, cell padding), `--space-5` 14px (gutter, sheet padding, gap between stacked blocks), `--space-6` 20px (hero padding, gap between main and side column, gap between sections; the largest gap).
- Fixed sizes: controls `--size-control` 22px tall, the search field and its button `--size-control-large` 32px, the search form `--size-search` 420px, channel orbs `--size-orb` 26px, rank chips `--size-rank` 18px, form labels `--size-label` 140px, the dialog `--size-dialog` 400px.
- Inner pages drop the search band and the three columns. In the sheet, top to bottom: `.ds-breadcrumb`, then `.ds-page-header` as a glass strip across the full width (title and one line on the left, buttons on the right), then `.ds-page__columns`: a main column (630px) and a right column of `--size-sidebar` (300px), 20px apart. `.ds-tabs` sit at the top of the main column above the list or table they switch; notices sit above the form; a dialog follows its table on a dimmed strip; the side column holds `.ds-sidebar` blocks, gadgets and at most one empty state.
- Views: the specimen is one start page of five screens, each a `.ds-view` inside the sheet, with header and footer written once around them. `home` is the start page (search band, the News, Channels, Mail, Today, Weather and Markets gadgets, a block of gadgets to add). `news` is the news front (tabs, a roomy list, pagination, most read). `story` is one story (long text, onward links, a traffic gadget with status badges). `mail` is the inbox (the button row in the page header, folder tabs, the table, the deleting dialog, folders, the key to the states, an empty state for drafts). `settings` changes the start page (notices, the form, gadgets to add). Links are plain `href="#name"`; the bar link of the showing view is lit by a `.ds-page:has(#name:target)` rule, the Start link also when no view is named, and the News link also on a story.

## Typography and colour roles

- Two families: `--font-body` and `--font-ui` are "Lucida Grande", "Lucida Sans Unicode", "Lucida Sans", Geneva, Verdana; `--font-heading` is "Gill Sans", "Gill Sans MT", "Trebuchet MS": system faces only, a humanist sans with Trebuchet MS where Gill Sans is not installed. `--font-mono` (Monaco) is for inline code.
- `--text-base` 12px at `--line-body` 1.5. `--text-small` 10px for times, tools, stat labels and the footer. `--text-ui` 12px for buttons, fields and tabs; buttons and bar links are bold.
- Headings are semibold (`--weight-heading`, `--weight-display` 600): `--text-display` 40px for the greeting in the search band, tightened by `--display-tracking` -0.5px; `--text-h1` 26px page titles and forecast figures; `--text-h2` 18px the brand name and empty-state titles; `--text-h3` 14px (bold) gadget titles, lead headlines, section titles. `--text-large` 16px for the lead and the search field. Nothing is uppercased.
- Links are not underlined at rest and underline on hover; the scopes and suggestions in the search band are the exception and show where the search goes.
- `--color-page` #060c1a with `--fill-page`: night blue with a cyan glow at the upper left and a violet one at the upper right.
- `--color-bar` #181818 with `--fill-bar` (black glass: a lit upper half ending at 50%) and near-white `--color-bar-text` with a glowing `--shadow-text`: the bar, the current tab, the dialog frame, channel orbs.
- `--color-bar-alt` #1b3d6b with `--fill-bar-alt` (blue glass lit from above): the small-print strip, gadget heads, table head, sidebar titles, the footer link strip.
- `--color-inverse` #123c8f with `--fill-inverse` (aurora: cyan rising from the lower left, violet from the upper right): the search band and the last strip of the footer, white text.
- `--color-text` #c9d6e6 on the dark surfaces `--color-canvas` #0b1526, `--color-surface` #101c33, `--color-surface-alt` #162740, `--color-surface-strong` #203552; `--color-heading` #ffffff; `--color-heading-alt` #62d8ff (cyan) section titles, prose h2, breadcrumb marks; `--color-text-muted` #989898.
- `--color-link` #62c8ff, `--color-link-quiet` #a3b8cf, visited #b59cff, hover #a8ecff, active #ffffff.
- `--color-button` #0f8fd0 with `--fill-button` (a gel capsule: a white sheen over the upper half, a faint second light at the foot) and a halo from `--shadow-control`: the primary button, the brand tile, arrow buttons, the current page number. `--color-button-secondary` #26374f with the same sheen: every other button.
- `--color-accent` #19c2f0 with dark `--color-accent-text`: pill badges and list bullets. `--color-accent-alt` #c86bff (violet): the New marker.
- `--color-fill-1` to `--color-fill-4` (deep blue, teal, violet, sea green): the lower end of forecast tiles and channel tiles.
- Borders: `--color-border` #2a4466 around gadgets, `--color-border-muted` #1c2f4a between rows, `--color-border-strong` #4d7fb3 for the sheet and the 2px dialog frame, `--color-input-border` #35557d.
- `--color-notice` #132c4a; `--color-danger` #e0483c on `--color-danger-surface` #3a1418; `--color-success` #2fa84a and `--color-warning` #f2b31c fill status badges.
- Surface: `--radius-control` 13px (capsule buttons and fields), `--radius-panel` 10px, `--radius-page` 16px, `--radius-pill` 20px. `--shadow-panel` is a lit inner top edge and an 18px drop; `--shadow-control` adds an 8px halo in `--color-focus` that grows to 14px on hover; `--shadow-dialog` is a 40px drop inside a 30px glow. `--fill-panel` is translucent and `--backdrop-blur` 10px blurs what lies behind it. `--transition` none.
- One-off shades are made with `color-mix()` over tokens: outlines of glossy fills (the fill's colour mixed with `--color-shadow`), the lit bar link (`--color-bar-text` with transparent), the sheet (90% `--color-canvas`), the halo of the search field (`--color-focus` with transparent), gloss on status fills (mixed with `--color-button-text`).
- The era's conventions for which text sits on which fill, kept by every layout and every token set of the era:
  - `--color-text`, `--color-text-muted`, `--color-heading`, `--color-heading-alt` and the link colours sit on `--color-canvas`, `--color-surface`, `--color-surface-alt`, `--color-surface-strong`, `--fill-panel` and the four `--color-fill-*` tints. A palette keeps page, canvas, surfaces and fills on one side (all light or all dark) so that these stay readable on every one of them.
  - `--color-bar-text` only on `--fill-bar`; `--color-bar-alt-text` only on `--fill-bar-alt`; `--color-inverse-text` only on `--fill-inverse`. Links on a bar take the bar's text colour, never `--color-link`.
  - `--color-button-text` on `--fill-button` and on any fill made from `--color-danger`, `--color-success` or `--color-accent-alt`; it is always the light colour and also supplies the gloss on those fills (mixed in with `color-mix()`).
  - `--color-accent-text` on `--fill-accent` and on `--color-warning`; `--color-button-secondary-text` on `--fill-button-secondary`; `--color-notice-text` on `--color-notice`; `--color-danger` as text on `--color-danger-surface` and on the surfaces.
  - Nothing but decoration sits on `--fill-page`: no text is set directly on the page.
  - `--color-focus` is the glow colour: the focus ring, the halo of the search field and of a hovered button. `--color-inverse-text` draws the bubbles on the sky band. `--color-shadow` tints every shadow and darkens the outline of every glossy fill.

## Components

- `.ds-page`: on `<body>`. Sets the page fill, base font and colour. `.ds-page__wrap` is the sheet; `.ds-page__gadgets` the three-column grid of the start page with `.ds-page__column` stacks; `.ds-page__columns` holds `.ds-page__main` and `.ds-page__aside` on inner pages. `.ds-section` with a cyan `.ds-section__title` separates blocks by 20px. `.ds-view` is one screen of the example site (`.ds-view--home` the first).
- `.ds-nav`: the header, full window width. `.ds-nav__top` is the small-print strip (`.ds-nav__date`, right-aligned `.ds-nav__tools` of bold `.ds-nav__tool` links); `.ds-nav__bar` the glass bar; each wraps a `.ds-nav__inner`. The bar holds the `.ds-brand`, `.ds-nav__links` of bold `.ds-nav__link` pills (current: `is-current`, a lit translucent pill; hover: `is-hover`) and a right-aligned `.ds-nav__more` line with a `.ds-nav__more-link`.
- `.ds-brand`: the site's mark and name at the left end of the bar, linking home; one per page. `.ds-brand__mark` is a round tile dressed like the primary button (`--fill-button`, a darker outline, `--shadow-control`) with the mark in `--color-button-text`; `.ds-brand__name` is the name at `--text-h2` in `--color-bar-text`. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo: replace both, keep the tile.
- `.ds-hero`: the search band at the top of the start page: a rounded block in `--fill-inverse`. `.ds-hero__main` holds the greeting as `.ds-hero__title` (40px) and one `.ds-hero__lead` line; `.ds-hero__search` on the right holds `.ds-hero__scopes` (links `.ds-hero__scope`, current `is-current`), `.ds-hero__action` (the glowing `.ds-hero__field` and a `.ds-button--large`), and a `.ds-hero__note` of suggestions (`.ds-hero__note-link`). One per site, on the start page only.
- `.ds-page-header`: the head of an inner page instead of the hero: a glass strip with `.ds-page-header__main` (`.ds-page-header__title`, 26px, the page's h1, and a grey `.ds-page-header__text` line) and `.ds-page-header__actions` at the right. One per inner page.
- `.ds-panel`: the gadget (`is-last` on the last of a column): `.ds-panel__head` is a glass strip with `.ds-panel__title`, an optional badge and right-aligned `.ds-panel__tools` (small `.ds-panel__tool` links such as Edit and x); `.ds-panel__body` holds the content, `.ds-panel__text`, a `.ds-panel__actions` row of buttons, a `.ds-panel__go` row (a field and its button) or a `.ds-panel__more` link.
- `.ds-stat`: one forecast tile; four sit in a `.ds-stats` row inside the Weather gadget. Each has a small `.ds-stat__label` (the day) over a 26px `.ds-stat__number` and a `.ds-stat__note`. `.ds-stats__place` above the row names the place, with `.ds-stats__now` for the current reading. Use the row for any three to five figures of one kind inside a gadget.
- `.ds-grid`: channel tiles, four to a row (`.ds-grid--wide` for a single row in a side column). Each `.ds-grid__cell` is a link (`--1` to `--4` choose the tint it fades into; hover: `is-hover`) with a `.ds-orb` over its name. `.ds-orb` is a 26px glossy ball with a line icon: `--fill-bar` by default, `.ds-orb--accent`, `.ds-orb--button`.
- `.ds-tabs`: a row of small pills over a hairline: `.ds-tabs__tab`, the current one `is-current` (a piece of the glass bar), `.ds-tabs__count` for a number. Inside a gadget or at the top of the main column.
- `.ds-list`: tight headline list. Each `.ds-list__item` (first: `is-first`) starts with a round bullet, then `.ds-list__body` with `.ds-list__title` and an optional `.ds-list__text`, then a right-aligned `.ds-list__meta` (a time or a badge). `is-lead` makes the first headline a heading; `.ds-list--roomy` is the full-page form with heading-size titles and 10px rows.
- `.ds-button`: the primary action: a blue gel capsule with a halo, white bold 12px, 22px tall. Variants: `.ds-button--secondary` (dark glass, the usual button), `.ds-button--danger` (red gel, only for deleting), `.ds-button--large` (32px, the search button), `.ds-button--arrow` (a round button holding only a `.ds-button__icon` arrow, before the link or after the field it belongs to). States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttons` is a row; `.ds-buttons--bar` a toolbar. One primary button per gadget, form or dialog.
- `.ds-sidebar`: a link block (`is-last` on the last of a column): `.ds-sidebar__title` on a glass strip, `.ds-sidebar__list` of `.ds-sidebar__item` rows (first: `is-first`) divided by hairlines, each with an optional `.ds-sidebar__rank` chip before the link and a right-aligned `.ds-sidebar__count`.
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
- `.ds-footer`: full window width: `.ds-footer__links`, a glass strip of centred small links (`.ds-footer__item`, first `is-first`, hairlines between, `.ds-footer__link`), then `.ds-footer__legal`, an aurora strip holding `.ds-footer__legal-text`, one centred page-width line of small print.
- `.ds-menu`: a glass drop-down under a navigation item (`.ds-menu__list` of `.ds-menu__link`), opened on hover or focus, never forced open.
- `.ds-progress`: the glossy accent progress bar, a native `<progress>` in a `.ds-progress__row` with label and value; for downloads, installs and storage.
- `.ds-carousel`: a picture viewer with previous and next orbs and dots (`.ds-carousel__stage`, `__track`, `__slide`, `__dots`); static, dots are not links.
- `.ds-avatar`: the framed account picture drawn as initials, beside a name in `.ds-person`; `--small` for tight places.

## Never

- `border-radius <= 16px`: corners are 10, 13 or 16px; only orbs, bullets and pills are rounder.
- `border-width <= 2px`: hairlines everywhere, 2px only around the dialog.
- `box-shadow-blur <= 40px`: glows and drops are soft, the dialog's 40px is the deepest.
- `gradient-fills <= 28%`: gloss belongs to bars, gadget heads, buttons, orbs, pills and tiles; text areas stay plain.
- `font-size <= 40px`: the greeting is the largest text.
- `font-size >= 10px`: nothing is smaller than the meta text.
- `font-weight <= 700`: bold is the heaviest weight.
- `font-families <= 3`: a body sans, a heading sans, a monospace for inline code.
- `line-height <= 1.5`: body text is 12px on 18px.
- `underlined-links <= 12%`: only the search band's scopes and suggestions are underlined at rest.
- `letter-spacing <= 0.5px`: only the greeting is tightened, by half a pixel; nothing is spaced out.
- `uppercase-text = 0%`: nothing is uppercased by CSS.
- `row-gap <= 20px`: list rows are divided by a hairline, not by space.
- `block-gap <= 20px`: gadgets are 14px apart, sections at most 20px.
- `content-width <= 980px`: text stays inside the fixed column.
- `palette-colours <= 46`: night blues, cyan, violet, a green, amber, red and greys.
- `transition = none`: states change at once.
- `animation = none`: nothing moves.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. Anything that shows live content is a gadget: a `.ds-panel` with a head strip, tools and a body, placed in a column. A new bar is `--fill-bar`, a strip `--fill-bar-alt`, a band `--fill-inverse`; anything round and pressable is dressed like `.ds-orb` or `.ds-button--arrow`. Keep to the conventions above for which text token sits on which fill, give glossy fills an outline mixed from their colour and `--color-shadow`, take glows from `--color-focus`, and make any extra shade with `color-mix()` over existing tokens.
