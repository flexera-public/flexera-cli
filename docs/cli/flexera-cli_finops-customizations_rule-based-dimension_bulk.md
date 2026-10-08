## flexera-cli finops-customizations rule-based-dimension bulk

Create or update rule-based dimensions from JSON input

```
flexera-cli finops-customizations rule-based-dimension bulk [flags]
```

### Options

```
      --continue-on-error   continue processing after a dimension fails
      --dry-run             validate and report dimensions without making API calls
      --file string         path to JSON input, or - to read JSON from stdin
  -h, --help                help for bulk
      --input string        inline JSON input (mutually exclusive with --file)
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

* [flexera-cli finops-customizations rule-based-dimension](flexera-cli_finops-customizations_rule-based-dimension.md)	 - Rule-Based Dimension operations (generated from the unified OpenAPI spec)

