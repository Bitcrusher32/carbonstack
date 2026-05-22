# CarbonStack OpenMLS Provider Boundary Implications

## Status

Classification: PHASE 2C DESIGN / PRE-INTEGRATION

This document records what the OpenMLS scratch probes imply for CarbonStack's provider boundary.

It does not implement a provider.

It does not wire OpenMLS into CarbonStackComms.

It does not authorize production security claims.

## Purpose

The OpenMLS scratch result proved that a local two-member MLS conversation can be created, joined, and used to protect/open application messages. The current scratch crate has also validated a two-message state-continuity probe inside one process.

The next question is not "how do we immediately wire this into Comms?"

The next question is:

What must CarbonStack's provider boundary represent so OpenMLS can fit without corrupting CarbonStack's trust model?

## Source Context

Relevant scratch result:

- `docs/29-openmls-scratch-result-v0.md`

Relevant implementation path:

- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal`

Relevant existing provider skeleton:

- `carbonstack-comms/internal/protocol`

## Core Design Shift

The provider is not a stateless encrypt/decrypt utility.

The provider is a stateful protocol engine.

OpenMLS showed that provider operations may:

- create local protocol identity material
- create public setup material
- create group/conversation state
- consume Welcome/join material
- mutate local state while processing messages
- expose epochs
- expose membership state
- require persistent local storage
- emit events or errors that CarbonStack must interpret

Therefore the provider boundary must explicitly model state.

## Required Concepts

### Provider Identity

CarbonStack needs a provider identity concept.

This should represent local cryptographic identity material.

It must not be confused with user display name.

Possible provider-side meaning:

- OpenMLS credential
- OpenMLS signing key
- public verification key
- local device identity material

CarbonStack mapping:

- device identity
- local provider identity
- trust-state record input

### Public Setup Material

CarbonStack needs a public setup bundle concept.

For OpenMLS this maps to:

- KeyPackage or equivalent public setup material

CarbonStack must treat this as public protocol material, not as proof of trust by itself.

Possible boundary name:

- `PublicBundle`
- `SetupBundle`
- `ProviderPublicBundle`

### Join Material

OpenMLS Welcome flow means CarbonStack needs a join-material concept.

For OpenMLS this maps to:

- Welcome
- possibly ratchet-tree-related material
- provider-specific join envelope

Possible boundary name:

- `JoinMaterial`
- `WelcomeMaterial`
- `ProviderJoinMaterial`

CarbonStackCypher may eventually route this as opaque provider material.

CarbonStack must not parse it as user plaintext.

### Conversation State

OpenMLS group state maps to CarbonStack conversation state.

This is not only a conversation ID.

It may include:

- MLS group state
- epoch
- member state
- pending commits
- secret tree state
- provider-specific storage references

Possible boundary name:

- `ConversationState`
- `ProviderConversationState`

### Provider Storage

OpenMLS scratch showed that provider storage is device-local protocol state.

Alice and Bob required separate provider/storage instances.

Shared provider/storage caused `GroupAlreadyExists`.

Provider storage must be modeled as local to:

- device
- account/device identity
- provider instance
- possibly conversation

CarbonStack must avoid global shared provider state across devices.

### Epoch

OpenMLS exposes group epoch.

CarbonStack likely needs to observe epoch for:

- debugging
- stale-message behavior
- trust history
- membership-change display
- replay/tamper tests
- validation scripts

Epoch should be represented as provider-originated state, not as user-trusted product policy by itself.

Possible boundary name:

- `ConversationEpoch`
- `ProviderEpoch`

### Membership Summary

OpenMLS can expose member state.

CarbonStack needs a way to inspect or summarize membership.

This may be needed for:

- loud membership changes
- server-hostile group state
- trust-list display
- reverify prompts
- revocation handling
- future MLS group UX

Possible boundary name:

- `MemberSummary`
- `ProviderMember`
- `ConversationMember`

### Provider Events

OpenMLS operations imply events.

CarbonStack needs a provider event stream or result metadata.

Possible provider events:

- identity created
- public bundle created
- conversation created
- member added
- Welcome produced
- Welcome staged
- conversation joined
- message protected
- message opened
- epoch advanced
- membership changed
- state updated
- stale epoch detected
- malformed message rejected
- decrypt/open failed
- pending commit produced
- pending commit merged

CarbonStack application logic decides how these become UX warnings, blocks, or trust-history entries.

## Operation Implications

### Create Identity

Provider operation:

- create local provider identity

Should return:

- provider identity handle or serialized local state reference
- public verification material
- provider event

Must not return:

- raw private keys to normal application logs
- user-facing trust claims

### Create Public Bundle

Provider operation:

- create public setup material for a device

OpenMLS equivalent:

- KeyPackage

Should return:

- public bundle bytes
- bundle fingerprint/hash reference if useful
- provider event

### Create Conversation

Provider operation:

- create a new conversation/group

OpenMLS equivalent:

- `MlsGroup::new_with_group_id`

Should return:

- conversation state
- conversation ID
- epoch
- member summary
- state-updated event

### Add Member

Provider operation:

- add another member using their public bundle

OpenMLS equivalent:

- `MlsGroup::add_members`

Should return:

- outbound commit material
- join/Welcome material for added member
- updated epoch or pending epoch
- membership-change event
- state-updated event

Open question:

- Should the provider merge immediately, or should CarbonStack explicitly approve/merge staged commits?

### Join Conversation

Provider operation:

- join a conversation using join material

OpenMLS equivalent:

- `StagedWelcome::new_from_welcome`
- `into_group`

Should return:

- conversation state
- epoch
- membership summary
- state-updated event

### Protect Message

Provider operation:

- protect plaintext application message

OpenMLS equivalent:

- `create_message`

Should return:

- protected provider message bytes
- content type / provider message type
- epoch
- state-updated event if applicable

Important lesson:

- OpenMLS `create_message` required mutable Alice group state in the scratch probe.
- Outbound protection may mutate provider state.
- Future provider persistence must checkpoint after sends, not only after receives.

### Open / Process Message

Provider operation:

- process protected provider message

OpenMLS equivalent:

- deserialize `MlsMessageIn`
- convert to `ProtocolMessage`
- `process_message`
- extract `ProcessedMessageContent`

Important lesson:

- processing/opening may mutate local group/provider state
- both outbound protection and inbound opening should be treated as persistence-relevant state transitions

Should return:

- opened plaintext if application message
- provider events
- updated epoch if changed
- state-updated marker
- non-application message classification if proposal/commit/etc.

CarbonStack must not treat this as a pure decrypt function.

## Error Implications

Provider errors should be typed enough for CarbonStack policy.

Useful error classes:

- malformed provider message
- wrong conversation
- stale epoch
- future epoch
- unknown sender
- untrusted member
- missing local state
- duplicate group
- join material invalid
- decrypt/open failed
- unsupported provider content
- provider storage failure
- state serialization failure

CarbonStack policy should decide:

- warn
- block
- retry
- mark device changed
- require reverify
- append trust event
- quarantine message
- fail closed

## Persistence Implications

The next OpenMLS experiment should test provider persistence/restart behavior.

Questions:

- What does OpenMLS store in the provider storage?
- Can provider/group state survive process restart?
- What must be persisted after identity creation?
- What must be persisted after KeyPackage creation?
- What must be persisted after group creation?
- What must be persisted after Welcome join?
- What must be persisted after message processing?
- Can state be exported as provider-owned bytes?
- Is a sidecar/provider database needed?
- How do we avoid logging secrets?
- How does this map to future secure vault design?

## Cypher Implications

CarbonStackCypher should remain an opaque relay.

It may eventually route:

- public setup bundles
- join/Welcome material
- protected application messages
- provider commit/proposal messages
- revocation notices

It must not become trusted identity truth.

It must not parse plaintext.

It must not silently rewrite membership state.

It must not silently replace keys without client-visible trust consequences.

## Trust-State Implications

CarbonStack trust logic remains above provider mechanics.

Provider can report:

- member added
- epoch advanced
- message opened
- join material consumed
- unknown sender
- malformed message
- state updated

CarbonStack decides:

- whether a change is user-visible
- whether to block sending
- whether to require verification
- whether to append trust history
- whether to mark identity changed
- whether to accept/reject a member

## Current Provider Skeleton Review Targets

Review current `carbonstack-comms/internal/protocol` for whether it needs:

- explicit `JoinMaterial`
- explicit `PublicBundle`
- explicit `ProviderMessage`
- explicit `ProviderEvent`
- explicit `ConversationEpoch`
- explicit `MemberSummary`
- explicit `StateUpdated` marker
- better separation between identity state and conversation state
- better error taxonomy
- clear persistence hooks

## Recommended Next Code Experiment

Before provider integration, run a persistence/restart scratch experiment.

Suggested flow:

1. Create Alice provider/state.
2. Create Bob provider/state.
3. Create Alice group.
4. Add Bob.
5. Bob joins from Welcome.
6. Alice sends message 1.
7. Bob opens message 1.
8. Simulate process restart or reload provider/group state.
9. Alice sends message 2.
10. Bob opens message 2.
11. Record exactly what persistence was required.

Current pre-restart result:

- A two-message in-memory state-continuity probe has already passed.
- The next experiment should move from in-memory continuity to real provider storage/export/reload behavior.

## Allowed Claims

Allowed:

- CarbonStack has derived provider-boundary implications from the OpenMLS scratch result.
- OpenMLS appears feasible enough to continue Phase 2C.
- The provider boundary must treat the provider as stateful.
- Open/process operations may mutate provider state.

## Not Allowed Claims

Not allowed:

- OpenMLS is integrated into CarbonStackComms messaging.
- CarbonStack has production E2EE.
- CarbonStack has selected OpenMLS as final provider.
- Provider persistence is solved.
- Cypher carries real MLS traffic.
- Hostile-server security is solved.
- Replay resistance is tested.
- Metadata privacy is solved.

