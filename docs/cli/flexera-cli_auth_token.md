## flexera-cli auth token

Mint an access token

### Options

```
  -h, --help   help for token
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

* [flexera-cli auth](flexera-cli_auth.md)	 - Acquire OAuth tokens from Flexera login
* [flexera-cli auth token client-credentials](flexera-cli_auth_token_client-credentials.md)	 - Exchange --client-id/--client-secret for an access token
* [flexera-cli auth token create](flexera-cli_auth_token_create.md)	 - Generate access token
* [flexera-cli auth token refresh](flexera-cli_auth_token_refresh.md)	 - Exchange a --refresh-token for an access token

