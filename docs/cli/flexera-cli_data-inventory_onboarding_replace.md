## flexera-cli data-inventory onboarding replace

Onboarding: Update

### Synopsis

Onboarding: Update

Update an existing cloud connector.

```
flexera-cli data-inventory onboarding replace [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli data-inventory onboarding replace --org-id ORG_ID --connector-id CONNECTOR_ID --provider PROVIDER --body @request.json
  flexera-cli data-inventory onboarding replace --org-id ORG_ID --connector-id CONNECTOR_ID --provider PROVIDER --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema data-inventory onboarding replace --example > request.json
```

### Options

```
      --account-type string             AccountType (body); Azure account type. Required for Azure BPC and CCO.; enum: ["EA","MCA","CSP"]
      --billing-account-id string       BillingAccountId (body); Billing Account ID.
      --body string                     raw JSON body (inline | @file | @-); overrides body field flags
      --body-client-id string           ClientId (body); Client ID.; minLength: 1
      --body-client-secret string       ClientSecret (body); Client secret.
      --connector-id string             connector_id (path, required); Connector (Schedule) ID; required by API
      --connector-name string           ConnectorName (body); Connector name.; minLength: 1
      --dry-run                         print the planned operation as JSON and exit without calling the API
      --exclude-region strings          ExcludeRegion (body); Regions to exclude from onboarding.; API default: []; CLI: comma-separated values or repeated flag
      --external-id string              ExternalId (body); External ID.
      --have-billing-access string      HaveBillingAccess (body); Indicates if the service principal has billing access.
  -h, --help                            help for replace
      --include-bpc string              IncludeBPC (body); Include BPC flag.
      --include-cost-and-usage string   IncludeCostAndUsage (body); Include CCO flag.
      --include-inventory string        IncludeInventory (body); Include inventory flag.
  -i, --interactive                     edit inputs in a terminal form, review a plan and approve with typed yes
      --partner-tenant-id string        PartnerTenantId (body); Partner Tenant ID (CSP only).; minLength: 1
      --provider string                 provider (query); Provider (aws/azure) - required; required by API
      --role-arn string                 RoleARN (body); Role ARN.; pattern: "^arn:(?:aws|aws-cn|aws-us-gov):iam::\\d{12}:role/.+$"
      --subscription-id string          SubscriptionId (body); Subscription ID.; minLength: 1
      --tenant-id string                TenantId (body); Azure Tenant ID.; minLength: 1
      --token-url string                TokenUrl (body); Token URL.
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

