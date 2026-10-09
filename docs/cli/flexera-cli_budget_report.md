## flexera-cli budget report

Provides a report comparing budget with actual spend

### Synopsis

Provides a report comparing budget with actual spend

Returns a report comparing the budget and forecast to the actual spend.

The response is similar to the rows returned by the `costs/aggregated` API.

Each row represents a specific combination of dimension values and timestamp, and includes the corresponding budget amount, forecast amount, actual spend amount and variance (between forecast and actual).

The API also expects the user to have cost access to the segments in the report. If the budget uses billing center level dimensions, access to each row/segment
is determined by the user's access to costs on the billing center. If the budget doesn't use billing center level dimensions, then the user is expected to have
cost access at the org level.

By default, the report contains one row for each combination of budget segment and year-month.
For example, for a budget with 10 budget segments covering 12 months,
the report would return 120 rows (i.e., 10 segments * 12 months).

### Budgeted and Unbudgeted spend
All rows include a `budgeted` dimension with a boolean value.
It is `false` only if the row contains new dimension values from the cost data that were not described by the budget segments.

For example, if a budget defines segments for 2 specific account values, for those rows the `budgeted` dimension will be `true`.

If a 3rd account appears in the billing data, its row will be returned with the `budgeted` dimension set to `false`.

This allows you to identify the unbudgeted spend from the 3rd account. You may then wish to either:
- update the budget segments to cover the new account,
- or modify the budget's filter expression to exclude the account.

### Filtering
To narrow the report to a specific range of year-months, specify the start and end dates in the filter query parameter.

By default, the budget report returns all the budgeted and unbudgeted rows.
- Set the `budgeted` parameter to `true` to return only budgeted rows.
- Set the `budgeted` parameter to `false` to return only unbudgeted rows.

### Aggregation across budget segments
The report API can sum up results based on the requested `dimensions`.

For example, consider a budget broken down by two dimensions: Vendor and Account.

- If the `dimensions` list includes both Vendor and Account, the report will display the full budget details, with rows for potentially dozens of accounts.
- If the `dimensions` list includes only Vendor, the report will return fewer rows, providing a summary of the budget, forecast, actual spend and variance for each vendor.
- If the `dimensions` list is empty, the report will return just one row for each year-month, summarizing the total budget, forecast, actual spend and variance across all budget segments.

### Aggregation across year-months

The report API can also sum up results across year-months.

By default, it returns one row for each budget segment and year-month combination.

However, by setting the `summarized` parameter to `true`,
the report will return only one row per budget segment,
representing its total budget, forecast, spend and variance across all the selected year-months.

```
flexera-cli budget report [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli budget report --org-id ORG_ID --id ID --start-at START_AT --end-at END_AT
```

### Options

```
      --dimensions strings   dimensions (query); The list of supported dimensions by which to report budget spend.; CLI: comma-separated values or repeated flag; illustrative example: ["ProviderName","SubAccountName"]
      --end-at string        endAt (query); Latest timestamp (exclusive) of the budget and spend. Consists of a year and month in YYYY-MM format. Will be interpreted as UTC, which is used for period boundaries. No records will be returned on or after this timestamp.; required by API; illustrative example: "2023-03"
      --filter true          filter (query); The filter query string parameter allows to filter reports. ***Note:*** The value for the filter query parameter must be a URL-encoded query string. | Filter | Required | Description | Allowed Operator | Behavior | Example | | --- | --- | --- | --- | --- | --- | | budgeted | No | Filter to limit returned data to just budgeted or unbudgeted rows. If true, o... (see cli schema); illustrative example: "budgeted eq true"
  -h, --help                 help for report
      --id string            id (path, required); Identifier of the budget; required by API; illustrative example: "2ed7db3"
      --start-at string      startAt (query); Earliest timestamp (inclusive) of the returned budget and spend. Consists of a year and month in YYYY-MM format. Will be interpreted as UTC, which is used for period boundaries.; required by API; illustrative example: "2023-01"
      --summarized false     summarized (query); Aggregates the rows for multiple year-months into a single row. Default false.; API default: false; illustrative example: true
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

* [flexera-cli budget](flexera-cli_budget.md)	 - Budget API

