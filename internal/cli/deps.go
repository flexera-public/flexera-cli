// Package cli holds the cobra+viper foundation shared by every flexera-cli
// command — both the spec-generated tag commands and the hand-written
// curated workflows. It centralizes configuration precedence
// (flag > env > config file > default), HTTP/auth client construction,
// request-body assembly, output rendering, and exit-code mapping so that
// generated leaf commands stay thin.
package cli

import (
	"context"
	"fmt"
	"io"

	cliconfig "github.com/flexera-public/flexera-cli/internal/config"
	cliflexera "github.com/flexera-public/flexera-cli/internal/flexera"
	flexera "github.com/flexera-public/unified-go-client"
)

// Deps is the shared dependency bundle threaded to every command's RunE via
// the cobra command context (see WithDeps/DepsFrom). It is populated once in
// the root command's PersistentPreRunE after viper has merged all config
// sources.
type Deps struct {
	// Config is the resolved common configuration (zone, auth, output, …).
	Config cliconfig.CommonConfig
	// HTTP is the (optionally debug-wrapped) low-level request doer.
	HTTP flexera.HttpRequestDoer
	// Stdout/Stderr are the command's output streams (cobra-injected so
	// tests can capture them).
	Stdout io.Writer
	Stderr io.Writer
	// Getenv reads environment variables. In production it is os.Getenv; in
	// tests it is the injected lookup. Curated commands that honor ad-hoc env
	// vars (e.g. FLEXERA_CLI_GRS_BASE_URL, FLEXERA_CLI_OPTIMA_BASE_URL) read
	// through this rather than os.Getenv directly.
	Getenv func(string) string
	// Printer renders API responses in the configured output format.
	Printer Printer
}

// APIClient builds an authenticated unified client from the resolved config.
func (d *Deps) APIClient() (*flexera.ClientWithResponses, error) {
	if d == nil {
		return nil, fmt.Errorf("cli: nil deps")
	}
	return cliflexera.Factory{HTTPClient: d.HTTP}.NewAPIClient(d.Config)
}

// Factory exposes the lower-level client factory for curated commands that
// need auth-token flows or custom (Optima/GRS) base URLs.
func (d *Deps) Factory() cliflexera.Factory {
	return cliflexera.Factory{HTTPClient: d.HTTP}
}

type depsKey struct{}

// WithDeps returns a child context carrying deps.
func WithDeps(ctx context.Context, deps *Deps) context.Context {
	return context.WithValue(ctx, depsKey{}, deps)
}

// DepsFrom extracts the Deps placed on the context by the root command.
func DepsFrom(ctx context.Context) *Deps {
	if d, ok := ctx.Value(depsKey{}).(*Deps); ok {
		return d
	}
	return nil
}
