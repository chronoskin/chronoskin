package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"html/template"
	"io/fs"
	"mime"
	"net/http"
	"path"
	"slices"
	"strings"
)

// webFS is the site's own files: templates, stylesheets, scripts. They are
// compiled into the binary, so the server is still one file to deploy.
//
//go:embed web
var webFS embed.FS

// Every stylesheet in web/static/css is joined into /static/app.css and
// every script in web/static/js into /static/app.js, so a new component is
// one new file and there is no build step. Stylesheets are joined by name,
// the foundations first and the touch sizes last. Where two components
// style one element, the more specific selector must decide, not file order.
var (
	stylesFirst = []string{"tokens.css", "base.css", "layout.css"}
	stylesLast  = []string{"touch.css"}
)

// Scripts are joined by path, the store before the components that use it.
var scriptsFirst = []string{"store.js"}

// singleAssets are served as they are, each under its own name.
var singleAssets = []string{"favicon.svg", "og.jpg", "theme.js", "vendor/alpine-csp.min.js"}

const (
	// cacheForever is for an address that names its content by a hash or by
	// the library version.
	cacheForever = "public, max-age=31536000, immutable"
	cacheHour    = "public, max-age=3600"
	cacheDay     = "public, max-age=86400"
)

// ordered lists the files of dir with the extension: first, the rest by name, last.
func ordered(files fs.FS, dir, ext string, first, last []string) []string {
	var rest []string
	err := fs.WalkDir(files, dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || path.Ext(p) != ext {
			return err
		}
		name := strings.TrimPrefix(p, dir+"/")
		if !slices.Contains(first, name) && !slices.Contains(last, name) {
			rest = append(rest, name)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	slices.Sort(rest)
	return slices.Concat(first, rest, last)
}

type asset struct {
	body, gzipped []byte
	contentType   string
	hash          string // of the body, shortened; the "v" of the asset's address
}

// assets holds every file served under /static/, by path.
type assets map[string]*asset

func loadAssets() assets {
	static, err := fs.Sub(webFS, "web/static")
	if err != nil {
		panic(err)
	}
	out := assets{}
	out.add("app.css", joined(static, "css", ordered(static, "css", ".css", stylesFirst, stylesLast)))
	out.add("app.js", joined(static, "js", ordered(static, "js", ".js", scriptsFirst, nil)))
	for _, name := range singleAssets {
		body, err := fs.ReadFile(static, name)
		if err != nil {
			panic(err)
		}
		out.add(name, body)
	}
	return out
}

func joined(files fs.FS, dir string, names []string) []byte {
	var b bytes.Buffer
	for _, name := range names {
		part, err := fs.ReadFile(files, path.Join(dir, name))
		if err != nil {
			panic(err)
		}
		b.Write(part)
		b.WriteByte('\n')
	}
	return b.Bytes()
}

func (a assets) add(name string, body []byte) {
	sum := sha256.Sum256(body)
	a[name] = &asset{
		body:        body,
		gzipped:     compress(body),
		contentType: mime.TypeByExtension(path.Ext(name)),
		hash:        hex.EncodeToString(sum[:6]),
	}
}

// url carries the content hash, so a browser may keep the file for good.
func (a assets) url(name string) (string, error) {
	f, ok := a[name]
	if !ok {
		return "", fs.ErrNotExist
	}
	return "/static/" + name + "?v=" + f.hash, nil
}

func (s *server) static(w http.ResponseWriter, r *http.Request) {
	s.serveAsset(w, r, r.PathValue("path"))
}

func (s *server) favicon(w http.ResponseWriter, r *http.Request) {
	s.serveAsset(w, r, "favicon.svg")
}

func (s *server) serveAsset(w http.ResponseWriter, r *http.Request, name string) {
	f, ok := s.assets[name]
	if !ok {
		http.NotFound(w, r)
		return
	}
	if r.URL.Query().Get("v") == f.hash {
		w.Header().Set("Cache-Control", cacheForever)
	} else {
		w.Header().Set("Cache-Control", cacheHour)
	}
	w.Header().Set("Content-Type", f.contentType)
	w.Header().Set("Vary", "Accept-Encoding")
	if acceptsGzip(r) {
		w.Header().Set("Content-Encoding", "gzip")
		w.Write(f.gzipped)
		return
	}
	w.Write(f.body)
}

func compress(body []byte) []byte {
	var b bytes.Buffer
	z, _ := gzip.NewWriterLevel(&b, gzip.BestCompression)
	z.Write(body)
	z.Close()
	return b.Bytes()
}

func acceptsGzip(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Accept-Encoding"), "gzip")
}

// templateSet is the parsed pages, by name.
type templateSet = map[string]*template.Template

// loadTemplates parses one set per page: the layout, the partials and that page.
func loadTemplates(a assets) templateSet {
	pages, err := fs.Glob(webFS, "web/templates/pages/*.html")
	if err != nil || len(pages) == 0 {
		panic("no page templates")
	}
	funcs := template.FuncMap{"asset": a.url}
	out := templateSet{}
	for _, page := range pages {
		t := template.New("layout").Funcs(funcs)
		t = template.Must(t.ParseFS(webFS, "web/templates/layout.html", "web/templates/partials/*.html", page))
		out[strings.TrimSuffix(path.Base(page), ".html")] = t
	}
	return out
}
