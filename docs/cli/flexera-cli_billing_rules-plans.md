## flexera-cli billing rules-plans

Create an adjustment plan rule

```
flexera-cli billing rules-plans [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli billing rules-plans --org-id ORG_ID --plan-id PLAN_ID --body @request.json
  flexera-cli billing rules-plans --org-id ORG_ID --plan-id PLAN_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema billing rules-plans --example > request.json
```

### Options

```
      --body string        raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run            print the planned operation as JSON and exit without calling the API
      --enabled            enabled (body)
      --end-after string   endAfter (body)
  -h, --help               help for rules-plans
  -i, --interactive        edit inputs in a terminal form, review a plan and approve with typed yes
      --name string        name (body)
      --plan-id string     planId (path, required)
      --start-on string    startOn (body)
      --type string        type (body)
      --yes                confirm the operation (required for destructive ops)
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

* [flexera-cli billing](flexera-cli_billing.md)	 - Billing operations (generated from the unified OpenAPI spec)

