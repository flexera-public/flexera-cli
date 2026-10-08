## flexera-cli billing-center billing-center-access-rules list

List roles of users, groups or service accounts

```
flexera-cli billing-center billing-center-access-rules list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli billing-center billing-center-access-rules list --org-id ORG_ID --billing-center BILLING_CENTER
```

### Options

```
      --billing-center string   billing_center (path, required)
  -h, --help                    help for list
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

* [flexera-cli billing-center billing-center-access-rules](flexera-cli_billing-center_billing-center-access-rules.md)	 - BillingCenterAccessRules operations (generated from the unified OpenAPI spec)

