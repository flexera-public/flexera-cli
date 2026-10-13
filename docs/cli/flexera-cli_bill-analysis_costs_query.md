## flexera-cli bill-analysis costs query

Query costs (auto-routes aggregated/select/export, auto-chunks long windows, auto-slices select)

```
flexera-cli bill-analysis costs query [flags]
```

### Examples

```
flexera-cli bill-analysis costs query --org-id 123 --start 2024-01 --end 2024-04
```

### Options

```
      --adj-dimension-granularity string   Adjustment dimension granularity: none|org-level (export only)
      --billing-center-ids string          Comma-separated BC IDs (default: all top-level BCs)
      --body string                        Raw JSON request body (inline | @file | @-); overrides query flags
      --dimensions string                  Comma-separated cost dimensions (auto-routes to /select when >10 dimensions, resource_id, or line-items)
      --end string                         Exclusive end (YYYY-MM or YYYY-MM-DD)
      --endpoint string                    auto (default), aggregated, select, or export (default "auto")
      --export                             Use /export/select endpoint (async, limited to 1 per day per org)
      --export-poll-interval duration      Polling interval for export status (default 5s)
      --filter string                      Filter expression as inline JSON or @file (SQL-like WHERE filter)
      --granularity string                 Period granularity: month|day (default "month")
  -h, --help                               help for query
      --limit int                          Row cap (for /select capped at 100,000)
      --line-items                         Force /select (with export, sets aggregate_to_requested_dimensions=false)
      --metrics string                     Comma-separated metrics (default: cost_amortized_unblended_adj)
      --no-chunk                           Disable automatic >24-month period chunking
      --no-slice                           Disable automatic slicing of /select requests hitting row cap
      --no-wait                            Return immediately after export starts without waiting for completion
      --optima-base-url string             Override the Optima base URL (env: FLEXERA_CLI_OPTIMA_BASE_URL)
      --period-type string                 Period type: charge_period|billing_period
      --progress-interval duration         Interval for progress updates during slicing/export (default 30s)
      --start string                       Inclusive start (YYYY-MM or YYYY-MM-DD)
      --summarized                         Collapse all buckets into one row (aggregated only)
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

