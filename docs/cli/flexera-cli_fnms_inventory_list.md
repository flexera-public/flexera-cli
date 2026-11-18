## flexera-cli fnms inventory list

List inventories

### Synopsis

List inventories

Flexera One IT Asset Management (ITAM) Inventory Devices API offers detailed information about inventoried devices. Users can filter devices by type, status, manufacturer, or location. The API provides data on the last inventory source, inventory date, enterprise group (location, cost center, department), assigned and logged-in user details, and linked assets.
By using the Flexera One ITAM Inventory Devices API, organizations can enhance visibility into their inventoried devices, improve asset management practices, ensure compliance with licensing and security policies, and streamline IT operations. This API is essential for maintaining an accurate and current inventory of IT assets throughout the enterprise and for updating hardware inventory data into the CMDB.<br>
**Benefits and Practical Applications:**
* **Visibility into Inventoried Devices:** The API provides detailed insights into all inventoried devices within the organization's IT infrastructure. This visibility includes hardware details, such as device types, models, configurations, and ownership from location and cost center perspectives.<br><br>
* **Improved Asset Management Practices:** By accessing comprehensive device information through the API, organizations can implement more robust asset management practices. This includes tracking device lifecycles, optimizing asset usage, and identifying opportunities for cost savings.<br><br>
* **Streamlined IT Operations:** Integrating the API into IT operations allows for streamlined processes. Organizations can automate inventory updates, perform audits efficiently, and allocate resources more effectively based on real-time data.<br><br>
* **Integration with CMDB:** One significant advantage is that the API provides current hardware inventory data with the business context into the Configuration Management Database (CMDB). This integration enhances the accuracy and completeness of the CMDB, supporting IT Service Management (ITSM) processes such as incident management, change management, and asset relationship mapping.<br><br>

```
flexera-cli fnms inventory list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli fnms inventory list --org-id ORG_ID
```

### Options

```
      --compliance-status string   complianceStatus (query); Filter inventories by compliance status; enum: ["Changed","Compliant","Ignored","New","Not Tracked"]; illustrative example: "Changed"
      --device-type string         deviceType (query); Filter inventories by device type; enum: ["Computer","Container Image","Mobile Device","Remote Device","VDI Template","Virtual Machine","VM Host"]; illustrative example: "Computer"
  -h, --help[=true|false]          help for list
      --limit int                  limit (query); Specify the number of records to return in the request response; format: int64; maximum: 10000; API default: 10000; illustrative example: 5000
      --location string            location (query); Filter inventories by location; illustrative example: "Asia/India/ Tamil Nadu/Chennai/WalLoc001"
      --manufacturer string        manufacturer (query); Filter inventories by manufacturer; illustrative example: "Dell"
      --offset int                 offset (query); Specify the first record to return in a request response. Note that the offset is zero-based (the first record is assigned index 0 instead of 1); format: int64; API default: 0; illustrative example: 500
      --status string              status (query); Filter inventories by status; enum: ["Active","Archived","Awaiting Inventory","Ignored"]; illustrative example: "Active"
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

* [flexera-cli fnms inventory](flexera-cli_fnms_inventory.md)	 - Inventory operations (generated from the unified OpenAPI spec)

