# Output-boundary inventory

| Category | Current contract / implementation |
| --- | --- |
| Generated structured API responses | Shared `Printer.Render`; JSON terminal policy or SDK table renderer. |
| FinOps JSON response bytes | Decode using `UseNumber`, render through printer. |
| FinOps non-JSON bytes | Confirmed JSON endpoints reject invalid JSON rather than silently treating it as a structured result. Empty success bodies stay empty. |
| Policy applied-policy log | Preserve text/Markdown bytes, including JSON-looking text; shaping rejected before config/body/client work. Existing table rejection retained. |
| Workflow results | Render using `deps.Config.Output`, not hard-coded JSON. |
| `curated list` | Preserve human text; shaping rejected before execution. |
| Generated 204/no-JSON success | Preserve `OK`; shaping rejected before execution, including with dry-run. |
| Generated dry-run | Shared JSON formatter, never shaping; configured style passed by new generator template. |
| FinOps dry-run | Shared JSON formatter using configured printer style, never shaping. |
| Auth-token-create dry-run | Redacted JSON; no HTTP; form-schema validation marked unsupported. |
| Runtime errors | Shared JSON formatter with configured printer style; never shaping. |
| Early Cobra errors | Typed execution-boundary JSON errors; never response shaping. |
| Help and completion | Preserve ordinary text; explicit shaping options return usage exit 2 without producing help/script bytes. |
| Binary/raw output contract | Explicit text/binary/mixed annotations reject shaping before execution. Non-JSON schemas use a byte-preserving response branch. |

Generated leaves carry output annotations derived from all selected success
responses; a mixed structured/no-JSON operation rejects shaping conservatively.
`--json-style` alone does not alter text or raw bytes. Structured JSON, errors
and safety plans retain their distinct printer boundaries.