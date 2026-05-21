# CarbonStack Protocol Feasibility Matrix

## Status

Classification: DRAFT / PHASE 2 FOUNDATION

This document compares candidate protocol foundations for CarbonStackComms and CarbonStackCypher.

This is not a final protocol decision and not an implementation document.

## Current Phase 2 Posture

CarbonStack Phase 2 should not begin by integrating a cryptographic library directly.

Phase 2 should begin with CarbonStack-specific trust machinery:

- local device identity records
- fake/dev fingerprints
- trust states
- verification commands
- key-change warnings
- revocation states
- send policy based on trust state
- append-only trust history

Only after that should the MockCryptoProvider be replaced or supplemented by a real protocol provider.

## User-Selected Phase 2 Defaults

Initial network size:

- 2 people

Initial device model:

- one device per person

Account and message recovery:

- no message recovery in Phase 2
- future identity or account namespace recovery is important but deferred
- lost device means re-enrollment and old local message access is gone unless future backup design exists

Unverified device policy:

- development mode: allow sending with loud warning
- mature mode: block sending until verified

Verification ceremony order:

- first: manual fingerprint comparison
- second: QR code
- later: hardware key, trusted existing device, shared phrase, or other recovery/enrollment ceremonies

AGPL posture:

- unresolved
- acceptable for experimental research branch if needed
- must be understood before permanent libsignal coupling

Rust posture:

- allowed if it is the cleanest protocol path
- Go remains preferred for current CLI and server plumbing

Metadata privacy:

- deferred until maturity or MLS/group transition stage

Server public key material:

- acceptable for the server to store public device/prekey material
- not acceptable for the server to be identity truth
- local trust history must override server state

Key-change policy:

- development mode: warn loudly and allow override
- mature mode: block sending until reverified

## Evaluation Criteria

Each protocol candidate is evaluated against:

1. Fit for 1:1 text messaging
2. Fit for future group messaging
3. Offline delivery support
4. Hostile-server suitability
5. Device identity support
6. Loud key-change compatibility
7. Revocation compatibility
8. Replay-resistance support
9. Forward secrecy and post-compromise recovery
10. CLI-first testability
11. Android and future CarbonStackOS path
12. Licensing risk
13. Implementation complexity
14. Risk of accidental custom cryptography
15. Fit for CarbonStack's constrained appliance thesis

## Candidate Summary

### Candidate A: Signal Protocol / libsignal

Intended role:

- strongest initial candidate for 1:1 messaging

Why it fits:

- built around asynchronous secure messaging
- aligns with identity keys, sessions, prekeys, safety-number-like UX, forward secrecy, and post-compromise recovery
- strong conceptual fit for CarbonStack's 1:1 Phase 2 direction
- good match for offline delivery

Strengths:

- best match for 1:1 private messaging
- mature design lineage
- strong fit for device identity and key-change warnings
- good conceptual path from MockCryptoProvider to SignalLikeProvider
- supports the idea that server stores public setup material but not plaintext

Weaknesses and risks:

- AGPL-3.0 license requires serious project decision before permanent coupling
- current public APIs are Java, Swift, and TypeScript, with Rust internals
- Go integration may require wrapper process, FFI, Rust sidecar, TypeScript prototype, or later Android-first integration
- not automatically a CarbonStack trust model
- could create false confidence if integrated before trust-state semantics are built
- group messaging path may still require MLS or separate design

Licensing note:

libsignal is AGPL-3.0. If CarbonStackComms links to or incorporates libsignal, assume the Comms repo or affected component may need AGPL-compatible licensing. Do not permanently commit to this until the license implications are understood.

Fit rating:

- 1:1 messaging: strong
- future groups: partial
- transfer channels: not primary
- Phase 2A trust-state scaffolding: not needed yet
- Phase 2B feasibility spike: strong candidate
- immediate implementation: not yet

Initial decision:

Signal/libsignal should be the first serious 1:1 feasibility candidate, but not before CarbonStack trust state is implemented with fake/dev fingerprints.

## Candidate B: MLS

Intended role:

- future group messaging candidate
- possible mature long-term conversation-state foundation

Why it fits:

- designed for secure group messaging
- supports asynchronous group keying
- designed around group membership, epochs, forward secrecy, and post-compromise security
- maps well to CarbonStack's future auditable group membership requirements

Strengths:

- best conceptual match for future groups
- standardized as RFC 9420
- strong future fit for membership changes and group epochs
- better long-term match than ad-hoc group design
- may eventually support a unified group-first model

Weaknesses and risks:

- probably too heavy for immediate 1:1 Phase 2
- library maturity and binding practicality must be evaluated
- more complex state machine than needed for two-person Phase 2
- does not automatically solve CarbonStack UX, verification ceremony, or device recovery policy
- may slow the project down if chosen too early

Fit rating:

- 1:1 messaging: possible but likely overkill
- future groups: strong
- transfer channels: not primary
- Phase 2A trust-state scaffolding: not needed yet
- Phase 2B feasibility spike: research candidate
- immediate implementation: no

Initial decision:

MLS should remain the likely future group candidate, but Phase 2 should not start with MLS implementation unless libsignal feasibility fails and MLS proves simpler than expected.

## Candidate C: Noise Framework

Intended role:

- future narrow secure channels
- transfer, provisioning, local device-to-device, or appliance handshake work

Why it fits:

- framework for building crypto protocols
- useful for authenticated handshakes and secure channels
- could fit USB transfer mode, local enrollment mode, provisioning, or recovery transport

Strengths:

- flexible
- good for narrow channels
- strong fit for future CarbonStackOS appliance transfer/provisioning flows
- simpler conceptual surface than full messenger protocols for transport-specific work

Weaknesses and risks:

- not a full messenger protocol
- does not provide asynchronous messaging state by itself
- does not define trust history, revocation, key-change UX, prekeys, or group state
- high risk of accidentally designing a custom messaging protocol if used as the main chat foundation

Fit rating:

- 1:1 messaging: partial and risky as main protocol
- future groups: weak
- transfer channels: strong
- Phase 2A trust-state scaffolding: not needed yet
- Phase 2B feasibility spike: later, for transfer/provisioning
- immediate implementation: no

Initial decision:

Noise should not be the main CarbonStackComms chat protocol. It is a strong future candidate for constrained transfer or appliance channels.

## Candidate D: libsodium / NaCl-style primitives

Intended role:

- utilities, experiments, file encryption, signed manifests, non-production fixtures, or narrow tooling

Why it fits:

- practical and accessible primitives
- useful for small controlled utilities
- may be helpful for local experiments or future backup/manifest tooling

Strengths:

- simpler than full protocol stacks
- useful building blocks
- good for utility cryptography when carefully scoped
- useful for testing concepts if clearly marked non-production

Weaknesses and risks:

- primitives are not a messaging protocol
- sealed boxes do not authenticate sender identity
- high risk of composing a bad custom protocol
- does not automatically provide forward secrecy, post-compromise recovery, prekeys, trust history, revocation, or replay handling

Fit rating:

- 1:1 messaging: weak as main protocol
- future groups: weak
- transfer channels: possible for narrow cases
- utilities: useful with care
- Phase 2A trust-state scaffolding: not needed
- immediate implementation: only for non-production experiments if necessary

Initial decision:

libsodium should not be the main CarbonStackComms protocol foundation. Use only for narrow utilities or explicitly non-production experiments.

## Feasibility Matrix

| Criterion | Signal / libsignal | MLS | Noise | libsodium primitives |
|---|---|---|---|---|
| 1:1 text messaging | Strong | Possible but heavy | Partial | Weak |
| Future groups | Partial | Strong | Weak | Weak |
| Offline delivery | Strong fit | Possible | Not complete alone | Not complete alone |
| Hostile-server content secrecy | Strong fit | Strong fit | Possible if designed | Possible if designed |
| Device identity support | Strong fit | Requires app mapping | Requires app design | Requires app design |
| Loud key-change UX | Strong conceptual fit | Requires app mapping | Requires app design | Requires app design |
| Revocation model | Partial, app-layer needed | Stronger group fit | App-layer needed | App-layer needed |
| Replay resistance | Strong protocol fit | Strong protocol fit | Pattern/design dependent | Must design |
| Forward secrecy | Strong | Strong | Pattern dependent | Not automatic |
| Post-compromise recovery | Strong | Strong | Pattern dependent | Not automatic |
| CLI-first feasibility | Unclear | Unclear | Good for narrow tests | Good for utilities |
| Android path | Stronger than Go path | Research needed | Possible | Possible |
| Go path | Unclear/harder | Depends on library | Likely possible | Good |
| Licensing risk | High due AGPL | Library-dependent | Usually low, implementation-dependent | Usually manageable |
| Risk of custom crypto mistakes | Medium if misused | Medium | High as main chat | High as main chat |
| CarbonStack appliance fit | Strong for chat | Strong for groups | Strong for transfer | Utility-only |

## Recommended Phase 2 Path

### Phase 2A: Trust State With Dev Fingerprints

Build CarbonStack trust machinery before real cryptography.

Implement using fake/dev fingerprints derived from stub identity key material.

Add:

- trust.json
- trust-events.jsonl
- device trust states
- fingerprint display
- manual verify command
- trust history command
- send warning or block policy
- simulated key-change behavior
- simulated revocation behavior

Goal:

Make CarbonStack's unique trust semantics real before protocol integration.

### Phase 2B: Protocol Provider Interface

Define provider boundary:

- MockCryptoProvider
- future SignalLikeProvider
- future MLSProvider
- future NoiseTransferProvider

Provider should expose cryptographic results.

CarbonStack application logic should own trust decisions.

### Phase 2C: Signal/libsignal Feasibility Spike

Research only after Phase 2A.

Required answers:

- Is AGPL acceptable for CarbonStackComms?
- Can libsignal be used from Go, Rust, TypeScript, Java, or a helper process?
- Can the CLI test harness exercise it?
- What does it require for identity keys, prekeys, sessions, and persistence?
- Can it support one-device-per-person Phase 2?
- How does it expose identity-key changes?
- What test vectors or fixtures exist?

### Phase 2D: MLS Feasibility Spike

Research after or alongside Signal feasibility.

Focus:

- future groups
- group epochs
- membership logs
- device revocation
- whether two-person rooms could be unified with future group model

### Phase 2E: Noise and Utility Primitives

Defer until transfer/provisioning/backup needs become concrete.

## Current Provisional Decision

Do not implement real protocol code yet.

Next implementation step:

- Phase 2A trust-state scaffolding with fake/dev fingerprints

Next research step:

- keep Signal/libsignal as first 1:1 feasibility candidate
- keep MLS as future group candidate
- keep Noise for transfer/provisioning
- keep libsodium for narrow utilities only

## Allowed Claims After This Document

Allowed:

- CarbonStack has a protocol feasibility matrix.
- CarbonStack has not selected a final cryptographic protocol.
- CarbonStack's Phase 2 default path is trust-state scaffolding before real crypto.
- Signal/libsignal is the first serious 1:1 candidate.
- MLS is the leading future group candidate.
- Noise is reserved for narrow transfer or appliance channels.
- libsodium is reserved for utilities or non-production experiments.

Not allowed:

- CarbonStack has implemented real encryption.
- CarbonStack is Signal-equivalent.
- CarbonStack has selected libsignal permanently.
- CarbonStack has selected MLS permanently.
- CarbonStack has solved hostile-server security.
- CarbonStack has solved metadata privacy.
- CarbonStack has solved replay resistance.
- CarbonStack has solved hardware-key identity.
- CarbonStack is production secure.

## Next Work

Implement Phase 2A in CarbonStackComms:

- local trust state files
- fake/dev fingerprints
- verify-device command
- trust-history command
- send policy warnings or blocks
- key-change simulation test
- revocation simulation test

Update the Phase 1 test runner or add a Phase 2 validation runner once Phase 2A exists.
