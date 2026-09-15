## flexera-cli billing-credits create

Create a credit assignment

```
flexera-cli billing-credits create [flags]
```

### Options

```
      --action string                        action (body)
      --bill-month string                    billMonth (body)
      --body string                          raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run                              print the planned operation as JSON and exit without calling the API
  -h, --help                                 help for create
      --name string                          name (body)
      --provider string                      provider (body)
      --target-billing-account-id string     targetBillingAccountId (body)
      --target-billing-account-name string   targetBillingAccountName (body)
      --target-customer-id int               targetCustomerId (body)
      --target-sub-account-id string         targetSubAccountId (body)
      --target-sub-account-name string       targetSubAccountName (body)
      --yes                                  confirm the operation (required for destructive ops)
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

* [flexera-cli billing-credits](flexera-cli_billing-credits.md)	 - Billing Credits operations (generated from the unified OpenAPI spec)

