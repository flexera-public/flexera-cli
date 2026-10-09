## flexera-cli saas usage-group costs

Update an existing cost for the specified usage group

### Synopsis

Update an existing cost for the specified usage group

Update an existing cost object for the specified Usage group object

```
flexera-cli saas usage-group costs [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas usage-group costs --org-id ORG_ID --usage-group-id USAGE_GROUP_ID --cost-id COST_ID --body @request.json
  flexera-cli saas usage-group costs --org-id ORG_ID --usage-group-id USAGE_GROUP_ID --cost-id COST_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema saas usage-group costs --example > request.json
```

### Options

```
      --body string             raw JSON body (inline | @file | @-); overrides body field flags
      --cost-id string          costId (path, required); Usage Cost Id; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "243412"
      --dry-run                 print the planned operation as JSON and exit without calling the API
      --ends-at string          endsAt (body); End date of the consumption period; format: date-time; illustrative example: "2022-02-01T00:00:00.000Z"
  -h, --help                    help for costs
  -i, --interactive             edit inputs in a terminal form, review a plan and approve with typed yes
      --overage-cost float      overageCost (body); Cost per consumption incurred after exceeding the total within the date range defined; format: double; illustrative example: 0.00075
      --starts-at string        startsAt (body); Start date of the consumption period; format: date-time; illustrative example: "2021-02-01T00:00:00.000Z"
      --total-purchased int     totalPurchased (body); Total amount of consumption allocated; format: int64; illustrative example: 15000
      --usage-group-id string   usageGroupID (path, required); ID of the usage group to which cost is associated; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "1105"
      --yes                     confirm the operation (required for destructive ops)
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

* [flexera-cli saas usage-group](flexera-cli_saas_usage-group.md)	 - Usage Group operations (generated from the unified OpenAPI spec)

