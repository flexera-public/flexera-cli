## flexera-cli finops-onboarding bill-connect aws iam-user

Update an AWS bill connect that was created using legacy AWS IAM User-based create method

### Synopsis

Update an AWS bill connect that was created using legacy AWS IAM User-based create method

Modifies an existing AWS bill connect that was created using legacy AWS IAM User-based create method.

```
flexera-cli finops-onboarding bill-connect aws iam-user [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-onboarding bill-connect aws iam-user --org-id ORG_ID --id ID --body @request.json
  flexera-cli finops-onboarding bill-connect aws iam-user --org-id ORG_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-onboarding bill-connect aws iam-user --example > request.json
```

### Options

```
      --access-key string          accessKey (body); The AWS IAM user's access key, to access the billAccountId; minLength: 1; illustrative example: "AKIAIOSFODNN7EXAMPLE"
      --body string                raw JSON body (inline | @file | @-); overrides body field flags
      --bucket-name string         bucketName (body); Name of the S3 bucket where bill files are saved; minLength: 1; illustrative example: "bills-bucket"
      --bucket-path string         bucketPath (body); Path to the bill files from the AWS S3 bucket root; minLength: 1; illustrative example: "billing/path/"
      --dry-run                    print the planned operation as JSON and exit without calling the API
      --effective-from string      effectiveFrom (body); The earliest billing month (UTC) from which to start processing data, formatted YYYY-MM. If omitted when creating an AWS bill connect, the current UTC month is used. If omitted when updating, the existing value is unchanged.; pattern: "^20[\\d]{2}-((0[1-9])|(1[012]))$"; illustrative example: "2025-10"
  -h, --help                       help for iam-user
      --id string                  id (path, required); Identifies a bill connect; required by API; illustrative example: "aws-20194320903"
  -i, --interactive                edit inputs in a terminal form, review a plan and approve with typed yes
      --secret-access-key string   secretAccessKey (body); The AWS IAM user's secret access key, to access the billAccountId
      --yes                        confirm the operation (required for destructive ops)
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

* [flexera-cli finops-onboarding bill-connect aws](flexera-cli_finops-onboarding_bill-connect_aws.md)	 - Bill Connect - AWS operations (generated from the unified OpenAPI spec)

