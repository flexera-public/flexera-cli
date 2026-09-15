## flexera-cli compliance control

Get Standard Control Details.

```
flexera-cli compliance control [flags]
```

### Options

```
      --accounts strings       accounts (body)
      --body string            raw JSON body (inline | @file | @-); overrides body field flags
      --category strings       category (body)
      --etime string           etime (body)
  -h, --help                   help for control
      --imc                    imc (body)
      --level int              level (body)
      --provider-type string   providerType (body)
      --providers strings      providers (body)
      --regions strings        regions (body)
      --services strings       services (body)
      --standard-name string   standard_name (path, required)
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

* [flexera-cli compliance](flexera-cli_compliance.md)	 - compliance operations (generated from the unified OpenAPI spec)

