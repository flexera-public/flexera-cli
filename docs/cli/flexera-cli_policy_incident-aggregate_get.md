## flexera-cli policy incident-aggregate get

Show an incident aggregate.

### Synopsis

Show an incident aggregate.

Show retrieves the details of an incident aggregate.

```
flexera-cli policy incident-aggregate get [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli policy incident-aggregate get --org-id ORG_ID --incident-aggregate-id INCIDENT_AGGREGATE_ID
```

### Options

```
  -h, --help                           help for get
      --incident-aggregate-id string   incidentAggregateId (path, required); The unique identifier for the incident aggregate.; required by API; illustrative example: "5b06ead5e0dacc007058c784"
      --view string                    view (query); View used to render incident aggregates.; enum: ["default"]; illustrative example: "default"
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

* [flexera-cli policy incident-aggregate](flexera-cli_policy_incident-aggregate.md)	 - Incident Aggregate operations (generated from the unified OpenAPI spec)

