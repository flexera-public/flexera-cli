## flexera-cli saas license purchases

Update purchase

### Synopsis

Update purchase

Updates a purchase for given purchase ID.

```
flexera-cli saas license purchases [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas license purchases --org-id ORG_ID --license-id LICENSE_ID --term-id TERM_ID --id ID --body @request.json
  flexera-cli saas license purchases --org-id ORG_ID --license-id LICENSE_ID --term-id TERM_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema saas license purchases --example > request.json
```

### Options

```
      --amount float            amount (body); Price associated with license purchase.; format: double; illustrative example: 123.123
      --body string             raw JSON body (inline | @file | @-); overrides body field flags
      --currency string         currency (body); Currency type for the amount.; illustrative example: "USD"
      --dry-run                 print the planned operation as JSON and exit without calling the API
      --effective-at string     effectiveAt (body); Start date associated with license term.; illustrative example: "2021-02-01 15:19:33"
      --ends-at string          endsAt (body); End date associated with license term.; illustrative example: "2021-12-31 23:59:59"
      --frequency-type string   frequencyType (body); Frequency of a payment for a license.; enum: ["once","monthly","quarterly","semiannually","annually","every18months","biannually","triennially"]; illustrative example: "monthly"
  -h, --help                    help for purchases
      --id string               id (path, required); Identifies a purchase associated with the license term.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "42214"
  -i, --interactive             edit inputs in a terminal form, review a plan and approve with typed yes
      --license-id string       licenseId (path, required); Unique identifier of license.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "38932"
      --purchase-id string      purchaseId (body); Unique ID associated to license purchase transaction.; illustrative example: "a3355453-86ab-4628-9670-6adb9ff2e444"
      --purchased-at string     purchasedAt (body); License term purchased date.; illustrative example: "2021-01-01 15:19:33"
      --term-id string          termId (path, required); Identifies a license term associated with the license.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "42214"
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

* [flexera-cli saas license](flexera-cli_saas_license.md)	 - License operations (generated from the unified OpenAPI spec)

