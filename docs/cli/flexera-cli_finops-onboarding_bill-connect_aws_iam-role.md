## flexera-cli finops-onboarding bill-connect aws iam-role

Update an AWS bill connect that was created using AWS IAM Role-based create method

### Synopsis

Update an AWS bill connect that was created using AWS IAM Role-based create method

Updates an AWS bill connect that was created using the AWS IAM Role-based create method.

```
flexera-cli finops-onboarding bill-connect aws iam-role [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-onboarding bill-connect aws iam-role --org-id ORG_ID --id ID --body @request.json
  flexera-cli finops-onboarding bill-connect aws iam-role --org-id ORG_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-onboarding bill-connect aws iam-role --example > request.json
```

### Options

```
      --body string                    raw JSON body (inline | @file | @-); overrides body field flags
      --bucket-name string             bucketName (body); Name of the S3 bucket where bill files are saved; minLength: 1; illustrative example: "bills-bucket"
      --bucket-path string             bucketPath (body); Path to the bill files from the AWS S3 bucket root; minLength: 1; illustrative example: "billing/path/"
      --dry-run                        print the planned operation as JSON and exit without calling the API
      --effective-from string          effectiveFrom (body); The earliest billing month (UTC) from which to start processing data, formatted YYYY-MM. If omitted when creating an AWS bill connect, the current UTC month is used. If omitted when updating, the existing value is unchanged.; pattern: "^20[\\d]{2}-((0[1-9])|(1[012]))$"; illustrative example: "2025-10"
  -h, --help                           help for iam-role
      --id string                      id (path, required); Identifies a bill connect; required by API; illustrative example: "aws-20194320903"
  -i, --interactive                    edit inputs in a terminal form, review a plan and approve with typed yes
      --sts-external-id string         stsExternalId (body); Unique identifier used when assuming a role in the customers account. This defaults to the customers org_id; minLength: 2; maxLength: 1224; pattern: "^[^ ]+$"; illustrative example: "abcdef1234567890"
      --sts-role-arn string            stsRoleArn (body); Amazon Resource Name (ARN) of the assumed sts role; minLength: 1; illustrative example: "arn:aws:iam::123456789012:role/bill_access_role"
      --sts-role-session-name string   stsRoleSessionName (body); Identifier for the assumed sts role session, this will default to flexera-finops if not supplied; illustrative example: "flexera-finops"
      --yes                            confirm the operation (required for destructive ops)
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

