## flexera-cli saas managed-application update

Update managed application

### Synopsis

Update managed application

Updates a managed application identified by ID.

```
flexera-cli saas managed-application update [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli saas managed-application update --org-id ORG_ID --managed-app-id MANAGED_APP_ID --body @request.json
  flexera-cli saas managed-application update --org-id ORG_ID --managed-app-id MANAGED_APP_ID --body @request.json --dry-run
Validated illustrative body, when available (review before use):
  flexera-cli cli schema saas managed-application update --example > request.json
```

### Options

```
      --activity-threshold int    activityThreshold (body); A global activity threshold for the managed application. An activity threshold defines the number of days that can pass before a user is considered inactive. For example, with a threshold of seven days, if a user logins in on January 1st, if no other activity is seen, on January 8th, the user will be shown as inactive.; enum: [1,7,15,30,45,60,90,120,180]; illustrative example: 90
      --body string               raw JSON body (inline | @file | @-); overrides body field flags
      --description string        description (body); Description of the managed application; illustrative example: "Office 365"
      --dry-run                   print the planned operation as JSON and exit without calling the API
  -h, --help                      help for update
  -i, --interactive               edit inputs in a terminal form, review a plan and approve with typed yes
      --is-active                 isActive (body); required by API; Is managed application active? An active managed application is considered actively managed. If isActive is false, everything within the managed application is frozen and exists only for historical purposes.; illustrative example: true
      --managed-app-id string     managedAppId (path, required); managedAppId identifies an managed application by given Id; required by API; pattern: "^[0-9a-f]+$"; illustrative example: "456"
      --name string               name (body); Name of the managed application; illustrative example: "Office 365"
      --point-of-contact string   pointOfContact (body); Point of contact for the managed application; illustrative example: "support@flexera.com"
      --yes                       confirm the operation (required for destructive ops)
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

* [flexera-cli saas managed-application](flexera-cli_saas_managed-application.md)	 - Managed Application operations (generated from the unified OpenAPI spec)

