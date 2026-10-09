# Personal home page: the free graphics page, 2000

## Summary

A "free stuff for your home page" site as hobbyists ran them between 1996 and 2001, built to look like two frames: a narrow menu frame of bevelled bars at the left on its own fill, a frame border, and the main frame on a tiled brick wall. Text is small white Verdana on deep navy with heavy uppercase orange and gold headings, sky-blue underlined links, raised 2px boxes with hard black shadows and glossy gradient bars and buttons. Shelves of downloads, numbered lessons, awards, a web ring and a submit-your-site form are its content; the kind faded when hosts and search engines took over after 2001.

## Layout

- Fixed, not fluid: `--size-page` is 784px, centred in an 800px window. `.ds-page__frame` is a two-column grid: `.ds-nav`, the `--size-menu` (150px) menu frame, and `.ds-page__body`, the main frame. Nothing is wider than its frame.
- The menu frame is filled with `--fill-bar-alt`, lettered in `--color-bar-alt-text`, runs the full height of the page and ends in a `--border-width-strong` frame border in `--color-border-strong`. `.ds-nav__frame` inside it is sticky, so the menu stays in the window: the `.ds-brand` (mark over name), a tagline, `.ds-nav__links` (a column of bevelled bars), a search field, the visitor counter and the "best viewed" note.
- The main frame has no fill: its content stands on `--fill-page`. It holds the views and, once, the `.ds-footer`.
- The specimen is one example site, a free graphics collection, of four views.
- Home view, `.ds-page__main`: the `.ds-hero` title box (the title at `--text-display`), the `.ds-stat` row of four figures, a `.ds-heading` bar over the `.ds-grid` of eight shelves (four to a row), `.ds-page__pair` (the `.ds-list` of updates at 3 parts, a `.ds-panel` at 2), the `.ds-construction` strip and a heading over the `.ds-awards`.
- Inner pages drop the title box and figures. `.ds-page__inner` opens with `.ds-page__head`: the `.ds-breadcrumb`, the `.ds-page-header` (title and one line at the left, a button at the right, a rule under) and, for a shelf, the `.ds-tabs` folder tabs. Below it one column: table, page numbers, a rainbow rule, a bulleted list; or notices, form and message box. A lesson uses `.ds-page__columns`: `.ds-prose` and a `--size-side` (160px) `.ds-sidebar` at the right.
- Views: `home`; `graphics` (a shelf: tabs, the table of sets with its badge key, page numbers, the rules); `tutorial` (a lesson with code, the sidebar, link states); `submit` (notices, the form, the confirm message box, the empty list of sites sent).
- Spacing is tight: `--space-1` 2px strip padding and the gap between figures; `--space-2` 4px cell padding and the gap between menu bars; `--space-3` 6px box padding, the gap between shelves and under a heading bar; `--space-4` 10px frame padding above, gaps between buttons; `--space-5` 14px between blocks and columns and at the sides of the main frame; `--space-6` 22px list indent.
- Fixed sizes: `--size-label` 130px, `--size-field` 220px, `--size-area` 320px, `--size-check` 14px, `--size-dialog` 340px, `--size-icon` 30px, `--size-award` 112px, `--size-date` 52px.

## Typography and colour roles

The conventions for which text sits on which fill are those of the era's first layout: text, heading, link and status colours are legible on page, canvas, surfaces and the four fills; bars, the inverse, the accent and notices bring their own text colour; status colours as fills are lettered in `--color-canvas`. This layout adds one: the menu frame is `--fill-bar-alt`, so everything on it that is not a button uses `--color-bar-alt-text`.

- `--font-body` Verdana 12px at line height 1.4, small print 10px. `--font-heading` Arial Black (weight 900), uppercase by `--heading-transform`: 14, 20 and 30px, the front-page title 44px with -1px tracking. `--font-ui` Arial 11px bold uppercase sets buttons, menu bars, table heads, tabs, the path and figure labels. `--font-mono` (Andale Mono, Courier New) sets code and the counter. Links in text are underlined; quiet links are not (`--link-decoration-quiet: none`).
- `--color-page` navy `#00004d` under the bricks, `--color-surface` `#000066` for boxes and rows, `--color-surface-alt` `#1a1a80` alternate rows, `--color-surface-strong` `#33338f` inactive tabs, the closing row of a table and the ground between figures.
- `--color-text` white; `--color-heading` orange `#ff9933` (titles, h1, h3); `--color-heading-alt` gold `#ffcc00` (figures, h2, dates, icons, form labels). Links sky `#66ccff`, visited lilac, hover yellow. `--color-accent-alt` aqua is the "Updated!" word and bullets.
- `--fill-bar` orange with navy text is section heading bars, table heads, the current tab and the dialog title. `--fill-bar-alt` brick red with pale gold text is the menu frame and title strips. `--fill-inverse` white with navy text is awards, the counter and the current page number. `--fill-accent` gold with navy text is the pressed menu bar and the "New!" badge. `--color-fill-1` to `-4` (maroon, brown, teal, plum) are the shelves and the table swatches and carry white text and sky links.
- Buttons are red `--color-button` with white text; the secondary is blue.
- Surface: `--border-style` `outset`, 2px for boxes, 4px for the frame border, the title box, awards and the dialog. Boxes cast a hard 3px black shadow (`--shadow-panel`), the title box and dialog 5px, buttons 1px; nothing is blurred. `--shadow-text` is a 1px hard shadow on bars, titles and the name. Bars and buttons are top-lit gradients; `--fill-input` has a 2px inner shade; `--fill-page` is a 48 by 32px brick pattern over `--color-page`. Radius 0, no transitions.
- Written literally: `outset` and `inset` bevels on buttons, menu bars, swatches and controls, `solid` row rules and the tab base line, italic for the tagline and lead lines, the `letter-spacing` of the counter (`--space-1`), and `.ds-blink`.

## Components

- `.ds-page`: on `<body>`. `.ds-page__frame` the two-frame grid; `.ds-page__body` the main frame; `.ds-page__main`, `.ds-page__inner` stack blocks; `.ds-page__head`; `.ds-page__columns` lesson plus sidebar; `.ds-page__pair` a 3 to 2 pair; `.ds-section` a heading bar with its block.
- `.ds-nav`: the menu frame. `.ds-nav__frame` the sticky inner column; `.ds-nav__tagline`; `.ds-nav__links` with `.ds-nav__item` / `.ds-nav__link`, bars in `--fill-button` with an outset bevel; the bar of the view that shows is pressed in on `--fill-accent`; `.ds-nav__note` small print; `.ds-nav__count` the boxed counter.
- `.ds-brand`: `.ds-brand__mark`, a tile dressed like the primary button, over `.ds-brand__name` at `--text-h3` (the frame is too narrow for the display size, which the front-page title uses). Name and mark are placeholders for the installing project's own.
- `.ds-hero`: the title box with `--shadow-dialog`: `.ds-hero__title` at `--text-display`, `.ds-hero__lead`, `.ds-hero__action`, `.ds-hero__hint`.
- `.ds-page-header`: `.ds-page-header__text` (`.ds-page-header__title`, `.ds-page-header__lead`) and `.ds-page-header__action`.
- `.ds-heading`: a section title on a full-width `--fill-bar` strip.
- `.ds-rule`, `.ds-rule--short`, `.ds-rule--rainbow`: the rule and the six-step colour bar.
- `.ds-construction` with `.ds-construction__text`; `.ds-blink`.
- `.ds-prose`: running text (`h1` to `h3`, `p`, lists, `strong`, `code`), `.ds-prose__lead`, `.ds-prose__code` for markup to copy.
- `.ds-link`: states `is-visited`, `is-hover`, `is-active`, `is-focus`; `.ds-link--quiet`, `.ds-link--strong`, `.ds-link--plain` (links in the menu frame); `.ds-links`; `.ds-icon`.
- `.ds-bullets`: `.ds-bullets__item`, `.ds-bullets__text`.
- `.ds-button`: `.ds-button--secondary`, `.ds-button--danger`, `.ds-button--large`; states `is-hover`, `is-pressed`, `is-focus`, `is-disabled`; `.ds-buttons`.
- `.ds-form`: `.ds-form__label`, `.ds-form__field`, `.ds-form__input` (`--short`, `is-error`), `.ds-form__select`, `.ds-form__textarea`, `.ds-form__check`, `.ds-form__hint`, `.ds-form__error`, `.ds-form__actions`.
- `.ds-table`: tight rows under a gradient head: `.ds-table__head` (`--num`), `.ds-table__cell` (`--num`, `--group`), `.ds-table__row--alt`, `.ds-table__desc`, `.ds-table__caption`, `.ds-table__swatch` (`--2`, `--3`, `--4`) a small sample square.
- `.ds-stat`: the row; four `.ds-stat__item` cells, each a heavy `.ds-stat__value` in `--color-heading-alt` over an uppercase `.ds-stat__label`.
- `.ds-list`: updates in a box with banded rows: `.ds-list__item` with a `.ds-list__meta` date and `.ds-list__body`; `.ds-list__more`.
- `.ds-panel`: `.ds-panel__title`, `.ds-panel__body`, `.ds-panel__text`, `.ds-panel__form` (also the search form of the menu frame). `.ds-ring` / `.ds-ring__item`: the web ring line.
- `.ds-grid`: four shelves to a row: `.ds-grid__cell` (`--2`, `--3`, `--4`) with a `.ds-grid__icon`, an uppercase `.ds-grid__title` link and `.ds-grid__text`.
- `.ds-tabs`: folder tabs on a thick line: `.ds-tabs__item` / `.ds-tabs__tab`, the current one (`is-current`) taller and in `--fill-bar`.
- `.ds-badge`: `.ds-badge--text`, `.ds-badge--count`, `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger`; `.ds-badges`.
- `.ds-awards`, `.ds-award` (`.ds-award__star`, `.ds-award__title`, `.ds-award__text`).
- `.ds-sidebar`: `.ds-sidebar__block` boxes with a `.ds-sidebar__title` strip and a `.ds-sidebar__list` of `.ds-sidebar__item` lines.
- `.ds-notice`: `.ds-notice__label`, `.ds-notice--error`; `.ds-notices`.
- `.ds-pagination`: `.ds-pagination__label`, `.ds-pagination__link` (`is-current` on `--fill-inverse`, `is-disabled`).
- `.ds-breadcrumb`: `.ds-breadcrumb__item`, the last `is-current`.
- `.ds-dialog`: the confirm message box in the page flow: `.ds-dialog__title`, `.ds-dialog__body`, `.ds-dialog__actions`.
- `.ds-empty`: `.ds-empty__title`, `.ds-empty__text` and a button.
- `.ds-footer`: at the foot of the main frame: rainbow bar, `.ds-footer__links` / `.ds-footer__item`, `.ds-footer__mail`, `.ds-footer__note`, `.ds-footer__legal`.
- `.ds-view`: one screen; `.ds-view--home` shows by default.

## Never

- `border-radius <= 0px`: every box and button is square.
- `box-shadow-blur <= 0px`: shadows are hard offsets.
- `border-width <= 4px`: the frame border is the thickest line.
- `font-size <= 44px`: the front-page title is the largest text.
- `font-size >= 10px`: small print is 10px.
- `font-weight >= 400`: regular, bold and black only.
- `font-families <= 4`: Verdana, Arial Black, Arial and one monospace.
- `content-width <= 784px`: both frames fit an 800px window.
- `transition = none`: nothing eases.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new page gets a bar in the menu frame; a new block in the main frame is a `.ds-heading` bar over a raised box; a new shelf is a `.ds-grid__cell`; a new figure joins `.ds-stat`. Keep the two frames, keep text small and bright on the dark tile, keep bars glossy and boxes square with hard shadows, and never add radius, blur, transitions or a menu across the top.
