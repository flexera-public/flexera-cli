## flexera-cli onboarding replace

Onboarding: Update

```
flexera-cli onboarding replace [flags]
```

### Options

```
      --body string                     raw JSON body (inline | @file | @-); overrides body field flags
      --connector-id string             connector_id (path, required)
      --connector-name string           ConnectorName (body)
      --dry-run                         print the planned operation as JSON and exit without calling the API
      --exclude-region strings          ExcludeRegion (body)
  -h, --help                            help for replace
      --include-bpc string              IncludeBPC (body)
      --include-cost-and-usage string   IncludeCostAndUsage (body)
      --include-inventory string        IncludeInventory (body)
      --provider string                 provider (query)
      --subscription-id string          SubscriptionId (body)
      --tenant-id string                TenantId (body)
      --yes                             confirm the operation (required for destructive ops)
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

* [flexera-cli onboarding](flexera-cli_onboarding.md)	 - Onboarding operations (generated from the unified OpenAPI spec)

