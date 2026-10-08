## flexera-cli risk misconfiguration-ui failed-asset

Failed Assets List.

```
flexera-cli risk misconfiguration-ui failed-asset [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli risk misconfiguration-ui failed-asset --org-id ORG_ID --body @request.json
Validated illustrative body, when available (review before use):
  flexera-cli cli schema risk misconfiguration-ui failed-asset --example > request.json
```

### Options

```
      --accounts strings         accounts (body)
      --args string              args (body)
      --body string              raw JSON body (inline | @file | @-); overrides body field flags
      --body-skip-token string   skipToken (body)
  -h, --help                     help for failed-asset
      --page-size int            pageSize (body)
      --providers strings        providers (body)
      --regions strings          regions (body)
      --rule-name string         ruleName (body)
      --show-suppressed string   showSuppressed (body)
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

* [flexera-cli risk misconfiguration-ui](flexera-cli_risk_misconfiguration-ui.md)	 - misconfiguration-ui operations (generated from the unified OpenAPI spec)

