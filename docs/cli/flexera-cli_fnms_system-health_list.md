## flexera-cli fnms system-health list

Get system health status

### Synopsis

Get system health status

The System Health API returns the same data that is visible from the System Health page in the UI. Users can use query parameters to access specific health components or combine multiple parameters to return multiple sets of data.

**Available Health Components:**
* **ARL Status** - Application Recognition Library status
* **SKU Library** - SKU Library status
* **PURL** - Package URL status
* **Data Platform** - Imported software status
* **License Reconciliation** - License reconciliation status
* **Beacon Status** - Beacon connectivity status
* **Inventory Data Inputs** - Data inputs with issues
* **Import Resolver Errors** - Failures processing discovery and inventory uploads
* **Business Adapters** - Business data inputs with issues
* **FNMEA Agent Report Log Rotation** - Engineering Applications agents with report log rotation issues

**Benefits:**
* **Comprehensive System Monitoring:** Get visibility into all critical system components in one API call
* **Flexible Data Retrieval:** Request specific health components or combine multiple components as needed
* **Integration Support:** Enable automated monitoring and alerting based on system health status
* **Operational Efficiency:** Programmatically access the same health data visible in the UI for reporting and dashboards

**Usage Examples:**
* Get all health data: /fnms/v1/orgs/12345/system-health
* Get specific components: /fnms/v1/orgs/12345/system-health?arlStatus&purlStatus&importResolverErrors
* Get single component: /fnms/v1/orgs/12345/system-health?beaconStatus

**Important Notes:**
* **Query Parameters:** Simply including a query parameter in the URL (e.g., ?arlStatus) is sufficient to request that component. No value is required.
* **Timestamps:** All timestamps returned by this API are in UTC format. The FNMS UI may convert these to the operator's local timezone for display.
* **Data Retrieval Status:** Each health component includes a dataRetrievalStatus field. Always check this field before using the component data. If the value is 'failed' or 'unknown', an error occurred during data collection and the returned fields contain default values that do not reflect the actual component status.

```
flexera-cli fnms system-health list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli fnms system-health list --org-id ORG_ID
```

### Options

```
      --arl-status string                               arlStatus (query); Includes Application Recognition Library status; illustrative example: "true"
      --beacon-status string                            beaconStatus (query); Includes Beacon status; illustrative example: "true"
      --business-adapters string                        businessAdapters (query); Includes Business Adapters status; illustrative example: "true"
      --data-platform string                            dataPlatform (query); Includes Data Platform status; illustrative example: "true"
      --fnmea-agent-report-log-rotation-status string   fnmeaAgentReportLogRotationStatus (query); Includes FNMEA Agent Report Log Rotation status; illustrative example: "true"
  -h, --help[=true|false]                               help for list
      --import-resolver-errors string                   importResolverErrors (query); Includes Import Resolver Errors status; illustrative example: "true"
      --input-data-feed string                          inputDataFeed (query); Includes Inventory Data Feed status; illustrative example: "true"
      --license-reconciliation string                   licenseReconciliation (query); Includes License Reconciliation status; illustrative example: "true"
      --purl-status string                              purlStatus (query); Includes Package URL status; illustrative example: "true"
      --sku-status string                               skuStatus (query); Includes SKU Library status; illustrative example: "true"
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

* [flexera-cli fnms system-health](flexera-cli_fnms_system-health.md)	 - System Health operations (generated from the unified OpenAPI spec)

