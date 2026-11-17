## flexera-cli finops-billing billing-orphans mappings replace

Replace an Azure cost assignment destination

### Synopsis

Replace an Azure cost assignment destination

Replace the only mutable field, targetSubAccountId, using the ETag from show.
Organization, source, billing account, frequency, optional month, and description
are immutable. Destination parentage is the caller's responsibility. A changed destination requires
normalization reprocessing; an unchanged replacement does not reset applied state.
The saved display metadata is immutable and is preserved and returned unchanged.
Missing If-Match is a bad request; an outdated ETag returns 412.

```
flexera-cli finops-billing billing-orphans mappings replace [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing billing-orphans mappings replace --org-id ORG_ID --id ID --body @request.json
  flexera-cli finops-billing billing-orphans mappings replace --org-id ORG_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema finops-billing billing-orphans mappings replace --example > request.json
```

### Options

```
      --body string                    raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run[=true|false]           print the planned operation as JSON and exit without calling the API
  -h, --help[=true|false]              help for replace
      --id string                      id (path, required); Identifier of the saved assignment, not a provider charge identifier.; required by API; format: uuid; illustrative example: "b2c3d4e5-f6a7-4901-acde-f12345678901"
  -i, --interactive[=true|false]       edit inputs in a terminal form, review a plan and approve with typed yes
      --target-sub-account-id string   targetSubAccountId (body); required by API; Destination subscription selected by the caller; billing-account parentage is not verified.; minLength: 1; illustrative example: "e9a0360e-2c18-428e-acd0-a76dcee0cc96"
      --yes[=true|false]               confirm the operation (required for destructive ops)
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

