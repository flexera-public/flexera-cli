## flexera-cli finops-onboarding bill-connect azure-csp create

Create an Azure CSP bill connect

### Synopsis

Create an Azure CSP bill connect

Creates an Azure CSP bill connect using the provided credentials.

```
flexera-cli finops-onboarding bill-connect azure-csp create [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-onboarding bill-connect azure-csp create --org-id ORG_ID --body @request.json
  flexera-cli finops-onboarding bill-connect azure-csp create --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-onboarding bill-connect azure-csp create --example > request.json
```

### Options

```
      --billing-account-id string     billingAccountId (body); The Billing Account ID in an Azure CSP bill; pattern: "[0-9a-fA-F]{8}-([0-9a-fA-F]{4}-){3}[0-9a-fA-F]{12}:[0-9a-fA-F]{8}-([0-9a-fA-F]{4}-){3}[0-9a-fA-F]{12}_2[0-9][0-9]{2}-([0][1-9]|[1][0-2])-([1-2][0-9]|[0][1-9]|[3][0-1])"; illustrative example: "23lopi-7875-b6d7-ploip-zx77-pppdf67fdfc7:32662f18-7ca4-4845-99e0-1213414d5bc4_2024-01-31"
      --body string                   raw JSON body (inline | @file | @-); overrides body field flags
      --body-client-id string         clientId (body); required by API; Client Id in Azure CSP. This is a GUID that uniquely identifies the app's registration in your Active Directory tenant; minLength: 1; illustrative example: "56fff7875-b6d7-7f5j-zx77-977df67fdfc7"
      --body-client-secret string     clientSecret (body); required by API; Client Secret in Azure CSP. References the Azure CSP SPN's Client Secret.
      --dry-run                       print the planned operation as JSON and exit without calling the API
  -h, --help                          help for create
  -i, --interactive                   edit inputs in a terminal form, review a plan and approve with typed yes
      --proxy string                  proxy (body); Proxy term used for partner integration routing; illustrative example: "example_proxy"
      --start-billing-period string   startBillingPeriod (body); Optional Parameter formatted as YYYYMM to let the service know what Billing Period to start ingest from. The value of startBillingPeriod is required to be on or after the enrollment started and on or before the current year month. If not provided, the ingest will begin from the billing period at the time of creation of the bill connect.; pattern: "^20[\\d]{2}((0[1-9])|(1[012]))$"; illustrative example: "202308"
      --tenant-id string              tenantId (body); required by API; Tenant ID in Azure CSP. This refers to the tenant/directory the app registration/SPN belongs to.; minLength: 1; illustrative example: "5837ffw33fg7-3g6t-52mt-4fht-5fd71ejf43fi"
      --yes                           confirm the operation (required for destructive ops)
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

* [flexera-cli finops-onboarding bill-connect azure-csp](flexera-cli_finops-onboarding_bill-connect_azure-csp.md)	 - Bill Connect - Azure CSP operations (generated from the unified OpenAPI spec)

