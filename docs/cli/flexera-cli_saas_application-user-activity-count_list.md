## flexera-cli saas application-user-activity-count list

List activity counts of the user

### Synopsis

List activity counts of the user

Retrieves the activity counts of the user for all the ranges. Possible ranges are:
                fifteen(0-15), thirty(16-30), fortyfive(31-45), sixty(46-60), ninety(61-90), onetwenty(91-120), onetwentyplus(121+)

                The activity count is calculated by number of days past the current date.

```
flexera-cli saas application-user-activity-count list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas application-user-activity-count list --org-id ORG_ID --body @request.json
Validated illustrative body, when available (review before use):
  flexera-cli cli schema saas application-user-activity-count list --example > request.json
```

### Options

```
      --body string             raw JSON body (inline | @file | @-); overrides body field flags
  -h, --help                    help for list
      --managed-app-id string   managedAppId (body); required by API; managed application id for which activity counts should be returned.; pattern: "^[0-9a-f]+$"; illustrative example: "345"
      --unique-id string        uniqueId (body); required by API; User's unique id for which activity counts should be returned.; illustrative example: "abc-xyz"
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

* [flexera-cli saas application-user-activity-count](flexera-cli_saas_application-user-activity-count.md)	 - Application User Activity Count operations (generated from the unified OpenAPI spec)

