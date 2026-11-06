## flexera-cli unified-onboarding onboarding snowflake create

Onboarding Snowflake: Create

### Synopsis

Onboarding Snowflake: Create

Onboard a new Snowflake connector. The programmatic access token is stored by the subscriber and never echoed back.

```
flexera-cli unified-onboarding onboarding snowflake create [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli unified-onboarding onboarding snowflake create --org-id ORG_ID --body @request.json
  flexera-cli unified-onboarding onboarding snowflake create --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema unified-onboarding onboarding snowflake create --example > request.json
```

### Options

```
      --account-id string                  AccountId (body); required by API; Snowflake account identifier.; pattern: "^[A-Za-z0-9]([A-Za-z0-9._-]{0,98}[A-Za-z0-9])?$"
      --body string                        raw JSON body (inline | @file | @-); overrides body field flags
      --connector-name string              ConnectorName (body); required by API; Connector name.; minLength: 3; pattern: "^[a-zA-Z0-9._-]+$"
      --dry-run[=true|false]               print the planned operation as JSON and exit without calling the API
  -h, --help[=true|false]                  help for create
      --include-cost-and-usage string      IncludeCostAndUsage (body); required by API; Include cost and usage flag ('true'/'false').
      --include-dco string                 IncludeDCO (body); required by API; Include DCO flag ('true'/'false').
  -i, --interactive[=true|false]           edit inputs in a terminal form, review a plan and approve with typed yes
      --programmatic-access-token string   ProgrammaticAccessToken (body); required by API; Snowflake programmatic access token (PAT).
      --role string                        Role (body); required by API; Snowflake role used for collection.; minLength: 1
      --user-name string                   UserName (body); required by API; Snowflake user name.; minLength: 1
      --warehouse-name string              WarehouseName (body); required by API; Snowflake warehouse used for collection.; minLength: 1
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

