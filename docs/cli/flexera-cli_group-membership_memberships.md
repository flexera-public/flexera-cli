## flexera-cli group-membership memberships

Add users to a group

```
flexera-cli group-membership memberships [flags]
```

### Options

```
      --body string    raw JSON body (inline | @file | @-); overrides body field flags
      --group-id int   groupId (path, required)
  -h, --help           help for memberships
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

* [flexera-cli group-membership](flexera-cli_group-membership.md)	 - Group Membership operations (generated from the unified OpenAPI spec)

