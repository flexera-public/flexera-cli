## flexera-cli finops-billing billing-credits vendor-credits

Index vendor credits for a bill month

### Synopsis

Index vendor credits for a bill month

List vendor credits eligible for assignment.

Each row represents the total credit amount across a unique combination of the following fields:

- charge description
- type
- customer
- billing account
- sub-account
- provider

```
flexera-cli finops-billing billing-credits vendor-credits [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing billing-credits vendor-credits --org-id ORG_ID
```

### Options

```
      --filter string       filter (query); Optional filter for returning vendor credits matching specific criteria. Filter parameters may be combined with 'and' and 'or' logical operators. ### Vendor Credits Filter Attributes - billMonth, the YYYYMM string representing the billing month of a credit row - chargeDescription, the charge description of a credit row - creditType, the CCO line item type of... (see cli schema); illustrative example: "billMonth eq '202601' and provider eq 'AWS' and chargeDescription co 'EDP'"
  -h, --help                help for vendor-credits
      --limit int           limit (query); Return no more than limit values per page; maximum: 200; API default: 20; illustrative example: 20
      --no-paginate         return only the first page (do not follow nextPage)
      --order-by string     orderBy (query); Optional orderBy query allows to specify an expression for determining what values are used to order the entities. Multiple expressions can be specified using comma separated values.; illustrative example: "amount desc, lineItemCount desc"
      --skip-token string   resume pagination from this token; An opaque token to be provided when requesting a subsequent page after receiving a partial response. Partial responses will include a "nextPage" attribute in their response body, which contains the URL of the next page including the appropriate skipToken.
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

* [flexera-cli finops-billing billing-credits](flexera-cli_finops-billing_billing-credits.md)	 - Billing Credits operations (generated from the unified OpenAPI spec)

