## flexera-cli customization list

Index customizations which have been applied to an org

```
flexera-cli customization list [flags]
```

### Options

```
      --filter string       filter (query)
  -h, --help                help for list
      --no-paginate         return only the first page (do not follow nextPage)
      --skip-token string   resume pagination from this token
      --view string         view (query)
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

* [flexera-cli customization](flexera-cli_customization.md)	 - Customization operations (generated from the unified OpenAPI spec)

