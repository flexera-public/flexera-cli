## flexera-cli risk compliance

compliance operations (generated from the unified OpenAPI spec)

### Options

```
  -h, --help   help for compliance
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

* [flexera-cli risk](flexera-cli_risk.md)	 - Risk Management API
* [flexera-cli risk compliance cis](flexera-cli_risk_compliance_cis.md)	 - Get CIS Benchmark Details.
* [flexera-cli risk compliance compliances](flexera-cli_risk_compliance_compliances.md)	 - Security Compliance.
* [flexera-cli risk compliance control](flexera-cli_risk_compliance_control.md)	 - Get Standard Control Details.
* [flexera-cli risk compliance export](flexera-cli_risk_compliance_export.md)	 - Export Compliance Chart Data.
* [flexera-cli risk compliance standard](flexera-cli_risk_compliance_standard.md)	 - Compliance Standards.
* [flexera-cli risk compliance trend](flexera-cli_risk_compliance_trend.md)	 - Compliance Favorites.
* [flexera-cli risk compliance update](flexera-cli_risk_compliance_update.md)	 - Toggle Compliance Standard Favorite Status.

