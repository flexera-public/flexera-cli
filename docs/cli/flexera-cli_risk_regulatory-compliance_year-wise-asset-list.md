## flexera-cli risk regulatory-compliance year-wise-asset-list

Regulatory Compliance Year-Wise Asset List

### Synopsis

Regulatory Compliance Year-Wise Asset List

Returns a paginated list of assets grouped by end-of-support year.

```
flexera-cli risk regulatory-compliance year-wise-asset-list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli risk regulatory-compliance year-wise-asset-list --org-id ORG_ID --body @request.json
Validated illustrative body, when available (review before use):
  flexera-cli cli schema risk regulatory-compliance year-wise-asset-list --example > request.json
```

### Options

```
      --body string       raw JSON body (inline | @file | @-); overrides body field flags
      --fields strings    fields (body); Fields to include in the response.; CLI: comma-separated values or repeated flag
      --filter string     filter (body); Filter expression like 'manufacturer co "Microsoft"'; API default: ""
  -h, --help              help for year-wise-asset-list
      --order-by string   orderBy (body); Order expression like 'manufacturer asc'; API default: ""
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

* [flexera-cli risk regulatory-compliance](flexera-cli_risk_regulatory-compliance.md)	 - regulatory-compliance operations (generated from the unified OpenAPI spec)

