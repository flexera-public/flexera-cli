## flexera-cli saas license allocations-all-2

Delete allocation

### Synopsis

Delete allocation

Delete an existing allocation.

```
flexera-cli saas license allocations-all-2 [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas license allocations-all-2 --org-id ORG_ID --license-id LICENSE_ID --term-id TERM_ID --purchase-id PURCHASE_ID --id ID
  flexera-cli saas license allocations-all-2 --org-id ORG_ID --license-id LICENSE_ID --term-id TERM_ID --purchase-id PURCHASE_ID --id ID --dry-run
```

### Options

```
      --dry-run              print the planned operation as JSON and exit without calling the API
  -h, --help                 help for allocations-all-2
      --id string            id (path, required); Unique identifier of the allocation.; required by API; illustrative example: "3421"
      --license-id string    licenseId (path, required); Unique identifier of the license.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "38932"
      --purchase-id string   purchaseId (path, required); Unique identifier of the purchase associated with the license term.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "93214"
      --term-id string       termId (path, required); Unique identifier of the license term associated with the license.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "42214"
      --yes                  confirm the operation (required for destructive ops)
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

