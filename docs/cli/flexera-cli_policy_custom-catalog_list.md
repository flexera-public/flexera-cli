## flexera-cli policy custom-catalog list

Index published templates with applied custom catalog settings

### Synopsis

Index published templates with applied custom catalog settings

Index returns a list of published templates which can be filtered by the catalogSource parameter.
The default setting only includes templates made available to all MSP child orgs.
Using different catalogSource value enables users to preview what templates will be available to their child orgs with other "policy-catalog-reference-org" settings.
This setting can be changed at the [Customization Value](https://developer.flexera.com/docs/api/policy/v1#/Customization%20Value) endpoint.

```
flexera-cli policy custom-catalog list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli policy custom-catalog list --org-id ORG_ID
```

### Options

```
      --catalog-source string   catalogSource (query); - "default": Includes templates based on the configuration set in the Policy Customization Service ("flexera-default", "msp-parent", or "msp-parent|flexera-default"). - "msp-parent": Includes only the custom templates created by the MSP parent organization. - "flexera-default": Includes only the templates published by Flexera. - "msp-parent|flexera-default":... (see cli schema); enum: ["default","msp-parent","flexera-default","msp-parent|flexera-default","debug","taggable-templates"]; API default: "default"; illustrative example: "msp-parent"
      --filter string           filter (query); Optional filter to retrieve published templates based on specific criteria. ### Supported Filter Keys | Filter | Type | Example | Description | | ------------------------- | ------ | -------------------------------------- | ------------------------------------------------------------------------------------- | | name | string | name in ['policy123', 'policy1... (see cli schema); illustrative example: "name eq 'foo'"
  -h, --help                    help for list
      --limit int               limit (query); Specifies a custom limit for pagination.; format: int64; illustrative example: 1000
      --no-paginate             return only the first page (do not follow nextPage)
      --order-by string         orderBy (query); Specifies the order to sort published templates by fields such as [name, shortDescription, longDescription, docLink, category, createdBy.email, createdAt, updatedAt, builtIn, hidden]. Ordering can be [asc] or [desc]. Default ordering is [asc] if no value is set.; illustrative example: "category asc, createdAt asc"
      --skip-token string       resume pagination from this token; Used in pagination to point to the next or previous set of records.
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

