// Package finops provides the curated Optima workflows that add behaviour on
// top of the generated Bill Analysis and Optima Recommendations commands:
//
//   - bill-analysis costs query: routes between /costs/aggregated and
//     /costs/select, defaults billing centers, and chunks long windows.
//   - recommendations list-usage-reduction|list-rate-reduction: auto-resolves
//     top-level billing centers and filters by recommendation category.
//
// Plain endpoint wrappers live in the generated command tree.
package finops

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	flexera "github.com/flexera-public/unified-go-client"
)

const flagOptimaBaseURL = "optima-base-url"

// Attach adds the curated workflows beneath their generated service commands.
func Attach(root *cobra.Command) error {
	if err := clipkg.AttachCommands(root, []string{"bill-analysis", "costs"}, newCostQueryCmd()); err != nil {
		return err
	}
	return clipkg.AttachCommands(root, []string{"recommendations"},
		newRecommendationListCmd("list-usage-reduction", "usage_reduction", "List usage-reduction recommendations (auto-resolves billing centers)"),
		newRecommendationListCmd("list-rate-reduction", "rate_reduction", "List rate-reduction recommendations (auto-resolves billing centers)"),
	)
}

func addOptimaBaseURLFlag(c *cobra.Command) {
	c.Flags().String(flagOptimaBaseURL, "", "Override the Optima base URL (env: FLEXERA_CLI_OPTIMA_BASE_URL)")
}

// optimaClient builds an Optima-routed client, requiring an org ID.
func optimaClient(cmd *cobra.Command, deps *clipkg.Deps) (*flexera.ClientWithResponses, error) {
	if err := deps.Config.RequireOrgID(); err != nil {
		return nil, err
	}
	baseURL, _ := cmd.Flags().GetString(flagOptimaBaseURL)
	return deps.Factory().NewOptimaClient(deps.Config, deps.Getenv, baseURL)
}

func newCostQueryCmd() *cobra.Command {
	var (
		startAt                 string
		startAtAlt              string
		endAt                   string
		endAtAlt                string
		granularity             string
		endpoint                string
		dimensions              string
		metrics                 string
		billingCenterIDs        string
		noChunk                 bool
		summarized              bool
		export                  bool
		noWait                  bool
		exportPollInterval      time.Duration
		lineItems               bool
		limit                   int64
		periodType              string
		adjDimensionGranularity string
		noSlice                 bool
		progressInterval        time.Duration
		filter                  string
		body                    string
	)
	c := &cobra.Command{
		Use:     "query",
		Short:   "Query costs (auto-routes aggregated/select/export, auto-chunks long windows, auto-slices select)",
		Example: "flexera-cli bill-analysis costs query --org-id 123 --start 2024-01 --end 2024-04",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())

			start := strings.TrimSpace(startAt)
			if start == "" {
				start = strings.TrimSpace(startAtAlt)
			}
			end := strings.TrimSpace(endAt)
			if end == "" {
				end = strings.TrimSpace(endAtAlt)
			}

			rawBody, err := clipkg.ResolveRawBody(body, cmd.InOrStdin())
			if err != nil {
				return err
			}
			if len(rawBody) > 0 && !json.Valid(rawBody) {
				return clipkg.Exit(2, errors.New("--body is not valid JSON"))
			}

			if len(rawBody) == 0 && (start == "" || end == "") {
				return clipkg.Exit(2, errors.New("--start and --end are required (or provide a JSON body via --body)"))
			}

			var ep flexera.BillAnalysisCostsEndpoint
			switch strings.ToLower(strings.TrimSpace(endpoint)) {
			case "", "auto":
				ep = flexera.CostsEndpointAuto
			case "aggregated":
				ep = flexera.CostsEndpointAggregated
			case "select":
				ep = flexera.CostsEndpointSelect
			case "export":
				ep = flexera.CostsEndpointExport
			default:
				return clipkg.Exit(2, fmt.Errorf("invalid --endpoint %q (expected auto, aggregated, select, or export)", endpoint))
			}

			gran := strings.ToLower(strings.TrimSpace(granularity))
			if gran != "" && gran != "month" && gran != "day" {
				return clipkg.Exit(2, fmt.Errorf("invalid --granularity %q (expected month or day)", granularity))
			}

			if lineItems && ep == flexera.CostsEndpointAggregated {
				return clipkg.Exit(2, errors.New("--line-items cannot be used with --endpoint aggregated"))
			}

			if periodType != "" && periodType != "charge_period" && periodType != "billing_period" {
				return clipkg.Exit(2, fmt.Errorf("invalid --period-type %q (expected charge_period or billing_period)", periodType))
			}

			if adjDimensionGranularity != "" && adjDimensionGranularity != "none" && adjDimensionGranularity != "org-level" {
				return clipkg.Exit(2, fmt.Errorf("invalid --adj-dimension-granularity %q (expected none or org-level)", adjDimensionGranularity))
			}

			if cmd.Flags().Changed("limit") && limit < 0 {
				return clipkg.Exit(2, errors.New("--limit must be non-negative"))
			}

			var filterObj *flexera.BillAnalysisFilterV1
			if strings.TrimSpace(filter) != "" {
				rawFilter, err := clipkg.ResolveRawBody(filter, cmd.InOrStdin())
				if err != nil {
					return err
				}
				var f flexera.BillAnalysisFilterV1
				if err := json.Unmarshal(rawFilter, &f); err != nil {
					return clipkg.Exit(2, fmt.Errorf("invalid --filter JSON: %w", err))
				}
				filterObj = &f
			}

			client, err := optimaClient(cmd, deps)
			if err != nil {
				return err
			}
			resolver := flexera.NewBillingCenterResolver(client)
			helper := flexera.NewBillAnalysisCostsQuery(client, resolver)

			req := flexera.BillAnalysisCostsQueryRequest{
				OrgID:                   deps.Config.OrgID,
				BillingCenterIDs:        splitCSV(billingCenterIDs),
				StartAt:                 start,
				EndAt:                   end,
				Granularity:             flexera.BillAnalysisCostsGranularity(gran),
				Dimensions:              splitCSV(dimensions),
				Metrics:                 splitCSV(metrics),
				Filter:                  filterObj,
				Endpoint:                ep,
				NoChunk:                 noChunk,
				Export:                  export,
				ExportNoWait:            noWait,
				LineItems:               lineItems,
				PeriodType:              periodType,
				AdjDimensionGranularity: adjDimensionGranularity,
				NoSlice:                 noSlice,
				Progress: func(msg string) {
					fmt.Fprintln(deps.Stderr, msg)
				},
				ProgressInterval: progressInterval,
			}
			if cmd.Flags().Changed("export-poll-interval") || exportPollInterval > 0 {
				req.ExportPollInterval = exportPollInterval
			}
			if cmd.Flags().Changed("limit") {
				req.Limit = &limit
			}
			if summarized {
				req.Summarized = &summarized
			}

			var resp *flexera.BillAnalysisCostsQueryResponse
			if len(rawBody) > 0 {
				resp, err = helper.QueryWithBody(cmd.Context(), deps.Config.OrgID, "application/json", bytes.NewReader(rawBody), req)
			} else {
				resp, err = helper.Query(cmd.Context(), req)
			}
			if err != nil {
				return err
			}
			return deps.Printer.Render(deps.Stdout, deps.Config.Output, resp)
		},
	}
	c.Flags().StringVar(&startAt, "start", "", "Inclusive start (YYYY-MM or YYYY-MM-DD)")
	c.Flags().StringVar(&startAtAlt, "start-at", "", "Inclusive start (alias for --start)")
	_ = c.Flags().MarkHidden("start-at")
	c.Flags().StringVar(&endAt, "end", "", "Exclusive end (YYYY-MM or YYYY-MM-DD)")
	c.Flags().StringVar(&endAtAlt, "end-at", "", "Exclusive end (alias for --end)")
	_ = c.Flags().MarkHidden("end-at")
	c.Flags().StringVar(&granularity, "granularity", "month", "Period granularity: month|day")
	c.Flags().StringVar(&endpoint, "endpoint", "auto", "auto (default), aggregated, select, or export")
	c.Flags().StringVar(&dimensions, "dimensions", "", "Comma-separated cost dimensions (auto-routes to /select when >10 dimensions, resource_id, or line-items)")
	c.Flags().StringVar(&metrics, "metrics", "", "Comma-separated metrics (default: cost_amortized_unblended_adj)")
	c.Flags().StringVar(&billingCenterIDs, "billing-center-ids", "", "Comma-separated BC IDs (default: all top-level BCs)")
	c.Flags().BoolVar(&noChunk, "no-chunk", false, "Disable automatic >24-month period chunking")
	c.Flags().BoolVar(&summarized, "summarized", false, "Collapse all buckets into one row (aggregated only)")
	c.Flags().BoolVar(&export, "export", false, "Use /export/select endpoint (async, limited to 1 per day per org)")
	c.Flags().BoolVar(&noWait, "no-wait", false, "Return immediately after export starts without waiting for completion")
	c.Flags().DurationVar(&exportPollInterval, "export-poll-interval", 5*time.Second, "Polling interval for export status")
	c.Flags().BoolVar(&lineItems, "line-items", false, "Force /select (with export, sets aggregate_to_requested_dimensions=false)")
	c.Flags().Int64Var(&limit, "limit", 0, "Row cap (for /select capped at 100,000)")
	c.Flags().StringVar(&periodType, "period-type", "", "Period type: charge_period|billing_period")
	c.Flags().StringVar(&adjDimensionGranularity, "adj-dimension-granularity", "", "Adjustment dimension granularity: none|org-level (export only)")
	c.Flags().BoolVar(&noSlice, "no-slice", false, "Disable automatic slicing of /select requests hitting row cap")
	c.Flags().DurationVar(&progressInterval, "progress-interval", 30*time.Second, "Interval for progress updates during slicing/export")
	c.Flags().StringVar(&filter, "filter", "", "Filter expression as inline JSON or @file (SQL-like WHERE filter)")
	c.Flags().StringVar(&body, "body", "", "Raw JSON request body (inline | @file | @-); overrides query flags")
	addOptimaBaseURLFlag(c)
	return c
}

func newRecommendationListCmd(use, kind, short string) *cobra.Command {
	var billingCenterIDs string
	c := &cobra.Command{
		Use: use, Short: short, Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			client, err := optimaClient(cmd, deps)
			if err != nil {
				return err
			}
			bcIDs := splitCSV(billingCenterIDs)
			if len(bcIDs) == 0 {
				resolver := flexera.NewBillingCenterResolver(client)
				resolved, rerr := resolver.TopLevelIDs(cmd.Context(), deps.Config.OrgID)
				if rerr != nil {
					return fmt.Errorf("auto-resolve billing centers: %w", rerr)
				}
				bcIDs = resolved
			}
			params := &flexera.OptimaRecommendationsRecommendationsIndexParams{BillingCenterIDs: &bcIDs}
			resp, err := client.OptimaRecommendationsRecommendationsIndexWithResponse(cmd.Context(), deps.Config.OrgID, params)
			if err != nil {
				return err
			}
			if resp.HTTPResponse == nil || resp.HTTPResponse.StatusCode != http.StatusOK {
				return optimaResponseError(resp.HTTPResponse, resp.Body)
			}
			filtered, ferr := flexera.FilterRecommendationsByCategory(resp.Body, kind)
			if ferr != nil {
				return renderJSONBody(deps, resp.Body)
			}
			return renderJSONBody(deps, filtered)
		},
	}
	c.Flags().StringVar(&billingCenterIDs, "billing-center-ids", "", "Comma-separated BC IDs (default: all top-level BCs)")
	addOptimaBaseURLFlag(c)
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

// renderJSONBody decodes confirmed JSON endpoints without the SDK's strict
// date decoding or float64 conversion. Empty success bodies stay empty.
func renderJSONBody(deps *clipkg.Deps, body []byte) error {
	if len(body) == 0 {
		return nil
	}
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.UseNumber()
	var value any
	if err := decoder.Decode(&value); err != nil {
		return fmt.Errorf("decoding Optima JSON response: %w", err)
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		if err == nil {
			return errors.New("Optima response must contain exactly one JSON value")
		}
		return fmt.Errorf("decoding Optima JSON response: %w", err)
	}
	return deps.Printer.Render(deps.Stdout, deps.Config.Output, value)
}
