// Package finops provides the hand-written "finops" command tree: Optima
// cost analytics + recommendations. All commands target the Optima zone host
// (api.optima*.flexeraeng.com) rather than the unified API gateway; the
// Optima client factory (internal/flexera) resolves the base URL and shares
// the unified OAuth bearer token.
package finops

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	cliflexera "github.com/flexera-public/flexera-cli/internal/flexera"
	flexera "github.com/flexera-public/unified-go-client"
	ba "github.com/flexera-public/unified-go-client/rightscale/bill_analysis"
	optrec "github.com/flexera-public/unified-go-client/rightscale/optima_recommendations"
)

const flagOptimaBaseURL = "optima-base-url"

// NewCmd builds the "finops" command tree.
func NewCmd() *cobra.Command {
	c := &cobra.Command{
		Use:     "finops",
		Short:   "Cost analytics, billing centers, and recommendations (Optima)",
		Example: "flexera-cli finops cost get --org-id 123 --start 2024-01 --end 2024-04",
	}
	// --optima-base-url is shared by every subcommand.
	c.PersistentFlags().String(flagOptimaBaseURL, "", "Override the Optima base URL (env: FLEXERA_CLI_OPTIMA_BASE_URL)")

	billingCenter := &cobra.Command{Use: "billing-center", Short: "Billing centers", RunE: parentRunE}
	billingCenter.AddCommand(newBillingCenterListCmd(), newBillingCenterGetCmd(), newBillingCenterAllocationTableCmd())

	cost := &cobra.Command{Use: "cost", Short: "Cost queries and exports", RunE: parentRunE}
	cost.AddCommand(
		newCostGetCmd(), newCostAggregatedCmd(), newCostSelectCmd(),
		newCostDimensionsCmd(), newCostMetricsCmd(),
		newCostExportSelectCmd(), newCostExportStatusCmd(),
	)

	recommendation := &cobra.Command{Use: "recommendation", Short: "Optima recommendations", RunE: parentRunE}
	recommendation.AddCommand(
		newRecommendationListCmd("list-usage-reduction", "usage_reduction", "List usage-reduction recommendations"),
		newRecommendationListCmd("list-rate-reduction", "rate_reduction", "List rate-reduction recommendations"),
	)

	billMonth := &cobra.Command{Use: "bill-month", Short: "Bill months", RunE: parentRunE}
	billMonth.AddCommand(newBillMonthListCmd())

	adjustment := &cobra.Command{Use: "adjustment", Short: "Adjustment definition", RunE: parentRunE}
	adjustment.AddCommand(newAdjustmentShowCmd(), newAdjustmentUpdateCmd())

	c.AddCommand(
		billingCenter, cost, recommendation, billMonth, adjustment,
		newReportCmd("anomaly-report", "Run an anomaly report", func(ctx context.Context, clients *cliflexera.OptimaClients, orgID int, body []byte) (*http.Response, []byte, error) {
			resp, err := clients.BillAnalysis.AnomaliesReportWithBodyWithResponse(ctx, int(orgID), nil, "application/json", strings.NewReader(string(body)))
			if err != nil {
				return nil, nil, err
			}
			return resp.HTTPResponse, resp.Body, nil
		}),
		newReportCmd("forecast-report", "Run a forecast report", func(ctx context.Context, clients *cliflexera.OptimaClients, orgID int, body []byte) (*http.Response, []byte, error) {
			resp, err := clients.BillAnalysis.ForecastsReportWithBodyWithResponse(ctx, int(orgID), nil, "application/json", strings.NewReader(string(body)))
			if err != nil {
				return nil, nil, err
			}
			return resp.HTTPResponse, resp.Body, nil
		}),
	)
	return c
}

func parentRunE(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return cmd.Help()
	}
	return fmt.Errorf("unknown command %q for %q", args[0], cmd.CommandPath())
}

// optimaClients builds the Optima multi-client bundle, requiring an org ID.
func optimaClients(cmd *cobra.Command, deps *clipkg.Deps) (*cliflexera.OptimaClients, error) {
	if err := deps.Config.RequireOrgID(); err != nil {
		return nil, err
	}
	baseURL, _ := cmd.Flags().GetString(flagOptimaBaseURL)
	return deps.Factory().NewOptimaClients(cmd.Context(), deps.Config, deps.Getenv, baseURL)
}

// --- billing-center ---

func newBillingCenterListCmd() *cobra.Command {
	return &cobra.Command{
		Use: "list", Short: "List billing centers", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			clients, err := optimaClients(cmd, deps)
			if err != nil {
				return err
			}
			resp, err := clients.BillingCenterService.BillingCentersIndexWithResponse(cmd.Context(), deps.Config.OrgID, nil)
			if err != nil {
				return err
			}
			return emitOptima(deps, resp.HTTPResponse, resp.Body)
		},
	}
}

func newBillingCenterGetCmd() *cobra.Command {
	var bcID string
	c := &cobra.Command{
		Use: "get", Short: "Show a billing center", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			if strings.TrimSpace(bcID) == "" {
				return errors.New("--id is required")
			}
			clients, err := optimaClients(cmd, deps)
			if err != nil {
				return err
			}
			resp, err := clients.BillingCenterService.BillingCentersShowWithResponse(cmd.Context(), deps.Config.OrgID, bcID, nil)
			if err != nil {
				return err
			}
			return emitOptima(deps, resp.HTTPResponse, resp.Body)
		},
	}
	c.Flags().StringVar(&bcID, "id", "", "Billing center ID (required)")
	return c
}

func newBillingCenterAllocationTableCmd() *cobra.Command {
	var bcID string
	c := &cobra.Command{
		Use: "allocation-table", Short: "Show a billing center's allocation table", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			if strings.TrimSpace(bcID) == "" {
				return errors.New("--id is required")
			}
			clients, err := optimaClients(cmd, deps)
			if err != nil {
				return err
			}
			resp, err := clients.BillingCenterService.BillingCentersShowAllocationTableWithResponse(cmd.Context(), deps.Config.OrgID, bcID, nil)
			if err != nil {
				return err
			}
			return emitOptima(deps, resp.HTTPResponse, resp.Body)
		},
	}
	c.Flags().StringVar(&bcID, "id", "", "Billing center ID (required)")
	return c
}

// --- cost ---

func newCostGetCmd() *cobra.Command {
	var startAt, endAt, granularity, endpoint, dimensions, metrics, billingCenterIDs string
	var noChunk, summarized bool
	c := &cobra.Command{
		Use: "get", Short: "Curated cost query (auto-routes aggregated/select, auto-chunks windows)", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			if strings.TrimSpace(startAt) == "" || strings.TrimSpace(endAt) == "" {
				return errors.New("--start and --end are required")
			}
			clients, err := optimaClients(cmd, deps)
			if err != nil {
				return err
			}
			resolver := flexera.NewBillingCenterResolver(clients.BillingCenterService)
			helper := flexera.New(clients.BillAnalysis, resolver)
			req := flexera.Request{
				OrgID:            deps.Config.OrgID,
				BillingCenterIDs: splitCSV(billingCenterIDs),
				StartAt:          startAt,
				EndAt:            endAt,
				Granularity:      flexera.Granularity(granularity),
				Dimensions:       splitCSV(dimensions),
				Metrics:          splitCSV(metrics),
				Endpoint:         flexera.Endpoint(endpoint),
				NoChunk:          noChunk,
			}
			if summarized {
				req.Summarized = &summarized
			}
			resp, err := helper.GetCost(cmd.Context(), req)
			if err != nil {
				return err
			}
			return deps.Printer.Render(deps.Stdout, deps.Config.Output, resp)
		},
	}
	c.Flags().StringVar(&startAt, "start", "", "Inclusive start (YYYY-MM or YYYY-MM-DD)")
	c.Flags().StringVar(&endAt, "end", "", "Exclusive end (YYYY-MM or YYYY-MM-DD)")
	c.Flags().StringVar(&granularity, "granularity", "month", "Period granularity: month|day")
	c.Flags().StringVar(&endpoint, "endpoint", "auto", "auto (default), aggregated, or select")
	c.Flags().StringVar(&dimensions, "dimensions", "", "Comma-separated cost dimensions (auto-routes to /select when resource_id is present)")
	c.Flags().StringVar(&metrics, "metrics", "", "Comma-separated metrics (default: cost_amortized_unblended_adj)")
	c.Flags().StringVar(&billingCenterIDs, "billing-center-ids", "", "Comma-separated BC IDs (default: all top-level BCs)")
	c.Flags().BoolVar(&noChunk, "no-chunk", false, "Disable automatic >24-month period chunking")
	c.Flags().BoolVar(&summarized, "summarized", false, "Collapse all buckets into one row (aggregated only)")
	return c
}

func newCostAggregatedCmd() *cobra.Command {
	return newPOSTBodyCmd("aggregated", "Raw POST /costs/aggregated", func(ctx context.Context, clients *cliflexera.OptimaClients, orgID int, body []byte) (*http.Response, []byte, error) {
		resp, err := clients.BillAnalysis.CostsAggregatedWithBodyWithResponse(ctx, int(orgID), nil, "application/json", strings.NewReader(string(body)))
		if err != nil {
			return nil, nil, err
		}
		return resp.HTTPResponse, resp.Body, nil
	})
}

func newCostSelectCmd() *cobra.Command {
	return newPOSTBodyCmd("select", "Raw POST /costs/select", func(ctx context.Context, clients *cliflexera.OptimaClients, orgID int, body []byte) (*http.Response, []byte, error) {
		resp, err := clients.BillAnalysis.CostsSelectWithBodyWithResponse(ctx, int(orgID), nil, "application/json", strings.NewReader(string(body)))
		if err != nil {
			return nil, nil, err
		}
		return resp.HTTPResponse, resp.Body, nil
	})
}

func newCostExportSelectCmd() *cobra.Command {
	return newPOSTBodyCmd("export-select", "Raw POST /costs/select/export", func(ctx context.Context, clients *cliflexera.OptimaClients, orgID int, body []byte) (*http.Response, []byte, error) {
		resp, err := clients.BillAnalysis.CostsExportSelectWithBodyWithResponse(ctx, int(orgID), nil, "application/json", strings.NewReader(string(body)))
		if err != nil {
			return nil, nil, err
		}
		return resp.HTTPResponse, resp.Body, nil
	})
}

func newCostDimensionsCmd() *cobra.Command {
	var dataset string
	c := &cobra.Command{
		Use: "dimensions", Short: "List available cost dimensions", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			clients, err := optimaClients(cmd, deps)
			if err != nil {
				return err
			}
			params := &ba.CostsDimensionsParams{}
			if d := strings.TrimSpace(dataset); d != "" {
				v := ba.CostsDimensionsParamsDataset(d)
				params.Dataset = &v
			}
			resp, err := clients.BillAnalysis.CostsDimensionsWithResponse(cmd.Context(), deps.Config.OrgID, params)
			if err != nil {
				return err
			}
			return emitOptima(deps, resp.HTTPResponse, resp.Body)
		},
	}
	c.Flags().StringVar(&dataset, "dataset", "", "Optional dataset: billing or cost")
	return c
}

func newCostMetricsCmd() *cobra.Command {
	var dataset string
	c := &cobra.Command{
		Use: "metrics", Short: "List available cost metrics", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			clients, err := optimaClients(cmd, deps)
			if err != nil {
				return err
			}
			params := &ba.CostsMetricsParams{}
			if d := strings.TrimSpace(dataset); d != "" {
				v := ba.CostsMetricsParamsDataset(d)
				params.Dataset = &v
			}
			resp, err := clients.BillAnalysis.CostsMetricsWithResponse(cmd.Context(), deps.Config.OrgID, params)
			if err != nil {
				return err
			}
			return emitOptima(deps, resp.HTTPResponse, resp.Body)
		},
	}
	c.Flags().StringVar(&dataset, "dataset", "", "Optional dataset: billing or cost")
	return c
}

func newCostExportStatusCmd() *cobra.Command {
	var exportID string
	c := &cobra.Command{
		Use: "export-status", Short: "Check an export's status", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			if strings.TrimSpace(exportID) == "" {
				return errors.New("--export-id is required")
			}
			clients, err := optimaClients(cmd, deps)
			if err != nil {
				return err
			}
			resp, err := clients.BillAnalysis.CostsExportSelectStatusWithResponse(cmd.Context(), deps.Config.OrgID, exportID, nil)
			if err != nil {
				return err
			}
			return emitOptima(deps, resp.HTTPResponse, resp.Body)
		},
	}
	c.Flags().StringVar(&exportID, "export-id", "", "Export ID returned by export-select (required)")
	return c
}

// --- recommendation ---

func newRecommendationListCmd(use, kind, short string) *cobra.Command {
	var billingCenterIDs string
	c := &cobra.Command{
		Use: use, Short: short, Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			clients, err := optimaClients(cmd, deps)
			if err != nil {
				return err
			}
			bcIDs := splitCSV(billingCenterIDs)
			if len(bcIDs) == 0 {
				resolver := flexera.NewBillingCenterResolver(clients.BillingCenterService)
				resolved, rerr := resolver.TopLevelIDs(cmd.Context(), deps.Config.OrgID)
				if rerr != nil {
					return fmt.Errorf("auto-resolve billing centers: %w", rerr)
				}
				bcIDs = resolved
			}
			params := &optrec.RecommendationsIndexParams{BillingCenterIDs: &bcIDs}
			resp, err := clients.OptimaRecommendations.RecommendationsIndexWithResponse(cmd.Context(), deps.Config.OrgID, params)
			if err != nil {
				return err
			}
			if resp.HTTPResponse == nil || resp.HTTPResponse.StatusCode != http.StatusOK {
				return optimaResponseError(resp.HTTPResponse, resp.Body)
			}
			filtered, ferr := flexera.FilterRecommendationsByCategory(resp.Body, kind)
			if ferr != nil {
				return writeRaw(deps.Stdout, resp.Body)
			}
			return writeRaw(deps.Stdout, filtered)
		},
	}
	c.Flags().StringVar(&billingCenterIDs, "billing-center-ids", "", "Comma-separated BC IDs (default: all top-level BCs)")
	return c
}

// --- bill-month ---

func newBillMonthListCmd() *cobra.Command {
	var orderBy string
	var limit, offset int64
	c := &cobra.Command{
		Use: "list", Short: "List bill months", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			clients, err := optimaClients(cmd, deps)
			if err != nil {
				return err
			}
			params := &ba.BillMonthsSearchParams{}
			if limit > 0 {
				params.Limit = &limit
			}
			if offset > 0 {
				params.Offset = &offset
			}
			if v := strings.TrimSpace(orderBy); v != "" {
				params.OrderBy = &v
			}
			// Raw client so empty-date fields don't trip the strict decoder.
			httpResp, err := clients.BillAnalysis.BillMonthsSearch(cmd.Context(), deps.Config.OrgID, params)
			if err != nil {
				return err
			}
			body, _ := io.ReadAll(httpResp.Body)
			_ = httpResp.Body.Close()
			if !statusOK(httpResp.StatusCode) {
				return optimaResponseError(httpResp, body)
			}
			return writeRaw(deps.Stdout, body)
		},
	}
	c.Flags().Int64Var(&limit, "limit", 0, "Maximum number of records (0 = API default)")
	c.Flags().Int64Var(&offset, "offset", 0, "Starting offset for pagination")
	c.Flags().StringVar(&orderBy, "order-by", "", "Order spec, e.g. month=desc")
	return c
}

// --- adjustment ---

func newAdjustmentShowCmd() *cobra.Command {
	return &cobra.Command{
		Use: "show", Short: "Show the adjustment definition", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			clients, err := optimaClients(cmd, deps)
			if err != nil {
				return err
			}
			httpResp, err := clients.BillAnalysis.AdjustmentDefinitionShow(cmd.Context(), deps.Config.OrgID, nil)
			if err != nil {
				return err
			}
			body, _ := io.ReadAll(httpResp.Body)
			_ = httpResp.Body.Close()
			if !statusOK(httpResp.StatusCode) {
				return optimaResponseError(httpResp, body)
			}
			return writeRaw(deps.Stdout, body)
		},
	}
}

func newAdjustmentUpdateCmd() *cobra.Command {
	var file string
	var yes, dryRun bool
	c := &cobra.Command{
		Use: "update", Short: "Update the adjustment definition", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			body, err := readBodyFile(file, cmd.InOrStdin())
			if err != nil {
				return err
			}
			plan := map[string]any{
				"op":     "finops adjustment update",
				"orgId":  deps.Config.OrgID,
				"body":   json.RawMessage(body),
				"method": "PUT /orgs/{org}/adjustments/definition",
			}
			done, err := resolveWriteOp(dryRun, yes, true, deps.Stdout, plan)
			if err != nil {
				return err
			}
			if done {
				return nil
			}
			clients, err := optimaClients(cmd, deps)
			if err != nil {
				return err
			}
			resp, err := clients.BillAnalysis.AdjustmentDefinitionUpdateWithBodyWithResponse(cmd.Context(), deps.Config.OrgID, nil, "application/json", strings.NewReader(string(body)))
			if err != nil {
				return err
			}
			if resp.HTTPResponse == nil || !statusOK(resp.HTTPResponse.StatusCode) {
				return optimaResponseError(resp.HTTPResponse, resp.Body)
			}
			return writeRaw(deps.Stdout, resp.Body)
		},
	}
	c.Flags().StringVar(&file, "file", "", "Path to JSON request body, or - for stdin (required)")
	c.Flags().BoolVar(&yes, "yes", false, "Confirm the operation; required for destructive ops")
	c.Flags().BoolVar(&dryRun, "dry-run", false, "Print the planned operation as JSON and exit without contacting the API")
	return c
}

// --- report (anomaly/forecast) ---

func newReportCmd(use, short string, call func(ctx context.Context, clients *cliflexera.OptimaClients, orgID int, body []byte) (*http.Response, []byte, error)) *cobra.Command {
	return newPOSTBodyCmdWithUse(use, short, call)
}

// --- shared POST-body command ---

func newPOSTBodyCmd(use, short string, call func(ctx context.Context, clients *cliflexera.OptimaClients, orgID int, body []byte) (*http.Response, []byte, error)) *cobra.Command {
	return newPOSTBodyCmdWithUse(use, short, call)
}

func newPOSTBodyCmdWithUse(use, short string, call func(ctx context.Context, clients *cliflexera.OptimaClients, orgID int, body []byte) (*http.Response, []byte, error)) *cobra.Command {
	var file string
	c := &cobra.Command{
		Use: use, Short: short, Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			body, err := readBodyFile(file, cmd.InOrStdin())
			if err != nil {
				return err
			}
			clients, err := optimaClients(cmd, deps)
			if err != nil {
				return err
			}
			httpResp, respBody, err := call(cmd.Context(), clients, deps.Config.OrgID, body)
			if err != nil {
				return err
			}
			if httpResp == nil || !statusOK(httpResp.StatusCode) {
				return optimaResponseError(httpResp, respBody)
			}
			return writeRaw(deps.Stdout, respBody)
		},
	}
	c.Flags().StringVar(&file, "file", "", "Path to JSON request body, or - for stdin (required)")
	return c
}

// --- helpers ---

func splitCSV(in string) []string {
	if strings.TrimSpace(in) == "" {
		return nil
	}
	parts := strings.Split(in, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if v := strings.TrimSpace(p); v != "" {
			out = append(out, v)
		}
	}
	return out
}

func statusOK(code int) bool { return code >= 200 && code < 300 }

func optimaResponseError(httpResp *http.Response, body []byte) error {
	status := 0
	if httpResp != nil {
		status = httpResp.StatusCode
	}
	return fmt.Errorf("Optima API returned status %d: %s", status, string(body))
}

// emitOptima validates a 200 response and writes the raw body.
func emitOptima(deps *clipkg.Deps, httpResp *http.Response, body []byte) error {
	if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		return optimaResponseError(httpResp, body)
	}
	return writeRaw(deps.Stdout, body)
}

// writeRaw emits a raw JSON byte slice with a trailing newline.
func writeRaw(stdout io.Writer, body []byte) error {
	if len(body) == 0 {
		return nil
	}
	if _, err := stdout.Write(body); err != nil {
		return err
	}
	if body[len(body)-1] != '\n' {
		_, _ = stdout.Write([]byte("\n"))
	}
	return nil
}

func readBodyFile(path string, stdin io.Reader) ([]byte, error) {
	if strings.TrimSpace(path) == "" {
		return nil, errors.New("--file is required (use - for stdin)")
	}
	if path == "-" {
		if stdin == nil {
			stdin = os.Stdin
		}
		return io.ReadAll(stdin)
	}
	return os.ReadFile(path)
}

// resolveWriteOp mirrors the package-main write-op confirmation flow: on
// --dry-run it prints a plan summary and returns done=true; destructive ops
// require --yes. Returns done=true when the caller should stop (dry-run).
func resolveWriteOp(dryRun, yes, destructive bool, stdout io.Writer, plan map[string]any) (bool, error) {
	if dryRun {
		summary := map[string]any{"dryRun": true, "destructive": destructive, "plan": plan}
		enc := json.NewEncoder(stdout)
		enc.SetIndent("", "  ")
		_ = enc.Encode(summary)
		return true, nil
	}
	if destructive && !yes {
		return false, errors.New("destructive operation requires --yes (or use --dry-run to preview)")
	}
	return false, nil
}
