## flexera-cli unified-onboarding onboarding delete

Onboarding: Delete

### Synopsis

Onboarding: Delete

Remove an existing cloud connector for AWS, Azure, or GCP.

```
flexera-cli unified-onboarding onboarding delete [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli unified-onboarding onboarding delete --org-id ORG_ID --provider PROVIDER --connector-id CONNECTOR_ID
  flexera-cli unified-onboarding onboarding delete --org-id ORG_ID --provider PROVIDER --connector-id CONNECTOR_ID --dry-run
```

### Options

```
      --connector-id string   connector_id (query); Connector ID; required by API
      --dry-run               print the planned operation as JSON and exit without calling the API
  -h, --help                  help for delete
      --provider string       provider (query); Provider (aws/azure/gcp); required by API
      --yes                   confirm the operation (required for destructive ops)
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

* [flexera-cli unified-onboarding onboarding](flexera-cli_unified-onboarding_onboarding.md)	 - Onboarding operations (generated from the unified OpenAPI spec)

