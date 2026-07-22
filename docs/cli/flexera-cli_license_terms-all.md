## flexera-cli license terms-all

Create license term

```
flexera-cli license terms-all [flags]
```

### Options

```
      --additional string                  additional (body)
      --body string                        raw JSON body (inline | @file | @-); overrides body field flags
      --effective-at string                effectiveAt (body)
      --effective-date-reminder-days int   effectiveDateReminderDays (body)
      --end-date-reminder-days int         endDateReminderDays (body)
      --ends-at string                     endsAt (body)
  -h, --help                               help for terms-all
      --hidden                             hidden (body)
      --is-active                          isActive (body)
      --is-vendor-set-provisioned-count    isVendorSetProvisionedCount (body)
      --license-id string                  licenseId (path, required)
      --managed-app-id string              managedAppId (body)
      --name string                        name (body)
      --sku string                         sku (body)
      --term-id string                     termId (body)
      --term-type string                   termType (body)
      --unique-id string                   uniqueId (body)
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

* [flexera-cli license](flexera-cli_license.md)	 - License operations (generated from the unified OpenAPI spec)

