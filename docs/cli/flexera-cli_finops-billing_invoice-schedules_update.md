## flexera-cli finops-billing invoice-schedules update

Update an invoice schedule

### Synopsis

Update an invoice schedule

Partially updates an invoice schedule using JSON Merge Patch (RFC 7396) semantics: only fields present in the request body are changed and omitted fields remain unchanged. Supports optimistic locking through If-Match.

Invoice numbering: invoiceNumberPrefix and invoiceNumberBase follow the same merge-patch rule, so omitting a field leaves it unchanged. A schedule created without numbering can have it set once through this endpoint, as long as it has not generated an invoice yet: invoiceNumberPrefix alone for prefix-only numbering, or invoiceNumberPrefix together with invoiceNumberBase for sequential numbering; invoiceNumberBase without a prefix is rejected with 400 invoice_number_prefix_required. Once the prefix, or prefix and base, is set, the two fields are frozen: a patch that carries either of them is rejected with 400 invoice_numbering_locked, whether it changes the value, clears it with "" or repeats the stored value, and whether or not an invoice was generated yet; leave both fields out to update anything else. Setting numbering on a schedule that already generated unnumbered invoices is rejected with the same error, so every invoice of a schedule carries the numbering mode it was created with.

An inactive schedule accepts exactly three fields through this endpoint: status, invoiceNumberPrefix and invoiceNumberBase. Any other field is rejected. Use the activate endpoint to reactivate it.

Setting status to inactive through this endpoint stops the schedule's future runs, exactly as the deactivate endpoint does, but it leaves endMonthYear unchanged. Use the deactivate endpoint when the schedule's end month should also be set to the current month in the schedule timezone.

```
flexera-cli finops-billing invoice-schedules update [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing invoice-schedules update --org-id ORG_ID --id ID --body @request.json
  flexera-cli finops-billing invoice-schedules update --org-id ORG_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-billing invoice-schedules update --example > request.json
```

### Options

```
      --body string                       raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run[=true|false]              print the planned operation as JSON and exit without calling the API
  -h, --help[=true|false]                 help for update
      --id string                         id (path, required); Identifier of the invoice schedule; required by API; illustrative example: "sch_1"
  -i, --interactive[=true|false]          edit inputs in a terminal form, review a plan and approve with typed yes
      --invoice-number-base string        invoiceNumberBase (body); First sequence number for the invoice numbers this schedule generates. Can be set only on a schedule whose numbering is not set yet and that has not generated an invoice, together with invoiceNumberPrefix (400 invoice_number_prefix_required without it). Once numbering is set, sending the field again at all, even with the same value or "", is rejected with 40... (see cli schema); maxLength: 9; pattern: "^$|^0*[1-9][0-9]*$"; illustrative example: "7k5"
      --invoice-number-prefix string      invoiceNumberPrefix (body); Prefix for the invoice numbers this schedule generates. Can be set only on a schedule whose numbering is not set yet and that has not generated an invoice; without a base it is the whole invoice number of every invoice the schedule generates. Once set, sending the field again at all, even with the same value or "", is rejected with 400 invoice_numbering_lock... (see cli schema); maxLength: 20; pattern: "^[A-Za-z0-9_-]*$"; illustrative example: "ztr"
      --invoice-template-id string        invoiceTemplateId (body); Existing invoice template to point the schedule at. Send isCreateTemplate=false with this field omitted to detach the schedule's template.; illustrative example: "tmpl_9002"
      --is-create-template[=true|false]   isCreateTemplate (body); Whether to create a new invoice template from the template properties in this request, rather than link an existing one.; illustrative example: false
      --schedule-name string              scheduleName (body); Updated name.; maxLength: 255; illustrative example: "adl"
      --status string                     status (body); Updated lifecycle status. PATCH accepts only inactive, which stops future runs without changing endMonthYear (the deactivate endpoint also sets it); use the activate endpoint to reactivate a schedule.; enum: ["inactive"]; illustrative example: "inactive"
      --yes[=true|false]                  confirm the operation (required for destructive ops)
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

* [flexera-cli finops-billing invoice-schedules](flexera-cli_finops-billing_invoice-schedules.md)	 - Invoice Schedules operations (generated from the unified OpenAPI spec)

