# Cypher Opaque MLS Artifact Relay Recon v0

Status: Recon
Component: CarbonStackCypher + CarbonStackComms OpenMLS sidecar
Phase: v0.2.47 Cypher minimal opaque MLS artifact relay recon
Previous docs:
- docs/91-openmls-sidecar-artifact-ownership-map-v0.md
- docs/92-openmls-sidecar-command-schema-matrix-v0.md
- docs/97-cypher-opaque-mls-artifact-relay-plan-v0.md

## 1. Purpose

This document records the initial relay recon for carrying OpenMLS sidecar artifacts through CarbonStackCypher.

The working model is:

    Cypher relays opaque MLS artifacts.
    Comms/sidecar creates and consumes MLS artifacts.
    Cypher does not parse MLS internals.
    Cypher does not receive local signer/provider storage.
    Cypher does not know plaintext.

## 2. Artifact ownership summary

### public-bundle.keypackage.bin

Producer:

    recipient/invitee device through public-bundle-export --write-artifact

Consumer:

    inviter/creator device through conversation-add-member --member-keypackage <path>

Relay direction:

    recipient/invitee -> inviter/creator

First Cypher meaning:

    public onboarding artifact available for another device to consume.

Cypher must not:

    parse as KeyPackage;
    validate MLS internals;
    infer final identity semantics from dev labels.

### welcome.bin

Producer:

    inviter/creator through conversation-add-member

Consumer:

    invitee through conversation-join --welcome <path>

Relay direction:

    inviter/creator -> invitee

First Cypher meaning:

    opaque invite/join artifact.

Cypher must not:

    parse Welcome;
    inspect bytes;
    treat it as plaintext;
    derive membership truth from it.

### application-message.bin

Producer:

    sender through message-protect

Consumer:

    receiver through message-open --message <path>

Relay direction:

    sender -> receiver

First Cypher meaning:

    opaque protected application message artifact.

Cypher must not:

    parse MLS message;
    inspect plaintext;
    decide trust-state;
    decide replay semantics beyond transport-level duplicate handling.

## 3. Metadata boundary

Minimum useful metadata:

    artifact_id
    artifact_type
    sender_device_ref
    recipient_device_ref
    conversation_ref
    artifact_label
    sha256
    size_bytes
    created_at
    delivery_state

Artifact types:

    keypackage
    welcome
    application-message

Delivery states:

    stored
    listed
    downloaded
    acknowledged
    expired

Open question:

    Whether `downloaded` should exist at all, or whether first implementation should only have `stored` and `acknowledged`.

## 4. Dev-label caveat

The current sidecar uses dev labels such as:

    carbonstack-alice-device
    carbonstack-bob-device
    carbonstack-test-conversation
    message-0001

Cypher recon may use these as temporary dev refs.

However, dev labels are not final production identity.

A future design must replace or wrap them with stable device/conversation refs that do not leak unnecessary user semantics.

## 5. Storage shape candidate

A minimal first storage object can be:

    artifacts/<artifact_id>/artifact.bin
    artifacts/<artifact_id>/metadata.json

Metadata fields:

    schema_version
    artifact_id
    artifact_type
    sender_device_ref
    recipient_device_ref
    conversation_ref
    artifact_label
    sha256
    size_bytes
    created_at
    delivery_state

For a database-backed implementation, the same fields can map to an artifact table.

For a filesystem-backed dev implementation, metadata + blob files are acceptable.

## 6. Integrity handling

Cypher can record:

    sha256
    size_bytes

Cypher can verify:

    bytes written match declared size;
    bytes read back match stored hash.

Cypher should not verify:

    OpenMLS cryptographic validity;
    group membership;
    sender authenticity;
    replay status at MLS layer.

Those remain Comms/sidecar responsibilities.

## 7. Routing model candidate

First model:

    mailbox-style pull

Flow:

    sender uploads artifact for recipient;
    recipient lists pending artifacts;
    recipient downloads artifact bytes;
    recipient passes artifact path to sidecar;
    recipient acknowledges after sidecar consumption succeeds.

Why mailbox-style:

- simple;
- self-hostable;
- does not require push notifications;
- easier to test from CLI;
- compatible with later Android polling or notification bridge.

## 8. Minimal API sketch

This is not final API.

Possible first routes:

    POST /v0/artifacts
    GET  /v0/artifacts?recipient=<recipient_device_ref>
    GET  /v0/artifacts/<artifact_id>
    POST /v0/artifacts/<artifact_id>/ack

POST body concept:

    metadata fields
    artifact bytes or multipart blob

GET list response:

    metadata only, no artifact bytes

GET artifact response:

    artifact bytes plus hash/size metadata

ACK body:

    recipient_device_ref
    local consume result reference or status

## 9. Security constraints for first implementation

Cypher must:

- reject path separators in artifact IDs and labels;
- generate server-side artifact IDs or strictly validate client-provided IDs;
- cap artifact size;
- store blobs outside executable/source paths;
- never write artifact bytes to arbitrary client-supplied paths;
- never log raw artifact bytes;
- never log signer/provider storage;
- never accept signer/provider storage as artifact types;
- reject unknown artifact types;
- keep artifact metadata separate from artifact bytes;
- preserve enough hash/size data for client-side integrity checks.

## 10. Test plan candidate

Minimum Cypher tests before claiming first relay scaffold:

- accepts valid keypackage artifact as opaque bytes;
- accepts valid welcome artifact as opaque bytes;
- accepts valid application-message artifact as opaque bytes;
- rejects unknown artifact type;
- rejects oversized artifact;
- rejects artifact label/path traversal attempts;
- list returns metadata only;
- download returns exact bytes;
- downloaded bytes hash matches uploaded hash;
- ack changes delivery state;
- signer.json/provider-storage.json artifact types are rejected;
- raw bytes are not printed in normal logs/test output.

Integration-style test later:

- sidecar produces `application-message.bin`;
- Cypher stores and returns exact bytes;
- receiver sidecar opens returned artifact successfully.

## 11. Implementation sequence candidate

Recommended next implementation ladder after recon:

1. Cypher doc/API skeleton for opaque artifact store.
2. Local filesystem-backed dev artifact store.
3. Unit tests for metadata/blob validation.
4. CLI or test helper to upload/list/download artifact bytes.
5. Sidecar-produced artifact roundtrip through Cypher test fixture.
6. Only then consider Comms runtime wiring.

## 12. Open questions

- Should artifact IDs be server-generated only?
- Should `artifact_label` remain separate from `message_label`?
- Should `conversation_ref` be opaque random ref even in dev?
- Should keypackage artifacts be public/profile-like or mailbox-delivered?
- Should downloaded artifacts stay available until explicit ack?
- Should Cypher support expiration/TTL in the first scaffold?
- Should Cypher store one artifact per recipient or one artifact with multiple recipient refs?
- Should ack be purely transport-level or include sidecar-consume status?

## 13. Current recommendation

The safest first Cypher rung is:

    filesystem-backed opaque artifact mailbox

with:

    strict artifact type enum
    server-generated artifact_id
    safe dev metadata
    sha256 + size_bytes
    list/download/ack tests
    no MLS parsing
    no Comms runtime wiring
    no trust-state mutation

This fits the current Phase 2D sidecar proof while keeping Cypher's responsibilities narrow.
