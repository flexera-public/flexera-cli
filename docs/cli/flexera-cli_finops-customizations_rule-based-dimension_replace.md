## flexera-cli finops-customizations rule-based-dimension replace

Creates/Replace a rules list

### Synopsis

Creates/Replace a rules list

Create or replace a single rules list for the given rule-based dimension and effectiveAt date.

```
flexera-cli finops-customizations rule-based-dimension replace [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-customizations rule-based-dimension replace --org-id ORG_ID --id ID --effective-at EFFECTIVE_AT --body @request.json
  flexera-cli finops-customizations rule-based-dimension replace --org-id ORG_ID --id ID --effective-at EFFECTIVE_AT --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-customizations rule-based-dimension replace --example > request.json
```

### Options

```
      --body string           raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run               print the planned operation as JSON and exit without calling the API
      --effective-at string   effectiveAt (path, required); The date (year-month) when this rules list takes effect, superseding any previous list. The list remains in effect until a subsequent list is defined to take effect at a later date.; required by API; pattern: "^\\d{4}-\\d{2}$"; illustrative example: "2023-01"
  -h, --help                  help for replace
      --id string             id (path, required); ID of the rule-based dimension; required by API; pattern: "^rbd_[\\S]*$"; illustrative example: "rbd_department"
  -i, --interactive           edit inputs in a terminal form, review a plan and approve with typed yes
      --yes                   confirm the operation (required for destructive ops)
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

* [flexera-cli finops-customizations rule-based-dimension](flexera-cli_finops-customizations_rule-based-dimension.md)	 - Rule-Based Dimension operations (generated from the unified OpenAPI spec)

