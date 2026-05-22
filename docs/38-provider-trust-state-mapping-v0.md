# CarbonStack Provider Trust-State Mapping v0

## Status

Classification: PHASE 2C PROVIDER-TRUST MAPPING / PRE-INTEGRATION

This document maps provider events and provider errors into candidate CarbonStack trust-state behavior.

It does not implement trust-state integration.

It does not wire OpenMLS into CarbonStackComms.

It does not route MLS traffic through CarbonStackCypher.

It does not claim production E2EE.

## Source Context

Relevant docs:

- `docs/20-identity-and-trust-state-v0.md`
- `docs/23-phase2a-trust-state-plan.md`
- `docs/30-openmls-provider-boundary-implications.md`
- `docs/35-openmls-provider-fixture-result-v0.md`
- `docs/36-provider-event-taxonomy-v0.md`
- `docs/37-provider-negative-fixture-result-v0.md`

Relevant Comms paths:

- `carbonstack-comms/internal/trust/trust.go`
- `carbonstack-comms/internal/trust/trust_test.go`
- `carbonstack-comms/internal/protocol/provider_events.go`
- `carbonstack-comms/internal/protocol/provider_events_test.go`
- `carbonstack-comms/internal/protocol/openmls_fixture_test.go`
- `carbonstack-comms/internal/protocol/openmls_negative_fixture_test.go`

## Purpose

Provider events describe protocol-provider behavior.

Trust state describes what CarbonStack should believe, warn about, block, or ask the user to verify.

This document defines the first mapping layer between them.

The goal is to avoid treating cryptographic provider failures as generic errors.

## Current Principle

Provider events should not directly mutate trust state until policy is explicit.

The safe pipeline is:

1. provider emits or implies event
2. CarbonStack classifies event
3. CarbonStack maps event to candidate trust action
4. trust-state code records or enforces policy
5. user-facing behavior is added only after tests/docs agree

Current work stops at step 3.

## Trust Action Vocabulary v0

Candidate trust actions:

- `none`
- `append_history`
- `debug_only`
- `warn_user`
- `block_send`
- `block_receive`
- `block_open`
- `quarantine_message`
- `require_reverify`
- `mark_identity_changed`
- `show_recovery_path`
- `stop_operation`
- `fatal_local_state`

These are policy candidates.

They are not yet implemented as final CLI behavior.

## Mapping Table v0

| Provider event | Class | Severity | Trust relevant | Candidate trust actions |
|---|---|---|---|---|
| `provider.fixture.started` | lifecycle | debug | no | `debug_only` |
| `provider.fixture.completed` | lifecycle | debug | no | `debug_only` |
| `provider.public_bundle.created` | public_setup | info | no | `append_history` |
| `conversation.created` | membership | notice | no | `append_history` |
| `conversation.welcome.created` | membership | notice | no | `append_history` |
| `conversation.welcome.staged` | membership | notice | no | `append_history` |
| `conversation.member_added` | membership | notice | yes-ish | `append_history`, `warn_user` when unexpected |
| `conversation.joined` | membership | notice | no | `append_history` |
| `conversation.loaded` | membership/storage-adjacent | notice | no | `append_history` |
| `message.protected` | message | info | no | `append_history`, `checkpoint_after_send` candidate |
| `message.opened` | message | info | no | `append_history`, `checkpoint_after_receive` candidate |
| `storage.missing` | storage_checkpoint | warning | no | `stop_operation`, `show_recovery_path` |
| `storage.corrupt` | storage_checkpoint | warning | maybe | `stop_operation`, `show_recovery_path`, `warn_user` |
| `checkpoint.failed` | storage_checkpoint | warning | maybe | `stop_operation`, `warn_user`, `show_recovery_path` |
| `provider.signature.invalid` | trust_security | security | yes | `block_open`, `warn_user`, `append_history`, `require_reverify` |
| `provider.identity.changed` | trust_security | security | yes | `mark_identity_changed`, `block_send`, `block_receive`, `require_reverify`, `warn_user` |
| `provider.identity.reverify.required` | trust_security | security | yes | `require_reverify`, `block_send`, `warn_user` |
| `provider.message.tamper.detected` | trust_security | security | yes | `block_open`, `quarantine_message`, `warn_user`, `append_history` |
| `provider.replay.detected` | trust_security | security | yes | `block_open`, `quarantine_message`, `warn_user`, `append_history` |
| `provider.epoch.stale` | trust_security | security | yes | `block_open`, `warn_user`, `append_history` |
| `provider.group.unrecoverable` | terminal_fatal | fatal | yes | `fatal_local_state`, `stop_operation`, `show_recovery_path` |
| `provider.secret.material.unavailable` | terminal_fatal | fatal | yes | `fatal_local_state`, `block_send`, `show_recovery_path` |
| `provider.invariant.violation` | terminal_fatal | fatal | yes | `fatal_local_state`, `stop_operation`, `warn_user` |
| unknown provider event | unknown | warning | no by default | `append_history`, `debug_only` |

## Current Positive Fixture Mappings

The current fixture stream includes:

- `provider.fixture.started`
- `provider.public_bundle.created`
- `conversation.created`
- `conversation.welcome.created`
- `conversation.member_added`
- `conversation.welcome.staged`
- `conversation.joined`
- `message.protected`
- `message.opened`
- `conversation.loaded`
- `provider.fixture.completed`

These should mostly map to history/debug behavior.

They should not trigger scary user-facing warnings in normal happy-path flow.

Exception:

- `conversation.member_added` can become trust-relevant if unexpected, policy-disallowed, or inconsistent with local intent.

## Current Negative Fixture Mappings

### Invalid signature

Source:

- `invalid-signature-error.json`

Provider event candidate:

- `provider.signature.invalid`

Meaning:

- A message failed signature validation.
- In the observed scratch failure, this happened when Alice used a fresh signer after reload.
- Bob correctly rejected the message.

Candidate trust behavior:

- block opening the message
- warn user
- append trust history
- require reverify if identity continuity is broken

This should not be treated as a generic message failure.

### Missing storage

Source:

- `missing-storage-error.json`

Provider event candidate:

- `storage.missing`

Meaning:

- Required local provider state is not available.

Candidate trust behavior:

- stop operation
- do not send
- show recovery path
- do not mutate provider state further

This is operational first, not necessarily remote-trust failure.

### Missing signer

Source:

- `missing-signer-error.json`

Provider event candidate:

- `provider.secret.material.unavailable`

Meaning:

- Required signing material is unavailable.

Candidate trust behavior:

- block send
- fatal local state
- preserve local evidence
- show recovery path
- require explicit identity recovery flow

This is security-relevant because signing identity continuity is required.

### Wrong group ID

Source:

- `wrong-group-error.json`

Provider event candidate:

- `provider.group.unrecoverable`

Meaning:

- Provider state cannot safely load or match the expected conversation/group.

Candidate trust behavior:

- stop operation
- do not send
- do not open message
- surface provider state mismatch
- show recovery path

### Malformed provider message

Source:

- `malformed-message-error.json`

Provider event candidate:

- `provider.message.tamper.detected`

Meaning:

- Provider material cannot be safely parsed or validated.
- Could be corruption, transmission failure, or malicious tampering.

Candidate trust behavior:

- block message
- quarantine message
- append trust event
- warn user
- retain sanitized diagnostic summary

## Mapping Boundaries

### Provider event classification is not enforcement

`DescribeProviderEvent` classifies events.

It does not enforce policy.

### Trust action mapping is not final UX

This document proposes candidate behavior.

It does not define final UI copy, final CLI output, or final recovery UX.

### Operational errors are not always trust failures

Missing storage may be a local recovery issue.

Invalid signatures and tamper candidates are more directly trust/security relevant.

### Unknown events should be visible but cautious

Unknown future provider events should be recorded and warned about in debug/developer surfaces.

They should not automatically become trust-relevant without review.

## Future Implementation Shape

A future pure mapping layer could live in `carbonstack-comms/internal/protocol` or `carbonstack-comms/internal/trust`.

Suggested shape:

- input: `ProviderEventDescriptor`
- output: `ProviderTrustDecision`

Example fields:

- `event`
- `class`
- `severity`
- `actions`
- `trust_relevant`
- `blocks_send`
- `blocks_open`
- `requires_reverify`
- `user_visible`
- `history_message`

Keep this pure and testable before wiring it into CLI flows.

## Implementation Guardrails

Do not immediately mutate `trust.json` from provider events.

First implement:

- pure mapping tests
- fixture-backed tests
- docs-aligned behavior
- no runtime Rust dependency
- no Cypher routing dependency

Then consider trust-state integration.

## Recommended Next Work

Recommended next step:

- add a tiny pure Go mapping test layer from provider events to candidate trust actions.

Suggested files:

- `carbonstack-comms/internal/protocol/provider_trust.go`
- `carbonstack-comms/internal/protocol/provider_trust_test.go`

Keep it pure.

Do not wire to CLI or trust storage yet.

## Allowed Claims

Allowed:

- CarbonStack has a provider-event-to-trust-action mapping draft.
- The mapping draft is based on existing fixture and negative-fixture work.
- The mapping is pre-integration and policy-shaping.

## Not Allowed Claims

Not allowed:

- trust-state consumes provider events.
- CLI blocks or warns based on provider events.
- OpenMLS is integrated into Comms.
- Cypher carries MLS payloads.
- production E2EE exists.
- hostile-server security is solved.
