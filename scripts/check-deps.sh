#!/bin/sh
# Fails if cmd/server depends on anything outside the standard library and
# this module.
set -eu

module=$(go list -m)
bad=$(go list -deps -f '{{if not .Standard}}{{.ImportPath}}{{end}}' ./cmd/server \
	| grep -v "^$module\(/\|\$\)" || true)

if [ -n "$bad" ]; then
	echo "cmd/server imports non-stdlib packages:" >&2
	echo "$bad" >&2
	exit 1
fi
echo "cmd/server: standard library only"
