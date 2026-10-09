## flexera-cli bill-upload create

POST /optima/orgs/{orgId}/billUploads

### Synopsis

POST /optima/orgs/{orgId}/billUploads

Creates a new bill upload, for a given [bill connect](https://reference.rightscale.com/optima-bill/#/CBIBillConnects) and billing period (month, yyyy-mm). Committing and processing a bill upload replaces that month of data for the bill connect/org.

**Required security scopes for JWTAuth**:
  * `optima:bill_upload:create+optima:bill_connect:create+common:org:own`

```
flexera-cli bill-upload create [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli bill-upload create --org-id ORG_ID --body @request.json
  flexera-cli bill-upload create --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema bill-upload create --example > request.json
```

### Options

```
      --bill-connect-id string   billConnectId (body); required by API; The bill connect to use for this bill upload; illustrative example: "cbi-integration1-1234567"
      --billing-period string    billingPeriod (body); required by API; The billing period covered by this bill upload, formatted as yyyy-mm; pattern: "^([0-9]{4})-([0-9]{2})$"; illustrative example: "2020-03"
      --body string              raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run                  print the planned operation as JSON and exit without calling the API
  -h, --help                     help for create
  -i, --interactive              edit inputs in a terminal form, review a plan and approve with typed yes
      --yes                      confirm the operation (required for destructive ops)
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

