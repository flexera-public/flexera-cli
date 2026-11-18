## flexera-cli fnms report execute-async status

Get async report execution status

### Synopsis

Get async report execution status

Returns the current status of an async report execution job. Poll this endpoint until the status is "Completed" or "Errored".

**Job statuses:**
* **Created**: The job has been accepted and is queued for processing.
* **In progress**: The job is actively running.
* **Completed**: The report executed successfully. Results are available via the retrieveUrl.
* **Errored**: The job failed. Check the errorMessage field for details.

When the job completes successfully, the response includes total, totalPages, and a retrieveUrl to start fetching results.

```
flexera-cli fnms report execute-async status [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli fnms report execute-async status --org-id ORG_ID --id ID --job-id JOB_ID
```

### Options

```
  -h, --help[=true|false]   help for status
      --id string           id (path, required); The report ID.; required by API; illustrative example: "2361"
      --job-id string       jobId (path, required); The job ID returned from the execute-async call.; required by API; illustrative example: "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
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

* [flexera-cli fnms report execute-async](flexera-cli_fnms_report_execute-async.md)	 - execute-async operations

