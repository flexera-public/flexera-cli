## flexera-cli saas usage-message-query list

Retrieve usage for the given period of time and for given usage group.

```
flexera-cli saas usage-message-query list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas usage-message-query list --org-id ORG_ID
```

### Options

```
      --ended-at string     endedAt (query, RFC3339 date-time with timezone)
      --filter string       filter (query)
      --group-by strings    groupBy (query)
  -h, --help                help for list
      --no-paginate         return only the first page (do not follow nextPage)
      --resolution string   resolution (query)
      --skip-token string   resume pagination from this token
      --started-at string   startedAt (query, RFC3339 date-time with timezone)
      --view string         view (query)
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

* [flexera-cli saas usage-message-query](flexera-cli_saas_usage-message-query.md)	 - Usage Message Query operations (generated from the unified OpenAPI spec)

