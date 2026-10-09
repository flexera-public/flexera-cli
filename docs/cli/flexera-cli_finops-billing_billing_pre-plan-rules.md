## flexera-cli finops-billing billing pre-plan-rules

Overwrite an adjustment rule that runs before all plans

### Synopsis

Overwrite an adjustment rule that runs before all plans

Overwrite an adjustment rule.

Note that omitting an optional attribute will generally reset its value.

```
flexera-cli finops-billing billing pre-plan-rules [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing billing pre-plan-rules --org-id ORG_ID --id ID --body @request.json
  flexera-cli finops-billing billing pre-plan-rules --org-id ORG_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-billing billing pre-plan-rules --example > request.json
```

### Options

```
      --body string        raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run            print the planned operation as JSON and exit without calling the API
      --enabled            enabled (body); required by API; A rule that is not enabled will not run.; illustrative example: true
      --end-after string   endAfter (body); The last date this rule applies.; format: date; illustrative example: "2023-05-17"
  -h, --help               help for pre-plan-rules
      --id string          id (path, required); id identifies this rule; required by API; format: uuid; illustrative example: "ae85f96e-6b9e-4d68-a741-d1ea7ca1fb28"
  -i, --interactive        edit inputs in a terminal form, review a plan and approve with typed yes
      --name string        name (body); required by API; Name of the rule.; illustrative example: "Margin"
      --start-on string    startOn (body); The first date this rule applies.; format: date; illustrative example: "2020-03-15"
      --type string        type (body); required by API; type controls what an adjustment rule does. The matching field (if any) can be populated with type-specific configuration. Current types include: - creditMemo - customUsageRate - fixedAmount - markupMarkdown - removeTax - upchargeDiscount Additional rule types (and matching fields) may be added in the future.; illustrative example: "markupMarkdown"
      --yes                confirm the operation (required for destructive ops)
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

