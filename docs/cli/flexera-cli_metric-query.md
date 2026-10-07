## flexera-cli metric-query

Metric Query operations (generated from the unified OpenAPI spec)

### Options

```
  -h, --help   help for metric-query
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

* [flexera-cli](flexera-cli.md)	 - Flexera One unified API command-line client
* [flexera-cli metric-query event-counts-by-type](flexera-cli_metric-query_event-counts-by-type.md)	 - Event counts by type
* [flexera-cli metric-query stored-metric-queries](flexera-cli_metric-query_stored-metric-queries.md)	 - Query metrics service
* [flexera-cli metric-query total-event-counts](flexera-cli_metric-query_total-event-counts.md)	 - Total event counts
* [flexera-cli metric-query user-counts-by-events-performed](flexera-cli_metric-query_user-counts-by-events-performed.md)	 - User counts by events performed

