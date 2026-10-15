## flexera-cli finops-billing invoices create

Create an invoice

### Synopsis

Create an invoice

Creates invoice exports for one or more customers.

The invoice configuration can be supplied in one of four ways:

- direct fields: provide exportType, periodType, costMetric, and displayLogoWithinPdf without a template source
- existing template: provide invoiceTemplateId; omitted root configuration fields use the saved template's values
- create saved template: provide isCreateTemplate=true and template with templateName; the new template is used and its ID is returned
- one-time template: provide template without templateName and with isCreateTemplate=false (or omitted); it supplies omitted root configuration fields but is not saved

Root configuration fields always override their corresponding template value when present. template cannot be combined with invoiceTemplateId. The four required root configuration fields (exportType, periodType, costMetric, displayLogoWithinPdf) are required only when neither invoiceTemplateId nor template is supplied. exportLayout is optional: PDF exports default to flatten, while nested selects the nested layout. Values are exact lowercase flatten or nested; omit exportLayout for CSV exports.

```
flexera-cli finops-billing invoices create [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing invoices create --org-id ORG_ID --body @request.json
  flexera-cli finops-billing invoices create --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-billing invoices create --example > request.json
```

### Options

```
      --body string                  raw JSON body (inline | @file | @-); overrides body field flags
      --cost-metric string           costMetric (body); Cost metric to use for export. Required only when neither invoiceTemplateId nor template is supplied.; enum: ["billedCost","effectiveCost","modifiedBilledCost","modifiedEffectiveCost"]; illustrative example: "effectiveCost"
      --dimensions strings           dimensions (body); Dimensions to include in export.; maxItems: 10; CLI: comma-separated values or repeated flag; illustrative example: ["Odio illo ut.","Est eos alias omnis maxime excepturi.","Velit fugit rerum non modi voluptatem qui."]
      --display-logo-within-pdf      displayLogoWithinPdf (body); Whether to show logo in PDF export. Required only when neither invoiceTemplateId nor template is supplied.; illustrative example: false
      --dry-run                      print the planned operation as JSON and exit without calling the API
      --due-date string              dueDate (body); Invoice due date.; format: date; illustrative example: "2001-06-25"
      --end-date string              endDate (body); required by API; Invoice range end date.; format: date; illustrative example: "1978-08-01"
      --export-layout string         exportLayout (body); PDF layout format. Exact lowercase flatten or nested; omitted PDF exports default to flatten. Must be omitted for CSV exports.; enum: ["flatten","nested"]; illustrative example: "flatten"
      --export-type string           exportType (body); Export type. Required only when neither invoiceTemplateId nor template is supplied.; enum: ["pdf","csv"]; illustrative example: "csv"
      --free-text string             freeText (body); Optional free text.; illustrative example: "Numquam quasi aut saepe dolorum sapiente molestias."
  -h, --help                         help for create
  -i, --interactive                  edit inputs in a terminal form, review a plan and approve with typed yes
      --invoice-template-id string   invoiceTemplateId (body); Optional saved invoice template to source default field values from. Cannot be combined with template. Opaque ID; resolved server-side and snapshotted into the invoice.; illustrative example: "Sed laborum corporis repellendus nisi."
      --is-create-template           isCreateTemplate (body); Whether to create and use a reusable invoice template from template properties.; API default: false; illustrative example: false
      --issue-date string            issueDate (body); Invoice issue date.; format: date; illustrative example: "1985-06-06"
      --logo string                  logo (body); Optional logo payload.; illustrative example: "Molestias quos harum veritatis beatae quis."
      --period-type string           periodType (body); Configured period type. Required only when neither invoiceTemplateId nor template is supplied.; enum: ["chargePeriod","billPeriod","billingPeriod"]; illustrative example: "chargePeriod"
      --start-date string            startDate (body); required by API; Invoice range start date.; format: date; illustrative example: "2014-03-24"
      --yes                          confirm the operation (required for destructive ops)
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

* [flexera-cli finops-billing invoices](flexera-cli_finops-billing_invoices.md)	 - Invoices operations (generated from the unified OpenAPI spec)

