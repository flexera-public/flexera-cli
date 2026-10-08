## flexera-cli saas metric-query event-counts-by-type

Event counts by type

```
flexera-cli saas metric-query event-counts-by-type [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas metric-query event-counts-by-type --org-id ORG_ID --managed-app-id MANAGED_APP_ID --days-since DAYS_SINCE --sub-app-id SUB_APP_ID
```

### Options

```
      --days-since int          daysSince (query)
  -h, --help                    help for event-counts-by-type
      --managed-app-id string   managedAppId (query)
      --resolution string       resolution (query)
      --sub-app-id string       subAppId (query)
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

* [flexera-cli saas metric-query](flexera-cli_saas_metric-query.md)	 - Metric Query operations (generated from the unified OpenAPI spec)

