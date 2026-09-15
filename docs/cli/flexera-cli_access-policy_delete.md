## flexera-cli access-policy delete

delete Access Policy

```
flexera-cli access-policy delete [flags]
```

### Options

```
      --access-policy-id string   accessPolicyId (path, required)
      --dry-run                   print the planned operation as JSON and exit without calling the API
  -h, --help                      help for delete
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
      --login-base-url string   override login base URL
      --org-id int              organization ID
  -o, --output string           output format (json|table)
      --refresh-token string    OAuth refresh token
      --zone string             API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli access-policy](flexera-cli_access-policy.md)	 - Access Policy operations (generated from the unified OpenAPI spec)

