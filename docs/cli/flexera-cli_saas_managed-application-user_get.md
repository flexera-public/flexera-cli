## flexera-cli saas managed-application-user get

Show managed application user

### Synopsis

Show managed application user

Retrieves a managed application user identified by user ID.

```
flexera-cli saas managed-application-user get [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas managed-application-user get --org-id ORG_ID --user-id USER_ID
```

### Options

```
  -h, --help             help for get
      --include-all      includeAll (query); Flag to determine whether or not inactive users should be included in the results.; API default: false; illustrative example: true
      --user-id string   userId (path, required); userId identifies an user by given ID.; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "123"
      --view string      view (query); View used to render managed application user.; enum: ["default","minimal","excludeEvent","entitlement","activityByRange"]; illustrative example: "default"
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

* [flexera-cli saas managed-application-user](flexera-cli_saas_managed-application-user.md)	 - Managed Application User operations (generated from the unified OpenAPI spec)

