## flexera-cli user-orgs list

List organizations the authenticated user (or --id) can access

```
flexera-cli user-orgs list [flags]
```

### Options

```
  -h, --help     help for list
      --id int   Flexera user ID (auto-detected from the access-token JWT when omitted)
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

* [flexera-cli user-orgs](flexera-cli_user-orgs.md)	 - List organizations a user can access (IAM)

