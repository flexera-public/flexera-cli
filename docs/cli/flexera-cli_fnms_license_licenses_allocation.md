## flexera-cli fnms license licenses allocation

Allocate a license to users/devices

### Synopsis

Allocate a license to users/devices

License allocation API aims to automate and streamline the allocation of licenses. This includes assigning one or more users or devices to a specific user or device license, including the license key, expiry date, and reason for the allocation.
The License Allocation API enables integration with IT Service Management (ITSM) workflows and other IT Asset Management (ITAM) automation processes. This streamlines the management of software licenses by automating allocation, thereby enhancing the organization's ability to efficiently manage license distribution, tracking, and compliance across their IT infrastructure.<br>

**Notes:**

* If a license allocation process request for allocation but not to consume until the requested application is installed and reported by inventory data, we recommend keeping the "Allocations consume license entitlements" setting unchecked on the license properties "User Rights & Rules" tab.
* The response will only include the allocations that were not successfully processed, along with the reason for failure. If all allocations were successful, allocationStatuses will be omitted or an empty array (the response body may be an empty JSON object).

**Benefits and Practical Applications:**
* **License Request Automation:** Allows users to request licenses through automated workflows.
* **Lifecycle Management:** Manages licenses from allocation to expiration, ensuring optimal utilization and compliance.
* **Efficiency:** Reduces manual effort and potential errors in license assignment/allocation.
* **Compliance:** Ensures that license allocations are tracked and documented correctly, aiding in compliance audits.
* **Integration:** Fits into existing ITSM and ITAM workflows, enhancing overall IT governance and asset management practices.

                    Note: Valid licenseExpiryDate format
                        YYYY-MM-DD
                            e.g. 2020-03-05

                        OR

                        YYYY-MM-DD HH:mm:SS
                            e.g. 2020-03-05 17:59:00

                    User-based allocation JSON request format:
                        [
                            {
                                "samAccountName": "qa.flexera.com\\user1",
                                "userEmailAddress": "user1@flexera.com",
                                "licenseKey": "ABC101",
                                "licenseExpiryDate": "2025-12-31",
                                "allocationReason": "Development",
                                "exemptionReason": ""
                            },
                            {
                                "samAccountName": "qa.flexera.com\\user2",
                                "userEmailAddress": "user2@flexera.com",
                                "licenseKey": "ABC105",
                                "licenseExpiryDate": "2025-12-31",
                                "allocationReason": "QA",
                                "exemptionReason": ""
                            }
                        ]

                    User-based allocation CSV request format:

                        samAccountName,userEmailAddress,licenseKey,licenseExpiryDate,allocationReason,exemptionReason
                        qa.flexera.com\user1,user1@flexera.com,ABC101,"2025-12-31",Development,
                        qa.flexera.com\user2,user2@flexera.com,ABC105,"2025-12-31",QA,


                    Device-based allocation JSON request format:
                        [
                            {
                                "deviceSerial": "GK0L9X1",
                                "deviceName": "test-machine",
                                "licenseKey": "ABC101",
                                "licenseExpiryDate": "2025-12-31",
                                "allocationReason": "Development",
                                "exemptionReason": ""
                            },
                            {
                                "deviceSerial": "AB0L9X45",
                                "deviceName": "redirection",
                                "licenseKey": "ABC102",
                                "licenseExpiryDate": "2025-12-31",
                                "allocationReason": "Development,Staging",
                                "exemptionReason": ""
                            }
                        ]

                    Device-based allocation CSV request format:

                        deviceSerial,deviceName,licenseKey,licenseExpiryDate,allocationReason,exemptionReason
                        GK0L9X1,test-machine,ABC101,"2025-12-31",Development,
                        AB0L9X45,redirection,ABC105,"2025-12-31",QA,

```
flexera-cli fnms license licenses allocation [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli fnms license licenses allocation --org-id ORG_ID --license-id LICENSE_ID
  flexera-cli fnms license licenses allocation --org-id ORG_ID --license-id LICENSE_ID --dry-run
```

### Options

```
      --dry-run[=true|false]   print the planned operation as JSON and exit without calling the API
  -h, --help[=true|false]      help for allocation
      --license-id string      licenseId (path, required); Globally unique identifier for the license.; required by API; illustrative example: "794ca625-4230-4c8d-9bfc-3849d9d89857"
      --yes[=true|false]       confirm the operation (required for destructive ops)
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

* [flexera-cli fnms license licenses](flexera-cli_fnms_license_licenses.md)	 - licenses operations

