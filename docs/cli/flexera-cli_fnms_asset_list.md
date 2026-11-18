## flexera-cli fnms asset list

List assets

### Synopsis

List assets

Flexera One IT Asset Management (ITAM) Asset API offers detailed information about Assets managed within Flexera One ITAM. Users can filter to retrieve relevant assets by type, status, or location. The API provides asset details data along with inventory date, enterprise group (location, cost center, department), assigned and logged-in user, and purchase details.
Overall, the Flexera One ITAM Asset API serves as a crucial tool for maintaining accurate records of IT assets, supporting efficient IT asset management processes, and enabling organizations to achieve greater control and compliance in managing their IT resources.<br>
**Benefits and Practical Applications:**
* **Integration Capabilities:** The API supports seamless integration with other systems and tools within the organization's IT environment. This integration facilitates automated workflows for asset management, compliance monitoring, and optimization.<br><br>
* **Comprehensive Asset Management:** By leveraging this API, organizations can achieve enhanced asset management practices, including tracking asset lifecycles, optimizing asset utilization, and ensuring compliance with licensing and security policies.<br><br>
* **Enhanced Visibility and Control:** Access to detailed asset information through the API enables IT teams to gain better visibility into their asset inventory, make informed decisions, and effectively manage resources across the enterprise.<br><br>

```
flexera-cli fnms asset list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli fnms asset list --org-id ORG_ID
```

### Options

```
      --asset-type string   assetType (query); Filters assets by asset type; enum: ["Workstation","Server","Monitor","Desk","Chair","Printer","Router","Switch","Telephone","Cell Phone","Laptop","Mobile Device","DataStorage","Network Appliance","Mainframe"]; illustrative example: "Desk"
  -h, --help[=true|false]   help for list
      --limit int           limit (query); Specify the number of records to return in the request response; format: int64; maximum: 10000; API default: 10000; illustrative example: 5000
      --location string     location (query); Filter assets by location; illustrative example: "Flexera/IL"
      --offset int          offset (query); Specify the first record to return in a request response. Note that the offset is zero-based (the first record is assigned index 0 instead of 1); format: int64; API default: 0; illustrative example: 500
      --status string       status (query); Filters assets by status; enum: ["Purchased","In Storage","Installed","Retired","Disposed","Other"]; illustrative example: "Purchased"
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

* [flexera-cli fnms asset](flexera-cli_fnms_asset.md)	 - Asset operations (generated from the unified OpenAPI spec)

