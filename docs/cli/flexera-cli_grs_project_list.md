## flexera-cli grs project list

List GRS projects for the org

```
flexera-cli grs project list [flags]
```

### Options

```
      --api-version string    Optional X-Api-Version header (defaults to 2.0)
      --grs-base-url string   Override the GRS base URL (default: zone-specific grs-front host)
  -h, --help                  help for list
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

* [flexera-cli grs project](flexera-cli_grs_project.md)	 - GRS projects

