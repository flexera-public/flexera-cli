## flexera-cli finops-billing shared-cost-rules update

Update a shared cost rule

### Synopsis

Update a shared cost rule

Updates an existing shared cost rule. All fields in the request body are optional; only supplied fields are changed. The rule name, if provided, must remain unique within the organization. A dynamicProportionateSplit rule takes new telemetry through a save that names a staged upload in telemetryUploadId, which may be the only field; the upload's applied.monthsTouched then lists the months to reprocess.

```
flexera-cli finops-billing shared-cost-rules update [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing shared-cost-rules update --org-id ORG_ID --id ID --body @request.json
  flexera-cli finops-billing shared-cost-rules update --org-id ORG_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-billing shared-cost-rules update --example > request.json
```

### Options

```
      --body string                  raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run[=true|false]         print the planned operation as JSON and exit without calling the API
      --effective-from string        effectiveFrom (body); Updated month (inclusive) from which the rule applies, formatted YYYY-MM. This is applicable on ChargePeriod only.; pattern: "^20[\\d]{2}-((0[1-9])|(1[012]))$"; illustrative example: "2025-07"
      --effective-to string          effectiveTo (body); Updated month (exclusive) before which the rule applies, formatted YYYY-MM. Must be > effectiveFrom. Omit this field for no end date; the rule applies indefinitely. This is applicable on ChargePeriod only.; pattern: "^20[\\d]{2}-((0[1-9])|(1[012]))$"; illustrative example: "2025-12"
  -h, --help[=true|false]            help for update
      --id string                    id (path, required); Unique identifier of the shared cost rule; required by API; illustrative example: "68f2a91c4b3d2e1f5a6b7c8d"
  -i, --interactive[=true|false]     edit inputs in a terminal form, review a plan and approve with typed yes
      --name string                  name (body); Updated display name for the shared cost rule. Must not contain "(", ")" or ">".; minLength: 1; maxLength: 255; pattern: "^[^()\u003e]+$"; illustrative example: "Equal split for AWS between Teams A, B and C - revised"
      --status string                status (body); Updated lifecycle status.; enum: ["active"]; illustrative example: "active"
      --telemetry-upload-id string   telemetryUploadId (body); A valid staged telemetry upload for this rule, applied in the same save. Required when the save changes the type to dynamicProportionateSplit or changes a dynamic rule's destination dimension; not allowed when the rule's type after the save is not dynamicProportionateSplit.; illustrative example: "tu_01JAB3K9Q2X7V5R8M4N6P0T1W3"
      --yes[=true|false]             confirm the operation (required for destructive ops)
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

* [flexera-cli finops-billing shared-cost-rules](flexera-cli_finops-billing_shared-cost-rules.md)	 - Shared Cost Rules operations (generated from the unified OpenAPI spec)

