## flexera-cli awssts-credential update

Update a Credential

```
flexera-cli awssts-credential update [flags]
```

### Options

```
      --access-key string          accessKey (body)
      --body string                raw JSON body (inline | @file | @-); overrides body field flags
      --description string         description (body)
      --dry-run                    print the planned operation as JSON and exit without calling the API
      --external-id string         externalId (body)
  -h, --help                       help for update
      --id string                  id (path, required)
      --max-minutes int            maxMinutes (body)
      --name string                name (body)
      --policy-json string         policyJson (body)
      --region string              region (body)
      --role-arn string            roleArn (body)
      --role-session-name string   roleSessionName (body)
      --secret-key string          secretKey (body)
      --service string             service (body)
      --version int                version (body)
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

* [flexera-cli awssts-credential](flexera-cli_awssts-credential.md)	 - AWS STS Credential operations (generated from the unified OpenAPI spec)

