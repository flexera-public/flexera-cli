## flexera-cli license fees-all-2

Delete license agreement fee

```
flexera-cli license fees-all-2 [flags]
```

### Options

```
      --dry-run             print the planned operation as JSON and exit without calling the API
      --fee-id string       feeId (path, required)
  -h, --help                help for fees-all-2
      --license-id string   licenseId (path, required)
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

* [flexera-cli license](flexera-cli_license.md)	 - License operations (generated from the unified OpenAPI spec)

