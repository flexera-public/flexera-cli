## flexera-cli finops-billing shared-cost-rules telemetry

Download a rule's telemetry

### Synopsis

Download a rule's telemetry

Downloads the telemetry a dynamicProportionateSplit rule holds now (Download History) as CSV with exactly the upload template's columns: Date (YYYY-MM-DD), the dimension's name and Metrics (>=0). Rows are sorted by date and then dimension value, with dates formatted YYYY-MM-DD; records deleted by an upload are not included. Dimension values starting with =, +, - or @ are written with a leading ' so spreadsheets do not run them as formulas. The file can be uploaded again unchanged. A rule with no stored records downloads as the header alone.

```
flexera-cli finops-billing shared-cost-rules telemetry [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing shared-cost-rules telemetry --org-id ORG_ID --id ID
```

### Options

```
  -h, --help[=true|false]   help for telemetry
      --id string           id (path, required); Unique identifier of the shared cost rule; required by API; illustrative example: "68f2a91c4b3d2e1f5a6b7c8d"
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

* [flexera-cli finops-billing shared-cost-rules](flexera-cli_finops-billing_shared-cost-rules.md)	 - Shared Cost Rules operations (generated from the unified OpenAPI spec)

