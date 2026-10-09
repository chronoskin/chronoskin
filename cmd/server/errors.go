package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/chronoskin/chronoskin/internal/library"
)

// errorPages is what each failure says to a person in a browser. A program
// gets the plain message alone.
var errorPages = map[int]struct{ title, text string }{
	http.StatusNotFound: {
		"Nothing here",
		"This address does not lead to a page or a style. The style may be from an older version of the library.",
	},
	http.StatusBadRequest: {
		"That request does not work",
		"Something in the address is not what the generator understands.",
	},
	http.StatusTooManyRequests: {
		"Slow down a little",
		"Too many requests came from your address in a short time. Wait a few seconds and try again.",
	},
}

// What a program is told for an unknown address, and when it asks too often.
const (
	notFoundMessage  = "404 page not found"
	rateLimitMessage = "rate limit exceeded"
)

func (s *server) fail(w http.ResponseWriter, r *http.Request, status int, message string) {
	page, known := errorPages[status]
	if !known || !wantsPage(r) {
		http.Error(w, message, status)
		return
	}
	// The page also shows the status and the plain message a program gets.
	detail := message
	if code := strconv.Itoa(status); !strings.HasPrefix(message, code) {
		detail = code + " " + message
	}
	s.render(w, r, "error", map[string]any{
		"Title":   page.title,
		"Nav":     "",
		"Status":  status,
		"Heading": page.title,
		"Text":    page.text,
		"Detail":  detail,
		"Retry":   status == http.StatusTooManyRequests,
	})
}

// failWith answers with the status an error stands for: not found for what
// the library lacks, a bad request otherwise.
func (s *server) failWith(w http.ResponseWriter, r *http.Request, err error) {
	status := http.StatusBadRequest
	if errors.Is(err, library.ErrNotFound) {
		status = http.StatusNotFound
	}
	s.fail(w, r, status, err.Error())
}

func (s *server) notFound(w http.ResponseWriter, r *http.Request) {
	s.fail(w, r, http.StatusNotFound, notFoundMessage)
}

// wantsPage tells a person reading pages in a browser from a program.
func wantsPage(r *http.Request) bool {
	if r.Method != http.MethodGet || !strings.Contains(r.Header.Get("Accept"), "text/html") {
		return false
	}
	p := r.URL.Path
	isFile := strings.HasSuffix(p, ".json") ||
		strings.HasSuffix(p, ".tar.gz") ||
		strings.HasPrefix(p, "/static/") ||
		strings.HasPrefix(p, "/preview/")
	if p == "/mcp" || isFile {
		return false
	}
	return r.URL.Query().Get("format") != "json"
}
