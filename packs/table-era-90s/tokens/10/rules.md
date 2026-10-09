## Colour roles

- Family: the dark home page of 1998, read from the 1998 community welcome page and the 1998 film site: a black page with white text, pure yellow links, lime, pure red and pure blue, and the greys of browser chrome. Nothing is tinted; every colour is a named colour of the period at full strength.
- `--color-page`, `--color-canvas` and `--color-surface` are black `#000000`. `--color-surface-alt` `#181818` (the darkest grey read from the screenshots) is alternate rows, quiet strips and the empty state; `--color-surface-strong` `#686868`.
- `--color-bar` `#0000ff` (pure blue, white text, yellow links) is the search band, the table head, panel titles and the header and footer rules. `--color-bar-alt` `#686868` (grey, white text, yellow links) is every other title strip and the notice border.
- `--color-inverse` `#c0c0c0` (silver) with black text is the light block of a dark page: a dialog title bar or a top bar. Links placed on it take its text colour.
- Text is white `--color-text`; `--color-text-muted` `#989898` is for times, counts and legal lines. `--color-heading` is white; `--color-heading-alt` `#00ff00` (lime) is the brand name, h3 and positive figures.
- Links are `--color-link` `#ffff00` (yellow), always underlined; `--color-link-quiet` is lime `#00ff00`, visited is silver `#c0c0c0`, active is red `#ff0000`; hover equals the link colour.
- `--color-accent` `#ffff00` carries black text (`--color-accent-text`): the solid badge is a yellow block. `--color-accent-alt` `#ff0000` is the "New!" word, prices and negative figures.
- `--color-fill-1` to `--color-fill-4` are `#181818`, `#000000`, `#0000ff`, `#686868`. They carry white text and yellow links.
- `--color-danger` `#ff0000` is error text on black and a fill under black text; `--color-success` `#00ff00` is a fill under black text; `--color-warning` `#ff5000` (the orange of the 1998 university page, the only amber in the new references) is a fill under white text. `--color-danger-surface` is black; notices sit on `#181818`.
- Controls are the grey of the operating system, as on any page of the period whatever its background: `--color-button` `#c0c0c0` and `--color-button-secondary` `#989898` with black text, white form fields with black text (`--color-input`, `--color-input-text`) in a silver bevel. Disabled controls are `#686868` with silver text. The two dark references show no form controls; these are the browser defaults read from the other captures.
- Borders: `--color-border` and `--color-border-muted` `#989898` for outlines, row rules and bevels, `--color-border-strong` `#ffffff` (the white outline the film site puts round its pictures). `--color-focus` is yellow.

## Typography roles

- One serif throughout, the browser default nobody changed: `--font-body`, `--font-heading` and `--font-ui` are Times New Roman, Times. `--font-mono` is Courier New.
- `--text-base` 16px; `--text-small` and `--text-ui` 13px, the smaller step of the film site, also used for controls and navigation.
- Headings are large and bold as on the community page: `--text-h3` 16px, `--text-h2` and `--text-large` 19px, `--text-h1` 24px, `--text-display` 32px for the brand name only. All bold is 700.
- Line height `normal`; no transform, no tracking. All three link decorations are `underline`.

## Surface roles

- Tables are framed the way `border=2` drew them: `--border-style` is `ridge`, `--border-width` 2px for cells, outlines and rules, `--border-width-strong` 4px for the header and footer rules and the heavy bevel of controls.
- `--fill-page` tiles a sparse field of grey points over `--color-page`, the starfield of a 1998 home page, which shows in the window beside the page column. `--fill-bar` fades from the bar colour at the left to a darker mix of it at the right. The sheet, panels, controls and the message box carry a 1px outline in `--color-border-strong` (`--shadow-panel`, `--shadow-control`, `--shadow-control-hover`, `--shadow-dialog`, a spread with no blur); the other shadows are `none`. Radius 0; every other fill is the flat palette colour and `--fill-panel` is `--color-surface`.
- `--focus-ring` is `1px solid` in `--color-focus`; `--transition` is `none`.

## Never

- `border-radius <= 0px`: no reference has a rounded corner.
- `box-shadow-blur <= 0px`: the only shadows are 1px outlines round the sheet, panels and controls; nothing is blurred.
- `text-shadow = none`: no reference has a text shadow.
- `gradient-fills <= 6%`: only the primary bar fades; every other fill is flat.
- `border-width <= 4px`: frames are 2px ridges; the widest border is the 4px bevel of a control.
- `transition = none`: nothing in the references changes on hover.
- `animation = none`: no reference has a CSS animation.
- `font-size >= 13px`: the smallest text measured on the two dark pages is 13px.
- `font-size <= 32px`: the brand name is the largest text; headings stop at 24px.
- `font-families <= 2`: Times and Courier.
- `font-weight >= 400`: only regular and bold were measured.
- `underlined-links >= 95%`: both references underline 100% of their link text.
- `letter-spacing = 0`: no reference tracks its text.
- `uppercase-text <= 0%`: no reference uses text-transform; capitals are typed.
