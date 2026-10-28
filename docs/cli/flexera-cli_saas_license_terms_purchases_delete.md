## flexera-cli saas license terms purchases delete

Delete purchase

### Synopsis

Delete purchase

Deletes a purchase associated with license term.

```
flexera-cli saas license terms purchases delete [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas license terms purchases delete --org-id ORG_ID --license-id LICENSE_ID --term-id TERM_ID --id ID
  flexera-cli saas license terms purchases delete --org-id ORG_ID --license-id LICENSE_ID --term-id TERM_ID --id ID --dry-run
```

### Options

```
      --dry-run             print the planned operation as JSON and exit without calling the API
  -h, --help                help for delete
      --id string           id (path, required); Identifies a purchase associated with license term ID.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "38342"
      --license-id string   licenseId (path, required); Unique identifier of license.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "38932"
      --term-id string      termId (path, required); Identifies a license term associated with the license.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "42214"
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

* [flexera-cli saas license terms purchases](flexera-cli_saas_license_terms_purchases.md)	 - purchases operations

