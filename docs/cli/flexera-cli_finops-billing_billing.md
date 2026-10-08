## flexera-cli finops-billing billing

Billing operations (generated from the unified OpenAPI spec)

### Options

```
  -h, --help   help for billing
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

* [flexera-cli finops-billing](flexera-cli_finops-billing.md)	 - FinOps Billing API
* [flexera-cli finops-billing billing create-plans](flexera-cli_finops-billing_billing_create-plans.md)	 - Create an adjustment plan rule
* [flexera-cli finops-billing billing create-rules](flexera-cli_finops-billing_billing_create-rules.md)	 - Create an enterprise adjustment rule
* [flexera-cli finops-billing billing customer-status](flexera-cli_finops-billing_billing_customer-status.md)	 - Show billing customer status
* [flexera-cli finops-billing billing delete-plans](flexera-cli_finops-billing_billing_delete-plans.md)	 - Remove an adjustment plan rule
* [flexera-cli finops-billing billing delete-rules](flexera-cli_finops-billing_billing_delete-rules.md)	 - Remove an enterprise adjustment rule
* [flexera-cli finops-billing billing list-plans](flexera-cli_finops-billing_billing_list-plans.md)	 - Index adjustment plan rules
* [flexera-cli finops-billing billing list-rules](flexera-cli_finops-billing_billing_list-rules.md)	 - Index enterprise adjustment rules
* [flexera-cli finops-billing billing plans](flexera-cli_finops-billing_billing_plans.md)	 - Overwrite an adjustment plan
* [flexera-cli finops-billing billing plans-all](flexera-cli_finops-billing_billing_plans-all.md)	 - Create an adjustment plan
* [flexera-cli finops-billing billing plans-all-2](flexera-cli_finops-billing_billing_plans-all-2.md)	 - Remove an adjustment plan
* [flexera-cli finops-billing billing plans-all-3](flexera-cli_finops-billing_billing_plans-all-3.md)	 - Show an adjustment plan
* [flexera-cli finops-billing billing plans-all-4](flexera-cli_finops-billing_billing_plans-all-4.md)	 - Index adjustment plans
* [flexera-cli finops-billing billing post-plan-rules](flexera-cli_finops-billing_billing_post-plan-rules.md)	 - Overwrite an adjustment rule that runs after all plans
* [flexera-cli finops-billing billing post-plan-rules-all](flexera-cli_finops-billing_billing_post-plan-rules-all.md)	 - Create an adjustment rule that runs after all plans
* [flexera-cli finops-billing billing post-plan-rules-all-2](flexera-cli_finops-billing_billing_post-plan-rules-all-2.md)	 - Remove an adjustment rule that runs after all plans
* [flexera-cli finops-billing billing post-plan-rules-all-3](flexera-cli_finops-billing_billing_post-plan-rules-all-3.md)	 - Index adjustment rules that run after all plans
* [flexera-cli finops-billing billing pre-plan-rules](flexera-cli_finops-billing_billing_pre-plan-rules.md)	 - Overwrite an adjustment rule that runs before all plans
* [flexera-cli finops-billing billing pre-plan-rules-all](flexera-cli_finops-billing_billing_pre-plan-rules-all.md)	 - Create an adjustment rule that runs before all plans
* [flexera-cli finops-billing billing pre-plan-rules-all-2](flexera-cli_finops-billing_billing_pre-plan-rules-all-2.md)	 - Remove an adjustment rule that runs before all plans
* [flexera-cli finops-billing billing pre-plan-rules-all-3](flexera-cli_finops-billing_billing_pre-plan-rules-all-3.md)	 - Index adjustment rules that run before all plans
* [flexera-cli finops-billing billing replace](flexera-cli_finops-billing_billing_replace.md)	 - Overwrite an adjustment plan rule
* [flexera-cli finops-billing billing replace-all](flexera-cli_finops-billing_billing_replace-all.md)	 - Overwrite enterprise adjustment rules
* [flexera-cli finops-billing billing replace-all-2](flexera-cli_finops-billing_billing_replace-all-2.md)	 - Overwrite an enterprise adjustment rule

