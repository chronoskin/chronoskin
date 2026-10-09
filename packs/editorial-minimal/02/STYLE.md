# Editorial minimalism magazine

## Summary

Editorial minimalism is the reading-first page of 2012 to 2018: large, carefully set text on a white or near-white page, hairline and strong rules instead of boxes, a single accent colour and no gloss, gradient or decoration. This layout is its magazine form: a centred masthead over a ruled row of section names, a lead story in very large type, and stories set in columns divided by vertical hairlines, as small independent magazines and essay publications of the period laid out their front pages. Its quality lives in the type: headline size, line height, the measure of the article and the space between rules.

## Layout

- The page is a fixed centred sheet of `--size-page` (1080px), designed for a 1280px viewport; `.ds-wrap` holds every row. Only the `.ds-topline` strip and the inverted footer paint edge to edge.
- Navigation is a masthead, not a bar: the thin `.ds-topline` strip (issue at the left, offer at the right), then the brand centred and large, then the section names centred in one row between a hairline above and a strong rule below. The current section is marked by an accent piece of that strong rule.
- The front page is three rows on the sheet. First the `.ds-hero`: lead story at the left over seven of twelve columns, its picture at the right over five. Then the `.ds-grid`: four stories in columns divided by vertical hairlines. Then the `.ds-trio`: three unequal columns (5, 3 and 4 parts) divided by vertical hairlines. Every row after the hero begins with a `.ds-rule`, a small capital label hanging under a strong rule. Rows are separated by `--space-6`.
- Inner pages have no hero. They start with the `.ds-breadcrumb`, then the `.ds-page-header` (an optional accent label, the title at `--text-display`, one line of text, one action at the right). A story then shows a wide `.ds-figure` across the sheet and the `.ds-article`: the measure of `--size-measure` (640px) in the middle, the byline column of `--size-side` (180px) at its left and a `.ds-sidebar` of the same width at its right. A list page puts `.ds-stat` directly under the page header, then `.ds-tabs`, then the table. A form page uses `.ds-split`: the form in the wide column, panels and a sidebar in the `--size-aside` (320px) column.
- Views. The specimen is a quarterly magazine about cities with four views. `home` is the front page: lead story with its drawing, the four stories of the issue, then the most read list, the contents of the issue and the subscription panel. `story` is one feature: breadcrumb, page header, wide figure, the text with its drop cap, part heading and pull quote between byline and related stories, then links to more. `issues` is the back issue list: the figures, year tabs, the table with pagination and the label legend, and the empty digital editions block. `subscribe` is the form page: notices, the address form, the button row and the open cancel dialog beside the two plans and the questions. Below 960px the hero stacks, the story grid goes to two columns and the side columns fall under the main one; below 600px everything is one column, the grid cells are divided by horizontal hairlines and a wide table scrolls inside `.ds-table__wrap` while the page itself never scrolls sideways.
- Spacing scale: `--space-1` 4px (inside badges, under a title), `--space-2` 8px (above a rule's label, between a label and its field), `--space-3` 16px (cell padding, navigation padding, gaps in a button row), `--space-4` 24px (paragraph gap, grid cell padding, panel padding, sheet side padding), `--space-5` 40px (between stacked blocks, trio column padding, the brand tile), `--space-6` 56px (between rows, hero padding and column gap), `--space-7` 88px (reserved for the textarea height and larger gaps).
- Other sizes: `--size-dialog` 440px, `--size-icon` 40px.
- On a phone (600px and below) there is no menu button: the brand stays centred over `.ds-nav__links`, which is one row with a `--space-3` gap, spread over the width and scrolling sideways if the words do not fit; the current link keeps its accent piece of the rule. At 600px and below a tooltip opens as a note fixed to the bottom of the window, `--space-3` from its edges.

## Typography and colour roles

- The type set decides the voice; the structure only assigns sizes. `--font-heading` sets the brand, headlines, list and grid titles, figures and pull quotes; `--font-body` the running text; `--font-ui` everything small: section names, labels, buttons, captions, footer. With this pack's own type set all three are one grotesque sans (Helvetica Neue, then Helvetica and Arial; the references used commercial grotesques), with headlines at weight 700 and tight tracking.
- Body is `--text-base` 18px at `--line-body` 1.55 in `--color-text`. Lead lines (`--text-large` 22px) use a line height of 1.45.
- Sizes: the lead headline is `--text-display` times 1.25 (80px), the largest type on any page; `--text-display` 64px is the page title of inner pages and the figures; `--text-h1` 44px the brand, a part heading in a story, the list numerals and panel prices; `--text-h2` 28px section headings, pull quotes, panel titles (grid titles use 85% and list titles 80% of it); `--text-h3` 19px small headings and the sidebar title; `--text-ui` 14px navigation, buttons, tabs; `--text-small` 12px labels, captions, meta. Tables, form fields, notices and side text use 85% of `--text-base`.
- Labels are small capitals of the UI face with letter spacing of a tenth of their size; they are part of the structure and do not follow `--ui-transform`, which sets navigation, buttons, tabs and form labels. `.ds-label--accent` colours the label of a lead story with `--color-accent`.
- The text and fill pairs are those of the era, set by its first layout: page, canvas and surface colours carry `--color-text`, `--color-heading`, `--color-heading-alt`, `--color-text-muted` and links; `--fill-bar` carries `--color-bar-text`; `--fill-bar-alt` carries `--color-bar-alt-text` (the topline); `--fill-inverse` carries `--color-inverse-text`, dimmed by opacity for quieter lines (the footer); `--fill-accent` carries `--color-accent-text`; `--color-fill-1` to `--color-fill-4` are picture tones that carry only a short label or a mark in `--color-inverse-text`; `--fill-button` carries `--color-button-text` inside a `--color-border-strong` border and `--fill-button-hover` carries `--color-button-hover-text`; `--fill-button-secondary` carries `--color-button-secondary-text` inside a `--color-border` hairline; buttons stand on page and surface colours only. Status colours are text and outline colours on the page.
- `--color-accent` is the single accent: the current section, tab and page number, the lead label, the drop cap and the rule of a pull quote. It is used as text on the page, so it must be readable there.
- Rules do the work of boxes. `--border-width-strong` in `--color-border-strong` closes the masthead and opens every section (`.ds-rule`), the figures and the byline column; hairlines of `--border-width` in `--color-border` divide columns; `--color-border-muted` divides rows inside a list, table or sidebar.
- Pictures are drawn, not loaded: `.ds-art` is a flat composition of bars in the picture tones on `--color-fill-1` with a disc in `--color-inverse-text`. `--radius-page` and `--shadow-panel` dress pictures and grid tiles, `--radius-panel` panels, the dialog and the empty block, `--radius-control` buttons, fields and the brand tile, `--radius-pill` badges.
- One-off values the vocabulary has no token for: the footer rule is `--color-inverse-text` at 30% (`color-mix()`); quiet footer text and the disc use opacity; bar heights and gaps inside `.ds-art` are percentages.

## Components

- `.ds-page`: on `<body>`. Sets the page fill and the body type. It has no width; `.ds-wrap` is the sheet.
- `.ds-wrap`: the 1080px sheet with 24px side padding. `.ds-section` gives a row 56px above and below (`--first` 40px above, `--flush` none above). `.ds-stack` stacks blocks 40px apart. `.ds-split` is a wide column with a 320px column beside it. `.ds-trio` with `.ds-trio__col` is the row of three ruled columns.
- `.ds-label`: a small capital line for a section name, byline or meta; `--accent` in the accent colour. `.ds-rule` is the section head: the same lettering under a strong rule, with an optional quiet link at its right.
- `.ds-topline`: the strip on `--fill-bar-alt` above the masthead: `.ds-topline__inner` with `.ds-topline__note` and `.ds-topline__link`.
- `.ds-brand`: the site's mark and name, centred in the masthead. `.ds-brand__mark` is a 40px tile dressed like the primary button (`--fill-button`, `--color-button-text`, a `--color-border-strong` border, `--radius-control`, `--shadow-control`); `.ds-brand__name` is the name at `--text-h1` in the display weight. Both are placeholders for the installing project's own name and logo.
- `.ds-nav`: the masthead on `--fill-bar`. `.ds-nav__inner` centres the brand over `.ds-nav__links`; a `.ds-nav__link` that is current (`is-current`) turns its piece of the strong rule to the accent. In the specimen the showing view marks its item with one `:has()` rule per view; a project sets `is-current` itself.
- `.ds-hero`: the lead story: `.ds-hero__copy` with a `.ds-label--accent`, `.ds-hero__title` in the largest type, `.ds-hero__lead` and `.ds-hero__actions` (one button and the byline), beside a `.ds-figure`.
- `.ds-figure`: a picture with its `.ds-figure__caption` under a hairline. `.ds-art` is the drawn picture (`--wide` for the figure across an article): `.ds-art__bar` blocks with `--2` and `--4` for the other tones, `--tall`, `--mid`, `--low` for height and `--slim` for width, and one `.ds-art__sun` disc.
- `.ds-page-header`: head of an inner page, after the breadcrumb: `.ds-page-header__copy` with an optional label, `.ds-page-header__title` and one `.ds-page-header__text` line; `.ds-page-header__actions` at the right with one button.
- `.ds-article`: the three columns of a story: `.ds-article__meta` (labels and short lines under a strong rule), `.ds-article__main` and `.ds-article__side`.
- `.ds-prose`: running text: `h1` (a part heading over a hairline), `h2`, `h3`, paragraphs, lists, `strong`, `code`. `.ds-prose__lead` is a larger paragraph; `.ds-prose__dropcap` is a span around the first letter of a paragraph that makes it a three-line initial in the accent colour. `.ds-quote` with `.ds-quote__text` is the pull quote on a strong accent rule at its left.
- `.ds-link`: text link. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `--quiet` for titles and meta; `--more` is a small capital link with a chevron. `.ds-linkrow` sets links in a wrapping row.
- `.ds-button`: a word in a box on `--fill-button` with a `--color-border-strong` border. Variants: `--secondary` (hairline border on `--fill-button-secondary`), `--danger` (outlined in `--color-danger`, filled on hover; destructive actions only), `--inverse` (on an inverted block), `--small`. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttonrow` spaces buttons 16px apart.
- `.ds-form`: vertical form. `.ds-form__row` sets fields side by side; `.ds-form__field` wraps `.ds-form__label` and one of `.ds-form__input`, `.ds-form__select`, `.ds-form__textarea`. Error: `is-error` on the input plus `.ds-form__error`. `.ds-form__hint`, `.ds-form__check` with `.ds-form__checkbox`, `.ds-form__actions`. `--inline` puts one field and its button on a line.
- `.ds-table`: text rows on faint hairlines, no box and no stripes; the head is small capitals over a strong rule. `.ds-table__num` right-aligns figures; `.ds-table__caption` is a note below. Wrap it in `.ds-table__wrap`; `.ds-table__foot` holds the pagination and a legend.
- `.ds-list`: the numbered list of most read stories: each `.ds-list__item` gets a large grey numeral from a counter, a `.ds-list__title` and a `.ds-list__meta` line.
- `.ds-panel`: a ruled box on `--fill-panel` for an offer or a plan: a label, `.ds-panel__title`, `.ds-panel__body` and an optional `.ds-panel__figure` price.
- `.ds-stat`: the row of figures under a strong rule. `.ds-stat` is the row; each `.ds-stat__item` is a `.ds-stat__figure` at `--text-display` over a `.ds-stat__label`, divided by vertical hairlines.
- `.ds-grid`: four story columns divided by vertical hairlines. Each `.ds-grid__cell` holds a `.ds-grid__tile` in a picture tone (`--2`, `--3`, `--4`) carrying its section name, a `.ds-grid__title`, a `.ds-grid__text` line and a byline label.
- `.ds-tabs`: words on a hairline; `.ds-tabs__tab` with `is-current` is underlined in the accent.
- `.ds-badge`: a small capital label on `--fill-accent`. `--alt` is outlined in `--color-accent-alt`, `--quiet` sits on `--color-surface-strong` for counts. Status variants are outlines in their own colour: `--success`, `--warning`, `--danger`. `.ds-badgerow` for several.
- `.ds-sidebar`: a titled list on faint rules: `.ds-sidebar__title`, `.ds-sidebar__list` of `.ds-sidebar__item`, each with an optional `.ds-sidebar__count` (a page number) at the right; the current row is `is-current`, bold and not a link.
- `.ds-notice`: a tinted line with a strong rule at its left; `--error` uses the danger colours. `.ds-notice__title` is the bold lead-in. `.ds-noticerow` stacks notices.
- `.ds-pagination`: a row of words and numbers, `.ds-pagination__link`; `is-current` is underlined in the accent, `is-disabled` is grey.
- `.ds-breadcrumb`: a small capital trail of `.ds-breadcrumb__item` separated by chevrons; the last is `.ds-breadcrumb__current`.
- `.ds-dialog`: a veil of `--color-overlay` centring a `.ds-dialog__box` with a strong hairline border: `.ds-dialog__head` (`.ds-dialog__title`, `.ds-dialog__close`), `.ds-dialog__body`, right-aligned `.ds-dialog__actions`.
- `.ds-empty`: a centred message on `--color-surface-alt`, standing where a list would be: `.ds-empty__icon`, `.ds-empty__title`, `.ds-empty__text` and one button.
- `.ds-footer`: the inverted block on `--fill-inverse`: `.ds-footer__cols` with the name (`.ds-footer__name`, `.ds-footer__text`) and three columns of `.ds-footer__title` and `.ds-footer__list` / `.ds-footer__item` / `.ds-footer__link`, then `.ds-footer__base`.
- `.ds-avatar`: a person, as initials in a `--size-avatar` circle on a picture tone (`--1`, `--4` for the other tones). Use it beside a name, never alone as decoration. `.ds-credit`: the row that holds an avatar, a `.ds-credit__name` and one quiet line.
- `.ds-progress`: a `--border-width-strong` line with a `.ds-progress__bar` in the accent (`--third`, `--most` set its length; a project sets the width itself) and a `.ds-progress__label` under it that says the same in words. Use it for how far a reader or a count has come, inside a margin, sidebar or panel.
- `.ds-accordion`: `<details class="ds-accordion__item">` rows on hairlines, each a `.ds-accordion__head` question and a `.ds-accordion__body` answer. Use it for a short list of questions or notes in a side column or panel.
- `.ds-tooltip`: a word with a hairline under it and a `.ds-tooltip__tip` shown on hover or focus (`--plain` drops the hairline, for a badge). Use it for a definition or a footnote of one sentence; it is never forced open.
- `.ds-switch`: an on or off preference: a hidden `.ds-switch__input` checkbox, the `.ds-switch__track` and the words. Use it only for a setting that takes effect by itself, such as email preferences; a statement to confirm stays a `.ds-form__check`.

## Never

- `box-shadow = none`: nothing is lifted off the page; rules divide it.
- `text-shadow = none`: text is flat.
- `gradient-fills = 0%`: every fill is one solid colour.
- `border-radius <= 0px`: every corner is square; only the drawn disc is round.
- `border-width <= 3px`: hairlines are 1px; the strong rule is 3px.
- `font-size <= 110px`: the lead headline is 80px; only the drop cap of a story is larger.
- `font-size >= 12px`: labels and captions are the smallest text.
- `font-weight <= 700`: headlines are bold, nothing is heavier.
- `font-families <= 2`: one family for everything, plus monospace for code.
- `line-height <= 1.55`: body text is set at 1.55.
- `letter-spacing <= 2px`: small capitals are spaced out by 1.2px; the largest headline is tightened by 2px.
- `content-width <= 1080px`: everything stays on the sheet.
- `block-gap <= 88px`: rows are 56px apart.
- `transition = none`: states change at once.
- `animation = none`: nothing moves.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new row is a `.ds-section` on the `.ds-wrap` sheet that begins with a `.ds-rule`; divide its columns with vertical hairlines as `.ds-grid` and `.ds-trio` do, and its rows with `--color-border-muted`. Use a box (`.ds-panel`) only for an offer. A new picture is a `.ds-art` or a tile in a picture tone, never an image. New headlines take the heading face at `--text-h2` or a `calc()` step of it; new small text is a `.ds-label`. Keep to the era's fill and text pairs so the component works with every palette: only labels and marks on picture tones, only `.ds-button--inverse` on an inverted block.
