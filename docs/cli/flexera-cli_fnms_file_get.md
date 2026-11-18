## flexera-cli fnms file get

Stream FNMS download bytes unchanged to stdout or a new file

### Synopsis

Stream FNMS download bytes unchanged to stdout or a new file

Flexera One ITAM File Show API provides a `Location` header containing the URL to download the file.

```
flexera-cli fnms file get [flags]
```

### Options

```
      --destination string   write raw bytes to a new file, or - for stdout; existing files are never overwritten (default "-")
      --filename string      FNMS download target (required); Specify the name of the file to retrieve; required by API; illustrative example: "b0db7056f3b382e9efbdf154f469d407980e98ffc39973dc95730226097e7419.zip"
  -h, --help[=true|false]    help for get
```

### Options inherited from parent commands

```
      --access-token string        static bearer access token
      --api-base-url string        override API base URL
      --client-id string           OAuth client ID
      --client-secret string       OAuth client secret
      --config string              config file (default $HOME/.flexera/config.yaml)
  -d, --debug[=true|false]         log HTTP requests/responses to stderr (Authorization redacted)
      --json-style string          JSON whitespace style (auto|pretty|compact) (default "auto")
      --login-base-url string      override login base URL
      --no-validate[=true|false]   skip API schema constraints (never JSON syntax or request data-loss checks)
      --org-id int                 organization ID
      --out-fields string          project JSON output fields (comma-separated paths)
      --out-jq string              shape JSON output with a jq expression
  -o, --output string              output format (json|table)
  -r, --raw-output[=true|false]    write jq string results without JSON quotes (requires --out-jq)
      --refresh-token string       OAuth refresh token
      --zone string                API zone (nam|eu|apac|test)
```

### SEE ALSO

* [flexera-cli fnms file](flexera-cli_fnms_file.md)	 - FNMS files

