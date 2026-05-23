# CarbonStack Unsupported Sidecar Command Event Result v0

## Status

Classification: PHASE 2D SIDECAR EVENT TAXONOMY RESULT / PRE-SECRET-COMMAND

This document records stabilization of the `provider.command.unsupported` provider event candidate.

It does not implement identity creation.

It does not implement public bundle export.

It does not implement conversation creation.

It does not implement message protect/open.

It does not wire OpenMLS into CarbonStackComms user-facing commands.

It does not route MLS traffic through CarbonStackCypher.

It does not mutate trust-state storage.

It does not claim production E2EE.

## Source Context

Relevant docs:

- `docs/36-provider-event-taxonomy-v0.md`
- `docs/38-provider-trust-state-mapping-v0.md`
- `docs/41-openmls-sidecar-json-envelope-plan-v0.md`
- `docs/42-openmls-sidecar-json-envelope-result-v0.md`

Relevant Comms paths:

- `carbonstack-comms/internal/protocol/provider_events.go`
- `carbonstack-comms/internal/protocol/provider_events_test.go`
- `carbonstack-comms/internal/protocol/provider_trust.go`
- `carbonstack-comms/internal/protocol/provider_trust_test.go`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`

## What Was Stabilized

The sidecar unsupported-command envelope uses:

- `provider.command.unsupported`

This event is now represented in the Go provider event taxonomy and trust-decision mapping.

## Event Meaning

`provider.command.unsupported` means:

- the sidecar received a command that is intentionally not implemented yet
- the sidecar returned a machine-readable JSON error envelope
- the sidecar exited with unsupported/bad-usage status
- no secret material was included
- no provider state was mutated

It does not mean:

- cryptographic verification failed
- identity changed
- message tampering was detected
- provider state is corrupt
- trust state should be revoked
- the user is under attack

## Classification v0

Current classification:

- class: `lifecycle`
- severity: `warning`
- trust relevant: `false`

Rationale:

- the event is operational and developer-facing
- it describes command-surface behavior, not MLS cryptographic state
- it should be visible in logs/history
- it should not trigger user trust warnings or reverify behavior

## Trust Decision v0

Current candidate trust actions:

- `append_history`
- `debug_only`

Current blocking behavior:

- does not block send
- does not block receive
- does not block open
- does not require reverify
- does not become user-visible trust warning

## Why This Matters

Unsupported commands are expected during Phase 2D.

Stabilizing this event prevents unsupported command failures from becoming unstructured process errors.

It also prevents unsupported command failures from being mistaken for cryptographic trust/security failures.

This keeps the sidecar boundary machine-readable while remaining cautious and non-alarming.

## Current Validation

Go tests now validate:

- `provider.command.unsupported` has a provider event constant
- event descriptor class/severity/trust relevance
- trust-decision behavior
- sidecar unsupported-command envelope compares against the provider event constant instead of a loose string

## Next Recommended Work

Next recommended step:

- plan `identity-create` before implementing it.

Candidate doc:

- `docs/44-openmls-sidecar-identity-create-plan-v0.md`

The plan should define:

- command input
- output envelope
- public summary shape
- local state write behavior
- storage path expectations
- private material rules
- provider events
- checkpoint behavior
- error cases
- validation flow

Do not jump directly to message protect/open.

## Allowed Claims

Allowed:

- `provider.command.unsupported` is a stable sidecar operational event candidate.
- Unsupported sidecar commands return machine-readable JSON error envelopes.
- Go taxonomy/trust-decision tests classify unsupported-command behavior.
- Unsupported sidecar commands are not trust/security failures by default.

## Not Allowed Claims

Not allowed:

- identity creation exists.
- public bundle export exists.
- message protect/open exists.
- unsupported command behavior is a cryptographic provider failure.
- trust-state storage consumes this event.
- Comms CLI consumes sidecar envelopes.
- Cypher routes MLS payloads.
- production E2EE exists.
