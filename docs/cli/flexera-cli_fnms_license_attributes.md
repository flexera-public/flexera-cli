## flexera-cli fnms license attributes

List license attributes

### Synopsis

List license attributes

Flexera One IT Asset Management (ITAM) Licenses Attributes API is intended to give organizations visibility, compliance, and optimization details for software licenses. This includes compliance status, metrics, duration, purchases, consumption, and usage, as well as category, serial number, and license key rules. This empowers IT teams to streamline the license management process, reduce compliance risks, and effectively optimize software spending.<br>
**Benefits and Practical Applications:**
* **Compliance and optimization details:** Provides insights into the actual used, consumed and purchased across the organization.<br><br>
* **Purchase and Renewal Details:**<br><br>
    * **Purchase Information:** Provides details related to the purchase of each license, including purchase date, cost, and renewal terms.<br><br>
    * **Renewal Alerts:** Alerts regarding upcoming renewals or expirations, helping organizations stay proactive in managing their software assets.<br><br>
* **Integration and Automation:** Enables seamless integration with other IT systems and tools, supporting automated workflows for license management, compliance monitoring, and reporting.
                            Facilitates real-time updates and data synchronization with external systems to ensure accurate and up-to-date license information.<br><br>
* **Risk Reduction and Cost Efficiency:** Empowers IT teams to reduce compliance risks by ensuring adherence to license agreements and regulations. Supports cost-efficient management of software assets by optimizing license allocations and identifying opportunities for cost savings.<br><br>

```
flexera-cli fnms license attributes [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli fnms license attributes --org-id ORG_ID
```

### Options

```
  -h, --help[=true|false]   help for attributes
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

