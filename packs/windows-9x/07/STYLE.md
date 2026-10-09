# Desktop grey chat and message boards, canary, 1998

## Summary

A community site laid out as the chat program of a late 1990s desktop: one window with a switch bar of window keys under the title bar, the open board or conversation at the left and the list of who is here in a rail at the right, closed by a status bar. This pack shows it in a canary scheme taken from the download sites of the time: pale yellow windows on a blue striped desktop, black title bars with yellow lettering, red selection and headings, every edge a flat black line with a second line inside it, and rounded hand-drawn headings over a 12px system sans. Support boards, hobby communities and web chat rooms took this form between 1996 and 2001, when visitors already knew a chat program and a newsgroup reader and a site could borrow both.

## Layout

- Fixed, not fluid: `--size-page` is 768px, one window (`.ds-page__window`, also a `.ds-window`) centred on the desktop of an 800px screen with a `--space-6` margin above and below.
- The specimen is one example site, the boards, of five views. `.ds-nav` (title bar and switch bar), the rail and `.ds-footer` (status bar) are written once, inside the window and outside the views; each view is a `.ds-view` and one shows at a time.
- Under the switch bar comes `.ds-page__columns`: `.ds-page__main` at the left, which holds the views, and `.ds-page__rail` at the right (`--size-rail`, 168px), `--space-2` apart and stretched to the same height. The rail is the only side column and it is at the right: a caption strip, the `.ds-nicks` list of names, which takes the spare height, and the `.ds-sidebar` block of links at its foot.
- Every view is a `.ds-stack`, a column of blocks `--space-4` apart.
- Home view (the lobby): the `.ds-hero` (a well with a band in the selection colours at its left, the title, a lead, one line and two keys), the `.ds-stat` line of four read-outs, the `.ds-grid` of six boards, two to a row, and a `.ds-split` of two `.ds-panel`s: the latest messages as a `.ds-list` and the house rules.
- Inner pages keep the title bar, the switch bar, the rail and the status bar and drop the hero, the read-outs and the boards. They open with the `.ds-breadcrumb` (small links), then the `.ds-page-header`: title and one line at the left, one key at the right, closed by a thick rule in the accent colour. Tabs come straight after the page header, with their `.ds-sheet`.
- Views: `home` is the lobby. `board` is one board: page header, four tabs over a sheet with the list view of topics (status labels, a selected row, a total row), then the pagination keys beside the key of status labels. `topic` is a conversation: `.ds-posts` with the name of the writer in a column at the left of each message, a line the program says itself, and the Say line under it. `write` is the form for a new message under two notices, with the confirmation dialog beside a panel of advice. `members` is a member's card beside the empty state of the private messages, over the list view of members. The key of the view that is showing stays pressed and bold; the Lobby key is marked when the address names no view.
- Fixed sizes: `--size-nick` 96px name column of a message; `--size-ctl` 16px caption button and `--size-glyph` 8px glyph and lamp; `--size-icon` 32px and `--size-icon-small` 16px icons; `--size-check` 13px check box and radio button; `--size-label` 96px form label column; `--size-field` 200px text field; `--size-dialog` 316px; `--size-button` 75px least key width; `--size-block` 8px progress block.
- Spacing scale: `--space-0` 1px is the hairline step (row padding in the list of names); `--space-1` 2px the gap between keys and the padding of bars; `--space-2` 4px the padding of keys, the gap between an icon and its word, between boards and between the working pane and the rail; `--space-3` 6px cell and field padding and the gap between keys in a row; `--space-4` 8px the gap between blocks and the padding of panels; `--space-5` 12px the padding of the hero, sheets and dialogs; `--space-6` 16px the margin of the desktop round the window.

## Typography and colour roles

- `--font-body` and `--font-ui` are the system sans at `--text-base` and `--text-ui` (12px); `--font-heading` is the rounded hand-drawn face, used for the hero title (`--text-display`, 30px), page titles (`--text-h1`, 20px), section titles (`--text-h2`, 15px), read-out figures and the brand name. `--font-mono` sets inline code and the times in the list of latest messages. `--text-small` (10px) is for dates, counts, roles and the breadcrumb. Bold is `--weight-bold`; nothing is lighter than `--weight-body`.
- Links in text are `--color-link`, underlined by `--link-decoration`, and lose the line under the pointer (`--link-decoration-hover`); quiet links (breadcrumb, status bar) are `--color-link-quiet` with `--link-decoration-quiet`.
- Which text sits on which fill, the same in every layout of the era:
  - The window face is `--color-canvas` (`--fill-panel` for windows, sheets, tabs and dialogs) and carries `--color-text`, `--color-text-muted`, `--color-heading`, `--color-heading-alt` and the link colours.
  - A sunken well (hero, read-outs, boards, panel bodies, the list of names, the sidebar, posts, list views, the empty state) is `--color-surface` and carries the same text and link tokens. `--color-surface-alt` is the alternate row, the name column of a message, inline code and the pressed switch bar key; `--color-surface-strong` is a total row. They carry `--color-text`.
  - The title bar is `--fill-bar` with `--color-bar-text`. The switch bar, caption strips and column heads are `--fill-bar-alt` with `--color-bar-alt-text`.
  - The selection (the band of the hero, the selected name and table row, progress blocks, the plain badge) is `--fill-accent` with `--color-accent-text`; links and icons inside a selection inherit that colour. `--color-accent` also draws the rule under a page header. `--color-accent-alt` is the line the program says itself.
  - Keys are `--fill-button` with `--color-button-text` (`--fill-button-hover` and `--color-button-hover-text` under the pointer); the lesser key is `--fill-button-secondary` with `--color-button-secondary-text`; a disabled key is `--color-disabled` with `--color-disabled-text`; the destructive key is `--fill-button` with a bold label in `--color-danger`.
  - Fields are `--fill-input` with `--color-input-text`; `--color-input-border` outlines the check box.
  - `--color-fill-1` to `--color-fill-4` are the icon blocks of the boards and carry `--color-text`.
  - Notices are `--color-notice` with `--color-notice-text` inside a line of `--color-border-strong`; the error notice is `--color-danger` on `--color-danger-surface`.
  - Status labels print `--color-success`, `--color-warning` and `--color-danger` as text and frame on `--color-surface`; the lamp of the switch bar is a block of `--color-success`.
  - Icons are inline SVG painted with `currentColor`: `--color-heading-alt` by default, the surrounding text colour on keys, bands, block colours and selections.
- Edges. Nothing has a border of its own except thin lines; every raised or sunken edge is a `box-shadow` token drawn inside the box:
  - `--shadow-panel` frames the window, `--shadow-dialog` a dialog. Both also throw a hard shadow.
  - `--shadow-control` is every raised key: buttons, switch bar keys, caption buttons, caption strips, column heads, tabs and their sheet, pagination keys, the brand tile. `--shadow-control-hover` replaces it under the pointer.
  - `--shadow-control-pressed` is everything sunken: fields, wells, read-outs, status cells, the check box, the progress bar, a pressed key, the current pagination key.
  - `--border-width-strong` is the thickness of those edges: padding that must clear an edge is written `calc(var(--border-width-strong) + var(--space-N))`.
  - `--border-width`, `--border-style` and `--color-border` draw the thin lines: the frame and icon block of a board, group boxes, the etched rule (over `--color-border-muted`), the line between messages (`--color-border-muted`), table row lines (`--color-surface-alt`).

## Components

The name `chronoskin` and its mark are placeholders for the installing project's own name and mark.

- `.ds-page`: the desktop, on `<body>`; `.ds-page__window` is the one window, `.ds-page__columns` the working pane (`.ds-page__main`) and the rail (`.ds-page__rail`).
- `.ds-window`: a raised frame; `.ds-window__title` is its title bar (`--inactive` for a window without the focus) with `.ds-window__caption` and the `.ds-window__ctl` caption buttons (`--min`, `--max`, `--close`).
- `.ds-nav`: the title bar and the switch bar. `.ds-nav__links` is the bar, each `.ds-nav__item` an equal share of it, `.ds-nav__link` a raised window key with an icon; the current one is pressed and bold. `.ds-nav__state` is the sunken connection read-out at the right, with its `.ds-nav__lamp`.
- `.ds-brand`: the mark and the name at the left of the title bar. `.ds-brand__mark` is a small key dressed like the default button, `.ds-brand__name` the name in the heading face.
- `.ds-hero`: the welcome of the lobby, a well. `.ds-hero__band` is the block in the selection colours with a large icon, `.ds-hero__body` holds `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__text` and `.ds-hero__action`. Home view only.
- `.ds-page-header`: the head of an inner page: `.ds-page-header__text` with `.ds-page-header__title` and `.ds-page-header__lead` at the left, `.ds-page-header__action` at the right, a thick accent rule under it.
- `.ds-view`: one screen of the example site; `.ds-view--home` shows when the address names none.
- `.ds-prose`: running text inside a message: `h1`, `h2`, `h3`, `p`, `ul`, `ol`, `li`, `strong`, `code`; `.ds-prose__lead` is the lead line.
- `.ds-link`: every link in text. States `is-visited`, `is-hover`, `is-active`, `is-focus`; `.ds-link--quiet` for tool and footer links, `.ds-link--strong` for a bold one. `.ds-links` wraps a line of links.
- `.ds-button`: a raised key at least `--size-button` wide; the plain class is the default key, inside one more line. `.ds-button--secondary` is every other key, `.ds-button--danger` the destructive one, `.ds-button--large` a bigger one. States `is-hover`, `is-pressed`, `is-focus`, `is-disabled` or `disabled`. `.ds-buttons` lays out a row; a `.ds-find__fill` spacer pushes keys apart.
- `.ds-form`: two columns, `.ds-form__label` and `.ds-form__field`. Controls: `.ds-form__input` (`--short`, `--wide`), `.ds-form__select`, `.ds-form__textarea`, `.ds-form__check` and `.ds-form__radio` inside a `.ds-form__option`. `.ds-form__hint` is a small line under a field, `.ds-form__error` a bold line in the danger colour, `.ds-form__actions` the row of keys.
- `.ds-table`: a list view: a well whose `.ds-table__head` cells are raised (`--num` right-aligned). `.ds-table__row` (`--alt` tinted, `is-selected` in the selection colours), `.ds-table__cell` (`--num`, `--group` for a total row), `.ds-table__caption`.
- `.ds-list`: lines of a log. Each `.ds-list__item` is a `.ds-list__meta` (the time, monospace), a `.ds-list__who` (the name, bold) and a `.ds-list__text`; an icon may stand in place of the first two.
- `.ds-panel`: a raised caption strip (`.ds-panel__title`) over a well (`.ds-panel__body`). `.ds-panel--fill` takes spare height.
- `.ds-stat`: the line of read-outs; each `.ds-stat__item` is a sunken cell with the `.ds-stat__value` (heading face, `--text-h1`) before its `.ds-stat__label` on one line.
- `.ds-grid`: the boards, two to a row. A `.ds-grid__cell` is a framed row: `.ds-grid__icon` is the block of colour (`.ds-grid__cell--2` to `--4` take the other block colours), `.ds-grid__body` holds `.ds-grid__title`, `.ds-grid__text` and `.ds-grid__meta`.
- `.ds-tabs`: a property sheet: `.ds-tabs__item` holds a `.ds-tabs__tab`; the current one (`is-current`) is taller, bold and joined to the `.ds-sheet` that always follows.
- `.ds-badge`: a small bold label in the selection colours. `.ds-badge--count` is a plain muted count. Status: `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger`. `.ds-badges` lays out a row.
- `.ds-sidebar`: a well at the foot of the rail with a `.ds-sidebar__title` and a `.ds-sidebar__list` of links.
- `.ds-notice`: the tip box inside a dark line; `.ds-notice__label` is the bold lead word, `.ds-notice--error` the error form. `.ds-notices` stacks several.
- `.ds-pagination`: a row of small keys: `.ds-pagination__label`, then `.ds-pagination__link`s; `is-current` is pressed and bold, `is-disabled` greyed.
- `.ds-breadcrumb`: the path in small type, the `.ds-breadcrumb__item`s parted by a double angle; the last is `is-current`.
- `.ds-dialog`: a small window `--size-dialog` wide, shown where it would open: `.ds-dialog__title`, `.ds-dialog__content` with a large icon and the `.ds-dialog__body`, `.ds-dialog__actions`.
- `.ds-empty`: an empty well with a large icon, `.ds-empty__title`, `.ds-empty__text` and one secondary key.
- `.ds-footer`: the status bar, a `.ds-status` row of sunken `.ds-status__cell`s (`--wide` takes the free width).
- `.ds-nicks`: the list of who is here: `.ds-nicks__group` is a small caption, `.ds-nicks__item` a name with its icon (`--host` bold, `--away` muted, `is-selected` in the selection colours).
- `.ds-posts`: a conversation in a well. Each `.ds-post` has `.ds-post__who` (`.ds-post__name`, `.ds-post__role`) at the left and `.ds-post__body` (`.ds-post__head`, then `.ds-post__text` or a `.ds-prose`) at the right; `.ds-post__event` is a line the program says itself.
- `.ds-icon` (`--large`, `--plain`), `.ds-well`, `.ds-rule`, `.ds-group` with `.ds-group__title`, `.ds-props` with `.ds-props__label` and `.ds-props__value`, `.ds-progress` with `.ds-progress__block`: the small parts every layout of the era shares.
- `.ds-heading`: a section title. `.ds-split`: two equal columns of the same height. `.ds-stack`: a column of blocks (`--tight` closer). `.ds-find`: a row of a label, a field and keys.

## Never

- `border-radius <= 0px`: the desktop had no rounded corner; the radio button is a circle.
- `box-shadow-blur <= 0px`: edges are hard lines and the shadow of a window is hard; nothing is blurred.
- `text-shadow = none`: no text shadow anywhere.
- `gradient-fills <= 3%`: the title bars and the stripes of the desktop are the only gradients.
- `border-width <= 2px`: thin lines are 1px; the thick stroke is the 2px rule under a page header.
- `font-size >= 10px`: dates and counts are 10px, nothing smaller.
- `font-size <= 30px`: the welcome of the lobby is the largest text; page titles stop at 20px.
- `font-weight >= 400`: only regular and bold.
- `font-families <= 3`: the system sans, the hand-drawn heading face and a monospace.
- `letter-spacing = 0`: no tracking.
- `uppercase-text <= 0%`: no text-transform.
- `row-gap <= 8px`: rows of lists and list views touch or sit 1px to 4px apart.
- `block-gap <= 16px`: blocks inside the window are 8px apart, never more than 16px.
- `content-width <= 768px`: everything is inside the one 768px window.
- `transition = none`: nothing eases on hover.
- `animation = none`: no CSS animation.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. Anything that holds content is a `.ds-panel` (caption strip over a well) or a `.ds-well`; anything that lists is a `.ds-table`, a `.ds-list` or the `.ds-nicks` list; a new kind of conversation is a `.ds-posts`. Anything pressed or typed into takes `--shadow-control-pressed`, anything pushed takes `--shadow-control`. Clear every edge with `calc(var(--border-width-strong) + ...)` padding, keep the rail at the right and the same on every view, keep text at `--text-base` or `--text-ui`, and never add a radius, a blurred shadow, a gradient outside the `--fill-*` tokens, a hover animation or a colour of your own.
