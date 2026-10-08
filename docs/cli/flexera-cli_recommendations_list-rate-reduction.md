## flexera-cli recommendations list-rate-reduction

List rate-reduction recommendations (auto-resolves billing centers)

```
flexera-cli recommendations list-rate-reduction [flags]
```

### Options

```
      --billing-center-ids string   Comma-separated BC IDs (default: all top-level BCs)
  -h, --help                        help for list-rate-reduction
      --optima-base-url string      Override the Optima base URL (env: FLEXERA_CLI_OPTIMA_BASE_URL)
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

* [flexera-cli recommendations](flexera-cli_recommendations.md)	 - Optima Recommendations API

