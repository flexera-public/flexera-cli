## flexera-cli saml2-domain delete

Delete an IdP's domain

```
flexera-cli saml2-domain delete [flags]
```

### Options

```
      --dry-run                       print the planned operation as JSON and exit without calling the API
  -h, --help                          help for delete
      --identity-provider-id string   identityProviderId (path, required)
      --name string                   name (path, required)
      --yes                           confirm the operation (required for destructive ops)
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

* [flexera-cli saml2-domain](flexera-cli_saml2-domain.md)	 - SAML2 Domain operations (generated from the unified OpenAPI spec)

