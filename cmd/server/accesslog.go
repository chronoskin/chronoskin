package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// logged writes one line per request to out in the combined log format,
// which log analysers read as it is:
//
//	203.0.113.7 - - [07/Oct/2026:14:03:11 +0000] "GET /s/v1-yk-3333 HTTP/1.1" 200 5123 "-" "curl/8.7"
//
// It is how installs that never run the site's script are counted. Nothing
// else about a visitor is written: no cookies, no other headers.
func (s *server) logged(next http.Handler, out io.Writer) http.Handler {
	var mu sync.Mutex
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &recorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)
		line := fmt.Sprintf("%s - - [%s] %q %d %d %q %q\n",
			s.remote(r),
			time.Now().UTC().Format(logTimeLayout),
			r.Method+" "+r.URL.RequestURI()+" "+r.Proto,
			rec.status,
			rec.bytes,
			dash(r.Referer()),
			dash(r.UserAgent()))
		mu.Lock()
		io.WriteString(out, line)
		mu.Unlock()
	})
}

const logTimeLayout = "02/Jan/2006:15:04:05 -0700"

// dash is the log's way of writing an empty field.
func dash(s string) string {
	if s == "" {
		return "-"
	}
	return s
}

func (s *server) remote(r *http.Request) string {
	if ip := s.visitor(r); ip != nil {
		return ip.String()
	}
	return r.RemoteAddr
}

// recorder notes the status and size of a response as it is written.
type recorder struct {
	http.ResponseWriter
	status int
	bytes  int
}

func (r *recorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func (r *recorder) Write(b []byte) (int, error) {
	n, err := r.ResponseWriter.Write(b)
	r.bytes += n
	return n, err
}

// Flush keeps streamed responses (the MCP endpoint) streaming.
func (r *recorder) Flush() {
	if f, ok := r.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}
