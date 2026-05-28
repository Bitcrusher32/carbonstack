# Cypher OpenMLS Envelope Content-Type Result v0

Status: Implementation result
Component: CarbonStackCypher
Phase: v0.2.48 Cypher OpenMLS envelope content-type implementation
Previous docs:
- docs/97-cypher-opaque-mls-artifact-relay-plan-v0.md
- docs/98-cypher-opaque-mls-artifact-relay-recon-v0.md
- docs/99-cypher-openmls-envelope-relay-recon-v0.md

## 1. Summary

This checkpoint records the first CarbonStackCypher implementation step toward OpenMLS sidecar artifact relay.

Cypher now accepts OpenMLS artifact content types through the existing `/v0/envelopes` opaque envelope route.

No new routes were added.

No database migration was added.

No Comms runtime wiring was added.

No MLS parsing was added.

## 2. Accepted content types

Existing stub content type remains accepted:

    carbonstack.message.text.stub.v0

New OpenMLS content types:

    carbonstack.mls.keypackage.v0
    carbonstack.mls.welcome.v0
    carbonstack.mls.application-message.v0

## 3. Protocol versions

Existing stub protocol version remains:

    stub-v0

OpenMLS sidecar artifact protocol version:

    carbonstack-openmls-sidecar-v0

The OpenMLS protocol version name is intentionally CarbonStack-specific. It is not a claim to generic OpenMLS standard compatibility or production security certification.

## 4. Route/storage model

The implementation reuses:

    POST /v0/envelopes
    GET /v0/devices/{device_id}/envelopes
    POST /v0/envelopes/{envelope_id}/ack

The existing `envelopes` table continues to store:

    content_type
    protocol_version
    ciphertext_b64
    sender_device_id
    recipient_device_id
    delivery_state

For this scaffold, `ciphertext_b64` means opaque envelope payload bytes encoded as base64.

For OpenMLS content types, those bytes may correspond to:

    public-bundle.keypackage.bin
    welcome.bin
    application-message.bin

## 5. Validation

Tests were added or updated to prove:

- keypackage content type is accepted;
- welcome content type is accepted;
- application-message content type is accepted;
- opaque bytes roundtrip exactly through `ciphertext_b64`;
- OpenMLS content types reject the old stub protocol version;
- stub content type rejects the OpenMLS sidecar protocol version;
- existing stub envelope lifecycle remains valid;
- unknown content types and invalid base64 remain rejected.

## 6. Non-goals

This checkpoint does not implement:

- Comms runtime bridge;
- sidecar invocation from Cypher;
- MLS parsing inside Cypher;
- signer/provider storage relay;
- plaintext relay;
- payload hash/size database migration;
- blob filesystem storage;
- Android app integration;
- production security certification.

## 7. Server-deployable pre-release direction

This work supports a future experimental server-deployable CarbonStack backbone.

That future pre-release should be framed as:

    experimental
    immature
    not externally audited
    not certified secure
    not production E2EE
    useful as a deployable OpenMLS/Cypher foundation for builders

The project should not claim security certification unless experienced external reviewers/auditors examine and harden the system.

## 8. Next rung

Next planned rung:

    v0.2.49 — Comms-side relay bridge recon.

Goal:

    design how CarbonStackComms should read sidecar artifact path hints, submit opaque artifact bytes through Cypher envelopes, retrieve bytes from Cypher, write them to safe local artifact paths, and pass them back into sidecar commands.
