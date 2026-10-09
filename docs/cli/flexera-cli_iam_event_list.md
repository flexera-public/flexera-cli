## flexera-cli iam event list

Index an org's events

### Synopsis

Index an org's events

Index returns a collection of events which occurred within an organization, ordered from oldest to
newest.

### Pagination

If the "filter" expression includes an end timestamp (e.g. "timestamp le {date}"), this endpoint will return a
paginated list of events which occurred before the end timestamp. After all events matching the filter have been
returned, the response will not include a "nextPage" attribute. However, if the request does not include a filter
expression or the filter expression does not set an end timestamp, then an empty page indicates that all available
events have been read. In this case, responses will continue to include a "nextPage" attribute to allow the reader
to keep their place while polling for new events.

Standard Flexera [rate limits](https://developer.flexera.com/docs/page/rate-limiting) apply.

### Retention

By default, events are retained for 3 months. Events older than 3 months will not be returned by this API. To request
a longer retention policy, please contact Flexera Customer Support.

```
flexera-cli iam event list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam event list --org-id ORG_ID
```

### Options

```
      --filter string       filter (query); An optional expression for filtering the collection of events returned. The following filters are supported: | Filter | Description | Allowed Operator | Behavior | Example | | --- | --- | --- | --- | --- | | timestamp | Filters on the event's timestamp | le | Less than or equal to - returns only events with a timestamp less than or equal to the provided valu... (see cli schema); illustrative example: "timestamp ge '2019-10-21T10:20:50Z' and timestamp le '2019-10-22T10:20:50Z'"
  -h, --help                help for list
      --limit int           limit (query); The maximum number of events to return per page; maximum: 200; API default: 50; illustrative example: 100
      --no-paginate         return only the first page (do not follow nextPage)
      --skip-token string   resume pagination from this token; An opaque token to be provided when requesting a subsequent page after receiving a partial response. Partial responses will include a "nextPage" attribute in their response body, which contains the URL of the next page including the appropriate skipToken.
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

* [flexera-cli iam event](flexera-cli_iam_event.md)	 - Event operations (generated from the unified OpenAPI spec)

