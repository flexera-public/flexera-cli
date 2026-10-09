## flexera-cli saas license fees-all-4

List license agreement fees

### Synopsis

List license agreement fees

Retrieves a collection of license agreement fees

```
flexera-cli saas license fees-all-4 [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas license fees-all-4 --org-id ORG_ID --license-id LICENSE_ID
```

### Options

```
      --filter string       filter (query); The date filter to query the licenses. Supported fields in the filter are [createdAt]; illustrative example: "createdAt gt 'ISO 8601 date'"
  -h, --help                help for fees-all-4
      --license-id string   licenseId (path, required); Identifies a license by given license ID.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "38932"
      --no-paginate         return only the first page (do not follow nextPage)
      --order-by string     orderBy (query); The order by filter to sort the licenses. Supported values are Supported fields in the orderBy are [createdAt]; API default: "createdAt desc"; illustrative example: "createdAt desc"
      --skip-token string   resume pagination from this token; An opaque token to be provided when requesting a subsequent page after receiving a partial response. Partial responses will include a "nextPage" attribute in their response body, which contains the URL of the next page including the appropriate skipToken.
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

* [flexera-cli saas license](flexera-cli_saas_license.md)	 - License operations (generated from the unified OpenAPI spec)

