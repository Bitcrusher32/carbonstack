# CarbonStack Protocol Threat Requirements v0

## Status

Classification: DRAFT / RESEARCH

This document defines the protocol requirements CarbonStackComms and CarbonStackCypher must satisfy before real cryptographic protocol integration begins.

This is not an implementation document. It is a selection and evaluation filter.

## Current Phase Boundary

Current Phase 1 behavior uses stub/base64 MockCryptoProvider behavior only.

That stub validates message lifecycle shape. It does not provide confidentiality, authentication, forward secrecy, post-compromise security, replay protection, metadata privacy, or hostile-server proof.

## Core Protocol Goal

Select or compose from mature, reviewed cryptographic protocol foundations without inventing custom cryptography.

The protocol must support CarbonStack's long-term model:

- small trusted groups
- 1:1 text first
- group messaging later
- hostile relay/server assumption
- loud trust changes
- explicit device enrollment
- explicit device revocation
- offline delivery
- auditable membership state
- testability from CLI before Android
- future Android/CarbonStackOS appliance integration

## Threat Assumptions

The protocol must assume:

- the server may be hostile
- the network may be hostile
- server storage may be inspected
- server routing metadata may be visible in early phases
- clients may go offline
- devices may be lost or replaced
- keys may change
- group membership may change
- compromised devices may need revocation
- replayed messages must be detectable or rejected
- trust state must not silently roll back

## Required Properties

### Message Content Confidentiality

The server must not be able to read plaintext message contents.

### Sender Authentication

Recipients must be able to verify which trusted device sent a message.

### Forward Secrecy

Compromise of current session material should not automatically reveal older messages.

### Post-Compromise Recovery

Where possible, future communication should recover after a device compromise is repaired or replaced.

### Replay Resistance

Replayed messages should be detectable or rejected.

### Loud Key Changes

Device identity key changes must be visible to users.

Silent server-driven key replacement is not acceptable.

### Device Enrollment Ceremony

Adding a device must be explicit and visible.

Future target:

- QR ceremony
- hardware-key approval
- trust history entry

### Device Revocation

Revocation must be visible and propagated.

Future target:

- hardware-key-signed revocation
- group-visible device removal
- forced re-verification after recovery

### Hostile-Server Group State

The server must not be able to silently add members, replace keys, roll back group state, or rewrite membership history without detection.

### CLI Testability

The protocol must be testable before Android.

Required:

- deterministic test vectors or fixture flows
- local CLI harness compatibility
- clear serialization format
- negative tests for replay/key-change/membership-tamper behavior

## Deferred Requirements

The following are important but may be deferred until after 1:1 protocol validation:

- full group messaging
- metadata privacy
- sealed sender-like behavior
- private contact discovery
- multi-device sync
- encrypted backups
- hardware-key-required release mode
- Android Keystore/StrongBox integration
- CarbonStackOS vault integration

## Explicit Non-Goals

Do not:

- invent a new cryptographic protocol
- claim Signal-equivalent security
- treat base64 or sealed boxes alone as a full messaging protocol
- rely on the server for identity truth
- hide key changes from users
- start Android before protocol behavior is testable from CLI
- implement groups before 1:1 lifecycle is well understood

## Candidate Families To Evaluate

- Signal Protocol / libsignal
- MLS / Messaging Layer Security
- Noise framework
- libsodium / NaCl-style primitives
- hybrid staged approach using simple primitives only for temporary non-production experiments

## Decision Output Required

Protocol research must produce:

- recommended Phase 2 protocol path
- rejected options with reasons
- licensing constraints
- implementation/library constraints
- testability assessment
- migration plan from MockCryptoProvider
- updated allowed/not-allowed security claims
