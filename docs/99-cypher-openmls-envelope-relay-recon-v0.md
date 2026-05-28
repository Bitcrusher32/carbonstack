# Cypher OpenMLS Envelope Relay Recon v0

Status: Recon result
Component: CarbonStackCypher + CarbonStackComms OpenMLS sidecar
Phase: v0.2.47 Cypher minimal opaque MLS artifact relay recon
Previous docs:
- docs/91-openmls-sidecar-artifact-ownership-map-v0.md
- docs/92-openmls-sidecar-command-schema-matrix-v0.md
- docs/96-openmls-sidecar-readme-current-state-cleanup-result-v0.md
- docs/97-cypher-opaque-mls-artifact-relay-plan-v0.md
- docs/98-cypher-opaque-mls-artifact-relay-recon-v0.md

## 1. Summary

Cypher repo recon found that CarbonStackCypher already has the right primitive for first OpenMLS artifact relay: an opaque envelope mailbox.

Therefore, the first OpenMLS relay implementation should not create a parallel artifact API yet.

Instead, it should extend the existing envelope relay to accept OpenMLS artifact content types.

## 2. Existing Cypher structure

The current Cypher repo is a Go + SQLite HTTP relay skeleton.

Important paths:

    cmd/cypher/main.go
    internal/config/config.go
    internal/db/db.go
    internal/httpapi/api.go
    internal/httpapi/api_test.go
    migrations/001_init.sql

Current server routing uses:

    http.NewServeMux

Current API routes include:

    GET  /v0/health
    POST /v0/dev/invites
    POST /v0/invites/claim
    POST /v0/devices/register
    GET  /v0/accounts/{account_id}/devices
    POST /v0/envelopes
    GET  /v0/devices/{device_id}/envelopes
    POST /v0/envelopes/{envelope_id}/ack

Current storage uses SQLite tables:

    invites
    accounts
    devices
    envelopes
    envelope_acks

## 3. Existing envelope model

The existing envelope table already stores:

    envelope_id
    sender_device_id
    recipient_device_id
    content_type
    protocol_version
    ciphertext_b64
    client_created_at
    server_received_at
    delivery_state

For first OpenMLS relay work, `ciphertext_b64` can carry opaque OpenMLS artifact bytes as base64.

This name is imperfect for KeyPackage and Welcome artifacts, but changing the schema is not required for the first relay scaffold.

## 4. Current implementation blocker

The current `submitEnvelope` handler only accepts:

    content_type = carbonstack.message.text.stub.v0
    protocol_version = stub-v0

It also verifies base64 and enforces the current Phase 1 size limit.

The first implementation should widen accepted content types while preserving opaque payload behavior.

## 5. Proposed OpenMLS content types

First accepted OpenMLS content types:

    carbonstack.mls.keypackage.v0
    carbonstack.mls.welcome.v0
    carbonstack.mls.application-message.v0

Existing stub content type should remain accepted:

    carbonstack.message.text.stub.v0

Recommended OpenMLS protocol version:

    carbonstack-openmls-sidecar-v0

This marks the relay contract as tied to the current promoted OpenMLS sidecar proof, not final production MLS protocol.

## 6. Artifact mapping

### KeyPackage

Sidecar artifact:

    public-bundle.keypackage.bin

Cypher envelope content type:

    carbonstack.mls.keypackage.v0

Direction:

    invitee/recipient device -> inviter/creator device

Server treatment:

    opaque base64 payload

### Welcome

Sidecar artifact:

    welcome.bin

Cypher envelope content type:

    carbonstack.mls.welcome.v0

Direction:

    inviter/creator device -> invitee/recipient device

Server treatment:

    opaque base64 payload

### Application message

Sidecar artifact:

    application-message.bin

Cypher envelope content type:

    carbonstack.mls.application-message.v0

Direction:

    sender device -> receiver device

Server treatment:

    opaque base64 payload

## 7. What Cypher must not handle

Cypher must not accept, route, parse, log, or store as relay artifacts:

    signer.json
    provider-storage.json
    raw MemoryStorage JSON
    raw OpenMLS group state
    plaintext
    private keys
    trust-state private material

These remain local-only.

## 8. Why not a separate artifact API yet

A separate `/v0/artifacts` API is not needed for the first implementation because:

- `/v0/envelopes` already models sender -> recipient opaque delivery;
- `/v0/devices/{device_id}/envelopes` already models mailbox-style retrieval;
- `/v0/envelopes/{envelope_id}/ack` already models recipient acknowledgement;
- SQLite storage already has the required routing fields;
- extending `content_type` is lower risk than adding a new route/storage model.

A future artifact/blob API may still be useful if payloads become too large for inline base64 storage or need filesystem-backed blobs.

## 9. Integrity plan

First implementation can prove exact byte relay in tests by:

1. taking fixture bytes;
2. base64 encoding them;
3. submitting through `/v0/envelopes`;
4. retrieving via device inbox;
5. base64 decoding returned `ciphertext_b64`;
6. comparing returned bytes with original bytes.

This is enough for the first scaffold.

Future migration may add:

    payload_sha256
    payload_size_bytes

but this is not required for the first content-type widening proof.

## 10. Test plan

Add tests to `internal/httpapi/api_test.go` or a focused adjacent test file.

Minimum tests:

- accepts `carbonstack.mls.keypackage.v0`;
- accepts `carbonstack.mls.welcome.v0`;
- accepts `carbonstack.mls.application-message.v0`;
- preserves exact bytes through submit/list;
- keeps existing `carbonstack.message.text.stub.v0` test passing;
- rejects unknown content type;
- rejects unsupported protocol version;
- rejects invalid base64;
- preserves recipient-only ack behavior.

Do not add tests that parse MLS internals.

## 11. Implementation recommendation

First Cypher implementation rung:

- add constants/helper for accepted content types;
- allow the three OpenMLS content types;
- allow `carbonstack-openmls-sidecar-v0` for OpenMLS artifact content types;
- keep `stub-v0` for `carbonstack.message.text.stub.v0`;
- preserve current base64 validation and size limit;
- add exact byte roundtrip tests.

No migration is required.

No route changes are required.

No Comms runtime integration is required.

## 12. Remaining open questions

- Should payload size limit remain 65536 for all MLS artifacts?
- Should KeyPackages and Welcome artifacts use the same protocol version as application messages?
- Should future schema add payload hash/size fields?
- Should future Comms bridge write downloaded artifacts to sidecar-compatible temp files?
- Should ack happen only after sidecar consumption succeeds?
- Should keypackage relay be mailbox-based or public-directory-like later?

## 13. Definition of done for next implementation rung

The next implementation rung is complete when:

- Cypher accepts the three OpenMLS content types;
- tests prove exact opaque byte roundtrip;
- unknown content types still fail;
- invalid base64 still fails;
- existing envelope lifecycle tests still pass;
- no MLS parsing is introduced;
- no signer/provider storage handling is introduced;
- no plaintext handling is introduced.
