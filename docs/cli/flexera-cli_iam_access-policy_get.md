## flexera-cli iam access-policy get

Get an access policy by ID.

### Synopsis

Get an access policy by ID.

Get an access policy by ID

```
flexera-cli iam access-policy get [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam access-policy get --org-id ORG_ID --access-policy-id ACCESS_POLICY_ID
```

### Options

```
      --access-policy-id string   accessPolicyId (path, required); Access Policy ID; required by API; illustrative example: "Officiis totam unde saepe cum."
  -h, --help                      help for get
      --view string               view (query); View used to render access policy; enum: ["default","extended"]; illustrative example: "default"
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

