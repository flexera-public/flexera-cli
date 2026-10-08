# flexera-cli

Command-line client for Flexera One APIs, including budgets, FinOps costs,
roles, policies, and other resources. Supports JSON/table output, offline
command discovery, request validation, and interactive write previews.

## Experimental Project

This project is currently considered **experimental**.

While we intend to minimize disruption, breaking changes may occur as we continue to evolve the design, APIs, and implementation. **Until the project reaches a stable v1.0.0 release, backward compatibility is not guaranteed.**

We welcome feedback and contributions, but recommend evaluating the current level of stability before adopting this project in production environments.

## Install

```sh
go install github.com/flexera-public/flexera-cli@latest
```

Current source requires **Go 1.26+**. Add `$(go env GOBIN)` to your `PATH` if
set; otherwise add `$(go env GOPATH)/bin` (usually `$HOME/go/bin`).
Prebuilt binaries and Homebrew are not available yet.
For a checkout, see [build from source](CONTRIBUTING.md#build-from-source).

## Quickstart

Help requires no credentials:

```sh
flexera-cli --help
flexera-cli iam --help
```

For live reads, use a token in local `$TOKEN` and replace `12345` with your org ID:

```sh
export FLEXERA_CLI_ACCESS_TOKEN="$TOKEN"
export FLEXERA_CLI_ORG_ID=12345
flexera-cli iam role list -o table
```

The region (API zone) defaults to `nam`. For another region, pass `--zone eu`,
`--zone apac`, or `--zone test`, or set `FLEXERA_CLI_ZONE`.
See [authentication options](docs/usage.md#authentication) for OAuth workflows.

Commands are grouped by API service: `flexera-cli <service> <resource> <action>`,
for example `bill-analysis`, `iam`, `finops-billing`, or `policy`. Run
`flexera-cli --help` for the service list; see [command layout](docs/usage.md#command-layout).

## Preview a write interactively

Generated JSON writes support interactive input and dry-run plans:

```sh
flexera-cli iam organization-invitation create -i --dry-run
```

This prompts for organization/body inputs and prints a redacted plan without
authentication or API calls. Non-interactive requests use typed flags or
`--body @file.json`. See [interactive mode](docs/usage.md#interactive-mode).

## Use it with an agent

[Agent guide](docs/agents.md) · [Optional skill](.github/skills/flexera-cli/SKILL.md)

## Next steps

- [Usage reference](docs/usage.md)
- [Command reference](docs/cli/flexera-cli.md)
- [Contributing and architecture](CONTRIBUTING.md)
- [Issues](https://github.com/flexera-public/flexera-cli/issues)
