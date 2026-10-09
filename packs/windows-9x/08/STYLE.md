# Desktop grey disc player shop, night, 1999

## Summary

A mail-order record shop laid out as the disc player program of a late 1990s desktop: one narrow window whose deck carries a display, a basket read-out and a row of numbered preset keys, with everything under it in full-width bands, the disc of the week in a large display first. This pack shows it in a night scheme: charcoal windows on a black squared desktop, a red title bar with a sheen, a pale display with dark lettering, amber selection and lamps, slab keys lit along the top edge and shaded along the bottom, and tall condensed capitals over a narrow sans. Music shops, radio stations and multimedia CD-ROM menus dressed themselves as the player on the visitor's own desktop from 1996 to 2001.

## Layout

- Fixed, not fluid: `--size-page` is 720px, one window (`.ds-page__window`, also a `.ds-window`) centred on the desktop of an 800px screen with a `--space-6` margin above and below.
- The specimen is one example site, the shop, of five views. `.ds-nav` (title bar, deck and preset keys) and `.ds-footer` (status bar) are written once, inside the window and outside the views; each view is a `.ds-view` and one shows at a time.
- Navigation is at the top, in three rows: the title bar; the deck `.ds-nav__deck` with the display at the left and the basket read-out with its key at the right; the row of five preset keys `.ds-nav__links`, equal in width, each a number and a lamp over a name.
- Every view is `.ds-page__body`: a single column of full-width bands `--space-4` apart. There is no side column, except in the catalogue.
- Home view: the `.ds-hero` (the large display: a sleeve, the title, a lead and one line, two keys stacked at the right), the `.ds-stat` row of four counters, a `.ds-heading`, the `.ds-grid` of four discs in their cases, and a `.ds-split` of two `.ds-panel`s: the best sellers as a `.ds-list` and how the shop works.
- Inner pages keep the title bar, the deck, the preset keys and the status bar and drop the large display, the counters and the cases. They open with the `.ds-breadcrumb` (small links), then the `.ds-page-header`: a raised band with the title and one line at the left and one key at the right. Tabs come after the page header, straight above the list view they filter; there is no sheet.
- Views: `home` is new this week. `catalogue` is one shelf: under the page header, `.ds-page__columns` puts the `.ds-sidebar` list box of shelves (`--size-side`, 148px) beside `.ds-page__main` with the band keys, the list view of discs and the pagination keys beside the key of stock labels. `disc` is one disc: `.ds-split--cover` with the sleeve (`--size-cover`, 208px) over a panel of label details at the left, the track list and the review in a well at the right. `basket` is two notices, the list view of the basket, the row of keys, then the empty state beside the confirmation dialog and a panel. `order` is `.ds-split--wide`: the form in a group box beside a panel with the sum. The preset key of the view that is showing stays down, its name bold, with its lamp lit; key 1 is marked when the address names no view.
- Fixed sizes: `--size-pos` 20px position field of the chart; `--size-ctl` 16px caption button and `--size-glyph` 8px glyph (a lamp is two glyphs wide); `--size-icon` 32px and `--size-icon-small` 16px icons (the sleeve of the large display is four icons wide); `--size-check` 13px check box and radio button; `--size-label` 104px form label column; `--size-field` 200px text field; `--size-dialog` 316px, also the narrow column of `.ds-split--wide`; `--size-button` 75px least key width; `--size-block` 8px progress block.
- Spacing scale: `--space-0` 1px is the hairline step; `--space-1` 2px the gap between preset keys and between a counter and its label; `--space-2` 4px the padding of keys and the height of a lamp; `--space-3` 6px cell and field padding, the padding of the deck and the gap between keys in a row; `--space-4` 8px the gap between bands, between cases and between columns; `--space-5` 12px the padding of the large display, panels and dialogs; `--space-6` 16px the margin of the desktop round the window and the gap inside the large display.

## Typography and colour roles

- `--font-body` and `--font-ui` are the narrow sans at `--text-base` and `--text-ui` (13px); keys are `--weight-ui`. `--font-heading` is the heavy condensed face, set in capitals by `--heading-transform` with `--heading-tracking`: the title of the large display (`--text-display`, 32px, `--display-tracking`), page titles (`--text-h1`, 22px), panel and section titles (`--text-h2`, 16px) and the brand name. `--font-mono` is the lettering of displays: both lines of the deck display, the small line of the large display, counters, key numbers and chart positions. `--text-small` (11px) is for artists, weeks, counts and the breadcrumb. Bold is `--weight-bold`; nothing is lighter than `--weight-body`.
- Links in text are `--color-link`, underlined by `--link-decoration`, and lose the line under the pointer (`--link-decoration-hover`); quiet links and the rows of the list box use `--link-decoration-quiet`.
- Which text sits on which fill, the same in every layout of the era:
  - The window face is `--color-canvas` (`--fill-panel` for windows, cases and dialogs) and carries `--color-text`, `--color-text-muted`, `--color-heading`, `--color-heading-alt` and the link colours.
  - A sunken well (panels, the list box, list views, the basket read-out, the review, the empty state) is `--color-surface` and carries the same text and link tokens. `--color-surface-alt` is the alternate row, a preset key held down, inline code and the position field of the chart; `--color-surface-strong` is a total row and a lamp that is off. They carry `--color-text`.
  - A display (the deck display, the large display, the figure of a counter) is `--fill-inverse` with `--color-inverse-text` and nothing else: no heading, muted or link colour is used on it.
  - The title bar is `--fill-bar` with `--color-bar-text`. The deck, the page header band and column heads are `--fill-bar-alt` with `--color-bar-alt-text`.
  - The selection (the current band key, the open shelf, a selected row, a lit lamp, progress blocks, the plain badge) is `--fill-accent` with `--color-accent-text`; links and icons inside a selection inherit that colour.
  - Keys are `--fill-button` with `--color-button-text` (`--fill-button-hover` and `--color-button-hover-text` under the pointer); the lesser key and the band keys are `--fill-button-secondary` with `--color-button-secondary-text`; a disabled key is `--color-disabled` with `--color-disabled-text`; the destructive key is `--fill-button` with a bold label in `--color-danger`.
  - Fields are `--fill-input` with `--color-input-text`; `--color-input-border` outlines the check box and the radio button.
  - `--color-fill-1` to `--color-fill-4` are the sleeves and carry `--color-text` (the drawing on them).
  - Prices are `--color-heading-alt`, bold. Notices are `--color-notice` with `--color-notice-text` inside a line of `--color-border-strong`; the error notice is `--color-danger` on `--color-danger-surface`.
  - Status labels print `--color-success`, `--color-warning` and `--color-danger` as text and frame on `--color-surface`.
  - Icons are inline SVG painted with `currentColor`: `--color-heading-alt` by default, the surrounding text colour on keys, sleeves and selections.
- Edges. Nothing has a border of its own except thin lines; every raised or sunken edge is a `box-shadow` token drawn inside the box:
  - `--shadow-panel` frames the window, `--shadow-dialog` a dialog. Both also throw a hard shadow.
  - `--shadow-control` is every raised part: buttons, preset keys, band keys, caption buttons, the page header band, column heads, the cases of the discs, pagination keys, the brand tile. `--shadow-control-hover` replaces it under the pointer.
  - `--shadow-control-pressed` is everything sunken: displays, counters, fields, wells, panels, the list box, status cells, the check box, the progress bar, a key held down, the current band key and pagination key.
  - `--border-width-strong` is the thickness of those edges: padding that must clear an edge is written `calc(var(--border-width-strong) + var(--space-N))`.
  - `--border-width`, `--border-style` and `--color-border` draw the thin lines: the line under a panel title, the frame of a lamp and of a chart position, group boxes, the etched rule (over `--color-border-muted`); a sleeve is framed in `--color-border-strong`; table row lines are `--color-surface-alt`.

## Components

The name `chronoskin` and its mark are placeholders for the installing project's own name and mark.

- `.ds-page`: the desktop, on `<body>`; `.ds-page__window` is the one window, `.ds-page__body` the column of bands of a view. `.ds-page__columns` with `.ds-page__main` is the one two-column arrangement, for a list box beside a list view.
- `.ds-window`: a raised frame; `.ds-window__title` is its title bar (`--inactive` for a window without the focus) with `.ds-window__caption` and the `.ds-window__ctl` caption buttons (`--min`, `--max`, `--close`).
- `.ds-nav`: the title bar, the deck and the preset keys. `.ds-nav__deck` holds `.ds-nav__display` (a `.ds-nav__small` line over a `.ds-nav__line`) and `.ds-nav__basket` (a sunken read-out with a secondary key). `.ds-nav__links` is the row of keys; each `.ds-nav__item` an equal share, `.ds-nav__link` a key with `.ds-nav__num` (the number and the `.ds-nav__lamp`) over `.ds-nav__name`. The current key is held down and its lamp is lit.
- `.ds-brand`: the mark and the name at the left of the title bar. `.ds-brand__mark` is a small key dressed like the default button, `.ds-brand__name` the name in the heading face.
- `.ds-hero`: the large display, home view only. `.ds-hero__sleeve` (a `.ds-sleeve`) at the left, `.ds-hero__body` with `.ds-hero__label` (monospace small print), `.ds-hero__title`, `.ds-hero__lead` and `.ds-hero__text`, and `.ds-hero__action`, a column of keys, at the right.
- `.ds-page-header`: the head of an inner page, a raised band: `.ds-page-header__text` with `.ds-page-header__title` and `.ds-page-header__lead`, and `.ds-page-header__action` at the right.
- `.ds-view`: one screen of the example site; `.ds-view--home` shows when the address names none.
- `.ds-prose`: running text in a well: `h1`, `h2`, `h3`, `p`, `ul`, `ol`, `li`, `strong`, `code`; `.ds-prose__lead` is the lead line.
- `.ds-link`: every link in text. States `is-visited`, `is-hover`, `is-active`, `is-focus`; `.ds-link--quiet` for tool and footer links, `.ds-link--strong` for a bold one. `.ds-links` wraps a line of links.
- `.ds-button`: a raised key at least `--size-button` wide; the plain class is the default key, inside one more line. `.ds-button--secondary` is every other key, `.ds-button--danger` the destructive one, `.ds-button--large` a bigger one. States `is-hover`, `is-pressed`, `is-focus`, `is-disabled` or `disabled`. `.ds-buttons` lays out a row; a `.ds-find__fill` spacer pushes keys apart.
- `.ds-form`: two columns, `.ds-form__label` and `.ds-form__field`. Controls: `.ds-form__input` (`--short`, `--wide`), `.ds-form__select`, `.ds-form__textarea`, `.ds-form__check` and `.ds-form__radio` inside a `.ds-form__option`. `.ds-form__hint` is a small line under a field, `.ds-form__error` a bold line in the danger colour, `.ds-form__actions` the row of keys.
- `.ds-table`: a list view: a well whose `.ds-table__head` cells are raised (`--num` right-aligned). `.ds-table__row` (`--alt` tinted, `is-selected` in the selection colours), `.ds-table__cell` (`--num`, `--group` for a total row), `.ds-table__caption`.
- `.ds-list`: a chart. Each `.ds-list__item` is a `.ds-list__pos` (the position in a small framed field), a `.ds-list__text` and a `.ds-list__meta`.
- `.ds-panel`: a sunken well with a `.ds-panel__title` in the heading face over a thin line; `.ds-panel__text` is a paragraph in it, `.ds-panel__body` a plain holder. `.ds-panel--fill` takes spare height.
- `.ds-stat`: the row of counters; each `.ds-stat__item` is a `.ds-stat__value` (a small display, monospace, right-aligned) over its `.ds-stat__label`.
- `.ds-grid`: discs in their cases, four to a row. A `.ds-grid__cell` is a raised card: a `.ds-sleeve`, `.ds-grid__title`, `.ds-grid__text` and a `.ds-grid__foot` with the `.ds-price` and a badge.
- `.ds-sleeve`: a square of block colour with a drawing on it (`--2` to `--4` take the other block colours).
- `.ds-tabs`: band keys joined in a row: `.ds-tabs__item` holds a `.ds-tabs__tab`; the current one (`is-current`) is held down in the selection colours. `.ds-tabs__fill` and `.ds-tabs__note` put a small count at the right end.
- `.ds-badge`: a small bold label in the selection colours. `.ds-badge--count` is a plain muted count. Status: `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger`. `.ds-badges` lays out a row.
- `.ds-sidebar`: a list box: `.ds-sidebar__title` over the sunken `.ds-sidebar__list`; each `.ds-sidebar__link` is a row with a `.ds-sidebar__count` at the right, the open one `is-current`.
- `.ds-notice`: the tip box inside a dark line; `.ds-notice__label` is the bold lead word, `.ds-notice--error` the error form. `.ds-notices` stacks several.
- `.ds-pagination`: a row of small keys: `.ds-pagination__label`, then `.ds-pagination__link`s; `is-current` is pressed and bold, `is-disabled` greyed.
- `.ds-breadcrumb`: the path in small type, the `.ds-breadcrumb__item`s parted by a slash; the last is `is-current`.
- `.ds-dialog`: a small window `--size-dialog` wide, shown where it would open: `.ds-dialog__title`, `.ds-dialog__content` with a large icon and the `.ds-dialog__body`, `.ds-dialog__actions`.
- `.ds-empty`: an empty well with a large icon, `.ds-empty__title`, `.ds-empty__text` and one secondary key.
- `.ds-footer`: the status bar, a `.ds-status` row of sunken `.ds-status__cell`s (`--wide` takes the free width).
- `.ds-icon` (`--large`, `--plain`), `.ds-well`, `.ds-rule`, `.ds-group` with `.ds-group__title`, `.ds-props` with `.ds-props__label` and `.ds-props__value`, `.ds-progress` with `.ds-progress__block`: the small parts every layout of the era shares.
- `.ds-heading`: a section title. `.ds-split`: two equal columns of the same height (`--cover` and `--wide` are the two uneven forms). `.ds-stack`: a column of blocks (`--tight` closer, `--fill` takes spare height). `.ds-find`: a row of a label, a field and keys.

## Never

- `border-radius <= 0px`: the desktop had no rounded corner; the radio button is a circle.
- `box-shadow-blur <= 0px`: edges are hard lines and the shadow of a window is hard; nothing is blurred.
- `text-shadow = none`: no text shadow anywhere.
- `gradient-fills <= 20%`: keys, bars, displays and lamps are shaded; wells, panels and the window face are flat.
- `border-width <= 3px`: thin lines are 1px; nothing is drawn thicker than the 3px edge of a key.
- `font-size >= 11px`: artists and counts are 11px, nothing smaller.
- `font-size <= 32px`: the title in the large display is the largest text; page titles stop at 22px.
- `font-weight >= 400`: only regular and bold.
- `font-families <= 3`: the narrow sans, the condensed heading face and a monospace.
- `letter-spacing <= 1px`: headings are tracked by half a pixel, the largest title by one.
- `uppercase-text <= 12%`: only headings are set in capitals.
- `row-gap <= 8px`: rows of lists and list views touch or sit 1px to 4px apart.
- `block-gap <= 16px`: bands inside the window are 8px apart, never more than 16px.
- `content-width <= 720px`: everything is inside the one 720px window.
- `transition = none`: nothing eases on hover.
- `animation = none`: no CSS animation.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new page is a column of full-width bands under a `.ds-page-header`; anything that holds content is a `.ds-panel` or a `.ds-well`; anything that lists is a `.ds-table` or a `.ds-list`; anything that reads out a figure or a line of status is a display on `--fill-inverse` with `--color-inverse-text` and the monospace face. Anything pressed or typed into takes `--shadow-control-pressed`, anything pushed takes `--shadow-control`. Clear every edge with `calc(var(--border-width-strong) + ...)` padding, keep navigation in the preset keys and out of side columns, and never add a radius, a blurred shadow, a gradient outside the `--fill-*` tokens, a hover animation or a colour of your own.
