## flexera-cli finops-onboarding processing-history bill-connects bill-months get

Get enterprise bill processing details

### Synopsis

Get enterprise bill processing details

Returns the ordered processing stages, customer-safe errors, and warnings for the selected enterprise bill month execution.

Only stages applicable to that execution are returned. Loading Bill Month Data is included when the execution loads or normalizes bill data and is omitted for reprocessing that starts at currency conversion. Currency conversion is omitted when confirmed unnecessary.

An inProgress stage identifies the next expected work and does not necessarily mean a worker is running. Warnings do not prevent a stage from completing. Finishing Up completes only when the processing request or importer execution has completed.

This operation is read-only and is available only to enterprise organizations without the Billing capability.

```
flexera-cli finops-onboarding processing-history bill-connects bill-months get [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-onboarding processing-history bill-connects bill-months get --org-id ORG_ID --bill-connect-id BILL_CONNECT_ID --bill-month BILL_MONTH
```

### Options

```
      --bill-connect-id string   billConnectId (path, required); Bill connect identifier; required by API; minLength: 1; illustrative example: "aws-91123443"
      --bill-month string        billMonth (path, required); Bill month in YYYYMM format; required by API; pattern: "^20[\\d]{2}((0[1-9])|(1[012]))$"; illustrative example: "202506"
  -h, --help[=true|false]        help for get
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

* [flexera-cli finops-onboarding processing-history bill-connects bill-months](flexera-cli_finops-onboarding_processing-history_bill-connects_bill-months.md)	 - bill-months operations

