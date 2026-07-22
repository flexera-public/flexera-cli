## flexera-cli basic-credential replace-project

Create a Credential

```
flexera-cli basic-credential replace-project [flags]
```

### Options

```
      --body string          raw JSON body (inline | @file | @-); overrides body field flags
      --description string   description (body)
      --dry-run              print the planned operation as JSON and exit without calling the API
  -h, --help                 help for replace-project
      --id string            id (path, required)
      --name string          name (body)
      --password string      password (body)
      --project-id int       projectId (path, required)
      --username string      username (body)
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

* [flexera-cli basic-credential](flexera-cli_basic-credential.md)	 - Basic Credential operations (generated from the unified OpenAPI spec)

