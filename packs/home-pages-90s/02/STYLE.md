# Personal home page: the family page, 1999

## Summary

A family or hobby home page as built at home between 1996 and 2001: a framed sheet centred on a pastel gingham tile, the family's name in a large playful face, a striped bar of links, a welcome box and a two-column table of "my pages". Text is dark purple in a rounded hand-drawn face on pink and white, headings are hot pink and teal, links are default blue and underlined, and every box has a double border with a hard offset shadow. Counters, "NEW!" marks, awards, web rings and an e-mail form are its furniture; it was common on free hosts until page builders replaced it after 2001.

## Layout

- Fixed, not fluid: `--size-page` is 640px. `.ds-page__frame` is a sheet in `--color-canvas` with a `--border-width-strong` border in `--color-border-strong` and `--shadow-panel`, centred with `--space-5` above and below, so the `--fill-page` tile shows 80px wide on each side of an 800px window.
- The specimen is one example site, a family's home page, of five views. `.ds-nav` and `.ds-footer` are written once, inside the sheet and outside the views.
- Home view: `.ds-nav` (`.ds-nav__top` with the `.ds-brand` at the left and the italic `.ds-nav__tagline` right-aligned beside it, then the `.ds-nav__links` bar across the sheet), then `.ds-page__main`: the centred `.ds-hero` welcome box, the `.ds-stat` row of three counters, and `.ds-page__columns`: a `.ds-page__column` (a `.ds-heading` over the two-column `.ds-grid` of pages, a heading over the bulleted `.ds-list` of news, a `.ds-panel`) and a `--size-side` (190px) `.ds-sidebar` of small boxes at the right, which ends near the foot of the main column.
- Inner pages keep header and footer and drop the welcome box and counters. `.ds-page__inner` opens with `.ds-page__head`: the `.ds-breadcrumb`, the `.ds-page-header` (title and italic line at the left, a button at the right, a thick rule under) and, where the page has parts, the `.ds-tabs` row of keys. Below it the page is one full-width column (table and page numbers, awards and `.ds-page__pair` boxes, or notices and form); the page of running text keeps `.ds-page__columns` with the sidebar.
- Views: `home`; `about` (running text, sidebar, link states); `recipes` (tabs, the recipe table with its badge key, page numbers, the empty state of a part with no recipes); `awards` (award badges, two web ring boxes, a bulleted list of rules); `mail` (notices, the form, the message box for leaving the mailing list).
- Spacing: `--space-1` 3px strip and cell padding; `--space-2` 6px between grid cells, tabs and page numbers; `--space-3` 10px box padding, gaps between buttons and form rows; `--space-4` 14px between sidebar boxes, paragraph spacing; `--space-5` 20px sheet padding, the gap between blocks and columns; `--space-6` 32px list indent.
- Fixed sizes: `--size-label` 140px, `--size-field` 230px, `--size-area` 330px, `--size-check` 16px, `--size-dialog` 360px, `--size-icon` 34px, `--size-award` 128px.

## Typography and colour roles

The conventions for which text sits on which fill are those of the era's first layout: text, heading, link and status colours are legible on page, canvas, surfaces and the four fills; bars, the inverse, the accent and notices bring their own text colour; status colours as fills are lettered in `--color-canvas`.

- `--font-body`, `--font-heading` and `--font-ui` are one rounded hand-drawn stack (Comic Sans MS, then Chalkboard SE and Comic Neue). Body 14px at line height 1.35; small print 11px; controls 12px bold; headings 16, 21 and 28px bold; the name 40px with 1px tracking. `--font-mono` (Lucida Console) sets counter digits, code and the textarea. Links are always underlined.
- `--color-page` pink `#ffd6ea`, `--color-canvas` `#fff0f7`, `--color-surface` white, `--color-surface-alt` sky `#d6efff` (alternate rows, the empty state), `--color-surface-strong` lilac `#e6d9ff` (the ground of the counters and of the grid).
- `--color-text` `#330066`; `--color-heading` hot pink `#cc0099` (name, titles, section headings, icons); `--color-heading-alt` teal `#008080` (tagline, lead lines, h2, form labels). Links `#0000ff`, visited `#800080`, hover `#ff0066`.
- `--fill-bar` purple with white text is the link bar, table heads, the current page number and the dialog title; `--fill-bar-alt` teal with white text is every title strip and the current tab; `--fill-inverse` deep purple with pale yellow text is counter digits and awards; `--fill-accent` pink with white text is the current link in the bar and the "NEW!" badge. `--color-fill-1` to `-4` (pink, sky, mint, lemon) fill the cells of the grid and carry body text and links.
- The primary button is yellow `--color-button` with purple text; the secondary is sky blue.
- Surface: `--border-style` `double`, 3px for boxes and cells, 5px for the sheet, the welcome box, heading underlines and awards. Boxes cast a hard 4px offset shadow in `--color-shadow` (`--shadow-panel`), buttons 2px, the welcome box and dialog 6px; nothing is blurred. Bars are candy-striped at 45 degrees, the inverse is pin-striped, buttons are a two-step fill; `--fill-page` is a 32px gingham of `--color-surface-alt` over `--color-page`. Radius 0, no text shadow, no transitions.
- Written literally: `outset` and `inset` bevels on buttons, tabs and controls, italic for taglines and lead lines, the round bullets of the news list (`border-radius: 50%`, a shape), the 1px-scale `letter-spacing` of counter digits (`--space-1`), and `.ds-blink`.

## Components

- `.ds-page`: on `<body>`; tile and body type. `.ds-page__frame` the sheet; `.ds-page__main`, `.ds-page__inner` stack blocks; `.ds-page__head`; `.ds-page__columns` with `.ds-page__column` and the sidebar; `.ds-page__pair` two equal boxes; `.ds-section` a heading with its block.
- `.ds-nav`: `.ds-nav__top` (brand and `.ds-nav__tagline`) over `.ds-nav__links`, a bar in `--fill-bar` of `.ds-nav__item` / `.ds-nav__link` with a diamond between them; the current view's link takes `--fill-accent`.
- `.ds-brand`: `.ds-brand__mark`, a tile dressed like the primary button, and `.ds-brand__name` at `--text-display` in `--color-heading`. Name and mark are placeholders for the installing project's own.
- `.ds-hero`: the centred welcome box with `--shadow-dialog`: `.ds-hero__title`, `.ds-hero__lead`, `.ds-hero__action`, `.ds-hero__hint`.
- `.ds-page-header`: `.ds-page-header__text` (`.ds-page-header__title`, `.ds-page-header__lead`) and `.ds-page-header__action`, closed by a thick rule.
- `.ds-heading`: a left-aligned section title with a thick underline.
- `.ds-rule`, `.ds-rule--short`, `.ds-rule--rainbow`: the carved rule and the six-step colour bar that opens the footer.
- `.ds-construction` with `.ds-construction__text`: the striped strip. `.ds-blink`: the slow blink for one small mark.
- `.ds-prose`: running text (`h1` to `h3`, `p`, lists, `strong`, `code`), `.ds-prose__lead`, `.ds-prose__code`.
- `.ds-link`: states `is-visited`, `is-hover`, `is-active`, `is-focus`; `.ds-link--quiet`, `.ds-link--strong`, `.ds-link--plain`; `.ds-links`; `.ds-icon`.
- `.ds-bullets`: a list with coloured diamonds, `.ds-bullets__item`, `.ds-bullets__text`.
- `.ds-button`: `.ds-button--secondary`, `.ds-button--danger`, `.ds-button--large`; states `is-hover`, `is-pressed`, `is-focus`, `is-disabled`; `.ds-buttons`.
- `.ds-form`: `.ds-form__label`, `.ds-form__field`, `.ds-form__input` (`--short`, `is-error`), `.ds-form__select`, `.ds-form__textarea`, `.ds-form__check`, `.ds-form__hint`, `.ds-form__error`, `.ds-form__actions`.
- `.ds-table`: collapsed cells, each with a double border: `.ds-table__head` (`--num`), `.ds-table__cell` (`--num`, `--group` on `--color-fill-4`), `.ds-table__row--alt`, `.ds-table__desc`, `.ds-table__caption`.
- `.ds-stat`: the row; three `.ds-stat__item` cells, each a `.ds-stat__value` (six digits on `--fill-inverse`) over a `.ds-stat__label`.
- `.ds-list`: news with round coloured bullets: `.ds-list__item`, `.ds-list__body`, `.ds-list__meta` (the date in brackets), `.ds-list__more`; `.ds-list__title` a bold coloured name.
- `.ds-panel`: `.ds-panel__title` strip, `.ds-panel__body`, `.ds-panel__text`, `.ds-panel__form`. `.ds-ring` / `.ds-ring__item`: the web ring line.
- `.ds-grid`: the two-column table of pages on `--color-surface-strong`: `.ds-grid__cell` (`--2`, `--3`, `--4`) with a `.ds-grid__icon`, `.ds-grid__title` and `.ds-grid__text`.
- `.ds-tabs`: a row of bevelled keys, `.ds-tabs__item` / `.ds-tabs__tab`; `is-current` is pressed in and takes `--fill-bar-alt`.
- `.ds-badge`: `.ds-badge--text`, `.ds-badge--count`, `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger`; `.ds-badges` the key line.
- `.ds-awards`, `.ds-award` (`.ds-award__star`, `.ds-award__title`, `.ds-award__text`): awards as badges.
- `.ds-sidebar`: `.ds-sidebar__block` boxes with a `.ds-sidebar__title` strip and a `.ds-sidebar__body` or a `.ds-sidebar__list` of `.ds-sidebar__item`.
- `.ds-notice`: `.ds-notice__label`, `.ds-notice--error`; `.ds-notices`.
- `.ds-pagination`: small outlined squares: `.ds-pagination__label`, `.ds-pagination__link` (`is-current`, `is-disabled`).
- `.ds-breadcrumb`: `.ds-breadcrumb__item`, the last `is-current`.
- `.ds-dialog`: `.ds-dialog__title`, `.ds-dialog__body`, `.ds-dialog__actions`, in the page flow.
- `.ds-empty`: `.ds-empty__title`, `.ds-empty__text` and a button.
- `.ds-footer`: rainbow bar, bracketed `.ds-footer__links` / `.ds-footer__item`, `.ds-footer__mail`, `.ds-footer__note`, `.ds-footer__legal`.
- `.ds-view`: one screen; `.ds-view--home` shows by default.

## Never

- `border-radius <= 0px`: boxes are square; only the round bullets are shapes.
- `box-shadow-blur <= 0px`: shadows are hard offsets, never soft.
- `text-shadow = none`: text is flat.
- `border-width <= 5px`: the thickest border is the 5px double line.
- `font-size <= 40px`: the name is the largest text.
- `font-size >= 11px`: small print is 11px.
- `font-weight >= 400`: regular and bold only.
- `font-families <= 2`: the hand-drawn face and one monospace.
- `uppercase-text <= 0%`: capitals are typed.
- `content-width <= 640px`: everything is inside the 640px sheet.
- `transition = none`: nothing eases.

## Extending

Derive a new component from the nearest one in the specimen and use tokens only. A new block in the main column is a `.ds-heading` over a bordered box; a new small box goes in the `.ds-sidebar`; a new page link is a `.ds-grid__cell`. Keep the sheet framed on the tile, the text dark on pastel, links blue and underlined, borders double with hard shadows, and never add radius, blur, transitions or grey.
