## flexera-cli billing rules-rules

Create an enterprise adjustment rule

```
flexera-cli billing rules-rules [flags]
```

### Options

```
      --body string        raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run            print the planned operation as JSON and exit without calling the API
      --enabled            enabled (body)
      --end-after string   endAfter (body)
  -h, --help               help for rules-rules
      --name string        name (body)
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
      --login-base-url string   override login base URL
      --org-id int              organization ID
  -o, --output string           output format (json|table)
      --refresh-token string    OAuth refresh token
      --zone string             API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli billing](flexera-cli_billing.md)	 - Billing operations (generated from the unified OpenAPI spec)

