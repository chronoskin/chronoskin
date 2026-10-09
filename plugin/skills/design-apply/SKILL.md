---
name: design-apply
description: Install a style pack into .design/ by Style ID.
argument-hint: "<style-id>"
---

Style ID: $ARGUMENTS

1. Call `get_style` with the Style ID.
2. Write the five returned files into `.design/` at the repository root: `style.json`, `STYLE.md`, `tokens.css`, `components.css`, `specimen.html`. Write them unchanged; the `FILE:` line that starts each block is not part of the file.
3. Add this pointer to `CLAUDE.md` and `AGENTS.md` (create them if missing, do not duplicate it):

   ```
   ## Design
   All UI in this repository follows the style pack in `.design/`.
   Read `.design/STYLE.md` and `.design/specimen.html` before writing any UI.
   ```
