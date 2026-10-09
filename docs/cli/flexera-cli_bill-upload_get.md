## flexera-cli bill-upload get

GET /optima/orgs/{orgId}/billUploads/{billUploadId}

### Synopsis

GET /optima/orgs/{orgId}/billUploads/{billUploadId}

Shows the details of a bill upload, including the list of successfully uploaded files so far.

**Required security scopes for JWTAuth**:
  * `optima:bill_upload:show+optima:bill_connect:show+common:org:own`

```
flexera-cli bill-upload get [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli bill-upload get --org-id ORG_ID --bill-upload-id BILL_UPLOAD_ID
```

### Options

```
      --bill-upload-id string   billUploadId (path, required); The identifier of the bill upload; required by API; format: uuid
  -h, --help                    help for get
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

