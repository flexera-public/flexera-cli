## flexera-cli iam access-policy

Access Policy operations (generated from the unified OpenAPI spec)

### Options

```
  -h, --help   help for access-policy
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
* [flexera-cli iam access-policy create](flexera-cli_iam_access-policy_create.md)	 - Create a new access policy
* [flexera-cli iam access-policy delete](flexera-cli_iam_access-policy_delete.md)	 - delete Access Policy
* [flexera-cli iam access-policy get](flexera-cli_iam_access-policy_get.md)	 - Get an access policy by ID.
* [flexera-cli iam access-policy list](flexera-cli_iam_access-policy_list.md)	 - List all access policies for an organization
* [flexera-cli iam access-policy replace](flexera-cli_iam_access-policy_replace.md)	 - Update an access policy
* [flexera-cli iam access-policy users](flexera-cli_iam_access-policy_users.md)	 - Get all access policies for a specific user

