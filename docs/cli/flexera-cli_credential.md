## flexera-cli credential

Credential API

### Synopsis

Commands for the Credential API (Flexera Credential API, v2).

Service id: cred

### Options

```
  -h, --help   help for credential
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

* [flexera-cli](flexera-cli.md)	 - Flexera One unified API command-line client
* [flexera-cli credential api-key-credential](flexera-cli_credential_api-key-credential.md)	 - API Key Credential operations (generated from the unified OpenAPI spec)
* [flexera-cli credential aws-credential](flexera-cli_credential_aws-credential.md)	 - AWS Credential operations (generated from the unified OpenAPI spec)
* [flexera-cli credential awssts-credential](flexera-cli_credential_awssts-credential.md)	 - AWS STS Credential operations (generated from the unified OpenAPI spec)
* [flexera-cli credential basic-credential](flexera-cli_credential_basic-credential.md)	 - Basic Credential operations (generated from the unified OpenAPI spec)
* [flexera-cli credential delete](flexera-cli_credential_delete.md)	 - Delete a Credential
* [flexera-cli credential delete-project](flexera-cli_credential_delete-project.md)	 - Delete a Credential
* [flexera-cli credential digest-credential](flexera-cli_credential_digest-credential.md)	 - Digest Credential operations (generated from the unified OpenAPI spec)
* [flexera-cli credential list](flexera-cli_credential_list.md)	 - Index a list of Credentials
* [flexera-cli credential list-project](flexera-cli_credential_list-project.md)	 - Index a list of Credentials
* [flexera-cli credential ntlm-credential](flexera-cli_credential_ntlm-credential.md)	 - NTLM Credential operations (generated from the unified OpenAPI spec)
* [flexera-cli credential o-auth2-credential](flexera-cli_credential_o-auth2-credential.md)	 - OAuth2 Credential operations (generated from the unified OpenAPI spec)
* [flexera-cli credential oracle-credential](flexera-cli_credential_oracle-credential.md)	 - Oracle Credential operations (generated from the unified OpenAPI spec)

