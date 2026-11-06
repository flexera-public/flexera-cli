## flexera-cli finops-billing shared-cost-rules telemetry-uploads create

Upload a telemetry file

### Synopsis

Upload a telemetry file

Validates a telemetry CSV for a dynamicProportionateSplit rule, compares it with the rule's saved telemetry and stages the result for 24 hours. Returns the summary and every error and warning in the file, without rows; showTelemetryUpload pages the rows. Staging changes nothing; an upload is applied only by a rule save that names it. The request body must not exceed 50 MB.

```
flexera-cli finops-billing shared-cost-rules telemetry-uploads create [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own upload file.
  flexera-cli finops-billing shared-cost-rules telemetry-uploads create --org-id ORG_ID --dimension-id DIMENSION_ID --file FILE
  flexera-cli finops-billing shared-cost-rules telemetry-uploads create --org-id ORG_ID --dimension-id DIMENSION_ID --file FILE --dry-run
```

### Options

```
      --dimension-id string    dimensionId (multipart form field); required by API; The destination dimension, rbd_* or tag_*. One the org doesn't have is reported as an unknown_dimension error in the 422 body.; illustrative example: "rbd_team"
      --dry-run[=true|false]   print the planned operation as JSON and exit without calling the API
      --file string            file (multipart file path); required by API; The CSV file.; format: binary
  -h, --help[=true|false]      help for create
      --rule-id string         ruleId (multipart form field); The rule the file is for. Omit it for a new rule.; illustrative example: "68f2a91c4b3d2e1f5a6b7c8d"
      --yes[=true|false]       confirm the operation (required for destructive ops)
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

