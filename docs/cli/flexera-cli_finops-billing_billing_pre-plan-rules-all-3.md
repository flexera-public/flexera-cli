## flexera-cli finops-billing billing pre-plan-rules-all-3

Index adjustment rules that run before all plans

### Synopsis

Index adjustment rules that run before all plans

Lists some or all adjustment rules in the organization that run before all plans

```
flexera-cli finops-billing billing pre-plan-rules-all-3 [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing billing pre-plan-rules-all-3 --org-id ORG_ID
```

### Options

```
      --filter string       filter (query); Optional filter for returning adjustments matching specific criteria. Filter parameters may be combined with 'and' and 'or' logical operators. Each parameter may compare one of the following attributes to a given value or null. ### Adjustment Filter Attributes - id, the unique identifier of an adjustment - name, the name of the adjustment - type, the type of... (see cli schema); illustrative example: "customerIds co 1234"
  -h, --help                help for pre-plan-rules-all-3
      --limit int           limit (query); Return no more than limit values per page; maximum: 100; API default: 20; illustrative example: 20
      --no-paginate         return only the first page (do not follow nextPage)
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

