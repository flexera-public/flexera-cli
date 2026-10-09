## flexera-cli risk misconfiguration-ui remediation-steps

Remediation Steps for Risk.

### Synopsis

Remediation Steps for Risk.

Endpoint that forwards request for remediation steps to the Secops UI API.

```
flexera-cli risk misconfiguration-ui remediation-steps [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli risk misconfiguration-ui remediation-steps --org-id ORG_ID --risk-id RISK_ID --remediation-method REMEDIATION_METHOD
```

### Options

```
  -h, --help                        help for remediation-steps
      --remediation-method string   remediationMethod (path, required); Remediation method (e.g., code, cli); required by API
      --risk-id string              riskId (path, required); Risk ID; required by API
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

