## flexera-cli risk misconfiguration-ui list

Risk Misconfiguration Rules List.

### Synopsis

Risk Misconfiguration Rules List.

Proxy endpoint that forwards api request for list of misconfiguration rules to the Secops UI API.

```
flexera-cli risk misconfiguration-ui list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli risk misconfiguration-ui list --org-id ORG_ID --body @request.json
Validated illustrative body, when available (review before use):
  flexera-cli cli schema risk misconfiguration-ui list --example > request.json
```

### Options

```
      --accounts strings             accounts (body); required by API; List of cloud account IDs or 'all'; minItems: 1; CLI: comma-separated values or repeated flag; illustrative example: ["837570591364","252277358118"]
      --body string                  raw JSON body (inline | @file | @-); overrides body field flags
      --compliance-standard string   complianceStandard (body); Required when featureType is 'compliance'; minLength: 1; illustrative example: "SOC_2"
      --control-id string            controlId (body); Required when featureType is 'compliance'; minLength: 1; illustrative example: "CC7.1"
      --date string                  date (body); Required when featureType is 'compliance'. Format: YYYY-MM-DD; minLength: 1; illustrative example: "2025-11-04"
      --event-id string              eventId (body); Required when featureType is 'event'; minLength: 1; illustrative example: "Configuration Change"
      --event-time string            eventTime (body); Required when featureType is 'event'. Format: YYYY-MM-DD HH:MM:SS; minLength: 1; illustrative example: "2025-06-15 14:30:00"
      --feature-type string          featureType (body); required by API; Feature type to filter by. Possible values: risk, compliance, event, policy; enum: ["risk","compliance","event","policy"]; illustrative example: "risk"
  -h, --help                         help for list
      --providers strings            providers (body); required by API; List of cloud providers or ['all']. Allowed values: 'aws', 'azure'; minItems: 1; CLI: comma-separated values or repeated flag; illustrative example: ["aws","azure"]
      --regions strings              regions (body); required by API; List of cloud regions to filter by or 'all'; minItems: 1; CLI: comma-separated values or repeated flag; illustrative example: ["us-east-1","us-east-2"]
      --services strings             services (body); required by API; List of cloud services to filter by or 'all'; minItems: 1; CLI: comma-separated values or repeated flag; illustrative example: ["ec2","rds","s3"]
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

