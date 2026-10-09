## flexera-cli billing-center replace

Update a BillingCenter

### Synopsis

Update a BillingCenter

Update a BillingCenter

**Required security scopes for GlobalSession**:
  * `optima:billing_center:update+common:org:own`

```
flexera-cli billing-center replace [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli billing-center replace --org-id ORG_ID --billing-center BILLING_CENTER --body @request.json
  flexera-cli billing-center replace --org-id ORG_ID --billing-center BILLING_CENTER --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema billing-center replace --example > request.json
```

### Options

```
      --billing-center string   billing_center (path, required); required by API
      --body string             raw JSON body (inline | @file | @-); overrides body field flags
      --description string      description (body); required by API; Description of the BillingCenter; illustrative example: "cloud resources used in marketing campaigns"
      --dry-run                 print the planned operation as JSON and exit without calling the API
  -h, --help                    help for replace
  -i, --interactive             edit inputs in a terminal form, review a plan and approve with typed yes
      --name string             name (body); required by API; Name of the BillingCenter; illustrative example: "Marketing"
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

* [flexera-cli billing-center](flexera-cli_billing-center.md)	 - Billing Center API

