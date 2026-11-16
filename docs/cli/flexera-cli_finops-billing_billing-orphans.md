## flexera-cli finops-billing billing-orphans

Billing Orphans operations (generated from the unified OpenAPI spec)

### Options

```
  -h, --help[=true|false]   help for billing-orphans
```

### Options inherited from parent commands

```
      --access-token string        static bearer access token
      --api-base-url string        override API base URL
      --client-id string           OAuth client ID
      --client-secret string       OAuth client secret
      --config string              config file (default $HOME/.flexera/config.yaml)
  -d, --debug[=true|false]         log HTTP requests/responses to stderr (Authorization redacted)
      --json-style string          JSON whitespace style (auto|pretty|compact) (default "auto")
      --login-base-url string      override login base URL
      --no-validate[=true|false]   skip API schema constraints (never JSON syntax or request data-loss checks)
      --org-id int                 organization ID
      --out-fields string          project JSON output fields (comma-separated paths)
      --out-jq string              shape JSON output with a jq expression
  -o, --output string              output format (json|table)
  -r, --raw-output[=true|false]    write jq string results without JSON quotes (requires --out-jq)
      --refresh-token string       OAuth refresh token
      --zone string                API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli finops-billing](flexera-cli_finops-billing.md)	 - FinOps Billing API
* [flexera-cli finops-billing billing-orphans billing-accounts](flexera-cli_finops-billing_billing-orphans_billing-accounts.md)	 - List Azure billing accounts with orphaned charges
* [flexera-cli finops-billing billing-orphans charges](flexera-cli_finops-billing_billing-orphans_charges.md)	 - List unassigned Azure orphaned charges
* [flexera-cli finops-billing billing-orphans mappings](flexera-cli_finops-billing_billing-orphans_mappings.md)	 - mappings operations

