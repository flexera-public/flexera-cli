## flexera-cli finops-billing billing-orphans mappings create

Create Azure cost assignments

### Synopsis

Create Azure cost assignments

Atomically create one assignment per selected charge description at the requested frequency.
Each charges entry contains an exact FOCUS ChargeDescription returned by
discovery, not a line-item type. Product-order and other display dimensions do not
narrow the selection.
Each entry includes the discovery-time billing amount,
currency, billing profile, product order, and customer metadata. Validate submitted
metadata against current eligible rows and persist it directly on the mapping. These
values are display metadata, not matching criteria or amounts applied by the pipeline.
Each charges entry supplies chargeFrequency. Every request requires billMonth to
validate selections against that month's current inventory. Only one-time mappings
store that month; recurring mappings never store it and remain month-independent.
Duplicate description/frequency pairs are invalid.
An existing mapping for any selected scoped description and
frequency returns 409 with no partial creation. Validate destination authorization
within the organization and billing account. Customer boundaries within that
billing account are not enforced.
No native subscription or charge ID is accepted from the client.
The returned count means configuration saved, not normalization completed.
A lost-response retry can conflict; refresh assignments before retrying.

```
flexera-cli finops-billing billing-orphans mappings create [flags]
```
### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing billing-orphans mappings create --org-id ORG_ID --body @request.json
  flexera-cli finops-billing billing-orphans mappings create --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-billing billing-orphans mappings create --example > request.json
```

### Options

```
      --bill-month string              billMonth (body); required by API; Discovery month or one-time assignment month in YYYYMM format. Omitted on recurring assignment resources.; pattern: "^\\d{4}(0[1-9]|1[0-2])$"; illustrative example: "202610"
      --bill-source string             billSource (body); required by API; Opaque source identifier returned by orphaned billing-account discovery.; minLength: 1; illustrative example: "azure-csp-source-a"
      --billing-account-id string      billingAccountId (body); required by API; Azure billing account identifier.; minLength: 1; illustrative example: "billing-account-a"
      --body string                    raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run[=true|false]           print the planned operation as JSON and exit without calling the API
  -h, --help[=true|false]              help for create
  -i, --interactive[=true|false]       edit inputs in a terminal form, review a plan and approve with typed yes
      --target-sub-account-id string   targetSubAccountId (body); required by API; Authorized destination subscription selected through Cost Explorer.; minLength: 1; illustrative example: "e9a0360e-2c18-428e-acd0-a76dcee0cc96"
      --yes[=true|false]               confirm the operation (required for destructive ops)
```

### Options inherited from parent commands

```
      --access-token string        static bearer access token
      --api-base-url string        override API base URL
      --client-id string           OAuth client ID
      --client-secret string       OAuth client secret
      --config string              config file (default $HOME/.flexera/config.yaml)
  -d, --debug[=true|false]         log HTTP requests/responses to stderr (Authorization redacted)
      --json-style string          JSON whitespace style (auto|pretty|compact) (default "auto")
      --login-base-url string      override login base URL
      --no-validate[=true|false]   skip API schema constraints (never JSON syntax or request data-loss checks)
      --org-id int                 organization ID
      --out-fields string          project JSON output fields (comma-separated paths)
      --out-jq string              shape JSON output with a jq expression
  -o, --output string              output format (json|table)
  -r, --raw-output[=true|false]    write jq string results without JSON quotes (requires --out-jq)
      --refresh-token string       OAuth refresh token
      --zone string                API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli finops-billing billing-orphans mappings](flexera-cli_finops-billing_billing-orphans_mappings.md)	 - mappings operations

