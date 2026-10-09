## flexera-cli risk misconfiguration-ui failed-asset

Failed Assets List.

### Synopsis

Failed Assets List.

Endpoint that forwards api request for listing of failed assets to the Secops UI API.

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
      --accounts strings         accounts (body); required by API; List of cloud account IDs or ['all']; minItems: 1; CLI: comma-separated values or repeated flag; illustrative example: ["837570591364","252277358118"]
      --args string              args (body); Rule arguments string (comma-separated or JSON array string); minLength: 1; illustrative example: "ingress,source"
      --body string              raw JSON body (inline | @file | @-); overrides body field flags
      --body-skip-token string   skipToken (body); Cursor-based pagination token for next page
  -h, --help                     help for failed-asset
      --page-size int            pageSize (body); Number of items per page; minimum: 1; maximum: 1000; illustrative example: 100
      --providers strings        providers (body); List of cloud providers or ['all']; minItems: 1; CLI: comma-separated values or repeated flag; illustrative example: ["aws","azure"]
      --regions strings          regions (body); List of cloud regions or ['all']; minItems: 1; CLI: comma-separated values or repeated flag; illustrative example: ["us-east-1","us-east-2"]
      --rule-name string         ruleName (body); required by API; Security rule name to filter assets by; minLength: 1; illustrative example: "vpc-default-network-acls-allow-all"
      --show-suppressed string   showSuppressed (body); Whether to include suppressed assets: 'true', 'false', or omit for all; minLength: 1; illustrative example: "false"
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

