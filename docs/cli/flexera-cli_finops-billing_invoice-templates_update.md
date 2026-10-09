## flexera-cli finops-billing invoice-templates update

Update an invoice template

### Synopsis

Update an invoice template

Partially updates an invoice template using JSON Merge Patch (RFC 7396) semantics: only fields present in the request body are changed and omitted fields remain unchanged. The template id is taken from the path and cannot be modified. Supports optimistic locking through If-Match.

```
flexera-cli finops-billing invoice-templates update [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing invoice-templates update --org-id ORG_ID --id ID --body @request.json
  flexera-cli finops-billing invoice-templates update --org-id ORG_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-billing invoice-templates update --example > request.json
```

### Options

```
      --body string               raw JSON body (inline | @file | @-); overrides body field flags
      --cost-metric string        costMetric (body); Updated cost metric.; enum: ["billedCost","effectiveCost","modifiedBilledCost","modifiedEffectiveCost"]; illustrative example: "effectiveCost"
      --dimensions strings        dimensions (body); Updated grouping dimensions.; maxItems: 10; CLI: comma-separated values or repeated flag; illustrative example: ["Temporibus nihil minima non asperiores quisquam.","Qui animi temporibus non excepturi.","Quo praesentium ea vitae molestias dolores et."]
      --display-logo-within-pdf   displayLogoWithinPdf (body); Updated PDF logo display flag.; illustrative example: true
      --dry-run                   print the planned operation as JSON and exit without calling the API
      --export-type string        exportType (body); Updated export type.; enum: ["pdf","csv"]; illustrative example: "pdf"
      --free-text string          freeText (body); Updated free-text company/address block.; illustrative example: "Ut libero animi porro illo."
  -h, --help                      help for update
      --id string                 id (path, required); Invoice template identifier; required by API; illustrative example: "tmpl_9001"
  -i, --interactive               edit inputs in a terminal form, review a plan and approve with typed yes
      --logo string               logo (body); Updated logo payload.; illustrative example: "Pariatur illum quibusdam quis."
      --period-type string        periodType (body); Updated period type.; enum: ["chargePeriod","billPeriod","billingPeriod"]; illustrative example: "billPeriod"
      --template-name string      templateName (body); Updated template name.; maxLength: 255; illustrative example: "561"
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

