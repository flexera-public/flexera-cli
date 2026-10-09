## flexera-cli risk notifications create

Create Notification Policy

### Synopsis

Create Notification Policy

Create a new notification policy for security alerts via gRPC backend service.

```
flexera-cli risk notifications create [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli risk notifications create --org-id ORG_ID --body @request.json
  flexera-cli risk notifications create --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema risk notifications create --example > request.json
```

### Options

```
      --body string                raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run                    print the planned operation as JSON and exit without calling the API
      --email-ids strings          emailIds (body); required by API; List of email addresses; minItems: 1; CLI: comma-separated values or repeated flag
      --expiry-time int            expiryTime (body); required by API; Expiry time in hours; minimum: 0
  -h, --help                       help for create
      --integration-tool string    integrationTool (body); required by API; Integration tool (e.g., 'email'); minLength: 1
  -i, --interactive                edit inputs in a terminal form, review a plan and approve with typed yes
      --interval-days int          intervalDays (body); required by API; Interval in days (must be 1, 7, 14, or 30); enum: [1,7,14,30]
      --notification-type string   notificationType (body); required by API; Notification type; enum: ["MISCONFIG_ALERT","VULN_ALERT"]
      --password string            password (body); required by API; Password
      --policy-name string         policyName (body); required by API; Policy name; minLength: 1
      --user-id string             userId (body); required by API; User ID; minLength: 1
      --user-name string           userName (body); User name (optional); API default: ""
      --yes                        confirm the operation (required for destructive ops)
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

* [flexera-cli risk notifications](flexera-cli_risk_notifications.md)	 - notifications operations (generated from the unified OpenAPI spec)

