## flexera-cli customer-group-type summary

Retrieve summary of customer groups for a given group type.

```
flexera-cli customer-group-type summary [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli customer-group-type summary --org-id ORG_ID --customer-group-type-id CUSTOMER_GROUP_TYPE_ID
```

### Options

```
      --customer-group-type-id string   customerGroupTypeId (path, required)
      --filter string                   filter (query)
  -h, --help                            help for summary
      --limit int                       limit (query)
      --no-paginate                     return only the first page (do not follow nextPage)
      --order-by string                 orderBy (query)
      --skip-token string               resume pagination from this token
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

* [flexera-cli customer-group-type](flexera-cli_customer-group-type.md)	 - Customer Group Type operations (generated from the unified OpenAPI spec)

