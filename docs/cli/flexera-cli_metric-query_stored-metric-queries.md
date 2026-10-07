## flexera-cli metric-query stored-metric-queries

Query metrics service

```
flexera-cli metric-query stored-metric-queries [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli metric-query stored-metric-queries --org-id ORG_ID --query-name QUERY_NAME --body @request.json
  flexera-cli metric-query stored-metric-queries --org-id ORG_ID --query-name QUERY_NAME --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema metric-query stored-metric-queries --example > request.json
```

### Options

```
      --body string          raw JSON body (inline | @file | @-); overrides body field flags
      --dimensions strings   dimensions (body)
      --dry-run              print the planned operation as JSON and exit without calling the API
      --filter string        filter (body)
      --granularity string   granularity (body)
  -h, --help                 help for stored-metric-queries
  -i, --interactive          edit inputs in a terminal form, review a plan and approve with typed yes
      --limit int            limit (body)
      --metrics strings      metrics (body)
      --query-name string    queryName (path, required)
      --yes                  confirm the operation (required for destructive ops)
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

* [flexera-cli metric-query](flexera-cli_metric-query.md)	 - Metric Query operations (generated from the unified OpenAPI spec)

