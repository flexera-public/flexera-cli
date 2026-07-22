## flexera-cli user-orgs

List organizations a user can access (IAM)

### Examples

```
flexera-cli user-orgs list --refresh-token <token>   # discover orgs your credentials can access
```

### Options

```
  -h, --help   help for user-orgs
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

* [flexera-cli](flexera-cli.md)	 - Flexera One unified API command-line client
* [flexera-cli user-orgs list](flexera-cli_user-orgs_list.md)	 - List organizations the authenticated user (or --id) can access

