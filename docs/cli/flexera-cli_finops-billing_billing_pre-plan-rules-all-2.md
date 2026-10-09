## flexera-cli finops-billing billing pre-plan-rules-all-2

Remove an adjustment rule that runs before all plans

### Synopsis

Remove an adjustment rule that runs before all plans

Remove an adjustment rule.

```
flexera-cli finops-billing billing pre-plan-rules-all-2 [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing billing pre-plan-rules-all-2 --org-id ORG_ID --id ID
  flexera-cli finops-billing billing pre-plan-rules-all-2 --org-id ORG_ID --id ID --dry-run
```

### Options

```
      --dry-run     print the planned operation as JSON and exit without calling the API
  -h, --help        help for pre-plan-rules-all-2
      --id string   id (path, required); id identifies this rule; required by API; format: uuid; illustrative example: "ae85f96e-6b9e-4d68-a741-d1ea7ca1fb28"
      --yes         confirm the operation (required for destructive ops)
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

* [flexera-cli finops-billing billing](flexera-cli_finops-billing_billing.md)	 - Billing operations (generated from the unified OpenAPI spec)

