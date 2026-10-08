# Contributing

## Build from source

Current source requires **Go 1.26+**, Git, and Make.

```sh
git clone https://github.com/flexera-public/flexera-cli
cd flexera-cli
make build
```

`make build` writes `./flexera-cli`, stamping version, commit, and UTC build date
via ldflags. Override with `make build VERSION=v1.2.3`, `COMMIT`, or `DATE`.

## Validation

```sh
make test                 # unit tests: no network, no credentials
go vet ./...
```

Release checks require the published SDK: disable workspaces and active local `replace` directives in `go.mod`.

```sh
GOWORK=off go test . -run '^TestPrecisionGenerated(Pagination|UntypedResponse)$' -count=1 -v
GOWORK=off go test ./...
```

`make test-integration` runs live, read-only API tests; requires local `FLEXERA_NAM_REFRESH_TOKEN`.

## Regeneration

Commands derive from `unified-openapi/openapi3.json`. **Do not manually edit**
`internal/commands/`, generated catalog artifacts, `docs/cli/`, or completion
scripts. Change the generator/source of truth, regenerate, and review the diff.

```sh
make generate                        # go generate ./... invokes cmd/regencli
make completions docs                # shell scripts + command reference
make update-unified-openapi           # pin upstream main
make update-unified-openapi REF=<branch>
make all                             # update dependencies/spec, then generate
```

Spec updates record repository, commit, and SHA-256 for an immutable snapshot in `unified-openapi/PIN`.
`make all` updates dependencies/spec; it is not validation. `cmd/gencli` generates commands;
`cmd/regencli` orchestrates publication, `cmd/gendocs` writes docs, and completions go to `completions/`.

## Architecture

The Cobra command tree calls the [unified Go SDK](https://github.com/flexera-public/unified-go-client).
The CLI handles configuration, discovery, validation, output, and interactive input.
Generated commands mirror APIs and are grouped by the `x-flexera-service` metadata in
the unified spec (`internal/commands/<service>/<tag>/`; CLI names, aliases, hoisting,
and renames live in `cmd/regencli/services.go`). Curated commands add behavior beyond a
single API call (auth, FinOps cost queries, policy resolution, GRS, organization
discovery, anomalies) and attach beneath their service via `Attach` in
`internal/app/root.go`; do not add curated commands that only duplicate a generated one. Use the SDK directly for other integrations.