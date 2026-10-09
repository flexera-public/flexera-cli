## flexera-cli finops-customizations saved-filters replace

Updates a saved filter

### Synopsis

Updates a saved filter

Updates a saved filter in the organization. Authorization is determined by the caller's privileges and the filter's ownership/visibility. Supports optimistic locking through If-Match. Billing Center IDs the caller is not entitled to see are dropped from the submitted billingCenterIds and Billing Center Level filter values, while stored Billing Center IDs the caller cannot see are preserved. The update is rejected with 403 when the requested replacement would drop Billing Center Level filter values the caller cannot see, and with 422 when the resulting billingCenterIds, including preserved Billing Centers the caller cannot see, would exceed 100 entries. The request fails if the caller's Billing Center entitlement cannot be verified.

```
flexera-cli finops-customizations saved-filters replace [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-customizations saved-filters replace --org-id ORG_ID --id ID --body @request.json
  flexera-cli finops-customizations saved-filters replace --org-id ORG_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-customizations saved-filters replace --example > request.json
```

### Options

```
      --billing-center-ids strings   billingCenterIds (body); Optional list of Billing Center IDs to store with the filter (max 100). Full-replace semantics apply to the caller's own Billing Centers; IDs the caller is not entitled to see are dropped from the request and preserved from the stored filter. The filter stores at most 100 Billing Centers in total, including preserved ones the caller cannot see: if preserved ... (see cli schema); maxItems: 100; CLI: comma-separated values or repeated flag; illustrative example: ["adZqre9KZci-oC9r_ZcZJw"]
      --body string                  raw JSON body (inline | @file | @-); overrides body field flags
      --description string           description (body); required by API; Description; maxLength: 1024; illustrative example: "Filters for Q4 cloud cost analysis"
      --dimensions strings           dimensions (body); Optional list of dimensions for GROUP BY (max 10).; maxItems: 10; CLI: comma-separated values or repeated flag; illustrative example: ["vendor","vendor_account"]
      --dry-run                      print the planned operation as JSON and exit without calling the API
  -h, --help                         help for replace
      --id string                    id (path, required); Identifier of the saved filter (UUID); required by API; format: uuid; illustrative example: "550e8400-e29b-41d4-a716-446655440000"
  -i, --interactive                  edit inputs in a terminal form, review a plan and approve with typed yes
      --name string                  name (body); required by API; Human-readable name for the filter; maxLength: 255; illustrative example: "My Q4 Cost Filter"
      --visibility string            visibility (body); required by API; Filter visibility scope; enum: ["private","shared"]; illustrative example: "private"
      --yes                          confirm the operation (required for destructive ops)
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

* [flexera-cli finops-customizations saved-filters](flexera-cli_finops-customizations_saved-filters.md)	 - Saved Filters operations (generated from the unified OpenAPI spec)

