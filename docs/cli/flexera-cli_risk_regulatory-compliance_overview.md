## flexera-cli risk regulatory-compliance overview

On-Prem BPC Misconfiguration Overview

### Synopsis

On-Prem BPC Misconfiguration Overview

Returns the BPC misconfiguration overview with findings count by severity (high, medium, low) for the latest run.

```
flexera-cli risk regulatory-compliance overview [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli risk regulatory-compliance overview --org-id ORG_ID --body @request.json
Validated illustrative body, when available (review before use):
  flexera-cli cli schema risk regulatory-compliance overview --example > request.json
```

### Options

```
      --body string         raw JSON body (inline | @file | @-); overrides body field flags
  -h, --help                help for overview
      --providers strings   providers (body); List of providers to filter. Use ['onprem'] for on-prem findings.; CLI: comma-separated values or repeated flag
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

* [flexera-cli risk regulatory-compliance](flexera-cli_risk_regulatory-compliance.md)	 - regulatory-compliance operations (generated from the unified OpenAPI spec)

