# Dark developer tools documentation

## Summary

This is the documentation site of a developer product from 2019 to 2026: a dark page in three columns, with the tree of pages at the left, the text in the middle and the headings of the page along a rule at the right. Code blocks with a title strip, mono parameter names, method badges and status dots carry most of the content, and a split page puts a reference table beside its request and response. In this set the page is a warm charcoal, headings are a serif, corners are generous and the one accent is an orange that also glows behind the opening of the first page.

## Layout

- Designed for a 1440px viewport, not fluid beyond it. `.ds-page` is a grid: `.ds-nav`, the top bar (`--size-top` 56px, sticky, on `--fill-bar`), across both columns; below it `.ds-tree` (`--size-tree` 264px, sticky, on `--fill-bar-alt`) at the left and the showing view beside it, with the footer line under the view. `.ds-content` is centred, at most `--size-page` (1100px) wide, padded by `--space-6` and `--space-7`, and stacks its children `--space-6` (32px) apart.
- Navigation is double. The top bar holds the brand with a mono `docs` tag, four section links (the current one underlined in `--color-accent-alt`), the search field with its key cap and one small button. The tree lists every page in groups under mono titles; the current page sits on a 12% wash.
- `.ds-doc` is the page body: the main column and a `--size-toc` (208px) `.ds-doc__aside` that stays in view while the text scrolls. `.ds-doc--split` is two equal columns for a reference or console page, explanation at the left and code at the right.
- The home view stacks: `.ds-hero` (label, title, lead and two large buttons at the left, a first request as a code block at the right, over the glow); `.ds-stat`, four figures between two rules; then `.ds-doc` with the `.ds-grid` of six guide cards, a panel of install commands and the `.ds-list` of updated pages, beside the "on this page" sidebar and the inverted `.ds-next` card.
- Inner pages have no hero. `.ds-content` opens with `.ds-page-header`: breadcrumb, a `--text-h1` title (with a method badge on a reference page), one line at `--text-large`, one secondary button at the right, and a rule below. Tabs are an underlined row directly over the code block they switch (`.ds-tabbed`). A guide ends with its notices and the pagination.
- Views. The specimen documents an e-mail API in four views. `overview` is the home view described above. `guide` is a how-to: tabbed code, prose, an information and an error notice, pagination, and the page's headings and related links at the right. `reference` is one endpoint: the parameter table and the label key at the left, request, response and the list of status codes at the right. `console` is a form that sends a test request, beside the empty response, the open dialog that revokes a key and the list of keys.
- Spacing scale: `--space-1` 4px (tree link vertical padding, inside key caps), `--space-2` 8px (button vertical padding, gaps between buttons and badges), `--space-3` 12px (panel head and table cell vertical padding, between tabs and code), `--space-4` 16px (panel, card and code padding, between cards), `--space-5` 24px (between stacked blocks, tree padding), `--space-6` 32px (between the sections of a view, split columns), `--space-7` 48px (content side padding, the gap to the right column).
- Other sizes: `--size-search` 280px, `--size-tile` 36px (the icon tile of a card), `--size-icon` 16px, `--size-mark` 24px, `--size-dot` 8px, `--size-glow` 380px, `--size-dialog` 440px.
- Below 1180px the right column goes under the text and split pages stack. Below 860px the tree becomes one row of links under the bar that scrolls sideways, the section links take a second row of the bar, cards stack, and a wide table scrolls inside `.ds-table__wrap`.
- On a phone (860px and below) the bar wraps: `.ds-nav__brand` and `.ds-nav__actions` stay on the first row, `.ds-nav__links` take a second row that scrolls sideways, and the tree shows only the `.ds-tree__list` that holds the current view. At 640px and below the mono badge of `.ds-nav__brand` and the `.ds-menu__label` are not shown, so the `.ds-menu` keeps only its avatar and still lists every view, the current one with the accent line; the links share the width equally and a tooltip in a panel title opens under its line.

## Typography and colour roles

The conventions for which text sits on which fill are those of the era's first layout and are kept here.

- `--color-text`, `--color-heading` and `--color-text-muted` on the page, on `--fill-panel`, `--color-surface-alt` (a hovered card, the dialog) and `--color-surface-strong` (key caps, inline code, quiet badges). `--color-bar-text` on `--fill-bar` (the top bar). `--color-bar-alt-text` on `--fill-bar-alt` (the tree, code heads, table heads, the footer), with `--color-heading` for the current page. `--color-input-text` on `--fill-input` (form controls, the search field and every code block). `--color-inverse-text` on `--fill-inverse`, used once per view, for `.ds-next`.
- `--color-fill-1` to `--color-fill-4` never carry text: they tint the icon tile of a guide card (16% fill, 35% hairline, the icon in 75% of the fill mixed into `--color-heading`), and `--color-fill-1` is the hero glow.
- `--color-accent-alt` is a line and text colour: the mono label of the hero, keywords in code, the underline of the current section and tab, the marker of the current heading in the sidebar, the first stat's rule.
- In code, `.ds-code__k` is `--color-accent-alt`, `.ds-code__s` `--color-success`, `.ds-code__n` `--color-link`, `.ds-code__c` `--color-text-muted`. Status colours are a dot, text or a 12% tint, never a solid fill.
- Type in this set: a grotesque for body and controls, a serif for headings at regular weight, a monospace for data. `--text-base` 16px at `--line-body` 1.65; `--text-ui` 14px (navigation, tree, buttons, forms, tables, code, card text); `--text-small` 13px (badges, mono labels, dates, footer); `--text-large` 18px (hero lead, page header line, prose lead). The hero title is 78% of `--text-display` (47px), because the text column is narrower than a landing page; `--text-h1` 40px page titles and prose h1; `--text-h2` 26px column heads, figures and prose h2; `--text-h3` 17px panel and card titles in `--weight-bold`.
- Surface in this set: `--radius-control` 10px, `--radius-panel` 14px, `--radius-page` 20px (guide cards, the hero code block, the dialog), pills for badges. Panels and cards lie on `--fill-panel`, a faint top-lit gradient, under `--shadow-panel`; buttons are a short gradient with a lit top edge and glow in their own colour when hovered; the hero code block and the dialog cast `--shadow-dialog`. `--fill-page` adds a wide, very faint glow at the top of the page.
- One-off values written with `calc()` and `color-mix()`: the hero title size, the hero glow (28% of `--color-fill-1`), tile tints, the wash behind the current tree link (12% of `--color-bar-alt-text`), status tints (12% and 32%), notice borders, key cap and inline code corners (half of `--radius-control`), mono label tracking. Section links are at `opacity` 0.66 until current. The empty state's outline is dashed.

## Components

- `.ds-page`: on `<body>`. Sets the font, text colour and page fill and makes the grid; `.ds-nav`, `.ds-tree`, the views and `.ds-footer` are its direct children.
- `.ds-content`, `.ds-doc` (`--split`), `.ds-doc__aside`, `.ds-stack`: the padded column of a view, the page body, its right column and a vertical stack. `.ds-panel--grow` fills spare height in a split page.
- `.ds-brand`: the mark on a 24px tile dressed like the primary button, and the name. Placeholders for the installing project's own name and logo.
- `.ds-nav`: the top bar: `.ds-nav__brand` (brand and a mono tag), `.ds-nav__links` of `.ds-nav__link` (`is-current`), `.ds-nav__actions` with `.ds-search` (`.ds-search__icon`, `.ds-search__text`, a `.ds-kbd`) and one small button.
- `.ds-tree`: the page tree: `.ds-tree__title` over `.ds-tree__list` of `.ds-tree__link` (`is-current`; `--view` marks the links the specimen highlights by view). A link may end in a badge or a mono method.
- `.ds-hero`: home view only: `.ds-hero__copy` with a `.ds-eyebrow`, `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__actions`; `.ds-hero__art` holds one `.ds-code`.
- `.ds-page-header`: `.ds-page-header__copy` (breadcrumb, `.ds-page-header__title`, `.ds-page-header__text`) and `.ds-page-header__actions`.
- `.ds-prose`: h1, h2, h3, p, lists, inline `code`, `strong`; `.ds-prose__lead`. It stands directly on the page, not in a panel.
- `.ds-code`: the code block: `.ds-code__head` with an optional dot, `.ds-code__name` and a note or key cap; `.ds-code__body` (a `pre`) with the spans `__k`, `__s`, `__c`, `__n`.
- `.ds-eyebrow`, `.ds-mono`, `.ds-kbd`, `.ds-dot` (`--success`, `--warning`, `--danger`, `--accent`, `--ring`), `.ds-avatar`: the small parts.
- `.ds-link`: `is-visited`, `is-hover`, `is-active`, `is-focus`; `--quiet`, `--more`. `.ds-linkrow`.
- `.ds-button`: `--secondary`, `--danger`, `--inverse` (on `.ds-next` only), `--small`, `--large`; `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttonrow`.
- `.ds-form`: `.ds-form__row`, `.ds-form__field`, `.ds-form__label`, `.ds-form__input` (`--mono` for keys and addresses), `.ds-form__select`, `.ds-form__textarea` (`--mono`), `is-error` with `.ds-form__error`, `.ds-form__hint`, `.ds-form__check`, `.ds-form__checkbox`, `.ds-form__actions`.
- `.ds-table`: the reference table in a panel: mono uppercase head, hairline rows, `.ds-table__name` for a mono parameter name, `.ds-table__num`, `.ds-table__state`. `.ds-table__wrap`, `.ds-table__foot`.
- `.ds-list`: updated pages between rules: `.ds-list__item` is a `.ds-list__body` (`.ds-list__title`, `.ds-list__text`) and a mono `.ds-list__meta` date.
- `.ds-panel`: `.ds-panel__head`, `.ds-panel__title`, `.ds-panel__body`, `.ds-panel__rows` of `.ds-panel__row` with `.ds-panel__key`.
- `.ds-stat`: four `.ds-stat__item` between two rules, each a `.ds-stat__figure` over a mono `.ds-stat__label` behind a 2px left rule.
- `.ds-grid`: guide cards, three across. `.ds-grid__cell` is a link: `.ds-grid__tile` (`--2`, `--3`, `--4`) with a 16px stroke icon, `.ds-grid__title`, `.ds-grid__text`. `.ds-colhead` with `.ds-colhead__title` names a section of the home view.
- `.ds-tabs`: underlined `.ds-tabs__tab` row; `is-current`. `.ds-tabbed` joins it to the code block below.
- `.ds-badge`: `--quiet`, `--outline`, `--new`, `--mono` (methods); status `--success`, `--warning`, `--danger`. `.ds-badgerow`.
- `.ds-sidebar`: "on this page": a mono `.ds-sidebar__title` over `.ds-sidebar__list` along a rule; `.ds-sidebar__item` (`--sub` indented), current with `is-current`.
- `.ds-next`: the inverted card: `.ds-next__title`, `.ds-next__text`, one `.ds-button--inverse`.
- `.ds-notice`: `--error`, `.ds-notice__title`, `.ds-noticerow`.
- `.ds-pagination`: previous, the numbered steps of a guide, next: `.ds-pagination__link`, `is-current`, `is-disabled`.
- `.ds-breadcrumb`: `.ds-breadcrumb__item`, `.ds-breadcrumb__current`.
- `.ds-dialog`: `.ds-dialog__box`, `.ds-dialog__head`, `.ds-dialog__title`, `.ds-dialog__close`, `.ds-dialog__body`, `.ds-dialog__actions`.
- `.ds-empty`: `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one small button, in a dashed box.
- `.ds-footer`: one line on `--fill-bar-alt` under the view: the copyright and `.ds-footer__links` of `.ds-footer__link`.
- `.ds-menu`: a drop-down behind a button in the bar (`.ds-menu__button`, `.ds-menu__list`); it lists every view, so on a phone it is the full navigation.
- `.ds-menu__item`: one link of the drop-down; the current view carries the accent line and the stronger fill.
- `.ds-menu__item--view`: an item of the menu that names a view of the site; the current view's item is marked by one `:has()` rule per view, like the bar links, so the menu is a second way to every view on a phone.
- `.ds-menu__title`: a small mono label over a group of links in the drop-down.
- `.ds-avatar`: initials in a circle for a person, in the account menu, quotes and lists.
- `.ds-tooltip`: a small question mark beside a title; `.ds-tooltip__tip` opens on hover or focus, never by default.
- `.ds-accordion`: questions as `details` items (`.ds-accordion__item`, `.ds-accordion__summary`, `.ds-accordion__body`); the first is open.
- `.ds-switch`: an on or off setting that takes effect at once, drawn over a checkbox (`.ds-switch__input`, `.ds-switch__track`, `.ds-switch__label`).
- `.ds-progress`: a thin bar for a build, a rollout or a quota, with a mono line above (`.ds-progress__head`, `.ds-progress__value`, `.ds-progress__bar`).

## Never

- `border-width <= 2px`: rules are 1px; 2px marks the current section and the stat rules.
- `border-radius <= 20px`: 10px on controls, 14px on panels, 20px on cards and the dialog; pills are shapes.
- `box-shadow-blur <= 50px`: the dialog and the hero code block cast the widest shadow.
- `text-shadow = none`: text is flat on every fill.
- `font-size >= 13px`: badges, mono labels and dates are the smallest text.
- `font-size <= 47px`: the hero title is the largest text.
- `font-weight <= 700`: serif headings are regular, small titles bold; nothing is black.
- `font-families <= 3`: a sans-serif, a serif for headings and a monospace.
- `line-height <= 1.7`: body text is set at 1.65.
- `uppercase-text <= 5%`: only mono labels and table heads are uppercase.
- `underlined-links <= 5%`: links are told apart by colour and underline on hover only.
- `letter-spacing <= 1px`: only uppercase mono labels are tracked, by under 1px.
- `gradient-fills <= 8%`: gradients are the glow, the faint light on panels and the primary button; no box is filled with a strong gradient.
- `animation = none`: nothing moves by itself.
- `block-gap <= 48px`: sections of a view are 32px apart, blocks inside them 24px.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new page is a link in `.ds-tree` and a view that opens with `.ds-page-header`; text goes straight on the page as `.ds-prose`, anything boxed is a `.ds-panel`, every example is a `.ds-code` with a head that names it, and names of parameters, methods, keys and codes are set in `--font-mono`. Keep to the era's conventions for which text sits on which fill, tint with the fills and the status colours instead of filling with them, and use one glow and one inverted card per view. Spacing comes from the `--space-*` scale, and any one-off shade is a `color-mix()` of tokens noted next to the rule.
