## flexera-cli policy custom-catalog replace

Create or update custom catalog tags for a published template under an MSP parent organization

### Synopsis

Create or update custom catalog tags for a published template under an MSP parent organization

Upsert enables an MSP parent organization to create or update custom catalog tags for a published template.
Custom catalog tags help define template visibility for child organizations, allowing MSP parents to control whether the templates are accessible to child organizations or restricted to the MSP parent itself.
Templates tagged as visible will only show up to child MSP organizations if the "policy-catalog-reference-org" settings for [Customization Value](https://developer.flexera.com/docs/api/policy/v1#/Customization%20Value) is set to "msp-parent" or "msp-parent|flexera-default"

```
flexera-cli policy custom-catalog replace [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli policy custom-catalog replace --org-id ORG_ID --published-template-id PUBLISHED_TEMPLATE_ID --body @request.json
  flexera-cli policy custom-catalog replace --org-id ORG_ID --published-template-id PUBLISHED_TEMPLATE_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema policy custom-catalog replace --example > request.json
```

### Options

```
      --body string                    raw JSON body (inline | @file | @-); overrides body field flags
      --child-orgs-visibility          childOrgsVisibility (body); required by API; Defines whether the published template is visible to child organizations under an MSP parent. If set to 'true', child organizations can access and use this template. If 'false', only the MSP parent organization can see it. By default, the template is not visible to child organizations, so any template that is not explicitly set to 'true' will remain hidden f... (see cli schema); illustrative example: true
      --dry-run                        print the planned operation as JSON and exit without calling the API
  -h, --help                           help for replace
  -i, --interactive                    edit inputs in a terminal form, review a plan and approve with typed yes
      --published-template-id string   publishedTemplateId (path, required); The unique identifier for the published template.; required by API; illustrative example: "5b06ead5e0dacc007058c784"
      --yes                            confirm the operation (required for destructive ops)
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

* [flexera-cli policy custom-catalog](flexera-cli_policy_custom-catalog.md)	 - CustomCatalog operations (generated from the unified OpenAPI spec)

