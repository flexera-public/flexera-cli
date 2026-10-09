# Agent setup

Requires terminal access to an installed `flexera-cli`; no MCP server is needed.

## Install the skill

Copy [.github/skills/flexera-cli/](../.github/skills/flexera-cli/SKILL.md), containing
only `SKILL.md`, into a supported location. CLI installation does not install it.

| Agent environment | Project location | Personal location |
|---|---|---|
| GitHub Copilot in VS Code | `.github/skills/flexera-cli/` | `~/.copilot/skills/flexera-cli/` |
| Agents supporting `.agents/skills` | `.agents/skills/flexera-cli/` | Agent-specific |
| Claude Code | `.claude/skills/flexera-cli/` | `~/.claude/skills/flexera-cli/` |

Keep the folder and filename unchanged. In VS Code, invoke `/flexera-cli` from
chat; this checkout already includes the project skill.

## Without skill support

Use this task instruction:

> Discover commands with `flexera-cli cli search`; use `--read-only` for reads.
> Inspect help/schema and confirm organization and zone. Use local credentials;
> never expose secrets. Preview writes with `--dry-run` where supported and obtain
> approval of the target and payload before execution. Do not blindly retry writes.

Use scoped credentials and terminal approval controls; a skill does not enforce
permissions. See [authentication](usage.md#authentication).

## Resolve command inputs

Read `--help` for flag guidance, then inspect `cli schema <command path>` for
the full input contract. `params[].name` is the API parameter name;
`params[].flag` is the CLI spelling. For typed body flags, use `bodyFields` to
map each flag to its exact JSON `property`, schema, and requiredness. A renamed
flag such as `--body-org-id` does not rename the JSON property `orgId`.

API defaults do not mean the CLI sends those values. Examples are illustrative,
not live resource IDs; `requestExampleSource` identifies upstream versus
synthesized bodies. Header metadata describes the API contract, not an
arbitrary-header CLI interface. Missing discovery relationships, filter
grammars, or prerequisites must not be invented from examples.