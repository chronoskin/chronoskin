# Wiki dictionary, masthead look-up and letter strip, 2003 to 2010

## Summary

This is the dictionary or phrasebook wiki of the middle 2000s in a "paper" skin: one framed page with a masthead that holds the name and a look-up box, a dark bar of sections, a strip of letters from A to Z and, under them, the sheet. An entry is a headword on a double rule, tabs for its faces, numbered senses with quotations, and translation boxes that fold; the front page shows a word of the day on a calendar leaf. Everything is set in Georgia on cream with brown rules and underlined ink-blue links, square and still, with only a one-pixel hard shadow under boxes.

## Layout

- The page is fluid. `--size-page` is `100%`; `.ds-page__frame` has `--space-5` (14px) above and `--space-6` (20px) at the sides. Designed at 1024px; it does not reflow for phones.
- Everything sits in one bordered box, `.ds-page__book`: the header `.ds-nav` (masthead `.ds-nav__mast` in the secondary bar fill, section bar `.ds-nav__menu` in the bar fill, letter strip `.ds-index` on `--color-surface-alt`), then the sheet `.ds-page__sheet` on the canvas, then the footer band in the inverse fill. No text sits on the page colour. There is no side column of site links: navigation is the bar and the letters.
- Views. The specimen is a four-screen example dictionary; the bar has one item per view and the current one takes the accent fill. `home` is the main page: the word of the day, the figures line, then a wide column (welcome box, newest entries) beside a `--size-side` (232px) margin column of link boxes, and the shelves grid. `entry` is one entry: trail, headword, tabs, the senses and translations beside a margin column with a small facts table and a notice. `index` is the index for one letter: second-letter strip, the table of headwords, paging and an empty section. `request` is the request form with its two notices, a tips box and the withdraw confirmation.
- Inner pages (every view but `home`) start with `.ds-breadcrumb`, then `.ds-page-header` in place of the word of the day (title and one muted line at the left, the action at the right, on a double rule), then `.ds-tabs` when the page has several faces, then the working blocks. Blocks in a view are in `.ds-view__body`, `--space-5` apart. Two-column areas use `.ds-layout`: a fluid column and the `--size-side` margin column at the right, `--space-6` apart.
- Form rows are a `--size-label` (160px) label and the field; inputs are `--size-field` (300px), the look-up field `--size-search` (240px), a text area `--size-editor` (96px) tall. The calendar leaf is `--size-leaf` (84px) wide. A dialog is `--size-dialog` (440px) wide.
- Spacing scale: `--space-1` 2px, `--space-2` 4px, `--space-3` 6px, `--space-4` 10px, `--space-5` 14px, `--space-6` 20px, `--space-7` 28px.
  - `--space-1`: vertical padding of buttons, tabs, strips; gap between letters.
  - `--space-2`: vertical padding of table cells, list rows and bar items.
  - `--space-3`: gap between buttons, padding of notices and panel bodies.
  - `--space-4`: horizontal padding of cells, strips and boxes; gap between grid cells.
  - `--space-5`: gap between blocks; horizontal padding of bar items and tabs.
  - `--space-6`: padding of the sheet and masthead, gap between the columns, margin above a language heading.
  - `--space-7`: list indent.

## Typography and colour roles

- Text on fills follows the era's conventions (see the first layout of this era): no text on `--color-page`; body, muted, heading and link colours, `--color-danger` and `--color-success` on the canvas, the three surfaces, `--fill-panel` and the four fills; only `--color-bar-text` on `--fill-bar`, `--color-bar-alt-text` on `--fill-bar-alt`, `--color-inverse-text` on `--fill-inverse`, `--color-accent-text` on `--fill-accent`, `--color-notice-text` on `--color-notice`, and `--color-text` on `--color-danger-surface`. Links on a bar take the bar's text colour and are told apart by underline.
- `--color-page` is kraft brown with fine vertical lines from `--fill-page`; the sheet `--color-canvas` is cream. `--color-surface` fills margin boxes, tables and the calendar leaf; `--color-surface-alt` the letter strip, idle tabs, language headings and alternate rows; `--color-surface-strong` panel titles, translation bars and the current letter.
- `--fill-bar` (dark brown) with `--color-bar-text` is the section bar, table heads and the dialog title. `--fill-bar-alt` (straw) with `--color-bar-alt-text` is the masthead and the titles of margin boxes. `--fill-inverse` is the footer band.
- `--color-heading` is for titles and headwords; `--color-heading-alt` (rust) for part-of-speech headings, panel titles, figures and the drop capitals of the grid. Links are `--color-link` (ink blue) and underlined; a missing entry is `.ds-link--new`.
- `--color-border-strong` frames the book and rules the page header (as a `double` line of `--border-width-strong`); `--color-border` outlines boxes; `--color-border-muted` rules rows.
- `--color-accent` with `--color-accent-text` is the current bar item, the month strip of the calendar leaf, the current page number and badges; `--color-accent-alt` is the top rule of a notice.
- `--color-fill-1` to `--color-fill-4` are the straw, sage, clay and stone of the shelf cells.
- One family: `--font-body`, `--font-heading` and `--font-ui` are all Georgia, headings at weight 400; italic marks parts of speech, labels and quotations. `--font-mono` is for respellings.
- Sizes: `--text-base` 14px; `--text-ui` 12px for bar, tabs, buttons, tables and margin boxes; `--text-small` 11px for meta and the footer; `--text-large` 16px for the definition of the day, panel titles and shelf names; `--text-h3` 15px, `--text-h2` 20px, `--text-h1` 25px, `--text-display` 30px for the wordmark and the word of the day.
- One-off choices outside the token set: `double` rules on the hero, page header and footer; `dotted` row rules in tables; the notice's top rule is `--border-width-strong`; status badges are tints made with `color-mix()`.

## Components

- `.ds-page`: on `<body>`. `.ds-page__frame` is the padded page, `.ds-page__book` the framed box, `.ds-page__sheet` the sheet. Views are `.ds-view` with a `.ds-view__body` column; `.ds-layout` with `.ds-layout__main` and `.ds-layout__side` makes the two columns.
- `.ds-nav`: the header. `.ds-nav__mast` holds the brand, the look-up form `.ds-nav__search` (`.ds-nav__label`, `.ds-nav__field`) and the `.ds-nav__personal` list of `.ds-nav__personal-link` items. `.ds-nav__menu` is the bar of `.ds-nav__link` items (`is-current` in the accent fill); `.ds-nav__aux` items at the right, after `.ds-nav__rest`, are never current.
- `.ds-brand`: mark and name at the left of the masthead. `.ds-brand__mark` is a square tile 1.5 times `--text-display`, dressed like the primary button; `.ds-brand__text` stacks `.ds-brand__name` and the italic `.ds-brand__tag`. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo: replace both, keep the tile.
- `.ds-index`: the letter strip: `.ds-index__label`, `.ds-index__letter` (`is-current`, `is-empty`), `.ds-index__rest` for links at the right. `.ds-index--boxed` is the same strip as a box inside a view.
- `.ds-breadcrumb`: a small italic trail with `.ds-breadcrumb__link` and `.ds-breadcrumb__current`.
- `.ds-hero`: the word of the day between double rules: the calendar leaf `.ds-hero__leaf` (`.ds-hero__month`, `.ds-hero__day`), `.ds-hero__text` with `.ds-hero__kicker`, `.ds-hero__title`, `.ds-hero__pron` and `.ds-hero__lead`, and a stack of buttons `.ds-hero__action`.
- `.ds-page-header`: the head of an inner page: `.ds-page-header__title`, `.ds-page-header__desc`, `.ds-page-header__action`.
- `.ds-tabs`: the faces of an entry on a strong rule under the header; `.ds-tabs__tab` with `is-current` (canvas, open at the bottom) and `is-missing`.
- `.ds-stat`: the row; a ledger line of `.ds-stat__item` figures, each a `.ds-stat__value` beside an italic `.ds-stat__label`.
- `.ds-section`: a heading on a hairline between blocks.
- `.ds-panel`: a titled box: italic `.ds-panel__title` strip, `.ds-panel__body` with `.ds-panel__line` paragraphs and a right-aligned `.ds-panel__more`.
- `.ds-list`: dictionary lines with a hanging indent: `.ds-list__item` with a bold `.ds-list__term`, an italic `.ds-list__pos`, the gloss and a `.ds-list__meta` line.
- `.ds-sidebar`: the margin column of link boxes: `.ds-sidebar__block`, `.ds-sidebar__title`, `.ds-sidebar__list`, `.ds-sidebar__item`.
- `.ds-grid`: four shelf cells in a row: `.ds-grid__cell` (`--2`, `--3`, `--4`) with a drop capital `.ds-grid__icon`, a `.ds-grid__name` link and `.ds-grid__meta`.
- `.ds-prose`: the entry text: h1, h2 (a language, in a tinted strip with a strong left edge), h3 (a part of speech, on a faint rule), p, ul, ol, strong, code. `.ds-edit` is the bracketed edit link in a heading. `.ds-headline` with `.ds-headline__word` is the headword line; `.ds-senses` the numbered senses (`.ds-senses__item`, `.ds-senses__label`, `.ds-senses__quote`).
- `.ds-trans`: a translation box: `.ds-trans__bar` with a bold `.ds-trans__gloss` and a bracketed show or hide link; an unfolded box also has the two-column `.ds-trans__body`.
- `.ds-table`: a ledger table: bar-fill head, dotted row rules; `.ds-table__num` right-aligns numbers, `.ds-table__row--alt` tints a row, `.ds-table__term` is a bold headword cell, `.ds-table__pos` an italic one, `.ds-table__nowrap` keeps a cell on one line.
- `.ds-notice`: a slip with a thick top rule in `--color-accent-alt`; `.ds-notice--error` has it in `--color-danger` on `--color-danger-surface`. `.ds-notice__title` is the bold first phrase, `.ds-notice__link` an underlined link in the text colour.
- `.ds-form`: rows `.ds-form__row` of a `.ds-form__label` and a `.ds-form__field` holding a `.ds-form__input` (`--wide`, `--auto`, `--text`; `is-invalid`, `is-disabled`) with `.ds-form__hint` or `.ds-form__error`; `.ds-form__check` and `.ds-form__checkbox`; `.ds-form__actions` with `.ds-form__spacer` pushing the destructive button to the right.
- `.ds-link`: every link in text. States `is-visited`, `is-hover`, `is-active`, `is-focus`; `.ds-link--new` (in `--color-danger`) for a page nobody has written, `.ds-link--quiet` for tool and outside links, `.ds-link--strong` for bold.
- `.ds-button`: primary action. `.ds-button--secondary` is the plain button, `.ds-button--danger` the destructive one (error text on the error surface), `.ds-button--small` a smaller one. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. Group with `.ds-buttons`.
- `.ds-badge`: a small label in the accent fill for counts and "new". `.ds-badge--success`, `.ds-badge--warning` and `.ds-badge--danger` are status labels: a 14 to 22% tint of the status colour over `--color-surface`, body text and a 1px outline in the status colour.
- `.ds-pagination`: a line of small bordered `.ds-pagination__link` steps (`is-current` in the accent fill, `is-disabled` plain) with a muted `.ds-pagination__label`.
- `.ds-dialog`: a box in the page flow, centred in the `.ds-dialog__backdrop` band (`--color-overlay`, no text on it): `.ds-dialog__title` in the bar fill, `.ds-dialog__body`, `.ds-dialog__text`, `.ds-dialog__actions`. It never floats over the page.
- `.ds-empty`: a dashed box on `--color-surface` with `.ds-empty__title`, `.ds-empty__text` and one secondary button.
- `.ds-footer`: the colophon band: `.ds-footer__text` and a `.ds-footer__row` of `.ds-footer__link` items.
- `.ds-clear` ends floats; `.ds-icon` sizes an inline icon and `.ds-sprite` hides a symbol sheet, for projects that add icons.
- `ds-accordion` (with `__summary`, `__body`): a collapsible box, a `<details>` section for contents, tips or rules.
- `ds-menu` (with `__summary`, `__list`, `__item`): a page-actions or language drop-down, a `<details>` opened by click.
- `ds-tooltip` (with `__tip`): a small hint box on a link, shown on hover or focus.
- `ds-progress` (inside `ds-meter`): a native progress bar for completeness or a drive.

## Never

- `border-radius <= 2px`: every box is square; only badges have 2px corners.
- `box-shadow-blur <= 0px`: shadows are hard one-pixel offsets, never soft.
- `text-shadow = none`: no text shadow was measured in any reference.
- `transition = none`: hover states switch at once.
- `animation = none`: nothing moves.
- `font-size <= 30px`: the word of the day, 30px, is the largest text.
- `font-size >= 11px`: meta text is 11px; nothing is smaller.
- `font-families <= 2`: one serif for everything, monospace for respellings.
- `border-width <= 3px`: rules are 1px; double rules and the notice edge are 3px.
- `letter-spacing <= 0px`: no reference tracks its text.
- `uppercase-text <= 0%`: no text is transformed to uppercase.
- `line-height <= 1.6`: body text is set at 1.55.
- `content-width >= 75%`: the page stretches with the window.

## Extending

Derive a new component from the nearest one in the specimen: a box in the margin from `.ds-sidebar__block`, a folded section from `.ds-trans`, a titled box from `.ds-panel`, a message from `.ds-notice`. Build it as a 1px `--color-border` outline on `--color-surface`, keep text on fills as the conventions above say, and use tokens only: no raw colours or lengths, no gradient or transition of its own, no font size outside the type tokens, and `--space-5` as the gap to the next block.
