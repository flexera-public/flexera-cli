## flexera-cli iam access-rule list

List access rules

### Synopsis

List access rules

Retrieve an organization's access rules.

```
flexera-cli iam access-rule list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam access-rule list --org-id ORG_ID
```

### Options

```
      --filter string   filter (query); Optional filter for returning access rules matching specific criteria. ### Supported Filter Keys | Filter | Required | Valid Operators | Example | Description | | --- | --- | --- | --- | --- | | subjectRef | no | eq | subjectRef eq 'ref:nam:::iam:user:123' | Return only access rules for the user with ID 123 | | scopeRef | no | eq | scopeRef eq 'ref:nam:::iam... (see cli schema); minLength: 1; illustrative example: "subjectRef eq 'ref:nam:::iam:user:12345'"
  -h, --help            help for list
      --view string     view (query); View used to render details of the subject; enum: ["default","extended"]; API default: "default"; illustrative example: "extended"
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

* [flexera-cli iam access-rule](flexera-cli_iam_access-rule.md)	 - Access Rule operations (generated from the unified OpenAPI spec)

