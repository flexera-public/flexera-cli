## flexera-cli finops cost

Cost queries and exports

```
flexera-cli finops cost [flags]
```

### Options

```
  -h, --help   help for cost
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

* [flexera-cli finops](flexera-cli_finops.md)	 - Cost analytics, billing centers, and recommendations (Optima)
* [flexera-cli finops cost aggregated](flexera-cli_finops_cost_aggregated.md)	 - Raw POST /costs/aggregated
* [flexera-cli finops cost dimensions](flexera-cli_finops_cost_dimensions.md)	 - List available cost dimensions
* [flexera-cli finops cost export-select](flexera-cli_finops_cost_export-select.md)	 - Raw POST /costs/select/export
* [flexera-cli finops cost export-status](flexera-cli_finops_cost_export-status.md)	 - Check an export's status
* [flexera-cli finops cost get](flexera-cli_finops_cost_get.md)	 - Curated cost query (auto-routes aggregated/select, auto-chunks windows)
* [flexera-cli finops cost metrics](flexera-cli_finops_cost_metrics.md)	 - List available cost metrics
* [flexera-cli finops cost select](flexera-cli_finops_cost_select.md)	 - Raw POST /costs/select

