## flexera-cli saved-filters list

Index saved filters

```
flexera-cli saved-filters list [flags]
```

### Options

```
      --favorite-ids string   favoriteIds (query)
  -h, --help                  help for list
      --limit int             limit (query)
      --no-paginate           return only the first page (do not follow nextPage)
      --order-by string       orderBy (query)
      --skip-token string     resume pagination from this token
      --visibility string     visibility (query)
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

* [flexera-cli saved-filters](flexera-cli_saved-filters.md)	 - Saved Filters operations (generated from the unified OpenAPI spec)

