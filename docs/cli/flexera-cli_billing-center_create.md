## flexera-cli billing-center create

Create a new BillingCenter

### Synopsis

Create a new BillingCenter

Create a new BillingCenter

**Required security scopes for GlobalSession**:
  * `optima:billing_center:create+common:org:own`

```
flexera-cli billing-center create [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli billing-center create --org-id ORG_ID --body @request.json
  flexera-cli billing-center create --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema billing-center create --example > request.json
```

### Options

```
      --body string          raw JSON body (inline | @file | @-); overrides body field flags
      --description string   description (body); required by API; Description of the BillingCenter; illustrative example: "cloud resources used in marketing campaigns"
      --dry-run              print the planned operation as JSON and exit without calling the API
  -h, --help                 help for create
  -i, --interactive          edit inputs in a terminal form, review a plan and approve with typed yes
      --name string          name (body); required by API; Name of the BillingCenter; illustrative example: "Marketing"
      --parent-href string   parent_href (body); API reference of parent billing center; pattern: "^/analytics/orgs/(\\d+)/billing_centers/([^/]+)$"; illustrative example: "/analytics/orgs/1/billing_centers/abc123"
      --yes                  confirm the operation (required for destructive ops)
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

