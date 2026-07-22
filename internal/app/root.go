// Package app wires the flexera-cli cobra command tree. It is kept separate
// from package main so that tooling (e.g. cmd/gendocs) can build the full
// command tree without pulling in main's process entrypoint.
package app

import (
	"io"

	"github.com/spf13/cobra"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	"github.com/flexera-public/flexera-cli/internal/commands"
	authcmd "github.com/flexera-public/flexera-cli/internal/curated/auth"
	finopscmd "github.com/flexera-public/flexera-cli/internal/curated/finops"
	grscmd "github.com/flexera-public/flexera-cli/internal/curated/grs"
	policycmd "github.com/flexera-public/flexera-cli/internal/curated/policy"
	userorgscmd "github.com/flexera-public/flexera-cli/internal/curated/userorgs"
	workflowscmd "github.com/flexera-public/flexera-cli/internal/curated/workflows"
	flexera "github.com/flexera-public/unified-go-client"
)

// NewRootCmd constructs the root command: persistent flags + viper config via
// internal/cli, the spec-generated tag commands (internal/commands), and the
// hand-written curated commands (auth, grs, curated workflows, finops, policy,
// user-orgs) that wrap library operations the spec can't express.
//
// It is the single source of truth for the command tree, shared by the binary
// (package main) and the docs generator (cmd/gendocs).
func NewRootCmd(stdout, stderr io.Writer, getenv func(string) string, base flexera.HttpRequestDoer, version string) (*cobra.Command, *clipkg.Deps) {
	root, deps := clipkg.NewRootCmd(clipkg.RootOptions{
		Use:      "flexera-cli",
		Short:    "Flexera One unified API command-line client",
		BaseHTTP: base,
		Version:  version,
		Getenv:   getenv,
	})
	root.SetOut(stdout)
	root.SetErr(stderr)

	commands.RegisterAll(root)

	for _, c := range []*cobra.Command{
		authcmd.NewCmd(),
		grscmd.NewCmd(),
		workflowscmd.NewCmd(),
		finopscmd.NewCmd(),
		policycmd.NewCmd(),
		userorgscmd.NewCmd(),
	} {
		root.AddCommand(c)
	}

	return root, deps
}
