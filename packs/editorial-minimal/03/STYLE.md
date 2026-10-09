# Editorial minimalism writing platform

## Summary

Editorial minimalism is the reading-first page of 2012 to 2018: a narrow column of carefully set text on a plain page, hairline rules, a single accent colour and almost no boxes or decoration. This layout is its application form, the writing platform and notes tool of those years: one low top bar, a feed of story cards that have no box around them, a reading view and an editor that both keep to the same measure. Chrome is kept to the least that works, so the look is carried by the type sizes, the line height and the white space between stories.

## Layout

- The app is a fixed centred column of `--size-page` (1040px), designed for a 1280px viewport. Reading and writing views narrow it to `--size-measure` (680px) with `.ds-wrap--narrow`. Only the top bar, the hero band and the footer strip paint edge to edge.
- Navigation is one bar `--size-nav` (56px) tall on `--fill-bar` with a hairline under it: brand and two or three section names at the left, and at the right a rounded search field, one small primary button and the reader's avatar. The current section has a `--border-width-strong` accent line on the bar's lower edge.
- The home view is the feed. Under the bar comes the inverted `.ds-hero` band, then a row of four editors' picks (`.ds-grid`), then `.ds-split`: the feed in a 680px column with its `.ds-tabs` on top, and a `--size-aside` (280px) column of a panel and sidebar lists at the right.
- Inner pages have no hero. They start with the `.ds-breadcrumb`, then a compact `.ds-page-header` (title, one grey line, one or two buttons at the right, an optional large avatar at the left, closed by a hairline). The story and the editor follow in the 680px measure; a profile keeps the full width, puts `.ds-stat` directly under the page header and then uses `.ds-split`: tabs, table and lists at the left, account panel and settings links at the right.
- Views. The specimen is a writing platform with four views. `home` is the feed: the hero, four picks, the tabbed list of stories and the side column with the week's prompt, topics, writers and drafts. `story` is the reading view: breadcrumb, page header, byline, the text with its drop cap and pull quote, then a row of links. `write` is the editor: the draft's page header with Preview and Publish, the notices, the compose form with its large title field, toolbar and text field, the button row and the open publish dialog. `profile` is the writer's page and settings: the figures, the table of stories with pagination and the label legend, the empty bookmarks box, and the account form, settings links and series at the right. Below 880px the side column falls under the main one and the picks go two across; below 600px the section names move to a second row of the bar, cards shrink their picture and a wide table scrolls inside `.ds-table__wrap` while the page itself never scrolls sideways.
- Spacing scale (denser than the era's other layouts): `--space-1` 4px (inside badges, under a title), `--space-2` 8px (between a label and its field, adjacent controls), `--space-3` 12px (under a byline or label, gaps in a button row), `--space-4` 20px (paragraph gap, panel padding, column side padding, between fields), `--space-5` 32px (between cards, between stacked blocks, grid gutter), `--space-6` 48px (between sections), `--space-7` 72px (hero padding).
- Other sizes: `--size-avatar` 32px, `--size-thumb` 136px (the picture of a card), `--size-dialog` 440px, `--size-icon` 40px.
- On a phone (600px and below) there is no menu button: the brand and `.ds-nav__aside` stay on the first row and `.ds-nav__links` takes a second row of its own, centred; `.ds-nav__search` is not shown from 880px down. The current link keeps its accent line. At 600px and below a tooltip opens as a note fixed to the bottom of the window, `--space-3` from its edges.

## Typography and colour roles

- `--font-heading` sets the brand, titles of stories, figures and dialog titles; `--font-body` the excerpts, the story and the fields a writer types in; `--font-ui` everything that is the application talking: navigation, buttons, tabs, bylines, labels, counts, footer. With this pack's own type set that is a slab serif for titles (Rockwell, then Roboto Slab), a transitional serif for text (Charter, then Georgia) and a plain humanist sans for the interface (Lucida Grande, then Verdana); the references used their own commercial faces.
- Body is `--text-base` 17px at `--line-body` 1.5 in `--color-text`. Lead lines (`--text-large` 20px) use a line height of 1.4 to 1.45; excerpts and page header lines are `--color-text-muted`.
- Sizes: `--text-display` 38px (hero title only), `--text-h1` 30px (page title, a part heading in a story, the editor's title field, figures), `--text-h2` 23px (card titles, section headings, pull quotes, dialog and empty titles), `--text-h3` 18px (brand, picks, panel titles), `--text-ui` 13px (navigation, buttons, tabs, sidebar lines), `--text-small` 11px (bylines, counts, labels, badges, footer). Tables, notices and form fields use the body face at 85% of `--text-base`.
- `.ds-label`, table heads, sidebar titles and badges are small capitals with letter spacing of a tenth of their size; navigation, buttons, tabs and form labels follow `--ui-transform` and carry 4% spacing.
- The text and fill pairs are those of the era, set by its first layout: page, canvas and surface colours and `--fill-panel` carry `--color-text`, `--color-heading`, `--color-heading-alt`, `--color-text-muted` and links; `--fill-bar` carries `--color-bar-text`; `--fill-bar-alt` carries `--color-bar-alt-text` (the editor toolbar and the footer); `--fill-inverse` carries `--color-inverse-text`, dimmed by opacity for the lead (the hero); `--fill-accent` carries `--color-accent-text`; `--color-fill-1` to `--color-fill-4` are picture tones (picks, card pictures, avatars) that carry only initials in `--color-inverse-text`; `--fill-button` carries `--color-button-text` inside a `--color-border-strong` border and `--fill-button-hover` carries `--color-button-hover-text`; `--fill-button-secondary` carries `--color-button-secondary-text` inside a `--color-border` hairline; `--fill-input` carries `--color-input-text`. Buttons stand on page, surface and bar colours only; on the hero band use `.ds-button--inverse` and `.ds-hero__link`.
- `--color-accent` marks the current section, tab and page number and fills the default badge. Author names in a byline are `--color-link`.
- Hairlines of `--border-width` in `--color-border` divide cards and close the page header and the bar; `--color-border-muted` divides table rows; `--border-width-strong` marks current items and the foot of a table head. `--radius-page` and `--shadow-panel` dress pictures, `--radius-panel` panels, the dialog and the empty box, `--radius-control` buttons, fields and the brand tile, `--radius-pill` badges and the search field. With this pack's surface set corners are 4px and panels, pictures and controls carry a faint 1px shadow.
- One-off values the vocabulary has no token for: avatars are circles (50%); the hero lead is dimmed by opacity 0.75; the pull quote is italic at a line height of 1.35.

## Components

- `.ds-page`: on `<body>`. Sets the page fill and the body type. It has no width; `.ds-wrap` is the column.
- `.ds-wrap`: the 1040px column with 20px side padding; `--narrow` is the 680px measure. `.ds-section` gives a block 48px above and below (`--first` 32px above, `--flush` none above). `.ds-stack` stacks blocks 32px apart. `.ds-split` is the 680px column with a 280px column at its right.
- `.ds-label`: a small capital line above a block. `.ds-avatar`: a round picture tone with initials (`--1`, `--3`, `--4` for the other tones, `--large` for a profile). `.ds-byline`: avatar, `.ds-byline__name` and a date line.
- `.ds-brand`: the site's mark and name at the left of the bar. `.ds-brand__mark` is a 28px tile dressed like the primary button (`--fill-button`, `--color-button-text`, a `--color-border-strong` border, `--radius-control`, `--shadow-control`); `.ds-brand__name` is the name at `--text-h3` in bold. Both are placeholders for the installing project's own name and logo.
- `.ds-nav`: the top bar. `.ds-nav__inner` holds the brand, `.ds-nav__links` of `.ds-nav__link` (current: `is-current`, an accent line on the bar's edge) and `.ds-nav__aside` with `.ds-nav__search`, a `.ds-button--small` and the avatar. In the specimen the showing view marks its item with one `:has()` rule per view; a project sets `is-current` itself.
- `.ds-hero`: the inverted band on the home view: `.ds-hero__inner` with `.ds-hero__title`, `.ds-hero__lead` and `.ds-hero__actions` (one `.ds-button--inverse` and one `.ds-hero__link`).
- `.ds-page-header`: head of an inner page, after the breadcrumb: an optional `.ds-avatar--large`, `.ds-page-header__copy` with `.ds-page-header__title` and one `.ds-page-header__text` line, and `.ds-page-header__actions` with at most one secondary and one primary button.
- `.ds-prose`: a story: `h1`, `h2`, `h3`, paragraphs, lists, `strong`, `code`. `.ds-prose__lead` is a larger paragraph; `.ds-prose__dropcap` is a span around the first letter of a paragraph that makes it a three-line initial. `.ds-quote` with `.ds-quote__text` is the centred italic pull quote, without rules.
- `.ds-link`: text link. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `--quiet` for titles in tables and side lists; `--more` is a small capital link with a chevron. `.ds-linkrow` is the row of links under a story, above a hairline.
- `.ds-button`: a word in a box on `--fill-button`. Variants: `--secondary` (hairline border on `--fill-button-secondary`), `--danger` (outlined in `--color-danger`, filled on hover; destructive actions only), `--inverse` (on the hero band), `--small` (in the bar). States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttonrow` spaces buttons 12px apart.
- `.ds-form`: vertical form. `.ds-form__row` sets fields side by side; `.ds-form__field` wraps `.ds-form__label` and one of `.ds-form__input`, `.ds-form__select`, `.ds-form__textarea`. Error: `is-error` on the input plus `.ds-form__error`. `.ds-form__hint`, `.ds-form__check` with `.ds-form__checkbox`, `.ds-form__actions`; `--inline` puts one field and its button on a line. For the editor: `.ds-form__input--title` is the large title field, `.ds-form__textarea--body` the tall text field, and `.ds-toolbar` the strip joined to its top with `.ds-toolbar__item` words (`--strong`, `--em`) and a `.ds-toolbar__count`.
- `.ds-table`: text rows on faint hairlines, no box and no stripes; the head is small capitals over a strong rule. `.ds-table__num` right-aligns figures; `.ds-table__caption` is a note below. Wrap it in `.ds-table__wrap`; `.ds-table__foot` holds the pagination and a legend.
- `.ds-list`: the feed. Each `.ds-list__item` is a card without a box, closed by a hairline: `.ds-list__main` with a `.ds-byline`, the `.ds-list__title`, a grey `.ds-list__text` excerpt and a `.ds-list__meta` row (a badge and counts), and an optional `.ds-list__thumb` picture (`--2`, `--3`) at the right.
- `.ds-panel`: a small card in the side column on `--fill-panel`: `.ds-panel__title` and `.ds-panel__body`.
- `.ds-stat`: a writer's figures in a row closed by a hairline. `.ds-stat` is the row; each `.ds-stat__item` is a `.ds-stat__figure` over a `.ds-stat__label`, set close together at the left, without dividers.
- `.ds-grid`: four picks across: each `.ds-grid__cell` a `.ds-grid__tile` picture tone (`--2`, `--3`, `--4`), a `.ds-grid__title` and a `.ds-grid__text` byline.
- `.ds-tabs`: words on a hairline; `.ds-tabs__tab` with `is-current` is underlined in the accent.
- `.ds-badge`: a small capital label on `--fill-accent`. `--alt` is outlined in `--color-accent-alt`, `--quiet` sits on `--color-surface-strong` for topics and counts. Status variants are outlines in their own colour: `--success`, `--warning`, `--danger`. `.ds-badgerow` for several.
- `.ds-sidebar`: a titled list in the side column: `.ds-sidebar__title` over a hairline, `.ds-sidebar__list` of `.ds-sidebar__item` with an optional `.ds-sidebar__count` at the right; the current row is `is-current`, bold and not a link.
- `.ds-notice`: a tinted line with a strong rule at its left; `--error` uses the danger colours. `.ds-notice__title` is the bold lead-in. `.ds-noticerow` stacks notices.
- `.ds-pagination`: a row of words and numbers, `.ds-pagination__link`; `is-current` is underlined in the accent, `is-disabled` is grey.
- `.ds-breadcrumb`: a small trail of `.ds-breadcrumb__item` separated by chevrons; the last is `.ds-breadcrumb__current`.
- `.ds-dialog`: a veil of `--color-overlay` centring a `.ds-dialog__box`: `.ds-dialog__head` (`.ds-dialog__title`, `.ds-dialog__close`), `.ds-dialog__body`, right-aligned `.ds-dialog__actions`.
- `.ds-empty`: a centred message in a hairline box where a list would be: `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one button.
- `.ds-footer`: one low strip on `--fill-bar-alt`: `.ds-footer__inner` with a line of text and `.ds-footer__links` of `.ds-footer__link`.
- `.ds-menu`: a `<details>` whose `.ds-menu__button` is a plain word (`--more` adds the plus sign) or the avatar, and whose `.ds-menu__list` of `.ds-menu__link` opens under it on a hairline sheet; `--end` aligns the list to the right edge. Use it for a few further places that do not earn a nav item.
- `.ds-progress`: a `--border-width-strong` line with a `.ds-progress__bar` in the accent (`--third`, `--most` set its length; a project sets the width itself) and a `.ds-progress__label` under it that says the same in words. Use it for how far a reader or a count has come, inside a margin, sidebar or panel.
- `.ds-tooltip`: a word with a hairline under it and a `.ds-tooltip__tip` shown on hover or focus (`--plain` drops the hairline, for a badge). Use it for a definition or a footnote of one sentence; it is never forced open.
- `.ds-switch`: an on or off preference: a hidden `.ds-switch__input` checkbox, the `.ds-switch__track` and the words. Use it only for a setting that takes effect by itself, such as email preferences; a statement to confirm stays a `.ds-form__check`.

## Never

- `box-shadow-blur <= 12px`: shadows are faint and close; nothing floats.
- `text-shadow = none`: text is flat.
- `gradient-fills = 0%`: every fill is one solid colour.
- `border-radius <= 4px`: corners are barely eased; only avatars are round.
- `border-width <= 2px`: rules are 1px hairlines; 2px marks the current item and a table head.
- `font-size <= 96px`: the three-line drop cap is the largest letter; the largest title is 38px.
- `font-size >= 11px`: bylines and counts are the smallest text.
- `font-weight <= 700`: titles are bold, nothing is heavier.
- `font-families <= 4`: a slab for titles, a serif for text, a sans for the interface, monospace for code.
- `line-height <= 1.5`: body text is set at 1.5.
- `letter-spacing <= 1.1px`: only small capitals are tracked.
- `content-width <= 1040px`: everything stays in the column.
- `block-gap <= 72px`: sections are 48px apart.
- `animation = none`: nothing moves.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new list is a set of unboxed rows closed by hairlines, as `.ds-list` does it; a new side block is a `.ds-sidebar` or, once per column, a `.ds-panel`. Anything the reader reads or the writer types takes the body face in the 680px measure; anything the application says takes `--font-ui` at `--text-ui` or `--text-small`. Pictures are picture tones, never images. Keep to the era's fill and text pairs so the component works with every palette: only initials on picture tones, only `.ds-button--inverse` and `.ds-hero__link` on the inverted band.
