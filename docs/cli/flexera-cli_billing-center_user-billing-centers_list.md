## flexera-cli billing-center user-billing-centers list

List BillingCenters

### Synopsis

List BillingCenters

List highest BillingCenters for which the user has the 'optima:billing_center:show' privilege.

**Required security scopes for SameUser**:
  * `common:org:affiliated`

```
flexera-cli billing-center user-billing-centers list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli billing-center user-billing-centers list --org-id ORG_ID --user USER
```

### Options

```
  -h, --help                           help for list
      --highest-accessible-b-cs-only   highestAccessibleBCsOnly (query); Returns the highest level of billing centers accessible to the user when true and all levels of billing centers accessible to the user when false.; API default: true
      --user int                       user (path, required); required by API
      --view string                    view (query); Returns only the id and the name of the billing centers if set to "compact"; enum: ["index","compact"]; API default: "index"
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

* [flexera-cli billing-center user-billing-centers](flexera-cli_billing-center_user-billing-centers.md)	 - UserBillingCenters operations (generated from the unified OpenAPI spec)

