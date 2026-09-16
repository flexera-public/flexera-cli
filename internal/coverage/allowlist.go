package coverage

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// AllowlistEntry records the CLI-coverage disposition of one hand-written
// unified-go-client symbol.
type AllowlistEntry struct {
	// Symbol is Symbol.Key(): "<package>#<name>".
	Symbol string `json:"symbol"`
	// Status is "covered" (a curated CLI command uses it) or "excluded"
	// (deliberately not exposed; Reason must explain why).
	Status string `json:"status"`
	// Command is the CLI command that covers this symbol. Required when
	// Status is "covered".
	Command string `json:"command,omitempty"`
	// Reason explains why an "excluded" symbol needs no CLI command.
	Reason string `json:"reason,omitempty"`
}

// Allowlist is the checked-in inventory of every hand-written symbol this
// repo has already triaged, keyed by AllowlistEntry.Symbol.
type Allowlist map[string]AllowlistEntry

// LoadAllowlist reads and decodes the JSON allowlist at path.
func LoadAllowlist(path string) (Allowlist, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read allowlist %s: %w", path, err)
	}
	var entries []AllowlistEntry
	if err := json.Unmarshal(b, &entries); err != nil {
		return nil, fmt.Errorf("parse allowlist %s: %w", path, err)
	}
	out := make(Allowlist, len(entries))
	for _, e := range entries {
		if e.Status != "covered" && e.Status != "excluded" {
			return nil, fmt.Errorf("allowlist %s: symbol %q has invalid status %q (want covered|excluded)", path, e.Symbol, e.Status)
		}
		if e.Status == "covered" && e.Command == "" {
			return nil, fmt.Errorf("allowlist %s: symbol %q is covered but has no command", path, e.Symbol)
		}
		if e.Status == "excluded" && e.Reason == "" {
			return nil, fmt.Errorf("allowlist %s: symbol %q is excluded but has no reason", path, e.Symbol)
		}
		out[e.Symbol] = e
	}
	return out, nil
}

// Diff compares found (the current scan of unified-go-client) against the
// allowlist and reports every symbol that needs a maintainer decision
// (new/unacknowledged) plus every allowlist entry that no longer exists
// upstream (stale, informational only).
type Diff struct {
	New   []Symbol
	Stale []AllowlistEntry
}

func DiffAgainstAllowlist(found []Symbol, allow Allowlist) Diff {
	var diff Diff
	seen := make(map[string]bool, len(found))
	for _, s := range found {
		seen[s.Key()] = true
		if _, ok := allow[s.Key()]; !ok {
			diff.New = append(diff.New, s)
		}
	}
	for key, entry := range allow {
		if !seen[key] {
			diff.Stale = append(diff.Stale, entry)
		}
	}
	sort.Slice(diff.Stale, func(i, j int) bool { return diff.Stale[i].Symbol < diff.Stale[j].Symbol })
	return diff
}
