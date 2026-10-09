## flexera-cli policy policy-manager update-template

Update policy manager to reference the latest published template version

### Synopsis

Update policy manager to reference the latest published template version

Updates the policy manager to use the latest version of its associated published template. The request can also include configuration updates such as schedule, log level, dry-run mode, severity, approval behavior, credentials, and options, with only the provided fields being updated while others remain unchanged. Updates to the options and credentials may be required if the newer version of the template introduces changes that make the previous configurations incompatible. The policy manager switches to the latest template immediately, but changes are gradually rolled out to child organizations, so updates may take some time to fully apply everywhere. If the policy manager is already using the latest template version, no changes are applied and a success response is returned.

```
flexera-cli policy policy-manager update-template [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli policy policy-manager update-template --org-id ORG_ID --id ID --body @request.json
  flexera-cli policy policy-manager update-template --org-id ORG_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema policy policy-manager update-template --example > request.json
```

### Options

```
      --body string        raw JSON body (inline | @file | @-); overrides body field flags
      --body-dry-run       dryRun (body); Flag for testing a policy without taking actions.; illustrative example: false
      --dry-run            print the planned operation as JSON and exit without calling the API
  -h, --help               help for update-template
      --id string          id (path, required); ID of the policy manager.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "def01234567890abcdef0123"
  -i, --interactive        edit inputs in a terminal form, review a plan and approve with typed yes
      --log-level string   logLevel (body); Defines the logging level.; enum: ["full","context","error"]; illustrative example: "full"
      --severity string    severity (body); Severity level of incidents raised by applied policies.; enum: ["low","medium","high","critical"]; illustrative example: "low"
      --skip-approvals     skipApprovals (body); Automatically apply policies without manual approval.; illustrative example: false
      --yes                confirm the operation (required for destructive ops)
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

