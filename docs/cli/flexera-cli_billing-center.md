## flexera-cli billing-center

Billing Center API

### Synopsis

Commands for the Billing Center API (RightScale Billing Center API).

Service id: billing_center_service

### Options

```
  -h, --help   help for billing-center
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

* [flexera-cli](flexera-cli.md)	 - Flexera One unified API command-line client
* [flexera-cli billing-center access-rules](flexera-cli_billing-center_access-rules.md)	 - AccessRules operations (generated from the unified OpenAPI spec)
* [flexera-cli billing-center allocation-table](flexera-cli_billing-center_allocation-table.md)	 - Update a BC AllocationTable or create it if it doesn't exist
* [flexera-cli billing-center allocation-table-all](flexera-cli_billing-center_allocation-table-all.md)	 - List the AllocationTable rules for a given BC
* [flexera-cli billing-center billing-center-access-rules](flexera-cli_billing-center_billing-center-access-rules.md)	 - BillingCenterAccessRules operations (generated from the unified OpenAPI spec)
* [flexera-cli billing-center create](flexera-cli_billing-center_create.md)	 - Create a new BillingCenter
* [flexera-cli billing-center delete](flexera-cli_billing-center_delete.md)	 - Delete a BillingCenter
* [flexera-cli billing-center get](flexera-cli_billing-center_get.md)	 - Show a single BillingCenter
* [flexera-cli billing-center list](flexera-cli_billing-center_list.md)	 - List all BillingCenters in a given Org.
* [flexera-cli billing-center org-allocation-table](flexera-cli_billing-center_org-allocation-table.md)	 - AllocationTable operations (generated from the unified OpenAPI spec)
* [flexera-cli billing-center replace](flexera-cli_billing-center_replace.md)	 - Update a BillingCenter
* [flexera-cli billing-center user-billing-centers](flexera-cli_billing-center_user-billing-centers.md)	 - UserBillingCenters operations (generated from the unified OpenAPI spec)

