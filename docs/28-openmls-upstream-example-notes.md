# CarbonStack OpenMLS Upstream Example Notes

## Status

Classification: DRAFT / PHASE 2C RESEARCH

This document records upstream OpenMLS concepts that matter before writing the first CarbonStack MLS scratch experiment.

This is not implementation code.

This does not integrate OpenMLS.

This does not authorize production security claims.

## Purpose

CarbonStack is preparing for a constrained OpenMLS feasibility spike.

Before adding any dependency, the project needs a plain-language map of the smallest upstream concepts required for the first local-only experiment.

## Upstream Grounding

OpenMLS is a Rust implementation of Messaging Layer Security as specified in RFC 9420.

OpenMLS documentation describes a high-level API for creating and managing MLS groups through `MlsGroup`.

OpenMLS also describes interchangeable components for:

- cryptographic provider
- key store
- random number generator

RFC 9420 defines MLS as an asynchronous group keying protocol with forward secrecy and post-compromise security.

For CarbonStack, the critical point is:

- MLS is group-shaped by design.
- CarbonStack has already decided that even 1:1 conversations are conceptually two-member conversations.
- Therefore MLS can be evaluated as a natural fit for CarbonStack's long-term conversation model.

## Minimal Upstream Concepts To Understand

### MlsGroup

`MlsGroup` is the central upstream concept for creating and managing an MLS group.

CarbonStack mapping:

- `MlsGroup` maps to provider-side conversation state.
- CarbonStack wrapper maps it to `ConversationState`.
- MLS epoch maps to `ConversationEpoch`.

### Credential / Identity Material

OpenMLS requires some form of credential or signing identity.

CarbonStack mapping:

- credential/signing identity maps to `ProviderIdentity`.
- public verification material maps to `PublicVerification`.
- CarbonStack trust state remains outside the provider.

### KeyPackage

MLS uses public setup material for adding members.

CarbonStack mapping:

- KeyPackage maps to `PublicBundle`.
- Cypher may eventually store public bundles.
- Cypher must not become identity truth.

### Welcome

When a member is added, a Welcome is used for the new member to join.

CarbonStack mapping:

- Welcome maps to provider-specific join material.
- It may later be routed through Cypher as opaque provider material.
- It is not user plaintext.

### StagedWelcome

OpenMLS documentation describes joining from a Welcome as a staged process.

CarbonStack implication:

- The provider may need a staged join operation.
- The provider boundary may eventually need to represent provider events during join.
- The first scratch experiment should not hide this complexity prematurely.

### Ratchet Tree

OpenMLS documentation notes that if the group configuration does not use the ratchet tree extension, the ratchet tree needs to be provided during Welcome processing.

CarbonStack implication:

- The first experiment should choose the simplest upstream-supported configuration.
- If possible, include ratchet tree extension or follow the simplest documented path.
- Do not invent a workaround.

### Application Message

MLS protects application messages after group state exists.

CarbonStack mapping:

- MLS application message maps to `ProtectedMessage`.
- MLS open/decrypt result maps to `OpenedMessage`.

## First Experiment Shape

The first experiment should be Rust-local only.

It should not use:

- CarbonStackCypher
- Go CLI commands
- production persistence
- Android
- CarbonStackOS
- hardware keys

The experiment should try to prove:

1. Alice credential/identity can be created.
2. Bob credential/identity can be created.
3. Bob KeyPackage/public setup material can be created.
4. Alice creates an MLS group.
5. Alice adds Bob using Bob's setup material.
6. Bob joins through Welcome/staged Welcome.
7. Alice protects an application message.
8. Bob opens the application message.
9. The opened plaintext matches.
10. Epoch or group state version can be inspected.
11. Membership can be inspected or at least inferred from the group state.
12. State serialization/persistence requirements are understood.

## Expected OpenMLS Questions During Code

The first code attempt should answer:

- Which OpenMLS crates are required?
- Which crypto provider is simplest?
- Which key store is simplest?
- Which credential type is simplest?
- What exact object creates a KeyPackage?
- What exact call adds Bob?
- What exact object contains the Welcome?
- What exact call stages and accepts the Welcome?
- What exact call protects an application message?
- What exact call opens/processes an application message?
- How is epoch inspected?
- How is membership inspected?
- How is group state persisted?

## CarbonStack Provider Boundary Notes

The current `internal/protocol` skeleton may need future changes after the scratch experiment.

Expected possible changes:

- explicit join material type
- explicit welcome material type
- clearer provider event shape
- state-updated signal after commits
- membership event representation
- epoch mismatch/stale epoch errors
- serialization boundary for provider state

Do not treat the current provider skeleton as final.

## Risks

### API drift

OpenMLS APIs may change.

Keep scratch code isolated and documented.

### Go/Rust boundary

The first experiment should be Rust-only.

Do not solve Go/Rust integration until the MLS flow itself is understood.

### Persistence complexity

MLS provider state may require more structured persistence than current mock provider state.

Record findings before integrating.

### Conceptual mismatch

If MLS provider state or Welcome handling does not map cleanly to CarbonStack's provider boundary, update the boundary instead of forcing an awkward mapping.

## Current Decision

Proceed with a Rust-only scratch experiment under the in-project experimental slot.

Intended path:

- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal`

Still prohibited:

- production integration
- OpenMLS provider wired into `comms send`
- user-facing MLS claims
- mainline security claims

## Allowed Claims

Allowed:

- CarbonStack has upstream OpenMLS research notes.
- CarbonStack understands the first OpenMLS concepts it needs to investigate.
- CarbonStack intends a Rust-only scratch experiment.

Not allowed:

- CarbonStack uses OpenMLS.
- CarbonStack has MLS encryption.
- CarbonStack has real E2EE.
- CarbonStack has selected a final MLS implementation.
