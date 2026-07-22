## flexera-cli msp-customer-tag tags

Sets a tag on an MSP's customer

```
flexera-cli msp-customer-tag tags [flags]
```

### Options

```
      --body string       raw JSON body (inline | @file | @-); overrides body field flags
      --customer-id int   customerId (path, required)
  -h, --help              help for tags
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

* [flexera-cli msp-customer-tag](flexera-cli_msp-customer-tag.md)	 - MSP Customer Tag operations (generated from the unified OpenAPI spec)

