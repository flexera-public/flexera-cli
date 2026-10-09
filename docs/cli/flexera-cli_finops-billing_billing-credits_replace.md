## flexera-cli finops-billing billing-credits replace

Overwrite a credit assignment

### Synopsis

Overwrite a credit assignment

Rewrite a credit assignment.

A credit assignment reassigns, hides, or shows a set of credits matched by the
CreditDescriptor entries in the assignment. Each descriptor matches billing rows by
description, subAccountId, and by the billMonth of the assignment.

The chosen action is applied to all matching credits.
A credit assignment may have up to 50 CreditDescriptor entries.

CreditDescriptor entries must be unique across all credit assignments for the same bill month.
A 409 Conflict is returned when a requested (description + subAccountId + billMonth) already exists.

Provide the ETag from a prior show response to guard against concurrent modifications.

```
flexera-cli finops-billing billing-credits replace [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing billing-credits replace --org-id ORG_ID --id ID --body @request.json
  flexera-cli finops-billing billing-credits replace --org-id ORG_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-billing billing-credits replace --example > request.json
```

### Options

```
      --action string                        action (body); required by API; Disposition of this credit assignment: one of reassign, show, or hide.; enum: ["reassign","show","hide"]; illustrative example: "reassign"
      --bill-month string                    billMonth (body); required by API; Bill month in YYYYMM format this assignment applies to.; pattern: "^\\d{4}(0[1-9]|1[0-2])$"; illustrative example: "202601"
      --body string                          raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run                              print the planned operation as JSON and exit without calling the API
  -h, --help                                 help for replace
      --id string                            id (path, required); Unique identifier of the credit assignment.; required by API; format: uuid; illustrative example: "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
  -i, --interactive                          edit inputs in a terminal form, review a plan and approve with typed yes
      --name string                          name (body); required by API; Human-readable name for this assignment.; illustrative example: "EDP Credits Q1"
      --provider string                      provider (body); The unique provider affected by the credit assignment.; illustrative example: "AWS"
      --target-billing-account-id string     targetBillingAccountId (body); Required when action is reassign.; illustrative example: "4"
      --target-billing-account-name string   targetBillingAccountName (body); Used when action is reassign.; illustrative example: "acme-target-us-east"
      --target-customer-id int               targetCustomerId (body); The customer (child org) identifier to reassign the credit to. Used when action is reassign.; illustrative example: 2001
      --target-sub-account-id string         targetSubAccountId (body); The sub-account to reassign the credit to. Required when action is reassign.; illustrative example: "111122223333"
      --target-sub-account-name string       targetSubAccountName (body); The sub-account to reassign the credit to. Used when action is reassign.; illustrative example: "acme-target-us-east"
      --yes                                  confirm the operation (required for destructive ops)
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

* [flexera-cli finops-billing billing-credits](flexera-cli_finops-billing_billing-credits.md)	 - Billing Credits operations (generated from the unified OpenAPI spec)

