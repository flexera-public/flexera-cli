## flexera-cli group-membership list

Index user memberships of a group in an organization

```
flexera-cli group-membership list [flags]
```

### Options

```
      --group-id int   groupId (path, required)
  -h, --help           help for list
      --view string    view (query)
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

