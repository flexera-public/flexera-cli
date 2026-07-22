## flexera-cli bill-connect-aws iam-user-all

Create an AWS bill connect using legacy IAM User-based authorization method

```
flexera-cli bill-connect-aws iam-user-all [flags]
```

### Options

```
      --access-key string          accessKey (body)
      --bill-account-id string     billAccountId (body)
      --body string                raw JSON body (inline | @file | @-); overrides body field flags
      --bucket-name string         bucketName (body)
      --bucket-path string         bucketPath (body)
      --dry-run                    print the planned operation as JSON and exit without calling the API
      --effective-from string      effectiveFrom (body)
  -h, --help                       help for iam-user-all
      --partition string           partition (body)
      --secret-access-key string   secretAccessKey (body)
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
      --login-base-url string   override login base URL
      --org-id int              organization ID
  -o, --output string           output format (json|table)
      --refresh-token string    OAuth refresh token
      --zone string             API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli bill-connect-aws](flexera-cli_bill-connect-aws.md)	 - Bill Connect - AWS operations (generated from the unified OpenAPI spec)

