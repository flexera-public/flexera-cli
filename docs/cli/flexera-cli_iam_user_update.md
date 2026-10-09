## flexera-cli iam user update

Update user name

### Synopsis

Update user name

Updates the first name and/or last name of a user affiliated to an org.

**Required privilege**: iam:org_user:update (org_owner or equivalent).

**Multi-org security**: The caller must have admin access in *every* org the target user
belongs to. If the user is affiliated with orgs A, B, and C, the caller needs access
in all three — not just the org specified in the request. This prevents an admin in one
org from modifying a shared user's name without consent from admins in other orgs.

**Sync guard**: A successful update sets an internal `updated_by` marker
(`admin:{callerID}`) on the user record. The User Management Service uses this
marker to reject subsequent AD/SAML sync writes that would overwrite the manually-set
name. The marker is cleared when the IdP profile master sync next runs with a value
that matches the current name.

**isNameManagedByIDP**: The response includes this flag. When `true`, the org's SAML
IDP has profile-master enabled, meaning the name edit may be overridden on the user's
next SSO login if the AD/SAML attribute mapping still points to a different value.

```
flexera-cli iam user update [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam user update --org-id ORG_ID --id ID --body @request.json
  flexera-cli iam user update --org-id ORG_ID --id ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema iam user update --example > request.json
```

### Options

```
      --body string         raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run             print the planned operation as JSON and exit without calling the API
      --first-name string   firstName (body); First name to set.; minLength: 1; illustrative example: "Jane"
  -h, --help                help for update
      --id int              id (path, required); ID of the user; required by API; minimum: 1; illustrative example: 12345
  -i, --interactive         edit inputs in a terminal form, review a plan and approve with typed yes
      --last-name string    lastName (body); Last name to set.; minLength: 1; illustrative example: "Smith"
      --yes                 confirm the operation (required for destructive ops)
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

* [flexera-cli iam user](flexera-cli_iam_user.md)	 - User operations (generated from the unified OpenAPI spec)

