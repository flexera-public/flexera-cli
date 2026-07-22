## flexera-cli finops

Cost analytics, billing centers, and recommendations (Optima)

### Examples

```
flexera-cli finops cost get --org-id 123 --start 2024-01 --end 2024-04
```

### Options

```
  -h, --help                     help for finops
      --optima-base-url string   Override the Optima base URL (env: FLEXERA_CLI_OPTIMA_BASE_URL)
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

* [flexera-cli](flexera-cli.md)	 - Flexera One unified API command-line client
* [flexera-cli finops adjustment](flexera-cli_finops_adjustment.md)	 - Adjustment definition
* [flexera-cli finops anomaly-report](flexera-cli_finops_anomaly-report.md)	 - Run an anomaly report
* [flexera-cli finops bill-month](flexera-cli_finops_bill-month.md)	 - Bill months
* [flexera-cli finops billing-center](flexera-cli_finops_billing-center.md)	 - Billing centers
* [flexera-cli finops cost](flexera-cli_finops_cost.md)	 - Cost queries and exports
* [flexera-cli finops forecast-report](flexera-cli_finops_forecast-report.md)	 - Run a forecast report
* [flexera-cli finops recommendation](flexera-cli_finops_recommendation.md)	 - Optima recommendations

