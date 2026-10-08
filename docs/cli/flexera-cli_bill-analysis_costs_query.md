## flexera-cli bill-analysis costs query

Query costs (auto-routes aggregated/select, auto-chunks long windows)

```
flexera-cli bill-analysis costs query [flags]
```

### Examples

```
flexera-cli bill-analysis costs query --org-id 123 --start 2024-01 --end 2024-04
```

### Options

```
      --billing-center-ids string   Comma-separated BC IDs (default: all top-level BCs)
      --dimensions string           Comma-separated cost dimensions (auto-routes to /select when resource_id is present)
      --end string                  Exclusive end (YYYY-MM or YYYY-MM-DD)
      --endpoint string             auto (default), aggregated, or select (default "auto")
      --granularity string          Period granularity: month|day (default "month")
  -h, --help                        help for query
      --metrics string              Comma-separated metrics (default: cost_amortized_unblended_adj)
      --no-chunk                    Disable automatic >24-month period chunking
      --optima-base-url string      Override the Optima base URL (env: FLEXERA_CLI_OPTIMA_BASE_URL)
      --start string                Inclusive start (YYYY-MM or YYYY-MM-DD)
      --summarized                  Collapse all buckets into one row (aggregated only)
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

