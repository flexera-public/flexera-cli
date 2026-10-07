## flexera-cli auth token create

Generate access token

### Synopsis

Generate a Flexera One access token using POST /oidc/token. JSON input is encoded as application/x-www-form-urlencoded. No prior authentication is required; inherited authentication flags are not request-body fields. Grant-specific field requirements are enforced by the login service.

```
flexera-cli auth token create [flags]
```

### Examples

```
flexera-cli auth token create --body @request.json --dry-run
```

### Options

```
      --body string                 JSON body (inline | @file | @-); overrides body field flags
      --body-client-id string       Service app client ID (required for client_credentials) (body)
      --body-client-secret string   Service app client secret (required for client_credentials) (body)
      --body-refresh-token string   Refresh token from a previous offline_access authorization code exchange (body)
      --code string                 Authorization code (required for authorization_code) (body)
      --dry-run                     preview the request with sensitive values redacted; no HTTP request
      --grant-type string           Required: authorization_code, client_credentials, or refresh_token (body)
  -h, --help                        help for create
      --redirect-uri string         Redirect URI (required for authorization_code) (body)
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

* [flexera-cli auth token](flexera-cli_auth_token.md)	 - Mint an access token

