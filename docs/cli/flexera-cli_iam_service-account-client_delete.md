## flexera-cli iam service-account-client delete

Delete a service account client

### Synopsis

Delete a service account client

Delete removes a service account client, invalidating the client's ID and secret. Other
clients belonging to the service account are unaffected.

```
flexera-cli iam service-account-client delete [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam service-account-client delete --org-id ORG_ID --service-account-id SERVICE_ACCOUNT_ID --target-client-id TARGET_CLIENT_ID
  flexera-cli iam service-account-client delete --org-id ORG_ID --service-account-id SERVICE_ACCOUNT_ID --target-client-id TARGET_CLIENT_ID --dry-run
```

### Options

```
      --dry-run                   print the planned operation as JSON and exit without calling the API
  -h, --help                      help for delete
      --service-account-id int    serviceAccountId (path, required); Unique identifier for the service account; required by API; minimum: 1; illustrative example: 1234
      --target-client-id string   clientId (path, required); Identifier for the client; required by API; minLength: 1; maxLength: 512; illustrative example: "1111aaaa2222bbbb3333cccc"
      --yes                       confirm the operation (required for destructive ops)
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

* [flexera-cli iam service-account-client](flexera-cli_iam_service-account-client.md)	 - Service Account Client operations (generated from the unified OpenAPI spec)

