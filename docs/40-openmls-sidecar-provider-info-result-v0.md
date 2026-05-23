# CarbonStack OpenMLS Sidecar Provider-Info Result v0

## Status

Classification: PHASE 2D SIDECAR BOOTSTRAP RESULT / PRE-INTEGRATION

This document records the first successful Go-tested OpenMLS sidecar command boundary.

It does not wire OpenMLS into CarbonStackComms user-facing commands.

It does not route MLS traffic through CarbonStackCypher.

It does not mutate trust-state storage.

It does not implement secret-bearing provider commands.

It does not claim production E2EE.

## Source Context

Relevant planning doc:

- `docs/39-phase2d-sidecar-command-surface-plan.md`

Relevant sidecar path:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar`

Relevant Go-side test:

- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`

Relevant Comms commit:

- `carbonstack-comms` `test: parse OpenMLS sidecar provider-info output`

## What Was Added

Phase 2D now has an experimental Rust sidecar crate:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar`

The sidecar supports one command:

- `provider-info`

The command prints JSON describing the experimental provider boundary.

## Provider-Info Output Shape

Current output includes:

- `provider`
- `implementation`
- `mode`
- `phase`
- `capabilities`
- `unsupported`
- `security_level`
- `private_material_included`
- `warnings`

Required current values:

- `provider`: `openmls`
- `implementation`: `carbonstack-openmls-sidecar`
- `mode`: `experimental-sidecar`
- `phase`: `phase2d-provider-info`
- `private_material_included`: `false`

Current supported capability:

- `provider-info`

Current intentionally unsupported commands:

- `identity-create`
- `public-bundle-export`
- `conversation-create`
- `conversation-add-member`
- `conversation-join`
- `message-protect`
- `message-open`
- `state-checkpoint`
- `state-load-check`

## What Was Validated

The Go-side test invokes the Rust sidecar command:

- `cargo run --quiet -- provider-info`

The test parses the JSON output and validates:

- provider is `openmls`
- implementation is `carbonstack-openmls-sidecar`
- mode is `experimental-sidecar`
- phase is `phase2d-provider-info`
- private material is not included
- `provider-info` appears in capabilities
- secret-bearing/state-mutating commands appear in unsupported
- warnings are present

## Why This Matters

This is the first Phase 2D runtime-boundary foothold.

It proves Go can invoke the Rust sidecar and parse a stable JSON response without wiring provider behavior into Comms CLI flows.

This keeps runtime integration incremental.

The sidecar boundary now exists as something testable, not just a doc concept.

## What This Does Not Prove

This does not prove:

- OpenMLS is wired into `comms send`
- OpenMLS is wired into `comms inbox`
- Cypher routes MLS payloads
- trust-state consumes provider events
- identity creation works
- public bundle export works
- conversation creation works
- message protect/open works through the sidecar
- state checkpointing works through the sidecar
- production storage is solved
- production E2EE exists

## Security Boundary

`provider-info` is intentionally non-secret.

It must not:

- create identity material
- export signing keys
- export provider storage
- mutate provider state
- read real user secrets
- write recovery material
- affect Comms CLI behavior
- affect Cypher routing
- mutate `trust.json` or `trust-events.jsonl`

## Current Design Decision

Sidecar-first remains the preferred early runtime-provider path.

FFI remains deferred.

Direct Go-provider implementation remains deferred.

The sidecar path is preferred for now because it:

- keeps Rust/OpenMLS isolated
- gives Go a stable command boundary to test
- avoids premature FFI coupling
- lets provider commands evolve independently before CLI integration
- keeps failure surfaces inspectable

## Next Recommended Work

Next work should still be small.

Recommended next step:

- define the sidecar JSON envelope/error-envelope shape before adding secret-bearing commands.

Candidate next doc:

- `docs/41-openmls-sidecar-json-envelope-plan-v0.md`

Candidate next code after that:

- teach `provider-info` to use the envelope shape
- or add a test-only `unsupported-command` error-envelope check

Do not jump directly to `identity-create`.

## Allowed Claims

Allowed:

- CarbonStack has a Phase 2D experimental OpenMLS sidecar crate.
- The sidecar currently supports `provider-info`.
- Go-side tests can invoke and parse the sidecar's `provider-info` JSON.
- The sidecar command boundary is now minimally test-protected.

## Not Allowed Claims

Not allowed:

- OpenMLS is integrated into CarbonStackComms messaging.
- Cypher routes MLS payloads.
- trust-state consumes sidecar/provider events.
- sidecar identity or messaging commands exist.
- sidecar output is production security behavior.
- production E2EE exists.
