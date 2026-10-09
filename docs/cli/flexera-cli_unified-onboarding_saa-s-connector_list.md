## flexera-cli unified-onboarding saa-s-connector list

List AICM (SaaS) connectors for the Organization

### Synopsis

List AICM (SaaS) connectors for the Organization

List all AICM connectors for the organization with optional filtering.

```
flexera-cli unified-onboarding saa-s-connector list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli unified-onboarding saa-s-connector list --org-id ORG_ID
```

### Options

```
      --filter string         filter (query); Advanced filter expression for AICM connectors. | Field | Operators | Description | |------------------|-----------|-------------| | connectorName | eq, co | Exact or contains match on connector name | | provider | in | n8n, anthropic, openAi, cursor, salesforceAgentforce, googleWorkspace | | status | in | enrollment status: enabled, failed, deleted, in-prog... (see cli schema)
  -h, --help                  help for list
      --limit int             limit (query); Maximum number of results; API default: 1000
      --offset int            offset (query); Number of results to skip; API default: 0
      --sort field:asc|desc   sort (query); Sort field, optionally with direction as field:asc|desc (default desc; NULLs last). Fields: connectorName, provider, enrollmentStatus, executionStatus, lastRunAt, onboardedAt, updatedAt. Example: `lastRunAt:desc`.
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

* [flexera-cli unified-onboarding saa-s-connector](flexera-cli_unified-onboarding_saa-s-connector.md)	 - SaaS Connector operations (generated from the unified OpenAPI spec)

