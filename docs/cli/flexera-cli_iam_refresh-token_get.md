## flexera-cli iam refresh-token get

Show a refresh token

### Synopsis

Show a refresh token

```
flexera-cli iam refresh-token get [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam refresh-token get --id ID
```

### Options

```
  -h, --help        help for get
      --id string   id (path, required); A short string that identifies the token but which cannot be used to gain access; required by API; pattern: "^[a-zA-Z0-9]+$"; illustrative example: "a1B2c3D4e5F6g7H8i9J0"
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

