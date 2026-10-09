# Desktop grey file manager, 1998

## Summary

A shareware catalogue dressed as the file manager of a late 1990s desktop: one silver-grey window on a teal desktop, with a navy title bar, a menu bar with underlined access keys, a toolbar of raised keys, a tree of folders at the left and icon and list views at the right, closed by a status bar of sunken cells. Every edge is a two-step bevel, light at the top left and dark at the bottom right, text is an 11px pixel-sized sans, and nothing is rounded, blurred or animated. Software sites, download libraries and intranet tools borrowed this look from the operating system between 1995 and 2001, when a page that looked like a program felt trustworthy.

## Layout

- Fixed, not fluid: `--size-page` is 760px, one window (`.ds-page__window`, also a `.ds-window`) centred on the desktop of an 800px screen with a `--space-6` margin above and below. Nothing is wider than the window and nothing reflows.
- The specimen is one example site, a software library, of four views. `.ds-nav` (title bar, menu bar, toolbar) and `.ds-footer` (status bar) are written once, inside the window and outside the views; each view is a `.ds-view` and one shows at a time.
- Home view, from the top inside `.ds-page__body` (blocks `--space-4` apart): the `.ds-hero` banner in a white well, the `.ds-stat` row of four read-outs, then `.ds-page__columns`: the `.ds-sidebar` tree pane (`--size-tree`, 184px) at the left and `.ds-page__main` at the right, `--space-2` apart and stretched to the same height. The main column holds a `.ds-pane` with the `.ds-grid` large-icon view, a `.ds-panel` group box with the `.ds-list` of new arrivals, and a `.ds-table` of the most downloaded programs.
- Inner pages keep the title bar, menu bar, toolbar and status bar and drop the hero, the read-outs and the tree. In their place: a `.ds-address` line with the `.ds-breadcrumb` path field directly under the toolbar, then `.ds-page__inner`, one full-width column that opens with the `.ds-page-header` (large icon, name and one line, the action at the right) and continues with the working components `--space-4` apart. Where a page has tabs they sit directly under the page header with their `.ds-sheet`.
- Views: `home` is the catalogue. `program` is one program: page header, a property sheet of four tabs holding the property rows, the status labels, a group box, the list view of files and the download dialog, then the read-me document in a well. `search` is a find window: a group box with the search row, a tip, the result list view with a selected row, the pagination keys and the empty state of a second pane. `submit` is a wizard page: the picture band with the steps at the left (`.ds-wizard`, `--size-band` 132px), notices, the form, an etched rule and the row of keys. The menu entry of the view that is showing is drawn in the selection colours; the home entry is marked when the address names no view.
- Spacing scale: `--space-0` 1px is the hairline step (tick inset, row padding in trees and list boxes); `--space-1` 2px the gap between keys and the padding of bars; `--space-2` 4px the padding of keys, the gap between an icon and its word and between panes; `--space-3` 6px cell and field padding and the gap between keys in a row; `--space-4` 8px the gap between blocks inside a window and the padding of group boxes; `--space-5` 12px the padding of sheets and dialogs and the gap between windows; `--space-6` 16px the margin of the desktop round a window and the gutter of two-column sheets.
- Fixed sizes: `--size-ctl` 16px caption button and `--size-glyph` 8px glyph; `--size-icon` 32px and `--size-icon-small` 16px icons; `--size-check` 13px check box; `--size-label` 116px form label column; `--size-field` 220px text field and `--size-search` 190px search field; `--size-dialog` 350px; `--size-button` 75px least key width; `--size-block` 8px progress block.

## Typography and colour roles

- One family: `--font-body`, `--font-heading` and `--font-ui` are the same pixel-sized system sans (MS Sans Serif, then Tahoma, Geneva, Verdana); `--font-mono` (Courier New) is for code. The references set their text in Arial and Verdana at 10px to 13px; the stack here is the desktop's own face, which the era's pages borrowed with the chrome.
- Sizes: `--text-base` and `--text-ui` 11px for everything; `--text-small` 10px for captions, counts and dates; `--text-large` 13px for a lead line; `--text-h2` 13px and `--text-h3` 11px bold for headings; `--text-h1` 16px for a page title and a read-out; `--text-display` 22px for the hero title only. Line height 1.35 in text, 1.2 in keys and bars. Bold is 700; keys are regular weight. No transform and no tracking.
- Colour: silver `--color-canvas` windows on a teal `--color-page`, white wells, navy title bar (the one gradient, navy to a lighter blue, in `--fill-bar`), navy selection, navy links with purple visited and red active, black text. Block colours are the pale yellow, teal, light blue and peach of the era's pages. Status colours are dark green, amber brown and dark red, all legible on white.
- The era's rule for which text sits on which fill, the same in every layout of the era and to be kept by every token set:
  - The desktop (`--fill-page` over `--color-page`) is always dark enough to carry `--color-bar-text`. Text written straight on the desktop (icon labels, small print) uses `--color-bar-text`.
  - The window face is `--color-canvas`; raised parts inside a window (sheets, notes, dialogs) are `--fill-panel`. Both carry `--color-text`, `--color-text-muted`, `--color-heading`, `--color-heading-alt` and the link colours.
  - A sunken well (list view, tree, document, read-out, empty list) is `--color-surface` and carries the same text and link tokens as the window face, so every palette keeps text and links legible on both. `--color-surface-alt` is the alternate row and the fill of inline code and of a pressed toolbar key; `--color-surface-strong` is a group or total row. They carry `--color-text`.
  - A title bar is `--fill-bar` with `--color-bar-text`; the title bar of a window without the focus is `--fill-inverse` with `--color-inverse-text`.
  - Menu bar, toolbar and column heads are `--fill-bar-alt` with `--color-bar-alt-text`.
  - The selection (current tree row, selected table row, open menu, selected list-box row, progress blocks, the plain badge) is `--fill-accent` with `--color-accent-text`. Links and icons inside a selection inherit that colour.
  - Keys are `--fill-button` with `--color-button-text` (`--color-button-hover-text` under the pointer); the lesser key is `--fill-button-secondary` with `--color-button-secondary-text`; a disabled key is `--color-disabled` with `--color-disabled-text`. The destructive key is `--fill-button` with a bold label in `--color-danger`.
  - Fields and the path field are `--fill-input` with `--color-input-text`; links inside the path field use the link tokens, so a palette keeps `--color-input` close to `--color-surface`. `--color-input-border` outlines the check box.
  - `--color-fill-1` to `--color-fill-4` carry `--color-text` (icons drawn on them) and, in the help layout, `--color-link`.
  - Notices are `--color-notice` with `--color-notice-text` inside a line of `--color-border-strong`; the error notice is `--color-danger` on `--color-danger-surface`.
  - Status labels print `--color-success`, `--color-warning` and `--color-danger` as text and frame on `--color-surface`, so each status colour must read on the well colour; none of them is used as a fill under text.
  - Icons are inline SVG painted with `currentColor`: `--color-heading-alt` by default, the surrounding text colour on keys, title bars, block colours and selections.
- Edges, the same in every layout of the era. Nothing has a border of its own except thin lines; every raised or sunken edge is a `box-shadow` token drawn inside the box:
  - `--shadow-panel`: the raised frame of a window. `--shadow-dialog`: the frame of a dialog.
  - `--shadow-control`: every raised key: buttons, toolbar keys, caption buttons, column heads, tabs and their sheet, pagination keys, the brand tile.
  - `--shadow-control-pressed`: everything sunken: fields, the path field, wells, read-outs, status cells, the check box, the progress bar, a pressed key, the current pagination key. Note that fields take the pressed token, not `--shadow-control`.
  - `--border-width-strong` is the thickness of those edges: padding that must clear an edge is written `calc(var(--border-width-strong) + var(--space-N))`.
  - `--border-width`, `--border-style` and `--color-border` draw the thin lines: group boxes, the etched rule (`--color-border` over `--color-border-muted`), toolbar separators, table row lines (`--color-surface-alt`), the frame of notices, badges and the default button (`--color-border-strong`).
  - `--color-border`, `--color-border-strong` and `--color-border-muted` are the three shades the edge tokens are built from: the mid shadow, the darkest outer line and the light edge. A palette sets them as a bevel set for its window face.
- Not expressible by a token and written literally in `components.css`: the `dotted` guide lines of the tree (and of topic lists), the `solid` strokes of the caption-button glyphs, and `clip-path` on the current tab, which cuts off its lower edge so that it joins the sheet.
- All radii are 0 and are still applied through `--radius-control`, `--radius-panel`, `--radius-pill` and `--radius-page`. `--shadow-text` is applied to title bars and buttons. `--transition` is `none`; `--backdrop-blur` is `0px` and nothing is translucent.

## Components

- `.ds-page`: on `<body>`; the desktop. `.ds-page__window` is the one window at `--size-page`; `.ds-page__body` stacks the blocks of a view; `.ds-page__columns` is the tree pane beside `.ds-page__main`; `.ds-page__inner` is the single column of an inner page.
- `.ds-nav`: the chrome of the window, written once: the title bar (a `.ds-window__title` holding the `.ds-brand`, the caption and the caption buttons), the menu bar `.ds-nav__links` of `.ds-nav__link` entries, each with one underlined `.ds-nav__key` letter, and the `.ds-toolbar`. The entry of the current view (`is-current`, or one selector per view) is drawn in the selection colours, as an open menu is.
- `.ds-brand`: the site's mark and name at the left of the title bar. `.ds-brand__mark` is the window's icon: a small tile dressed like the default button (`--fill-button`, `--color-button-text`, a line of `--color-border-strong`, `--radius-control`, `--shadow-control`) holding the mark as inline SVG with a 3 unit stroke and square caps. `.ds-brand__name` is the name in bold `--font-heading` at `--text-ui`, in the bar's text colour. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo.
- `.ds-hero`: the banner of the open folder, the era's page introduction: a white well with a large icon, `.ds-hero__text` (`.ds-hero__title` at `--text-display` in `--color-heading-alt`, `.ds-hero__lead`), and `.ds-hero__action`, a labelled search field (`.ds-hero__row`, `.ds-hero__field`) with the default button. Never a tall banner.
- `.ds-page-header`: the head of an inner page: a large icon, `.ds-page-header__text` with `.ds-page-header__title` (`--text-h1`) and `.ds-page-header__lead`, and `.ds-page-header__action` with one key at the right. No fill of its own; it stands on the window face.
- `.ds-sidebar`: the tree pane: `.ds-sidebar__title`, a raised strip like a column head, over `.ds-sidebar__body`, a white well holding the `.ds-tree` of folder links. `.ds-pane` with `.ds-pane__title` is the same strip over any other pane.
- `.ds-grid`: the large-icon view of a folder: a white well of four columns. Each `.ds-grid__cell` is a link with a `.ds-grid__icon` tile (a 32px icon on one of the four block colours: `--2`, `--3`, `--4` on the cell), a `.ds-grid__title` and a small `.ds-grid__text` count.
- `.ds-list`: rows on the window face: a small icon, `.ds-list__text` (a link and one line) and a right-aligned `.ds-list__meta` (size and date). `.ds-list__item` is the row.
- `.ds-panel`: a group box, the same thing as `.ds-group`: a `<fieldset>` with its `.ds-panel__title` legend cut into the frame and a `.ds-panel__body`.
- `.ds-wizard`: a wizard page: `.ds-wizard__band`, the sunken picture band in the inverse colours listing each `.ds-wizard__step` (`is-current` bold and underlined), beside `.ds-wizard__main`.
- `.ds-footer`: the status bar of the window, a `.ds-status` row: one wide cell with the site line and quiet links, then short cells.
- `.ds-window`: the raised frame of a window (`--fill-panel`, `--shadow-panel`). `.ds-window__title` is the title bar (`--fill-bar`, bold `--font-ui` in `--color-bar-text`), `.ds-window__title--inactive` the bar of a window without the focus, `.ds-window__caption` the title text, `.ds-window__controls` the group of caption buttons. `.ds-window__ctl` is one caption button, a small raised square with its glyph drawn in CSS: `--min`, `--max`, `--close`. `.ds-window__body` pads plain content.
- `.ds-view`: one screen of the example site; only the one named in the address shows, and `.ds-view--home` when none is named. No box of its own.
- `.ds-prose`: running text: styles `h1`, `h2`, `h3`, `p`, `ul`, `ol`, `li`, `strong` and `code`; `.ds-prose__lead` is the lead line at `--text-large`. Used inside a well or on a sheet.
- `.ds-link`: every link in text, underlined. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--quiet` (`--color-link-quiet`, `--link-decoration-quiet`) for tool and footer links, `.ds-link--strong` for a bold link. `.ds-links` wraps a line of links.
- `.ds-button`: a raised key at least `--size-button` wide, on a `<button>` or on a link that leads to another view. The plain class is the default button of its window: one more line of `--color-border-strong` round the key. `.ds-button--secondary` is every other key, `.ds-button--danger` the destructive one (bold `--color-danger` label, never a red key), `.ds-button--large` a bigger one. States `is-hover`, `is-pressed` (the edge turns inward), `is-focus` (dotted ring inside the key), `is-disabled` or `disabled`. `.ds-buttons` lays out a row; a `.ds-find__fill` spacer pushes the confirming keys to the right.
- `.ds-form`: two columns, `.ds-form__label` at the left (`--size-label`) and `.ds-form__field` at the right. Controls: `.ds-form__input` (`--short`, `--wide`), `.ds-form__select`, `.ds-form__textarea`, `.ds-form__check` inside a `.ds-form__option` label. All are sunken fields. `.ds-form__hint` is a small muted line under a field, `.ds-form__error` a bold line in `--color-danger` under it (the field itself does not change), `.ds-form__actions` a right-aligned row of keys.
- `.ds-table`: a list view: a sunken well of `--color-surface` whose column heads (`.ds-table__head`, `--num` for a right-aligned one) are raised keys. Rows are `.ds-table__row`; `.ds-table__row--alt` tints a row, `is-selected` puts it in the selection colours. `.ds-table__cell`, with `--num` for figures and `--group` for a total or group row. `.ds-table__caption` is the small line under the well.
- `.ds-stat`: the row of summary figures; each `.ds-stat__item` is a sunken read-out with a `.ds-stat__value` (`--text-h1`, `--color-heading-alt`) over a `.ds-stat__label` (`--text-small`, muted). Three or four in a row, never large numerals or cards.
- `.ds-tabs`: a property sheet. `.ds-tabs__item` holds a `.ds-tabs__tab`, a raised key whose lower edge slides under the sheet; the current tab (`is-current`) is taller, bold and joined to the sheet. `.ds-sheet` is the raised sheet under the tabs and always follows them.
- `.ds-badge`: a small bold label in the selection colours with a dark line round it. `.ds-badge--count` is a plain muted count in brackets. Status: `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger`, the status colour as text and frame on the well colour. `.ds-badges` lays out a row of them.
- `.ds-notice`: the pale tip box inside a dark line; `.ds-notice__label` is the bold lead word. `.ds-notice--error` is the error form. `.ds-notices` stacks several.
- `.ds-pagination`: a row of small raised keys: `.ds-pagination__label`, then `.ds-pagination__link` for previous, the page numbers and next; `is-current` is pressed and bold, `is-disabled` greyed.
- `.ds-breadcrumb`: the path, written in a sunken field with a backslash between the `.ds-breadcrumb__item`s; the last is `is-current` and bold.
- `.ds-dialog`: a small window `--size-dialog` wide, shown in the page flow where it would open. `.ds-dialog__title` is its title bar (with a `.ds-window__ctl--close`), `.ds-dialog__content` holds a large icon and the `.ds-dialog__body` message, `.ds-dialog__actions` centres the keys. No backdrop.
- `.ds-empty`: an empty well with `.ds-empty__title`, `.ds-empty__text` and one secondary key.
- `.ds-icon`: a 16px inline SVG pixel icon; `.ds-icon--large` is 32px, `.ds-icon--plain` takes the surrounding text colour.
- `.ds-well`: a sunken field of `--color-surface` for a document or a list. `.ds-rule`: the etched line. `.ds-group` with `.ds-group__title`: a group box, a `<fieldset>` whose `<legend>` is cut into the thin frame. `.ds-props` with `.ds-props__label` and `.ds-props__value`: rows of a property page. `.ds-progress` with `.ds-progress__block` (`--empty` for the rest): a progress bar made of blocks.
- `.ds-tree`: the tree list: `.ds-tree__item` rows, `.ds-tree__link` (icon and name, not underlined, `is-current` in the selection colours), nested `.ds-tree__list` joined by a dotted guide line.
- `.ds-toolbar`: a row of `.ds-toolbar__button` keys (icon and word; `is-pressed`, `is-disabled`) with `.ds-toolbar__sep` etched separators. `.ds-address` puts an `.ds-address__label` before the path field. `.ds-status` is a row of sunken `.ds-status__cell`s (`--wide` takes the free width).
- `.ds-heading`: a section title at `--text-h2`. `.ds-split`: two equal columns. `.ds-stack`: a column of blocks `--space-4` apart. `.ds-find`: a row of labels, fields and a key.
- `.ds-menu`: a pull-down under a menu-bar entry: an `<li class="ds-menu">` holding a `<details>` whose `<summary class="ds-nav__link">` is the entry and whose `.ds-menu__list` of `.ds-menu__item` links opens on a click, in the selection colours while open. `.ds-menu--up` opens it upward, from a task bar.
- `.ds-tooltip`: a hint on hover or focus: the element gets the class, and holds a `.ds-tooltip__text` (`role="tooltip"`) drawn as the pale yellow tool tip; `.ds-tooltip--up` shows it above.

## Never

- `border-radius <= 0px`: the desktop had no rounded corner.
- `box-shadow-blur <= 0px`: every shadow is a hard 1px or 2px bevel line; nothing is blurred.
- `text-shadow = none`: no text shadow anywhere.
- `gradient-fills <= 3%`: the title bars are the only gradient.
- `border-width <= 2px`: thin lines are 1px; the widest stroke is the 2px top of a caption glyph.
- `font-size >= 10px`: captions are 10px, nothing smaller.
- `font-size <= 22px`: the hero title is the largest text; page titles stop at 16px.
- `font-weight >= 400`: only regular and bold.
- `font-families <= 2`: the system sans and a monospace face.
- `letter-spacing = 0`: no tracking.
- `uppercase-text <= 0%`: no text-transform.
- `row-gap <= 8px`: rows of lists, trees and list views touch or sit 1px to 4px apart.
- `block-gap <= 16px`: blocks inside the window are 8px apart, never more than 16px.
- `content-width <= 760px`: everything is inside the one 760px window.
- `transition = none`: nothing eases on hover.
- `animation = none`: no CSS animation.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. Anything that holds content is a `.ds-window` with a title bar or a `.ds-group` with its caption; anything that lists is a `.ds-table`, a `.ds-tree` or a list in a `.ds-well`; anything pressed or typed into takes `--shadow-control-pressed`, anything pushed takes `--shadow-control`. Clear every edge with `calc(var(--border-width-strong) + ...)` padding, keep text at `--text-base` or `--text-ui`, and never add a radius, a blurred shadow, a gradient outside `--fill-*` tokens, a hover animation or a colour of your own.
