## flexera-cli saas license notes

Update note

### Synopsis

Update note

Update the details of a note associated to resource.

```
flexera-cli saas license notes [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas license notes --org-id ORG_ID --license-id LICENSE_ID --id ID --body @request.json
  flexera-cli saas license notes --org-id ORG_ID --license-id LICENSE_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema saas license notes --example > request.json
```

### Options

```
      --body string         raw JSON body (inline | @file | @-); overrides body field flags
      --details string      details (body); required by API; Details of the note.; maxLength: 500; illustrative example: "license had a discount of 10."
      --dry-run             print the planned operation as JSON and exit without calling the API
  -h, --help                help for notes
      --id string           id (path, required); Unique identifier of note.; required by API; illustrative example: "3421"
  -i, --interactive         edit inputs in a terminal form, review a plan and approve with typed yes
      --license-id string   licenseId (path, required); Unique identifier of license that note is associated to.; required by API; illustrative example: "34521"
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

* [flexera-cli saas license](flexera-cli_saas_license.md)	 - License operations (generated from the unified OpenAPI spec)

