## flexera-cli graphql query

Execute a raw GraphQL query against /explore/graphql

```
flexera-cli graphql query [flags]
```

### Options

```
      --body string   GraphQL request body: inline JSON, @file, or @- for stdin (required)
  -h, --help          help for query
      --path string   override the default /explore/graphql path
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

* [flexera-cli graphql](flexera-cli_graphql.md)	 - graphql operations (generated from the unified OpenAPI spec)

