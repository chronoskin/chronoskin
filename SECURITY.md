# Security

If you find a way to make the site at chrono.skin, or a server built from this repository, do something it should not (run script from a style pack on the site's own origin, read or write files outside the library, get round the rate limit, crash it with a request), please report it privately instead of opening a public issue.

Use GitHub's private reporting: the **Security** tab of this repository, then **Report a vulnerability**. Say what you did, what happened and what you expected.

What is in scope: the server in `cmd/server`, the pack tool in `cmd/pack`, and the packs themselves (a pack is HTML and CSS that the site shows inside a sandboxed frame with no script of its own, no network and no origin; a pack that escapes that is a finding).

What is not: reports that need a visitor's own browser extensions or a modified client, missing headers that change nothing in practice, and volume attacks.
