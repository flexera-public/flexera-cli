## flexera-cli iam refresh-token list

Index a user's refresh tokens

### Synopsis

Index a user's refresh tokens

Index refresh tokens belonging to a user. Supports optional query parameters orgId and userId.
Org owners can list refresh tokens for a specific user by providing both orgId and userId in the query.
The isDeletable flag in the token response tells the org owner whether they can delete the user token or not.

```
flexera-cli iam refresh-token list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam refresh-token list
```

### Options

```
  -h, --help          help for list
      --user-id int   userId (query); User ID; illustrative example: 67890
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

* [flexera-cli iam refresh-token](flexera-cli_iam_refresh-token.md)	 - Refresh Token operations (generated from the unified OpenAPI spec)

