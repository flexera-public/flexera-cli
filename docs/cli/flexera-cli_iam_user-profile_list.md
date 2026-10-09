## flexera-cli iam user-profile list

Show current user's profile

### Synopsis

Show current user's profile

Returns the profile of the currently authenticated user.

**isNameManagedByIDP**: When `true`, the user last authenticated via an external SSO/SAML
identity provider. This means name edits made via PATCH /iam/v1/users/me may be overridden
on the user's next SSO login if the IdP attribute mapping still points to different values.
The UI should display a warning when this flag is true.

**Authentication**: Only the authenticated user's own profile is returned. There is no
admin variant of this endpoint — use GET /iam/v1/orgs/{orgId}/users/{id} for viewing
another user's profile.

```
flexera-cli iam user-profile list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam user-profile list
```

### Options

```
  -h, --help          help for list
      --view string   view (query); View used to render the user's profile; enum: ["default","extended"]; API default: "default"; illustrative example: "extended"
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

* [flexera-cli iam user-profile](flexera-cli_iam_user-profile.md)	 - User Profile operations (generated from the unified OpenAPI spec)

