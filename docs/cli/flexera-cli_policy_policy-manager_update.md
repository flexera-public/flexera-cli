## flexera-cli policy policy-manager update

Update a policy manager

### Synopsis

Update a policy manager

Updates the properties of the managed applied policies or the criteria used to determine which child orgs should have a managed applied policy running in them. These changes will be applied to manager immediately and adjustments will be made at the child org level asynchronously.

```
flexera-cli policy policy-manager update [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli policy policy-manager update --org-id ORG_ID --id ID --body @request.json
  flexera-cli policy policy-manager update --org-id ORG_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema policy policy-manager update --example > request.json
```

### Options

```
      --allow-delete-policy      allowDeletePolicy (body); When true, allows deletion of policies at the child organization level.; illustrative example: false
      --allow-edit-policy        allowEditPolicy (body); When true, allows editing of policy options at the child organization level.; illustrative example: false
      --body string              raw JSON body (inline | @file | @-); overrides body field flags
      --body-dry-run             dryRun (body); Flag for testing a policy without taking actions.; illustrative example: false
      --description string       description (body); Human readable description for this specific application of the policy.; illustrative example: "Delete unattached volumes after 24 hours in US-East."
      --dry-run                  print the planned operation as JSON and exit without calling the API
  -h, --help                     help for update
      --id string                id (path, required); ID of the policy manager.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "def01234567890abcdef0123"
  -i, --interactive              edit inputs in a terminal form, review a plan and approve with typed yes
      --log-level string         logLevel (body); Defines the logging level.; enum: ["full","context","error"]; illustrative example: "full"
      --name string              name (body); Name of the policy manager.; illustrative example: "AWS Oversized Instances Recommendations"
      --org-tags-filter string   orgTagsFilter (body); Filter expression for targeting organizations by tags for policy application. ### Operators | Operator | Description | Example | | -------- | ----------- | ------- | | co | Contains | tags co 'service:level' | | eq | Equal | tags eq 'region:us-east' | | ne | Not Equal | tags ne 'service:level:basic' | ### Usage Notes * Operators are case-sensitive (lowercase... (see cli schema); illustrative example: "(tags co 'service:level:basic' or tags eq 'region:us-east')"
      --severity string          severity (body); Severity level of incidents raised by applied policies.; enum: ["low","medium","high","critical"]; illustrative example: "low"
      --skip-approvals           skipApprovals (body); Automatically apply policies without manual approval.; illustrative example: false
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

* [flexera-cli policy policy-manager](flexera-cli_policy_policy-manager.md)	 - Policy Manager operations (generated from the unified OpenAPI spec)

