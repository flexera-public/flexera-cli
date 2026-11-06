## flexera-cli finops-onboarding processing-history bill-months-download

Download enterprise bill months in CSV

### Synopsis

Download enterprise bill months in CSV

Download enterprise bill months in CSV format. Currently only CSV format is supported.

```
flexera-cli finops-onboarding processing-history bill-months-download [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-onboarding processing-history bill-months-download --org-id ORG_ID --download-token DOWNLOAD_TOKEN
```

### Options

```
      --download-token string   downloadToken (query); Download token for enterprise bill months.; required by API
      --format string           format (query); Download format (default: csv); enum: ["csv"]; API default: "csv"; illustrative example: "csv"
  -h, --help[=true|false]       help for bill-months-download
```

### Options inherited from parent commands

```
      --access-token string        static bearer access token
      --api-base-url string        override API base URL
      --client-id string           OAuth client ID
      --client-secret string       OAuth client secret
      --config string              config file (default $HOME/.flexera/config.yaml)
  -d, --debug[=true|false]         log HTTP requests/responses to stderr (Authorization redacted)
      --json-style string          JSON whitespace style (auto|pretty|compact) (default "auto")
      --login-base-url string      override login base URL
      --no-validate[=true|false]   skip API schema constraints (never JSON syntax or request data-loss checks)
      --org-id int                 organization ID
      --out-fields string          project JSON output fields (comma-separated paths)
      --out-jq string              shape JSON output with a jq expression
  -o, --output string              output format (json|table)
  -r, --raw-output[=true|false]    write jq string results without JSON quotes (requires --out-jq)
      --refresh-token string       OAuth refresh token
      --zone string                API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli finops-onboarding processing-history](flexera-cli_finops-onboarding_processing-history.md)	 - Processing History operations (generated from the unified OpenAPI spec)

