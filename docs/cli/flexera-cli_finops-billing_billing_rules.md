## flexera-cli finops-billing billing rules

rules operations

### Options

```
  -h, --help   help for rules
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

* [flexera-cli finops-billing billing](flexera-cli_finops-billing_billing.md)	 - Billing operations (generated from the unified OpenAPI spec)
* [flexera-cli finops-billing billing rules create](flexera-cli_finops-billing_billing_rules_create.md)	 - Create an enterprise adjustment rule
* [flexera-cli finops-billing billing rules delete](flexera-cli_finops-billing_billing_rules_delete.md)	 - Remove an enterprise adjustment rule
* [flexera-cli finops-billing billing rules list](flexera-cli_finops-billing_billing_rules_list.md)	 - Index enterprise adjustment rules
* [flexera-cli finops-billing billing rules replace](flexera-cli_finops-billing_billing_rules_replace.md)	 - Overwrite an enterprise adjustment rule
* [flexera-cli finops-billing billing rules replace-all](flexera-cli_finops-billing_billing_rules_replace-all.md)	 - Overwrite enterprise adjustment rules

