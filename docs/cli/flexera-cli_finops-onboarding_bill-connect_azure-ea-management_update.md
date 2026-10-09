## flexera-cli finops-onboarding bill-connect azure-ea-management update

Update an Azure EA (Enterprise Agreement) Management bill connect

### Synopsis

Update an Azure EA (Enterprise Agreement) Management bill connect

Modifies an existing Azure EA (Enterprise Agreement) Management bill connect associated with a given bill connect ID.
Bill Connects provisioned through Unified Onboarding are read-only in this API. Update and delete operations will be rejected with a 403 Forbidden. Use Unified Onboarding to manage these resources.
You can identify these Bill Connects by the onboardingOrigin field:
- "platform": created via Unified Onboarding
- "finops": created and managed through this API (default)

```
flexera-cli finops-onboarding bill-connect azure-ea-management update [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-onboarding bill-connect azure-ea-management update --org-id ORG_ID --id ID --body @request.json
  flexera-cli finops-onboarding bill-connect azure-ea-management update --org-id ORG_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-onboarding bill-connect azure-ea-management update --example > request.json
```

### Options

```
      --billing-account-id string     billingAccountId (body); Azure Billing Account ID cannot be updated; illustrative example: "Ducimus saepe error ea."
      --body string                   raw JSON body (inline | @file | @-); overrides body field flags
      --body-client-id string         clientId (body); required by API; client id in an Azure EA; minLength: 1; illustrative example: "56fff7875-b6d7-7f5j-zx77-977df67fdfc7"
      --body-client-secret string     clientSecret (body); required by API; client secret in an Azure EA
      --dry-run                       print the planned operation as JSON and exit without calling the API
  -h, --help                          help for update
      --id string                     id (path, required); Identifies a bill connect; required by API; illustrative example: "cbi-oi-azure-ea-12345678"
  -i, --interactive                   edit inputs in a terminal form, review a plan and approve with typed yes
      --start-billing-period string   startBillingPeriod (body); Optional Parameter formatted as YYYYMM to let the service know what Billing Period to start ingest from. The value of startBillingPeriod is required to be on or after the enrollment started and on or before the current year month. If missing, the startBillingPeriod will remain unchanged. (The startBillingPeriod can not be unset once set. It can however, be u... (see cli schema); pattern: "^20[\\d]{2}((0[1-9])|(1[012]))$"; illustrative example: "202308"
      --tenant-id string              tenantId (body); required by API; tenant id in an Azure EA; minLength: 1; illustrative example: "5837ffw33fg7-3g6t-52mt-4fht-5fd71ejf43fi"
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

* [flexera-cli finops-onboarding bill-connect azure-ea-management](flexera-cli_finops-onboarding_bill-connect_azure-ea-management.md)	 - Bill Connect - Azure EA Management operations (generated from the unified OpenAPI spec)

