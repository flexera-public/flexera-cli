# Agent-first implementation progress

Required phases **1–7 are implemented and generator-published**. Release remains
blocked on upstream publication/repinning of the verified SDK pagination fix.
Optional endpoint diffs (phase 8) remain deferred. See the plan's closeout section.

## Approved decisions

- Implement phases 1–7; defer optional endpoint-verified diffs.
- Use `huh`, an always-visible root help hint, and uncompressed catalog JSON.
- Sanitized uncompressed catalog: 4,051,683 bytes; approved guard: 5 MiB.
- No jq output/result caps. Use private temporary-file spooling to avoid
  accumulating streams in memory and withhold stdout until evaluation succeeds.
  Disk exhaustion is an error; infinite expressions require cancellation.
- Target service-account-client path IDs use `--target-client-id`, distinct
  from inherited OAuth `--client-id`.
- Choose the lowest documented numeric 2xx as the primary response schema;
  render the actual supported success response.
- Keep `Auth_Token_token` as spec-backed `auth token create`, not an exclusion.
- The sibling SDK may be changed and tested locally; publication and a new
  dependency pin remain necessary before release.

## Implemented foundations

- Shared auto/pretty/compact JSON formatter; destination-writer terminal detection.
- Config/environment-backed `--json-style` and printer initialization.
- JSON response migration for FinOps and workflow output; policy logs preserve
  their explicit text/Markdown contract, even when bytes look like JSON.
- Execution-boundary errors use shared JSON formatting without response shaping.
  Typed flag/argument/config errors exit 2; unrelated runtime errors remain exit 1.
  Unknown-command tokens are captured from argv/Cobra argument resolution, with
  offline top-three search suggestions and original Cobra text retained.
  Wrapped validation details/schema and jq parse positions serialize top-level.
- Shared dry-run encoding; the generator now passes its printer style.
  Checked-in generated commands have been regenerated.
- Private disk-spooled jq rendering and field projection helpers, including
  lossless numeric decoding and no partial output after evaluation failure.
  Root `--out-jq`, `--out-fields`, and `--raw-output` are now CLI-only persistent
  flags; jq compiles once and field paths validate before command execution.
  The known `values` envelope is selected from catalog metadata when available.
- Transactional generated-command staging, symbol verification, publication
  rollback, and injected-failure tests.
- Generator metadata from assigned operations, operation annotations, safe
  illustrative help, 202/other supported success responses and collision fallback.
- Catalog data model, parsing, reference checks, reference closure and final
  bill-connect path mapping builder. Catalog, coverage report, command tree and
  registry publish together with rollback on publication failure.
- Complete coverage: 550 generated operations plus 20 exact curated operations,
  representing all 570 annotated spec operations with no exclusions.
- Isolated staged compilation/live-tree audit and permanent two-way catalog/tree
  tests; zero inherited long-name/shorthand collisions in the generated tree.
- Scalar date/date-time query flags supported with SDK-compatible parsing;
  `usage-message-query list` and eight previously omitted 202-only operations
  are now registered. Full regeneration is byte-for-byte deterministic.
- `auth token create` form adapter, redacted dry-run and mock HTTP tests.
- Sibling SDK lossless pagination conversion/count accumulation and tests.
- Offline `cli schema`: live-tree/alias resolution, depth-limited cycle-safe
  schema expansion, request/response/params selection, and validated non-sensitive
  example-only output. Implicit config and auth initialization are skipped;
  explicit config and output/style flag/env precedence are tested.
- Offline `cli search`: lazy live-tree BM25 index, weighted lexical fields,
  stemming/synonyms, tag/action/read-only filters, shell-safe usage synopses,
  JSON/table output through the shared printer, and root help hint.
  Golden queries pass; full-tree index build measured about 13.4 ms on Apple M4
  Pro (warm catalog), below the 20 ms target. Curated-only side effects require
  explicit read-only classification. `--raw-output` now has shorthand `-r`.
- Generated text/structured/mixed output annotations and pre-execution shaping
  guards: text, binary, mixed, human curated listings, policy logs, help and
  completion reject output shaping with exit 2 before config/body/API work.
  JSON-style alone preserves bytes; JSON dry-run plans never undergo shaping.
- Raw request validation helpers: cached local OAS schema resolution, exactly-one
  JSON parsing with lossless numbers, request-mode multi-error constraints,
  structured value-free diagnostics, and typed decoding/re-encoding safeguards
  against dropped fields, explicit-null loss and rounded IDs. Schema skipping
  does not bypass syntax or data-loss checks. Real budget SDK fixture tested.
  Hand-written generator template now invokes validation/data-loss checks before
  client creation, and uses effective decoded body bytes for safe JSON plans.
  CLI-only `--no-validate` is declared, but commands without the new validation
  annotation reject it explicitly. Commands are now generator-published with
  validation and credential-free preview ordering.
- `Plan` and `ConfirmPlan` helpers preserve dryRun/destructive/plan top-level
  keys, add ok/skipped/unsupported validation status, redact schema writeOnly/
  password fields plus sensitive names in body and params, and mark non-replayable
  previews redacted. Raw upload bytes are omitted; request data remains intact.
  Temporary generated budget fixtures prove invalid input and previews require
  no client/auth/HTTP, skipped validation cannot bypass data loss, and the preview
  body matches the mocked SDK wire body. Generated artifacts were subsequently
  published through generators only, following explicit user approval.
- Parameter validation helpers enforce flag/config presence separately from
  zero/false values, schema enum/format/numeric/array constraints, and required
  inputs even when constraints are skipped. Root dependencies retain org-ID
  source presence. Hand-written templates call parameter checks before bodies
  and API clients and reuse effective values in plans; temporary SDK fixtures
  pass. Checked-in commands now use the validated effective parameter values.
- Human `Plan.RenderHuman` helper shows every effective parameter and complete
  body fields with the same secret redaction as JSON previews. It labels supplied
  fields rather than server-side diffs and does not print nonexistent body files
  or secret-bearing replay commands. Short safe bodies have shell-quoted invocation
  examples; other bodies may be explicitly saved to a new private `0600` file.
- Local JSON-write `-i` form mode uses huh behind a scripted Prompter interface,
  terminal guards on actual stdin/stderr, editable prefill, required/optional field
  widgets, safe editor fallback, redacted plans, exact yes approval, --yes bypass,
  and dual-output credential-free dry-runs. No implicit prompts without -i.
- Example safety strips unsafe/invalid request samples and secret schema samples
  before publication. Nested help/schema-export paths share the registry mapping.
- Safe HTTP debug traces omit headers/queries/bodies; Ctrl-C/SIGTERM cancellation
  reaches jq projection/evaluation and cleans private spools. Strict rational
  numeric constraints and explicit UUID/email/URI/calendar validators complement
  OAS structural checks. Untyped response JSON is decoded losslessly from wire bytes.

## Audit evidence and remaining work

The initial live tree had 603 runnable commands. Collisions included 16 Azure/
Databricks body auth fields, five service-account-client path IDs, four query
org IDs, and one body org ID. Regeneration resolved these collisions and the
final full-tree gate passes. See the migration note below.

The baseline generator audit selected 560 of 570 annotated spec operations
before curated tag suppression. It omitted the form-body auth operation,
eight 202-only operations and the usage-over-time query with date/time params.
All these gaps are now represented by registered operations and recorded in
`internal/catalog/coverage_gen.json`.
Do not classify a multi-step curated workflow as an exact operation merely
because its tag resembles a spec tag.

The embedded catalog now contains 570 entries and 838 transitively referenced
schemas. Its bootstrap test has been replaced by published-artifact and live-tree
completeness tests. SCIM properties literally named `$ref` are handled as property
definitions rather than schema-reference keywords.

## Release blocker

The pinned SDK `92b58ffbc17f` still rounds merged page numbers. Full CLI tests pass
using the isolated local workspace `/tmp/flexera-cli-finish.work` with the sibling
SDK fix, including generated pagination/field/jq regressions above 2^53. The ordinary
pinned suite deliberately exposes the three merged-pagination failures until an
upstream version with the fix is published and pinned. No local replacement is
checked into go.mod. Do not ship by skipping those regression tests.

No upstream commits, dependency repinning or live API testing have been performed.

## Flag migration note

- Azure CSP, Azure EA Management, Azure MCA and Databricks bill-connect body
  credentials use `--body-client-id` and `--body-client-secret`; bare names
  remain OAuth authentication configuration. JSON keys do not change.
- Service-account-client target path IDs use `--target-client-id`, not the
  authentication `--client-id`.
- Refresh-token query org IDs and user-setting-blob query org IDs use inherited
  `--org-id`/environment/config. User-setting-blob body org ID uses `--body-org-id`.
- No ambiguous old-spelling aliases are retained. `regulatory-compliance --fields`
  and GraphQL `--query` retain their original API/body meanings.