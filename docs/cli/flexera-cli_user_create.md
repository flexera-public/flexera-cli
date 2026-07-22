## flexera-cli user create

Create and affiliate a user to an org

```
flexera-cli user create [flags]
```

### Options

```
      --body string         raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run             print the planned operation as JSON and exit without calling the API
      --email string        email (body)
      --first-name string   first_name (body)
  -h, --help                help for create
      --last-name string    last_name (body)
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

* [flexera-cli user](flexera-cli_user.md)	 - User operations (generated from the unified OpenAPI spec)

