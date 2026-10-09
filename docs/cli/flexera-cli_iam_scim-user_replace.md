## flexera-cli iam scim-user replace

Replace a user's attributes

### Synopsis

Replace a user's attributes

Replaces a user's attributes in an org.

```
flexera-cli iam scim-user replace [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam scim-user replace --org-id ORG_ID --id ID --body @request.json
  flexera-cli iam scim-user replace --org-id ORG_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema iam scim-user replace --example > request.json
```

### Options

```
      --active             active (body); The user's active status.; API default: true; illustrative example: true
      --body string        raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run            print the planned operation as JSON and exit without calling the API
  -h, --help               help for replace
      --id string          id (path, required); Unique identifier for the user.; required by API; illustrative example: "12345"
  -i, --interactive        edit inputs in a terminal form, review a plan and approve with typed yes
      --schemas strings    schemas (body); List of URIs of the SCIM schemas supported.; CLI: comma-separated values or repeated flag; illustrative example: ["urn:ietf:params:scim:schemas:core:2.0:User"]
      --user-name string   userName (body); The user name of the user.; illustrative example: "jsmith"
      --yes                confirm the operation (required for destructive ops)
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

* [flexera-cli iam scim-user](flexera-cli_iam_scim-user.md)	 - SCIM User operations (generated from the unified OpenAPI spec)

