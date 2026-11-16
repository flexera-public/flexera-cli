## flexera-cli finops-billing billing-orphans mappings delete

Delete an Azure orphan cost mapping

### Synopsis

Delete an Azure orphan cost mapping

Delete an individual mapping using the ETag from show.
Eligible charges become discoverable again once the mapping is deleted.
Deletion does not immediately revert attribution in processed billing data;
normalization reprocessing is required. Existing work-order snapshots remain unchanged.
Missing If-Match is a bad request; an outdated ETag returns 412.

```
flexera-cli finops-billing billing-orphans mappings delete [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing billing-orphans mappings delete --org-id ORG_ID --id ID
  flexera-cli finops-billing billing-orphans mappings delete --org-id ORG_ID --id ID --dry-run
```

### Options

```
      --dry-run[=true|false]   print the planned operation as JSON and exit without calling the API
  -h, --help[=true|false]      help for delete
      --id string              id (path, required); Identifier of the saved assignment, not a provider charge identifier.; required by API; format: uuid; illustrative example: "b2c3d4e5-f6a7-4901-acde-f12345678901"
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

* [flexera-cli finops-billing billing-orphans mappings](flexera-cli_finops-billing_billing-orphans_mappings.md)	 - mappings operations

