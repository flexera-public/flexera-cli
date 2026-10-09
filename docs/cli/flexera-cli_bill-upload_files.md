## flexera-cli bill-upload files

POST /optima/orgs/{orgId}/billUploads/{billUploadId}/files/{fileId}

### Synopsis

POST /optima/orgs/{orgId}/billUploads/{billUploadId}/files/{fileId}

Uploads a file, adding it to the given bill upload.
Note, by default, you can only upload 500 files max per bill upload, and each must be smaller than 1000 MB.
If the upload is successful, a MD5 hash is returned, so that you can verify its integrity.


**Required security scopes for JWTAuth**:
  * `optima:bill_upload:create+optima:bill_connect:create+common:org:own`

Request body: The raw content of the file to upload

```
flexera-cli bill-upload files [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli bill-upload files --org-id ORG_ID --bill-upload-id BILL_UPLOAD_ID --file-id FILE_ID --body @request.json
  flexera-cli bill-upload files --org-id ORG_ID --bill-upload-id BILL_UPLOAD_ID --file-id FILE_ID --body @request.json --dry-run
```

### Options

```
      --bill-upload-id string   billUploadId (path, required); The identifier of the bill upload; required by API; format: uuid
      --body string             raw request body (@file | @-)
      --dry-run                 print the planned operation as JSON and exit without calling the API
      --file-id string          fileId (path, required); The basename of the uploaded file (100 characters max), supported extensions: .csv, .json, .jsonl, gzipped or not; required by API
  -h, --help                    help for files
      --yes                     confirm the operation (required for destructive ops)
```

### Options inherited from parent commands

```
      --access-token string     static bearer access token
      --api-base-url string     override API base URL
      --client-id string        OAuth client ID
      --client-secret string    OAuth client secret
      --config string           config file (default $HOME/.flexera/config.yaml)
  -d, --debug                   log HTTP requests/responses to stderr (Authorization redacted)
      --json-style string       JSON whitespace style (auto|pretty|compact) (default "auto")
      --login-base-url string   override login base URL
      --no-validate             skip API schema constraints (never JSON syntax or request data-loss checks)
      --org-id int              organization ID
      --out-fields string       project JSON output fields (comma-separated paths)
      --out-jq string           shape JSON output with a jq expression
  -o, --output string           output format (json|table)
  -r, --raw-output              write jq string results without JSON quotes (requires --out-jq)
      --refresh-token string    OAuth refresh token
      --zone string             API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli bill-upload](flexera-cli_bill-upload.md)	 - Bill Upload API

