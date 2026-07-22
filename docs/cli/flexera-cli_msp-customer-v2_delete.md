## flexera-cli msp-customer-v2 delete

Delete an MSP's customer (v2)

```
flexera-cli msp-customer-v2 delete [flags]
```

### Options

```
      --customer-id int   customerId (path, required)
      --dry-run           print the planned operation as JSON and exit without calling the API
  -h, --help              help for delete
      --yes               confirm the operation (required for destructive ops)
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

* [flexera-cli msp-customer-v2](flexera-cli_msp-customer-v2.md)	 - MSP Customer V2 operations (generated from the unified OpenAPI spec)

