## flexera-cli managed-application app-events

List managed application's events

```
flexera-cli managed-application app-events [flags]
```

### Options

```
  -h, --help                      help for app-events
      --managed-app-id string     managedAppId (path, required)
      --most-recent-events-only   mostRecentEventsOnly (query)
      --no-paginate               return only the first page (do not follow nextPage)
      --skip-token string         resume pagination from this token
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

* [flexera-cli managed-application](flexera-cli_managed-application.md)	 - Managed Application operations (generated from the unified OpenAPI spec)

