## flexera-cli compliance standard

Compliance Standards.

```
flexera-cli compliance standard [flags]
```

### Options

```
      --accounts strings    accounts (body)
      --body string         raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run             print the planned operation as JSON and exit without calling the API
      --etime string        etime (body)
  -h, --help                help for standard
      --providers strings   providers (body)
      --regions strings     regions (body)
      --services strings    services (body)
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
      --login-base-url string   override login base URL
      --org-id int              organization ID
  -o, --output string           output format (json|table)
      --refresh-token string    OAuth refresh token
      --zone string             API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli compliance](flexera-cli_compliance.md)	 - compliance operations (generated from the unified OpenAPI spec)

