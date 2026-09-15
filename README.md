# flexera-cli

A command-line client for the **Flexera One unified API**.

`flexera-cli` is a thin, exemplary consumer of the unified Flexera Go client library ([`github.com/flexera-public/unified-go-client`](https://github.com/flexera-public/unified-go-client)). 

All business logic, helpers, and curated operations live in that library; this binary only wires a [cobra](https://github.com/spf13/cobra) command tree — **most of it generated directly from the unified OpenAPI spec** — to it.  If you are building another integration (a Terraform provider, a web app, an internal tool), prefer importing the Go client library directly; this CLI shows how.

- **Generated commands** mirror the API surface (`budget`, `role`,
  `published-template`, `vulnerability`, …) with consistent verbs
  (`list`, `get`, `create`, `update`, `replace`, `delete`).
- **Curated commands** wrap multi-step or non-spec workflows: `auth`,
  `finops` (Optima cost analytics), `policy` (with project auto-resolution),
  `grs`, `user-orgs`, and `curated` (e.g. AI cost-anomaly investigation).

## Experimental Project

This project is currently considered **experimental**.

While we intend to minimize disruption, breaking changes may occur as we continue to evolve the design, APIs, and implementation. **Until the project reaches a stable v1.0.0 release, backward compatibility is not guaranteed.**

We welcome feedback and contributions, but recommend evaluating the current level of stability before adopting this project in production environments.

---

## Install

### `go install`

```sh
go install github.com/flexera-public/flexera-cli@latest
```

### Build from source

```sh
git clone https://github.com/flexera-public/flexera-cli
cd flexera-cli
make build            # version-stamped binary in ./flexera-cli
./flexera-cli --version
```

`make build` stamps the version/commit/build-date via ldflags. Override the
version explicitly with `make build VERSION=v1.2.3`.

### Release binaries / Homebrew

> Pre-built release binaries and a Homebrew tap are not yet published. For now,
> use `go install` or `make build`.

---

## Quick start

```sh
# Authenticate-free help is always available:
flexera-cli --help
flexera-cli budget --help

# List budgets for an org using a static access token:
flexera-cli budget list --org-id 12345 --access-token "$TOKEN"

# Same, but reading config from the environment (see Configuration):
export FLEXERA_CLI_ORG_ID=12345
export FLEXERA_CLI_ACCESS_TOKEN="$TOKEN"
flexera-cli budget list
```

---

## Authentication

Every command except `auth` accepts one of:

- **Static bearer token** — `--access-token` (or `FLEXERA_CLI_ACCESS_TOKEN`).
- **OAuth client credentials** — `--client-id` + `--client-secret`. The CLI
  obtains an access token automatically and refreshes it during longer
  workflows. Supplying only one of the pair fails fast.
- **OAuth refresh token** — `--refresh-token`.

The `auth token` command mints a token you can reuse:

```sh
flexera-cli auth token client-credentials --client-id "$ID" --client-secret "$SECRET"
flexera-cli auth token refresh --refresh-token "$REFRESH"
```

The API **zone** defaults to `nam`; override with `--zone nam|eu|apac|test`
(or `FLEXERA_CLI_ZONE`).

---

## Configuration

Configuration is resolved with the precedence **flag > environment variable >
config file > built-in default**.

### Config file

By default the CLI reads `~/.flexera/config.yaml` (override with `--config
<path>`). Keys are the long flag names:

```yaml
# ~/.flexera/config.yaml
zone: nam
org-id: 12345
output: json
access-token: "your-token"        # or use client-id / client-secret / refresh-token
```

### Environment variables

Every persistent flag has an environment variable: prefix `FLEXERA_CLI_`,
upper-cased, with `-` replaced by `_`. For example `--org-id` →
`FLEXERA_CLI_ORG_ID`, `--api-base-url` → `FLEXERA_CLI_API_BASE_URL`.

| Variable | Flag | Purpose |
|---|---|---|
| `FLEXERA_CLI_ZONE` | `--zone` | API zone (`nam`/`eu`/`apac`/`test`) |
| `FLEXERA_CLI_ORG_ID` | `--org-id` | Organization ID |
| `FLEXERA_CLI_OUTPUT` | `--output` | `json` or `table` |
| `FLEXERA_CLI_ACCESS_TOKEN` | `--access-token` | Static bearer token |
| `FLEXERA_CLI_CLIENT_ID` | `--client-id` | OAuth client ID |
| `FLEXERA_CLI_CLIENT_SECRET` | `--client-secret` | OAuth client secret |
| `FLEXERA_CLI_REFRESH_TOKEN` | `--refresh-token` | OAuth refresh token |
| `FLEXERA_CLI_API_BASE_URL` | `--api-base-url` | Override the API gateway base URL |
| `FLEXERA_CLI_LOGIN_BASE_URL` | `--login-base-url` | Override the login/OAuth base URL |
| `FLEXERA_CLI_OPTIMA_BASE_URL` | `--optima-base-url` | Override the Optima host (`finops`) |
| `FLEXERA_CLI_GRS_BASE_URL` | `--grs-base-url` | Override the GRS host (`grs`, project resolution) |
| `FLEXERA_CLI_DEBUG` | `--debug`/`-d` | Log HTTP requests/responses (Authorization redacted) |

---

## Output formats

Most commands emit JSON by default. Use `--output table` (or `-o table`) where a
human-readable table renderer exists (primarily list commands):

```sh
flexera-cli budget list --org-id 12345 -o table
```

Paginated list endpoints follow `nextPage` automatically and merge results; pass
`--no-paginate` to fetch only the first page, or `--skip-token <token>` to
resume.

---

## Shell completion

Cobra-powered completion is built in. Generate a script with
`flexera-cli completion <bash|zsh|fish|powershell>`.

```sh
# Bash (current shell)
source <(flexera-cli completion bash)
# Bash (persistent, Linux)
flexera-cli completion bash | sudo tee /etc/bash_completion.d/flexera-cli >/dev/null

# Zsh (current shell)
source <(flexera-cli completion zsh)
# Zsh (persistent) — ensure `compinit` is loaded, then:
flexera-cli completion zsh > "${fpath[1]}/_flexera-cli"

# Fish
flexera-cli completion fish | source
flexera-cli completion fish > ~/.config/fish/completions/flexera-cli.fish

# PowerShell
flexera-cli completion powershell | Out-String | Invoke-Expression
```

`make completions` writes all four scripts to `completions/`.

---

## Command reference

The full, per-command reference (flags, usage, subcommands) is generated from
the binary into [`docs/cli/`](docs/cli/flexera-cli.md). Regenerate it with:

```sh
make docs        # runs: go run ./cmd/gendocs  ->  docs/cli/*.md
```

Or explore interactively with `flexera-cli <command> --help`.

---

## Examples

```sh
# Read-only / generated commands
flexera-cli role list --org-id 12345 --access-token "$TOKEN"
flexera-cli vulnerability list --org-id 12345 --access-token "$TOKEN" -o table

# Write commands support a dry-run preview and require --yes for destructive ops
flexera-cli budget create --org-id 12345 --body @budget.json --dry-run
flexera-cli budget delete --org-id 12345 --id b-1 --yes

# Request bodies: typed flags where the schema is simple, or --body for the rest
flexera-cli tag-dimension create --org-id 12345 --name "Environment" --body @dim.json

# Curated: FinOps cost analytics (Optima) — auto-resolves billing centers,
# auto-routes /aggregated vs /select, and chunks >24-month windows
flexera-cli finops cost get --org-id 12345 --start 2024-01 --end 2024-04

# Curated: policy commands auto-resolve the project via GRS when --project-id is omitted
flexera-cli policy applied-policy list --org-id 12345 --access-token "$TOKEN"

# Curated: discover which orgs your credentials can access (user ID auto-detected from the JWT)
flexera-cli user-orgs list --refresh-token "$REFRESH"

# Curated: AI-powered cost anomaly investigation
flexera-cli curated anomaly-investigation --org-id 12345 --input '{"...":"..."}'
```

Add `--debug` to any command to trace HTTP requests/responses on stderr
(the `Authorization` header is redacted).

---

## Testing & contributing

```sh
make test                 # unit tests (no network, no credentials)
go vet ./...
make docs completions     # regenerate reference docs + completion scripts
```

The generated command tree is produced from the unified OpenAPI spec via
`go generate ./...` (see `cmd/gencli` and `cmd/regencli`); do not hand-edit
files under `internal/commands/`. Live, read-only integration tests are
available via `make test-integration` (requires `FLEXERA_NAM_REFRESH_TOKEN`).

The CLI is a thin shell over the unified Go client library at
[`github.com/flexera-public/unified-go-client`](https://github.com/flexera-public/unified-go-client); prefer
importing that library directly when building other integrations.
