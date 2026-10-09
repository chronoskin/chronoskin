# Wiki hosting front page, centred sheet and create box, 2003 to 2010

## Summary

This is the front door of a wiki hosting service of the later 2000s, the kind that promised a class, a club or a team its own wiki in a minute: one rounded white sheet centred on a strongly coloured page, with a dark strip, the name and a short menu of rounded items at the top. The front page is a pitch, not a wiki page: an inverted band with a headline, three ticked promises and a small "create your wiki" form, then three large figures, four tinted reason cards and a list of busy wikis beside what people say; the inner pages are a tour, a table of plans, a directory with topics at the left and the sign-up form. Text is 13px Arial with bold geometric sans headings and underlined blue links, bars and buttons are glossy two-step gradients, corners are 5 to 10px and boxes cast a short soft shadow, with no motion.

## Layout

- The page is fluid. `--size-page` is `86%`: the sheet `.ds-page__sheet` is centred and the page colour shows at both sides, `--space-5` (16px) above and `--space-7` (36px) below. Designed at 1024px; it does not reflow for phones.
- Everything is inside the sheet, which clips to `--radius-page`. From the top: the header `.ds-nav` (the strip `.ds-nav__strip` in the bar fill with a tagline and the sign-in links; then `.ds-nav__top`, on the sheet, with the brand at the left and the menu `.ds-nav__menu` at the right over a hairline); the current view; the footer band in the secondary bar fill. No text sits on the page colour. There is no side column of site links and no tab row above the sheet.
- Views. The specimen is a five-screen example hosting service; the menu has one item per view and the current one is a pill in the accent fill. `home` is the pitch: the band `.ds-hero` (a fluid column and the `--size-create`, 318px, create box) from edge to edge of the sheet, then three figures, four reason cards and a 3 to 2 pair of the busy list and a quotes panel. `tour` explains a wiki in three numbered steps beside a `--size-side` (214px) column of panels. `plans` is a notice and the table of plans. `wikis` is the directory: topic boxes at the left, then tabs, the list, paging and an empty "your wikis" section. `create` is the sign-up form with two notices and the start-over confirmation, beside a column of panels.
- Inner pages (every view but `home`) start with `.ds-breadcrumb`, then `.ds-page-header` in place of the band (title and one muted line at the left, the action at the right, no rule), then `.ds-tabs` when the page has several faces, then the working blocks. A view's blocks are in `.ds-view__body`, `--space-5` apart with `--space-6` (24px) at the sides. `.ds-layout` puts a narrow column at the right, `.ds-layout--side` at the left; `.ds-stack` is a column of blocks `--space-4` apart.
- Form labels sit above their fields and `.ds-form__pair` puts two rows side by side; a lone input may take `--size-field` (300px); a text area is `--size-editor` (84px) tall. A dialog is `--size-dialog` (430px) wide, centred in a band in the flow.
- Spacing scale: `--space-1` 3px, `--space-2` 5px, `--space-3` 8px, `--space-4` 12px, `--space-5` 16px, `--space-6` 24px, `--space-7` 36px.
  - `--space-1`: vertical padding of buttons, menu items, tabs and inputs; gap under a label.
  - `--space-2`: vertical padding of table cells, list rows and the strip.
  - `--space-3`: gap between buttons, between an icon and its text, padding of notices.
  - `--space-4`: padding of cards and panels, gap between cards and between stacked blocks.
  - `--space-5`: gap between the blocks of a view, padding of the form and the create box.
  - `--space-6`: side padding of the sheet's regions, gap between columns.
  - `--space-7`: gap inside the band, left padding of the band, space under the sheet.

## Typography and colour roles

- The era's conventions for which text sits on which fill hold here as in every layout of the era.
  - `--color-page`: no text at all.
  - `--color-canvas`, `--color-surface`, `--color-surface-alt`, `--color-surface-strong`, `--fill-panel` and `--color-fill-1` to `--color-fill-4`: `--color-text`, `--color-text-muted`, `--color-heading`, `--color-heading-alt`, all link colours, `--color-danger` and `--color-success`.
  - `--fill-bar`: `--color-bar-text` only. `--fill-bar-alt`: `--color-bar-alt-text` only. `--fill-inverse`: `--color-inverse-text` only. Links on these fills take the same text colour and are underlined.
  - `--fill-accent`: `--color-accent-text`. `--color-notice`: `--color-notice-text`. `--color-danger-surface`: `--color-text`, with `--color-danger` for the title.
  - `--fill-button`: `--color-button-text`; `--fill-button-secondary`: `--color-button-secondary-text`; `--fill-input`: `--color-input-text`; `--color-disabled`: `--color-disabled-text`.
  - `--color-overlay` carries no text; a dialog on it is a `--color-canvas` box.
- Here `--color-page` is orange (#e9701a), lighter at the top, and the sheet white. `--color-surface` fills the form, tables, topic boxes and page numbers; `--color-surface-alt` alternate rows, idle tabs, list initials and inline code; `--color-surface-strong` the title strips of topic boxes.
- `--color-bar` is charcoal with white text: the top strip, the table head, the current tab and page number, the dialog title, and the thick rule under the tabs. `--color-bar-alt` (pale orange) is the footer band. `--color-inverse` (deep green) is the band of the front page, closed by a `--border-width-strong` rule in `--color-accent`.
- Links are `--color-link` (blue), underlined until hovered, when they turn `--color-link-hover` (orange); `--color-link-visited` is purple. A link to a page not yet written is `.ds-link--new` in `--color-danger`. Menu items, the breadcrumb and outside links are not underlined.
- `--color-border` outlines cards, boxes and tables; `--color-border-strong` the sheet, the create box, buttons and round icons; `--color-border-muted` divides rows.
- `--color-fill-1` to `--color-fill-4` (peach, lime, sky, rose) are the four reason cards. `--color-accent` (lime) is the current menu item, the round icons and step numbers, the badge and the left rule of a quotation. `--color-heading-alt` (burnt orange) is for figures, panel titles and h3; a notice is outlined in `--color-warning`.
- `--font-body` and `--font-ui` are the Arial stack; `--font-heading` is a geometric sans (Century Gothic, with Futura as the fallback), bold, for the name, titles, figures, card and panel titles and the table head; `--font-mono` is for wiki markup and addresses.
- Sizes: `--text-base` 13px text; `--text-ui` 12px menu, buttons, inputs, tabs, tables; `--text-small` 11px meta, strip, footer; `--text-large` 16px lead, the large button and the large input; `--text-h3` 14px; `--text-h2` 18px; `--text-h1` 24px page titles and figures; `--text-display` 30px name and headline, tracked by `--display-tracking`. Line height `--line-body` 1.5.
- One-off choices outside the token set: round icons, step numbers and notice marks have a `50%` radius; status badges are a 16 to 22% `color-mix()` of the status colour over `--color-surface`; the menu item has a `transparent` outline until hovered; the empty state has a `dashed` outline.

## Components

- `.ds-page`: on `<body>`. `.ds-page__sheet` is the centred sheet. `.ds-view` sections hold a `.ds-view__body`; `.ds-stack` is a column of blocks.
- `.ds-nav`: the header. `.ds-nav__strip` is the dark strip with `.ds-nav__account` and its underlined `.ds-nav__aux` links; `.ds-nav__top` holds the brand and `.ds-nav__menu`, whose `.ds-nav__link` items lead to views and take `is-current`.
- `.ds-brand`: the mark and name in a row. `.ds-brand__mark` is a tile 1.35 times `--text-display`, dressed like the primary button (`--fill-button`, `--color-button-text`, a `--color-border-strong` outline, `--radius-control`, `--shadow-control`); `.ds-brand__name` is in the heading colour. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo: replace both, keep the tile.
- `.ds-hero`: the band of the front page: `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__points` of `.ds-hero__point` lines with a `.ds-hero__tick`, the `.ds-hero__more` line with `.ds-hero__link`, and the create box `.ds-hero__box` (a sheet-coloured form with `.ds-hero__boxtitle`, the `.ds-hero__address` field and `.ds-hero__suffix`, the main button and `.ds-hero__note`). Only the front page has it.
- `.ds-page-header`: the head of an inner page: `.ds-page-header__title` and `.ds-page-header__desc` at the left, `.ds-page-header__action` at the right.
- `.ds-breadcrumb`: the small trail above the title, with `.ds-breadcrumb__link`, `.ds-breadcrumb__sep` and `.ds-breadcrumb__current`.
- `.ds-tabs`: rounded tabs on a thick bar-coloured rule. `.ds-tabs__tab` with `is-current` (bar fill, bar text).
- `.ds-prose`: explanatory text: h1, h2 (may start with a round `.ds-step` number), h3, p, ul, ol, strong, code. `.ds-sample` shows markup and its result in two `.ds-sample__side` halves (`--typed` for the markup), each with a `.ds-sample__label`.
- `.ds-link`: every link in text. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--new` for a missing page, `.ds-link--quiet` for outside links, `.ds-link--strong` for bold.
- `.ds-stat`: the row; three `.ds-stat__item` figures between two faint rules, each a large `.ds-stat__value` and a `.ds-stat__label` on one line.
- `.ds-grid`: four reason cards in a row, each `.ds-grid__cell` (`--2`, `--3`, `--4` for the other fills) with a round `.ds-grid__icon`, a `.ds-grid__name` and `.ds-grid__text`.
- `.ds-columns`: a 3 to 2 pair of columns. `.ds-layout` and `.ds-layout--side`: a working column and a narrow one. `.ds-section` is a heading on a hairline; `.ds-legend` a small muted line.
- `.ds-list`: wikis: `.ds-list__item` with a square `.ds-list__initial`, a bold `.ds-list__title` link over a `.ds-list__meta` line, and a `.ds-list__count` at the right.
- `.ds-panel`: a rounded card with a `.ds-panel__title` and a `.ds-panel__body`; `.ds-panel__line` is a paragraph, `.ds-panel__quote` and `.ds-panel__by` a quotation and its author, `.ds-panel__items` a list.
- `.ds-sidebar`: the column of topic boxes: `.ds-sidebar__block`, `.ds-sidebar__title`, `.ds-sidebar__list`, and `.ds-sidebar__item` rows with a `.ds-sidebar__count`.
- `.ds-table`: plans side by side: a bar-coloured head, row headings in bold, `.ds-table__num` for figures, `.ds-table__row--alt` for alternate rows and `.ds-table__row--foot` for the row of buttons.
- `.ds-badge`: a small accent label. `.ds-badge--success`, `.ds-badge--warning` and `.ds-badge--danger` are tinted status labels with a 1px outline in the status colour.
- `.ds-notice`: a rounded message outlined in `--color-warning` with a round `.ds-notice__mark`; `.ds-notice--error` is outlined in `--color-danger` on `--color-danger-surface`. `.ds-notice__title` is the bold first phrase, `.ds-notice__link` an underlined link in the text colour.
- `.ds-pagination`: centred rounded steps: `.ds-pagination__label`, `.ds-pagination__link` with `is-current` and `is-disabled`.
- `.ds-button`: primary action. `.ds-button--secondary` is the pale button, `.ds-button--danger` the destructive one (error text on the error surface), `.ds-button--large` the create button, `.ds-button--small` for a box. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. Group with `.ds-buttons`.
- `.ds-form`: a boxed form. `.ds-form__pair` holds two `.ds-form__row` columns; a row is a `.ds-form__label`, a `.ds-form__input` (`--large`, `--field`, `--text` for the text area; `is-invalid`, `is-disabled`) and a `.ds-form__hint` or `.ds-form__error`. `.ds-form__check` with `.ds-form__checkbox`; `.ds-form__actions` is the button row, with `.ds-form__spacer` pushing the destructive button to the right.
- `.ds-dialog`: a confirmation box in the page flow inside `.ds-dialog__backdrop`: `.ds-dialog__title`, `.ds-dialog__body`, `.ds-dialog__text`, `.ds-dialog__actions`. It never floats over the page.
- `.ds-empty`: a dashed box with `.ds-empty__title`, `.ds-empty__text` and one button.
- `.ds-footer`: the pale band that closes the sheet: `.ds-footer__text` and a `.ds-footer__row` of underlined `.ds-footer__link` items.
- `ds-accordion` (with `__summary`, `__body`): a collapsible box, a `<details>` section for contents, tips or rules.
- `ds-menu` (with `__summary`, `__list`, `__item`): a page-actions or language drop-down, a `<details>` opened by click.
- `ds-tooltip` (with `__tip`): a small hint box on a link, shown on hover or focus.
- `ds-progress` (inside `ds-meter`): a native progress bar for completeness or a drive.

## Never

- `border-radius <= 10px`: corners are 5 to 10px; only round icons and pills are rounder.
- `box-shadow-blur <= 8px`: shadows are short, 1 to 8px of blur.
- `text-shadow = none`: no text shadow was measured in any reference.
- `transition = none`: hover states switch at once.
- `animation = none`: nothing moves.
- `font-size <= 30px`: the name and the headline, 30px, are the largest text.
- `font-size >= 11px`: meta text is 11px; nothing is smaller.
- `font-families <= 3`: Arial, one geometric sans for headings, monospace for markup.
- `border-width <= 3px`: outlines are 1px; the rule under the band and the tabs is 3px.
- `letter-spacing <= 0.5px`: text is never spaced out; the headline is tightened by half a pixel.
- `uppercase-text <= 0%`: no text is transformed to uppercase.
- `line-height <= 1.6`: body text is set at 1.5.
- `content-width >= 60%`: the sheet takes most of the window.

## Extending

Derive a new component from the nearest one in the specimen: a card from `.ds-grid__cell` or `.ds-panel`, a box of links from `.ds-sidebar__block`, a message from `.ds-notice`, a row of things from `.ds-list__item`. Build it as a 1px `--color-border` outline on `--color-surface` or one of the four fills with `--radius-panel` and `--shadow-panel`, keep text on fills as the conventions above say, and use tokens only: no raw colours or lengths, no transition, no font size outside the type tokens, and `--space-5` as the gap to the next block.
