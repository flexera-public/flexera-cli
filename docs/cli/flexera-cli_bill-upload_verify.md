## flexera-cli bill-upload verify

Verify local CBI bill-upload CSV file(s) before uploading

```
flexera-cli bill-upload verify [flags]
```

### Options

```
      --file strings   local CSV file to verify (repeatable, required)
  -h, --help           help for verify
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

* [flexera-cli bill-upload](flexera-cli_bill-upload.md)	 - BillUpload operations (generated from the unified OpenAPI spec)

