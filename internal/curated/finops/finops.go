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
	var startAt, endAt, granularity, endpoint, dimensions, metrics, billingCenterIDs string
	var noChunk, summarized bool
	c := &cobra.Command{
		Use:     "query",
		Short:   "Query costs (auto-routes aggregated/select, auto-chunks long windows)",
		Example: "flexera-cli bill-analysis costs query --org-id 123 --start 2024-01 --end 2024-04",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			if strings.TrimSpace(startAt) == "" || strings.TrimSpace(endAt) == "" {
				return errors.New("--start and --end are required")
			}
			client, err := optimaClient(cmd, deps)
			if err != nil {
				return err
			}
			resolver := flexera.NewBillingCenterResolver(client)
			helper := flexera.New(client, resolver)
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
