## flexera-cli fnms installed-software-v2 create

Create a file containing active devices with installed software

### Synopsis

Create a file containing active devices with installed software

The Flexera One ITAM Installed Software V2 API generates a report file containing all active devices and their installed software. Report generation is asynchronous, and the output is downloaded as a ZIP file.

There are three steps to get the report:

* **Step 1: Start report generation**
  Send a POST request (no request body) to:
  "/fnms/v2/orgs/{orgId}/devices/installed-software"
  The request initiates report generation and returns a job status of **Created**.

* **Step 2: Wait for completion (two options)**
  1. **Poll the generation endpoint**
     Repeat the same POST request to "/fnms/v2/orgs/{orgId}/devices/installed-software" until the response status becomes **Completed** and includes a **completionDate**.
  2. **Check the generated file endpoint**
     Use the file path returned in the Step 1 response and send a GET request, for example:
     "/fnms/v1/orgs/{orgId}/files/{filename}"
     When the file is ready, the response includes a **Location** header containing the download URL.

* **Step 3: Download the ZIP**
  Retrieve the ZIP from the URL provided in the **Location** header. This will typically be a long URL, for example, "https://api.flexera.com/fnmsdata/v1/b0db7056f3b382e9efbdf154f469d407980e98ffc39973dc95730226097e7419.zip".
```
{
  "assignedUser": "Galaxy Admin",
  "computerId": 38183,
  "computerName": "LTP-EMEA-102830",
  "deviceStatus": "Active",
  "hostNameForTheVM": "ws01e",
  "hostedIn": "AWS",
  "inventoryDate": "2006-09-15",
  "inventoryDeviceRole": "Production",
  "inventoryDeviceType": "Virtual Machine",
  "ipAddress": "10.0.1.2",
  "lastInventorySource": "Flexera Saas",
  "lastLoggedInUser": "csa-test-user-1",
  "manufacturer": "Dell",
  "model": "Optiplex",
  "numberOfCores": 1,
  "numberOfProcessors": 1,
  "os": "Microsoft Windows Server 2008 R2 Enterprise",
  "processorType": "Intel",
  "serialNo": "8NOGS1S",
  "softwareTitles": [
    {
      "accessMode": "ZenDesktop",
      "category": "Software",
      "classification": "Freeware",
      "discoveryDate": "2006-01-15",
      "edition": "Enterprise",
      "lastUsedDate": "2006-10-18",
      "productName": "SQL Server",
      "publisher": "Microsoft",
      "softwareStatus": "[Authorized]",
      "softwareTitleName": "Windows Server 2008 R2 Enterprise",
      "version": "2008 R2"
    },
    {
      "accessMode": "ZenDesktop",
      "category": "Software",
      "classification": "Freeware",
      "discoveryDate": "2006-01-15",
      "edition": "Enterprise",
      "lastUsedDate": "2006-10-18",
      "productName": "SQL Server",
      "publisher": "Microsoft",
      "softwareStatus": "[Authorized]",
      "softwareTitleName": "Windows Server 2008 R2 Enterprise",
      "version": "2008 R2"
    },
    {
      "accessMode": "ZenDesktop",
      "category": "Software",
      "classification": "Freeware",
      "discoveryDate": "2006-01-15",
      "edition": "Enterprise",
      "lastUsedDate": "2006-10-18",
      "productName": "SQL Server",
      "publisher": "Microsoft",
      "softwareStatus": "[Authorized]",
      "softwareTitleName": "Windows Server 2008 R2 Enterprise",
      "version": "2008 R2"
    },
    {
      "accessMode": "ZenDesktop",
      "category": "Software",
      "classification": "Freeware",
      "discoveryDate": "2006-01-15",
      "edition": "Enterprise",
      "lastUsedDate": "2006-10-18",
      "productName": "SQL Server",
      "publisher": "Microsoft",
      "softwareStatus": "[Authorized]",
      "softwareTitleName": "Windows Server 2008 R2 Enterprise",
      "version": "2008 R2"
    }
  ],
  "totalMemory": 1073741824
}
```

```
flexera-cli fnms installed-software-v2 create [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli fnms installed-software-v2 create --org-id ORG_ID
  flexera-cli fnms installed-software-v2 create --org-id ORG_ID --dry-run
```

### Options

```
      --dry-run[=true|false]      print the planned operation as JSON and exit without calling the API
  -h, --help[=true|false]         help for create
      --os strings                os (query); Limit the request response to devices with the specified operating system installed; CLI: comma-separated values or repeated flag; illustrative example: ["linux","commodore"]
      --publisher strings         publisher (query); Limit the request response to devices with installed software created by the specified publisher; CLI: comma-separated values or repeated flag; illustrative example: ["microsoft","ibm"]
      --software-status strings   softwareStatus (query); Limit the request response to devices with installed software with the specified software status; CLI: comma-separated values or repeated flag; illustrative example: ["unmanaged","retired"]
      --software-title strings    softwareTitle (query); Limit the request response to devices with the specified software titles installed; CLI: comma-separated values or repeated flag; illustrative example: ["adobe","notepad","goland"]
      --yes[=true|false]          confirm the operation (required for destructive ops)
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

* [flexera-cli fnms installed-software-v2](flexera-cli_fnms_installed-software-v2.md)	 - Installed Software V2 operations (generated from the unified OpenAPI spec)

