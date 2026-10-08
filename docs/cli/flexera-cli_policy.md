## flexera-cli policy

Policy API

### Synopsis

Commands for the Policy API (Flexera Policy API, v1).

Service id: policy

### Options

```
  -h, --help   help for policy
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

* [flexera-cli](flexera-cli.md)	 - Flexera One unified API command-line client
* [flexera-cli policy action-status](flexera-cli_policy_action-status.md)	 - Policy action statuses (project-scoped)
* [flexera-cli policy applied-policy](flexera-cli_policy_applied-policy.md)	 - Applied policies (project-scoped)
* [flexera-cli policy archived-incident](flexera-cli_policy_archived-incident.md)	 - Policy archived incidents (project-scoped)
* [flexera-cli policy custom-catalog](flexera-cli_policy_custom-catalog.md)	 - CustomCatalog operations (generated from the unified OpenAPI spec)
* [flexera-cli policy customization-type](flexera-cli_policy_customization-type.md)	 - Customization Type operations (generated from the unified OpenAPI spec)
* [flexera-cli policy customization-value](flexera-cli_policy_customization-value.md)	 - Customization Value operations (generated from the unified OpenAPI spec)
* [flexera-cli policy incident-aggregate](flexera-cli_policy_incident-aggregate.md)	 - Incident Aggregate operations (generated from the unified OpenAPI spec)
* [flexera-cli policy meta](flexera-cli_policy_meta.md)	 - Relationship-aware applied-policy meta operations (project-scoped)
* [flexera-cli policy policy-aggregate](flexera-cli_policy_policy-aggregate.md)	 - Policy Aggregate operations (generated from the unified OpenAPI spec)
* [flexera-cli policy policy-manager](flexera-cli_policy_policy-manager.md)	 - Policy Manager operations (generated from the unified OpenAPI spec)
* [flexera-cli policy policy-template](flexera-cli_policy_policy-template.md)	 - Policy templates (project-scoped)
* [flexera-cli policy published-template](flexera-cli_policy_published-template.md)	 - Published Template operations (generated from the unified OpenAPI spec)
* [flexera-cli policy unmanaged-applied-policies](flexera-cli_policy_unmanaged-applied-policies.md)	 - Unmanaged Applied Policies operations (generated from the unified OpenAPI spec)
* [flexera-cli policy unmanaged-incidents](flexera-cli_policy_unmanaged-incidents.md)	 - Unmanaged Incidents operations (generated from the unified OpenAPI spec)

