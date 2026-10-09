.PHONY: new-layout check-era release test check build run lint views library

test:
	go vet ./...
	go test ./...

check: test
	sh scripts/check-deps.sh

build:
	CGO_ENABLED=0 go build -trimpath -o bin/server ./cmd/server
	CGO_ENABLED=0 go build -trimpath -o bin/pack ./cmd/pack

run: library
	go run ./cmd/server

lint:
	go run ./cmd/pack lint packs/*/[0-9][0-9]

# Clicks through every specimen in a browser: links, menus, phone width,
# contrast, centring, leftover rules. Needs Chrome or Chromium; takes a
# few minutes.
views: build
	bin/pack views packs/*/[0-9][0-9]

# The library's version names one build of it. It is part of every preview
# picture's address, so a release must have a version of its own or browsers
# keep the pictures of the last one.
VERSION ?= dev

# The library in library/ is generated from the packs, so it is rebuilt
# rather than edited.
library:
	rm -rf library
	go run ./cmd/pack build-library -version $(VERSION) -out library -previews packs

# One archive to copy to the server: the server for Linux and the library.
release:
	$(MAKE) library VERSION=$$(date +%Y.%m.%d-%H%M)
	rm -rf dist && mkdir -p dist/chronoskin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o dist/chronoskin/server ./cmd/server
	cp -R library dist/chronoskin/library
	if [ -d deploy ]; then cp -R deploy dist/chronoskin/deploy; fi
	COPYFILE_DISABLE=1 tar --no-xattrs --no-mac-metadata -C dist -czf dist/chronoskin-linux-amd64.tar.gz chronoskin

# A new layout for an era, started as a copy of the era's first pack.
# usage: make new-layout ERA=terminal
new-layout:
	sh scripts/new-layout.sh $(ERA)

# Every check a new layout or token set must pass, for one era.
# usage: make check-era ERA=terminal [LAYOUT=07]
check-era: build
	sh scripts/check-era.sh $(ERA) $(LAYOUT)
