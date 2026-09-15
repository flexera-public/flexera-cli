## flexera-cli recommendations list

List all recommendations

```
flexera-cli recommendations list [flags]
```

### Options

```
      --billing-center-i-ds strings   billingCenterIDs (query)
  -h, --help                          help for list
      --org-id int                    orgID (path, required)
      --statuses strings              statuses (query)
      --view string                   view (query)
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
  -o, --output string           output format (json|table)
      --refresh-token string    OAuth refresh token
      --zone string             API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli recommendations](flexera-cli_recommendations.md)	 - Recommendations operations (generated from the unified OpenAPI spec)

