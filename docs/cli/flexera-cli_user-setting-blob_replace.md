## flexera-cli user-setting-blob replace

Generates a signed URL to store a user settings object at a specified key, which may represent a page ID or a combination of page ID and prefix ID, used for retrieving user-specific settings.

```
flexera-cli user-setting-blob replace [flags]
```

### Options

```
      --body string   raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run       print the planned operation as JSON and exit without calling the API
      --expiry int    expiry (body)
  -h, --help          help for replace
      --id string     id (body)
      --org-id int    orgId (body)
      --type string   type (body)
      --yes           confirm the operation (required for destructive ops)
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

* [flexera-cli user-setting-blob](flexera-cli_user-setting-blob.md)	 - User Setting Blob operations (generated from the unified OpenAPI spec)

