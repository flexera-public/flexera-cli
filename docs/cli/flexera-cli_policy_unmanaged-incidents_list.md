## flexera-cli policy unmanaged-incidents list

Index unmanaged incidents.

### Synopsis

Index unmanaged incidents.

Index retrieves the list of unmanaged incidents in an organization.

```
flexera-cli policy unmanaged-incidents list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli policy unmanaged-incidents list --org-id ORG_ID
```

### Options

```
      --filter string       filter (query); Optional filter to retrieve unmanaged incidents based on specific criteria. ### Supported Filter Keys | Filter | Type | Example | Description | | ------------------- | ------ | ---------------------------------------- | ---------------------------------------------------------------------------------- | | project.id | number | project.id in [10017, 10134] | ... (see cli schema); illustrative example: "updatedAt lt '2024-01-09T00:00:00Z'"
  -h, --help                help for list
      --limit int           limit (query); Specifies a custom limit for pagination.; format: int64; illustrative example: 1000
      --no-paginate         return only the first page (do not follow nextPage)
      --order-by string     orderBy (query); Specifies the order to sort unmanaged incidents by fields such as [project.id, appliedPolicy.name, createdAt, updatedAt, resolvedAt, resolvedBy.email, actionFailed, actionPending, severity, category, dryRun]. Ordering can be [asc] or [desc]. Default ordering is [asc] if no value is set.; illustrative example: "appliedPolicy.name asc, updatedAt"
      --skip-token string   resume pagination from this token; Used in pagination to point to the next or previous set of records.
      --view string         view (query); View used to render unmanaged incidents.; enum: ["default","index"]; illustrative example: "default"
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

* [flexera-cli policy unmanaged-incidents](flexera-cli_policy_unmanaged-incidents.md)	 - Unmanaged Incidents operations (generated from the unified OpenAPI spec)

