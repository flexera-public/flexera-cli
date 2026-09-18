## flexera-cli rule-based-dimension from_csv generate

Generate rule-based-dimension specs from CSV

```
flexera-cli rule-based-dimension from_csv generate [flags]
```

### Options

```
      --case-insensitive          make generated conditions case-insensitive (default true)
      --column-config string      JSON map of output-column overrides ({header:{id,name,skip}})
      --csv string                alias for --file
      --effective-at string       effective month for generated rules (for example 2024-01)
  -f, --file string               CSV file path, or - for stdin (required)
  -h, --help                      help for generate
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
      --login-base-url string   override login base URL
      --org-id int              organization ID
  -o, --output string           output format (json|table)
      --refresh-token string    OAuth refresh token
      --zone string             API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli rule-based-dimension from_csv](flexera-cli_rule-based-dimension_from_csv.md)	 - Generate and apply rule-based dimensions from CSV

