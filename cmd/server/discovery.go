package main

import (
	"encoding/xml"
	"net/http"
	"strings"
)

// robots, sitemap and llms are generated from the library, so they are
// right for whatever eras the server was started with.
func (s *server) robots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", cacheHour)
	// Generating is an action, not a page.
	w.Write([]byte("User-agent: *\nAllow: /\nDisallow: /api/\n\nSitemap: " + s.base(r) + "/sitemap.xml\n"))
}

// sitemap lists the fixed pages and every layout with its own tokens. The
// mixes are countless and all reachable from those.
func (s *server) sitemap(w http.ResponseWriter, r *http.Request) {
	type url struct {
		Loc string `xml:"loc"`
	}
	set := struct {
		XMLName xml.Name `xml:"urlset"`
		Xmlns   string   `xml:"xmlns,attr"`
		URLs    []url    `xml:"url"`
	}{Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9"}
	base := s.base(r)
	paths := []string{"/", "/docs"}
	for _, e := range s.eraList {
		paths = append(paths, "/eras/"+e.Slug)
	}
	for _, path := range paths {
		set.URLs = append(set.URLs, url{base + path})
	}
	for _, e := range s.eraList {
		for _, l := range e.LayoutViews {
			set.URLs = append(set.URLs, url{base + l.Link})
		}
	}
	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	w.Header().Set("Cache-Control", cacheHour)
	w.Write([]byte(xml.Header))
	xml.NewEncoder(w).Encode(set)
}

// llms is /llms.txt: how a model or an agent uses the site, in plain Markdown.
func (s *server) llms(w http.ResponseWriter, r *http.Request) {
	base := s.base(r)
	eras, layouts := s.totals()
	var b strings.Builder
	b.WriteString("# chronoskin\n\n")
	b.WriteString("> " + siteDescription + "\n\n")
	b.WriteString("A style is a folder called `.design/` with five files: `STYLE.md` (the rules in words), `tokens.css` (colours, type and surface as CSS custom properties), `components.css` (component classes), `specimen.html` (a small example site that uses every component) and `style.json`. A coding agent reads them before it writes any interface and builds only with the pack's tokens and classes.\n\n")
	b.WriteString("The library has " + plural(eras, "era") + " and " + plural(layouts, "layout") + ". A style combines one layout with one palette, one type set and one surface set of the same era. Every combination has a Style ID such as `" + s.exampleID() + "`, and the same ID always gives the same files.\n\n")
	b.WriteString("## Pages\n\n")
	b.WriteString("- [Eras](" + base + "/): pick an era and get a style\n")
	b.WriteString("- [Docs](" + base + "/docs): what a pack contains, the Style ID, how to install\n\n")
	b.WriteString("## For agents and tools\n\n")
	b.WriteString("- `GET " + base + "/api/generate?format=json`: a new style as JSON. Parameters: `era`, `archetype`, `mode` (`mix` or `pure`), `seed`, `from` and `lock` to regenerate, `density`, and one parameter per palette colour to set by hand\n")
	b.WriteString("- `GET " + base + "/s/<id>.json`: the five files of a style as JSON\n")
	b.WriteString("- `GET " + base + "/s/<id>.tar.gz`: the `.design/` folder as an archive (`curl -s " + base + "/s/<id>.tar.gz | tar xz`)\n")
	b.WriteString("- `GET " + base + "/s/<id>/specimen.html`: the example site on its own\n")
	b.WriteString("- `POST " + base + "/mcp`: an MCP server with the tools `list_eras`, `generate_style`, `get_style` and `lint_css`\n\n")
	b.WriteString("## Eras\n\n")
	for _, e := range s.eraList {
		b.WriteString("- [" + e.Name + "](" + base + e.Link + ") (" + e.Years + ", `" + e.Slug + "`): " + e.Description + "\n")
	}
	w.Header().Set("Content-Type", "text/markdown; charset=utf-8")
	w.Header().Set("Cache-Control", cacheHour)
	w.Write([]byte(b.String()))
}

func (s *server) exampleID() string {
	return s.lib.Short(s.firstStyle())
}
