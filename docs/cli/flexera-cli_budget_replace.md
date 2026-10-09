## flexera-cli budget replace

Updates a budget

### Synopsis

Updates a budget

Replaces a specific budget for the given organization and budget ID.

```
flexera-cli budget replace [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli budget replace --org-id ORG_ID --id ID --body @request.json
  flexera-cli budget replace --org-id ORG_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema budget replace --example > request.json
```

### Options

```
      --body string             raw JSON body (inline | @file | @-); overrides body field flags
      --dimensions strings      dimensions (body); required by API; Dimensions used to break down this budget. Each budget segment will be defined as a unique combination of these dimension values.; CLI: comma-separated values or repeated flag; illustrative example: ["bc_level_1","ProviderName"]
      --dry-run                 print the planned operation as JSON and exit without calling the API
  -h, --help                    help for replace
      --id string               id (path, required); Identifier of the budget; required by API; illustrative example: "2ed7db3"
  -i, --interactive             edit inputs in a terminal form, review a plan and approve with typed yes
      --metric string           metric (body); required by API; The cost metric used for the budget; enum: ["cost_nonamortized_unblended_adj","cost_amortized_unblended_adj","cost_nonamortized_blended_adj","cost_amortized_blended_adj","BilledCost","ModifiedBilledCost","EffectiveCost","Mo... (see cli schema); API default: "cost_amortized_unblended_adj"; illustrative example: "BilledCost"
      --name string             name (body); required by API; A descriptive name to uniquely identify the budget; maxLength: 70; illustrative example: "Engineering Budget"
      --year-months "2023-01"   yearMonths (body); required by API; An array of year-month strings like "2023-01", defining the period covered by the budget.; CLI: comma-separated values or repeated flag; illustrative example: ["2023-01"]
      --yes                     confirm the operation (required for destructive ops)
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

* [flexera-cli budget](flexera-cli_budget.md)	 - Budget API

