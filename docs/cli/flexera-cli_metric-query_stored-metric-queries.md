## flexera-cli metric-query stored-metric-queries

Query metrics service

```
flexera-cli metric-query stored-metric-queries [flags]
```

### Options

```
      --body string          raw JSON body (inline | @file | @-); overrides body field flags
      --dimensions strings   dimensions (body)
      --filter string        filter (body)
      --granularity string   granularity (body)
  -h, --help                 help for stored-metric-queries
      --limit int            limit (body)
      --metrics strings      metrics (body)
      --query-name string    queryName (path, required)
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

* [flexera-cli metric-query](flexera-cli_metric-query.md)	 - Metric Query operations (generated from the unified OpenAPI spec)

