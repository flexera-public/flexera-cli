## flexera-cli iam organization-invitation create

Create an invitation to an org

### Synopsis

Create an invitation to an org

Create an invitation to an organization, which granting roles or membership to groups within an org.
The invitation may also grant access to select projects within the org.

```
flexera-cli iam organization-invitation create [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam organization-invitation create --org-id ORG_ID --body @request.json
  flexera-cli iam organization-invitation create --org-id ORG_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema iam organization-invitation create --example > request.json
```

### Options

```
      --body string                 raw JSON body (inline | @file | @-); overrides body field flags
      --dry-run                     print the planned operation as JSON and exit without calling the API
  -h, --help                        help for create
  -i, --interactive                 edit inputs in a terminal form, review a plan and approve with typed yes
      --invitee-email string        inviteeEmail (body); required by API; Email of the invitee; pattern: "^[a-zA-Z0-9!#$%\u0026'*+/=?^_.`{|}~-]+@(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\\.)+[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9]$"; illustrative example: "taylor@example.com"
      --invitee-first-name string   inviteeFirstName (body); First name of the invitee. If skipEmailNotification=True, this first name will be used to create the user. If the user already exists, the user's first name will not be updated.; maxLength: 256; illustrative example: "Jane"
      --invitee-last-name string    inviteeLastName (body); Last name of the invitee. If skipEmailNotification=True, this last name will be used to create the user. If the user already exists, the user's last name will not be updated.; maxLength: 256; illustrative example: "Smith"
      --is-sso                      isSso (body); Used to denote if this is a SSO user invite; illustrative example: true
      --skip-email-notification     skipEmailNotification (body); When true, no email notification will be sent to the invitee. If the email is not sent, the invitation is still created within Flexera One and can be accepted by the user in the user settings section. When this parameter is true, the inviteeFirstName and inviteeLast Name are required.; API default: false; illustrative example: false
      --yes                         confirm the operation (required for destructive ops)
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

* [flexera-cli iam organization-invitation](flexera-cli_iam_organization-invitation.md)	 - Organization Invitation operations (generated from the unified OpenAPI spec)

