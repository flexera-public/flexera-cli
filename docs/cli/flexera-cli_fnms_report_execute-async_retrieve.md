## flexera-cli fnms report execute-async retrieve

Retrieve async report execution results

### Synopsis

Retrieve async report execution results

Retrieves paginated results from a completed async report execution.

Each page is a pre-computed chunk stored during job execution. Retrieval does not re-execute the query. It simply fetches the corresponding pre-computed page. The page size is fixed at the limit value specified when the job was initiated and cannot be changed.

On the first call, you can omit the page parameter or set page=1 to receive the first page. For JSON responses, the body includes pagination metadata. For CSV responses, use the Link header to navigate pages; when the Link header is absent, all pages have been retrieved.
This endpoint streams the response body; clients should read the response body until the stream completes.

**Calling retrieve before the job is complete:** If this endpoint is called while the job status is still "Created" or "In progress", it returns 409 Conflict. Poll the status endpoint (/{id}/execute-async/{jobId}/status) until the status is "Completed" before calling this endpoint.

```
flexera-cli fnms report execute-async retrieve [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli fnms report execute-async retrieve --org-id ORG_ID --id ID --job-id JOB_ID
```

### Options

```
  -h, --help[=true|false]   help for retrieve
      --id string           id (path, required); The report ID.; required by API; illustrative example: "2361"
      --job-id string       jobId (path, required); The job ID returned from the execute-async call.; required by API; illustrative example: "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
      --page int            page (query); Page number to retrieve (1-based). Omit or set to 1 for the first page.; format: int64; API default: 1; illustrative example: 2654421279973593000
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

