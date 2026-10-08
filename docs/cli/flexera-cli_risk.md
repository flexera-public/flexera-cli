## flexera-cli risk

Risk Management API

### Synopsis

Commands for the Risk Management API (Flexera Risk Management API, v1).

Service id: risk

### Options

```
  -h, --help   help for risk
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

* [flexera-cli](flexera-cli.md)	 - Flexera One unified API command-line client
* [flexera-cli risk compliance](flexera-cli_risk_compliance.md)	 - compliance operations (generated from the unified OpenAPI spec)
* [flexera-cli risk misconfiguration-ui](flexera-cli_risk_misconfiguration-ui.md)	 - misconfiguration-ui operations (generated from the unified OpenAPI spec)
* [flexera-cli risk notifications](flexera-cli_risk_notifications.md)	 - notifications operations (generated from the unified OpenAPI spec)
* [flexera-cli risk regulatory-compliance](flexera-cli_risk_regulatory-compliance.md)	 - regulatory-compliance operations (generated from the unified OpenAPI spec)
* [flexera-cli risk vulnerability](flexera-cli_risk_vulnerability.md)	 - vulnerability operations (generated from the unified OpenAPI spec)
* [flexera-cli risk vulnerability-local](flexera-cli_risk_vulnerability-local.md)	 - vulnerability-local operations (generated from the unified OpenAPI spec)

