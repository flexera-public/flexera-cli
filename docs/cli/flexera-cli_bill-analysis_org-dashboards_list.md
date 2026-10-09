## flexera-cli bill-analysis org-dashboards list

index org_dashboards

### Synopsis

index org_dashboards

Lists all Dashboards for a given Org.
Requires user to have `optima:public_dashboard:index` on a billing center within the org.

```
flexera-cli bill-analysis org-dashboards list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli bill-analysis org-dashboards list --org-id ORG_ID
```

### Options

```
      --area string         area (query); Optional area filter; illustrative example: "bc-index"
  -h, --help                help for list
      --visibility string   visibility (query); Optional visibility filter; enum: ["default","public"]; illustrative example: "default"
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

