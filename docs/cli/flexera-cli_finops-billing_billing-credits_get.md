## flexera-cli finops-billing billing-credits get

Show a credit assignment

### Synopsis

Show a credit assignment

Returns a single credit assignment by ID.

```
flexera-cli finops-billing billing-credits get [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing billing-credits get --org-id ORG_ID --id ID
```

### Options

```
  -h, --help        help for get
      --id string   id (path, required); Unique identifier of the credit assignment.; required by API; format: uuid; illustrative example: "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
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

* [flexera-cli finops-billing billing-credits](flexera-cli_finops-billing_billing-credits.md)	 - Billing Credits operations (generated from the unified OpenAPI spec)

