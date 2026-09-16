// Package graphql provides the non-OpenAPI GraphQL commands: the raw
// /explore/graphql query helper (flexera.(*Client).GraphQL) and the query
// generate/modify helper (service/graphql/v1's GenerateQuery).
package graphql

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	graphqllib "github.com/flexera-public/unified-go-client/service/graphql/v1"
	"github.com/spf13/cobra"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	flexera "github.com/flexera-public/unified-go-client"
)

// baseClient extracts the concrete *flexera.Client so its Server, HTTP doer,
// and RequestEditors can be reused by the standalone graphql sub-packages
// and hand-written helpers that live outside the generated ClientInterface.
func baseClient(deps *clipkg.Deps) (*flexera.Client, error) {
	apiClient, err := deps.APIClient()
	if err != nil {
		return nil, err
	}
	base, ok := apiClient.ClientInterface.(*flexera.Client)
	if !ok {
		return nil, fmt.Errorf("graphql: unsupported API client implementation")
	}
	return base, nil
}

// NewQueryCmd builds the raw GraphQL query command backed by
// flexera.(*Client).GraphQL, which POSTs to /explore/graphql. This endpoint
// is not part of the OpenAPI spec, so no generated command covers it.
func NewQueryCmd() *cobra.Command {
	var (
		body string
		path string
	)
	c := &cobra.Command{
		Use:   "query",
		Short: "Execute a raw GraphQL query against /explore/graphql",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			raw, err := clipkg.ResolveBody(body, nil, cmd.InOrStdin())
			if err != nil {
				return err
			}
			if len(raw) == 0 {
				return fmt.Errorf("--body is required (JSON, e.g. {\"query\":\"...\"})")
			}
			base, err := baseClient(deps)
			if err != nil {
				return err
			}
			resp, err := base.GraphQL(cmd.Context(), flexera.GraphQLRequest{
				Body: raw,
				Path: path,
			})
			if err != nil {
				return err
			}
			if resp.StatusCode < 200 || resp.StatusCode >= 300 {
				return flexera.ResponseError(resp.StatusCode, resp.Body)
			}
			var result any
			if err := json.Unmarshal(resp.Body, &result); err != nil {
				return fmt.Errorf("decoding response body: %w", err)
			}
			return deps.Printer.Render(deps.Stdout, deps.Config.Output, result)
		},
	}
	c.Flags().StringVar(&body, "body", "", "GraphQL request body: inline JSON, @file, or @- for stdin (required)")
	c.Flags().StringVar(&path, "path", "", "override the default /explore/graphql path")
	return c
}

// NewGenerateCmd builds the GraphQL query-generation command.
func NewGenerateCmd() *cobra.Command {
	var (
		prompt       string
		query        string
		modifyPrompt string
		indent       bool
	)
	c := &cobra.Command{
		Use:   "generate",
		Short: "Generate or modify a GraphQL query",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			if strings.TrimSpace(prompt) == "" && strings.TrimSpace(query) == "" {
				return fmt.Errorf("one of --prompt or --query is required")
			}
			base, err := baseClient(deps)
			if err != nil {
				return err
			}
			client, err := graphqllib.NewClient(
				base.Server,
				graphqllib.WithHTTPClient(base.Client),
				graphqllib.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
					for _, editor := range base.RequestEditors {
						if err := editor(ctx, req); err != nil {
							return err
						}
					}
					return nil
				}),
			)
			if err != nil {
				return err
			}
			result, err := client.GenerateQuery(cmd.Context(), int64(deps.Config.OrgID), graphqllib.GenerateQueryRequestBody{
				Prompt:       prompt,
				Query:        query,
				ModifyPrompt: modifyPrompt,
				Indent:       indent,
			})
			if err != nil {
				return err
			}
			return deps.Printer.Render(deps.Stdout, deps.Config.Output, result)
		},
	}
	c.Flags().StringVar(&prompt, "prompt", "", "natural-language description of the desired query")
	c.Flags().StringVar(&query, "query", "", "existing GraphQL query to modify")
	c.Flags().StringVar(&modifyPrompt, "modify-prompt", "", "instruction for modifying --query")
	c.Flags().BoolVar(&indent, "indent", false, "indent the generated query")
	return c
}

// Attach adds the non-OpenAPI query and generate commands below the
// generated graphql command.
func Attach(root *cobra.Command) error {
	for _, command := range root.Commands() {
		if command.Name() == "graphql" {
			command.AddCommand(NewQueryCmd())
			command.AddCommand(NewGenerateCmd())
			return nil
		}
	}
	return fmt.Errorf("graphql generated command is not registered")
}
