## flexera-cli finops-onboarding bill-connect snowflake delete

Delete a Snowflake bill connect

### Synopsis

Delete a Snowflake bill connect

Deletes a Snowflake bill connect.
Bill Connects provisioned through Unified Onboarding are read-only in this API. Update and delete operations will be rejected with a 403 Forbidden. Use Unified Onboarding to manage these resources.
You can identify these Bill Connects by the onboardingOrigin field:
- "platform": created via Unified Onboarding
- "finops": created and managed through this API (default)

```
flexera-cli finops-onboarding bill-connect snowflake delete [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-onboarding bill-connect snowflake delete --org-id ORG_ID --id ID
  flexera-cli finops-onboarding bill-connect snowflake delete --org-id ORG_ID --id ID --dry-run
```

### Options

```
      --dry-run[=true|false]   print the planned operation as JSON and exit without calling the API
  -h, --help[=true|false]      help for delete
      --id string              id (path, required); Identifies a bill connect; required by API; illustrative example: "paas-snowflake-TYTMXEN-RS68134"
      --yes[=true|false]       confirm the operation (required for destructive ops)
```

### Options inherited from parent commands

```
      --access-token string        static bearer access token
      --api-base-url string        override API base URL
      --client-id string           OAuth client ID
      --client-secret string       OAuth client secret
      --config string              config file (default $HOME/.flexera/config.yaml)
  -d, --debug[=true|false]         log HTTP requests/responses to stderr (Authorization redacted)
      --json-style string          JSON whitespace style (auto|pretty|compact) (default "auto")
      --login-base-url string      override login base URL
      --no-validate[=true|false]   skip API schema constraints (never JSON syntax or request data-loss checks)
      --org-id int                 organization ID
      --out-fields string          project JSON output fields (comma-separated paths)
      --out-jq string              shape JSON output with a jq expression
  -o, --output string              output format (json|table)
  -r, --raw-output[=true|false]    write jq string results without JSON quotes (requires --out-jq)
      --refresh-token string       OAuth refresh token
      --zone string                API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli finops-onboarding bill-connect snowflake](flexera-cli_finops-onboarding_bill-connect_snowflake.md)	 - Bill Connect - Snowflake operations (generated from the unified OpenAPI spec)

