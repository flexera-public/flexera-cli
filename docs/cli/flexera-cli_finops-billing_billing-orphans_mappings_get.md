## flexera-cli finops-billing billing-orphans mappings get

Show an Azure cost assignment

### Synopsis

Show an Azure cost assignment

Return an individual assignment. Recurring resources omit billMonth even when opened from a month-filtered list.

```
flexera-cli finops-billing billing-orphans mappings get [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing billing-orphans mappings get --org-id ORG_ID --id ID
```

### Options

```
  -h, --help[=true|false]   help for get
      --id string           id (path, required); Identifier of the saved assignment, not a provider charge identifier.; required by API; format: uuid; illustrative example: "b2c3d4e5-f6a7-4901-acde-f12345678901"
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

* [flexera-cli finops-billing billing-orphans mappings](flexera-cli_finops-billing_billing-orphans_mappings.md)	 - mappings operations

