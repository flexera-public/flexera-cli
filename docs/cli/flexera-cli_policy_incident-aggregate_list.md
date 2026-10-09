## flexera-cli policy incident-aggregate list

Index incident aggregates.

### Synopsis

Index incident aggregates.

Index retrieves the list of incident aggregates in an organization.

```
flexera-cli policy incident-aggregate list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli policy incident-aggregate list --org-id ORG_ID
```

### Options

```
      --filter string       filter (query); Optional filter to retrieve incident aggregates based on specific criteria. ### Supported Filter Keys | Filter | Type | Example | Description | | ---------- | ------ | ------------------------------------ | ---------------------------------------------------------------- | | severity | string | severity in ['high', 'critical'] | Returns incident aggregates w... (see cli schema); illustrative example: "category eq 'cost'"
  -h, --help                help for list
      --limit int           limit (query); Specifies a custom limit for pagination.; format: int64; illustrative example: 1000
      --no-paginate         return only the first page (do not follow nextPage)
      --order-by string     orderBy (query); Specifies the order to sort incident aggregates by fields such as [severity, createdAt, updatedAt, category, dryRun]. Ordering can be [asc] or [desc]. Default ordering is [asc] if no value is set.; illustrative example: "severity asc, createdAt asc"
      --skip-token string   resume pagination from this token; Used in pagination to point to the next or previous set of records.
      --view string         view (query); View used to render incident aggregates.; enum: ["default"]; illustrative example: "default"
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

* [flexera-cli policy incident-aggregate](flexera-cli_policy_incident-aggregate.md)	 - Incident Aggregate operations (generated from the unified OpenAPI spec)

