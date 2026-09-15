## flexera-cli bill-upload files

POST /optima/orgs/{orgId}/billUploads/{billUploadId}/files/{fileId}

```
flexera-cli bill-upload files [flags]
```

### Options

```
      --bill-upload-id string   billUploadId (path, required)
      --file-id string          fileId (path, required)
  -h, --help                    help for files
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

