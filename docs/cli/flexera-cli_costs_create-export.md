## flexera-cli costs create-export

exportSelect costs

```
flexera-cli costs create-export [flags]
```

### Options

```
      --adj-dimension-granularity string    adj_dimension_granularity (body)
      --aggregate-to-requested-dimensions   aggregate_to_requested_dimensions (body)
      --billing-center-ids strings          billing_center_ids (body)
      --body string                         raw JSON body (inline | @file | @-); overrides body field flags
      --dimensions strings                  dimensions (body)
      --dry-run                             print the planned operation as JSON and exit without calling the API
      --end-at string                       end_at (body)
      --granularity string                  granularity (body)
  -h, --help                                help for create-export
      --metrics strings                     metrics (body)
      --start-at string                     start_at (body)
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
      --login-base-url string   override login base URL
      --org-id int              organization ID
  -o, --output string           output format (json|table)
      --refresh-token string    OAuth refresh token
      --zone string             API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli costs](flexera-cli_costs.md)	 - costs operations (generated from the unified OpenAPI spec)

