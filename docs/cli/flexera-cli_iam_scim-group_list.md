## flexera-cli iam scim-group list

Index an org's groups.

### Synopsis

Index an org's groups.

Index returns a list of groups for an org.

This API returns the members attribute as an empty array in index responses. To retrieve members, call GET /scim/v2/orgs/{orgId}/Groups/{id}.

```
flexera-cli iam scim-group list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam scim-group list --org-id ORG_ID
```

### Options

```
      --filter string   filter (query); A filter to narrow the number of groups to return. Supported fields in the filter are [id, displayName] The following filters are supported: | Filter | Description | Allowed Operator | Behavior | Example | | --- | ---| --- | --- | --- | | id | Filters on the group's ID | eq | Equal - The attribute and operator values must be identical for a match. | id eq '1... (see cli schema); illustrative example: "id eq '586859'"
  -h, --help            help for list
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

* [flexera-cli iam scim-group](flexera-cli_iam_scim-group.md)	 - SCIM Group operations (generated from the unified OpenAPI spec)

