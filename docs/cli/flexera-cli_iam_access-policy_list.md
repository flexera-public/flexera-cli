## flexera-cli iam access-policy list

List all access policies for an organization

### Synopsis

List all access policies for an organization

List all access policies for an organization.

```
flexera-cli iam access-policy list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam access-policy list --org-id ORG_ID
```

### Options

```
      --filter string   filter (query); Filter used to narrow down access policies; illustrative example: "Qui voluptatibus dolores a qui eum."
  -h, --help            help for list
      --view string     view (query); View used to render access policy; enum: ["default","extended"]; illustrative example: "default"
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

* [flexera-cli iam access-policy](flexera-cli_iam_access-policy.md)	 - Access Policy operations (generated from the unified OpenAPI spec)

