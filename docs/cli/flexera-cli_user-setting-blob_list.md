## flexera-cli user-setting-blob list

Generates a signed URL to retrieve a user settings object from a specified key, which may represent a page ID, a combination of page ID and prefix ID, or another identifier for accessing user-specific settings

```
flexera-cli user-setting-blob list [flags]
```

### Options

```
      --expiry int    expiry (query)
  -h, --help          help for list
      --id string     id (query)
      --org-id int    orgId (query)
      --type string   type (query)
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

