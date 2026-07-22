## flexera-cli msp-customer-tag list-all

Index an MSP's customers based on tag filter

```
flexera-cli msp-customer-tag list-all [flags]
```

### Options

```
      --filter string       filter (query)
  -h, --help                help for list-all
      --no-paginate         return only the first page (do not follow nextPage)
      --skip-token string   resume pagination from this token
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

