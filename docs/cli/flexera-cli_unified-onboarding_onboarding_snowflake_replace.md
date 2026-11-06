## flexera-cli unified-onboarding onboarding snowflake replace

Onboarding Snowflake: Update

### Synopsis

Onboarding Snowflake: Update

Update an existing Snowflake connector. Supply `ProgrammaticAccessToken` to rotate the stored credential. `AccountId` cannot be changed.

```
flexera-cli unified-onboarding onboarding snowflake replace [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli unified-onboarding onboarding snowflake replace --org-id ORG_ID --connector-id CONNECTOR_ID --body @request.json
  flexera-cli unified-onboarding onboarding snowflake replace --org-id ORG_ID --connector-id CONNECTOR_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema unified-onboarding onboarding snowflake replace --example > request.json
```

### Options

```
      --body string                        raw JSON body (inline | @file | @-); overrides body field flags
      --connector-id string                connector_id (path, required); Connector ID; required by API
      --connector-name string              ConnectorName (body); Connector name.; minLength: 3; pattern: "^[a-zA-Z0-9._-]+$"
      --dry-run[=true|false]               print the planned operation as JSON and exit without calling the API
  -h, --help[=true|false]                  help for replace
      --include-cost-and-usage string      IncludeCostAndUsage (body); Include cost and usage flag.
      --include-dco string                 IncludeDCO (body); Include DCO flag.
  -i, --interactive[=true|false]           edit inputs in a terminal form, review a plan and approve with typed yes
      --programmatic-access-token string   ProgrammaticAccessToken (body); Replacement programmatic access token (PAT).
      --role string                        Role (body); Snowflake role.; minLength: 1
      --user-name string                   UserName (body); Snowflake user name.; minLength: 1
      --warehouse-name string              WarehouseName (body); Snowflake warehouse.; minLength: 1
      --yes[=true|false]                   confirm the operation (required for destructive ops)
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

* [flexera-cli unified-onboarding onboarding snowflake](flexera-cli_unified-onboarding_onboarding_snowflake.md)	 - snowflake operations

