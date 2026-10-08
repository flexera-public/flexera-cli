## flexera-cli data-inventory onboarding create

Onboarding: Create

```
flexera-cli data-inventory onboarding create [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli data-inventory onboarding create --org-id ORG_ID --body @request.json
  flexera-cli data-inventory onboarding create --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema data-inventory onboarding create --example > request.json
```

### Options

```
      --account-type string             AccountType (body)
      --billing-account-id string       BillingAccountId (body)
      --body string                     raw JSON body (inline | @file | @-); overrides body field flags
      --body-client-id string           ClientId (body)
      --body-client-secret string       ClientSecret (body)
      --connector-name string           ConnectorName (body)
      --dry-run                         print the planned operation as JSON and exit without calling the API
      --exclude-region strings          ExcludeRegion (body)
      --external-id string              ExternalId (body)
      --have-billing-access string      HaveBillingAccess (body)
  -h, --help                            help for create
      --include-bpc string              IncludeBPC (body)
      --include-cost-and-usage string   IncludeCostAndUsage (body)
      --include-inventory string        IncludeInventory (body)
  -i, --interactive                     edit inputs in a terminal form, review a plan and approve with typed yes
      --partner-tenant-id string        PartnerTenantId (body)
      --provider string                 Provider (body)
      --role-arn string                 RoleARN (body)
      --subscription-id string          SubscriptionId (body)
      --tenant-id string                TenantId (body)
      --token-url string                TokenUrl (body)
      --yes                             confirm the operation (required for destructive ops)
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

* [flexera-cli data-inventory onboarding](flexera-cli_data-inventory_onboarding.md)	 - Onboarding operations (generated from the unified OpenAPI spec)

