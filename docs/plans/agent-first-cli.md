# Plan: agent-first usability for flexera-cli

Status: **required phases 1–7 implemented and generator-published; release blocked on upstream SDK pagination publication/repin**
Inspired by: [Introducing cf: the agentic CLI for the entire Cloudflare API](https://blog.cloudflare.com/cloudflare-cf-cli-launch/)

## Implementation closeout — 2026-10-06

- The registered/catalog coverage contract is complete: **570 operations**, comprising
  **550 generated** leaves and **20 exact curated** wrappers, with no exclusions.
  `internal/catalog/coverage_gen.json` records every disposition. Search/schema use
  actual registered paths; multi-step workflows remain unannotated.
- Output policy, shaping, offline discovery, early JSON errors, schema/body/parameter
  validation, redacted JSON/human plans, and local JSON-write form mode are published.
  Generated files were updated **only by project generators**, never manually.
- Interactive mode uses `huh` behind `Prompter`, checks terminal stdin and stderr,
  supports editable prefill/optional fields and exact typed `yes`, and permits
  redirected stdout. Deep/unsupported schemas use a private editor file only when
  safe; otherwise they fail with `--body @file` guidance. Accessible mode is opt-in
  through `FLEXERA_CLI_ACCESSIBLE=1`. Editor configuration is an executable path,
  not a shell command. Explicit body export creates a new file with mode `0600`.
- Examples are sanitized and locally validated before catalog publication. The
  current generation omits 63 unsafe/invalid request samples and removes 167 unsafe
  schema sample annotations rather than advertising them as applicable input.
  The sanitized catalog is **4,051,683 bytes**; the approved limit is **5 MiB**.
- Tests include golden searches, both catalog/tree directions, inherited long-name
  and shorthand collisions, staged publication/rollback, temporary SDK fixtures,
  live scripted forms, byte-preserving downloads, wire/preview equality, secret-safe
  tracing, cancellation/spool cleanup, and paginated/non-paginated IDs above $2^{53}$.
  Index construction measured approximately 13.4 ms on Apple M4 Pro (warm catalog).
- **Release gate:** `go.mod` still pins SDK `92b58ffbc17f`, whose pagination merge
  rounds numbers. The approved sibling SDK fix passes the full suite through an
  isolated local Go workspace; the ordinary pinned suite intentionally fails the
  three merged-pagination regression cases until that fix is published and repinned.
  No temporary replacement is checked into `go.mod`; no upstream commit/push or
  live credentialed API tests were performed.

### Approved deviations and deferred work

- No jq result/output caps: results spool to a private temporary file, commit only
  after successful evaluation, and clean up on errors/cancellation. Disk exhaustion
  is an error, infinite expressions require cancellation, and jq expressions that
  build large individual values still consume memory. Signal cancellation is wired.
- Policy applied-policy logs retain their documented text/Markdown contract even
  when bytes resemble JSON; shaping/table are rejected. FinOps JSON endpoints fail
  on invalid JSON rather than silently treating raw bytes as structured results.
- The catalog's primary success schema is the lowest documented numeric 2xx;
  commands render their actual supported returned success status.
- Conflicting service-account target IDs use `--target-client-id`; body auth fields
  use `--body-client-id`/`--body-client-secret`, with JSON keys unchanged.
- Optional endpoint-verified update/replace diffs (phase 8), first-run state files,
  MCP/agent files, and future `x-flexera-keywords`/`x-flexera-example` annotations
  remain deferred. No deletions are inferred from omitted fields.

The design below retains the original baseline and requirements for review; the
closeout above and `agent-first-implementation.md` describe the implemented state.

## Goals

Make `flexera-cli` easy to use for coding agents that have never seen it before, and keep it pleasant for humans:

1. **Command search.** `flexera-cli cli search "<task>"` finds the right command among every registered generated operation and searchable curated command. The current spec has 570 `x-flexera-action` operations, but the live tree has about 536 generated operation leaves; do not imply every spec operation is currently emitted until the coverage audit closes that gap.
2. **JSON sized for the reader.** JSON is pretty-printed when stdout is a terminal and compact when it is piped or redirected.
3. **Guided input from the API schema:**
   - examples in `--help`
   - `cli schema` introspection
   - client-side body validation
   - an interactive form mode with a Terraform-style plan and a typed `yes` before applying.
4. **Client-side output shaping.** `--out-jq` (a jq expression, using gojq) and a simpler `--out-fields` projection. These names distinguish output shaping from API query/body fields.

## Non-goals

- Replacing the server-side `--filter` query params. They stay as they are.
- Adding a jq dependency to `unified-go-client`. All output shaping happens in this repo, in `internal/cli`; if needed, change the client's pagination merge only to preserve JSON number fidelity, never to perform output shaping.
- An MCP server, AGENTS.md or skill files. These can come later, but they don't block this plan.
- Search using an LLM or embeddings. Search is lexical, offline and deterministic.

## Original baseline (before implementation)

| Area | Where | Notes |
|------|-------|-------|
| Rendering | [internal/cli/printer.go](../../internal/cli/printer.go) → `flexera.Write` in unified-go-client `render_json.go` | JSON is always indented (`SetIndent("", "  ")`). Errors go through `WriteErrorJSON`, which is also indented. |
| Dry-run / confirm | [internal/cli/bodyflag.go](../../internal/cli/bodyflag.go) `ConfirmWrite` | Uses its own indenting encoder. Destructive commands without `--yes` return an error. |
| Body input | `ResolveBody`: `--body` inline, `@file`, `@-`; typed flags for scalar top-level fields | There is no schema validation. A body is decoded straight into the generated request type. |
| Codegen | [cmd/regencli](../../cmd/regencli/main.go) → [cmd/gencli](../../cmd/gencli/main.go) per tag → `internal/commands/<pkg>/cmd_gen.go` | `regencli` currently runs `RemoveAll(internal/commands)` before generation and continues after per-tag failures; regeneration is not transactional. `gencli.successCode` currently recognizes only 200, 201 and 204. |
| Spec | `unified-openapi/openapi3.json` (OAS **3.0.3**, about 8 MB) | 570 `x-flexera-action` operations, about 10.9k `example` keys, about 130 `examples` keys. |
| Curated renderers | `internal/curated/*` | Several use `deps.Printer.Render`, but `finops.go` also calls `writeRaw` at multiple response sites, `policy.go` writes logs directly, `workflows/curated.go` hard-codes `"json"`, and `curated list` emits text. Generated 204/no-JSON commands print `OK`. |
| Docs | `cmd/gendocs` → `docs/cli/` | Generated from the cobra tree, so it picks up `Example` automatically. |
| Errors | `internal/cli/errors.go` → `Execute` | Unknown-command/flag errors are plain text and can happen before `PersistentPreRunE` constructs `deps.Printer`; `main_smoke_test.go` tests this behavior. |
| Pagination | unified-go-client `CollectPages` | Returns an object with a merged `values` array plus envelope metadata, **not** a top-level array. |

### Flag-name collisions to guard against

Local flags shadow inherited persistent flags in cobra. Check both long names and shorthands against *all* inherited flag sets, not just the root. The output flags are `--out-jq` and `--out-fields`, **not** `--jq`, `--fields`, or `--query`; the generated `regulatory-compliance --fields` body flag and both `graphql --query` flags keep their existing meanings.

The current tree already has about 32 generated local flags that shadow inherited root flags, including `org-id`, `client-id` and `client-secret`. Before requiring a zero-collision tree, generate a report of every current collision with its command path, flag source and effective meaning. Recommended resolution: when a generated path/query `org-id` has the same meaning and type as the root persistent flag, use the inherited flag/config value rather than declaring a local shadow; when a generated body property conflicts with an inherited auth flag such as `client-id` or `client-secret`, expose it as `--body-<name>` while preserving the JSON key and reserve the bare name for auth. Document this intentional CLI spelling change and its migration note; do not preserve an ambiguous old spelling as an alias. Resolve any other collision by the same semantic rule, not by whichever flag Cobra happens to prefer. Do not treat the existing set as permanent exceptions: the final tree test should fail on unresolved collisions and on every newly introduced collision. Re-run it after each regeneration.

Before adding each new flag, build the real Cobra tree and assert there are no long-name or shorthand collisions. Do **not** rename unrelated existing flags unnecessarily.

## Design

### 0. Foundation: the generated operation catalog

Search and schema introspection need operation metadata at runtime. Generate metadata only for **successfully emitted** operations, then attach it to the command paths users can actually invoke.

- **New package: `internal/catalog/`**. It lives outside `internal/commands` so that `regencli`'s `RemoveAll` doesn't delete it.
  - `catalog.go` is hand-written. It embeds `catalog_gen.json`, parses it lazily (`sync.Once`), and exposes `Lookup(operationID)`, `All()` and `Schema(ref)`.
  - `catalog_gen.json` is written by `regencli` after registration has been generated. A missing or unreadable catalog fails the build rather than silently returning an empty index.
- **Entry shape (one per emitted operation):**
  ```json
  {
    "operationId": "Budget_Budget_create",
    "command": ["budget", "create"],
    "method": "POST", "path": "/finops-analytics/v1/orgs/{orgId}/budgets",
    "tag": "Budget", "resource": "budget", "action": "create",
    "summary": "...", "description": "...",
    "destructive": false, "paginated": false, "responseEnvelope": "none",
    "params": [{"flag": "org-id", "source": "config", "in": "path", "type": "integer", "required": true, "enum": null, "description": "..."}],
    "bodyFlags": ["name", "metric"],
    "requestSchema": {"$ref": "#/components/schemas/Budget_CreateRequestBody"},
    "responseSchema": {"$ref": "#/components/schemas/<actual-response-schema>"},
    "requestExample": { }
  }
  ```
  - Illustrative optional fields are omitted if the operation has no such schema. The request schema above is real; resolve the response schema from the selected success response, not from the example placeholder. `destructive` means "requires `--yes`" in today's generator, **not** "is an HTTP write". `responseEnvelope` is `values` only for a known paginated list; otherwise `none`.
  - A top-level `"schemas"` map holds **only the component schemas those entries reference**, found by walking `$ref`s transitively (including nested refs, with cycle detection). Preserve OAS 3.0 constructs such as `nullable`, `allOf`, `additionalProperties` and `required`. Measure the actual file before fixing the size budget.
- **Two-phase generation (no import cycle):**
  1. Refactor `cmd/gencli` so its `renderOp` values (after `assignVerbs`) drive both the Go template and a temporary metadata JSON file per tag, keyed by `operationId`. Record emitted flag names, CLI verb, request/response schema references, chosen 2xx response, and spec path/method. Do not independently rederive CLI verbs by scanning spec tags in `regencli`.
  2. First establish an operation coverage contract: every `x-flexera-action` operation must be either emitted as a registered generated leaf, intentionally represented by an exact curated operation, or explicitly excluded with a recorded reason. The current audit found about 536 registered generated leaves from 570 spec operations; in particular, eight operations have only a 202 success response, which `gencli.successCode` currently does not select. Extend success-response selection to supported 2xx responses, including 202 where the generated client has the corresponding response type, and test response-body/schema selection from the OAS `responses[status].content[mediaType].schema` shape. Keep intentional curated ownership distinct from an unexpected generator skip.
  3. `regencli` constructs special bill-connect nesting in `writeRegister`. Generate into a temporary staging tree; verify every expected tag, generated Go symbol, registration path, catalog entry and schema ref there before touching checked-in outputs. Any unexpected tag-generation failure, unsupported operation, duplicate ID/path, dangling ref or missing mapping aborts the run and leaves the current generated tree, registry and catalog unchanged. Intentional curated-owned tags must have an explicit disposition in the coverage report. After all checks pass, publish the staged generated files, registration and catalog together, with rollback on publication failure. Apply **the same** registration mapping when converting per-tag verbs to final command paths (including bill-connect child names). Write deterministic, sorted JSON outside `internal/commands`.
  4. `gencli` sets `Annotations: {"flexera.operationId": "<id>"}` on each emitted leaf. Add a tree-walk test using `app.NewRootCmd` that proves **both directions**: every annotated leaf has exactly one catalog entry with its actual `CommandPath` (strip the root executable name before comparison), and every catalog entry has an annotated leaf. This catches registry changes made later without introducing a `regencli` → `app` build cycle.
  5. Curated commands can opt into an ID only when they implement precisely that API operation. Multi-step workflows and curated-only routes remain searchable by Cobra help/flags but have no spec-backed schema or validation; omit their `schema` link.
- **Example synthesis** (shared helper in `cmd/gencli`, called from catalog and help generation). Precedence:
  1. media-type `example`
  2. first `examples[*].value` in sorted key order
  3. schema `example`
  4. per-property synthesis of **required writable** fields only; each uses the first of `example`, `default`, `enum[0]`, a format placeholder, or a type placeholder. Optional fields are discoverable in `cli schema`, not fabricated into a "minimal" sample.
  - Limits: depth 4, arrays get one element, and `oneOf`/`anyOf` use the first variant; resolve `$ref` safely, detect cycles and exclude read-only properties. Examples are *illustrative*, not a promise of valid IDs or server acceptance. Test each real example against its request schema; label failing/placeholder examples clearly rather than advertising them as ready to apply.
- **Guards added to `regencli`:**
  - Fail if the final top-level command name collides with `cli`, including curated aliases, and fail if the size of `catalog_gen.json` exceeds a measured budget (3 MB is a starting hypothesis, not an established limit). Verify catalog generation determinism by regenerating twice and comparing bytes.
  - Run the operation-coverage and flag-collision audits before publishing staged output. Fail on any unclassified operation or unresolved collision; emit a concise report of intentional curated ownership and exclusions so changes in coverage are reviewable.

### 1. `flexera-cli cli search`

A new top-level curated `cli` group, reserved for meta-commands about the CLI itself. The name avoids colliding with any future Flexera platform search.

```
flexera-cli cli search <words...> [--limit 10] [--tag <tag>] [--action list|get|create|...] [--read-only]
```

- **Corpus.** Walk the *already constructed* root Cobra tree once (pass it to the search builder; do not recursively call `app.NewRootCmd` from `cli search`). Index every runnable leaf; exclude `cli search` itself, hidden help/completion commands and duplicate aliases. If a leaf has the `flexera.operationId` annotation, enrich it from the catalog; otherwise use `Short`, `Long` and flag usage. Curated-only commands are searchable but have no schema link. Indexing must not run any leaf `RunE` or create an API client.
- **Ranking.** BM25 with a weight per field:
  - command path words ×3
  - summary ×2
  - tag/resource ×2
  - operationId split on camelCase ×2
  - flag/param names ×1
  - description ×1
- **Normalisation:**
  - lowercase, and split on non-alphanumerics and camelCase
  - light suffix stemming (`s`, `es`, `ing`, `ed`)
  - a small hand-maintained synonym map in `internal/catalog/synonyms.go`. Examples:
    - remove/destroy → delete
    - show/describe/fetch/read → get
    - add/new/make → create
    - edit/modify/patch/change → update
    - Domain terms: spend → cost, "cloud account" → connector, etc.
- **Output.** JSON by default, in this shape. The `usage` field includes required flags only.
  ```json
  [{"command": "flexera-cli budget create", "summary": "...", "action": "create",
    "destructive": false, "score": 7.4,
    "usage": "flexera-cli budget create --org-id ORG_ID --body BODY_JSON_OR_@FILE_OR_@-",
    "schema": "flexera-cli cli schema budget create"}]
  ```
  `usage` is a synopsis, not a copy-ready command; keep placeholders shell-safe and explain that users must replace them.
  `-o table` prints a compact list for humans. Use the *final* path and flag metadata; omit `schema` for curated-only commands. Include required `--org-id` (sourced from flag/env/config) and request-body requirement in `usage`. `--read-only` must filter by actual HTTP read/write semantics **plus curated side effects**, not by `destructive` (which means "requires `--yes`"). Allow missing tags/actions on curated entries instead of guessing.
- **Build cost.** The index is built at runtime, lazily, from the live registered command tree (currently roughly 600 generated and curated leaves). Target: under 20 ms, checked by a benchmark. No prebuilt index artifact.
- **Discovery hints** (cf appends the same kind of hint to `--help`):
  - The root `Long` gets one line: `Find commands for a task: flexera-cli cli search "<what you want to do>"`.
  - Unknown-command JSON suggestions use the shared search index; see the early-error contract in §2b.

### 2. JSON sized for the reader

- Move JSON response rendering policy into `internal/cli.Printer`. It becomes a struct built in `PersistentPreRunE` and keeps calling `flexera.WriteTable` for tables.
  ```go
  type Printer struct {
      Style  JSONStyle // auto|pretty|compact
      IsTTY  func(io.Writer) bool
      JQ     *gojq.Code // nil when --out-jq is absent
      Fields []fieldPath // parsed --out-fields paths
      RawOutput bool
  }
  ```
- **`auto` (the default)** pretty-prints when the writer is an `*os.File` for which `term.IsTerminal(fd)` is true (`golang.org/x/term`), and prints compact JSON otherwise.
  - Test buffers therefore get compact JSON. Tests that assert on indented output need updating, or they can set `Style: pretty`.
- **Override:** a new root persistent key `--json-style auto|pretty|compact` with env var `FLEXERA_CLI_JSON_STYLE` and a config-file key. Add it to `internal/cli/config.go`'s bound keys, validate the value at startup, and pass it to the printer. It follows existing viper precedence for normal API commands; the offline `cli search`/`cli schema` path accepts flag/env/default without reading the implicit config file (see §4b). Style applies to JSON on the *actual* destination writer (stdout for results, stderr for errors).
- **Output boundary (audit each call site before claiming coverage):**
  | Output | Policy |
  |--------|--------|
  | Structured JSON API results and `cli search` | `deps.Printer.Render`: TTY-aware JSON, `--out-fields`, `--out-jq`, and JSON/table format where supported. |
  | JSON-shaped curated responses | Decode valid JSON first and use the printer; migrate **all** `finops.writeRaw` response call sites and `policy applied-policy log` if they return JSON. When intentionally non-JSON, keep bytes unchanged and explicitly reject shaping/style/table flags instead of pretending to transform them. Preserve responses requiring raw bytes. |
  | Generated 204/no-JSON `OK`, human `curated list`, help, completion, and binary downloads | Preserve their existing text/byte contract; reject `--out-jq`, `--out-fields`, and `--raw-output` before executing the command, or add an explicitly documented structured JSON mode in a separate change. No silent no-ops. `--json-style` alone has no effect on text/binary. |
  | JSON dry-run plans | Route through the shared JSON formatter but **never** through response shaping; dry-run is a safety/inspection contract, not an API response. |
  | Errors on stderr | Use the shared JSON formatter (when applicable), **never** response shaping. See §2b for errors that occur before the printer exists. |
  Migrate `workflows/curated.go` to use `deps.Config.Output`; review `policy.go`, `finops.go`, `curated list`, generated 204 branches and every other direct stdout write using a code search. Add a tree-walk/command test for each output category. Do not silently turn binary data into JSON or change the meaning of `-o table` for commands that already reject it.
- `SetEscapeHTML(false)` is kept.
- Pretty and compact output are identical JSON values, so switching the default is non-breaking for anything that parses the output. It is a byte-level change for anything that greps it (release note).

### 2b. Early errors are not printer errors

`internal/cli/errors.go`'s `Execute` receives Cobra parse/unknown-command errors **before** `PersistentPreRunE`; `deps.Printer` and Viper config may not exist. Implement a small execution-level error writer which takes the root command's stderr, reads `--json-style` safely from already-parsed flags when available (otherwise defaults to `auto`; do not assume invalid config can be loaded), and uses the same JSON encoder as `Printer.RenderError` without depending on `Deps`. Keep exit code 1 for unrelated runtime errors; make flag syntax/config, unknown command, unsupported output options, schema lookup and jq compilation errors usage errors (exit 2). Prefer typed usage errors for new call sites; classify Cobra parse failures at the execution boundary rather than matching arbitrary API-error strings. Cobra's parse-error text should remain in the `error` value.

- For unknown command, create a typed execution error containing the failing parent path, unmatched token(s), and original Cobra error text; capture tokens at command resolution from the original argv rather than recovering them by parsing Cobra's message. Retain Cobra's error text and attach `suggestions` with the top three ranked command strings. If there is no match, emit `suggestions: []`. Only build the offline search index on this error path; do not require authentication, API access or a config file. Escape JSON in both the error and suggestions.
- For structured validation errors, define a typed error carrying `details` and `schema`; serialize these as top-level JSON properties instead of stringifying them inside `error`. Preserve `errors.As` through `ExitError`.
- `--help` still exits 0 and prints ordinary Cobra help. Update `main_smoke_test.go`'s existing plain-text unknown-command expectation intentionally, and test malformed flags, invalid config, unknown subcommands under curated parents, and errors before `PersistentPreRunE`.

### 3. `--out-jq` and `--out-fields`

Both are root persistent flags, applied in `Printer.Render` **before** formatting. Do not add aliases `--jq`, `--fields`, or `--query`. `--raw-output` / `-r` is persistent and meaningful only with `--out-jq`; reject it with exit 2 otherwise. For `-o table`, accept one structured result only; reject multiple results or raw-output with exit 2 rather than emitting an ambiguous table stream.

- **`--out-jq '<expr>'`** is compiled once with `github.com/itchyny/gojq` in `PersistentPreRunE`.
  - A compile error is a usage error and exits with code 2. The error JSON includes the gojq position.
  - Input: the response value round-tripped through `json.Marshal` / `json.Unmarshal` into `any`. The round trip only happens when `--out-jq` or `--out-fields` is set. Verify gojq's numeric input handling and use a lossless decoding strategy where supported so projecting IDs does not lose large-integer precision; test values exceeding 2⁵³. If that is not possible, reject unsafe numeric input rather than silently rounding IDs.
  - Preserve numeric precision before this printer boundary as well: `CollectPages` currently merges page values through a `float64`-based JSON map conversion. Update that merge boundary to preserve `json.Number`/raw numeric values where feasible, or detect and reject unsafe integer values before shaping. Tests must cover both paginated and non-paginated responses with IDs above 2⁵³; a printer-only round-trip test is insufficient.
  - Output: zero results produce empty stdout; one result renders as a normal JSON value; multiple results render one JSON value per line (NDJSON). For NDJSON, each item is compact even on a TTY; an `--json-style pretty` request with multiple results is rejected (exit 2) rather than producing invalid NDJSON. `--raw-output` / `-r` prints string results without JSON quotes; non-strings render as JSON. Example: `... --out-jq '.values[].id' -r | xargs ...`.
  - Iterate through **all** jq results, collecting output in memory before writing stdout; if evaluation produces an error, discard buffered results and return JSON error, exit 1. Document the memory trade-off and add a bounded-result/output guard with a clear error instead of risking unbounded memory from a jq expression.
  - No `input`, env or `$__loc__` extensions. Use a plain `gojq.Compile` without `WithEnvironLoader`, so a query can't read environment variables (which may hold tokens).
- **`--out-fields id,name,owner.email`** is parsed as a comma-separated list of dot paths and compiled to a gojq program:
  - For an array value, project each element. For an ordinary object, project that object. For a **known** `responseEnvelope: values` operation, project each element of `.values` while retaining other envelope keys (`nextPage`, `count`, `total`, etc.) unchanged. Never treat an arbitrary object that happens to have `values` as an envelope. Example `budget list --out-fields id,name` → `{ "values": [{"id": "...", "name": "..."}], ...metadata... }`.
  - For JSON keys that are not jq identifiers, quote/escape each segment when compiling (do not concatenate user input into raw jq syntax). Reject empty/invalid segments and conflicting paths such as `owner` with `owner.email` or duplicate paths, exit 2. Missing paths become `null`; null/missing parents do not crash. Preserve output ordering only where JSON guarantees it (not map key order).
- If both flags are set, `--out-fields` is applied first and `--out-jq` second. For lists use `--out-jq '.values[].id'`; after `--out-fields id,name`, `--out-jq '.values[].id'` still works. For non-paginated arrays use `.[].id` instead. Document and test this order.

**Resolving the collisions** (see the table above):
- Add `out-jq`, `out-fields`, `raw-output`, `json-style` and `no-validate` to `gencli`'s `reservedFlag`; reserve `interactive` for future body flags even though it is local to writes. `reservedFlag` currently **skips** colliding body properties, so update `extractBodyFields` to emit `--body-<name>` (checking that this fallback is also unused) and keep the original JSON key. Existing `regulatory-compliance --fields` remains unchanged with the new output flag names.
- Add a unit test that walks the full tree and fails if any leaf's local long flag or shorthand shadows **any** inherited persistent flag, including those on curated ancestors. Also test the new fallback in `cmd/gencli/main_test.go`.

### 4. Guided input from the API schema

#### 4a. Examples in `--help`

- `gencli` fills cobra's `Example` on every generated leaf with:
  1. A minimal invocation with **every required CLI/config param** (including `--org-id` where applicable) as explicit, shell-safe replacement tokens, for example `--org-id ORG_ID --id RESOURCE_UUID`. Never use unquoted `<int>`/`<uuid>` placeholders, which the shell interprets as redirection. Label the tokens as placeholders to replace; include no invented credentials or resource IDs.
  2. For body operations, an `--body '...JSON...'` invocation only if the **whole example** fits the help limit, is safe to display, and is valid shell quoting. Otherwise show `--body @request.json` plus `flexera-cli cli schema <path> --example > request.json`. Do not truncate *inside* JSON or shell quoting: a 20-line `…` snippet would look copyable but be invalid.
  3. For write operations, a `--dry-run` variant; distinguish `--yes` from dry-run, and never include `--yes` on an unreviewed example.
- Prefer real spec examples over synthesized placeholders, but label both as illustrative. Reject or flag examples that fail schema validation, require a real secret, or supply invalid identifiers. Never print real credential examples. Keep help concise even for deep schemas; `cli schema` is the full source of detail.
- `gendocs` then picks the examples up in `docs/cli/*.md` with no extra work.

#### 4b. `flexera-cli cli schema <command path...>`

```
flexera-cli cli schema budget create                 # full JSON: params, request/response schema, example
flexera-cli cli schema budget create --example       # only the JSON body (pipe into --body @-)
flexera-cli cli schema budget create --part request|response|params [--depth N]
```

- Resolve the requested path by traversing the live Cobra tree (same aliases/registration as actual execution), then look up its leaf annotation. An unknown/ambiguous path or unannotated curated command returns exit 2 with an actionable message; neither search nor schema needs auth, a network call, or an API client. In `internal/cli/root.go`, keep a single root `PersistentPreRunE` that branches for `cli search`/`cli schema` and skips normal API/config resolution on that path, while still reading relevant `-o` and `--json-style` flags (and their env values) so a missing default config file never blocks offline discovery. Do not add a child `PersistentPreRunE` that accidentally suppresses root initialization for ordinary commands; if child hooks become necessary, explicitly enable and test traversal/order. An explicitly supplied `--config` is honored (and malformed contents fail with exit 2); the implicit default config file is not read on this path. `--example` rejects combinations with `--part`, `-o table`, `--out-jq`, `--out-fields`, or `-r` (exit 2), so its stdout remains the exact request body.
- Schemas are emitted with `$ref`s resolved to `--depth` (default 3), with a visited-ref set to avoid cycles; deeper refs remain `{"$ref": ...}`. Preserve `required`, `nullable`, `additionalProperties`, `oneOf`/`allOf` and descriptive metadata. Print one JSON value to stdout, no comments or banners. `--example` prints **only** valid JSON with a trailing newline; when an example is unavailable, return exit 2 rather than outputting a non-JSON placeholder. `--body @-` consumes that single value.

#### 4c. Client-side validation

- **Order for generated JSON writes:** gather CLI flags (without failing on interactive missing values) → reject incompatible interactive stdin and perform the TTY guard → resolve `--body`/typed flags to raw JSON → parse exactly one JSON value if supplied (`-i` with no body starts from an empty object, non-interactive mode reports the existing missing-body error) → optionally prompt and merge edits (§4d) → validate the final raw JSON against the catalog schema → decode into the generated Go request type → create/render plan or apply. Validate **before** the generated-type `json.Unmarshal`: otherwise unknown keys can disappear and a successful preview can differ from the request the client actually sends. Move `deps.APIClient()` until after validation and dry-run, so invalid input and previews require no credentials or network.
- Use `github.com/getkin/kin-openapi/openapi3` (`Schema.VisitJSON` with `openapi3.MultiErrors()`) for OAS 3.0 `nullable`/composition semantics. Resolve catalog refs into `SchemaRef` objects once and cache the result; compare errors against a small set of actual generated request bodies before enabling validation by default. Do **not** validate `application/octet-stream` via JSON schema; keep existing byte handling and label it `validation: {"status":"unsupported"}` in dry runs.
- On failure: exit code 2, and a structured error:
  ```json
  {"error": "request body failed validation",
   "details": [{"path": "/name", "message": "property \"name\" is missing"},
               {"path": "/metric", "message": "value is not one of the allowed values [\"cost\",\"usage\"]"}],
   "schema": "flexera-cli cli schema budget create"}
  ```
- The dry-run plan keeps today's top-level `dryRun`, `destructive` and `plan` keys; add `"validation": {"status":"ok"}` when checked, `"skipped"` for `--no-validate`, or `"unsupported"` for raw uploads. Render with the shared JSON formatter but never apply output shaping to plans. Today dry runs include the body: redact secret-valued fields in **both** human and JSON previews (use schema `writeOnly`/password metadata, plus reviewed sensitive-name fallbacks) and set an explicit `redacted: true` marker when the JSON preview is not replayable. Document this value-level compatibility change; do not redact the in-memory body used for the actual request.
- Escape hatch: root persistent `--no-validate`, for incorrect specs; it skips **schema constraints**, not JSON syntax, generated-type decoding, terminal guards or required CLI flags. It must never claim `ok` after being skipped. New flags need an explicit config/env policy: `--no-validate`, `--out-jq`, `--out-fields`, `--raw-output` are CLI-only (not bound to Viper); only `--json-style` is config/env-backed.
- Unknown JSON properties: when the schema explicitly forbids them, validation fails. When it permits them, first check whether decoding/re-encoding the generated request type would drop them. If so, **fail** with a message identifying those keys and suggesting a fix; never warn-and-send-a-different-body. `--no-validate` does not override silent data loss. If the generated client can transmit the exact raw body, that can be considered separately, with a wire-level test. Do not reject `additionalProperties` when the schema permits them and the generated type preserves them.
- Check path/query params against the catalog's `required`, enum, format and numeric constraints **before** making an API call. A required string is not merely `TrimSpace != ""`; required numeric/bool/array params require `cmd.Flags().Changed` or their documented config source, so `0`/`false` can be legitimate values. Preserve the SDK's injected `Api-Version` headers instead of inventing CLI flags. Remove generated ad-hoc checks only after representative tests confirm equivalent behavior; skip schema validation for curated workflows without an exact matching operation.

#### 4d. Form mode with plan and apply (humans)

- **Flag:** `--interactive` / `-i` is a **local flag on generated JSON write commands only**, not a root persistent flag. Curated writes and raw uploads can opt in after they have an input/validation adapter. Reject its use on other commands as an unknown flag; keep the generated body-field name reserved so it cannot shadow local `--interactive`.
- **Terminal requirement.** Check the actual prompt input (`cmd.InOrStdin()`) **and prompt output (`cmd.ErrOrStderr()`)** for terminals before reading `@-`, requiring body/params, constructing a client, or displaying a prompt. No stdout TTY requirement: `-i` can legitimately pipe JSON success output. Fail immediately with exit 2: `interactive mode requires terminal input and stderr`; test piped stdin and redirected stderr independently. Never automatically prompt for a missing `--yes` without `-i` (including on a TTY).
- **Pre-fill and execution order.** Inspect flags for supplied values; if `--body @-` would consume the prompt input, reject `-i --body @-` (use `@file` or inline JSON instead). Load inline/file JSON and typed flags before prompting; keep the existing `--body` precedence, show supplied values editable, and only prompt for absent required values. After edits, validate and decode *once* as in §4c. Ensure the same final params/body are used in the displayed plan and in the eventual API request. For `-i --dry-run`, gather missing inputs, show the plan, then stop without an API call or `yes` question.
- **Prompt order:**
  1. Missing required path and query params.
  2. Required body properties.
  3. "Configure optional fields?", a multi-select of optional properties.
- **Widgets by schema type:**

  | Schema | Widget |
  |--------|--------|
  | `enum` | select |
  | `boolean` | confirm |
  | `string` with `format` | input + validator (uuid, date-time, email, uri) |
  | `integer` / `number` | input + numeric/min/max validator |
  | `array` of scalars | repeated input ("add another?") |
  | `object` (depth ≤ 3) | recurse with a breadcrumb title |
  | `oneOf`/`anyOf` | variant select, then recurse |
  | deeper or unsupported | open `$EDITOR` with the synthesised example, then validate |

  Do not add secret values to an editor file or display them in a terminal plan. If `$EDITOR` is unavailable or the form schema cannot be represented safely, fail with guidance to use `--body @file` rather than sending a partial body. Apply the same final validation after editor input.

- **Plan review (Terraform-style)** is rendered to stderr:
  ```
  flexera-cli will perform the following action:

    # budget create  (POST /finops-analytics/v1/orgs/{orgId}/budgets)
    org: 12345
    + name   = "Q4 cloud"
    + metric = "BilledCost"
    + yearMonths = ["2023-01"]
    ... additional request fields ...

  Save the complete body to a private file to get an equivalent command.

  Do you want to perform this action?
    Only 'yes' will be accepted to approve.

    Enter a value:
  ```
  - The display above is abbreviated to illustrate layout; real plans show **every effective input** (with secrets redacted), not `...` omissions. `+` marks a field that is added, **not** a server-side diff. Make the command reproducible: print an inline shell-quoted JSON body only when it is safe, short, and has no sensitive fields; otherwise print `--body @request.json` **only after** writing that exact body to a user-chosen file with `0600` permissions, or offer an explicit "save plan/body" action. Never print a hypothetical file that does not exist. Redact secrets in the displayed plan, debug traces, errors and shell examples; do not put secrets in shell history or reusable commands. If no safe reproducible form exists, omit the line with an explanatory note.
  - `yes` must match exactly; anything else prints `Apply cancelled.` to stderr and exits 1 without an API call.
  - `-i --yes` skips only the final prompt, like `terraform apply -auto-approve`.
  - `-i --dry-run` gathers missing inputs, shows the human plan on stderr **and** prints the existing JSON dry-run plan on stdout for automation, then stops before the confirmation prompt. This makes both outputs testable and keeps the no-network dry-run contract.
- **Diffs for update/replace (phase 4d-2):**
  - **Deferred, opt-in only** after the endpoint's replace-vs-patch semantics are documented and verified. Merely matching a `get` by resource/path params is insufficient: check permissions, response shape, defaults/read-only fields and whether omission actually deletes anything. Never label an omitted field `-` without proof that it will be removed. A preflight GET must not be required for the write or silently overwrite more recent server state; note that its snapshot may be stale and use ETags/conditional requests where supported. If fetch fails, show a no-diff plan with a note.
- **Refactor.** Replace `ConfirmWrite`'s `map[string]any` with a `Plan` struct:
  ```go
  type Plan struct {
      Command, Method, Path string
      OrgID       int
      Params      map[string]any
      Body        json.RawMessage
      Destructive bool
      Validation  *ValidationResult
      Current     json.RawMessage // optional, opt-in diff snapshot
  }
  ```
  - It has two renderers: JSON for `--dry-run`, keeping the existing top-level keys, and a human renderer for the Terraform-style view. Store all effective path/query flags including `org-id` and the final body, but redact marked secret fields only in display/export—not in the request.
  - `gencli` templates build a `Plan` instead of `writePlan`.
- **Prompt library.** `github.com/charmbracelet/huh`, behind a small `Prompter` interface so tests can script answers. Its accessible mode gives screen-reader-friendly output.

## Phasing / PRs

| # | PR | Depends on | Notes |
|---|----|-----------|-------|
| 1 | Output inventory + Printer refactor + TTY-aware JSON + `--json-style` | — | Decide structured vs. text/binary contracts for *every* direct stdout writer; migrate JSON paths in `finops`, `policy`, `workflows`, generated responses and dry runs. Keep binary/text output explicit. |
| 2 | `--out-jq`, `--out-fields`, `-r` + flag-collision resolution + collision test | 1, 3 | Project `values` only for known paginated envelopes; no current `--fields` rename. Resolve and test the existing collision inventory before enforcing a zero-collision tree. Test unsupported output commands and stream/table semantics. |
| 3 | Catalog generation (`internal/catalog`, annotations, example synthesis, registry reconciliation, size guard) | — | First complete operation-coverage and existing-flag-collision audits. Generation must be staged and leave checked-in outputs unchanged on any error. After those prerequisites, it can run in parallel with 1. Establish the two-way tree test before consumers; fail partial regen. |
| 4 | `--help` examples + `cli schema` | 3 | Preserve valid JSON/shell syntax; regenerate docs and completions. |
| 5 | `cli search` + root help hint + execution-level early-error writer + unknown-command suggestions | 3 | Golden-query tests, pre-init errors and revised smoke expectation; do not put this logic in `Printer`. |
| 6 | Raw JSON validation + `--no-validate` + `Plan` struct | 1, 3 | Validate before Go decoding and compare effective body; dry-run keeps old top-level keys. |
| 7 | Local `--interactive` form mode + plan/apply prompt | 6 | Add `huh` behind `Prompter`, TTY guard on stdin/stderr, no implicit prompt without `-i`. |
| 8 | Endpoint-verified update/replace diffs | 7 | Optional, opt-in stretch; do not infer deletions solely from omitted fields. |

## Testing

- **Printer:**
  - table tests for auto/pretty/compact with an injected `IsTTY`
  - `--out-jq` zero, single, multiple, raw results, compile and runtime errors (including no partial stdout); explicit pretty+multi and table+multi errors
  - `--out-fields` for arrays, objects, paginated `values` envelopes, nested paths, missing/null paths, special JSON keys, conflicting/duplicate paths, and large integer IDs; include paginated IDs greater than 2⁵³ and prove they are preserved or rejected before rounding
  - combined `--out-fields` + `--out-jq`; `-r` without `--out-jq` exits 2
  - test all output categories: JSON through Printer, dry-run without shaping, text/binary unchanged with shaping rejected, early errors on stderr, and redirected stderr/stdout independently.
- **gencli** ([cmd/gencli/main_test.go](../../cmd/gencli/main_test.go)):
  - example-synthesis precedence, real example validation, secret exclusion and depth limits; help examples use shell-safe replacement tokens and are complete shell commands or file instructions, never truncated JSON or unquoted angle-bracket placeholders
  - `reservedFlag` fallback emits `--body-<name>` for future collisions, while existing `regulatory-compliance --fields` remains unchanged
  - metadata and Go code derive from the same `renderOp`; catalog entry shape, response-schema selection and transitive ref closure.
- **Catalog/regeneration:** load/size/determinism tests, duplicate ID/path failures and a two-way comparison against final registered Cobra paths, including bill-connect nesting, curated ownership and non-catalog curated leaves. Assert every spec operation has an emitted, exact-curated, or explicitly excluded disposition; cover 202-only responses and response-schema extraction. Inject a generation/verification failure and assert staged regeneration leaves the current generated tree, registration and catalog byte-for-byte unchanged.
- **Search:** a golden table of natural-language queries, each with the expected command in the top 3. Examples:
  - "delete a budget"
  - "list cloud bills"
  - "who has access to org"
  - "run a policy"

  Also a benchmark for building the index.
- **Validation:** fixture bodies (valid, missing required, bad enum, wrong type, explicit/implicit additional properties, nullable, composed schemas, raw upload). Test required `false`/`0`, config-sourced org ID, unknown-field loss after Go decoding, `--no-validate`, structured details, and that invalid bodies/`--dry-run` do not construct a client or make HTTP calls.
- **Form mode:**
  - scripted `Prompter` for missing required flags, editable prefilled body, happy path, cancel, `-i --yes`, `-i --dry-run`, secret redaction and same-body-for-plan-and-send
  - stdin/stderr non-TTY guard, stdout redirected, `-i --body @-` rejected and no interactive prompts without `-i`.
- **Smoke:** extend [main_smoke_test.go](../../main_smoke_test.go) with offline `cli search`, `cli schema --example`, and `--out-jq '.values[].id'` against a stubbed doer. Update the existing unknown-command test for JSON suggestions/exit 2; add pre-initialization parse errors.
- **Regression:** tree-walk tests for local vs. inherited long/shorthand flag collisions and for catalog completeness. Include a reviewed inventory of pre-existing `org-id`, `client-id`, and `client-secret` collisions; the final clean-tree gate fails for unresolved baseline items as well as newly introduced collisions.

## Documentation

- README sections:
  - "For agents": search, then schema, then dry-run, then `--yes`; plus `--out-jq '.values[].id'`, `--out-fields`, `-r`, compact JSON and unsupported-output behavior.
  - "Interactive mode": opt-in only, terminal input/stderr requirement, no secrets in reproducible examples, pipe-able stdout.
- Regenerate `docs/cli/` and `completions/` in each PR that changes the command tree.

## Open decisions

1. **Destructive commands on a terminal without `--yes`.** Keep today's error unless the user explicitly passes `-i`; automatic prompting is a separate behavior change that requires its own review and tests. Non-terminal behavior always errors.
2. **First-run help hint.** cf shows the search hint on the *first* `--help`. That needs a state file (`~/.flexera/state.json`). The recommendation is to show the hint every time in the root help only, with no state file.
3. **Catalog size and binary growth.** Accept up to about 3 MB of embedded JSON? The alternative is gzip plus lazy decompression (a smaller binary, but a harder-to-review diff).
4. **Validation strictness.** Respect explicit `additionalProperties`; for unspecified behavior follow kin-openapi's OAS 3.0 behavior, but always fail if the generated Go type would discard a supplied field before sending. A warning alone is not safe. `--no-validate` cannot bypass silent body loss.
5. **Prompt library.** `charmbracelet/huh` (rich, pulls in the bubbletea dependency tree) vs. a minimal `bufio` line prompter (no dependencies, fewer widgets). The recommendation is `huh`.
6. **Future spec annotations.** Should `x-flexera-keywords` (for search) and `x-flexera-example` (a curated request example) be added to the unified spec pipeline, so generated help doesn't depend on synthesised examples?
