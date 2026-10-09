## flexera-cli finops-customizations saved-filters list

Index saved filters

### Synopsis

Index saved filters

Lists all saved filters in the organization. Returns the caller's private filters and all organization-shared filters (visibility = "shared"). Billing Center entitlement is applied to every item: Billing Center group-by dimensions are omitted when the caller has no accessible Billing Center, and filterExpressionCount and groupByExpressionCount count only what the caller is entitled to see. Nothing in the response indicates that values were omitted. The request fails if the caller's Billing Center entitlement cannot be verified.

```
flexera-cli finops-customizations saved-filters list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-customizations saved-filters list --org-id ORG_ID
```

### Options

```
      --favorite-ids string   favoriteIds (query); Comma-separated saved filter IDs marked as favorites, max 50.; illustrative example: "Qui aut quasi nihil qui consequuntur et."
  -h, --help                  help for list
      --limit int             limit (query); Page size (default 10, max 200); minimum: 1; maximum: 200; API default: 10; illustrative example: 47
      --no-paginate           return only the first page (do not follow nextPage)
      --order-by string       orderBy (query); Optional orderBy query allows to specify an expression for determining what values are used to order the entities. Multiple expressions can be specified using comma separated values.; illustrative example: "visibility desc,name asc"
      --skip-token string     resume pagination from this token; An opaque token to be provided when requesting a subsequent page after receiving a partial response. Partial responses will include a "nextPage" attribute in their response body, which contains the URL of the next page including the appropriate skipToken.
      --visibility string     visibility (query); Optional visibility filter for results.; enum: ["shared","private"]; illustrative example: "shared"
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

* [flexera-cli finops-customizations saved-filters](flexera-cli_finops-customizations_saved-filters.md)	 - Saved Filters operations (generated from the unified OpenAPI spec)

