## flexera-cli policy meta

Relationship-aware applied-policy meta operations (project-scoped)

```
flexera-cli policy meta [flags]
```

### Options

```
  -h, --help   help for meta
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

* [flexera-cli policy](flexera-cli_policy.md)	 - Project-scoped policy operations with GRS project auto-resolution
* [flexera-cli policy meta audit](flexera-cli_policy_meta_audit.md)	 - Audit status of every applied policy in a project
* [flexera-cli policy meta discover](flexera-cli_policy_meta_discover.md)	 - Discover applied-policy parent/child relationships
* [flexera-cli policy meta terminate-children](flexera-cli_policy_meta_terminate-children.md)	 - Delete all unambiguous children of a parent applied policy
* [flexera-cli policy meta terminate-orphaned](flexera-cli_policy_meta_terminate-orphaned.md)	 - Delete applied policies whose parent reference is orphaned

