## flexera-cli finops-customizations tag-observation list

List tag observations

### Synopsis

List tag observations

Lists observed tag keys, optionally filtered by provider, source type, or tag-key fields using the filter query parameter. Returns a snapshot-consistent page of matching observations.

```
flexera-cli finops-customizations tag-observation list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-customizations tag-observation list --org-id ORG_ID
```

### Options

```
      --filter string       filter (query); Optional filter for Tag Observations. The following filters are supported: | Filter | Allowed Operators | Example | | --- | --- | --- | | provider | eq ne in | provider in ['aws', 'azure'] | | sourceType | eq ne in | sourceType eq 'awsCostAllocationTag' | | qualifiedKey | co eq ne in | qualifiedKey co 'user:Customer' | | key | co eq ne in | key co 'Customer'... (see cli schema); maxLength: 8192; illustrative example: "provider in ['aws', 'azure'] and (qualifiedKey co 'customer' or key co 'customer' or normalizedKey co 'customer')"
  -h, --help                help for list
      --limit int           limit (query); Page size (default 50, max 200).; minimum: 1; maximum: 200; API default: 50; illustrative example: 15
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

* [flexera-cli finops-customizations tag-observation](flexera-cli_finops-customizations_tag-observation.md)	 - Tag Observation operations (generated from the unified OpenAPI spec)

