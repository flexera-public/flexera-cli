## flexera-cli recommendations list

List all recommendations

### Synopsis

List all recommendations

List all recommendations.
User must have the 'optima:recommendation:index' privilege on the org or on the specified billing center(s) to make this call.

**Required security scopes for GlobalSession**:
  * `common:org:affiliated+common:org:own`

```
flexera-cli recommendations list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli recommendations list --org-id ORG_ID
```

### Options

```
      --billing-center-i-ds strings   billingCenterIDs (query); IDs of BillingCenters to get recommendations for. It is not allowed for any of the BillingCenterIDs to be an ancestor of another specified BillingCenterID.; CLI: comma-separated values or repeated flag
  -h, --help                          help for list
      --statuses strings              statuses (query); Recommendation statuses to get; API default: ["active"]; CLI: comma-separated values or repeated flag; items.enum: ["active","snoozed","rejected","realized"]
      --view details                  view (query); An optional parameter that controls the level of detail returned about policy violation data. Depending on it the data is returned in the details or `detailsExtended` field; enum: ["default","extended"]; API default: "default"
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

* [flexera-cli recommendations](flexera-cli_recommendations.md)	 - Optima Recommendations API

