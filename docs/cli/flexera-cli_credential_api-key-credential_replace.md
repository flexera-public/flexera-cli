## flexera-cli credential api-key-credential replace

Create a Credential

### Synopsis

Create a Credential

Create a Credential that uses the 'API Key' scheme.

```
flexera-cli credential api-key-credential replace [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli credential api-key-credential replace --org-id ORG_ID --id ID --body @request.json
  flexera-cli credential api-key-credential replace --org-id ORG_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema credential api-key-credential replace --example > request.json
```

### Options

```
      --body string          raw JSON body (inline | @file | @-); overrides body field flags
      --description string   description (body); Credentials description; illustrative example: "The AWS Oregon region (us-west-2) development credentials."
      --dry-run              print the planned operation as JSON and exit without calling the API
      --field string         field (body); Name of either the authorization header key or the query parameter field; API default: "Authorization"; illustrative example: "Authorization"
  -h, --help                 help for replace
      --id string            id (path, required); Credentials id; required by API; pattern: "^[_a-zA-Z0-9][-_a-zA-Z0-9]{0,127}$"; illustrative example: "abcdefghij-123456790"
  -i, --interactive          edit inputs in a terminal form, review a plan and approve with typed yes
      --key string           key (body); required by API; Static API key; illustrative example: "sumthingSeekrit"
      --location string      location (body); Location of the authorization; enum: ["header","query"]; API default: "header"; illustrative example: "query"
      --name string          name (body); required by API; Credentials name used in UI; illustrative example: "Development Credentials"
      --type string          type (body); Type of authorization header, prefixes the key value; API default: "Bearer"; illustrative example: "Bearer"
      --yes                  confirm the operation (required for destructive ops)
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

* [flexera-cli credential api-key-credential](flexera-cli_credential_api-key-credential.md)	 - API Key Credential operations (generated from the unified OpenAPI spec)

