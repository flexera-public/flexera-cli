## flexera-cli policy unmanaged-applied-policies list

Index unmanaged applied policies

### Synopsis

Index unmanaged applied policies

Retrieves applied policies that are not part of a policy aggregate or meta child policies.
        Only applied policies are applied from the project-scoped Template endpoint (as opposed to the org-wide Catalog) are part of this view.
        These template-based policies should largely be only used for development and testing purposes.
        Meta Child Policies (applied policies created by other policies passing their id as a reference) are also not included.

```
flexera-cli policy unmanaged-applied-policies list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli policy unmanaged-applied-policies list --org-id ORG_ID
```

### Options

```
      --filter string       filter (query); Optional filter to retrieve unmanaged policies based on specific criteria. ### Supported Filter Keys | Filter | Type | Example | Description | | --------------------| ------ | ------------------------------------ | --------------------------------------------------------------- | | name | string | name in ['policy123', 'policy124'] | Returns unmanaged polici... (see cli schema); illustrative example: "name eq 'foo'"
  -h, --help                help for list
      --limit int           limit (query); Specifies a custom limit for pagination.; format: int64; illustrative example: 1000
      --no-paginate         return only the first page (do not follow nextPage)
      --order-by string     orderBy (query); Specifies the order to sort unmanaged items by fields such as [name, status, createdAt, updatedAt, createdBy.email, policyTemplate.name, schedule, dryRun, category]. Ordering can be [asc] or [desc]. Default ordering is [asc] if no value is set.; illustrative example: "status asc, createdAt asc"
      --skip-token string   resume pagination from this token; Token used for pagination to navigate to the next set of records.
      --view string         view (query); View used to render unmanaged items.; enum: ["default","index"]; illustrative example: "default"
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

* [flexera-cli policy unmanaged-applied-policies](flexera-cli_policy_unmanaged-applied-policies.md)	 - Unmanaged Applied Policies operations (generated from the unified OpenAPI spec)

