#!/usr/bin/env bash
# check_coverage.sh — enforce 100% line coverage on a curated allowlist of
# read-only command handlers after a successful `make test-integration` run.
#
# Usage:
#   bash scripts/check_coverage.sh <coverage.out> <readonly_symbols.txt>
#
# Exit codes:
#   0 — every allowlisted symbol reported 100.0% coverage.
#   1 — usage error (missing/unreadable inputs).
#   2 — at least one allowlisted symbol was missing from the coverage profile.
#   3 — at least one allowlisted symbol reported < 100.0% coverage.

set -euo pipefail

if [[ $# -ne 2 ]]; then
    echo "usage: $0 <coverage-profile> <symbols-allowlist>" >&2
    exit 1
fi

PROFILE="$1"
ALLOWLIST="$2"

if [[ ! -r "$PROFILE" ]]; then
    echo "error: coverage profile not readable: $PROFILE" >&2
    exit 1
fi
if [[ ! -r "$ALLOWLIST" ]]; then
    echo "error: allowlist not readable: $ALLOWLIST" >&2
    exit 1
fi

# Build a temporary FUNC-coverage report once. `go tool cover -func` emits
# lines of the form:
#   github.com/.../path/file.go:LINE:	FuncName	100.0%
tmp_func=$(mktemp)
trap 'rm -f "$tmp_func"' EXIT
go tool cover -func="$PROFILE" > "$tmp_func"

missing=()
below=()
ok=0

# Read allowlist line-by-line, ignoring comments and blank lines.
while IFS= read -r raw || [[ -n "$raw" ]]; do
    # Strip leading/trailing whitespace.
    sym="${raw#"${raw%%[![:space:]]*}"}"
    sym="${sym%"${sym##*[![:space:]]}"}"
    [[ -z "$sym" || "${sym:0:1}" == "#" ]] && continue

    # Match column-2 exactly. Coverage rows are tab-separated, but `go tool
    # cover -func` uses runs of whitespace; awk handles both.
    line=$(awk -v target="$sym" '$2 == target { print; exit }' "$tmp_func")
    if [[ -z "$line" ]]; then
        missing+=("$sym")
        continue
    fi

    pct=$(awk '{print $NF}' <<<"$line" | tr -d '%')
    if [[ "$pct" != "100.0" ]]; then
        below+=("$sym ($pct%)")
        continue
    fi

    ok=$((ok + 1))
done < "$ALLOWLIST"

echo "coverage gate: $ok allowlisted symbols at 100.0%"

fail=0
if [[ ${#missing[@]} -gt 0 ]]; then
    echo
    echo "MISSING from coverage profile (handler never executed):" >&2
    printf '  %s\n' "${missing[@]}" >&2
    fail=2
fi
if [[ ${#below[@]} -gt 0 ]]; then
    echo
    echo "BELOW 100% coverage:" >&2
    printf '  %s\n' "${below[@]}" >&2
    fail=3
fi

exit "$fail"
