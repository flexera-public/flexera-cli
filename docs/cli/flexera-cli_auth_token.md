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
      --login-base-url string   override login base URL
      --org-id int              organization ID
  -o, --output string           output format (json|table)
      --refresh-token string    OAuth refresh token
      --zone string             API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli auth](flexera-cli_auth.md)	 - Acquire OAuth tokens from Flexera login
* [flexera-cli auth token client-credentials](flexera-cli_auth_token_client-credentials.md)	 - Exchange --client-id/--client-secret for an access token
* [flexera-cli auth token refresh](flexera-cli_auth_token_refresh.md)	 - Exchange a --refresh-token for an access token

