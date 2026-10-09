## flexera-cli saas usage-message-query list

Retrieve usage for the given period of time and for given usage group.

### Synopsis

Retrieve usage for the given period of time and for given usage group.

Retrieves usage for the given period of time and usage group.

```
flexera-cli saas usage-message-query list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas usage-message-query list --org-id ORG_ID
```

### Options

```
      --ended-at string     endedAt (query, RFC3339 date-time with timezone); End date of the usage consumption period; format: date-time; illustrative example: "2022-06-28T00:00:00.000Z"
      --filter string       filter (query); The date filter to query the usage groups. Supported fields in the filter are [managedAppId, sourceId, appUser, usageGroup] | Attribute | Description | Allowed Operators | Example | |---------------|-------------------------------------------------------------|-------------------|----------------------------| | managedAppId | Filter usage consumption by thei... (see cli schema); illustrative example: "Rerum ut saepe."
      --group-by strings    groupBy (query); List of fields to group results by.; CLI: comma-separated values or repeated flag; illustrative example: ["usageGroups","usage"]
  -h, --help                help for list
      --no-paginate         return only the first page (do not follow nextPage)
      --resolution string   resolution (query); resolution specifies the granularity to query results. The boundaries align to the timestamp representation.; enum: ["all","none","second","minute","fifteen_minute","thirty_minute","hour","day","week","month","quarter","year"]; illustrative example: "day"
      --skip-token string   resume pagination from this token; An opaque token to be provided when requesting a subsequent page after receiving a partial response. Partial responses will include a "nextPage" attribute in their response body, which contains the URL of the next page including the appropriate skipToken.
      --started-at string   startedAt (query, RFC3339 date-time with timezone); Start date of the usage consumption period; format: date-time; illustrative example: "2021-06-28T00:00:00.000Z"
      --view string         view (query); View used to render usage group data.; enum: ["default","extended"]; illustrative example: "default"
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

* [flexera-cli saas usage-message-query](flexera-cli_saas_usage-message-query.md)	 - Usage Message Query operations (generated from the unified OpenAPI spec)

