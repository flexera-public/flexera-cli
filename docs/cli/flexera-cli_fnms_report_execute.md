## flexera-cli fnms report execute

Start an asynchronous FNMS report execution and wait for completion

```
flexera-cli fnms report execute [flags]
```

### Options

```
      --dry-run[=true|false]               preview without making API requests
      --filter-params string               report filters as a JSON object (deepObject query parameters)
      --force-new-execution[=true|false]   skip job deduplication
  -h, --help[=true|false]                  help for execute
      --id string                          saved report ID (required)
      --limit int                          rows per result page (5000–10000) (default 10000)
      --poll-interval duration             interval between status requests (default 2s)
      --poll-timeout duration              maximum time to wait (default 5m0s)
      --search-text string                 text to filter report results
      --yes[=true|false]                   confirm operation (not required for non-destructive report jobs)
```

### Options inherited from parent commands

```
      --access-token string        static bearer access token
      --api-base-url string        override API base URL
      --client-id string           OAuth client ID
      --client-secret string       OAuth client secret
      --config string              config file (default $HOME/.flexera/config.yaml)
  -d, --debug[=true|false]         log HTTP requests/responses to stderr (Authorization redacted)
      --json-style string          JSON whitespace style (auto|pretty|compact) (default "auto")
      --login-base-url string      override login base URL
      --no-validate[=true|false]   skip API schema constraints (never JSON syntax or request data-loss checks)
      --org-id int                 organization ID
      --out-fields string          project JSON output fields (comma-separated paths)
      --out-jq string              shape JSON output with a jq expression
  -o, --output string              output format (json|table)
  -r, --raw-output[=true|false]    write jq string results without JSON quotes (requires --out-jq)
      --refresh-token string       OAuth refresh token
      --zone string                API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli fnms report](flexera-cli_fnms_report.md)	 - Report operations (generated from the unified OpenAPI spec)

