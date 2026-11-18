# Contributing

## Build from source

Current source requires **Go 1.27.3 or newer**, Git, and Make.

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

### Generator failures

`regencli` preflights all annotated, non-curated operations before generating any
packages. Unsupported operations are reported together with HTTP method, path,
reason, and remediation. `gencli` also validates before writing either source or
metadata. To check a spec without generating files:

```sh
go run ./cmd/gencli -spec unified-openapi/openapi3.json -check
```

For incorrect/missing operation IDs, actions, parameters, or responses, fix the
upstream OpenAPI source/annotations and regenerate unified-openapi and
unified-go-client together, then run `make refresh`. Concrete 2xx responses are
supported; wildcard/default responses cannot select SDK response fields. GET
redirects declaring a `Location` header require an upstream SDK download helper
and an exact curated CLI wrapper rather than generated 2xx/JSON handling.
API-specific redirect handling, polling, and report filter serialization belong
in unified-go-client; CLI wrappers should only adapt flags, confirmation, and
output. Publish the SDK helpers and update the CLI dependency before wiring
those wrappers. For unsupported shapes, extend the relevant collector
and renderer in `cmd/gencli` with regression tests. If the operation needs custom
behavior, implement an exact curated wrapper and register its operation ID in
`curatedOperations` in `cmd/regencli/coverage.go`; do not silently exclude it or
edit generated commands. Unknown service errors identify the `serviceCommands`
entry to add in `cmd/regencli/services.go`.

## Architecture

The Cobra command tree calls the [unified Go SDK](https://github.com/flexera-public/unified-go-client).
The CLI handles configuration, discovery, validation, output, and interactive input.
Generated commands mirror APIs and are grouped by the `x-flexera-service` metadata in
the unified spec (`internal/commands/<service>/<tag>/`; CLI names, aliases, hoisting,
and renames live in `cmd/regencli/services.go`). Operations beneath a tag are named
`[<sub-resource>...] <action>` from their HTTP path by `cmd/gencli/naming.go`; generation fails
rather than guess when names collide, so resolve conflicts there with `commandOverrides`.
Also use `commandOverrides` when a derived CRUD verb misdescribes an operation (for example
`resend` or `reprocess` instead of `update`). When renaming an existing command, keep the old
spelling working: add it to `commandAliases` (leaves) or `TagAliases` (tag commands);
generation rejects aliases that collide with sibling commands. Curated commands add behavior beyond a
single API call (auth, FinOps cost queries, policy resolution, organization
discovery, anomalies) and attach beneath their service via `Attach` in
`internal/app/root.go`; do not add curated commands that only duplicate a generated one. Use the SDK directly for other integrations.