## flexera-cli policy meta terminate-children

Delete all unambiguous children of a parent applied policy

```
flexera-cli policy meta terminate-children [flags]
```

### Options

```
      --dry-run            Preview matched children without deleting them
  -h, --help               help for terminate-children
      --parent-id string   Parent applied policy ID
      --project-id int     Project ID (optional; resolved from GRS for the org when omitted)
      --yes                Confirm the operation; required unless --dry-run
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

