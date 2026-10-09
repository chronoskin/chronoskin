# Board file archive, 2002 to 2007

## Summary

This is the download section that boards of the early 2000s grew beside their forums: a database of skins, button sets and add-ons, laid out with the same ruled boxes and title bars as the board itself. One bar carries the name and the sections, a featured file opens the index, and under it the categories stand as a directory of cells across the full width, followed by the newest files beside a numbered list of the most downloaded. The skin is an earth one: rust brown bars with a light top edge on sand and cream cells, serif titles over small Lucida text, square boxes with a hard one-pixel shadow, and no motion.

## Layout

- The page is fluid. `--size-page` is `100%`; `.ds-page` has `--space-5` (12px) above and below and `--space-6` (18px) at the sides, so the sheet `.ds-page__frame` takes about 96% of the viewport. Designed at 1024px. Never centre a narrow column.
- Order from the top, all inside the sheet: `.ds-nav__head`, one primary bar with the brand on the left and the sections as tabs standing on its lower edge at the right; `.ds-nav__sub`, a secondary strip with the visitor line on the left and the search box on the right; `.ds-page__body` with the views; the footer band; one line of small print `.ds-page__legal`. No text sits on the page colour.
- A view is a `.ds-view` section whose `.ds-view__body` stacks its blocks `--space-5` apart. Full-width blocks come first; where two columns are needed, `.ds-layout` puts `.ds-layout__main` (at least `--size-main`, 440px) beside `.ds-layout__rail` (`--size-rail`, 220px) on the right, and lets the rail drop under the main column when the window is too narrow for both.
- The category directory `.ds-grid` is as many equal columns as fit at `--size-cell` (250px) each: three at 1024px. A file row is at least `--size-row` (30px); bars and title strips are `--size-bar` (24px); count columns of a table are `--size-count` (64px) and its date column `--size-date` (104px).
- Spacing scale: `--space-1` 1px, `--space-2` 3px, `--space-3` 5px, `--space-4` 8px, `--space-5` 12px, `--space-6` 18px.
  - `--space-1`: the gap between cells and rows inside any bordered block.
  - `--space-2`: vertical padding of bars, strips, rows and buttons; gap between stacked small lines.
  - `--space-3`: vertical padding of cells and tabs, gap between inline items.
  - `--space-4`: horizontal padding of cells and strips, padding of directory cells and panel bodies, gap between buttons.
  - `--space-5`: body padding and the gap between neighbouring blocks. This is the only block gap.
  - `--space-6`: side padding of the page, list indents, padding of dialog and empty-state bodies, horizontal padding of tabs.
- Inner pages keep the bar, the strip and the footer and drop the featured file. From the top: `.ds-breadcrumb`, then `.ds-page-header` in place of the hero (title and one line on the left, the action or the file buttons on the right, no box), then `.ds-tabs` when the list has several orders, then the working block. A category is one full-width `.ds-table` with `.ds-toolbar` under it; a file, the upload form and the member's files use `.ds-layout`, with details, notes or the empty favourites box in the rail.
- Views. The specimen is a five-screen example archive; the bar has one tab per view and marks the current one. `files` (home) is the archive index: featured file, notice, category directory, latest files beside the most downloaded, the figures strip and who is browsing. `category` is one category: sort tabs, the file table with its two groups, the count line and page numbers. `file` is one file: download buttons, the delete confirmation that the Delete button opens, screenshots, the description as long text, file details and more by the author in the rail, then the comments. `upload` is the upload form as it came back with an error, beside the upload notes. `mine` is the member's files: figures, the list of uploads with their status, and an empty favourites box.
- Form rows are a `--size-label` (200px) label cell and a field cell that takes the rest; text inputs are `--size-field` (300px) unless `--wide` or `--auto`. A dialog is `--size-dialog` (60%) wide, centred in the flow on a band of the overlay colour.

## Typography and colour roles

- Three families. `--font-body` (Lucida Sans Unicode, then Lucida Grande and Verdana) for all running text; `--font-heading` (Georgia) for the wordmark, page titles and every title strip; `--font-ui` (Tahoma) for navigation, tabs, buttons, inputs, table heads and badges. `--font-mono` is for inline code only.
- Sizes: `--text-small` 10px for descriptions, meta, strips, counts and most of the page; `--text-base` and `--text-ui` 11px for running text, buttons and navigation; `--text-h3` 12px; `--text-large` and `--text-h2` 13px for file and category names and for title strips; `--text-h1` 18px for the page title and the featured file; `--text-display` 22px for the wordmark only.
- `--line-body` is 1.5, `--line-heading` 1.2. Headings are bold Georgia, never uppercase, with no tracking; the wordmark is regular weight (`--weight-display` 400).
- Links in running text are `--color-link` and underlined, and lose the underline on hover. Names of files and categories, breadcrumb links and member names use `--color-link-quiet` without underline (`--link-decoration-quiet`). Links on the secondary strip, in a notice, in a list head and in the footer take the colour of their bar and are always underlined. Hover turns a link `--color-link-hover`.
- Surfaces: `--color-page` outside the sheet, `--color-canvas` for the sheet. `--color-surface` is the main cell; `--color-surface-alt` the count and date cells, form labels, the featured file's side cells and the figures; `--color-surface-strong` the group rows, panel and rail titles and the form action row. Alternate rows use `--color-fill-3`. `--fill-panel` shows in the 1px gaps between cells.
- Bars: `--fill-bar` with `--color-bar-text` for the top bar, table heads, the directory title, list heads, form and dialog titles, the current tab of `.ds-tabs`. `--fill-bar-alt` with `--color-bar-alt-text` for the strip under the bar, idle tabs, comment heads, rail foot strips and quiet badges. `--fill-inverse` with `--color-inverse-text` for the footer. The wordmark stands on the primary bar, so it is `--color-bar-text`.
- The current section tab is a `--color-surface` cell with `--color-heading` text and a `--color-border-strong` outline on three sides.
- Borders: blocks have a 1px `--color-border` outline; `--color-border-muted` divides rows inside a rail box and outlines the breadcrumb; `--color-border-strong` outlines the sheet and buttons, and at `--border-width-strong` (2px) rules the bottom of the bar, the tab row, the top of the footer and the dialog.
- Markers: `--color-accent` for folder and file icons and the badge; `--color-accent-alt` for the kicker over the featured file and "new" remarks; `--color-heading-alt` for rail titles and rank numbers. Status badges are a `color-mix()` of 30% of the status colour over `--color-surface` with `--color-text` and an outline in the status colour. The destructive button is `--color-danger` on `--color-danger-surface`.
- `--color-fill-1` to `--color-fill-4` tint the directory cells; in this palette they are four steps of sand.
- One-off values outside the token set: the status badge tints above; the notice edge is an equal `color-mix()` of `--color-accent` and `--color-notice`; the screenshot placeholder `.ds-thumb` is drawn at a 4 to 3 ratio with rows sized in percent; always-underlined links on bars write `underline` directly.
- The references set Verdana with Lucida and Lucida Grande as fallbacks; this pack leads with the Lucida faces of the same stacks and takes Georgia from the one reference stack that names it. All are system fonts.

## Components

- `ds-progress`: a quota or poll bar, a native `<progress>` with a label above (private-message quota).
- `ds-accordion`: a collapsible `<details>` entry for a forum jump or an FAQ answer.
- `ds-avatar`: the square picture of a member, placed under or beside the name.
- `.ds-page`: on `<body>`. Wrap everything in `.ds-page__frame`; `.ds-page__body` holds the views. Each view is a `.ds-view` section (`.ds-view--home` for the index) and its blocks go in `.ds-view__body`. `.ds-page__legal` is the small print at the foot of the sheet. `.ds-layout` with `.ds-layout__main` and `.ds-layout__rail` makes the two-column area. `.ds-icon` sizes an inline SVG icon, `.ds-sprite` hides the symbol sheet, `.ds-block` puts a span on its own line, `.ds-scroll` lets a wide table scroll inside its own box.
- `.ds-nav`: the site header. `.ds-nav__head` is the primary bar holding the `.ds-brand` and `.ds-nav__menu`, a row of `.ds-nav__link` tabs; the current one takes `is-current`, and in the specimen one `:has()` rule per view applies the same dress. `.ds-nav__sub` is the secondary strip with `.ds-nav__sublink` links and the `.ds-nav__search` form, whose input also takes `.ds-nav__field`.
- `.ds-brand`: the site's mark and name at the left of the bar; one per page. `.ds-brand__mark` is a square tile dressed like the primary button (`--fill-button`, `--color-button-text`, a `--color-border-strong` outline, `--radius-control`, `--shadow-control`), 1.3 times `--text-display`. `.ds-brand__name` is the name in `--font-heading` at `--text-display`. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo: replace both, keep the tile.
- `.ds-breadcrumb`: one pale bordered line at the top of an inner page, with `.ds-breadcrumb__link`, `.ds-breadcrumb__sep` and `.ds-breadcrumb__current`.
- `.ds-hero`: the featured file at the top of the index, one ruled block of three cells: `.ds-hero__shot` with a `.ds-thumb`, `.ds-hero__text` with `.ds-hero__kicker`, `.ds-hero__title`, `.ds-hero__lead` and the download button in `.ds-hero__action`, and `.ds-hero__aside` with the file's figures. Never a banner.
- `.ds-thumb`: a screenshot placeholder, a small drawing of a board built from `.ds-thumb__bar`, `.ds-thumb__strip` and `.ds-thumb__row` (`--alt`). Replace it with a real image in a project. `.ds-shots` is a row of `.ds-shots__item` thumbs with captions.
- `.ds-page-header`: the head of an inner page. `.ds-page-header__text` holds `.ds-page-header__title` and one `.ds-page-header__desc` line; `.ds-page-header__action` or a `.ds-buttons` group sits at the right.
- `.ds-prose`: long text such as a file description: h1, h2, h3, p, ul, strong, code, hr.
- `.ds-link`: every link in text. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--quiet` for names and tools, `.ds-link--strong` for bold.
- `.ds-button`: the primary action, brown with white bold text. `.ds-button--secondary` is the plain sand button. `.ds-button--danger` is for a destructive action such as deleting a file; use it once per confirmation, beside a secondary button. `is-hover`, `is-active`, `is-focus`, `is-disabled` (or `disabled`). Group buttons in `.ds-buttons`.
- `.ds-form`: a ruled block of rows. `.ds-form__title`, `.ds-form__row` with `.ds-form__label` (plus `.ds-form__hint`) and `.ds-form__field`; controls take `.ds-form__input` (`--wide`, `--auto`, `is-invalid`, `is-disabled`), errors `.ds-form__error`, checkboxes `.ds-form__check` and `.ds-form__checkbox`, buttons in `.ds-form__actions`.
- `.ds-table`: the file list of a category and any data table. `th` is the primary bar (`.ds-table__head--left` for text columns); `.ds-table__group` is a group row; `.ds-table__row--alt` the alternate row; cells `.ds-table__num` and `.ds-table__last`; inside the main cell `.ds-table__title`, `.ds-table__desc`, `.ds-table__meta`.
- `.ds-list`: file rows outside a table. `.ds-list__head` (with an optional `.ds-list__more` link at its right), then `.ds-list__item` (`is-alt` for the alternate row) with a `.ds-list__icon`, `.ds-list__title`, `.ds-list__meta` and `.ds-list__aside` for size and date.
- `.ds-panel`: any titled box. `.ds-panel__title` (add `--bar` for a main block), `.ds-panel__sub`, `.ds-panel__body` (`--alt`), `.ds-panel__line`. `.ds-spec` is the label and value list of a file inside a panel: `.ds-spec__label`, `.ds-spec__value`.
- `.ds-stat`: summary figures as one ruled strip of equal cells. `.ds-stat__title` is optional; each `.ds-stat__item` has a `.ds-stat__value` over a `.ds-stat__label`. Use three to six.
- `.ds-grid`: the category directory, cells with 1px gaps in one bordered block. `.ds-grid__title`, `.ds-grid__cell` (`--2`, `--3`, `--4` for the other fills) with `.ds-grid__icon`, `.ds-grid__name`, `.ds-grid__count`, `.ds-grid__meta` and a `.ds-grid__subs` line of sub-category links.
- `.ds-tabs`: a row of `.ds-tabs__tab` items on a 2px rule; `is-current` takes the primary bar fill. Use for the orders or filters of one list.
- `.ds-badge`: a small square label in the accent fill. `.ds-badge--quiet` is the secondary bar. `.ds-badge--success` (approved, updated), `.ds-badge--warning` (awaiting check, beta) and `.ds-badge--danger` (broken, sent back) are tints with an outline. `.ds-new` is bold text for a "new" remark.
- `.ds-sidebar`: the rail's boxes. Each `.ds-sidebar__block` has a `.ds-sidebar__title`, a `.ds-sidebar__list` of ruled `.ds-sidebar__item` rows (`.ds-sidebar__rank` for a number, `.ds-sidebar__name` for the link, `.ds-sidebar__count` for a figure) and an optional `.ds-sidebar__foot` with a `.ds-sidebar__footlink`.
- `.ds-notice`: a pale yellow line for archive news; `.ds-notice--error` is the error box with a `.ds-notice__title`. `.ds-notice__link` is a link inside either.
- `.ds-pagination`: `.ds-pagination__label`, then small bordered `.ds-pagination__link` boxes; `is-current` for the current page. `.ds-toolbar` puts a count line on the left and the page numbers on the right.
- `.ds-comment`: comments under a file, one ruled block: a `.ds-comment__head` strip per comment, then `.ds-comment__body` (`--alt`) with `.ds-comment__text` and an optional `.ds-comment__sig` signature line.
- `.ds-dialog`: a confirmation box in the page flow with a 2px outline and a hard shadow: `.ds-dialog__title`, `.ds-dialog__body`, `.ds-dialog__text`, `.ds-dialog__actions`, on a `.ds-dialog__backdrop` band. It never floats over the page.
- `.ds-empty`: a ruled block with one centred message and one button for a list with nothing in it: `.ds-empty__title`, `.ds-empty__body`, `.ds-empty__text`.
- `.ds-footer`: the dark band closing the sheet: a `.ds-footer__row` with a `.ds-footer__group` of `.ds-footer__link` items and `.ds-footer__sep` marks on the left and the time line on the right.

## Never

- `border-radius <= 0px`: every box is square in all references.
- `box-shadow-blur <= 0px`: the only shadows are hard one-pixel offsets, as the references drew them with images.
- `text-shadow = none`: no text shadow was measured in the period references.
- `transition = none`: states switch at once.
- `animation = none`: nothing moves.
- `font-size <= 22px`: only the wordmark is 22px.
- `font-size >= 10px`: 10px is the smallest size and the most common.
- `font-weight <= 700`: bold is the heaviest weight.
- `font-families <= 4`: Lucida, Georgia, Tahoma and a monospace for code.
- `border-width <= 2px`: 1px everywhere, 2px on emphasised edges.
- `letter-spacing <= 0px`: no tracking.
- `uppercase-text <= 0%`: nothing is transformed to uppercase.
- `underlined-links >= 30%`: links in text are underlined; only names and titles are not.
- `gradient-fills <= 20%`: gradients are on bars, strips and buttons only.
- `row-gap <= 1px`: rows touch or are separated by a 1px line.
- `block-gap <= 20px`: neighbouring blocks are 12px apart.
- `content-width >= 85%`: the sheet takes about 96% of the viewport.

## Extending

Derive a new component from the nearest one in the specimen: a new titled box from `.ds-panel`, a new row layout from `.ds-list` or `.ds-table`, a new strip from `.ds-nav__sub`, a new figure list from `.ds-spec`. Build it as a 1px `--color-border` outline on `--fill-panel`, with cells in the three surface colours `--space-1` apart and a title in `--fill-bar` or `--color-surface-strong`. Keep text tokens on the fills named above (bar text on bars, heading text on the strong surface, quiet links on cells). Use tokens only: no raw colours or lengths, no new radius, shadow or transition, no font size outside the type tokens, and `--space-5` as the gap to the next block.
