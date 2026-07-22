## flexera-cli policy applied-policy evaluate

Request evaluation of an applied policy

```
flexera-cli policy applied-policy evaluate [flags]
```

### Options

```
  -h, --help             help for evaluate
      --id string        Applied policy ID
      --project-id int   Project ID (optional; resolved from GRS for the org when omitted)
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

* [flexera-cli policy applied-policy](flexera-cli_policy_applied-policy.md)	 - Applied policies (project-scoped)

