## flexera-cli org-dashboards create

create org_dashboards

```
flexera-cli org-dashboards create [flags]
```

### Options

```
      --area string         area (body)
      --body string         raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run             print the planned operation as JSON and exit without calling the API
  -h, --help                help for create
      --name string         name (body)
      --scope strings       scope (body)
      --visibility string   visibility (body)
      --yes                 confirm the operation (required for destructive ops)
```

### Options inherited from parent commands

```
      --access-token string     static bearer access token
      --api-base-url string     override API base URL
      --client-id string        OAuth client ID
      --client-secret string    OAuth client secret
      --config string           config file (default $HOME/.flexera/config.yaml)
  -d, --debug                   log HTTP requests/responses to stderr (Authorization redacted)
      --login-base-url string   override login base URL
      --org-id int              organization ID
  -o, --output string           output format (json|table)
      --refresh-token string    OAuth refresh token
      --zone string             API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli org-dashboards](flexera-cli_org-dashboards.md)	 - org_dashboards operations (generated from the unified OpenAPI spec)

