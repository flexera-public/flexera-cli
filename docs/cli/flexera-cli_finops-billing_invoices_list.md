## flexera-cli finops-billing invoices list

Index invoices

### Synopsis

Index invoices

Lists invoice export lifecycle records in the organization.

```
flexera-cli finops-billing invoices list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing invoices list --org-id ORG_ID
```

### Options

```
      --bill-month string          billMonth (query); Optional bill month filter.; pattern: "^\\d{4}-\\d{2}$"; illustrative example: "Cruz Fadel"
      --customer-name string       customerName (query); Optional customer name filter.; illustrative example: "Officia sed."
      --export-type string         exportType (query); Optional export type filter.; enum: ["pdf","csv"]; illustrative example: "csv"
      --expression-filter string   expressionFilter (query); OData-style expression evaluated against each returned resource. Attribute names are case-sensitive and operators are case-insensitive. String eq, ne, and in comparisons are case-sensitive; co performs case-insensitive substring matching. Date values use YYYY-MM-DD; date-time values use ISO-8601 and are normalized to UTC; billMonth uses YYYY-MM. Missing fiel... (see cli schema); illustrative example: "customerName co 'Acme' and exportType eq 'pdf'"
      --filter string              filter (query); Legacy substring search of the stored invoice generation filter JSON. This parameter does not accept OData expressions.; illustrative example: "Facere dolorum est velit adipisci tenetur."
      --group-by-count int         groupByCount (query); Optional group-by dimension count filter.; illustrative example: 17118634271918162000
  -h, --help                       help for list
      --limit int                  limit (query); Page size (default 10, max 200); minimum: 1; maximum: 200; API default: 10; illustrative example: 7
      --no-paginate                return only the first page (do not follow nextPage)
      --order-by string            orderBy (query); Optional orderBy query allows to specify an expression for determining what values are used to order the entities. Multiple expressions can be specified using comma separated values.; illustrative example: "updatedAt desc"
      --skip-token string          resume pagination from this token; An opaque token to be provided when requesting a subsequent page after receiving a partial response. Partial responses will include a "nextPage" attribute in their response body, which contains the URL of the next page including the appropriate skipToken.
      --status string              status (query); Optional status filter.; illustrative example: "Est ullam."
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

* [flexera-cli finops-billing invoices](flexera-cli_finops-billing_invoices.md)	 - Invoices operations (generated from the unified OpenAPI spec)

