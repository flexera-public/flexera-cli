## flexera-cli iam saml2-domain get

Show an IdP's domain

### Synopsis

Show an IdP's domain

Show an existing domain belonging to an IDP.

```
flexera-cli iam saml2-domain get [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam saml2-domain get --org-id ORG_ID --identity-provider-id IDENTITY_PROVIDER_ID --name NAME
```

### Options

```
  -h, --help                          help for get
      --identity-provider-id string   identityProviderId (path, required); ID for the identity provider; required by API; pattern: "^[0-9a-fA-F]{24}$"; illustrative example: "1111aaaa2222bbbb3333cccc"
      --name string                   name (path, required); Name of the domain. See also "RFC 1035".; required by API; maxLength: 255; pattern: "^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\\.)+[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9]$"; illustrative example: "flexera.com"
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

* [flexera-cli iam saml2-domain](flexera-cli_iam_saml2-domain.md)	 - SAML2 Domain operations (generated from the unified OpenAPI spec)

