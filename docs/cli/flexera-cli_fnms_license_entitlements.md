## flexera-cli fnms license entitlements

List license entitlements

### Synopsis

List license entitlements

Flexera One IT Asset Management (ITAM) Licenses Entitlements API is specifically designed to provide organizations with detailed insights into the entitlements configured on each software license the organization is managing within ITAM.<br>
**Benefits and Practical Applications:**
* **Integration and Automation:** Supports integration with other IT systems and tools, facilitating automated workflows for license management, compliance monitoring, and reporting. Enables real-time data synchronization and updates with external systems to ensure accurate and up-to-date entitlement information.<br><br>
* **Visibility and Control:** Facilitates proactive management of entitlements, allowing for optimization of license allocations and cost management.<br><br>
* **Purchased and entitlement limits:** Understanding and effectively managing purchased licenses, entitlement limits, and added extra entitlements. This ensures compliance, cost efficiency, and strategic alignment of software resources with organizational goals.<br><br>

```
flexera-cli fnms license entitlements [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli fnms license entitlements --org-id ORG_ID
```

### Options

```
  -h, --help[=true|false]   help for entitlements
      --limit int           limit (query); Specify the number of records to return in the request response; format: int64; maximum: 10000; API default: 10000; illustrative example: 5000
      --offset int          offset (query); Specify the first record to return in a request response. Note that the offset is zero-based (the first record is assigned index 0 instead of 1); format: int64; API default: 0; illustrative example: 500
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

* [flexera-cli fnms license](flexera-cli_fnms_license.md)	 - License operations (generated from the unified OpenAPI spec)

