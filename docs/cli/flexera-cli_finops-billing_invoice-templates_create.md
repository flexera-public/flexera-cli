## flexera-cli finops-billing invoice-templates create

Create an invoice template

### Synopsis

Create an invoice template

Creates an invoice template in the organization.

```
flexera-cli finops-billing invoice-templates create [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing invoice-templates create --org-id ORG_ID --body @request.json
  flexera-cli finops-billing invoice-templates create --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-billing invoice-templates create --example > request.json
```

### Options

```
      --body string               raw JSON body (inline | @file | @-); overrides body field flags
      --cost-metric string        costMetric (body); required by API; Configured cost metric.; enum: ["billedCost","effectiveCost","modifiedBilledCost","modifiedEffectiveCost"]; illustrative example: "billedCost"
      --dimensions strings        dimensions (body); Optional grouping dimensions.; maxItems: 10; CLI: comma-separated values or repeated flag; illustrative example: ["service","account","region"]
      --display-logo-within-pdf   displayLogoWithinPdf (body); required by API; Whether the logo is rendered within PDF exports.; illustrative example: true
      --dry-run                   print the planned operation as JSON and exit without calling the API
      --export-type string        exportType (body); required by API; Invoice export type.; enum: ["pdf","csv"]; illustrative example: "pdf"
      --free-text string          freeText (body); Optional free-text company/address block.; illustrative example: "Aut accusamus."
  -h, --help                      help for create
  -i, --interactive               edit inputs in a terminal form, review a plan and approve with typed yes
      --logo string               logo (body); Optional logo payload. Ignored when exportType=csv.; illustrative example: "\u003cbase64-or-reference-payload\u003e"
      --period-type string        periodType (body); required by API; Configured period type.; enum: ["chargePeriod","billPeriod","billingPeriod"]; illustrative example: "chargePeriod"
      --template-name string      templateName (body); required by API; Human-readable template name.; maxLength: 255; illustrative example: "Standard MSP Invoice"
      --yes                       confirm the operation (required for destructive ops)
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

* [flexera-cli finops-billing invoice-templates](flexera-cli_finops-billing_invoice-templates.md)	 - Invoice Templates operations (generated from the unified OpenAPI spec)

