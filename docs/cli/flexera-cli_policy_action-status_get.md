## flexera-cli policy action-status get

Show an action status

### Synopsis

Show an action status

Show retrieves the details of an action status.

```
flexera-cli policy action-status get [flags]
```

### Options

```
  -h, --help             help for get
      --id string        Action status ID; The unique identifier for the action status.; required by API; illustrative example: "5b06ead5e0dacc007058c784"
      --project-id int   Project ID (optional; resolved from GRS for the org when omitted); The unique identifier for the project; required by API; format: int64; minimum: 1; illustrative example: 60073
      --view string      Optional Policy action-status view; View used to render action statuses.; enum: ["default","extended"]; illustrative example: "default"
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

* [flexera-cli policy action-status](flexera-cli_policy_action-status.md)	 - Policy action statuses (project-scoped)

