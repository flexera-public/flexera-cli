## flexera-cli saas metric-query user-counts-by-events-performed

User counts by events performed

### Synopsis

User counts by events performed

Retrieves the counts of users grouped into ranges by the number of events of each type they have performed against a given sub application on a managed application.

```
flexera-cli saas metric-query user-counts-by-events-performed [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas metric-query user-counts-by-events-performed --org-id ORG_ID --managed-app-id MANAGED_APP_ID --days-since DAYS_SINCE --sub-app-id SUB_APP_ID
```

### Options

```
      --days-since int          daysSince (query); The total number of days worth of events to be included within the returned counts.; required by API; enum: [90,180,365,730]; illustrative example: 90
  -h, --help                    help for user-counts-by-events-performed
      --managed-app-id string   managedAppId (query); The managed application unique identifier for which application events count to be returned.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "34556"
      --sub-app-id string       subAppId (query); The unique identifier of the sub application for which the events counts is requested.; required by API; illustrative example: "incident_management"
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

* [flexera-cli saas metric-query](flexera-cli_saas_metric-query.md)	 - Metric Query operations (generated from the unified OpenAPI spec)

