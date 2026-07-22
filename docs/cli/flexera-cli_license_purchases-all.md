## flexera-cli license purchases-all

Create purchase

```
flexera-cli license purchases-all [flags]
```

### Options

```
      --amount float            amount (body)
      --body string             raw JSON body (inline | @file | @-); overrides body field flags
      --currency string         currency (body)
      --effective-at string     effectiveAt (body)
      --ends-at string          endsAt (body)
      --frequency-type string   frequencyType (body)
  -h, --help                    help for purchases-all
      --id string               id (body)
      --license-id string       licenseId (path, required)
      --purchase-id string      purchaseId (body)
      --purchased-at string     purchasedAt (body)
      --term-id string          termId (path, required)
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

* [flexera-cli license](flexera-cli_license.md)	 - License operations (generated from the unified OpenAPI spec)

