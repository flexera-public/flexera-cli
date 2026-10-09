## flexera-cli bill-analysis anomalies report

report anomalies

### Synopsis

report anomalies

Generate an anomaly report for the provided dimensions

```
flexera-cli bill-analysis anomalies report [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli bill-analysis anomalies report --org-id ORG_ID --body @request.json
Validated illustrative body, when available (review before use):
  flexera-cli cli schema bill-analysis anomalies report --example > request.json
```

### Options

```
      --billing-center-ids strings   billingCenterIds (body); required by API; IDs of BillingCenters to get data for. It is not allowed for any of the BillingCenterIDs to be an ancestor of another specified BillingCenterID.; minItems: 1; CLI: comma-separated values or repeated flag; illustrative example: ["unallocated"]
      --body string                  raw JSON body (inline | @file | @-); overrides body field flags
      --detection-method string      detectionMethod (body); Specifies the detection method to use: either Bollinger Band or AI Model.; enum: ["bollinger_band","ai_model"]; API default: "bollinger_band"; illustrative example: "bollinger_band"
      --dimensions strings           dimensions (body); The list of supported dimensions by which to roll up the costs.; maxItems: 10; CLI: comma-separated values or repeated flag; illustrative example: ["vendor","category","service","instance_type","billing_center_id"]
      --end-at string                endAt (body); required by API; Latest timestamp (exclusive) of the costs. For month granularity: consists of a year and month in YYYY-MM format. For day granularity: consists of a year, month, and day in YYYY-MM-DD format. Will be interpreted as UTC, which is used for period boundaries. No records will be returned on or after this timestamp.; pattern: "^\\d{4}-\\d{2}(-\\d{2})?$"; illustrative example: "2019-01"
      --granularity string           granularity (body); Indicates which data source to query, having costs already aggregated up to this granularity. Choosing this granularity wisely can improve performance, as choosing to fetch 1 month of costs with 'month' granularity will be faster than fetching the same 31 days at 'day' granularity.; enum: ["day","month"]; API default: "month"; illustrative example: "month"
  -h, --help                         help for report
      --limit int                    limit (body); limit number of records to return.; format: int32; minimum: 1; illustrative example: 10
      --metric string                metric (body); required by API; Metric to perform anomaly detection on. Currently only cost metrics supported.; enum: ["cost_nonamortized_unblended_adj","cost_amortized_unblended_adj","cost_nonamortized_blended_adj","cost_amortized_blended_adj","BilledCost","ModifiedBilledCost","EffectiveCost","Mo... (see cli schema); illustrative example: "cost_amortized_blended_adj"
      --standard-deviations float    standardDeviations (body); required by API; number of standard deviations to use for bollinger band calculations; format: double; minimum: 0; API default: 2; illustrative example: 2
      --start-at string              startAt (body); required by API; Earliest timestamp (inclusive) of the returned costs. For month granularity: consists of a year and month in YYYY-MM format. For day granularity: consists of a year, month, and day in YYYY-MM-DD format. Will be interpreted as UTC, which is used for period boundaries.; pattern: "^\\d{4}-\\d{2}(-\\d{2})?$"; illustrative example: "2018-01"
      --window-size int              windowSize (body); required by API; window size to use for bollinger bands; format: int64; minimum: 1; API default: 10; illustrative example: 10
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

