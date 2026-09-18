# Makefile for flexera-cli
#
# Common workflows:
#   make all                 — update Go deps, then regenerate from the pinned OpenAPI spec
#   make generate            — regenerate all commands from the OpenAPI spec
#   make build               — build the flexera-cli binary (version-stamped)
#   make refresh             — regenerate commands and build the binary
#   make test                — unit tests, no network, no credentials
#   make test-integration    — live read-only test suite (requires FLEXERA_NAM_REFRESH_TOKEN)
#   make coverage-report     — print line/func coverage from the integration run
#   make coverage-html       — open coverage.html in the default browser
#   make coverage-check      — enforce 100% on read-only handler allowlist
#   make check-client-coverage — fail if unified-go-client gained hand-written
#                                symbols with no triaged CLI-coverage decision
#   make install             — go install the binary into $GOPATH/bin
#   make completions         — generate shell-completion scripts (completions/)
#   make docs                — regenerate the Markdown command reference (docs/cli/)
#
#
# Notes:
#   * test-integration runs subtests in parallel (-parallel=8) to stay below
#     Flexera API rate limits while finishing quickly. Adjust JOBS to tune.
#   * Coverage is captured only from the integration run; unit tests touch a
#     much smaller surface and are not merged into the read-only coverage
#     profile.

PKG            ?= ./...
JOBS           ?= 8
# Comma-separated org IDs for test-integration; defaults to the orgs
# integration_registry_test.go uses when FLEXERA_NAM_ORG_IDS is unset (6,1105).
ORG_IDS        ?=
COVER_PROFILE  ?= coverage.integration.out
COVER_HTML     ?= coverage.html
SCRIPTS_DIR    := scripts

# Version stamping (override per build: `make build VERSION=v1.2.3`).
VERSION        ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT         ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
DATE           ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS        := -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.buildDate=$(DATE)

.PHONY: all update-deps update-unified-openapi build install test test-integration coverage-report coverage-html coverage-check check-client-coverage completions docs clean

# all updates Go dependencies and regenerates the CLI from the committed,
# pinned OpenAPI spec.
all: update-deps generate
	@echo "Updated dependencies and regenerated flexera-cli source"

# update-deps upgrades all Go module dependencies to their latest versions
# and tidies/vendors go.mod & go.sum accordingly.
update-deps: update-unified-openapi
	go get -u github.com/flexera-public/unified-go-client@latest
	go mod tidy

# update-unified-openapi downloads an immutable snapshot of the requested
# upstream branch (main by default) and records its commit and SHA-256.
update-unified-openapi:
	./scripts/update-unified-openapi $(REF)

generate: 
	go generate ./...

build:
	go build -ldflags "$(LDFLAGS)" -o flexera-cli .

refresh: generate build
	@echo "Regenerated flexera-cli source, and built flexera-cli binary"

install:
	go install -ldflags "$(LDFLAGS)" .
	$(shell go env GOPATH)/bin/flexera-cli --version

# completions generates shell-completion scripts for the common shells.
completions: build
	mkdir -p completions
	./flexera-cli completion bash       > completions/flexera-cli.bash
	./flexera-cli completion zsh        > completions/flexera-cli.zsh
	./flexera-cli completion fish       > completions/flexera-cli.fish
	./flexera-cli completion powershell > completions/flexera-cli.ps1
	@echo "wrote completions/flexera-cli.{bash,zsh,fish,ps1}"

# docs regenerates the Markdown command reference under docs/cli/.
docs:
	go run ./cmd/gendocs
	@echo "wrote docs/cli/"

test:
	go test $(PKG)

test-integration:
	@if [ -z "$$FLEXERA_NAM_REFRESH_TOKEN" ]; then \
		echo "error: FLEXERA_NAM_REFRESH_TOKEN must be exported to run live integration tests" >&2; \
		exit 1; \
	fi
	FLEXERA_NAM_ORG_IDS=$${FLEXERA_NAM_ORG_IDS:-$(ORG_IDS)} \
		go test -tags=integration \
			-run TestReadOnlyCommandsLive \
			-timeout=20m \
			-parallel=$(JOBS) \
			-covermode=atomic \
			-coverprofile=$(COVER_PROFILE) \
			-v $(PKG)

coverage-report: $(COVER_PROFILE)
	go tool cover -func=$(COVER_PROFILE) | tail -n 40

coverage-html: $(COVER_PROFILE)
	go tool cover -html=$(COVER_PROFILE) -o $(COVER_HTML)
	@echo "wrote $(COVER_HTML)"

coverage-check: $(COVER_PROFILE)
	bash $(SCRIPTS_DIR)/check_coverage.sh $(COVER_PROFILE) $(SCRIPTS_DIR)/readonly_symbols.txt

# check-client-coverage scans unified-go-client for hand-written (non
# generated) exported symbols and fails if any lack a triage entry in
# internal/coverage/unified_client_allowlist.json. This is the
# drift-prevention check for upstream helper functions/types with no
# OpenAPI operation behind them, which cmd/gencli therefore cannot see.
check-client-coverage:
	go run ./cmd/checkcoverage

clean:
	rm -f flexera-cli $(COVER_PROFILE) $(COVER_HTML)
	rm -rf completions
