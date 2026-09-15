## flexera-cli user-memberships privileges-report

Show a user's privileges report

```
flexera-cli user-memberships privileges-report [flags]
```

### Options

```
  -h, --help                 help for privileges-report
      --id int               id (path, required)
      --prefix string        prefix (query)
      --scope-refs strings   scopeRefs (query)
      --view string          view (query)
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

* [flexera-cli user-memberships](flexera-cli_user-memberships.md)	 - User Memberships operations (generated from the unified OpenAPI spec)

