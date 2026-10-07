## flexera-cli policy-manager

Policy Manager operations (generated from the unified OpenAPI spec)

### Options

```
  -h, --help   help for policy-manager
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

* [flexera-cli](flexera-cli.md)	 - Flexera One unified API command-line client
* [flexera-cli policy-manager create](flexera-cli_policy-manager_create.md)	 - Create a Policy Manager
* [flexera-cli policy-manager delete](flexera-cli_policy-manager_delete.md)	 - Delete a policy manager
* [flexera-cli policy-manager get](flexera-cli_policy-manager_get.md)	 - Get a policy manager
* [flexera-cli policy-manager list](flexera-cli_policy-manager_list.md)	 - List policy managers
* [flexera-cli policy-manager summary](flexera-cli_policy-manager_summary.md)	 - Retrieve Summary by Organization
* [flexera-cli policy-manager update](flexera-cli_policy-manager_update.md)	 - Update a policy manager
* [flexera-cli policy-manager update-template](flexera-cli_policy-manager_update-template.md)	 - Update policy manager to reference the latest published template version

