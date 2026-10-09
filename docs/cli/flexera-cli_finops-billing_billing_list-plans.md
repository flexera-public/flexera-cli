## flexera-cli finops-billing billing list-plans

Index adjustment plan rules

### Synopsis

Index adjustment plan rules

Lists some or all adjustment plan rules in an adjustment plan.

```
flexera-cli finops-billing billing list-plans [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing billing list-plans --org-id ORG_ID --plan-id PLAN_ID
```

### Options

```
      --filter string       filter (query); Optional filter for returning rules matching specific criteria. Filter parameters may be combined with 'and' and 'or' logical operators. Each parameter may compare one of the following attributes to a given value or null. ### Rule Filter Attributes - id, the unique identifier of a rule - name, the type of the rule - type, the type of the rule - enabled, whet... (see cli schema); illustrative example: "id eq 'ae85f96e-6b9e-4d68-a741-d1ea7ca1fb28'"
  -h, --help                help for list-plans
      --limit int           limit (query); Return no more than limit values per page; maximum: 50; API default: 20; illustrative example: 20
      --no-paginate         return only the first page (do not follow nextPage)
      --plan-id string      planId (path, required); Unique identifier for the plan this rule is part of.; required by API; format: uuid; illustrative example: "c6671c74-513a-4127-b8df-81adfb65bf7f"
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

* [flexera-cli finops-billing billing](flexera-cli_finops-billing_billing.md)	 - Billing operations (generated from the unified OpenAPI spec)

