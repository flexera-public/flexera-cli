## flexera-cli policy-manager create

Create a Policy Manager

```
flexera-cli policy-manager create [flags]
```

### Options

```
      --allow-delete-policy      allowDeletePolicy (body)
      --allow-edit-policy        allowEditPolicy (body)
      --body string              raw JSON body (inline | @file | @-); overrides body field flags
      --description string       description (body)
      --dry-run                  print the planned operation as JSON and exit without calling the API
  -h, --help                     help for create
      --log-level string         logLevel (body)
      --name string              name (body)
      --org-tags-filter string   orgTagsFilter (body)
      --severity string          severity (body)
      --skip-approvals           skipApprovals (body)
      --template-ref string      templateRef (body)
      --yes                      confirm the operation (required for destructive ops)
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

* [flexera-cli policy-manager](flexera-cli_policy-manager.md)	 - Policy Manager operations (generated from the unified OpenAPI spec)

