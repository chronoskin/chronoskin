# Assets

The chronoskin mark and the icons of the web UI, as files of their own, for
use outside the site: a README, a slide, a link preview, a plugin listing.

| Path | What |
| --- | --- |
| `logo/mark.svg`, `logo/mark-white.svg` | The mark alone, dark and light |
| `logo/mark-on-black.svg`, `logo/mark-on-white.svg` | The mark on a rounded square, for avatars and app icons |
| `logo/lockup.svg`, `logo/lockup-white.svg` | The mark with the name, dark and light. The name is set in the system sans, as on the site |
| `logo/favicon.svg` | The site's icon; it follows the browser's light or dark theme |
| `logo/avatar-black-*.png`, `logo/avatar-white-*.png` | The mark on a full square, for profile pictures: 1024 for GitHub, 240 for Product Hunt |
| `logo/social-card.png` | The picture shown when a link to the site is shared, 2400 by 1260. The site serves it as `og.jpg` |
| `logo/github-social.jpg` | The same card drawn two to one, 2560 by 1280, for a repository's social preview on GitHub |
| `logo/eras/<era>.svg` | The example mark of each era: the logo its specimens carry beside the name, one drawing per era in the manner of logos of those years. Drawn with `currentColor` on a 24 unit grid |
| `icons/*.svg` | The icons of the web UI, drawn with `currentColor` on a 16, 20 or 24 unit grid |

The mark is two open rings around a centre. It is the mark of chronoskin
itself and does not change. A style pack's specimen is a site of the same
name with a mark of its era instead (`logo/eras/`, see
`docs/pack-format.md`), so that it shows what a logo of that period looked
like.

The site does not read these files. It draws the mark and the icons from
its own templates (`cmd/server/web/templates/partials/`), which the server
embeds, and an era's mark from that era's packs. A test in `cmd/server`
fails when a file here and its source drift apart, so change both together.
