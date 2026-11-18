## flexera-cli fnms license licenses deallocation

Deallocate a license from users/devices.

### Synopsis

Deallocate a license from users/devices.

The License De-allocation API serves as a counterpart to the License Allocation API, focusing on automating the process of removing or de-allocating licenses from users or devices within an IT environment.
The API automates the license de-allocation process, which involves removing assigned licenses from one or more users or devices. Provides interface to integrate with IT Service Management (ITSM) workflows or other IT automation processes to ensure seamless management of license removals.<br>
**Note:** The response will only include the de-allocations that were not successfully processed, along with the reason for failure. If all de-allocations were successful, deallocationStatuses will be omitted or an empty array (the response body may be an empty JSON object).<br>
**Benefits and Practical Applications:**
* **License Reclamation:** Automates the process of reclaiming licenses from users or devices that no longer require them, optimizing license utilization.
* **Cost Management:** Helps organizations manage costs by ensuring licenses are not wasted on unused.
* **Efficiency:** Reduces administrative overhead and potential errors associated with manual de-allocation processes.
* **Compliance:** Ensures that licenses are properly reclaimed when no longer needed, helping organizations maintain compliance with licensing agreements.
* **Integration:** Integrates with existing ITSM and ITAM workflows, providing a holistic approach to license management and governance.

                    User-based de-allocation JSON request format:
                        [
                            {
                                "samAccountName": "qa.flexera.com\\user1",
                                "userEmailAddress": "user1@flexera.com",
                            },
                            {
                                "samAccountName": "qa.flexera.com\\user2",
                                "userEmailAddress": "user2@flexera.com",
                            }
                        ]

                    User-based de-allocation CSV request format:

                        samAccountName,userEmailAddress
                        qa.flexera.com\user1,user1@flexera.com
                        qa.flexera.com\user2,user2@flexera.com


                    Device-based de-allocation JSON request format:
                        [
                            {
                                "deviceSerial": "GK0L9X1",
                                "deviceName": "test-machine",
                            },
                            {
                                "deviceSerial": "AB0L9X45",
                                "deviceName": "redirection",
                            }
                        ]

                    Device-based de-allocation CSV request format:

                        deviceSerial,deviceName
                        GK0L9X1,test-machine
                        AB0L9X45,redirection

```
flexera-cli fnms license licenses deallocation [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli fnms license licenses deallocation --org-id ORG_ID --license-id LICENSE_ID
  flexera-cli fnms license licenses deallocation --org-id ORG_ID --license-id LICENSE_ID --dry-run
```

### Options

```
      --dry-run[=true|false]   print the planned operation as JSON and exit without calling the API
  -h, --help[=true|false]      help for deallocation
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

