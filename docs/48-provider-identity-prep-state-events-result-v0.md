# CarbonStack Identity Prep State Events Result v0

## Status

Classification: PHASE 2D SIDECAR EVENT TAXONOMY RESULT / PRE-SECRET-GENERATION

This document records stabilization of identity prep state events used by the OpenMLS sidecar state-skeleton rung.

It does not implement real identity creation.

It does not generate OpenMLS signing keys.

It does not generate OpenMLS credentials.

It does not generate KeyPackages.

It does not write OpenMLS provider storage.

It does not export public bundles.

It does not create conversations.

It does not protect/open messages.

It does not wire OpenMLS into CarbonStackComms user-facing commands.

It does not route MLS traffic through CarbonStackCypher.

It does not mutate trust-state storage.

It does not claim production E2EE.

## Source Context

Relevant docs:

- `docs/42-openmls-sidecar-json-envelope-result-v0.md`
- `docs/43-provider-command-unsupported-event-result-v0.md`
- `docs/44-openmls-sidecar-identity-create-plan-v0.md`
- `docs/45-openmls-sidecar-identity-create-prep-result-v0.md`
- `docs/46-provider-command-invalid-not-implemented-events-result-v0.md`
- `docs/47-openmls-sidecar-identity-prep-state-result-v0.md`

Relevant Comms paths:

- `carbonstack-comms/internal/protocol/provider_events.go`
- `carbonstack-comms/internal/protocol/provider_events_test.go`
- `carbonstack-comms/internal/protocol/provider_trust.go`
- `carbonstack-comms/internal/protocol/provider_trust_test.go`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/main.rs`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/state.rs`

## Events Stabilized

This rung stabilizes the event vocabulary around the dev-only identity prep state skeleton:

- `provider.identity.prep_state_written`
- `provider.identity.exists`
- `checkpoint.failed`

These events are now represented in Go provider-event taxonomy and trust-decision mapping.

## `provider.identity.prep_state_written`

Meaning:

- `identity-create --device-label <safe-label>` validated the label.
- The sidecar created an ignored dev-only per-device state directory.
- The sidecar wrote a non-secret `identity-prep.json` manifest.
- No OpenMLS identity material was generated.
- No provider storage was written.
- No private material was printed.

Current classification:

- class: `storage/checkpoint`
- severity: `info`
- trust relevant: `false`

Current trust-decision behavior:

- append history
- debug only
- do not block send
- do not block receive
- do not block open
- do not require reverify
- do not show as a user-visible trust/security warning

## `provider.identity.exists`

Meaning:

- `identity-create --device-label <safe-label>` found existing prep/identity state.
- The sidecar refused overwrite.
- Existing state was not replaced.
- No private material was printed.

Current classification:

- class: `storage/checkpoint`
- severity: `warning`
- trust relevant: `false`

Current trust-decision behavior:

- append history
- debug only
- do not block send
- do not block receive
- do not block open
- do not require reverify
- do not show as a user-visible trust/security warning

Rationale:

- Duplicate identity/prep state is an operational guardrail.
- It is not automatically evidence of cryptographic compromise.
- It should be preserved in developer/audit history without alarming users.

## `checkpoint.failed`

Meaning:

- The sidecar attempted to write required state/checkpoint material and failed.
- For the current identity prep state skeleton, this means the non-secret prep manifest write failed.
- Future state-mutating commands may reuse this event for failed checkpoint/state-write operations.

Current classification:

- class: `storage/checkpoint`
- severity: `warning`
- trust relevant: `false`

Current trust-decision behavior:

- stop operation
- show recovery path
- block current send/outgoing state mutation
- history relevant
- do not require reverify by default
- do not mark identity as compromised by default

Receive/open behavior is intentionally not newly decided by this rung.

Rationale:

- Failed state persistence is operationally serious and should stop the current state-mutating operation.
- It is not automatically a cryptographic trust failure.
- It should not imply identity compromise without additional evidence.

## Current Validation

Go tests now validate:

- `ProviderEventIdentityPrepStateWritten` exists.
- `ProviderEventIdentityExists` exists.
- `ProviderEventCheckpointFailed` exists.
- `provider.identity.prep_state_written` is classified as storage/checkpoint, info, non-trust-relevant.
- `provider.identity.exists` is classified as storage/checkpoint, warning, non-trust-relevant.
- `checkpoint.failed` remains classified as storage/checkpoint, warning, non-trust-relevant.
- prep-state-written maps to append-history/debug-only behavior.
- identity-exists maps to append-history/debug-only behavior.
- checkpoint-failed maps to stop-operation/show-recovery-path behavior.
- checkpoint-failed blocks current outgoing/send mutation.
- checkpoint-failed is history relevant.
- checkpoint-failed does not require reverify by default.
- sidecar Go tests compare prep-state and duplicate-state events against typed constants.

## Implementation Notes

During this rung, `checkpoint.failed` was found to already exist in the older storage/checkpoint event group.

The final approach preserves that existing storage/checkpoint handling rather than duplicating `checkpoint.failed` as a separate switch case.

The trust test also avoids forcing a new receive/open blocking policy for `checkpoint.failed`. This checkpoint stabilizes the current-operation/send-mutation behavior without prematurely deciding all future receive/open behavior.

## Why This Matters

The sidecar has now crossed into controlled local state writes.

That makes event semantics more important than they were during command-surface-only work.

This rung ensures that identity prep state write, duplicate state refusal, and checkpoint failure are typed and policy-mapped before the project attempts real OpenMLS identity material generation.

The project now has a safer ladder:

- command exists
- command validates labels
- command writes non-secret ignored prep state
- state events are typed and trust-mapped
- only then attempt real dev-only identity generation

## Security Boundary

This result is not real identity creation.

It does not:

- create OpenMLS signing keys
- create OpenMLS credentials
- create KeyPackages
- write signer JSON
- write MemoryStorage JSON
- write OpenMLS provider storage
- export public bundles
- create conversations
- protect/open messages
- mutate `trust.json`
- mutate `trust-events.jsonl`
- affect Comms CLI behavior
- affect Cypher routing

## Next Recommended Work

Recommended next step:

- implement minimal dev-only real `identity-create` material generation.

The first real implementation should:

- write only under ignored sidecar dev-state
- refuse overwrite by default
- generate the smallest viable OpenMLS identity/signing material needed for later public bundle work
- return sanitized JSON only
- keep `private_material_included: false`
- include explicit `identity_created`, `state_written`, and `provider_storage_written` fields
- avoid printing signer JSON, MemoryStorage JSON, provider storage JSON, private keys, seeds, or recovery material
- keep Comms CLI, Cypher routing, and trust-store mutation out of scope

Do not jump directly to public bundle export, conversation lifecycle, message protect/open, or Comms runtime integration.

## Allowed Claims

Allowed:

- `provider.identity.prep_state_written` is a typed provider event.
- `provider.identity.exists` is a typed provider event.
- `checkpoint.failed` is confirmed/stabilized for this state-skeleton path.
- Identity prep state events are classified and trust-mapped.
- The sidecar can write non-secret dev-only prep state and report that through typed events.
- Duplicate prep state refusal is represented as typed non-trust-relevant behavior.

## Not Allowed Claims

Not allowed:

- identity creation is implemented.
- OpenMLS signer/credential generation exists in the sidecar.
- KeyPackage generation exists in the sidecar.
- public bundle export exists.
- conversation creation exists.
- message protect/open exists.
- provider state persistence is solved.
- production storage exists.
- Comms CLI consumes sidecar identity state.
- Cypher routes MLS payloads.
- trust-state storage consumes these provider events.
- production E2EE exists.
