## flexera-cli finops anomaly-report

Run an anomaly report

```
flexera-cli finops anomaly-report [flags]
```

### Options

```
      --file string   Path to JSON request body, or - for stdin (required)
  -h, --help          help for anomaly-report
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

* [flexera-cli finops](flexera-cli_finops.md)	 - Cost analytics, billing centers, and recommendations (Optima)

