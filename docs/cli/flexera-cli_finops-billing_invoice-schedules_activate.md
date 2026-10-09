## flexera-cli finops-billing invoice-schedules activate

Activate an invoice schedule

### Synopsis

Activate an invoice schedule

Activates an invoice schedule (reactivating a previously deactivated schedule) and restores its future runs. Supply endMonthYear to bound the reactivated schedule; omit it to make the schedule open-ended. While inactive, schedules cannot be modified via PATCH; activate first. Supports optimistic locking through If-Match.

```
flexera-cli finops-billing invoice-schedules activate [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli finops-billing invoice-schedules activate --org-id ORG_ID --id ID
  flexera-cli finops-billing invoice-schedules activate --org-id ORG_ID --id ID --dry-run
```

### Options

```
      --dry-run                 print the planned operation as JSON and exit without calling the API
      --end-month-year string   endMonthYear (query); Optional last active billing month for the reactivated schedule (YYYY-MM). Omit it to reactivate indefinitely.; pattern: "^\\d{4}-\\d{2}$"; illustrative example: "2027-08"
  -h, --help                    help for activate
      --id string               id (path, required); Identifier of the invoice schedule; required by API; illustrative example: "sch_1"
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

* [flexera-cli finops-billing invoice-schedules](flexera-cli_finops-billing_invoice-schedules.md)	 - Invoice Schedules operations (generated from the unified OpenAPI spec)

