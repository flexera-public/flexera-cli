## flexera-cli saas usage-group costs

costs operations

### Options

```
  -h, --help   help for costs
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

* [flexera-cli saas usage-group](flexera-cli_saas_usage-group.md)	 - Usage Group operations (generated from the unified OpenAPI spec)
* [flexera-cli saas usage-group costs create](flexera-cli_saas_usage-group_costs_create.md)	 - Create a usage cost for the specified usage group
* [flexera-cli saas usage-group costs delete](flexera-cli_saas_usage-group_costs_delete.md)	 - Delete an existing usage cost for the specified Usage Cost object
* [flexera-cli saas usage-group costs get](flexera-cli_saas_usage-group_costs_get.md)	 - Show an existing usage cost for the specified Usage group object
* [flexera-cli saas usage-group costs replace](flexera-cli_saas_usage-group_costs_replace.md)	 - Update an existing cost for the specified usage group

