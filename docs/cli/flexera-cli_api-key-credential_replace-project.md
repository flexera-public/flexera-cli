## flexera-cli api-key-credential replace-project

Create a Credential

```
flexera-cli api-key-credential replace-project [flags]
```

### Options

```
      --body string          raw JSON body (inline | @file | @-); overrides body field flags
      --description string   description (body)
      --dry-run              print the planned operation as JSON and exit without calling the API
      --field string         field (body)
  -h, --help                 help for replace-project
      --id string            id (path, required)
      --key string           key (body)
      --location string      location (body)
      --name string          name (body)
      --project-id int       projectId (path, required)
      --type string          type (body)
      --yes                  confirm the operation (required for destructive ops)
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

* [flexera-cli api-key-credential](flexera-cli_api-key-credential.md)	 - API Key Credential operations (generated from the unified OpenAPI spec)

