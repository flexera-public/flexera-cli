## flexera-cli

Flexera One unified API command-line client

### Synopsis

Flexera One unified API command-line client.

Find commands for a task: flexera-cli cli search "<what you want to do>"

### Options

```
      --access-token string     static bearer access token
      --api-base-url string     override API base URL
      --client-id string        OAuth client ID
      --client-secret string    OAuth client secret
      --config string           config file (default $HOME/.flexera/config.yaml)
  -d, --debug                   log HTTP requests/responses to stderr (Authorization redacted)
  -h, --help                    help for flexera-cli
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

* [flexera-cli auth](flexera-cli_auth.md)	 - Acquire OAuth tokens from Flexera login
* [flexera-cli bill-analysis](flexera-cli_bill-analysis.md)	 - Bill Analysis API
* [flexera-cli bill-upload](flexera-cli_bill-upload.md)	 - Bill Upload API
* [flexera-cli billing-center](flexera-cli_billing-center.md)	 - Billing Center API
* [flexera-cli budget](flexera-cli_budget.md)	 - Budget API
* [flexera-cli cli](flexera-cli_cli.md)	 - Discover CLI commands and API schemas
* [flexera-cli credential](flexera-cli_credential.md)	 - Credential API
* [flexera-cli data-inventory](flexera-cli_data-inventory.md)	 - Data Inventory API
* [flexera-cli finops-billing](flexera-cli_finops-billing.md)	 - FinOps Billing API
* [flexera-cli finops-customizations](flexera-cli_finops-customizations.md)	 - FinOps Customizations API
* [flexera-cli finops-onboarding](flexera-cli_finops-onboarding.md)	 - FinOps Onboarding API
* [flexera-cli graphql](flexera-cli_graphql.md)	 - GraphQL API
* [flexera-cli grs](flexera-cli_grs.md)	 - Global Resource Service API
* [flexera-cli iam](flexera-cli_iam.md)	 - Identity and Access Management API
* [flexera-cli it-visibility](flexera-cli_it-visibility.md)	 - IT Visibility Insights API
* [flexera-cli policy](flexera-cli_policy.md)	 - Policy API
* [flexera-cli recommendations](flexera-cli_recommendations.md)	 - Optima Recommendations API
* [flexera-cli risk](flexera-cli_risk.md)	 - Risk Management API
* [flexera-cli saas](flexera-cli_saas.md)	 - SaaS API
* [flexera-cli unified-onboarding](flexera-cli_unified-onboarding.md)	 - Unified Onboarding API

