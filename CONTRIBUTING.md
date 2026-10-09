# Contributing

The most useful thing to add is a layout: a new kind of page for an era that has few. Token sets and whole eras are added the same way. None of it needs Go; a pack is five plain files.

## Before you start

Read [docs/pack-format.md](docs/pack-format.md). It is the contract: the tokens every pack defines, the components every specimen shows, how a specimen is split into views, and how Never rules are written. The linter enforces it.

Then look at two or three packs of the era you want to add to. Each era is a folder, `packs/<era>/`, and its packs are numbered inside it: `01`, `02`, `03` and so on.

You need Go 1.26 to build the checks, and Google Chrome or Chromium for the ones that render a pack to check it and to take the screenshots you look at.

## A new layout

```
make new-layout ERA=terminal
```

This copies the era's first pack to the next free number, renamed, so you start from files that already pass the linter. Now make it a different page:

1. **A new kind of page.** Not a rearrangement of an existing layout: a different skeleton (where the navigation sits, how the page is divided) and a different example site. If the era has a landing page and a blog, a dashboard or a shop is welcome.
2. **Its own token set.** A palette, a type set and a surface set in `tokens.css` that differ clearly from every set the era has: another colour family, another type pairing, another treatment of borders and shadows.
3. **Three to six views.** The specimen is a small click-through site, one view showing at a time, with every component of the checklist somewhere it would really be used.
4. **The site is chronoskin.** Every specimen is the same site, named `chronoskin`, with the `brand` component. Do not invent other product or company names.
5. **True to the period.** A style is a blend of several real sites of its years and should not be recognisable as any single one.

## Mixing is the hard part

Any layout of an era is shown with any palette, type set and surface set of that era. So your layout must look right with every existing token set, and your token set on every existing layout:

- Never write a colour, font, border style, radius or shadow into `components.css`; use the tokens.
- Never assume a palette is light or dark.
- Follow the era's first pack for which text token sits on which fill; its `STYLE.md` says so.

## Check it

```
make check-era ERA=terminal LAYOUT=07
```

This lints each pack of the era, checks its Never rules against its own specimen, clicks through the views of each specimen in a browser (at 390px too, for a pack marked `"responsive": true`), checks that the era's token sets differ from one another, and composes your layout with every token set. Screenshots of every combination are written to `shots/`; look at them, at the design width and at 390px. Leave out `LAYOUT` to compose the whole era, which is what a new token set needs and takes several minutes.

`make run` then shows the library with your pack in it at `http://localhost:8080`.

## A new era

An era is a folder `packs/<slug>/` with its first pack in `01/` and an `era.json` beside it: the era's name, a description and a two-character code that no other era uses. Its years come from the first pack's `style.json`. Bring at least three layouts, so that mixing has something to mix.

## Rules of the repository

- No Node.js, npm, bundlers or CSS frameworks.
- `cmd/server` imports the Go standard library only; `make check` fails otherwise.
- The site's own interface is black, white and grey.
- Changes to the pack format change `docs/pack-format.md` and `internal/pack/vocab.go` together; a test keeps them in step.

## Changing the site

[docs/development.md](docs/development.md) says how to build and run everything and what is where; [docs/server.md](docs/server.md) describes how the web UI is put together: one template per page, one stylesheet and one script per component. `make check` runs the tests.

## Conduct and security

Be kind and specific: [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md). Found a way to make the site or a pack do something it should not? Report it privately: [SECURITY.md](SECURITY.md).

## Licence of contributions

By contributing you agree that your contribution is licensed under the project's [MIT licence](LICENSE).
