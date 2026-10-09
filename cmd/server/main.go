// Command server is the production server: web UI, HTTP API and MCP endpoint.
// It imports the standard library only; scripts/check-deps.sh enforces this.
//
//	server [-addr :8080] [-library dir] [-base-url https://host] [-repo URL]
//	       [-trust-proxy] [-access-log FILE] [-analytics URL]
//	server mcp --stdio [-library dir] [-base-url https://host]
//
// The fallbacks of the flags and the limits the server runs within are the
// constants at the top of server.go.
package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/chronoskin/chronoskin/internal/library"
)

func main() {
	log.SetFlags(0)
	args := os.Args[1:]
	stdio := len(args) > 0 && args[0] == "mcp"
	if stdio {
		args = args[1:]
	}
	fs := flag.NewFlagSet("server", flag.ExitOnError)
	addr := fs.String("addr", defaultAddr, "listen address")
	libDir := fs.String("library", defaultLibraryDir, "library directory")
	baseURL := fs.String("base-url", "", "public URL of this server, used in permalinks (default: taken from each request)")
	useStdio := fs.Bool("stdio", false, "with the mcp subcommand: serve MCP on stdin and stdout")
	repo := fs.String("repo", defaultRepo, "source repository linked from the header")
	trustProxy := fs.Bool("trust-proxy", false, "rate-limit by the last X-Forwarded-For address; set only behind a reverse proxy that appends it")
	accessLog := fs.String("access-log", "", "file to append one line per request to, in the combined log format (default: none)")
	analytics := fs.String("analytics", "", "URL of a GoatCounter instance to count visits with, such as https://stats.example.com (default: none)")
	fs.Parse(args)

	lib, err := library.Load(*libDir)
	if err != nil {
		log.Fatalf("loading library: %v", err)
	}
	s := newServer(lib, *baseURL)
	s.trustProxy = *trustProxy
	s.repo = *repo
	s.analytics = strings.TrimRight(*analytics, "/")

	if stdio {
		if !*useStdio {
			log.Fatal("usage: server mcp --stdio")
		}
		if err := s.serveStdio(os.Stdin, os.Stdout); err != nil {
			log.Fatal(err)
		}
		return
	}

	handler := s.handler()
	if *accessLog != "" {
		// Opened for appending, so that logrotate can truncate it in place.
		f, err := os.OpenFile(*accessLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, accessLogMode)
		if err != nil {
			log.Fatalf("access log: %v", err)
		}
		handler = s.logged(handler, f)
	}
	srv := &http.Server{
		Addr:              *addr,
		Handler:           handler,
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}
	log.Printf("library %s (%d eras), listening on %s", lib.Version, len(lib.Eras), *addr)
	log.Fatal(srv.ListenAndServe())
}
