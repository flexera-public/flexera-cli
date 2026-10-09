## flexera-cli bill-analysis bill-months list

search bill-months

### Synopsis

search bill-months

Search bill months with filters and sorting

```
flexera-cli bill-analysis bill-months list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli bill-analysis bill-months list --org-id ORG_ID
```

### Options

```
      --filter string             filter (query); Supports operators: eq, ne, sw (starts with), ge, le, and, or. Use parentheses for grouping.; illustrative example: "(billSource sw 'cbi-oi-gcp') and (billMonth eq '202507' or billMonth eq '202506') and (status eq 'processing' or status eq 'locked')"
  -h, --help                      help for list
      --limit int                 limit (query); Maximum number of records to return; format: int64; minimum: 1; maximum: 1000; API default: 100; illustrative example: 250
      --offset int                offset (query); Starting offset for pagination (if provided, uses offset-based pagination; if omitted, uses cursor-based). Ignored if skipToken is present.; format: int64; minimum: 0; maximum: 10000; illustrative example: 0
      --order-by string           orderBy (query); Format: field=direction [and field=direction]. Directions: asc, desc; illustrative example: "billMonth=asc and billSource=desc"
      --query-skip-token string   skip_token (query); Base64-encoded pagination token returned from previous search response
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

* [flexera-cli bill-analysis bill-months](flexera-cli_bill-analysis_bill-months.md)	 - bill-months operations (generated from the unified OpenAPI spec)

