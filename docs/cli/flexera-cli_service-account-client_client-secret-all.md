## flexera-cli service-account-client client-secret-all

Rotate a service account client secret

```
flexera-cli service-account-client client-secret-all [flags]
```

### Options

```
      --client-id string         clientId (path, required)
  -h, --help                     help for client-secret-all
      --service-account-id int   serviceAccountId (path, required)
```

### Options inherited from parent commands

```
      --access-token string     static bearer access token
      --api-base-url string     override API base URL
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

* [flexera-cli service-account-client](flexera-cli_service-account-client.md)	 - Service Account Client operations (generated from the unified OpenAPI spec)

