## flexera-cli fnms report execute-async

execute-async operations

### Options

```
  -h, --help[=true|false]   help for execute-async
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
* [flexera-cli fnms report execute-async create](flexera-cli_fnms_report_execute-async_create.md)	 - Start an asynchronous FNMS report execution
* [flexera-cli fnms report execute-async retrieve](flexera-cli_fnms_report_execute-async_retrieve.md)	 - Retrieve async report execution results
* [flexera-cli fnms report execute-async status](flexera-cli_fnms_report_execute-async_status.md)	 - Get async report execution status

