## flexera-cli connector-status list

Get status for all enabled products for a given connector

```
flexera-cli connector-status list [flags]
```

### Options

```
      --account-id string     account_id (query)
      --connector-id string   connector_id (query)
  -h, --help                  help for list
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

* [flexera-cli connector-status](flexera-cli_connector-status.md)	 - ConnectorStatus operations (generated from the unified OpenAPI spec)

