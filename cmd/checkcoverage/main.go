// Command checkcoverage scans unified-go-client for hand-written
// (non-generated) exported symbols and fails if any lack a corresponding
// entry in internal/coverage/unified_client_allowlist.json. It is the
// drift-prevention check for functions/types added upstream that have no
// OpenAPI operation -- and therefore no generated CLI command -- backing
// them.
//
// Run via `make check-coverage` or `go run ./cmd/checkcoverage`. On finding
// new symbols, add a curated command and an allowlist entry with
// status "covered", or, if the symbol is infrastructure that needs no CLI
// command, add an entry with status "excluded" and a Reason.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"encoding/json"

	"github.com/flexera-public/flexera-cli/internal/coverage"
)

func main() {
	modulePath := flag.String("module", "github.com/flexera-public/unified-go-client", "module to scan")
	dir := flag.String("dir", "", "local checkout to scan instead of resolving -module from go.mod (for testing pre-publish upstream changes)")
	allowlistPath := flag.String("allowlist", "internal/coverage/unified_client_allowlist.json", "path to the checked-in allowlist")
	flag.Parse()

	moduleDir := *dir
	if moduleDir == "" {
		resolved, err := resolveModuleDir(*modulePath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "checkcoverage:", err)
			os.Exit(2)
		}
		moduleDir = resolved
	}

	symbols, err := coverage.Scan(moduleDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "checkcoverage: scan:", err)
		os.Exit(2)
	}

	allow, err := coverage.LoadAllowlist(*allowlistPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "checkcoverage:", err)
		os.Exit(2)
	}

	diff := coverage.DiffAgainstAllowlist(symbols, allow)

	if len(diff.Stale) > 0 {
		fmt.Println("stale allowlist entries (no longer found upstream; safe to remove):")
		for _, e := range diff.Stale {
			fmt.Printf("  - %s\n", e.Symbol)
		}
		fmt.Println()
	}

	if len(diff.New) == 0 {
		fmt.Printf("checkcoverage: OK -- %d hand-written symbols, all triaged\n", len(symbols))
		return
	}

	fmt.Printf("checkcoverage: FAIL -- %d hand-written unified-go-client symbol(s) are not in %s:\n\n", len(diff.New), *allowlistPath)
	for _, s := range diff.New {
		fmt.Printf("  - %s (%s, declared in %s)\n", s.Key(), s.Kind, s.File)
	}
	fmt.Println(`Each of these is new (or newly hand-written) API in unified-go-client with
no OpenAPI operation behind it, so cmd/gencli cannot generate a command for
it. For each one, either:

  1. Add a curated CLI command under internal/curated/ that uses it, then
     add an allowlist entry: {"symbol": "...", "status": "covered", "command": "..."}
  2. If it's infrastructure with no standalone CLI use, add an allowlist
     entry: {"symbol": "...", "status": "excluded", "reason": "..."}`)
	os.Exit(1)
}

// resolveModuleDir shells out to `go list -m -json <modulePath>` to find the
// on-disk directory for the version currently pinned in go.mod (module
// cache or a replace target). This is deliberately not go/packages so the
// tool has no extra dependency beyond the stdlib.
func resolveModuleDir(modulePath string) (string, error) {
	cmd := exec.Command("go", "list", "-m", "-json", modulePath)
	out, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return "", fmt.Errorf("go list -m -json %s: %s", modulePath, strings.TrimSpace(string(ee.Stderr)))
		}
		return "", fmt.Errorf("go list -m -json %s: %w", modulePath, err)
	}
	var mod struct {
		Dir string
	}
	if err := json.Unmarshal(out, &mod); err != nil {
		return "", fmt.Errorf("parse go list output: %w", err)
	}
	if mod.Dir == "" {
		return "", fmt.Errorf("module %s has no resolved Dir (not downloaded?); run `go mod download`", modulePath)
	}
	return filepath.Clean(mod.Dir), nil
}
