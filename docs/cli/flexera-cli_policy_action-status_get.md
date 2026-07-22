## flexera-cli policy action-status get

Show an action status

```
flexera-cli policy action-status get [flags]
```

### Options

```
  -h, --help             help for get
      --id string        Action status ID
      --project-id int   Project ID (optional; resolved from GRS for the org when omitted)
      --view string      Optional Policy action-status view
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

* [flexera-cli policy action-status](flexera-cli_policy_action-status.md)	 - Policy action statuses (project-scoped)

