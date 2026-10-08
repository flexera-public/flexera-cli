# Plan: service-grouped command tree + curated de-duplication

Status: **draft — for review**

## Problem

1. **Flat namespace.** ~120 generated tag commands plus 8 curated roots all live at
   the top level. `flexera-cli --help` is a wall of resources with no indication of
   which API service owns them.
2. **Global collision hacks.** Because every tag shares one namespace, the generator
   appends service suffixes to leaves: `project list-iam` / `project list-grs`,
   `customization-type list-policy` / `get-iam`, `onboarding delete-uobs` /
   `delete-data-inventory`, `user orgs` (GRS) vs `user ...` (IAM).
3. **Curated duplicates.** Several hand-written commands (mainly `finops`, plus
   `curated list`) re-wrap the *same* SDK operation the generated tree already
   exposes, with fewer features (no `--dry-run` plan, no schema/body flags, no
   interactive mode, no catalog identity).

## Goals

- Top level = **API service** (as named in `unified-openapi/specs.yaml`), then
  resource (OpenAPI tag), then action — i.e. `flexera-cli <service> <resource> <action>`.
- Curated commands keep existing **only** when they add behavior beyond the
  generated/SDK operation (multi-step, auto-resolution, local file processing,
  filtering). They are mounted *inside* the owning service group.
- No manual edits to generated files; everything flows from `cmd/regencli`.

## Non-goals

- Renaming awkward generated leaves (`create-all-7`, `plans-all-2`,
  `billing-center-i-ds`). Tracked separately; only stutter created by this change
  is addressed (see "Hoisting").

---

## 0. Upstream spec metadata (unified-openapi)

Rather than infer the service from operationId prefixes, the unified-openapi
merge step (which already reads `specs.yaml`) stamps service metadata into the
unified spec. Both `unified-go-client` and `flexera-cli` then consume it as data.
This follows the existing `x-flexera-auth`, `x-flexera-pagination`,
`x-flexera-action`, and `x-flexera-resource` extensions.

**Why per-operation:** tags are not service-unique in the merged spec —
`Onboarding` (Divnt, Uobs), `User` and `Project` (Grs, Iam), and
`Customization Type` (Iam, Policy) each span two services. Tag-level annotation
alone cannot place those operations.

### Root: service registry

```jsonc
"x-flexera-services": {
  "bill_analysis": {
    "title": "Bill Analysis",                // from specs.yaml name, vendor prefix stripped
    "description": "Cost queries, anomalies, forecasts, dashboards, dimensions",
    "vendor": "rightscale",                  // specs.yaml vendor
    "version": "",                           // specs.yaml version
    "source": "https://…/openapi.json",      // specs.yaml source.url
    "pathPrefix": "/bill-analysis",
    "operationIdPrefix": "BillAnalysis",
    "docsUrl": "https://developer.flexera.com/…",
    "cli": { "name": "bill-analysis", "aliases": ["ba"] }  // optional; CLI may override
  },
  "iam": { "title": "Identity and Access Management", "cli": { "name": "iam" }, … }
}
```

Optionally also emit the Redoc-standard `x-tagGroups`
(`[{ "name": "Bill Analysis", "tags": [...] }]`) so rendered API docs group the
same way. Because of the shared tags above, this needs service-qualified
tag names (see below) to be exact.

### Per operation

```jsonc
"post": {
  "operationId": "BillAnalysis_costs_aggregated",
  "tags": ["Costs"],
  "x-flexera-service": "bill_analysis",   // NEW — key into x-flexera-services
  "x-flexera-resource": "cost-aggregated",
  "x-flexera-action": "create"
}
```

### Per tag (optional)

`"x-flexera-service": "iam"` on tag objects whose operations are all one service,
and `"x-flexera-services": ["grs","iam"]` on the 4 shared tags. Alternative
(cleaner, but a breaking rename for SDK users): split shared tags into
`IAM Project` / `GRS Project`, etc.

### Consumers

| Consumer | Use |
|---|---|
| `unified-go-client` | Service constant/enum + `ServiceOf(operationId)`; per-service base path / server overrides (replaces ad-hoc Optima/GRS base-URL plumbing); optional doc grouping. |
| `cmd/gencli` / `cmd/regencli` | Group by `x-flexera-service`; group `Short`/`Long` from `title`/`description`; names/aliases from `cli` (CLI-side override table wins). |
| `internal/catalog` | Add `service` + `serviceTitle` to each entry → `cli search --service iam`, service shown in `cli schema`, synonyms ("Bill Analysis", "IAM", "FinOps Billing") resolve to groups. |

Fallback: if `x-flexera-service` is missing on an operation, regencli fails
(no silent prefix inference), keeping the spec as the source of truth.

---

## 1. Service groups

Service comes from `x-flexera-service` (§0). Display names/aliases come from
`x-flexera-services[*].cli`, with a CLI-side override table in `cmd/regencli` for
anything not yet upstream. Current mapping:

| operationId prefix      | specs.yaml `service`     | New top-level            | Aliases          | Ops |
|-------------------------|--------------------------|--------------------------|------------------|----:|
| BillAnalysis            | bill_analysis            | `bill-analysis`          | `ba`             | 43 |
| BillUpload              | bill_upload              | `bill-upload`            |                  |  6 |
| BillingCenterService    | billing_center_service   | `billing-center`         | `bc`             | 16 |
| Budget                  | budget                   | `budget`                 |                  |  7 |
| Cred                    | cred                     | `credential`             | `cred`           | 52 |
| Divnt                   | divnt                    | `data-inventory`         | `divnt`          |  3 |
| FinopsBilling           | finops_billing           | `finops-billing`         |                  | 38 |
| FinopsCustomizations    | finops_customizations    | `finops-customizations`  |                  | 25 |
| FinopsOnboarding        | finops_onboarding        | `finops-onboarding`      |                  | 47 |
| Graphql                 | graphql                  | `graphql`                |                  |  2 |
| Grs                     | grs                      | `grs`                    |                  |  2 |
| Iam                     | iam                      | `iam`                    |                  |129 |
| OptimaRecommendations   | optima_recommendations   | `recommendations`        | `optima-recommendations` | 2 |
| Policy                  | policy / governance      | `policy`                 |                  | 51 |
| Risk                    | risk                     | `risk`                   |                  | 53 |
| Saas                    | saas                     | `saas`                   |                  | 70 |
| Uobs                    | uobs                     | `unified-onboarding`     | `uobs`           | 16 |
| Vis                     | vis                      | `it-visibility`          | `vis`            |  7 |
| Auth (curated)          | —                        | `auth`                   |                  |  1 |

Cross-cutting roots stay: `auth`, `cli` (search/schema), `completion`, `help`.

Top-level count: **~128 → 21**.

### Hoisting (avoid stutter)

When a service has a tag whose kebab name equals the service name (singular/plural
insensitive), that tag's leaves are hoisted into the service group:

- `bill-upload bill-upload list` → `bill-upload list`
- `budget budget get` → `budget get`
- `recommendations recommendations list` → `recommendations list`
- `graphql graphql graphql` → `graphql graphql` (leaf rename deferred)

### Suffix removal

Leaf disambiguation suffixes are now computed **per service**, not globally, so
the cross-service ones disappear:

| Before                              | After                                     |
|-------------------------------------|-------------------------------------------|
| `project list-iam`                  | `iam project list`                        |
| `project list-grs`                  | `grs project list`                        |
| `customization-type list-iam`       | `iam customization-type list`             |
| `customization-type get-policy`     | `policy customization-type get`           |
| `onboarding delete-uobs`            | `unified-onboarding onboarding delete`    |
| `onboarding delete-data-inventory`  | `data-inventory onboarding delete`        |
| `user orgs` (GRS)                   | `grs user orgs`                           |

---

## 2. Curated command disposition

Rule: **drop** if it calls exactly one SDK operation that the generated tree
already exposes and adds nothing beyond it; **keep & move** if it composes calls,
auto-resolves IDs, processes local files, or filters results.

### Drop (pure duplicates)

| Curated command                          | SDK operation                              | Generated replacement                         |
|------------------------------------------|--------------------------------------------|-----------------------------------------------|
| `finops billing-center list`             | BillingCenterService_BillingCenters_index  | `billing-center billing-centers list`         |
| `finops billing-center get`              | BillingCenterService_BillingCenters_show   | `billing-center billing-centers get`          |
| `finops billing-center allocation-table` | …_show_allocation_table                    | `billing-center billing-centers allocation-table-all` |
| `finops cost aggregated`                 | BillAnalysis_costs_aggregated              | `bill-analysis costs aggregated`              |
| `finops cost select`                     | BillAnalysis_costs_select                  | `bill-analysis costs create-select`           |
| `finops cost export-select`              | BillAnalysis_costs_exportSelect            | `bill-analysis costs create-export`           |
| `finops cost export-status`              | BillAnalysis_costs_exportSelectStatus      | `bill-analysis costs get`                     |
| `finops cost dimensions` ¹               | BillAnalysis_costs_dimensions              | `bill-analysis costs dimensions`              |
| `finops cost metrics` ¹                  | BillAnalysis_costs_metrics                 | `bill-analysis costs metrics`                 |
| `finops bill-month list` ²               | BillAnalysis_bill_months_search            | `bill-analysis bill-months list`              |
| `finops adjustment show`                 | BillAnalysis_adjustment_definition_show    | `bill-analysis adjustment-definition list`    |
| `finops adjustment update`               | BillAnalysis_adjustment_definition_update  | `bill-analysis adjustment-definition replace` |
| `finops anomaly-report`                  | BillAnalysis_anomalies_report              | `bill-analysis anomalies report`              |
| `finops forecast-report`                 | BillAnalysis_forecasts_report              | `bill-analysis forecasts create`              |
| `curated list`                           | (static text)                              | `cli search` / `--help`                       |

¹ Curated added `--dataset` (billing|cost) via a request editor. Upstream
`bill_analysis` neither documents nor reads it (front_service ignores `dataset`),
so the flag was a no-op and is dropped rather than patched into the spec.
² **Prerequisite:** curated used the raw client to tolerate empty date fields. Add an
acceptance test with an empty-date fixture against the generated leaf first.

Also removed: the `--optima-base-url` persistent flag (verify generated commands
honor the existing per-service base URL config before removal).

### Keep & move

| Before                                         | After                                                        | Why it stays |
|------------------------------------------------|--------------------------------------------------------------|--------------|
| `finops cost get`                              | `bill-analysis costs query`                                  | auto-routes aggregated/select, chunks >24 mo, resolves BCs |
| `finops recommendation list-usage-reduction`   | `recommendations list-usage-reduction`                       | BC auto-resolve + category filter |
| `finops recommendation list-rate-reduction`    | `recommendations list-rate-reduction`                        | same |
| `curated anomaly-investigation`                | `bill-analysis anomalies investigate`                        | multi-step AI workflow |
| `bill-upload push` / `verify`                  | `bill-upload push` / `verify` (unchanged path)               | create+upload+commit; local CSV check |
| `graphql query` / `generate`                   | `graphql query` / `generate`                                 | different endpoint / query builder |
| `rule-based-dimension bulk` / `from_csv …`     | `finops-customizations rule-based-dimension bulk` / `from-csv …` | bulk + CSV processing |
| `policy applied-policy|action-status|archived-incident|policy-template …` | `policy …` (unchanged path) | GRS project auto-resolution |
| `policy meta …`                                | `policy meta …`                                              | relationship-aware workflows |
| `grs project list`                             | `grs project list-for-org`                                   | org-scoped resolver (generated is user-scoped) |
| `user-orgs list`                               | `iam user-memberships orgs` (curated **overrides** generated leaf) | defaults `--id` from access token |
| `auth token create|client-credentials|refresh` | unchanged                                                    | cross-cutting |

Curated wrappers that wrap an exact operation (policy, user-memberships orgs)
keep their `flexera.operationId` annotation so catalog coverage stays 1:1.

---

## 3. Before / after examples

### Root help

```text
BEFORE                                   AFTER
flexera-cli                              flexera-cli
├── access-policy                        ├── auth
├── access-rule                          ├── bill-analysis
├── access-rules                         ├── bill-upload
├── adjustment-definition                ├── billing-center
├── allocation-table                     ├── budget
├── anomalies                            ├── credential
├── api-event                            ├── data-inventory
├── api-key-credential                   ├── finops-billing
├── … (~110 more)                        ├── finops-customizations
├── costs                                ├── finops-onboarding
├── curated                              ├── graphql
├── finops                               ├── grs
├── grs                                  ├── iam
├── policy                               ├── it-visibility
├── user-orgs                            ├── policy
└── cli                                  ├── recommendations
                                         ├── risk
                                         ├── saas
                                         ├── unified-onboarding
                                         └── cli
```

### `bill-analysis` (absorbs `finops` + `curated`)

```text
BEFORE                                         AFTER
costs aggregated                               bill-analysis
costs create-export                            ├── costs
costs create-select                            │   ├── aggregated
costs get                                      │   ├── create-export
costs dimensions                               │   ├── create-select
costs metrics                                  │   ├── get
finops cost get            (curated)           │   ├── dimensions
finops cost aggregated     (dup)               │   ├── metrics
finops cost select         (dup)               │   └── query        (curated, was finops cost get)
finops cost export-select  (dup)               ├── anomalies
finops cost export-status  (dup)               │   ├── aggregated | anomalies | get | report
finops cost dimensions     (dup)               │   └── investigate  (curated, was curated anomaly-investigation)
finops cost metrics        (dup)               ├── bill-months  list | download | update
finops bill-month list     (dup)               ├── adjustment-definition  list | replace
finops adjustment show     (dup)               ├── forecasts  create
finops adjustment update   (dup)               ├── billing-settings | currency-setting | commitment-reallocation-setting
finops anomaly-report      (dup)               ├── custom-dashboards | org-dashboards
finops forecast-report     (dup)               ├── custom-dimension | custom-dimensions
anomalies aggregated|anomalies|get|report      └── cloud-vendor-accounts
curated list               (dup)
curated anomaly-investigation (curated)
bill-months … / adjustment-definition … / forecasts …
```

### `iam` (129 ops, previously 37 top-level roots)

```text
BEFORE                                 AFTER
access-policy list                     iam access-policy list
group create                           iam group create
project list-iam                       iam project list
customization-type list-iam            iam customization-type list
user-memberships orgs --id 42          iam user-memberships orgs          (id from token)
user-orgs list                         iam user-memberships orgs
service-account-client client-secret   iam service-account-client client-secret
saml2-identity-provider get            iam saml2-identity-provider get
scim-user list                         iam scim-user list
```

### `policy` (curated project-scoped + generated org-scoped merged)

```text
BEFORE                                 AFTER
policy applied-policy list             policy applied-policy list           (curated, unchanged)
policy meta audit                      policy meta audit                    (curated, unchanged)
published-template list                policy published-template list
policy-manager summary                 policy policy-manager summary
policy-aggregate list                  policy policy-aggregate list
incident-aggregate get                 policy incident-aggregate get
customization-type list-policy         policy customization-type list
customization-value get                policy customization-value get
custom-catalog tag                     policy custom-catalog tag
unmanaged-incidents list               policy unmanaged-incidents list
```

### Smaller services

```text
BEFORE                                 AFTER
bill-upload list                       bill-upload list                     (hoisted)
bill-upload push                       bill-upload push                     (curated)
budget report                          budget report                        (hoisted)
cloud-vendor-account list              budget cloud-vendor-account list
billing-centers list                   billing-center billing-centers list
finops billing-center list             (removed → above)
recommendations list                   recommendations list                 (hoisted)
finops recommendation list-rate-…      recommendations list-rate-reduction  (curated)
rule-based-dimension bulk              finops-customizations rule-based-dimension bulk
rule-based-dimension from_csv generate finops-customizations rule-based-dimension from-csv generate
bill-connect aws list                  finops-onboarding bill-connect aws list
billing plans                          finops-billing billing plans
aws-credential get                     credential aws-credential get
license list                           saas license list
vulnerability list                     risk vulnerability list
device estimate                        it-visibility device estimate
connector list                         unified-onboarding connector list
onboarding onboarding-all              data-inventory onboarding list
grs project list                       grs project list-for-org             (curated)
project list-grs                       grs project list
user orgs                              grs user orgs
```

---

## 4. Implementation phases

0. **Upstream (unified-openapi)** — emit `x-flexera-services` registry and
   per-operation `x-flexera-service` from `specs.yaml`; optional `x-tagGroups`.
   Bump `unified-openapi/PIN`; regenerate `unified-go-client` (service enum /
   `ServiceOf`) and bump `go.mod`.
1. **Spec gaps** — add empty-date bill-months fixture test. Regenerate.
2. **Generator** (`cmd/regencli`, `cmd/gencli`)
   - Group operations by `x-flexera-service` (then tag); emit
     `internal/commands/<service>/<tag>/cmd_gen.go` and
     `internal/commands/<service>/group_gen.go` (`NewCmd()` with Short from
     `specs.yaml` `name`, aliases from the service table).
   - Per-service leaf de-duplication (removes `-iam`/`-grs`/`-policy` suffixes).
   - Hoisting rule for service-named tags.
   - Replace `curatedCanonical` (tag-level, global) with a
     `curatedOwned map[service]map[tag]set[leaf]` so curated can own a leaf
     (`iam user-memberships orgs`) without skipping the whole tag.
   - `register_gen.go` registers ~18 service groups instead of ~120 tags.
3. **Curated re-homing** — change `Attach` helpers to take a path
   (`Attach(root, "bill-analysis", "costs")`); move kept commands per §2; delete
   `internal/curated/finops` duplicate leaves (keep `cost query` + recommendations
   in new packages, e.g. `internal/curated/billanalysis`, `internal/curated/recommendations`);
   delete `internal/curated/workflows` `list`, `internal/curated/userorgs` (folded into
   the IAM override), and the empty `internal/curated/rulebaseddimensions/` dir.
4. **Catalog/discovery** — regenerate `catalog_gen.json` / `coverage_gen.json`
   (command paths gain a service prefix; operation count unchanged at 570, curated
   exact wrappers drop from 20 → ~13); update `internal/catalog/synonyms.go` and
   golden search tests.
5. **Compatibility (optional, see Q1)** — hidden, `Deprecated:` shim commands for
   old top-level paths that print the new path and forward args, for one release.
6. **Docs** — regenerate `docs/cli/`, completions; update README, `docs/usage.md`,
   `docs/agents.md`, `SKILL.md` examples.

## 5. Verification

- `go generate ./... && go test ./...`; regencli is idempotent (second run no diff).
- Coverage test: every spec operation maps to exactly one leaf; no operation lost.
- Tree test: no top-level command other than service groups + `auth`/`cli`/builtins.
- For each dropped curated leaf, an acceptance test proves the replacement path
  sends the identical request (method, path, query, body) against the fake server.

## 6. Open questions

1. **Back-compat:** deprecated shims for one release, or clean break (pre-1.0)?
2. **Names:** friendly (`credential`, `it-visibility`, `unified-onboarding`,
   `data-inventory`) with spec ids as aliases — or spec ids as primary?
3. **Optima-era services** (`bill-analysis`, `billing-center`, `recommendations`,
   `bill-upload`): fold under a single `finops` umbrella, or keep flat per service?
4. **`grs project list-for-org`** vs. making the curated resolver the canonical
   `grs project list` and renaming the user-scoped generated one.
5. `auth token client-credentials` / `refresh` are flag sugar over `auth token create`
   — keep or drop?
6. **Shared tags:** annotate with `x-flexera-services: [...]`, or split into
   service-qualified tags upstream (breaking for SDK tag-based grouping)?
7. Should CLI names/aliases (`x-flexera-services[*].cli`) live upstream in the
   spec, or stay CLI-local so the spec stays client-agnostic?

## Implementation status

Implemented across unified-openapi, unified-go-client, and flexera-cli. Decisions taken:

- Clean break: no deprecated aliases for old flat paths.
- Friendly service command names; spec service ids/short forms are aliases.
- No `finops` umbrella; curated `finops`/`curated`/`user-orgs` roots removed.
- Shared tags are generated once per owning service (no cross-service dedupe).
- CLI names, aliases, hoisting, and renames live in `cmd/regencli/services.go`.
- `billing-center` AllocationTable renamed to `org-allocation-table` to avoid a
  sibling collision (caught by the new staged-audit name/alias check).
- Curated GRS project listing kept as `grs project list-for-org` (token-derived
  user, org filter, legacy account-id mapping) next to generated `grs project list`.
- `FLEXERA_CLI_OPTIMA_BASE_URL` now also applies to generated Optima-backed commands.
- `cli search --service` filters by service command, alias, or spec id.
- `gendocs` prunes stale pages before regenerating.
- No merge-time parameter patches: `dataset` is not part of the Bill Analysis API.

Follow-up: replace the local `go.mod` replace of `unified-go-client` with a
version bump once the spec and client changes are published.
