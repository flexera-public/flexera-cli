## flexera-cli finops-billing shared-cost-rules telemetry-uploads get

Show a telemetry upload

### Synopsis

Show a telemetry upload

Returns a staged telemetry upload with the requested page of rows. summary and warnings always cover the whole file, whatever the page or filter. The file's errors are returned only by uploadTelemetry; each row still lists its own error and warning codes. Rows are paged by offset and limit because a staged upload never changes, so pages are stable.

```
flexera-cli finops-billing shared-cost-rules telemetry-uploads get [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing shared-cost-rules telemetry-uploads get --org-id ORG_ID --upload-id UPLOAD_ID
```

### Options

```
      --filter string       filter (query); all returns every row, issues the rows with any error or warning, changes the rows that are added, modified or deleted.; enum: ["all","issues","changes"]; API default: "all"; illustrative example: "all"
  -h, --help[=true|false]   help for get
      --limit int           limit (query); Maximum number of rows to return.; minimum: 1; maximum: 100; API default: 50; illustrative example: 50
      --offset int          offset (query); Index of the first row to return, counted after the filter. A staged upload never changes during its 24 hours and holds at most one 50 MB file, so an offset always returns the same rows and clients can jump straight to any page.; minimum: 0; API default: 0; illustrative example: 0
      --upload-id string    uploadId (path, required); Identifier of the staged telemetry upload.; required by API; illustrative example: "tu_01JAB3K9Q2X7V5R8M4N6P0T1W3"
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

* [flexera-cli finops-billing shared-cost-rules telemetry-uploads](flexera-cli_finops-billing_shared-cost-rules_telemetry-uploads.md)	 - telemetry-uploads operations

