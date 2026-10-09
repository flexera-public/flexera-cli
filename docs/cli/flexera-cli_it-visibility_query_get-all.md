## flexera-cli it-visibility query get-all

Execute a query on the delta export data

### Synopsis

Execute a query on the delta export data

Executes the specified query on the delta export data. Supports optional gzip compression for output files by including the query parameter "outputCompression=gzip".

```
flexera-cli it-visibility query get-all [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli it-visibility query get-all --org-id ORG_ID --body @request.json
Validated illustrative body, when available (review before use):
  flexera-cli cli schema it-visibility query get-all --example > request.json
```

### Options

```
      --body string                 raw JSON body (inline | @file | @-); overrides body field flags
  -h, --help                        help for get-all
      --output-compression string   outputCompression (query); If set to 'gzip', the server will return a gzip-compressed CSV. Supported value: gzip; enum: ["gzip"]; illustrative example: "gzip"
      --query-name string           queryName (body); required by API; The query to execute.; enum: ["hardware_business_services","hardware_contextualized","hardware_evidence","hardware_inventory","hardware_inventory_source","hardware_technopedia","hardware_technopedia_lifecycle"... (see cli schema); illustrative example: "software_inventory"
      --resume-token string         resumeToken (query); An opaque token to be provided when requesting a subsequent set of delta changes for the same query.
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

* [flexera-cli it-visibility query](flexera-cli_it-visibility_query.md)	 - Query operations (generated from the unified OpenAPI spec)

