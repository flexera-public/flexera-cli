## flexera-cli saas usage-group costs-all-2

Delete an existing usage cost for the specified Usage Cost object

### Synopsis

Delete an existing usage cost for the specified Usage Cost object

```
flexera-cli saas usage-group costs-all-2 [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas usage-group costs-all-2 --org-id ORG_ID --usage-group-id USAGE_GROUP_ID --cost-id COST_ID
  flexera-cli saas usage-group costs-all-2 --org-id ORG_ID --usage-group-id USAGE_GROUP_ID --cost-id COST_ID --dry-run
```

### Options

```
      --cost-id string          costId (path, required); Object id of the cost object.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "243412"
      --dry-run                 print the planned operation as JSON and exit without calling the API
  -h, --help                    help for costs-all-2
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

