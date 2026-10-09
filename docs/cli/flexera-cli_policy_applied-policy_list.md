## flexera-cli policy applied-policy list

List applied policies

### Synopsis

List applied policies

Index retrieves the list of applied policies in a project.

```
flexera-cli policy applied-policy list [flags]
```

### Options

```
      --filter string       Optional filter expression; Optional filter for returning applied policies matching specific criteria. ### Supported Filter Keys | Filter | Type | Example | Description | | ---------------------- | ------ | -------------------------------------------------- | -------------------------------------------------------------------------------------------------- | | name | string | name in [... (see cli schema); illustrative example: "name eq 'foo' or name eq 'foo 2'"
  -h, --help                help for list
      --limit int           Optional page size; Custom pagination limit to be used.; format: int64; illustrative example: 1000
      --no-paginate         return only the requested page (do not follow nextPage)
      --order-by string     Optional sort expression; Optional order by for returning applied policies in a specific order. ### Supported Order by Keys | Filter | Description | | ---------------------- | ------------------------------------------------------------------------------------------------- | | name | Returns applied policies ordered by 'name' property | | metaParentPolicyId | Returns applied policies... (see cli schema); illustrative example: "createdAt asc"
      --project-id int      Project ID (optional; resolved from GRS for the org when omitted); The unique identifier for the project; required by API; format: int64; minimum: 1; illustrative example: 60073
      --skip-token string   Optional pagination token; resume from this position; Used in pagination to point to the next or previous set of records.
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

* [flexera-cli policy applied-policy](flexera-cli_policy_applied-policy.md)	 - Applied policies (project-scoped)

