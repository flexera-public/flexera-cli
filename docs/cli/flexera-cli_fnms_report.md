## flexera-cli fnms report

Report operations (generated from the unified OpenAPI spec)

### Options

```
  -h, --help[=true|false]   help for report
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

* [flexera-cli fnms](flexera-cli_fnms.md)	 - IT Asset Management API
* [flexera-cli fnms report download-export](flexera-cli_fnms_report_download-export.md)	 - Stream FNMS download bytes unchanged to stdout or a new file
* [flexera-cli fnms report execute](flexera-cli_fnms_report_execute.md)	 - Start an asynchronous FNMS report execution and wait for completion
* [flexera-cli fnms report execute-async](flexera-cli_fnms_report_execute-async.md)	 - execute-async operations
* [flexera-cli fnms report export](flexera-cli_fnms_report_export.md)	 - Start an FNMS report export and wait for completion
* [flexera-cli fnms report exports](flexera-cli_fnms_report_exports.md)	 - exports operations
* [flexera-cli fnms report list](flexera-cli_fnms_report_list.md)	 - List ITAM saved and built-in reports
* [flexera-cli fnms report wait-execution](flexera-cli_fnms_report_wait-execution.md)	 - Wait for an existing report execution (does not retrieve result pages)
* [flexera-cli fnms report wait-export](flexera-cli_fnms_report_wait-export.md)	 - Wait for an existing report export and return its download path

