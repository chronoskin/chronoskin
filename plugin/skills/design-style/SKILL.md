---
name: design-style
description: Use for any UI work (HTML, CSS, components, pages, layouts). Makes the UI follow the style pack in .design/ instead of default styling.
---

Before writing any UI:

1. If `.design/` exists at the repository root, read `.design/STYLE.md` and `.design/specimen.html` in full. The specimen is the reference: when in doubt, do what it does. Its site name and mark (`chronoskin`) are placeholders: use the project's own name and logo in the `.ds-brand` slot, dressed the same way.
2. Load `.design/tokens.css` and then `.design/components.css` in every page, before any stylesheet of your own. Do not copy their contents into your own files and do not edit them.
3. Build pages from the pack's component classes, with the markup the specimen uses for each component. Lay the page out the way the specimen lays its page out.
4. In your own CSS use only the pack's tokens: every colour, font, font size, radius, shadow and spacing is a `var(--…)` from the pack. No raw colours, no new fonts, no new sizes.
5. Obey every rule in the "Never" section of `STYLE.md`. They are measured limits, not suggestions.
6. A component the pack lacks is derived from the nearest specimen component and uses tokens only.
7. Do not improve the style. Small text, tight rows, underlined links, square corners and strong colours are the style; keep them even where you would normally choose otherwise.
8. Before finishing, run the `lint_css` tool on the CSS you wrote, with the ID from `.design/style.json`, and fix every violation.

If `.design/` does not exist, ask which era to use, or call `generate_style` and then apply the result with `/chronoskin:design-apply`.
