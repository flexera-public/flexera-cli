## flexera-cli finops-customizations report-subscriptions create

Create a report subscription

### Synopsis

Create a report subscription

Creates a report subscription in the organization. The owner is derived from the JWT subject claim; the caller becomes the owner.

```
flexera-cli finops-customizations report-subscriptions create [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-customizations report-subscriptions create --org-id ORG_ID --body @request.json
  flexera-cli finops-customizations report-subscriptions create --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-customizations report-subscriptions create --example > request.json
```

### Options

```
      --body string              raw JSON body (inline | @file | @-); overrides body field flags
      --dashboard-id string      dashboardId (body); required by API; Identifier of the dashboard to render and send.; illustrative example: "dash-8f21"
      --dashboard-scope string   dashboardScope (body); required by API; Which dashboard store resolves the target dashboard.; enum: ["user","org"]; illustrative example: "org"
      --dry-run                  print the planned operation as JSON and exit without calling the API
      --enabled                  enabled (body); Whether the schedule is active.; API default: true; illustrative example: false
  -h, --help                     help for create
  -i, --interactive              edit inputs in a terminal form, review a plan and approve with typed yes
      --name string              name (body); required by API; Human-readable name for the subscription; maxLength: 255; illustrative example: "Weekly Cloud Cost"
      --visibility string        visibility (body); Visibility scope.; enum: ["private","shared"]; API default: "private"; illustrative example: "private"
      --yes                      confirm the operation (required for destructive ops)
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

* [flexera-cli finops-customizations report-subscriptions](flexera-cli_finops-customizations_report-subscriptions.md)	 - Report Subscriptions operations (generated from the unified OpenAPI spec)

