## flexera-cli policy meta audit

Audit status of every applied policy in a project

```
flexera-cli policy meta audit [flags]
```

### Options

```
  -h, --help             help for audit
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

* [flexera-cli policy meta](flexera-cli_policy_meta.md)	 - Relationship-aware applied-policy meta operations (project-scoped)

