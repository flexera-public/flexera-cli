## flexera-cli misconfiguration-ui get-risk

Remediation Steps for Risk.

```
flexera-cli misconfiguration-ui get-risk [flags]
```

### Options

```
  -h, --help                        help for get-risk
      --remediation-method string   remediationMethod (path, required)
      --risk-id string              riskId (path, required)
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

* [flexera-cli misconfiguration-ui](flexera-cli_misconfiguration-ui.md)	 - misconfiguration-ui operations (generated from the unified OpenAPI spec)

