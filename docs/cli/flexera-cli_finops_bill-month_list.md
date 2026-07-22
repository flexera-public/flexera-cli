## flexera-cli finops bill-month list

List bill months

```
flexera-cli finops bill-month list [flags]
```

### Options

```
  -h, --help              help for list
      --limit int         Maximum number of records (0 = API default)
      --offset int        Starting offset for pagination
      --order-by string   Order spec, e.g. month=desc
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

* [flexera-cli finops bill-month](flexera-cli_finops_bill-month.md)	 - Bill months

