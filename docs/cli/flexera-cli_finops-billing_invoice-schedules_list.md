## flexera-cli finops-billing invoice-schedules list

Index invoice schedules

### Synopsis

Index invoice schedules

Lists invoice schedules in the organization.

```
flexera-cli finops-billing invoice-schedules list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing invoice-schedules list --org-id ORG_ID
```

### Options

```
      --customer-id int        customerId (query); Optional filter: only schedules covering this customer.; illustrative example: 53125
      --filter string          filter (query); OData-style expression evaluated against each returned resource. Attribute names are case-sensitive and operators are case-insensitive. String eq, ne, and in comparisons are case-sensitive; co performs case-insensitive substring matching. Date values use YYYY-MM-DD; date-time values use ISO-8601 and are normalized to UTC; billMonth uses YYYY-MM. Missing fiel... (see cli schema); illustrative example: "status eq 'active' and customerId eq 53125"
  -h, --help                   help for list
      --limit int              limit (query); Page size (default 50, max 200); minimum: 1; maximum: 200; API default: 50; illustrative example: 7
      --no-paginate            return only the first page (do not follow nextPage)
      --order-by string        orderBy (query); Optional orderBy query allows to specify an expression for determining what values are used to order the entities. Multiple expressions can be specified using comma separated values.; illustrative example: "updatedAt desc"
      --schedule-name string   scheduleName (query); Optional exact-match filter on schedule name.; illustrative example: "Ut sint."
      --skip-token string      resume pagination from this token; An opaque token to be provided when requesting a subsequent page after receiving a partial response. Partial responses will include a "nextPage" attribute in their response body, which contains the URL of the next page including the appropriate skipToken.
      --status string          status (query); Optional status filter.; enum: ["active","inactive"]; illustrative example: "active"
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

* [flexera-cli finops-billing invoice-schedules](flexera-cli_finops-billing_invoice-schedules.md)	 - Invoice Schedules operations (generated from the unified OpenAPI spec)

