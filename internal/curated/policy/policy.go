// Package policy provides the project-scoped Policy resources (applied-policy,
// action-status, archived-incident, policy-template) and the meta workflows,
// attached beneath the generated `policy` service. They are hand-written
// because of GRS project auto-resolution (--project-id may be omitted),
// envelope pagination, and write-op confirmation.
package policy

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	flexera "github.com/flexera-public/unified-go-client"
)

// Attach adds the project-scoped resource groups and relationship-aware meta
// workflows beneath the generated `policy` service command, which already
// carries the org-scoped Policy resources.
func Attach(root *cobra.Command) error {
	return clipkg.AttachCommands(root, []string{"policy"}, newGroups()...)
}

// newGroups builds the curated policy resource groups.
func newGroups() []*cobra.Command {
	c := &cobra.Command{}
	c.AddCommand(
		group("applied-policy", "Applied policies (project-scoped)",
			newAppliedPolicyListCmd(), newAppliedPolicyGetCmd(), newAppliedPolicyCreateCmd(),
			newAppliedPolicyUpdateCmd(), newAppliedPolicyDeleteCmd(), newAppliedPolicyEvaluateCmd(),
			newAppliedPolicyLogCmd(), newAppliedPolicyStatusCmd()),
		group("action-status", "Policy action statuses (project-scoped)",
			newActionStatusListCmd(), newActionStatusGetCmd()),
		group("archived-incident", "Policy archived incidents (project-scoped)",
			newArchivedIncidentListCmd(), newArchivedIncidentGetCmd()),
		group("policy-template", "Policy templates (project-scoped)",
			newPolicyTemplateListCmd(), newPolicyTemplateGetCmd(), newPolicyTemplateCreateCmd(),
			newPolicyTemplateValidateCmd(), newPolicyTemplateUpdateCmd(), newPolicyTemplateDeleteCmd(),
			newPolicyTemplateEvaluateCmd()),
		group("meta", "Relationship-aware applied-policy meta operations (project-scoped)",
			newMetaDiscoverCmd(), newMetaAuditCmd(), newMetaTerminateChildrenCmd(),
			newMetaTerminateOrphanedCmd()),
	)
	annotateDirectOperations(c)
	groups := c.Commands()
	for _, g := range groups {
		c.RemoveCommand(g)
	}
	return groups
}

// annotateDirectOperations adds catalog identities only to direct API leaves,
// not to grouping commands or relationship-aware meta workflows.
func annotateDirectOperations(c *cobra.Command) {
	operations := map[string]map[string]string{
		"applied-policy": {
			"list":     "Policy_Applied_Policy_index",
			"get":      "Policy_Applied_Policy_show",
			"create":   "Policy_Applied_Policy_create",
			"update":   "Policy_Applied_Policy_update",
			"delete":   "Policy_Applied_Policy_delete",
			"evaluate": "Policy_Applied_Policy_evaluate",
			"log":      "Policy_Applied_Policy_showLog",
			"status":   "Policy_Applied_Policy_showStatus",
		},
		"action-status": {
			"list": "Policy_Action_Status_index",
			"get":  "Policy_Action_Status_show",
		},
		"archived-incident": {
			"list": "Policy_Archived_Incident_index",
			"get":  "Policy_Archived_Incident_show",
		},
		"policy-template": {
			"list":     "Policy_Policy_Template_index",
			"get":      "Policy_Policy_Template_show",
			"create":   "Policy_Policy_Template_create",
			"validate": "Policy_Policy_Template_validate",
			"update":   "Policy_Policy_Template_update",
			"delete":   "Policy_Policy_Template_delete",
			"evaluate": "Policy_Policy_Template_evaluate",
		},
	}
	for _, resource := range c.Commands() {
		for _, leaf := range resource.Commands() {
			operationID, ok := operations[resource.Name()][leaf.Name()]
			if !ok || leaf.HasSubCommands() {
				continue
			}
			if leaf.Annotations == nil {
				leaf.Annotations = make(map[string]string)
			}
			leaf.Annotations["flexera.operationId"] = operationID
		}
	}
}

func group(use, short string, children ...*cobra.Command) *cobra.Command {
	c := &cobra.Command{Use: use, Short: short, RunE: parentRunE}
	c.AddCommand(children...)
	return c
}

// parentRunE makes a grouping command show help when bare and error on an
// unknown subcommand (instead of cobra's default of printing help + exit 0).
func parentRunE(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return cmd.Help()
	}
	return fmt.Errorf("unknown command %q for %q", args[0], cmd.CommandPath())
}

// --- shared helpers ---

func tableUnsupported(deps *clipkg.Deps, op string) error {
	if deps.Config.Output == "table" {
		return fmt.Errorf("table output is not supported for %s", op)
	}
	return nil
}

func requireFlag(value, msg string) error {
	if strings.TrimSpace(value) == "" {
		return errors.New(msg)
	}
	return nil
}

// loadJSONInput reads a JSON request body from filePath ("-" for stdin) and
// strictly decodes it into target. Mirrors the legacy package-main helper.
func loadJSONInput(filePath string, stdin io.Reader, target any) error {
	trimmedPath := strings.TrimSpace(filePath)
	if trimmedPath == "" {
		return errors.New("input file is required")
	}
	var raw []byte
	var err error
	if trimmedPath == "-" {
		raw, err = io.ReadAll(stdin)
	} else {
		raw, err = os.ReadFile(trimmedPath)
	}
	if err != nil {
		return err
	}
	if len(bytes.TrimSpace(raw)) == 0 {
		return errors.New("input payload is empty")
	}
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		return err
	}
	if decoder.More() {
		return errors.New("input payload must contain exactly one JSON object")
	}
	return nil
}

// resolvePolicyProjectID returns --project-id when >0, otherwise resolves the
// org's single project via the library's flexera.ProjectResolver (erroring on
// zero or multiple). GRS is legacy but is the only source for project
// enumeration; the client is pointed at the GRS host.
func resolvePolicyProjectID(ctx context.Context, deps *clipkg.Deps, projectID int) (int, error) {
	if projectID > 0 {
		return projectID, nil
	}
	client, err := deps.Factory().NewAPIClientForBaseURL(deps.Config, deps.GRSBaseURL(""))
	if err != nil {
		return 0, fmt.Errorf("failed to resolve project ID via GRS: %w", err)
	}
	opts, err := deps.Factory().ProjectResolverOptions(deps.Config)
	if err != nil {
		return 0, fmt.Errorf("failed to resolve project ID via GRS: %w", err)
	}
	resolver, err := flexera.NewProjectResolver(client, opts...)
	if err != nil {
		return 0, fmt.Errorf("failed to resolve project ID via GRS: %w", err)
	}
	projects, err := resolver.ListProjects(ctx, int64(deps.Config.OrgID))
	if err != nil {
		return 0, fmt.Errorf("failed to resolve project ID via GRS: %w", err)
	}
	switch len(projects) {
	case 0:
		return 0, fmt.Errorf("no projects found for org %d via GRS; supply --project-id explicitly", deps.Config.OrgID)
	case 1:
		return int(projects[0].ID), nil
	default:
		ids := make([]string, 0, len(projects))
		for _, p := range projects {
			ids = append(ids, fmt.Sprintf("%d (%s)", p.ID, p.Name))
		}
		return 0, errors.New("org " + fmt.Sprint(deps.Config.OrgID) + " has multiple projects; pass --project-id to disambiguate: " + strings.Join(ids, ", "))
	}
}

func render(deps *clipkg.Deps, v any) error {
	return deps.Printer.Render(deps.Stdout, deps.Config.Output, v)
}

func collectPolicyPages[T any](ctx context.Context, noPaginate bool, initial *string, table bool, fetch flexera.PageFetcher) (any, error) {
	result, err := flexera.CollectPages(ctx, noPaginate, initial, fetch)
	if err != nil {
		return nil, err
	}
	if table {
		if merged, ok := result.(map[string]any); ok {
			data, err := json.Marshal(merged)
			if err != nil {
				return nil, err
			}
			var typed T
			if err := json.Unmarshal(data, &typed); err != nil {
				return nil, fmt.Errorf("decoding merged policy page: %w", err)
			}
			return &typed, nil
		}
	}
	return result, nil
}

// ===================== applied-policy (project-scoped) =====================

func newAppliedPolicyListCmd() *cobra.Command {
	var projectID, limit int
	var filter, orderBy, skipToken string
	var noPaginate bool
	c := &cobra.Command{Use: "list", Short: "List applied policies", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			var params *flexera.PolicyAppliedPolicyIndexParams
			if strings.TrimSpace(filter) != "" || strings.TrimSpace(orderBy) != "" || strings.TrimSpace(skipToken) != "" || limit > 0 {
				params = &flexera.PolicyAppliedPolicyIndexParams{}
				if t := strings.TrimSpace(filter); t != "" {
					params.Filter = &t
				}
				if t := strings.TrimSpace(orderBy); t != "" {
					params.OrderBy = &t
				}
				if t := strings.TrimSpace(skipToken); t != "" {
					params.SkipToken = &t
				}
				if limit > 0 {
					lv := int64(limit)
					params.Limit = &lv
				}
			}
			var initialSkipToken *string
			if token := strings.TrimSpace(skipToken); token != "" {
				initialSkipToken = &token
			}
			result, err := collectPolicyPages[flexera.PolicyAppliedPolicyList](cmd.Context(), noPaginate, initialSkipToken, deps.Config.Output == "table", func(ctx context.Context, token *string) (any, error) {
				pageParams := flexera.PolicyAppliedPolicyIndexParams{}
				if params != nil {
					pageParams = *params
				}
				pageParams.SkipToken = token
				response, callErr := client.PolicyAppliedPolicyIndexWithResponse(ctx, int64(deps.Config.OrgID), int64(pid), &pageParams)
				if callErr != nil {
					return nil, callErr
				}
				if response.JSON200 == nil {
					return nil, flexera.ResponseError(response.StatusCode(), response.Body)
				}
				return response.JSON200, nil
			})
			if err != nil {
				return err
			}
			return render(deps, result)
		}}
	addProjectFlag(c, &projectID)
	addListFlags(c, &filter, &orderBy, &skipToken, &limit)
	c.Flags().BoolVar(&noPaginate, "no-paginate", false, "return only the requested page (do not follow nextPage)")
	return c
}

func newAppliedPolicyGetCmd() *cobra.Command {
	var projectID int
	var id, view string
	c := &cobra.Command{Use: "get", Short: "Show an applied policy", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := requireFlag(id, "applied policy ID is required; use --id"); err != nil {
				return err
			}
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			var params *flexera.PolicyAppliedPolicyShowParams
			if t := strings.TrimSpace(view); t != "" {
				params = &flexera.PolicyAppliedPolicyShowParams{}
				v := flexera.PolicyAppliedPolicyShowParamsView(t)
				params.View = &v
			}
			resp, err := client.PolicyAppliedPolicyShowWithResponse(cmd.Context(), int64(deps.Config.OrgID), int64(pid), strings.TrimSpace(id), params)
			if err != nil {
				return err
			}
			if resp.JSON200 == nil {
				return flexera.ResponseError(resp.StatusCode(), resp.Body)
			}
			return render(deps, resp.JSON200)
		}}
	addProjectFlag(c, &projectID)
	c.Flags().StringVar(&id, "id", "", "Applied policy ID")
	c.Flags().StringVar(&view, "view", "", "Optional Policy applied-policy view")
	return c
}

func newAppliedPolicyCreateCmd() *cobra.Command {
	var projectID int
	var filePath string
	c := &cobra.Command{Use: "create", Short: "Create an applied policy", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			var body flexera.PolicyAppliedPolicyCreateJSONRequestBody
			if err := loadJSONInput(filePath, cmd.InOrStdin(), &body); err != nil {
				return fmt.Errorf("invalid applied-policy create payload: %w", err)
			}
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			if err := tableUnsupported(deps, "policy applied-policy create"); err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			resp, err := client.PolicyAppliedPolicyCreateWithResponse(cmd.Context(), int64(deps.Config.OrgID), int64(pid), body)
			if err != nil {
				return err
			}
			if err := flexera.ExpectStatus(resp.StatusCode(), resp.Body, 201); err != nil {
				return err
			}
			result := map[string]any{"status": "created", "name": body.Name, "templateRef": body.TemplateRef, "projectId": pid}
			if resp.JSON201 != nil {
				result["appliedPolicyId"] = resp.JSON201.Id
			}
			return render(deps, result)
		}}
	addProjectFlag(c, &projectID)
	c.Flags().StringVar(&filePath, "file", "", "Path to JSON payload file, or - to read from stdin")
	return c
}

func newAppliedPolicyUpdateCmd() *cobra.Command {
	var projectID int
	var id, filePath string
	c := &cobra.Command{Use: "update", Short: "Update an applied policy", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := requireFlag(id, "applied policy ID is required; use --id"); err != nil {
				return err
			}
			var body flexera.PolicyAppliedPolicyUpdateJSONRequestBody
			if err := loadJSONInput(filePath, cmd.InOrStdin(), &body); err != nil {
				return fmt.Errorf("invalid applied-policy update payload: %w", err)
			}
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			if err := tableUnsupported(deps, "policy applied-policy update"); err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			tid := strings.TrimSpace(id)
			resp, err := client.PolicyAppliedPolicyUpdateWithResponse(cmd.Context(), int64(deps.Config.OrgID), int64(pid), tid, body)
			if err != nil {
				return err
			}
			if err := flexera.ExpectStatus(resp.StatusCode(), resp.Body, 204); err != nil {
				return err
			}
			result := map[string]any{"status": "updated", "appliedPolicyId": tid, "projectId": pid}
			if body.Name != nil {
				result["name"] = *body.Name
			}
			return render(deps, result)
		}}
	addProjectFlag(c, &projectID)
	c.Flags().StringVar(&id, "id", "", "Applied policy ID")
	c.Flags().StringVar(&filePath, "file", "", "Path to JSON payload file, or - to read from stdin")
	return c
}

func newAppliedPolicyDeleteCmd() *cobra.Command {
	var projectID int
	var id string
	c := &cobra.Command{Use: "delete", Short: "Delete an applied policy", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := requireFlag(id, "applied policy ID is required; use --id"); err != nil {
				return err
			}
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			if err := tableUnsupported(deps, "policy applied-policy delete"); err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			tid := strings.TrimSpace(id)
			resp, err := client.PolicyAppliedPolicyDeleteWithResponse(cmd.Context(), int64(deps.Config.OrgID), int64(pid), tid)
			if err != nil {
				return err
			}
			if err := flexera.ExpectStatus(resp.StatusCode(), resp.Body, 204); err != nil {
				return err
			}
			return render(deps, map[string]any{"status": "deleted", "appliedPolicyId": tid, "projectId": pid})
		}}
	addProjectFlag(c, &projectID)
	c.Flags().StringVar(&id, "id", "", "Applied policy ID")
	return c
}

func newAppliedPolicyEvaluateCmd() *cobra.Command {
	var projectID int
	var id string
	c := &cobra.Command{Use: "evaluate", Short: "Request evaluation of an applied policy", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := requireFlag(id, "applied policy ID is required; use --id"); err != nil {
				return err
			}
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			if err := tableUnsupported(deps, "policy applied-policy evaluate"); err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			tid := strings.TrimSpace(id)
			resp, err := client.PolicyAppliedPolicyEvaluateWithResponse(cmd.Context(), int64(deps.Config.OrgID), int64(pid), tid)
			if err != nil {
				return err
			}
			if err := flexera.ExpectStatus(resp.StatusCode(), resp.Body, 202); err != nil {
				return err
			}
			return render(deps, map[string]any{"status": "evaluation_requested", "appliedPolicyId": tid, "projectId": pid})
		}}
	addProjectFlag(c, &projectID)
	c.Flags().StringVar(&id, "id", "", "Applied policy ID")
	return c
}

func newAppliedPolicyLogCmd() *cobra.Command {
	var projectID int
	var id string
	c := &cobra.Command{Use: "log", Short: "Show an applied policy's log", Args: cobra.NoArgs, Annotations: map[string]string{"flexera.output": "text"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := requireFlag(id, "applied policy ID is required; use --id"); err != nil {
				return err
			}
			deps := clipkg.DepsFrom(cmd.Context())
			// The API contract is text/markdown, even if a log looks like JSON.
			if err := tableUnsupported(deps, "policy applied-policy log"); err != nil {
				return err
			}
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			resp, err := client.PolicyAppliedPolicyShowLogWithResponse(cmd.Context(), int64(deps.Config.OrgID), int64(pid), strings.TrimSpace(id), nil)
			if err != nil {
				return err
			}
			if resp.StatusCode() != 200 {
				return flexera.ResponseError(resp.StatusCode(), resp.Body)
			}
			if _, err := deps.Stdout.Write(resp.Body); err != nil {
				return err
			}
			if len(resp.Body) == 0 || resp.Body[len(resp.Body)-1] != '\n' {
				_, err := fmt.Fprintln(deps.Stdout)
				return err
			}
			return nil
		}}
	addProjectFlag(c, &projectID)
	c.Flags().StringVar(&id, "id", "", "Applied policy ID")
	return c
}

func newAppliedPolicyStatusCmd() *cobra.Command {
	var projectID int
	var id string
	c := &cobra.Command{Use: "status", Short: "Show an applied policy's status", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := requireFlag(id, "applied policy ID is required; use --id"); err != nil {
				return err
			}
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			resp, err := client.PolicyAppliedPolicyShowStatusWithResponse(cmd.Context(), int64(deps.Config.OrgID), int64(pid), strings.TrimSpace(id))
			if err != nil {
				return err
			}
			if resp.JSON200 == nil {
				return flexera.ResponseError(resp.StatusCode(), resp.Body)
			}
			return render(deps, resp.JSON200)
		}}
	addProjectFlag(c, &projectID)
	c.Flags().StringVar(&id, "id", "", "Applied policy ID")
	return c
}

// ===================== action-status (project-scoped) =====================

func newActionStatusListCmd() *cobra.Command {
	var projectID, limit int
	var filter, orderBy, skipToken, view string
	var noPaginate bool
	c := &cobra.Command{Use: "list", Short: "List action statuses", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			var params *flexera.PolicyActionStatusIndexParams
			if strings.TrimSpace(filter) != "" || strings.TrimSpace(orderBy) != "" || strings.TrimSpace(skipToken) != "" || strings.TrimSpace(view) != "" || limit > 0 {
				params = &flexera.PolicyActionStatusIndexParams{}
				if t := strings.TrimSpace(filter); t != "" {
					params.Filter = &t
				}
				if t := strings.TrimSpace(orderBy); t != "" {
					params.OrderBy = &t
				}
				if t := strings.TrimSpace(skipToken); t != "" {
					params.SkipToken = &t
				}
				if t := strings.TrimSpace(view); t != "" {
					v := flexera.PolicyActionStatusIndexParamsView(t)
					params.View = &v
				}
				if limit > 0 {
					lv := int64(limit)
					params.Limit = &lv
				}
			}
			var initial *string
			if token := strings.TrimSpace(skipToken); token != "" {
				initial = &token
			}
			result, err := collectPolicyPages[flexera.PolicyActionStatusList](cmd.Context(), noPaginate, initial, deps.Config.Output == "table", func(ctx context.Context, token *string) (any, error) {
				page := flexera.PolicyActionStatusIndexParams{}
				if params != nil {
					page = *params
				}
				page.SkipToken = token
				resp, err := client.PolicyActionStatusIndexWithResponse(ctx, int64(deps.Config.OrgID), int64(pid), &page)
				if err != nil {
					return nil, err
				}
				if resp.JSON200 == nil {
					return nil, flexera.ResponseError(resp.StatusCode(), resp.Body)
				}
				return resp.JSON200, nil
			})
			if err != nil {
				return err
			}
			return render(deps, result)
		}}
	addProjectFlag(c, &projectID)
	addListFlags(c, &filter, &orderBy, &skipToken, &limit)
	c.Flags().StringVar(&view, "view", "", "Optional Policy action-status view")
	c.Flags().BoolVar(&noPaginate, "no-paginate", false, "return only the requested page (do not follow nextPage)")
	return c
}

func newActionStatusGetCmd() *cobra.Command {
	var projectID int
	var id, view string
	c := &cobra.Command{Use: "get", Short: "Show an action status", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := requireFlag(id, "action status ID is required; use --id"); err != nil {
				return err
			}
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			var params *flexera.PolicyActionStatusShowParams
			if t := strings.TrimSpace(view); t != "" {
				params = &flexera.PolicyActionStatusShowParams{}
				v := flexera.PolicyActionStatusShowParamsView(t)
				params.View = &v
			}
			resp, err := client.PolicyActionStatusShowWithResponse(cmd.Context(), int64(deps.Config.OrgID), int64(pid), strings.TrimSpace(id), params)
			if err != nil {
				return err
			}
			if resp.JSON200 == nil {
				return flexera.ResponseError(resp.StatusCode(), resp.Body)
			}
			return render(deps, resp.JSON200)
		}}
	addProjectFlag(c, &projectID)
	c.Flags().StringVar(&id, "id", "", "Action status ID")
	c.Flags().StringVar(&view, "view", "", "Optional Policy action-status view")
	return c
}

// ===================== archived-incident (project-scoped) =====================

func newArchivedIncidentListCmd() *cobra.Command {
	var projectID, limit int
	var filter, orderBy, skipToken, view string
	var noPaginate bool
	c := &cobra.Command{Use: "list", Short: "List archived incidents", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			var params *flexera.PolicyArchivedIncidentIndexParams
			if strings.TrimSpace(filter) != "" || strings.TrimSpace(orderBy) != "" || strings.TrimSpace(skipToken) != "" || strings.TrimSpace(view) != "" || limit > 0 {
				params = &flexera.PolicyArchivedIncidentIndexParams{}
				if t := strings.TrimSpace(filter); t != "" {
					params.Filter = &t
				}
				if t := strings.TrimSpace(orderBy); t != "" {
					params.OrderBy = &t
				}
				if t := strings.TrimSpace(skipToken); t != "" {
					params.SkipToken = &t
				}
				if t := strings.TrimSpace(view); t != "" {
					v := flexera.PolicyArchivedIncidentIndexParamsView(t)
					params.View = &v
				}
				if limit > 0 {
					lv := int64(limit)
					params.Limit = &lv
				}
			}
			var initial *string
			if token := strings.TrimSpace(skipToken); token != "" {
				initial = &token
			}
			result, err := collectPolicyPages[flexera.PolicyArchivedIncidentList](cmd.Context(), noPaginate, initial, deps.Config.Output == "table", func(ctx context.Context, token *string) (any, error) {
				page := flexera.PolicyArchivedIncidentIndexParams{}
				if params != nil {
					page = *params
				}
				page.SkipToken = token
				resp, err := client.PolicyArchivedIncidentIndexWithResponse(ctx, int64(deps.Config.OrgID), int64(pid), &page)
				if err != nil {
					return nil, err
				}
				if resp.JSON200 == nil {
					return nil, flexera.ResponseError(resp.StatusCode(), resp.Body)
				}
				return resp.JSON200, nil
			})
			if err != nil {
				return err
			}
			return render(deps, result)
		}}
	addProjectFlag(c, &projectID)
	addListFlags(c, &filter, &orderBy, &skipToken, &limit)
	c.Flags().StringVar(&view, "view", "", "Optional Policy archived-incident view")
	c.Flags().BoolVar(&noPaginate, "no-paginate", false, "return only the requested page (do not follow nextPage)")
	return c
}

func newArchivedIncidentGetCmd() *cobra.Command {
	var projectID int
	var id, view string
	c := &cobra.Command{Use: "get", Short: "Show an archived incident", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := requireFlag(id, "archived incident ID is required; use --id"); err != nil {
				return err
			}
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			var params *flexera.PolicyArchivedIncidentShowParams
			if t := strings.TrimSpace(view); t != "" {
				params = &flexera.PolicyArchivedIncidentShowParams{}
				v := flexera.PolicyArchivedIncidentShowParamsView(t)
				params.View = &v
			}
			resp, err := client.PolicyArchivedIncidentShowWithResponse(cmd.Context(), int64(deps.Config.OrgID), int64(pid), strings.TrimSpace(id), params)
			if err != nil {
				return err
			}
			if resp.JSON200 == nil {
				return flexera.ResponseError(resp.StatusCode(), resp.Body)
			}
			return render(deps, resp.JSON200)
		}}
	addProjectFlag(c, &projectID)
	c.Flags().StringVar(&id, "id", "", "Archived incident ID")
	c.Flags().StringVar(&view, "view", "", "Optional Policy archived-incident view")
	return c
}

// ===================== policy-template (project-scoped) =====================

func newPolicyTemplateListCmd() *cobra.Command {
	var projectID, limit int
	var filter, orderBy, skipToken, view string
	var noPaginate bool
	c := &cobra.Command{Use: "list", Short: "List policy templates", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			var params *flexera.PolicyPolicyTemplateIndexParams
			if strings.TrimSpace(filter) != "" || strings.TrimSpace(orderBy) != "" || strings.TrimSpace(skipToken) != "" || strings.TrimSpace(view) != "" || limit > 0 {
				params = &flexera.PolicyPolicyTemplateIndexParams{}
				if t := strings.TrimSpace(filter); t != "" {
					params.Filter = &t
				}
				if t := strings.TrimSpace(orderBy); t != "" {
					params.OrderBy = &t
				}
				if t := strings.TrimSpace(skipToken); t != "" {
					params.SkipToken = &t
				}
				if t := strings.TrimSpace(view); t != "" {
					v := flexera.PolicyPolicyTemplateIndexParamsView(t)
					params.View = &v
				}
				if limit > 0 {
					lv := int64(limit)
					params.Limit = &lv
				}
			}
			var initial *string
			if token := strings.TrimSpace(skipToken); token != "" {
				initial = &token
			}
			result, err := collectPolicyPages[flexera.PolicyPolicyTemplateList](cmd.Context(), noPaginate, initial, deps.Config.Output == "table", func(ctx context.Context, token *string) (any, error) {
				page := flexera.PolicyPolicyTemplateIndexParams{}
				if params != nil {
					page = *params
				}
				page.SkipToken = token
				resp, err := client.PolicyPolicyTemplateIndexWithResponse(ctx, int64(deps.Config.OrgID), int64(pid), &page)
				if err != nil {
					return nil, err
				}
				if resp.JSON200 == nil {
					return nil, flexera.ResponseError(resp.StatusCode(), resp.Body)
				}
				return resp.JSON200, nil
			})
			if err != nil {
				return err
			}
			return render(deps, result)
		}}
	addProjectFlag(c, &projectID)
	addListFlags(c, &filter, &orderBy, &skipToken, &limit)
	c.Flags().StringVar(&view, "view", "", "Optional Policy template view")
	c.Flags().BoolVar(&noPaginate, "no-paginate", false, "return only the requested page (do not follow nextPage)")
	return c
}

func newPolicyTemplateGetCmd() *cobra.Command {
	var projectID int
	var id, view string
	c := &cobra.Command{Use: "get", Short: "Show a policy template", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := requireFlag(id, "policy template ID is required; use --id"); err != nil {
				return err
			}
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			var params *flexera.PolicyPolicyTemplateShowParams
			if t := strings.TrimSpace(view); t != "" {
				params = &flexera.PolicyPolicyTemplateShowParams{}
				v := flexera.PolicyPolicyTemplateShowParamsView(t)
				params.View = &v
			}
			resp, err := client.PolicyPolicyTemplateShowWithResponse(cmd.Context(), int64(deps.Config.OrgID), int64(pid), strings.TrimSpace(id), params)
			if err != nil {
				return err
			}
			if resp.JSON200 == nil {
				return flexera.ResponseError(resp.StatusCode(), resp.Body)
			}
			return render(deps, resp.JSON200)
		}}
	addProjectFlag(c, &projectID)
	c.Flags().StringVar(&id, "id", "", "Policy template ID")
	c.Flags().StringVar(&view, "view", "", "Optional Policy template view")
	return c
}

func newPolicyTemplateCreateCmd() *cobra.Command {
	var projectID int
	var filePath string
	c := &cobra.Command{Use: "create", Short: "Create a policy template", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			var body flexera.PolicyPolicyTemplateCreateJSONRequestBody
			if err := loadJSONInput(filePath, cmd.InOrStdin(), &body); err != nil {
				return fmt.Errorf("invalid policy-template create payload: %w", err)
			}
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			if err := tableUnsupported(deps, "policy policy-template create"); err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			resp, err := client.PolicyPolicyTemplateCreateWithResponse(cmd.Context(), int64(deps.Config.OrgID), int64(pid), body)
			if err != nil {
				return err
			}
			if err := flexera.ExpectStatus(resp.StatusCode(), resp.Body, 201); err != nil {
				return err
			}
			result := map[string]any{"status": "created", "projectId": pid, "filename": body.Filename}
			if resp.JSON201 != nil {
				result["location"] = resp.JSON201.Location
			} else if resp.HTTPResponse != nil {
				if location := strings.TrimSpace(resp.HTTPResponse.Header.Get("Location")); location != "" {
					result["location"] = location
				}
			}
			return render(deps, result)
		}}
	addProjectFlag(c, &projectID)
	c.Flags().StringVar(&filePath, "file", "", "Path to JSON payload file, or - to read from stdin")
	return c
}

func newPolicyTemplateValidateCmd() *cobra.Command {
	var projectID int
	var filePath string
	c := &cobra.Command{Use: "validate", Short: "Validate a policy template", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			var body flexera.PolicyPolicyTemplateValidateJSONRequestBody
			if err := loadJSONInput(filePath, cmd.InOrStdin(), &body); err != nil {
				return fmt.Errorf("invalid policy-template validate payload: %w", err)
			}
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			if err := tableUnsupported(deps, "policy policy-template validate"); err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			resp, err := client.PolicyPolicyTemplateValidateWithResponse(cmd.Context(), int64(deps.Config.OrgID), int64(pid), body)
			if err != nil {
				return err
			}
			if err := flexera.ExpectStatus(resp.StatusCode(), resp.Body, 204); err != nil {
				return err
			}
			return render(deps, map[string]any{"status": "validated", "projectId": pid, "filename": body.Filename})
		}}
	addProjectFlag(c, &projectID)
	c.Flags().StringVar(&filePath, "file", "", "Path to JSON payload file, or - to read from stdin")
	return c
}

func newPolicyTemplateUpdateCmd() *cobra.Command {
	var projectID int
	var id, filePath string
	c := &cobra.Command{Use: "update", Short: "Update a policy template", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := requireFlag(id, "policy template ID is required; use --id"); err != nil {
				return err
			}
			var body flexera.PolicyPolicyTemplateUpdateJSONRequestBody
			if err := loadJSONInput(filePath, cmd.InOrStdin(), &body); err != nil {
				return fmt.Errorf("invalid policy-template update payload: %w", err)
			}
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			if err := tableUnsupported(deps, "policy policy-template update"); err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			tid := strings.TrimSpace(id)
			resp, err := client.PolicyPolicyTemplateUpdateWithResponse(cmd.Context(), int64(deps.Config.OrgID), int64(pid), tid, body)
			if err != nil {
				return err
			}
			if err := flexera.ExpectStatus(resp.StatusCode(), resp.Body, 204); err != nil {
				return err
			}
			return render(deps, map[string]any{"status": "updated", "projectId": pid, "policyTemplateId": tid, "filename": body.Filename})
		}}
	addProjectFlag(c, &projectID)
	c.Flags().StringVar(&id, "id", "", "Policy template ID")
	c.Flags().StringVar(&filePath, "file", "", "Path to JSON payload file, or - to read from stdin")
	return c
}

func newPolicyTemplateDeleteCmd() *cobra.Command {
	var projectID int
	var id string
	c := &cobra.Command{Use: "delete", Short: "Delete a policy template", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := requireFlag(id, "policy template ID is required; use --id"); err != nil {
				return err
			}
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			if err := tableUnsupported(deps, "policy policy-template delete"); err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			tid := strings.TrimSpace(id)
			resp, err := client.PolicyPolicyTemplateDeleteWithResponse(cmd.Context(), int64(deps.Config.OrgID), int64(pid), tid)
			if err != nil {
				return err
			}
			if err := flexera.ExpectStatus(resp.StatusCode(), resp.Body, 204); err != nil {
				return err
			}
			return render(deps, map[string]any{"status": "deleted", "projectId": pid, "policyTemplateId": tid})
		}}
	addProjectFlag(c, &projectID)
	c.Flags().StringVar(&id, "id", "", "Policy template ID")
	return c
}

func newPolicyTemplateEvaluateCmd() *cobra.Command {
	var projectID int
	var id, filePath string
	c := &cobra.Command{Use: "evaluate", Short: "Evaluate a policy template", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := requireFlag(id, "policy template ID is required; use --id"); err != nil {
				return err
			}
			var body flexera.PolicyPolicyTemplateEvaluateJSONRequestBody
			if err := loadJSONInput(filePath, cmd.InOrStdin(), &body); err != nil {
				return fmt.Errorf("invalid policy-template evaluate payload: %w", err)
			}
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			if err := tableUnsupported(deps, "policy policy-template evaluate"); err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			resp, err := client.PolicyPolicyTemplateEvaluateWithResponse(cmd.Context(), int64(deps.Config.OrgID), int64(pid), strings.TrimSpace(id), body)
			if err != nil {
				return err
			}
			if resp.JSON200 == nil {
				return flexera.ResponseError(resp.StatusCode(), resp.Body)
			}
			return render(deps, resp.JSON200)
		}}
	addProjectFlag(c, &projectID)
	c.Flags().StringVar(&id, "id", "", "Policy template ID")
	c.Flags().StringVar(&filePath, "file", "", "Path to JSON payload file, or - to read from stdin")
	return c
}

// ===================== meta (relationship-aware, project-scoped) =====================

func newMetaDiscoverCmd() *cobra.Command {
	var projectID int
	c := &cobra.Command{Use: "discover", Short: "Discover applied-policy parent/child relationships", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			wf := flexera.NewPolicyMetaWorkflow(client)
			snapshot, err := wf.Discover(cmd.Context(), int64(deps.Config.OrgID), int64(pid))
			if err != nil {
				return err
			}
			return render(deps, snapshot)
		}}
	addProjectFlag(c, &projectID)
	return c
}

func newMetaAuditCmd() *cobra.Command {
	var projectID int
	c := &cobra.Command{Use: "audit", Short: "Audit status of every applied policy in a project", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			wf := flexera.NewPolicyMetaWorkflow(client)
			out, err := wf.Audit(cmd.Context(), int64(deps.Config.OrgID), int64(pid))
			if err != nil {
				return err
			}
			return render(deps, out)
		}}
	addProjectFlag(c, &projectID)
	return c
}

func newMetaTerminateChildrenCmd() *cobra.Command {
	var projectID int
	var parentID string
	var yes, dryRun bool
	c := &cobra.Command{Use: "terminate-children", Short: "Delete all unambiguous children of a parent applied policy", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := requireFlag(parentID, "parent applied policy ID is required; use --parent-id"); err != nil {
				return err
			}
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			if !dryRun {
				if err := tableUnsupported(deps, "policy meta terminate-children"); err != nil {
					return err
				}
				if !yes {
					return errors.New("destructive operation requires --yes (or use --dry-run to preview)")
				}
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			wf := flexera.NewPolicyMetaWorkflow(client)
			out, err := wf.TerminateChildren(cmd.Context(), flexera.PolicyMetaTerminationInput{
				OrgID:     int64(deps.Config.OrgID),
				ProjectID: int64(pid),
				ParentID:  strings.TrimSpace(parentID),
				DryRun:    dryRun,
			})
			if err != nil {
				return err
			}
			return render(deps, out)
		}}
	addProjectFlag(c, &projectID)
	c.Flags().StringVar(&parentID, "parent-id", "", "Parent applied policy ID")
	c.Flags().BoolVar(&yes, "yes", false, "Confirm the operation; required unless --dry-run")
	c.Flags().BoolVar(&dryRun, "dry-run", false, "Preview matched children without deleting them")
	return c
}

func newMetaTerminateOrphanedCmd() *cobra.Command {
	var projectID int
	var yes, dryRun bool
	c := &cobra.Command{Use: "terminate-orphaned", Short: "Delete applied policies whose parent reference is orphaned", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			pid, err := resolvePolicyProjectID(cmd.Context(), deps, projectID)
			if err != nil {
				return err
			}
			if !dryRun {
				if err := tableUnsupported(deps, "policy meta terminate-orphaned"); err != nil {
					return err
				}
				if !yes {
					return errors.New("destructive operation requires --yes (or use --dry-run to preview)")
				}
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			wf := flexera.NewPolicyMetaWorkflow(client)
			out, err := wf.TerminateOrphaned(cmd.Context(), flexera.PolicyMetaTerminationInput{
				OrgID:     int64(deps.Config.OrgID),
				ProjectID: int64(pid),
				DryRun:    dryRun,
			})
			if err != nil {
				return err
			}
			return render(deps, out)
		}}
	addProjectFlag(c, &projectID)
	c.Flags().BoolVar(&yes, "yes", false, "Confirm the operation; required unless --dry-run")
	c.Flags().BoolVar(&dryRun, "dry-run", false, "Preview matched orphaned policies without deleting them")
	return c
}

// --- flag helpers ---

func addListFlags(c *cobra.Command, filter, orderBy, skipToken *string, limit *int) {
	c.Flags().StringVar(filter, "filter", "", "Optional filter expression")
	c.Flags().StringVar(orderBy, "order-by", "", "Optional sort expression")
	c.Flags().StringVar(skipToken, "skip-token", "", "Optional pagination token; resume from this position")
	c.Flags().IntVar(limit, "limit", 0, "Optional page size")
}

func addProjectFlag(c *cobra.Command, projectID *int) {
	c.Flags().IntVar(projectID, "project-id", 0, "Project ID (optional; resolved from GRS for the org when omitted)")
}
