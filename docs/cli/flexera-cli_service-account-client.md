## flexera-cli service-account-client

Service Account Client operations (generated from the unified OpenAPI spec)

### Options

```
  -h, --help   help for service-account-client
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
* [flexera-cli service-account-client client-secret](flexera-cli_service-account-client_client-secret.md)	 - Delete old service account client secret
* [flexera-cli service-account-client client-secret-all](flexera-cli_service-account-client_client-secret-all.md)	 - Rotate a service account client secret
* [flexera-cli service-account-client client-secrets](flexera-cli_service-account-client_client-secrets.md)	 - Index a service account client's secrets
* [flexera-cli service-account-client clients](flexera-cli_service-account-client_clients.md)	 - Create a service account client
* [flexera-cli service-account-client delete](flexera-cli_service-account-client_delete.md)	 - Delete a service account client
* [flexera-cli service-account-client get](flexera-cli_service-account-client_get.md)	 - Show service account client details
* [flexera-cli service-account-client list](flexera-cli_service-account-client_list.md)	 - Index a service account's clients

