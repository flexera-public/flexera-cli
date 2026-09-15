## flexera-cli access-rules list-groups

List group's roles in every BillingCenter.

```
flexera-cli access-rules list-groups [flags]
```

### Options

```
      --group int   group (path, required)
  -h, --help        help for list-groups
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

* [flexera-cli access-rules](flexera-cli_access-rules.md)	 - AccessRules operations (generated from the unified OpenAPI spec)

