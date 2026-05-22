# CarbonStack OpenMLS Provider Fixture Contract Plan

## Status

Classification: PHASE 2C PROVIDER-CONTRACT PLAN / PRE-INTEGRATION

This document defines the next safe step after the OpenMLS MemoryStorage persistence result.

It does not wire OpenMLS into CarbonStackComms.

It does not route MLS traffic through CarbonStackCypher.

It does not claim production E2EE.

## Context

CarbonStack has validated a Rust-only OpenMLS scratch ladder:

- credential and KeyPackage setup
- Alice group creation
- Bob add from KeyPackage
- Bob join from Welcome
- application message protect/open
- two-message state continuity
- same-process provider-storage reload
- MemoryStorage file save/load
- post-reload message protect/open with reloaded Alice signer

The next step is not direct integration.

The next step is a provider fixture contract.

## Purpose

The purpose is to define what a real provider must produce and consume before Go/Rust integration begins.

A fixture contract lets CarbonStack learn the provider boundary shape without coupling the Go CLI to Rust OpenMLS too early.

## Why Fixtures First

Fixtures are lower-risk than direct integration.

They allow CarbonStack to validate:

- provider data shapes
- event vocabulary
- persistence checkpoint expectations
- error cases
- trust-state mapping
- Cypher opaque-envelope assumptions

without requiring:

- sidecar design
- FFI
- Go/Rust build integration
- production storage
- real user secrets
- CLI send/inbox rewiring

## Fixture Output Directory

Proposed scratch output path:

- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/fixtures/dev/`

These fixtures must be treated as dev-only.

If committed later, they must contain no real secrets and must be clearly labeled as generated scratch fixtures.

If any fixture contains signing keys, provider storage, or secret material, do not commit it.

## Fixture Set v0

### 1. Provider summary

File:

- `provider-summary.json`

Should contain:

- provider name
- provider implementation
- scratch crate version/commit note
- ciphersuite
- group ID label
- phase/run mode
- warnings that fixtures are dev-only

### 2. Device summaries

Files:

- `alice-device-summary.json`
- `bob-device-summary.json`

Should contain:

- device label
- public identity label
- public KeyPackage hash/reference length
- member index if available
- no private key material

### 3. Public setup bundle summary

Files:

- `alice-public-setup-summary.json`
- `bob-public-setup-summary.json`

Should contain:

- public setup type
- ciphersuite
- hash/reference length
- serialized length if emitted
- no private key material

### 4. Join material summary

File:

- `welcome-summary.json`

Should contain:

- join material type
- encrypted secret count
- serialized length if emitted
- group info present flag
- ratchet tree extension flag if available

### 5. Conversation state summary

Files:

- `conversation-before-message.json`
- `conversation-after-message-one.json`
- `conversation-after-reload.json`
- `conversation-after-message-two.json`

Should contain:

- group ID label
- Alice epoch
- Bob epoch
- Alice member count
- Bob member count
- checkpoint label
- no secret state

### 6. Protected message summary

Files:

- `message-one-summary.json`
- `message-two-summary.json`

Should contain:

- message label
- content type
- epoch
- sender summary
- protected byte length if emitted
- plaintext hash or length
- plaintext sample only if explicitly safe/dev text

### 7. Provider event stream

File:

- `provider-events.jsonl`

Should contain JSONL events such as:

- `provider.identity.created`
- `provider.public_bundle.created`
- `conversation.created`
- `conversation.member_added`
- `conversation.welcome.created`
- `conversation.welcome.staged`
- `conversation.joined`
- `message.protected`
- `message.opened`
- `storage.saved`
- `storage.loaded`
- `conversation.loaded`
- `signature.invalid`
- `checkpoint.required`

### 8. Error fixture

File:

- `invalid-signature-error.json`

Should represent the known failure:

- phase-B fresh signer caused Bob to reject message two
- OpenMLS error class: `ValidationError(InvalidSignature)`
- CarbonStack mapping candidate: identity/signature validation failure
- likely trust action: block, warn, require reverify, append trust event

## Provider Boundary Concepts To Validate

Fixture work should validate these concepts before integration:

- `ProviderPublicBundle`
- `ProviderJoinMaterial`
- `ProviderProtectedMessage`
- `ProviderPlaintext`
- `ProviderEvent`
- `ProviderError`
- `ProviderCheckpoint`
- `ProviderStorageState`
- `ProviderIdentityState`
- `ConversationEpoch`
- `ConversationMemberSummary`

## Trust-State Mapping Questions

The fixture contract should answer:

- Which provider events should enter trust history?
- Which errors should block sending?
- Which errors should block opening?
- Which errors require reverify?
- Which events are normal protocol lifecycle?
- Which events are suspicious under a hostile server?

Known example:

- `ValidationError(InvalidSignature)` after signer mismatch should not be hidden as a generic message failure.

## Cypher Mapping Questions

Cypher should remain an opaque relay.

Fixture work should clarify which byte/string envelopes Cypher may carry later:

- public setup bundle
- Welcome/join material
- protected application message
- provider commit/proposal material
- provider event summaries only if client-local and non-secret

Cypher must not parse plaintext.

Cypher must not become identity truth.

Cypher must not silently rewrite provider material without client-visible consequences.

## Output Safety Rules

Do not commit:

- real private keys
- signing keys intended for real use
- provider storage files with secret material
- MemoryStorage JSON files
- temp signer JSON
- real recovery material
- generated Cargo target output

Possible to commit later if clearly dev-only and sanitized:

- summaries
- public lengths/hashes
- fake plaintext samples
- event names
- error-shape examples
- schema examples

## Recommended Implementation Shape

Extend the Rust scratch crate with a fixture mode:

- `cargo run -- fixtures`

The fixture mode should:

1. run the same local Alice/Bob persistence flow
2. write sanitized summaries to `fixtures/dev`
3. avoid writing secret provider storage into the repo
4. write provider event JSONL
5. write invalid-signature error shape from the known failure path, if easy

## Success Criteria

The fixture spike succeeds if:

- fixture mode runs locally
- generated fixture summaries contain no secret material
- provider event vocabulary is concrete enough to map into Go docs/tests
- provider error shape includes invalid signature
- no Cargo artifacts or secret state are tracked
- full local validation still passes

## Failure Criteria

Stop and document if:

- fixture generation requires secret material
- fixture outputs are too ambiguous
- event vocabulary is unclear
- OpenMLS API exposure makes safe summaries difficult
- repo hygiene becomes risky

## Next Work After Fixtures

After fixtures, choose between:

- Go-side provider-contract tests using fixture-shaped JSON
- Rust sidecar command prototype
- further OpenMLS negative-path tests
- mls-rs comparison if OpenMLS fixture/integration shape becomes ugly

## Allowed Claims

Allowed:

- CarbonStack is defining a provider fixture contract before integration.
- Fixture work is intended to reduce integration risk.
- OpenMLS scratch work remains isolated.

## Not Allowed Claims

Not allowed:

- OpenMLS is integrated into Comms.
- fixture summaries are real secure messaging.
- fixture mode is production provider behavior.
- CarbonStack has production E2EE.
- Cypher routes MLS payloads.
- trust-state integration is complete.
