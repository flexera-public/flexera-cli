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