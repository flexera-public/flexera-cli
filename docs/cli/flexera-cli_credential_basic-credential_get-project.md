## flexera-cli credential basic-credential get-project

Show a Credential

### Synopsis

Show a Credential

Show a Credential that uses the 'Basic' scheme.

```
flexera-cli credential basic-credential get-project [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli credential basic-credential get-project --project-id PROJECT_ID --id ID
```

### Options

```
  -h, --help             help for get-project
      --id string        id (path, required); Credentials id; required by API; pattern: "^[_a-zA-Z0-9][-_a-zA-Z0-9]{0,127}$"; illustrative example: "abcdefghij-123456790"
      --project-id int   projectId (path, required); Identifies the Project that owns the Credential.; required by API; format: int64; minimum: 1; illustrative example: 2345
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

* [flexera-cli credential basic-credential](flexera-cli_credential_basic-credential.md)	 - Basic Credential operations (generated from the unified OpenAPI spec)

