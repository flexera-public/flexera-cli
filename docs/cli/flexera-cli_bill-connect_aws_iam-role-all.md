## flexera-cli bill-connect aws iam-role-all

Creates an AWS bill connect using the IAM Role-based authorization method

```
flexera-cli bill-connect aws iam-role-all [flags]
```

### Options

```
      --bill-account-id string         billAccountId (body)
      --body string                    raw JSON body (inline | @file | @-); overrides body field flags
      --bucket-name string             bucketName (body)
      --bucket-path string             bucketPath (body)
      --dry-run                        print the planned operation as JSON and exit without calling the API
      --effective-from string          effectiveFrom (body)
  -h, --help                           help for iam-role-all
      --sts-external-id string         stsExternalId (body)
      --sts-role-arn string            stsRoleArn (body)
      --sts-role-session-name string   stsRoleSessionName (body)
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
      --login-base-url string   override login base URL
      --org-id int              organization ID
  -o, --output string           output format (json|table)
      --refresh-token string    OAuth refresh token
      --zone string             API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli bill-connect aws](flexera-cli_bill-connect_aws.md)	 - Bill Connect - AWS operations (generated from the unified OpenAPI spec)

