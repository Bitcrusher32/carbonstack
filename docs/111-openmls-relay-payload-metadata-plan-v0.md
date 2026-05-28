# OpenMLS Relay Payload Metadata Plan v0

Status: Planning / migration design
Component: CarbonStackCypher + CarbonStackComms relay helper + OpenMLS sidecar
Phase: v0.2.59 payload metadata/hash/size planning
Previous docs:
- docs/108-real-cypher-server-openmls-relay-lifecycle-result-v0.md
- docs/109-openmls-real-cypher-relay-smoke-harness-result-v0.md
- docs/110-openmls-relay-ack-semantics-result-v0.md

## 1. Summary

This checkpoint plans payload metadata for Cypher envelopes carrying OpenMLS sidecar artifacts.

Current relay behavior stores opaque artifact bytes in:

    ciphertext_b64

This is intentionally inherited from the existing Cypher envelope model. It is sufficient for the first OpenMLS relay proofs, but it is awkward for deployability and debugging because the server does not currently expose decoded payload size or decoded payload hash metadata.

The planned metadata fields are:

    payload_size_bytes
    payload_sha256

These fields should describe the decoded bytes carried in `ciphertext_b64`.

## 2. Current state

Current Cypher `envelopes` table stores:

    envelope_id
    sender_device_id
    recipient_device_id
    content_type
    protocol_version
    ciphertext_b64
    client_created_at
    server_received_at
    delivery_state

Current Cypher submit behavior:

    trims request fields
    validates sender/recipient/content_type/protocol_version/ciphertext_b64 presence
    validates supported content type
    validates supported protocol version for content type
    validates ciphertext_b64 size limit
    validates base64 decoding succeeds
    verifies sender and recipient devices exist
    stores the original ciphertext_b64 string

Current Cypher inbox behavior returns queued envelopes with the same stored payload string, but no decoded payload metadata.

Current Comms relay behavior decodes `CiphertextB64` and writes exact bytes to a sidecar-compatible artifact path.

## 3. Planned metadata semantics

`payload_size_bytes`:

    decimal/integer count of decoded payload bytes from `ciphertext_b64`.

`payload_sha256`:

    lowercase hexadecimal SHA-256 digest of decoded payload bytes from `ciphertext_b64`.

Both fields describe decoded bytes, not the base64 text length.

For OpenMLS relay artifacts, this means:

    public-bundle.keypackage.bin bytes
    welcome.bin bytes
    application-message.bin bytes

## 4. Server-side computation rule

CarbonStackCypher should compute metadata server-side during `submitEnvelope`.

The server should:

    decode `ciphertext_b64`;
    reject invalid base64 as it already does;
    compute `payload_size_bytes = len(decoded_payload)`;
    compute `payload_sha256 = sha256(decoded_payload)`;
    store both metadata fields with the envelope;
    return both metadata fields in submit response;
    return both metadata fields in inbox response.

Clients should not be trusted to provide these fields in the first implementation.

If future client-provided metadata is allowed, it should be optional and must be checked against server-computed metadata before storage.

## 5. Database migration plan

Add a new migration after `001_init.sql`, likely:

    002_envelope_payload_metadata.sql

The migration should add nullable metadata columns first:

    ALTER TABLE envelopes ADD COLUMN payload_sha256 TEXT;
    ALTER TABLE envelopes ADD COLUMN payload_size_bytes INTEGER;

For new envelopes, both fields should be populated.

For existing dev DB envelopes, the first implementation may leave old rows null. Current dev database compatibility is acceptable because CarbonStackCypher is pre-release and not production-deployed.

A later cleanup may backfill old rows if needed.

## 6. API response shape

Submit response should become:

    envelope_id
    delivery_state
    server_received_at
    payload_size_bytes
    payload_sha256

Inbox envelope response should become:

    envelope_id
    sender_device_id
    recipient_device_id
    content_type
    protocol_version
    ciphertext_b64
    payload_size_bytes
    payload_sha256
    client_created_at
    server_received_at
    delivery_state

Ack response does not need payload metadata.

## 7. Comms client update plan

CarbonStackComms client envelope structs should add:

    PayloadSizeBytes int64  `json:"payload_size_bytes"`
    PayloadSHA256    string `json:"payload_sha256"`

SubmitEnvelopeResponse may also add the same fields if Cypher returns them.

The Comms relay helper should validate payload metadata when writing artifacts from envelope records.

Recommended validation order:

    decode CiphertextB64;
    if PayloadSizeBytes is present/nonzero, compare with decoded byte length;
    if PayloadSHA256 is present/non-empty, compare with lowercase SHA-256 hex digest;
    if metadata mismatches, fail before writing artifact;
    write artifact only after metadata validation succeeds.

For the first implementation, missing metadata can remain tolerated only if compatibility with pre-migration tests is needed. The real-server proof should require metadata once Cypher migration is implemented.

## 8. Test matrix

Cypher tests should cover:

    submit returns payload_size_bytes and payload_sha256 for stub envelope;
    inbox returns payload_size_bytes and payload_sha256 for stub envelope;
    OpenMLS KeyPackage opaque bytes return correct metadata;
    OpenMLS Welcome opaque bytes return correct metadata;
    OpenMLS application-message opaque bytes return correct metadata;
    invalid base64 remains rejected;
    unknown content type remains rejected;
    wrong protocol version remains rejected;
    existing ack behavior remains unchanged.

Comms tests should cover:

    WriteOpenMLSArtifactFromEnvelope accepts matching metadata;
    WriteOpenMLSArtifactFromEnvelope rejects mismatched payload_size_bytes;
    WriteOpenMLSArtifactFromEnvelope rejects mismatched payload_sha256;
    real-server OpenMLS lifecycle validates metadata before sidecar consume.

## 9. Security boundary

Payload metadata improves deployability, debugging, and storage/transport sanity checks.

It does not prove:

    production E2EE;
    hostile-server safety;
    metadata privacy;
    payload authenticity against a malicious server;
    OpenMLS semantic validity;
    Android readiness;
    external audit or certification.

The real cryptographic authenticity/integrity boundary remains OpenMLS sidecar consumption:

    conversation-add-member
    conversation-join
    message-open

`payload_sha256` is a server-computed relay/storage metadata field, not a trust root.

## 10. Recommendation

Do not treat v0.2.59 as a blind migration implementation.

The next implementation rung should deliberately update:

    carbonstack-cypher migration/schema/API/tests;
    carbonstack-comms client structs;
    carbonstack-comms relay metadata validation;
    real-server smoke proof expectations;
    docs and README claim wording.

## 11. Next rung

Next planned rung:

    v0.2.60 — payload metadata implementation or deployability docs/runbook adjustment.

If preserving the original ladder, v0.2.60 may become a metadata implementation rung and deployability docs/runbook can move one rung later.

If avoiding ladder churn, v0.2.59 can be extended into implementation after this plan is reviewed.
