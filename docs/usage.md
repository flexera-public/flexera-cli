# Usage guide

See the [quickstart](../README.md#quickstart).

Also use `--help` flag to discover available commands; see the [command reference](cli/flexera-cli.md) and documentation for details.

## Authentication

| Method | Flags | Behavior |
|---|---|---|
| OAuth refresh token | `--refresh-token` | Uses a User Refresh Token credential |
| OAuth client credentials | `--client-id` + `--client-secret` | Uses a Service Account Client ID and Client Secret credential pair |
| Bearer token | `--access-token` | Uses the supplied bearer token without OAuth exchange |

For most use-cases, `--refresh-token` or `--client-id` + `--client-secret` is recommended.

## Configuration

Precedence: **flag > environment > config > default**; discovery has [exceptions](#offline-configuration-safeguards).

### Config file

Default: `~/.flexera/config.yaml`; override with `--config <path>`. Keys are long flag names:

```yaml
zone: nam
org-id: 12345
output: json
# Optional credential keys: access-token, client-id, client-secret, refresh-token.
```

### Environment variables

Persistent flags map to `FLEXERA_CLI_` + uppercase name with `-` replaced by `_`.
Output shaping flags (`--out-jq`, `--out-fields`, `--raw-output`) are CLI-only.

| Variable | Flag | Purpose |
|---|---|---|
| `FLEXERA_CLI_CONFIG` | `--config` | Explicit configuration file path |
| `FLEXERA_CLI_ZONE` | `--zone` | API zone (`nam`/`eu`/`apac`/`test`) |
| `FLEXERA_CLI_ORG_ID` | `--org-id` | Organization ID |
| `FLEXERA_CLI_OUTPUT` | `--output` | `json` or `table` |
| `FLEXERA_CLI_JSON_STYLE` | `--json-style` | `auto`, `pretty`, or `compact` |
| `FLEXERA_CLI_ACCESS_TOKEN` | `--access-token` | Static bearer token |
| `FLEXERA_CLI_CLIENT_ID` | `--client-id` | OAuth client ID |
| `FLEXERA_CLI_CLIENT_SECRET` | `--client-secret` | OAuth client secret |
| `FLEXERA_CLI_REFRESH_TOKEN` | `--refresh-token` | OAuth refresh token |
| `FLEXERA_CLI_API_BASE_URL` | `--api-base-url` | Override the API gateway base URL |
| `FLEXERA_CLI_LOGIN_BASE_URL` | `--login-base-url` | Override the login/OAuth base URL |
| `FLEXERA_CLI_OPTIMA_BASE_URL` | `--optima-base-url` | Override the Optima host (`finops`) |
| `FLEXERA_CLI_GRS_BASE_URL` | `--grs-base-url` | Override the GRS host (`grs`, project resolution) |
| `FLEXERA_CLI_DEBUG` | `--debug`/`-d` | HTTP method/host/path/status to stderr; excludes headers, queries, bodies |

## CLI Helpers

There are several CLI helpers available to simplify common tasks such as: cli search to help find commands, output format and shaping to control command output and display, automatic pagination, and more.  See the sections below for details.

### Search Command

```sh
flexera-cli cli search "list budgets" --read-only
```

Search is lexical, deterministic, and offline. Results include command, summary,
score, confirmation requirement, usage synopsis, and schema link when available.

- `--limit 10` caps search results; `-o table` selects table output.
- `--tag Budget` and `--action delete` filter spec metadata.
- `--read-only` uses HTTP semantics and classified curated side effects, not
  confirmation requirements or `--yes`; unclassified curated workflows are excluded.
  Authentication/output may still involve token or temporary-file activity.
- Uppercase usage placeholders require substitution, including configured org IDs;
  usage strings are not replay commands. Curated-only commands have no schema link.
- JSON results support output shaping.

### Schema Command

```sh
flexera-cli cli schema budget create
```

Inspects parameters, request/response schemas, and examples without auth or network.
Paths resolve against the live tree, including aliases and bill-connect nesting.

- `--part request`, `--part response`, or `--part params` selects one part.
- `--depth` defaults to 3; `--depth 0` preserves references. Cycles/deeper refs remain `$ref`.
- `--example` emits a validated, non-sensitive JSON request body with a newline.
  Missing/invalid/sensitive examples fail; identifiers and server acceptance are not guaranteed.
- `--example` rejects `--part`, table output, and shaping. Introspection is JSON-only.
- Curated workflows without exact operation annotations have no schema; use `--help`.

### Output format flag

Most commands default to JSON. `--output table` / `-o table` requires a table
renderer (primarily list commands). `--json-style auto` is pretty on a terminal
and compact when redirected; `pretty` or `compact` overrides it.

### Output shaping flag

Shaping applies only to structured JSON: `--out-fields` projects before
`--out-jq` evaluates. `--raw-output` / `-r` requires `--out-jq`.

```sh
flexera-cli cli search "list budgets" --read-only --out-fields command,summary
```

### Automatic Pagination flag

- CLI will follow paginated lists [`nextPage`] and return full results automatically.
- `--no-paginate` fetches only the a single page.
- API `--limit`, where exposed, sets page size, not a total cap or pagination stop.

### Errors and exit codes

Errors are unshaped JSON on stderr, including pre-configuration/auth failures.

| Exit code | Meaning |
|---|---|
| **0** | Success/help (help is text) |
| **1** | Runtime/API failure |
| **2** | Syntax/configuration, schema lookup, or jq compilation failure |

### Interactive mode

Generate JSON write commands support local `--interactive` / `-i`:

```sh
flexera-cli tag-dimension create -i --dry-run
```

- Input and stderr must be terminals; stdout may be redirected. Required parameters
  come first; inline/file bodies and typed flags prefill editable fields.
- `-i --body @-` is rejected: stdin is reserved for prompts. Reads, raw uploads,
  and unadapted curated writes reject `-i`; prompting is never implicit.
- The complete redacted plan appears on stderr. Exact lowercase `yes` proceeds;
  other input cancels without an API call. `-i --yes` skips only final confirmation.
- `-i --dry-run` emits human stderr and JSON stdout plans without auth, API-client
  construction, or final confirmation.
- Unsupported/deep schemas require `$EDITOR` as a single executable path. Private
  editor files exclude secret-valued fields; unsafe schemas require `--body @file`.
- `FLEXERA_CLI_ACCESSIBLE=1` enables accessible prompts. Sensitive plans omit shell
  replay commands; explicit save writes the exact body to a new file with mode `0600`.

## Shell completion

Load completion in the current shell:

```sh
# Bash (current shell)
source <(flexera-cli completion bash)
# Zsh (current shell)
source <(flexera-cli completion zsh)
# Fish
flexera-cli completion fish | source
# PowerShell
flexera-cli completion powershell | Out-String | Invoke-Expression
```

`make completions` writes all four scripts to `completions/`; see
[contributor regeneration guidance](../CONTRIBUTING.md#regeneration).

## Examples

Set local authentication and replace `12345` with your org ID:

```sh
export FLEXERA_CLI_REFRESH_TOKEN="$TOKEN"
export FLEXERA_CLI_ORG_ID=12345
flexera-cli role list
flexera-cli role list -o table
flexera-cli policy applied-policy list
```

Commands accept typed flags or sometimes `--body`. `budget.json` is a user-supplied request file:

```sh
flexera-cli budget create --body @budget.json --dry-run
flexera-cli budget delete --id BUDGET_ID --dry-run
```