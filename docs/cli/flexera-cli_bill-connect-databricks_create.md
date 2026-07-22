## flexera-cli bill-connect-databricks create

Create a Databricks bill connect

```
flexera-cli bill-connect-databricks create [flags]
```

### Options

```
      --account-id string         accountId (body)
      --body string               raw JSON body (inline | @file | @-); overrides body field flags
      --client-id string          clientId (body)
      --client-secret string      clientSecret (body)
      --dry-run                   print the planned operation as JSON and exit without calling the API
  -h, --help                      help for create
      --sql-warehouse-id string   sqlWarehouseId (body)
      --workspace-url string      workspaceUrl (body)
      --yes                       confirm the operation (required for destructive ops)
```

### Options inherited from parent commands

```
      --access-token string     static bearer access token
      --api-base-url string     override API base URL
      --config string           config file (default $HOME/.flexera/config.yaml)
  -d, --debug                   log HTTP requests/responses to stderr (Authorization redacted)
      --login-base-url string   override login base URL
      --org-id int              organization ID
  -o, --output string           output format (json|table)
      --refresh-token string    OAuth refresh token
      --zone string             API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli bill-connect-databricks](flexera-cli_bill-connect-databricks.md)	 - Bill Connect - Databricks operations (generated from the unified OpenAPI spec)

