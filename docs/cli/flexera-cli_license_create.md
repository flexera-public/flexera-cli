## flexera-cli license create

Create license

```
flexera-cli license create [flags]
```

### Options

```
      --body string                     raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run                         print the planned operation as JSON and exit without calling the API
  -h, --help                            help for create
      --managed-app-id string           managedAppId (body)
      --name string                     name (body)
      --point-of-contact-email string   pointOfContactEmail (body)
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

* [flexera-cli license](flexera-cli_license.md)	 - License operations (generated from the unified OpenAPI spec)

