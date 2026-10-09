# Bulletin board forum, member profile and private messages, 2002 to 2007

## Summary

This is the member's side of a bulletin board in the first half of the 2000s: the profile page other members see, and the private message centre with its folders, inbox table and compose form. The profile opens with a head box carrying the avatar, name and rank, a ruled line of post figures, and then two equal columns of titled boxes; the messenger pages put a folder menu on the left of the working block. Brick red title bars with a darker lower half stand on warm grey cells, text is Trebuchet MS at 13px and 11px, every box is square and ruled, and nothing moves.

## Layout

- The page is fluid. `--size-page` is `100%`; `.ds-page` has `--space-6` (15px) at the sides and the foot and none at the top, so the sheet `.ds-page__frame` hangs from the top edge of the window under a `--border-width-strong` rule in `--color-border-strong`. Designed at 1024px. Never centre a narrow column.
- Order from the top, all inside the sheet: the masthead `.ds-nav__head` (brand and tagline on the left, the signed-in member's box `.ds-nav__member`, `--size-member` 300px wide, on the right); the navigation `.ds-nav__menu`, one bar whose items share the full width equally and are centred, 1px apart; `.ds-page__body`, which holds the views; the footer band; the small print `.ds-page__legal`. No text sits on the page colour.
- The profile (home view) is one column of blocks `--space-5` apart: an optional `.ds-notice`, the profile head `.ds-hero`, the `.ds-stat` line, then `.ds-columns`, two equal columns of boxes with `--space-5` between them and between the boxes of a column. Fill both columns so they end near each other. There is no sidebar on the profile.
- Inner pages (inbox, a message, the compose form, the buddy lists, help) keep the masthead, navigation and footer and drop the hero and the two columns. From the top: `.ds-breadcrumb` (a small open line on a faint rule), `.ds-page-header` in place of the hero (a surface box with a strong left edge: title and one line of description on the left, the main action on the right), an error `.ds-notice` when there is one, then `.ds-layout`: the menu `.ds-sidebar` in a `--size-menu` (190px) column on the left and the working blocks in `.ds-layout__main`, `--space-5` apart. In the inbox that is `.ds-tabs`, the folder `.ds-table`, and a `.ds-toolbar` with the folder actions on the left and `.ds-pagination` on the right. A message is one `.ds-post`, the confirmation `.ds-dialog` in the flow under it, and a `.ds-toolbar`. Give the menu enough blocks to end near the foot of the main column.
- Views. The specimen is a six-screen example; the navigation bar has one item per view and marks the current one. `profile` (home) is the member's profile: notice, head box, post figures, forum information and latest posts on the left, contact box with a quick message line, buddies and signature on the right. `inbox` is the private message inbox: folder menu with the storage meter, folder tabs, the message table, the folder actions and page numbers. `message` is one message with its delete confirmation and reply buttons. `compose` is the compose form as it came back with an unknown recipient. `buddies` is the buddy list as a grid of member cards over an empty ignore list. `help` is the help page: contents menu beside the long text and a line of related links.
- Folder table columns: checkbox and status `--size-check` (26px) each, subject (fluid), sender `--size-from` (130px), date `--size-date` (150px), size `--size-count` (64px, right-aligned). Rows are at least `--size-row` (28px); bars and title strips `--size-bar` (24px). A post has a `--size-from` author cell. The avatar tile is `--size-avatar` (84px); `.ds-avatar--small` is 40% of it.
- Spacing scale: `--space-1` 1px, `--space-2` 2px, `--space-3` 4px, `--space-4` 7px, `--space-5` 11px, `--space-6` 15px.
  - `--space-1`: the ruled gap between cells inside every bordered block, and between navigation items.
  - `--space-2`: vertical padding of bars, strips, table cells, buttons and inputs; gap between tabs and page numbers.
  - `--space-3`: vertical padding of fact rows, list items and menu items; gap between inline items.
  - `--space-4`: horizontal padding of cells and strips, padding of panel bodies, gap between buttons.
  - `--space-5`: body padding, padding of the head box, and the gap between neighbouring blocks and columns. This is the only block gap.
  - `--space-6`: page gutter, list indent, padding of the dialog and the empty state.
- Form rows are a right-aligned `--size-label` (30%) label cell and a field cell; text inputs are `--size-field` (260px) unless `--wide`; the action row starts under the fields. A dialog is `--size-dialog` (50%) wide, centred in the flow. A meter is `--size-meter` (10px) high.

## Typography and colour roles

- One family: `--font-body`, `--font-heading` and `--font-ui` are the same Trebuchet MS stack. `--font-mono` is for inline code only.
- Sizes: `--text-small` 11px for meta, table and fact cells, menus and strips; `--text-base` 13px for running text, titles in lists and tables; `--text-ui` 12px for buttons, inputs, navigation and tabs; `--text-h2` 13px for bar titles; `--text-large` 14px for figures and the sender's name; `--text-h1` 16px for the page heading; `--text-display` 18px for the wordmark and the member's name in the profile head. Nothing is larger than 18px.
- `--line-body` is 1.3. Headings are bold, never uppercase, no tracking. `--weight-ui` is 700 for navigation, tabs and the primary button; the secondary button and inputs use `--weight-body`.
- Links in text, titles of posts and messages and member names in cards are `--color-link`, underlined (`--link-decoration`). Tool links, sender names in the table and footer-strip links inside boxes use `.ds-link--quiet` (`--color-link-quiet`, `--link-decoration-quiet`, no underline until hovered). Hover turns any link `--color-link-hover`.
- Surfaces: `--color-page` outside the sheet, `--color-canvas` the sheet. `--color-surface` is the main cell, the head box, the page header and the member box; `--color-surface-alt` the label cells, status, date and size cells, the author cell, the figures line and the toolbar; `--color-surface-strong`, always with `--color-heading` text, the group rows, quiet title strips, closing strips, idle tabs, form action row and current page number. `--fill-panel` shows in the 1px gaps.
- Bars: `--fill-bar` with `--color-bar-text` for the navigation, box titles, the current tab and the dialog title. `--fill-bar-alt` with `--color-bar-alt-text` for the table head row, the date strip of a message and quiet badges. `--fill-inverse` with `--color-inverse-text` for the footer band. The current navigation item is a `--color-surface` cell with `--color-heading` text.
- Borders: blocks have a 1px `--color-border` outline; `--color-border-muted` is for rules inside a box and under the breadcrumb; `--color-border-strong` is the sheet's top rule, the left edge of the head box and page header, the rule under the tabs, the dialog outline and button outlines. Inputs use `--color-input-border`.
- Markers: `--fill-accent` with `--color-accent-text` for the badge and the meter bar; `--color-accent` for the unread icon; `--color-accent-alt` for a bold remark (`.ds-new`). Status badges are a `color-mix()` of 28% of the status colour over `--color-surface`, with `--color-text` and a 1px solid outline in the status colour. The destructive button is `--color-danger` text on `--color-danger-surface` with a `--color-danger` outline.
- `--color-fill-1` to `--color-fill-4` colour the member cards of `.ds-grid`; `--color-fill-3` is also the alternate list row.
- One-off values outside the token set: the status badge tints above; the notice edge is a `color-mix()` of 45% `--color-accent` in `--color-notice`; the fact label column is 36%; the footer separators are at 0.6 opacity.
- The stack names Trebuchet MS with Verdana, Arial and Helvetica as fallbacks, all system fonts.

## Components

- `ds-progress`: a quota or poll bar, a native `<progress>` with a label above (private-message quota).
- `ds-accordion`: a collapsible `<details>` entry for a forum jump or an FAQ answer.
- `.ds-page`: on `<body>`. Wrap everything in `.ds-page__frame`; `.ds-page__body` holds the views, each a `.ds-view` section (`.ds-view--home` for the profile) whose blocks go in `.ds-view__body`. `.ds-page__legal` is the small print. `.ds-columns` with two `.ds-columns__col` makes the two equal columns; `.ds-layout` with `.ds-layout__main` the menu-plus-main area. `.ds-icon` sizes an inline icon, `.ds-sprite` hides the symbol sheet, `.ds-block` puts a span on its own line.
- `.ds-nav`: the site header. `.ds-nav__head` holds `.ds-nav__brand` (the `.ds-brand` with a `.ds-nav__tagline` under it) and `.ds-nav__member`, the signed-in member's box with a small avatar, a bold `.ds-nav__who` line and the unread count. `.ds-nav__menu` is the bar of equal `.ds-nav__item` cells, each a `.ds-nav__link`; the current one takes `is-current`.
- `.ds-brand`: the site's mark and name at the left of the masthead, linking home; one per page. `.ds-brand__mark` is a square tile dressed like the primary button (`--fill-button`, `--color-button-text`, a `--color-border-strong` outline, `--radius-control`, `--shadow-control`), 1.7 times `--text-display`. `.ds-brand__name` is the name in `--font-heading` at `--text-display`, coloured `--color-heading`. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo: replace both, keep the tile.
- `.ds-breadcrumb`: a small open line with `.ds-breadcrumb__link`, `.ds-breadcrumb__sep` and a bold `.ds-breadcrumb__current`.
- `.ds-hero`: the head of a profile or of any page about one person or thing. A `.ds-avatar` tile, then `.ds-hero__title` (the name, display size), `.ds-hero__rank` (rank and badges), `.ds-hero__lead` (one line of description) and `.ds-hero__action` (the main action and one secondary button); `.ds-hero__aside` is a right-aligned column of small facts. A surface box with a strong left edge; it takes `--shadow-dialog` and `--radius-page`. Never make it a banner.
- `.ds-avatar`: a square tile holding the member's picture, here a silhouette drawn in `--color-text-muted` on `--color-surface-alt`. `.ds-avatar--small` for cards, posts and the member box.
- `.ds-page-header`: the head of an inner page, used instead of `.ds-hero`. `.ds-page-header__title`, `.ds-page-header__desc`, and `.ds-page-header__action` with one button on the right.
- `.ds-stat`: the row of summary figures, one ruled line of equal cells. Each `.ds-stat__item` sets a bold `.ds-stat__value` and its `.ds-stat__label` side by side on one line. `.ds-stat` is the row; use three to six items.
- `.ds-panel`: any titled box. `.ds-panel__title` is a primary bar; `.ds-panel__title--quiet` a pale strong strip for secondary boxes. `.ds-panel__body` (`--alt` for a small tinted row), `.ds-panel__line` paragraphs, `.ds-panel__quick` a one-line form of a field and a button. Put a `.ds-facts` list straight inside a panel for profile details.
- `.ds-facts`: label and value rows (`.ds-facts__label`, `.ds-facts__value`) for details of a member or record.
- `.ds-meter`: a thin bordered track with a `.ds-meter__bar` in the accent fill, for storage used.
- `.ds-list`: posts or messages outside a table. `.ds-list__head` (primary bar), `.ds-list__item` (`is-alt` for the alternate row) with `.ds-list__title`, `.ds-list__excerpt` and `.ds-list__meta`, closed by a `.ds-list__foot` strip of quiet links.
- `.ds-grid`: member cards in four columns (`.ds-grid--2` for two, inside a column), 1px apart in one bordered block. `.ds-grid__title`, `.ds-grid__cell` (`--2`, `--3`, `--4` for the other fills) with a small avatar, `.ds-grid__name` and `.ds-grid__meta`, and a `.ds-grid__foot` strip.
- `.ds-tabs`: folder tabs on a strong rule. `.ds-tabs__tab` is a pale strong cell with heading text; `is-current` takes the primary bar. `.ds-tabs__count` is the count in normal weight.
- `.ds-table`: a message folder or any data table. `th` is the secondary bar (`.ds-table__head--num` right-aligned); `.ds-table__group` a group row; cells `.ds-table__check`, `.ds-table__status` (`is-read`), `.ds-table__from`, `.ds-table__date`, `.ds-table__num`; `.ds-table__title` is the subject link (`is-unread` bold), `.ds-table__meta` a muted second line.
- `.ds-toolbar`: the tinted row under a table or message: `.ds-buttons` on the left, `.ds-pagination` or links on the right.
- `.ds-post`: one message. `.ds-post__head` is the secondary strip with date and number; `.ds-post__author` the tinted cell with `.ds-post__name` and `.ds-post__meta` lines; `.ds-post__body` holds `.ds-post__text`, a `.ds-post__quote` with its `.ds-post__cite`, and the `.ds-post__sig`; `.ds-post__foot` is the strip of quiet tool links.
- `.ds-prose`: long text such as help: h1 (with a faint rule), h2, h3, p, ul, strong, code.
- `.ds-link`: every link in text. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--quiet` for tool and meta links; `.ds-link--strong` for bold.
- `.ds-button`: the primary action. `.ds-button--secondary` is the plain grey button; `.ds-button--danger` the destructive action, once per confirmation; `is-hover`, `is-active`, `is-focus`, `is-disabled` (or `disabled`). Group buttons in `.ds-buttons`.
- `.ds-form`: a bordered block of rows. `.ds-form__title`, `.ds-form__row` with `.ds-form__label` (plus `.ds-form__hint`) and `.ds-form__field`; controls take `.ds-form__input` (`--wide`, `--auto`, `--grow` inside a flex row, `is-invalid`, `is-disabled`), errors `.ds-form__error`, checkboxes `.ds-form__check` and `.ds-form__checkbox`, buttons in `.ds-form__actions`.
- `.ds-badge`: a small square label in the accent fill. `.ds-badge--quiet` is the secondary bar. The status variants `.ds-badge--success` (online, answered), `.ds-badge--warning` (away, receipt pending) and `.ds-badge--danger` (banned, moderator warning) are a pale tint with a 1px outline. `.ds-new` is a bold remark in the second accent.
- `.ds-sidebar`: the menu column of an inner page. Each `.ds-sidebar__block` has a `.ds-sidebar__title` bar and either a `.ds-sidebar__list` of `.ds-sidebar__item` rows (`is-current`; a `.ds-sidebar__count` on the right) or a `.ds-sidebar__body` for a meter or a small form.
- `.ds-notice`: a pale line for information; `.ds-notice--error` is the error box with a `.ds-notice__title`.
- `.ds-pagination`: `.ds-pagination__label`, then small bordered `.ds-pagination__link` boxes; `is-current` for the current page.
- `.ds-dialog`: a confirmation box in the page flow, strong outline, `.ds-dialog__title` bar, `.ds-dialog__body` with `.ds-dialog__text` and right-aligned `.ds-dialog__actions`. `.ds-dialog__backdrop` is a flat band in the overlay colour. It never floats over the page.
- `.ds-empty`: a bordered block with a `.ds-empty__title` strip, and in `.ds-empty__body` one centred `.ds-empty__text` and one button.
- `.ds-footer`: the inverse band closing the sheet: a `.ds-footer__row` of `.ds-footer__link` items with `.ds-footer__sep` between them on the left, the time line on the right.

## Never

- `border-radius <= 0px`: every box is square in all references.
- `box-shadow-blur <= 0px`: the only shadow is the hard inner line of a button; nothing is soft.
- `text-shadow = none`: no text shadow was measured in the period references.
- `transition = none`: states switch at once.
- `animation = none`: nothing moves.
- `font-size <= 18px`: the wordmark and the member's name are the largest text.
- `font-size >= 11px`: 11px is the smallest size of this type set.
- `font-weight <= 700`: bold is the heaviest weight.
- `font-families <= 2`: Trebuchet MS, plus a monospace for code.
- `border-width <= 3px`: 1px everywhere, 3px on the emphasised edges.
- `letter-spacing <= 0px`: no tracking.
- `uppercase-text <= 0%`: nothing is transformed to uppercase.
- `underlined-links >= 50%`: links in text and titles are underlined; 88% to 94% in the references this follows.
- `gradient-fills <= 15%`: gradients are on bars and the primary button only.
- `row-gap <= 1px`: rows touch or are separated by a 1px line.
- `block-gap <= 20px`: neighbouring blocks are 11px apart.
- `content-width >= 90%`: the sheet fills the viewport.

## Extending

Derive a new component from the nearest one in the specimen: a new profile box from `.ds-panel` with `.ds-facts`, a new folder view from `.ds-table`, a new menu block from `.ds-sidebar__block`, a new card set from `.ds-grid`. Build it as a 1px `--color-border` outline on `--fill-panel` with `--space-1` of padding and gap, cover it with cells in the three surface colours, and give it a title in `--fill-bar` or a strip in `--color-surface-strong` with `--color-heading` text. Keep text on the fills it already sits on in the specimen. Use tokens only: no raw colours or lengths, no new radius, shadow or transition, no font size outside the type tokens, and `--space-5` as the gap to the next block. On a profile, new boxes go into the shorter of the two columns; on inner pages, new navigation goes into the menu column.
