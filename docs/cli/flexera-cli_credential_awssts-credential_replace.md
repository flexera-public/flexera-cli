## flexera-cli credential awssts-credential replace

Create a Credential

### Synopsis

Create a Credential

Create a Credential that uses the 'AWS STS' scheme.

```
flexera-cli credential awssts-credential replace [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli credential awssts-credential replace --org-id ORG_ID --id ID --body @request.json
  flexera-cli credential awssts-credential replace --org-id ORG_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema credential awssts-credential replace --example > request.json
```

### Options

```
      --access-key string          accessKey (body); AWS Access Key ID; illustrative example: "KEYFROMAWSINCAPSANDDIGITS123"
      --body string                raw JSON body (inline | @file | @-); overrides body field flags
      --body-version int           version (body); AWS Signature version; format: int64; enum: [4]; API default: 4; illustrative example: 4
      --description string         description (body); Credentials description; illustrative example: "The AWS Oregon region (us-west-2) development credentials."
      --dry-run                    print the planned operation as JSON and exit without calling the API
      --external-id string         externalId (body); Unique identifier that might be required when assuming a role in another account.; illustrative example: "abcdefghij-123456790"
  -h, --help                       help for replace
      --id string                  id (path, required); Credentials id; required by API; pattern: "^[_a-zA-Z0-9][-_a-zA-Z0-9]{0,127}$"; illustrative example: "abcdefghij-123456790"
  -i, --interactive                edit inputs in a terminal form, review a plan and approve with typed yes
      --max-minutes int            maxMinutes (body); Maximum time, in minutes, that signer token should remain valid before needing refresh. Administrative setting may differ on the allowed time limit.; format: int64; minimum: 30; maximum: 720; API default: 60; illustrative example: 620
      --name string                name (body); required by API; Credentials name used in UI; illustrative example: "Development Credentials"
      --policy-json string         policyJson (body); Restrictive IAM policy in JSON format; illustrative example: "{\"Statement\": [{\"Action\": [\"s3:List*\", \"s3:Get*\"], \"Effect\": \"Allow\", \"Resource\": [\"arn:aws:s3:::some-bucket/*\", \"arn:aws:s3:::some-bucket\"], \"Sid\": \"\"}], \"... (see cli schema)
      --region string              region (body); AWS region hosting service endpoint; API default: "inferred"; illustrative example: "us-east-1"
      --role-arn string            roleArn (body); required by API; Amazon Resource Name (ARN) of the role to assume; minLength: 1; illustrative example: "arn:aws:iam::012345678912:role/some-role-name"
      --role-session-name string   roleSessionName (body); required by API; Identifier for the assumed role session; minLength: 1; illustrative example: "abcdefghij-123456790"
      --secret-key string          secretKey (body); AWS Secret Key
      --service string             service (body); AWS Service for which request is signed; API default: "inferred"; illustrative example: "ec2"
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

* [flexera-cli credential awssts-credential](flexera-cli_credential_awssts-credential.md)	 - AWS STS Credential operations (generated from the unified OpenAPI spec)

