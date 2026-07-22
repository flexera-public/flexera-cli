## flexera-cli managed-application delete

Delete managed application

```
flexera-cli managed-application delete [flags]
```

### Options

```
      --dry-run                 print the planned operation as JSON and exit without calling the API
  -h, --help                    help for delete
      --managed-app-id string   managedAppId (path, required)
      --yes                     confirm the operation (required for destructive ops)
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

* [flexera-cli managed-application](flexera-cli_managed-application.md)	 - Managed Application operations (generated from the unified OpenAPI spec)

