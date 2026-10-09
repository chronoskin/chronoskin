package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"net/http"
	"strings"
	"time"

	"github.com/chronoskin/chronoskin/internal/library"
	"github.com/chronoskin/chronoskin/internal/pack"
)

type styleInfo struct {
	ID        string            `json:"id"`      // the short form, the one to pass around
	LongID    string            `json:"long_id"` // the form that names the era and each part
	Name      string            `json:"name"`
	Era       library.Era       `json:"era"`
	Archetype string            `json:"archetype"`
	Mode      string            `json:"mode"`
	Viewport  pack.Viewport     `json:"viewport"`
	Density   string            `json:"density,omitempty"`
	Colours   map[string]string `json:"colours,omitempty"`
	Summary   string            `json:"summary"`
	URL       string            `json:"url"`
	Preview   string            `json:"preview_url"`
}

func (s *server) describe(id pack.ID, base string) styleInfo {
	st, _ := s.lib.Style(id)
	url := base + "/s/" + s.lib.Short(id)
	custom, _ := pack.DecodeColours(id.Colours)
	colours := map[string]string{}
	for token, c := range custom {
		colours[token] = c.CSS()
	}
	return styleInfo{
		ID:        s.lib.Short(id),
		LongID:    id.String(),
		Name:      st.Name,
		Era:       st.Era,
		Archetype: st.Archetype,
		Mode:      st.Mode,
		Viewport:  st.Viewport,
		Density:   id.Density,
		Colours:   colours,
		Summary:   st.Description,
		URL:       url,
		Preview:   url + "/specimen.html",
	}
}

func (s *server) packFiles(id pack.ID, base string) (map[string]string, error) {
	return s.lib.Pack(id, base+"/s/"+s.lib.Short(id))
}

// styleKinds are the forms a style is served in besides its page.
var styleKinds = []struct{ suffix, kind string }{
	{".tar.gz", "tar"},
	{".json", "json"},
	{"/specimen.html", "specimen"},
}

func splitKind(rest string) (id, kind string) {
	for _, k := range styleKinds {
		if id, ok := strings.CutSuffix(rest, k.suffix); ok {
			return id, k.kind
		}
	}
	return rest, "page"
}

// style serves /s/<id>, /s/<id>.json, /s/<id>.tar.gz and /s/<id>/specimen.html.
func (s *server) style(w http.ResponseWriter, r *http.Request) {
	text, kind := splitKind(r.PathValue("rest"))
	id, ok := s.lib.Parse(text)
	if !ok {
		s.notFound(w, r)
		return
	}
	if _, err := s.lib.Style(id); err != nil {
		s.failWith(w, r, err)
		return
	}
	base := s.base(r)
	key := kind + " " + base + " " + id.String()
	if s.baseURL == "" {
		// The permalink inside these responses is built from the request's host.
		w.Header().Add("Vary", "Host")
	}
	switch kind {
	case "page":
		s.stylePage(w, r, id, base)
	case "json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		s.revalidated(w, r, s.cached(key, func() []byte { return s.styleJSON(id, base) }))
	case "tar":
		w.Header().Set("Content-Type", "application/gzip")
		w.Header().Set("Content-Disposition", `attachment; filename="`+s.lib.Short(id)+`.tar.gz"`)
		s.revalidated(w, r, s.cached(key, func() []byte {
			files, _ := s.packFiles(id, base)
			return tarball(files)
		}))
	case "specimen":
		s.specimen(w, r, id, base)
	}
}

func (s *server) styleJSON(id pack.ID, base string) []byte {
	files, _ := s.packFiles(id, base)
	var buf bytes.Buffer
	encodeJSON(&buf, struct {
		styleInfo
		Files map[string]string `json:"files"`
	}{s.describe(id, base), files})
	return buf.Bytes()
}

// specimenPolicy is the Content-Security-Policy of a specimen, which is
// generated markup. It gets no network and no origin of its own, only this
// site may frame it, and the one script it may run is the line this server
// adds (see menuScript).
var specimenPolicy = "sandbox allow-scripts; default-src 'none'; style-src 'unsafe-inline'; img-src data:; " +
	"script-src '" + menuScriptHash + "'; frame-ancestors 'self'; base-uri 'none'; form-action 'none'"

// specimen serves the example site. Shown inside the site ("?frame") it
// draws no scrollbar; opened on its own it is the pack's file as it is.
func (s *server) specimen(w http.ResponseWriter, r *http.Request, id pack.ID, base string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Security-Policy", specimenPolicy)
	framed := r.URL.Query().Has("frame")
	key := "specimen " + id.String()
	if framed {
		key = "framed " + id.String()
	}
	s.revalidated(w, r, s.cached(key, func() []byte {
		files, _ := s.packFiles(id, base)
		page := served(files)
		if framed {
			page = strings.Replace(page, "</head>", framedStyle+"</head>", 1)
		}
		return []byte(page)
	}))
}

// revalidated writes a body the browser may keep but must check before
// using again: a library can be rebuilt under the same name, and a stale
// specimen would then show a style that no longer exists.
func (s *server) revalidated(w http.ResponseWriter, r *http.Request, body []byte) {
	sum := sha256.Sum256(body)
	etag := `"` + hex.EncodeToString(sum[:8]) + `"`
	w.Header().Set("ETag", etag)
	w.Header().Set("Cache-Control", "no-cache, no-transform")
	if r.Header.Get("If-None-Match") == etag {
		w.WriteHeader(http.StatusNotModified)
		return
	}
	w.Write(body)
}

// epoch is the modification time of every archived file, fixed so that the
// same style always gives the same bytes.
var epoch = time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)

// tarball packs the files as .design/, byte-identical on every call.
func tarball(files map[string]string) []byte {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gz)
	tw.WriteHeader(&tar.Header{Name: pack.Dir + "/", Typeflag: tar.TypeDir, Mode: 0o755, ModTime: epoch})
	for _, name := range pack.Files {
		content := files[name]
		tw.WriteHeader(&tar.Header{Name: pack.Dir + "/" + name, Mode: 0o644, Size: int64(len(content)), ModTime: epoch})
		tw.Write([]byte(content))
	}
	tw.Close()
	gz.Close()
	return buf.Bytes()
}

// menuScript is the one script a specimen runs. A pack's views change
// through ":target", so the page never reloads and a "<details>" menu would
// stay open over the view its link led to: the first line closes it. The
// second passes the site's shortcut keys up from the preview, where they
// would otherwise be lost once the visitor has clicked into it.
const menuScript = `addEventListener("click",function(e){var a=e.target.closest&&e.target.closest('a[href^="#"]'),d=a&&a.closest("details");if(d)d.open=false});` +
	`addEventListener("keydown",function(e){if(parent===window||e.ctrlKey||e.metaKey||e.altKey)return;var t=e.target;if(t&&t.closest&&t.closest("input,textarea,select"))return;var k=e.key.length===1?e.key.toLowerCase():e.key;if(k==="r"||k==="a"||k==="ArrowLeft"||k==="ArrowRight"){e.preventDefault();parent.postMessage({chronoskinKey:k},"*")}})`

// menuScriptHash names menuScript in the specimen's Content-Security-Policy.
var menuScriptHash = func() string {
	sum := sha256.Sum256([]byte(menuScript))
	return "sha256-" + base64.StdEncoding.EncodeToString(sum[:])
}()

// served is a pack's specimen as one self-contained page, menuScript at its end.
func served(files map[string]string) string {
	page := pack.InlineSpecimen(files)
	script := "<script>" + menuScript + "</script>"
	if i := strings.LastIndex(page, "</body>"); i >= 0 {
		return page[:i] + script + page[i:]
	}
	return page + script
}

// framedStyle hides the scrollbar of a specimen shown in the site's preview.
const framedStyle = "<style>html{scrollbar-width:none;-ms-overflow-style:none}::-webkit-scrollbar{display:none}</style>"
