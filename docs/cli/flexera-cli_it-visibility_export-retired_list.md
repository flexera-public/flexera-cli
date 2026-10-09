## flexera-cli it-visibility export-retired list

List all inventory exports

### Synopsis

List all inventory exports

Index returns a list of inventory exports to which an org has access.
        Retrieves file list by filter params

```
flexera-cli it-visibility export-retired list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli it-visibility export-retired list --org-id ORG_ID
```

### Options

```
      --filter string       filter (query); Optional filter for returning requests matching specific criteria. The following filters are supported: | Filter | Description | Allowed Operator | Behavior | Example | | --- | ---| --- | --- | --- | | id | Filters on the export id | eq | Equal - The attribute and operator values must be identical for a match | id eq '12345' | | name | Filters on the export ... (see cli schema); illustrative example: "type eq 'snapshot' and size ge 0"
  -h, --help                help for list
      --no-paginate         return only the first page (do not follow nextPage)
      --skip-token string   resume pagination from this token; An opaque token to be provided when requesting a subsequent page after receiving a partial response. Partial responses will include a "nextPage" attribute in their response body, which contains the URL of the next page including the appropriate skipToken.
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

* [flexera-cli it-visibility export-retired](flexera-cli_it-visibility_export-retired.md)	 - Export (Retired) operations (generated from the unified OpenAPI spec)

