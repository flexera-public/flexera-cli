// Package grs provides the hand-written "grs project" command for the Flexera
// Governance / Resource Service.
//
// GRS is a legacy API: IAM endpoints are preferred wherever they exist. The
// only GRS capability retained here is listing an org's projects, which has no
// IAM equivalent. The project-listing logic lives in the library
// (flexera.ProjectResolver); this command is thin cobra wiring over it. The
// former `grs user orgs list` was replaced by the IAM-backed `user-orgs`
// command.
package grs

import (
	"fmt"

	"github.com/spf13/cobra"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	flexera "github.com/flexera-public/unified-go-client"
)

// NewCmd builds the "grs" command tree.
func NewCmd() *cobra.Command {
	c := &cobra.Command{
		Use:     "grs",
		Short:   "Governance / Resource Service (legacy; projects only)",
		Example: "flexera-cli grs project list --org-id 123",
		RunE:    parentRunE,
	}
	project := &cobra.Command{Use: "project", Short: "GRS projects", RunE: parentRunE}
	project.AddCommand(newProjectListCmd())
	c.AddCommand(project)
	return c
}

func parentRunE(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return cmd.Help()
	}
	return fmt.Errorf("unknown command %q for %q", args[0], cmd.CommandPath())
}

func newProjectListCmd() *cobra.Command {
	var grsBaseURL, apiVersion string
	c := &cobra.Command{
		Use:   "list",
		Short: "List GRS projects for the org",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			client, err := deps.Factory().NewAPIClientForBaseURL(deps.Config, deps.GRSBaseURL(grsBaseURL))
			if err != nil {
				return err
			}
			opts := []flexera.ProjectResolverOption{}
			if apiVersion != "" {
				opts = append(opts, flexera.WithProjectsAPIVersion(apiVersion))
			}
			resolver, err := flexera.NewProjectResolver(client, opts...)
			if err != nil {
				return err
			}
			projects, err := resolver.ListProjects(cmd.Context(), int64(deps.Config.OrgID))
			if err != nil {
				return err
			}
			return deps.Printer.Render(deps.Stdout, deps.Config.Output, projects)
		},
	}
	c.Flags().StringVar(&grsBaseURL, "grs-base-url", "", "Override the GRS base URL (default: zone-specific grs-front host)")
	c.Flags().StringVar(&apiVersion, "api-version", "", "Optional X-Api-Version header (defaults to 2.0)")
	return c
}
