## flexera-cli unified-onboarding connector status

Get status for all enabled products for a given connector

### Synopsis

Get status for all enabled products for a given connector

Get status for all enabled products for a given connector along with errors, if any.

```
flexera-cli unified-onboarding connector status [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli unified-onboarding connector status --org-id ORG_ID --connector-id CONNECTOR_ID
```

### Options

```
      --account-id string     account_id (query); Account ID (optional)
      --connector-id string   connector_id (query); Connector ID (required); required by API
  -h, --help                  help for status
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

* [flexera-cli unified-onboarding connector](flexera-cli_unified-onboarding_connector.md)	 - Connector operations (generated from the unified OpenAPI spec)

