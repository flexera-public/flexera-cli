## flexera-cli saas object-definition list

List Object Definitions

### Synopsis

List Object Definitions

This endpoint retrieves a list of object definitions.

```
flexera-cli saas object-definition list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas object-definition list --org-id ORG_ID
```

### Options

```
      --filter string     filter (query); The filter to query the query definitions for specific vendor or app.. Supported fields in the filter are [vendorTechnopediaId, appTechnopediaId, appId]; illustrative example: "appTechnopediaId eq '89ccc2c8-f1dc-47e1-bfd6-dc1847f41869'"
  -h, --help              help for list
      --order-by string   orderBy (query); The order by filter to sort the managed applications. Supported fields in the orderBy are [modifiedAt]; API default: "modifiedAt desc"; illustrative example: "modifiedAt desc"
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

* [flexera-cli saas object-definition](flexera-cli_saas_object-definition.md)	 - Object Definition operations (generated from the unified OpenAPI spec)

