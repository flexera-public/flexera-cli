## flexera-cli recommendations export

Export recommendations as CSV

### Synopsis

Export recommendations as CSV

Exports recommendations as a CSV file: every recommendation the paged list returns for the same `filter` and `orderBy`, with the same billing center rules, one row each, with the given columns.

**Required security scopes for GlobalSession**:
  * `common:org:affiliated+common:org:own`

```
flexera-cli recommendations export [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli recommendations export --org-id ORG_ID --columns COLUMNS
```

### Options

```
      --columns strings   columns (query); The columns of the file, in order, each as <field>:<header>: the field to write and the text of its header. Everything after the first ':' is the header. ### Supported Fields | Field | Value | | ------------------------- | --------------------------------------------------------------------------------------------- | | id, recommendation, type, policySet, so... (see cli schema); required by API; minItems: 1; CLI: comma-separated values or repeated flag
      --filter string     filter (query); Optional filter for returning recommendations matching specific criteria. ### Supported Filter Keys | Filter | Type | Example | Description | | --------------------- | ------ | ------------------------------------------------- | ------------------------------------------------------------------------------------ | | billingCenterID | string | billingCenterID... (see cli schema)
  -h, --help              help for export
      --order-by string   orderBy (query); Optional order by for returning recommendations in a specific order. ### Supported Order by Keys | Key | Description | | -------------------- | ------------------------------------------------------------------------------- | | recommendation | Returns recommendations ordered by 'recommendation' property | | type | Returns recommendations ordered by 'type' p... (see cli schema)
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

