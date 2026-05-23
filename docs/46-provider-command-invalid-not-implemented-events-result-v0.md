# CarbonStack Identity-Create Prep Command Events Result v0

## Status

Classification: PHASE 2D SIDECAR EVENT TAXONOMY RESULT / PRE-SECRET-GENERATION

This document records stabilization of two identity-create prep command events:

- `provider.command.invalid`
- `provider.command.not_implemented`

It does not implement identity creation.

It does not generate OpenMLS identity material.

It does not write provider storage.

It does not export public bundles.

It does not create conversations.

It does not protect/open messages.

It does not wire OpenMLS into CarbonStackComms user-facing commands.

It does not route MLS traffic through CarbonStackCypher.

It does not mutate trust-state storage.

It does not claim production E2EE.

## Source Context

Relevant docs:

- `docs/41-openmls-sidecar-json-envelope-plan-v0.md`
- `docs/42-openmls-sidecar-json-envelope-result-v0.md`
- `docs/43-provider-command-unsupported-event-result-v0.md`
- `docs/44-openmls-sidecar-identity-create-plan-v0.md`
- `docs/45-openmls-sidecar-identity-create-prep-result-v0.md`

Relevant Comms paths:

- `carbonstack-comms/internal/protocol/provider_events.go`
- `carbonstack-comms/internal/protocol/provider_events_test.go`
- `carbonstack-comms/internal/protocol/provider_trust.go`
- `carbonstack-comms/internal/protocol/provider_trust_test.go`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/main.rs`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/labels.rs`

## Events Stabilized

### `provider.command.invalid`

Meaning:

- the sidecar recognized the command
- the command syntax or arguments were invalid
- example: `identity-create` without `--device-label`
- example: `identity-create --device-label "../bad"`

Current classification:

- class: `lifecycle`
- severity: `warning`
- trust relevant: `false`

Current trust-decision behavior:

- append history
- debug only
- do not block send
- do not block receive
- do not block open
- do not require reverify
- do not show as user-visible trust/security warning

### `provider.command.not_implemented`

Meaning:

- the sidecar recognized the command
- arguments passed validation
- the command is intentionally not implemented yet
- example: `identity-create --device-label carbonstack-alice-device`

Current classification:

- class: `lifecycle`
- severity: `warning`
- trust relevant: `false`

Current trust-decision behavior:

- append history
- debug only
- do not block send
- do not block receive
- do not block open
- do not require reverify
- do not show as user-visible trust/security warning

## Why This Matters

`identity-create` is now a recognized sidecar command surface, but it does not generate secrets yet.

These events let the sidecar report pre-secret command behavior without turning operational command-state into cryptographic trust/security failure.

This preserves a clean distinction between:

- invalid command invocation
- recognized-but-not-implemented command
- unsupported/future command
- actual provider cryptographic errors

## Current Validation

Go tests now validate:

- provider-event descriptors for both events
- provider-trust decisions for both events
- sidecar identity-create missing-label envelope uses `provider.command.invalid`
- sidecar identity-create invalid-label envelope uses `provider.command.invalid`
- sidecar identity-create safe-label not-implemented envelope uses `provider.command.not_implemented`
- these events are warning/non-trust-relevant
- these events do not block send/receive/open or require reverify

## Boundary

These events are not:

- invalid signature
- identity changed
- tamper detected
- provider state corrupt
- message replay
- epoch stale
- user trust failure

They are operational sidecar command-surface events.

## Next Recommended Work

Recommended next step:

- implement dev-only `identity-create` state creation carefully.

Implementation should:

- create real OpenMLS identity material only if the code path can keep stdout sanitized
- write only under ignored sidecar dev state
- refuse overwrite by default
- emit machine-readable JSON envelopes
- include `private_material_included: false`
- include explicit `state_written` and `identity_created` fields
- keep Comms CLI, Cypher routing, and trust-store mutation out of scope

Candidate next result:

- `docs/47-openmls-sidecar-identity-create-result-v0.md`

Do not jump directly to public bundle export, conversation lifecycle, or message protect/open.

## Allowed Claims

Allowed:

- `provider.command.invalid` is a typed sidecar operational event.
- `provider.command.not_implemented` is a typed sidecar operational event.
- Both events classify as lifecycle/warning/non-trust-relevant.
- Both events map to non-blocking debug/history trust-decision behavior.
- Identity-create prep error envelopes are now tied to typed provider events.

## Not Allowed Claims

Not allowed:

- identity creation is implemented.
- OpenMLS signer/credential generation exists in the sidecar.
- public bundle export exists.
- conversation creation exists.
- message protect/open exists.
- provider state persistence is solved.
- production storage exists.
- Comms CLI consumes sidecar identity state.
- Cypher routes MLS payloads.
- production E2EE exists.
