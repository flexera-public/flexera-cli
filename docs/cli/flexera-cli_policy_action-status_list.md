## flexera-cli policy action-status list

List action statuses

### Synopsis

List action statuses

Index retrieves a list of action statuses in a project.

```
flexera-cli policy action-status list [flags]
```

### Options

```
      --filter string       Optional filter expression; Optional filter to retrieve records based on specific criteria. ### Supported Filter Keys | Filter | Type | Example | Description | | ---------------- | ------ | ---------------------------------------------- | ---------------------------------------------------------------------------------- | | action.id | string | action.id eq '662bc1783a9036b3b8081da3' |... (see cli schema); illustrative example: "incident.id eq '5b06ead5e0dacc007058c784' and appliedPolicy.id eq '6b06ebb5e0dacc010058c784'"
  -h, --help                help for list
      --limit int           Optional page size; Specifies a custom limit for pagination.; format: int64; illustrative example: 1000
      --no-paginate         return only the requested page (do not follow nextPage)
      --order-by string     Optional sort expression; Specifies the order to sort action statuses by fields such as [action.id, incident.id, appliedPolicy.id, type, status, runBy.email, runBy.name, name, startedAt, finishedAt]. Ordering can be [asc] or [desc]. Default ordering is [asc] if no value is set.; illustrative example: "incident.id asc, appliedPolicy.id asc"
      --project-id int      Project ID (optional; resolved from GRS for the org when omitted); The unique identifier for the project; required by API; format: int64; minimum: 1; illustrative example: 60073
      --skip-token string   Optional pagination token; resume from this position; Used in pagination to point to the next or previous set of records.
      --view string         Optional Policy action-status view; View used to render action statuses.; enum: ["default","extended"]; illustrative example: "default"
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

* [flexera-cli policy action-status](flexera-cli_policy_action-status.md)	 - Policy action statuses (project-scoped)

