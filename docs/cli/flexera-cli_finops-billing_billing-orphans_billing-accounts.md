## flexera-cli finops-billing billing-orphans billing-accounts

List Azure billing accounts with orphaned charges

### Synopsis

List Azure billing accounts with orphaned charges

Return the complete, unpaginated list of bill source and billing account pairs with eligible orphaned charges in the requested month.
Accounts remain visible even when all their eligible charges have stored assignments.
Each entry contains only billSource, billingAccountId, and the billing account name when available.

```
flexera-cli finops-billing billing-orphans billing-accounts [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing billing-orphans billing-accounts --org-id ORG_ID --bill-month BILL_MONTH
```
### Options

```
      --bill-month string    billMonth (query); Discovery month or one-time assignment month in YYYYMM format. Omitted on recurring assignment resources.; required by API; pattern: "^\\d{4}(0[1-9]|1[0-2])$"; illustrative example: "202610"
      --bill-source string   billSource (query); Opaque source identifier returned by orphaned billing-account discovery.; minLength: 1; illustrative example: "azure-csp-source-a"
      --filter string        filter (query); Filter on billingAccountId or billingAccountName using eq, co, and logical and/or.; illustrative example: "billingAccountId eq 'billing-account-a'"
  -h, --help[=true|false]    help for billing-accounts
      --order-by string      orderBy (query); Optional orderBy query allows to specify an expression for determining what values are used to order the entities. Multiple expressions can be specified using comma separated values.; illustrative example: "billingAccountId asc, billSource asc"
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

* [flexera-cli finops-billing billing-orphans](flexera-cli_finops-billing_billing-orphans.md)	 - Billing Orphans operations (generated from the unified OpenAPI spec)

