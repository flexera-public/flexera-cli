## flexera-cli fnms report list

List ITAM saved and built-in reports

### Synopsis

List ITAM saved and built-in reports

Flexera One ITAM Reports List API helps retrieve a list of built-in and custom reports, including the report title (name), report ID, and additional details for automating the Report Execute API.<br>
**Benefits and Practical Applications:**
* **Enhanced Decision-Making:** Customers can get detailed and tailored list of built-in and custom reports including additional details.<br><br>
* **Integration and Automation:** Enables seamless integration with other IT systems and tools, supporting automated workflows for report generation.<br><br>
* **Parameter Support Integrating:** This feature, allows for more effective filtering of report data within the Execute and Export API.Please note that this functionality is primarily intended for built-in reports.

```
flexera-cli fnms report list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli fnms report list --org-id ORG_ID
```

### Options

```
  -h, --help[=true|false]   help for list
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

* [flexera-cli fnms report](flexera-cli_fnms_report.md)	 - Report operations (generated from the unified OpenAPI spec)

