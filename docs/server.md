# The server

```
bin/server [-addr :8080] [-library library] [-base-url https://host] [-trust-proxy] [-access-log file] [-analytics https://stats.host]
bin/server mcp --stdio [-library library]
```

## Settings

Nothing is read from a configuration file. What differs from one machine to the next is a flag or an environment variable, and each has a default that works for a local run.

| Setting | Default | What it is |
| --- | --- | --- |
| `-addr` | `:8080` | Where the server listens |
| `-library` | `library` | The built library to serve (`make library` writes it) |
| `-base-url` | taken from each request | The public address, used in permalinks. Set it in production; over stdio, where no request names a host, permalinks fall back to `https://chrono.skin` |
| `-repo` | `https://github.com/chronoskin/chronoskin` | The source repository linked from the header |
| `-trust-proxy` | off | Rate-limit by the last `X-Forwarded-For` address. Only behind a reverse proxy that appends it |
| `-access-log` | none | A file to append one line per request to, in the combined log format |
| `-analytics` | none | The address of a GoatCounter instance that counts visits |
| `CHROME` (environment) | found by itself | The Chromium-based browser `bin/pack` drives for preview pictures, `views`, `shot` and `adhere`. Without it, Google Chrome or Chromium in their usual places, then `chromium`, `chromium-browser` or `google-chrome` on the `PATH` |
| `SHOTS` (environment) | `shots` | Where the checks in `scripts/` write their screenshots |
| `VERSION` (make) | `dev` | The version `make library` stamps on the library; `make release` uses the date and time |

## Routes

| Route | Returns |
| --- | --- |
| `/` | Every era as a card to choose, with a bar that narrows them by page type and years and generates a style |
| `/eras/<slug>` | One era: what it is and every layout it has. `/eras` itself leads to the home page |
| `/saved`, `/docs` | Styles saved in this browser; how to use a pack |
| `/api/generate` | `era`, `archetype`, `mode`, `seed`, `from`, `lock`; redirects to `/s/<id>`, or JSON with `format=json` |
| `/s/<id>` | One style on one screen: the specimen at its design width, and beside it the eras and page types to regenerate among, the four parts with Keep and arrows for each, Regenerate and the install commands |
| `/preview/<era>/<n>.jpg` | Preview image of a structure with its own tokens |
| `/s/<id>/specimen.html` | The specimen alone, stylesheets inlined |
| `/s/<id>.json`, `/s/<id>.tar.gz` | The pack as JSON, or as a `.design/` tarball |
| `/mcp` | MCP over HTTP: `list_eras`, `generate_style`, `get_style`, `lint_css` |
| `/healthz` | Liveness |

Behind a reverse proxy, set `-base-url` to the public address and `-trust-proxy` so that the rate limit sees the client and not the proxy. `lint_css` accepts at most 256 KB and reports at most 200 violations. `-access-log` appends one line per request in the combined log format, and `-analytics` names a GoatCounter instance that counts visits; both are off unless set.

A library's Style IDs are only stable for that library: parts are numbered in the order the packs are given to `build-library`. A new release that adds packs must use a new `-prefix` (`v2`, …), or old IDs would point at different styles while clients still hold cached copies.

## How the web UI is built

Everything the site shows is in `cmd/server/web/` and compiled into the binary:

| Path | What |
| --- | --- |
| `templates/layout.html` | The document every page shares |
| `templates/partials/` | Pieces used by more than one page |
| `templates/pages/` | One file per page; each defines `main` |
| `static/css/` | One stylesheet per component. Every file here is joined into `/static/app.css`: `tokens`, `base` and `layout` first, `touch` last, the rest by name |
| `static/js/` | The Alpine store and one file per component, all joined into `/static/app.js`, the store first |
| `static/vendor/` | Third-party files, unchanged |

The pages are rendered on the server and every control is a link or a form field, so the site works without script. Alpine.js (its CSP build, one vendored file) adds what needs the browser. Pages are sent with a strict Content-Security-Policy: no inline script or style. ## Adding to the web UI

- **A component**: one stylesheet in `static/css/`, and one script in `static/js/components/` if it needs the browser. Both are picked up by their presence; there is no list to edit. Where your stylesheet and another style the same element, make your selector more specific rather than counting on file order.
- **Behaviour**: register it with `Alpine.data('name', () => ({ ... }))` and name it in the template with `x-data="name"`. The CSP build of Alpine evaluates no expressions, so a template may only name a property or a method: `x-on:click="toggle"`, `x-text="label"`, never `x-on:click="open = !open"`. Declare every property the component sets in its data object; one first set in `init()` lands on the outer component.
- **A page**: a template in `templates/pages/` that defines `main`, a handler that calls `s.render`, and a route in `server.go`.
- **An icon**: a `{{define "name"}}` in `templates/partials/icons.html`, and the same drawing as a file in `assets/icons/`; a test keeps the two the same.
- **A template's shape**: one element to a line; a long tag has one attribute to a line, plain attributes first, then `data-*`, then `x-*`. Begin each template with a comment saying what it is for and what data it takes.

No inline script or style: the Content-Security-Policy forbids both. The site's own interface is black, white and grey.

After a change, `make check` runs the tests, and `make run` serves the site at `http://localhost:8080`.
