## flexera-cli credential list

Index a list of Credentials

### Synopsis

Index a list of Credentials

Index a list of Credentials using optional filters.

```
flexera-cli credential list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli credential list --org-id ORG_ID
```

### Options

```
      --filter string   filter (query); Optional filter for Credentials. The following filters are supported: | Filter | Allowed Operators | Example | | --- | --- | --- | | createdAt | eq le lt ge gt | createdAt ge '2019-08-07T23:59:59Z' | | id | co eq ne | id eq 'azure-ro' | | name | co eq ne | name co 'Azure' | | scheme | eq ne | scheme eq 'api-key' | | tags._tag-name_ | co eq ne | tags.cloud-pr... (see cli schema); illustrative example: "(id co 'azure-' and scheme eq 'oauth2')"
  -h, --help            help for list
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

* [flexera-cli credential](flexera-cli_credential.md)	 - Credential API

