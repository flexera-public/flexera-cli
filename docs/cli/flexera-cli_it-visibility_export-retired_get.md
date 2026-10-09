## flexera-cli it-visibility export-retired get

Download specified export file

### Synopsis

Download specified export file

Download returns the specified export file.

```
flexera-cli it-visibility export-retired get [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli it-visibility export-retired get --org-id ORG_ID --id ID
```

### Options

```
  -h, --help        help for get
      --id string   id (path, required); The Identifier of the file as returned from the INDEX endpoint; required by API; maxLength: 36; illustrative example: "7d795afa-c508-4d47-92ae-248113a792d1"
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

* [flexera-cli it-visibility export-retired](flexera-cli_it-visibility_export-retired.md)	 - Export (Retired) operations (generated from the unified OpenAPI spec)

