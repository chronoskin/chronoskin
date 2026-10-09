package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/chronoskin/chronoskin/internal/library"
)

const (
	defaultAddr       = ":8080"
	defaultLibraryDir = "library"
	defaultRepo       = "https://github.com/chronoskin/chronoskin" // linked from the header
	defaultSite       = "https://chrono.skin"                      // host of permalinks over stdio, where no request names one
)

// These limits suit any deployment, so none is a flag.
const (
	readHeaderTimeout = 10 * time.Second
	readTimeout       = 15 * time.Second
	writeTimeout      = 30 * time.Second
	idleTimeout       = 90 * time.Second
	maxHeaderBytes    = 16 << 10

	// A client may make rateBurst costly requests at once, and
	// ratePerSecond more each second after that.
	// One Regenerate is three requests (the pick, the page, the specimen),
	// so this lets a visitor press it a hundred times in a row and about
	// three times a second for as long as they like. Everything is served
	// from memory; the limit is against scraping, not against use.
	rateBurst     = 300
	ratePerSecond = 10

	maxCachedResponses = 512

	// Keeps visitor addresses from other users of the machine.
	accessLogMode = 0o640
)

type server struct {
	lib        *library.Library
	baseURL    string
	trustProxy bool
	repo       string // where the header's source link points
	analytics  string // a GoatCounter instance that counts visits; empty for none
	limiter    *limiter

	// The library never changes while the server runs, so what is derived
	// from it alone is computed once.
	eraList    []eraView
	archetypes []string
	homeView   map[string]any
	assets     assets
	pages      templateSet

	// Composed responses, keyed by kind, base URL and Style ID. A pack never
	// changes for an ID; the cache only saves composing it again.
	cacheMu sync.Mutex
	cache   map[string][]byte
}

func newServer(lib *library.Library, baseURL string) *server {
	s := &server{
		lib:     lib,
		baseURL: strings.TrimRight(baseURL, "/"),
		repo:    defaultRepo,
		limiter: newLimiter(rateBurst, ratePerSecond),
		cache:   map[string][]byte{},
	}
	s.assets = loadAssets()
	s.pages = loadTemplates(s.assets)
	s.eraList, s.archetypes = s.eraViews()
	s.homeView = s.homeData()
	return s
}

// cached returns the bytes for key, building them on first use. The cache
// is emptied when it grows large: entries are cheap to rebuild.
func (s *server) cached(key string, build func() []byte) []byte {
	s.cacheMu.Lock()
	b, ok := s.cache[key]
	s.cacheMu.Unlock()
	if ok {
		return b
	}
	b = build()
	s.cacheMu.Lock()
	if len(s.cache) >= maxCachedResponses {
		clear(s.cache)
	}
	s.cache[key] = b
	s.cacheMu.Unlock()
	return b
}

func (s *server) handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthz)
	mux.HandleFunc("GET /{$}", s.home)
	// The list of eras is the home page; each era has a page of its own.
	mux.HandleFunc("GET /eras", redirectHome)
	mux.HandleFunc("GET /eras/{slug}", s.era)
	mux.HandleFunc("GET /saved", s.saved)
	mux.HandleFunc("GET /robots.txt", s.robots)
	mux.HandleFunc("GET /sitemap.xml", s.sitemap)
	mux.HandleFunc("GET /llms.txt", s.llms)
	mux.HandleFunc("GET /docs", s.docs)
	mux.HandleFunc("GET /static/{path...}", s.static)
	mux.HandleFunc("GET /favicon.svg", s.favicon)
	mux.HandleFunc("GET /preview/{era}/{file}", s.preview)
	mux.HandleFunc("GET /api/generate", s.generate)
	mux.HandleFunc("GET /s/{rest...}", s.style)
	mux.HandleFunc("/mcp", s.mcpHTTP)
	mux.HandleFunc("/", s.notFound)
	return s.guarded(mux)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("ok\n"))
}

func redirectHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// guarded sets the security headers of every response and applies the rate limit.
func (s *server) guarded(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()
		h.Set("X-Content-Type-Options", "nosniff")
		h.Set("Referrer-Policy", "no-referrer")
		h.Set("Cross-Origin-Opener-Policy", "same-origin")
		if !servedFromMemory(r.URL.Path) && !s.limiter.allow(s.clientIP(r)) {
			h.Set("Retry-After", "1")
			s.fail(w, r, http.StatusTooManyRequests, rateLimitMessage)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// servedFromMemory paths are exempt from the rate limit, which protects
// what costs work: composing styles and the MCP endpoint. Images and the
// site's own files cost nothing, and one page asks for dozens at once.
func servedFromMemory(path string) bool {
	return path == "/healthz" ||
		path == "/favicon.svg" ||
		path == "/robots.txt" ||
		strings.HasPrefix(path, "/static/") ||
		strings.HasPrefix(path, "/preview/")
}

// base is the public URL of the server as seen by this request.
func (s *server) base(r *http.Request) string {
	if s.baseURL != "" {
		return s.baseURL
	}
	scheme := "http"
	if r.TLS != nil || (s.trustProxy && r.Header.Get("X-Forwarded-Proto") == "https") {
		scheme = "https"
	}
	return scheme + "://" + r.Host
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	encodeJSON(w, v)
}

// encodeJSON writes v indented, with "<", ">" and "&" as they are: the
// files of a pack are read by people and agents, not put into a page.
func encodeJSON(w io.Writer, v any) {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	enc.Encode(v)
}
