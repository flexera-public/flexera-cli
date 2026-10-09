## flexera-cli iam contracts list

Index an org's contracts

### Synopsis

Index an org's contracts

Retrieve a list of contracts with the following filtering options:

            - **Only orgId**:
            - Returns all contracts where orgId is either initiator or target

            - **init and target parameters**:
            - If both init and target are provided → returns only contracts between those two organizations.
            - If only target is provided → returns contracts between orgId (as initiator) and the target.
            - If only init is provided → returns contracts between orgId (as target) and the init.

            - **latest_only parameter**:
            - When set to true with init or target  → returns only the most recent contract between those two orgs.
            - When set to true with only orgId → returns the latest contract for each pair involving that org.

```
flexera-cli iam contracts list [flags]
```

### Examples

```
Illustrative only: replace uppercase tokens; provide your own request.json for body input.
  flexera-cli iam contracts list --org-id ORG_ID
```

### Options

```
  -h, --help          help for list
      --init int      init (query); Initiator org ID; illustrative example: 1234
      --latest-only   latest_only (query); When set to true, returns only the most recent contract between the given organizations; illustrative example: true
      --target int    target (query); Target org ID; illustrative example: 1234
      --view string   view (query); View used to render the contract; enum: ["default","index","extended"]; illustrative example: "extended"
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

* [flexera-cli iam contracts](flexera-cli_iam_contracts.md)	 - Contracts operations (generated from the unified OpenAPI spec)

