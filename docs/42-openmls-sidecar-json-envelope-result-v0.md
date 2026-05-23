# CarbonStack OpenMLS Sidecar JSON Envelope Result v0

## Status

Classification: PHASE 2D SIDECAR ENVELOPE RESULT / PRE-SECRET-COMMAND

This document records the first implemented JSON success/error envelope for the experimental OpenMLS sidecar.

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

- `docs/39-phase2d-sidecar-command-surface-plan.md`
- `docs/40-openmls-sidecar-provider-info-result-v0.md`
- `docs/41-openmls-sidecar-json-envelope-plan-v0.md`

Relevant sidecar path:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar`

Relevant Go-side test:

- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`

Relevant Comms commit:

- `carbonstack-comms` `test: envelope OpenMLS sidecar provider-info output`

## What Was Added

The OpenMLS sidecar now emits a common JSON envelope for:

- successful `provider-info`
- unsupported command failure

This replaces the earlier raw provider-info JSON object with an explicit envelope shape.

## Success Envelope Validated

Command:

- `cargo run -- provider-info`

Current success output includes:

- `ok: true`
- `command: provider-info`
- `provider: openmls`
- `implementation: carbonstack-openmls-sidecar`
- `mode: experimental-sidecar`
- `phase: phase2d-provider-info`
- `data.capabilities`
- `data.unsupported`
- `data.security_level`
- `events`
- `warnings`
- `private_material_included: false`

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

## Unsupported Command Envelope Validated

Command:

- `cargo run -- identity-create`

Current behavior:

- prints JSON error envelope
- exits with code `2`

Current error envelope includes:

- `ok: false`
- `command: identity-create`
- `provider: openmls`
- `implementation: carbonstack-openmls-sidecar`
- `mode: experimental-sidecar`
- `phase: phase2d-provider-info`
- `error.code: unsupported_command`
- `error.message: unsupported command: identity-create`
- `error.provider_event: provider.command.unsupported`
- `error.severity: warning`
- `error.trust_relevant: false`
- `events[0].event: provider.command.unsupported`
- `private_material_included: false`

The nonzero exit code is expected and desirable.

The important behavior is that the sidecar still prints machine-readable JSON before exiting.

## Go-Side Validation

The Go-side provider-info test now validates both:

- success envelope for `provider-info`
- unsupported-command error envelope for `identity-create`

The test validates:

- provider identity fields
- sidecar mode/phase fields
- capability list
- unsupported command list
- private-material flag
- warning presence
- unsupported command exit code `2`
- unsupported command error code/event/severity/trust relevance

## Why This Matters

This establishes the first stable sidecar command response contract.

Future sidecar commands should not invent their own unrelated output shape.

The Go side now has a predictable parse target for both success and failure.

This also prevents unsupported commands from collapsing into plain stderr text or unstructured process failures.

## Security Boundary

The envelope work remains non-secret and non-state-mutating.

It must not:

- create identity material
- export public bundles
- create or join conversations
- protect or open messages
- checkpoint provider state
- print private keys
- print signer material
- print provider storage
- mutate `trust.json`
- mutate `trust-events.jsonl`
- affect Comms CLI behavior
- affect Cypher routing

## Provider Event Note

The result introduces or preserves a new operational provider event candidate:

- `provider.command.unsupported`

Current meaning:

- sidecar received a recognized unsupported/future command request
- this is not a cryptographic trust failure
- severity is warning
- trust relevance is false

Future provider-event taxonomy docs/code may need to add this event explicitly if it becomes part of the stable sidecar event vocabulary.

## Next Recommended Work

Recommended next step:

- add or update Go provider-event taxonomy to include `provider.command.unsupported`, if the event is intended to be stable.

Then proceed to a careful pre-implementation plan for the first secret-adjacent command.

Candidate next docs/code:

- document `provider.command.unsupported` in the provider event taxonomy
- add Go constant/classification for `provider.command.unsupported`
- add a tiny test mapping it to warning/non-trust-relevant behavior
- only after that, plan `identity-create`

Do not jump directly to message protect/open.

## Allowed Claims

Allowed:

- CarbonStack has an implemented JSON success/error envelope for the experimental OpenMLS sidecar.
- `provider-info` returns an envelope with `ok: true`.
- unsupported commands return an envelope with `ok: false` and exit code `2`.
- Go-side tests parse both success and unsupported-command envelope paths.

## Not Allowed Claims

Not allowed:

- identity creation exists.
- public bundle export exists.
- conversation creation exists.
- message protect/open exists.
- Comms CLI consumes sidecar envelopes.
- Cypher routes MLS payloads.
- trust-state consumes sidecar events.
- production E2EE exists.
