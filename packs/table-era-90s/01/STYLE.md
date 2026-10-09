# Table-era portal, 1998

## Summary

A directory-and-search portal home page as built between 1996 and 2000: a fixed 600px column centred in an 800px window, laid out as tight cells with flat colour fills and no decoration beyond rules and bevelled form controls. Nearly all text is an underlined dark blue link, set at 13px in a serif body face with bold sans-serif headings and 10px sans-serif small print, on a white page with one flat teal band, yellow title strips and pale yellow boxes. It was the standard look of search engines, directories and start pages until wider, stylesheet-driven layouts replaced it after 2000.

## Layout

- The page is fixed, not fluid: `--size-page` is 600px, centred by `.ds-page__frame`, which carries a `--space-3` (8px) margin inside it. Nothing is wider than 600px and nothing stretches with the window.
- The specimen is one example site, a portal, of five views. `.ds-nav` and `.ds-footer` are written once, outside the views; each view is a `.ds-view` section and one shows at a time.
- Order from the top on the home view: `.ds-nav` (wordmark left, tagline, 10px tool links right, then one centred row of text links separated by hyphens), the `.ds-hero` search band at full column width, the `.ds-tabs` strip, a 2px `.ds-rule`, then `.ds-page__columns`, then the `.ds-footer` opened by another `.ds-rule`.
- `.ds-page__columns` is two columns: a main column and a `--size-sidebar` (190px) `.ds-sidebar` on the right, separated by a `--space-4` (10px) gutter and the sidebar's 1px left rule. The main column opens with the two-column `.ds-directory`; the sidebar carries enough short sections to end near the foot of the main column.
- Inner pages (a category, a result list, a settings page, a help page) keep the `.ds-nav` and the `.ds-footer` and drop the search band. In its place, inside `.ds-page__head` with no gaps between them: the `.ds-breadcrumb`, the `.ds-page-header` (title and one line at the left, one button at the right), the `.ds-tabs` strip for the views of that page, and the 2px `.ds-rule`. Below the rule inner pages are one full 600px column, `.ds-page__inner`, with no sidebar and `--space-5` between blocks: `.ds-stat` first when the page has figures, then the `.ds-table` or `.ds-list` with its `.ds-pagination`, then running text or the `.ds-form`. What the sidebar held on the home page becomes a `.ds-panel` at the foot of the column if it is needed at all.
- Views: `home` is the portal front page (search band, channel strip, directory, headlines, grid, markets panel, sidebar). `category` is one directory category: figures, the table of listed sites with its status key, result pages and a panel of related categories. `search` is a result list: category matches, site matches with result pages, and the empty state for a kind of match that found nothing. `mypage` is the settings screen: notices, the form, the table of sections on the member's page and the message box that removing a section opens. `help` is a page of running text, including the line that shows the link states, with a boxed panel of further questions. The current view's link in `.ds-nav__links` is bold, in the text colour and not underlined; the home link is marked when the address names no view.
- Blocks are separated by rules, title strips and flat fills, never by cards with shadows or wide gutters.
- Spacing scale: `--space-1` 2px is the spacing between table and grid cells and the padding of controls and strips; `--space-2` 4px is cell padding, the padding around separators and the gap between a title and its body; `--space-3` 8px is the page margin, paragraph spacing, box padding and the gap between buttons; `--space-4` 10px is the gutter between columns; `--space-5` 16px is the gap between neighbouring blocks of the main column and the side padding of the search band; `--space-6` 40px is the list indent (and the width of a short input).
- Fixed sizes: `--size-search` 250px search field, `--size-label` 110px form label column, `--size-field` 180px text input, `--size-area` 240px textarea, `--size-check` 13px checkbox, `--size-dialog` 320px message box.

## Typography and colour roles

- Two families and a code face. `--font-body` (Times New Roman, Times) sets running text, link lists, table data and the footer. `--font-heading` and `--font-ui` (Arial, Helvetica) set the wordmark, headings, title strips, the search band, forms, buttons, notices and all 10px small print. `--font-mono` (Courier New) is for code and textareas. The references used Times, Arial, Geneva, Helvetica and Verdana as installed system fonts; no substitute was needed.
- Sizes: `--text-base` and `--text-ui` 13px; `--text-small` 10px for tool links, tab strip, captions, meta and legal text; `--text-large` and `--text-h2` 16px for directory titles, panel titles and section headings; `--text-h1` 18px for a page title; `--text-h3` 13px bold; `--text-display` 28px for the wordmark only. Line height is `normal` everywhere. Bold is 700, the wordmark included.
- No uppercase transform and no letter spacing. Headings are distinguished by family, weight and size only.
- Links: `--color-link` (dark blue) and always underlined, including navigation, titles and quiet links (`--link-decoration`, `--link-decoration-quiet` and `--link-decoration-hover` are all `underline`). `--color-link-visited` is a greyed violet, `--color-link-active` is red, `--color-link-hover` equals the link colour because the era had no hover effect. `--color-link-quiet` (bright blue) is for promotional links set in the sans face.
- Text is `--color-text` black; `--color-text-muted` grey is for times and counts only. `--color-heading` is black; `--color-heading-alt` (navy) is the wordmark, tagline, h3 and positive figures.
- `--color-success` (dark green, the one green the references print) and `--color-warning` (the title-strip yellow, the era's amber) are flat fills for status labels. Green also works as text on white; the yellow is a fill only and carries black text.
- Fills are flat. `--color-bar` (teal) is the search band, table head and the 2px rules. `--color-bar-alt` (yellow) is every panel and sidebar title strip and the border of notices. `--color-inverse` (black, white text) is the dialog title bar only. `--color-surface-alt` (light grey) is the tab strip, alternate table rows and the empty state. `--color-fill-1` to `--color-fill-4` (pale yellow, light grey, orange, teal) colour grid cells; `--color-fill-3` is also the 2px underline of `.ds-heading`.
- `--color-accent` (crimson) is the solid badge; `--color-accent-alt` (red) is the "New!" word and negative figures. `--color-danger` is error text and the error notice border on `--color-danger-surface` (pale yellow).
- Borders: 1px `--color-border` grey for table cells, the outline of the search band and of grid cells, the sidebar rule and thin rules; 1px `--color-border-strong` black for boxed panels and dialogs; 2px bevels on controls. Radius is 0 and every shadow token is `none`. Every one of these lines is drawn with `--border-width`, `--border-width-strong` and `--border-style`, and the sheet, panels and dialogs take `--shadow-panel` and `--shadow-dialog`, so another surface set changes the frames, rules and bars of the whole page.
- Not expressible by a token, written literally in `components.css`: the bevel of buttons and selects is `border-style: outset` and of inputs, checkboxes and pressed buttons `inset`, drawn in `--color-border-muted` and `--color-input-border`. `--border-style` (`solid`) is used for everything else.

## Components

- `.ds-menu`: a jump menu in the sidebar, a `<select>` "Go to a section..." with a `.ds-menu__go` button beside it; use it where a page offers a quick way to a section.
- `.ds-page`: on `<body>`; white page, body type. `.ds-page__frame` is the 600px column, `.ds-page__columns` the main-plus-sidebar grid, `.ds-page__main` the main column with `--space-5` between blocks. On inner pages `.ds-page__head` stacks breadcrumb, page header, tabs and rule, and `.ds-page__inner` is the single full-width column. The page margin is padding inside `.ds-page__frame`, not on `<body>`: when a palette gives `--color-canvas` a colour of its own the frame shows as a sheet on the page and its content keeps that margin from the sheet edge. `.ds-page__frame` takes `--shadow-panel` as well, so a surface set that outlines its boxes also outlines the sheet.
- `.ds-nav`: header. `.ds-nav__top` holds the `.ds-brand`, `.ds-nav__tagline` and `.ds-nav__tools`; `.ds-nav__links` is the centred row of `.ds-nav__item` / `.ds-nav__link`, hyphen separated; the link of the view that is showing is bold, black and not underlined (`components.css` has one selector per view for this, and `is-current` does the same by hand). A link to a section that has no view of its own carries `.ds-nav__link--elsewhere`: it opens the nearest fitting view and is never marked as the view that is showing.
- `.ds-brand`: the site's mark and name, at the left of `.ds-nav__top`. `.ds-brand__mark` is a small square tile dressed like the button (`--fill-button` over `--color-button`, `--color-button-text` for the drawing, the 2px `outset` bevel in `--color-border-muted`, `--radius-control`, `--shadow-control`), so every palette and surface reskins it; it holds the mark as inline SVG painted with `currentColor`, drawn with square line caps, a 3 unit stroke and crisp edges. `.ds-brand__name` is the name in the heading family at `--text-display`, with the heading weight, tracking and transform, in `--color-heading-alt`. The name `chronoskin` and its mark are placeholders for the installing project's own name and logo.
- `.ds-hero`: the search band, the era's page introduction, filled with `--fill-bar` inside a `--border-width` outline. `.ds-hero__title` (bold 13px), `.ds-hero__lead`, `.ds-hero__more` (10px links), `.ds-hero__action` (select, `.ds-hero__field`, button) and `.ds-hero__hint`. Never a tall banner with a large heading.
- `.ds-page-header`: the head of an inner page, used instead of `.ds-hero`. `.ds-page-header__text` holds `.ds-page-header__title` (18px bold sans) and `.ds-page-header__lead` (one 13px serif line); `.ds-page-header__action` holds one button, bottom-aligned at the right. No fill and no border of its own: the tab strip and 2px rule under it close it.
- `.ds-prose`: running text on inner pages; styles `h1`, `h2`, `h3`, `p`, `ul`, `ol`, `strong`, `code`; `.ds-prose__lead` for a 16px lead line.
- `.ds-link`: every link. States `is-visited`, `is-hover`, `is-active`, `is-focus`. `.ds-link--quiet` for promotional links, `.ds-link--strong` for bold links. `.ds-links` wraps a line of links.
- `.ds-button`: grey bevelled button with 13px sans label, on a `<button>` or, when it leads to another page, on a link. `.ds-button--secondary` is a darker grey for the lesser action, `.ds-button--danger` is the destructive action (delete, remove): the same bevel on `--color-surface` with a bold label in `--color-danger`, never a red button, `.ds-button--large` uses 16px. States `is-hover`, `is-pressed` (bevel inverts), `is-focus`, `is-disabled`. `.ds-buttons` lays out a row.
- `.ds-form`: two-column form, right-aligned bold `.ds-form__label` and `.ds-form__field`. Controls: `.ds-form__input` (`--short` for 40px fields), `.ds-form__select`, `.ds-form__textarea`, `.ds-form__check`. `.ds-form__hint` for 10px help, `.ds-form__error` for a bold red message under the field (the control itself does not change), `.ds-form__actions` for the button row. A select is a bevelled key in `--fill-button` with `--color-button-text`; a checkbox is a small inset field in `--color-input` whose tick is a block of `--color-input-text`.
- `.ds-table`: data table with 1px bordered cells and 2px spacing. `.ds-table__head` (teal, centred sans), `.ds-table__cell` with `--num`, `--up`, `--down` and `--group` (pale yellow full-width row), `.ds-table__row--alt` for grey rows, `.ds-table__caption` for the 10px line beneath.
- `.ds-stat`: summary figures as one row of bordered pale yellow cells with 2px spacing inside a 1px frame, built like a table row. Each `.ds-stat__item` holds a `.ds-stat__value` (18px bold sans) over a `.ds-stat__label` (10px sans). Three to five items; never large numerals or separate cards.
- `.ds-list`: bulleted headline list; `.ds-list__item`, `.ds-list__meta` (grey 10px time in brackets), `.ds-list__more` (right-aligned bold 10px link).
- `.ds-directory`: two-column category index: `.ds-directory__title` (16px bold sans link) over `.ds-directory__subs` (13px serif links, comma separated, ending in an ellipsis). The centrepiece of a portal home page.
- `.ds-panel`: a module: `.ds-panel__title` yellow strip with an optional `.ds-panel__tool` link at the right, then `.ds-panel__body`. `.ds-panel--boxed` adds a 1px black outline.
- `.ds-grid`: two-column grid of flat coloured, outlined cells with 2px spacing: `.ds-grid__cell` and `--2`, `--3`, `--4` for the four fills, `.ds-grid__title`, `.ds-grid__text`.
- `.ds-tabs`: the era's tabs, a grey strip of 10px links separated by bars, used for the channels under the search band and for the views of an inner page under its page header; `.ds-tabs__item`, `.ds-tabs__tab`, current `is-current`.
- `.ds-badge`: small solid crimson label. `.ds-badge--text` is the common form, a bold red word beside a link; `.ds-badge--count` is a grey count in brackets. Status labels are the same solid block in a flat status colour, with the word typed in capitals: `.ds-badge--success` (green, white text), `.ds-badge--warning` (yellow, black text), `.ds-badge--danger` (red, white text). `.ds-badges` is a 10px key line that explains them under a table.
- `.ds-sidebar`: right column behind a 1px rule. `.ds-sidebar__date`, then sections of `.ds-sidebar__title` (yellow strip), `.ds-sidebar__list` / `.ds-sidebar__item`, `.ds-sidebar__more`; `.ds-sidebar__lookup` with `.ds-sidebar__field` for a one-field form.
- `.ds-notice`: pale yellow box with a 1px yellow border; `.ds-notice__label` bold lead word, `.ds-notice--center` for a promotional line, `.ds-notice--error` in red. `.ds-notices` stacks several.
- `.ds-pagination`: one centred line: `.ds-pagination__label`, bracketed previous and next, numbered `.ds-pagination__link`, current `is-current`, `is-disabled`.
- `.ds-breadcrumb`: bold sans path separated by ">"; `.ds-breadcrumb__item`, last one `is-current`.
- `.ds-dialog`: a 320px box in the page flow with a 1px black border, black `.ds-dialog__title` bar, centred `.ds-dialog__body` and `.ds-dialog__actions`. No backdrop, no shadow.
- `.ds-empty`: grey block with `.ds-empty__title`, `.ds-empty__text` and a button.
- `.ds-footer`: centred link rows `.ds-footer__links` of `.ds-footer__item` (hyphen separated; `--small` is 10px and bar separated), `.ds-footer__label`, `.ds-footer__legal`.
- `.ds-view`: one screen of the example site; only the one named in the address shows, and `.ds-view--home` shows when none is named. It has no box of its own.
- `.ds-rule`: 2px teal rule; `.ds-rule--thin` 1px grey. `.ds-heading`: section title on a 2px orange underline.

## Never

- `border-radius <= 0px`: no reference has a rounded corner.
- `box-shadow = none`: no reference has a box shadow.
- `text-shadow = none`: no reference has a text shadow.
- `gradient-fills <= 0%`: every fill is a flat colour.
- `border-width <= 2px`: borders are 1px lines or 2px bevels and rules.
- `font-size <= 28px`: the wordmark is the largest text; headings stop at 18px.
- `font-size >= 10px`: small print is 10px, nothing smaller.
- `font-weight >= 400`: no light weights existed.
- `font-families <= 3`: a serif, a sans-serif and a monospace face.
- `underlined-links >= 95%`: links are always underlined; only a current item loses the underline.
- `letter-spacing = 0`: no tracking anywhere.
- `uppercase-text <= 0%`: no text-transform; capitals are typed.
- `row-gap <= 8px`: list and table rows sit close: 2px between table rows, at most 8px between directory entries.
- `block-gap <= 16px`: neighbouring blocks are at most 16px apart.
- `content-width <= 600px`: the page is a fixed 600px column.
- `transition = none`: nothing animates on hover.
- `animation = none`: no CSS animation.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new module is a `.ds-panel` with a yellow title strip; a new index is a `.ds-directory` or a `.ds-list`; a new band is a flat fill of `--fill-bar` or one of `--color-fill-*` at full column width; anything tabular is a `.ds-table` with 2px spacing. Keep text at 13px or 10px, keep links underlined and blue, separate with 1px or 2px rules, and never add radius, shadow, gradient, hover effects or more whitespace than `--space-5`.
