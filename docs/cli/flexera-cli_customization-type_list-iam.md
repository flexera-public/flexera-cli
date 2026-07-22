## flexera-cli customization-type list-iam

Returns types of customizations that are available to an org

```
flexera-cli customization-type list-iam [flags]
```

### Options

```
      --filter string       filter (query)
  -h, --help                help for list-iam
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

* [flexera-cli customization-type](flexera-cli_customization-type.md)	 - Customization Type operations (generated from the unified OpenAPI spec)

