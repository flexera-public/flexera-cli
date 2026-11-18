## flexera-cli fnms report execute-async create

Start an asynchronous FNMS report execution

### Synopsis

Start an asynchronous FNMS report execution

Async version of the execute endpoint. Starts an asynchronous execution of a saved or built-in ITAM report. The report is executed in the background, and results are partitioned into pages and stored for efficient retrieval.

**Benefits and Practical Applications:**
* **No Timeouts:** Large reports that would otherwise time out are processed in the background.<br><br>
* **Efficient Retrieval:** Results are pre-partitioned into pages, so each retrieval call is fast regardless of total result size.<br><br>
* **Deduplication:** If an identical execution is already running, the API returns the existing job details (409 Conflict) unless forceNewExecution=true is specified.<br><br>

**Important:** The limit parameter is fixed at job creation time and defaults to 10000 if not specified. Results are pre-partitioned into immutable pages in the background based on that limit value. It cannot be changed after the job has been initiated.

**Steps:**
* **Step 1:** POST to .../reports/{id}/execute-async with optional limit and forceNewExecution parameters. Returns a jobId and statusUrl.
* **Step 2:** GET {statusUrl} to poll until status is "Completed".
* **Step 3:** Use the retrieveUrl from the completed status response to fetch the first page of results. Use the page URL from each response to fetch subsequent pages.

```
flexera-cli fnms report execute-async create [flags]
```

### Options

```
      --dry-run[=true|false]               preview without making API requests
      --filter-params string               report filters as a JSON object (deepObject query parameters); Query parameters
      --force-new-execution[=true|false]   skip job deduplication; Skip deduplication check and force a new execution.; API default: false; illustrative example: false
  -h, --help[=true|false]                  help for create
      --id string                          saved report ID (required); The saved report ID to execute asynchronously.; required by API; illustrative example: "2361"
      --limit int                          rows per result page (5000–10000); Number of rows per result page. Defaults to 10000 if not specified. Must be between 5000 and 10000.; format: int64; minimum: 5000; maximum: 10000; API default: 10000; illustrative example: 10000 (default 10000)
      --search-text string                 text to filter report results; Text to filter report results.; illustrative example: "dev"
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

* [flexera-cli fnms report execute-async](flexera-cli_fnms_report_execute-async.md)	 - execute-async operations

