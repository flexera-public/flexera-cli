## flexera-cli saas import-job create

Create import job

### Synopsis

Create import job

This resource creates an import job for loading data into the managed application. SaaS data objects must be first created before an import job can be submitted.

```
flexera-cli saas import-job create [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas import-job create --org-id ORG_ID --body @request.json
  flexera-cli saas import-job create --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema saas import-job create --example > request.json
```

### Options

```
      --body string             raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run                 print the planned operation as JSON and exit without calling the API
  -h, --help                    help for create
  -i, --interactive             edit inputs in a terminal form, review a plan and approve with typed yes
      --managed-app-id string   managedAppId (body); required by API; managedAppId identifies a managed application by given ID.; pattern: "^[0-9a-f]+$"; illustrative example: "123"
      --yes                     confirm the operation (required for destructive ops)
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

* [flexera-cli saas import-job](flexera-cli_saas_import-job.md)	 - Import Job operations (generated from the unified OpenAPI spec)

