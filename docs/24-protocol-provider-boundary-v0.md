# CarbonStack Protocol Provider Boundary v0

## Status

Classification: DRAFT / PHASE 2B FOUNDATION

This document defines the protocol-provider boundary for CarbonStackComms.

It does not implement real cryptography.

It defines how CarbonStack should separate application-level trust semantics from cryptographic protocol machinery.

## Current Decision

CarbonStack will use an MLS-shaped, provider-neutral architecture.

Signal/libsignal remains an important reference and fallback, but CarbonStack should not permanently couple mainline development to libsignal at this stage.

Current direction:

- CarbonStack owns trust semantics.
- Protocol providers own cryptographic machinery.
- Conversations are conceptually groups, even when they contain only two people.
- MLS is the preferred long-term protocol shape.
- Signal/libsignal is used as a reference for 1:1 messaging UX and design lessons.
- Noise remains reserved for transfer, provisioning, or appliance channels.
- libsodium remains reserved for narrow utilities or non-production experiments.

## Why This Boundary Exists

CarbonStack must avoid two failure modes:

1. Treating a cryptographic library as if it automatically solves product trust semantics.
2. Inventing custom cryptography because the application model was not clearly separated from protocol machinery.

Protocol providers can answer cryptographic questions.

CarbonStack must answer user-trust questions.

## CarbonStack Owns

CarbonStack application logic owns:

- local trust state
- contact and display labels
- manual verification UX
- trust history
- key-change warnings
- send/block policy
- revocation policy
- compromised-device policy
- local persistence policy
- text normalization policy
- allowed/not-allowed security claims
- user-facing warnings
- development vs strict mode behavior

CarbonStack must not delegate these decisions blindly to a provider.

## Protocol Provider Owns

A protocol provider owns:

- cryptographic identity material
- public key packages or public setup bundles
- conversation/session/group state
- message protection
- message opening
- protocol-level replay detection where supported
- protocol-level epoch/session errors
- protocol-level membership state where supported
- provider-specific serialization
- provider-specific persistence blobs

A provider may report security-relevant events, but CarbonStack decides what those events mean for UX and policy.

## Provider-Neutral Vocabulary

### ProviderIdentity

Local cryptographic identity controlled by a device.

In Phase 2A this is fake/dev material.

In a real provider this may map to:

- MLS credential/signature key material
- Signal identity key material
- other provider-specific identity material

### ProviderPublicBundle

Public material published for other devices to establish or join a conversation.

It may map to:

- MLS KeyPackage
- Signal prekey bundle
- other provider-specific setup material

The server may store this public material, but the server is not identity truth.

### ConversationID

CarbonStack-level conversation identifier.

A 1:1 chat is still conceptually a conversation with membership state.

### ConversationEpoch

Provider-reported conversation state version.

In MLS this maps naturally to group epoch.

In Signal-like providers this may be less direct and may map to session generation or local state version.

### ProtectedMessage

Provider-produced encrypted/protected message payload.

CarbonStackCypher should route this as opaque envelope material.

### OpenedMessage

Provider-opened message result.

It should include:

- plaintext
- sender/device identity signal
- conversation ID
- epoch/session state if relevant
- provider warnings or errors

### ProviderTrustSignal

Security-relevant provider signal that CarbonStack should evaluate.

Examples:

- identity changed
- unknown sender
- replay detected
- stale epoch
- revoked or removed member
- malformed protected message
- decryption failed
- membership changed
- provider state needs persistence

## Provider Result Philosophy

Providers should report facts.

CarbonStack should decide policy.

Example:

Provider reports:

- recipient identity changed

CarbonStack decides:

- dev mode warns and allows
- strict mode blocks until reverified
- trust history records the event

Example:

Provider reports:

- sender not in conversation membership

CarbonStack decides:

- reject message
- record warning
- surface local conflict state

## MLS Mapping

MLS is the preferred long-term architecture shape.

Approximate mapping:

- ProviderIdentity -> MLS credential and signature key material
- ProviderPublicBundle -> MLS KeyPackage
- ConversationID -> MLS Group ID or CarbonStack wrapper ID
- ConversationEpoch -> MLS group epoch
- Conversation membership -> MLS group membership
- Add member -> MLS Add proposal / Commit / Welcome
- Remove member -> MLS Remove proposal / Commit
- Message -> MLS application message
- Membership change -> epoch transition
- Device revocation -> remove member plus CarbonStack trust event

Why this fits CarbonStack:

- conversations can be group-shaped from the beginning
- 1:1 conversations are just two-member groups
- future groups do not require total conceptual rewrite
- membership changes can become loud, inspectable events
- device removal maps to group state changes
- epochs align with auditable trust history

MLS still does not solve everything.

CarbonStack must still define:

- user-facing verification
- local trust records
- display labels
- server conflict handling
- revocation UX
- storage/vault policy
- metadata privacy posture
- hostile-server behavior tests

## Signal/libsignal Mapping

Signal/libsignal remains a reference and fallback.

Approximate mapping:

- ProviderIdentity -> Signal identity key
- ProviderPublicBundle -> prekey bundle
- ConversationID -> pairwise session or CarbonStack wrapper conversation
- ConversationEpoch -> provider/session state version, less directly mapped
- Message -> Signal ciphertext
- Identity change -> safety-number/key-change warning equivalent

Why Signal remains useful:

- mature reference for 1:1 secure messaging
- useful design lessons for identity changes
- useful design lessons for prekey/offline delivery
- useful design lessons for safety-number UX

Why Signal is not mainline now:

- AGPL dependency risk
- integration complexity
- public repo states outside-of-Signal use is unsupported
- could force CarbonStack into Signal-shaped assumptions before MLS feasibility is tested

## Noise Mapping

Noise is not the main chat provider.

Future possible use:

- USB transfer mode
- local device provisioning
- recovery ceremony channel
- appliance-to-appliance bootstrap
- constrained secure transport

Noise should not be used as a replacement for a full asynchronous messaging protocol.

## libsodium Mapping

libsodium and NaCl-style primitives are not the main chat provider.

Possible future use:

- signed manifests
- local utility encryption
- non-production fixtures
- backup utility experiments
- narrow tooling

Do not compose a custom messenger protocol out of primitives.

## Proposed Provider Interface Shape

This is conceptual, not final Go code.

Required provider responsibilities:

- create local provider identity
- export public verification material
- export public setup bundle
- create or join conversation
- protect message
- open message
- report provider trust signals
- serialize provider state
- restore provider state

Conceptual operations:

- CreateIdentity()
- PublicBundle()
- PublicVerification()
- CreateConversation()
- JoinConversation()
- ProtectMessage()
- OpenMessage()
- ApplyProviderEvent()
- ExportState()
- ImportState()

## Provider Event Types

Initial provider events should include:

- identity_created
- public_bundle_created
- conversation_created
- conversation_joined
- message_protected
- message_opened
- identity_changed
- membership_changed
- member_added
- member_removed
- replay_detected
- stale_epoch
- malformed_message
- decrypt_failed
- state_updated

These are not all required in Phase 2B implementation.

They define the vocabulary the provider boundary should be able to grow into.

## CarbonStack Trust-State Interaction

Current Phase 2A trust states remain authoritative:

- unknown
- unverified
- verified
- changed
- revoked
- compromised

Provider output may cause transitions, but CarbonStack owns the transition policy.

Example:

Provider reports identity_changed.

CarbonStack transition:

- verified -> changed
- append trust event
- dev mode warns/allows
- strict mode blocks until reverified

Example:

CarbonStack user verifies new fingerprint.

CarbonStack transition:

- changed -> verified
- append trust event
- strict send allowed again

## Server Boundary

CarbonStackCypher should remain a dumb relay for as long as possible.

Cypher may store:

- account IDs
- device IDs
- public provider bundles
- encrypted/protected envelopes
- delivery state
- revocation announcements in later phases

Cypher must not be trusted for:

- plaintext
- private keys
- identity truth
- silent key replacement
- membership truth without client verification

## Phase 2B Code Direction

Add provider-neutral skeleton in:

- carbonstack-comms/internal/protocol

Later experimental provider path:

- carbonstack-comms/internal/protocol/mls

The MLS path may use Rust or another implementation strategy if needed.

The first code should be interfaces and test doubles only.

Do not integrate OpenMLS, libsignal, or any real provider yet.

## Phase 2B Minimal Implementation Target

Minimum code target:

- protocol.Provider interface
- protocol.MockProvider or StubProvider adapter around current MockCryptoProvider
- protocol result/event structs
- app code remains behaviorally unchanged
- validation runner still passes

This creates a stable replacement seam before real protocol work.

## Phase 2C MLS Feasibility Target

Future experimental target:

- two-member conversation
- local-only provider harness
- Alice creates identity
- Bob creates identity
- Alice creates conversation
- Bob joins
- Alice protects text
- Bob opens text
- membership/epoch state can be inspected
- provider state can be serialized/restored
- no production security claims

## Explicit Non-Goals

This document does not authorize:

- custom cryptography
- production encryption claims
- libsignal integration
- MLS implementation from scratch
- OpenMLS integration
- Android work
- CarbonStackOS work
- groups beyond conceptual two-member conversation modeling
- metadata privacy claims
- hostile-server proof claims

## Current Decision Summary

CarbonStack will proceed with:

- MLS-shaped provider-neutral architecture
- Signal/libsignal as reference and fallback
- no AGPL dependency in mainline unless absolutely necessary
- Rust allowed inside protocol/provider modules if needed
- no real provider integration until provider boundary skeleton is committed and validated

## Next Work

Implement provider-neutral skeleton in `carbonstack-comms/internal/protocol`.

Keep current MockCryptoProvider behavior working.

Local validation must continue passing.
