## flexera-cli policy policy-manager list

List policy managers

### Synopsis

List policy managers

Lists all policy managers for the given organization.

```
flexera-cli policy policy-manager list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli policy policy-manager list --org-id ORG_ID
```

### Options

```
      --filter string         filter (query); Optional filter to retrieve policy managers based on specific criteria. ### Supported Filter Keys | Filter | Type | Example | Description | | ---------------------- | ------ | ------------------------------------ | --------------------------------------------------------------- | | name | string | name in ['policy123', 'policy124'] | Returns policy managers ... (see cli schema); illustrative example: "name eq 'foo' or dryRun eq true"
  -h, --help                  help for list
      --include-terminating   includeTerminating (query); When true, includes policy managers that are in the process of termination.; API default: false; illustrative example: false
      --limit int             limit (query); Custom pagination limit to be used.; format: int64; illustrative example: 1000
      --no-paginate           return only the first page (do not follow nextPage)
      --order-by string       orderBy (query); Specifies the order to sort policy managers by fields such as [name, status, createdAt, updatedAt, createdBy.email, template.name, schedule, dryRun, category]. Ordering can be [asc] or [desc]. Default ordering is [asc] if no value is set.; illustrative example: "status asc, createdAt asc"
      --skip-token string     resume pagination from this token; Used in pagination to point to the next or previous set of records.
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

* [flexera-cli policy policy-manager](flexera-cli_policy_policy-manager.md)	 - Policy Manager operations (generated from the unified OpenAPI spec)

