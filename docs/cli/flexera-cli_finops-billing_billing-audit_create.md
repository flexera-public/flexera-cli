## flexera-cli finops-billing billing-audit create

Aggregate rule audit data for a month

### Synopsis

Aggregate rule audit data for a month

Group and sum audit-view rows by the requested columns for a single month (as defined by periodType), returning billed and effective cost impact per group.

Request body: Audit query body.

```
flexera-cli finops-billing billing-audit create [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing billing-audit create --org-id ORG_ID --body @request.json
  flexera-cli finops-billing billing-audit create --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-billing billing-audit create --example > request.json
```

### Options

```
      --aggregation-columns strings   aggregationColumns (body); required by API; Columns to group by. Order is preserved in the response.; minItems: 1; CLI: comma-separated values or repeated flag; items.enum: ["billingAccountId","billingAccountName","subAccountId","subAccountName","serviceName","serviceCategory","chargeCategory","customer","billingCenterLevel1","billingCenterLevel2","bi... (see cli schema); illustrative example: ["billingAccountId","billingAccountName","subAccountId","subAccountName","ruleName"]
      --body string                   raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run                       print the planned operation as JSON and exit without calling the API
  -h, --help                          help for create
  -i, --interactive                   edit inputs in a terminal form, review a plan and approve with typed yes
      --period-type yearMonth         periodType (body); Which timestamp column yearMonth applies to: `chargePeriod` (when the resource was consumed) or `billPeriod` (when the charge appeared on the bill). Mirrors the UI's View Settings → Period Type radio. Defaults to `chargePeriod` when omitted.; enum: ["chargePeriod","billPeriod"]; API default: "chargePeriod"; illustrative example: "chargePeriod"
      --year-month string             yearMonth (body); required by API; Year-month in YYYYMM (audit grain is monthly); pattern: "^\\d{4}(0[1-9]|1[0-2])$"; illustrative example: "202605"
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

* [flexera-cli finops-billing billing-audit](flexera-cli_finops-billing_billing-audit.md)	 - Billing Audit operations (generated from the unified OpenAPI spec)

