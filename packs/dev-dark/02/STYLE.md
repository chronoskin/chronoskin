# Dark developer tools issue tracker

## Summary

This is the signed-in side of a developer tool from 2019 to 2026: an issue tracker on a blue-black page, with a rail of icon links down the left edge, a top bar whose search field is the entry to a command palette, and a work area of hairline boxes. Rows are dense and led by a status dot; ids, commits, counts and dates are set in monospace, and here the headings and controls are monospace too. The one accent colour marks the primary action and the progress of the running cycle, over a single faint glow in the first card.

## Layout

- Designed for a 1440px viewport, not fluid beyond it. `.ds-page` is a grid of two columns: the rail (`.ds-nav`, `--size-rail` 232px) for the whole height, and beside it `.ds-topbar` (`--size-top` 48px, sticky), the showing view and the footer line. `.ds-content` is centred in the work area, at most `--size-page` (1120px) wide, padded by `--space-5` and `--space-6`, and stacks its children `--space-5` (24px) apart.
- Navigation is the rail on `--fill-bar-alt`: `.ds-nav__head` with the brand and a key cap, groups of links under mono group names, each link an icon or a team dot, a label and an optional count; the current link sits on a 14% wash of the rail's text colour. The account is at the foot. `.ds-topbar` on `--fill-bar` holds a mono location, `.ds-search` (the palette field with its `Ctrl K` key cap) in the middle and two small buttons at the right.
- The home view stacks: `.ds-hero`, one card with the state of the cycle and a meter of small squares; `.ds-stat`, four figures in one ruled box; `.ds-cols` (2 to 1) with the issue list panel beside a panel of deployments and the inverted `.ds-tip`; and `.ds-grid`, four project cards under a `.ds-colhead`.
- Inner pages have no hero. `.ds-content` opens with `.ds-page-header`: the breadcrumb, a `--text-h2` title and a meta line at the left, at most two buttons at the right, and a rule below. Then the work: `.ds-cols--side` puts a `--size-side` (300px) column of properties at the right, `.ds-cols--even` halves, `.ds-cols--menu` puts a `.ds-subnav` of sections at the left of a settings page. Tabs are underlined and sit on the head of the panel that holds the table; the pagination is in its `.ds-table__foot`.
- Views. The specimen has four. `issues` is the home view described above. `issue` is one issue: its description as prose, the activity list and the open delete dialog in the main column; the properties sidebar, the action buttons and linked items at the right. `deploys` is the table of deployments with tabs and pagination, then an empty state beside the environments. `settings` is the workspace form under its notices, beside the section menu, with the label key and the danger zone below.
- Spacing scale: `--space-1` 4px (rail link and small button vertical padding, meter gaps), `--space-2` 8px (list row and button vertical padding, gaps between buttons and badges), `--space-3` 12px (panel head and table cell vertical padding), `--space-4` 16px (panel, stat and card padding), `--space-5` 24px (between blocks, hero padding), `--space-6` 32px (work area side padding, around a dialog), `--space-7` 48px (below the last block).
- Other sizes: `--size-search` 420px, `--size-id` 68px (the id column of an issue row), `--size-icon` 16px, `--size-mark` 22px, `--size-dot` 8px, `--size-glow` 420px, `--size-dialog` 440px.
- Below 1000px the rail becomes a block across the top with its links in one row that scrolls sideways, cards stand two across and column pairs stack. Below 640px everything is one column, an issue row wraps its labels under its title, and a wide table scrolls inside `.ds-table__wrap`.
- On a phone the rail is a bar across the top from 1000px down: `.ds-nav__links--main` is one row that scrolls sideways, `.ds-nav__group`, `.ds-nav__foot`, the `.ds-kbd` and the other link lists are not shown, and `.ds-subnav` becomes a scrolling row whose current link has an `--color-accent-alt` line under it. At 640px and below each `.ds-nav__link` keeps only its `.ds-nav__label` (icon, badge and code are not shown) with tighter padding, the `.ds-menu` in `.ds-topbar__actions` keeps only its avatar (`.ds-menu__label` is not shown) and sits at the right, and a tooltip in a panel title opens under its line.

## Typography and colour roles

The conventions for which text sits on which fill are those of the era's first layout and are kept here.

- `--color-text`, `--color-heading` and `--color-text-muted` on the page, on `--fill-panel`, `--color-surface-alt` (list group rows, hovered rows, the dialog) and `--color-surface-strong` (key caps, quiet badges, avatars, empty meter cells). `--color-bar-text` on `--fill-bar` (the top bar). `--color-bar-alt-text` on `--fill-bar-alt` (the rail, table heads, code heads, the footer line), with `--color-heading` for the current link. `--color-inverse-text` on `--fill-inverse`, used once, for `.ds-tip`.
- `--color-fill-1` to `--color-fill-4` never carry text: they tint the top of a project card at 12%, colour its progress bar at full strength, and `--color-fill-1` is the hero glow and the done cells of the meter. Blocked cells use `--color-warning`.
- `--color-accent-alt` is a line and text colour: the mono label of the hero, the underline of the current tab, the marker of the current menu entry, a team dot.
- Status is a dot before a word (`.ds-dot`, `.ds-table__state`) or a tinted badge; `--color-success`, `--color-warning` and `--color-danger` are never a solid fill. The destructive button is a 12% tint with a hairline and text in `--color-danger`.
- Type in this set: `--font-body` is a humanist sans for descriptions and rows; `--font-heading`, `--font-ui` and `--font-mono` are one monospace family. `--text-base` 14px at `--line-body` 1.55, `--text-ui` 12px (rail, buttons, forms, tables, rows), `--text-small` 11px (mono data, badges, table heads). The hero title and prose h1 are `--text-h1` 28px, the page title, figures and prose h2 `--text-h2` 20px, panel and card titles `--text-h3` 14px bold. `--text-display` is not used by this layout. Links in running text are underlined in this set.
- Surface in this set: almost square, `--radius-control` 2px, `--radius-panel` and `--radius-page` 4px, `--radius-pill` 2px so badges are tags. Buttons, the search field and the brand tile carry `--shadow-control`, a 2px inset lip; panels have no shadow; the dialog casts `--shadow-dialog`. `--fill-page` draws a faint grid. `--border-width-strong` 3px is the underline of the current tab and the ring of a status dot.
- One-off values written with `calc()` and `color-mix()`: the washes behind a hovered and a current rail link (8% and 14% of `--color-bar-alt-text`), the hero glow (30% of `--color-fill-1`), in-progress meter cells (40% of `--color-fill-1` in `--color-surface-strong`), card tints (12%), status tints (12% and 32%), notice borders, key cap corners (half of `--radius-control`), mono label tracking. Bar widths are percentages. The empty state's outline is dashed.

## Components

- `.ds-page`: on `<body>`. Sets the font, text colour and page fill and makes the two-column grid; `.ds-nav`, `.ds-topbar`, the views and `.ds-footer` are its direct children.
- `.ds-content`, `.ds-cols` (`--side`, `--even`, `--menu`), `.ds-stack`: the padded column of a view, the column pairs and a vertical stack. `.ds-panel--grow` fills spare height.
- `.ds-brand`: the mark on a 22px tile dressed like the primary button, and the name. Placeholders for the installing project's own name and logo.
- `.ds-nav`: the rail. `.ds-nav__head`, `.ds-nav__group`, `.ds-nav__links` (`--main` for the links that follow the view) of `.ds-nav__link` with `.ds-nav__icon`, `.ds-nav__label` and an optional badge or mono code; `.ds-nav__foot` with an avatar and `.ds-nav__who`. A project sets `is-current` on the current link.
- `.ds-topbar`: `.ds-topbar__where`, `.ds-search` (`.ds-search__icon`, `.ds-search__text`, a `.ds-kbd`) and `.ds-topbar__actions`.
- `.ds-hero`: home view only, one card: `.ds-hero__copy` with a `.ds-eyebrow`, `.ds-hero__title`, `.ds-hero__lead` and `.ds-hero__actions`; `.ds-hero__art` holds `.ds-meter`, a grid of `.ds-meter__cell` (`--done`, `--doing`, `--late`) over `.ds-meter__key`.
- `.ds-page-header`: `.ds-page-header__copy` (breadcrumb, `.ds-page-header__title`, `.ds-page-header__text`, which may start with a status badge) and `.ds-page-header__actions`.
- `.ds-prose`: h1, h2, h3, p, lists, `code`, `strong`, inside a panel; `.ds-prose__lead`.
- `.ds-code`: code block with `.ds-code__head`, `.ds-code__name`, `.ds-code__body` and the spans `__k`, `__s`, `__c`, `__n`.
- `.ds-eyebrow`, `.ds-mono`, `.ds-kbd`, `.ds-dot` (`--success`, `--warning`, `--danger`, `--accent`, `--ring`), `.ds-avatar`: the small parts every row is made of.
- `.ds-link`: `is-visited`, `is-hover`, `is-active`, `is-focus`; `--quiet`, `--more`. `.ds-linkrow`.
- `.ds-button`: `--secondary`, `--danger`, `--inverse` (on `.ds-tip` only), `--small`, `--large`; `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttonrow`.
- `.ds-form`: `.ds-form__row`, `.ds-form__field`, `.ds-form__label`, `.ds-form__input` (`--mono`), `.ds-form__select`, `.ds-form__textarea` (`--mono`), `is-error` with `.ds-form__error`, `.ds-form__hint`, `.ds-form__check`, `.ds-form__checkbox`, `.ds-form__actions`.
- `.ds-table`: in a panel under the tabs: mono uppercase head on `--fill-bar-alt`, hairline rows, `.ds-table__num`, `.ds-table__state`. `.ds-table__wrap`, `.ds-table__foot`.
- `.ds-list`: issue rows in a panel. `.ds-list__group` heads a state; `.ds-list__item` is a dot, the mono `.ds-list__id`, the `.ds-list__title` link and `.ds-list__meta` (label badge, avatar, date).
- `.ds-panel`: `.ds-panel__head`, `.ds-panel__title`, `.ds-panel__body`, `.ds-panel__rows` of `.ds-panel__row` with `.ds-panel__key`.
- `.ds-stat`: the ruled box of four `.ds-stat__item`: a mono `.ds-stat__label` led by a dot, the `.ds-stat__figure` and a `.ds-stat__delta` (`--down`).
- `.ds-grid`: four project cards. `.ds-grid__cell` (`--2`, `--3`, `--4`) is a link holding a mono key, `.ds-grid__title`, `.ds-grid__text`, `.ds-grid__bar` (`--50`, `--75`, `--90`) and `.ds-grid__foot`. `.ds-colhead` with `.ds-colhead__title` names the row.
- `.ds-tabs`: underlined `.ds-tabs__tab` row on a panel head, each with an optional mono count; `is-current`.
- `.ds-badge`: `--quiet`, `--outline`, `--new`, `--mono`; status `--success`, `--warning`, `--danger`. `.ds-badgerow`.
- `.ds-sidebar`: the properties box: `.ds-sidebar__title`, `.ds-sidebar__list` of `.ds-sidebar__item`, each a `.ds-sidebar__key` and a link or value.
- `.ds-subnav`: the sections of a settings page, `.ds-subnav__link` along a rule; `is-current`.
- `.ds-activity`: `.ds-activity__item` rows of avatar, `.ds-activity__text` and a mono time.
- `.ds-tip`: the inverted card: `.ds-tip__title`, a paragraph, one `.ds-button--inverse`.
- `.ds-notice`: `--error`, `.ds-notice__title`, `.ds-noticerow`.
- `.ds-pagination`: `.ds-pagination__link`, `is-current`, `is-disabled`.
- `.ds-breadcrumb`: `.ds-breadcrumb__item`, `.ds-breadcrumb__current`.
- `.ds-dialog`: `.ds-dialog__box`, `.ds-dialog__head`, `.ds-dialog__title`, `.ds-dialog__close`, `.ds-dialog__body`, `.ds-dialog__actions`.
- `.ds-empty`: `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one small button, in a dashed box where a panel would be.
- `.ds-footer`: one mono line on `--fill-bar-alt`: `.ds-footer__status`, the build, `.ds-footer__links` of `.ds-footer__link`.
- `.ds-menu`: a drop-down behind a button in the bar (`.ds-menu__button`, `.ds-menu__list`); it lists every view, so on a phone it is the full navigation.
- `.ds-menu__item`: one link of the drop-down; the current view carries the accent line and the stronger fill.
- `.ds-menu__item--view`: an item of the menu that names a view of the site; the current view's item is marked by one `:has()` rule per view, like the bar links, so the menu is a second way to every view on a phone.
- `.ds-menu__title`: a small mono label over a group of links in the drop-down.
- `.ds-tooltip`: a small question mark beside a title; `.ds-tooltip__tip` opens on hover or focus, never by default.
- `.ds-switch`: an on or off setting that takes effect at once, drawn over a checkbox (`.ds-switch__input`, `.ds-switch__track`, `.ds-switch__label`).
- `.ds-progress`: a thin bar for a build, a rollout or a quota, with a mono line above (`.ds-progress__head`, `.ds-progress__value`, `.ds-progress__bar`).

## Never

- `border-width <= 3px`: rules are 1px; 3px marks the current tab and rings a status dot.
- `border-radius <= 4px`: 2px on controls and tags, 4px on boxes.
- `box-shadow-blur <= 25px`: only the dialog casts a blurred shadow.
- `text-shadow = none`: text is flat on every fill.
- `font-size >= 11px`: mono data and badges are the smallest text.
- `font-size <= 28px`: the hero title is the largest text; an app screen has no display size.
- `font-weight <= 700`: titles are semibold or bold, never black.
- `font-families <= 2`: one sans-serif for running text, one monospace for everything else.
- `line-height <= 1.6`: body text is set at 1.55.
- `uppercase-text <= 5%`: only mono labels and table heads are uppercase.
- `underlined-links <= 15%`: only links in running text and panels are underlined; rows, navigation and tabs are not.
- `letter-spacing <= 1px`: only uppercase mono labels are tracked, by under 1px.
- `gradient-fills <= 3%`: gradients are the hero glow and the tint at the top of a project card.
- `animation = none`: nothing moves by itself.
- `transition = none`: states change at once.
- `block-gap <= 32px`: blocks are 24px apart.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. Anything new in the work area is a `.ds-panel` with a `.ds-panel__head`, 24px from its neighbours; a new kind of record is a `.ds-list__item` or a table row led by a `.ds-dot`; a new destination is a `.ds-nav__link` with a 16px stroke icon; anything a developer would copy is `.ds-mono` or a `.ds-kbd`. Keep to the era's conventions for which text sits on which fill, tint with the fills and the status colours instead of filling with them, and use one glow and one inverted card per view. Spacing comes from the `--space-*` scale, and any one-off shade is a `color-mix()` of tokens noted next to the rule.
