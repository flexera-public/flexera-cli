## flexera-cli bill-analysis bill-months download

download bill-months

### Synopsis

download bill-months

Download bill months data in CSV format using a download token

```
flexera-cli bill-analysis bill-months download [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli bill-analysis bill-months download --org-id ORG_ID --download-token DOWNLOAD_TOKEN
```

### Options

```
      --download-token string   download_token (query); Token for downloading the complete result set; required by API
      --format string           format (query); Download format (default: csv); enum: ["csv"]; illustrative example: "csv"
  -h, --help                    help for download
```

### Options inherited from parent commands

```
      --access-token string     static bearer access token
      --api-base-url string     override API base URL
      --client-id string        OAuth client ID
      --client-secret string    OAuth client secret
      --config string           config file (default $HOME/.flexera/config.yaml)
  -d, --debug                   log HTTP requests/responses to stderr (Authorization redacted)
      --json-style string       JSON whitespace style (auto|pretty|compact) (default "auto")
      --login-base-url string   override login base URL
      --no-validate             skip API schema constraints (never JSON syntax or request data-loss checks)
      --org-id int              organization ID
      --out-fields string       project JSON output fields (comma-separated paths)
      --out-jq string           shape JSON output with a jq expression
  -o, --output string           output format (json|table)
  -r, --raw-output              write jq string results without JSON quotes (requires --out-jq)
      --refresh-token string    OAuth refresh token
      --zone string             API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli bill-analysis bill-months](flexera-cli_bill-analysis_bill-months.md)	 - bill-months operations (generated from the unified OpenAPI spec)

