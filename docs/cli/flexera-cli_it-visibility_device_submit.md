## flexera-cli it-visibility device submit

Delete device evidence

### Synopsis

Delete device evidence

Submits a request to delete device evidence for the specified device IDs.

A datasource is included if it matches any datasource selector. If no datasource selectors are provided, evidence from every datasource associated with the specified devices is included.

Submitting a deletion request does not guarantee that all selected evidence will be deleted because inventory may change while the request is processed. If a device remains active in the environment, its evidence may reappear after new inventory data is collected.

Repeating a deletion request is allowed and does not return an error when the selected evidence has already been deleted.

### Client guidance
* Submit no more than 500 unique device IDs per request. This recommendation is not enforced, but larger requests provide no processing benefit and may increase latency or timeout risk.
* Split larger workloads into batches and process one batch at a time per organization. Use the same datasource selectors for every batch in one logical operation.
* There is currently no specific concurrency limit or hard cap in place. We recommend that customers test and tune their request volume based on their network conditions, performance requirements, and business needs.
* The API does not automatically retry failed deletion requests. Clients must not automatically retry timed-out requests because inventory may change and the original deletion may already have been submitted.

```
flexera-cli it-visibility device submit [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli it-visibility device submit --org-id ORG_ID --body @request.json
  flexera-cli it-visibility device submit --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema it-visibility device submit --example > request.json
```

### Options

```
      --body string          raw JSON body (inline | @file | @-); overrides body field flags
      --device-ids strings   deviceIds (body); required by API; Device IDs whose evidence should be included. Send no more than 500 unique IDs per request; this recommendation is not enforced by the API.; CLI: comma-separated values or repeated flag; illustrative example: ["4e0e3471-ccbc-41c7-a21e-c384f2eb5669"]
      --dry-run              print the planned operation as JSON and exit without calling the API
  -h, --help                 help for submit
  -i, --interactive          edit inputs in a terminal form, review a plan and approve with typed yes
      --yes                  confirm the operation (required for destructive ops)
```

### Options inherited from parent commands

```
      --access-token string     static bearer access token
      --api-base-url string     override API base URL
      --client-id string        OAuth client ID
      --client-secret string    OAuth client secret
      --config string           config file (default $HOME/.flexera/config.yaml)
  -d, --debug                   log HTTP requests/responses to stderr (Authorization redacted)
      --json-style string       JSON whitespace style (auto|pretty|compact) (default "auto")
      --login-base-url string   override login base URL
      --no-validate             skip API schema constraints (never JSON syntax or request data-loss checks)
      --org-id int              organization ID
      --out-fields string       project JSON output fields (comma-separated paths)
      --out-jq string           shape JSON output with a jq expression
  -o, --output string           output format (json|table)
  -r, --raw-output              write jq string results without JSON quotes (requires --out-jq)
      --refresh-token string    OAuth refresh token
      --zone string             API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli it-visibility device](flexera-cli_it-visibility_device.md)	 - Device operations (generated from the unified OpenAPI spec)

