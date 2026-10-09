## flexera-cli grs user list

List organizations the authenticated user can access

### Synopsis

List organizations the authenticated user can access

Returns the Flexera One organizations that the user identified by `{userId}` has access to with the caller's credentials. This endpoint is the canonical 'who am I and what orgs do I belong to?' lookup for user-issued tokens (refresh-token grant) and is the recommended first call to validate that a CLI/SDK is configured correctly.

This endpoint requires a user-scoped access token. Service-account tokens (subject prefix `sa-`) do not have a user identity and cannot call this route — callers should detect that case before issuing the request and surface a clear error.

NOTE: Served by the legacy GRS service, which is being deprecated. Prefer newer Flexera One identity APIs when available.

Deprecated in the upstream API.

```
flexera-cli grs user list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli grs user list --user-id USER_ID
```

### Options

```
  -h, --help                 help for list
      --user-id u-{userId}   userId (path, required); Flexera user ID (numeric). For a user-issued access token this is the integer suffix of the JWT 'sub' claim (u-{userId}) or the legacy 'user' claim. Service-account tokens cannot call this endpoint.; required by API; format: int64; minimum: 1
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

* [flexera-cli grs user](flexera-cli_grs_user.md)	 - User operations (generated from the unified OpenAPI spec)

