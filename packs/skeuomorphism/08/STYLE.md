# Skeuomorphism blog (kitchen notebook, paper slips on a table)

## Summary

A personal journal from the years when web pages imitated physical materials: no bar at all, only the site's mark and name centred on a table top, a ruled line of labels under them, and below that paper: a taped sheet for this week's entry, slips with a calendar leaf for the earlier ones, index cards down the side. It was a common dress for recipe, craft, travel and diary sites between about 2009 and 2013, when a blog wanted to look like the notebook it replaced. Text is a bookish serif at a comfortable size, and the richness is in the paper: punched holes, a red margin line, card dividers with coloured tabs, labels that look printed.

## Layout

- Fixed centred column: `--size-page` is 880px (`.ds-page__wrap` and `.ds-nav__inner`). Only the page texture and the footer board run the full window width.
- The header (`.ds-nav`) is not a bar: the brand centred on the page texture `--space-6` from the top, one italic `.ds-nav__tagline`, then `.ds-nav__items`, a centred line of labels between two scored rules. The footer is a dark board across the window. Both are written once.
- Order of the home view (`notebook`): the `.ds-hero` sheet across the column; the `.ds-stats` line of figures on the page texture; `.ds-page__columns`, with `.ds-page__aside` (`--size-aside`, 220px) on the left holding the `.ds-search` field and three `.ds-sidebar` cards, and `.ds-page__main` (630px) on the right, `--size-gutter` (30px) apart, holding a `.ds-page__title`, the `.ds-list` of entry slips and the `.ds-pagination` centred under them; then a second `.ds-page__title` and the `.ds-grid` of four divider cards across the column. The side column and the slips end near each other; keep three slips to three cards.
- Inner pages drop the hero, the figures and the side column. In their place: `.ds-page__trail` (the breadcrumb on the page texture) and one `.ds-page__sheet`, a punched sheet with a margin line, as wide as the column. The sheet opens with `.ds-page-header` (title and one italic line left, one or two keys right, a double rule under it) and continues in `.ds-section` parts, later ones with `.ds-section--ruled`. A section may be split by `.ds-section__split` into `.ds-section__main` and a `--size-side` (240px) `.ds-section__side` of stacked panels on the right. Tabs stand at the top of the section they sort.
- Views. The specimen is a four-screen site, one `.ds-view` showing at a time. `notebook` is the home view: this week's recipe, the figures, the latest three entries with the side cards, the recipe box. `recipe` is one entry: the table of ingredients, the method as running text, two panels beside them, and the empty card where readers' notes would be. `index` is the recipe index: divider tabs, the table of recipes with its testing labels, a count and the key to the labels. `send` is the form for sending a recipe, with its notices and keys, two panels beside it, and under it the dialog that the destructive key opens. The labels lead to `notebook`, `index` and `send`; `recipe` is reached from any title and keeps Recipe index marked.
- Spacing scale: `--space-1` 4px (inside labels, between lines of a card), `--space-2` 7px (cell padding, between keys), `--space-3` 10px (card padding, under titles), `--space-4` 15px (between slips and cards, key side padding), `--space-5` 20px (sheet padding, between side cards), `--space-6` 30px (above the masthead and the hero, the sheet's right padding), `--space-7` 44px (above the footer).
- Control sizes: `--size-pill` 22px (pagination keys, small keys, a divider tab), `--size-control` 30px (keys, fields, labels of the header), `--size-control-large` 42px (the hero key and the brand tile). `--size-leaf` (56px) is the calendar leaf, `--size-plate` (300px) and `--size-dish` (170px) the photograph in the hero. `--size-hair` (1px) is the offset of every highlight.

## Typography and colour roles

- `--font-body` for running text, tables, fields and card text, with italics for taglines, hints and the footer note; `--font-heading` for the name, the hero title, page, section and entry titles, figures and the day on a calendar leaf; `--font-ui` for labels, keys, tabs, table heads, meta lines and counts; `--font-mono` for code.
- Sizes: `--text-base` body, `--text-small` meta, labels, card text and the footer, `--text-ui` controls, `--text-large` the hero's lead and its key, `--text-display` the hero title only, `--text-h1` page titles, the name and the figures, `--text-h2` section and entry titles, `--text-h3` card titles.
- `--fill-page` is the table top; text that lies straight on it (the name, tagline, figures, breadcrumb, `.ds-page__title`) uses `--color-heading`, `--color-heading-alt` and `--color-text-muted`, with a one-pixel highlight below mixed from `--color-canvas`. `--color-canvas` is the punched sheet; `--color-surface` (through `--fill-panel`) slips, cards, the hero and table rows; `--color-surface-alt` alternate rows, the empty card and inline code; `--color-surface-strong` the count label.
- `--fill-bar` with `--color-bar-text` is small and rare: the current label of the header, the month strip of a calendar leaf, the table head. `--fill-bar-alt` with `--color-bar-alt-text` (titles in `--color-heading-alt`) is the title strip of a card, a tab that is not current and the dialog's title. `--fill-inverse` with `--color-inverse-text` is the footer board only.
- `--color-fill-1` to `--color-fill-4` colour the tabs and upper edges of the four divider cards (text on a tab is `--color-inverse-text`); `--color-fill-2` is also the cloth in the hero photograph, `--color-inverse-text` its check and the plate, and `--color-accent-alt` what is on the plate and the tape.
- `--color-button` is the one primary key; `--fill-button-secondary` the other keys, the select and pagination. `--color-accent` is the printed label; `--color-success`, `--color-warning` and `--color-danger` the testing labels, the destructive key, the upper edge of the dialog and, thinned, the margin line of a sheet. Text on the warning label is `--color-warning` darkened most of the way to `--color-shadow`.
- One-off values built in `components.css` from tokens: the three punched holes and the margin line of a sheet (gradients of `--color-page`, `--color-shadow` and `--color-danger`), the check of the cloth, the dish, the tape (rotated two degrees), the double rule under a page header (a `--border-width-strong` offset shadow of `--color-border`), `dashed` for the empty card, `uppercase` and 1px tracking for the hero's kicker and the month, `italic`, and scored rules mixed from `--color-shadow` and `--color-canvas`.

## Components

- `.ds-page`: on `<body>`; paints the table top. `.ds-page__wrap` is the 880px column; `.ds-page__columns` with `.ds-page__aside` (left) and `.ds-page__main`; `.ds-page__title` a heading on the table top (`--spaced` when it follows a block); `.ds-page__trail` the breadcrumb's place; `.ds-page__sheet` the punched sheet of an inner page. `.ds-section` is one part of a sheet (`--ruled` after another), with `.ds-section__title`, `.ds-section__split` into `.ds-section__main` and `.ds-section__side`, and `.ds-section__foot` (a `.ds-section__note` left, labels right).
- `.ds-view`: one screen of the site, a `<section>` with an id; only the one named in the address shows, and `.ds-view--home` when none is named. Always a plain block.
- `.ds-nav`: the masthead. `.ds-nav__inner` centres the brand, the `.ds-nav__tagline` and `.ds-nav__items`, a line of `.ds-nav__item` labels. The current one is a small stitched label: `is-current`, or in the specimen one selector per view. `is-hover` lays a pale patch under a label.
- `.ds-brand`: the mark and name, centred at the top. `.ds-brand__mark` is a 42px tile dressed as the primary key; `.ds-brand__name` the name in the heading font, always bold. Both are placeholders for the installing project's own name and logo.
- `.ds-hero`: this week's entry, once, at the top of the home view: a sheet held by a strip of tape. `.ds-hero__plate` is the photograph (`.ds-hero__cloth`, `.ds-hero__dish`, `.ds-hero__soup`, all drawn); `.ds-hero__copy` holds `.ds-hero__kicker`, `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__actions` (the one `.ds-button--large` and at most one link) and `.ds-hero__note`.
- `.ds-page-header`: the head of a sheet in place of the hero: `.ds-page-header__copy` with `.ds-page-header__title` and one `.ds-page-header__text`; `.ds-page-header__actions` with at most two keys, the primary last.
- `.ds-prose`: running text: h1, h2, h3, paragraphs, lists, `code`, links.
- `.ds-link`: text link, underlined; states `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--quiet` for breadcrumb and meta links. `.ds-links` is a wrapping row.
- `.ds-button`: an enamelled label. `.ds-button--secondary`; `.ds-button--danger` for a destructive action, set apart with `.ds-buttons__end`; `.ds-button--large` for the hero; `.ds-button--small`. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-button__icon` is a 16px glyph. `.ds-buttons` is a row.
- `.ds-form`: `.ds-form__row` with a `.ds-form__label` column and a `.ds-form__field`. `.ds-form__input` (`--short` for a number) and `.ds-form__textarea` are fields with a line to write on, `.ds-form__select` a secondary key. `is-error` with `.ds-form__error`; `.ds-form__hint`; `.ds-form__check` with `.ds-form__checkbox`; `.ds-form__actions` under a rule. `.ds-search` is one field and one key at the head of the side column.
- `.ds-table`: framed, ruled table. `.ds-table__head` cells are the stitched band; `.ds-table__row` with `is-first` and `is-alt`; `.ds-table__cell`, `.ds-table__cell--name`, `is-numeric`.
- `.ds-list`: the entries, each `.ds-list__item` its own slip: a `.ds-list__date` calendar leaf (`.ds-list__month`, `.ds-list__day`) and a `.ds-list__body` with `.ds-list__title`, `.ds-list__meta`, `.ds-list__text` and a `.ds-list__foot` (labels left, one link right). Slips lie `--space-4` apart.
- `.ds-panel`: an index card with a `.ds-panel__title` strip and `.ds-panel__body`; `.ds-panel__text`, or `.ds-panel__rows` of `.ds-panel__row` (`is-first`) with `.ds-panel__key` and `.ds-panel__value`.
- `.ds-stat`: one figure: `.ds-stat__value` over a `.ds-stat__label`, centred. Three to five sit in the `.ds-stats` line between two scored rules on the table top; `is-first` on the first.
- `.ds-grid`: four `.ds-grid__cell` divider cards across the column (`--2`, `--3`, `--4` for the other fills and tab positions). Each has a `.ds-grid__tab` standing on its upper edge, a `.ds-grid__title`, a `.ds-grid__text` and a `.ds-grid__more` link.
- `.ds-tabs`: card dividers standing on a rule: `.ds-tabs__tab`; the `is-current` one is paper coloured, a little taller, and joins the sheet below.
- `.ds-badge`: a small printed label. `.ds-badge--count` is a number on a sunk patch; `.ds-badge--success` (tested), `.ds-badge--warning` (still testing) and `.ds-badge--danger` (withdrawn) are the testing labels. `.ds-badges` is a row.
- `.ds-sidebar`: an index card in the side column: `.ds-sidebar__title` strip, optional `.ds-sidebar__text`, and a `.ds-sidebar__list` of ruled `.ds-sidebar__item` rows (`is-first`), each a link and an optional `.ds-sidebar__meta`.
- `.ds-notice`: a note stuck on the sheet; `.ds-notice__title` for the bold lead-in; `.ds-notice--error`. `.ds-notices` wraps several.
- `.ds-pagination`: small keys centred under the entries, `.ds-pagination__link`; `is-current` is pressed, `is-disabled` is flat.
- `.ds-breadcrumb`: on the table top above a sheet, of quiet links, `.ds-breadcrumb__sep` and `.ds-breadcrumb__current`.
- `.ds-dialog`: a card with a red upper edge: `.ds-dialog__title` strip with `.ds-dialog__close`, `.ds-dialog__body` with `.ds-dialog__heading` and `.ds-dialog__text`, `.ds-dialog__actions` with the confirming key last. `.ds-dialog__backdrop` is the dimmed area behind it.
- `.ds-empty`: a blank card with a dashed edge: `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one secondary key.
- `.ds-footer`: the dark board: centred `.ds-footer__links` of `.ds-footer__link` and an italic `.ds-footer__note`.
- `.ds-switch`: glossy sliding ON/OFF switch over a checkbox (`.ds-switch__track`, inset, with a raised knob), set as rows of a `.ds-switches` list (`.ds-switches__row`, `.ds-switches__name`). Here: the Promise row of the send form.
- `.ds-progress`: inset `<progress>` bar with a glossy fill and a `.ds-progress__label` line. Here: the Photograph row of the send form.
- `.ds-tooltip`: `?` key with a dark glossy bubble and arrow (`.ds-tooltip__trigger`, `.ds-tooltip__bubble`) shown on hover or focus. Here: on the Promise label.
- `.ds-avatar`: framed photo with an inner shadow (`.ds-avatar__photo`, `.ds-avatar__glyph`), used in a `.ds-testimonial` (`.ds-testimonial__text`, `.ds-testimonial__who`). Here: a reader's box in the notebook side column.
- `.ds-menu`: popover list with an arrow, hanging from a toolbar button (`<details>`, `.ds-menu__button`, `.ds-menu__list`, `.ds-menu__item`). Here: the header toolbar of a recipe.

## Never

- `border-radius <= 5px`: paper has nearly square corners; keys are 5px.
- `border-width <= 3px`: hairlines, and a 3px coloured edge on a divider card and the dialog.
- `box-shadow-blur <= 10px`: paper casts a short, close shadow.
- `font-size <= 38px`: the hero title is the largest text.
- `font-size >= 12px`: meta text is the smallest.
- `font-weight >= 400`: no light weights.
- `font-families <= 4`: a book serif, a typewriter face for titles, a humanist sans for labels and a monospace.
- `line-height <= 1.65`: body text is set at 1.6.
- `uppercase-text <= 5%`: capitals only on the kicker and the calendar month.
- `letter-spacing <= 1px`: nothing is tracked beyond the kicker.
- `transition = none`: states change at once.
- `animation = none`: nothing moves.
- `content-width <= 900px`: everything stays inside the 880px column.
- `row-gap <= 20px`: slips lie 15px apart; table rows touch.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new thing on the table is paper: `--fill-panel`, a 1px border in `--color-border-strong`, `--radius-panel` and `--shadow-panel`; if it has a title, a `--fill-bar-alt` strip. Something that must stand out is a label: a small `--fill-accent`, status or `--fill-bar` patch with text in the label font. A long inner page is another `.ds-page__sheet` with sections; do not put a bar across the window, do not round corners beyond a key's, and do not add soft floating shadows, new font families or new colours.
