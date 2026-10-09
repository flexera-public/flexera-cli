## flexera-cli finops-billing billing list-rules

Index enterprise adjustment rules

### Synopsis

Index enterprise adjustment rules

Lists some or all enterprise adjustment rules in the organization.

```
flexera-cli finops-billing billing list-rules [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing billing list-rules --org-id ORG_ID
```

### Options

```
      --filter string       filter (query); Optional filter for returning enterprise adjustment rules matching specific criteria. Filter parameters may be combined with 'and' and 'or' logical operators. Each parameter may compare one of the following attributes to a given value or null. ### Enterprise Rule Filter Attributes - id, the unique identifier of a rule - name, the name of the rule - type, the... (see cli schema); illustrative example: "id eq 'd5f2f3eb-5f91-4f10-89de-40eaf3cf4a6f'"
  -h, --help                help for list-rules
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

