## flexera-cli policy policy-template

Policy templates (project-scoped)

```
flexera-cli policy policy-template [flags]
```

### Options

```
  -h, --help   help for policy-template
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

* [flexera-cli policy](flexera-cli_policy.md)	 - Policy API
* [flexera-cli policy policy-template create](flexera-cli_policy_policy-template_create.md)	 - Create a policy template
* [flexera-cli policy policy-template delete](flexera-cli_policy_policy-template_delete.md)	 - Delete a policy template
* [flexera-cli policy policy-template evaluate](flexera-cli_policy_policy-template_evaluate.md)	 - Evaluate a policy template
* [flexera-cli policy policy-template get](flexera-cli_policy_policy-template_get.md)	 - Show a policy template
* [flexera-cli policy policy-template list](flexera-cli_policy_policy-template_list.md)	 - List policy templates
* [flexera-cli policy policy-template update](flexera-cli_policy_policy-template_update.md)	 - Update a policy template
* [flexera-cli policy policy-template validate](flexera-cli_policy_policy-template_validate.md)	 - Validate a policy template

