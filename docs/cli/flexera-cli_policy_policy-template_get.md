## flexera-cli policy policy-template get

Show a policy template

```
flexera-cli policy policy-template get [flags]
```

### Options

```
  -h, --help             help for get
      --id string        Policy template ID
      --project-id int   Project ID (optional; resolved from GRS for the org when omitted)
      --view string      Optional Policy template view
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

* [flexera-cli policy policy-template](flexera-cli_policy_policy-template.md)	 - Policy templates (project-scoped)

