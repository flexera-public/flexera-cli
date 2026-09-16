## flexera-cli bill-upload

BillUpload operations (generated from the unified OpenAPI spec)

### Options

```
  -h, --help   help for bill-upload
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

* [flexera-cli](flexera-cli.md)	 - Flexera One unified API command-line client
* [flexera-cli bill-upload create](flexera-cli_bill-upload_create.md)	 - POST /optima/orgs/{orgId}/billUploads
* [flexera-cli bill-upload delete](flexera-cli_bill-upload_delete.md)	 - DELETE /optima/orgs/{orgId}/billUploads/{billUploadId}
* [flexera-cli bill-upload files](flexera-cli_bill-upload_files.md)	 - POST /optima/orgs/{orgId}/billUploads/{billUploadId}/files/{fileId}
* [flexera-cli bill-upload get](flexera-cli_bill-upload_get.md)	 - GET /optima/orgs/{orgId}/billUploads/{billUploadId}
* [flexera-cli bill-upload list](flexera-cli_bill-upload_list.md)	 - GET /optima/orgs/{orgId}/billUploads
* [flexera-cli bill-upload operations](flexera-cli_bill-upload_operations.md)	 - POST /optima/orgs/{orgId}/billUploads/{billUploadId}/operations
* [flexera-cli bill-upload push](flexera-cli_bill-upload_push.md)	 - Create, upload, and commit a bill upload

