// Command flexera-cli is a thin, exemplary consumer of the Flexera unified
// Go client library (github.com/flexera-public/unified-go-client). All business logic,
// helpers, and curated operations live in the library; this binary only
// wires a cobra command tree (mostly generated from the OpenAPI spec) to it.
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/flexera-public/flexera-cli/internal/app"
	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	flexera "github.com/flexera-public/unified-go-client"
)

// Stamped via -ldflags at build time (see Makefile).
var (
	version   = "dev"
	commit    = "unknown"
	buildDate = "unknown"
)

func main() {
	os.Exit(run(context.Background(), os.Args[1:], os.Stdout, os.Stderr, os.Getenv, &http.Client{}))
}

// run builds the cobra root and executes it. It is the testable entrypoint:
// args/streams/getenv/httpClient are injected.
func run(ctx context.Context, args []string, stdout, stderr io.Writer, getenv func(string) string, httpClient flexera.HttpRequestDoer) int {
	root, deps := app.NewRootCmd(stdout, stderr, getenv, httpClient, versionString())
	root.SetArgs(args)
	return clipkg.Execute(ctx, root, deps)
}

// versionString composes the --version output from the ldflag-stamped vars.
func versionString() string {
	return fmt.Sprintf("%s (commit %s, built %s)", version, commit, buildDate)
}
