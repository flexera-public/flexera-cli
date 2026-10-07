## flexera-cli policy-manager update

Update a policy manager

```
flexera-cli policy-manager update [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli policy-manager update --org-id ORG_ID --id ID --body @request.json
  flexera-cli policy-manager update --org-id ORG_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema policy-manager update --example > request.json
```

### Options

```
      --allow-delete-policy      allowDeletePolicy (body)
      --allow-edit-policy        allowEditPolicy (body)
      --body string              raw JSON body (inline | @file | @-); overrides body field flags
      --body-dry-run             dryRun (body)
      --description string       description (body)
      --dry-run                  print the planned operation as JSON and exit without calling the API
  -h, --help                     help for update
      --id string                id (path, required)
  -i, --interactive              edit inputs in a terminal form, review a plan and approve with typed yes
      --log-level string         logLevel (body)
      --name string              name (body)
      --org-tags-filter string   orgTagsFilter (body)
      --severity string          severity (body)
      --skip-approvals           skipApprovals (body)
      --yes                      confirm the operation (required for destructive ops)
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

* [flexera-cli policy-manager](flexera-cli_policy-manager.md)	 - Policy Manager operations (generated from the unified OpenAPI spec)

