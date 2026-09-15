## flexera-cli recommendations replace

updateStatus Recommendations

```
flexera-cli recommendations replace [flags]
```

### Options

```
      --body string                  raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run                      print the planned operation as JSON and exit without calling the API
  -h, --help                         help for replace
      --id string                    id (body)
      --org-id int                   orgID (path, required)
      --snoozed-target-date string   snoozedTargetDate (body)
      --status string                status (body)
      --status-reason string         statusReason (body)
      --yes                          confirm the operation (required for destructive ops)
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
  -o, --output string           output format (json|table)
      --refresh-token string    OAuth refresh token
      --zone string             API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli recommendations](flexera-cli_recommendations.md)	 - Recommendations operations (generated from the unified OpenAPI spec)

