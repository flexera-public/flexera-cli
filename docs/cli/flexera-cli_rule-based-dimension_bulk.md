## flexera-cli rule-based-dimension bulk

Create or update rule-based dimensions from JSON input

```
flexera-cli rule-based-dimension bulk [flags]
```

### Options

```
      --continue-on-error   continue processing after a dimension fails
      --dry-run             validate and report dimensions without making API calls
      --file string         path to JSON input, or - to read JSON from stdin
  -h, --help                help for bulk
      --input string        inline JSON input (mutually exclusive with --file)
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

* [flexera-cli rule-based-dimension](flexera-cli_rule-based-dimension.md)	 - Rule-Based Dimension operations (generated from the unified OpenAPI spec)
