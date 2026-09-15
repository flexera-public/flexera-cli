## flexera-cli anomalies report

report anomalies

```
flexera-cli anomalies report [flags]
```

### Options

```
      --billing-center-ids strings   billingCenterIds (body)
      --body string                  raw JSON body (inline | @file | @-); overrides body field flags
      --detection-method string      detectionMethod (body)
      --dimensions strings           dimensions (body)
      --dry-run                      print the planned operation as JSON and exit without calling the API
      --end-at string                endAt (body)
      --granularity string           granularity (body)
  -h, --help                         help for report
      --limit int                    limit (body)
      --metric string                metric (body)
      --standard-deviations float    standardDeviations (body)
      --start-at string              startAt (body)
      --window-size int              windowSize (body)
      --yes                          confirm the operation (required for destructive ops)
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

* [flexera-cli anomalies](flexera-cli_anomalies.md)	 - anomalies operations (generated from the unified OpenAPI spec)

