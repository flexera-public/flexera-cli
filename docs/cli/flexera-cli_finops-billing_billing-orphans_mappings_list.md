## flexera-cli finops-billing billing-orphans mappings list

List Azure cost assignments

### Synopsis

List Azure cost assignments

List individual recurring assignments and one-time assignments for the requested month.
Return a flat list of matching assignments without grouping by destination.
Filterable fields are chargeDescription, chargeFrequency, targetSubAccountId,
billingAccountId, and billSource. Order assignments by targetSubAccountId.
Pagination applies to individual assignments. totalCount counts all matching
assignments before pagination. Return the persisted metadata directly on each assignment, including
billing amounts and customer IDs as observed at creation, not live billing data.

```
flexera-cli finops-billing billing-orphans mappings list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing billing-orphans mappings list --org-id ORG_ID --bill-month BILL_MONTH
```

### Options

```
      --bill-month string          billMonth (query); Discovery month or one-time assignment month in YYYYMM format. Omitted on recurring assignment resources.; required by API; pattern: "^\\d{4}(0[1-9]|1[0-2])$"; illustrative example: "202610"
      --filter string              filter (query); Filter stored assignments using eq, co, in, and logical and/or.; illustrative example: "chargeFrequency eq 'recurring'"
  -h, --help[=true|false]          help for list
      --limit int                  limit (query); Return no more than limit values per page; maximum: 100; API default: 50; illustrative example: 50
      --no-paginate[=true|false]   return only the first page (do not follow nextPage)
      --order-by string            orderBy (query); Optional orderBy query allows to specify an expression for determining what values are used to order the entities. Multiple expressions can be specified using comma separated values.; illustrative example: "targetSubAccountId asc"
      --skip-token string          resume pagination from this token; An opaque token to be provided when requesting a subsequent page after receiving a partial response. Partial responses will include a "nextPage" attribute in their response body, which contains the URL of the next page including the appropriate skipToken.
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

* [flexera-cli finops-billing billing-orphans mappings](flexera-cli_finops-billing_billing-orphans_mappings.md)	 - mappings operations

