---
name: design-new
description: Generate a new design style for an era and page type, preview it, then regenerate or apply it.
argument-hint: "[era] [archetype]"
---

Arguments: $ARGUMENTS (an era slug, then a page type; both optional).

1. If no era was given, call `list_eras` and ask which one to use.
2. Call `generate_style` with the era and archetype.
3. Show the style's summary, Style ID and preview URL.
4. Offer two next steps: regenerate (call `generate_style` again with `from` set to the current ID and any `lock` the user asks for), or apply it with `/chronoskin:design-apply <style-id>`. `generate_style` also takes `density` (`compact`, `normal`, `roomy`) and `colours` (an object of palette token and colour, such as `{"--color-accent": "#2b55e0"}`) when the user asks for tighter or looser spacing or for their own brand colours; pass them with `from` and all four parts locked to adjust the current style without changing it otherwise.
