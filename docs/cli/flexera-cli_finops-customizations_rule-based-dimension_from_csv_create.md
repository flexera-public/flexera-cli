## flexera-cli finops-customizations rule-based-dimension from_csv create

Create or update rule-based dimensions from CSV (create)

```
flexera-cli finops-customizations rule-based-dimension from_csv create [flags]
```

### Options

```
      --case-insensitive          make generated conditions case-insensitive (default true)
      --column-config string      JSON map of output-column overrides ({header:{id,name,skip}})
      --continue-on-error         continue processing after a dimension fails
      --csv string                alias for --file
      --dry-run                   validate and report dimensions without making API calls
      --effective-at string       effective month for generated rules (for example 2024-01)
  -f, --file string               CSV file path, or - for stdin (required)
  -h, --help                      help for create
      --id-template string        Go template for generated IDs
      --name-template string      Go template for generated names
      --separator-header string   CSV header separating condition columns from output columns
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

* [flexera-cli finops-customizations rule-based-dimension from_csv](flexera-cli_finops-customizations_rule-based-dimension_from_csv.md)	 - Generate and apply rule-based dimensions from CSV

