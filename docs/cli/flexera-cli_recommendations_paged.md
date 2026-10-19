## flexera-cli recommendations paged

List recommendations (filtered, sorted, cursor-paginated)

### Synopsis

List recommendations (filtered, sorted, cursor-paginated)

List recommendations, filtered and sorted via a generic filter/orderBy expression, cursor-paginated.
billingCenterID and status scoping are done via `filter` (e.g. `billingCenterID in ['bc1'] and status in ['active']`), not separate discrete parameters, matching explorer_service's own contract for this endpoint.
User must have the 'optima:recommendation:index' privilege on every billing center the filter's billingCenterID clause(s) select with eq/in. If it selects none (no billingCenterID clause, or only exclusions such as ne/not), the request covers every billing center the user can access: the whole org for a user with org-level access, otherwise the billing centers they have been granted (403 if none). Results are always confined to the billing centers the call was authorized for, whatever else the filter says.

**Required security scopes for GlobalSession**:
  * `common:org:affiliated+common:org:own`

```
flexera-cli recommendations paged [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli recommendations paged --org-id ORG_ID
```

### Options

```
      --filter string       filter (query); Optional filter for returning recommendations matching specific criteria. ### Supported Filter Keys | Filter | Type | Example | Description | | --------------------- | ------ | ------------------------------------------------- | ------------------------------------------------------------------------------------ | | billingCenterID | string | billingCenterID... (see cli schema)
  -h, --help                help for paged
      --limit int           limit (query); Custom pagination limit to be used.
      --no-paginate         return only the first page (do not follow nextPage)
      --order-by string     orderBy (query); Optional order by for returning recommendations in a specific order. ### Supported Order by Keys | Key | Description | | -------------------- | ------------------------------------------------------------------------------- | | recommendation | Returns recommendations ordered by 'recommendation' property | | type | Returns recommendations ordered by 'type' p... (see cli schema)
      --skip-token string   resume pagination from this token; An opaque token to be provided when requesting a subsequent page after receiving a partial response. Partial responses will include a "nextPage" attribute in their response body, which contains the URL of the next page including the appropriate skipToken.
      --view details        view (query); An optional parameter that controls the level of detail returned about policy violation data. Depending on it the data is returned in the details or `detailsExtended` field; enum: ["default","extended"]; API default: "default"
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

* [flexera-cli recommendations](flexera-cli_recommendations.md)	 - Optima Recommendations API

