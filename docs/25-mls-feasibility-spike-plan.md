# CarbonStack MLS Feasibility Spike Plan

## Status

Classification: DRAFT / PHASE 2C PLANNING

This document defines the MLS feasibility spike for CarbonStack.

It does not implement MLS.

It does not authorize production cryptography claims.

It exists to prevent uncontrolled protocol exploration.

## Current Baseline

CarbonStack currently has:

- Phase 1 local relay/client lifecycle
- CarbonStackCypher local opaque-envelope relay
- CarbonStackComms CLI client
- Phase 2A trust-state scaffold
- trust-list
- changed-device reverify lifecycle
- provider-neutral protocol boundary doc
- provider-neutral protocol skeleton in CarbonStackComms

CarbonStack does not yet have:

- real cryptography
- MLS integration
- OpenMLS integration
- libsignal integration
- replay resistance
- metadata privacy
- hostile-server proof
- secure local vault
- hardware-key identity
- Android client
- CarbonStackOS implementation

## Current Protocol Direction

CarbonStack uses an MLS-shaped, provider-neutral architecture.

Design consensus:

- every conversation is conceptually group-shaped
- 1:1 conversations are two-member conversations
- MLS is the preferred long-term architecture shape
- Signal/libsignal remains a reference and fallback
- avoid AGPL dependencies in mainline unless necessary
- Rust is acceptable inside provider modules if it serves the project
- no custom cryptography

## Purpose of the MLS Spike

The MLS spike should answer one question:

Can MLS practically fit CarbonStack's provider boundary and trust model?

It should not try to build the full messenger.

## Non-Goals

The MLS spike must not include:

- production messaging
- Android
- CarbonStackOS
- hardware-key flows
- metadata privacy
- secure local vault
- public deployment
- production authentication
- group UX beyond local test cases
- custom cryptographic implementation
- claims of Signal-equivalent security
- claims of production-ready E2EE

## Candidate Implementation Direction

Preferred candidate to investigate:

- OpenMLS

Reason:

- Rust implementation
- aligns with RFC 9420 direction
- likely better fit for provider-module isolation than forcing libsignal into the Go CLI directly

Important:

- Do not import OpenMLS into mainline until the spike plan is clear.
- Do not implement MLS from scratch.
- Do not permanently commit to OpenMLS before feasibility is proven.

## Spike Location

Initial preferred path:

- carbonstack-comms/internal/protocol/mls

Rationale:

- keeps the provider near the existing protocol boundary
- makes later integration easier
- keeps experimental MLS work visibly isolated
- allows the mock provider and MLS provider to share vocabulary

If Rust integration becomes messy, fallback options:

- carbonstack-protocol-lab separate repo
- separate Rust sidecar binary
- temporary research branch
- local-only prototype outside mainline

## Minimal Feasibility Questions

The spike must answer:

1. Can Alice create an MLS identity or credential?
2. Can Bob create an MLS identity or credential?
3. Can Alice publish setup material equivalent to a provider public bundle?
4. Can Bob publish setup material equivalent to a provider public bundle?
5. Can Alice create a two-member conversation?
6. Can Bob join that conversation?
7. Can Alice protect a text message?
8. Can Bob open that text message?
9. Can provider state be serialized and restored?
10. Can conversation epoch or state version be inspected?
11. Can membership state be inspected?
12. Can CarbonStack map MLS events into provider events?
13. Can CarbonStack trust-state logic remain outside the provider?
14. Can the spike run locally and repeatably?

## Minimal Test Flow

Required local-only flow:

1. Alice provider identity created.
2. Bob provider identity created.
3. Alice public bundle exported.
4. Bob public bundle exported.
5. Alice creates a conversation.
6. Bob joins the conversation.
7. Alice protects plaintext:
   - `hello from Alice over MLS spike`
8. Bob opens plaintext.
9. Test confirms plaintext matches.
10. Test confirms conversation has two members.
11. Test confirms epoch/state version is available.
12. Test exports provider state.
13. Test imports provider state.
14. Test still opens or protects another message after restore, if supported.

## CarbonStack Provider Mapping

The spike should map MLS concepts into current provider vocabulary.

Expected mapping:

- MLS credential/signature key -> ProviderIdentity
- MLS KeyPackage -> PublicBundle
- MLS group ID or wrapper ID -> ConversationID
- MLS epoch -> ConversationEpoch
- MLS application message -> ProtectedMessage
- MLS opened message -> OpenedMessage
- MLS commit/welcome/membership event -> ProviderEvent
- MLS state serialization -> provider state persistence

## Trust-State Boundary Rule

MLS provider code may report provider facts.

CarbonStack trust code decides policy.

Provider may report:

- member added
- member removed
- stale epoch
- malformed message
- decrypt failed
- identity or credential mismatch
- state updated

CarbonStack decides:

- whether to warn
- whether to block send
- whether to mark device changed
- whether to require reverify
- whether to append trust history
- whether to reject revoked devices

## Success Criteria

The MLS spike is successful if:

- local two-member conversation works
- text can be protected/opened
- provider state can be serialized/restored
- epoch or state version is inspectable
- membership is inspectable
- provider events can map cleanly to CarbonStack provider events
- trust-state logic remains outside the MLS provider
- the local validation path still passes
- no production claims are introduced

## Failure Criteria

The MLS spike should be considered blocked or failed if:

- library integration is too unstable for local tests
- provider state cannot be reasonably persisted
- Go/Rust boundary becomes too complex for current project stage
- identity and membership mapping does not fit CarbonStack's trust model
- test harness cannot be made repeatable
- implementation requires custom cryptography
- implementation forces premature Android or OS work
- implementation pulls in unacceptable licensing constraints

## Decision Outcomes

The spike should produce one of these outcomes:

### Outcome A: MLS feasible now

Continue toward experimental MLS provider behind `internal/protocol`.

### Outcome B: MLS feasible but Rust/sidecar needed

Define a Rust provider module or helper boundary.

### Outcome C: MLS conceptually fits but implementation is too heavy now

Keep MLS as long-term shape and continue improving provider boundary/tests.

### Outcome D: MLS does not fit CarbonStack's near-term needs

Re-evaluate Signal-like 1:1 provider or other mature protocol options.

## Required Documentation After Spike

After any MLS spike attempt, update or add:

- MLS feasibility result doc
- provider boundary doc
- validation matrix
- allowed/not-allowed claims
- LogDoc checkpoint

## Allowed Claims Before Spike

Allowed:

- CarbonStack has an MLS-shaped provider-neutral architecture.
- CarbonStack has a provider skeleton.
- CarbonStack intends to investigate MLS feasibility.
- CarbonStack has not integrated MLS yet.

Not allowed:

- CarbonStack uses MLS.
- CarbonStack has real encryption.
- CarbonStack has real protocol security.
- CarbonStack has group messaging.
- CarbonStack has production-ready E2EE.

## Next Work

Next code step should be a tiny experimental MLS research scaffold only if implementation dependencies are understood.

Before importing anything, identify:

- likely Rust crate
- license
- minimum local example
- build requirements on Windows
- whether a Go/Rust sidecar boundary is cleaner than direct integration
