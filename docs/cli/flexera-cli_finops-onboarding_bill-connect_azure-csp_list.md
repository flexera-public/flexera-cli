## flexera-cli finops-onboarding bill-connect azure-csp list

Validate the credentials stored for an Azure CSP bill connect

### Synopsis

Validate the credentials stored for an Azure CSP bill connect

Validates the credentials stored for an Azure CSP bill connect associated with the provided bill connect ID.

```
flexera-cli finops-onboarding bill-connect azure-csp list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-onboarding bill-connect azure-csp list --org-id ORG_ID --id ID
```

### Options

```
  -h, --help        help for list
      --id string   id (path, required); Identifies a bill connect; required by API; illustrative example: "cbi-oi-azure-csp-605fd6aa9c95c35bc2ddff0c651b182aefc900fa2bc83ea035fed434caef97b5"
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

* [flexera-cli finops-onboarding bill-connect azure-csp](flexera-cli_finops-onboarding_bill-connect_azure-csp.md)	 - Bill Connect - Azure CSP operations (generated from the unified OpenAPI spec)

