## flexera-cli billing rules-plans

Create an adjustment plan rule

```
flexera-cli billing rules-plans [flags]
```

### Options

```
      --body string        raw JSON body (inline | @file | @-); overrides body field flags
      --enabled            enabled (body)
      --end-after string   endAfter (body)
  -h, --help               help for rules-plans
      --name string        name (body)
      --plan-id string     planId (path, required)
      --start-on string    startOn (body)
      --type string        type (body)
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

