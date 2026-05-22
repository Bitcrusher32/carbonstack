# CarbonStack OpenMLS Provider Fixture Result v0

## Status

Classification: PHASE 2C PROVIDER-CONTRACT FIXTURE RESULT / PRE-INTEGRATION

This document records the first successful sanitized OpenMLS provider fixture generation.

It does not wire OpenMLS into CarbonStackComms.

It does not route MLS traffic through CarbonStackCypher.

It does not claim production E2EE.

## Source Context

Relevant plan:

- `docs/34-openmls-provider-fixture-contract-plan.md`

Relevant scratch path:

- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal`

Fixture command:

- `cargo run -- fixtures`

Fixture output path:

- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/fixtures/dev`

## What Was Validated

The Rust-only OpenMLS scratch crate now has a fixture mode.

The fixture mode runs a local provider-contract sample flow and writes sanitized summaries.

It validated:

- Alice and Bob setup material summary generation.
- Public setup summary generation.
- Welcome/join summary generation.
- Conversation checkpoint summary generation.
- Message summary generation.
- Provider event JSONL generation.
- Invalid-signature error fixture generation.
- No provider storage files are written into the fixture directory.
- No signer JSON is written into the fixture directory.
- No private key material is intentionally included.

## Generated Fixture Files

Current fixture set:

- `provider-summary.json`
- `alice-device-summary.json`
- `bob-device-summary.json`
- `alice-public-setup-summary.json`
- `bob-public-setup-summary.json`
- `conversation-before-message.json`
- `welcome-summary.json`
- `conversation-after-message-one.json`
- `conversation-after-reload.json`
- `message-one-summary.json`
- `message-two-summary.json`
- `conversation-after-message-two.json`
- `invalid-signature-error.json`
- `provider-events.jsonl`

## Provider Event Vocabulary

Current JSONL event examples include:

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

## Invalid Signature Fixture

The fixture set includes:

- `invalid-signature-error.json`

This preserves the known failure where phase-B originally used a fresh Alice signer after reload and Bob rejected the message with:

- `ValidationError(InvalidSignature)`

Candidate CarbonStack mapping:

- `provider.signature.invalid`

Candidate trust actions:

- block message
- warn user
- append trust event
- require reverify if identity changed

## Why This Matters

This is the first bridge from OpenMLS scratch behavior to CarbonStack provider-contract shape.

It gives the Go-side provider boundary something concrete to learn from without linking Rust into the CLI yet.

This reduces integration risk.

## What This Still Does Not Prove

This does not prove:

- OpenMLS is integrated into CarbonStackComms.
- Cypher routes MLS payloads.
- the fixture mode is secure messaging.
- fixture summaries are production provider events.
- production storage is solved.
- secure vault storage is solved.
- hostile-server security is solved.
- replay resistance is solved.
- metadata privacy is solved.

## Next Work

Recommended next work:

1. Add Go-side provider-contract fixture review/tests.
2. Define concrete provider event structs or docs.
3. Map fixture event names to CarbonStack trust-state candidates.
4. Define which fixture fields belong in Cypher envelopes later.
5. Continue negative-path fixture work before sidecar integration.

## Allowed Claims

Allowed:

- CarbonStack has a dev-only OpenMLS provider fixture mode.
- Fixture mode emits sanitized provider-contract summaries.
- Fixture mode gives Go-side provider boundary work concrete shapes to consume.

## Not Allowed Claims

Not allowed:

- OpenMLS is integrated into CarbonStackComms messaging.
- fixture mode is production encryption.
- fixture files contain real production security material.
- trust-state integration is complete.
- Cypher carries MLS traffic.
