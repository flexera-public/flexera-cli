## flexera-cli saas license fees

Update license agreement fee

### Synopsis

Update license agreement fee

Updates a license fee associated with license agreement.

```
flexera-cli saas license fees [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas license fees --org-id ORG_ID --license-id LICENSE_ID --fee-id FEE_ID --body @request.json
  flexera-cli saas license fees --org-id ORG_ID --license-id LICENSE_ID --fee-id FEE_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema saas license fees --example > request.json
```

### Options

```
      --body string                raw JSON body (inline | @file | @-); overrides body field flags
      --currency string            currency (body); Type of currency.; illustrative example: "USD"
      --dry-run                    print the planned operation as JSON and exit without calling the API
      --effective-at string        effectiveAt (body); Start date associated with license agreement with time varying price.; illustrative example: "2019-10-15T16:05:36.1000Z"
      --ends-at string             endsAt (body); End date associated with license agreement with time varying price.; illustrative example: "2018-12-30T00:00:00.1000Z"
      --fee-id string              feeId (path, required); Unique ID associated with license agreement fee.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "32123"
      --frequency-type string      frequencyType (body); Frequency of a payment for miscellaneous fees.; enum: ["once","monthly","quarterly","semiannually","annually","every18months","biannually","triennially"]; illustrative example: "semiannually"
  -h, --help                       help for fees
  -i, --interactive                edit inputs in a terminal form, review a plan and approve with typed yes
      --license-id string          licenseId (path, required); Unique identifier of license.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "38932"
      --miscellaneous-fees float   miscellaneousFees (body); Miscellaneous fee associated with a license agreement.; format: double; illustrative example: 213.5
      --name string                name (body); Name of the fee to indicate the miscellaneous spend associated with licenses.; illustrative example: "Workday"
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

* [flexera-cli saas license](flexera-cli_saas_license.md)	 - License operations (generated from the unified OpenAPI spec)

