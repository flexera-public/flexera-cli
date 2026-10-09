## flexera-cli iam customization-type list

Returns types of customizations that are available to an org

### Synopsis

Returns types of customizations that are available to an org

Get available customization types.

```
flexera-cli iam customization-type list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam customization-type list --org-id ORG_ID
```

### Options

```
      --filter string       filter (query); Optional filter for selecting the list of customization types returned. The following filters are supported: | Filter | Description | Allowed Operator | Behavior | Example | | --- | ---| --- | --- | --- | | id | Filters on the customization type's id | co | Contains - The entire operator value must be a substring of the attribute value for a match. | id co '... (see cli schema); illustrative example: "id co 'navbar'"
  -h, --help                help for list
      --no-paginate         return only the first page (do not follow nextPage)
      --skip-token string   resume pagination from this token; An opaque token to be provided when requesting a subsequent page after receiving a partial response. Partial responses will include a "nextPage" attribute in their response body, which contains the URL of the next page including the appropriate skipToken.
      --view string         view (query); View used to render the customization type; enum: ["default"]; API default: "default"; illustrative example: "default"
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

* [flexera-cli iam customization-type](flexera-cli_iam_customization-type.md)	 - Customization Type operations (generated from the unified OpenAPI spec)

