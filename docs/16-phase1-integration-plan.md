# CarbonStack Phase 1 Integration Plan

## Status

Classification: PLANNED / NOT IMPLEMENTED

This document defines the first implementation-oriented vertical slice across CarbonStackComms and CarbonStackCypher.

Phase 1 does not validate production security. It does not select or implement a final cryptographic protocol. It does not start CarbonStackOS work.

## Goal

Prove the minimal relay/client system shape:

1. Invite-only registration exists.
2. Accounts can register one or more devices.
3. Devices can publish public/stub key material.
4. A client can submit an encrypted-envelope-shaped message to CarbonStackCypher.
5. Another client can retrieve the envelope.
6. The server stores and routes opaque envelopes without needing plaintext.
7. The client lifecycle is testable through a CLI before Android implementation.

## Components

### CarbonStackCypher

The relay/server stack.

Responsibilities:

- invite-only registration
- account records
- device records
- opaque envelope storage
- envelope retrieval
- delivery acknowledgement
- basic rate limiting later
- no plaintext message access

### CarbonStackComms

The client stack.

Phase 1 client target:

- CLI test client
- local state file/database
- mock identity/device state
- mock crypto provider
- envelope send/retrieve lifecycle

### CarbonStackOS

Deferred.

CarbonStackOS must not be started during Phase 1 except as documentation context.

## Non-Goals

Phase 1 does not include:

- Android app implementation
- CarbonStackOS build work
- production cryptography
- custom cryptography
- group messaging
- attachments
- media
- hardware-key enrollment
- local vault hardening
- metadata privacy claims
- production deployment claims

## Phase 1 Success Criteria

Phase 1 is successful when a local operator can:

1. Start CarbonStackCypher locally.
2. Create or seed an invite.
3. Register two accounts/devices.
4. Send an envelope from Client A to Client B.
5. Retrieve the envelope as Client B.
6. Acknowledge delivery.
7. Confirm that the server did not require plaintext message contents.

## Phase 1 Failure Criteria

Phase 1 fails if:

- the implementation requires plaintext message contents on the server
- the project starts inventing cryptography
- Android/OS work blocks the relay/client skeleton
- groups or attachments become required for the first slice
- docs and implementation disagree without explicit revision
- security claims exceed validation

## Implementation Order

1. CarbonStackCypher data model.
2. CarbonStackCypher API contract.
3. CarbonStackCypher local server skeleton.
4. CarbonStackComms CLI lifecycle doc.
5. CarbonStackComms CLI stub.
6. Integration test: two local clients, one local relay.
7. Update Security Assurance Matrix with results.

## Validation Level

After Phase 1, allowed claims:

- CarbonStack has a working relay/client skeleton.
- CarbonStackCypher can store and route opaque envelopes.
- CarbonStackComms can exercise a basic message lifecycle through a CLI.
- The implementation remains protocol-pluggable.

Not allowed claims:

- CarbonStack is production secure.
- CarbonStack is audited.
- CarbonStack is Signal-equivalent.
- CarbonStack prevents endpoint compromise.
- CarbonStack has selected or validated its final cryptographic protocol.
