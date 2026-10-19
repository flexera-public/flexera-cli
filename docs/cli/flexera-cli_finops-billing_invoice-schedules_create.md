## flexera-cli finops-billing invoice-schedules create

Create an invoice schedule

### Synopsis

Create an invoice schedule

Creates one invoice schedule covering every customer in the request. Validation applies to the whole request: if any customer is invalid the entire request fails and nothing is created.

The schedule's invoice template is selected one of three ways:

- link an existing template: isCreateTemplate=false with invoiceTemplateId set
- create one: isCreateTemplate=true with template set, and the created template's id is returned as the schedule's invoiceTemplateId
- no template: isCreateTemplate=false with neither field, and the created schedule has no invoiceTemplateId

Any other combination is rejected with 422.

Automatic invoice numbering is optional. Supply invoiceNumberBase, with an optional invoiceNumberPrefix, to number every invoice the schedule generates; an invoiceNumberPrefix without an invoiceNumberBase is rejected with 400 invoice_number_prefix_requires_base. A number is allocated only after bill-lock passes: a failed generation keeps its number and a retry reuses it. If the sequence is exhausted the whole call fails with 409 invoice_number_sequence_exhausted (not retryable) and work already completed for earlier customers is retained. The schedule's counter is retained when the schedule is soft-deleted.

```
flexera-cli finops-billing invoice-schedules create [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing invoice-schedules create --org-id ORG_ID --body @request.json
  flexera-cli finops-billing invoice-schedules create --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-billing invoice-schedules create --example > request.json
```

### Options

```
      --body string                    raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run                        print the planned operation as JSON and exit without calling the API
  -h, --help                           help for create
  -i, --interactive                    edit inputs in a terminal form, review a plan and approve with typed yes
      --invoice-number-base string     invoiceNumberBase (body); Optional first sequence number for the invoice numbers this schedule generates.; maxLength: 9; pattern: "^0*[1-9][0-9]*$"; illustrative example: "uf4"
      --invoice-number-prefix string   invoiceNumberPrefix (body); Optional prefix for the invoice numbers this schedule generates. Requires invoiceNumberBase.; maxLength: 20; pattern: "^[A-Za-z0-9_-]*$"; illustrative example: "zel"
      --invoice-template-id string     invoiceTemplateId (body); Identifier of an existing invoice template to share across every customer on this schedule. Omit both this and template to create a schedule with no template.; illustrative example: "tmpl_9001"
      --is-create-template             isCreateTemplate (body); Whether to create a new invoice template from the template properties in this request, rather than link an existing one.; API default: false; illustrative example: false
      --schedule-name string           scheduleName (body); required by API; Human-readable name for the schedule.; maxLength: 255; illustrative example: "Monthly Q3 Invoice"
      --yes                            confirm the operation (required for destructive ops)
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

* [flexera-cli finops-billing invoice-schedules](flexera-cli_finops-billing_invoice-schedules.md)	 - Invoice Schedules operations (generated from the unified OpenAPI spec)

