# Frutiger Aero travel journal

## Summary

A fixed-width travel journal in the meadow variety of the Frutiger Aero look: a grass-green page crossed by sunbeams, with the name on a deep green plate, a row of separate glass keys for navigation, and every block of the page a crisp white card afloat on the green. Personal journals and photo diaries of this kind were common from about 2005 to 2010, when a trip was told day by day in words and digital photos and friends answered in a guestbook. The band of sky at the top carries three photo prints tossed at an angle, cards are opaque with a white inner rim and a contact shadow, corners are nearly square (2 to 4px), and the type pairs an Arial body with regular Georgia headings and small bold Verdana controls.

## Layout

- Page width is fixed: `--size-page` = 920px, centred in a 1024px window. It does not stretch and does not reflow for phones.
- There is no sheet. `.ds-page__wrap` (920px) is a stack of separate cards on the page fill, `--space-5` (20px) apart: the header row, the cards of the showing view, the footer card. No text sits on the page itself.
- The header row is three loose pieces: the name plate `.ds-nav__plate` at the left, the navigation keys `.ds-nav__links` pushed to the right (each `--size-nav` 30px tall), and the search field (`--size-search` 140px).
- The home page runs: the hero band across the full width; then `.ds-page__columns`, a main column (610px) holding one `.ds-page__card` with the entries, and a right column `.ds-page__aside` of `--size-side` (290px) holding the figures, a panel and a link card; then one full-width card with the four albums.
- Spacing scale: `--space-1` 3px (picture frames, gaps under titles), `--space-2` 5px (inside controls, between keys, fact rows), `--space-3` 9px (gaps between buttons, table cells, comment rows), `--space-4` 14px (panel padding, list rows, gaps between album cells), `--space-5` 20px (card padding, gutter, gap between cards), `--space-6` 30px (hero padding, the dimmed strip; the largest gap).
- Fixed sizes: controls `--size-control` 26px tall, hero buttons `--size-control-large` 38px, the brand gel `--size-mark` 30px, list pictures `--size-thumb` 112px, hero prints `--size-print` 172px in a block of `--size-prints` 330px, comment tiles `--size-avatar` 32px, form labels `--size-label` 120px, the dialog `--size-dialog` 420px.
- Inner pages drop the hero. Top to bottom: one head card (`.ds-page__card--head`) with the `.ds-breadcrumb` and under it the `.ds-page-header` (title and one line on the left, one or two buttons on the right); then `.ds-page__columns` with the same two columns: cards in the main column, panels, link cards and at most one empty state at the side. `.ds-tabs` sit at the top of a card directly above the text or table they switch; pagination closes the card of its table; notices sit above the form inside its card; a dialog follows the columns across the full width on its dimmed strip.
- Views: the specimen is one journal of four screens, each a `.ds-view` between the header row and the footer card, which are written once. `home` is the journal front (hero with prints, latest entries, the trip in figures, where we are, where we have been, albums). `entry` is one entry (tabs, the text with a strip of pictures, comments, the leg of the trip, more entries, the day's photos, onward links). `archive` lists a part of the trip (tabs, table with status labels, pagination, the key to the labels, an empty state for the part not yet begun). `guestbook` is the signing form (notices, the form with its buttons, latest messages, ways to follow, the clearing dialog). Links are plain `href="#name"`; the key of the showing view takes the primary button's fill by a `.ds-page:has(#name:target)` rule, the Journal key also when no view is named.

## Typography and colour roles

- Three families and a monospace: `--font-body` is Arial, `--font-heading` Georgia, `--font-ui` Verdana, all system faces of the period; `--font-mono` (Courier New) only for a file name in running text.
- `--text-base` 13px at `--line-body` 1.5. `--text-small` 11px for meta lines, hints, captions and the footer. `--text-ui` 11px bold Verdana (`--weight-ui` 700) for keys, buttons, tabs and form labels.
- Headings are regular-weight Georgia: `--text-display` 38px hero title, tightened by `--display-tracking` -0.5px; `--text-h1` 27px page titles, figures and the brand name; `--text-h2` 20px card titles and entry titles; `--text-h3` 15px panel, album, dialog and empty-state titles. `--text-large` 15px hero lead and the large button. Nothing is uppercased.
- Links in running text and plain lists are underlined (`--link-decoration` underline); titles (`.ds-link--title`, bold), quiet links, keys and tabs are not.
- `--color-page` #5aa51c with `--fill-page`: grass green, lit from the top, with two slanted sunbeams, paler at the foot.
- `--color-bar` #3f7d12 with `--fill-bar` (satin: a bright line at the top, light to dark without a break) and white `--color-bar-text` under `--shadow-text`: the name plate.
- `--color-bar-alt` #d6e9a6 with `--fill-bar-alt`: the figures card, table heads, the dialog's title strip.
- `--color-inverse` #1c6fb0 with `--fill-inverse` (sky blue, darker at the foot, the same sunbeams): the hero and the footer card, white text.
- `--color-heading` #2f6b0a card and page titles; `--color-heading-alt` #1d74b8 (sky) panel and link-card titles, prose h2, breadcrumb marks. `--color-text` #33382b, `--color-text-muted` #687058.
- `--color-link` #1b6fb5, `--color-link-quiet` #55703a, visited #7a5a9a, hover #e0561a, active #2f6b0a.
- `--color-button` #1e88d2 with `--fill-button` (sky glass): the primary button, the brand gel, the current key and page number. `--color-button-secondary` #e3ead3 with `--fill-button-secondary`: every other key and button, comment tiles.
- `--color-accent` #f26a2e (poppy) with `--fill-accent`: labels and the stroke under the current tab. `--color-accent-alt` #9a52c4: the New marker and the dusk sky of the town picture.
- `--color-fill-1` to `--color-fill-4` (lime, sky, sun, petal): the lower end of each album cell's fade and the skies and sands of the pictures; never text or borders.
- Pictures are drawn from the palette: `--color-inverse` sky and sea, `--color-success` hills, `--color-warning` sun, `--color-text-muted` rock, `--color-shadow` roofs. They stand in for the project's own photographs.
- Borders are 1px solid, 3px (`--border-width-strong`) for the dialog's rim, the stroke at the left of a notice and under the current tab. `--color-border` #c4dc9a, `--color-border-muted` #e2edc9 inside a card, `--color-border-strong` #5c9a2a for the dialog and (at 60%) the cards.
- `--color-notice` #e3f1fb for notices; `--color-danger` #b81d2a on `--color-danger-surface` #fbe7e6 for errors; `--color-success` #2e8b3d and `--color-warning` #e8a200 fill the status labels.
- Surface: `--radius-control` 2px, `--radius-panel` 3px, `--radius-pill` 2px, `--radius-page` 4px. `--shadow-panel` is a white inner rim and a short contact shadow under the card; `--shadow-dialog` the deeper one under prints and the dialog; a hovered key gains a rim and glow in `--color-focus`. `--fill-panel` is opaque and `--backdrop-blur` 0px: nothing is frosted. `--transition` none.
- One-off shades are made with `color-mix()` over tokens: outlines of glass fills (the fill's colour mixed with `--color-shadow`), the cards (92% `--color-canvas`), the notice border, roofs and rock in the pictures, gloss on danger, success and warning fills (mixed with `--color-button-text`). Prints and the blank print of the empty state are turned by a few degrees with `transform`.
- The era's conventions for which text sits on which fill are kept: body, heading and link colours on canvas, surfaces, `--fill-panel` and the four tints; `--color-bar-text` only on `--fill-bar`, `--color-bar-alt-text` only on `--fill-bar-alt`, `--color-inverse-text` only on `--fill-inverse`; `--color-button-text` on `--fill-button` and on fills made from `--color-danger`, `--color-success` or `--color-accent-alt`; `--color-accent-text` on `--fill-accent` and on `--color-warning`; `--color-button-secondary-text` on `--fill-button-secondary`; `--color-notice-text` on `--color-notice`. No text is set directly on the page.

## Components

- `.ds-page`: on `<body>`. `.ds-page__wrap` is the 920px column; `.ds-page__card` a white card (`--head` for the head card of an inner page, `is-last` on the last card of a column); `.ds-page__columns` holds `.ds-page__main` and `.ds-page__aside`. `.ds-section__title` (20px over a hairline) heads a card, `.ds-section__more` closes one with a link. `.ds-view` is one screen (`.ds-view--home` the first).
- `.ds-nav`: the header row. `.ds-nav__plate` holds the `.ds-brand`; `.ds-nav__links` is the row of `.ds-nav__link` keys (current: `is-current`, the primary button's fill; hover: `is-hover`); `.ds-nav__search` holds `.ds-nav__input`. Three to five keys.
- `.ds-brand`: the site's mark and name on the plate, linking home; one per page. `.ds-brand__mark` is a round gel dressed like the primary button; `.ds-brand__name` is the name at `--text-h1` in `--color-bar-text`. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo: replace both, keep the gel.
- `.ds-hero`: the journal's opening band of sky, home page only. `.ds-hero__main` holds a small `.ds-hero__kicker`, `.ds-hero__title`, `.ds-hero__lead` and `.ds-hero__action` (a `.ds-button--large` and a secondary one). `.ds-hero__prints` on the right holds three `.ds-print` frames (`--1` to `--3` place and turn them), each a `.ds-picture` over a `.ds-print__caption`.
- `.ds-picture`: a 4 to 3 drawing that stands in for a photograph: `--hills`, `--sea`, `--peak`, `--town`. Used in prints, `.ds-list__thumb`, album cells and `.ds-strip`.
- `.ds-page-header`: the head of an inner page instead of the hero, inside the head card under the breadcrumb: `.ds-page-header__main` with `.ds-page-header__title` (27px) and one grey `.ds-page-header__text` line; `.ds-page-header__actions` at the right. One per inner page.
- `.ds-stat`: one figure, a 27px `.ds-stat__number` over an 11px `.ds-stat__label`; four sit two by two in `.ds-stats`, a pale card at the top of the side column, divided by hairlines.
- `.ds-grid`: four album cells in a row, 14px apart. `.ds-grid__cell` (`--1` to `--4` choose the tint it fades into) frames a `.ds-picture` over a `.ds-grid__title` and a grey `.ds-grid__text` line.
- `.ds-prose`: long text: h1, h2 (sky blue), h3, p, ul, ol, code, strong and links inside it are styled. `.ds-strip` sets three framed pictures (`.ds-strip__frame`) in a row inside it; `.ds-strip--pair` sets them two to a row in a panel.
- `.ds-link`: any link. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--title` bold and not underlined, `.ds-link--quiet` for minor links. `.ds-links` lays several out in a row.
- `.ds-button`: the primary action: sky glass, white bold 11px, 26px tall, 2px corners; may end with a `.ds-button__icon`. Variants: `.ds-button--secondary` (pale green glass, the usual button), `.ds-button--danger` (red glass, only for an action that removes or clears), `.ds-button--large` (38px, the hero). States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttons` is a row. One primary button per card or dialog.
- `.ds-form`: label-left form: `.ds-form__row` with a bold `.ds-form__label` and a `.ds-form__field` holding `.ds-form__input`, `.ds-form__select` or `.ds-form__textarea`; `.ds-form__hint` and `.ds-form__error` under a control; `is-error`, `is-focus`, `is-disabled` on the control; `.ds-form__check` with `.ds-form__checkbox`; `.ds-form__actions` for the buttons, under a hairline and aligned with the fields.
- `.ds-table`: data table in a 1px border: `.ds-table__head` cells form a pale strip; `.ds-table__row` with `is-alt` on alternate rows; `.ds-table__cell` with a hairline above; `is-numeric` right-aligns, `is-date` is small and grey; `.ds-table__note` is a second grey line.
- `.ds-list`: the list of entries. Each `.ds-list__item` is a framed `.ds-list__thumb` picture beside `.ds-list__body`: a 20px `.ds-list__title`, a grey `.ds-list__meta` line and a `.ds-list__text`. Rows are divided by hairlines.
- `.ds-panel`: a card in the side column (`is-last` on the last block): sky-blue `.ds-panel__title`, `.ds-panel__body` with `.ds-panel__text`, a `.ds-panel__facts` list of `.ds-panel__fact` rows (name in `.ds-panel__fact-name`, value at the right) and `.ds-panel__actions`.
- `.ds-tabs`: words on a hairline; `.ds-tabs__tab`, the current one `is-current` on a 3px stroke in `--color-accent`. Two to five tabs at the top of a card.
- `.ds-badge`: a small glass label, poppy by default. `.ds-badge--count` (flat pale), `.ds-badge--new` (violet). Status variants: `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger`. `.ds-badges` lays several out in a row.
- `.ds-sidebar`: a link card in the side column (`is-last` on the last): `.ds-sidebar__title` on a pale strip, `.ds-sidebar__list` of `.ds-sidebar__item` rows divided by hairlines (current: `is-current`), with an optional right-aligned `.ds-sidebar__count`.
- `.ds-comment`: one comment or guestbook message: a `.ds-comment__avatar` tile with an initial, then `.ds-comment__body` with a bold `.ds-comment__name`, a grey `.ds-comment__date` and the `.ds-comment__text`.
- `.ds-notice`: a slip with a 3px stroke at its left: `.ds-notice__text` and a bold `.ds-notice__title`. Variant `.ds-notice--error`.
- `.ds-pagination`: separate small keys: `.ds-pagination__link`, `is-current` (sky glass), `is-disabled`, `.ds-pagination__gap` for an ellipsis. Newer to the left, older to the right.
- `.ds-breadcrumb`: 11px trail of `.ds-link` items with sky-blue `.ds-breadcrumb__sep` guillemets and a bold `.ds-breadcrumb__current`.
- `.ds-dialog`: a card with a 3px rim: `.ds-dialog__title` on a pale strip (with a small red `.ds-dialog__close`), `.ds-dialog__body` with `.ds-dialog__text`, and `.ds-dialog__actions` with the buttons at the right. It casts `--shadow-dialog` and sits on `.ds-dialog__backdrop`, a dimmed strip in the page flow.
- `.ds-empty`: a pale card with a blank, tilted `.ds-empty__print`, `.ds-empty__title`, `.ds-empty__text` and one button.
- `.ds-footer`: the last card, a band of sky: the `.ds-footer__legal` line on the left and a `.ds-footer__list` of `.ds-footer__link` items on the right.
- `.ds-menu`: a glass drop-down under a navigation item (`.ds-menu__list` of `.ds-menu__link`), opened on hover or focus, never forced open.
- `.ds-progress`: the glossy accent progress bar, a native `<progress>` in a `.ds-progress__row` with label and value; for downloads, installs and storage.
- `.ds-carousel`: a picture viewer with previous and next orbs and dots (`.ds-carousel__stage`, `__track`, `__slide`, `__dots`); static, dots are not links.
- `.ds-avatar`: the framed account picture drawn as initials, beside a name in `.ds-person`; `--small` for tight places.

## Never

- `border-radius <= 4px`: corners are 1, 2, 3 or 4px; only the round brand gel is rounder.
- `border-width <= 3px`: hairlines, and 3px for the dialog's rim, notice strokes, the current tab and the blank print.
- `box-shadow-blur <= 20px`: shadows are short contact shadows; only prints and the dialog reach 20px.
- `gradient-fills <= 40%`: gloss belongs to the plate, keys, labels, cells and the drawn pictures; text areas stay plain.
- `font-size <= 38px`: the hero title is the largest text.
- `font-size >= 11px`: nothing is smaller than the meta text.
- `font-weight <= 700`: bold is the heaviest weight; headings are regular.
- `font-families <= 4`: body, heading and control faces, a monospace only for a file name.
- `line-height <= 1.5`: body text is 13px on about 20px.
- `underlined-links <= 45%`: only links in running text, plain lists and the footer are underlined; titles, keys and tabs are not.
- `letter-spacing <= 0.5px`: only the hero title is tightened, by half a pixel; nothing is spaced out.
- `uppercase-text = 0%`: nothing is uppercased by CSS.
- `row-gap <= 30px`: rows are divided by a hairline, not by space.
- `block-gap <= 30px`: cards are 20px apart.
- `content-width <= 920px`: text stays inside the cards.
- `palette-colours <= 64`: greens and limes, sky blues, poppy, violet, amber and red, and the tints of the pictures.
- `transition = none`: states change at once.
- `animation = none`: nothing moves.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new block of content is a `.ds-page__card` in the main column or a `.ds-panel` or `.ds-sidebar` at the side, never text on the page; a dark plate is `--fill-bar`, a pale strip `--fill-bar-alt`, a band of sky `--fill-inverse`; anything pressable is dressed like `.ds-nav__link` or `.ds-button`. Keep to the conventions above for which text token sits on which fill, give glass fills an outline mixed from their colour and `--color-shadow`, frame every picture in a white border with the panel shadow, and make any extra shade with `color-mix()` over existing tokens.
