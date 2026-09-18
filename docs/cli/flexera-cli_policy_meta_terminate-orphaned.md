## flexera-cli policy meta terminate-orphaned

Delete applied policies whose parent reference is orphaned

```
flexera-cli policy meta terminate-orphaned [flags]
```

### Options

```
      --dry-run          Preview matched orphaned policies without deleting them
  -h, --help             help for terminate-orphaned
      --project-id int   Project ID (optional; resolved from GRS for the org when omitted)
      --yes              Confirm the operation; required unless --dry-run
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

