## flexera-cli policy policy-template evaluate

Evaluate a policy template

### Synopsis

Evaluate a policy template

Evaluates the datasources or resources requested and returns the result. If the datasource or resource requires parameters or
        credentials they must be also provided if there is no default value.

```
flexera-cli policy policy-template evaluate [flags]
```

### Options

```
      --file string      Path to JSON payload file, or - to read from stdin
  -h, --help             help for evaluate
      --id string        Policy template ID; The unique identifier for the policy template.; required by API; illustrative example: "5b06ead5e0dacc007058c784"
      --project-id int   Project ID (optional; resolved from GRS for the org when omitted); The unique identifier for the project; required by API; format: int64; minimum: 1; illustrative example: 60073
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

* [flexera-cli policy policy-template](flexera-cli_policy_policy-template.md)	 - Policy templates (project-scoped)

