## flexera-cli saas customer-group-type list

List customer group types

### Synopsis

List customer group types

Retrieves a collection of customer group types.

```
flexera-cli saas customer-group-type list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas customer-group-type list --org-id ORG_ID
```

### Options

```
      --filter string       filter (query); The filter to query the customer group types. Supported fields in the filter are [name,displayName,applicationId]; illustrative example: "displayName eq 'Department'"
  -h, --help                help for list
      --no-paginate         return only the first page (do not follow nextPage)
      --order-by string     orderBy (query); The order by filter to sort the customer group types. Supported fields in the orderBy are [createdAt]; API default: "createdAt desc"; illustrative example: "createdAt desc"
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

* [flexera-cli saas customer-group-type](flexera-cli_saas_customer-group-type.md)	 - Customer Group Type operations (generated from the unified OpenAPI spec)

