## flexera-cli saas license terms-all

Create license term

### Synopsis

Create license term

Creates a license term associated with license agreement.

```
flexera-cli saas license terms-all [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas license terms-all --org-id ORG_ID --license-id LICENSE_ID --body @request.json
  flexera-cli saas license terms-all --org-id ORG_ID --license-id LICENSE_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema saas license terms-all --example > request.json
```

### Options

```
      --additional string                  additional (body); JSON encoded map of additional license term information (if any).; format: binary; illustrative example: "{\"product_name\":\"Sample Product\"}"
      --body string                        raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run                            print the planned operation as JSON and exit without calling the API
      --effective-at string                effectiveAt (body); Start date associated with license agreement.; illustrative example: "2019-10-15T16:05:36.1000Z"
      --effective-date-reminder-days int   effectiveDateReminderDays (body); Effective date reminder day count.; maximum: 10950; illustrative example: 5
      --end-date-reminder-days int         endDateReminderDays (body); End date reminder day count.; maximum: 10950; illustrative example: 14
      --ends-at string                     endsAt (body); End date associated with license agreement.; illustrative example: "2020-12-30T00:00:00.1000Z"
  -h, --help                               help for terms-all
      --hidden                             hidden (body); Boolean value to check license term is explicitly removed by customer or not.; illustrative example: true
  -i, --interactive                        edit inputs in a terminal form, review a plan and approve with typed yes
      --is-active                          isActive (body); Determines whether the license term is active.; illustrative example: true
      --is-vendor-set-provisioned-count    isVendorSetProvisionedCount (body); Flag to indicate if 'provisionedCount' is determined through vendor api for auto-generated licenses.; illustrative example: true
      --license-id string                  licenseId (path, required); Unique identifier of license.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "38932"
      --managed-app-id string              managedAppId (body); required by API; It identifies a managed application by given ID.; pattern: "^[0-9a-f]+$"; illustrative example: "38932"
      --name string                        name (body); required by API; Name of the licensed product.; illustrative example: "Workday"
      --sku string                         sku (body); License term SKU.; illustrative example: "c7df2760-2c81-4ef7-b578-5b5392b571df"
      --term-id string                     termId (body); Identifies a license term associated with the license.; pattern: "^[0-9a-f]+$"; illustrative example: "42214"
      --term-type string                   termType (body); required by API; Term type is to distinguish types of license terms.; enum: ["perUser","perLogin","perNode","perLocation","perOrg"]; illustrative example: "perUser"
      --unique-id string                   uniqueId (body); Vendor unique ID associated with license term.; illustrative example: "48a80680-7326-48cd-9935-b556b81d3"
      --yes                                confirm the operation (required for destructive ops)
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

