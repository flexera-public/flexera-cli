## flexera-cli finops-customizations report-subscriptions update

Update a report subscription

### Synopsis

Update a report subscription

Partially updates a report subscription (owner only) using JSON Merge Patch (RFC 7396) semantics: only fields present in the request body are changed and omitted fields remain unchanged. The subscription id is taken from the path and cannot be modified. Supports optimistic locking through If-Match.

```
flexera-cli finops-customizations report-subscriptions update [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-customizations report-subscriptions update --org-id ORG_ID --id ID --body @request.json
  flexera-cli finops-customizations report-subscriptions update --org-id ORG_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-customizations report-subscriptions update --example > request.json
```

### Options

```
      --body string         raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run             print the planned operation as JSON and exit without calling the API
      --enabled             enabled (body); Updated active state.; illustrative example: true
  -h, --help                help for update
      --id string           id (path, required); Identifier of the report subscription (UUID); required by API; format: uuid; illustrative example: "550e8400-e29b-41d4-a716-446655440000"
  -i, --interactive         edit inputs in a terminal form, review a plan and approve with typed yes
      --name string         name (body); Updated name.; maxLength: 255; illustrative example: "Monthly Cloud Cost"
      --visibility string   visibility (body); Updated visibility scope.; enum: ["private","shared"]; illustrative example: "private"
      --yes                 confirm the operation (required for destructive ops)
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

* [flexera-cli finops-customizations report-subscriptions](flexera-cli_finops-customizations_report-subscriptions.md)	 - Report Subscriptions operations (generated from the unified OpenAPI spec)

