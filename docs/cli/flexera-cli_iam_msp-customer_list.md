## flexera-cli iam msp-customer list

Index an MSP's customers

### Synopsis

Index an MSP's customers

Index a managed service provider's list of customer tenants.

```
flexera-cli iam msp-customer list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam msp-customer list --org-id ORG_ID
```

### Options

```
      --filter string   filter (query); Optional filter for filtering list of customers returned. The following filters are supported: | Filter | Description | Allowed Operator | Behavior | Example | | --- | ---| --- | --- | --- | | name | Filters on the organization name | co | Contains - The entire operator value must be a substring of the attribute value for a match. | name co 'SAP' | | | | eq ... (see cli schema); minLength: 1; illustrative example: "(name co 'SAP' or name co 'HP')"
  -h, --help            help for list
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

