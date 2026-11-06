## flexera-cli unified-onboarding onboarding snowflake-validate

Snowflake Onboarding: Validate connection

### Synopsis

Snowflake Onboarding: Validate connection

Synchronously test **submitted** Snowflake credentials. Returns HTTP 200 for every outcome: `success:true` when Snowflake accepts them, otherwise `success:false` with the vendor HTTP `code` (or 502 unreachable / 504 timeout), a normalized `reason` and a specific `detail`. A malformed body is a 422.

```
flexera-cli unified-onboarding onboarding snowflake-validate [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli unified-onboarding onboarding snowflake-validate --org-id ORG_ID --body @request.json
  flexera-cli unified-onboarding onboarding snowflake-validate --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema unified-onboarding onboarding snowflake-validate --example > request.json
```

### Options

```
      --account-id string                  AccountId (body); required by API; Snowflake account identifier.; pattern: "^[A-Za-z0-9]([A-Za-z0-9._-]{0,98}[A-Za-z0-9])?$"
      --body string                        raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run[=true|false]               print the planned operation as JSON and exit without calling the API
  -h, --help[=true|false]                  help for snowflake-validate
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

* [flexera-cli unified-onboarding onboarding](flexera-cli_unified-onboarding_onboarding.md)	 - Onboarding operations (generated from the unified OpenAPI spec)

