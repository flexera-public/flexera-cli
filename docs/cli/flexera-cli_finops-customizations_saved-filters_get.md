## flexera-cli finops-customizations saved-filters get

Show a saved filter

### Synopsis

Show a saved filter

Shows a specified saved filter in the organization. Billing Center entitlement is applied before the filter is returned: billingCenterIds contains only Billing Centers the caller is entitled to see, Billing Center Level filter values the caller cannot see are removed from filter, and Billing Center group-by dimensions are omitted when the caller has no accessible Billing Center. Nothing in the response indicates that values were omitted. The request fails if the caller's Billing Center entitlement cannot be verified.

```
flexera-cli finops-customizations saved-filters get [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-customizations saved-filters get --org-id ORG_ID --id ID
```

### Options

```
  -h, --help        help for get
      --id string   id (path, required); Identifier of the saved filter (UUID); required by API; format: uuid; illustrative example: "550e8400-e29b-41d4-a716-446655440000"
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

* [flexera-cli finops-customizations saved-filters](flexera-cli_finops-customizations_saved-filters.md)	 - Saved Filters operations (generated from the unified OpenAPI spec)

