## flexera-cli recommendations groups

Group recommendations into a tree

### Synopsis

Group recommendations into a tree

Groups recommendations into a tree, one level per `groupBy` field, with the number of recommendations and the total savings of every group.
The recommendations are selected by `filter`, in the same way as for the paged list, and the same billing center rules apply: User must have the 'optima:recommendation:index' privilege on every billing center the filter's billingCenterID clause(s) select with eq/in. If it selects none, the request covers every billing center the user can access: the whole org for a user with org-level access, otherwise the billing centers they have been granted (403 if none). Results are always confined to the billing centers the call was authorized for, whatever else the filter says.
When `groupBy` includes billingCenterID, each recommendation is grouped under the authorized billing center it belongs to.

**Required security scopes for GlobalSession**:
  * `common:org:affiliated+common:org:own`

```
flexera-cli recommendations groups [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli recommendations groups --org-id ORG_ID --group-by GROUP_BY
```

### Options

```
      --filter string     filter (query); Optional filter for returning recommendations matching specific criteria. ### Supported Filter Keys | Filter | Type | Example | Description | | --------------------- | ------ | ------------------------------------------------- | ------------------------------------------------------------------------------------ | | billingCenterID | string | billingCenterID... (see cli schema)
      --group-by string   groupBy (query); Comma separated fields to group the recommendations by. Each field adds one level to the tree, in the given order. ### Supported Group by Fields | Field | Description | | --------------- | --------------------------------------------------------------------------------------- | | billingCenterID | Groups by the authorized billing center the recommendation be... (see cli schema); required by API
  -h, --help              help for groups
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

