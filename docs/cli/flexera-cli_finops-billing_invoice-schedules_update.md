## flexera-cli finops-billing invoice-schedules update

Update an invoice schedule

### Synopsis

Update an invoice schedule

Partially updates an invoice schedule using JSON Merge Patch (RFC 7396) semantics: only fields present in the request body are changed and omitted fields remain unchanged. Supports optimistic locking through If-Match.

Invoice numbering: invoiceNumberPrefix and invoiceNumberBase follow the same merge-patch rule, so omitting a field leaves it unchanged. Sending "" for both clears numbering and disables it. Sending "" for invoiceNumberPrefix alone, with an existing base, removes the prefix only. Re-adding numbering later continues from the last used number + 1. An invoiceNumberPrefix without a base is rejected with 400 invoice_number_prefix_requires_base, and changing the base after the first number was allocated is rejected with 400 invoice_number_base_locked.

An inactive schedule accepts exactly three fields through this endpoint: status, invoiceNumberPrefix and invoiceNumberBase. Any other field is rejected. Use the activate endpoint to reactivate it.

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
      --body string                    raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run                        print the planned operation as JSON and exit without calling the API
  -h, --help                           help for update
      --id string                      id (path, required); Identifier of the invoice schedule; required by API; illustrative example: "sch_1"
  -i, --interactive                    edit inputs in a terminal form, review a plan and approve with typed yes
      --invoice-number-base string     invoiceNumberBase (body); Optional updated first sequence number for the invoice numbers this schedule generates.; maxLength: 9; pattern: "^$|^0*[1-9][0-9]*$"; illustrative example: "n3b"
      --invoice-number-prefix string   invoiceNumberPrefix (body); Optional updated prefix for the invoice numbers this schedule generates. Requires invoiceNumberBase.; maxLength: 20; pattern: "^[A-Za-z0-9_-]*$"; illustrative example: "xkx"
      --invoice-template-id string     invoiceTemplateId (body); Existing invoice template to point the schedule at. Send isCreateTemplate=false with this field omitted to detach the schedule's template.; illustrative example: "tmpl_9002"
      --is-create-template             isCreateTemplate (body); Whether to create a new invoice template from the template properties in this request, rather than link an existing one.; illustrative example: false
      --schedule-name string           scheduleName (body); Updated name.; maxLength: 255; illustrative example: "dq5"
      --status string                  status (body); Updated lifecycle status. PATCH accepts only inactive; use the activate endpoint to reactivate a schedule.; enum: ["inactive"]; illustrative example: "inactive"
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

