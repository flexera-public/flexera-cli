## flexera-cli iam msp-customer-tag list-all

Index an MSP's customers based on tag filter

### Synopsis

Index an MSP's customers based on tag filter

Index a managed service provider's customers based on tag(s).

```
flexera-cli iam msp-customer-tag list-all [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam msp-customer-tag list-all --org-id ORG_ID
```

### Options

```
      --filter string       filter (query); Tags filter for filtering list of customers returned. The following filters are supported: | Filter | Description | Allowed Operator | Behavior | Example | | --- | ---| --- | --- | --- | | tags | Filters on the tag(s) | co | Contains - The entire operator value must be a substring of the attribute value for a match. | tags co 'SAP' | | | | eq | Equal - The a... (see cli schema); minLength: 1; illustrative example: "(name co 'SAP' or name co 'HP')"
  -h, --help                help for list-all
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

* [flexera-cli iam msp-customer-tag](flexera-cli_iam_msp-customer-tag.md)	 - MSP Customer Tag operations (generated from the unified OpenAPI spec)

