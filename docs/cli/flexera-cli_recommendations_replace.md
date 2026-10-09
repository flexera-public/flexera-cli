## flexera-cli recommendations replace

updateStatus Recommendations

### Synopsis

updateStatus Recommendations

Update the status of a recommendation

**Required security scopes for GlobalSession**:
  * `common:org:affiliated+common:org:own`

```
flexera-cli recommendations replace [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli recommendations replace --org-id ORG_ID --body @request.json
  flexera-cli recommendations replace --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema recommendations replace --example > request.json
```

### Options

```
      --body string                  raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run                      print the planned operation as JSON and exit without calling the API
  -h, --help                         help for replace
      --id string                    id (body); required by API; Recommendation unique ID; illustrative example: "676f0f2c838a081da3f7cd61_vol-06fa6937d41c139eb"
  -i, --interactive                  edit inputs in a terminal form, review a plan and approve with typed yes
      --snoozed-target-date string   snoozedTargetDate (body); Target date to move recommendation from snoozed to active again; illustrative example: "2025-01-30"
      --status string                status (body); required by API; New status of recommendation; enum: ["active","snoozed","rejected","realized"]; illustrative example: "realized"
      --status-reason string         statusReason (body); Reason of new status of recommendation; illustrative example: "Recommendation isn't relevant right now"
      --yes                          confirm the operation (required for destructive ops)
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

* [flexera-cli recommendations](flexera-cli_recommendations.md)	 - Optima Recommendations API

