## flexera-cli grs project list

List projects accessible by the authenticated user

### Synopsis

List projects accessible by the authenticated user

Returns the projects accessible by the user identified by `{userId}`. The authoritative project ID for legacy Flexera APIs is `legacy.account_id` when present (fall back to top-level `id` when `legacy.account_id` is absent).

NOTE: This endpoint is served by the legacy GRS service, which is being deprecated. Prefer newer Flexera One APIs when available.

Deprecated in the upstream API.

```
flexera-cli grs project list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli grs project list --user-id USER_ID
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

* [flexera-cli grs project](flexera-cli_grs_project.md)	 - Project operations (generated from the unified OpenAPI spec)

