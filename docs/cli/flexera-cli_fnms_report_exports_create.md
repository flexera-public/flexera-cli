## flexera-cli fnms report exports create

Start an FNMS report export

### Synopsis

Start an FNMS report export

Flexera One ITAM Report Export API creates a file containing dynamic report data for the built-in and custom saved report. This API is capable of asynchronously generating a report data which can be downloaded as a zip file.<br>
**Benefits and Practical Applications:**
* **Enhanced Decision-Making:** Customers can generate detailed and tailored reports that provide insights into IT asset management, compliance statuses, and other data needed by business owners.<br><br>
* **Efficiency and Automation:** By automating the report generation process and using the API for data retrieval, organizations save time and reduce manual effort.<br><br>
* **Integration:** The flexibility of the Report Execute API allows seamless integration with other IT systems, enabling data-driven decision-making and comprehensive IT asset management practices.<br><br>
* **Built-In Report with Parameter Support**: We can now retrieve data from built-in reports with parameter support. This feature allows for more effective data retrieval and filtering within the Execute and Export API. Please note that this functionality is primarily intended for built-in reports.<br><br>
* **Export File Availability (1 hour):** The generated export file is available for download for 1 hour after it is created; after that you must request a new export.<br><br>

**Steps:**
* **Step 1:** Send a POST request (without a data payload) to "/fnms/v1/orgs/{orgId}/reports/{id}/exports". This starts the generation and returns an export ID.
* **Step 2:** Send a GET request to "/fnms/v1/orgs/{orgId}/reports/exports/{exportId}" with the export ID returned in Step 1. This should eventually return a status of "Completed", along with a "path" and "completionDate".
* **Step 3:** Use the "path" to download the content. This will be a long Flexera One download URL using the base domain for your organization's zone, such as "https://api.flexera.com/fnmsdata/v1/b0db7056f3b382e9efbdf154f46..." for NAM.

```
flexera-cli fnms report exports create [flags]
```

### Options

```
      --accept string          export format: application/json or text/csv (default "application/json")
      --dry-run[=true|false]   preview without making API requests
      --filter-params string   report filters as a JSON object (deepObject query parameters); Query parameters
  -h, --help[=true|false]      help for create
      --id string              saved report ID (required); The saved report ID against which data will be exported in a zip file.; required by API; illustrative example: "100534"
      --search-text string     text to filter report results; Text/Keyword to search from report attributes in order to filter and limit the results returned.; illustrative example: "dev"
      --yes[=true|false]       confirm operation (not required for non-destructive report jobs)
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

