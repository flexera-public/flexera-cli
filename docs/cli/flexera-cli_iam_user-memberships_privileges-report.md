## flexera-cli iam user-memberships privileges-report

Show a user's privileges report

### Synopsis

Show a user's privileges report

Returns a detailed report of the user's privileges across requested scopes (organizations and projects). This allows a user to discover what privileges they have been granted through various roles across different scopes.

```
flexera-cli iam user-memberships privileges-report [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam user-memberships privileges-report --id ID --scope-refs SCOPE_REFS
```

### Options

```
  -h, --help                 help for privileges-report
      --id int               id (path, required); ID of the user; required by API; minimum: 1; illustrative example: 12345
      --prefix string        prefix (query); Optional prefix to filter privilege IDs. Examples: 'cm:', 'ss:', 'platform:roles:', 'governance:'; minLength: 1; illustrative example: "cm:"
      --scope-refs strings   scopeRefs (query); Array of scope refs (org or project) to filter privileges. At least one scope ref must be provided.; required by API; minItems: 1; CLI: comma-separated values or repeated flag; illustrative example: ["ref:nam:::iam:org:1234","ref:nam:1234::iam:project:2345"]
      --view string          view (query); View used to render privilege report details; enum: ["default","extended"]; API default: "default"; illustrative example: "extended"
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

* [flexera-cli iam user-memberships](flexera-cli_iam_user-memberships.md)	 - User Memberships operations (generated from the unified OpenAPI spec)

