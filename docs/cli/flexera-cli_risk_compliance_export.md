## flexera-cli risk compliance export

Export Compliance Chart Data.

### Synopsis

Export Compliance Chart Data.

Exports compliance chart data in the specified format (xlsx, csv, pdf).

```
flexera-cli risk compliance export [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli risk compliance export --org-id ORG_ID --body @request.json
Validated illustrative body, when available (review before use):
  flexera-cli cli schema risk compliance export --example > request.json
```

### Options

```
      --accounts strings             accounts (body); required by API; List of account IDs. Use 'all' to include all accounts.; minItems: 1; CLI: comma-separated values or repeated flag; illustrative example: ["837570591364","252277358118"]
      --body string                  raw JSON body (inline | @file | @-); overrides body field flags
      --category strings             category (body); List of control categories to filter by; CLI: comma-separated values or repeated flag; illustrative example: ["Identity And Access Management","Storage","Logging","Monitoring","Networking"]
      --compliance-standard string   complianceStandard (body); Compliance standard name (e.g., 'NIST_800-82_rev2'). If not provided, summary export will be generated.; illustrative example: "NIST_800-82_rev2"
      --format string                format (body); required by API; Export format - 'json' or 'xlsx'; illustrative example: "xlsx"
  -h, --help                         help for export
      --imc                          imc (body); required by API; IMC filter flag; illustrative example: false
      --level int                    level (body); CIS level filter. Allowed values: 1 or 2; API default: 1; illustrative example: 1
      --providers strings            providers (body); required by API; List of cloud providers. Allowed values: 'aws', 'azure', 'all'; minItems: 1; CLI: comma-separated values or repeated flag; illustrative example: ["aws","azure"]
      --regions strings              regions (body); required by API; List of regions. Use 'all' to include all regions.; minItems: 1; CLI: comma-separated values or repeated flag; illustrative example: ["us-east-1","us-east-2","us-west-1","us-west-2"]
      --services strings             services (body); required by API; List of services. Use 'all' to include all services.; minItems: 1; CLI: comma-separated values or repeated flag; illustrative example: ["cloudfront","cloudtrail","cloudwatch","config","dynamodb","ec2","ecs","eks","elb","elbv2","iam","kms","lambda","rds","s3","ses","sns","sqs","vpc"]
      --snapback string              snapback (body); required by API; Snapback date in YYYY-MM-DD format; illustrative example: "2025-11-02"
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

* [flexera-cli risk compliance](flexera-cli_risk_compliance.md)	 - compliance operations (generated from the unified OpenAPI spec)

