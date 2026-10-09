## flexera-cli iam user-setting-blob replace

Generates a signed URL to store a user settings object at a specified key, which may represent a page ID or a combination of page ID and prefix ID, used for retrieving user-specific settings.

### Synopsis

Generates a signed URL to store a user settings object at a specified key, which may represent a page ID or a combination of page ID and prefix ID, used for retrieving user-specific settings.

Retrieves a signed URL for saving a user settings object at a specified key, which may represent a page ID, a combination of page ID and prefix ID, or another identifier for user-specific settings.

```
flexera-cli iam user-setting-blob replace [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam user-setting-blob replace --body @request.json
  flexera-cli iam user-setting-blob replace --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema iam user-setting-blob replace --example > request.json
```

### Options

```
      --body string       raw JSON body (inline | @file | @-); overrides body field flags
      --body-org-id int   orgId (body); The unique identifier for the organization; minimum: 1; illustrative example: 100
      --dry-run           print the planned operation as JSON and exit without calling the API
      --expiry int        expiry (body); required by API; expriy time in seconds in signed url; minimum: 1; maximum: 1800; illustrative example: 1800
  -h, --help              help for replace
      --id string         id (body); required by API; The key may be a combination of a Page ID and a Prefix ID. For example: page-settings:optima/cloud-dashboards/82783782372230, where page-settings:optima/cloud-dashboards represents the Page ID and 82783782372230 is the Prefix ID.; minLength: 1; maxLength: 100; pattern: "^[-._a-zA-Z0-9:/=?\u0026]+$"; illustrative example: "preferredLoginMethod"
  -i, --interactive       edit inputs in a terminal form, review a plan and approve with typed yes
      --type string       type (body); The scope/context of the setting, either 'org' or 'global'.; enum: ["global","org"]; illustrative example: "global"
      --yes               confirm the operation (required for destructive ops)
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

* [flexera-cli iam user-setting-blob](flexera-cli_iam_user-setting-blob.md)	 - User Setting Blob operations (generated from the unified OpenAPI spec)

