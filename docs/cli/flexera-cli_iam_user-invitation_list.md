## flexera-cli iam user-invitation list

Index user invitations

### Synopsis

Index user invitations

Index all invitations belonging to a user.

```
flexera-cli iam user-invitation list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam user-invitation list
```

### Options

```
  -h, --help            help for list
      --status string   status (query); Status of invitation; enum: ["pending","accepted","expired","declined"]; illustrative example: "pending"
      --view string     view (query); View used to render invitations; enum: ["default","tiny"]; illustrative example: "default"
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

* [flexera-cli iam user-invitation](flexera-cli_iam_user-invitation.md)	 - User Invitation operations (generated from the unified OpenAPI spec)

