## flexera-cli msp-customer-domain verify

Verify an MSP's customer's IdP's domain

```
flexera-cli msp-customer-domain verify [flags]
```

### Options

```
      --customer-id int               customerId (path, required)
  -h, --help                          help for verify
      --identity-provider-id string   identityProviderId (path, required)
      --name string                   name (path, required)
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

* [flexera-cli msp-customer-domain](flexera-cli_msp-customer-domain.md)	 - MSP Customer Domain operations (generated from the unified OpenAPI spec)

