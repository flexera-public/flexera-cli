## flexera-cli iam customization get

Show an org's customization

### Synopsis

Show an org's customization

Returns the effective customization for an org. If the org has its own value, that is returned (source: "child"). If the org has no value but a parent org (partner or MSP) has one with shouldInherit=true (or unset), the parent's value is returned (source: "parent"). Returns 404 if no effective value exists for the org (no child override and no inheritable parent value).

```
flexera-cli iam customization get [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam customization get --org-id ORG_ID --id ID
```

### Options

```
  -h, --help          help for get
      --id string     id (path, required); The customization's unique identifier, which matches the Id of the customization type implemented by this customization.; required by API; pattern: "^[a-z-]+$"; illustrative example: "navbar-logo-url"
      --view string   view (query); View used to render the customization; enum: ["default"]; API default: "default"; illustrative example: "default"
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

* [flexera-cli iam customization](flexera-cli_iam_customization.md)	 - Customization operations (generated from the unified OpenAPI spec)

