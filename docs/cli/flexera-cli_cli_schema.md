## flexera-cli cli schema

Inspect a command's API schemas, parameters and illustrative example

### Synopsis

Inspect an exact spec-backed operation without authentication or network access. Examples are illustrative, not a promise of server acceptance. --example emits only a schema-validated, non-sensitive request body; unavailable or invalid examples are errors.

```
flexera-cli cli schema <command path...> [flags]
```

### Options

```
      --depth int     reference expansion depth; 0 preserves references (default 3)
      --example       emit only a validated illustrative request body
  -h, --help          help for schema
      --part string   schema part (request|response|params)
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

* [flexera-cli cli](flexera-cli_cli.md)	 - Discover CLI commands and API schemas

