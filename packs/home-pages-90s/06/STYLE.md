# Personal home page: the game shrine, 2000

## Summary

A shrine to one game, as fans kept them between 1996 and 2001: a narrow top frame across the whole window with the name and the links, and under it a dark page of two unequal columns, an "altar" with a framed picture and fact boxes at the left and the title, the updates and the newest screenshots at the right. Text is small pale Trebuchet on black-red with gold and rose lowercase headings in a condensed poster face, red underlined links, thin 1px boxes with a second line drawn round them and bars that darken toward the foot. A walkthrough in chapters, an item table, a gallery, a web ring and a guestbook are its content; the kind faded when wikis and video took over after 2001.

## Layout

- Fixed, not fluid: `--size-page` is 740px, centred in an 800px window. `.ds-page__frame` has no fill and no border: everything below the top frame stands on `--fill-page`, the tile.
- `.ds-nav` is written once, outside the frame: a strip on `--fill-bar` as wide as the window, closed by a `--border-width-strong` line in `--color-border-strong`, that stays at the top of the window as a frame did. `.ds-nav__frame` inside it is 740px: the `.ds-brand` at `--text-h1` at the left, `.ds-nav__links` at the right.
- The specimen is one example site, a shrine to a made-up role-playing game, of five views.
- Home view, `.ds-page__main`, is two columns `--space-5` apart. Left, `.ds-altar` (`--size-altar` 236px): the `.ds-altar__frame` with its square picture and caption, then three `.ds-panel` boxes (the fact file, the web ring, a short list). Right, `.ds-page__column`: the `.ds-hero` with no box (the title at `--text-display`, centred on the tile), the `.ds-stat` line of three plaques, a `.ds-heading` over the `.ds-list` of dated updates, a heading over the `.ds-grid` of four framed thumbnails, and the `.ds-construction` strip. The two columns end near each other.
- Inner pages keep the top frame and the footer and drop the altar and the hero. `.ds-page__inner` is one column 740px wide and opens with `.ds-page__head`: the `.ds-breadcrumb`, the `.ds-page-header` (title and one italic line at the left, a button at the right, a bar under that fades out to the right) and, in the gallery, the centred `.ds-tabs`. Below it: `.ds-page__columns`, a `--size-side` (180px) `.ds-sidebar` of chapters at the left and `.ds-prose` at the right, then a fading rule and chapter numbers; or the grid of eight thumbnails, page numbers and the empty state; or the table; or entries, notices, form and message box.
- Views: `home`; `walkthrough` (one chapter: the sidebar of chapters, running text, link states, chapter numbers); `gallery` (tabs, eight thumbnails, page numbers, the empty state of a part with no pictures); `items` (the table of weapons with its badge key); `guestbook` (entries, notices, the form to sign, the delete message box).
- Spacing is tight: `--space-1` 2px strip padding and the height of a fading bar; `--space-2` 4px caption and title-strip padding; `--space-3` 6px cell padding, between a title and its block, between tabs and form rows; `--space-4` 10px box padding, paragraph spacing, between buttons; `--space-5` 14px between blocks, columns and thumbnails; `--space-6` 24px list indent and the space above the footer.
- Fixed sizes: `--size-label` 150px form labels, `--size-field` 240px inputs, `--size-area` 340px textarea, `--size-check` 14px, `--size-dialog` 360px, `--size-icon` 36px, `--size-thumb` 76px height of a thumbnail, `--size-date` 78px date column of the updates.

## Typography and colour roles

The conventions for which text sits on which fill are those of the era's first layout: text, heading, link and status colours are legible on page, canvas, surfaces and the four fills; bars, the inverse, the accent and notices bring their own text colour; status colours as fills are lettered in `--color-canvas`. This layout adds one: the name and the links stand on `--fill-bar`, so the brand and the top-frame links use `--color-bar-text`.

- `--font-body` and `--font-ui` are Trebuchet MS: body 12px at line height 1.45, small print 10px, controls 11px bold. `--font-heading` is Impact (then Haettenschweiler) at weight 400 with 1px tracking, set lowercase by `--heading-transform`: 15, 22 and 34px, the front-page title 52px with 3px tracking. Buttons, links of the top frame, table heads, tabs and badges are lowercase too (`--ui-transform`). `--font-mono` (Fixedsys, Lucida Console) sets plaques, dates and numeric cells. Links in text are underlined and gain an overline under the pointer; quiet links have no line.
- `--color-page` `#140006` under a diamond lattice, `--color-canvas` `#260010`, `--color-surface` `#33000f` (boxes, rows), `--color-surface-alt` `#4d0a1c` (alternate rows, code, the empty state), `--color-surface-strong` `#661428` (the plaque line, the closing row of a table).
- `--color-text` pale `#f0e0d0`; `--color-heading` gold `#ffd24d` (titles, the picture, names in the guestbook, numeric cells, h1 and h3); `--color-heading-alt` rose `#ff6680` (section titles, lead lines, dates, fact names, form labels). Links red `#ff4d4d`, visited copper, hover white. `--color-accent-alt` pale steel is the "Updated!" word and list marks.
- `--fill-bar` crimson with white text is the top frame, table heads and the dialog title. `--fill-bar-alt` black with gold text is every title strip. `--fill-inverse` silver with dark text is the plaques and the current tab. `--fill-accent` gold with dark text is the current link of the top frame, the "NEW!" badge and the current page number. `--color-fill-1` to `-4` (wine, umber, violet, midnight) are the grounds of the pictures and carry gold drawings.
- Buttons are wine `--color-button` with white text; the secondary is charcoal with silver text. Form fields are black with pale text.
- `--color-success`, `--color-warning` and `--color-danger` double as fills: a status badge and the destructive button are filled with one of them and lettered in `--color-canvas`. `--color-notice` with `--color-notice-text` is notices and the two stripes of the construction strip; `--color-danger-surface` with `--color-danger` is the error notice.
- Written literally in `components.css`: `outset` and `inset` bevels on buttons, the brand tile and form controls (drawn in the control's own colour), `dotted` rules inside lists and boxes, italic for taglines and lead lines, and the slow two-step `.ds-blink`, which only runs inside `prefers-reduced-motion: no-preference`.
- Surface: `--border-style` `solid`, 1px for boxes and 3px for the frame line, the picture frame, heading marks and the dialog. `--shadow-panel` is not a shadow but a second frame: a 2px ring of `--color-canvas` and a 1px line of `--color-border` round every box; `--shadow-dialog` is the same in `--color-border-strong` round the picture and the dialog; a hovered button gains a 1px ring. Nothing is blurred. `--fill-bar` darkens from top to foot, `--fill-bar-alt` lightens, buttons have a 1px light line at the top. `--fill-page` is a 14px diamond lattice of `--color-border-muted` over `--color-page`. Radius 0, no text shadow, no transitions.
- Also written literally: the `solid` left mark of a heading, and the `linear-gradient` of the fading bars, which runs from `--color-border-strong` to `transparent`.

## Components

- `.ds-page`: on `<body>`; the tiled page and body type. `.ds-page__frame` the centred 740px column; `.ds-page__main` the two columns of the front page and `.ds-page__column` the right one; `.ds-page__inner` the single column of an inner page; `.ds-page__head` stacks path, page header and tabs; `.ds-page__columns` chapter list plus text; `.ds-section` a heading with its block.
- `.ds-nav`: the top frame. `.ds-nav__frame` centres its content; `.ds-nav__links` is the line of `.ds-nav__item` / `.ds-nav__link` links with a slash between; the link of the view that shows takes `--fill-accent` (one selector per view, or `is-current`).
- `.ds-brand`: mark and name at the left of the top frame. `.ds-brand__mark` is a tile dressed like the button (`--fill-button`, `--color-button-text`, outset bevel, `--radius-control`, `--shadow-control`) holding the mark as inline SVG with square caps; `.ds-brand__name` is the name at `--text-h1` in `--color-bar-text`. Name and mark are placeholders for the installing project's own.
- `.ds-altar`: the left column of the front page. `.ds-altar__frame` is the framed picture: `.ds-altar__pic` (a square on `--color-fill-1` with a line drawing) over `.ds-altar__caption`. `.ds-facts` with `.ds-facts__item` and `.ds-facts__name` is a fact file inside a panel.
- `.ds-hero`: the introduction, with no box: `.ds-hero__title` at `--text-display`, `.ds-hero__lead`, `.ds-hero__action` (a `.ds-button--large` and a secondary button), `.ds-hero__hint`.
- `.ds-page-header`: `.ds-page-header__text` with `.ds-page-header__title` and the italic `.ds-page-header__lead`, `.ds-page-header__action` at the right, and a fading bar under.
- `.ds-heading`: a section title in `--color-heading-alt` with a thick mark at the left and a thin line under.
- `.ds-rule`: a plain rule; `.ds-rule--fade` is a bar that fades out at both ends, the usual divider.
- `.ds-table`: striped rows with a line under each: `.ds-table__head` (`--num`), `.ds-table__cell` (`--num` in the monospace and `--color-heading`, `--group` a closing row on `--color-surface-strong`), `.ds-table__row--alt`, `.ds-table__desc`, `.ds-table__caption`.
- `.ds-stat`: one line of plaques on `--color-surface-strong`; each `.ds-stat__item` is a `.ds-stat__value` on `--fill-inverse` beside a `.ds-stat__label`.
- `.ds-list`: dated updates in a box: `.ds-list__item` with a `.ds-list__meta` date and a `.ds-list__body`; `.ds-list__more` under it. `.ds-list--entries` is the guestbook: `.ds-list__title`, meta, `.ds-list__text`.
- `.ds-panel`: a box with a `.ds-panel__title` strip on `--fill-bar-alt`, `.ds-panel__body` and `.ds-panel__text`. `.ds-ring` is the previous, random, next line of a web ring.
- `.ds-grid`: four framed thumbnails to a row: `.ds-grid__cell` (`--2`, `--3`, `--4` pick the fill) holds a `.ds-grid__pic` block with a line drawing over a `.ds-grid__title` link and a `.ds-grid__text` file size.
- `.ds-tabs`: a centred row of boxes, `.ds-tabs__item` / `.ds-tabs__tab`, the current one (`is-current`) on `--fill-inverse`.
- `.ds-sidebar`: a column of `.ds-sidebar__block` boxes, each a `.ds-sidebar__title` strip over a `.ds-sidebar__list` of `.ds-sidebar__item` lines, the chapter that shows `is-current`.
- `.ds-footer`: a fading rule, `.ds-footer__links` with `.ds-footer__item`, the italic `.ds-footer__note` and `.ds-footer__legal`, centred on the tile.
- `.ds-prose`: running text; styles `h1`, `h2`, `h3`, `p`, `ul`, `ol`, `strong`, `code`; `.ds-prose__lead` the italic opening line, `.ds-prose__code` a block to copy.
- `.ds-link`: every link; states `is-visited`, `is-hover`, `is-active`, `is-focus`; `.ds-link--quiet` small, `.ds-link--strong` bold, `.ds-link--plain` on a coloured fill. `.ds-links` wraps a line of them. `.ds-icon` holds a small drawing before a link.
- `.ds-bullets`: a link list with small square marks, `.ds-bullets__item` holding a `.ds-bullets__text`.
- `.ds-button`: a bevelled key; `.ds-button--secondary`, `.ds-button--danger`, `.ds-button--large`; states `is-hover`, `is-pressed`, `is-focus`, `is-disabled`. `.ds-buttons` lays out a row.
- `.ds-form`: two columns, right-aligned `.ds-form__label` and `.ds-form__field`; `.ds-form__input` (`--short`, `is-error`), `.ds-form__select`, `.ds-form__textarea`, `.ds-form__check`, `.ds-form__hint`, `.ds-form__error`, `.ds-form__actions`.
- `.ds-badge`: a solid "NEW!" label on `--fill-accent`; `.ds-badge--text` an italic coloured word, `.ds-badge--count` a muted count, `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger` status labels; `.ds-badges` the key line under a table.
- `.ds-notice`: a line on `--color-notice`; `.ds-notice__label`; `.ds-notice--error`. `.ds-notices` stacks them.
- `.ds-pagination`: a centred line: `.ds-pagination__label`, `.ds-pagination__link` with `is-current` (on `--fill-accent`) and `is-disabled`.
- `.ds-breadcrumb`: a small path, `.ds-breadcrumb__item` separated by a coloured mark, the last `is-current`.
- `.ds-dialog`: a message box in the page flow: `.ds-dialog__title` bar on `--fill-bar`, `.ds-dialog__body`, `.ds-dialog__actions`. No backdrop.
- `.ds-empty`: `.ds-empty__title`, `.ds-empty__text` and a button on `--color-surface-alt`.
- `.ds-construction`: the diagonal-striped strip with `.ds-construction__text` on it. `.ds-blink`: add to one small mark for the slow blink.
- `.ds-view`: one screen of the example site; `.ds-view--home` shows when the address names none.

## Never

- `border-radius <= 0px`: every box and button is square.
- `box-shadow-blur <= 0px`: the second frames are hard rings, never soft.
- `text-shadow = none`: text is flat.
- `border-width <= 3px`: the frame line is the thickest.
- `font-size <= 52px`: the front-page title is the largest text.
- `font-size >= 10px`: small print is 10px.
- `font-weight >= 400`: regular and bold only.
- `font-families <= 3`: Trebuchet, the poster face of the headings and one monospace.
- `uppercase-text <= 0%`: nothing is set in capitals.
- `content-width <= 740px`: both columns fit an 800px window.
- `transition = none`: nothing eases on hover.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new page gets a link in the top frame; a new box of facts is a `.ds-panel` in the altar; a new block at the right is a `.ds-heading` over a thin-framed box; a new picture is a `.ds-grid__cell`; a new figure joins `.ds-stat`. Keep the top frame, keep the picture at the left of the front page, keep text small and pale on the dark tile and links red and underlined, and never add radius, blur, transitions, a sheet behind the columns or a menu down the side.
