## flexera-cli credential delete-project

Delete a Credential

### Synopsis

Delete a Credential

Delete a Credential that uses the given scheme.
If the scheme is incorrect then the API will return not found.

```
flexera-cli credential delete-project [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli credential delete-project --project-id PROJECT_ID --scheme SCHEME --id ID
  flexera-cli credential delete-project --project-id PROJECT_ID --scheme SCHEME --id ID --dry-run
```

### Options

```
      --dry-run          print the planned operation as JSON and exit without calling the API
  -h, --help             help for delete-project
      --id string        id (path, required); Credentials id; required by API; pattern: "^[_a-zA-Z0-9][-_a-zA-Z0-9]{0,127}$"; illustrative example: "abcdefghij-123456790"
      --project-id int   projectId (path, required); Identifies the Project that owns the Credential.; required by API; format: int64; minimum: 1; illustrative example: 2345
      --scheme string    scheme (path, required); The name of the security scheme.; required by API; enum: ["api-key","aws","aws-sts","basic","digest","ntlm","oauth2","oracle"]; illustrative example: "ntlm"
      --yes              confirm the operation (required for destructive ops)
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

* [flexera-cli credential](flexera-cli_credential.md)	 - Credential API

