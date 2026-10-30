## flexera-cli it-visibility device delete estimate

Estimate device evidence deletion

### Synopsis

Estimate device evidence deletion

Estimates the device evidence that would be deleted for the specified device IDs, grouped by datasource.

A datasource is included if it matches any datasource selector. If no datasource selectors are provided, evidence from every datasource associated with the specified devices is included.

The estimate reflects the inventory available when the request is made. Because inventory may change, it does not guarantee that a later request will delete the same evidence.

### Client guidance
* Submit no more than 500 unique device IDs per request. This recommendation is not enforced, but larger requests provide no processing benefit and may increase latency or timeout risk.
* Split larger workloads into batches and process them sequentially. Use the same datasource selectors for every batch in one logical operation.
* There is currently no specific concurrency limit or hard cap in place. We recommend that customers test and tune their request volume based on their network conditions, performance requirements, and business needs.
* The API does not automatically retry failed estimate requests. Clients may retry transient failures using exponential backoff with jitter.

```
flexera-cli it-visibility device delete estimate [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli it-visibility device delete estimate --org-id ORG_ID --body @request.json
Validated illustrative body, when available (review before use):
  flexera-cli cli schema it-visibility device delete estimate --example > request.json
```

### Options

```
      --body string          raw JSON body (inline | @file | @-); overrides body field flags
      --device-ids strings   deviceIds (body); required by API; Device IDs whose evidence should be included. Send no more than 500 unique IDs per request; this recommendation is not enforced by the API.; CLI: comma-separated values or repeated flag; illustrative example: ["4e0e3471-ccbc-41c7-a21e-c384f2eb5669"]
  -h, --help                 help for estimate
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

* [flexera-cli it-visibility device delete](flexera-cli_it-visibility_device_delete.md)	 - delete operations

