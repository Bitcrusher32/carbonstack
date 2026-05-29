# Envelope Lifecycle Semantics v0

Status: current semantics document
Component: CarbonStack + CarbonStackComms + CarbonStackCypher
Phase: v0.2.64 inbox/ack/general semantics + schema standardization

## 1. Purpose

This document defines the current envelope lifecycle language for the CarbonStack experimental backbone.

It does not define a production protocol.

It does not prove production security.

It does not prove hostile-server safety.

It does not prove metadata privacy.

The goal is narrower: every current public doc, test harness, and component README should use the same words for submit, inbox, payload metadata, sidecar consume, and ack.

## 2. Current proof boundary

The current validated artifact is the Cypher + Comms OpenMLS relay backbone.

It proves this local development path:

1. Bob exports an OpenMLS KeyPackage artifact.
2. Cypher relays the KeyPackage as an opaque envelope.
3. Alice retrieves the envelope payload.
4. Comms validates payload metadata before writing artifact bytes.
5. Alice consumes the KeyPackage through the OpenMLS sidecar.
6. Alice creates a Welcome artifact.
7. Cypher relays the Welcome as an opaque envelope.
8. Bob retrieves the envelope payload.
9. Comms validates payload metadata before writing artifact bytes.
10. Bob consumes the Welcome through the OpenMLS sidecar.
11. Alice protects an application-message artifact.
12. Cypher relays the application-message as an opaque envelope.
13. Bob retrieves the envelope payload.
14. Comms validates payload metadata before writing artifact bytes.
15. Bob consumes the application-message through the OpenMLS sidecar.
16. The plaintext matches.
17. Comms acknowledges each envelope only after the relevant sidecar consume command succeeds.

## 3. Envelope

An envelope is the server-visible relay unit.

Cypher stores it.

Comms submits it.

The recipient Comms side retrieves it.

The OpenMLS sidecar does not receive the envelope directly. It receives an artifact file written from the envelope payload.

Current envelope fields include:

- `envelope_id`
- `sender_device_id`
- `recipient_device_id`
- `content_type`
- `protocol_version`
- `ciphertext_b64`
- `payload_sha256`
- `payload_size_bytes`
- `client_created_at`
- `server_received_at`
- `delivery_state`

## 4. Opaque payload

The envelope payload is stored in `ciphertext_b64`.

For OpenMLS artifacts, this name is historical and imperfect.

Current meaning:

    ciphertext_b64 = opaque payload bytes encoded as base64

Cypher must not parse OpenMLS internals.

Cypher must not parse plaintext.

Cypher must not infer trust from payload contents.

## 5. Content type

`content_type` identifies the kind of opaque payload.

Current OpenMLS relay content types:

    carbonstack.mls.keypackage.v0
    carbonstack.mls.welcome.v0
    carbonstack.mls.application-message.v0

Existing stub content type:

    carbonstack.message.text.stub.v0

The content type is routing/schema metadata. It is not proof that the payload is semantically valid OpenMLS.

## 6. Protocol version

`protocol_version` identifies the expected envelope-level protocol family.

Current OpenMLS relay protocol version:

    carbonstack-openmls-sidecar-v0

Existing stub protocol version:

    stub-v0

The protocol version is not a claim of generic OpenMLS standard compatibility.

It is a CarbonStack compatibility label for the current sidecar relay path.

## 7. Payload metadata

Cypher computes:

    payload_sha256
    payload_size_bytes

Both fields describe decoded `ciphertext_b64` bytes.

`payload_sha256` is the lowercase SHA-256 hex digest of decoded payload bytes.

`payload_size_bytes` is the decoded payload byte length.

Comms validates these fields before writing downloaded artifact bytes.

Payload metadata proves only transport/storage consistency against the metadata returned by the server.

It does not prove:

- OpenMLS authenticity;
- hostile-server safety;
- metadata privacy;
- sender identity;
- semantic validity;
- production E2EE;
- trust root status.

A malicious server can lie about server-returned metadata.

The OpenMLS sidecar consume step remains the cryptographic validity gate.

## 8. Inbox

The inbox route returns queued envelopes for a recipient device.

Current route:

    GET /v0/devices/{device_id}/envelopes

Current semantics:

    inbox returns delivery_state = queued envelopes only.

It does not return acknowledged envelopes.

Inbox retrieval does not mean the envelope has been consumed.

Inbox retrieval does not mean the payload has been written.

Inbox retrieval does not mean the recipient has accepted the artifact.

## 9. Queued

`queued` means Cypher currently holds the envelope as pending for the recipient device.

It does not mean delivered.

It does not mean read.

It does not mean decrypted.

It does not mean OpenMLS-consumed.

It means available through the recipient inbox route.

## 10. Sidecar consume

Sidecar consume means the relevant OpenMLS sidecar command accepted the artifact and completed successfully.

Current consume commands:

- KeyPackage:
  - `conversation-add-member`

- Welcome:
  - `conversation-join`

- application-message:
  - `message-open`

For the current proof, sidecar consume is the point where Comms is allowed to ack the envelope.

## 11. Ack

Ack means the recipient device marks an envelope handled in Cypher.

Current route:

    POST /v0/envelopes/{envelope_id}/ack

Current request field:

    recipient_device_id

Ack is idempotent for the same recipient.

First same-recipient ack:

- records an ack event;
- sets `delivery_state = acknowledged`;
- returns `delivery_state = acknowledged`.

Second same-recipient ack:

- returns `delivery_state = acknowledged`;
- does not create a new semantic state;
- exists to make client retries safe.

Wrong-recipient ack is rejected.

Unknown-envelope ack is rejected.

Missing-recipient ack is rejected.

## 12. Acknowledged

`acknowledged` means Cypher has accepted a recipient-device ack for the envelope.

In the current Comms proof, ack is sent only after sidecar consume succeeds.

Cypher itself does not know that sidecar consume succeeded.

Cypher only knows that the recipient device sent a valid ack request.

That distinction matters.

Correct wording:

    Comms acks after sidecar consume succeeds.

Incorrect wording:

    Cypher proves the sidecar consumed the artifact.

## 13. What Cypher knows

Cypher knows:

- envelope ID;
- sender device ID;
- recipient device ID;
- content type;
- protocol version;
- base64 payload string;
- payload size/hash over decoded bytes;
- server receive time;
- delivery state;
- ack request identity.

Cypher does not know:

- plaintext;
- OpenMLS group state;
- OpenMLS semantic validity;
- whether a human read a message;
- whether a device is trustworthy;
- whether the server is honest;
- whether the payload is safe beyond basic base64/hash/size metadata.

## 14. What Comms knows

Comms knows:

- which sidecar artifact kind it is relaying;
- which Cypher content type maps to that artifact;
- whether payload metadata matches decoded bytes;
- whether writing the artifact file succeeded;
- whether the sidecar command succeeded;
- whether to ack after consume success.

Comms must not treat server metadata as a trust root.

Comms must not ack merely because an envelope was downloaded.

Comms must not ack merely because bytes were written.

## 15. What the OpenMLS sidecar proves

The sidecar proves whether the artifact can be consumed by the current development OpenMLS state.

For the current proof:

- `conversation-add-member` consumes a KeyPackage;
- `conversation-join` consumes a Welcome;
- `message-open` consumes an application-message and recovers expected plaintext.

The sidecar does not prove the whole CarbonStack product is production-ready.

It does not prove secure vault storage.

It does not prove Android readiness.

It does not prove external audit or certification.

## 16. Current safe rule

The current safe rule is:

    download is not ack;
    artifact write is not ack;
    payload metadata validation is not ack;
    sidecar consume success permits ack.

That rule is the current CarbonStack experimental backbone lifecycle boundary.
