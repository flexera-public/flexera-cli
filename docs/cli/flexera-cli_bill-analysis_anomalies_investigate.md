## flexera-cli bill-analysis anomalies investigate

AI-powered cost anomaly investigation (multi-step workflow)

```
flexera-cli bill-analysis anomalies investigate [flags]
```

### Examples

```
flexera-cli bill-analysis anomalies investigate --org-id 123 --input '{"granularity":"day","recency":"P7D","increase_only":true}'
```

### Options

```
      --file string              Path to JSON input, or - for stdin
  -h, --help                     help for investigate
      --input string             Inline JSON input (mutually exclusive with --file)
      --optima-base-url string   Override the Optima base URL
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

