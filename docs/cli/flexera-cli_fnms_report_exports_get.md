## flexera-cli fnms report exports get

Status of the report export

### Synopsis

Status of the report export

Flexera One ITAM Report Export Status API accepts the "Export ID" and returns the status of report being generated for the export.
Once a report export starts, the Report Export Status API uses "Export ID" for tracking the status of report being generated for exporting.<br>
**Benefits and Practical Applications:**
* **Enhanced Decision-Making:** Customers can track the status of report being generated for exporting.<br><br>
* **Efficiency and Automation:** By using the API for status retrieval, organizations save time and reduce manual effort.<br><br>

```
flexera-cli fnms report exports get [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli fnms report exports get --org-id ORG_ID --export-id EXPORT_ID
```

### Options

```
      --export-id string    exportId (path, required); The export ID returned when export report was requested.; required by API; illustrative example: "278900"
  -h, --help[=true|false]   help for get
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

* [flexera-cli fnms report exports](flexera-cli_fnms_report_exports.md)	 - exports operations

