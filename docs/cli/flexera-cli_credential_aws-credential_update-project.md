## flexera-cli credential aws-credential update-project

Update a Credential

### Synopsis

Update a Credential

Update a Credential that uses the 'AWS' scheme.

```
flexera-cli credential aws-credential update-project [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli credential aws-credential update-project --project-id PROJECT_ID --id ID --body @request.json
  flexera-cli credential aws-credential update-project --project-id PROJECT_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema credential aws-credential update-project --example > request.json
```

### Options

```
      --access-key string    accessKey (body); AWS Access Key ID; illustrative example: "KEYFROMAWSINCAPSANDDIGITS123"
      --body string          raw JSON body (inline | @file | @-); overrides body field flags
      --body-version int     version (body); AWS Signature version; format: int64; enum: [4]; illustrative example: 4
      --description string   description (body); Credentials description; illustrative example: "The AWS Oregon region (us-west-2) development credentials."
      --dry-run              print the planned operation as JSON and exit without calling the API
  -h, --help                 help for update-project
      --id string            id (path, required); Credentials id; required by API; pattern: "^[_a-zA-Z0-9][-_a-zA-Z0-9]{0,127}$"; illustrative example: "abcdefghij-123456790"
  -i, --interactive          edit inputs in a terminal form, review a plan and approve with typed yes
      --name string          name (body); Credentials name used in UI; illustrative example: "Development Credentials"
      --project-id int       projectId (path, required); Identifies the Project that owns the Credential.; required by API; format: int64; minimum: 1; illustrative example: 2345
      --region string        region (body); AWS region hosting service endpoint; illustrative example: "us-east-1"
      --secret-key string    secretKey (body); AWS Secret Key
      --service string       service (body); AWS Service for which request is signed; illustrative example: "ec2"
      --yes                  confirm the operation (required for destructive ops)
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

* [flexera-cli credential aws-credential](flexera-cli_credential_aws-credential.md)	 - AWS Credential operations (generated from the unified OpenAPI spec)

