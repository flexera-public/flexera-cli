## flexera-cli policy policy-template evaluate

Evaluate a policy template

```
flexera-cli policy policy-template evaluate [flags]
```

### Options

```
      --file string      Path to JSON payload file, or - to read from stdin
  -h, --help             help for evaluate
      --id string        Policy template ID
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

* [flexera-cli policy policy-template](flexera-cli_policy_policy-template.md)	 - Policy templates (project-scoped)

