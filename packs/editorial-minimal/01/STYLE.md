# Editorial minimalism journal

## Summary

Editorial minimalism is the reading-first page of 2012 to 2018: one narrow column of large serif text on a white or near-white page, a thin header, hairline rules and a single accent colour, with almost no boxes, fills or shadows. It was the look of writing platforms, single-author essay sites, small magazines and the blogs of design studios once web fonts and high-density screens made long text comfortable to read. Because there is so little chrome, the style lives in its typography: the size of the text, its line height, the measure and the white space around each block.

## Layout

- The page is a fixed centred column of `--size-page` (700px), about 68 characters of body text, designed for a 1280px viewport. Everything sits in this column through `.ds-wrap`: header, hero, lists, forms and footer text. Nothing is wider than the column except the inverted `.ds-letters` band and the pale footer strip, whose backgrounds run edge to edge while their content stays in the column.
- Navigation is one thin bar, `--size-nav` (72px) tall, with a hairline under it: the brand at the left edge of the column, three or four words at the right edge. There is no second row, no search and no button in it.
- The home page is a stack in the column: `.ds-hero`, then sections that each begin with a `.ds-rule` (a small label centred on a hairline), then the `.ds-letters` band and the footer. White space separates the sections, never a fill.
- Inner pages have no hero. They start with the `.ds-breadcrumb`, then the `.ds-page-header` (title, one italic line, one action at the right, closed by a hairline), then the working content in the same column: `.ds-tabs` come first under the page header, then the table or the text. An essay page widens `.ds-wrap` with `--margins`: the text keeps its 700px column and two `--size-margin` (200px) note columns appear beside it, the left one for the date and length (`.ds-margin--left`, set right-aligned against the text), the right one for a `.ds-sidebar`. Below 1240px the notes fall back into the column, before and after the text.
- Views. The specimen is a single-author journal of walking essays with four views. `home` is the journal front: the hero, the four most recent essays, the four series, the book panel and the letters band. `essay` is one essay: breadcrumb, page header, the text with its drop cap, pull quote, parts and footnotes between the two margin columns, then further reading. `archive` is the list page: tabs, the table of essays with pagination and the label legend, and the empty translations section. `about` is the author page: a short text, the figures, then the letter form with its notices, the button row and the open discard dialog. Below 760px the navigation wraps under the brand, the grid folds to two columns, form rows stack and a wide table scrolls inside `.ds-table__wrap` while the page itself never scrolls sideways.
- Spacing scale: `--space-1` 4px (inside badges, under a title), `--space-2` 8px (between a label and its field, between adjacent controls), `--space-3` 16px (cell padding, gaps in a button row), `--space-4` 24px (paragraph gap, list row padding, column side padding, grid gutter), `--space-5` 40px (between stacked blocks, under a rule or tabs, panel padding, the margin column gap), `--space-6` 64px (between sections), `--space-7` 104px (above the hero).
- Other sizes: `--size-dialog` 460px (dialog and subscription form), `--size-icon` 40px.
- On a phone (760px and below) there is no menu button: `.ds-nav__links` takes a row of its own under the brand and its links spread over the full width; the `.ds-menu` among them keeps its hairline sheet but aligns it to its right edge. The current link keeps its underline. At 600px and below a tooltip opens as a note fixed to the bottom of the window, `--space-3` from its edges.

## Typography and colour roles

- Body and headings share one book serif (`--font-body`, `--font-heading`); the references used commercial text faces and open ones such as Cardo, so the stack is Iowan Old Style, Palatino, then EB Garamond and Georgia. `--font-ui` is a humanist sans used only small: navigation, buttons, labels, meta lines, table heads, footer. `--font-mono` is for inline code.
- Body is `--text-base` 20px at `--line-body` 1.6 in `--color-text`. Headings are the same face at regular weight (`--weight-heading` 400): hierarchy comes from size and white space, not from bold or colour. `--weight-bold` 700 is for `strong` only; `--weight-ui` 600 for the small sans.
- Sizes: `--text-display` 44px (hero title only), `--text-h1` 36px (page title, part titles in an essay, figures), `--text-h2` 26px (list titles, section headings, pull quotes, panel and dialog titles), `--text-h3` 20px (brand name, grid titles), `--text-large` 24px (hero lead, first paragraph, page header line), `--text-small` 13px (labels and meta), `--text-ui` 12px (navigation, buttons, tabs, form labels). Tables, form fields, notices and footnotes use the body face at 85% or 80% of `--text-base`, a one-off size made with `calc()`.
- Labels and meta lines are always small capitals of the sans: `text-transform: uppercase` with letter spacing of a tenth of the size (`calc(var(--text-small) * 0.1)`). This is part of the structure and does not follow `--ui-transform`, which sets navigation, buttons, tabs and form labels; those carry a spacing of 4% of their size. An `h3` in running text is such a label, in `--color-heading-alt`.
- The drop cap (`.ds-prose__dropcap`) is three lines deep: its size is computed from `--text-base` and `--line-body`, and it shares a baseline with an unseen strut in the text face set two lines down, so it stands on the third baseline with any type set. The pull quote is `--text-h2` at a line height of 1.35.
- Links in text are `--color-link` with `--link-decoration`; hover changes to `--color-link-hover`. Quiet links (`--color-link-quiet`, `--link-decoration-quiet`) are for titles, breadcrumbs and margin lists.
- Which text sits on which fill (the convention of this era, kept by every layout and token set):
  - `--color-page`, `--color-canvas`, `--color-surface`, `--color-surface-alt`, `--color-surface-strong` and `--fill-panel` carry `--color-text`, `--color-heading`, `--color-heading-alt`, `--color-text-muted` and the link colours. The surfaces are the page colour or a pale step from it.
  - `--fill-bar` carries `--color-bar-text` (navigation and table head); the brand name is `--color-bar-text`. `--fill-bar-alt` carries `--color-bar-alt-text` only.
  - `--fill-inverse` carries `--color-inverse-text`; quieter text on it is the same colour at an opacity of 0.65 to 0.75. The only button allowed on it is `.ds-button--inverse`.
  - `--fill-accent` carries `--color-accent-text`. `--color-danger` used as a fill also carries `--color-accent-text`.
  - `--color-fill-1` to `--color-fill-4` are picture tones: they stand where a photograph or cover would be. They carry nothing but a numeral, mark or short title in `--color-inverse-text`, never running text.
  - `--fill-button` carries `--color-button-text` inside a `--color-border-strong` border; `--fill-button-hover` carries `--color-button-hover-text`. `--fill-button-secondary` carries `--color-button-secondary-text` inside a `--color-border` hairline. Buttons stand on page and surface colours only.
  - `--fill-input` carries `--color-input-text` inside `--color-input-border`; `--color-notice` carries `--color-notice-text`; `--color-danger-surface` carries `--color-danger`; `--color-disabled` carries `--color-disabled-text`.
  - `--color-success`, `--color-warning` and `--color-danger` are text and outline colours on the page (status labels, the destructive button), so each must be readable on `--color-page`.
- `--color-accent` is the single accent: the underline of the current navigation item, tab and page number, the checkbox, the default badge. `--color-accent-alt` only outlines the `--alt` badge.
- Rules: hairlines are `--border-width` in `--color-border`; faint row dividers use `--color-border-muted`; `--border-width-strong` in `--color-border-strong` marks the top of a pull quote, the figures and the foot of a table head. `--radius-page` and `--shadow-panel` dress the picture tones, `--radius-panel` the panel and dialog, `--radius-control` buttons, fields and the brand tile, `--radius-pill` badges.
- One-off values the vocabulary has no token for: the slash of the breadcrumb is `--color-border-strong` at an opacity of 0.4; text on the inverted band is dimmed by opacity.

## Components

- `.ds-page`: on `<body>`. Sets the page fill, the body face, size and line height. It has no width; `.ds-wrap` is the column.
- `.ds-wrap`: the 700px column with 24px side padding. `--margins` turns it into three columns for an essay: `.ds-margin--left`, `.ds-wrap__main`, `.ds-margin--right`; `.ds-margin__text` is a note under a label. `.ds-section` gives a block 64px above and below (`--first` 40px above, directly under a page header; `--flush` none above, after another section). `.ds-stack` stacks blocks 40px apart.
- `.ds-label`: a small tracked capital line: a kicker above a title, a date in the margin. `.ds-rule` is the section head: the same lettering centred on a hairline.
- `.ds-brand`: the site's mark and name at the left of the navigation. `.ds-brand__mark` is a 28px tile dressed like the primary button (`--fill-button`, `--color-button-text`, a `--color-border-strong` border, `--radius-control`, `--shadow-control`); `.ds-brand__name` is the name in the heading face at `--text-h3`. Both are placeholders for the installing project's own name and logo.
- `.ds-nav`: the thin bar on `--fill-bar` with a hairline below. `.ds-nav__inner` is the column; `.ds-nav__links` holds `.ds-nav__link` items; the current one (`is-current`) has a `--border-width-strong` underline in `--color-accent`. In the specimen the showing view marks its item with one `:has()` rule per view; a project sets `is-current` itself.
- `.ds-hero`: the statement of the site at the top of the home page, unboxed: a `.ds-label`, `.ds-hero__title` at `--text-display`, `.ds-hero__lead` at `--text-large` and `.ds-hero__actions` with one button and one `.ds-link--more`.
- `.ds-page-header`: head of an inner page, after the breadcrumb: `.ds-page-header__copy` with `.ds-page-header__title` (`--text-h1`) and one italic `.ds-page-header__text` line, `.ds-page-header__actions` at the right with a single button, and a hairline below.
- `.ds-prose`: running text: `h1` (a centred part title), `h2`, `h3` (a label), paragraphs, lists, `strong`, `code`, `sup`. `.ds-prose__lead` is a larger first paragraph; `.ds-prose__dropcap` is a span around the first letter of a paragraph that makes it a three-line initial. `.ds-quote` with `.ds-quote__text` is the pull quote between a strong and a hairline rule. `.ds-footnotes` with `.ds-footnotes__item` is the numbered note list under an essay.
- `.ds-link`: text link. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `--quiet` for titles and meta; `--more` is a small capital link with a chevron. `.ds-linklist` with `.ds-linklist__item` sets links as rows on faint rules.
- `.ds-button`: a small-capital word in a bordered box on `--fill-button`; with this pack's palette it is an outline that fills on hover. Variants: `--secondary` (hairline border, accent text), `--danger` (outlined in `--color-danger`, filled on hover; destructive actions only), `--inverse` (on the inverted band), `--small`. States `is-hover`, `is-active`, `is-focus`, `is-disabled`. `.ds-buttonrow` spaces buttons 16px apart.
- `.ds-form`: vertical form, 24px between fields. `.ds-form__row` sets fields side by side; `.ds-form__field` wraps `.ds-form__label` and one of `.ds-form__input`, `.ds-form__select`, `.ds-form__textarea`. Error: `is-error` on the input plus `.ds-form__error`. `.ds-form__hint`, `.ds-form__check` with `.ds-form__checkbox`, `.ds-form__actions`. `--inline` puts one field and its button on a line.
- `.ds-table`: text rows on faint hairlines, no box and no stripes; the head is small capitals over a strong rule. `.ds-table__num` right-aligns figures; `.ds-table__caption` is a note below. Wrap it in `.ds-table__wrap`; `.ds-table__foot` is the row under it with the pagination and a legend.
- `.ds-list`: the journal: `.ds-list__item` rows on faint rules, each a `.ds-list__title` (`--text-h2`), an italic `.ds-list__text` line and a small capital `.ds-list__meta` line. `.ds-list__more` holds the link to the archive.
- `.ds-panel`: the one bordered box of a page, on `--fill-panel` with 40px padding: `.ds-panel__title` and `.ds-panel__body`. `.ds-panel__split` puts a `.ds-panel__art` cover (a picture tone with its title) beside `.ds-panel__copy`.
- `.ds-stat`: the row of figures between a strong rule above and a hairline below. `.ds-stat` is the row; each `.ds-stat__item` is a `.ds-stat__figure` (`--text-h1`) over a `.ds-stat__label`, divided by vertical hairlines.
- `.ds-grid`: four equal columns of `.ds-grid__cell`, each a `.ds-grid__tile` in a picture tone (`--2`, `--3`, `--4` for the others) holding a numeral, with `.ds-grid__title` and `.ds-grid__text` below. The cell has no border.
- `.ds-tabs`: words on a hairline; `.ds-tabs__tab` with `is-current` is underlined in the accent.
- `.ds-badge`: a small capital label on `--fill-accent`. `--alt` is outlined in `--color-accent-alt`, `--quiet` sits on `--color-surface-strong` for counts. Status variants are outlines in their own colour: `--success`, `--warning`, `--danger`. `.ds-badgerow` for several.
- `.ds-sidebar`: a short list in the margin: `.ds-sidebar__title` over a hairline, `.ds-sidebar__list` of `.ds-sidebar__item`; the current row is `is-current`, bold and not a link.
- `.ds-notice`: a tinted line with a strong rule at its left; `--error` uses the danger colours. `.ds-notice__title` is the bold lead-in. `.ds-noticerow` stacks notices.
- `.ds-pagination`: a row of small words and numbers, `.ds-pagination__link`; `is-current` is underlined in the accent, `is-disabled` is grey.
- `.ds-breadcrumb`: a small capital trail of `.ds-breadcrumb__item` separated by slashes; the last is `.ds-breadcrumb__current`.
- `.ds-dialog`: a veil of `--color-overlay` centring a plain `.ds-dialog__box` with a strong hairline border: `.ds-dialog__head` (`.ds-dialog__title`, `.ds-dialog__close`), `.ds-dialog__body`, right-aligned `.ds-dialog__actions`.
- `.ds-empty`: a centred message closed by a hairline, standing where a list would be: `.ds-empty__icon`, `.ds-empty__title`, an italic `.ds-empty__text` and one button.
- `.ds-letters`: the inverted band that closes the home page: a label, `.ds-letters__title`, `.ds-letters__text` and an inline `.ds-letters__form`.
- `.ds-footer`: one pale strip on `--fill-bar-alt`: `.ds-footer__inner` holds a line of text and `.ds-footer__links` of `.ds-footer__link`.
- `.ds-menu`: a `<details>` whose `.ds-menu__button` is a plain word (`--more` adds the plus sign) or the avatar, and whose `.ds-menu__list` of `.ds-menu__link` opens under it on a hairline sheet; `--end` aligns the list to the right edge. Use it for a few further places that do not earn a nav item.
- `.ds-avatar`: a person, as initials in a `--size-avatar` circle on a picture tone (`--1`, `--4` for the other tones). Use it beside a name, never alone as decoration. `.ds-credit`: the row that holds an avatar, a `.ds-credit__name` and one quiet line.
- `.ds-progress`: a `--border-width-strong` line with a `.ds-progress__bar` in the accent (`--third`, `--most` set its length; a project sets the width itself) and a `.ds-progress__label` under it that says the same in words. Use it for how far a reader or a count has come, inside a margin, sidebar or panel.
- `.ds-tooltip`: a word with a hairline under it and a `.ds-tooltip__tip` shown on hover or focus (`--plain` drops the hairline, for a badge). Use it for a definition or a footnote of one sentence; it is never forced open.
- `.ds-switch`: an on or off preference: a hidden `.ds-switch__input` checkbox, the `.ds-switch__track` and the words. Use it only for a setting that takes effect by itself, such as email preferences; a statement to confirm stays a `.ds-form__check`.

## Never

- `box-shadow = none`: nothing is lifted off the page; blocks are separated by white space or a hairline.
- `text-shadow = none`: text is flat.
- `gradient-fills = 0%`: every fill is one solid colour.
- `border-radius <= 3px`: only controls are eased; panels and pictures are square.
- `border-width <= 2px`: rules are 1px hairlines; 2px marks the current item, a pull quote and a table head.
- `font-size <= 110px`: the three-line drop cap is the largest letter; the largest title is 44px.
- `font-size >= 12px`: the small sans of navigation and buttons is the smallest text.
- `font-weight <= 700`: headings are regular; bold is for emphasis in text.
- `font-families <= 3`: one serif, one small sans, monospace for code.
- `line-height <= 1.6`: body text is set at 1.6.
- `letter-spacing <= 1.3px`: only small capitals are tracked, by a tenth of their size.
- `content-width <= 700px`: text stays in the reading column.
- `block-gap <= 104px`: the largest space is the one above the hero.
- `animation = none`: nothing moves.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new section is a `.ds-section` in the `.ds-wrap` column that begins with a `.ds-rule`; separate things with white space from the `--space-*` scale first, a hairline second, and a box only once per page (`.ds-panel`). New small text is a `.ds-label`; new titles take `--text-h2` in the heading face at `--weight-heading`. Keep to the fill and text pairs listed above so that the component works with every palette of the era: never put running text on a picture tone, and put only `.ds-button--inverse` on an inverted band.
