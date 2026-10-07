## flexera-cli finops adjustment

Adjustment definition

```
flexera-cli finops adjustment [flags]
```

### Options

```
  -h, --help   help for adjustment
```

### Options inherited from parent commands

```
      --access-token string      static bearer access token
      --api-base-url string      override API base URL
      --client-id string         OAuth client ID
      --client-secret string     OAuth client secret
      --config string            config file (default $HOME/.flexera/config.yaml)
  -d, --debug                    log HTTP requests/responses to stderr (Authorization redacted)
      --json-style string        JSON whitespace style (auto|pretty|compact) (default "auto")
      --login-base-url string    override login base URL
      --no-validate              skip API schema constraints (never JSON syntax or request data-loss checks)
      --optima-base-url string   Override the Optima base URL (env: FLEXERA_CLI_OPTIMA_BASE_URL)
      --org-id int               organization ID
      --out-fields string        project JSON output fields (comma-separated paths)
      --out-jq string            shape JSON output with a jq expression
  -o, --output string            output format (json|table)
  -r, --raw-output               write jq string results without JSON quotes (requires --out-jq)
      --refresh-token string     OAuth refresh token
      --zone string              API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli finops](flexera-cli_finops.md)	 - Cost analytics, billing centers, and recommendations (Optima)
* [flexera-cli finops adjustment show](flexera-cli_finops_adjustment_show.md)	 - Show the adjustment definition
* [flexera-cli finops adjustment update](flexera-cli_finops_adjustment_update.md)	 - Update the adjustment definition

