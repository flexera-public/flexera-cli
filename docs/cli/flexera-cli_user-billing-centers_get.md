## flexera-cli user-billing-centers get

List BillingCenters with privileges.

```
flexera-cli user-billing-centers get [flags]
```

### Options

```
      --billing-center string   billing_center (path, required)
  -h, --help                    help for get
      --user int                user (path, required)
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

* [flexera-cli user-billing-centers](flexera-cli_user-billing-centers.md)	 - UserBillingCenters operations (generated from the unified OpenAPI spec)

