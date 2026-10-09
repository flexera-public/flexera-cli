## flexera-cli iam msp-customer create

Create a new MSP customer

### Synopsis

Create a new MSP customer

Create provisions a new customer tenant for a managed service provider.

```
flexera-cli iam msp-customer create [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam msp-customer create --org-id ORG_ID --body @request.json
  flexera-cli iam msp-customer create --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema iam msp-customer create --example > request.json
```

### Options

```
      --body string          raw JSON body (inline | @file | @-); overrides body field flags
      --description string   description (body); Optional text describing the customer; maxLength: 4096; illustrative example: "MSP Customer requires services X, Y, Z."
      --dry-run              print the planned operation as JSON and exit without calling the API
      --external-id string   externalId (body); Identifier of the organization used in your external system; maxLength: 256; illustrative example: "W12345"
  -h, --help                 help for create
  -i, --interactive          edit inputs in a terminal form, review a plan and approve with typed yes
      --name string          name (body); required by API; Friendly name for the customer; minLength: 1; maxLength: 512; illustrative example: "MSPCustomer Inc."
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

* [flexera-cli iam msp-customer](flexera-cli_iam_msp-customer.md)	 - MSP Customer operations (generated from the unified OpenAPI spec)

