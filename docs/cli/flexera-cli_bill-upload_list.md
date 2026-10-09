## flexera-cli bill-upload list

GET /optima/orgs/{orgId}/billUploads

### Synopsis

GET /optima/orgs/{orgId}/billUploads

Lists existing bill uploads. You can filter by bill upload and/or billing period, too..

**Required security scopes for JWTAuth**:
  * `optima:bill_upload:index+optima:bill_connect:index+common:org:own`

```
flexera-cli bill-upload list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli bill-upload list --org-id ORG_ID
```

### Options

```
      --bill-connect-id string   billConnectId (query); Optional filter by bill connect ID
      --billing-period string    billingPeriod (query); Optional filter by billing period; pattern: "^([0-9]{4})-([0-9]{2})$"
  -h, --help                     help for list
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

