## flexera-cli bill-analysis org-dashboards create

create org_dashboards

### Synopsis

create org_dashboards

Creates a Dashboard for a given Org.

```
flexera-cli bill-analysis org-dashboards create [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli bill-analysis org-dashboards create --org-id ORG_ID --body @request.json
  flexera-cli bill-analysis org-dashboards create --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema bill-analysis org-dashboards create --example > request.json
```

### Options

```
      --area string         area (body); required by API; UI Area; illustrative example: "bc-index"
      --body string         raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run             print the planned operation as JSON and exit without calling the API
  -h, --help                help for create
  -i, --interactive         edit inputs in a terminal form, review a plan and approve with typed yes
      --name string         name (body); required by API; Display name of dashboard; pattern: "^[\\p{L}0-9._'/(),!@#$%\u0026=+|:? -]*$"; illustrative example: "My DB Cost Summary"
      --scope strings       scope (body); Describes whether this dashboard is scoped to a particular billing center; CLI: comma-separated values or repeated flag; illustrative example: ["wGl5eUB2zV1e8Cb1FM_ZYf"]
      --visibility string   visibility (body); A flag describing whether an org dashboard is public or default; enum: ["default","public"]; API default: "public"; illustrative example: "public"
      --yes                 confirm the operation (required for destructive ops)
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

* [flexera-cli bill-analysis org-dashboards](flexera-cli_bill-analysis_org-dashboards.md)	 - org_dashboards operations (generated from the unified OpenAPI spec)

