## flexera-cli iam user

User operations (generated from the unified OpenAPI spec)

### Options

```
  -h, --help   help for user
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

* [flexera-cli iam](flexera-cli_iam.md)	 - Identity and Access Management API
* [flexera-cli iam user create](flexera-cli_iam_user_create.md)	 - Create and affiliate a user to an org
* [flexera-cli iam user delete](flexera-cli_iam_user_delete.md)	 - Delete user
* [flexera-cli iam user get](flexera-cli_iam_user_get.md)	 - Show an individual org user
* [flexera-cli iam user groups](flexera-cli_iam_user_groups.md)	 - Get user groups
* [flexera-cli iam user list](flexera-cli_iam_user_list.md)	 - Index an org's users
* [flexera-cli iam user update](flexera-cli_iam_user_update.md)	 - Update user name

