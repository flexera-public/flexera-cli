## flexera-cli finops-onboarding

FinOps Onboarding API

### Synopsis

Commands for the FinOps Onboarding API (Flexera FinOps Onboarding API, v1).

Service id: finops_onboarding

### Options

```
  -h, --help   help for finops-onboarding
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
* [flexera-cli finops-onboarding bill-connect](flexera-cli_finops-onboarding_bill-connect.md)	 - Bill Connect operations (generated from the unified OpenAPI spec)
* [flexera-cli finops-onboarding processing-history](flexera-cli_finops-onboarding_processing-history.md)	 - Processing History operations (generated from the unified OpenAPI spec)

