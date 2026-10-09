## flexera-cli finops-customizations saved-filters create

Creates a saved filter

### Synopsis

Creates a saved filter

Creates a saved filter in the organization. Owner user ID is derived from the JWT subject claim. Billing Center IDs the caller is not entitled to see are dropped silently, from both billingCenterIds and Billing Center Level filter values. The request fails if the caller's Billing Center entitlement cannot be verified.

```
flexera-cli finops-customizations saved-filters create [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-customizations saved-filters create --org-id ORG_ID --body @request.json
  flexera-cli finops-customizations saved-filters create --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-customizations saved-filters create --example > request.json
```

### Options

```
      --billing-center-ids strings   billingCenterIds (body); Optional list of Billing Center IDs to store with the filter (max 100). IDs of Billing Centers the caller is not entitled to see are dropped silently.; maxItems: 100; CLI: comma-separated values or repeated flag; illustrative example: ["adZqre9KZci-oC9r_ZcZJw","kQm2Y7pVQoqxN4M8fCyTig"]
      --body string                  raw JSON body (inline | @file | @-); overrides body field flags
      --description string           description (body); Optional description; maxLength: 1024; illustrative example: "Compare cloud provider costs across all regions"
      --dimensions strings           dimensions (body); Optional list of dimensions for GROUP BY (max 10).; maxItems: 10; CLI: comma-separated values or repeated flag; illustrative example: ["provider","region"]
      --dry-run                      print the planned operation as JSON and exit without calling the API
  -h, --help                         help for create
  -i, --interactive                  edit inputs in a terminal form, review a plan and approve with typed yes
      --name string                  name (body); required by API; Human-readable name for the filter; maxLength: 255; illustrative example: "Cloud Compare by Region 2027-03-14"
      --visibility string            visibility (body); required by API; Filter visibility scope; enum: ["private","shared"]; illustrative example: "shared"
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

