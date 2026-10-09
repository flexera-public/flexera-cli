## flexera-cli saas managed-application-user list

List managed application users

### Synopsis

List managed application users

Retrieves a collection of managed application users.

```
flexera-cli saas managed-application-user list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas managed-application-user list --org-id ORG_ID
```

### Options

```
      --filter string       filter (query); The date filter to query the managed application users. Supported fields in the filter are [managedAppId]; illustrative example: "managedAppId eq 123"
  -h, --help                help for list
      --include-all         includeAll (query); Flag to determine whether or not inactive users should be included in the results; API default: false; illustrative example: true
      --no-paginate         return only the first page (do not follow nextPage)
      --skip-token string   resume pagination from this token; An opaque token to be provided when requesting a subsequent page after receiving a partial response. Partial responses will include a "nextPage" attribute in their response body, which contains the URL of the next page including the appropriate skipToken.
      --view string         view (query); View used to render integration template; enum: ["default","minimal","excludeEvent","entitlement"]; illustrative example: "default"
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

* [flexera-cli saas managed-application-user](flexera-cli_saas_managed-application-user.md)	 - Managed Application User operations (generated from the unified OpenAPI spec)

