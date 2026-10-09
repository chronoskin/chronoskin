package main

import (
	"bytes"
	"net/http"
)

// siteDescription is what a page without a description of its own says of
// itself to search engines and in shared links.

const siteDescription = "Give your coding agent a look to build in: installable style packs drawn from eras of web design, from table layouts to bento grids."

// render writes a page. It renders to a buffer first, so that a template
// error cannot leave half a page.
func (s *server) render(w http.ResponseWriter, r *http.Request, page string, data map[string]any) {
	s.addShared(r, data)
	var buf bytes.Buffer
	if err := s.pages[page].ExecuteTemplate(&buf, "layout", data); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	h := w.Header()
	h.Set("Content-Type", "text/html; charset=utf-8")
	h.Set("Content-Security-Policy", s.pagePolicy())
	// A page is made for each request. "no-transform" also tells a proxy in
	// front (Cloudflare) to pass it on as it is: left to itself it adds a
	// script of its own, which the policy above would block and browsers
	// report as an error.
	h.Set("Cache-Control", "no-cache, no-transform")
	h.Set("Vary", "Accept-Encoding")
	body := buf.Bytes()
	if acceptsGzip(r) {
		h.Set("Content-Encoding", "gzip")
		body = compress(body)
	}
	// An error page says so in its status, not only in its words.
	if status, _ := data["Status"].(int); status != 0 {
		w.WriteHeader(status)
	}
	w.Write(body)
}

func (s *server) addShared(r *http.Request, data map[string]any) {
	data["Repo"] = s.repo
	data["Analytics"] = s.analytics
	// The address a page is known by: without the state a visitor adds,
	// and for a page that is one of many alike, the one that stands for
	// them (CanonicalPath).
	canonical := r.URL.Path
	if p, _ := data["CanonicalPath"].(string); p != "" {
		canonical = p
	}
	data["Canonical"] = s.base(r) + canonical
	// A share image must be a full address. A page without a picture of
	// its own shares the site's card.
	image, _ := data["Image"].(string)
	if image == "" {
		image, _ = s.assets.url("og.jpg")
	}
	data["Image"] = s.base(r) + image
	if _, ok := data["PageTitle"]; !ok {
		data["PageTitle"] = data["Title"]
	}
	if _, ok := data["Description"]; !ok {
		data["Description"] = siteDescription
	}
}

// pagePolicy is the Content-Security-Policy of the site's pages: nothing
// inline and nothing from elsewhere, except the visit counter when one is
// set (its script, and the request it reports a visit with).
func (s *server) pagePolicy() string {
	csp := "default-src 'self'; frame-src 'self'; base-uri 'none'; form-action 'self'; frame-ancestors 'none'"
	if s.analytics != "" {
		csp += "; script-src 'self' " + s.analytics +
			"; connect-src 'self' " + s.analytics +
			"; img-src 'self' " + s.analytics
	}
	return csp
}
