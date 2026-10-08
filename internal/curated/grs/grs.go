// Package grs provides `grs project list-for-org`, a workflow over
// flexera.ProjectResolver: it derives the user from the access token, calls
// the zone-specific GRS host, filters projects to the org, and reports legacy
// account IDs (the project IDs Policy APIs expect). The generated
// `grs project list` exposes the raw endpoint.
package grs

import (
	"github.com/spf13/cobra"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	flexera "github.com/flexera-public/unified-go-client"
)

// Attach adds the curated project workflow beneath the generated
// `grs project` command.
func Attach(root *cobra.Command) error {
	return clipkg.AttachCommands(root, []string{"grs", "project"}, newProjectListCmd())
}

func newProjectListCmd() *cobra.Command {
	var grsBaseURL, apiVersion string
	c := &cobra.Command{
		Use:   "list-for-org",
		Short: "List the org's projects for the authenticated user (policy-ready project IDs)",
		Long: "List the projects the authenticated user can access in --org-id. The user is derived " +
			"from the access token, results are filtered to the org, and each project ID is the " +
			"legacy account ID that the Policy APIs accept as --project-id.",
		Example: "flexera-cli grs project list-for-org --org-id 123",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			client, err := deps.Factory().NewAPIClientForBaseURL(deps.Config, deps.GRSBaseURL(grsBaseURL))
			if err != nil {
				return err
			}
			opts, err := deps.Factory().ProjectResolverOptions(deps.Config)
			if err != nil {
				return err
			}
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
