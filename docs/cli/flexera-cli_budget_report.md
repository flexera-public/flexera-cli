## flexera-cli budget report

Provides a report comparing budget with actual spend

```
flexera-cli budget report [flags]
```

### Options

```
      --dimensions strings   dimensions (query)
      --end-at string        endAt (query)
      --filter string        filter (query)
  -h, --help                 help for report
      --id string            id (path, required)
      --start-at string      startAt (query)
      --summarized           summarized (query)
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

* [flexera-cli budget](flexera-cli_budget.md)	 - Budget operations (generated from the unified OpenAPI spec)

