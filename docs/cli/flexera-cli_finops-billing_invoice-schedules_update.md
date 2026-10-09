## flexera-cli finops-billing invoice-schedules update

Update an invoice schedule

### Synopsis

Update an invoice schedule

Partially updates an invoice schedule using JSON Merge Patch (RFC 7396) semantics: only fields present in the request body are changed and omitted fields remain unchanged. Supports optimistic locking through If-Match.

An inactive schedule cannot be modified through this endpoint. Use the activate endpoint to reactivate it.

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
      --body string                  raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run                      print the planned operation as JSON and exit without calling the API
  -h, --help                         help for update
      --id string                    id (path, required); Identifier of the invoice schedule; required by API; illustrative example: "sch_1"
  -i, --interactive                  edit inputs in a terminal form, review a plan and approve with typed yes
      --invoice-template-id string   invoiceTemplateId (body); Existing invoice template to point the schedule at. Send isCreateTemplate=false with this field omitted to detach the schedule's template.; illustrative example: "tmpl_9002"
      --is-create-template           isCreateTemplate (body); Whether to create a new invoice template from the template properties in this request, rather than link an existing one.; illustrative example: false
      --schedule-name string         scheduleName (body); Updated name.; maxLength: 255; illustrative example: "p4y"
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

* [flexera-cli finops-billing invoice-schedules](flexera-cli_finops-billing_invoice-schedules.md)	 - Invoice Schedules operations (generated from the unified OpenAPI spec)

