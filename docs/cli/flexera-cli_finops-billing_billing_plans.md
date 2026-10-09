## flexera-cli finops-billing billing plans

Overwrite an adjustment plan

### Synopsis

Overwrite an adjustment plan

Overwrite an adjustment plan, including the order of its rules.

Note that omitting an optional attribute will generally reset its value.

```
flexera-cli finops-billing billing plans [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing billing plans --org-id ORG_ID --id ID --body @request.json
  flexera-cli finops-billing billing plans --org-id ORG_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-billing billing plans --example > request.json
```

### Options

```
      --body string          raw JSON body (inline | @file | @-); overrides body field flags
      --description string   description (body); Description of the adjustment plan; illustrative example: "Gold Tier Customers enjoy our highest discounts"
      --dry-run              print the planned operation as JSON and exit without calling the API
  -h, --help                 help for plans
      --id string            id (path, required); ID of the billing plan; required by API; format: uuid; illustrative example: "c6671c74-513a-4127-b8df-81adfb65bf7f"
  -i, --interactive          edit inputs in a terminal form, review a plan and approve with typed yes
      --name string          name (body); required by API; Display name for the adjustment plan; illustrative example: "Gold Tier"
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

* [flexera-cli finops-billing billing](flexera-cli_finops-billing_billing.md)	 - Billing operations (generated from the unified OpenAPI spec)

