## flexera-cli saas managed-application app-users

List managed application's users

### Synopsis

List managed application's users

Retrieves a managed application's users identified by managed application Id.

```
flexera-cli saas managed-application app-users [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas managed-application app-users --org-id ORG_ID --managed-app-id MANAGED_APP_ID
```

### Options

```
  -h, --help                    help for app-users
      --include-all             includeAll (query); Flag to determine whether or not inactive users should be included in the results; API default: false; illustrative example: false
      --managed-app-id string   managedAppId (path, required); managedAppId identifies an managed application by given Id.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "123"
      --no-paginate             return only the first page (do not follow nextPage)
      --skip-token string       resume pagination from this token; An opaque token to be provided when requesting a subsequent page after receiving a partial response. Partial responses will include a "nextPage" attribute in their response body, which contains the URL of the next page including the appropriate skipToken.
      --view string             view (query); View used to render managed application users; enum: ["default","minimal","excludeEvent","entitlement"]; illustrative example: "default"
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

* [flexera-cli saas managed-application](flexera-cli_saas_managed-application.md)	 - Managed Application operations (generated from the unified OpenAPI spec)

