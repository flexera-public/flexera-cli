## flexera-cli policy applied-policy

Applied policies (project-scoped)

```
flexera-cli policy applied-policy [flags]
```

### Options

```
  -h, --help   help for applied-policy
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

* [flexera-cli policy](flexera-cli_policy.md)	 - Project-scoped policy operations with GRS project auto-resolution
* [flexera-cli policy applied-policy create](flexera-cli_policy_applied-policy_create.md)	 - Create an applied policy
* [flexera-cli policy applied-policy delete](flexera-cli_policy_applied-policy_delete.md)	 - Delete an applied policy
* [flexera-cli policy applied-policy evaluate](flexera-cli_policy_applied-policy_evaluate.md)	 - Request evaluation of an applied policy
* [flexera-cli policy applied-policy get](flexera-cli_policy_applied-policy_get.md)	 - Show an applied policy
* [flexera-cli policy applied-policy list](flexera-cli_policy_applied-policy_list.md)	 - List applied policies
* [flexera-cli policy applied-policy log](flexera-cli_policy_applied-policy_log.md)	 - Show an applied policy's log
* [flexera-cli policy applied-policy status](flexera-cli_policy_applied-policy_status.md)	 - Show an applied policy's status
* [flexera-cli policy applied-policy update](flexera-cli_policy_applied-policy_update.md)	 - Update an applied policy

