## flexera-cli iam user-setting-blob list

Generates a signed URL to retrieve a user settings object from a specified key, which may represent a page ID, a combination of page ID and prefix ID, or another identifier for accessing user-specific settings

```
flexera-cli iam user-setting-blob list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam user-setting-blob list --id ID --expiry EXPIRY
```

### Options

```
      --expiry int    expiry (query)
  -h, --help          help for list
      --id string     id (query)
      --type string   type (query)
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

* [flexera-cli iam user-setting-blob](flexera-cli_iam_user-setting-blob.md)	 - User Setting Blob operations (generated from the unified OpenAPI spec)

