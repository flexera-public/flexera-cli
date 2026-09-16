## flexera-cli bill-upload push

Create, upload, and commit a bill upload

```
flexera-cli bill-upload push [flags]
```

### Options

```
      --bill-connect-id string   bill connect ID (required)
      --billing-period string    billing period in yyyy-mm format (required)
      --file strings             local bill file to upload (repeatable, required)
  -h, --help                     help for push
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

