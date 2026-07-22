## flexera-cli onboarding delete

Onboarding: Delete

```
flexera-cli onboarding delete [flags]
```

### Options

```
      --connector-id string   connector_id (path, required)
      --dry-run               print the planned operation as JSON and exit without calling the API
  -h, --help                  help for delete
      --provider string       provider (query)
      --yes                   confirm the operation (required for destructive ops)
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

* [flexera-cli onboarding](flexera-cli_onboarding.md)	 - Onboarding operations (generated from the unified OpenAPI spec)

