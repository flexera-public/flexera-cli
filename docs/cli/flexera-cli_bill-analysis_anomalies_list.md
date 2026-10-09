## flexera-cli bill-analysis anomalies list

index anomalies

### Synopsis

index anomalies

List all anomalies for a given org, filtered by the supported dimensions and metric. The startAt and endAt times are specified in YYYY-MM-DD format. User must have the 'common:org:own' privilege to make this call.

```
flexera-cli bill-analysis anomalies list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli bill-analysis anomalies list --org-id ORG_ID --body @request.json
Validated illustrative body, when available (review before use):
  flexera-cli cli schema bill-analysis anomalies list --example > request.json
```

### Options

```
      --body string          raw JSON body (inline | @file | @-); overrides body field flags
      --end-at string        endAt (body); required by API; Latest timestamp (exclusive) of the anomaly. Consists of a year, month, and day in YYYY-MM-DD format. Will be interpreted as UTC, which is used for period boundaries. No records will be returned on or after this timestamp.; pattern: "^\\d{4}-\\d{2}-\\d{2}$"; illustrative example: "2025-02-17"
  -h, --help                 help for list
      --limit int            limit (body); Pagination limit; format: int64; minimum: 1; maximum: 1000; illustrative example: 10
      --metric string        metric (body); required by API; The metric used for the anomaly; enum: ["BilledCost","EffectiveCost"]; illustrative example: "BilledCost"
      --offset int           offset (body); Pagination offset; format: int64; minimum: 0; illustrative example: 0
      --sort-column string   sortColumn (body); Sort column; API default: "costImpact"; illustrative example: "costImpact"
      --sort-order string    sortOrder (body); Sort order; enum: ["asc","desc"]; API default: "desc"; illustrative example: "desc"
      --start-at string      startAt (body); required by API; Earliest timestamp (inclusive) of the returned anomaly. Consists of a year, month, and day in YYYY-MM-DD format. Will be interpreted as UTC, which is used for period boundaries.; pattern: "^\\d{4}-\\d{2}-\\d{2}$"; illustrative example: "2025-01-18"
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

* [flexera-cli bill-analysis anomalies](flexera-cli_bill-analysis_anomalies.md)	 - anomalies operations (generated from the unified OpenAPI spec)

