# CarbonStack Protocol Candidate Evaluation

## Status

Classification: DRAFT / RESEARCH

This document evaluates candidate protocol foundations for CarbonStackComms and CarbonStackCypher.

It is not a final protocol decision.

## Current Recommendation Summary

Initial direction:

- 1:1 messaging: investigate Signal Protocol / libsignal first.
- Group messaging: investigate MLS as the likely long-term group protocol candidate.
- Transport or narrow device-transfer channels: investigate Noise.
- Primitive-only fallback: libsodium may be useful for tools or temporary experiments, but sealed boxes alone are not a full messaging protocol.

## Candidate 1: Signal Protocol / libsignal

### Why It Fits

Signal-style protocols are designed for asynchronous private messaging and are directly aligned with 1:1 secure chat goals.

libsignal contains platform-agnostic APIs used by official Signal clients and servers, exposed for Java, Swift, and TypeScript, with underlying Rust implementations. :contentReference[oaicite:1]{index=1}

### Strengths

- strong alignment with 1:1 messaging
- mature ecosystem
- forward secrecy and post-compromise recovery model
- practical reference point for key-change UX
- good conceptual fit for offline delivery and prekey-style setup

### Risks / Questions

- license compatibility must be checked before integration
- binding fit for Go CLI/server and future Android client must be evaluated
- may be complex to integrate correctly
- group messaging story may require separate design or later MLS path
- must avoid overclaiming Signal-equivalent behavior

The libsignal repository license is AGPL-3.0 according to the public GitHub license file, so license implications must be treated as a real design constraint. :contentReference[oaicite:2]{index=2}

### Initial Fit

1:1 Phase 2 candidate: STRONG

Group Phase 2 candidate: PARTIAL / RESEARCH REQUIRED

## Candidate 2: MLS / Messaging Layer Security

### Why It Fits

MLS is designed for group key agreement. RFC 9420 describes asynchronous group keying with forward secrecy and post-compromise security. :contentReference[oaicite:3]{index=3}

The RFC Editor summary describes MLS as efficient asynchronous group key establishment with forward secrecy and post-compromise security for groups ranging from two to thousands. :contentReference[oaicite:4]{index=4}

### Strengths

- standardized protocol
- designed for group messaging
- strong match for auditable group epochs and membership changes
- likely relevant to future CarbonStack group state

### Risks / Questions

- may be too heavy for Phase 2 1:1-first work
- implementation/library maturity must be researched
- requires careful identity/device model integration
- may not solve all application-layer trust UX by itself
- likely requires substantial protocol design around CarbonStack-specific ceremonies and revocation UX

### Initial Fit

1:1 Phase 2 candidate: POSSIBLE BUT PROBABLY OVERKILL

Group future candidate: STRONG

## Candidate 3: Noise Framework

### Why It Fits

Noise is a framework for building crypto protocols based on Diffie-Hellman key agreement, supporting authentication patterns, forward secrecy, and related secure-channel properties. :contentReference[oaicite:5]{index=5}

The official Noise site describes it as a framework for building crypto protocols with mutual/optional authentication, identity hiding, forward secrecy, and zero round-trip encryption. :contentReference[oaicite:6]{index=6}

### Strengths

- good for narrow secure channel construction
- useful for custom transport handshakes if needed
- strong fit for future transfer channels or device-to-device protocols
- flexible and relatively understandable compared to full messenger protocols

### Risks / Questions

- framework, not a complete messaging protocol
- does not automatically provide full asynchronous messaging state, trust UX, group membership, or revocation model
- easy to misuse if treated as a complete Signal replacement

### Initial Fit

1:1 chat protocol candidate: PARTIAL

Transport/transfer protocol candidate: STRONG

## Candidate 4: libsodium / NaCl-style Primitives

### Why It Fits

libsodium is a portable library providing cryptographic operations such as encryption, signatures, password hashing, and related primitives. :contentReference[oaicite:7]{index=7}

### Strengths

- practical, well-known primitives
- useful for tooling and experiments
- potentially useful for file encryption, signatures, sealed test envelopes, or local utility code
- simpler than full protocol stacks

### Risks / Questions

- primitives are not a full messaging protocol
- sealed boxes alone do not authenticate the sender
- protocol composition risk is high

libsodium sealed boxes are designed to anonymously send messages to a recipient public key, but the recipient cannot verify sender identity from the sealed box alone. :contentReference[oaicite:8]{index=8}

### Initial Fit

Temporary experimental envelope candidate: WEAK / NON-PRODUCTION ONLY

Utility primitive candidate: USEFUL WITH CARE

## Initial Ranking

### For Phase 2 1:1 Messaging

1. Signal Protocol / libsignal
2. MLS only if library/architecture makes 1:1 practical
3. Noise only for narrow channel use, not full messenger state
4. libsodium primitives only for non-production scaffolding or utilities

### For Future Group Messaging

1. MLS
2. Signal-style sender-key/group approach only if explicitly justified
3. Custom group design rejected unless based on mature reviewed components

### For Transfer / Appliance Channels

1. Noise
2. libsodium primitives if the channel is narrow, documented, and heavily tested

## Open Research Questions

- What libsignal bindings are practical for CarbonStackComms CLI and future Android?
- Is AGPL acceptable for this project family?
- Is MLS implementation maturity sufficient for future groups?
- What Go/Rust/Android libraries exist and are maintained?
- How should CarbonStack model identity keys, device keys, prekeys, and trust history?
- What should the Phase 2 test harness prove before any security claims change?
- How can the protocol remain CLI-testable before Android integration?
- How should key changes and device revocation be represented in local state?

## Provisional Decision

Do not implement protocol code yet.

Next step:

- research libsignal integration feasibility
- research MLS implementation feasibility
- define CarbonStack identity/trust state model
- define Phase 2 protocol test vectors
- update allowed/not-allowed security claims
