## flexera-cli finops-onboarding bill-connect databricks create

Create a Databricks bill connect

### Synopsis

Create a Databricks bill connect

Creates a new Databricks bill connect using Account API credentials.

```
flexera-cli finops-onboarding bill-connect databricks create [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-onboarding bill-connect databricks create --org-id ORG_ID --body @request.json
  flexera-cli finops-onboarding bill-connect databricks create --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-onboarding bill-connect databricks create --example > request.json
```

### Options

```
      --account-id string           accountId (body); required by API; The Databricks account ID; minLength: 1; illustrative example: "e9b1c2d3-4567-8901-2345-6789abcdef01"
      --body string                 raw JSON body (inline | @file | @-); overrides body field flags
      --body-client-id string       clientId (body); required by API; The Databricks service principal Application ID; minLength: 1; illustrative example: "061a44ab-a1bb-4b73-9548-330acadd7cd8"
      --body-client-secret string   clientSecret (body); required by API; The Databricks service principal secret
      --dry-run                     print the planned operation as JSON and exit without calling the API
  -h, --help                        help for create
  -i, --interactive                 edit inputs in a terminal form, review a plan and approve with typed yes
      --sql-warehouse-id string     sqlWarehouseId (body); required by API; Databricks SQL warehouse ID for cost and usage queries; minLength: 1; illustrative example: "5d31479ca2dfbd90"
      --workspace-url string        workspaceUrl (body); required by API; Databricks workspace URL; minLength: 1; illustrative example: "https://dbc-6d0b35df-baa3.cloud.databricks.com"
      --yes                         confirm the operation (required for destructive ops)
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

* [flexera-cli finops-onboarding bill-connect databricks](flexera-cli_finops-onboarding_bill-connect_databricks.md)	 - Bill Connect - Databricks operations (generated from the unified OpenAPI spec)

