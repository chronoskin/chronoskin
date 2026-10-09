# Development

Requires Go 1.26. No Node.js and no other tooling.

A Chromium-based browser (Google Chrome, Chromium) is optional. When one is installed, `make library` draws a preview picture of every layout for the site, and the checks for new packs can render them (`make check-era`). Without one the site runs the same, its cards without pictures.

```
make run      # build the library from the packs, then serve on :8080
make check    # vet, tests, and the rule that the server imports the standard library only
make lint     # check every pack against the pack format
make views    # click through every specimen in a browser (needs Chrome or Chromium)
make build    # bin/server and bin/pack
```

Settings (the port, the folders, the browser the checks drive) are flags and environment variables with working defaults; [server.md](server.md) lists them all in one table.

The server is one binary that reads a built library from a folder. `make library` builds that folder, `library/`, from `packs/`; it is generated and not committed. [server.md](server.md) has the routes, the flags and how the web UI is put together.

`Dockerfile` builds an image of the server with the library in it; run `make library` first.

## What is where

| Path | What |
| --- | --- |
| `cmd/server` | Production server: web UI, HTTP API, MCP. Standard library only |
| `cmd/pack` | Offline tool for packs: the linter, the measurements, building the library |
| `internal/pack` | Style pack format (`.design/`): vocabulary and linter |
| `docs/pack-format.md` | The frozen pack format |
| `packs/<era>/` | One folder per era. `era.json` has its name, description and short code; `01`, `02` and so on are its packs, one layout with its own tokens each, a blend of the period and not a copy of one site; `tokens/NN` are further palette, type and surface sets without a layout |
| `internal/browser` | Headless Chromium over the DevTools Protocol pipe: capture, feature extraction, pixel colours |
| `internal/features` | Feature vector, distance between two pages, Never-rule checks |
| `internal/library` | Library on disk and in memory: building it from packs, composing styles from parts |
| `library/` | The built library the server reads: generated from the packs by `make library`, not committed, never edited by hand |
| `plugin/`, `.claude-plugin/` | Claude Code plugin: skills and MCP configuration |
| `scripts/` | The checks for packs, and the scaffold for a new layout |
| `docs/` | The pack format, the server, development, dependencies |
| `assets/` | The chronoskin mark and the web UI's icons as files, for use outside the site |
