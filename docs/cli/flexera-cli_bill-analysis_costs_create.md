## flexera-cli bill-analysis costs create

exportSelect costs

### Synopsis

exportSelect costs

Initiates an export of cost data and returns a unique export ID for tracking the export status.

```
flexera-cli bill-analysis costs create [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli bill-analysis costs create --org-id ORG_ID --body @request.json
  flexera-cli bill-analysis costs create --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema bill-analysis costs create --example > request.json
```

### Options

```
      --adj-dimension-granularity string    adj_dimension_granularity (body); Indicates which druid datasource to query. - 'none' means no adjustment dimensions are included in the datasource, so any adjustments will already be aggregated into the final costs. - 'org-level' means org-level adjustment dimensions are included in the datasource, so there will be separate rows for raw costs and for any adjustments.; enum: ["none","org-level"]; API default: "none"; illustrative example: "none"
      --aggregate-to-requested-dimensions   aggregate_to_requested_dimensions (body); When true, the export will aggregate (sum) metrics, grouping by only the dimensions explicitly listed in the 'dimensions' parameter, resulting in fewer rows with summed metric values. When false (default), the export returns row-level data without aggregation, preserving the original granularity of the underlying cost data.; API default: false; illustrative example: false
      --billing-center-ids strings          billing_center_ids (body); required by API; IDs of BillingCenters to get cost data for. It is not allowed for any of the BillingCenterIDs to be an ancestor of another specified BillingCenterID.; minItems: 1; CLI: comma-separated values or repeated flag; illustrative example: ["unallocated"]
      --body string                         raw JSON body (inline | @file | @-); overrides body field flags
      --dimensions strings                  dimensions (body); required by API; The list of supported dimensions to return.; minItems: 1; CLI: comma-separated values or repeated flag; illustrative example: ["vendor","category","service","instance_type","billing_center_id"]
      --dry-run                             print the planned operation as JSON and exit without calling the API
      --end-at string                       end_at (body); required by API; Latest timestamp (exclusive) of the costs. For month granularity: consists of a year and month in YYYY-MM format. For day granularity: consists of a year, month, and day in YYYY-MM-DD format. Will be interpreted be UTC, which is used for period boundaries. No records will be returned on or after this timestamp.; pattern: "^\\d{4}-\\d{2}(-\\d{2})?$"; illustrative example: "2019-01"
      --granularity string                  granularity (body); Indicates which data source to query, having costs already aggregated up to this granularity. Choosing this granularity wisely can improve performance, as choosing to fetch 1 month of costs with 'month' granularity will be faster than fetching the same 31 days at 'day' granularity.; enum: ["day","month"]; API default: "month"; illustrative example: "month"
  -h, --help                                help for create
  -i, --interactive                         edit inputs in a terminal form, review a plan and approve with typed yes
      --metrics strings                     metrics (body); required by API; Metrics to return. When metric 'usage_amount' is requested, dimension 'usage_unit' must be requested in the dimensions parameter.; minItems: 1; CLI: comma-separated values or repeated flag; illustrative example: ["cost_amortized_blended_adj","cost_amortized_unblended_adj","cost_nonamortized_blended_adj","cost_nonamortized_unblended_adj","list_price_amortized_adj","list_price_nonamortized_a... (see cli schema)
      --period-type string                  period_type (body); Determines which date column to use for filtering. - 'charge_period' (default) filters by ChargePeriodStart/ChargePeriodMonth. - 'billing_period' filters by BillingPeriodStart/BillingPeriodMonth. For V1 orgs, only 'charge_period' is supported. V2 orgs support both 'charge_period' and 'billing_period'. If 'billing_period' is selected but the billing data does... (see cli schema); enum: ["charge_period","billing_period"]; API default: "charge_period"; illustrative example: "billing_period"
      --start-at string                     start_at (body); required by API; Earliest timestamp (inclusive) of the returned costs. For month granularity: consists of a year and month in YYYY-MM format. For day granularity: consists of a year, month, and day in YYYY-MM-DD format. Will be interpreted as UTC, which is used for period boundaries.; pattern: "^\\d{4}-\\d{2}(-\\d{2})?$"; illustrative example: "2018-01"
      --yes                                 confirm the operation (required for destructive ops)
```

### Options inherited from parent commands

```
      --access-token string     static bearer access token
      --api-base-url string     override API base URL
      --client-id string        OAuth client ID
      --client-secret string    OAuth client secret
      --config string           config file (default $HOME/.flexera/config.yaml)
  -d, --debug                   log HTTP requests/responses to stderr (Authorization redacted)
      --json-style string       JSON whitespace style (auto|pretty|compact) (default "auto")
      --login-base-url string   override login base URL
      --no-validate             skip API schema constraints (never JSON syntax or request data-loss checks)
      --org-id int              organization ID
      --out-fields string       project JSON output fields (comma-separated paths)
      --out-jq string           shape JSON output with a jq expression
  -o, --output string           output format (json|table)
  -r, --raw-output              write jq string results without JSON quotes (requires --out-jq)
      --refresh-token string    OAuth refresh token
      --zone string             API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli bill-analysis costs](flexera-cli_bill-analysis_costs.md)	 - costs operations (generated from the unified OpenAPI spec)

