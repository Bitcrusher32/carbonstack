# CarbonStack OpenMLS Minimal Example Plan

## Status

Classification: DRAFT / PHASE 2C OPENMLS PLANNING

This document defines the smallest OpenMLS experiment CarbonStack should attempt.

It does not import OpenMLS.

It does not implement MLS.

It does not authorize production security claims.

## Current Baseline

CarbonStack currently has:

- Phase 2A trust-state scaffold
- provider-neutral protocol skeleton
- MLS-shaped provider boundary
- MLS feasibility spike plan
- MLS implementation candidate notes
- reserved experimental MLS provider slot in CarbonStackComms

Current experimental slot:

- carbonstack-comms/internal/protocol/mls

## Why This Document Exists

OpenMLS is a real cryptographic protocol implementation.

The first contact with it must be constrained.

This document prevents the project from drifting into broad integration, side quests, or premature production claims.

## Current Candidate

Primary candidate:

- OpenMLS

Reasons:

- Rust implementation
- RFC 9420 oriented
- high-level group API
- aligns with CarbonStack's MLS-shaped conversation model
- apparently compatible with an in-project provider-module approach

Alternate candidate:

- mls-rs

Reasons to keep as alternate:

- advertises RFC 9420 conformance
- advertises multiple identity/group support
- advertises configurable storage including in-memory and SQLite
- may become easier if OpenMLS ergonomics are poor

## Minimal Experiment Rule

The first OpenMLS experiment must be local-only.

It must not use:

- CarbonStackCypher
- CarbonStackComms send/inbox commands
- real user state
- network delivery
- production persistence
- hardware keys
- Android
- CarbonStackOS

It may use:

- local test fixtures
- temporary in-memory state
- temporary files
- Rust-only scratch code
- provider-boundary vocabulary from `internal/protocol`

## Minimal Success Flow

The smallest useful OpenMLS flow is:

1. Alice identity/credential exists.
2. Bob identity/credential exists.
3. Alice setup material exists.
4. Bob setup material exists.
5. Alice creates a group/conversation.
6. Bob is added or joins through the required MLS flow.
7. Alice protects text:
   - `hello from Alice over OpenMLS spike`
8. Bob opens the text.
9. Test confirms plaintext matches.
10. Test confirms membership includes Alice and Bob.
11. Test confirms epoch or state version can be observed.
12. Test can export/restore state if the library path is straightforward.

## Absolute Non-Goals

Do not attempt:

- polished CLI integration
- Cypher envelope delivery
- app-level trust-state integration
- revocation mapping
- multi-device
- multi-group UI
- persistence design
- secure vault design
- hostile-server tests
- metadata privacy
- performance work
- Android work
- OS work

## Research Questions Before Code

Before adding an OpenMLS dependency, answer:

1. What crate or crates are required?
2. What license applies to each required crate?
3. What is the smallest current upstream example?
4. What crypto provider is expected?
5. What credential type is simplest for a local test?
6. What storage/key-store provider is simplest for local test?
7. Can the example build on Windows?
8. Does the example require nightly Rust?
9. How does state serialization/export work?
10. How does membership inspection work?
11. How does epoch inspection work?
12. What errors are exposed for malformed/stale/decrypt-failure cases?

## Suggested First Code Shape

The first code should be a Rust-only scratch experiment, not Go integration.

Preferred location if kept in-project:

- carbonstack-comms/internal/protocol/mls/research/openmls-minimal

This keeps the spike near the provider boundary while avoiding premature integration.

The first code may be deleted or rewritten later.

## Why Not Go/Rust Integration First

Go/Rust integration is a second problem.

The first problem is whether MLS fits CarbonStack's model at all.

Therefore:

- first prove the OpenMLS local flow
- then decide sidecar vs FFI vs other integration
- only later wire provider behavior into Go Comms

## Provider Boundary Mapping Target

The local OpenMLS result should eventually map to:

- OpenMLS credential/key material -> ProviderIdentity
- OpenMLS KeyPackage/setup material -> PublicBundle
- OpenMLS group -> ConversationState
- OpenMLS epoch -> ConversationEpoch
- OpenMLS application message -> ProtectedMessage
- OpenMLS opened application message -> OpenedMessage
- OpenMLS state changes -> ProviderEvent

## Success Criteria

This planning stage is successful when:

- the minimal OpenMLS example path is understood
- required crates and licenses are known
- build requirements are known
- smallest experiment location is chosen
- no production claims are added
- local validation still passes

## Failure / Pause Criteria

Pause before code if:

- required dependencies have unacceptable licenses
- upstream examples are too unstable
- Windows build requirements are unclear
- Rust toolchain requirements are excessive
- the minimal group flow cannot be understood from docs/examples
- the project starts drifting into full integration before a local scratch proof

## Current Recommendation

Continue with OpenMLS as the first spike candidate.

Keep mls-rs as alternate.

Next concrete step:

- inspect OpenMLS examples/docs
- identify minimal crate list
- identify simplest credential/provider/key-store setup
- then create a local Rust-only scratch experiment if feasible

## Allowed Claims

Allowed:

- CarbonStack has an OpenMLS minimal example plan.
- CarbonStack is preparing a constrained OpenMLS feasibility experiment.
- OpenMLS is the first intended spike candidate.
- mls-rs remains an alternate candidate.

Not allowed:

- CarbonStack uses OpenMLS.
- CarbonStack has MLS encryption.
- CarbonStack has real E2EE.
- CarbonStack has selected a final MLS implementation.
- CarbonStack has production-ready secure messaging.
