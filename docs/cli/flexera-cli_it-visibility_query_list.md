## flexera-cli it-visibility query list

Download query results

### Synopsis

Download query results

Download the specified query results.

```
flexera-cli it-visibility query list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli it-visibility query list --org-id ORG_ID --id ID
```

### Options

```
  -h, --help                      help for list
      --id string                 id (path, required); The unique query identifier; required by API; maxLength: 36; pattern: "^[{]?[0-9a-fA-F]{8}-([0-9a-fA-F]{4}-){3}[0-9a-fA-F]{12}[}]?$"; illustrative example: "7d795afa-c508-4d47-92ae-248113a792d1"
      --max-results int           maxResults (query); The maximum number of rows to return for uncompressed CSV downloads. The maxResults parameter cannot be used with outputCompression=gzip. If both parameters are specified in the same request, the API will return a 400 Bad Request error.; maximum: 100000; API default: 0; illustrative example: 100000
      --query-skip-token string   skipToken (query); An opaque token to be provided when requesting a subsequent page after receiving a partial response. Partial responses will include a "nextPage" attribute in their response body, which contains the URL of the next page including the appropriate skipToken.
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

* [flexera-cli it-visibility query](flexera-cli_it-visibility_query.md)	 - Query operations (generated from the unified OpenAPI spec)

