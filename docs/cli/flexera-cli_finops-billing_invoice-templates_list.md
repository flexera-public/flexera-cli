## flexera-cli finops-billing invoice-templates list

Index invoice templates

### Synopsis

Index invoice templates

Lists invoice templates in the organization, newest first.

```
flexera-cli finops-billing invoice-templates list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing invoice-templates list --org-id ORG_ID
```

### Options

```
      --filter string       filter (query); OData-style filter expression used to filter invoice templates.; illustrative example: "templateName co 'summary'"
  -h, --help                help for list
      --limit int           limit (query); Page size (default 50, max 200); minimum: 1; maximum: 200; API default: 50; illustrative example: 189
      --no-paginate         return only the first page (do not follow nextPage)
      --order-by string     orderBy (query); Optional orderBy query allows to specify an expression for determining what values are used to order the entities. Multiple expressions can be specified using comma separated values.; illustrative example: "updatedAt desc"
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

* [flexera-cli finops-billing invoice-templates](flexera-cli_finops-billing_invoice-templates.md)	 - Invoice Templates operations (generated from the unified OpenAPI spec)

