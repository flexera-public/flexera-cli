## flexera-cli finops cost export-select

Raw POST /costs/select/export

```
flexera-cli finops cost export-select [flags]
```

### Options

```
      --file string   Path to JSON request body, or - for stdin (required)
  -h, --help          help for export-select
```

### Options inherited from parent commands

```
      --access-token string      static bearer access token
      --api-base-url string      override API base URL
      --client-id string         OAuth client ID
      --client-secret string     OAuth client secret
      --config string            config file (default $HOME/.flexera/config.yaml)
  -d, --debug                    log HTTP requests/responses to stderr (Authorization redacted)
      --login-base-url string    override login base URL
      --optima-base-url string   Override the Optima base URL (env: FLEXERA_CLI_OPTIMA_BASE_URL)
      --org-id int               organization ID
  -o, --output string            output format (json|table)
      --refresh-token string     OAuth refresh token
      --zone string              API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli finops cost](flexera-cli_finops_cost.md)	 - Cost queries and exports

