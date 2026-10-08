// Command gendocs generates the flexera-cli command reference as Markdown
// under docs/cli/. Run via `make docs` or `go run ./cmd/gendocs`.
package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"

	"github.com/flexera-public/flexera-cli/internal/app"
)

const outDir = "docs/cli"

func main() {
	root, _ := app.NewRootCmd(io.Discard, io.Discard, os.Getenv, &http.Client{}, "dev")

	// Stable output: no per-file generation timestamp.
	disableAutoGenTag(root)

	if err := os.MkdirAll(outDir, 0o755); err != nil {
		log.Fatalf("gendocs: %v", err)
	}
	// Every page is generated; prune so renamed or removed commands do not
	// leave stale pages behind.
	stale, err := filepath.Glob(filepath.Join(outDir, root.Name()+"*.md"))
	if err != nil {
		log.Fatalf("gendocs: %v", err)
	}
	for _, path := range stale {
		if err := os.Remove(path); err != nil {
			log.Fatalf("gendocs: %v", err)
		}
	}
	if err := doc.GenMarkdownTree(root, outDir); err != nil {
		log.Fatalf("gendocs: %v", err)
	}
}

// disableAutoGenTag turns off cobra's "Auto generated ... on <date>" footer on
// every command so the generated Markdown is deterministic.
func disableAutoGenTag(cmd *cobra.Command) {
	cmd.DisableAutoGenTag = true
	for _, c := range cmd.Commands() {
		disableAutoGenTag(c)
	}
}
