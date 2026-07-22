## flexera-cli msp-customer-v2 create

Create a new MSP customer (v2)

```
flexera-cli msp-customer-v2 create [flags]
```

### Options

```
      --body string          raw JSON body (inline | @file | @-); overrides body field flags
      --description string   description (body)
      --dry-run              print the planned operation as JSON and exit without calling the API
      --external-id string   externalId (body)
  -h, --help                 help for create
      --name string          name (body)
      --yes                  confirm the operation (required for destructive ops)
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

