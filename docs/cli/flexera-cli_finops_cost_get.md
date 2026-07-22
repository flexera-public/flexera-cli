## flexera-cli finops cost get

Curated cost query (auto-routes aggregated/select, auto-chunks windows)

```
flexera-cli finops cost get [flags]
```

### Options

```
      --billing-center-ids string   Comma-separated BC IDs (default: all top-level BCs)
      --dimensions string           Comma-separated cost dimensions (auto-routes to /select when resource_id is present)
      --end string                  Exclusive end (YYYY-MM or YYYY-MM-DD)
      --endpoint string             auto (default), aggregated, or select (default "auto")
      --granularity string          Period granularity: month|day (default "month")
  -h, --help                        help for get
      --metrics string              Comma-separated metrics (default: cost_amortized_unblended_adj)
      --no-chunk                    Disable automatic >24-month period chunking
      --start string                Inclusive start (YYYY-MM or YYYY-MM-DD)
      --summarized                  Collapse all buckets into one row (aggregated only)
```

### Options inherited from parent commands

```
      --access-token string      static bearer access token
      --api-base-url string      override API base URL
      --client-id string         OAuth client ID
      --client-secret string     OAuth client secret
      --config string            config file (default $HOME/.flexera/config.yaml)
  -d, --debug                    log HTTP requests/responses to stderr (Authorization redacted)
      --login-base-url string    override login base URL
      --optima-base-url string   Override the Optima base URL (env: FLEXERA_CLI_OPTIMA_BASE_URL)
      --org-id int               organization ID
  -o, --output string            output format (json|table)
      --refresh-token string     OAuth refresh token
      --zone string              API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli finops cost](flexera-cli_finops_cost.md)	 - Cost queries and exports

