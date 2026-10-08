## flexera-cli bill-analysis costs

costs operations (generated from the unified OpenAPI spec)

### Options

```
  -h, --help   help for costs
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

* [flexera-cli bill-analysis](flexera-cli_bill-analysis.md)	 - Bill Analysis API
* [flexera-cli bill-analysis costs aggregated](flexera-cli_bill-analysis_costs_aggregated.md)	 - aggregated costs
* [flexera-cli bill-analysis costs create](flexera-cli_bill-analysis_costs_create.md)	 - exportSelect costs
* [flexera-cli bill-analysis costs dimensions](flexera-cli_bill-analysis_costs_dimensions.md)	 - dimensions costs
* [flexera-cli bill-analysis costs get-export](flexera-cli_bill-analysis_costs_get-export.md)	 - exportSelectStatus costs
* [flexera-cli bill-analysis costs get-select](flexera-cli_bill-analysis_costs_get-select.md)	 - select costs
* [flexera-cli bill-analysis costs metrics](flexera-cli_bill-analysis_costs_metrics.md)	 - metrics costs
* [flexera-cli bill-analysis costs query](flexera-cli_bill-analysis_costs_query.md)	 - Query costs (auto-routes aggregated/select, auto-chunks long windows)

