## flexera-cli finops-billing billing-orphans charges

List unassigned Azure orphaned charges

### Synopsis

List unassigned Azure orphaned charges

Query the orphaned-costs inventory and cross-reference stored assignments before filtering, counting, and pagination.
Eligible rows have cco_line_item_type = 'Purchase' and an originally null or empty
native SubscriptionId. chargeDescription is the exact FOCUS ChargeDescription,
not cco_line_item_type or ProductName. The line-item type is only an eligibility
filter, never an assignment selector.
Exclude matching recurring assignments regardless of discovery month, and exclude
one-time assignments only for the matching month. Match exact description and
frequency within organization, bill source, and billing account. Exclusion takes
effect on save, even before normalization applies the mapping. Clients cannot
disable this exclusion. Fail the request if assignment lookup is unavailable.

Omitting billSource queries all sources within the organization, month, and billing
account. Creating mappings requires a specific billSource and validates the current
eligible rows for that source and the requested month.

Display rows may be split by currency, billing profile, product order, and customer.
Selecting a description/frequency selects all eligible matching rows in scope,
including rows hidden by display filters. Never sum different currencies.
Orderable/filterable fields are chargeDescription, chargeFrequency, billedCost,
billingProfileId, productOrderId, commitmentDiscountType, and customerId.
Pagination uses offsets over the current inventory; changes between pages can
cause rows to be skipped or repeated.

```
flexera-cli finops-billing billing-orphans charges [flags]
```
### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing billing-orphans charges --org-id ORG_ID --bill-month BILL_MONTH --billing-account-id BILLING_ACCOUNT_ID
```

### Options

```
      --bill-month string           billMonth (query); Discovery month or one-time assignment month in YYYYMM format. Omitted on recurring assignment resources.; required by API; pattern: "^\\d{4}(0[1-9]|1[0-2])$"; illustrative example: "202610"
      --bill-source string          billSource (query); Opaque source identifier returned by orphaned billing-account discovery.; minLength: 1; illustrative example: "azure-csp-source-a"
      --billing-account-id string   billingAccountId (query); Azure billing account identifier.; required by API; minLength: 1; illustrative example: "billing-account-a"
      --filter string               filter (query); Filter eligible unassigned charge rows using eq, co, in, and logical and/or.; illustrative example: "chargeFrequency eq 'recurring'"
  -h, --help[=true|false]           help for charges
      --limit int                   limit (query); Return no more than limit values per page; maximum: 100; API default: 50; illustrative example: 50
      --no-paginate[=true|false]    return only the first page (do not follow nextPage)
      --order-by string             orderBy (query); Optional orderBy query allows to specify an expression for determining what values are used to order the entities. Multiple expressions can be specified using comma separated values.; illustrative example: "chargeDescription asc"
      --skip-token string           resume pagination from this token; An opaque token to be provided when requesting a subsequent page after receiving a partial response. Partial responses will include a "nextPage" attribute in their response body, which contains the URL of the next page including the appropriate skipToken.
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

