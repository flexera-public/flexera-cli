## flexera-cli refresh-token list

Index a user's refresh tokens

```
flexera-cli refresh-token list [flags]
```

### Options

```
  -h, --help          help for list
      --org-id int    orgId (query)
      --user-id int   userId (query)
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
  -o, --output string           output format (json|table)
      --refresh-token string    OAuth refresh token
      --zone string             API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli refresh-token](flexera-cli_refresh-token.md)	 - Refresh Token operations (generated from the unified OpenAPI spec)

