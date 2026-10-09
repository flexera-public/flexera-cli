## flexera-cli policy policy-manager create

Create a Policy Manager

### Synopsis

Create a Policy Manager

Creates a Policy Manager. After the manager is created an asynchronous process will start and discover which child orgs the manager will manage applied policies in. It will then create applied policies in them. The operation is asynchronous, and will continue in the background after the request is accepted. It may take some time to start. The status of the managed applied policies can be accessed by viewing the [summary](#/Policy%20Manager/Policy%20Manager%23summary) of the policy manager. This will also include information about applied policies the manager was unable to create and what the error was. An example might be the creator of the manager does not have the needed permissions in the child org to create an applied policy there.

```
flexera-cli policy policy-manager create [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli policy policy-manager create --org-id ORG_ID --body @request.json
  flexera-cli policy policy-manager create --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema policy policy-manager create --example > request.json
```

### Options

```
      --allow-delete-policy      allowDeletePolicy (body); When true, allows deletion of policies at the child organization level.; illustrative example: false
      --allow-edit-policy        allowEditPolicy (body); When true, allows editing of policy options at the child organization level.; illustrative example: false
      --body string              raw JSON body (inline | @file | @-); overrides body field flags
      --body-dry-run             dryRun (body); Flag for testing a policy without taking actions.; illustrative example: false
      --description string       description (body); Human readable description for this specific application of the policy.; illustrative example: "Delete unattached volumes after 24 hours in US-East."
      --dry-run                  print the planned operation as JSON and exit without calling the API
  -h, --help                     help for create
  -i, --interactive              edit inputs in a terminal form, review a plan and approve with typed yes
      --log-level string         logLevel (body); Defines the logging level.; enum: ["full","context","error"]; illustrative example: "full"
      --name string              name (body); required by API; Name of the policy manager.; illustrative example: "AWS Oversized Instances Recommendations"
      --org-tags-filter string   orgTagsFilter (body); Filter expression for targeting organizations by tags for policy application. ### Operators | Operator | Description | Example | | -------- | ----------- | ------- | | co | Contains | tags co 'service:level' | | eq | Equal | tags eq 'region:us-east' | | ne | Not Equal | tags ne 'service:level:basic' | ### Usage Notes * Operators are case-sensitive (lowercase... (see cli schema); illustrative example: "(tags co 'service:level:basic' or tags eq 'region:us-east')"
      --severity string          severity (body); Severity level of incidents raised by applied policies.; enum: ["low","medium","high","critical"]; illustrative example: "low"
      --skip-approvals           skipApprovals (body); Automatically apply policies without manual approval.; illustrative example: false
      --template-ref string      templateRef (body); required by API; The \"templateRef\" is used to specify the published template that is to be applied. This field accepts a published template reference like: \"ref::3::policy:published-template:5b06ead5e0dacc007058c784\".; illustrative example: "ref:nam:1234::policy:published-template:d41d8cd98f00b204e9800998ecf8427e"
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

