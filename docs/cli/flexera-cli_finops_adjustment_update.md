## flexera-cli finops adjustment update

Update the adjustment definition

```
flexera-cli finops adjustment update [flags]
```

### Options

```
      --dry-run       Print the planned operation as JSON and exit without contacting the API
      --file string   Path to JSON request body, or - for stdin (required)
  -h, --help          help for update
      --yes           Confirm the operation; required for destructive ops
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

* [flexera-cli finops adjustment](flexera-cli_finops_adjustment.md)	 - Adjustment definition

