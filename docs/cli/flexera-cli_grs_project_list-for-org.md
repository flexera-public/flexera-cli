## flexera-cli grs project list-for-org

List the org's projects for the authenticated user (policy-ready project IDs)

### Synopsis

List the projects the authenticated user can access in --org-id. The user is derived from the access token, results are filtered to the org, and each project ID is the legacy account ID that the Policy APIs accept as --project-id.

```
flexera-cli grs project list-for-org [flags]
```

### Examples

```
flexera-cli grs project list-for-org --org-id 123
```

### Options

```
      --api-version string    Optional X-Api-Version header (defaults to 2.0)
      --grs-base-url string   Override the GRS base URL (default: zone-specific grs-front host)
  -h, --help                  help for list-for-org
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

* [flexera-cli grs project](flexera-cli_grs_project.md)	 - Project operations (generated from the unified OpenAPI spec)

