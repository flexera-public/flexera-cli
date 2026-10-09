## flexera-cli policy policy-aggregate update

Update a policy aggregate

### Synopsis

Update a policy aggregate

Update enables users to change attributes of the managed applied policies and to add or remove projects to manage
        applied policies in. Changes to the list of projects will result in policies being created or terminated. Changes to configuration
        options will result in all managed policies being updated. This will happen asynchronously. Changes to the the schedule will result in
        applied policies being rescheduled.

```
flexera-cli policy policy-aggregate update [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli policy policy-aggregate update --org-id ORG_ID --policy-aggregate-id POLICY_AGGREGATE_ID --body @request.json
  flexera-cli policy policy-aggregate update --org-id ORG_ID --policy-aggregate-id POLICY_AGGREGATE_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema policy policy-aggregate update --example > request.json
```

### Options

```
      --body string                  raw JSON body (inline | @file | @-); overrides body field flags
      --body-dry-run                 dryRun (body); dryRun is a flag used for testing a policy so that an incident can be raised without performing an action.; API default: false; illustrative example: false
      --description string           description (body); description provides a human readable description for this specific application of the policy.; minLength: 1; maxLength: 256; illustrative example: "Delete unattached volumes after 24 hours in US-East."
      --dry-run                      print the planned operation as JSON and exit without calling the API
  -h, --help                         help for update
  -i, --interactive                  edit inputs in a terminal form, review a plan and approve with typed yes
      --log-level string             logLevel (body); logLevel defines the amount of information written to log. It must be either 'full', to log everything, 'context', to log everything except the body responses, or 'error', to only log errors and the information on when the evaluation started and ended.; enum: ["full","context","error"]; API default: "full"; illustrative example: "full"
      --name string                  name (body); name provides a name for this specific application of the policy.; minLength: 1; maxLength: 128; illustrative example: "us_east_unattached_volumes"
      --policy-aggregate-id string   policyAggregateId (path, required); The unique identifier for the policy aggregate; required by API; illustrative example: "5b06ead5e0dacc007058c784"
      --severity string              severity (body); Severity describes the excepted level of concern / urgency which incidents raised by the applied policy should be addressed within the organization. This is set at the template level, but can be set to a different value during policy application. If a value isn't passed during policy application, the template's value will be used.; enum: ["low","medium","high","critical"]; illustrative example: "low"
      --skip-approvals               skipApprovals (body); When enabled, the policy will bypass any manual approval steps and automatically take actions as defined in the policy. This feature is designed to streamline the policy enforcement process and reduce the need for manual intervention. However, enabling automatic actions may result in changes being applied without human review.; API default: false; illustrative example: false
      --yes                          confirm the operation (required for destructive ops)
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

* [flexera-cli policy policy-aggregate](flexera-cli_policy_policy-aggregate.md)	 - Policy Aggregate operations (generated from the unified OpenAPI spec)

