# Style pack format v1

A style pack is five text files in a `.design/` directory. This document is the frozen contract for them. `bin/pack lint <dir>` checks the rules below that a script can check (the vocabulary, the token-only rules, what a file may load or contain, the checklist); the vocabularies live in `internal/pack/vocab.go`.

| File | Owner part | Purpose |
| --- | --- | --- |
| `style.json` | server | Manifest |
| `STYLE.md` | all four | Rules for the agent |
| `tokens.css` | palette, type, surface | Every colour, font, size, border, radius and shadow |
| `components.css` | structure | Layout tokens and component classes, built only on tokens |
| `specimen.html` | structure | One page that shows every component |

No file makes a network request. A pack ships no images and no fonts.

## The idea: fixed names, free values

Every pack defines the same token names and the same component classes. Only the values and the CSS behind the classes differ. This is what lets mix mode take the structure of one style and the palette, type and surface of others, and what lets a script check an agent's CSS against the pack.

## style.json

```json
{
  "id": "v1.bulletin-board-forums.forum.s17.p03.t11.u08",
  "library": "2026.10",
  "era": { "slug": "bulletin-board-forums", "years": [2002, 2007] },
  "archetype": "forum",
  "mode": "mix",
  "markup": "modern",
  "viewport": { "width": 1024, "fluid": false },
  "source": "https://<host>/s/v1-bb-h3b8",
  "short": "v1-bb-h3b8"
}
```

- `id` is `v1.<era-slug>.<archetype>.sNN.pNN.tNN.uNN`: structure, palette, type and surface numbers within the era. In `pure` mode the four numbers are equal. Two adjustments may follow: `.dc` or `.dr` (compact or roomy spacing) and `.c<text>` (colours set by hand), in that order.
- `short` is the same ID in the form people pass around: `v1-<era code>-<four digits>`, where the era code is two characters fixed for the library and each digit is one part number in base 36 (1 to 9, then a to z). The adjustments follow as `-dc` or `-dr` and `-c<text>`. Both forms name the same style everywhere an ID is accepted.
- `density` (`compact` or `roomy`) is present only when the spacing was adjusted. Spacing scales every `--space-*` token of `components.css` (by 0.8 or 1.25) and moves the Never rules about gaps with it.
- `colours` is present only when palette colours were set by hand: an object of token name and colour (`#rrggbb` or `rgba()`), replacing those tokens in `tokens.css`. In the ID they are written as the text after `c`: for each token set, in vocabulary order, one byte for its position in the palette (with the top bit set when the colour is translucent), three bytes of red, green and blue and, when translucent, one of alpha; the bytes in base 32 with the alphabet a to z, 2 to 7, without padding. A colour equal to the palette's own is not written, so a style has one ID.
- `archetype` is one of `forum`, `blog`, `portal`, `shop`, `landing`, `app`, `docs`.
- `mode` is `pure` or `mix`. `markup` is `modern`.
- `viewport.width` is the width the style was designed for; `fluid` says whether the page stretches beyond it. `viewport.responsive` (optional, `true`) says the layout rearranges itself down to a phone 390px wide; `bin/pack views <dir>` then also clicks through the specimen at that width. Leave it out for a style of the years before phones, whose pages had one width.
- `source` ends with `/s/` and the ID in either form.
- No other fields.

## tokens.css

Three `:root` blocks, each introduced by a marker comment, in this order. Each block defines exactly the tokens of its part: none missing, none added.

```css
/* @part palette */
:root { --color-page: #e5e5e5; ... }

/* @part type */
:root { --font-body: Verdana, Arial, Helvetica, sans-serif; ... }

/* @part surface */
:root { --border-width: 1px; ... }
```

Raw colours appear only in the palette part. The type and surface parts refer to colours through `var(--color-*)`; the surface part may also use `color-mix()` with `white`, `black` or `transparent` to derive gradient stops, and `data:` URIs for patterns.

### Palette

| Token | Role |
| --- | --- |
| `--color-page` | Page background, outside the content |
| `--color-canvas` | Content frame or sheet that sits on the page and holds the surfaces. Equal to `--color-page` when the style has no such frame |
| `--color-surface` | Background of panels, table cells, list rows |
| `--color-surface-alt` | Alternate rows, secondary panels |
| `--color-surface-strong` | Third, strongest surface level: group rows, panel title strips |
| `--color-bar` | Primary bar: site navigation, table head |
| `--color-bar-text` | Text on the primary bar |
| `--color-bar-alt` | Secondary bar: sub-navigation, toolbar, footer strip |
| `--color-bar-alt-text` | Text on the secondary bar |
| `--color-inverse` | Inverted block: a dark footer or band on a light style, or the reverse |
| `--color-inverse-text` | Text on an inverted block |
| `--color-text` | Body text |
| `--color-text-muted` | Secondary text: dates, counts, captions |
| `--color-heading` | Headings |
| `--color-heading-alt` | Second heading colour: panel and sidebar titles |
| `--color-link` | Link |
| `--color-link-quiet` | Secondary links: tools, meta, navigation inside content |
| `--color-link-visited` | Visited link |
| `--color-link-hover` | Hovered link |
| `--color-link-active` | Active link |
| `--color-border` | Ordinary borders and rules |
| `--color-border-strong` | Outer or emphasised borders |
| `--color-border-muted` | Faint dividers inside a box |
| `--color-accent` | Highlights: current item, badges |
| `--color-accent-text` | Text on the accent colour |
| `--color-accent-alt` | Second highlight: emphasis that is not an error, "new" markers |
| `--color-fill-1`, `--color-fill-2`, `--color-fill-3`, `--color-fill-4` | Decorative block fills for sections, cards and grid cells. A style with fewer block colours repeats its surface colours here |
| `--color-button` | Primary button |
| `--color-button-text` | Text on the primary button |
| `--color-button-hover-text` | Text on the hovered primary button |
| `--color-button-secondary` | Secondary button |
| `--color-button-secondary-text` | Text on the secondary button |
| `--color-input` | Form control background |
| `--color-input-text` | Form control text |
| `--color-input-border` | Form control border |
| `--color-disabled` | Disabled control background |
| `--color-disabled-text` | Disabled control text |
| `--color-danger` | Error text and error borders |
| `--color-danger-surface` | Background of an error message |
| `--color-success` | Positive status: returned, paid, online. The era's green, or its nearest colour |
| `--color-warning` | Cautionary status: due soon, pending. The era's amber, or its nearest colour |
| `--color-notice` | Notice background |
| `--color-notice-text` | Notice text |
| `--color-focus` | Focus indicator |
| `--color-shadow` | Colour of every shadow |
| `--color-overlay` | Backdrop behind a dialog |

### Type

| Token | Role |
| --- | --- |
| `--font-body` | Body font stack |
| `--font-heading` | Heading font stack |
| `--font-ui` | Buttons, inputs, navigation |
| `--font-mono` | Code |
| `--text-base` | Body size |
| `--text-small` | Secondary text size |
| `--text-ui` | Size in buttons, inputs, navigation |
| `--text-large` | Lead paragraphs, large buttons |
| `--text-display` | Hero heading or wordmark. Equal to `--text-h1` when the style has no display size |
| `--text-h1`, `--text-h2`, `--text-h3` | Heading sizes |
| `--line-body`, `--line-heading`, `--line-display` | Line heights |
| `--weight-body`, `--weight-bold`, `--weight-heading`, `--weight-display`, `--weight-ui` | Weights |
| `--heading-transform`, `--heading-tracking` | `text-transform` and `letter-spacing` of headings |
| `--display-tracking` | `letter-spacing` of display text |
| `--figure-shift` | How far a figure set alone in a small box (a step number, a score, a count) is lifted to sit in its middle, in `em`; `0em` for most faces, a little more for one whose figures hang below the line, such as Georgia's old-style figures |
| `--ui-transform` | `text-transform` of buttons and navigation |
| `--link-decoration`, `--link-decoration-hover` | `text-decoration` of links in running text |
| `--link-decoration-quiet` | `text-decoration` of quiet links, titles and navigation |

Font stacks name system fonts and openly licensed fonts only.

### Surface

| Token | Role |
| --- | --- |
| `--border-width`, `--border-width-strong` | Border widths |
| `--border-style` | Border style |
| `--radius-control` | Radius of buttons and inputs |
| `--radius-panel` | Radius of panels, tables, dialogs |
| `--radius-pill` | Radius of badges and tags |
| `--radius-page` | Radius of the largest containers: content sheet, hero, grid cells |
| `--shadow-panel` | `box-shadow` of panels |
| `--shadow-control` | `box-shadow` of buttons and inputs |
| `--shadow-control-hover` | `box-shadow` of a hovered button |
| `--shadow-control-pressed` | `box-shadow` of a pressed button |
| `--shadow-dialog` | `box-shadow` of dialogs |
| `--shadow-text` | `text-shadow` on bars and buttons |
| `--fill-page` | `background` of the page, including any pattern |
| `--fill-bar` | `background` of the primary bar |
| `--fill-bar-alt` | `background` of the secondary bar |
| `--fill-inverse` | `background` of inverted blocks |
| `--fill-accent` | `background` of accent elements such as badges |
| `--fill-panel` | `background` of panels |
| `--fill-button`, `--fill-button-hover` | `background` of the primary button |
| `--fill-button-secondary` | `background` of the secondary button |
| `--fill-input` | `background` of form controls |
| `--focus-ring` | `outline` of a focused control |
| `--transition` | `transition` of interactive elements; `none` in eras without them |
| `--backdrop-blur` | How strongly a translucent surface blurs what lies behind it; `0px` in a style without glass |

A style that has no shadow sets the token to `none`; no radius is `0`. The token is still defined.

## components.css

1. The file starts with one `:root` block of **layout tokens**. Their names begin with `--size-` or `--space-`. `--size-page` (the content width) is required; the rest are the style's own spacing scale. These belong to the structure part, so they are not in `tokens.css`.
2. Every other rule is scoped by a `.ds-` class. No bare element selectors: the file must be safe to load into an existing project.
3. Classes are named `ds-<component>`, `ds-<component>__<element>`, `ds-<component>--<variant>`. State classes are `is-<state>` and are only used together with a `ds-` class.
4. Outside `:root`, declarations use tokens only:
   - no raw colours and no colour keywords except `transparent`, `currentColor` and `inherit`;
   - no raw lengths in `px`, `em`, `rem` or `pt` except `0` (percentages, `ch` and `calc()` over tokens are fine);
   - `font-family` through a `--font-*` token; no `font` shorthand;
   - no `url()`, no custom property definitions.

   Still allowed: unitless numbers (`opacity`, `line-height`, `font-weight`), percentages, `ch`, `vw` and `vh`, `calc()` over tokens, and `color-mix()` whose colours are all tokens or `transparent`. Use these for a one-off shade or size the vocabulary has no token for; say so in STYLE.md.
5. No `@import`, no `@font-face`. `@media` and `@keyframes` are allowed.
6. Interactive states are written twice, as a pseudo-class and as a state class (`.ds-link:hover, .ds-link.is-hover`), so the specimen can show them statically.

## specimen.html

- Starts with `<!doctype html>`. Loads exactly two stylesheets with relative links, `tokens.css` then `components.css`. The server inlines both when it serves the specimen on its own.
- No `<style>`, no `style=""`, no `<script>`, no `<img>` or other embedded media. Inline `<svg>` is allowed; it paints only with `currentColor` or `none`, so its colour comes from a token.
- Uses only classes defined in `components.css`. `<body>` carries `ds-page`.
- Is one small working example site with several **views** (screens), of which one shows at a time; see Views below. A component may appear in more than one view, and only one instance carries the `data-component` marker.
- Contains every component of the checklist below, whatever the reference page showed. Each is marked with `data-component="<id>"` on its outermost element, or on a wrapper when the checklist asks for several instances (links, notices, buttons).
- Placeholder content is neutral text that fits the archetype. Nothing is copied from the reference: no text, names, logos or images.
- The site in every specimen is called `chronoskin`, in the title, the `brand` component, the footer and the copy. No other product or site name is invented. The brand is written like this:

```html
<a class="ds-brand" data-component="brand" href="#">
  <span class="ds-brand__mark"><svg viewBox="0 0 24 24" fill="currentColor" aria-hidden="true">...</svg></span>
  <span class="ds-brand__name">chronoskin</span>
</a>
```

  The name is always `chronoskin`. The mark is the era's own: one drawing shared by every layout of an era, in the manner of the logos of its years (a pixel icon, a glossy badge, a folded sheet, a prompt sign), so that a specimen shows what a logo of that period looked like. It is a single inline `<svg>` with `viewBox="0 0 24 24"`, built from plain shapes, coloured only through `currentColor` (two tones through `opacity`), with no gradients, filters, text or classes, and it must read at 16px. A new layout takes the mark of its era's first pack; a new era brings a mark of its own. The name and mark are placeholders for the installing project's own: STYLE.md says so in its Components section.

### Checklist

| `data-component` | Base class | Must show |
| --- | --- | --- |
| `shell` | `.ds-page` | Page container at the style's width |
| `nav` | `.ds-nav` | Site header and navigation with a current item |
| `hero` | `.ds-hero` | Page introduction in the era's form: title, a line of lead text and the main action |
| `brand` | `.ds-brand` | The site's mark and name, in the header: `.ds-brand__mark` holds the chronoskin mark as inline SVG, `.ds-brand__name` holds the name. The mark sits on a small tile dressed from tokens the way the style dresses a primary button (fill, text colour, border, radius, shadow), so every palette and surface reskins it |
| `page-header` | `.ds-page-header` | The head of an inner page that has no hero: page title, one line of description, an action on the right |
| `typography` | `.ds-prose` | h1, h2, h3, paragraphs, a list, inline code |
| `links` | `.ds-link` | Normal, visited, hover and active |
| `buttons` | `.ds-button` | Primary, secondary (`.ds-button--secondary`), destructive (`.ds-button--danger`), disabled |
| `form` | `.ds-form` | Text input, select, checkbox, textarea, a field with a validation error |
| `table` | `.ds-table` | Header row, at least four rows, a numeric column |
| `list` | `.ds-list` | Item or post list, at least three items with meta text |
| `panel` | `.ds-panel` | Panel or card with a title and a body |
| `stats` | `.ds-stat` | A row of at least three summary figures, each a number with a label, in the era's form. `.ds-stat` may be the row or the single figure; STYLE.md says which and names the other class |
| `grid` | `.ds-grid` | Grid of at least four cards or cells, using the `--color-fill-*` tokens where the style has block colours |
| `tabs` | `.ds-tabs` | Tab row or its era equivalent, with a current tab |
| `badge` | `.ds-badge` | Badge, tag or count, and the three status variants: `.ds-badge--success`, `.ds-badge--warning`, `.ds-badge--danger` |
| `sidebar` | `.ds-sidebar` | A sidebar block with a heading and links |
| `notice` | `.ds-notice` | An information notice and an error notice |
| `pagination` | `.ds-pagination` | Previous, numbered pages with a current page, next |
| `breadcrumb` | `.ds-breadcrumb` | Three levels |
| `dialog` | `.ds-dialog` | A dialog or the era's equivalent, shown open in the page flow |
| `empty` | `.ds-empty` | Empty state with a message and an action |
| `footer` | `.ds-footer` | Footer |

The variant classes named above are required in every pack: an application needs a destructive action and status labels whatever the era, and an era that had no such thing still gets one built from its own borders, fills and type.

### Views

The specimen is not a sheet of components and not two pages stacked: it is an example site that a visitor can click through. It has between three and six views, each a screen that site would really have: a home view, then inner pages such as a list, a detail or thread, a form or settings screen. What the site is about decides the views (a forum has its board index, a topic list and a thread; a product page has its landing, pricing and sign-up).

- The header, navigation and footer are written once, outside the views. Each view is `<section class="ds-view" id="<name>">`, with the class first and a lowercase id; the first is the home view and carries `class="ds-view ds-view--home"`.
- Navigation is by plain links: `href="#<name>"`. Every view is reachable from the specimen's own navigation, and links inside views lead to other views where a real site would (a row to its detail, a "new" button to the form, a breadcrumb back).
- `components.css` contains these rules as written, so that exactly one view shows, the home view when the address names none, and a newly opened view starts at the top:

```css
.ds-view { display: none; }
.ds-view:target, .ds-page:not(:has(.ds-view:target)) .ds-view--home { display: block; }
.ds-view:target { scroll-margin-top: 100vh; }
```

  A visible view is always `display: block`; an inner wrapper does its layout. The navigation marks the current view with one rule per view, for example `.ds-page:has(#topics:target) .ds-nav__link[href="#topics"]`, and marks the home link when no view is named.
- Inner views start as that era started a page below the home page: breadcrumb, `page-header`, then the working components (tabs, table, form), without the hero.
- Every checklist component sits in the view where the example site would use it. A dialog shows in the view that would open it, an empty state where a list could be empty. No view is a dump of leftover components.
- Each view fills its screen sensibly on its own: columns end near each other, and nothing is left as a tall empty area.

A pack may add more `ds-` classes for things its era needs, such as a category row in a forum.

## STYLE.md

The first line is `# <style name>`. Then exactly these level-2 sections, in order:

1. `## Summary`. Three sentences on what the style is and when it was common.
2. `## Layout`. Page width, columns, navigation position, the spacing scale and which `--space-*` token goes where. It also says how **inner pages** are laid out: what replaces the hero, which columns remain, where tabs and the page header go.
3. `## Typography and colour roles`. Which token is used where.
4. `## Components`. One entry per class, naming the class in backticks (`` `.ds-button` ``), when to use it and its variants. Every base class of the checklist must appear.
5. `## Never`. Machine-checkable rules, at least three.
6. `## Extending`. A new component is derived from the nearest specimen component and uses tokens only.

### Never rules

Each rule is one bullet of the form

```
- `<metric> <op> <value>`: <explanation>
```

where `<op>` is `<=`, `>=`, `=` or `!=`, and `<value>` is a number, a number with `px` or `%`, or `none`. The explanation follows the colon.

| Metric | Measures |
| --- | --- |
| `border-radius` | Corner radius of any box. Pills and circles (rounded to half their shorter side) are shapes and do not count |
| `border-width` | Width of any border |
| `box-shadow` | Presence of box shadows (`= none`) |
| `box-shadow-blur` | Blur radius of any box shadow |
| `text-shadow` | Presence of text shadows (`= none`) |
| `gradient-fills` | Share of boxes with a gradient background, in % |
| `font-size` | Any font size |
| `font-weight` | Any font weight |
| `font-families` | Number of distinct font families in use |
| `line-height` | Line height of body text, as a ratio |
| `underlined-links` | Share of link text that is underlined, in % |
| `letter-spacing` | Letter spacing of any text |
| `uppercase-text` | Share of text set in uppercase by `text-transform`, in % |
| `row-gap` | Gap between consecutive rows of a table or list; an upper bound is checked against the 90th percentile, a lower bound against the 10th |
| `block-gap` | Gap between consecutive wide blocks that share a parent, checked like `row-gap` |
| `content-width` | Width of the span that holds the middle 90% of the text, in px, or in % of the viewport |
| `palette-colours` | Number of distinct colours painted |
| `transition` | Presence of CSS transitions (`= none`) |
| `animation` | Presence of CSS animations (`= none`) |

Rules say what is absent in the era, as numbers a tool can measure. `bin/pack adhere -pack <dir> <dir>/specimen.html` checks them; a rule the pack's own specimen does not satisfy is reported as unusable and must be fixed or removed.
