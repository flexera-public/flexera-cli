## flexera-cli cli search

Find commands for a task

### Synopsis

Search registered commands offline using deterministic lexical ranking. Usage values are synopses, not copy-ready commands: replace uppercase placeholders with your own values. Read-only excludes unclassified curated side effects; it is not a guarantee against normal authentication or output-file activity.

```
flexera-cli cli search <words...> [flags]
```

### Options

```
      --action string    restrict to a spec action (list|get|create|update|replace|delete|action)
  -h, --help             help for search
      --limit int        maximum results (default 10)
      --read-only        only HTTP reads and explicitly classified read-only curated commands
      --service string   restrict to an API service (command, alias, or service id, e.g. bill-analysis, ba, bill_analysis)
      --tag string       restrict to a spec tag
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

* [flexera-cli cli](flexera-cli_cli.md)	 - Discover CLI commands and API schemas

