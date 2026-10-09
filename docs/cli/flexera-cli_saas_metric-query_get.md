## flexera-cli saas metric-query get

Query metrics service

### Synopsis

Query metrics service

The query endpoint takes in a metric query request, pulls the relevant raw data from the
            metrics service and, marshals the result into a query result.

## Stored Queries

Here is a description of each default stored query available.
| queryName                        | Description                                                                                                                            |
|---------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------|
| departmentSpend                  | The amount spent per department within the scope of the managed application. This query requires organizational data to be available.  |
| applicationUsage                 | The number of active/inactive/never active app users within the activity threshold.                                                    |
| applicationUsageOverTime         | The number of active/inactive/never active app users over a given time period.                                                         |
| daysSinceActiveRange             | The number of active/inactive/never application users within the given time range.                                                     |
| licenseUsage                     | The number of active/inactive/never application users per managed application per license.                                             |
| licenseUsageOverTime             | The number of active/inactive/never application users per managed application per license over a given time period .                   |
| totalNumberOfUsers             | The number of total app users within the activity threshold.                                                                           |
| totalNumberOfUsersOverTime        | The number of total app users over a given time period.                                                                                |
| orgInfo                          | The estimated amount spent, number of active managed apps, licenses and total inactive, never active users for entire org are returned.|
| orgGroupSpend                    | The amount spent per department for each department in the org, sorted by amount desc                                                  |
| orgApplicationUsage              | The number of active/inactive/never active app users within the activity threshold for all apps in the org, sorted by total users desc.|

## Dimensions
Dimensions that the query can be split upon.
| Dimensions     | Description                                                                                                                                                             |
|----------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| resourceType   | If there is more than one resourceType can be a value such as appUser. If there are more resourceTypes available, they show up in the resourceType dimension construct. |
| group          | The groups as defined by the organizational data. For example, if department is available, the data will be split by department.                                        |
| resourceStatus | If the resource, such as AppUser has a status of active, never active, or inactive, the result will be split along those lines.                                            |

## Granularity
| Granularity  | Description                                                                  |
|--------------|------------------------------------------------------------------------------|
| day          | Each data point represents 24 hours in UTC.                                  |
| month        | Each data point represents an entire calendar month of data.                 |
| firstOfMonth | Each data point represents the first of the month of each month of the year. |

## Metrics
| Granularity  | Description                                                                  |
|--------------|------------------------------------------------------------------------------|
| day          | Each data point represents 24 hours in UTC.                                  |
| month        | Each data point represents an entire calendar month of data.                 |
| firstOfMonth | Each data point represents the first of the month of each month of the year. |

## Filter
| Filter       | Description                                                       |
|--------------|-------------------------------------------------------------------|
| managedAppId | The managed application id                                        |
| date         | A date in the format YYYY-MM-DD                                   |
| sku          | This value allows the filtering down to a particular license sku. |

## Currency

The currency (default is USD) can be set globally in Administration -> SaaS Settings -> General within app.flexera.com.

```
flexera-cli saas metric-query get [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas metric-query get --org-id ORG_ID --query-name QUERY_NAME --body @request.json
Validated illustrative body, when available (review before use):
  flexera-cli cli schema saas metric-query get --example > request.json
```

### Options

```
      --body string          raw JSON body (inline | @file | @-); overrides body field flags
      --dimensions strings   dimensions (body); An array of dimensions desired for the result.; minItems: 1; CLI: comma-separated values or repeated flag; items.enum: ["resourceType","group","resourceStatus"]; illustrative example: ["group"]
      --filter string        filter (body); A filter query string. The filter string is required. The available fields are [managedAppId, date, sku] | Filter | Description | Allowed operators | Example | |--------------|--------------------------------------------|-------------------|-----------------------------------------------------------------------------| | managedAppId | Filter to a single mana... (see cli schema); illustrative example: "date ge '2020-02-01T00:00:00Z' AND date lt '2020-10-01T00:00:00Z' AND managedAppId eq '1234'"
      --granularity string   granularity (body); The granularity desired for the metric result.; enum: ["day","month","firstOfMonth"]; API default: "month"; illustrative example: "month"
  -h, --help                 help for get
      --limit int            limit (body); A limit field limits the number of top records returned to the specified number. This field is valid only for orgGroupSpend and orgApplicationUsagequeries. Max limit is 100.; maximum: 100; illustrative example: 8
      --metrics strings      metrics (body); The metrics to calculate. costTotalAmortized will amortize the cost across all items. numOfTotal will returnthe total number of items.; CLI: comma-separated values or repeated flag; items.enum: ["numOfTotal","costTotalAmortized"]; illustrative example: ["numOfTotal","numOfTotal","numOfTotal"]
      --query-name string    queryName (path, required); queryName identifies a query to execute and return results.; required by API; enum: ["departmentSpend","applicationUsage","applicationUsageOverTime","daysSinceActiveRange","licenseUsage","licenseUsageOverTime","totalNumberOfUsers","totalNumberOfUsersOverTime","org... (see cli schema); illustrative example: "departmentSpend"
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

* [flexera-cli saas metric-query](flexera-cli_saas_metric-query.md)	 - Metric Query operations (generated from the unified OpenAPI spec)

