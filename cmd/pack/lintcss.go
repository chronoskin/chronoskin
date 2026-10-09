package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/chronoskin/chronoskin/internal/pack"
)

// lintCSS is the lint_css tool as a command, for agents without MCP.
func lintCSS(args []string) (int, error) {
	fs := flag.NewFlagSet("lint-css", flag.ExitOnError)
	packDir := fs.String("pack", defaultPackDir, "pack directory")
	fs.Parse(args)
	if fs.NArg() == 0 {
		return 2, fmt.Errorf("usage: pack lint-css [-pack .design] <file.css>...")
	}
	tokens, err := os.ReadFile(filepath.Join(*packDir, "tokens.css"))
	if err != nil {
		return 2, err
	}
	status := 0
	for _, path := range fs.Args() {
		css, err := os.ReadFile(path)
		if err != nil {
			return 2, err
		}
		violations, err := pack.LintCSS(string(css), string(tokens))
		if err != nil {
			return 2, fmt.Errorf("%s: %w", path, err)
		}
		for _, v := range violations {
			fmt.Printf("%s: %s\n", path, v)
			status = 1
		}
		if len(violations) == 0 {
			fmt.Printf("%s: ok\n", path)
		}
	}
	return status, nil
}
