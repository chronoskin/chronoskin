# Dependencies

The Go module has no third-party dependencies: `go.mod` requires nothing, and both commands build from the standard library alone.

`scripts/check-deps.sh`, run by `make check` and in CI, fails if `cmd/server` imports anything outside the standard library and this module. A dependency for `cmd/pack` needs a written reason here first.

No WebSocket client is needed to drive the browser: `internal/browser` talks to Chromium through `--remote-debugging-pipe`, which is plain pipes.

## Front end

The site uses one third-party script, vendored as a single file and served by the server itself. Nothing is installed or built.

| File | Version | Reason |
| --- | --- | --- |
| `cmd/server/web/static/vendor/alpine-csp.min.js` | Alpine.js 3.14.9, CSP build (`@alpinejs/csp`, `dist/cdn.min.js`) | Small reactive behaviour in the pages (copy buttons, saved styles, preview fitting) without a framework or a build step. The CSP build runs without `unsafe-eval`, so the pages keep a strict Content-Security-Policy |

To update it, replace the file with the same path from a newer release and check the pages for console errors.
