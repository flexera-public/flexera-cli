## flexera-cli bill-connect-azure-mca update

Update an Azure MCA bill connect

```
flexera-cli bill-connect-azure-mca update [flags]
```

### Options

```
      --billing-account-id string     billingAccountId (body)
      --body string                   raw JSON body (inline | @file | @-); overrides body field flags
      --client-id string              clientId (body)
      --client-secret string          clientSecret (body)
      --dry-run                       print the planned operation as JSON and exit without calling the API
  -h, --help                          help for update
      --id string                     id (path, required)
      --start-billing-period string   startBillingPeriod (body)
      --tenant-id string              tenantId (body)
      --yes                           confirm the operation (required for destructive ops)
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

* [flexera-cli bill-connect-azure-mca](flexera-cli_bill-connect-azure-mca.md)	 - Bill Connect - Azure MCA operations (generated from the unified OpenAPI spec)

