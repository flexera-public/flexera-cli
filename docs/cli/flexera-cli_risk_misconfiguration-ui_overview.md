## flexera-cli risk misconfiguration-ui overview

Risk Misconfiguration Overview.

### Synopsis

Risk Misconfiguration Overview.

Proxy endpoint that forwards api request for list of misconfiguration overview to the Secops UI API.

```
flexera-cli risk misconfiguration-ui overview [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli risk misconfiguration-ui overview --org-id ORG_ID --body @request.json
Validated illustrative body, when available (review before use):
  flexera-cli cli schema risk misconfiguration-ui overview --example > request.json
```

### Options

```
      --accounts strings    accounts (body); required by API; List of cloud account IDs or 'all'; minItems: 1; CLI: comma-separated values or repeated flag; illustrative example: ["837570591364","252277358118"]
      --body string         raw JSON body (inline | @file | @-); overrides body field flags
  -h, --help                help for overview
      --providers strings   providers (body); required by API; List of cloud providers or ['all']. Allowed values: 'aws', 'azure'; minItems: 1; CLI: comma-separated values or repeated flag; illustrative example: ["aws","azure"]
      --regions strings     regions (body); required by API; List of cloud regions to filter by or 'all'; minItems: 1; CLI: comma-separated values or repeated flag; illustrative example: ["us-east-1","us-east-2"]
      --services strings    services (body); required by API; List of cloud services to filter by or 'all'; minItems: 1; CLI: comma-separated values or repeated flag; illustrative example: ["ec2","rds","s3"]
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

* [flexera-cli risk misconfiguration-ui](flexera-cli_risk_misconfiguration-ui.md)	 - misconfiguration-ui operations (generated from the unified OpenAPI spec)

