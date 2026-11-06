## flexera-cli iam msp-customer-tag tags delete

Delete a tag for an MSP's customer

### Synopsis

Delete a tag for an MSP's customer

Delete removes a tag from the managed service provider's customer tenant.

```
flexera-cli iam msp-customer-tag tags delete [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam msp-customer-tag tags delete --org-id ORG_ID --customer-id CUSTOMER_ID --tag TAG
  flexera-cli iam msp-customer-tag tags delete --org-id ORG_ID --customer-id CUSTOMER_ID --tag TAG --dry-run
```

### Options

```
      --customer-id int        customerId (path, required); ID of the managed service provider's customer tenant; required by API; minimum: 1; illustrative example: 1234
      --dry-run[=true|false]   print the planned operation as JSON and exit without calling the API
  -h, --help[=true|false]      help for delete
      --tag string             tag (path, required); Name of the tag; required by API; minLength: 1; maxLength: 256; pattern: "[a-zA-Z0-1\\:]"; illustrative example: "alpha"
      --yes[=true|false]       confirm the operation (required for destructive ops)
```

### Options inherited from parent commands

```
      --access-token string        static bearer access token
      --api-base-url string        override API base URL
      --client-id string           OAuth client ID
      --client-secret string       OAuth client secret
      --config string              config file (default $HOME/.flexera/config.yaml)
  -d, --debug[=true|false]         log HTTP requests/responses to stderr (Authorization redacted)
      --json-style string          JSON whitespace style (auto|pretty|compact) (default "auto")
      --login-base-url string      override login base URL
      --no-validate[=true|false]   skip API schema constraints (never JSON syntax or request data-loss checks)
      --org-id int                 organization ID
      --out-fields string          project JSON output fields (comma-separated paths)
      --out-jq string              shape JSON output with a jq expression
  -o, --output string              output format (json|table)
  -r, --raw-output[=true|false]    write jq string results without JSON quotes (requires --out-jq)
      --refresh-token string       OAuth refresh token
      --zone string                API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli iam msp-customer-tag tags](flexera-cli_iam_msp-customer-tag_tags.md)	 - tags operations

