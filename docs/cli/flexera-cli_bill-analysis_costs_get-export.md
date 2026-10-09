## flexera-cli bill-analysis costs get-export

exportSelectStatus costs

### Synopsis

exportSelectStatus costs

Checks the status of an export initiated by the `costs/exportSelect` action.
Returns the status of the export, and if it is complete, a link to download the results.

```
flexera-cli bill-analysis costs get-export [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli bill-analysis costs get-export --org-id ORG_ID --export-id EXPORT_ID
```

### Options

```
      --export-id string   exportId (path, required); The ID of the export to check status for; required by API; illustrative example: "SGVsbG8gV29ybGQh"
  -h, --help               help for get-export
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

* [flexera-cli bill-analysis costs](flexera-cli_bill-analysis_costs.md)	 - costs operations (generated from the unified OpenAPI spec)

