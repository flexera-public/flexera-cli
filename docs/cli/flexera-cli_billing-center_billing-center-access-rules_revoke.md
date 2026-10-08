## flexera-cli billing-center billing-center-access-rules revoke

Revoke a given Role from the Subject on this BillingCenter

```
flexera-cli billing-center billing-center-access-rules revoke [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli billing-center billing-center-access-rules revoke --org-id ORG_ID --billing-center BILLING_CENTER --body @request.json
  flexera-cli billing-center billing-center-access-rules revoke --org-id ORG_ID --billing-center BILLING_CENTER --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema billing-center billing-center-access-rules revoke --example > request.json
```

### Options

```
      --billing-center string   billing_center (path, required)
      --body string             raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run                 print the planned operation as JSON and exit without calling the API
  -h, --help                    help for revoke
  -i, --interactive             edit inputs in a terminal form, review a plan and approve with typed yes
      --yes                     confirm the operation (required for destructive ops)
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

