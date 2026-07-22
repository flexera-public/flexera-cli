## flexera-cli saml2-identity-provider create

Create an identity provider

```
flexera-cli saml2-identity-provider create [flags]
```

### Options

```
      --body string                              raw JSON body (inline | @file | @-); overrides body field flags
      --certificate-public-key string            certificatePublicKey (body)
      --discovery-hint string                    discoveryHint (body)
      --dry-run                                  print the planned operation as JSON and exit without calling the API
      --group-sync-policy string                 groupSyncPolicy (body)
  -h, --help                                     help for create
      --issuer-uri string                        issuerUri (body)
      --jit-provisioning-enabled string          jitProvisioningEnabled (body)
      --logout-redirect-url string               logoutRedirectUrl (body)
      --name string                              name (body)
      --request-binding string                   requestBinding (body)
      --request-signature-algorithm string       requestSignatureAlgorithm (body)
      --response-signature-algorithm string      responseSignatureAlgorithm (body)
      --response-signature-verification string   responseSignatureVerification (body)
      --sign-authn-requests string               signAuthnRequests (body)
      --sso-url string                           ssoUrl (body)
      --yes                                      confirm the operation (required for destructive ops)
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

* [flexera-cli saml2-identity-provider](flexera-cli_saml2-identity-provider.md)	 - SAML2 Identity Provider operations (generated from the unified OpenAPI spec)

