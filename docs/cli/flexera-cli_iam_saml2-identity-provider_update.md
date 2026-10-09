## flexera-cli iam saml2-identity-provider update

Update an identity provider

### Synopsis

Update an identity provider

Update the given properties of an existing identity provider.

```
flexera-cli iam saml2-identity-provider update [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam saml2-identity-provider update --org-id ORG_ID --id ID --body @request.json
  flexera-cli iam saml2-identity-provider update --org-id ORG_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema iam saml2-identity-provider update --example > request.json
```

### Options

```
      --body string                              raw JSON body (inline | @file | @-); overrides body field flags
      --certificate-public-key string            certificatePublicKey (body); The PEM or DER encoded public key certificate of the Identity Provider used to verify SAML message and assertion signatures.; pattern: "^-----BEGIN CERTIFICATE-----\\r?\\n([0-9a-zA-Z+/]+\\r?\\n)*([0-9a-zA-Z+/]+={0,3}\\r?\\n)-----END CERTIFICATE-----\\r?\\n?$"; illustrative example: "-----BEGIN CERTIFICATE-----\nAb12Cd/Ef+Gh\n0123456789==\n-----END CERTIFICATE-----"
      --discovery-hint string                    discoveryHint (body); String that users may enter to allow them to authenticate with this identity provider. The company's domain name is a good choice. The hint is optional but must be unique if provided.; maxLength: 255; pattern: "^[:()_+-.@a-zA-Z0-9 ]*$"; illustrative example: "my.company.domain.com"
      --dry-run                                  print the planned operation as JSON and exit without calling the API
      --group-sync-policy string                 groupSyncPolicy (body); Action to take during JIT flow, for groups configured for the identity provider, if JIT provisioning is enabled. * NONE - Group memberships are not modified. (default) * APPEND - Only adds the user to the specified groups, if any. Only valid if JIT Provisioning is enabled. * SYNC - Performs a full sync. If groups are provided by the identity provider, that g... (see cli schema); enum: ["NONE","APPEND","SYNC"]; illustrative example: "APPEND"
  -h, --help                                     help for update
      --id string                                id (path, required); ID of the identity provider; required by API; pattern: "^[0-9a-fA-F]{24}$"; illustrative example: "1111aaaa2222bbbb3333cccc"
  -i, --interactive                              edit inputs in a terminal form, review a plan and approve with typed yes
      --issuer-uri string                        issuerUri (body); This value is usually the SAML Metadata EntityID of the IdP; minLength: 1; maxLength: 2000; illustrative example: "mysamlprovider.com/entities/123"
      --jit-provisioning-enabled string          jitProvisioningEnabled (body); If enabled, performs JIT actions when a user logs in; enum: ["false","true"]; illustrative example: "false"
      --logout-redirect-url string               logoutRedirectUrl (body); The URL to redirect the user to upon logout, for example, the organization's Identity Provider page. By default, upon logout, the user is redirected to the Flexera One login page.; maxLength: 2000; illustrative example: "https://my.idp.com"
      --name string                              name (body); Display name for this identity provider; minLength: 1; maxLength: 512; illustrative example: "My trusted IDP"
      --request-binding string                   requestBinding (body); The SAML Authentication Request Protocol binding used to send SAML AuthnRequest messages to the IdP.; enum: ["HTTP-POST","HTTP-REDIRECT"]; illustrative example: "HTTP-POST"
      --request-signature-algorithm string       requestSignatureAlgorithm (body); Specifies the signature algorithm used to sign SAML AuthnRequest messages sent to the IdP.; enum: ["SHA-1","SHA-256"]; illustrative example: "SHA-256"
      --request-signing-key-id string            requestSigningKeyId (body); ID for last known active signing key.; pattern: "^[0-9a-fA-F]{24}$"; illustrative example: "2222bbbb3333cccc4444dddd"
      --response-signature-algorithm string      responseSignatureAlgorithm (body); Specifies the minimum signature algorithm when validating SAML assertions issued by the IdP.; enum: ["SHA-1","SHA-256"]; illustrative example: "SHA-256"
      --response-signature-verification string   responseSignatureVerification (body); The protocol to use when authenticating users from this identity provider; enum: ["response","assertion","response or assertion"]; illustrative example: "response"
      --sign-authn-requests string               signAuthnRequests (body); Specifies whether to sign SAML 2 AuthnRequest messages.; enum: ["false","true"]; illustrative example: "true"
      --sso-url string                           ssoUrl (body); The binding-specific IdP Authentication Request Protocol endpoint that receives SAML AuthnRequest messages.; minLength: 1; maxLength: 2000; illustrative example: "mysamlprovider.com/apps/456"
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
      --json-style string       JSON whitespace style (auto|pretty|compact) (default "auto")
      --login-base-url string   override login base URL
      --no-validate             skip API schema constraints (never JSON syntax or request data-loss checks)
      --org-id int              organization ID
      --out-fields string       project JSON output fields (comma-separated paths)
      --out-jq string           shape JSON output with a jq expression
  -o, --output string           output format (json|table)
  -r, --raw-output              write jq string results without JSON quotes (requires --out-jq)
      --refresh-token string    OAuth refresh token
      --zone string             API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli iam saml2-identity-provider](flexera-cli_iam_saml2-identity-provider.md)	 - SAML2 Identity Provider operations (generated from the unified OpenAPI spec)

