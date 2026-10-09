# Dark developer tools changelog and blog

## Summary

This is the changelog and engineering blog of a developer product from 2019 to 2026: a dark page with a plain bar, a mono ticker of the current versions under it, and one centred column that a dated rule runs down. Releases hang on that rule as entries with a version tag, featured posts sit in bordered cards under tinted line drawings, and a release page shows its configuration change as a diff with tinted lines. In this set the page is a deep navy, type is a tight bold grotesque, tags are squared and the one accent is a cyan that also glows behind the newest release.

## Layout

- Designed for a 1440px viewport, not fluid beyond it. Content sits in `.ds-wrap`, at most `--size-page` (1120px) wide with `--space-5` (24px) of side padding. Running text stays under `--size-narrow` (700px).
- Navigation is `.ds-nav`, a bar `--size-top` (60px) tall on `--fill-bar` that scrolls with the page: brand at the left, four links whose current one has a `--border-width-strong` line in `--color-accent` along the bottom of the bar, a secondary and a primary small button at the right. Directly under it `.ds-ticker`, a mono strip on `--fill-bar-alt` between two hairlines, names the release lines. There is no left rail.
- The home view stacks: `.ds-hero` (label, title, lead and an inline subscribe field at the left; `.ds-release`, the newest version as a card, at the right, over the glow); a `.ds-colhead` and `.ds-grid`, one tall lead card beside four small ones; `.ds-statband`, a full-width band on `--color-surface-alt` holding `.ds-stat`; then `.ds-cols`, the timeline `.ds-list` beside a `--size-side` (296px) column with a panel, the inverted `.ds-callout` and a code block.
- Inner pages have no hero. They open with `.ds-page-header`: breadcrumb, a `--text-h1` title, an optional row of labels (`.ds-page-header__meta`) and one muted line at the left, at most two buttons at the right, over a hairline. Below it `.ds-cols` (main and side), `.ds-cols--even` (halves) or `.ds-cols--menu` (a 213px column of sidebars at the left). Tabs are a row of chips in a `.ds-toolbar` directly over the table they filter; the pagination is in the `.ds-table__foot`.
- Views. The specimen is the changelog of a query engine in four views. `latest` is the home view described above. `release` is one release: prose, the diff, the table of packages and the pager to the neighbouring releases in the main column; facts, the headings of the page, contributors with a row of links, an upgrade command and the inverted callout at the right. `archive` lists every release: sidebars by line and year and the label key at the left, chips, the table with pagination and an empty state at the right. `subscribe` has the e-mail form under its notices, beside the feed addresses, the state of the subscription with its buttons and the open dialog that unsubscribes.
- Spacing scale: `--space-1` 4px (inside chips and key caps, between tally cells), `--space-2` 8px (button vertical padding, between adjacent buttons and badges, sidebar rows), `--space-3` 12px (panel head and table cell vertical padding), `--space-4` 16px (panel padding, between grid cards, between form rows), `--space-5` 24px (between stacked blocks, side padding, card body), `--space-6` 36px (between timeline entries, below a page header), `--space-7` 56px (between the sections of a view, between the main and the side column), `--space-8` 80px (hero top, above the footer).
- Other sizes: `--size-rail` 124px (the date column of the timeline), `--size-cover` 128px (the drawing of a small card), `--size-glow` 460px, `--size-mark` 26px, `--size-dot` 8px, `--size-dialog` 440px.
- Below 980px the side column goes under, the lead card spans the row and figures stand two across. Below 640px everything is one column, the links of the bar move to a second row that scrolls sideways, the date of an entry sits over it on the same rule, and a wide table scrolls inside `.ds-table__wrap`; the page itself never scrolls sideways.
- On a phone (640px and below) `.ds-nav__inner` wraps: the brand and `.ds-nav__actions` stay on the first row, `.ds-nav__links` take a second row that scrolls sideways and spread across it, and the secondary button of the actions is not shown. The `.ds-menu` stays and lists every view, the current one with the accent line.

## Typography and colour roles

The conventions for which text sits on which fill are those of the era's first layout and are kept here.

- `--color-text`, `--color-heading` and `--color-text-muted` sit on the page, on `--fill-panel`, `--color-surface-alt` (the stat band, a hovered row, the dialog) and `--color-surface-strong` (the current chip, key caps, quiet badges, empty tally cells). `--color-bar-text` sits on `--fill-bar` (the top bar). `--color-bar-alt-text` sits on `--fill-bar-alt` (the ticker, the head of the release card, code heads, table heads, the footer), with `--color-heading` for a link or the brand on that fill. `--color-input-text` sits on `--fill-input` (form controls and every code block). `--color-inverse-text` sits on `--fill-inverse`, used once per view, for `.ds-callout`.
- `--color-fill-1` to `--color-fill-4` never carry text. They tint the cover of a card (a gradient from 30% to 6% of the fill over transparent) and colour its drawing (70% of the fill mixed into `--color-heading`); `--color-fill-1` is the hero glow.
- `--color-accent` marks the current link of the bar and fills the default badge through `--fill-accent`. `--color-accent-alt` is a line and text colour: the mono label of the hero, the node of each timeline entry, the dot of the current sidebar row, the "changed" cells of a tally, keywords in code.
- `--color-success`, `--color-warning` and `--color-danger` are a dot, text, the lit cells of a tally, or a 12% tint with a 32% hairline and the colour as text. Added and removed lines of a diff are a 13% tint with text mixed 55% into `--color-input-text`. Nothing else is filled with them; the destructive button is a tint too.
- Borders: `--color-border` on boxes and bars, `--color-border-strong` on the release card, secondary buttons, the timeline rule, the current chip and the dialog, `--color-border-muted` between rows inside a box. All `--border-width`; `--border-width-strong` is the line under the current link and the focus offset.
- Type. One family for body, headings and controls, `--font-mono` for versions, dates, counts, the ticker and code. `--text-base` 15px at `--line-body` 1.5; `--text-ui` 14px for navigation, buttons, forms, tables and card text; `--text-small` 12px for badges, mono labels and dates; `--text-large` 18px for the hero lead and timeline titles. `--text-display` 64px is the hero title only; `--text-h1` 38px page titles, figures and the version on the release card; `--text-h2` 26px column heads and the lead card; `--text-h3` 16px panel and card titles. Headings are `--weight-heading` 700 with negative tracking.
- Uppercase appears only on mono labels (`.ds-eyebrow`, table heads, sidebar titles), tracked by 6% of the font size. Links are `--color-link` without underline and gain one on hover; quiet links are `--color-link-quiet` and turn `--color-heading`.
- Surface. `--radius-control` 8px for buttons, inputs and the brand tile, `--radius-panel` 12px for panels, code, notices, the callout and the pager, `--radius-page` 18px for the release card, grid cards and the dialog box, `--radius-pill` 4px for badges and chips. `--shadow-panel` is on panels and grid cards, `--shadow-control` (a glow in the button colour) on buttons and the brand tile, `--shadow-dialog` on the release card and the dialog.
- One-off values written with `calc()` and `color-mix()`: the glow (24% of a fill), cover tints (30% and 6%), drawing colours (70%), status tints (12% and 32%), diff lines (13% and 55%), the notice border (22% and 40%), key cap corners (half of `--radius-control`), tally cell corners (half of `--radius-pill`), mono label tracking, the left column of `.ds-cols--menu` (72% of `--size-side`). Navigation links are at `opacity` 0.62 until current, callout text at 0.78. The empty state's outline is dashed.

## Components

- `.ds-page`: on `<body>`. Sets the font, text colour and `--fill-page`. `.ds-nav`, the views and `.ds-footer` are its direct children.
- `.ds-wrap`, `.ds-section`, `.ds-stack`, `.ds-cols`, `.ds-colhead`: the centred column, a section 56px below the last, a vertical stack 24px apart, the column pairs (`--even`, `--menu`) and a section title with a link at its right (`.ds-colhead__title`).
- `.ds-brand`: the site's mark and name. `.ds-brand__mark` is a 26px tile dressed like the primary button holding the mark as inline SVG; `.ds-brand__name` is the name. Both are placeholders for the installing project's own name and logo.
- `.ds-nav`: the top bar. `.ds-nav__inner` holds the brand, `.ds-nav__links` of `.ds-nav__link` (current: `is-current`) and `.ds-nav__actions`; `.ds-ticker` with `.ds-ticker__inner`, `.ds-ticker__item` (`--end`) and `.ds-ticker__link` follows inside the same header. In the specimen the showing view marks its link with one `:has()` rule per view; a project sets `is-current`.
- `.ds-hero`: home view only. `.ds-hero__inner` is two columns: `.ds-eyebrow`, `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__form` (one input and the main action) and `.ds-hero__note`; then `.ds-release`. The glow is its `::before`.
- `.ds-release`: the newest version as a card: `.ds-release__head`, `.ds-release__body` with `.ds-release__version`, `.ds-release__title` and `.ds-tally`, then `.ds-release__foot`.
- `.ds-tally`: counts as rows of small cells: `.ds-tally__row` (`--add`, `--fix`) holds `.ds-tally__name`, `.ds-tally__cells` of `.ds-tally__cell` (`is-on`) and `.ds-tally__count`.
- `.ds-page-header`: head of an inner page: `.ds-page-header__inner` with `.ds-page-header__copy` (breadcrumb, `.ds-page-header__title`, `.ds-page-header__meta`, `.ds-page-header__text`) and `.ds-page-header__actions`.
- `.ds-prose`: running text: h1, h2, h3, p, lists, `code`, `strong`; `.ds-prose__lead` for the first paragraph.
- `.ds-code`: a code block on `--fill-input`: optional `.ds-code__head` with `.ds-code__name`, then `.ds-code__body` (a `pre`). Spans `.ds-code__k`, `__s`, `__c`, `__n` colour keywords, strings, comments and names. For a diff every line is a `.ds-diff__line` (`--add`, `--del`, `--hunk`).
- `.ds-eyebrow`, `.ds-mono`, `.ds-kbd`, `.ds-dot`, `.ds-avatar`: the mono label (`--accent`), mono data text, a key cap, a status dot (`--success`, `--warning`, `--danger`, `--accent`, `--ring`) and an initials disc; `.ds-avatars` rows them.
- `.ds-link`: text link with `is-visited`, `is-hover`, `is-active`, `is-focus`; `--quiet`, `--more` (appends an arrow). `.ds-linkrow` lays links in a row.
- `.ds-button`: 14px medium text, 8px corners, a glow. `--secondary` (translucent fill, strong hairline), `--danger` (tint, hairline and text in the danger colour), `--inverse` (on the callout only), `--small`, `--large`. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttonrow` spaces them 8px apart.
- `.ds-form`: labels above fields. `.ds-form__row`, `.ds-form__field`, `.ds-form__label`, `.ds-form__input` (`--mono`), `.ds-form__select`, `.ds-form__textarea` (`--mono`), `is-error` plus `.ds-form__error`, `.ds-form__hint`, `.ds-form__check` with `.ds-form__checkbox`, `.ds-form__actions`.
- `.ds-table`: inside a panel, edge to edge: a mono uppercase head on `--fill-bar-alt`, hairline rows, `is-hover` rows on `--color-surface-alt`, `.ds-table__num` for right-aligned mono numbers. Wrap in `.ds-table__wrap`; `.ds-table__foot` closes the panel and holds the pagination.
- `.ds-list`: the timeline. `.ds-list__item` is `.ds-list__meta` (mono date over a version badge, right-aligned) beside `.ds-list__body`, which draws the rule and the node and holds `.ds-list__title`, `.ds-list__text` and `.ds-list__tags`.
- `.ds-panel`: hairline box on `--fill-panel`: `.ds-panel__head` with `.ds-panel__title`, `.ds-panel__body`; `.ds-panel__rows` of `.ds-panel__row` (`.ds-panel__key`) for facts.
- `.ds-stat`: the row of figures inside `.ds-statband`. Each `.ds-stat__item` is a `.ds-stat__figure` beside a two-line `.ds-stat__label`, parted by vertical hairlines.
- `.ds-grid`: featured posts. `.ds-grid__cell` (`--lead` for the tall one) is a link holding `.ds-grid__cover` (`--2`, `--3`, `--4` for the other fills) with one inline SVG drawn in strokes, and `.ds-grid__body` with `.ds-grid__kicker`, `.ds-grid__title` and `.ds-grid__text`.
- `.ds-callout`: the inverted block: `.ds-callout__title`, `.ds-callout__text` and one `.ds-button--inverse`.
- `.ds-tabs`: a row of chips, `.ds-tabs__tab`; the current one (`is-current`) is on `--color-surface-strong`. `.ds-toolbar` puts it in a row with a `.ds-toolbar__note`.
- `.ds-badge`: small squared label on `--fill-accent`. `--quiet`, `--outline`, `--new` (accent-alt tint), `--mono`. Status variants: `--success`, `--warning`, `--danger`. `.ds-badgerow` for several.
- `.ds-sidebar`: a mono `.ds-sidebar__title` over `.ds-sidebar__list`; each `.ds-sidebar__item` is a row with a count at the right, the current one (`is-current`) bright behind a dot.
- `.ds-pager`: two boxes, `.ds-pager__link` (`--next`) with a mono `.ds-pager__label`, to the older and the newer release.
- `.ds-notice`: message strip on `--color-notice`; `--error` in the danger colours. `.ds-notice__title`; `.ds-noticerow` stacks notices.
- `.ds-pagination`: `.ds-pagination__link`; `is-current` on `--color-surface-strong`, `is-disabled` grey.
- `.ds-breadcrumb`: `.ds-breadcrumb__item` parted by slashes; the last is `.ds-breadcrumb__current`.
- `.ds-dialog`: a block in `--color-overlay` centring `.ds-dialog__box` with `--shadow-dialog`: `.ds-dialog__head` (`.ds-dialog__title`, `.ds-dialog__close`), `.ds-dialog__body`, `.ds-dialog__actions`.
- `.ds-empty`: a dashed box: `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one small button.
- `.ds-footer`: on `--fill-bar-alt`: `.ds-footer__inner` with the brand, `.ds-footer__links` of `.ds-footer__link` and one small button; `.ds-footer__line` is the mono last line.
- `.ds-menu`: a drop-down behind a button in the bar (`.ds-menu__button`, `.ds-menu__list`); it lists every view, so on a phone it is the full navigation.
- `.ds-menu__item`: one link of the drop-down; the current view carries the accent line and the stronger fill.
- `.ds-menu__item--view`: an item of the menu that names a view of the site; the current view's item is marked by one `:has()` rule per view, like the bar links, so the menu is a second way to every view on a phone.
- `.ds-menu__title`: a small mono label over a group of links in the drop-down.
- `.ds-tooltip`: a small question mark beside a title; `.ds-tooltip__tip` opens on hover or focus, never by default.
- `.ds-switch`: an on or off setting that takes effect at once, drawn over a checkbox (`.ds-switch__input`, `.ds-switch__track`, `.ds-switch__label`).
- `.ds-progress`: a thin bar for a build, a rollout or a quota, with a mono line above (`.ds-progress__head`, `.ds-progress__value`, `.ds-progress__bar`).

## Never

- `border-width <= 2px`: rules and outlines are 1px hairlines; 2px is the line under the current link.
- `border-radius <= 18px`: 8px on controls, 12px on panels, 18px on the largest boxes.
- `box-shadow-blur <= 70px`: only the release card and the dialog cast a wide shadow.
- `text-shadow = none`: text is flat on every fill.
- `font-size >= 12px`: badges, mono labels and dates are the smallest text.
- `font-size <= 64px`: the hero title is the largest text.
- `font-weight <= 700`: headings are bold; nothing is heavier.
- `font-families <= 2`: one grotesque and one monospace.
- `line-height <= 1.5`: body text is set at 1.5.
- `uppercase-text <= 5%`: only mono labels are uppercase.
- `underlined-links <= 5%`: links are told apart by colour and underline on hover only.
- `letter-spacing <= 3px`: the hero title is pulled together by 2.4px; nothing is spaced wider than that.
- `animation = none`: nothing moves by itself.
- `block-gap <= 80px`: sections are 56px apart, the footer 80px below the last.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new block is a `.ds-panel` or a `.ds-grid__cell` with a 1px border in `--color-border`, placed in `.ds-wrap`; anything dated goes on the timeline as a `.ds-list__item` with its version as a mono badge; a change to a file is shown as a `.ds-code` of `.ds-diff__line` rows. Keep to the conventions above for which text sits on which fill: never put text on a full-strength `--color-fill-*` or status colour, tint it instead. Use one glow and one inverted block per view. Spacing comes from the `--space-*` scale, and any one-off shade is a `color-mix()` of tokens noted next to the rule.
