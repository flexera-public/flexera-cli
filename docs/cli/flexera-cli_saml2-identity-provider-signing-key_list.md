## flexera-cli saml2-identity-provider-signing-key list

Index an IdP's signing keys

```
flexera-cli saml2-identity-provider-signing-key list [flags]
```

### Options

```
  -h, --help                          help for list
      --identity-provider-id string   identityProviderId (path, required)
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

* [flexera-cli saml2-identity-provider-signing-key](flexera-cli_saml2-identity-provider-signing-key.md)	 - SAML2 Identity Provider Signing Key operations (generated from the unified OpenAPI spec)

