## flexera-cli saas license allocations-all

Create allocation

### Synopsis

Create allocation

Creates a new allocation

```
flexera-cli saas license allocations-all [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas license allocations-all --org-id ORG_ID --license-id LICENSE_ID --term-id TERM_ID --purchase-id PURCHASE_ID --body @request.json
  flexera-cli saas license allocations-all --org-id ORG_ID --license-id LICENSE_ID --term-id TERM_ID --purchase-id PURCHASE_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema saas license allocations-all --example > request.json
```

### Options

```
      --body string          raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run              print the planned operation as JSON and exit without calling the API
  -h, --help                 help for allocations-all
  -i, --interactive          edit inputs in a terminal form, review a plan and approve with typed yes
      --license-id string    licenseId (path, required); Unique identifier of the license.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "38932"
      --match-type string    matchType (body); required by API; Specifies the condition type for the allocation rule.; enum: ["all","any"]; API default: "any"; illustrative example: "all"
      --purchase-id string   purchaseId (path, required); Unique identifier of the purchase associated with the license term.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "93214"
      --term-id string       termId (path, required); Unique identifier of the license term associated with the license.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "42214"
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

* [flexera-cli saas license](flexera-cli_saas_license.md)	 - License operations (generated from the unified OpenAPI spec)

